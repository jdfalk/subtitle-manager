// file: pkg/webserver/search_test.go
package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSearchHandlerMethodValidation(t *testing.T) {
	handler := searchHandler()

	// Test PUT method - should return 405
	req, err := http.NewRequest("PUT", "/api/search", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestSearchHandlerPostValidation(t *testing.T) {
	handler := searchHandler()

	// Test with empty body - should return 400
	req, err := http.NewRequest("POST", "/api/search", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSearchHandlerPostValidRequest(t *testing.T) {
	handler := searchHandler()

	// Set TEST_SAFE_MEDIA_DIR so the handler accepts the test path
	oldEnv := os.Getenv("TEST_SAFE_MEDIA_DIR")
	os.Setenv("TEST_SAFE_MEDIA_DIR", "/")
	defer os.Setenv("TEST_SAFE_MEDIA_DIR", oldEnv)

	// Create a valid search request
	searchReq := SearchRequest{
		Providers: []string{"embedded"},
		MediaPath: "/nonexistent/path.mkv", // This will fail but should validate request structure
		Language:  "en",
	}

	reqBody, err := json.Marshal(searchReq)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/search", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Should return 404 because file doesn't exist, but request structure is valid
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestSearchHandlerGetValidation(t *testing.T) {
	handler := searchHandler()

	// Test GET without required path parameter - should return 400
	req, err := http.NewRequest("GET", "/api/search", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSearchPreviewHandlerMethodValidation(t *testing.T) {
	handler := searchPreviewHandler()

	// Test POST method - should return 405
	req, err := http.NewRequest("POST", "/api/search/preview", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestSearchPreviewHandlerUrlValidation(t *testing.T) {
	handler := searchPreviewHandler()

	// Test GET without URL parameter - should return 400
	req, err := http.NewRequest("GET", "/api/search/preview", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCalculateNameMatch(t *testing.T) {
	tests := []struct {
		subtitleName string
		mediaPath    string
		expected     float64
	}{
		{"movie 2023 1080p srt", "/path/to/movie 2023 1080p.mkv", 1.0},      // All words match
		{"different movie srt", "/path/to/movie 2023 1080p.mkv", 1.0 / 3.0}, // 1 out of 3 words match (movie)
		{"", "/path/to/movie.mkv", 0.0},
		{"movie", "", 0.0},
		{"movie.2023.1080p.srt", "/path/to/movie.2023.1080p.mkv", 1.0}, // Exact match (treated as single word)
	}

	for _, test := range tests {
		result := calculateNameMatch(test.subtitleName, test.mediaPath)
		if abs(result-test.expected) > 0.001 { // Use tolerance for floating point comparison
			t.Errorf("calculateNameMatch(%q, %q) = %v; want %v",
				test.subtitleName, test.mediaPath, result, test.expected)
		}
	}
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

func TestExtractNameFromURL(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"http://example.com/subtitle.srt", "subtitle.srt"},
		{"http://example.com/path/to/subtitle.srt?param=value", "subtitle.srt"},
		{"http://example.com/", "Subtitle"},
		{"", "Subtitle"},
	}

	for _, test := range tests {
		result := extractNameFromURL(test.url)
		if result != test.expected {
			t.Errorf("extractNameFromURL(%q) = %q; want %q",
				test.url, result, test.expected)
		}
	}
}
