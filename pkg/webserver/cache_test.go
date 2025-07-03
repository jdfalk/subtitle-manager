// file: pkg/webserver/cache_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174011

package webserver

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/cache"
)

// TestCacheStatsHandler tests the cache statistics endpoint.
func TestCacheStatsHandler(t *testing.T) {
	// Initialize a test cache manager
	config := cache.DefaultConfig()
	manager, err := cache.NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()
	cacheManager = manager

	// Add some test data
	ctx := context.Background()
	err = manager.SetProviderSearchResults(ctx, "test", []byte("data"))
	if err != nil {
		t.Fatalf("failed to set test data: %v", err)
	}

	req, err := http.NewRequest("GET", "/api/cache/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := cacheStatsHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body contains stats
	var stats cache.Stats
	if err := json.Unmarshal(rr.Body.Bytes(), &stats); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if stats.Entries != 1 {
		t.Errorf("expected 1 entry, got %d", stats.Entries)
	}
}

// TestCacheStatsHandler_NotInitialized tests the stats endpoint when cache is not initialized.
func TestCacheStatsHandler_NotInitialized(t *testing.T) {
	// Set cacheManager to nil to simulate uninitialized state
	originalManager := cacheManager
	cacheManager = nil
	defer func() { cacheManager = originalManager }()

	req, err := http.NewRequest("GET", "/api/cache/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := cacheStatsHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

// TestCacheClearHandler tests the cache clear endpoint.
func TestCacheClearHandler(t *testing.T) {
	// Initialize a test cache manager
	config := cache.DefaultConfig()
	manager, err := cache.NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()
	cacheManager = manager

	// Add some test data
	ctx := context.Background()
	err = manager.SetProviderSearchResults(ctx, "test1", []byte("data1"))
	if err != nil {
		t.Fatalf("failed to set test data: %v", err)
	}
	err = manager.SetTMDBMetadata(ctx, "test2", []byte("data2"))
	if err != nil {
		t.Fatalf("failed to set test data: %v", err)
	}

	// Test clearing all cache
	req, err := http.NewRequest("POST", "/api/cache/clear", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := cacheClearHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify cache is cleared
	_, err = manager.GetProviderSearchResults(ctx, "test1")
	if err != cache.ErrNotFound {
		t.Error("expected cache to be cleared")
	}

	_, err = manager.GetTMDBMetadata(ctx, "test2")
	if err != cache.ErrNotFound {
		t.Error("expected cache to be cleared")
	}
}

// TestCacheClearHandler_WithPrefix tests clearing cache by prefix.
func TestCacheClearHandler_WithPrefix(t *testing.T) {
	// Initialize a test cache manager
	config := cache.DefaultConfig()
	manager, err := cache.NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()
	cacheManager = manager

	// Add some test data
	ctx := context.Background()
	err = manager.SetProviderSearchResults(ctx, "test1", []byte("data1"))
	if err != nil {
		t.Fatalf("failed to set test data: %v", err)
	}
	err = manager.SetTMDBMetadata(ctx, "test2", []byte("data2"))
	if err != nil {
		t.Fatalf("failed to set test data: %v", err)
	}

	// Test clearing only provider cache
	requestBody := map[string]string{"prefix": "provider:"}
	body, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/api/cache/clear", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := cacheClearHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify only provider cache is cleared
	_, err = manager.GetProviderSearchResults(ctx, "test1")
	if err != cache.ErrNotFound {
		t.Error("expected provider cache to be cleared")
	}

	// TMDB cache should still exist
	_, err = manager.GetTMDBMetadata(ctx, "test2")
	if err != nil {
		t.Error("TMDB cache should not be cleared")
	}
}

// TestCacheConfigHandler tests the cache configuration endpoint.
func TestCacheConfigHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cache/config", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := cacheConfigHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body contains config
	var config cache.Config
	if err := json.Unmarshal(rr.Body.Bytes(), &config); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	// Config should have some default values
	if config.Backend == "" {
		t.Error("expected backend to be set in config")
	}
}

// TestCacheHealthHandler tests the cache health endpoint.
func TestCacheHealthHandler(t *testing.T) {
	// Initialize a test cache manager
	config := cache.DefaultConfig()
	manager, err := cache.NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()
	cacheManager = manager

	req, err := http.NewRequest("GET", "/api/cache/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := cacheHealthHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if response["status"] != "healthy" {
		t.Errorf("expected status to be 'healthy', got %v", response["status"])
	}
}

// TestCacheTypedOperationsHandler tests cache operations on specific types.
func TestCacheTypedOperationsHandler(t *testing.T) {
	// Initialize a test cache manager
	config := cache.DefaultConfig()
	manager, err := cache.NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()
	cacheManager = manager

	// Add some test data
	ctx := context.Background()
	err = manager.SetProviderSearchResults(ctx, "test1", []byte("data1"))
	if err != nil {
		t.Fatalf("failed to set test data: %v", err)
	}
	err = manager.SetTMDBMetadata(ctx, "test2", []byte("data2"))
	if err != nil {
		t.Fatalf("failed to set test data: %v", err)
	}

	// Test clearing provider cache
	req, err := http.NewRequest("POST", "/api/cache/types/provider/clear", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := cacheTypedOperationsHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify only provider cache is cleared
	_, err = manager.GetProviderSearchResults(ctx, "test1")
	if err != cache.ErrNotFound {
		t.Error("expected provider cache to be cleared")
	}

	// TMDB cache should still exist
	_, err = manager.GetTMDBMetadata(ctx, "test2")
	if err != nil {
		t.Error("TMDB cache should not be cleared")
	}

	// Check response body
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if response["status"] != "success" {
		t.Errorf("expected status to be 'success', got %v", response["status"])
	}

	if response["type"] != "provider" {
		t.Errorf("expected type to be 'provider', got %v", response["type"])
	}
}

// TestCacheTypedOperationsHandler_InvalidType tests error handling for invalid cache types.
func TestCacheTypedOperationsHandler_InvalidType(t *testing.T) {
	// Initialize a test cache manager
	config := cache.DefaultConfig()
	manager, err := cache.NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()
	cacheManager = manager

	req, err := http.NewRequest("POST", "/api/cache/types/invalid/clear", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := cacheTypedOperationsHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
