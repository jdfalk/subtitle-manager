// file: pkg/scanner/scanner.go
package scanner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sourcegraph/conc/pool"

	"github.com/jdfalk/subtitle-manager/pkg/database"
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
	// Validate the language code to ensure it conforms to expected patterns
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
	out, err := security.ValidateSubtitleOutputPath(path, lang)
	if err != nil {
		logger.Warnf("invalid subtitle output path: %v", err)
		return err
	}
	if !upgrade {
		// Validate the output path again immediately before file stat to prevent path injection
		validatedOut, err := security.ValidateAndSanitizePath(out)
		if err != nil {
			logger.Warnf("invalid output path before stat: %v", err)
			return err
		}
		if _, err := os.Stat(validatedOut); err == nil {
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
		return err
	}
	if upgrade {
		// Validate the output path again immediately before file read to prevent path injection
		validatedOut, err := security.ValidateAndSanitizePath(out)
		if err != nil {
			logger.Warnf("invalid output path before read: %v", err)
			return err
		}
		if oldData, err := os.ReadFile(validatedOut); err == nil {
			if len(data) <= len(oldData) {
				logger.Debugf("existing subtitle %s is higher quality", out)
				return nil
			}
		}
	}
	// Validate the output path again immediately before file write to prevent path injection
	validatedOut, err := security.ValidateAndSanitizePath(out)
	if err != nil {
		logger.Warnf("invalid output path before write: %v", err)
		return err
	}
	if err := os.WriteFile(validatedOut, data, 0644); err != nil {
		logger.Warnf("write %s: %v", validatedOut, err)
		return err
	}
	logger.Infof("downloaded subtitle %s", validatedOut)
	if store != nil {
		_ = store.InsertDownload(&database.DownloadRecord{File: out, VideoFile: path, Provider: providerName, Language: lang})
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
