package updater

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// githubAPIBaseURL is the base URL for GitHub API requests. It can be
// overridden in tests using SetGitHubAPIBaseURL.
var githubAPIBaseURL = "https://api.github.com"

// httpClient is used for network requests. Tests can replace it with a custom
// client.
var httpClient = &http.Client{Timeout: 30 * time.Second}

// exePathFunc returns the path to the current executable. It is overridable for
// testing.
var exePathFunc = os.Executable

// restartFunc is called after a successful update to launch the new binary.
// The default implementation starts the new process and exits.
var restartFunc = func() error {
	exe, err := exePathFunc()
	if err != nil {
		return err
	}
	cmd := exec.Command(exe, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	if err := cmd.Start(); err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

// SetGitHubAPIBaseURL overrides the GitHub API base URL. This is primarily used
// for tests.
func SetGitHubAPIBaseURL(u string) { githubAPIBaseURL = u }

// SetHTTPClient overrides the HTTP client used for requests. Used in tests.
func SetHTTPClient(c *http.Client) { httpClient = c }

// SetExecutablePathFunc overrides the function used to determine the
// executable path. Used in tests.
func SetExecutablePathFunc(fn func() (string, error)) { exePathFunc = fn }

// SetRestartFunc overrides the restart behavior. Used in tests to avoid
// exiting the process.
func SetRestartFunc(fn func() error) { restartFunc = fn }

// Release represents the subset of GitHub release information used by the
// updater.
type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name string `json:"name"`
		URL  string `json:"browser_download_url"`
	} `json:"assets"`
}

// checkForUpdate retrieves the latest release for repo and compares it with the
// current version. It returns the release when a newer version is available.
func checkForUpdate(ctx context.Context, repo, current string) (*Release, bool, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, githubAPIBaseURL+"/repos/"+repo+"/releases/latest", nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var rel Release
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return nil, false, err
	}
	latest := strings.TrimPrefix(rel.TagName, "v")
	if latest == current {
		return &rel, false, nil
	}
	return &rel, true, nil
}

// findAsset returns the download URL for the release asset matching the current
// OS and architecture.
const projectName = "subtitle-manager"
func findAsset(rel *Release) (string, error) {
	name := fmt.Sprintf("%s-%s-%s", projectName, runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	for _, a := range rel.Assets {
		if a.Name == name {
			return a.URL, nil
		}
	}
	return "", fmt.Errorf("asset %s not found", name)
}

// downloadBinary retrieves the binary from url and returns its contents.
func downloadBinary(ctx context.Context, url string) ([]byte, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return io.ReadAll(resp.Body)
}

// replaceExecutable writes data to the current executable path.
func replaceExecutable(data []byte) error {
	path, err := exePathFunc()
	if err != nil {
		return err
	}
	tmp := path + ".new"
	if err := os.WriteFile(tmp, data, 0755); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		return err
	}
	return nil
}

// SelfUpdate checks for a newer release in repo and, if available, downloads the
// appropriate binary, replaces the current executable, and restarts the
// application.
func SelfUpdate(ctx context.Context, repo, current string) error {
	rel, newer, err := checkForUpdate(ctx, repo, current)
	if err != nil || !newer {
		return err
	}
	url, err := findAsset(rel)
	if err != nil {
		return err
	}
	data, err := downloadBinary(ctx, url)
	if err != nil {
		return err
	}
	if err := replaceExecutable(data); err != nil {
		return err
	}
	return restartFunc()
}

// frequencyToDuration converts a textual frequency like "daily" to a
// time.Duration. Unknown values default to 24 hours.
func frequencyToDuration(freq string) time.Duration {
	switch strings.ToLower(freq) {
	case "hourly":
		return time.Hour
	case "daily":
		return 24 * time.Hour
	case "weekly":
		return 7 * 24 * time.Hour
	default:
		d, err := time.ParseDuration(freq)
		if err != nil {
			return 24 * time.Hour
		}
		return d
	}
}

// StartPeriodic updates periodically according to frequency. The repo and
// current version are used for update checks. The provided context controls the
// lifetime of the background goroutine.
func StartPeriodic(ctx context.Context, repo, current, frequency string) {
	interval := frequencyToDuration(frequency)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				_ = SelfUpdate(context.Background(), repo, current)
			}
		}
	}()
}
