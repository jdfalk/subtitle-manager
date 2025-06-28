// Package syncer provides synchronization utilities for subtitle files.
// It supports batch processing, audio transcription alignment, and embedded subtitle synchronization.
//
// This package is used by subtitle-manager to align subtitle timing with media files.
package syncer

import (
	"os"
	"time"

	"github.com/asticode/go-astisub"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/security"
)

// BatchItem specifies a single subtitle synchronization request.
// Media is the video file path, Subtitle is the subtitle to adjust and
// Output is the destination file written in SRT format. When Output is
// empty the Subtitle path is overwritten.
type BatchItem struct {
	Media    string
	Subtitle string
	Output   string
}

// SyncBatch synchronizes multiple subtitle files in sequence using the given
// options. The returned slice contains an error entry for each item processed.
// A nil value indicates the file was synchronized successfully.
func SyncBatch(items []BatchItem, opts Options) []error {
	logger := logging.GetLogger("syncer.batch")
	logger.Infof("starting batch sync of %d items", len(items))
	start := time.Now()

	errs := make([]error, len(items))
	for i, it := range items {
		itemStart := time.Now()
		logger.Infof("[%d/%d] syncing %s", i+1, len(items), it.Subtitle)

		outPath := it.Output
		if outPath == "" {
			outPath = it.Subtitle
		}
		var err error
		outPath, err = security.ValidateAndSanitizePath(outPath)
		if err != nil {
			logger.Warnf("[%d/%d] path validation failed for %s: %v", i+1, len(items), it.Subtitle, err)
			errs[i] = err
			continue
		}

		result, err := Sync(it.Media, it.Subtitle, opts)
		if err != nil {
			logger.Warnf("[%d/%d] sync failed for %s: %v", i+1, len(items), it.Subtitle, err)
			errs[i] = err
			continue
		}

		sub := astisub.Subtitles{Items: result}
		f, ferr := os.Create(outPath)
		if ferr != nil {
			logger.Warnf("[%d/%d] file creation failed for %s: %v", i+1, len(items), outPath, ferr)
			errs[i] = ferr
			continue
		}
		if err := sub.WriteToSRT(f); err != nil {
			logger.Warnf("[%d/%d] SRT write failed for %s: %v", i+1, len(items), outPath, err)
			errs[i] = err
		} else {
			itemDuration := time.Since(itemStart)
			logger.Infof("[%d/%d] ✅ completed %s -> %s in %v", i+1, len(items), it.Subtitle, outPath, itemDuration)
		}
		f.Close()
	}

	totalDuration := time.Since(start)
	logger.Infof("batch sync completed in %v", totalDuration)
	return errs
}
