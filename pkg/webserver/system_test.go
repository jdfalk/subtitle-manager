package webserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	auth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// TestSystemHandlers verifies /api/logs and /api/system endpoints.
func TestSystemHandlers(t *testing.T) {
	skipIfNoSQLite(t)

	db := testutil.GetTestDB(t)
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
	if err := json.NewDecoder(r.Body).Decode(&lines); err != nil {
		t.Fatalf("failed to decode request body: %v", err)
	}
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
	if err := json.NewDecoder(r2.Body).Decode(&info); err != nil {
		t.Fatalf("failed to decode request body: %v", err)
	}
	r2.Body.Close()
	if info.GoVersion == "" {
		t.Fatalf("no info")
	}

	req3, _ := http.NewRequest("GET", srv.URL+"/api/announcements", nil)
	req3.Header.Set("X-API-Key", key)
	r3, err := srv.Client().Do(req3)
	if err != nil || r3.StatusCode != http.StatusOK {
		t.Fatalf("announcements: %v %d", err, r3.StatusCode)
	}
	var ann []map[string]any
	if err := json.NewDecoder(r3.Body).Decode(&ann); err != nil {
		t.Fatalf("failed to decode request body: %v", err)
	}
	r3.Body.Close()
	if len(ann) == 0 {
		t.Fatalf("no announcements")
	}
}
