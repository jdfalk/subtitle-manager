package webserver

import (
	"errors"
	"io/fs"
	"net/http"
	"strings"
)

// spaFileServer serves static files from the given filesystem and falls back to
// index.html for unknown paths when it exists. This enables client-side routing
// for the React application.
func spaFileServer(fsys fs.FS) http.Handler {
	fsHandler := http.FileServer(http.FS(fsys))
	// Detect if index.html is present so we don't break during tests.
	_, err := fs.Stat(fsys, "index.html")
	hasIndex := err == nil

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hasIndex {
			path := strings.TrimPrefix(r.URL.Path, "/")
			if path != "" {
				if _, err := fs.Stat(fsys, path); err != nil && errors.Is(err, fs.ErrNotExist) {
					r = r.Clone(r.Context())
					r.URL.Path = "/index.html"
				}
			}
		}
		fsHandler.ServeHTTP(w, r)
	})
}
