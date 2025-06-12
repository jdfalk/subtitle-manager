package bazarr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Settings represents a subset of Bazarr's configuration used for import.
type Settings map[string]any

// FetchSettings retrieves Bazarr settings from the given baseURL using the provided API key.
// The baseURL should include scheme and host, for example "http://localhost:6767".
// A non-200 status code results in an error.
func FetchSettings(baseURL, apiKey string) (Settings, error) {
	url := fmt.Sprintf("%s/api/system/settings", baseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", apiKey)

	resp, err := http.DefaultClient.Do(req)
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
