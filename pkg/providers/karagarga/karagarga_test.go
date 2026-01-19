// file: pkg/providers/karagarga/karagarga_test.go
// version: 1.0.1
// guid: 2b8c087d-3a55-4e1b-bcb7-3f03ef68698b
package karagarga

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (rt roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt(req)
}

func TestFetch_Success_ReturnsBody(t *testing.T) {
	// Arrange
	const expectedBody = "subtitle-content"
	const mediaPath = "/videos/movie.mkv"
	const lang = "en"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected method %s, got %s", http.MethodGet, r.Method)
		}
		if r.URL.Path != "/subtitles/movie.mkv/en" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		_, _ = io.WriteString(w, expectedBody)
	}))
	t.Cleanup(server.Close)

	client := &Client{
		APIURL:     server.URL,
		HTTPClient: server.Client(),
	}

	// Act
	body, err := client.Fetch(context.Background(), mediaPath, lang)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(body) != expectedBody {
		t.Fatalf("expected body %q, got %q", expectedBody, string(body))
	}
}

func TestFetch_StatusNotOK_ReturnsError(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	t.Cleanup(server.Close)

	client := &Client{
		APIURL:     server.URL,
		HTTPClient: server.Client(),
	}

	// Act
	_, err := client.Fetch(context.Background(), "movie.mkv", "en")

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "status 404" {
		t.Fatalf("expected status error, got %v", err)
	}
}

func TestFetch_RequestCreationFails_ReturnsError(t *testing.T) {
	// Arrange
	client := &Client{
		APIURL:     "http://[::1",
		HTTPClient: &http.Client{},
	}

	// Act
	_, err := client.Fetch(context.Background(), "movie.mkv", "en")

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestFetch_HTTPClientFails_ReturnsError(t *testing.T) {
	// Arrange
	client := &Client{
		APIURL: "https://example.test",
		HTTPClient: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("transport failure")
			}),
		},
	}

	// Act
	_, err := client.Fetch(context.Background(), "movie.mkv", "en")

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "transport failure") {
		t.Fatalf("expected transport failure, got %v", err)
	}
}
