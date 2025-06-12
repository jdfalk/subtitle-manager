package webserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"subtitle-manager/pkg/auth"
	"subtitle-manager/pkg/database"

	"github.com/spf13/viper"
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

// TestScanHandlers verifies /api/scan and /api/scan/status.
func TestScanHandlers(t *testing.T) {
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
	// fake subtitle server
	subSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("sub"))
	}))
	defer subSrv.Close()
	viper.Set("providers.generic.api_url", subSrv.URL)
	defer viper.Reset()
	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()
	// create video file
	dir := t.TempDir()
	vid := filepath.Join(dir, "file.mkv")
	os.WriteFile(vid, []byte("x"), 0644)
	// trigger scan
	body := strings.NewReader(`{"provider":"generic","directory":"` + dir + `","lang":"en"}`)
	req, _ := http.NewRequest("POST", srv.URL+"/api/scan", body)
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil || resp.StatusCode != http.StatusAccepted {
		t.Fatalf("scan start: %v %d", err, resp.StatusCode)
	}
	// poll status until not running
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		req2, _ := http.NewRequest("GET", srv.URL+"/api/scan/status", nil)
		req2.Header.Set("X-API-Key", key)
		r2, err := srv.Client().Do(req2)
		if err != nil {
			t.Fatalf("status: %v", err)
		}
		var s struct {
			Running   bool `json:"running"`
			Completed int  `json:"completed"`
		}
		json.NewDecoder(r2.Body).Decode(&s)
		r2.Body.Close()
		if !s.Running {
			if s.Completed != 1 {
				t.Fatalf("completed %d", s.Completed)
			}
			return
		}
	}
	t.Fatalf("scan did not finish")
}

// TestExtract verifies that POST /api/extract returns subtitle items.
func TestExtract(t *testing.T) {
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

	dir := t.TempDir()
	script := filepath.Join(dir, "ffmpeg")
	data := "#!/bin/sh\ncp ../../testdata/simple.srt \"$6\"\n"
	if err := os.WriteFile(script, []byte(data), 0755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", dir+":"+oldPath)
	defer os.Setenv("PATH", oldPath)

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	body := strings.NewReader(`{"path":"dummy.mkv"}`)
	req, _ := http.NewRequest("POST", srv.URL+"/api/extract", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", key)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	var items []any
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(items) == 0 {
		t.Fatalf("no items returned")
	}
}

// TestSetup verifies the initial setup workflow.
func TestSetup(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	// status should report setup needed
	resp, err := http.Get(srv.URL + "/api/setup/status")
	if err != nil {
		t.Fatalf("status: %v", err)
	}
	var st struct {
		Needed bool `json:"needed"`
	}
	json.NewDecoder(resp.Body).Decode(&st)
	resp.Body.Close()
	if !st.Needed {
		t.Fatalf("expected setup needed")
	}

	body := strings.NewReader(`{"server_name":"Test","admin_user":"a","admin_pass":"p"}`)
	resp2, err := http.Post(srv.URL+"/api/setup", "application/json", body)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp2.StatusCode != http.StatusNoContent {
		t.Fatalf("status %d", resp2.StatusCode)
	}

	// now status should be false
	resp3, err := http.Get(srv.URL + "/api/setup/status")
	if err != nil {
		t.Fatalf("status2: %v", err)
	}
	json.NewDecoder(resp3.Body).Decode(&st)
	resp3.Body.Close()
	if st.Needed {
		t.Fatalf("setup still needed")
	}
}
