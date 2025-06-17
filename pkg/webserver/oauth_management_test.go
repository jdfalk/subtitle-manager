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

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
	"github.com/spf13/viper"
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
	var creds OAuthCredentials
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

// TestGitHubOAuthRegenerate verifies client secret regeneration and persistence.
func TestGitHubOAuthRegenerate(t *testing.T) {
	db, err := database.Open(":memory:")
	testutil.MustNoError(t, "open db", err)
	defer db.Close()

	testutil.MustNoError(t, "create admin", auth.CreateUser(db, "admin", "p", "", "admin"))
	key := testutil.MustGet(t, "api key", func() (string, error) { return auth.GenerateAPIKey(db, 1) })

	tmp := filepath.Join(t.TempDir(), "cfg.yaml")
	viper.SetConfigFile(tmp)
	viper.Set("github_client_id", "gh_id")
	viper.Set("github_client_secret", "ghs_old")
	viper.Set("github_redirect_url", "http://cb")
	testutil.MustNoError(t, "write config", viper.WriteConfig())
	defer viper.Reset()

	h, err := Handler(db)
	testutil.MustNoError(t, "handler", err)
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("POST", srv.URL+"/api/oauth/github/regenerate", nil)
	req.Header.Set("X-API-Key", key)
	resp := testutil.MustGet(t, "request", func() (*http.Response, error) { return srv.Client().Do(req) })
	testutil.MustEqual(t, "status", http.StatusOK, resp.StatusCode)

	var creds OAuthCredentials
	_ = json.NewDecoder(resp.Body).Decode(&creds)
	testutil.MustEqual(t, "client id", "gh_id", creds.ClientID)
	testutil.MustEqual(t, "redirect", "http://cb", creds.RedirectURL)
	if !strings.HasPrefix(creds.ClientSecret, "ghs_") {
		t.Fatalf("secret prefix")
	}
	testutil.MustNotEqual(t, "secret changed", "ghs_old", creds.ClientSecret)
	testutil.MustEqual(t, "viper", creds.ClientSecret, viper.GetString("github_client_secret"))
	data := testutil.MustGet(t, "read file", func() ([]byte, error) { return os.ReadFile(tmp) })
	testutil.MustContain(t, "persist", string(data), creds.ClientSecret)
}
