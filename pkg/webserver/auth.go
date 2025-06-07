package webserver

import (
	"database/sql"
	"net/http"

	"subtitle-manager/pkg/auth"
)

// authMiddleware verifies session cookies or API keys.
func authMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("session")
		if err == nil {
			if _, err := auth.ValidateSession(db, token.Value); err == nil {
				next.ServeHTTP(w, r)
				return
			}
		}
		key := r.Header.Get("X-API-Key")
		if key != "" {
			if _, err := auth.ValidateAPIKey(db, key); err == nil {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	})
}
