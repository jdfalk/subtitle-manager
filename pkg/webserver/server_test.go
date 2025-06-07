package webserver

import (
	"net/http/httptest"
	"testing"
)

// TestHandler verifies that the handler serves index.html at root.
func TestHandler(t *testing.T) {
	h, err := Handler()
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()
	resp, err := srv.Client().Get(srv.URL + "/")
	if err != nil {
		t.Fatalf("http get: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}
