// file: pkg/webserver/cache.go
// version: 1.0.0
// guid: 0f9e8d7c-6b5a-1f0e-4a3b-7c6d5e4f3d2c

package webserver

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// CacheEntry represents a cached HTTP response
type CacheEntry struct {
	Body        []byte
	Headers     http.Header
	StatusCode  int
	CachedAt    time.Time
	ExpiresAt   time.Time
}

// IsExpired checks if the cache entry has expired
func (ce *CacheEntry) IsExpired() bool {
	return time.Now().After(ce.ExpiresAt)
}

// InMemoryCache provides a thread-safe in-memory cache for HTTP responses
type InMemoryCache struct {
	entries   map[string]*CacheEntry
	mutex     sync.RWMutex
	defaultTTL time.Duration
}

// NewInMemoryCache creates a new in-memory cache with the specified default TTL
func NewInMemoryCache(defaultTTL time.Duration) *InMemoryCache {
	cache := &InMemoryCache{
		entries:   make(map[string]*CacheEntry),
		defaultTTL: defaultTTL,
	}
	
	// Start cleanup goroutine to remove expired entries
	go cache.cleanupExpiredEntries()
	
	return cache
}

// Get retrieves a cache entry by key, returning nil if not found or expired
func (c *InMemoryCache) Get(key string) *CacheEntry {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	entry, exists := c.entries[key]
	if !exists || entry.IsExpired() {
		return nil
	}
	
	return entry
}

// Set stores a cache entry with the default TTL
func (c *InMemoryCache) Set(key string, entry *CacheEntry) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	entry.CachedAt = time.Now()
	entry.ExpiresAt = time.Now().Add(c.defaultTTL)
	c.entries[key] = entry
}

// SetWithTTL stores a cache entry with a custom TTL
func (c *InMemoryCache) SetWithTTL(key string, entry *CacheEntry, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	entry.CachedAt = time.Now()
	entry.ExpiresAt = time.Now().Add(ttl)
	c.entries[key] = entry
}

// Delete removes a cache entry
func (c *InMemoryCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	delete(c.entries, key)
}

// Clear removes all cache entries
func (c *InMemoryCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.entries = make(map[string]*CacheEntry)
}

// Size returns the number of entries in the cache
func (c *InMemoryCache) Size() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	return len(c.entries)
}

// cleanupExpiredEntries periodically removes expired cache entries
func (c *InMemoryCache) cleanupExpiredEntries() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.entries {
			if entry.IsExpired() {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}

// Global cache instance
var globalCache = NewInMemoryCache(5 * time.Minute)

// generateCacheKey creates a cache key from the request method and URL
func generateCacheKey(r *http.Request) string {
	key := fmt.Sprintf("%s:%s", r.Method, r.URL.String())
	
	// Include query parameters in the hash for more precise caching
	if r.URL.RawQuery != "" {
		key += "?" + r.URL.RawQuery
	}
	
	// Hash the key to keep it a reasonable length
	hash := md5.Sum([]byte(key))
	return fmt.Sprintf("%x", hash)
}

// isCacheable determines if a request/response should be cached
func isCacheable(r *http.Request, statusCode int) bool {
	// Only cache GET requests
	if r.Method != http.MethodGet {
		return false
	}
	
	// Only cache successful responses
	if statusCode < 200 || statusCode >= 300 {
		return false
	}
	
	// Cache API endpoints that are safe to cache
	cachablePaths := []string{
		"/api/system",
		"/api/history",
		"/api/tags",
		"/api/library/browse",
		"/api/widgets",
		"/api/providers",
		"/api/media",
	}
	
	for _, path := range cachablePaths {
		if strings.HasPrefix(r.URL.Path, path) {
			return true
		}
	}
	
	return false
}

// getCacheTTL returns the appropriate cache TTL for different endpoints
func getCacheTTL(r *http.Request) time.Duration {
	switch {
	case strings.HasPrefix(r.URL.Path, "/api/system"):
		return 30 * time.Second // System info changes rarely
	case strings.HasPrefix(r.URL.Path, "/api/widgets"):
		return 2 * time.Minute // Widget data can be cached longer
	case strings.HasPrefix(r.URL.Path, "/api/tags"):
		return 1 * time.Minute // Tags change infrequently
	case strings.HasPrefix(r.URL.Path, "/api/providers"):
		return 5 * time.Minute // Provider info rarely changes
	case strings.HasPrefix(r.URL.Path, "/api/library/browse"):
		return 30 * time.Second // Directory listings can change
	default:
		return 30 * time.Second // Default cache time
	}
}

// CacheMiddleware provides HTTP response caching for GET requests to reduce database load.
//
// This middleware caches responses for frequently accessed API endpoints to improve
// response times and reduce database/backend load. It provides:
//
//   - Automatic caching of GET requests for safe endpoints
//   - Configurable TTL per endpoint type
//   - Memory-efficient storage with automatic cleanup
//   - Cache invalidation and management
//   - Thread-safe concurrent access
//
// Cached endpoints include:
//   - System information (/api/system) - 30s TTL
//   - Widget data (/api/widgets) - 2m TTL  
//   - Tag information (/api/tags) - 1m TTL
//   - Provider configuration (/api/providers) - 5m TTL
//   - Library browsing (/api/library/browse) - 30s TTL
//
// The cache automatically excludes:
//   - Non-GET requests (POST, PUT, DELETE, etc.)
//   - Error responses (4xx, 5xx status codes)
//   - Authentication/session endpoints
//   - Real-time data endpoints
func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if this request should be cached
		if !isCacheable(r, 200) { // Pre-check, will verify status code later
			next.ServeHTTP(w, r)
			return
		}
		
		// Generate cache key
		cacheKey := generateCacheKey(r)
		
		// Try to get from cache
		if cached := globalCache.Get(cacheKey); cached != nil {
			// Serve from cache
			for key, values := range cached.Headers {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}
			w.Header().Set("X-Cache", "HIT")
			w.WriteHeader(cached.StatusCode)
			w.Write(cached.Body)
			return
		}
		
		// Create response recorder to capture the response
		recorder := &cacheResponseRecorder{
			ResponseWriter: w,
			statusCode:     200,
			headers:        make(http.Header),
		}
		
		// Serve the request
		next.ServeHTTP(recorder, r)
		
		// Check if the response should be cached
		if isCacheable(r, recorder.statusCode) {
			// Create cache entry
			entry := &CacheEntry{
				Body:       recorder.body,
				Headers:    recorder.headers,
				StatusCode: recorder.statusCode,
			}
			
			// Store in cache with appropriate TTL
			ttl := getCacheTTL(r)
			globalCache.SetWithTTL(cacheKey, entry, ttl)
			
			// Add cache header
			w.Header().Set("X-Cache", "MISS")
		}
		
		// Write the response to the client
		for key, values := range recorder.headers {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(recorder.statusCode)
		w.Write(recorder.body)
	})
}

// cacheResponseRecorder captures response data for caching
type cacheResponseRecorder struct {
	http.ResponseWriter
	statusCode int
	headers    http.Header
	body       []byte
}

// Header returns the headers map
func (crr *cacheResponseRecorder) Header() http.Header {
	return crr.headers
}

// WriteHeader captures the status code
func (crr *cacheResponseRecorder) WriteHeader(statusCode int) {
	crr.statusCode = statusCode
}

// Write captures the response body
func (crr *cacheResponseRecorder) Write(b []byte) (int, error) {
	crr.body = append(crr.body, b...)
	return len(b), nil
}

// InvalidateCache provides cache invalidation functionality for specific patterns
func InvalidateCache(pattern string) {
	// This could be enhanced to support pattern-based invalidation
	// For now, we'll clear the entire cache when needed
	globalCache.Clear()
}

// CacheStats returns statistics about the cache
func CacheStats() map[string]interface{} {
	return map[string]interface{}{
		"size": globalCache.Size(),
		"ttl":  globalCache.defaultTTL.String(),
	}
}