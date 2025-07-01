// file: pkg/webserver/whisper.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174002

package webserver

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/jdfalk/subtitle-manager/pkg/transcriber"
)

// whisperContainerStatusHandler returns the status of the Whisper container.
func whisperContainerStatusHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		container, err := transcriber.NewWhisperContainer()
		if err != nil {
			http.Error(w, "Failed to create container client", http.StatusInternalServerError)
			return
		}
		defer container.Close()

		status, err := container.GetContainerStatus(r.Context())
		if err != nil {
			http.Error(w, "Failed to get container status", http.StatusInternalServerError)
			return
		}

		running, err := container.IsContainerRunning(r.Context())
		if err != nil {
			http.Error(w, "Failed to check container running state", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"status":  status,
			"running": running,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}

// whisperContainerStartHandler starts the Whisper container.
func whisperContainerStartHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		container, err := transcriber.NewWhisperContainer()
		if err != nil {
			http.Error(w, "Failed to create container client", http.StatusInternalServerError)
			return
		}
		defer container.Close()

		err = container.StartContainer(r.Context())
		if err != nil {
			http.Error(w, "Failed to start container: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "started"})
	})
}

// whisperContainerStopHandler stops the Whisper container.
func whisperContainerStopHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		container, err := transcriber.NewWhisperContainer()
		if err != nil {
			http.Error(w, "Failed to create container client", http.StatusInternalServerError)
			return
		}
		defer container.Close()

		err = container.StopContainer(r.Context())
		if err != nil {
			http.Error(w, "Failed to stop container: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "stopped"})
	})
}

// whisperTranscribeHandler starts a transcription job.
func whisperTranscribeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			FilePath string `json:"file_path"`
			Language string `json:"language"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.FilePath == "" {
			http.Error(w, "file_path is required", http.StatusBadRequest)
			return
		}

		// Validate file extension
		ext := filepath.Ext(req.FilePath)
		validExts := []string{".mp3", ".wav", ".m4a", ".flac", ".ogg", ".mp4", ".avi", ".mov", ".mkv"}
		isValid := false
		for _, validExt := range validExts {
			if ext == validExt {
				isValid = true
				break
			}
		}
		if !isValid {
			http.Error(w, "Unsupported file format", http.StatusBadRequest)
			return
		}

		container, err := transcriber.NewWhisperContainer()
		if err != nil {
			http.Error(w, "Failed to create container client", http.StatusInternalServerError)
			return
		}
		defer container.Close()

		task, err := container.TranscribeFile(r.Context(), req.FilePath, req.Language)
		if err != nil {
			http.Error(w, "Failed to start transcription: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"task_id": task.ID,
			"status":  "started",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}

// whisperModelsHandler returns supported Whisper models.
func whisperModelsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := map[string]interface{}{
			"models": transcriber.SupportedModels,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}
