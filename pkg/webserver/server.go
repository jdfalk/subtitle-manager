package webserver

import (
	"database/sql"
	"encoding/json"
	"io/fs"
	"net/http"
	"time"

	"github.com/spf13/viper"

	"subtitle-manager/pkg/auth"
	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/providers"
	"subtitle-manager/pkg/scanner"
	"subtitle-manager/pkg/subtitles"
	"subtitle-manager/webui"
)

// Handler returns an http.Handler that serves the embedded web UI.
func Handler(db *sql.DB) (http.Handler, error) {
	f, err := fs.Sub(webui.FS, "dist")
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.Handle("/api/login", loginHandler(db))
	mux.Handle("/api/oauth/github/login", githubLoginHandler(db))
	mux.Handle("/api/oauth/github/callback", githubCallbackHandler(db))
	mux.Handle("/api/config", authMiddleware(db, "basic", configHandler()))
	mux.Handle("/api/scan", authMiddleware(db, "basic", scanHandler()))
	mux.Handle("/api/scan/status", authMiddleware(db, "basic", scanStatusHandler()))
	mux.Handle("/api/extract", authMiddleware(db, "basic", extractHandler()))
	mux.Handle("/api/scan", authMiddleware(db, "basic", scanHandler()))
	fsHandler := http.FileServer(http.FS(f))
	mux.Handle("/", authMiddleware(db, "read", fsHandler))
	return mux, nil
}

// StartServer starts an HTTP server on the given address serving the embedded UI.
func StartServer(addr string) error {
	db, err := database.Open(viper.GetString("db_path"))
	if err != nil {
		return err
	}
	defer db.Close()
	h, err := Handler(db)
	if err != nil {
		return err
	}
	return http.ListenAndServe(addr, h)
}

// loginHandler authenticates a user and sets a session cookie.
func loginHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		id, err := auth.AuthenticateUser(db, username, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, err := auth.GenerateSession(db, id, 24*time.Hour)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "session", Value: token, Path: "/", HttpOnly: true})
	})
}

// configHandler returns the complete configuration as JSON.
// configHandler handles reading and updating configuration values.
//
// GET requests return the current configuration as JSON. POST requests accept a
// JSON body of key/value pairs which are written into Viper and persisted to the
// config file if one is in use.
func configHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(viper.AllSettings())
		case http.MethodPost:
			var vals map[string]any
			if err := json.NewDecoder(r.Body).Decode(&vals); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			for k, v := range vals {
				viper.Set(k, v)
			}
			if cfg := viper.ConfigFileUsed(); cfg != "" {
				if err := viper.WriteConfig(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// extractHandler extracts subtitles from the provided media file.
//
// POST requests expect a JSON body `{"path":"/file.mkv"}`. The subtitle items
// extracted by subtitles.ExtractFromMedia are returned as JSON.
func extractHandler() http.Handler {
	type req struct {
		Path string `json:"path"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var q req
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil || q.Path == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		items, err := subtitles.ExtractFromMedia(q.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(items)
	})
}

// scanHandler runs a library scan using the scanner package.
// It expects a JSON body containing provider, dir, lang and optional upgrade.
// On success the scan is executed synchronously and HTTP 204 is returned.
func scanHandler() http.Handler {
	type scanRequest struct {
		Provider string `json:"provider"`
		Dir      string `json:"dir"`
		Lang     string `json:"lang"`
		Upgrade  bool   `json:"upgrade"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req scanRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p, err := providers.Get(req.Provider, viper.GetString("opensubtitles.api_key"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		workers := viper.GetInt("scan_workers")
		var store database.SubtitleStore
		if dbPath := viper.GetString("db_path"); dbPath != "" {
			backend := viper.GetString("db_backend")
			if s, err := database.OpenStore(dbPath, backend); err == nil {
				store = s
				defer s.Close()
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		if err := scanner.ScanDirectory(r.Context(), req.Dir, req.Lang, req.Provider, p, req.Upgrade, workers, store); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
