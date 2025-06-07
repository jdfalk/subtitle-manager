// file: pkg/providers/embedded/embedded.go
package embedded

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"
)

// Client implements the providers.Provider interface for Embedded.
// It performs a simple HTTP GET to download subtitles.
type Client struct {
	// APIURL is the base URL of the Embedded API.
	APIURL string
	// HTTPClient is used to make requests.
	HTTPClient *http.Client
}

// New returns a Client configured with reasonable defaults.
func New() *Client {
	return &Client{
		APIURL:     "https://api.embedded.com",
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
	}
}

// Fetch downloads the subtitle for mediaPath in lang.
// It returns the subtitle bytes or an error.
func (c *Client) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
	name := filepath.Base(mediaPath)
	url := fmt.Sprintf("%s/subtitles/%s/%s", c.APIURL, name, lang)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
