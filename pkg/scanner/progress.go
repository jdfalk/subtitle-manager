// file: pkg/scanner/progress.go
// version: 1.1.0
// guid: 30d76902-4260-48c3-8dbb-acbbdc9bcea7
package scanner

import (
	"context"
	"os"
	"path/filepath"

	"github.com/sourcegraph/conc/pool"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/metadata"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/security"
)

// ProgressFunc is called with each processed video file path.
type ProgressFunc func(file string)

// ScanDirectoryProgress walks through dir and downloads subtitles like
// ScanDirectory, invoking cb for each processed file.
func ScanDirectoryProgress(ctx context.Context, dir, lang, providerName string,
	p providers.Provider, upgrade bool, workers int, store database.SubtitleStore, cb ProgressFunc) error {
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
		f := path
		work.Go(func() error {
			logger.Debugf("process %s", f)
			if err := ProcessFile(ctx, f, lang, providerName, p, upgrade, store); err == nil {
				if cb != nil {
					cb(f)
				}
			}
			return nil
		})
		return nil
	})
	if err != nil {
		return err
	}
	if err := work.Wait(); err != nil {
		return err
	}
	if store != nil {
		if err := metadata.ScanLibrary(ctx, sanitizedDir, store); err != nil {
			logger.Warnf("scan library: %v", err)
		}
	}
	return nil
}
