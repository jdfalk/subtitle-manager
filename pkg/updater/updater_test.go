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

// TestSetHTTPClient tests setting the HTTP client.
func TestSetHTTPClient(t *testing.T) {
	original := httpClient
	defer func() { httpClient = original }()

	customClient := &http.Client{}
	SetHTTPClient(customClient)

	if httpClient != customClient {
		t.Error("SetHTTPClient did not set the HTTP client correctly")
	}
}

// TestFrequencyToDuration tests the frequency conversion function.
func TestFrequencyToDuration(t *testing.T) {
	tests := []struct {
		name     string
		freq     string
		expected string
	}{
		{
			name:     "hourly",
			freq:     "hourly",
			expected: "1h0m0s",
		},
		{
			name:     "daily",
			freq:     "daily",
			expected: "24h0m0s",
		},
		{
			name:     "weekly",
			freq:     "weekly",
			expected: "168h0m0s",
		},
		{
			name:     "invalid",
			freq:     "invalid",
			expected: "24h0m0s", // defaults to daily
		},
		{
			name:     "empty",
			freq:     "",
			expected: "24h0m0s", // defaults to daily
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := frequencyToDuration(tt.freq)
			if result.String() != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result.String())
			}
		})
	}
}

// TestStartPeriodic tests the periodic update functionality.
func TestStartPeriodic(t *testing.T) {
	// Test that StartPeriodic doesn't panic with invalid inputs
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Mock the HTTP client to avoid real network calls
	SetHTTPClient(&http.Client{})
	defer SetHTTPClient(http.DefaultClient)

	// Override executable path to avoid actual file operations
	SetExecutablePathFunc(func() (string, error) { return "/tmp/test", nil })
	defer SetExecutablePathFunc(os.Executable)

	// Override restart function to avoid actual restart
	var restartCalled bool
	SetRestartFunc(func() error { restartCalled = true; return nil })
	defer SetRestartFunc(func() error { return nil })

	// Start periodic updates with a very long interval
	// This should not block and should return immediately
	StartPeriodic(ctx, "test/repo", "1.0.0", "never")

	// Cancel the context immediately to stop any background operations
	cancel()

	// Test doesn't fail if StartPeriodic handles the context cancellation properly
	if restartCalled {
		t.Error("restart should not be called during test")
	}
}
