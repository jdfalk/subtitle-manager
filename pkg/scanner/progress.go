// file: pkg/scanner/progress.go
package scanner

import (
	"context"
	"os"
	"path/filepath"

	"github.com/sourcegraph/conc/pool"

	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/logging"
	"subtitle-manager/pkg/providers"
)

// ProgressFunc is called with each processed video file path.
type ProgressFunc func(file string)

// ScanDirectoryProgress walks through dir and downloads subtitles like
// ScanDirectory, invoking cb for each processed file.
func ScanDirectoryProgress(ctx context.Context, dir, lang, providerName string,
	p providers.Provider, upgrade bool, workers int, store database.SubtitleStore, cb ProgressFunc) error {
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
	return work.Wait()
}
