package scanner

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/sourcegraph/conc/pool"

	"subtitle-manager/pkg/database"
// using provider p for the given language. providerName is stored in download
// history. If upgrade is false existing
func ScanDirectory(ctx context.Context, dir, lang string, providerName string, p providers.Provider, upgrade bool, workers int, store database.SubtitleStore) error {
			return ProcessFile(ctx, f, lang, providerName, p, upgrade, store)
// ProcessFile downloads a subtitle for path using providerName for history
// tracking. The subtitle is saved next to the media file with the language
// code appended before the extension. If upgrade is false an existing subtitle
// file is left untouched.
func ProcessFile(ctx context.Context, path, lang string, providerName string, p providers.Provider, upgrade bool, store database.SubtitleStore) error {
	if store != nil {
		_ = store.InsertDownload(&database.DownloadRecord{File: out, VideoFile: path, Provider: providerName, Language: lang})
	}
// subtitle files are skipped. Subtitles are saved next to the video with the
// language code appended before the extension.
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
	out := strings.TrimSuffix(path, filepath.Ext(path)) + "." + lang + ".srt"
	if !upgrade {
		if _, err := os.Stat(out); err == nil {
			return nil
		}
	}
	data, err := p.Fetch(ctx, path, lang)
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
