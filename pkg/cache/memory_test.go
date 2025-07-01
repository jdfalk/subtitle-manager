// file: pkg/cache/memory_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174005

package cache

import (
	"context"
	"testing"
	"time"
)

// TestMemoryCache_BasicOperations tests basic cache operations.
func TestMemoryCache_BasicOperations(t *testing.T) {
	config := MemoryConfig{
		MaxEntries:      100,
		MaxMemory:       1024 * 1024, // 1MB
		DefaultTTL:      1 * time.Hour,
		CleanupInterval: 1 * time.Minute,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

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

// TestMemoryCache_TTLExpiration tests TTL expiration behavior.
func TestMemoryCache_TTLExpiration(t *testing.T) {
	config := MemoryConfig{
		MaxEntries:      100,
		MaxMemory:       1024 * 1024,
		DefaultTTL:      1 * time.Hour,
		CleanupInterval: 100 * time.Millisecond,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

	// Set entry with short TTL
	key := "expire-test"
	value := []byte("expire-value")
	shortTTL := 200 * time.Millisecond

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
	time.Sleep(300 * time.Millisecond)

	// Verify entry is expired
	_, err = cache.Get(ctx, key)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound for expired entry, got %v", err)
	}

	// Verify TTL returns -1 for non-existent key
	ttl, err := cache.TTL(ctx, key)
	if err != nil {
		t.Fatalf("failed to get TTL: %v", err)
	}
	if ttl != -1 {
		t.Errorf("expected TTL -1 for non-existent key, got %v", ttl)
	}
}

// TestMemoryCache_NoExpiration tests entries with no expiration.
func TestMemoryCache_NoExpiration(t *testing.T) {
	config := MemoryConfig{
		MaxEntries:      100,
		MaxMemory:       1024 * 1024,
		DefaultTTL:      1 * time.Hour,
		CleanupInterval: 100 * time.Millisecond,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

	key := "no-expire-test"
	value := []byte("no-expire-value")

	// Set entry with no expiration (TTL = 0)
	err := cache.Set(ctx, key, value, 0)
	if err != nil {
		t.Fatalf("failed to set cache entry: %v", err)
	}

	// Check TTL
	ttl, err := cache.TTL(ctx, key)
	if err != nil {
		t.Fatalf("failed to get TTL: %v", err)
	}
	if ttl != -2 {
		t.Errorf("expected TTL -2 for no expiration, got %v", ttl)
	}

	// Wait a bit and verify entry still exists
	time.Sleep(100 * time.Millisecond)

	retrieved, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("failed to get cache entry: %v", err)
	}
	if string(retrieved) != string(value) {
		t.Errorf("expected %s, got %s", string(value), string(retrieved))
	}
}

// TestMemoryCache_MaxEntries tests entry count limit enforcement.
func TestMemoryCache_MaxEntries(t *testing.T) {
	config := MemoryConfig{
		MaxEntries:      3, // Small limit for testing
		MaxMemory:       1024 * 1024,
		DefaultTTL:      1 * time.Hour,
		CleanupInterval: 1 * time.Minute,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

	// Add entries up to the limit
	for i := 0; i < 3; i++ {
		key := string(rune('a' + i))
		value := []byte("value" + key)
		err := cache.Set(ctx, key, value, 1*time.Hour)
		if err != nil {
			t.Fatalf("failed to set cache entry %s: %v", key, err)
		}
	}

	// Add one more entry, which should trigger eviction
	err := cache.Set(ctx, "d", []byte("valued"), 1*time.Hour)
	if err != nil {
		t.Fatalf("failed to set cache entry d: %v", err)
	}

	// Get stats to verify entry count
	stats, err := cache.Stats(ctx)
	if err != nil {
		t.Fatalf("failed to get cache stats: %v", err)
	}

	if stats.Entries > 3 {
		t.Errorf("expected at most 3 entries, got %d", stats.Entries)
	}

	if stats.Evictions == 0 {
		t.Error("expected some evictions to have occurred")
	}
}

// TestMemoryCache_Clear tests clearing all cache entries.
func TestMemoryCache_Clear(t *testing.T) {
	config := MemoryConfig{
		MaxEntries:      100,
		MaxMemory:       1024 * 1024,
		DefaultTTL:      1 * time.Hour,
		CleanupInterval: 1 * time.Minute,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

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
	stats, err := cache.Stats(ctx)
	if err != nil {
		t.Fatalf("failed to get cache stats: %v", err)
	}
	if stats.Entries != 5 {
		t.Errorf("expected 5 entries, got %d", stats.Entries)
	}

	// Clear cache
	err = cache.Clear(ctx)
	if err != nil {
		t.Fatalf("failed to clear cache: %v", err)
	}

	// Verify cache is empty
	stats, err = cache.Stats(ctx)
	if err != nil {
		t.Fatalf("failed to get cache stats: %v", err)
	}
	if stats.Entries != 0 {
		t.Errorf("expected 0 entries after clear, got %d", stats.Entries)
	}
	if stats.MemoryUsage != 0 {
		t.Errorf("expected 0 memory usage after clear, got %d", stats.MemoryUsage)
	}
}

// TestMemoryCache_Stats tests cache statistics.
func TestMemoryCache_Stats(t *testing.T) {
	config := MemoryConfig{
		MaxEntries:      100,
		MaxMemory:       1024 * 1024,
		DefaultTTL:      1 * time.Hour,
		CleanupInterval: 1 * time.Minute,
	}

	cache := NewMemoryCache(config)
	defer cache.Close()

	ctx := context.Background()

	// Add an entry
	err := cache.Set(ctx, "test", []byte("value"), 1*time.Hour)
	if err != nil {
		t.Fatalf("failed to set cache entry: %v", err)
	}

	// Get the entry (cache hit)
	_, err = cache.Get(ctx, "test")
	if err != nil {
		t.Fatalf("failed to get cache entry: %v", err)
	}

	// Try to get non-existent entry (cache miss)
	_, err = cache.Get(ctx, "nonexistent")
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}

	// Get stats
	stats, err := cache.Stats(ctx)
	if err != nil {
		t.Fatalf("failed to get cache stats: %v", err)
	}

	if stats.Entries != 1 {
		t.Errorf("expected 1 entry, got %d", stats.Entries)
	}

	if stats.Hits != 1 {
		t.Errorf("expected 1 hit, got %d", stats.Hits)
	}

	if stats.Misses != 1 {
		t.Errorf("expected 1 miss, got %d", stats.Misses)
	}

	if stats.HitRate != 0.5 {
		t.Errorf("expected hit rate 0.5, got %f", stats.HitRate)
	}

	if stats.MemoryUsage <= 0 {
		t.Errorf("expected positive memory usage, got %d", stats.MemoryUsage)
	}

	if stats.MemoryLimit != config.MaxMemory {
		t.Errorf("expected memory limit %d, got %d", config.MaxMemory, stats.MemoryLimit)
	}

	if stats.Uptime <= 0 {
		t.Errorf("expected positive uptime, got %v", stats.Uptime)
	}
}