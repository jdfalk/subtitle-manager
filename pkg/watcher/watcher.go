package watcher

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"

	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
)

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

// WatchDirectory monitors dir for new video files and downloads subtitles using
// provider p for the given language. Subtitles are written next to the media
// file with the language code appended before the extension.
func WatchDirectory(ctx context.Context, dir, lang string, p providers.Provider) error {
	logger := logging.GetLogger("watcher")
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer w.Close()

	if err := w.Add(dir); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-w.Errors:
			logger.Warnf("watch error: %v", err)
		case ev := <-w.Events:
			if ev.Op&(fsnotify.Create|fsnotify.Rename|fsnotify.Write) != 0 && isVideoFile(ev.Name) {
				out := strings.TrimSuffix(ev.Name, filepath.Ext(ev.Name)) + "." + lang + ".srt"
				if _, err := os.Stat(out); err == nil {
					continue
				}
				data, err := p.Fetch(ctx, ev.Name, lang)
				if err != nil {
					logger.Warnf("fetch %s: %v", ev.Name, err)
					continue
				}
				if err := os.WriteFile(out, data, 0644); err != nil {
					logger.Warnf("write %s: %v", out, err)
					continue
				}
				logger.Infof("downloaded subtitle %s", out)
			}
		}
	}
}
