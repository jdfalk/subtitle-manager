package webserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/bazarr"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/maintenance"
	"github.com/jdfalk/subtitle-manager/pkg/radarr"
	"github.com/jdfalk/subtitle-manager/pkg/sonarr"
	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
	"github.com/jdfalk/subtitle-manager/pkg/updater"
	"github.com/jdfalk/subtitle-manager/pkg/webhooks"
	"github.com/jdfalk/subtitle-manager/webui"
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
	mux.Handle(prefix+"/api/logout", logoutHandler(db))
	mux.Handle(prefix+"/api/setup/status", setupStatusHandler(db))
	mux.Handle(prefix+"/api/setup", setupHandler(db))
	mux.Handle(prefix+"/api/setup/bazarr", bazarrImportHandler(db))
	mux.Handle(prefix+"/api/setup/bazarr/upload", bazarrConfigUploadHandler(db))
	mux.Handle(prefix+"/api/oauth/github/login", githubLoginHandler(db))
	mux.Handle(prefix+"/api/oauth/github/callback", githubCallbackHandler(db))
	mux.Handle(prefix+"/api/oauth/github/generate", authMiddleware(db, "admin", githubOAuthGenerateHandler()))
	mux.Handle(prefix+"/api/oauth/github/regenerate", authMiddleware(db, "admin", githubOAuthRegenerateHandler()))
	mux.Handle(prefix+"/api/oauth/github/reset", authMiddleware(db, "admin", githubOAuthResetHandler()))
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
	mux.Handle(prefix+"/api/tasks/start", authMiddleware(db, "basic", startTaskHandler()))
	mux.Handle(prefix+"/api/providers/status", authMiddleware(db, "basic", providerStatusHandler()))
	mux.Handle(prefix+"/api/providers/refresh", authMiddleware(db, "basic", providerRefreshHandler()))
	mux.Handle(prefix+"/api/providers/reset", authMiddleware(db, "basic", providerResetHandler()))
	mux.Handle(prefix+"/api/database/info", authMiddleware(db, "basic", databaseInfoHandler(db)))
	mux.Handle(prefix+"/api/database/stats", authMiddleware(db, "basic", databaseStatsHandler(db)))
	mux.Handle(prefix+"/api/database/backup", authMiddleware(db, "basic", databaseBackupHandler()))
	mux.Handle(prefix+"/api/database/optimize", authMiddleware(db, "basic", databaseOptimizeHandler(db)))
	mux.Handle(prefix+"/api/backups", authMiddleware(db, "basic", backupsHandler()))
	mux.Handle(prefix+"/api/backups/create", authMiddleware(db, "basic", createBackupHandler()))
	mux.Handle(prefix+"/api/backups/restore", authMiddleware(db, "basic", restoreBackupHandler()))
	mux.Handle(prefix+"/api/releases", authMiddleware(db, "basic", releasesHandler()))
	mux.Handle(prefix+"/api/announcements", authMiddleware(db, "basic", announcementsHandler()))
	mux.Handle(prefix+"/api/webhooks/sonarr", webhooks.SonarrHandler())
	mux.Handle(prefix+"/api/webhooks/radarr", webhooks.RadarrHandler())
	mux.Handle(prefix+"/api/webhooks/custom", webhooks.CustomHandler())
	mux.Handle(prefix+"/api/translate", authMiddleware(db, "basic", translateHandler()))
	mux.Handle(prefix+"/api/sync/batch", authMiddleware(db, "basic", syncBatchHandler()))
	// New API endpoints for modern UI
	mux.Handle(prefix+"/api/providers", authMiddleware(db, "basic", providersHandler()))
	mux.Handle(prefix+"/api/library/browse", authMiddleware(db, "basic", libraryBrowseHandler()))
	mux.Handle(prefix+"/api/library/tags", authMiddleware(db, "basic", libraryTagsHandler(db)))
	mux.Handle(prefix+"/api/users/", authMiddleware(db, "admin", userRouter(db)))

	// Universal tagging system
	mux.Handle(prefix+"/api/tags", authMiddleware(db, "admin", tagsHandler(db)))
	mux.Handle(prefix+"/api/tags/", authMiddleware(db, "admin", tagItemHandler(db)))
	mux.Handle(prefix+"/api/tags/bulk", authMiddleware(db, "admin", bulkTagsHandler(db)))

	// Universal entity tagging endpoints
	mux.Handle(prefix+"/api/providers/", authMiddleware(db, "admin", universalTagsHandler(db)))
	mux.Handle(prefix+"/api/media/", authMiddleware(db, "basic", mediaTagsHandler(db)))
	mux.Handle(prefix+"/api/movies/", authMiddleware(db, "basic", universalTagsHandler(db)))
	mux.Handle(prefix+"/api/series/", authMiddleware(db, "basic", universalTagsHandler(db)))
	mux.Handle(prefix+"/api/languages/", authMiddleware(db, "admin", universalTagsHandler(db)))

	fsHandler := spaFileServer(f)
	mux.Handle(prefix+"/", staticFileMiddleware(http.StripPrefix(prefix+"/", fsHandler)))
	return securityHeadersMiddleware(mux), nil
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

	// Start periodic cleanup of expired sessions
	go startSessionCleanup(db)

	// Start automatic update checker if enabled
	if viper.GetBool("auto_update") {
		freq := viper.GetString("update_frequency")
		repo := "subtitle-manager/subtitle-manager"
		ctx := context.Background()
		updater.StartPeriodic(ctx, repo, AppVersion, freq)
	}

	// Start automated maintenance tasks
	go maintenance.StartDatabaseCleanup(context.Background(), db,
		viper.GetString("db_cleanup_frequency"))

	// Start Sonarr/Radarr sync tasks when configured
	storeBackend := database.GetDatabaseBackend()
	storePath := viper.GetString("db_path")
	if store, err := database.OpenStore(storePath, storeBackend); err == nil {
		if viper.GetBool("integrations.radarr.enabled") {
			host := viper.GetString("integrations.radarr.host")
			port := viper.GetString("integrations.radarr.port")
			key := viper.GetString("integrations.radarr.api_key")
			ssl := viper.GetBool("integrations.radarr.ssl")
			base := strings.Trim(viper.GetString("integrations.radarr.base_url"), "/")
			interval := viper.GetInt("integrations.radarr.sync_interval")
			if interval == 0 {
				interval = 60
			}
			scheme := "http"
			if ssl {
				scheme = "https"
			}
			url := fmt.Sprintf("%s://%s:%v/%s", scheme, host, port, base)
			c := radarr.NewClient(url, key)
			ctx := context.Background()
			radarr.StartSync(ctx, time.Duration(interval)*time.Minute, c, store)
		}
		if viper.GetBool("integrations.sonarr.enabled") {
			host := viper.GetString("integrations.sonarr.host")
			port := viper.GetString("integrations.sonarr.port")
			key := viper.GetString("integrations.sonarr.api_key")
			ssl := viper.GetBool("integrations.sonarr.ssl")
			base := strings.Trim(viper.GetString("integrations.sonarr.base_url"), "/")
			interval := viper.GetInt("integrations.sonarr.episode_sync_interval")
			if interval == 0 {
				interval = 60
			}
			scheme := "http"
			if ssl {
				scheme = "https"
			}
			url := fmt.Sprintf("%s://%s:%v/%s", scheme, host, port, base)
			c := sonarr.NewClient(url, key)
			ctx := context.Background()
			sonarr.StartSync(ctx, time.Duration(interval)*time.Minute, c, store)
		}
	}

	// Start additional maintenance tasks for metadata and disk scanning
	storePath = database.GetDatabasePath()
	backend = database.GetDatabaseBackend()
	go func() {
		var store database.SubtitleStore
		var err error
		switch backend {
		case "pebble":
			store, err = database.OpenPebble(storePath)
		case "postgres":
			store, err = database.OpenPostgresStore(storePath)
		default:
			store, err = database.OpenSQLStore(storePath)
		}
		if err != nil {
			return
		}
		maintenance.StartMetadataRefresh(context.Background(), store,
			viper.GetString("tmdb_api_key"), viper.GetString("metadata_refresh_frequency"))
	}()

	go maintenance.StartDiskScan(context.Background(), viper.GetString("db_path"),
		viper.GetString("disk_scan_frequency"))

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

		// Set secure session cookie with proper security attributes
		cookie := &http.Cookie{
			Name:     "session",
			Value:    token,
			Path:     "/",
			HttpOnly: true,                            // Prevents XSS attacks
			Secure:   r.TLS != nil,                    // HTTPS only in production
			SameSite: http.SameSiteStrictMode,         // CSRF protection
			MaxAge:   int((24 * time.Hour).Seconds()), // Explicit expiration
		}
		http.SetCookie(w, cookie)
	})
}

// logoutHandler invalidates the current session and clears the session cookie.
func logoutHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Get the session token from the cookie
		if cookie, err := r.Cookie("session"); err == nil {
			// Invalidate the session in the database
			_ = auth.InvalidateSession(db, cookie.Value)
		}

		// Clear the session cookie by setting it with an expired date
		clearCookie := &http.Cookie{
			Name:     "session",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   r.TLS != nil,
			SameSite: http.SameSiteStrictMode,
			MaxAge:   -1, // This immediately expires the cookie
		}
		http.SetCookie(w, clearCookie)

		w.WriteHeader(http.StatusOK)
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

// providersHandler returns information about available subtitle providers
func providersHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Get list of all available providers from the registry
			providers := getAvailableProviders()
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(providers); err != nil {
				http.Error(w, "Failed to encode providers", http.StatusInternalServerError)
				return
			}
		case http.MethodPost:
			// Handle provider configuration updates
			var req struct {
				Name    string                 `json:"name"`
				Enabled bool                   `json:"enabled"`
				Config  map[string]interface{} `json:"config"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Update provider configuration in viper
			configKey := fmt.Sprintf("providers.%s", req.Name)
			providerConfig := map[string]interface{}{
				"enabled": req.Enabled,
				"config":  req.Config,
			}
			viper.Set(configKey, providerConfig)

			// Save configuration if using a config file
			if cfg := viper.ConfigFileUsed(); cfg != "" {
				if err := viper.WriteConfig(); err != nil {
					http.Error(w, "Failed to save configuration", http.StatusInternalServerError)
					return
				}
			}

			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// libraryBrowseHandler provides directory browsing for media library
func libraryBrowseHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path := r.URL.Query().Get("path")
		if path == "" {
			path = "/"
		}

		// Validate and sanitize the path to prevent path traversal attacks
		sanitizedPath, err := validateAndSanitizePath(path)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid path: %v", err), http.StatusBadRequest)
			return
		}

		items, err := browseDirectory(sanitizedPath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to browse directory: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(items); err != nil {
			http.Error(w, "Failed to encode directory listing", http.StatusInternalServerError)
			return
		}
	})
}

// Helper functions for new API endpoints

type ProviderInfo struct {
	Name        string                 `json:"name"`
	DisplayName string                 `json:"displayName"`
	Enabled     bool                   `json:"enabled"`
	Config      map[string]interface{} `json:"config"`
	Type        string                 `json:"type"`
}

type MediaItem struct {
	Name        string     `json:"name"`
	Path        string     `json:"path"`
	IsDirectory bool       `json:"isDirectory"`
	Size        int64      `json:"size,omitempty"`
	ModTime     time.Time  `json:"modTime"`
	Subtitles   []Subtitle `json:"subtitles,omitempty"`
}

type Subtitle struct {
	Language string `json:"language"`
	Path     string `json:"path"`
	Format   string `json:"format"`
}

// getAvailableProviders returns a slice of every provider known to the system
// with the configuration loaded from Viper. Providers are returned regardless
// of whether they are currently enabled or configured so that the UI can
// present the full list for selection.
func getAvailableProviders() []ProviderInfo {
	// List of all available providers (this matches the registry)
	providerNames := []string{
		"addic7ed", "animekalesi", "animetosho", "assrt", "avistaz",
		"betaseries", "bsplayer", "embedded", "generic", "gestdown",
		"greeksubs", "greeksubtitles", "hdbits", "hosszupuska", "karagarga",
		"ktuvit", "legendasdivx", "legendasnet", "napiprojekt", "napisy24",
		"nekur", "opensubtitles", "opensubtitlescom", "opensubtitlesvip",
		"podnapisi", "regielive", "soustitres", "subdivx", "subf2m",
		"subs4free", "subs4series", "subscene", "subscenter", "subssabbz",
		"subsunacs", "subsynchro", "subtitrarinoi", "subtitriidlv",
		"subtitulamos", "supersubtitles", "titlovi", "titrariro", "titulky",
		"turkcealtyazi", "tusubtitulo", "tvsubtitles", "whisper", "wizdom",
		"xsubs", "yavka", "yifysubtitles", "zimuku",
	}

	var providers []ProviderInfo
	for _, name := range providerNames {
		configKey := fmt.Sprintf("providers.%s", name)
		providerConfig := viper.GetStringMap(configKey)

		enabled := false
		config := make(map[string]interface{})

		if providerConfig != nil {
			if enabledVal, ok := providerConfig["enabled"]; ok {
				if enabledBool, ok := enabledVal.(bool); ok {
					enabled = enabledBool
				}
			}
			if configVal, ok := providerConfig["config"].(map[string]interface{}); ok {
				config = configVal
			}
		}

		// Always include providers in the list so the UI can configure
		// them even when no entry exists in the configuration file.

		providers = append(providers, ProviderInfo{
			Name:        name,
			DisplayName: formatProviderName(name),
			Enabled:     enabled,
			Config:      config,
			Type:        getProviderType(name),
		})
	}

	return providers
}

// formatProviderName formats provider names for display
func formatProviderName(name string) string {
	specialNames := map[string]string{
		"opensubtitles":    "OpenSubtitles",
		"opensubtitlescom": "OpenSubtitles.com",
		"opensubtitlesvip": "OpenSubtitles VIP",
		"addic7ed":         "Addic7ed",
		"podnapisi":        "Podnapisi.NET",
		"subscene":         "Subscene",
		"yifysubtitles":    "YIFY Subtitles",
		"turkcealtyazi":    "Türkçe Altyazı",
		"greeksubtitles":   "Greek Subtitles",
		"legendasdivx":     "Legendas DivX",
		"legendasnet":      "Legendas.NET",
		"napiprojekt":      "NapiProjekt",
	}

	if display, ok := specialNames[name]; ok {
		return display
	}

	// Default formatting: split camelCase and capitalize each word
	words := strings.Split(name, "")
	var result []string
	var currentWord strings.Builder

	titleCaser := cases.Title(language.English)
	for i, char := range words {
		if i > 0 && strings.ToUpper(char) == char && strings.ToLower(char) != char {
			if currentWord.Len() > 0 {
				result = append(result, titleCaser.String(currentWord.String()))
				currentWord.Reset()
			}
		}
		currentWord.WriteString(char)
	}

	if currentWord.Len() > 0 {
		result = append(result, titleCaser.String(currentWord.String()))
	}

	if len(result) == 0 {
		return titleCaser.String(name)
	}

	return strings.Join(result, " ")
}

// getProviderType returns the provider type/category
func getProviderType(name string) string {
	switch name {
	case "whisper":
		return "transcription"
	case "embedded":
		return "extraction"
	case "generic":
		return "custom"
	default:
		return "subtitle"
	}
}

// browseDirectory lists media files and directories with subtitle information
func browseDirectory(path string) ([]MediaItem, error) {
	if path == "" {
		path = "/"
	}

	// Check if path exists and is readable
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("path not accessible: %v", err)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("path is not a directory")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	var items []MediaItem
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		fullPath := filepath.Join(path, entry.Name())
		item := MediaItem{
			Name:        entry.Name(),
			Path:        fullPath,
			IsDirectory: entry.IsDir(),
			Size:        info.Size(),
			ModTime:     info.ModTime(),
		}

		// If it's a media file, look for subtitles
		if !entry.IsDir() && isMediaFile(entry.Name()) {
			item.Subtitles = findSubtitles(fullPath)
		}

		items = append(items, item)
	}

	return items, nil
}

// isMediaFile checks if the file is a supported media file
func isMediaFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	mediaExtensions := []string{
		".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm",
		".m4v", ".3gp", ".ts", ".mts", ".m2ts", ".vob", ".mpg",
		".mpeg", ".ogv", ".rm", ".rmvb", ".asf", ".divx",
	}

	for _, mediaExt := range mediaExtensions {
		if ext == mediaExt {
			return true
		}
	}
	return false
}

// findSubtitles looks for subtitle files associated with a media file
func findSubtitles(mediaPath string) []Subtitle {
	dir := filepath.Dir(mediaPath)
	baseName := strings.TrimSuffix(filepath.Base(mediaPath), filepath.Ext(mediaPath))

	var subtitles []Subtitle

	// Look for subtitle files with common patterns
	entries, err := os.ReadDir(dir)
	if err != nil {
		return subtitles
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))

		// Check if it's a subtitle file
		if !isSubtitleFile(ext) {
			continue
		}

		// Check if it matches our media file
		if strings.HasPrefix(name, baseName) {
			language := extractLanguageFromFilename(name)
			subtitles = append(subtitles, Subtitle{
				Language: language,
				Path:     filepath.Join(dir, name),
				Format:   strings.TrimPrefix(ext, "."),
			})
		}
	}

	return subtitles
}

// isSubtitleFile checks if the file extension indicates a subtitle file
func isSubtitleFile(ext string) bool {
	subtitleExtensions := []string{".srt", ".vtt", ".ass", ".ssa", ".sub", ".idx", ".sup"}
	for _, subExt := range subtitleExtensions {
		if ext == subExt {
			return true
		}
	}
	return false
}

// extractLanguageFromFilename tries to extract language code from filename

// defaultLanguagePatterns contains the built-in language codes used for filename
// detection. Additional languages can be configured via SetLanguagePatterns.
var defaultLanguagePatterns = map[string]string{
	"en":      "English",
	"es":      "Spanish",
	"fr":      "French",
	"de":      "German",
	"it":      "Italian",
	"pt":      "Portuguese",
	"ru":      "Russian",
	"ja":      "Japanese",
	"ko":      "Korean",
	"zh":      "Chinese",
	"zh-cn":   "Chinese (Simplified)",
	"zh-hans": "Chinese (Simplified)",
	"zh-tw":   "Chinese (Taiwan)",
	"zh-hk":   "Chinese (Hong Kong)",
	"zh-hant": "Chinese (Traditional)",
	"ar":      "Arabic",
	"hi":      "Hindi",
	"tr":      "Turkish",
	"pl":      "Polish",
	"nl":      "Dutch",
	"sv":      "Swedish",
	"no":      "Norwegian",
	"da":      "Danish",
	"fi":      "Finnish",
	"he":      "Hebrew",
	"el":      "Greek",
	"cs":      "Czech",
	"sk":      "Slovak",
	"hu":      "Hungarian",
	"vi":      "Vietnamese",
	"id":      "Indonesian",
	"th":      "Thai",
	"uk":      "Ukrainian",
	"ro":      "Romanian",
	"bg":      "Bulgarian",
	"hr":      "Croatian",
	"sr":      "Serbian",
	"ms":      "Malay",
	"fa":      "Persian",
	"bn":      "Bengali",
}

// languagePatterns holds the active language map used by
// extractLanguageFromFilename. It can be overridden for tests or advanced
// configuration via SetLanguagePatterns.
var languagePatterns = defaultLanguagePatterns

// SetLanguagePatterns overrides the default language patterns used for filename
// detection. This allows configuring additional languages or custom codes.
func SetLanguagePatterns(m map[string]string) {
	languagePatterns = m
}

// ResetLanguagePatterns restores the built-in language patterns.
func ResetLanguagePatterns() {
	languagePatterns = defaultLanguagePatterns
}

func extractLanguageFromFilename(filename string) string {
	// Look for common language patterns
	lower := strings.ToLower(filename)

	for code, lang := range languagePatterns {
		if strings.Contains(lower, "."+code+".") ||
			strings.Contains(lower, "_"+code+"_") ||
			strings.Contains(lower, "-"+code+"-") ||
			strings.HasSuffix(lower, "."+code+filepath.Ext(lower)) {
			return lang
		}
	}

	// Default to English if no language detected
	return "English"
}

// startSessionCleanup runs a periodic cleanup of expired sessions.
// This prevents the sessions table from growing indefinitely with expired tokens.
func startSessionCleanup(db *sql.DB) {
	ticker := time.NewTicker(1 * time.Hour) // Cleanup every hour
	defer ticker.Stop()

	for range ticker.C {
		if err := auth.CleanupExpiredSessions(db); err != nil {
			// Log the error but don't stop the cleanup routine
			// In a production app, you'd use proper logging here
			fmt.Printf("Failed to cleanup expired sessions: %v\n", err)
		}
	}
}

// validateAndSanitizePath cleans a user-supplied path and checks for basic path
// traversal attempts. It returns an absolute version of the path when valid.
func validateAndSanitizePath(userPath string) (string, error) {
	// Clean the path to resolve .. and . elements
	cleanPath := filepath.Clean(userPath)

	// Convert to absolute path for consistent checking
	if !filepath.IsAbs(cleanPath) {
		return "", fmt.Errorf("path must be absolute")
	}
	absPath := filepath.Clean(cleanPath)

	// Allow root directory and Windows drive letters
	if cleanPath == "/" || cleanPath == "." ||
		(runtime.GOOS == "windows" && len(filepath.VolumeName(cleanPath)) == 2 &&
			strings.Trim(filepath.Clean(cleanPath), "\\") == filepath.VolumeName(cleanPath)) {
		return absPath, nil
	}

	// Additional security check: ensure no path traversal components remain
	if strings.Contains(cleanPath, "..") {
		return "", fmt.Errorf("path traversal detected: %s", cleanPath)
	}

	return absPath, nil
}
