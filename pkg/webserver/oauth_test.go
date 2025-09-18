package webserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"golang.org/x/oauth2"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	auth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
)

// TestGitHubCallbackHandler verifies OAuth2 callback creates a session cookie.
func TestGitHubCallbackHandler(t *testing.T) {
	skipIfNoSQLite(t)

	tokenResp := `{"access_token":"t","token_type":"bearer"}`
	apiResp := `{"login":"tester","email":"tester@example.com"}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/token":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, tokenResp)
		case "/user":
			fmt.Fprint(w, apiResp)
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	oauthCfg = &oauth2.Config{
		ClientID:     "id",
		ClientSecret: "sec",
		RedirectURL:  "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  srv.URL + "/auth",
			TokenURL: srv.URL + "/token",
		},
	}
	SetGitHubAPIURL(srv.URL + "/user")
	defer func() { oauthCfg = nil; SetGitHubAPIURL("https://api.github.com/user") }()

	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	h := githubCallbackHandler(db)
	r := httptest.NewRequest("GET", "/callback?code=c", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusFound {
		t.Fatalf("status %d", w.Code)
	}
	if len(w.Result().Cookies()) == 0 {
		t.Fatalf("session cookie missing")
	}
}

// TestGitHubOAuthReset verifies that the reset endpoint clears configuration.
func TestGitHubOAuthReset(t *testing.T) {
	skipIfNoSQLite(t)

	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	if err := auth.CreateUser(db, "admin", "p", "", "admin"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	keyObj, err := auth.GenerateAPIKey(db, 1)
	if err != nil {
		t.Fatalf("api key: %v", err)
	}

	tmp := filepath.Join(t.TempDir(), "cfg.yaml")
	viper.SetConfigFile(tmp)
	viper.Set("github_client_id", "id")
	viper.Set("github_client_secret", "sec")
	viper.Set("github_redirect_url", "url")
	viper.Set("github_oauth_enabled", true)
	if err := viper.WriteConfig(); err != nil {
		t.Fatalf("write config: %v", err)
	}
	defer viper.Reset()

	h, err := Handler(db)
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	srv := httptest.NewServer(h)
	defer srv.Close()

	req, _ := http.NewRequest("POST", srv.URL+"/api/oauth/github/reset", nil)
	req.Header.Set("X-API-Key", keyObj.GetId())
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
	resp.Body.Close()

	if viper.GetString("github_client_id") != "" ||
		viper.GetString("github_client_secret") != "" ||
		viper.GetString("github_redirect_url") != "" ||
		viper.GetBool("github_oauth_enabled") {
		t.Fatalf("config not reset")
	}
}
