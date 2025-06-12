package webserver

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"

	"subtitle-manager/pkg/auth"
	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/subtitles"
	"subtitle-manager/webui"
)

// setupNeeded returns true when no user accounts exist.
func setupNeeded(db *sql.DB) (bool, error) {
	var c int
	row := db.QueryRow(`SELECT COUNT(1) FROM users`)
	if err := row.Scan(&c); err != nil {
		return false, err
	}
	return c == 0, nil
}

// setupStatusHandler reports whether initial setup is required.
func setupStatusHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		needed, err := setupNeeded(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]bool{"needed": needed})
	})
}

// setupHandler performs the initial configuration and creates the admin user.
func setupHandler(db *sql.DB) http.Handler {
	type req struct {
		ServerName   string         `json:"server_name"`
		ReverseProxy bool           `json:"reverse_proxy"`
		AdminUser    string         `json:"admin_user"`
		AdminPass    string         `json:"admin_pass"`
		Integrations map[string]any `json:"integrations"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		needed, err := setupNeeded(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !needed {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var q req
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if q.AdminUser == "" || q.AdminPass == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := auth.CreateUser(db, q.AdminUser, q.AdminPass, "", "admin"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if q.ServerName != "" {
			viper.Set("server_name", q.ServerName)
		}
		viper.Set("reverse_proxy", q.ReverseProxy)
		for k, v := range q.Integrations {
			viper.Set("integrations."+k, v)
		}
		if cfg := viper.ConfigFileUsed(); cfg != "" {
			if err := viper.WriteConfig(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

// Handler returns an http.Handler that serves the embedded web UI.
func Handler(db *sql.DB) (http.Handler, error) {
	f, err := fs.Sub(webui.FS, "dist")
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.Handle("/api/login", loginHandler(db))
	mux.Handle("/api/setup/status", setupStatusHandler(db))
	mux.Handle("/api/setup", setupHandler(db))
	mux.Handle("/api/oauth/github/login", githubLoginHandler(db))
	mux.Handle("/api/oauth/github/callback", githubCallbackHandler(db))
	mux.Handle("/api/config", authMiddleware(db, "basic", configHandler()))
	mux.Handle("/api/scan", authMiddleware(db, "basic", scanHandler()))
	mux.Handle("/api/scan/status", authMiddleware(db, "basic", scanStatusHandler()))
	mux.Handle("/api/convert", authMiddleware(db, "basic", convertHandler()))
	mux.Handle("/api/extract", authMiddleware(db, "basic", extractHandler()))
	mux.Handle("/api/history", authMiddleware(db, "read", historyHandler(db)))
	mux.Handle("/api/logs", authMiddleware(db, "basic", logsHandler()))
	mux.Handle("/api/system", authMiddleware(db, "basic", systemHandler()))
	mux.Handle("/api/tasks", authMiddleware(db, "basic", tasksHandler()))
	mux.Handle("/api/translate", authMiddleware(db, "basic", translateHandler()))
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

// convertHandler converts an uploaded subtitle file to SRT format.
//
// POST requests must contain a multipart form with the file field named "file".
// The converted subtitle bytes are returned in the response body.
func convertHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		f, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer f.Close()
		tmp, err := os.CreateTemp("", "upload-*.srt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer os.Remove(tmp.Name())
		if _, err := io.Copy(tmp, f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tmp.Close()
		data, err := subtitles.ConvertToSRT(tmp.Name())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(data)
	})
}
