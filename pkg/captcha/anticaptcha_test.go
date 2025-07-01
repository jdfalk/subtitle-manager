// file: pkg/captcha/anticaptcha_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174009

package captcha

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTwoCaptcha(t *testing.T) {
	solver := NewTwoCaptcha("test-api-key")
	assert.NotNil(t, solver)
	assert.Equal(t, "test-api-key", solver.APIKey)
	assert.Equal(t, "https://2captcha.com", solver.BaseURL)
	assert.NotNil(t, solver.client)
}

func TestTwoCaptcha_Solve_Success(t *testing.T) {
	// Mock server that simulates successful captcha solving
	handler := http.NewServeMux()
	// Handle submit request
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		assert.Equal(t, "test-api-key", r.FormValue("key"))
		assert.Equal(t, "userrecaptcha", r.FormValue("method"))
		assert.Equal(t, "test-site-key", r.FormValue("googlekey"))
		assert.Equal(t, "https://example.com", r.FormValue("pageurl"))
		assert.Equal(t, "1", r.FormValue("json"))
		response := map[string]interface{}{
			"status":  1,
			"request": "123456789",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	// Handle result request
	handler.HandleFunc("/res.php", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		assert.Equal(t, "test-api-key", r.URL.Query().Get("key"))
		assert.Equal(t, "get", r.URL.Query().Get("action"))
		assert.Equal(t, "123456789", r.URL.Query().Get("id"))
		assert.Equal(t, "1", r.URL.Query().Get("json"))
		response := map[string]interface{}{
			"status":  1,
			"request": "solved-captcha-token",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = server.URL
	solver.client = server.Client()
	result, err := solver.Solve(context.Background(), "test-site-key", "https://example.com")
	require.NoError(t, err)
	assert.Equal(t, "solved-captcha-token", result)
}

func TestTwoCaptcha_Solve_SubmitError(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":  0,
			"request": "ERROR_ZERO_BALANCE",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = server.URL
	solver.client = server.Client()
	result, err := solver.Solve(context.Background(), "test-site-key", "https://example.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ERROR_ZERO_BALANCE")
	assert.Empty(t, result)
}

func TestTwoCaptcha_Solve_NotReady(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":  1,
			"request": "123456789",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	handler.HandleFunc("/res.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":  0,
			"request": "CAPCHA_NOT_READY",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = server.URL
	solver.client = server.Client()
	ctx, cancel := context.WithTimeout(context.Background(), 100*1000) // Short timeout for test
	defer cancel()
	result, err := solver.Solve(ctx, "test-site-key", "https://example.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
	assert.Empty(t, result)
}

func TestTwoCaptcha_Solve_ContextCanceled(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":  1,
			"request": "123456789",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = server.URL
	solver.client = server.Client()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result, err := solver.Solve(ctx, "test-site-key", "https://example.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context canceled")
	assert.Empty(t, result)
}

func TestTwoCaptcha_Solve_InvalidJSON(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("invalid json"))
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = server.URL
	solver.client = server.Client()
	result, err := solver.Solve(context.Background(), "test-site-key", "https://example.com")
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestTwoCaptcha_Solve_HTTPError(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = server.URL
	solver.client = server.Client()
	result, err := solver.Solve(context.Background(), "test-site-key", "https://example.com")
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestTwoCaptcha_SolverInterface(t *testing.T) {
	var solver Solver = NewTwoCaptcha("test-api-key")
	assert.NotNil(t, solver)
}

func TestTwoCaptcha_Solve_PollingHappyPath(t *testing.T) {
	// First poll returns not ready, second poll returns solution
	pollCount := 0
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":  1,
			"request": "test-captcha-id",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	handler.HandleFunc("/res.php", func(w http.ResponseWriter, r *http.Request) {
		pollCount++
		if pollCount == 1 {
			response := map[string]interface{}{
				"status":  0,
				"request": "CAPCHA_NOT_READY",
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)
		} else {
			response := map[string]interface{}{
				"status":  1,
				"request": "solution-token",
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)
		}
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = server.URL
	solver.client = server.Client()
	ctx := context.Background()
	result, err := solver.Solve(ctx, "test-site-key", "https://example.com")
	assert.NoError(t, err)
	assert.Equal(t, "solution-token", result)
	assert.GreaterOrEqual(t, pollCount, 2)
}

func TestTwoCaptcha_Solve_PollingError(t *testing.T) {
	// First poll returns not ready, second poll returns error
	pollCount := 0
	handler := http.NewServeMux()
	handler.HandleFunc("/in.php", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":  1,
			"request": "test-captcha-id",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	handler.HandleFunc("/res.php", func(w http.ResponseWriter, r *http.Request) {
		pollCount++
		if pollCount == 1 {
			response := map[string]interface{}{
				"status":  0,
				"request": "CAPCHA_NOT_READY",
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)
		} else {
			response := map[string]interface{}{
				"status":  0,
				"request": "ERROR_CAPTCHA_UNSOLVABLE",
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)
		}
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = server.URL
	solver.client = server.Client()
	ctx := context.Background()
	result, err := solver.Solve(ctx, "test-site-key", "https://example.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ERROR_CAPTCHA_UNSOLVABLE")
	assert.Empty(t, result)
	assert.GreaterOrEqual(t, pollCount, 2)
}

func TestTwoCaptcha_Solve_NetworkError(t *testing.T) {
	solver := NewTwoCaptcha("test-api-key")
	solver.BaseURL = "http://127.0.0.1:0" // Unreachable port
	solver.client = &http.Client{}
	ctx := context.Background()
	_, err := solver.Solve(ctx, "test-site-key", "https://example.com")
	assert.Error(t, err)
}
