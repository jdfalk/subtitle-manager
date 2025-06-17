// file: pkg/webserver/oauth_management_test.go
package webserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// TestGitHubOAuthGenerate verifies POST /api/oauth/github/generate creates new credentials.
func TestGitHubOAuthGenerate(t *testing.T) {
	db, err := database.Open(":memory:")
	testutil.MustNoError(t, "open db", err)
	defer db.Close()

	testutil.MustNoError(t, "create admin", auth.CreateUser(db, "admin", "p", "", "admin"))
	key := testutil.MustGet(t, "api key", func() (string, error) { return auth.GenerateAPIKey(db, 1) })

	tmp := filepath.Join(t.TempDir(), "cfg.yaml")
	viper.SetConfigFile(tmp)
	testutil.MustNoError(t, "write config", viper.WriteConfig())
	defer viper.Reset()

	h, err := Handler(db)
	testutil.MustNoError(t, "handler", err)
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("POST", srv.URL+"/api/oauth/github/generate", nil)
	req.Header.Set("X-API-Key", key)
	req.Header.Set("Origin", srv.URL)
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
	var creds struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RedirectURL  string `json:"redirect_url"`
	}
	json.NewDecoder(resp.Body).Decode(&creds)
	resp.Body.Close()

	if !strings.HasPrefix(creds.ClientID, "gh_") || !strings.HasPrefix(creds.ClientSecret, "ghs_") {
		t.Fatalf("unexpected credentials: %#v", creds)
	}
	if viper.GetString("github_client_id") != creds.ClientID {
		t.Fatalf("client id not stored")
	}
	data := testutil.MustGet(t, "read file", func() ([]byte, error) { return os.ReadFile(tmp) })
	if !strings.Contains(string(data), creds.ClientID) {
		t.Fatalf("config not written")
	}
}
