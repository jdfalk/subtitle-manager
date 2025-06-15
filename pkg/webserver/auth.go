package webserver

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
)

// userIDKey is a custom type for context keys to avoid collisions
type userIDKey string

const userIDContextKey userIDKey = "userID"

// authMiddleware verifies session cookies or API keys.
func authMiddleware(db *sql.DB, perm string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id int64
		if t, err := r.Cookie("session"); err == nil {
			if uid, err := auth.ValidateSession(db, t.Value); err == nil {
				id = uid
			}
		}
		if id == 0 {
			if key := r.Header.Get("X-API-Key"); key != "" {
				if uid, err := auth.ValidateAPIKey(db, key); err == nil {
					id = uid
				}
			}
		}
		if id == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ok, err := auth.CheckPermission(db, id, perm)
		if err != nil || !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), userIDContextKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// staticFileMiddleware allows unrestricted access to static files (HTML, CSS, JS, etc.)
// This ensures the login page and app resources are always accessible.
func staticFileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always allow access to static files - no authentication required
		// This enables users to load the login form and app resources
		next.ServeHTTP(w, r)
	})
}
