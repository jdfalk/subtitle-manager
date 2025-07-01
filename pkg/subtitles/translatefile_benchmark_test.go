package subtitles

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/translator"
)

func BenchmarkTranslateFileToSRT(b *testing.B) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"data":{"translations":[{"translatedText":"hola"}]}}`)
	}))
	defer srv.Close()
	translator.SetGoogleAPIURL(srv.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	// Get absolute path to test data
	wd, err := os.Getwd()
	if err != nil {
		b.Fatalf("getwd: %v", err)
	}
	inPath := filepath.Join(wd, "../../testdata/simple.srt")

	for i := 0; i < b.N; i++ {
		out := filepath.Join(b.TempDir(), fmt.Sprintf("out-%d.srt", i))
		if err := TranslateFileToSRT(inPath, out, "es", "google", "k", "", ""); err != nil {
			b.Fatalf("translate: %v", err)
		}
	}
}

func BenchmarkTranslateFileToSRTBatch(b *testing.B) {
	// Mock server that supports batch translation
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate batch response - this is simplified for benchmark
		fmt.Fprint(w, `{"data":{"translations":[{"translatedText":"hola"},{"translatedText":"adiós"}]}}`)
	}))
	defer srv.Close()
	translator.SetGoogleAPIURL(srv.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	for i := 0; i < b.N; i++ {
		out := filepath.Join(b.TempDir(), fmt.Sprintf("out-batch-%d.srt", i))
		if err := TranslateFileToSRT("../../testdata/simple.srt", out, "es", "google", "test-key", "", ""); err != nil {
			b.Fatalf("translate batch: %v", err)
		}
	}
}

func BenchmarkTranslateFileToSRTNonGoogle(b *testing.B) {
	// Test non-Google service for comparison
	for i := 0; i < b.N; i++ {
		out := filepath.Join(b.TempDir(), fmt.Sprintf("out-gpt-%d.srt", i))
		// This will fail due to no API key, but measures the overhead
		_ = TranslateFileToSRT("../../testdata/simple.srt", out, "es", "gpt", "", "fake-key", "")
	}
}
