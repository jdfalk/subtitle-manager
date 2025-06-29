// file: pkg/webserver/download.go
package webserver

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jdfalk/subtitle-manager/pkg/database"
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
func downloadHandler(db *sql.DB) http.Handler {
	type req struct {
		Provider string `json:"provider"`
		Path     string `json:"path"`
		Lang     string `json:"lang"`
	}
	type resp struct {
		File string `json:"file"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var q req
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil || q.Path == "" || q.Lang == "" {
			metrics.APIRequests.WithLabelValues("/api/download", "POST", "400").Inc()
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate and sanitize the file path to prevent path injection
		validatedPath, err := security.ValidateAndSanitizePath(q.Path)
		if err != nil {
			http.Error(w, "Invalid file path", http.StatusBadRequest)
			return
		}

		// Validate the language code to ensure it conforms to expected patterns
		if err := security.ValidateLanguageCode(q.Lang); err != nil {
			http.Error(w, "Invalid language code", http.StatusBadRequest)
			return
		}
		// Additional validation to ensure language code does not contain path traversal characters
		if strings.Contains(q.Lang, "/") || strings.Contains(q.Lang, "\\") || strings.Contains(q.Lang, "..") {
			http.Error(w, "Invalid language code", http.StatusBadRequest)
			return
		}

		// Validate provider name if provided
		if q.Provider != "" {
			if err := security.ValidateProviderName(q.Provider); err != nil {
				http.Error(w, "Invalid provider name", http.StatusBadRequest)
				return
			}
		}
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
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		if err := scanner.ProcessFile(r.Context(), validatedPath, q.Lang, name, p, false, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Construct output path using validated inputs
		out, outErr := security.ValidateSubtitleOutputPath(validatedPath, q.Lang)
		if outErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if db != nil {
			_ = database.InsertDownload(db, out, validatedPath, name, q.Lang)
		}
		metrics.APIRequests.WithLabelValues("/api/download", "POST", "200").Inc()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp{File: out})
	})
}
