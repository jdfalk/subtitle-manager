// file: pkg/providers/opensubtitles/opensubtitles.go
package opensubtitles

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client implements the providers.Provider interface for OpenSubtitles.
type Client struct {
	// APIURL allows overriding the REST endpoint, mainly for testing.
	APIURL string
	// UserAgent identifies this application to the OpenSubtitles API.
	UserAgent  string
	HTTPClient *http.Client
	APIKey     string
}

// New returns a new Client with the provided API key.
func New(apiKey string) *Client {
	return &Client{
		APIURL:     "https://rest.opensubtitles.org",
		UserAgent:  "subtitle-manager/0.1",
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
		APIKey:     apiKey,
	}
}

// Fetch downloads the first matching subtitle for mediaPath in lang.
func (c *Client) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
	hash, size, err := fileHashFunc(mediaPath)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/search/moviehash-%x/moviebytesize-%d/sublanguageid-%s", c.APIURL, hash, size, lang)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	if c.APIKey != "" {
		req.Header.Set("Api-Key", c.APIKey)
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search status %d", resp.StatusCode)
	}
	var results []struct {
		SubDownloadLink string `json:"SubDownloadLink"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no subtitles found")
	}
	dlReq, err := http.NewRequestWithContext(ctx, http.MethodGet, results[0].SubDownloadLink, nil)
	if err != nil {
		return nil, err
	}
	dlReq.Header.Set("User-Agent", c.UserAgent)
	dlResp, err := c.HTTPClient.Do(dlReq)
	if err != nil {
		return nil, err
	}
	defer dlResp.Body.Close()
	if dlResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download status %d", dlResp.StatusCode)
	}
	return io.ReadAll(dlResp.Body)
}
