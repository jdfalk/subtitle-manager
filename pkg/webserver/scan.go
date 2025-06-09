// file: pkg/webserver/scan.go
package webserver

import (
    "context"
    "encoding/json"
    "net/http"
    "sync"

    "github.com/spf13/viper"

    "subtitle-manager/pkg/providers"
    "subtitle-manager/pkg/scanner"
)

// scanStatus tracks progress for an active scan.
type scanStatus struct {
    Running   bool     `json:"running"`
    Completed int      `json:"completed"`
    Files     []string `json:"files"`
}

var (
    scanMu sync.Mutex
    status = scanStatus{}
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
        status = scanStatus{Running: true}
        scanMu.Unlock()
        go func() {
            key := viper.GetString("opensubtitles.api_key")
            p, err := providers.Get(q.Provider, key)
            if err != nil {
                scanMu.Lock()
                status.Running = false
                scanMu.Unlock()
                return
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
