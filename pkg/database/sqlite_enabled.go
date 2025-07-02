//go:build sqlite
// +build sqlite

// file: pkg/database/sqlite_enabled.go
// version: 1.0.0
// guid: 7e6f5a4b-3c2d-8e7f-1a0b-4c3d2e1f0a9b

package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// OpenSQLStore opens or creates an SQLite database and returns a SQLStore.
// This function is only available when building with the 'sqlite' tag.
func OpenSQLStore(path string) (*SQLStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database at %s: %w", path, err)
	}
	if err := initSchema(db); err != nil {
		return nil, err
	}
	return &SQLStore{db: db}, nil
}

// Open maintains backward compatibility by returning the raw *sql.DB.
// This function is only available when building with the 'sqlite' tag.
func Open(path string) (*sql.DB, error) {
	s, err := OpenSQLStore(path)
	if err != nil {
		return nil, err
	}
	return s.db, nil
}

func initSchema(db *sql.DB) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS subtitles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        file TEXT NOT NULL,
        video_file TEXT,
        release TEXT,
        language TEXT NOT NULL,
        service TEXT NOT NULL,
        embedded INTEGER NOT NULL DEFAULT 0,
        source_url TEXT,
        provider_metadata TEXT,
        confidence_score REAL,
        parent_id INTEGER,
        modification_type TEXT,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS downloads (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        file TEXT NOT NULL,
        video_file TEXT NOT NULL,
        provider TEXT NOT NULL,
        language TEXT NOT NULL,
        search_query TEXT,
        match_score REAL,
        download_attempts INTEGER DEFAULT 1,
        error_message TEXT,
        response_time_ms INTEGER,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS media_items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        path TEXT NOT NULL,
        title TEXT NOT NULL,
        season INTEGER,
        episode INTEGER,
        release_group TEXT,
        alt_titles TEXT,
        field_locks TEXT,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}
	// Add new subtitle metadata columns
	if err := addColumnIfNotExists(db, "subtitles", "source_url", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'source_url' to 'subtitles': %w", err)
	}
	if err := addColumnIfNotExists(db, "subtitles", "provider_metadata", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'provider_metadata' to 'subtitles': %w", err)
	}
	if err := addColumnIfNotExists(db, "subtitles", "confidence_score", "REAL"); err != nil {
		return fmt.Errorf("failed to add column 'confidence_score' to 'subtitles': %w", err)
	}
	if err := addColumnIfNotExists(db, "subtitles", "parent_id", "INTEGER"); err != nil {
		return fmt.Errorf("failed to add column 'parent_id' to 'subtitles': %w", err)
	}
	if err := addColumnIfNotExists(db, "subtitles", "modification_type", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'modification_type' to 'subtitles': %w", err)
	}

	// Add new download metadata columns
	if err := addColumnIfNotExists(db, "downloads", "search_query", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'search_query' to 'downloads': %w", err)
	}
	if err := addColumnIfNotExists(db, "downloads", "match_score", "REAL"); err != nil {
		return fmt.Errorf("failed to add column 'match_score' to 'downloads': %w", err)
	}
	if err := addColumnIfNotExists(db, "downloads", "download_attempts", "INTEGER DEFAULT 1"); err != nil {
		return fmt.Errorf("failed to add column 'download_attempts' to 'downloads': %w", err)
	}
	if err := addColumnIfNotExists(db, "downloads", "error_message", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'error_message' to 'downloads': %w", err)
	}
	if err := addColumnIfNotExists(db, "downloads", "response_time_ms", "INTEGER"); err != nil {
		return fmt.Errorf("failed to add column 'response_time_ms' to 'downloads': %w", err)
	}

	if err := addColumnIfNotExists(db, "media_items", "release_group", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'release_group' to 'media_items': %w", err)
	}
	if err := addColumnIfNotExists(db, "media_items", "alt_titles", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'alt_titles' to 'media_items': %w", err)
	}
	if err := addColumnIfNotExists(db, "media_items", "field_locks", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'field_locks' to 'media_items': %w", err)
	}

	// Subtitle sources table for tracking provider performance and metadata
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS subtitle_sources (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        source_hash TEXT UNIQUE NOT NULL,
        original_url TEXT NOT NULL,
        provider TEXT NOT NULL,
        title TEXT,
        release_info TEXT,
        file_size INTEGER,
        download_count INTEGER DEFAULT 0,
        success_count INTEGER DEFAULT 0,
        avg_rating REAL,
        last_seen TIMESTAMP,
        metadata TEXT,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS dashboard_prefs (
        user_id INTEGER PRIMARY KEY,
        layout TEXT NOT NULL,
        updated_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        email TEXT,
        role TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS api_keys (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        key TEXT UNIQUE NOT NULL,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS sessions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        token TEXT UNIQUE NOT NULL,
        expires_at TIMESTAMP NOT NULL,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS login_tokens (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        token TEXT UNIQUE NOT NULL,
        expires_at TIMESTAMP NOT NULL,
        used INTEGER NOT NULL DEFAULT 0,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS permissions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        role TEXT NOT NULL,
        permission TEXT NOT NULL
    )`); err != nil {
		return err
	}

	// Enhanced tags table with universal support for all entity types
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tags (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT UNIQUE NOT NULL,
        type TEXT NOT NULL DEFAULT 'user',
        entity_type TEXT DEFAULT 'all',
        color TEXT,
        description TEXT,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	// Monitored items table for automatic subtitle monitoring
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS monitored_items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        media_id TEXT NOT NULL,
        path TEXT NOT NULL,
        languages TEXT NOT NULL,
        last_checked TIMESTAMP,
        status TEXT NOT NULL DEFAULT 'pending',
        retry_count INTEGER NOT NULL DEFAULT 0,
        max_retries INTEGER NOT NULL DEFAULT 3,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	// Universal tag associations table for polymorphic relationships
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tag_associations (
        tag_id INTEGER NOT NULL,
        entity_type TEXT NOT NULL,
        entity_id TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        PRIMARY KEY (tag_id, entity_type, entity_id),
        FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
    )`); err != nil {
		return err
	}

	// Legacy tables for backward compatibility - these will be migrated
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS user_tags (
        user_id INTEGER NOT NULL,
        tag_id INTEGER NOT NULL,
        UNIQUE(user_id, tag_id)
    )`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS media_tags (
        media_id INTEGER NOT NULL,
        tag_id INTEGER NOT NULL,
        UNIQUE(media_id, tag_id)
    )`); err != nil {
		return err
	}

	// Language profiles table for Bazarr-style language management
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS language_profiles (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        config TEXT NOT NULL,
        cutoff_score INTEGER NOT NULL DEFAULT 80,
        is_default INTEGER NOT NULL DEFAULT 0,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	// Media profile assignments table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS media_profile_assignments (
        media_id TEXT PRIMARY KEY,
        profile_id TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

	// Seed default roles and permissions if empty
	var count int
	row := db.QueryRow(`SELECT COUNT(1) FROM permissions`)
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		if _, err := db.Exec(`INSERT INTO permissions (role, permission) VALUES
            ('admin', 'all'),
            ('user', 'read'),
            ('user', 'download')`); err != nil {
			return err
		}
	}

	// Seed default language profile if empty
	var profileCount int
	profileRow := db.QueryRow(`SELECT COUNT(1) FROM language_profiles`)
	if err := profileRow.Scan(&profileCount); err != nil {
		return err
	}
	if profileCount == 0 {
		if _, err := db.Exec(`INSERT INTO language_profiles (id, name, config, cutoff_score, is_default, created_at, updated_at) VALUES
            ('default', 'Default English', '[{"language":"en","priority":1,"forced":false,"hi":false}]', 75, TRUE, datetime('now'), datetime('now'))`); err != nil {
			return err
		}
	}

	return nil
}

// addColumnIfNotExists attempts to add a column to a table.
// It ignores the error if the column already exists.
func addColumnIfNotExists(db *sql.DB, table, column, typ string) error {
	stmt := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", table, column, typ)
	if _, err := db.Exec(stmt); err != nil {
		if strings.Contains(err.Error(), "duplicate column name") {
			return nil
		}
		return err
	}
	return nil
}
