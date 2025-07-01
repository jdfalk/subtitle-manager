// file: pkg/cache/memory.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174001

package cache

import (
	"context"
	"sync"
	"time"
)

// cacheEntry represents a single cache entry with value and expiration.
type cacheEntry struct {
	value     []byte
	expiresAt time.Time
	size      int64
}

// isExpired checks if the entry has expired.
func (e *cacheEntry) isExpired() bool {
	return !e.expiresAt.IsZero() && time.Now().After(e.expiresAt)
}

// MemoryCache implements an in-memory cache with TTL support and LRU eviction.
// It provides thread-safe operations using read-write mutex.
type MemoryCache struct {
	config     MemoryConfig
	data       map[string]*cacheEntry
	accessTime map[string]time.Time // for LRU tracking
	mu         sync.RWMutex
	stats      *Stats
	startTime  time.Time
	stopCh     chan struct{}
	stopped    bool
}

// NewMemoryCache creates a new in-memory cache with the given configuration.
func NewMemoryCache(config MemoryConfig) *MemoryCache {
	if config.CleanupInterval == 0 {
		config.CleanupInterval = 10 * time.Minute
	}
	if config.DefaultTTL == 0 {
		config.DefaultTTL = 1 * time.Hour
	}

	cache := &MemoryCache{
		config:     config,
		data:       make(map[string]*cacheEntry),
		accessTime: make(map[string]time.Time),
		stats: &Stats{
			MemoryLimit: config.MaxMemory,
		},
		startTime: time.Now(),
		stopCh:    make(chan struct{}),
	}

	// Start cleanup goroutine
	go cache.cleanupLoop()

	return cache
}

// Get retrieves a value by key.
func (m *MemoryCache) Get(ctx context.Context, key string) ([]byte, error) {
	m.mu.RLock()
	entry, exists := m.data[key]
	m.mu.RUnlock()

	if !exists {
		m.mu.Lock()
		m.stats.Misses++
		m.mu.Unlock()
		return nil, ErrNotFound
	}

	if entry.isExpired() {
		// Remove expired entry
		m.mu.Lock()
		delete(m.data, key)
		delete(m.accessTime, key)
		m.stats.Entries--
		m.stats.MemoryUsage -= entry.size
		m.stats.Misses++
		m.mu.Unlock()
		return nil, ErrNotFound
	}

	// Update access time for LRU
	m.mu.Lock()
	m.accessTime[key] = time.Now()
	m.stats.Hits++
	m.mu.Unlock()

	// Return a copy to prevent modification
	result := make([]byte, len(entry.value))
	copy(result, entry.value)
	return result, nil
}

// Set stores a value with the specified TTL.
func (m *MemoryCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	// Create entry
	entry := &cacheEntry{
		value: make([]byte, len(value)),
		size:  int64(len(value) + len(key) + 64), // approximate overhead
	}
	copy(entry.value, value)

	// Handle TTL: 0 means no expiration, < 0 means use default TTL
	if ttl == 0 {
		// No expiration - leave expiresAt as zero time
	} else if ttl < 0 {
		// Use default TTL
		if m.config.DefaultTTL > 0 {
			entry.expiresAt = time.Now().Add(m.config.DefaultTTL)
		}
	} else {
		// Use specified TTL
		entry.expiresAt = time.Now().Add(ttl)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if we need to evict existing entry
	if existing, exists := m.data[key]; exists {
		m.stats.MemoryUsage -= existing.size
		m.stats.Entries--
	}

	// Check memory limit
	if m.config.MaxMemory > 0 && m.stats.MemoryUsage+entry.size > m.config.MaxMemory {
		m.evictLRU(entry.size)
	}

	// Check entry limit
	if m.config.MaxEntries > 0 && len(m.data) >= m.config.MaxEntries {
		m.evictLRUEntries(1)
	}

	// Store entry
	m.data[key] = entry
	m.accessTime[key] = time.Now()
	m.stats.MemoryUsage += entry.size
	m.stats.Entries++

	return nil
}

// Delete removes a key from the cache.
func (m *MemoryCache) Delete(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if entry, exists := m.data[key]; exists {
		delete(m.data, key)
		delete(m.accessTime, key)
		m.stats.MemoryUsage -= entry.size
		m.stats.Entries--
	}

	return nil
}

// Clear removes all entries from the cache.
func (m *MemoryCache) Clear(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data = make(map[string]*cacheEntry)
	m.accessTime = make(map[string]time.Time)
	m.stats.MemoryUsage = 0
	m.stats.Entries = 0

	return nil
}

// Exists checks if a key exists and is not expired.
func (m *MemoryCache) Exists(ctx context.Context, key string) (bool, error) {
	m.mu.RLock()
	entry, exists := m.data[key]
	m.mu.RUnlock()

	if !exists {
		return false, nil
	}

	if entry.isExpired() {
		// Remove expired entry
		m.mu.Lock()
		delete(m.data, key)
		delete(m.accessTime, key)
		m.stats.Entries--
		m.stats.MemoryUsage -= entry.size
		m.mu.Unlock()
		return false, nil
	}

	return true, nil
}

// TTL returns the remaining time-to-live for a key.
func (m *MemoryCache) TTL(ctx context.Context, key string) (time.Duration, error) {
	m.mu.RLock()
	entry, exists := m.data[key]
	m.mu.RUnlock()

	if !exists {
		return -1, nil
	}

	if entry.expiresAt.IsZero() {
		return -2, nil // no expiration
	}

	if entry.isExpired() {
		// Remove expired entry
		m.mu.Lock()
		delete(m.data, key)
		delete(m.accessTime, key)
		m.stats.Entries--
		m.stats.MemoryUsage -= entry.size
		m.mu.Unlock()
		return -1, nil
	}

	return time.Until(entry.expiresAt), nil
}

// Close stops the cleanup goroutine and releases resources.
func (m *MemoryCache) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.stopped {
		close(m.stopCh)
		m.stopped = true
	}

	return nil
}

// Stats returns current cache statistics.
func (m *MemoryCache) Stats(ctx context.Context) (*Stats, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	stats := *m.stats
	stats.Uptime = time.Since(m.startTime)

	// Calculate hit rate
	total := stats.Hits + stats.Misses
	if total > 0 {
		stats.HitRate = float64(stats.Hits) / float64(total)
	}

	return &stats, nil
}

// evictLRU evicts least recently used entries to free up the specified amount of memory.
func (m *MemoryCache) evictLRU(needed int64) {
	type lruEntry struct {
		key        string
		accessTime time.Time
		size       int64
	}

	var entries []lruEntry
	for key, accessTime := range m.accessTime {
		if entry, exists := m.data[key]; exists {
			entries = append(entries, lruEntry{
				key:        key,
				accessTime: accessTime,
				size:       entry.size,
			})
		}
	}

	// Sort by access time (oldest first)
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].accessTime.After(entries[j].accessTime) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}

	freed := int64(0)
	for _, entry := range entries {
		if freed >= needed {
			break
		}

		delete(m.data, entry.key)
		delete(m.accessTime, entry.key)
		m.stats.MemoryUsage -= entry.size
		m.stats.Entries--
		m.stats.Evictions++
		freed += entry.size
	}
}

// evictLRUEntries evicts the specified number of least recently used entries.
func (m *MemoryCache) evictLRUEntries(count int) {
	type lruEntry struct {
		key        string
		accessTime time.Time
	}

	var entries []lruEntry
	for key, accessTime := range m.accessTime {
		entries = append(entries, lruEntry{
			key:        key,
			accessTime: accessTime,
		})
	}

	// Sort by access time (oldest first)
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].accessTime.After(entries[j].accessTime) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}

	evicted := 0
	for _, entry := range entries {
		if evicted >= count {
			break
		}

		if dataEntry, exists := m.data[entry.key]; exists {
			delete(m.data, entry.key)
			delete(m.accessTime, entry.key)
			m.stats.MemoryUsage -= dataEntry.size
			m.stats.Entries--
			m.stats.Evictions++
			evicted++
		}
	}
}

// cleanupLoop periodically removes expired entries.
func (m *MemoryCache) cleanupLoop() {
	ticker := time.NewTicker(m.config.CleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.cleanup()
		case <-m.stopCh:
			return
		}
	}
}

// cleanup removes expired entries from the cache.
func (m *MemoryCache) cleanup() {
	m.mu.Lock()
	defer m.mu.Unlock()

	var keysToDelete []string
	for key, entry := range m.data {
		if entry.isExpired() {
			keysToDelete = append(keysToDelete, key)
		}
	}

	for _, key := range keysToDelete {
		entry := m.data[key]
		delete(m.data, key)
		delete(m.accessTime, key)
		m.stats.MemoryUsage -= entry.size
		m.stats.Entries--
	}
}