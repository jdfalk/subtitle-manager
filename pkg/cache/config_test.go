// file: pkg/cache/config_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174009

package cache

import (
	"testing"
	"time"

	"github.com/spf13/viper"
)

// TestConfigFromViper tests creating cache configuration from viper.
func TestConfigFromViper(t *testing.T) {
	// Create a new viper instance for testing
	v := viper.New()

	// Set test values
	v.Set("cache.backend", "redis")
	v.Set("cache.memory.max_entries", 5000)
	v.Set("cache.memory.max_memory", 52428800) // 50MB
	v.Set("cache.memory.default_ttl", "2h")
	v.Set("cache.memory.cleanup_interval", "5m")
	v.Set("cache.redis.address", "redis:6379")
	v.Set("cache.redis.password", "secret")
	v.Set("cache.redis.database", 1)
	v.Set("cache.redis.pool_size", 20)
	v.Set("cache.redis.min_idle_conns", 5)
	v.Set("cache.redis.key_prefix", "test:")
	v.Set("cache.redis.dial_timeout", "10s")
	v.Set("cache.redis.read_timeout", "5s")
	v.Set("cache.redis.write_timeout", "5s")
	v.Set("cache.ttls.provider_search_results", "10m")
	v.Set("cache.ttls.tmdb_metadata", "48h")
	v.Set("cache.ttls.translation_results", "1h")
	v.Set("cache.ttls.user_sessions", "12h")
	v.Set("cache.ttls.api_responses", "15m")

	// Replace the global viper instance temporarily
	originalViper := viper.GetViper()
	defer func() {
		// Restore original viper
		for key := range v.AllSettings() {
			viper.GetViper().Set(key, originalViper.Get(key))
		}
	}()

	// Copy test settings to global viper
	for key, value := range v.AllSettings() {
		viper.Set(key, value)
	}

	// Test ConfigFromViper
	config, err := ConfigFromViper()
	if err != nil {
		t.Fatalf("failed to create config from viper: %v", err)
	}

	// Verify backend
	if config.Backend != "redis" {
		t.Errorf("expected backend 'redis', got %s", config.Backend)
	}

	// Verify memory config
	if config.Memory.MaxEntries != 5000 {
		t.Errorf("expected max entries 5000, got %d", config.Memory.MaxEntries)
	}

	if config.Memory.MaxMemory != 52428800 {
		t.Errorf("expected max memory 52428800, got %d", config.Memory.MaxMemory)
	}

	if config.Memory.DefaultTTL != 2*time.Hour {
		t.Errorf("expected default TTL 2h, got %v", config.Memory.DefaultTTL)
	}

	if config.Memory.CleanupInterval != 5*time.Minute {
		t.Errorf("expected cleanup interval 5m, got %v", config.Memory.CleanupInterval)
	}

	// Verify Redis config
	if config.Redis.Address != "redis:6379" {
		t.Errorf("expected address 'redis:6379', got %s", config.Redis.Address)
	}

	if config.Redis.Password != "secret" {
		t.Errorf("expected password 'secret', got %s", config.Redis.Password)
	}

	if config.Redis.Database != 1 {
		t.Errorf("expected database 1, got %d", config.Redis.Database)
	}

	if config.Redis.PoolSize != 20 {
		t.Errorf("expected pool size 20, got %d", config.Redis.PoolSize)
	}

	if config.Redis.MinIdleConns != 5 {
		t.Errorf("expected min idle conns 5, got %d", config.Redis.MinIdleConns)
	}

	if config.Redis.KeyPrefix != "test:" {
		t.Errorf("expected key prefix 'test:', got %s", config.Redis.KeyPrefix)
	}

	if config.Redis.DialTimeout != 10*time.Second {
		t.Errorf("expected dial timeout 10s, got %v", config.Redis.DialTimeout)
	}

	if config.Redis.ReadTimeout != 5*time.Second {
		t.Errorf("expected read timeout 5s, got %v", config.Redis.ReadTimeout)
	}

	if config.Redis.WriteTimeout != 5*time.Second {
		t.Errorf("expected write timeout 5s, got %v", config.Redis.WriteTimeout)
	}

	// Verify TTL config
	if config.TTLs.ProviderSearchResults != 10*time.Minute {
		t.Errorf("expected provider search results TTL 10m, got %v", config.TTLs.ProviderSearchResults)
	}

	if config.TTLs.TMDBMetadata != 48*time.Hour {
		t.Errorf("expected TMDB metadata TTL 48h, got %v", config.TTLs.TMDBMetadata)
	}

	if config.TTLs.TranslationResults != 1*time.Hour {
		t.Errorf("expected translation results TTL 1h, got %v", config.TTLs.TranslationResults)
	}

	if config.TTLs.UserSessions != 12*time.Hour {
		t.Errorf("expected user sessions TTL 12h, got %v", config.TTLs.UserSessions)
	}

	if config.TTLs.APIResponses != 15*time.Minute {
		t.Errorf("expected API responses TTL 15m, got %v", config.TTLs.APIResponses)
	}
}

// TestConfigFromViper_Defaults tests that defaults are used when values are not set.
func TestConfigFromViper_Defaults(t *testing.T) {
	// Create a new viper instance with no settings
	originalViper := viper.GetViper()
	defer func() {
		// Restore original viper
		viper.Reset()
		for key, value := range originalViper.AllSettings() {
			viper.Set(key, value)
		}
	}()

	// Clear all settings
	viper.Reset()

	// Test ConfigFromViper with empty settings (should use zero values)
	config, err := ConfigFromViper()
	if err != nil {
		t.Fatalf("failed to create config from viper: %v", err)
	}

	// Expect default configuration values when none are set
	if config.Backend != "memory" {
		t.Errorf("expected backend 'memory', got %s", config.Backend)
	}

	if config.Memory.MaxEntries != 10000 {
		t.Errorf("expected max entries 10000, got %d", config.Memory.MaxEntries)
	}

	if config.Memory.MaxMemory != 100*1024*1024 {
		t.Errorf("expected max memory 100MB, got %d", config.Memory.MaxMemory)
	}

	if config.Redis.Address != "localhost:6379" {
		t.Errorf("expected address 'localhost:6379', got %s", config.Redis.Address)
	}
}

// TestConfigFromViper_InvalidDurations tests handling of invalid duration strings.
func TestConfigFromViper_InvalidDurations(t *testing.T) {
	// Create a new viper instance for testing
	v := viper.New()

	// Set invalid duration values
	v.Set("cache.memory.default_ttl", "invalid")
	v.Set("cache.memory.cleanup_interval", "not-a-duration")
	v.Set("cache.redis.dial_timeout", "bad-timeout")
	v.Set("cache.ttls.provider_search_results", "malformed")

	// Replace the global viper instance temporarily
	originalViper := viper.GetViper()
	defer func() {
		// Restore original viper
		viper.Reset()
		for key, value := range originalViper.AllSettings() {
			viper.Set(key, value)
		}
	}()

	// Copy test settings to global viper
	for key, value := range v.AllSettings() {
		viper.Set(key, value)
	}

	// Test ConfigFromViper - should not fail, should use fallback values
	config, err := ConfigFromViper()
	if err != nil {
		t.Fatalf("failed to create config from viper: %v", err)
	}

	// Verify fallback values are used
	if config.Memory.DefaultTTL != 1*time.Hour {
		t.Errorf("expected fallback default TTL 1h, got %v", config.Memory.DefaultTTL)
	}

	if config.Memory.CleanupInterval != 10*time.Minute {
		t.Errorf("expected fallback cleanup interval 10m, got %v", config.Memory.CleanupInterval)
	}

	if config.Redis.DialTimeout != 5*time.Second {
		t.Errorf("expected fallback dial timeout 5s, got %v", config.Redis.DialTimeout)
	}

	if config.TTLs.ProviderSearchResults != 5*time.Minute {
		t.Errorf("expected fallback provider search results TTL 5m, got %v", config.TTLs.ProviderSearchResults)
	}
}
