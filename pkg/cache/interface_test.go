// file: pkg/cache/interface_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174004

package cache

import (
	"testing"
	"time"
)

// TestDefaultConfig verifies the default configuration values.
func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.Backend != "memory" {
		t.Errorf("expected default backend to be 'memory', got %s", config.Backend)
	}

	if config.Memory.MaxEntries != 10000 {
		t.Errorf("expected default max entries to be 10000, got %d", config.Memory.MaxEntries)
	}

	if config.Memory.MaxMemory != 100*1024*1024 {
		t.Errorf("expected default max memory to be 100MB, got %d", config.Memory.MaxMemory)
	}

	if config.Memory.DefaultTTL != 1*time.Hour {
		t.Errorf("expected default TTL to be 1 hour, got %v", config.Memory.DefaultTTL)
	}

	if config.Memory.CleanupInterval != 10*time.Minute {
		t.Errorf("expected default cleanup interval to be 10 minutes, got %v", config.Memory.CleanupInterval)
	}

	if config.Redis.Address != "localhost:6379" {
		t.Errorf("expected default Redis address to be 'localhost:6379', got %s", config.Redis.Address)
	}

	if config.Redis.KeyPrefix != "subtitle-manager:" {
		t.Errorf("expected default Redis key prefix to be 'subtitle-manager:', got %s", config.Redis.KeyPrefix)
	}

	if config.TTLs.ProviderSearchResults != 5*time.Minute {
		t.Errorf("expected provider search results TTL to be 5 minutes, got %v", config.TTLs.ProviderSearchResults)
	}

	if config.TTLs.TMDBMetadata != 24*time.Hour {
		t.Errorf("expected TMDB metadata TTL to be 24 hours, got %v", config.TTLs.TMDBMetadata)
	}

	if config.TTLs.TranslationResults != 0 {
		t.Errorf("expected translation results TTL to be 0 (permanent), got %v", config.TTLs.TranslationResults)
	}

	if config.TTLs.UserSessions != 24*time.Hour {
		t.Errorf("expected user sessions TTL to be 24 hours, got %v", config.TTLs.UserSessions)
	}

	if config.TTLs.APIResponses != 30*time.Minute {
		t.Errorf("expected API responses TTL to be 30 minutes, got %v", config.TTLs.APIResponses)
	}
}
