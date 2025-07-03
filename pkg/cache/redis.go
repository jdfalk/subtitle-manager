// file: pkg/cache/redis.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174002

package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache implements a Redis-based cache backend.
// It provides distributed caching with TTL support.
type RedisCache struct {
	client    *redis.Client
	config    RedisConfig
	stats     *Stats
	startTime time.Time
}

// NewRedisCache creates a new Redis cache with the given configuration.
func NewRedisCache(config RedisConfig) *RedisCache {
	opts := &redis.Options{
		Addr:         config.Address,
		Password:     config.Password,
		DB:           config.Database,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	// Set defaults if not configured
	if opts.PoolSize == 0 {
		opts.PoolSize = 10
	}
	if opts.MinIdleConns == 0 {
		opts.MinIdleConns = 2
	}
	if opts.DialTimeout == 0 {
		opts.DialTimeout = 5 * time.Second
	}
	if opts.ReadTimeout == 0 {
		opts.ReadTimeout = 3 * time.Second
	}
	if opts.WriteTimeout == 0 {
		opts.WriteTimeout = 3 * time.Second
	}

	client := redis.NewClient(opts)

	return &RedisCache{
		client:    client,
		config:    config,
		stats:     &Stats{},
		startTime: time.Now(),
	}
}

// prefixKey adds the configured prefix to a cache key.
func (r *RedisCache) prefixKey(key string) string {
	if r.config.KeyPrefix == "" {
		return key
	}
	return r.config.KeyPrefix + key
}

// Get retrieves a value by key from Redis.
func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	prefixedKey := r.prefixKey(key)
	result := r.client.Get(ctx, prefixedKey)

	if result.Err() == redis.Nil {
		r.stats.Misses++
		return nil, ErrNotFound
	}

	if result.Err() != nil {
		r.stats.Misses++
		return nil, result.Err()
	}

	r.stats.Hits++
	return []byte(result.Val()), nil
}

// Set stores a value with the specified TTL in Redis.
func (r *RedisCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	prefixedKey := r.prefixKey(key)
	return r.client.Set(ctx, prefixedKey, value, ttl).Err()
}

// Delete removes a key from Redis.
func (r *RedisCache) Delete(ctx context.Context, key string) error {
	prefixedKey := r.prefixKey(key)
	return r.client.Del(ctx, prefixedKey).Err()
}

// Clear removes all keys with the configured prefix from Redis.
// WARNING: This operation can be expensive for large datasets.
func (r *RedisCache) Clear(ctx context.Context) error {
	if r.config.KeyPrefix == "" {
		// If no prefix is configured, we can't safely clear all keys
		// as it might affect other applications using the same Redis instance
		return r.client.FlushDB(ctx).Err()
	}

	// Use SCAN to find and delete keys with our prefix
	pattern := r.config.KeyPrefix + "*"
	iter := r.client.Scan(ctx, 0, pattern, 0).Iterator()

	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return err
	}

	if len(keys) > 0 {
		return r.client.Del(ctx, keys...).Err()
	}

	return nil
}

// Exists checks if a key exists in Redis.
func (r *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	prefixedKey := r.prefixKey(key)
	result := r.client.Exists(ctx, prefixedKey)
	if result.Err() != nil {
		return false, result.Err()
	}
	return result.Val() > 0, nil
}

// TTL returns the remaining time-to-live for a key in Redis.
func (r *RedisCache) TTL(ctx context.Context, key string) (time.Duration, error) {
	prefixedKey := r.prefixKey(key)
	result := r.client.TTL(ctx, prefixedKey)
	if result.Err() != nil {
		return -1, result.Err()
	}

	ttl := result.Val()
	if ttl == -2*time.Second {
		// Key doesn't exist
		return -1, nil
	}
	if ttl == -1*time.Second {
		// Key exists but has no expiration
		return -2, nil
	}

	return ttl, nil
}

// Close closes the Redis connection.
func (r *RedisCache) Close() error {
	return r.client.Close()
}

// Stats returns current cache statistics.
// Note: Redis statistics are limited compared to memory cache.
func (r *RedisCache) Stats(ctx context.Context) (*Stats, error) {
	// Get Redis INFO stats
	info := r.client.Info(ctx, "stats", "memory")
	if info.Err() != nil {
		return nil, info.Err()
	}

	stats := *r.stats
	stats.Uptime = time.Since(r.startTime)

	// Calculate hit rate
	total := stats.Hits + stats.Misses
	if total > 0 {
		stats.HitRate = float64(stats.Hits) / float64(total)
	}

	// Try to get keyspace info for entry count
	keyspaceInfo := r.client.Info(ctx, "keyspace")
	if keyspaceInfo.Err() == nil {
		// Parse keyspace info if available
		// This is a simplified approach - in production you might want
		// more sophisticated parsing of the INFO keyspace output
		if r.config.KeyPrefix != "" {
			// Count keys with our prefix
			pattern := r.config.KeyPrefix + "*"
			scanResult := r.client.Scan(ctx, 0, pattern, 0)
			if scanResult.Err() == nil {
				keys, _ := scanResult.Val()
				stats.Entries = int64(len(keys))
			}
		}
	}

	return &stats, nil
}

// Ping tests the Redis connection.
func (r *RedisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}
