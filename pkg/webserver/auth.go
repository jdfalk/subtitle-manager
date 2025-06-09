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
