// file: pkg/webserver/oauth_management.go
package webserver

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

// oauthCredentials represents OAuth2 credentials returned to clients.
type oauthCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url,omitempty"`
}

// generateSecureToken returns a cryptographically secure random hex string.
func generateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// githubOAuthGenerateHandler handles POST /api/oauth/github/generate requests.
// It generates new GitHub OAuth2 credentials and persists them via Viper.
func githubOAuthGenerateHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		id, err := generateSecureToken(16)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		secret, err := generateSecureToken(32)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		creds := oauthCredentials{
			ClientID:     "gh_" + id,
			ClientSecret: "ghs_" + secret,
			RedirectURL:  r.Header.Get("Origin") + "/api/oauth/github/callback",
		}

		viper.Set("github_client_id", creds.ClientID)
		viper.Set("github_client_secret", creds.ClientSecret)
		viper.Set("github_redirect_url", creds.RedirectURL)
		if cfg := viper.ConfigFileUsed(); cfg != "" {
			if err := viper.WriteConfig(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(creds)
	})
}

// githubOAuthResetHandler resets GitHub OAuth configuration values.
//
// POST /api/oauth/github/reset
func githubOAuthResetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Reset configuration fields
		viper.Set("github_client_id", "")
		viper.Set("github_client_secret", "")
		viper.Set("github_redirect_url", "")
		viper.Set("github_oauth_enabled", false)

		if cfg := viper.ConfigFileUsed(); cfg != "" {
			if err := viper.WriteConfig(); err != nil {
				http.Error(w, "Failed to save configuration", http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	})
}
