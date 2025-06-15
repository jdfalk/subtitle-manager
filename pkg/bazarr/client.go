package bazarr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Settings represents a subset of Bazarr's configuration used for import.
type Settings map[string]any

// validateBaseURL validates that the baseURL is safe for making requests.
// It prevents SSRF attacks while allowing legitimate Bazarr connections.
func validateBaseURL(baseURL string) error {
	// Parse the URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %v", err)
	}

	// Only allow HTTP and HTTPS schemes
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("only HTTP and HTTPS schemes are allowed, got: %s", u.Scheme)
	}

	// Validate hostname is not empty
	if u.Hostname() == "" {
		return fmt.Errorf("hostname cannot be empty")
	}

	// Prevent requests to sensitive local services (but allow general localhost for Bazarr)
	hostname := strings.ToLower(u.Hostname())

	// Block cloud metadata services and other dangerous endpoints
	blockedHosts := []string{
		"169.254.169.254",          // AWS metadata service
		"metadata.google.internal", // GCP metadata service
		"metadata",                 // Generic cloud metadata
	}

	for _, blocked := range blockedHosts {
		if hostname == blocked {
			return fmt.Errorf("hostname %s is not allowed", hostname)
		}
	}

	// Validate port is in reasonable range for Bazarr (typically 6767, but allow common ports)
	if u.Port() != "" {
		// This is a basic check - Bazarr could run on various ports
		// We mainly want to prevent port 22 (SSH), 3389 (RDP), etc.
		blockedPorts := []string{"22", "23", "3389", "5900", "5901"}
		port := u.Port()
		for _, blocked := range blockedPorts {
			if port == blocked {
				return fmt.Errorf("port %s is not allowed", port)
			}
		}
	}

	// Ensure the URL path starts with /
	if !strings.HasPrefix(u.Path, "/") {
		u.Path = "/" + u.Path
	}

	return nil
}

// FetchSettings retrieves Bazarr settings from the given baseURL using the provided API key.
// The baseURL should include scheme and host, for example "http://localhost:6767".
// A non-200 status code results in an error.
// This function validates the baseURL to prevent SSRF attacks.
func FetchSettings(baseURL, apiKey string) (Settings, error) {
	// Validate the base URL to prevent SSRF attacks
	if err := validateBaseURL(baseURL); err != nil {
		return nil, fmt.Errorf("invalid baseURL: %v", err)
	}

	url := fmt.Sprintf("%s/api/system/settings", baseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
