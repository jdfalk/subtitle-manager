// file: pkg/providers/opensubtitles/opensubtitles.go
package opensubtitles

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// LoginResponse represents the response from the OpenSubtitles login API
type LoginResponse struct {
	User struct {
		AllowedTranslations int    `json:"allowed_translations"`
		AllowedDownloads    int    `json:"allowed_downloads"`
		Level              string `json:"level"`
		UserID             int    `json:"user_id"`
		ExtInstalled       bool   `json:"ext_installed"`
		VIP                bool   `json:"vip"`
	} `json:"user"`
	BaseURL string `json:"base_url"`
	Token   string `json:"token"`
	Status  string `json:"status"`
}

// SearchResult represents a subtitle search result
type SearchResult struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		SubtitleID       string `json:"subtitle_id"`
		Language         string `json:"language"`
		DownloadCount    int    `json:"download_count"`
		NewDownloadCount int    `json:"new_download_count"`
		HearingImpaired  bool   `json:"hearing_impaired"`
		HD               bool   `json:"hd"`
		FPS              float64 `json:"fps"`
		Votes            int    `json:"votes"`
		Ratings          float64 `json:"ratings"`
		FromTrusted      bool   `json:"from_trusted"`
		ForeignPartsOnly bool   `json:"foreign_parts_only"`
		AutoTranslated   bool   `json:"auto_translated"`
		MachineTranslated bool  `json:"machine_translated"`
		UploadDate       string `json:"upload_date"`
		Release          string `json:"release"`
		Comments         string `json:"comments"`
		LegacySubtitleID int    `json:"legacy_subtitle_id"`
		Uploader         struct {
			UploaderID int    `json:"uploader_id"`
			Name       string `json:"name"`
			Rank       string `json:"rank"`
		} `json:"uploader"`
		FeatureDetails struct {
			FeatureID    int    `json:"feature_id"`
			FeatureType  string `json:"feature_type"`
			Year         int    `json:"year"`
			Title        string `json:"title"`
			MovieName    string `json:"movie_name"`
			ImdbID       int    `json:"imdb_id"`
			TmdbID       int    `json:"tmdb_id"`
		} `json:"feature_details"`
		URL      string `json:"url"`
		RelatedLinks struct {
			Label       string `json:"label"`
			URL         string `json:"url"`
			ImgURL      string `json:"img_url"`
		} `json:"related_links"`
		Files []struct {
			FileID   int    `json:"file_id"`
			CDNumber int    `json:"cd_number"`
			FileName string `json:"file_name"`
		} `json:"files"`
	} `json:"attributes"`
}

// SearchResponse represents the response from the search API
type SearchResponse struct {
	TotalPages int            `json:"total_pages"`
	TotalCount int            `json:"total_count"`
	PerPage    int            `json:"per_page"`
	Page       int            `json:"page"`
	Data       []SearchResult `json:"data"`
}

// DownloadResponse represents the download link response
type DownloadResponse struct {
	Link       string `json:"link"`
	FileName   string `json:"file_name"`
	Requests   int    `json:"requests"`
	Remaining  int    `json:"remaining"`
	Message    string `json:"message"`
	ResetTime  string `json:"reset_time"`
}

// Client implements the providers.Provider interface for OpenSubtitles.
type Client struct {
	// APIURL allows overriding the REST endpoint, mainly for testing.
	APIURL string
	// UserAgent identifies this application to the OpenSubtitles API.
	UserAgent  string
	HTTPClient *http.Client

	// Authentication
	username string
	password string
	token    string
	tokenMu  sync.RWMutex
	tokenExp time.Time
}

// New returns a new Client configured with username/password from viper config.
func New(_ string) *Client {
	apiURL := viper.GetString("opensubtitles.api_url")
	if apiURL == "" {
		apiURL = "https://api.opensubtitles.com/api/v1"
	}
	ua := viper.GetString("opensubtitles.user_agent")
	if ua == "" {
		ua = "subtitle-manager v1.0"
	}

	username := viper.GetString("opensubtitles.username")
	password := viper.GetString("opensubtitles.password")

	return &Client{
		APIURL:     apiURL,
		UserAgent:  ua,
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
		username:   username,
		password:   password,
	}
}

// login authenticates with OpenSubtitles and stores the session token
func (c *Client) login(ctx context.Context) error {
	if c.username == "" || c.password == "" {
		return fmt.Errorf("OpenSubtitles username and password not configured")
	}

	loginData := map[string]string{
		"username": c.username,
		"password": c.password,
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return fmt.Errorf("failed to marshal login data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.APIURL+"/login", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create login request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Api-Key", "") // Empty API key for username/password login

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("login request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("login failed with status %d: %s", resp.StatusCode, string(body))
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return fmt.Errorf("failed to decode login response: %w", err)
	}

	c.tokenMu.Lock()
	c.token = loginResp.Token
	c.tokenExp = time.Now().Add(24 * time.Hour) // Tokens typically expire in 24 hours
	c.tokenMu.Unlock()

	return nil
}

// getToken returns a valid authentication token, logging in if necessary
func (c *Client) getToken(ctx context.Context) (string, error) {
	c.tokenMu.RLock()
	if c.token != "" && time.Now().Before(c.tokenExp) {
		token := c.token
		c.tokenMu.RUnlock()
		return token, nil
	}
	c.tokenMu.RUnlock()

	// Need to login or refresh token
	if err := c.login(ctx); err != nil {
		return "", err
	}

	c.tokenMu.RLock()
	token := c.token
	c.tokenMu.RUnlock()

	return token, nil
}

// Search returns download URLs for matching subtitles without downloading them.
func (c *Client) Search(ctx context.Context, mediaPath, lang string) ([]string, error) {
	token, err := c.getToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	hash, size, err := fileHashFunc(mediaPath)
	if err != nil {
		return nil, err
	}

	// Use the new REST API v1 endpoint
	url := fmt.Sprintf("%s/subtitles?moviehash=%x&moviebytesize=%d&languages=%s", c.APIURL, hash, size, lang)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		// Token might be expired, try to login again
		c.tokenMu.Lock()
		c.token = ""
		c.tokenMu.Unlock()

		token, err = c.getToken(ctx)
		if err != nil {
			return nil, fmt.Errorf("re-authentication failed: %w", err)
		}

		// Retry the request with new token
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("search failed with status %d: %s", resp.StatusCode, string(body))
	}

	var searchResp SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	urls := make([]string, 0, len(searchResp.Data))
	for _, result := range searchResp.Data {
		if result.Attributes.SubtitleID != "" {
			// We'll need to get download links for each subtitle
			downloadURL := fmt.Sprintf("%s/download", c.APIURL)
			urls = append(urls, downloadURL+"?file_id="+result.Attributes.SubtitleID)
		}
	}

	return urls, nil
}

// Fetch downloads the first matching subtitle for mediaPath in lang.
func (c *Client) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
	token, err := c.getToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	// First search for subtitles
	hash, size, err := fileHashFunc(mediaPath)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/subtitles?moviehash=%x&moviebytesize=%d&languages=%s", c.APIURL, hash, size, lang)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("search failed with status %d: %s", resp.StatusCode, string(body))
	}

	var searchResp SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	if len(searchResp.Data) == 0 {
		return nil, fmt.Errorf("no subtitles found")
	}

	// Get download link for the first result
	firstResult := searchResp.Data[0]
	if len(firstResult.Attributes.Files) == 0 {
		return nil, fmt.Errorf("no files available for subtitle")
	}

	fileID := firstResult.Attributes.Files[0].FileID
	downloadReq := map[string]interface{}{
		"file_id": fileID,
	}

	downloadData, err := json.Marshal(downloadReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal download request: %w", err)
	}

	dlReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.APIURL+"/download", bytes.NewBuffer(downloadData))
	if err != nil {
		return nil, err
	}

	dlReq.Header.Set("User-Agent", c.UserAgent)
	dlReq.Header.Set("Authorization", "Bearer "+token)
	dlReq.Header.Set("Content-Type", "application/json")

	dlResp, err := c.HTTPClient.Do(dlReq)
	if err != nil {
		return nil, err
	}
	defer dlResp.Body.Close()

	if dlResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(dlResp.Body)
		return nil, fmt.Errorf("download request failed with status %d: %s", dlResp.StatusCode, string(body))
	}

	var downloadResp DownloadResponse
	if err := json.NewDecoder(dlResp.Body).Decode(&downloadResp); err != nil {
		return nil, fmt.Errorf("failed to decode download response: %w", err)
	}

	// Now download the actual subtitle file
	fileReq, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadResp.Link, nil)
	if err != nil {
		return nil, err
	}

	fileReq.Header.Set("User-Agent", c.UserAgent)

	fileResp, err := c.HTTPClient.Do(fileReq)
	if err != nil {
		return nil, err
	}
	defer fileResp.Body.Close()

	if fileResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("file download failed with status %d", fileResp.StatusCode)
	}

	return io.ReadAll(fileResp.Body)
}
