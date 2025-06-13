package webserver

import (
	"encoding/json"
	"net/http"
	"runtime"

	"subtitle-manager/pkg/logging"
)

// logsHandler returns recent log lines captured in memory.
func logsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(logging.Hook.Logs())
	})
}

// systemHandler exposes basic system information.
func systemHandler() http.Handler {
	type info struct {
		GoVersion  string `json:"go_version"`
		OS         string `json:"os"`
		Arch       string `json:"arch"`
		Goroutines int    `json:"goroutines"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := info{
			GoVersion:  runtime.Version(),
			OS:         runtime.GOOS,
			Arch:       runtime.GOARCH,
			Goroutines: runtime.NumGoroutine(),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	})
}

// tasksHandler reports status for background tasks such as scanning.
func tasksHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scanMu.Lock()
		st := status
		scanMu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"scan": st})
	})
}
