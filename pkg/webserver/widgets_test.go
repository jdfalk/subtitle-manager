// file: pkg/webserver/widgets_test.go
package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	auth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// TestDashboardLayout verifies storing and retrieving layout preferences.
func TestDashboardLayout(t *testing.T) {
	skipIfNoSQLite(t)

	db := testutil.GetTestDB(t)
	defer db.Close()

	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	keyObj, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	layout := `{"widgets":[]}`
	body, _ := json.Marshal(map[string]string{"layout": layout})
	req, _ := http.NewRequest(http.MethodPost, srv.URL+"/api/widgets/layout", bytes.NewReader(body))
	req.Header.Set("X-API-Key", keyObj.GetId())
	req.Header.Set("Content-Type", "application/json")
	r1, err := srv.Client().Do(req)
	if err != nil || r1.StatusCode != http.StatusNoContent {
		t.Fatalf("post: %v %d", err, r1.StatusCode)
	}

	req2, _ := http.NewRequest(http.MethodGet, srv.URL+"/api/widgets/layout", nil)
	req2.Header.Set("X-API-Key", keyObj.GetId())
	r2, err := srv.Client().Do(req2)
	if err != nil || r2.StatusCode != http.StatusOK {
		t.Fatalf("get: %v %d", err, r2.StatusCode)
	}
	var out struct{ Layout string }
	if err := json.NewDecoder(r2.Body).Decode(&out); err != nil {
		t.Fatalf("decode: %v", err)
	}
	r2.Body.Close()
	if out.Layout != layout {
		t.Fatalf("expected %s got %s", layout, out.Layout)
	}
}

// TestDashboardLayout_InvalidMethodAndBody covers negative paths for layout handler.
func TestDashboardLayout_InvalidMethodAndBody(t *testing.T) {
	skipIfNoSQLite(t)

	db := testutil.GetTestDB(t)
	defer db.Close()

	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	keyObj, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	// Invalid method
	req, _ := http.NewRequest(http.MethodDelete, srv.URL+"/api/widgets/layout", nil)
	req.Header.Set("X-API-Key", keyObj.GetId())
	r1, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("delete: %v", err)
	}
	if r1.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", r1.StatusCode)
	}
	r1.Body.Close()

	// Invalid JSON body
	req2, _ := http.NewRequest(http.MethodPost, srv.URL+"/api/widgets/layout", bytes.NewReader([]byte("{")))
	req2.Header.Set("X-API-Key", keyObj.GetId())
	req2.Header.Set("Content-Type", "application/json")
	r2, err := srv.Client().Do(req2)
	if err != nil {
		t.Fatalf("post invalid: %v", err)
	}
	if r2.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", r2.StatusCode)
	}
	r2.Body.Close()

	// Missing layout field
	badBody, _ := json.Marshal(map[string]string{"layout": ""})
	req3, _ := http.NewRequest(http.MethodPost, srv.URL+"/api/widgets/layout", bytes.NewReader(badBody))
	req3.Header.Set("X-API-Key", keyObj.GetId())
	req3.Header.Set("Content-Type", "application/json")
	r3, err := srv.Client().Do(req3)
	if err != nil {
		t.Fatalf("post missing layout: %v", err)
	}
	if r3.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", r3.StatusCode)
	}
	r3.Body.Close()
}

// TestWidgetsList ensures the API returns available dashboard widgets.
func TestWidgetsList(t *testing.T) {
	skipIfNoSQLite(t)

	db := testutil.GetTestDB(t)
	defer db.Close()

	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	keyObj, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/api/widgets", nil)
	req.Header.Set("X-API-Key", keyObj.GetId())
	r, err := srv.Client().Do(req)
	if err != nil || r.StatusCode != http.StatusOK {
		t.Fatalf("get: %v %d", err, r.StatusCode)
	}
	var widgets []Widget
	if err := json.NewDecoder(r.Body).Decode(&widgets); err != nil {
		t.Fatalf("decode: %v", err)
	}
	r.Body.Close()
	if len(widgets) == 0 {
		t.Fatalf("expected widgets list")
	}
}

// TestWidgetsListMethodNotAllowed verifies the handler rejects unsupported methods.
func TestWidgetsListMethodNotAllowed(t *testing.T) {
	skipIfNoSQLite(t)

	db := testutil.GetTestDB(t)
	defer db.Close()

	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	keyObj, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("key: %v", err)
	}

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodPost, srv.URL+"/api/widgets", nil)
	req.Header.Set("X-API-Key", keyObj.GetId())
	r, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if r.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 got %d", r.StatusCode)
	}
	r.Body.Close()
}
