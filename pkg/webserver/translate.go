// file: pkg/webserver/translate.go
package webserver

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

// translateHandler handles translating an uploaded subtitle file.
//
// POST requests must use multipart/form-data containing a "file" part and
// a "lang" field specifying the target language. Optional fields "service"
// and "grpc" override the configured translation service and gRPC address.
// The translated SRT data is returned with content type "application/x-subrip".
func translateHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		file, hdr, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()
		lang := r.FormValue("lang")
		if lang == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		service := r.FormValue("service")
		if service == "" {
			service = viper.GetString("translate_service")
			if service == "" {
				service = "google"
			}
		}
		grpcAddr := r.FormValue("grpc")
		if grpcAddr == "" {
			grpcAddr = viper.GetString("grpc_addr")
		}
		gKey := viper.GetString("google_api_key")
		gptKey := viper.GetString("openai_api_key")

		in, err := os.CreateTemp("", "in-*"+filepath.Ext(hdr.Filename))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer func() { _ = in.Close() }()
		defer os.Remove(in.Name())
		if _, err := io.Copy(in, file); err != nil {
			_ = in.Close()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_ = in.Close()

		out, err := os.CreateTemp("", "out-*.srt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer func() { _ = out.Close() }()
		defer os.Remove(out.Name())
		if err := out.Close(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := subtitles.TranslateFileToSRT(in.Name(), out.Name(), lang, service, gKey, gptKey, grpcAddr); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := os.ReadFile(out.Name())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/x-subrip")
		_, _ = w.Write(data)
	})
}
