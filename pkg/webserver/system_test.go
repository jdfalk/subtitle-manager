package webserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

// TestSystemHandlers verifies /api/logs and /api/system endpoints.
func TestSystemHandlers(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	key, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("key: %v", err)
	}

	logger := logging.GetLogger("test")
	logger.Info("hello")

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("GET", srv.URL+"/api/logs", nil)
	req.Header.Set("X-API-Key", key)
	r, err := srv.Client().Do(req)
	if err != nil || r.StatusCode != http.StatusOK {
		t.Fatalf("logs: %v %d", err, r.StatusCode)
	}
	var lines []string
	json.NewDecoder(r.Body).Decode(&lines)
	r.Body.Close()
	if len(lines) == 0 {
		t.Fatalf("no logs returned")
	}

	req2, _ := http.NewRequest("GET", srv.URL+"/api/system", nil)
	req2.Header.Set("X-API-Key", key)
	r2, err := srv.Client().Do(req2)
	if err != nil || r2.StatusCode != http.StatusOK {
		t.Fatalf("system: %v %d", err, r2.StatusCode)
	}
	var info struct {
		GoVersion string `json:"go_version"`
	}
	json.NewDecoder(r2.Body).Decode(&info)
	r2.Body.Close()
	if info.GoVersion == "" {
		t.Fatalf("no info")
	}
}
