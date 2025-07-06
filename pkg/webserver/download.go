// file: pkg/webserver/download.go
package webserver

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/metrics"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
	"github.com/jdfalk/subtitle-manager/pkg/security"
)

// downloadHandler downloads a subtitle using a provider and stores history.
//
// POST requests expect a JSON body {"provider":"generic","path":"/file.mkv","lang":"en"}.
// The subtitle is written next to the media file and the resulting path is
// returned as JSON {"file":"/file.en.srt"}.
//
// Improvements:
// - Extensive documentation for handler and logic
// - Consistent JSON error responses
// - Robust validation and error handling
// - Comments for maintainability
// - Detailed logging for observability
func downloadHandler(db *sql.DB) http.Handler {
	type req struct {
		Provider string `json:"provider"`
		Path     string `json:"path"`
		Lang     string `json:"lang"`
	}
	type resp struct {
		File string `json:"file"`
	}
	type apiError struct {
		Error string `json:"error"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetLogger("webserver.download")
		if r.Method != http.MethodPost {
			logger.Debugf("method %s not allowed", r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			_ = json.NewEncoder(w).Encode(apiError{Error: "Method not allowed"})
			return
		}

		var q req
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil || q.Path == "" || q.Lang == "" {
			logger.Warnf("invalid request body: %v", err)
			metrics.APIRequests.WithLabelValues("/api/download", "POST", "400").Inc()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(apiError{Error: "Invalid request body, path, or language"})
			return
		}

		// Validate and sanitize the file path to prevent path injection
		validatedPath, err := security.ValidateAndSanitizePath(q.Path)
		if err != nil {
			logger.Warnf("invalid file path: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(apiError{Error: "Invalid file path: " + err.Error()})
			return
		}

		// Validate the language code to ensure it conforms to expected patterns
		if err := security.ValidateLanguageCode(q.Lang); err != nil {
			logger.Warnf("invalid language code: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(apiError{Error: "Invalid language code: " + err.Error()})
			return
		}

		// Validate provider name if provided
		if q.Provider != "" {
			if err := security.ValidateProviderName(q.Provider); err != nil {
				logger.Warnf("invalid provider name: %v", err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(apiError{Error: "Invalid provider name: " + err.Error()})
				return
			}
		}

		// Provider resolution
		var p providers.Provider
		var name string
		var providerErr error
		if q.Provider != "" {
			if inst, ok := providers.GetInstance(q.Provider); ok {
				p, providerErr = providers.Get(inst.Name, "")
				name = inst.ID
			} else {
				p, providerErr = providers.Get(q.Provider, "")
				name = q.Provider
			}
			if providerErr != nil {
				logger.Warnf("provider resolution failed: %v", providerErr)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(apiError{Error: "Provider not found or error: " + providerErr.Error()})
				return
			}
		}

		// Process the file using the scanner
		if err := scanner.ProcessFile(r.Context(), validatedPath, q.Lang, name, p, false, nil); err != nil {
			logger.Errorf("failed to process file %s: %v", validatedPath, err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(apiError{Error: "Failed to process file: " + err.Error()})
			return
		}

		// Construct output path using validated inputs
		out, outErr := security.ValidateSubtitleOutputPath(validatedPath, q.Lang)
		if outErr != nil {
			logger.Errorf("failed to construct output path for %s: %v", validatedPath, outErr)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(apiError{Error: "Failed to construct output path: " + outErr.Error()})
			return
		}

		// Insert download record into database if available
		if db != nil {
			if err := database.InsertDownload(db, out, validatedPath, name, q.Lang); err != nil {
				logger.Warnf("failed to record download: %v", err)
			}
		}

		metrics.APIRequests.WithLabelValues("/api/download", "POST", "200").Inc()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp{File: out})
		logger.Infof("downloaded subtitle %s for %s", out, validatedPath)
	})
}
