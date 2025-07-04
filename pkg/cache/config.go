// file: pkg/cache/config.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174008

package cache

import (
	"time"

	"github.com/spf13/viper"
)

func init() {
	// Ensure a default backend is set for tests and dev
	if viper.GetString("cache.backend") == "" {
		viper.SetDefault("cache.backend", "memory")
	}
}

// ConfigFromViper creates a cache configuration from viper settings.
// This allows the cache to be configured via command-line flags, environment
// variables, or configuration files.
func ConfigFromViper() (*Config, error) {
	config := DefaultConfig()

	// Override defaults only when values are set in viper
	if viper.IsSet("cache.backend") {
		config.Backend = viper.GetString("cache.backend")
	}
	if viper.IsSet("cache.memory.max_entries") {
		config.Memory.MaxEntries = viper.GetInt("cache.memory.max_entries")
	}
	if viper.IsSet("cache.memory.max_memory") {
		config.Memory.MaxMemory = viper.GetInt64("cache.memory.max_memory")
	}
	if viper.IsSet("cache.redis.address") {
		config.Redis.Address = viper.GetString("cache.redis.address")
	}
	if viper.IsSet("cache.redis.password") {
		config.Redis.Password = viper.GetString("cache.redis.password")
	}
	if viper.IsSet("cache.redis.database") {
		config.Redis.Database = viper.GetInt("cache.redis.database")
	}
	if viper.IsSet("cache.redis.pool_size") {
		config.Redis.PoolSize = viper.GetInt("cache.redis.pool_size")
	}
	if viper.IsSet("cache.redis.min_idle_conns") {
		config.Redis.MinIdleConns = viper.GetInt("cache.redis.min_idle_conns")
	}
	if viper.IsSet("cache.redis.key_prefix") {
		config.Redis.KeyPrefix = viper.GetString("cache.redis.key_prefix")
	}

	// Parse duration strings for memory config
	if defaultTTLStr := viper.GetString("cache.memory.default_ttl"); defaultTTLStr != "" {
		if duration, err := time.ParseDuration(defaultTTLStr); err == nil {
			config.Memory.DefaultTTL = duration
		} else {
			config.Memory.DefaultTTL = 1 * time.Hour // fallback
		}
	}

	if cleanupIntervalStr := viper.GetString("cache.memory.cleanup_interval"); cleanupIntervalStr != "" {
		if duration, err := time.ParseDuration(cleanupIntervalStr); err == nil {
			config.Memory.CleanupInterval = duration
		} else {
			config.Memory.CleanupInterval = 10 * time.Minute // fallback
		}
	}

	// Redis cache configuration

	// Parse duration strings for Redis config
	if dialTimeoutStr := viper.GetString("cache.redis.dial_timeout"); dialTimeoutStr != "" {
		if duration, err := time.ParseDuration(dialTimeoutStr); err == nil {
			config.Redis.DialTimeout = duration
		} else {
			config.Redis.DialTimeout = 5 * time.Second // fallback
		}
	}

	if readTimeoutStr := viper.GetString("cache.redis.read_timeout"); readTimeoutStr != "" {
		if duration, err := time.ParseDuration(readTimeoutStr); err == nil {
			config.Redis.ReadTimeout = duration
		} else {
			config.Redis.ReadTimeout = 3 * time.Second // fallback
		}
	}

	if writeTimeoutStr := viper.GetString("cache.redis.write_timeout"); writeTimeoutStr != "" {
		if duration, err := time.ParseDuration(writeTimeoutStr); err == nil {
			config.Redis.WriteTimeout = duration
		} else {
			config.Redis.WriteTimeout = 3 * time.Second // fallback
		}
	}

	// TTL configuration
	if providerTTLStr := viper.GetString("cache.ttls.provider_search_results"); providerTTLStr != "" {
		if duration, err := time.ParseDuration(providerTTLStr); err == nil {
			config.TTLs.ProviderSearchResults = duration
		} else {
			config.TTLs.ProviderSearchResults = 5 * time.Minute // fallback
		}
	}

	if tmdbTTLStr := viper.GetString("cache.ttls.tmdb_metadata"); tmdbTTLStr != "" {
		if duration, err := time.ParseDuration(tmdbTTLStr); err == nil {
			config.TTLs.TMDBMetadata = duration
		} else {
			config.TTLs.TMDBMetadata = 24 * time.Hour // fallback
		}
	}

	if translationTTLStr := viper.GetString("cache.ttls.translation_results"); translationTTLStr != "" {
		if duration, err := time.ParseDuration(translationTTLStr); err == nil {
			config.TTLs.TranslationResults = duration
		} else {
			config.TTLs.TranslationResults = 0 // permanent - fallback
		}
	}

	if sessionTTLStr := viper.GetString("cache.ttls.user_sessions"); sessionTTLStr != "" {
		if duration, err := time.ParseDuration(sessionTTLStr); err == nil {
			config.TTLs.UserSessions = duration
		} else {
			config.TTLs.UserSessions = 24 * time.Hour // fallback
		}
	}

	if apiTTLStr := viper.GetString("cache.ttls.api_responses"); apiTTLStr != "" {
		if duration, err := time.ParseDuration(apiTTLStr); err == nil {
			config.TTLs.APIResponses = duration
		} else {
			config.TTLs.APIResponses = 30 * time.Minute // fallback
		}
	}

	return config, nil
}

// NewFromViper creates a new cache instance using viper configuration.
func NewFromViper() (Cache, error) {
	config, err := ConfigFromViper()
	if err != nil {
		return nil, err
	}
	return New(config)
}

// NewManagerFromViper creates a new cache manager using viper configuration.
func NewManagerFromViper() (*Manager, error) {
	config, err := ConfigFromViper()
	if err != nil {
		return nil, err
	}
	return NewManager(config)
}
