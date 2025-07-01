package webserver

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"syscall"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/backups"
	"github.com/jdfalk/subtitle-manager/pkg/errors"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/tasks"
)

// isValidTaskName validates that a task name contains only safe characters
// to prevent injection attacks or invalid task names.
func isValidTaskName(name string) bool {
	// Allow only alphanumeric characters, hyphens, underscores, and dots
	// Maximum length of 50 characters
	if len(name) == 0 || len(name) > 50 {
		return false
	}

	validNamePattern := regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)
	return validNamePattern.MatchString(name)
}

// logsHandler returns recent log lines captured in memory.
func logsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(logging.Hook.Logs())
	})
}

// systemHandler exposes basic system information.
func systemHandler() http.Handler {
	type info struct {
		GoVersion  string `json:"go_version"`
		OS         string `json:"os"`
		Arch       string `json:"arch"`
		Goroutines int    `json:"goroutines"`
		DiskFree   uint64 `json:"disk_free"`
		DiskTotal  uint64 `json:"disk_total"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var statfs syscall.Statfs_t
		root := "/"
		_ = syscall.Statfs(root, &statfs)
		data := info{
			GoVersion:  runtime.Version(),
			OS:         runtime.GOOS,
			Arch:       runtime.GOARCH,
			Goroutines: runtime.NumGoroutine(),
			DiskFree:   statfs.Bfree * uint64(statfs.Bsize),
			DiskTotal:  statfs.Blocks * uint64(statfs.Bsize),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	})
}

// tasksHandler reports status for background tasks such as scanning.
func tasksHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(tasks.List())
	})
}

// startTaskHandler launches a named task and returns its initial state.
func startTaskHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate task name to prevent potential injection
		if !isValidTaskName(name) {
			http.Error(w, "Invalid task name", http.StatusBadRequest)
			return
		}
		t := tasks.Start(r.Context(), name, func(ctx context.Context) error {
			// Simulate a short running task
			for i := 0; i <= 10; i++ {
				tasks.Update(name, i*10)
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(100 * time.Millisecond):
				}
			}
			return nil
		})
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(t)
	})
}

// providerStatusHandler returns the current provider availability map.
func providerStatusHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(providers.List())
	})
}

// providerRefreshHandler refreshes status information for all providers.
func providerRefreshHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providers.Refresh(r.Context(), []string{"opensubtitles", "subscene"})
		w.WriteHeader(http.StatusAccepted)
	})
}

// providerResetHandler clears provider status data.
func providerResetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providers.Reset()
		w.WriteHeader(http.StatusNoContent)
	})
}

// backupsHandler lists known backups.
func backupsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(backups.List())
	})
}

// createBackupHandler creates a new backup and returns it.
func createBackupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		b := backups.Create()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(b)
	})
}

// restoreBackupHandler restores the latest backup.
func restoreBackupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := backups.Restore()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"restored": name})
	})
}

// releasesHandler fetches release information from GitHub.
func releasesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet,
			"https://api.github.com/repos/subtitle-manager/subtitle-manager/releases/latest", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, _ = io.Copy(w, resp.Body)
	})
}

// announcementsHandler returns messages from announcements.json.
func announcementsHandler() http.Handler {
	type ann struct {
		Date    string `json:"date"`
		Message string `json:"message"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open(filepath.Join("..", "..", "announcements.json"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		defer f.Close()
		var data []ann
		if err := json.NewDecoder(f).Decode(&data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	})
}

// errorStatsHandler returns error statistics for the admin dashboard.
func errorStatsHandler() http.Handler {
	dashboard := errors.GetDashboard()
	return dashboard.StatsHandler()
}

// errorRecentHandler returns recent error events.
func errorRecentHandler() http.Handler {
	dashboard := errors.GetDashboard()
	return dashboard.RecentHandler()
}

// errorTopHandler returns the most frequent errors.
func errorTopHandler() http.Handler {
	dashboard := errors.GetDashboard()
	return dashboard.TopErrorsHandler()
}

// errorHealthHandler returns overall error health status.
func errorHealthHandler() http.Handler {
	dashboard := errors.GetDashboard()
	return dashboard.HealthHandler()
}
