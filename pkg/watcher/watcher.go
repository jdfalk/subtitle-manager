package watcher

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
	"github.com/jdfalk/subtitle-manager/pkg/security"
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
func WatchDirectory(ctx context.Context, dir, lang, providerName string, p providers.Provider, store database.SubtitleStore) error {
	logger := logging.GetLogger("watcher")

	sanitizedDir, err := security.ValidateAndSanitizePath(dir)
	if err != nil {
		return err
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer w.Close()

	if err := w.Add(sanitizedDir); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-w.Errors:
			logger.Warnf("watch error: %v", err)
		case ev := <-w.Events:
			sanitizedName, err := security.ValidateAndSanitizePath(ev.Name)
			if err != nil {
				logger.Warnf("ignore invalid path %s: %v", ev.Name, err)
				continue
			}
			if ev.Op&(fsnotify.Create|fsnotify.Rename|fsnotify.Write) != 0 && isVideoFile(sanitizedName) {
				if err := scanner.ProcessFile(ctx, sanitizedName, lang, providerName, p, false, store); err != nil {
					logger.Warnf("process %s: %v", sanitizedName, err)
				}
			}
		}
	}
}

// WatchDirectoryRecursive works like WatchDirectory but monitors dir and all
// of its subdirectories. New directories created while watching are added
// automatically.
func WatchDirectoryRecursive(ctx context.Context, dir, lang, providerName string, p providers.Provider, store database.SubtitleStore) error {
	logger := logging.GetLogger("watcher")

	sanitizedDir, err := security.ValidateAndSanitizePath(dir)
	if err != nil {
		return err
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer w.Close()

	if err := filepath.WalkDir(sanitizedDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return w.Add(path)
		}
		return nil
	}); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-w.Errors:
			logger.Warnf("watch error: %v", err)
		case ev := <-w.Events:
			sanitizedName, err := security.ValidateAndSanitizePath(ev.Name)
			if err != nil {
				logger.Warnf("ignore invalid path %s: %v", ev.Name, err)
				continue
			}
			if ev.Op&fsnotify.Create != 0 {
				if info, err := os.Stat(sanitizedName); err == nil && info.IsDir() {
					_ = w.Add(sanitizedName)
				}
			}
			if ev.Op&(fsnotify.Create|fsnotify.Rename|fsnotify.Write) != 0 && isVideoFile(sanitizedName) {
				if err := scanner.ProcessFile(ctx, sanitizedName, lang, providerName, p, false, store); err != nil {
					logger.Warnf("process %s: %v", sanitizedName, err)
				}
			}
		}
	}
}
