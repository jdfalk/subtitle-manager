// file: sdks/go/subtitleclient/client_test.go
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440027

package subtitleclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	config := Config{
		BaseURL:   "http://test.example.com",
		APIKey:    "test-key",
		Timeout:   60 * time.Second,
		UserAgent: "test-agent",
	}

	client := NewClient(config)
	assert.Equal(t, "http://test.example.com", client.baseURL)
	assert.Equal(t, "test-key", client.apiKey)
	assert.Equal(t, "test-agent", client.userAgent)
	assert.Equal(t, 60*time.Second, client.httpClient.Timeout)
}

func TestNewDefaultClient(t *testing.T) {
	client := NewDefaultClient("http://test.example.com", "test-key")
	assert.Equal(t, "http://test.example.com", client.baseURL)
	assert.Equal(t, "test-key", client.apiKey)
	assert.Equal(t, "subtitle-manager-go-sdk/1.0.0", client.userAgent)
}

func setupTestServer(t *testing.T, handler http.HandlerFunc) (*httptest.Server, *Client) {
	server := httptest.NewServer(handler)
	client := NewDefaultClient(server.URL, "test-api-key")
	return server, client
}

func TestGetSystemInfo(t *testing.T) {
	systemInfo := SystemInfo{
		GoVersion:  "go1.24.0",
		OS:         "linux",
		Arch:       "amd64",
		Goroutines: 10,
		DiskFree:   1000000,
		DiskTotal:  2000000,
	}

	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/system", r.URL.Path)
		assert.Equal(t, "test-api-key", r.Header.Get("X-API-Key"))
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(systemInfo)
	})
	defer server.Close()

	result, err := client.GetSystemInfo(context.Background())
	require.NoError(t, err)
	assert.Equal(t, systemInfo.GoVersion, result.GoVersion)
	assert.Equal(t, systemInfo.OS, result.OS)
	assert.Equal(t, systemInfo.Arch, result.Arch)
	assert.Equal(t, systemInfo.Goroutines, result.Goroutines)
	assert.Equal(t, 50.0, result.DiskUsagePercent())
}

func TestLogin(t *testing.T) {
	loginResponse := LoginResponse{
		UserID:   1,
		Username: "testuser",
		Role:     "admin",
	}

	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/login", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Equal(t, "testuser", req.Username)
		assert.Equal(t, "password", req.Password)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(loginResponse)
	})
	defer server.Close()

	result, err := client.Login(context.Background(), "testuser", "password")
	require.NoError(t, err)
	assert.Equal(t, loginResponse.UserID, result.UserID)
	assert.Equal(t, loginResponse.Username, result.Username)
	assert.Equal(t, loginResponse.Role, result.Role)
	assert.True(t, result.IsAdmin())
	assert.True(t, result.HasBasicAccess())
	assert.True(t, result.HasReadAccess())
}

func TestGetLogs(t *testing.T) {
	logs := []LogEntry{
		{
			Timestamp: time.Now(),
			Level:     "info",
			Component: "webserver",
			Message:   "Server started",
			Fields:    map[string]interface{}{"port": 8080},
		},
		{
			Timestamp: time.Now(),
			Level:     "error",
			Component: "auth",
			Message:   "Login failed",
			Fields:    map[string]interface{}{"username": "baduser"},
		},
	}

	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/logs", r.URL.Path)
		assert.Equal(t, "level=info", r.URL.RawQuery)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(logs)
	})
	defer server.Close()

	result, err := client.GetLogs(context.Background(), LogParams{
		Level: "info",
		Limit: 50,
	})
	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "webserver", result[0].Component)
	assert.False(t, result[0].IsError())
	assert.True(t, result[1].IsError())
}

func TestConvertSubtitle(t *testing.T) {
	srtContent := "1\n00:00:01,000 --> 00:00:03,000\nHello World\n\n"

	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/convert", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		assert.Contains(t, r.Header.Get("Content-Type"), "multipart/form-data")

		// Parse multipart form
		err := r.ParseMultipartForm(10 << 20) // 10MB
		require.NoError(t, err)

		file, header, err := r.FormFile("file")
		require.NoError(t, err)
		defer file.Close()

		assert.Equal(t, "test.vtt", header.Filename)

		w.Header().Set("Content-Type", "application/x-subrip")
		w.Write([]byte(srtContent))
	})
	defer server.Close()

	inputContent := "WEBVTT\n\n00:01.000 --> 00:03.000\nHello World"
	result, err := client.ConvertSubtitle(context.Background(), "test.vtt", strings.NewReader(inputContent))
	require.NoError(t, err)
	assert.Equal(t, srtContent, string(result))
}

func TestDownloadSubtitles(t *testing.T) {
	downloadResult := DownloadResult{
		Success:      true,
		SubtitlePath: stringPtr("/path/to/movie.en.srt"),
		Provider:     stringPtr("opensubtitles"),
	}

	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/download", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var req DownloadRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Equal(t, "/movies/example.mkv", req.Path)
		assert.Equal(t, "en", req.Language)
		assert.Equal(t, []string{"opensubtitles", "subscene"}, req.Providers)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(downloadResult)
	})
	defer server.Close()

	result, err := client.DownloadSubtitles(context.Background(), DownloadRequest{
		Path:      "/movies/example.mkv",
		Language:  "en",
		Providers: []string{"opensubtitles", "subscene"},
	})
	require.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, "/path/to/movie.en.srt", *result.SubtitlePath)
	assert.Equal(t, "opensubtitles", *result.Provider)
}

func TestGetHistory(t *testing.T) {
	historyResponse := HistoryResponse{
		Items: []HistoryItem{
			{
				ID:           1,
				Type:         "download",
				FilePath:     "/movies/example.mkv",
				SubtitlePath: stringPtr("/movies/example.en.srt"),
				Language:     stringPtr("en"),
				Provider:     stringPtr("opensubtitles"),
				Status:       "success",
				CreatedAt:    time.Now(),
				UserID:       1,
			},
			{
				ID:           2,
				Type:         "convert",
				FilePath:     "/subtitles/test.vtt",
				SubtitlePath: stringPtr("/subtitles/test.srt"),
				Status:       "success",
				CreatedAt:    time.Now(),
				UserID:       1,
			},
		},
		Total: 2,
		Page:  1,
		Limit: 20,
	}

	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/history", r.URL.Path)
		assert.Contains(t, r.URL.RawQuery, "page=1")
		assert.Contains(t, r.URL.RawQuery, "limit=20")
		assert.Contains(t, r.URL.RawQuery, "type=download")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(historyResponse)
	})
	defer server.Close()

	result, err := client.GetHistory(context.Background(), HistoryParams{
		Page:  1,
		Limit: 20,
		Type:  "download",
	})
	require.NoError(t, err)
	assert.Equal(t, 2, result.Total)
	assert.Equal(t, 1, result.Page)
	assert.Len(t, result.Items, 2)
	assert.False(t, result.HasNextPage())
	assert.False(t, result.HasPreviousPage())
	assert.Equal(t, 1, result.TotalPages())

	firstItem := result.Items[0]
	assert.Equal(t, "download", firstItem.Type)
	assert.True(t, firstItem.IsSuccess())
	assert.False(t, firstItem.IsFailed())
	assert.Equal(t, "opensubtitles", *firstItem.Provider)
}

func TestGetScanStatus(t *testing.T) {
	scanStatus := ScanStatus{
		Scanning:       true,
		Progress:       0.75,
		CurrentPath:    stringPtr("/movies/subfolder"),
		FilesProcessed: intPtr(150),
		FilesTotal:     intPtr(200),
		StartTime:      timePtr(time.Now()),
	}

	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/scan/status", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(scanStatus)
	})
	defer server.Close()

	result, err := client.GetScanStatus(context.Background())
	require.NoError(t, err)
	assert.True(t, result.Scanning)
	assert.Equal(t, 75.0, result.ProgressPercent())
	assert.Equal(t, 150, *result.FilesProcessed)
	assert.Equal(t, 50, result.RemainingFiles())
}

func TestAPIError(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "unauthorized",
			"message": "Authentication required",
		})
	})
	defer server.Close()

	_, err := client.GetSystemInfo(context.Background())
	require.Error(t, err)

	apiErr, ok := err.(*APIError)
	require.True(t, ok)
	assert.Equal(t, 401, apiErr.StatusCode)
	assert.Equal(t, "unauthorized", apiErr.Code)
	assert.Equal(t, "Authentication required", apiErr.Message)
	assert.True(t, apiErr.IsAuthenticationError())
	assert.False(t, apiErr.IsAuthorizationError())
}

func TestHealthCheck(t *testing.T) {
	t.Run("healthy", func(t *testing.T) {
		server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(SystemInfo{
				GoVersion: "go1.24.0",
				OS:        "linux",
				Arch:      "amd64",
			})
		})
		defer server.Close()

		healthy := client.HealthCheck(context.Background())
		assert.True(t, healthy)
	})

	t.Run("unhealthy", func(t *testing.T) {
		server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})
		defer server.Close()

		healthy := client.HealthCheck(context.Background())
		assert.False(t, healthy)
	})
}

func TestHistoryIterator(t *testing.T) {
	// Mock server that returns different pages
	page := 0
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		page++

		var response HistoryResponse
		if page == 1 {
			response = HistoryResponse{
				Items: []HistoryItem{
					{ID: 1, Type: "download", FilePath: "/movie1.mkv", Status: "success"},
					{ID: 2, Type: "download", FilePath: "/movie2.mkv", Status: "success"},
				},
				Total: 3,
				Page:  1,
				Limit: 2,
			}
		} else {
			response = HistoryResponse{
				Items: []HistoryItem{
					{ID: 3, Type: "download", FilePath: "/movie3.mkv", Status: "success"},
				},
				Total: 3,
				Page:  2,
				Limit: 2,
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	defer server.Close()

	iterator := client.GetHistoryIterator(context.Background(), HistoryParams{Limit: 2})

	var items []HistoryItem
	for iterator.Next(context.Background()) {
		items = append(items, *iterator.Item())
	}

	require.NoError(t, iterator.Err())
	assert.Len(t, items, 3)
	assert.Equal(t, int64(1), items[0].ID)
	assert.Equal(t, int64(2), items[1].ID)
	assert.Equal(t, int64(3), items[2].ID)
	assert.Equal(t, 3, iterator.Total())
}

func TestRateLimiting(t *testing.T) {
	// Create client with very low rate limit for testing
	client := NewClient(Config{
		BaseURL:   "http://test.example.com",
		APIKey:    "test-key",
		RateLimit: 1, // 1 request per second
	})

	// Mock server
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SystemInfo{GoVersion: "go1.24.0"})
	}))
	defer server.Close()

	client.baseURL = server.URL

	// Make multiple requests quickly
	start := time.Now()
	ctx := context.Background()

	_, err := client.GetSystemInfo(ctx)
	require.NoError(t, err)

	_, err = client.GetSystemInfo(ctx)
	require.NoError(t, err)

	elapsed := time.Since(start)

	// Second request should be delayed by rate limiter
	assert.GreaterOrEqual(t, elapsed, time.Second)
	assert.Equal(t, 2, requestCount)
}

// Helper functions for creating pointers to basic types
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func timePtr(t time.Time) *time.Time {
	return &t
}
