// file: pkg/webserver/convert.go
package webserver

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

// convertHandler handles POST /api/convert requests.
// It expects a multipart form with a "file" field containing a subtitle file.
// The subtitle is converted to SRT and returned as a downloadable attachment.
func convertHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		f, hdr, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer f.Close()

		// Preserve the original file extension for astisub format detection
		ext := filepath.Ext(hdr.Filename)
		tmp, err := os.CreateTemp("", "convert-*"+ext)
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
		w.Header().Set("Content-Type", "application/x-subrip")
		w.Header().Set("Content-Disposition", "attachment; filename=\"converted.srt\"")
		_, _ = w.Write(data)
	})
}
