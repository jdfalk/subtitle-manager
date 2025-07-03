// file: pkg/captcha/anticaptcha_test.go
package captcha

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestNewTwoCaptcha verifies that NewTwoCaptcha creates a properly initialized instance.
func TestNewTwoCaptcha(t *testing.T) {
	apiKey := "test-api-key"
	solver := NewTwoCaptcha(apiKey)

	if solver.APIKey != apiKey {
		t.Errorf("expected API key %s, got %s", apiKey, solver.APIKey)
	}

	if solver.BaseURL != "https://2captcha.com" {
		t.Errorf("expected BaseURL https://2captcha.com, got %s", solver.BaseURL)
	}

	if solver.client == nil {
		t.Error("expected client to be initialized, got nil")
	}

	if solver.client.Timeout != 60*time.Second {
		t.Errorf("expected timeout 60s, got %v", solver.client.Timeout)
	}
}

// TestTwoCaptchaSolveSuccess verifies successful captcha solving.
func TestTwoCaptchaSolveSuccess(t *testing.T) {
	captchaID := "test-captcha-id"
	solutionToken := "test-solution-token"

	// Create mock server
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Errorf("expected form content type, got %s", r.Header.Get("Content-Type"))
		}

		response := map[string]any{
			"status":  1,
			"request": captchaID,
		}
		json.NewEncoder(w).Encode(response)
	})

	handler.HandleFunc("/res.php", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}

		response := map[string]any{
			"status":  1,
			"request": solutionToken,
		}
		json.NewEncoder(w).Encode(response)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// Create solver with test server
	solver := &TwoCaptcha{
		APIKey:  "test-key",
		BaseURL: server.URL,
		client:  server.Client(),
	}

	ctx := context.Background()
	result, err := solver.Solve(ctx, "test-site-key", "http://example.com")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != solutionToken {
		t.Errorf("expected solution %s, got %s", solutionToken, result)
	}
}

// TestTwoCaptchaSolveSubmissionError verifies error handling during submission.
func TestTwoCaptchaSolveSubmissionError(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status":  0,
			"request": "ERROR_ZERO_BALANCE",
		}
		json.NewEncoder(w).Encode(response)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	solver := &TwoCaptcha{
		APIKey:  "test-key",
		BaseURL: server.URL,
		client:  server.Client(),
	}

	ctx := context.Background()
	_, err := solver.Solve(ctx, "test-site-key", "http://example.com")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "ERROR_ZERO_BALANCE") {
		t.Errorf("expected error to contain ERROR_ZERO_BALANCE, got %v", err)
	}
}

// TestTwoCaptchaSolvePollingNotReady verifies handling of not-ready responses.
func TestTwoCaptchaSolvePollingNotReady(t *testing.T) {
	captchaID := "test-captcha-id"
	solutionToken := "test-solution-token"
	pollCount := 0

	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status":  1,
			"request": captchaID,
		}
		json.NewEncoder(w).Encode(response)
	})

	handler.HandleFunc("/res.php", func(w http.ResponseWriter, r *http.Request) {
		pollCount++

		if pollCount == 1 {
			// First poll: not ready
			response := map[string]any{
				"status":  0,
				"request": "CAPCHA_NOT_READY",
			}
			json.NewEncoder(w).Encode(response)
		} else {
			// Second poll: ready with solution
			response := map[string]any{
				"status":  1,
				"request": solutionToken,
			}
			json.NewEncoder(w).Encode(response)
		}
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	solver := &TwoCaptcha{
		APIKey:  "test-key",
		BaseURL: server.URL,
		client:  server.Client(),
	}

	ctx := context.Background()
	result, err := solver.Solve(ctx, "test-site-key", "http://example.com")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != solutionToken {
		t.Errorf("expected solution %s, got %s", solutionToken, result)
	}

	if pollCount < 2 {
		t.Errorf("expected at least 2 polls, got %d", pollCount)
	}
}

// TestTwoCaptchaSolvePollingError verifies handling of polling errors.
func TestTwoCaptchaSolvePollingError(t *testing.T) {
	captchaID := "test-captcha-id"

	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status":  1,
			"request": captchaID,
		}
		json.NewEncoder(w).Encode(response)
	})

	handler.HandleFunc("/res.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status":  0,
			"request": "ERROR_CAPTCHA_UNSOLVABLE",
		}
		json.NewEncoder(w).Encode(response)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	solver := &TwoCaptcha{
		APIKey:  "test-key",
		BaseURL: server.URL,
		client:  server.Client(),
	}

	ctx := context.Background()
	_, err := solver.Solve(ctx, "test-site-key", "http://example.com")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "ERROR_CAPTCHA_UNSOLVABLE") {
		t.Errorf("expected error to contain ERROR_CAPTCHA_UNSOLVABLE, got %v", err)
	}
}

// TestTwoCaptchaSolveContextCancellation verifies context cancellation handling.
func TestTwoCaptchaSolveContextCancellation(t *testing.T) {
	captchaID := "test-captcha-id"

	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status":  1,
			"request": captchaID,
		}
		json.NewEncoder(w).Encode(response)
	})

	handler.HandleFunc("/res.php", func(w http.ResponseWriter, r *http.Request) {
		// Always return not ready to force polling
		response := map[string]any{
			"status":  0,
			"request": "CAPCHA_NOT_READY",
		}
		json.NewEncoder(w).Encode(response)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	solver := &TwoCaptcha{
		APIKey:  "test-key",
		BaseURL: server.URL,
		client:  server.Client(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := solver.Solve(ctx, "test-site-key", "http://example.com")

	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}

	if err != context.DeadlineExceeded {
		t.Errorf("expected context deadline exceeded, got %v", err)
	}
}

// TestTwoCaptchaSolveNetworkError verifies network error handling.
func TestTwoCaptchaSolveNetworkError(t *testing.T) {
	solver := &TwoCaptcha{
		APIKey:  "test-key",
		BaseURL: "http://nonexistent.example.com",
		client:  &http.Client{Timeout: 1 * time.Second},
	}

	ctx := context.Background()
	_, err := solver.Solve(ctx, "test-site-key", "http://example.com")

	if err == nil {
		t.Fatal("expected network error, got nil")
	}
}

// TestTwoCaptchaSolveInvalidJSON verifies handling of invalid JSON responses.
func TestTwoCaptchaSolveInvalidJSON(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("invalid json"))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	solver := &TwoCaptcha{
		APIKey:  "test-key",
		BaseURL: server.URL,
		client:  server.Client(),
	}

	ctx := context.Background()
	_, err := solver.Solve(ctx, "test-site-key", "http://example.com")

	if err == nil {
		t.Fatal("expected JSON decode error, got nil")
	}
}
