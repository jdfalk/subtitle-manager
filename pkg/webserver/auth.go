package webserver

import (
	"context"
	"database/sql"
	"net/http"

	"subtitle-manager/pkg/auth"
)

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
		ctx := context.WithValue(r.Context(), "userID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// setupAwareMiddleware conditionally applies authentication based on setup status.
// If setup is needed, it bypasses authentication to allow access to the setup UI.
// Otherwise, it applies normal authentication.
func setupAwareMiddleware(db *sql.DB, perm string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if setup is needed
		needed, err := setupNeeded(db)
		if err != nil {
			// If we can't determine setup status, err on the side of caution
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// If setup is needed, bypass authentication
		if needed {
			next.ServeHTTP(w, r)
			return
		}

		// Otherwise, apply normal authentication
		authMiddleware(db, perm, next).ServeHTTP(w, r)
	})
}
