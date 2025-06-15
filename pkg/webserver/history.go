// file: pkg/webserver/history.go
package webserver

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// historyHandler returns translation and download history as JSON.
// The optional `lang` query parameter filters results by language.
func historyHandler(db *sql.DB) http.Handler {
	type resp struct {
		Translations []database.SubtitleRecord `json:"translations"`
		Downloads    []database.DownloadRecord `json:"downloads"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		lang := r.URL.Query().Get("lang")
		subs, err := database.ListSubtitles(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		downloads, err := database.ListDownloads(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if lang != "" {
			filtered := subs[:0]
			for _, s := range subs {
				if s.Language == lang {
					filtered = append(filtered, s)
				}
			}
			subs = filtered
			fdl := downloads[:0]
			for _, d := range downloads {
				if d.Language == lang {
					fdl = append(fdl, d)
				}
			}
			downloads = fdl
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp{Translations: subs, Downloads: downloads})
	})
}
