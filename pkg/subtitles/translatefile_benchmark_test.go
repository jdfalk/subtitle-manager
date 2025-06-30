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
