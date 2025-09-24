// file: pkg/webserver/search_integration_test.go
// version: 1.3.0
// guid: a1b2c3d4-e5f6-7g8h-9i0j-1k2l3m4n5o6p

package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSearchHandlerPositiveValidation tests that search handler properly validates
// and processes well-formed requests, even when no providers find results
func TestSearchHandlerPositiveValidation(t *testing.T) {
	// Set up test environment - create a temporary test file
	testDir := t.TempDir()
	testFile := testDir + "/validation_test.mkv"
	require.NoError(t, os.WriteFile(testFile, []byte("fake video content"), 0644))

	// Set TEST_SAFE_MEDIA_DIR to allow access to our test file
	originalEnv := os.Getenv("TEST_SAFE_MEDIA_DIR")
	os.Setenv("TEST_SAFE_MEDIA_DIR", testDir)
	defer os.Setenv("TEST_SAFE_MEDIA_DIR", originalEnv)

	testCases := []struct {
		name     string
		request  SearchRequest
		wantCode int
	}{
		{
			name: "valid_basic_request",
			request: SearchRequest{
				Providers: []string{"embedded"}, // Known provider that's safe for testing
				MediaPath: "validation_test.mkv",
				Language:  "en",
			},
			wantCode: http.StatusOK, // Should process successfully even if no results
		},
		{
			name: "valid_request_with_episode_info",
			request: SearchRequest{
				Providers: []string{"embedded"},
				MediaPath: "validation_test.mkv",
				Language:  "en",
				Season:    1,
				Episode:   1,
			},
			wantCode: http.StatusOK,
		},
		{
			name: "valid_request_with_metadata",
			request: SearchRequest{
				Providers:    []string{"embedded"},
				MediaPath:    "validation_test.mkv",
				Language:     "en",
				Year:         2023,
				ReleaseGroup: "TEST",
			},
			wantCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tc.request)
			require.NoError(t, err)

			req, err := http.NewRequest("POST", "/api/search", bytes.NewBuffer(reqBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := searchHandler()
			handler.ServeHTTP(rr, req)

			require.Equal(t, tc.wantCode, rr.Code)

			// If successful, verify response structure
			if rr.Code == http.StatusOK {
				var response SearchResponse
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				require.NoError(t, err)

				// Verify basic response structure
				require.NotNil(t, response.Results)
				require.Equal(t, tc.request, response.Query)
			}
		})
	}
}

// TestSearchHandlerResponseStructure verifies the search response contains all expected fields
func TestSearchHandlerResponseStructure(t *testing.T) {
	testDir := t.TempDir()
	testFile := testDir + "/structure_test.mkv"
	require.NoError(t, os.WriteFile(testFile, []byte("test content"), 0644))

	originalEnv := os.Getenv("TEST_SAFE_MEDIA_DIR")
	os.Setenv("TEST_SAFE_MEDIA_DIR", testDir)
	defer os.Setenv("TEST_SAFE_MEDIA_DIR", originalEnv)

	request := SearchRequest{
		Providers: []string{"embedded"},
		MediaPath: "structure_test.mkv",
		Language:  "en",
	}

	reqBody, err := json.Marshal(request)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/search", bytes.NewBuffer(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := searchHandler()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var response SearchResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	// Verify response structure
	require.NotNil(t, response.Results)
	require.GreaterOrEqual(t, response.Total, 0)
	require.Equal(t, request, response.Query)

	// If there are results, verify they have the expected structure
	for _, result := range response.Results {
		require.NotEmpty(t, result.Provider)
		require.NotEmpty(t, result.Language)
		require.GreaterOrEqual(t, result.Score, 0.0)
		// Note: Not checking SubtitleData as it might not be populated in search results
	}
}

// TestSearchHandlerMultipleProviders tests search with multiple providers
func TestSearchHandlerMultipleProviders(t *testing.T) {
	testDir := t.TempDir()
	testFile := testDir + "/multi_provider_test.mkv"
	require.NoError(t, os.WriteFile(testFile, []byte("test content"), 0644))

	originalEnv := os.Getenv("TEST_SAFE_MEDIA_DIR")
	os.Setenv("TEST_SAFE_MEDIA_DIR", testDir)
	defer os.Setenv("TEST_SAFE_MEDIA_DIR", originalEnv)

	request := SearchRequest{
		Providers: []string{"embedded", "generic"}, // Multiple safe providers
		MediaPath: "multi_provider_test.mkv",
		Language:  "en",
	}

	reqBody, err := json.Marshal(request)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/search", bytes.NewBuffer(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := searchHandler()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var response SearchResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	// Verify response from multiple providers
	require.Equal(t, request, response.Query)
	require.NotNil(t, response.Results)
}

// TestSearchHandlerLanguageVariations tests searching with different languages
func TestSearchHandlerLanguageVariations(t *testing.T) {
	testDir := t.TempDir()
	testFile := testDir + "/lang_test.mkv"
	require.NoError(t, os.WriteFile(testFile, []byte("test content"), 0644))

	originalEnv := os.Getenv("TEST_SAFE_MEDIA_DIR")
	os.Setenv("TEST_SAFE_MEDIA_DIR", testDir)
	defer os.Setenv("TEST_SAFE_MEDIA_DIR", originalEnv)

	languages := []string{"en", "es", "fr", "de"}

	for _, lang := range languages {
		t.Run("language_"+lang, func(t *testing.T) {
			request := SearchRequest{
				Providers: []string{"embedded"},
				MediaPath: "lang_test.mkv",
				Language:  lang,
			}

			reqBody, err := json.Marshal(request)
			require.NoError(t, err)

			req, err := http.NewRequest("POST", "/api/search", bytes.NewBuffer(reqBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := searchHandler()
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)

			var response SearchResponse
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			require.NoError(t, err)

			require.Equal(t, lang, response.Query.Language)
		})
	}
}
