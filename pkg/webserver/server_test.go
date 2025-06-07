package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"subtitle-manager/pkg/auth"
	"subtitle-manager/pkg/database"
)

// TestHandler verifies that the handler serves index.html at root.
func TestHandler(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	// create test user and api key
	if err := auth.CreateUser(db, "test", "pass", "", "admin"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/", nil)
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("http get: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
}
