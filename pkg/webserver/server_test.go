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

// TestRBAC verifies that permissions are enforced for protected routes.
func TestRBAC(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	// viewer role should not access /api/config
	if err := auth.CreateUser(db, "viewer", "p", "", "viewer"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	vkey, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("viewer key: %v", err)
	}
	// admin role can access
	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	akey, err := auth.GenerateAPIKey(db, 2)
	if err != nil {
		t.Fatalf("admin key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("GET", srv.URL+"/api/config", nil)
	req.Header.Set("X-API-Key", vkey)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("http get: %v", err)
	}
	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("viewer status %d", resp.StatusCode)
	}

	req2, _ := http.NewRequest("GET", srv.URL+"/api/config", nil)
	req2.Header.Set("X-API-Key", akey)
	resp2, err := srv.Client().Do(req2)
	if err != nil {
		t.Fatalf("admin get: %v", err)
	}
	if resp2.StatusCode != http.StatusOK {
		t.Fatalf("admin status %d", resp2.StatusCode)
	}
}
