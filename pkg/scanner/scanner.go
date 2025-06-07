package scanner

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
)

// ScanDirectory traverses dir looking for video files and downloads subtitles
// using provider p for the given language. If upgrade is false existing
// subtitle files are skipped. Subtitles are saved next to the video with the
// language code appended before the extension.
func ScanDirectory(ctx context.Context, dir, lang string, p providers.Provider, upgrade bool) error {
	logger := logging.GetLogger("scanner")
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !isVideoFile(path) {
			return nil
		}
		out := strings.TrimSuffix(path, filepath.Ext(path)) + "." + lang + ".srt"
		if !upgrade {
			if _, err := os.Stat(out); err == nil {
				return nil
			}
		}
		data, err := p.Fetch(ctx, path, lang)
		if err != nil {
			logger.Warnf("fetch %s: %v", path, err)
			return nil
		}
		if err := os.WriteFile(out, data, 0644); err != nil {
			logger.Warnf("write %s: %v", out, err)
			return nil
		}
		logger.Infof("downloaded subtitle %s", out)
		return nil
	})
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
