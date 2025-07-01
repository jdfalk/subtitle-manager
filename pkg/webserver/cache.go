// file: pkg/webserver/cache.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174010

package webserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jdfalk/subtitle-manager/pkg/cache"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

// Global cache manager instance
var cacheManager *cache.Manager

// InitializeCache initializes the global cache manager from configuration.
func InitializeCache() error {
	manager, err := cache.NewManagerFromViper()
	if err != nil {
		return err
	}
	cacheManager = manager
	return nil
}

// GetCacheManager returns the global cache manager instance.
func GetCacheManager() *cache.Manager {
	return cacheManager
}

// cacheStatsHandler returns cache statistics.
func cacheStatsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetLogger("webserver.cache")

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if cacheManager == nil {
			http.Error(w, "Cache not initialized", http.StatusInternalServerError)
			return
		}

		stats, err := cacheManager.Stats(r.Context())
		if err != nil {
			logger.Errorf("failed to get cache stats: %v", err)
			http.Error(w, "Failed to get cache statistics", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(stats); err != nil {
			logger.Errorf("failed to encode cache stats response: %v", err)
		}
	})
}

// cacheClearHandler clears cache entries.
func cacheClearHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetLogger("webserver.cache")

		if r.Method != http.MethodPost && r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if cacheManager == nil {
			http.Error(w, "Cache not initialized", http.StatusInternalServerError)
			return
		}

		// Parse request body for optional prefix
		var req struct {
			Prefix string `json:"prefix,omitempty"`
		}

		if r.Header.Get("Content-Type") == "application/json" {
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				logger.Warnf("failed to decode clear cache request: %v", err)
				// Continue with empty prefix (clear all)
			}
		}

		var err error
		if req.Prefix != "" {
			// Clear by prefix
			err = cacheManager.ClearByPrefix(r.Context(), req.Prefix)
			logger.Infof("cleared cache entries with prefix: %s", req.Prefix)
		} else {
			// Clear all cache
			err = cacheManager.Clear(r.Context())
			logger.Info("cleared all cache entries")
		}

		if err != nil {
			logger.Errorf("failed to clear cache: %v", err)
			http.Error(w, "Failed to clear cache", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "success", "message": "Cache cleared"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Errorf("failed to encode clear cache response: %v", err)
		}
	})
}

// cacheConfigHandler returns current cache configuration.
func cacheConfigHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetLogger("webserver.cache")

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		config, err := cache.ConfigFromViper()
		if err != nil {
			logger.Errorf("failed to get cache config: %v", err)
			http.Error(w, "Failed to get cache configuration", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(config); err != nil {
			logger.Errorf("failed to encode cache config response: %v", err)
		}
	})
}

// cacheTypedOperationsHandler handles operations on specific cache types.
func cacheTypedOperationsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetLogger("webserver.cache")

		if cacheManager == nil {
			http.Error(w, "Cache not initialized", http.StatusInternalServerError)
			return
		}

		// Extract cache type from URL path
		// Expected path: /api/cache/types/{type}/clear
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) < 4 {
			http.Error(w, "Invalid cache type path", http.StatusBadRequest)
			return
		}

		cacheType := parts[3] // Should be like "provider", "tmdb", "translation", etc.
		operation := ""
		if len(parts) > 4 {
			operation = parts[4]
		}

		// Map cache types to prefixes
		var prefix string
		switch cacheType {
		case "provider":
			prefix = "provider:"
		case "tmdb":
			prefix = "tmdb:"
		case "translation":
			prefix = "translation:"
		case "session":
			prefix = "session:"
		case "api":
			prefix = "api:"
		default:
			http.Error(w, "Unknown cache type", http.StatusBadRequest)
			return
		}

		switch operation {
		case "clear":
			if r.Method != http.MethodPost && r.Method != http.MethodDelete {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			err := cacheManager.ClearByPrefix(r.Context(), prefix)
			if err != nil {
				logger.Errorf("failed to clear %s cache: %v", cacheType, err)
				http.Error(w, "Failed to clear cache", http.StatusInternalServerError)
				return
			}

			logger.Infof("cleared %s cache entries", cacheType)
			w.Header().Set("Content-Type", "application/json")
			response := map[string]string{
				"status":  "success",
				"message": "Cache cleared for " + cacheType,
				"type":    cacheType,
			}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				logger.Errorf("failed to encode typed cache clear response: %v", err)
			}

		default:
			http.Error(w, "Unknown operation", http.StatusBadRequest)
		}
	})
}

// cacheHealthHandler checks cache health and connectivity.
func cacheHealthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetLogger("webserver.cache")

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if cacheManager == nil {
			http.Error(w, "Cache not initialized", http.StatusInternalServerError)
			return
		}

		// Test cache connectivity by doing a simple set/get operation
		testKey := "health-check"
		testValue := []byte("ok")

		err := cacheManager.SetAPIResponse(r.Context(), testKey, testValue)
		if err != nil {
			logger.Errorf("cache health check failed (set): %v", err)
			response := map[string]interface{}{
				"status":  "unhealthy",
				"message": "Failed to write to cache",
				"error":   err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(response)
			return
		}

		retrievedValue, err := cacheManager.GetAPIResponse(r.Context(), testKey)
		if err != nil {
			logger.Errorf("cache health check failed (get): %v", err)
			response := map[string]interface{}{
				"status":  "unhealthy",
				"message": "Failed to read from cache",
				"error":   err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(response)
			return
		}

		if string(retrievedValue) != string(testValue) {
			logger.Error("cache health check failed: value mismatch")
			response := map[string]interface{}{
				"status":  "unhealthy",
				"message": "Cache value mismatch",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Clean up test data
		cacheManager.Delete(r.Context(), "api:"+testKey)

		// Get stats for additional health info
		stats, err := cacheManager.Stats(r.Context())
		if err != nil {
			logger.Warnf("failed to get stats for health check: %v", err)
		}

		response := map[string]interface{}{
			"status":  "healthy",
			"message": "Cache is operational",
		}

		if stats != nil {
			response["stats"] = stats
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Errorf("failed to encode cache health response: %v", err)
		}
	})
}