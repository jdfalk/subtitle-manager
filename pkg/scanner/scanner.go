package scanner

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/sourcegraph/conc/pool"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
)

// ScanDirectory walks through the directory and downloads subtitles for video files
// using provider p for the given language. providerName is stored in download
// history. If upgrade is false existing subtitle files are skipped.
func ScanDirectory(ctx context.Context, dir, lang string, providerName string, p providers.Provider, upgrade bool, workers int, store database.SubtitleStore) error {
	logger := logging.GetLogger("scanner")
	work := pool.New().WithErrors().WithMaxGoroutines(workers)
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !isVideoFile(path) {
			return nil
		}
		f := path
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
// file is left untouched.
func ProcessFile(ctx context.Context, path, lang string, providerName string, p providers.Provider, upgrade bool, store database.SubtitleStore) error {
	logger := logging.GetLogger("scanner")
	
	// Validate the lang parameter to ensure it contains only alphanumeric characters
	if !isValidLang(lang) {
		logger.Warnf("invalid language code: %s", lang)
		return fmt.Errorf("invalid language code: %s", lang)
	}
	
	out := strings.TrimSuffix(path, filepath.Ext(path)) + "." + lang + ".srt"
	if !upgrade {
		if _, err := os.Stat(out); err == nil {
			return nil
		}
	}
	var data []byte
	var err error
	if p != nil {
		data, err = p.Fetch(ctx, path, lang)
	} else {
		data, providerName, err = providers.FetchFromAll(ctx, path, lang, "")
	}
	if err != nil {
		logger.Warnf("fetch %s: %v", path, err)
		return err
	}
	if err := os.WriteFile(out, data, 0644); err != nil {
		logger.Warnf("write %s: %v", out, err)
		return err
	}
	logger.Infof("downloaded subtitle %s", out)
	if store != nil {
		_ = store.InsertDownload(&database.DownloadRecord{File: out, VideoFile: path, Provider: providerName, Language: lang})
	}
	return nil
}

// isValidLang checks if the language code contains only alphanumeric characters
func isValidLang(lang string) bool {
	for _, r := range lang {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
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
