// file: sdks/go/subtitleclient/client.go
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440025

// Package subtitleclient provides a Go SDK for the Subtitle Manager API.
//
// This package offers type-safe access to all API endpoints with automatic retry,
// error handling, pagination support, and Go-idiomatic patterns.
package subtitleclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

// Client provides access to the Subtitle Manager API.
type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
	userAgent  string
	limiter    *rate.Limiter
	maxRetries int
}

// Config holds configuration options for the client.
type Config struct {
	BaseURL    string        // API base URL
	APIKey     string        // API key for authentication  
	Timeout    time.Duration // HTTP request timeout
	MaxRetries int           // Maximum retry attempts
	UserAgent  string        // Custom user agent
	RateLimit  rate.Limit    // Rate limit (requests per second)
}

// NewClient creates a new Subtitle Manager API client.
func NewClient(config Config) *Client {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}
	if config.UserAgent == "" {
		config.UserAgent = "subtitle-manager-go-sdk/1.0.0"
	}
	if config.RateLimit == 0 {
		config.RateLimit = rate.Limit(10) // 10 requests per second
	}
	if config.APIKey == "" {
		config.APIKey = os.Getenv("SUBTITLE_MANAGER_API_KEY")
	}

	return &Client{
		baseURL: strings.TrimSuffix(config.BaseURL, "/"),
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		apiKey:     config.APIKey,
		userAgent:  config.UserAgent,
		limiter:    rate.NewLimiter(config.RateLimit, 1),
		maxRetries: config.MaxRetries,
	}
}

// NewDefaultClient creates a client with default configuration.
func NewDefaultClient(baseURL, apiKey string) *Client {
	return NewClient(Config{
		BaseURL: baseURL,
		APIKey:  apiKey,
	})
}

// doRequest performs an HTTP request with retry logic and rate limiting.
func (c *Client) doRequest(ctx context.Context, method, path string, body io.Reader, headers map[string]string) (*http.Response, error) {
	// Wait for rate limiter
	if err := c.limiter.Wait(ctx); err != nil {
		return nil, err
	}

	url := c.baseURL + path
	var lastErr error

	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, method, url, body)
		if err != nil {
			return nil, fmt.Errorf("creating request: %w", err)
		}

		// Set default headers
		req.Header.Set("User-Agent", c.userAgent)
		req.Header.Set("Accept", "application/json")
		
		if c.apiKey != "" {
			req.Header.Set("X-API-Key", c.apiKey)
		}

		// Set custom headers
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				time.Sleep(time.Duration(attempt+1) * time.Second)
				continue
			}
			return nil, fmt.Errorf("request failed after %d attempts: %w", c.maxRetries+1, err)
		}

		// Check for retryable status codes
		if resp.StatusCode >= 500 || resp.StatusCode == 429 {
			resp.Body.Close()
			lastErr = fmt.Errorf("HTTP %d", resp.StatusCode)
			if attempt < c.maxRetries {
				backoff := time.Duration(attempt+1) * time.Second
				if resp.StatusCode == 429 {
					if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
						if seconds, err := strconv.Atoi(retryAfter); err == nil {
							backoff = time.Duration(seconds) * time.Second
						}
					}
				}
				time.Sleep(backoff)
				continue
			}
		}

		return resp, nil
	}

	return nil, lastErr
}

// handleAPIError processes API error responses and returns appropriate errors.
func (c *Client) handleAPIError(resp *http.Response) error {
	defer resp.Body.Close()

	var apiErr APIError
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("HTTP %d", resp.StatusCode),
		}
	}

	apiErr.StatusCode = resp.StatusCode
	return &apiErr
}

// get performs a GET request and unmarshals the response into result.
func (c *Client) get(ctx context.Context, path string, params url.Values, result interface{}) error {
	if params != nil && len(params) > 0 {
		path += "?" + params.Encode()
	}

	resp, err := c.doRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return c.handleAPIError(resp)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

// post performs a POST request with JSON body.
func (c *Client) post(ctx context.Context, path string, body interface{}, result interface{}) error {
	var bodyReader io.Reader
	headers := make(map[string]string)

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
		headers["Content-Type"] = "application/json"
	}

	resp, err := c.doRequest(ctx, "POST", path, bodyReader, headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return c.handleAPIError(resp)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

// postFile performs a POST request with multipart file upload.
func (c *Client) postFile(ctx context.Context, path string, filename string, fileData io.Reader, fields map[string]string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add form fields
	for key, value := range fields {
		if err := writer.WriteField(key, value); err != nil {
			return nil, fmt.Errorf("writing field %s: %w", key, err)
		}
	}

	// Add file
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, fmt.Errorf("creating form file: %w", err)
	}

	if _, err := io.Copy(part, fileData); err != nil {
		return nil, fmt.Errorf("copying file data: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("closing multipart writer: %w", err)
	}

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}

	resp, err := c.doRequest(ctx, "POST", path, body, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, c.handleAPIError(resp)
	}

	return io.ReadAll(resp.Body)
}

// Authentication methods

// Login authenticates with username and password.
func (c *Client) Login(ctx context.Context, username, password string) (*LoginResponse, error) {
	body := LoginRequest{
		Username: username,
		Password: password,
	}

	var result LoginResponse
	if err := c.post(ctx, "/api/login", body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Logout invalidates the current session.
func (c *Client) Logout(ctx context.Context) error {
	return c.post(ctx, "/api/logout", nil, nil)
}

// System information methods

// GetSystemInfo retrieves system information.
func (c *Client) GetSystemInfo(ctx context.Context) (*SystemInfo, error) {
	var result SystemInfo
	if err := c.get(ctx, "/api/system", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLogs retrieves application logs.
func (c *Client) GetLogs(ctx context.Context, params LogParams) ([]LogEntry, error) {
	values := make(url.Values)
	if params.Level != "" {
		values.Set("level", params.Level)
	}
	if params.Limit > 0 {
		values.Set("limit", strconv.Itoa(params.Limit))
	}

	var result []LogEntry
	if err := c.get(ctx, "/api/logs", values, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Configuration methods

// GetConfig retrieves application configuration.
func (c *Client) GetConfig(ctx context.Context) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := c.get(ctx, "/api/config", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Subtitle operations

// ConvertSubtitle converts a subtitle file to SRT format.
func (c *Client) ConvertSubtitle(ctx context.Context, filename string, fileData io.Reader) ([]byte, error) {
	return c.postFile(ctx, "/api/convert", filename, fileData, nil)
}

// TranslateSubtitle translates a subtitle file to the target language.
func (c *Client) TranslateSubtitle(ctx context.Context, filename string, fileData io.Reader, language string, provider string) ([]byte, error) {
	fields := map[string]string{
		"language": language,
	}
	if provider != "" {
		fields["provider"] = provider
	}

	return c.postFile(ctx, "/api/translate", filename, fileData, fields)
}

// ExtractSubtitles extracts embedded subtitles from a video file.
func (c *Client) ExtractSubtitles(ctx context.Context, filename string, fileData io.Reader, language string, track int) ([]byte, error) {
	fields := make(map[string]string)
	if language != "" {
		fields["language"] = language
	}
	if track > 0 {
		fields["track"] = strconv.Itoa(track)
	}

	return c.postFile(ctx, "/api/extract", filename, fileData, fields)
}

// Download and scanning

// DownloadSubtitles downloads subtitles for a media file.
func (c *Client) DownloadSubtitles(ctx context.Context, req DownloadRequest) (*DownloadResult, error) {
	var result DownloadResult
	if err := c.post(ctx, "/api/download", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// StartLibraryScan starts a library scan.
func (c *Client) StartLibraryScan(ctx context.Context, req ScanRequest) (*ScanResult, error) {
	var result ScanResult
	if err := c.post(ctx, "/api/scan", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetScanStatus retrieves the current scan status.
func (c *Client) GetScanStatus(ctx context.Context) (*ScanStatus, error) {
	var result ScanStatus
	if err := c.get(ctx, "/api/scan/status", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// History methods

// GetHistory retrieves operation history with pagination.
func (c *Client) GetHistory(ctx context.Context, params HistoryParams) (*HistoryResponse, error) {
	values := make(url.Values)
	if params.Page > 0 {
		values.Set("page", strconv.Itoa(params.Page))
	}
	if params.Limit > 0 {
		values.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.Type != "" {
		values.Set("type", params.Type)
	}
	if !params.StartDate.IsZero() {
		values.Set("start_date", params.StartDate.Format(time.RFC3339))
	}
	if !params.EndDate.IsZero() {
		values.Set("end_date", params.EndDate.Format(time.RFC3339))
	}

	var result HistoryResponse
	if err := c.get(ctx, "/api/history", values, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// OAuth2 management (admin only)

// GenerateGitHubOAuth generates GitHub OAuth2 credentials.
func (c *Client) GenerateGitHubOAuth(ctx context.Context) (*OAuthCredentials, error) {
	var result OAuthCredentials
	if err := c.post(ctx, "/api/oauth/github/generate", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RegenerateGitHubOAuth regenerates GitHub OAuth2 client secret.
func (c *Client) RegenerateGitHubOAuth(ctx context.Context) (*OAuthCredentials, error) {
	var result OAuthCredentials
	if err := c.post(ctx, "/api/oauth/github/regenerate", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ResetGitHubOAuth resets GitHub OAuth2 configuration.
func (c *Client) ResetGitHubOAuth(ctx context.Context) error {
	return c.post(ctx, "/api/oauth/github/reset", nil, nil)
}

// Utility methods

// HealthCheck checks if the API is accessible.
func (c *Client) HealthCheck(ctx context.Context) bool {
	_, err := c.GetSystemInfo(ctx)
	return err == nil
}

// WaitForScanCompletion waits for a library scan to complete.
func (c *Client) WaitForScanCompletion(ctx context.Context, checkInterval time.Duration) (*ScanStatus, error) {
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			status, err := c.GetScanStatus(ctx)
			if err != nil {
				return nil, err
			}
			if !status.Scanning {
				return status, nil
			}
		}
	}
}

// HistoryIterator provides pagination over history items.
type HistoryIterator struct {
	client *Client
	params HistoryParams
	page   int
	items  []HistoryItem
	index  int
	total  int
	err    error
}

// GetHistoryIterator returns an iterator for paginating through history.
func (c *Client) GetHistoryIterator(ctx context.Context, params HistoryParams) *HistoryIterator {
	if params.Limit == 0 {
		params.Limit = 20
	}
	return &HistoryIterator{
		client: c,
		params: params,
		page:   1,
	}
}

// Next advances the iterator to the next item.
func (it *HistoryIterator) Next(ctx context.Context) bool {
	if it.err != nil {
		return false
	}

	// If we have items in current page, return the next one
	if it.index < len(it.items) {
		it.index++
		return it.index <= len(it.items)
	}

	// Need to fetch next page
	it.params.Page = it.page
	resp, err := it.client.GetHistory(ctx, it.params)
	if err != nil {
		it.err = err
		return false
	}

	it.items = resp.Items
	it.total = resp.Total
	it.index = 1
	it.page++

	return len(it.items) > 0
}

// Item returns the current history item.
func (it *HistoryIterator) Item() *HistoryItem {
	if it.index > 0 && it.index <= len(it.items) {
		return &it.items[it.index-1]
	}
	return nil
}

// Err returns any error that occurred during iteration.
func (it *HistoryIterator) Err() error {
	return it.err
}

// Total returns the total number of items across all pages.
func (it *HistoryIterator) Total() int {
	return it.total
}