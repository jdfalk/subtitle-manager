package subtitles

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/asticode/go-astisub"
	"github.com/sourcegraph/conc/pool"

	"github.com/jdfalk/subtitle-manager/pkg/translator"
)

// TranslateFileToSRT translates the subtitle file at inPath using the
// specified translation service and writes an SRT file to outPath.
// googleKey, gptKey and grpcAddr are passed to the underlying provider
// depending on the service selected.
func TranslateFileToSRT(inPath, outPath, lang, service, googleKey, gptKey, grpcAddr string) error {
	sub, err := astisub.OpenFile(inPath)
	if err != nil {
		return err
	}
	for _, item := range sub.Items {
		t, err := translator.Translate(service, item.String(), lang, googleKey, gptKey, grpcAddr)
		if err != nil {
			return err
		}
		item.Lines = []astisub.Line{{Items: []astisub.LineItem{{Text: t}}}}
	}
	buf := &bytes.Buffer{}
	if err := sub.WriteToSRT(buf); err != nil {
		return err
	}
	return os.WriteFile(outPath, buf.Bytes(), 0644)
}

// TranslateFilesToSRT concurrently translates each file in paths using
// TranslateFileToSRT. Output files are written next to the inputs with the
// language code appended before the extension. The number of worker
// goroutines is limited by workers.
func TranslateFilesToSRT(paths []string, lang, service, googleKey, gptKey, grpcAddr string, workers int) error {
	p := pool.New().WithErrors().WithMaxGoroutines(workers)
	for _, in := range paths {
		in := in
		out := strings.TrimSuffix(in, filepath.Ext(in)) + "." + lang + ".srt"
		p.Go(func() error {
			return TranslateFileToSRT(in, out, lang, service, googleKey, gptKey, grpcAddr)
		})
	}
	return p.Wait()
}
