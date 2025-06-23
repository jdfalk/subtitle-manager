package syncer

import (
	"os"

	"github.com/asticode/go-astisub"
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
	errs := make([]error, len(items))
	for i, it := range items {
		outPath := it.Output
		if outPath == "" {
			outPath = it.Subtitle
		}
		var err error
		outPath, err = security.ValidateAndSanitizePath(outPath)
		if err != nil {
			errs[i] = err
			continue
		}
		result, err := Sync(it.Media, it.Subtitle, opts)
		if err != nil {
			errs[i] = err
			continue
		}
		sub := astisub.Subtitles{Items: result}
		f, ferr := os.Create(outPath)
		if ferr != nil {
			errs[i] = ferr
			continue
		}
		if err := sub.WriteToSRT(f); err != nil {
			errs[i] = err
		}
		f.Close()
	}
	return errs
}
