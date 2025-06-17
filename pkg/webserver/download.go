// file: pkg/webserver/download.go
package webserver

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var p providers.Provider
		var name string
		var err error
		if q.Provider != "" {
			p, err = providers.Get(q.Provider, "")
			name = q.Provider
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		if err := scanner.ProcessFile(r.Context(), q.Path, q.Lang, name, p, false, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		out := strings.TrimSuffix(q.Path, filepath.Ext(q.Path)) + "." + q.Lang + ".srt"
		if db != nil {
			_ = database.InsertDownload(db, out, q.Path, name, q.Lang)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp{File: out})
	})
}
