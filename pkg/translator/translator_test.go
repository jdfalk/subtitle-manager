package translator

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGoogleTranslate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		r.Body.Close()
		s := string(body)
		if !(strings.Contains(s, "q=hello") && strings.Contains(s, "target=es") && strings.Contains(s, "key=test")) {
			t.Fatalf("unexpected request body: %s", s)
		}
		fmt.Fprint(w, `{"data":{"translations":[{"translatedText":"hola"}]}}`)
	}))
	defer srv.Close()

	origURL := googleAPIURL
	SetGoogleAPIURL(srv.URL)
	defer SetGoogleAPIURL(origURL)

	got, err := GoogleTranslate("hello", "es", "test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hola" {
		t.Fatalf("expected hola, got %s", got)
	}
}

func TestUnsupportedServiceError(t *testing.T) {
	if ErrUnsupportedService.Error() == "" {
		t.Fatal("error string empty")
	}
}
