package webserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/spf13/viper"
)

// TestDatabaseHandlers verifies the new database related endpoints.
func TestDatabaseHandlers(t *testing.T) {
	skipIfNoSQLite(t)

	dir := t.TempDir()
	dbPath := filepath.Join(dir, "app.db")
	db, err := database.Open(dbPath)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	viper.Set("db_backend", "sqlite")
	viper.Set("db_path", dir)
	viper.Set("sqlite3_filename", "app.db")
	defer viper.Reset()

	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("GET", srv.URL+"/api/database/info", nil)
	req.Header.Set("X-API-Key", key)
	r1, err := srv.Client().Do(req)
	if err != nil || r1.StatusCode != http.StatusOK {
		t.Fatalf("info: %v %d", err, r1.StatusCode)
	}
	var info struct{ Type string }
	if err := json.NewDecoder(r1.Body).Decode(&info); err != nil {
		t.Fatalf("failed to decode request body: %v", err)
	}
	r1.Body.Close()
	if info.Type != "sqlite" {
		t.Fatalf("unexpected type %s", info.Type)
	}

	req2, _ := http.NewRequest("GET", srv.URL+"/api/database/stats", nil)
	req2.Header.Set("X-API-Key", key)
	r2, err := srv.Client().Do(req2)
	if err != nil || r2.StatusCode != http.StatusOK {
		t.Fatalf("stats: %v %d", err, r2.StatusCode)
	}
	r2.Body.Close()

	req3, _ := http.NewRequest("POST", srv.URL+"/api/database/backup", nil)
	req3.Header.Set("X-API-Key", key)
	r3, err := srv.Client().Do(req3)
	if err != nil || r3.StatusCode != http.StatusOK {
		t.Fatalf("backup: %v %d", err, r3.StatusCode)
	}
	if ct := r3.Header.Get("Content-Type"); ct != "application/gzip" {
		t.Fatalf("content type %s", ct)
	}
	buf := make([]byte, 1)
	if n, _ := r3.Body.Read(buf); n == 0 {
		t.Fatalf("empty backup")
	}
	r3.Body.Close()

	// simulate missing database file to trigger error path
	viper.Set("sqlite3_filename", "missing.db")
	req4, _ := http.NewRequest("POST", srv.URL+"/api/database/backup", nil)
	req4.Header.Set("X-API-Key", key)
	r4, err := srv.Client().Do(req4)
	if err != nil {
		t.Fatalf("backup err: %v", err)
	}
	if r4.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500 got %d", r4.StatusCode)
	}
	r4.Body.Close()
}
