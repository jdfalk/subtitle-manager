// file: pkg/providers/greeksubs/greeksubs.go
package greeksubs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"
)

type Client struct {
	APIURL     string
	HTTPClient *http.Client
}

func New() *Client {
	return &Client{
		APIURL:     "https://api.greeksubs.com",
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
	}
}

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
