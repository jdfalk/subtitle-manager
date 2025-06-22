package bazarr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/security"
)

// Settings represents a subset of Bazarr's configuration used for import.
type Settings map[string]any

// FetchSettings retrieves Bazarr settings from the given baseURL using the provided API key.
// The baseURL should include scheme and host, for example "http://localhost:6767".
// A non-200 status code results in an error.
// This function validates the baseURL to prevent SSRF attacks.
func FetchSettings(baseURL, apiKey string) (Settings, error) {
	// Validate the base URL to prevent SSRF attacks
	if _, err := security.ValidateURL(baseURL); err != nil {
		return nil, fmt.Errorf("invalid baseURL: %v", err)
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid baseURL: %v", err)
	}
	endpoint, err := url.Parse("/api/system/settings")
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint URL: %v", err)
	}
	fullURL := base.ResolveReference(endpoint)

	req, err := http.NewRequest(http.MethodGet, fullURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", apiKey)

	// Create a client with timeouts to prevent hanging requests
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var cfg Settings
	if err := json.NewDecoder(resp.Body).Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
