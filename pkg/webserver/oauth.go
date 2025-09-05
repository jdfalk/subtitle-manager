package webserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

	auth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
)

var githubAPIURL = "https://api.github.com/user"

// SetGitHubAPIURL overrides the GitHub API URL for testing purposes.
func SetGitHubAPIURL(u string) { githubAPIURL = u }

var oauthCfg *oauth2.Config

func initOAuthConfig() {
	oauthCfg = &oauth2.Config{
		ClientID:     viper.GetString("github_client_id"),
		ClientSecret: viper.GetString("github_client_secret"),
		RedirectURL:  viper.GetString("github_redirect_url"),
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email"},
	}
}

// githubLoginHandler starts the GitHub OAuth2 flow.
func githubLoginHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if oauthCfg == nil {
			initOAuthConfig()
		}
		state := uuid.NewString()
		url := oauthCfg.AuthCodeURL(state)
		http.Redirect(w, r, url, http.StatusFound)
	})
}

// githubCallbackHandler handles the OAuth2 callback and creates a session.
func githubCallbackHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if oauthCfg == nil {
			initOAuthConfig()
		}
		code := r.FormValue("code")
		if code == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token, err := oauthCfg.Exchange(context.Background(), code)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		client := oauthCfg.Client(context.Background(), token)
		resp, err := client.Get(githubAPIURL)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()
		var u struct {
			Email string `json:"email"`
			Login string `json:"login"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if u.Email == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		id, err := auth.GetOrCreateUser(db, u.Login, u.Email, "user")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		session, err := auth.GenerateSession(db, id, 24*time.Hour)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Extract session token from gcommon Session
		tokenStr := session.GetId()

		// Set secure session cookie with proper security attributes
		cookie := &http.Cookie{
			Name:     "session",
			Value:    tokenStr,
			Path:     "/",
			HttpOnly: true,                            // Prevents XSS attacks
			Secure:   r.TLS != nil,                    // HTTPS only in production
			SameSite: http.SameSiteStrictMode,         // CSRF protection
			MaxAge:   int((24 * time.Hour).Seconds()), // Explicit expiration
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
