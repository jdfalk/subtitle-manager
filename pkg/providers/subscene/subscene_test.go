package subscene

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestClientFetch verifies that the client downloads subtitles using the expected URL.
func TestClientFetch(t *testing.T) {
	var gotURL string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotURL = r.URL.String()
		fmt.Fprint(w, "data")
	}))
	defer srv.Close()

	c := New()
	c.APIURL = srv.URL
	b, err := c.Fetch(context.Background(), "/path/movie.mkv", "en")
	if err != nil {
		t.Fatalf("fetch: %v", err)
	}
	if string(b) != "data" {
		t.Fatalf("unexpected body: %s", b)
	}
	if gotURL != "/subtitles/movie.mkv/en" {
		t.Fatalf("unexpected url: %s", gotURL)
	}
}
