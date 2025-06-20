package captcha

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestSolveImage verifies that the client solves image captchas
// using the Anti-Captcha API endpoints.
func TestSolveImage(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/createTask", func(w http.ResponseWriter, r *http.Request) {
		var req map[string]any
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		json.NewEncoder(w).Encode(map[string]any{"errorId": 0, "taskId": 1})
	})
	handler.HandleFunc("/getTaskResult", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"errorId":  0,
			"status":   "ready",
			"solution": map[string]string{"text": "abc"},
		})
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()

	SetAPIURL(srv.URL)
	defer SetAPIURL("https://api.anti-captcha.com")

	c := New("key")
	c.HTTPClient = srv.Client()
	ans, err := c.SolveImage(context.Background(), "base64")
	if err != nil {
		t.Fatalf("solve: %v", err)
	}
	if ans != "abc" {
		t.Fatalf("expected abc, got %s", ans)
	}
}

// TestSolveRecaptchaV2 verifies solving a recaptcha returns the token.
func TestSolveRecaptchaV2(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/createTask", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"errorId": 0, "taskId": 2})
	})
	handler.HandleFunc("/getTaskResult", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"errorId":  0,
			"status":   "ready",
			"solution": map[string]string{"gRecaptchaResponse": "token"},
		})
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()

	SetAPIURL(srv.URL)
	defer SetAPIURL("https://api.anti-captcha.com")

	c := New("key")
	c.HTTPClient = srv.Client()
	tok, err := c.SolveRecaptchaV2(context.Background(), "http://x", "k")
	if err != nil {
		t.Fatalf("solve recaptcha: %v", err)
	}
	if tok != "token" {
		t.Fatalf("expected token, got %s", tok)
	}
}
