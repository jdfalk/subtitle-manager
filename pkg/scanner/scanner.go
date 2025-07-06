// file: pkg/scanner/scanner.go
package scanner

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sourcegraph/conc/pool"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/events"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/security"
)

// ScanDirectory walks through the directory and downloads subtitles for video files
// using provider p for the given language. providerName is stored in download
// history. If upgrade is false existing subtitle files are skipped.
func ScanDirectory(ctx context.Context, dir, lang string, providerName string, p providers.Provider, upgrade bool, workers int, store database.SubtitleStore) error {
	logger := logging.GetLogger("scanner")
	sanitizedDir, err := security.ValidateAndSanitizePath(dir)
	if err != nil {
		logger.Warnf("invalid path: %v", err)
		return err
	}
	work := pool.New().WithErrors().WithMaxGoroutines(workers)
	err = filepath.WalkDir(sanitizedDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !isVideoFile(path) {
			return nil
		}
		f := filepath.Clean(path)
		work.Go(func() error {
			logger.Debugf("process %s", f)
			return ProcessFile(ctx, f, lang, providerName, p, upgrade, store)
		})
		return nil
	})
	if err != nil {
		return err
	}
	return work.Wait()
}

// ProcessFile downloads a subtitle for path using providerName for history
// tracking. The subtitle is saved next to the media file with the language
// code appended before the extension. If upgrade is false an existing subtitle
// file is left untouched. When upgrade is true and a subtitle already exists,
// the new subtitle replaces it only if the file size is larger, indicating
// potentially better quality.
func ProcessFile(ctx context.Context, path, lang string, providerName string, p providers.Provider, upgrade bool, store database.SubtitleStore) error {
	logger := logging.GetLogger("scanner")

	// Validate and sanitize all user inputs
	sanitizedPath, err := security.ValidateAndSanitizePath(path)
	if err != nil {
		logger.Warnf("invalid path: %v", err)
		return err
	}
	path = sanitizedPath

	// Validate the language code to prevent path traversal attacks
	if err := security.ValidateLanguageCode(lang); err != nil {
		logger.Warnf("invalid language code: %v", err)
		return err
	}
	// Ensure the language code does not contain any path traversal characters
	if strings.Contains(lang, "/") || strings.Contains(lang, "\\") || strings.Contains(lang, "..") {
		logger.Warnf("language code contains invalid characters")
		return fmt.Errorf("invalid language code")
	}

	// Validate provider name if provided
	if err := security.ValidateProviderName(providerName); err != nil {
		logger.Warnf("invalid provider name: %v", err)
		return err
	}

	// Construct and validate the output path securely
	validatedOutputPath, err := security.ValidateSubtitleOutputPath(path, lang)
	if err != nil {
		logger.Warnf("invalid subtitle output path: %v", err)
		return err
	}

	if !upgrade {
		if _, err := os.Stat(validatedOutputPath); err == nil {
			return nil
		}
	}
	var data []byte
	if p != nil {
		data, err = p.Fetch(ctx, path, lang)
	} else {
		data, providerName, err = providers.FetchFromAll(ctx, path, lang, "")
	}
	if err != nil {
		logger.Warnf("fetch %s: %v", path, err)

		// Send event for subtitle fetch failure
		events.PublishSubtitleFailed(ctx, events.SubtitleFailedData{
			FilePath:  path,
			Language:  lang,
			Provider:  providerName,
			Error:     err.Error(),
			Timestamp: time.Now(),
		})

		return err
	}
	var wasUpgrade bool
	if upgrade {
		if oldData, err := os.ReadFile(validatedOutputPath); err == nil {
			if len(data) <= len(oldData) {
				logger.Debugf("existing subtitle %s is higher quality", validatedOutputPath)
				return nil
			}
			wasUpgrade = true
		}
	}
	if err := os.WriteFile(validatedOutputPath, data, 0644); err != nil {
		logger.Warnf("write %s: %v", validatedOutputPath, err)

		// Send event for file write failure
		events.PublishSubtitleFailed(ctx, events.SubtitleFailedData{
			FilePath:  path,
			Language:  lang,
			Provider:  providerName,
			Error:     "Failed to write subtitle file: " + err.Error(),
			Timestamp: time.Now(),
		})
		return err
	}
	logger.Infof("downloaded subtitle %s", validatedOutputPath)

	// Get file size for webhook event
	var fileSize int64
	if stat, err := os.Stat(validatedOutputPath); err == nil {
		fileSize = stat.Size()
	}

	// Send appropriate event
	if wasUpgrade {
		events.PublishSubtitleUpgraded(ctx, events.SubtitleUpgradedData{
			FilePath:        path,
			NewSubtitlePath: validatedOutputPath,
			Language:        lang,
			NewProvider:     providerName,
			NewScore:        1.0, // Default score, could be enhanced
			Timestamp:       time.Now(),
		})
	} else {
		events.PublishSubtitleDownloaded(ctx, events.SubtitleDownloadedData{
			FilePath:     path,
			SubtitlePath: validatedOutputPath,
			Language:     lang,
			Provider:     providerName,
			Score:        1.0, // Default score, could be enhanced
			Size:         fileSize,
			Timestamp:    time.Now(),
		})
	}
	if store != nil {
		_ = store.InsertDownload(&database.DownloadRecord{File: validatedOutputPath, VideoFile: path, Provider: providerName, Language: lang})
	}
	return nil
}

var videoExtensions = []string{".mkv", ".mp4", ".avi", ".mov"}

func isVideoFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	for _, e := range videoExtensions {
		if ext == e {
			return true
		}
	}
	return false
}

// ScanDirectoryWithProfiles walks through the directory and downloads subtitles for video files
// using language profiles. Each video file's profile is determined by its media_profiles assignment.
func ScanDirectoryWithProfiles(ctx context.Context, dir string, db *sql.DB, upgrade bool, workers int, store database.SubtitleStore) error {
	logger := logging.GetLogger("scanner")
	sanitizedDir, err := security.ValidateAndSanitizePath(dir)
	if err != nil {
		logger.Warnf("invalid path: %v", err)
		return err
	}
	work := pool.New().WithErrors().WithMaxGoroutines(workers)
	err = filepath.WalkDir(sanitizedDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !isVideoFile(path) {
			return nil
		}
		f := filepath.Clean(path)
		work.Go(func() error {
			logger.Debugf("process %s with profiles", f)
			return ProcessFileWithProfile(ctx, f, db, upgrade, store)
		})
		return nil
	})
	if err != nil {
		return err
	}
	return work.Wait()
}

// ProcessFileWithProfile downloads subtitles using the language profile assigned to the media file.
func ProcessFileWithProfile(ctx context.Context, path string, db *sql.DB, upgrade bool, store database.SubtitleStore) error {
	logger := logging.GetLogger("scanner")

	// Validate path
	sanitizedPath, err := security.ValidateAndSanitizePath(path)
	if err != nil {
		logger.Warnf("invalid path: %v", err)
		return err
	}

	// Use profile-based fetch to get subtitles
	data, providerName, actualLang, err := providers.FetchWithProfile(ctx, db, sanitizedPath, "")
	if err != nil {
		logger.Warnf("fetch with profile %s: %v", sanitizedPath, err)
		return err
	}

	// Construct and validate the output path securely using the actual language found
	out, err := security.ValidateSubtitleOutputPath(sanitizedPath, actualLang)
	if err != nil {
		logger.Warnf("invalid subtitle output path: %v", err)
		return err
	}

	if !upgrade {
		if _, err := os.Stat(out); err == nil {
			logger.Debugf("subtitle already exists: %s", out)
			return nil
		}
	}

	if upgrade {
		if oldData, err := os.ReadFile(out); err == nil {
			if len(data) <= len(oldData) {
				logger.Debugf("existing subtitle %s is higher quality", out)
				return nil
			}
		}
	}

	if err := os.WriteFile(out, data, 0644); err != nil {
		logger.Warnf("write: %v", err)
		return err
	}
	logger.Infof("downloaded %s subtitle %s using profile", actualLang, out)
	if store != nil {
		_ = store.InsertDownload(&database.DownloadRecord{File: out, VideoFile: sanitizedPath, Provider: providerName, Language: actualLang})
	}
	return nil
}
