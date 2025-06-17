package updater

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// TestSelfUpdate verifies that SelfUpdate downloads the new binary and replaces
// the executable without restarting when restartFunc is overridden.
func TestSelfUpdate(t *testing.T) {
	tmpDir := t.TempDir()
	exePath := filepath.Join(tmpDir, "app")
	os.WriteFile(exePath, []byte("old"), 0755)

	SetExecutablePathFunc(func() (string, error) { return exePath, nil })
	defer SetExecutablePathFunc(os.Executable)

	var restarted bool
	SetRestartFunc(func() error { restarted = true; return nil })
	defer SetRestartFunc(func() error { return nil })

	var url string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/repos/test/repo/releases/latest":
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, `{"tag_name":"v1.1.0","assets":[{"name":"subtitle-manager-`+runtime.GOOS+`-`+runtime.GOARCH+`","browser_download_url":"`+url+`/bin"}]}`)
		case "/bin":
			w.Write([]byte("new"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	url = srv.URL
	defer srv.Close()

	SetGitHubAPIBaseURL(srv.URL)
	defer SetGitHubAPIBaseURL("https://api.github.com")

	if err := SelfUpdate(context.Background(), "test/repo", "1.0.0"); err != nil {
		t.Fatalf("self update: %v", err)
	}
	data, _ := os.ReadFile(exePath)
	if string(data) != "new" {
		t.Fatalf("binary not replaced")
	}
	if !restarted {
		t.Fatalf("restartFunc not called")
	}
}
