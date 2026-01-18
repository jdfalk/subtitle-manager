// file: pkg/providers/gestdown/gestdown_test.go
// version: 1.0.0
// guid: 16f526e0-ae08-40c4-bc7e-d4dfc47f7c60

package gestdown

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func TestNewClientDefaults(t *testing.T) {
	client := New()

	if client == nil {
		t.Fatal("expected client")
	}

	if client.APIURL != "https://api.gestdown.com" {
		t.Fatalf("expected default APIURL, got %q", client.APIURL)
	}

	if client.HTTPClient == nil {
		t.Fatal("expected HTTP client")
	}
}

func TestClientFetch(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/subtitles/movie.mkv/en" {
				t.Fatalf("expected request path to include filename and language, got %q", req.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("subtitle-data"))
		}))
		defer server.Close()

		client := &Client{
			APIURL: server.URL,
			HTTPClient: &http.Client{
				Transport: server.Client().Transport,
			},
		}

		data, err := client.Fetch(ctx, "/media/movie.mkv", "en")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if string(data) != "subtitle-data" {
			t.Fatalf("expected subtitle data, got %q", string(data))
		}
	})

	t.Run("invalid base url", func(t *testing.T) {
		client := &Client{
			APIURL:     "http://[::1",
			HTTPClient: &http.Client{},
		}

		_, err := client.Fetch(ctx, "/media/movie.mkv", "en")
		if err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("transport error", func(t *testing.T) {
		transportErr := errors.New("transport failure")
		client := &Client{
			APIURL: "https://example.invalid",
			HTTPClient: &http.Client{
				Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
					return nil, transportErr
				}),
			},
		}

		_, err := client.Fetch(ctx, "/media/movie.mkv", "en")
		if !errors.Is(err, transportErr) {
			t.Fatalf("expected transport error, got %v", err)
		}
	})

	t.Run("non-200 status", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = io.WriteString(w, "not found")
		}))
		defer server.Close()

		client := &Client{
			APIURL: server.URL,
			HTTPClient: &http.Client{
				Transport: server.Client().Transport,
			},
		}

		_, err := client.Fetch(ctx, "/media/movie.mkv", "en")
		if err == nil {
			t.Fatal("expected error")
		}
	})
}
