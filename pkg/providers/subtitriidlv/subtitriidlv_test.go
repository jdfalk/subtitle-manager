// file: pkg/providers/subtitriidlv/subtitriidlv_test.go
// version: 1.0.0
// guid: 2a448c2f-0130-4121-a13f-16ab6ea740a5
package subtitriidlv

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

// TestNew_Defaults_ReturnsConfiguredClient verifies the default configuration values.
func TestNew_Defaults_ReturnsConfiguredClient(t *testing.T) {
	// Arrange

	// Act
	client := New()

	// Assert
	if client == nil {
		t.Fatal("expected non-nil client")
	}
	if client.APIURL != "https://api.subtitriidlv.com" {
		t.Fatalf("expected default API URL, got %q", client.APIURL)
	}
	if client.HTTPClient == nil {
		t.Fatal("expected HTTP client to be set")
	}
	if client.HTTPClient.Timeout != 15*time.Second {
		t.Fatalf("expected timeout 15s, got %s", client.HTTPClient.Timeout)
	}
}

// TestClientFetch_Success_ReturnsBody verifies a successful subtitle download.
func TestClientFetch_Success_ReturnsBody(t *testing.T) {
	// Arrange
	mediaPath := filepath.Join("/tmp", "movie.mkv")
	lang := "en"
	body := []byte("subtitle-data")
	var gotPath string

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		gotPath = request.URL.Path
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(body)
	}))
	defer server.Close()

	client := &Client{
		APIURL:     server.URL,
		HTTPClient: server.Client(),
	}

	// Act
	data, err := client.Fetch(context.Background(), mediaPath, lang)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(data) != string(body) {
		t.Fatalf("expected %q, got %q", body, data)
	}
	expectedPath := "/subtitles/movie.mkv/en"
	if gotPath != expectedPath {
		t.Fatalf("expected path %q, got %q", expectedPath, gotPath)
	}
}

// TestClientFetch_StatusNotOK_ReturnsError verifies non-200 responses return an error.
func TestClientFetch_StatusNotOK_ReturnsError(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := &Client{
		APIURL:     server.URL,
		HTTPClient: server.Client(),
	}

	// Act
	data, err := client.Fetch(context.Background(), "/tmp/video.mp4", "fr")

	// Assert
	if err == nil {
		t.Fatal("expected error for non-OK status")
	}
	if data != nil {
		t.Fatalf("expected no data, got %q", data)
	}
}

// TestClientFetch_RequestError_ReturnsError verifies transport errors are surfaced.
func TestClientFetch_RequestError_ReturnsError(t *testing.T) {
	// Arrange
	transportErr := errors.New("transport failure")
	client := &Client{
		APIURL: "https://example.invalid",
		HTTPClient: &http.Client{
			Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
				return nil, transportErr
			}),
		},
	}

	// Act
	data, err := client.Fetch(context.Background(), "/tmp/video.mp4", "es")

	// Assert
	if err == nil {
		t.Fatal("expected error from transport")
	}
	if !errors.Is(err, transportErr) {
		t.Fatalf("expected transport error, got %v", err)
	}
	if data != nil {
		t.Fatalf("expected no data, got %q", data)
	}
}

// TestClientFetch_InvalidURL_ReturnsError verifies invalid URL input fails early.
func TestClientFetch_InvalidURL_ReturnsError(t *testing.T) {
	// Arrange
	client := &Client{
		APIURL:     "http://[::1",
		HTTPClient: &http.Client{},
	}

	// Act
	data, err := client.Fetch(context.Background(), "/tmp/video.mp4", "de")

	// Assert
	if err == nil {
		t.Fatal("expected error from invalid URL")
	}
	if data != nil {
		t.Fatalf("expected no data, got %q", data)
	}
}
