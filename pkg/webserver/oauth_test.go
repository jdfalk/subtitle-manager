package webserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/oauth2"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// TestGitHubCallbackHandler verifies OAuth2 callback creates a session cookie.
func TestGitHubCallbackHandler(t *testing.T) {
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
