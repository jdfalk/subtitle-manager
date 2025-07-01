// file: pkg/cache/redis_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174007

package cache

import (
	"context"
	"testing"
	"time"
)

// TestRedisCache_BasicOperations tests basic Redis cache operations.
// This test requires a Redis server running on localhost:6379.
func TestRedisCache_BasicOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Redis integration test in short mode")
	}

	config := RedisConfig{
		Address:   "localhost:6379",
		Database:  1, // Use a different database for testing
		KeyPrefix: "test:",
	}

	cache := NewRedisCache(config)

	// Test Redis connection
	ctx := context.Background()
	if err := cache.Ping(ctx); err != nil {
		t.Skipf("Redis not available, skipping test: %v", err)
	}

	defer cache.Close()

	// Clean up any existing test data
	cache.Clear(ctx)

	// Test Set and Get
	key := "test-key"
	value := []byte("test-value")

	err := cache.Set(ctx, key, value, 1*time.Hour)
	if err != nil {
		t.Fatalf("failed to set cache entry: %v", err)
	}

	retrieved, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("failed to get cache entry: %v", err)
	}

	if string(retrieved) != string(value) {
		t.Errorf("expected %s, got %s", string(value), string(retrieved))
	}

	// Test Exists
	exists, err := cache.Exists(ctx, key)
	if err != nil {
		t.Fatalf("failed to check if key exists: %v", err)
	}
	if !exists {
		t.Error("expected key to exist")
	}

	// Test TTL
	ttl, err := cache.TTL(ctx, key)
	if err != nil {
		t.Fatalf("failed to get TTL: %v", err)
	}
	if ttl <= 0 {
		t.Errorf("expected positive TTL, got %v", ttl)
	}

	// Test Delete
	err = cache.Delete(ctx, key)
	if err != nil {
		t.Fatalf("failed to delete cache entry: %v", err)
	}

	// Verify deletion
	_, err = cache.Get(ctx, key)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

// TestRedisCache_TTLExpiration tests TTL expiration behavior with Redis.
func TestRedisCache_TTLExpiration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Redis integration test in short mode")
	}

	config := RedisConfig{
		Address:   "localhost:6379",
		Database:  1,
		KeyPrefix: "test:",
	}

	cache := NewRedisCache(config)

	ctx := context.Background()
	if err := cache.Ping(ctx); err != nil {
		t.Skipf("Redis not available, skipping test: %v", err)
	}

	defer cache.Close()

	// Clean up
	cache.Clear(ctx)

	// Set entry with short TTL
	key := "expire-test"
	value := []byte("expire-value")
	shortTTL := 2 * time.Second

	err := cache.Set(ctx, key, value, shortTTL)
	if err != nil {
		t.Fatalf("failed to set cache entry: %v", err)
	}

	// Verify entry exists
	retrieved, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("failed to get cache entry: %v", err)
	}
	if string(retrieved) != string(value) {
		t.Errorf("expected %s, got %s", string(value), string(retrieved))
	}

	// Wait for expiration
	time.Sleep(3 * time.Second)

	// Verify entry is expired
	_, err = cache.Get(ctx, key)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound for expired entry, got %v", err)
	}
}

// TestRedisCache_NoExpiration tests entries with no expiration in Redis.
func TestRedisCache_NoExpiration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Redis integration test in short mode")
	}

	config := RedisConfig{
		Address:   "localhost:6379",
		Database:  1,
		KeyPrefix: "test:",
	}

	cache := NewRedisCache(config)

	ctx := context.Background()
	if err := cache.Ping(ctx); err != nil {
		t.Skipf("Redis not available, skipping test: %v", err)
	}

	defer cache.Close()

	// Clean up
	cache.Clear(ctx)

	key := "no-expire-test"
	value := []byte("no-expire-value")

	// Set entry with no expiration (TTL = 0)
	err := cache.Set(ctx, key, value, 0)
	if err != nil {
		t.Fatalf("failed to set cache entry: %v", err)
	}

	// Check TTL (should be -2 for no expiration)
	ttl, err := cache.TTL(ctx, key)
	if err != nil {
		t.Fatalf("failed to get TTL: %v", err)
	}
	if ttl != -2 {
		t.Errorf("expected TTL -2 for no expiration, got %v", ttl)
	}

	// Wait a bit and verify entry still exists
	time.Sleep(1 * time.Second)

	retrieved, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("failed to get cache entry: %v", err)
	}
	if string(retrieved) != string(value) {
		t.Errorf("expected %s, got %s", string(value), string(retrieved))
	}
}

// TestRedisCache_Clear tests clearing all cache entries with prefix.
func TestRedisCache_Clear(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Redis integration test in short mode")
	}

	config := RedisConfig{
		Address:   "localhost:6379",
		Database:  1,
		KeyPrefix: "test:",
	}

	cache := NewRedisCache(config)

	ctx := context.Background()
	if err := cache.Ping(ctx); err != nil {
		t.Skipf("Redis not available, skipping test: %v", err)
	}

	defer cache.Close()

	// Clean up first
	cache.Clear(ctx)

	// Add some entries
	for i := 0; i < 5; i++ {
		key := string(rune('a' + i))
		value := []byte("value" + key)
		err := cache.Set(ctx, key, value, 1*time.Hour)
		if err != nil {
			t.Fatalf("failed to set cache entry %s: %v", key, err)
		}
	}

	// Verify entries exist
	for i := 0; i < 5; i++ {
		key := string(rune('a' + i))
		exists, err := cache.Exists(ctx, key)
		if err != nil {
			t.Fatalf("failed to check existence of key %s: %v", key, err)
		}
		if !exists {
			t.Errorf("expected key %s to exist", key)
		}
	}

	// Clear cache
	err := cache.Clear(ctx)
	if err != nil {
		t.Fatalf("failed to clear cache: %v", err)
	}

	// Verify entries are gone
	for i := 0; i < 5; i++ {
		key := string(rune('a' + i))
		exists, err := cache.Exists(ctx, key)
		if err != nil {
			t.Fatalf("failed to check existence of key %s: %v", key, err)
		}
		if exists {
			t.Errorf("expected key %s to be gone after clear", key)
		}
	}
}

// TestNew_RedisBackend tests creating a Redis cache through the factory.
func TestNew_RedisBackend(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Redis integration test in short mode")
	}

	config := &Config{
		Backend: "redis",
		Redis: RedisConfig{
			Address:   "localhost:6379",
			Database:  1,
			KeyPrefix: "test:",
		},
	}

	cache, err := New(config)
	if err != nil {
		// If Redis is not available, skip the test
		t.Skipf("Redis not available, skipping test: %v", err)
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

	// Clean up
	cache.Clear(ctx)
}