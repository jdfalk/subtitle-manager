package opensubtitles

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

type mockHandler struct{}

func (mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/search") {
		fmt.Fprintf(w, `[{"SubDownloadLink":"http://%s/download"}]`, r.Host)
		return
	}
	if r.URL.Path == "/download" {
		fmt.Fprint(w, "sub data")
		return
	}
	w.WriteHeader(404)
}

func TestFetch(t *testing.T) {
	srv := httptest.NewServer(mockHandler{})
	defer srv.Close()
	c := New("")
	c.APIURL = srv.URL
	c.HTTPClient = srv.Client()
	// override fileHash to avoid reading a real file
	orig := fileHashFunc
	fileHashFunc = func(string) (uint64, int64, error) { return 1, 1, nil }
	defer func() { fileHashFunc = orig }()

	data, err := c.Fetch(context.Background(), "dummy.mkv", "en")
	if err != nil {
		t.Fatalf("fetch error: %v", err)
	}
	if string(data) != "sub data" {
		t.Fatalf("unexpected data: %s", data)
	}
}

// TestNewUsesConfig verifies that viper settings override defaults.
func TestNewUsesConfig(t *testing.T) {
	viper.Set("opensubtitles.api_url", "http://api")
	viper.Set("opensubtitles.user_agent", "ua")
	defer viper.Reset()
	c := New("k")
	if c.APIURL != "http://api" {
		t.Fatalf("expected api_url http://api, got %s", c.APIURL)
	}
	if c.UserAgent != "ua" {
		t.Fatalf("expected user_agent ua, got %s", c.UserAgent)
	}
}
