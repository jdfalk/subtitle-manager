package webserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
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

// TestConfigUpdate verifies that POST /api/config updates and persists values.
func TestConfigUpdate(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	akey, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("admin key: %v", err)
	}

	tmp := filepath.Join(t.TempDir(), "cfg.yaml")
	viper.SetConfigFile(tmp)
	viper.Set("test_key", "old")
	if err := viper.WriteConfig(); err != nil {
		t.Fatalf("write: %v", err)
	}
	defer viper.Reset()

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	body := strings.NewReader(`{"test_key":"new"}`)
	req, _ := http.NewRequest("POST", srv.URL+"/api/config", body)
	req.Header.Set("X-API-Key", akey)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status %d", resp.StatusCode)
	}

	if viper.GetString("test_key") != "new" {
		t.Fatalf("viper not updated")
	}
	data, err := os.ReadFile(tmp)
	if err != nil {
		t.Fatalf("read cfg: %v", err)
	}
	if !strings.Contains(string(data), "test_key: new") {
		t.Fatalf("config not written")
	}
}

// TestScanEndpoint verifies that POST /api/scan runs a directory scan.
func TestScanEndpoint(t *testing.T) {
	dir := t.TempDir()
	vid := filepath.Join(dir, "movie.mkv")
	if err := os.WriteFile(vid, []byte("x"), 0644); err != nil {
		t.Fatalf("create video: %v", err)
	}
	subSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "data")
	}))
	defer subSrv.Close()

	viper.Set("providers.generic.api_url", subSrv.URL)
	viper.Set("providers.generic.username", "")
	viper.Set("providers.generic.password", "")
	viper.Set("providers.generic.api_key", "")
	viper.Set("db_path", "")
	viper.Set("scan_workers", 1)
	defer viper.Reset()

	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create admin: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("api key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	body := strings.NewReader(fmt.Sprintf(`{"provider":"generic","dir":"%s","lang":"en"}`, dir))
	req, _ := http.NewRequest("POST", srv.URL+"/api/scan", body)
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status %d", resp.StatusCode)
	}
	b, err := os.ReadFile(filepath.Join(dir, "movie.en.srt"))
	if err != nil {
		t.Fatalf("subtitle not created: %v", err)
	}
	if string(b) != "data" {
		t.Fatalf("unexpected subtitle %q", b)
	}
}
