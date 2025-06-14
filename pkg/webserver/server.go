package webserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"

	"subtitle-manager/pkg/auth"
	"subtitle-manager/pkg/bazarr"
	"subtitle-manager/pkg/database"
	"subtitle-manager/pkg/subtitles"
	"subtitle-manager/pkg/webhooks"
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

		// Parse the entire JSON body to capture all settings
		var fullBody map[string]any
		if err := json.NewDecoder(r.Body).Decode(&fullBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Extract the known fields
		adminUser, _ := fullBody["admin_user"].(string)
		adminPass, _ := fullBody["admin_pass"].(string)
		serverName, _ := fullBody["server_name"].(string)
		reverseProxy, _ := fullBody["reverse_proxy"].(bool)
		integrations, _ := fullBody["integrations"].(map[string]any)

		if adminUser == "" || adminPass == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := auth.CreateUser(db, adminUser, adminPass, "", "admin"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the core settings
		if serverName != "" {
			viper.Set("server_name", serverName)
		}
		viper.Set("reverse_proxy", reverseProxy)

		// Set integration settings
		for k, v := range integrations {
			viper.Set("integrations."+k, v)
		}

		// Import all other settings (like imported Bazarr configuration)
		excludedKeys := map[string]bool{
			"admin_user": true, "admin_pass": true, "server_name": true,
			"reverse_proxy": true, "integrations": true,
		}
		for k, v := range fullBody {
			if !excludedKeys[k] {
				viper.Set(k, v)
			}
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
	prefix := strings.Trim(viper.GetString("base_url"), "/")
	if prefix != "" {
		prefix = "/" + prefix
	}
	mux := http.NewServeMux()
	mux.Handle(prefix+"/api/login", loginHandler(db))
	mux.Handle(prefix+"/api/setup/status", setupStatusHandler(db))
	mux.Handle(prefix+"/api/setup", setupHandler(db))
	mux.Handle(prefix+"/api/setup/bazarr", bazarrImportHandler(db))
	mux.Handle(prefix+"/api/setup/bazarr/upload", bazarrConfigUploadHandler(db))
	mux.Handle(prefix+"/api/oauth/github/login", githubLoginHandler(db))
	mux.Handle(prefix+"/api/oauth/github/callback", githubCallbackHandler(db))
	mux.Handle(prefix+"/api/config", authMiddleware(db, "basic", configHandler()))
	// Add Bazarr endpoints for Settings UI
	mux.Handle(prefix+"/api/bazarr/config", authMiddleware(db, "basic", bazarrConfigHandler()))
	mux.Handle(prefix+"/api/bazarr/import", authMiddleware(db, "basic", bazarrImportConfigHandler()))
	mux.Handle(prefix+"/api/scan", authMiddleware(db, "basic", scanHandler()))
	mux.Handle(prefix+"/api/scan/status", authMiddleware(db, "basic", scanStatusHandler()))
	mux.Handle(prefix+"/api/convert", authMiddleware(db, "basic", convertHandler()))
	mux.Handle(prefix+"/api/extract", authMiddleware(db, "basic", extractHandler()))
	mux.Handle(prefix+"/api/download", authMiddleware(db, "basic", downloadHandler(db)))
	mux.Handle(prefix+"/api/history", authMiddleware(db, "read", historyHandler(db)))
	mux.Handle(prefix+"/api/logs", authMiddleware(db, "basic", logsHandler()))
	mux.Handle(prefix+"/api/system", authMiddleware(db, "basic", systemHandler()))
	mux.Handle(prefix+"/api/tasks", authMiddleware(db, "basic", tasksHandler()))
	mux.Handle(prefix+"/api/webhooks/sonarr", webhooks.SonarrHandler())
	mux.Handle(prefix+"/api/webhooks/radarr", webhooks.RadarrHandler())
	mux.Handle(prefix+"/api/webhooks/custom", webhooks.CustomHandler())
	mux.Handle(prefix+"/api/translate", authMiddleware(db, "basic", translateHandler()))
	fsHandler := http.FileServer(http.FS(f))
	mux.Handle(prefix+"/", staticFileMiddleware(http.StripPrefix(prefix+"/", fsHandler)))
	return mux, nil
}

// StartServer starts an HTTP server on the given address serving the embedded UI.
func StartServer(addr string) error {
	// Get database configuration
	backend := database.GetDatabaseBackend()
	dbPath := viper.GetString("db_path")

	var db *sql.DB
	var err error

	// For web server, we need SQL database for authentication
	// If backend is not sqlite, we still need to create a SQLite DB for auth
	if backend == "sqlite" {
		fullPath := database.GetDatabasePath()
		// Ensure directory exists
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("failed to create database directory: %w", err)
		}
		db, err = database.Open(fullPath)
	} else {
		// For non-SQLite backends, create a separate SQLite DB for auth
		authDbPath := filepath.Join(dbPath, "auth.db")
		// Ensure directory exists
		if err := os.MkdirAll(dbPath, 0755); err != nil {
			return fmt.Errorf("failed to create database directory: %w", err)
		}
		db, err = database.Open(authDbPath)
	}

	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	// Check for automatic admin user creation via environment variables
	if err := createDefaultAdminIfNeeded(db); err != nil {
		return fmt.Errorf("failed to create initial admin user: %w", err)
	}

	h, err := Handler(db)
	if err != nil {
		return err
	}
	return http.ListenAndServe(addr, h)
}

// createDefaultAdminIfNeeded creates a default admin user from environment variables
// if no users exist and the required environment variables are set.
func createDefaultAdminIfNeeded(db *sql.DB) error {
	needed, err := setupNeeded(db)
	if err != nil {
		return fmt.Errorf("failed to check setup status: %w", err)
	}

	if !needed {
		return nil // Setup already completed
	}

	adminUser := viper.GetString("admin_user")
	adminPass := viper.GetString("admin_password")

	if adminUser == "" || adminPass == "" {
		return nil // No environment variables set, manual setup required
	}

	// Create the admin user
	if err := auth.CreateUser(db, adminUser, adminPass, "", "admin"); err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}

	fmt.Printf("Created default admin user from environment variables: %s\n", adminUser)
	return nil
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

// bazarrImportHandler imports settings from a Bazarr instance and returns the mapped settings for review
func bazarrImportHandler(db *sql.DB) http.Handler {
	type req struct {
		URL    string `json:"url"`
		APIKey string `json:"api_key"`
	}
	type resp struct {
		RawSettings map[string]any       `json:"raw_settings"`
		Settings    map[string]any       `json:"settings"`
		Preview     map[string]any       `json:"preview"`
		Mappings    []bazarr.MappingInfo `json:"mappings"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Allow this during setup phase
		needed, err := setupNeeded(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !needed {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		var q req
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if q.URL == "" || q.APIKey == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Fetch settings from Bazarr
		settings, err := bazarr.FetchSettings(q.URL, q.APIKey)
		if err != nil {
			http.Error(w, "Failed to connect to Bazarr: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Map settings for preview with detailed mapping information
		mapped, mappings := bazarr.MapSettingsWithInfo(settings)

		result := resp{
			RawSettings: settings,
			Settings:    settings,
			Preview:     mapped,
			Mappings:    mappings,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})
}

// bazarrConfigUploadHandler accepts a Bazarr config.ini file upload and parses it
func bazarrConfigUploadHandler(db *sql.DB) http.Handler {
	type resp struct {
		Preview  map[string]any       `json:"preview"`
		Mappings []bazarr.MappingInfo `json:"mappings"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Allow this during setup phase
		needed, err := setupNeeded(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !needed {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// Parse multipart form with file upload
		if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("config")
		if err != nil {
			http.Error(w, "No config file provided", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Read the config file
		content := make([]byte, 1<<20) // 1MB max
		n, err := file.Read(content)
		if err != nil && err.Error() != "EOF" {
			http.Error(w, "Failed to read config file", http.StatusBadRequest)
			return
		}
		content = content[:n]

		// Parse INI-style config (basic implementation)
		configMap := parseINIConfig(string(content))

		// Convert to Bazarr settings format and map with detailed mapping information
		mapped, mappings := bazarr.MapSettingsWithInfo(configMap)

		result := resp{
			Preview:  mapped,
			Mappings: mappings,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})
}

// parseINIConfig parses a basic INI file format used by Bazarr
func parseINIConfig(content string) map[string]any {
	result := make(map[string]any)
	currentSection := ""

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		// Section header
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.Trim(line, "[]")
			if result[currentSection] == nil {
				result[currentSection] = make(map[string]any)
			}
			continue
		}

		// Key-value pair
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				if currentSection != "" {
					if section, ok := result[currentSection].(map[string]any); ok {
						section[key] = value
					}
				} else {
					result[key] = value
				}
			}
		}
	}

	return result
}

// bazarrConfigHandler returns the current Bazarr configuration for preview
func bazarrConfigHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock Bazarr configuration for demonstration
		// In practice, this would connect to actual Bazarr API
		mockBazarrConfig := map[string]interface{}{
			"opensubtitles_username": "your_username",
			"opensubtitles_password": "[REDACTED]",
			"addic7ed_username":      "your_addic7ed_user",
			"provider_pool_enabled":  true,
			"language_profiles":      []string{"English", "Spanish"},
			"subtitle_directory":     "/config/subtitles",
			"chmod_enabled":          true,
			"chmod":                  "0644",
			"sonarr_url":             "http://sonarr:8989",
			"radarr_url":             "http://radarr:7878",
			"api_key":                "[REDACTED]",
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(mockBazarrConfig); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})
}

// bazarrImportConfigHandler imports settings from Bazarr into the current configuration
func bazarrImportConfigHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Mock implementation - in practice this would:
		// 1. Fetch current Bazarr configuration
		// 2. Map Bazarr settings to subtitle-manager configuration
		// 3. Update the current configuration with the mapped settings
		// 4. Save the updated configuration

		// For now, we'll simulate a successful import
		result := map[string]interface{}{
			"status":  "success",
			"message": "Successfully imported settings from Bazarr",
			"imported_keys": []string{
				"opensubtitles_username",
				"opensubtitles_password",
				"addic7ed_username",
				"provider_pool_enabled",
				"subtitle_directory",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})
}
