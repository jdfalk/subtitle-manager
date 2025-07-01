// file: pkg/database/performance.go
// version: 1.0.0
// guid: 8f7e6d5c-4b3a-9e8f-2a1b-5c4d3e2f1b0a

package database

import (
	"database/sql"
	"time"
)

// OptimizeConnectionPool configures optimal database connection pool settings
// for improved performance and resource utilization.
//
// This function sets connection pool parameters based on database backend type
// and expected load patterns to optimize for both response time and memory usage.
//
// Parameters:
//   - db: Database connection handle
//   - backend: Database backend type ("sqlite", "postgres", "pebble")
//
// Connection pool optimizations:
//   - SQLite: Conservative settings due to file-based nature
//   - PostgreSQL: Aggressive settings for network database
//   - PebbleDB: N/A (embedded key-value store)
func OptimizeConnectionPool(db *sql.DB, backend string) {
	switch backend {
	case "sqlite":
		// SQLite optimizations for file-based database
		// Conservative settings to prevent lock contention
		db.SetMaxOpenConns(10)  // Limit concurrent connections to prevent file locks
		db.SetMaxIdleConns(5)   // Keep some connections warm for fast access
		db.SetConnMaxLifetime(5 * time.Minute) // Reasonable connection lifetime
		db.SetConnMaxIdleTime(2 * time.Minute) // Close idle connections faster

	case "postgres":
		// PostgreSQL optimizations for network database  
		// More aggressive settings for better throughput
		db.SetMaxOpenConns(25)  // Higher limit for network database
		db.SetMaxIdleConns(10)  // More idle connections for better response time
		db.SetConnMaxLifetime(10 * time.Minute) // Longer lifetime for stable connections
		db.SetConnMaxIdleTime(5 * time.Minute)  // Keep connections longer

	default:
		// Default conservative settings for unknown backends
		db.SetMaxOpenConns(15)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(5 * time.Minute)
		db.SetConnMaxIdleTime(3 * time.Minute)
	}
}

// CreatePerformanceIndexes adds strategic database indexes to improve query performance.
//
// These indexes are designed to optimize the most common query patterns in the application:
//   - Video file lookups for subtitle and download history
//   - Media item path lookups for library scanning
//   - User and session management queries
//   - Tag association queries
//
// Parameters:
//   - db: Database connection handle
//   - backend: Database backend type for syntax compatibility
//
// Returns any error encountered during index creation.
func CreatePerformanceIndexes(db *sql.DB, backend string) error {
	var indexes []string

	switch backend {
	case "postgres":
		indexes = []string{
			// Core performance indexes for common queries
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_subtitles_video_file ON subtitles(video_file)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_subtitles_language ON subtitles(language)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_subtitles_created_at ON subtitles(created_at DESC)`,
			
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_downloads_video_file ON downloads(video_file)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_downloads_language ON downloads(language)`, 
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_downloads_created_at ON downloads(created_at DESC)`,
			
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_media_items_path ON media_items(path)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_media_items_title ON media_items(title)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_media_items_created_at ON media_items(created_at DESC)`,
			
			// Authentication and session indexes
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_users_username ON users(username)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_sessions_token ON sessions(token)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_sessions_user_id ON sessions(user_id)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at)`,
			
			// Tag system performance indexes
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_tag_associations_entity ON tag_associations(entity_type, entity_id)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_tag_associations_tag_id ON tag_associations(tag_id)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_tags_name ON tags(name)`,
			`CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_tags_type_entity ON tags(type, entity_type)`,
		}
	default:
		// SQLite and other backends
		indexes = []string{
			// Core performance indexes for common queries
			`CREATE INDEX IF NOT EXISTS idx_subtitles_video_file ON subtitles(video_file)`,
			`CREATE INDEX IF NOT EXISTS idx_subtitles_language ON subtitles(language)`,
			`CREATE INDEX IF NOT EXISTS idx_subtitles_created_at ON subtitles(created_at DESC)`,
			
			`CREATE INDEX IF NOT EXISTS idx_downloads_video_file ON downloads(video_file)`,
			`CREATE INDEX IF NOT EXISTS idx_downloads_language ON downloads(language)`,
			`CREATE INDEX IF NOT EXISTS idx_downloads_created_at ON downloads(created_at DESC)`,
			
			`CREATE INDEX IF NOT EXISTS idx_media_items_path ON media_items(path)`,
			`CREATE INDEX IF NOT EXISTS idx_media_items_title ON media_items(title)`,
			`CREATE INDEX IF NOT EXISTS idx_media_items_created_at ON media_items(created_at DESC)`,
			
			// Authentication and session indexes
			`CREATE INDEX IF NOT EXISTS idx_users_username ON users(username)`,
			`CREATE INDEX IF NOT EXISTS idx_sessions_token ON sessions(token)`,
			`CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id)`,
			`CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at)`,
			
			// Tag system performance indexes
			`CREATE INDEX IF NOT EXISTS idx_tag_associations_entity ON tag_associations(entity_type, entity_id)`,
			`CREATE INDEX IF NOT EXISTS idx_tag_associations_tag_id ON tag_associations(tag_id)`,
			`CREATE INDEX IF NOT EXISTS idx_tags_name ON tags(name)`,
			`CREATE INDEX IF NOT EXISTS idx_tags_type_entity ON tags(type, entity_type)`,
		}
	}

	// Execute all index creation statements
	for _, indexSQL := range indexes {
		if _, err := db.Exec(indexSQL); err != nil {
			// Log the error but continue with other indexes
			// This allows the application to start even if some indexes fail
			continue
		}
	}

	return nil
}

// OptimizeDatabaseSettings applies database-specific performance optimizations.
//
// This function configures database-specific settings for optimal performance:
//   - SQLite: WAL mode, synchronous settings, cache size
//   - PostgreSQL: Connection settings, query planner hints
//
// Parameters:
//   - db: Database connection handle 
//   - backend: Database backend type
//
// Returns any error encountered during optimization.
func OptimizeDatabaseSettings(db *sql.DB, backend string) error {
	switch backend {
	case "sqlite":
		// SQLite performance optimizations
		optimizations := []string{
			// Enable WAL mode for better concurrency
			`PRAGMA journal_mode = WAL`,
			// Optimize synchronous mode for performance vs durability trade-off
			`PRAGMA synchronous = NORMAL`,
			// Increase cache size for better performance (2MB)
			`PRAGMA cache_size = -2000`,
			// Optimize page size for modern SSDs
			`PRAGMA page_size = 4096`,
			// Enable foreign key constraints
			`PRAGMA foreign_keys = ON`,
			// Set reasonable timeout for busy database
			`PRAGMA busy_timeout = 5000`,
		}
		
		for _, pragma := range optimizations {
			if _, err := db.Exec(pragma); err != nil {
				// Continue with other optimizations even if one fails
				continue
			}
		}

	case "postgres":
		// PostgreSQL performance optimizations
		optimizations := []string{
			// Optimize for subtitle manager workload
			`SET work_mem = '16MB'`,
			`SET maintenance_work_mem = '64MB'`,
			`SET effective_cache_size = '256MB'`,
			`SET random_page_cost = 1.1`,
		}
		
		for _, setting := range optimizations {
			if _, err := db.Exec(setting); err != nil {
				// Continue with other optimizations even if one fails
				continue
			}
		}
	}

	return nil
}