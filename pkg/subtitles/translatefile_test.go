package subtitles

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"subtitle-manager/pkg/translator"
)

func TestTranslateFileToSRT(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"data":{"translations":[{"translatedText":"hola"}]}}`)
	}))
	defer srv.Close()
	translator.SetGoogleAPIURL(srv.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	out := filepath.Join(t.TempDir(), "out.srt")
	err := TranslateFileToSRT("../../testdata/simple.srt", out, "es", "google", "k", "", "")
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if !strings.Contains(string(data), "hola") {
		t.Fatalf("expected translated text, got %s", data)
	}
}

func TestTranslateFilesToSRT(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"data":{"translations":[{"translatedText":"hola"}]}}`)
	}))
	defer srv.Close()
	translator.SetGoogleAPIURL(srv.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	dir := t.TempDir()
	inputs := []string{}
	for i := 0; i < 2; i++ {
		in := filepath.Join(dir, fmt.Sprintf("in%d.srt", i))
		b, _ := os.ReadFile("../../testdata/simple.srt")
		os.WriteFile(in, b, 0644)
		inputs = append(inputs, in)
	}
	if err := TranslateFilesToSRT(inputs, "es", "google", "k", "", "", 2); err != nil {
		t.Fatalf("batch: %v", err)
	}
	for _, in := range inputs {
		out := strings.TrimSuffix(in, filepath.Ext(in)) + ".es.srt"
		data, err := os.ReadFile(out)
		if err != nil {
			t.Fatalf("read out: %v", err)
		}
		if !strings.Contains(string(data), "hola") {
			t.Fatalf("missing translation in %s", out)
		}
	}
}
