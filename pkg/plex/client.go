// file: pkg/plex/client.go
package plex

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client provides minimal access to the Plex HTTP API.
// BaseURL should include protocol and host (e.g. http://localhost:32400).
// Token is the user's Plex API token.
type Client struct {
	BaseURL string
	Token   string
	client  *http.Client
}

// NewClient returns a configured Plex API client.
func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL: strings.TrimRight(baseURL, "/"),
		Token:   token,
		client:  &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *Client) buildURL(path string) string {
	if strings.Contains(path, "?") {
		return fmt.Sprintf("%s%s&X-Plex-Token=%s", c.BaseURL, path, url.QueryEscape(c.Token))
	}
	return fmt.Sprintf("%s%s?X-Plex-Token=%s", c.BaseURL, path, url.QueryEscape(c.Token))
}

// LibraryItem represents a single media file in Plex.
type LibraryItem struct {
	Path    string // absolute file path
	Title   string // movie or show title
	Season  int    // season number if episode
	Episode int    // episode number if episode
	Key     string // Plex ratingKey
}

type mediaContainer struct {
	MediaContainer struct {
		Metadata []struct {
			Type             string `json:"type"`
			Title            string `json:"title"`
			RatingKey        string `json:"ratingKey"`
			GrandparentTitle string `json:"grandparentTitle"`
			ParentIndex      int    `json:"parentIndex"`
			Index            int    `json:"index"`
			Media            []struct {
				Part []struct {
					File string `json:"file"`
				} `json:"Part"`
			} `json:"Media"`
		} `json:"Metadata"`
	} `json:"MediaContainer"`
}

// AllItems retrieves all movies and episodes from the Plex library.
func (c *Client) AllItems(ctx context.Context) ([]LibraryItem, error) {
	u := c.buildURL("/library/sections/all")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	var mc mediaContainer
	if err := json.NewDecoder(resp.Body).Decode(&mc); err != nil {
		return nil, err
	}
	var items []LibraryItem
	for _, m := range mc.MediaContainer.Metadata {
		if len(m.Media) == 0 || len(m.Media[0].Part) == 0 {
			continue
		}
		it := LibraryItem{Path: m.Media[0].Part[0].File, Key: m.RatingKey}
		if m.Type == "episode" {
			it.Title = m.GrandparentTitle
			it.Season = m.ParentIndex
			it.Episode = m.Index
		} else {
			it.Title = m.Title
		}
		items = append(items, it)
	}
	return items, nil
}

// RefreshLibrary triggers a full Plex library refresh.
func (c *Client) RefreshLibrary(ctx context.Context) error {
	u := c.buildURL("/library/sections/all/refresh")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	return nil
}

// RefreshItem triggers a scan for a single library item identified by ratingKey.
func (c *Client) RefreshItem(ctx context.Context, key string) error {
	path := fmt.Sprintf("/library/metadata/%s/refresh", url.PathEscape(key))
	u := c.buildURL(path)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	return nil
}
