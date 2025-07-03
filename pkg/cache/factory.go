// file: pkg/cache/factory.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174003

package cache

import (
	"context"
	"fmt"
	"strings"
)

// New creates a new cache instance based on the configuration.
// Supported backends: "memory", "redis"
func New(config *Config) (Cache, error) {
	if config == nil {
		config = DefaultConfig()
	}

	backend := strings.ToLower(config.Backend)
	switch backend {
	case "memory", "":
		return NewMemoryCache(config.Memory), nil
	case "redis":
		cache := NewRedisCache(config.Redis)
		// Test Redis connection
		ctx := context.Background()
		if err := cache.Ping(ctx); err != nil {
			cache.Close()
			return nil, fmt.Errorf("redis connection failed: %w", err)
		}
		return cache, nil
	default:
		return nil, fmt.Errorf("unsupported cache backend: %s", backend)
	}
}

// NewWithStatsProvider creates a cache instance that also implements StatsProvider.
func NewWithStatsProvider(config *Config) (interface {
	Cache
	StatsProvider
}, error) {
	cache, err := New(config)
	if err != nil {
		return nil, err
	}

	// All our cache implementations also implement StatsProvider
	if statsCache, ok := cache.(StatsProvider); ok {
		return struct {
			Cache
			StatsProvider
		}{cache, statsCache}, nil
	}

	return nil, fmt.Errorf("cache backend %s does not support statistics", config.Backend)
}

// Manager provides a high-level interface for cache management with type-specific operations.
type Manager struct {
	cache  Cache
	config *Config
}

// NewManager creates a new cache manager with the given configuration.
func NewManager(config *Config) (*Manager, error) {
	cache, err := New(config)
	if err != nil {
		return nil, err
	}

	return &Manager{
		cache:  cache,
		config: config,
	}, nil
}

// GetProviderSearchResults retrieves cached provider search results.
func (m *Manager) GetProviderSearchResults(ctx context.Context, key string) ([]byte, error) {
	return m.cache.Get(ctx, "provider:"+key)
}

// SetProviderSearchResults caches provider search results with configured TTL.
func (m *Manager) SetProviderSearchResults(ctx context.Context, key string, data []byte) error {
	return m.cache.Set(ctx, "provider:"+key, data, m.config.TTLs.ProviderSearchResults)
}

// GetTMDBMetadata retrieves cached TMDB/OMDb metadata.
func (m *Manager) GetTMDBMetadata(ctx context.Context, key string) ([]byte, error) {
	return m.cache.Get(ctx, "tmdb:"+key)
}

// SetTMDBMetadata caches TMDB/OMDb metadata with configured TTL.
func (m *Manager) SetTMDBMetadata(ctx context.Context, key string, data []byte) error {
	return m.cache.Set(ctx, "tmdb:"+key, data, m.config.TTLs.TMDBMetadata)
}

// GetTranslationResult retrieves cached translation results.
func (m *Manager) GetTranslationResult(ctx context.Context, key string) ([]byte, error) {
	return m.cache.Get(ctx, "translation:"+key)
}

// SetTranslationResult caches translation results with configured TTL (permanent if TTL is 0).
func (m *Manager) SetTranslationResult(ctx context.Context, key string, data []byte) error {
	return m.cache.Set(ctx, "translation:"+key, data, m.config.TTLs.TranslationResults)
}

// GetUserSession retrieves cached user session data.
func (m *Manager) GetUserSession(ctx context.Context, key string) ([]byte, error) {
	return m.cache.Get(ctx, "session:"+key)
}

// SetUserSession caches user session data with configured TTL.
func (m *Manager) SetUserSession(ctx context.Context, key string, data []byte) error {
	return m.cache.Set(ctx, "session:"+key, data, m.config.TTLs.UserSessions)
}

// GetAPIResponse retrieves cached API response data.
func (m *Manager) GetAPIResponse(ctx context.Context, key string) ([]byte, error) {
	return m.cache.Get(ctx, "api:"+key)
}

// SetAPIResponse caches API response data with configured TTL.
func (m *Manager) SetAPIResponse(ctx context.Context, key string, data []byte) error {
	return m.cache.Set(ctx, "api:"+key, data, m.config.TTLs.APIResponses)
}

// ClearByPrefix removes all cache entries with the specified prefix.
func (m *Manager) ClearByPrefix(ctx context.Context, prefix string) error {
	// This is a simplified implementation. For production use,
	// you might want to implement prefix-based deletion more efficiently.
	switch cache := m.cache.(type) {
	case *MemoryCache:
		return m.clearMemoryCacheByPrefix(ctx, cache, prefix)
	case *RedisCache:
		return m.clearRedisCacheByPrefix(ctx, cache, prefix)
	default:
		return fmt.Errorf("prefix-based clearing not supported for this cache type")
	}
}

// clearMemoryCacheByPrefix clears memory cache entries by prefix.
func (m *Manager) clearMemoryCacheByPrefix(ctx context.Context, cache *MemoryCache, prefix string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	var keysToDelete []string
	for key := range cache.data {
		if strings.HasPrefix(key, prefix) {
			keysToDelete = append(keysToDelete, key)
		}
	}

	for _, key := range keysToDelete {
		if entry, exists := cache.data[key]; exists {
			delete(cache.data, key)
			delete(cache.accessTime, key)
			cache.stats.MemoryUsage -= entry.size
			cache.stats.Entries--
		}
	}

	return nil
}

// clearRedisCacheByPrefix clears Redis cache entries by prefix.
func (m *Manager) clearRedisCacheByPrefix(ctx context.Context, cache *RedisCache, prefix string) error {
	// Build the full pattern including the Redis key prefix
	pattern := cache.prefixKey(prefix + "*")
	iter := cache.client.Scan(ctx, 0, pattern, 0).Iterator()

	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return err
	}

	if len(keys) > 0 {
		return cache.client.Del(ctx, keys...).Err()
	}

	return nil
}

// Stats returns cache statistics if supported.
func (m *Manager) Stats(ctx context.Context) (*Stats, error) {
	if statsProvider, ok := m.cache.(StatsProvider); ok {
		return statsProvider.Stats(ctx)
	}
	return nil, fmt.Errorf("cache backend does not support statistics")
}

// Close closes the underlying cache.
func (m *Manager) Close() error {
	return m.cache.Close()
}

// Clear removes all entries from the cache.
func (m *Manager) Clear(ctx context.Context) error {
	return m.cache.Clear(ctx)
}

// Delete removes a key from the cache.
func (m *Manager) Delete(ctx context.Context, key string) error {
	return m.cache.Delete(ctx, key)
}
