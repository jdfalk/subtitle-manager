// file: pkg/providers/opensubtitles/opensubtitles.go
package opensubtitles

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/viper"
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
	apiURL := viper.GetString("opensubtitles.api_url")
	if apiURL == "" {
		apiURL = "https://rest.opensubtitles.org"
	}
	ua := viper.GetString("opensubtitles.user_agent")
	if ua == "" {
		ua = "github.com/jdfalk/subtitle-manager/0.1"
	}
	return &Client{
		APIURL:     apiURL,
		UserAgent:  ua,
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
		APIKey:     apiKey,
	}
}

// Search returns download URLs for matching subtitles without downloading them.
func (c *Client) Search(ctx context.Context, mediaPath, lang string) ([]string, error) {
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
	urls := make([]string, len(results))
	for i, r := range results {
		urls[i] = r.SubDownloadLink
	}
	return urls, nil
}

// Fetch downloads the first matching subtitle for mediaPath in lang.
func (c *Client) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
	urls, err := c.Search(ctx, mediaPath, lang)
	if err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, fmt.Errorf("no subtitles found")
	}
	dlReq, err := http.NewRequestWithContext(ctx, http.MethodGet, urls[0], nil)
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
