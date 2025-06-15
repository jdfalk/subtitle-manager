package cmd

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

// TestBatchCmd verifies the batch command translates multiple files concurrently.
func TestBatchCmd(t *testing.T) {
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
		data, _ := os.ReadFile("../testdata/simple.srt")
		os.WriteFile(in, data, 0644)
		inputs = append(inputs, in)
	}

	viper.Set("translate_service", "google")
	viper.Set("google_api_key", "k")
	viper.Set("openai_api_key", "")
	viper.Set("grpc_addr", "")
	viper.Set("batch_workers", 2)
	defer viper.Reset()
	if err := batchCmd.RunE(batchCmd, append([]string{"es"}, inputs...)); err != nil {
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
