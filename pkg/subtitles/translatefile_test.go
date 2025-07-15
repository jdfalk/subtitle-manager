package subtitles

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/translator"
	"github.com/spf13/viper"
)

func TestTranslateFileToSRT(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()["q"]
		if len(qs) == 0 {
			qs = []string{""}
		}
		parts := make([]string, len(qs))
		for i := range qs {
			parts[i] = `{"translatedText":"hola"}`
		}
		fmt.Fprintf(w, `{"data":{"translations":[%s]}}`, strings.Join(parts, ","))
	}))
	defer srv.Close()
	translator.SetGoogleAPIURL(srv.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	// Get absolute path to test data
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	// Use the secure approach: copy test data to temp directory and set media_directory
	src := filepath.Join(wd, "../../testdata/simple.srt")
	inPath := filepath.Join(t.TempDir(), "in.srt")
	inData, _ := os.ReadFile(src)
	os.WriteFile(inPath, inData, 0644)
	viper.Set("media_directory", filepath.Dir(inPath))
	t.Cleanup(viper.Reset)
	out := filepath.Join(t.TempDir(), "out.srt")
	err = TranslateFileToSRT(inPath, out, "es", "google", "k", "", "")
	if err != nil {
		t.Fatalf("translate: %v", err)
	}
	outData, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if !strings.Contains(string(outData), "hola") {
		t.Fatalf("expected translated text, got %s", outData)
	}
}

func TestTranslateFilesToSRT(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()["q"]
		if len(qs) == 0 {
			qs = []string{""}
		}
		parts := make([]string, len(qs))
		for i := range qs {
			parts[i] = `{"translatedText":"hola"}`
		}
		fmt.Fprintf(w, `{"data":{"translations":[%s]}}`, strings.Join(parts, ","))
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

func TestTranslateFileToSRTCache(t *testing.T) {
	count := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		qs := r.URL.Query()["q"]
		if len(qs) == 0 {
			qs = []string{""}
		}
		parts := make([]string, len(qs))
		for i := range qs {
			parts[i] = `{"translatedText":"hola"}`
		}
		fmt.Fprintf(w, `{"data":{"translations":[%s]}}`, strings.Join(parts, ","))
	}))
	defer srv.Close()
	translator.SetGoogleAPIURL(srv.URL)
	defer translator.SetGoogleAPIURL("https://translation.googleapis.com/language/translate/v2")

	dir := t.TempDir()
	viper.Set("media_directory", dir)
	t.Cleanup(viper.Reset)
	in := filepath.Join(dir, "in.srt")
	data := "1\n00:00:01,000 --> 00:00:02,000\nHello\n\n2\n00:00:02,500 --> 00:00:03,500\nHello\n"
	if err := os.WriteFile(in, []byte(data), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	out := filepath.Join(dir, "out.srt")
	if err := TranslateFileToSRT(in, out, "es", "google", "k", "", ""); err != nil {
		t.Fatalf("translate: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 request, got %d", count)
	}
}
