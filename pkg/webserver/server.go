package webserver

import (
	"io/fs"
	"net/http"

	"subtitle-manager/webui"
)

// Handler returns an http.Handler that serves the embedded web UI.
func Handler() (http.Handler, error) {
	f, err := fs.Sub(webui.FS, "dist")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(f)), nil
}

// StartServer starts an HTTP server on the given address serving the embedded UI.
func StartServer(addr string) error {
	h, err := Handler()
	if err != nil {
		return err
	}
	return http.ListenAndServe(addr, h)
}
