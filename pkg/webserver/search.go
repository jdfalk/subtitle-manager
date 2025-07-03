// file: pkg/webserver/search.go
package webserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/providers"
	"github.com/spf13/viper"
)

// SearchRequest represents the request payload for subtitle search
type SearchRequest struct {
	Providers    []string `json:"providers"`
	MediaPath    string   `json:"mediaPath"`
	Language     string   `json:"language"`
	Season       int      `json:"season,omitempty"`
	Episode      int      `json:"episode,omitempty"`
	Year         int      `json:"year,omitempty"`
	ReleaseGroup string   `json:"releaseGroup,omitempty"`
}

// SearchResult represents a single subtitle search result
type SearchResult struct {
	ID          string  `json:"id"`
	Provider    string  `json:"provider"`
	Name        string  `json:"name"`
	Language    string  `json:"language"`
	Score       float64 `json:"score"`
	Downloads   int     `json:"downloads"`
	DownloadURL string  `json:"downloadUrl"`
	PreviewURL  string  `json:"previewUrl"`
	FileSize    int64   `json:"fileSize,omitempty"`
	UploadDate  string  `json:"uploadDate,omitempty"`
	IsHI        bool    `json:"isHI,omitempty"`        // Hearing Impaired
	FromTrusted bool    `json:"fromTrusted,omitempty"` // From trusted uploader
}

// SearchResponse represents the complete search response
type SearchResponse struct {
	Results []SearchResult `json:"results"`
	Total   int            `json:"total"`
	Query   SearchRequest  `json:"query"`
}

// PreviewResponse represents subtitle content preview
type PreviewResponse struct {
	Content  string `json:"content"`
	Language string `json:"language"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

// SearchHistoryItem represents a search history entry
type SearchHistoryItem struct {
	ID        int           `json:"id"`
	Query     SearchRequest `json:"query"`
	Timestamp time.Time     `json:"timestamp"`
	Results   int           `json:"results"`
}

// isValidURL checks if the provided URL is a valid HTTP/HTTPS URL and not localhost or private IP
func isValidURL(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	// Prevent SSRF: block localhost and private IPs
	host := u.Hostname()
	if host == "localhost" || host == "127.0.0.1" || host == "::1" {
		return false
	}
	if strings.HasPrefix(host, "10.") || strings.HasPrefix(host, "192.168.") || strings.HasPrefix(host, "172.") {
		return false
	}
	return true
}

// searchHandler handles manual subtitle search requests
func searchHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleSearch(w, r)
		case http.MethodGet:
			handleSearchQuery(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// handleSearch processes POST requests for new searches
func handleSearch(w http.ResponseWriter, r *http.Request) {
	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.MediaPath == "" {
		http.Error(w, "Media path is required", http.StatusBadRequest)
		return
	}
	if req.Language == "" {
		req.Language = "en" // Default to English
	}
	if len(req.Providers) == 0 {
		http.Error(w, "At least one provider must be selected", http.StatusBadRequest)
		return
	}

	// Use a configurable safe directory for media files, allow override for tests
	safeDir := viper.GetString("media.safe_dir")
	if testSafeDir := os.Getenv("TEST_SAFE_MEDIA_DIR"); testSafeDir != "" {
		safeDir = testSafeDir
	}
	if safeDir == "" {
		safeDir = "/media/" // fallback default
	}
	// Clean and resolve the absolute path
	joined := filepath.Join(safeDir, req.MediaPath)
	absMediaPath, err := filepath.Abs(joined)
	if err != nil || !strings.HasPrefix(absMediaPath, safeDir) {
		http.Error(w, "Invalid media path", http.StatusBadRequest)
		return
	}

	// Check if media file exists
	if _, err := os.Stat(absMediaPath); os.IsNotExist(err) {
		http.Error(w, "Media file not found", http.StatusNotFound)
		return
	}

	// TODO: Complete implementation for https://github.com/jdfalk/subtitle-manager/issues/36
	// - Add more robust error handling, logging, and test coverage
	// - Consider caching, rate limiting, and provider fallback logic
	// - Optimize parallel search and scoring for large provider sets

	results := performParallelSearch(r.Context(), req)
	scoredResults := calculateScores(results, req)

	response := SearchResponse{
		Results: scoredResults,
		Total:   len(scoredResults),
		Query:   req,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleSearchQuery processes GET requests for simple searches
func handleSearchQuery(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	mediaPath := r.URL.Query().Get("path")
	language := r.URL.Query().Get("lang")
	provider := r.URL.Query().Get("provider")

	if mediaPath == "" {
		http.Error(w, "Media path is required", http.StatusBadRequest)
		return
	}
	if language == "" {
		language = "en"
	}

	// Convert single provider to array for compatibility
	providers := []string{}
	if provider != "" {
		providers = append(providers, provider)
	} else {
		// Default to common providers if none specified
		providers = []string{"opensubtitles", "subscene", "addic7ed"}
	}

	req := SearchRequest{
		Providers: providers,
		MediaPath: mediaPath,
		Language:  language,
	}

	// Parse optional parameters
	if season := r.URL.Query().Get("season"); season != "" {
		if s, err := strconv.Atoi(season); err == nil {
			req.Season = s
		}
	}
	if episode := r.URL.Query().Get("episode"); episode != "" {
		if e, err := strconv.Atoi(episode); err == nil {
			req.Episode = e
		}
	}
	if year := r.URL.Query().Get("year"); year != "" {
		if y, err := strconv.Atoi(year); err == nil {
			req.Year = y
		}
	}
	if releaseGroup := r.URL.Query().Get("releaseGroup"); releaseGroup != "" {
		req.ReleaseGroup = releaseGroup
	}

	// Perform search and return results
	results := performParallelSearch(r.Context(), req)
	scoredResults := calculateScores(results, req)

	// For backward compatibility with existing UI, return simple URL array if single provider
	if len(providers) == 1 {
		urls := make([]string, len(scoredResults))
		for i, result := range scoredResults {
			urls[i] = result.DownloadURL
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(urls)
		return
	}

	response := SearchResponse{
		Results: scoredResults,
		Total:   len(scoredResults),
		Query:   req,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// performParallelSearch executes search across multiple providers concurrently
func performParallelSearch(ctx context.Context, req SearchRequest) []SearchResult {
	var wg sync.WaitGroup
	resultsChan := make(chan []SearchResult, len(req.Providers))

	// Search each provider in parallel
	for _, providerName := range req.Providers {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()

			provider, err := providers.Get(name, "")
			if err != nil {
				return
			}

			// Check if provider supports search
			searcher, ok := provider.(providers.Searcher)
			if !ok {
				return
			}

			urls, err := searcher.Search(ctx, req.MediaPath, req.Language)
			if err != nil {
				return
			}

			// Convert URLs to SearchResult objects
			results := make([]SearchResult, 0, len(urls))
			for i, url := range urls {
				result := SearchResult{
					ID:          fmt.Sprintf("%s_%d", name, i),
					Provider:    name,
					Name:        extractNameFromURL(url),
					Language:    req.Language,
					DownloadURL: url,
					PreviewURL:  fmt.Sprintf("/api/search/preview?url=%s", url),
					Score:       0.5, // Default score, will be calculated later
				}
				results = append(results, result)
			}

			resultsChan <- results
		}(providerName)
	}

	// Close channel when all goroutines complete
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect all results
	var allResults []SearchResult
	for results := range resultsChan {
		allResults = append(allResults, results...)
	}

	return allResults
}

// calculateScores assigns relevance scores to search results
func calculateScores(results []SearchResult, req SearchRequest) []SearchResult {
	for i := range results {
		score := 0.5 // Base score

		// Score based on provider reliability
		switch results[i].Provider {
		case "opensubtitles":
			score += 0.3
		case "subscene":
			score += 0.2
		case "addic7ed":
			score += 0.25
		default:
			score += 0.1
		}

		// Score based on name matching
		if nameMatchScore := calculateNameMatch(results[i].Name, req.MediaPath); nameMatchScore > 0 {
			score += nameMatchScore * 0.2
		}

		// Score based on additional criteria
		if results[i].FromTrusted {
			score += 0.1
		}
		if results[i].Downloads > 1000 {
			score += 0.05
		}

		// Cap score at 1.0
		if score > 1.0 {
			score = 1.0
		}

		results[i].Score = score
	}

	// Sort by score descending
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].Score < results[j].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	return results
}

// calculateNameMatch calculates similarity between subtitle name and media path
func calculateNameMatch(subtitleName, mediaPath string) float64 {
	// Simple name matching - can be enhanced with more sophisticated algorithms
	if subtitleName == "" || mediaPath == "" {
		return 0
	}

	mediaName := strings.ToLower(strings.TrimSuffix(filepath.Base(mediaPath), filepath.Ext(mediaPath)))
	subName := strings.ToLower(subtitleName)

	// Check for common words
	mediaWords := strings.Fields(mediaName)
	subWords := strings.Fields(subName)

	if len(mediaWords) == 0 {
		return 0
	}

	matches := 0
	for _, mWord := range mediaWords {
		for _, sWord := range subWords {
			if strings.Contains(sWord, mWord) || strings.Contains(mWord, sWord) {
				matches++
				break
			}
		}
	}

	return float64(matches) / float64(len(mediaWords))
}

// extractNameFromURL extracts a readable name from download URL
func extractNameFromURL(url string) string {
	// Simple extraction - can be enhanced based on provider URL patterns
	if url == "" {
		return "Subtitle"
	}

	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		name := parts[len(parts)-1]
		// Remove query parameters
		if idx := strings.Index(name, "?"); idx > 0 {
			name = name[:idx]
		}
		if name != "" {
			return name
		}
	}
	return "Subtitle"
}

// searchPreviewHandler handles subtitle content preview requests
func searchPreviewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		urlStr := r.URL.Query().Get("url")
		if urlStr == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
			return
		}
		if !isValidURL(urlStr) {
			http.Error(w, "Invalid or unsafe URL", http.StatusBadRequest)
			return
		}

		// Fetch subtitle content (should use appropriate provider, not direct user URL)
		resp, err := http.Get(urlStr)
		if err != nil {
			http.Error(w, "Failed to fetch subtitle", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Failed to download subtitle", http.StatusBadGateway)
			return
		}

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read subtitle content", http.StatusInternalServerError)
			return
		}

		// Return preview (first 1000 characters)
		previewContent := string(content)
		if len(previewContent) > 1000 {
			previewContent = previewContent[:1000] + "..."
		}

		preview := PreviewResponse{
			Content:  previewContent,
			Language: r.URL.Query().Get("lang"),
			Name:     extractNameFromURL(urlStr),
			Provider: r.URL.Query().Get("provider"),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(preview)
	})
}

// searchHistoryHandler handles search history persistence
func searchHistoryHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetSearchHistory(w, r, db)
		case http.MethodPost:
			handleSaveSearchHistory(w, r, db)
		case http.MethodDelete:
			handleDeleteSearchHistory(w, r, db)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// handleGetSearchHistory retrieves search history
func handleGetSearchHistory(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// For now, return empty array - can be implemented later with proper DB schema
	history := []SearchHistoryItem{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

// handleSaveSearchHistory saves a search to history
func handleSaveSearchHistory(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// For now, just return success - can be implemented later
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

// handleDeleteSearchHistory removes search history
func handleDeleteSearchHistory(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// For now, just return success - can be implemented later
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
