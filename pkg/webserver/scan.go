// file: pkg/webserver/scan.go
package webserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/metadata"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/radarr"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
	"github.com/jdfalk/subtitle-manager/pkg/sonarr"
)

// scanStatus tracks progress for an active scan.
type scanStatus struct {
	Running   bool     `json:"running"`
	Completed int      `json:"completed"`
	Files     []string `json:"files"`
}

var (
	scanMu sync.Mutex
	status = scanStatus{Files: []string{}}
)

// libScanStatus tracks progress for library scanning.
type libScanStatus struct {
	Running   bool     `json:"running"`
	Completed int      `json:"completed"`
	Files     []string `json:"files"`
}

var (
	libMu     sync.Mutex
	libStatus = libScanStatus{Files: []string{}}
)

func scanHandler() http.Handler {
	type req struct {
		Provider  string `json:"provider"`
		Directory string `json:"directory"`
		Lang      string `json:"lang"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var q req
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		scanMu.Lock()
		if status.Running {
			scanMu.Unlock()
			w.WriteHeader(http.StatusConflict)
			return
		}
		status = scanStatus{Running: true, Files: []string{}}
		scanMu.Unlock()
		go func() {
			var p providers.Provider
			var err error
			if q.Provider != "" {
				p, err = providers.Get(q.Provider, "")
				if err != nil {
					scanMu.Lock()
					status.Running = false
					scanMu.Unlock()
					return
				}
			}
			cb := func(f string) {
				scanMu.Lock()
				status.Completed++
				status.Files = append(status.Files, f)
				scanMu.Unlock()
			}
			_ = scanner.ScanDirectoryProgress(context.Background(), q.Directory, q.Lang, q.Provider, p, false, 2, nil, cb)
			scanMu.Lock()
			status.Running = false
			scanMu.Unlock()
		}()
		w.WriteHeader(http.StatusAccepted)
	})
}

func scanStatusHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		scanMu.Lock()
		defer scanMu.Unlock()
		_ = json.NewEncoder(w).Encode(status)
	})
}

func libraryScanHandler(db *sql.DB) http.Handler {
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
		libMu.Lock()
		if libStatus.Running {
			libMu.Unlock()
			w.WriteHeader(http.StatusConflict)
			return
		}
		libStatus = libScanStatus{Running: true, Files: []string{}}
		libMu.Unlock()
		go func() {
			backend := database.GetDatabaseBackend()
			storePath := database.GetDatabasePath()
			store, err := database.OpenStore(storePath, backend)
			if err != nil {
				libMu.Lock()
				libStatus.Running = false
				libMu.Unlock()
				return
			}
			defer store.Close()
			cb := func(f string) {
				libMu.Lock()
				libStatus.Completed++
				libStatus.Files = append(libStatus.Files, f)
				libMu.Unlock()
			}
			_ = metadata.ScanLibraryProgress(context.Background(), q.Path, store, cb)

			if viper.GetBool("integrations.radarr.enabled") {
				host := viper.GetString("integrations.radarr.host")
				port := viper.GetString("integrations.radarr.port")
				key := viper.GetString("integrations.radarr.api_key")
				ssl := viper.GetBool("integrations.radarr.ssl")
				base := strings.Trim(viper.GetString("integrations.radarr.base_url"), "/")
				scheme := "http"
				if ssl {
					scheme = "https"
				}
				url := fmt.Sprintf("%s://%s:%s/%s", scheme, host, port, base)
				c := radarr.NewClient(url, key)
				_ = radarr.Sync(context.Background(), c, store)
			}
			if viper.GetBool("integrations.sonarr.enabled") {
				host := viper.GetString("integrations.sonarr.host")
				port := viper.GetString("integrations.sonarr.port")
				key := viper.GetString("integrations.sonarr.api_key")
				ssl := viper.GetBool("integrations.sonarr.ssl")
				base := strings.Trim(viper.GetString("integrations.sonarr.base_url"), "/")
				scheme := "http"
				if ssl {
					scheme = "https"
				}
				url := fmt.Sprintf("%s://%s:%s/%s", scheme, host, port, base)
				c := sonarr.NewClient(url, key)
				_ = sonarr.Sync(context.Background(), c, store)
			}
			libMu.Lock()
			libStatus.Running = false
			libMu.Unlock()
		}()
		w.WriteHeader(http.StatusAccepted)
	})
}

func libraryScanStatusHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		libMu.Lock()
		defer libMu.Unlock()
		_ = json.NewEncoder(w).Encode(libStatus)
	})
}
