// file: pkg/webhooks/webhooks.go
package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/jdfalk/subtitle-manager/pkg/scanner"
	"github.com/jdfalk/subtitle-manager/pkg/security"
)

// Dispatcher sends webhook events to a list of URLs.
type Dispatcher struct {
	URLs   []string
	client *http.Client
}

// New returns a new Dispatcher with the given URLs after validating them.
func New(urls []string) (*Dispatcher, error) {
	// Validate all webhook URLs to prevent SSRF attacks
	for _, url := range urls {
		if err := validateWebhookURL(url); err != nil {
			return nil, fmt.Errorf("invalid webhook URL %s: %v", url, err)
		}
	}

	return &Dispatcher{URLs: urls, client: &http.Client{Timeout: 10 * time.Second}}, nil
}

// Send delivers an event with optional payload to all configured URLs.
func (d *Dispatcher) Send(ctx context.Context, event string, payload any) error {
	body, err := json.Marshal(map[string]any{"event": event, "payload": payload})
	if err != nil {
		return err
	}
	for _, u := range d.URLs {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(body))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		if _, err := d.client.Do(req); err != nil {
			return err
		}
	}
	return nil
}

// event describes a webhook payload with the file path and subtitle parameters.
type event struct {
	Path     string `json:"path"`
	Lang     string `json:"lang"`
	Provider string `json:"provider"`
}

// handle processes a webhook event by fetching a subtitle for the provided file.
func handle(w http.ResponseWriter, r *http.Request, ev event) {
	logger := logging.GetLogger("webhook")

	// Validate required fields
	if ev.Path == "" || ev.Lang == "" {
		logger.Warn("missing required fields in webhook event")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate and sanitize the file path
	if _, err := security.ValidateAndSanitizePath(ev.Path); err != nil {
		logger.Warnf("invalid file path in webhook: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate language code
	if err := security.ValidateLanguageCode(ev.Lang); err != nil {
		logger.Warnf("invalid language code in webhook: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate provider name if specified
	if err := security.ValidateProviderName(ev.Provider); err != nil {
		logger.Warnf("invalid provider name in webhook: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var p providers.Provider
	var name string
	var err error
	if ev.Provider != "" {
		p, err = providers.Get(ev.Provider, "")
		name = ev.Provider
		if err != nil {
			logger.Warnf("failed to get provider %s: %v", ev.Provider, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if err := scanner.ProcessFile(r.Context(), ev.Path, ev.Lang, name, p, true, nil); err != nil {
		logger.Warnf("process %s: %v", ev.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// SonarrHandler handles webhook events from Sonarr.
func SonarrHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var ev event
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handle(w, r, ev)
	})
}

// RadarrHandler handles webhook events from Radarr.
func RadarrHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var ev event
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handle(w, r, ev)
	})
}

// CustomHandler accepts generic webhook events with the same payload format.
func CustomHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var ev event
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handle(w, r, ev)
	})
}

// validateWebhookURL validates that a webhook URL is safe to use and prevents SSRF attacks
func validateWebhookURL(rawURL string) error {
	if rawURL == "" {
		return nil // Empty URLs are allowed (feature disabled)
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %v", err)
	}

	// Only allow HTTPS for webhooks (security best practice)
	if parsedURL.Scheme != "https" {
		return fmt.Errorf("only HTTPS URLs are allowed for webhooks")
	}

	// Block private/internal IP ranges and localhost
	host := parsedURL.Hostname()
	if isPrivateOrLocalhost(host) {
		return fmt.Errorf("webhooks to private/internal addresses are not allowed")
	}

	return nil
}

// isPrivateOrLocalhost checks if a hostname is a private IP or localhost
func isPrivateOrLocalhost(host string) bool {
	// Check for localhost variations
	if host == "localhost" || host == "127.0.0.1" || host == "::1" {
		return true
	}

	// Check for private IP ranges (simplified check)
	privatePatterns := []string{
		"10.",
		"192.168.",
		"172.16.", "172.17.", "172.18.", "172.19.", "172.20.",
		"172.21.", "172.22.", "172.23.", "172.24.", "172.25.",
		"172.26.", "172.27.", "172.28.", "172.29.", "172.30.", "172.31.",
		"169.254.", // Link-local
	}

	for _, pattern := range privatePatterns {
		if strings.HasPrefix(host, pattern) {
			return true
		}
	}

	return false
}
