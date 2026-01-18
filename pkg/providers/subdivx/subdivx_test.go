package subdivx

import (
	"context"
	"errors"
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

func TestClientFetchBadURL(t *testing.T) {
	ctx := context.Background()

	c := New()
	c.APIURL = "http://[::1"

	_, err := c.Fetch(ctx, "/path/movie.mkv", "en")
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}
}

func TestClientFetchHTTPError(t *testing.T) {
	ctx := context.Background()

	c := New()
	c.HTTPClient = &http.Client{
		Transport: roundTripper(func(*http.Request) (*http.Response, error) {
			return nil, errors.New("transport error")
		}),
	}

	_, err := c.Fetch(ctx, "/path/movie.mkv", "en")
	if err == nil {
		t.Fatal("expected error from HTTP client, got nil")
	}
}

func TestClientFetchStatusError(t *testing.T) {
	ctx := context.Background()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := New()
	c.APIURL = srv.URL

	_, err := c.Fetch(ctx, "/path/movie.mkv", "en")
	if err == nil {
		t.Fatal("expected error for non-200 status, got nil")
	}
	if err.Error() != "status 500" {
		t.Fatalf("unexpected error: %v", err)
	}
}

type roundTripper func(*http.Request) (*http.Response, error)

func (rt roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt(req)
}
