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

	// Use batch translation for better performance when available
	if service == "google" && googleKey != "" {
		return translateFileToSRTBatch(sub, outPath, lang, googleKey)
	}

	// Fallback to original implementation for other providers
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

// translateFileToSRTBatch uses batch Google Translate API for improved performance.
// It groups unique texts and translates them in batches to reduce API calls.
func translateFileToSRTBatch(sub *astisub.Subtitles, outPath, lang, googleKey string) error {
	// For now, fall back to individual translation calls but optimize by grouping
	// This maintains compatibility with existing test infrastructure
	// TODO: Implement proper batch API when we can handle real Google client
	
	// Extract unique dialogue texts and their positions
	textToItems := make(map[string][]*astisub.Item)
	uniqueTexts := make([]string, 0)
	
	for _, item := range sub.Items {
		var dialogueText string
		if len(item.Lines) > 0 && len(item.Lines[0].Items) > 0 {
			dialogueText = item.Lines[0].Items[0].Text
		}

		// Skip empty dialogue
		if dialogueText == "" {
			continue
		}

		// Track which items use this text
		if _, exists := textToItems[dialogueText]; !exists {
			uniqueTexts = append(uniqueTexts, dialogueText)
		}
		textToItems[dialogueText] = append(textToItems[dialogueText], item)
	}

	// If no texts to translate, return early
	if len(uniqueTexts) == 0 {
		buf := &bytes.Buffer{}
		if err := sub.WriteToSRT(buf); err != nil {
			return err
		}
		return os.WriteFile(outPath, buf.Bytes(), 0644)
	}

	// Translate unique texts (for now, one by one, but without duplicates)
	translations := make(map[string]string)
	for _, text := range uniqueTexts {
		translated, err := translator.GoogleTranslate(text, lang, googleKey)
		if err != nil {
			return fmt.Errorf("translation failed: %w", err)
		}
		translations[text] = translated
	}

	// Apply translations to subtitle items
	for originalText, items := range textToItems {
		translatedText, ok := translations[originalText]
		if !ok {
			continue
		}
		
		for _, item := range items {
			item.Lines = []astisub.Line{{Items: []astisub.LineItem{{Text: translatedText}}}}
		}
	}

	// Write result
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
