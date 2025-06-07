package opensubtitles

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
