package subtitles

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/asticode/go-astisub"
	"github.com/sourcegraph/conc/pool"

	"github.com/jdfalk/subtitle-manager/pkg/security"
	"github.com/jdfalk/subtitle-manager/pkg/translator"
)

// TranslateFileToSRT translates the subtitle file at inPath using the
// specified translation service and writes an SRT file to outPath.
// googleKey, gptKey and grpcAddr are passed to the underlying provider
// depending on the service selected. Results are cached in-memory so
// identical lines are only translated once per invocation.
func TranslateFileToSRT(inPath, outPath, lang, service, googleKey, gptKey, grpcAddr string) error {
	// Validate and sanitize input paths to prevent path injection attacks
	validatedInPath, err := security.ValidateAndSanitizePath(inPath)
	if err != nil {
		return fmt.Errorf("invalid input path: %w", err)
	}

	validatedOutPath, err := security.ValidateAndSanitizePath(outPath)
	if err != nil {
		return fmt.Errorf("invalid output path: %w", err)
	}

	sub, err := astisub.OpenFile(validatedInPath)
	if err != nil {
		return err
	}
	cache := make(map[string]string, len(sub.Items))
	for _, item := range sub.Items {
		// Extract just the dialogue text for translation and caching,
		// avoiding timestamps and sequence numbers that prevent deduplication
		var dialogueText string
		if len(item.Lines) > 0 && len(item.Lines[0].Items) > 0 {
			dialogueText = item.Lines[0].Items[0].Text
		}

		// Skip empty dialogue
		if dialogueText == "" {
			continue
		}

		// Check cache for existing translation
		t, ok := cache[dialogueText]
		if !ok {
			t, err = translator.Translate(service, dialogueText, lang, googleKey, gptKey, grpcAddr)
			if err != nil {
				return err
			}
			cache[dialogueText] = t
		}
		item.Lines = []astisub.Line{{Items: []astisub.LineItem{{Text: t}}}}
	}
	buf := &bytes.Buffer{}
	if err := sub.WriteToSRT(buf); err != nil {
		return err
	}
	return os.WriteFile(validatedOutPath, buf.Bytes(), 0644)
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
