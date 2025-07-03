// file: pkg/webserver/whisper_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174004

package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWhisperModelsHandler(t *testing.T) {
	handler := whisperModelsHandler()

	req := httptest.NewRequest("GET", "/api/whisper/models", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	models, ok := response["models"].([]interface{})
	if !ok {
		t.Fatal("Expected models field to be an array")
	}

	if len(models) == 0 {
		t.Fatal("Expected at least one model")
	}

	// Check that it contains expected models
	modelStrings := make([]string, len(models))
	for i, m := range models {
		modelStrings[i] = m.(string)
	}

	expectedModels := []string{"tiny", "base", "small", "medium", "large"}
	for _, expected := range expectedModels {
		found := false
		for _, actual := range modelStrings {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected model %q not found in response", expected)
		}
	}
}

func TestWhisperModelsHandlerMethodNotAllowed(t *testing.T) {
	handler := whisperModelsHandler()

	req := httptest.NewRequest("POST", "/api/whisper/models", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

func TestWhisperTranscribeHandlerInvalidMethod(t *testing.T) {
	handler := whisperTranscribeHandler()

	req := httptest.NewRequest("GET", "/api/whisper/transcribe", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

func TestWhisperTranscribeHandlerInvalidJSON(t *testing.T) {
	handler := whisperTranscribeHandler()

	req := httptest.NewRequest("POST", "/api/whisper/transcribe", bytes.NewReader([]byte("invalid json")))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestWhisperTranscribeHandlerMissingFilePath(t *testing.T) {
	handler := whisperTranscribeHandler()

	body := map[string]string{
		"language": "en",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/api/whisper/transcribe", bytes.NewReader(jsonBody))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestWhisperTranscribeHandlerInvalidFileExtension(t *testing.T) {
	handler := whisperTranscribeHandler()

	body := map[string]string{
		"file_path": "/path/to/file.txt",
		"language":  "en",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/api/whisper/transcribe", bytes.NewReader(jsonBody))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}
