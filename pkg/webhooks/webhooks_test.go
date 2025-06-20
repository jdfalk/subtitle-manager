// file: pkg/webhooks/webhooks_test.go
package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestDispatcherSend verifies that Send posts the event and payload to all URLs.
func TestDispatcherSend(t *testing.T) {
	var got struct {
		Event   string            `json:"event"`
		Payload map[string]string `json:"payload"`
	}
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		json.NewDecoder(r.Body).Decode(&got)
	}))
	defer srv.Close()

	d := &Dispatcher{URLs: []string{srv.URL}, client: srv.Client()}
	if err := d.Send(context.Background(), "test", map[string]string{"a": "b"}); err != nil {
		t.Fatalf("send: %v", err)
	}
	if got.Event != "test" || got.Payload["a"] != "b" {
		t.Fatalf("unexpected payload: %#v", got)
	}
}

// TestNewInvalidURL verifies that New validates URLs and returns an error for HTTP.
func TestNewInvalidURL(t *testing.T) {
	if _, err := New([]string{"http://example.com"}); err == nil {
		t.Fatalf("expected error for http URL")
	}
}

// TestValidateWebhookURL checks validation rules for allowed and disallowed URLs.
func TestValidateWebhookURL(t *testing.T) {
	if err := validateWebhookURL("https://example.com"); err != nil {
		t.Fatalf("valid url rejected: %v", err)
	}
	if err := validateWebhookURL("http://example.com"); err == nil {
		t.Fatalf("http scheme should fail")
	}
	if err := validateWebhookURL("https://192.168.1.1"); err == nil {
		t.Fatalf("private address should fail")
	}
}

// TestIsPrivateOrLocalhost ensures detection of localhost and private addresses.
func TestIsPrivateOrLocalhost(t *testing.T) {
	cases := []struct {
		host string
		want bool
	}{
		{"localhost", true},
		{"127.0.0.1", true},
		{"192.168.0.2", true},
		{"example.com", false},
	}
	for _, c := range cases {
		c := c // capture range variable
		t.Run(c.host, func(t *testing.T) {
			if isPrivateOrLocalhost(c.host) != c.want {
				t.Fatalf("%s expected %v", c.host, c.want)
			}
		})
	}
}

// TestHandlersMethod verifies that non-POST methods are rejected.
func TestHandlersMethod(t *testing.T) {
	handlers := []struct {
		name string
		h    http.Handler
	}{
		{"sonarr", SonarrHandler()},
		{"radarr", RadarrHandler()},
		{"custom", CustomHandler()},
	}
	for _, tc := range handlers {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			tc.h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/", nil))
			if w.Code != http.StatusMethodNotAllowed {
				t.Fatalf("%s: expected 405, got %d", tc.name, w.Code)
			}
		})
	}
}

// TestHandlersBadBody ensures invalid JSON payloads return 400.
func TestHandlersBadBody(t *testing.T) {
	handlers := []http.Handler{SonarrHandler(), RadarrHandler(), CustomHandler()}
	for i, h := range handlers {
		h := h // capture range variable
		t.Run(fmt.Sprintf("handler_%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			h.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{")))
			if w.Code != http.StatusBadRequest {
				t.Fatalf("expected 400, got %d", w.Code)
			}
		})
	}
}

// TestHandleInvalidProvider checks that unknown providers return 400.
func TestHandleInvalidProvider(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	handle(w, r, event{Path: "file", Lang: "en", Provider: "unknown"})
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

// TestHandleInvalidLang verifies that scanner errors result in a 500 response.
func TestHandleInvalidLang(t *testing.T) {
	dir := t.TempDir()
	file := dir + "/video.mkv"
	if err := os.WriteFile(file, []byte("x"), 0644); err != nil {
		t.Fatalf("write video: %v", err)
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	handle(w, r, event{Path: file, Lang: "??"})
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
