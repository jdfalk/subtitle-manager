// file: pkg/cache/comprehensive_cache_test.go
// version: 1.0.0
// guid: f9e8d7c6-b5a4-3291-8f7e-6d5c4b3a2109

package cache

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to create a test memory cache with reasonable defaults
func newTestMemoryCache() Cache {
	config := MemoryConfig{
		MaxEntries:      1000,
		MaxMemory:       10 * 1024 * 1024, // 10MB
		DefaultTTL:      5 * time.Minute,
		CleanupInterval: 1 * time.Minute,
	}
	return NewMemoryCache(config)
}

// Test cache hit/miss scenarios for all cache implementations

func TestCacheHitMissScenarios(t *testing.T) {
	testCases := []struct {
		name    string
		factory func() Cache
	}{
		{"MemoryCache", func() Cache {
			config := MemoryConfig{
				MaxEntries:      1000,
				MaxMemory:       10 * 1024 * 1024, // 10MB
				DefaultTTL:      5 * time.Minute,
				CleanupInterval: 1 * time.Minute,
			}
			return NewMemoryCache(config)
		}},
		{"RedisCache", func() Cache {
			// Skip Redis tests if Redis is not available
			config := RedisConfig{
				Address:  "localhost:6379",
				Password: "",
				Database: 0,
				PoolSize: 10,
			}
			cache := NewRedisCache(config)
			// Test if Redis is actually available
			ctx := context.Background()
			if err := cache.Set(ctx, "test_connection", []byte("test"), 1*time.Second); err != nil {
				t.Skip("Redis not available, skipping Redis cache tests")
			}
			return cache
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := tc.factory()
			if cache == nil {
				t.Skip("Cache implementation not available")
			}
			defer cache.Close()

			ctx := context.Background()

			// Test cache miss on non-existent key
			_, err := cache.Get(ctx, "nonexistent_key")
			assert.Equal(t, ErrNotFound, err, "Should return ErrNotFound for non-existent key")

			// Test cache set and hit
			testKey := "test_key_1"
			testValue := []byte("test_value_1")
			err = cache.Set(ctx, testKey, testValue, 1*time.Minute)
			require.NoError(t, err, "Should successfully set cache entry")

			// Verify cache hit
			retrievedValue, err := cache.Get(ctx, testKey)
			require.NoError(t, err, "Should successfully retrieve cached value")
			assert.Equal(t, testValue, retrievedValue, "Retrieved value should match stored value")

			// Test cache existence
			exists, err := cache.Exists(ctx, testKey)
			require.NoError(t, err, "Should successfully check key existence")
			assert.True(t, exists, "Key should exist in cache")

			// Test non-existent key existence
			exists, err = cache.Exists(ctx, "nonexistent_key")
			require.NoError(t, err, "Should successfully check non-existent key")
			assert.False(t, exists, "Non-existent key should not exist")
		})
	}
}

func TestCacheEvictionPolicies(t *testing.T) {
	testCases := []struct {
		name    string
		factory func() Cache
	}{
		{"MemoryCache", func() Cache {
			config := MemoryConfig{
				MaxEntries:      1000,
				MaxMemory:       10 * 1024 * 1024, // 10MB
				DefaultTTL:      5 * time.Minute,
				CleanupInterval: 1 * time.Minute,
			}
			return NewMemoryCache(config)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := tc.factory()
			defer cache.Close()

			ctx := context.Background()

			// Test TTL-based eviction
			testKey := "expiring_key"
			testValue := []byte("expiring_value")

			// Set with very short TTL
			err := cache.Set(ctx, testKey, testValue, 50*time.Millisecond)
			require.NoError(t, err, "Should successfully set expiring entry")

			// Verify entry exists initially
			retrievedValue, err := cache.Get(ctx, testKey)
			require.NoError(t, err, "Should retrieve value before expiration")
			assert.Equal(t, testValue, retrievedValue, "Value should match before expiration")

			// Wait for expiration
			time.Sleep(100 * time.Millisecond)

			// Verify entry is expired/evicted
			_, err = cache.Get(ctx, testKey)
			assert.Equal(t, ErrNotFound, err, "Should return ErrNotFound after expiration")

			// Test manual eviction
			permanentKey := "permanent_key"
			permanentValue := []byte("permanent_value")

			err = cache.Set(ctx, permanentKey, permanentValue, 0) // No expiration
			require.NoError(t, err, "Should set permanent entry")

			// Verify entry exists
			_, err = cache.Get(ctx, permanentKey)
			require.NoError(t, err, "Should retrieve permanent entry")

			// Manually delete entry
			err = cache.Delete(ctx, permanentKey)
			require.NoError(t, err, "Should successfully delete entry")

			// Verify entry is deleted
			_, err = cache.Get(ctx, permanentKey)
			assert.Equal(t, ErrNotFound, err, "Should return ErrNotFound after deletion")
		})
	}
}

func TestCacheConcurrency(t *testing.T) {
	cache := newTestMemoryCache()
	defer cache.Close()

	ctx := context.Background()
	numGoroutines := 50
	numOperationsPerGoroutine := 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Test concurrent reads and writes
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < numOperationsPerGoroutine; j++ {
				key := fmt.Sprintf("key_%d_%d", goroutineID, j)
				value := []byte(fmt.Sprintf("value_%d_%d", goroutineID, j))

				// Set operation
				err := cache.Set(ctx, key, value, 1*time.Minute)
				assert.NoError(t, err, "Concurrent set should not fail")

				// Get operation
				retrievedValue, err := cache.Get(ctx, key)
				assert.NoError(t, err, "Concurrent get should not fail")
				assert.Equal(t, value, retrievedValue, "Retrieved value should match")

				// Exists operation
				exists, err := cache.Exists(ctx, key)
				assert.NoError(t, err, "Concurrent exists should not fail")
				assert.True(t, exists, "Key should exist")
			}
		}(i)
	}

	wg.Wait()
}

func TestCacheTTLBehavior(t *testing.T) {
	cache := newTestMemoryCache()
	defer cache.Close()

	ctx := context.Background()

	// Test TTL functionality
	testKey := "ttl_test_key"
	testValue := []byte("ttl_test_value")
	ttl := 200 * time.Millisecond

	err := cache.Set(ctx, testKey, testValue, ttl)
	require.NoError(t, err, "Should set entry with TTL")

	// Check TTL immediately after setting
	remainingTTL, err := cache.TTL(ctx, testKey)
	require.NoError(t, err, "Should get TTL for existing key")
	assert.Greater(t, remainingTTL, time.Duration(0), "TTL should be positive")
	assert.LessOrEqual(t, remainingTTL, ttl, "TTL should not exceed set TTL")

	// Wait partial time and check TTL again
	time.Sleep(50 * time.Millisecond)
	remainingTTL, err = cache.TTL(ctx, testKey)
	require.NoError(t, err, "Should get updated TTL")
	assert.Greater(t, remainingTTL, time.Duration(0), "TTL should still be positive")
	assert.Less(t, remainingTTL, ttl, "TTL should have decreased")

	// Wait for expiration
	time.Sleep(200 * time.Millisecond)

	// Check TTL after expiration
	remainingTTL, err = cache.TTL(ctx, testKey)
	require.NoError(t, err, "TTL check should not error")
	assert.Equal(t, time.Duration(-1), remainingTTL, "TTL should be -1 for expired/non-existent key")

	// Test no expiration (TTL = 0)
	permanentKey := "permanent_ttl_key"
	permanentValue := []byte("permanent_value")

	err = cache.Set(ctx, permanentKey, permanentValue, 0)
	require.NoError(t, err, "Should set permanent entry")

	ttlResult, err := cache.TTL(ctx, permanentKey)
	require.NoError(t, err, "Should get TTL for permanent entry")
	assert.Equal(t, time.Duration(-2), ttlResult, "TTL should be -2 for entries with no expiration")
}

func TestCacheClearOperation(t *testing.T) {
	cache := newTestMemoryCache()
	defer cache.Close()

	ctx := context.Background()

	// Add multiple entries
	entries := map[string][]byte{
		"key1": []byte("value1"),
		"key2": []byte("value2"),
		"key3": []byte("value3"),
		"key4": []byte("value4"),
		"key5": []byte("value5"),
	}

	for key, value := range entries {
		err := cache.Set(ctx, key, value, 1*time.Minute)
		require.NoError(t, err, "Should set entry %s", key)
	}

	// Verify all entries exist
	for key, expectedValue := range entries {
		value, err := cache.Get(ctx, key)
		require.NoError(t, err, "Should retrieve entry %s", key)
		assert.Equal(t, expectedValue, value, "Value should match for key %s", key)
	}

	// Clear cache
	err := cache.Clear(ctx)
	require.NoError(t, err, "Should clear cache successfully")

	// Verify all entries are cleared
	for key := range entries {
		_, err := cache.Get(ctx, key)
		assert.Equal(t, ErrNotFound, err, "Entry %s should be cleared", key)

		exists, err := cache.Exists(ctx, key)
		require.NoError(t, err, "Exists check should not error for key %s", key)
		assert.False(t, exists, "Key %s should not exist after clear", key)
	}
}

func TestCacheErrorScenarios(t *testing.T) {
	cache := newTestMemoryCache()
	defer cache.Close()

	ctx := context.Background()

	// Test operations on closed cache
	cache.Close()

	// These operations should handle closed cache gracefully
	err := cache.Set(ctx, "test_key", []byte("test_value"), 1*time.Minute)
	assert.Error(t, err, "Set operation should fail on closed cache")

	_, err = cache.Get(ctx, "test_key")
	assert.Error(t, err, "Get operation should fail on closed cache")

	err = cache.Delete(ctx, "test_key")
	assert.Error(t, err, "Delete operation should fail on closed cache")

	_, err = cache.Exists(ctx, "test_key")
	assert.Error(t, err, "Exists operation should fail on closed cache")

	_, err = cache.TTL(ctx, "test_key")
	assert.Error(t, err, "TTL operation should fail on closed cache")

	err = cache.Clear(ctx)
	assert.Error(t, err, "Clear operation should fail on closed cache")
}

func TestCacheValueSizes(t *testing.T) {
	cache := newTestMemoryCache()
	defer cache.Close()

	ctx := context.Background()

	testCases := []struct {
		name      string
		valueSize int
	}{
		{"SmallValue", 10},
		{"MediumValue", 1024},
		{"LargeValue", 1024 * 1024}, // 1MB
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key := fmt.Sprintf("size_test_%s", tc.name)
			value := make([]byte, tc.valueSize)

			// Fill with predictable pattern
			for i := range value {
				value[i] = byte(i % 256)
			}

			err := cache.Set(ctx, key, value, 1*time.Minute)
			require.NoError(t, err, "Should set %s successfully", tc.name)

			retrievedValue, err := cache.Get(ctx, key)
			require.NoError(t, err, "Should retrieve %s successfully", tc.name)
			assert.Equal(t, len(value), len(retrievedValue), "Length should match for %s", tc.name)
			assert.Equal(t, value, retrievedValue, "Content should match for %s", tc.name)
		})
	}
}

func TestCacheContextCancellation(t *testing.T) {
	cache := newTestMemoryCache()
	defer cache.Close()

	// Test context cancellation
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	err := cache.Set(cancelCtx, "cancelled_key", []byte("cancelled_value"), 1*time.Minute)
	if err != nil {
		assert.Equal(t, context.Canceled, err, "Should handle cancelled context appropriately")
	}

	// Test context timeout
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	time.Sleep(1 * time.Millisecond) // Ensure timeout

	err = cache.Set(timeoutCtx, "timeout_key", []byte("timeout_value"), 1*time.Minute)
	if err != nil {
		assert.Equal(t, context.DeadlineExceeded, err, "Should handle context timeout appropriately")
	}
}

// Benchmark tests for cache performance

func BenchmarkCacheOperations(b *testing.B) {
	cache := newTestMemoryCache()
	defer cache.Close()

	ctx := context.Background()
	testKey := "benchmark_key"
	testValue := []byte("benchmark_value_with_reasonable_size_for_testing_performance")

	b.Run("Set", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("%s_%d", testKey, i)
			cache.Set(ctx, key, testValue, 1*time.Minute)
		}
	})

	// Pre-populate cache for get benchmark
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%s_%d", testKey, i)
		cache.Set(ctx, key, testValue, 1*time.Minute)
	}

	b.Run("Get", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("%s_%d", testKey, i%1000)
			cache.Get(ctx, key)
		}
	})

	b.Run("Exists", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("%s_%d", testKey, i%1000)
			cache.Exists(ctx, key)
		}
	})

	b.Run("Delete", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("%s_%d", testKey, i)
			cache.Delete(ctx, key)
		}
	})
}

func BenchmarkCacheConcurrency(b *testing.B) {
	cache := newTestMemoryCache()
	defer cache.Close()

	ctx := context.Background()
	testValue := []byte("concurrent_benchmark_value")

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("concurrent_key_%d", i)

			// Mix of operations
			switch i % 4 {
			case 0:
				cache.Set(ctx, key, testValue, 1*time.Minute)
			case 1:
				cache.Get(ctx, key)
			case 2:
				cache.Exists(ctx, key)
			case 3:
				cache.Delete(ctx, key)
			}
			i++
		}
	})
}
