// file: pkg/cache/factory_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174006

package cache

import (
	"context"
	"testing"
	"time"
)

// TestNew_MemoryBackend tests creating a memory cache through the factory.
func TestNew_MemoryBackend(t *testing.T) {
	config := &Config{
		Backend: "memory",
		Memory: MemoryConfig{
			MaxEntries:      10,
			MaxMemory:       1024,
			DefaultTTL:      1 * time.Hour,
			CleanupInterval: 1 * time.Minute,
		},
	}

	cache, err := New(config)
	if err != nil {
		t.Fatalf("failed to create memory cache: %v", err)
	}
	defer cache.Close()

	// Test basic operations
	ctx := context.Background()
	err = cache.Set(ctx, "test", []byte("value"), 1*time.Hour)
	if err != nil {
		t.Fatalf("failed to set cache entry: %v", err)
	}

	value, err := cache.Get(ctx, "test")
	if err != nil {
		t.Fatalf("failed to get cache entry: %v", err)
	}

	if string(value) != "value" {
		t.Errorf("expected 'value', got %s", string(value))
	}
}

// TestNew_DefaultConfig tests creating a cache with default configuration.
func TestNew_DefaultConfig(t *testing.T) {
	cache, err := New(nil)
	if err != nil {
		t.Fatalf("failed to create cache with default config: %v", err)
	}
	defer cache.Close()

	// Test that it defaults to memory cache
	if _, ok := cache.(*MemoryCache); !ok {
		t.Error("expected MemoryCache with default config")
	}
}

// TestNew_UnsupportedBackend tests error handling for unsupported backends.
func TestNew_UnsupportedBackend(t *testing.T) {
	config := &Config{
		Backend: "unsupported",
	}

	_, err := New(config)
	if err == nil {
		t.Error("expected error for unsupported backend")
	}

	expectedError := "unsupported cache backend: unsupported"
	if err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%s'", expectedError, err.Error())
	}
}

// TestNewManager tests the cache manager functionality.
func TestNewManager(t *testing.T) {
	config := &Config{
		Backend: "memory",
		Memory: MemoryConfig{
			MaxEntries:      100,
			MaxMemory:       1024 * 1024,
			DefaultTTL:      1 * time.Hour,
			CleanupInterval: 1 * time.Minute,
		},
		TTLs: TTLConfig{
			ProviderSearchResults: 5 * time.Minute,
			TMDBMetadata:          24 * time.Hour,
			TranslationResults:    0, // permanent
			UserSessions:          24 * time.Hour,
			APIResponses:          30 * time.Minute,
		},
	}

	manager, err := NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()

	ctx := context.Background()

	// Test provider search results
	err = manager.SetProviderSearchResults(ctx, "search1", []byte("provider-data"))
	if err != nil {
		t.Fatalf("failed to set provider search results: %v", err)
	}

	data, err := manager.GetProviderSearchResults(ctx, "search1")
	if err != nil {
		t.Fatalf("failed to get provider search results: %v", err)
	}
	if string(data) != "provider-data" {
		t.Errorf("expected 'provider-data', got %s", string(data))
	}

	// Test TMDB metadata
	err = manager.SetTMDBMetadata(ctx, "movie123", []byte("movie-metadata"))
	if err != nil {
		t.Fatalf("failed to set TMDB metadata: %v", err)
	}

	data, err = manager.GetTMDBMetadata(ctx, "movie123")
	if err != nil {
		t.Fatalf("failed to get TMDB metadata: %v", err)
	}
	if string(data) != "movie-metadata" {
		t.Errorf("expected 'movie-metadata', got %s", string(data))
	}

	// Test translation results
	err = manager.SetTranslationResult(ctx, "text123", []byte("translated-text"))
	if err != nil {
		t.Fatalf("failed to set translation result: %v", err)
	}

	data, err = manager.GetTranslationResult(ctx, "text123")
	if err != nil {
		t.Fatalf("failed to get translation result: %v", err)
	}
	if string(data) != "translated-text" {
		t.Errorf("expected 'translated-text', got %s", string(data))
	}

	// Test user sessions
	err = manager.SetUserSession(ctx, "session123", []byte("session-data"))
	if err != nil {
		t.Fatalf("failed to set user session: %v", err)
	}

	data, err = manager.GetUserSession(ctx, "session123")
	if err != nil {
		t.Fatalf("failed to get user session: %v", err)
	}
	if string(data) != "session-data" {
		t.Errorf("expected 'session-data', got %s", string(data))
	}

	// Test API responses
	err = manager.SetAPIResponse(ctx, "api123", []byte("api-response"))
	if err != nil {
		t.Fatalf("failed to set API response: %v", err)
	}

	data, err = manager.GetAPIResponse(ctx, "api123")
	if err != nil {
		t.Fatalf("failed to get API response: %v", err)
	}
	if string(data) != "api-response" {
		t.Errorf("expected 'api-response', got %s", string(data))
	}
}

// TestManager_ClearByPrefix tests clearing cache entries by prefix.
func TestManager_ClearByPrefix(t *testing.T) {
	config := &Config{
		Backend: "memory",
		Memory: MemoryConfig{
			MaxEntries:      100,
			MaxMemory:       1024 * 1024,
			DefaultTTL:      1 * time.Hour,
			CleanupInterval: 1 * time.Minute,
		},
		TTLs: TTLConfig{
			ProviderSearchResults: 5 * time.Minute,
			TMDBMetadata:          24 * time.Hour,
		},
	}

	manager, err := NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()

	ctx := context.Background()

	// Add entries with different prefixes
	err = manager.SetProviderSearchResults(ctx, "search1", []byte("provider-data1"))
	if err != nil {
		t.Fatalf("failed to set provider search results: %v", err)
	}

	err = manager.SetProviderSearchResults(ctx, "search2", []byte("provider-data2"))
	if err != nil {
		t.Fatalf("failed to set provider search results: %v", err)
	}

	err = manager.SetTMDBMetadata(ctx, "movie123", []byte("movie-metadata"))
	if err != nil {
		t.Fatalf("failed to set TMDB metadata: %v", err)
	}

	// Clear provider search results
	err = manager.ClearByPrefix(ctx, "provider:")
	if err != nil {
		t.Fatalf("failed to clear by prefix: %v", err)
	}

	// Verify provider entries are gone
	_, err = manager.GetProviderSearchResults(ctx, "search1")
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound for cleared entry, got %v", err)
	}

	_, err = manager.GetProviderSearchResults(ctx, "search2")
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound for cleared entry, got %v", err)
	}

	// Verify TMDB metadata still exists
	data, err := manager.GetTMDBMetadata(ctx, "movie123")
	if err != nil {
		t.Fatalf("failed to get TMDB metadata after prefix clear: %v", err)
	}
	if string(data) != "movie-metadata" {
		t.Errorf("expected 'movie-metadata', got %s", string(data))
	}
}

// TestManager_Stats tests getting statistics from the manager.
func TestManager_Stats(t *testing.T) {
	config := &Config{
		Backend: "memory",
		Memory: MemoryConfig{
			MaxEntries:      100,
			MaxMemory:       1024 * 1024,
			DefaultTTL:      1 * time.Hour,
			CleanupInterval: 1 * time.Minute,
		},
	}

	manager, err := NewManager(config)
	if err != nil {
		t.Fatalf("failed to create cache manager: %v", err)
	}
	defer manager.Close()

	ctx := context.Background()

	// Add some entries
	err = manager.SetProviderSearchResults(ctx, "search1", []byte("data"))
	if err != nil {
		t.Fatalf("failed to set provider search results: %v", err)
	}

	// Get stats
	stats, err := manager.Stats(ctx)
	if err != nil {
		t.Fatalf("failed to get cache stats: %v", err)
	}

	if stats.Entries != 1 {
		t.Errorf("expected 1 entry, got %d", stats.Entries)
	}

	if stats.MemoryLimit != config.Memory.MaxMemory {
		t.Errorf("expected memory limit %d, got %d", config.Memory.MaxMemory, stats.MemoryLimit)
	}
}
