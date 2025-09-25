// file: pkg/cache/interface.go
// version: 1.2.0
// guid: 123e4567-e89b-12d3-a456-426614174000

package cache

import (
	"context"
	"errors"
	"time"
)

// Common errors for cache operations
var (
	ErrNotFound    = errors.New("cache key not found")
	ErrExpired     = errors.New("cache entry expired")
	ErrCacheClosed = errors.New("cache is closed")
)

// Cache provides a unified interface for caching operations with TTL support.
// Implementations can be memory-based, Redis-based, or other storage backends.
type Cache interface {
	// Get retrieves a value by key. Returns ErrNotFound if key doesn't exist or expired.
	Get(ctx context.Context, key string) ([]byte, error)

	// Set stores a value with the specified TTL. A TTL of 0 means no expiration.
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error

	// Delete removes a key from the cache.
	Delete(ctx context.Context, key string) error

	// Clear removes all entries from the cache.
	Clear(ctx context.Context) error

	// Exists checks if a key exists and is not expired.
	Exists(ctx context.Context, key string) (bool, error)

	// TTL returns the remaining time-to-live for a key.
	// Returns -1 if key doesn't exist, -2 if key exists but has no expiration.
	TTL(ctx context.Context, key string) (time.Duration, error)

	// Close releases any resources held by the cache implementation.
	Close() error
}

// Stats provides cache usage statistics.
type Stats struct {
	// Total number of cached entries
	Entries int64 `json:"entries"`

	// Cache hit rate (0.0 to 1.0)
	HitRate float64 `json:"hit_rate"`

	// Total number of cache hits
	Hits int64 `json:"hits"`

	// Total number of cache misses
	Misses int64 `json:"misses"`

	// Memory usage in bytes (for memory cache)
	MemoryUsage int64 `json:"memory_usage"`

	// Maximum memory limit in bytes (for memory cache)
	MemoryLimit int64 `json:"memory_limit"`

	// Number of evicted entries
	Evictions int64 `json:"evictions"`

	// Uptime since cache initialization
	Uptime time.Duration `json:"uptime"`
}

// StatsProvider provides cache statistics for monitoring and management.
type StatsProvider interface {
	// Stats returns current cache statistics.
	Stats(ctx context.Context) (*Stats, error)
}

// Config holds configuration for cache implementations.
type Config struct {
	// Backend specifies the cache backend: "memory" or "redis"
	Backend string `json:"backend" yaml:"backend" mapstructure:"backend"`

	// Memory cache configuration
	Memory MemoryConfig `json:"memory" yaml:"memory" mapstructure:"memory"`

	// Redis cache configuration
	Redis RedisConfig `json:"redis" yaml:"redis" mapstructure:"redis"`

	// TTL configurations for different data types
	TTLs TTLConfig `json:"ttls" yaml:"ttls" mapstructure:"ttls"`
}

// MemoryConfig configures the in-memory cache implementation.
type MemoryConfig struct {
	// MaxEntries is the maximum number of entries. 0 means unlimited.
	MaxEntries int `json:"max_entries" yaml:"max_entries" mapstructure:"max_entries"`

	// MaxMemory is the maximum memory usage in bytes. 0 means unlimited.
	MaxMemory int64 `json:"max_memory" yaml:"max_memory" mapstructure:"max_memory"`

	// DefaultTTL is the default TTL for entries without explicit TTL.
	DefaultTTL time.Duration `json:"default_ttl" yaml:"default_ttl" mapstructure:"default_ttl"`

	// CleanupInterval is how often expired entries are cleaned up.
	CleanupInterval time.Duration `json:"cleanup_interval" yaml:"cleanup_interval" mapstructure:"cleanup_interval"`
}

// RedisConfig configures the Redis cache implementation.
type RedisConfig struct {
	// Address is the Redis server address (host:port).
	Address string `json:"address" yaml:"address" mapstructure:"address"`

	// Password for Redis authentication.
	Password string `json:"password" yaml:"password" mapstructure:"password"`

	// Database number to use.
	Database int `json:"database" yaml:"database" mapstructure:"database"`

	// PoolSize is the maximum number of socket connections.
	PoolSize int `json:"pool_size" yaml:"pool_size" mapstructure:"pool_size"`

	// MinIdleConns is the minimum number of idle connections.
	MinIdleConns int `json:"min_idle_conns" yaml:"min_idle_conns" mapstructure:"min_idle_conns"`

	// DialTimeout for establishing new connections.
	DialTimeout time.Duration `json:"dial_timeout" yaml:"dial_timeout" mapstructure:"dial_timeout"`

	// ReadTimeout for socket reads.
	ReadTimeout time.Duration `json:"read_timeout" yaml:"read_timeout" mapstructure:"read_timeout"`

	// WriteTimeout for socket writes.
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout" mapstructure:"write_timeout"`

	// KeyPrefix to prepend to all cache keys.
	KeyPrefix string `json:"key_prefix" yaml:"key_prefix" mapstructure:"key_prefix"`
}

// TTLConfig defines TTL values for different types of cached data.
type TTLConfig struct {
	// ProviderSearchResults caches subtitle provider search results.
	ProviderSearchResults time.Duration `json:"provider_search_results" yaml:"provider_search_results" mapstructure:"provider_search_results"`

	// SearchResults caches manual search results.
	SearchResults time.Duration `json:"search_results" yaml:"search_results" mapstructure:"search_results"`

	// TMDBMetadata caches TMDB/OMDb metadata.
	TMDBMetadata time.Duration `json:"tmdb_metadata" yaml:"tmdb_metadata" mapstructure:"tmdb_metadata"`

	// TranslationResults caches translation service results.
	TranslationResults time.Duration `json:"translation_results" yaml:"translation_results" mapstructure:"translation_results"`

	// UserSessions caches user session data.
	UserSessions time.Duration `json:"user_sessions" yaml:"user_sessions" mapstructure:"user_sessions"`

	// APIResponses caches general API responses.
	APIResponses time.Duration `json:"api_responses" yaml:"api_responses" mapstructure:"api_responses"`
}

// DefaultConfig returns a default cache configuration.
func DefaultConfig() *Config {
	return &Config{
		Backend: "memory",
		Memory: MemoryConfig{
			MaxEntries:      10000,
			MaxMemory:       100 * 1024 * 1024, // 100MB
			DefaultTTL:      1 * time.Hour,
			CleanupInterval: 10 * time.Minute,
		},
		Redis: RedisConfig{
			Address:      "localhost:6379",
			Database:     0,
			PoolSize:     10,
			MinIdleConns: 2,
			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
			KeyPrefix:    "subtitle-manager:",
		},
		TTLs: TTLConfig{
			ProviderSearchResults: 5 * time.Minute,
			SearchResults:         5 * time.Minute,
			TMDBMetadata:          24 * time.Hour,
			TranslationResults:    0, // permanent
			UserSessions:          24 * time.Hour,
			APIResponses:          30 * time.Minute,
		},
	}
}
