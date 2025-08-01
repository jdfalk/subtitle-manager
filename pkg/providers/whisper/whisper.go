// file: pkg/providers/whisper/whisper.go
package whisper

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// Client implements the providers.Provider interface for Whisper.
// It performs a simple HTTP GET to download subtitles.
type Client struct {
	// APIURL is the base URL of the Whisper API.
	APIURL string
	// HTTPClient is used to make requests.
	HTTPClient *http.Client
}

// New returns a Client configured with reasonable defaults.
func New() *Client {
	return &Client{
		APIURL:     viper.GetString("providers.whisper.api_url"),
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
