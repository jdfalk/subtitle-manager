// Package database provides data storage and retrieval functionality for subtitle-manager.
// It supports SQLite and other database backends for managing subtitle metadata and operations.
//
// This package is the core data layer for subtitle-manager, handling persistence and queries.
// It defines the data models, database schema, and CRUD operations for subtitles, media items,
// downloads, and user data. It also manages tag-based organization and retrieval of entities.
//
// Usage:
//
//	import (
//	    "github.com/yourusername/subtitle-manager/database"
//	)
//
//	func main() {
//	    // Open the database
//	    db, err := database.Open("path/to/your/database/file")
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    defer db.Close()
//
//	    // Insert a new subtitle record
//	    err = db.InsertSubtitle(&database.SubtitleRecord{
//	        File:      "example.srt",
//	        VideoFile: "example.mp4",
//	        Release:   "Example Release",
//	        Language:  "en",
//	        Service:   "subscene",
//	        Embedded:  false,
//	    })
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    // List all subtitles
//	    subs, err := db.ListSubtitles()
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    for _, sub := range subs {
//	        fmt.Println(sub.File, sub.Language)
//	    }
//	}
package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

// GetDatabasePath returns the full database path based on the backend type.
// For SQLite, it combines db_path with sqlite3_filename.
// For other backends (pebble, postgres), it returns db_path as-is.
func GetDatabasePath() string {
	backend := viper.GetString("db_backend")
	dbPath := viper.GetString("db_path")

	if backend == "sqlite" {
		filename := viper.GetString("sqlite3_filename")
		return filepath.Join(dbPath, filename)
	}

	return dbPath
}

// GetDatabaseBackend returns the configured database backend
func GetDatabaseBackend() string {
	return viper.GetString("db_backend")
}

// SubtitleRecord represents a subtitle file that has been processed.
// VideoFile is the path to the media file the subtitle belongs to.
// Release denotes the original release name including group when known.
type SubtitleRecord struct {
	ID        string
	File      string
	VideoFile string
	Release   string
	Language  string
	Service   string
	Embedded  bool
	CreatedAt time.Time
}

// DownloadRecord represents a downloaded subtitle file.
// File is the path to the subtitle file stored on disk.
// VideoFile is the media file the subtitle corresponds to.
type DownloadRecord struct {
	ID        string
	File      string
	VideoFile string
	Provider  string
	Language  string
	CreatedAt time.Time
}

// MediaItem represents a video file discovered in the library.
// Path is the absolute location on disk. Title is the parsed show or movie name.
// Season and Episode provide optional numbering for TV episodes.
type MediaItem struct {
	ID           string
	Path         string
	Title        string
	Season       int
	Episode      int
	ReleaseGroup string
	AltTitles    string
	FieldLocks   string
	CreatedAt    time.Time
}

// Tag represents a universal tag that can be associated with any entity type.
// Supports media, users, providers, language profiles, and media profiles.
// Type indicates the tag classification (system, user, custom).
// EntityType specifies which types of entities this tag can be applied to.
// Color and Description provide additional metadata for UI display and organization.
type Tag struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`        // 'system', 'user', 'custom'
	EntityType  string    `json:"entity_type"` // 'media', 'language', 'movie', 'series', 'provider', 'user', 'all'
	Color       string    `json:"color"`       // hex color for UI display (optional)
	Description string    `json:"description"` // optional tag description
	CreatedAt   time.Time `json:"created_at"`
}

// TagAssociation represents a relationship between a tag and any entity.
// Uses polymorphic association pattern with EntityType and EntityID.
type TagAssociation struct {
	TagID      string    `json:"tag_id"`
	EntityType string    `json:"entity_type"` // 'movie', 'series', 'user', 'provider', etc.
	EntityID   string    `json:"entity_id"`   // polymorphic entity identifier
	CreatedAt  time.Time `json:"created_at"`
}

// TaggedEntity interface defines methods that all taggable entities must implement.
type TaggedEntity interface {
	GetEntityType() string
	GetEntityID() string
	GetTags() []string
	SetTags([]string)
}

// SQLStore implements SubtitleStore using an SQLite database.
type SQLStore struct {
	db *sql.DB
}

// OpenSQLStore opens or creates an SQLite database and returns a SQLStore.
func OpenSQLStore(path string) (*SQLStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err := initSchema(db); err != nil {
		db.Close()
		return nil, err
	}
	return &SQLStore{db: db}, nil
}

// Open maintains backward compatibility by returning the raw *sql.DB.
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
	if err := addColumnIfNotExists(db, "media_items", "release_group", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'release_group' to 'media_items': %w", err)
	}
	if err := addColumnIfNotExists(db, "media_items", "alt_titles", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'alt_titles' to 'media_items': %w", err)
	}
	if err := addColumnIfNotExists(db, "media_items", "field_locks", "TEXT"); err != nil {
		return fmt.Errorf("failed to add column 'field_locks' to 'media_items': %w", err)
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

	return nil
}

// InsertDownload stores a download record.
func (s *SQLStore) InsertDownload(rec *DownloadRecord) error {
	_, err := s.db.Exec(`INSERT INTO downloads (file, video_file, provider, language, created_at) VALUES (?, ?, ?, ?, ?)`,
		rec.File, rec.VideoFile, rec.Provider, rec.Language, time.Now())
	return err
}

// ListDownloads retrieves download records ordered by most recent.
func (s *SQLStore) ListDownloads() ([]DownloadRecord, error) {
	rows, err := s.db.Query(`SELECT id, file, video_file, provider, language, created_at FROM downloads ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []DownloadRecord
	for rows.Next() {
		var r DownloadRecord
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Provider, &r.Language, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// ListDownloadsByVideo retrieves download history for a specific video file.
func (s *SQLStore) ListDownloadsByVideo(video string) ([]DownloadRecord, error) {
	rows, err := s.db.Query(`SELECT id, file, video_file, provider, language, created_at FROM downloads WHERE video_file = ? ORDER BY id DESC`, video)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []DownloadRecord
	for rows.Next() {
		var r DownloadRecord
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Provider, &r.Language, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// DeleteDownload removes download records matching file from the database.
func (s *SQLStore) DeleteDownload(file string) error {
	_, err := s.db.Exec(`DELETE FROM downloads WHERE file = ?`, file)
	return err
}

// (Removed duplicate InsertMediaItem method - keeping the version with created_at)

// (Removed duplicate ListMediaItems method - keeping the version with created_at)

// (Removed duplicate DeleteMediaItem method - keeping the version with created_at)

// InsertDownload stores a download record using a raw *sql.DB.
func InsertDownload(db *sql.DB, file, video, provider, lang string) error {
	_, err := db.Exec(`INSERT INTO downloads (file, video_file, provider, language, created_at) VALUES (?, ?, ?, ?, ?)`,
		file, video, provider, lang, time.Now())
	return err
}

// ListDownloads retrieves download records using a raw *sql.DB.
func ListDownloads(db *sql.DB) ([]DownloadRecord, error) {
	rows, err := db.Query(`SELECT id, file, video_file, provider, language, created_at FROM downloads ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []DownloadRecord
	for rows.Next() {
		var r DownloadRecord
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Provider, &r.Language, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// ListDownloadsByVideo retrieves download history for a specific video file using a raw *sql.DB.
func ListDownloadsByVideo(db *sql.DB, video string) ([]DownloadRecord, error) {
	rows, err := db.Query(`SELECT id, file, video_file, provider, language, created_at FROM downloads WHERE video_file = ? ORDER BY id DESC`, video)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []DownloadRecord
	for rows.Next() {
		var r DownloadRecord
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Provider, &r.Language, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// DeleteDownload removes download records matching file using a raw *sql.DB.
func DeleteDownload(db *sql.DB, file string) error {
	_, err := db.Exec(`DELETE FROM downloads WHERE file = ?`, file)
	return err
}

// InsertMediaItem stores a media library record.
func (s *SQLStore) InsertMediaItem(rec *MediaItem) error {
	_, err := s.db.Exec(`INSERT INTO media_items (path, title, season, episode, release_group, alt_titles, field_locks, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		rec.Path, rec.Title, rec.Season, rec.Episode, rec.ReleaseGroup, rec.AltTitles, rec.FieldLocks, time.Now())
	return err
}

// ListMediaItems retrieves all media items sorted by creation time.
func (s *SQLStore) ListMediaItems() ([]MediaItem, error) {
	rows, err := s.db.Query(`SELECT id, path, title, season, episode, release_group, alt_titles, field_locks, created_at FROM media_items ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []MediaItem
	for rows.Next() {
		var r MediaItem
		var id int64
		if err := rows.Scan(&id, &r.Path, &r.Title, &r.Season, &r.Episode, &r.ReleaseGroup, &r.AltTitles, &r.FieldLocks, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// DeleteMediaItem removes a media item matching path.
func (s *SQLStore) DeleteMediaItem(path string) error {
	_, err := s.db.Exec(`DELETE FROM media_items WHERE path = ?`, path)
	return err
}

// CountSubtitles returns the number of subtitle records.
func (s *SQLStore) CountSubtitles() (int, error) {
	row := s.db.QueryRow(`SELECT COUNT(*) FROM subtitles`)
	var n int
	err := row.Scan(&n)
	return n, err
}

// CountDownloads returns the number of download records.
func (s *SQLStore) CountDownloads() (int, error) {
	row := s.db.QueryRow(`SELECT COUNT(*) FROM downloads`)
	var n int
	err := row.Scan(&n)
	return n, err
}

// CountMediaItems returns the number of media items.
func (s *SQLStore) CountMediaItems() (int, error) {
	row := s.db.QueryRow(`SELECT COUNT(*) FROM media_items`)
	var n int
	err := row.Scan(&n)
	return n, err
}

// SetMediaReleaseGroup updates the release group for a media item.
func (s *SQLStore) SetMediaReleaseGroup(path, group string) error {
	_, err := s.db.Exec(`UPDATE media_items SET release_group = ? WHERE path = ?`, group, path)
	return err
}

// SetMediaAltTitles updates alternate titles for a media item.
func (s *SQLStore) SetMediaAltTitles(path string, titles []string) error {
	data, err := json.Marshal(titles)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(`UPDATE media_items SET alt_titles = ? WHERE path = ?`, string(data), path)
	return err
}

// SetMediaFieldLocks updates field locks for a media item.
func (s *SQLStore) SetMediaFieldLocks(path string, locks string) error {
	_, err := s.db.Exec(`UPDATE media_items SET field_locks = ? WHERE path = ?`, locks, path)
	return err
}

// SetMediaTitle updates the title for a media item.
func (s *SQLStore) SetMediaTitle(path, title string) error {
	_, err := s.db.Exec(`UPDATE media_items SET title = ? WHERE path = ?`, title, path)
	return err
}

// DB returns the underlying *sql.DB for compatibility with existing code.
func (s *SQLStore) DB() *sql.DB {
	return s.db
}

// Close closes the underlying database connection.
func (s *SQLStore) Close() error {
	return s.db.Close()
}

// InsertSubtitle stores a new subtitle record with associated metadata.
func (s *SQLStore) InsertSubtitle(rec *SubtitleRecord) error {
	_, err := s.db.Exec(`INSERT INTO subtitles (file, video_file, release, language, service, embedded, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		rec.File, rec.VideoFile, rec.Release, rec.Language, rec.Service, boolToInt(rec.Embedded), time.Now())
	return err
}

// ListSubtitles retrieves subtitle records ordered by most recent.
func (s *SQLStore) ListSubtitles() ([]SubtitleRecord, error) {
	rows, err := s.db.Query(`SELECT id, file, video_file, release, language, service, embedded, created_at FROM subtitles ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recs []SubtitleRecord
	for rows.Next() {
		var r SubtitleRecord
		var embedded int
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Release, &r.Language, &r.Service, &embedded, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		r.Embedded = embedded == 1
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// ListSubtitlesByVideo retrieves subtitle history for a specific video file.
func (s *SQLStore) ListSubtitlesByVideo(video string) ([]SubtitleRecord, error) {
	rows, err := s.db.Query(`SELECT id, file, video_file, release, language, service, embedded, created_at FROM subtitles WHERE video_file = ? ORDER BY id DESC`, video)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []SubtitleRecord
	for rows.Next() {
		var r SubtitleRecord
		var embedded int
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Release, &r.Language, &r.Service, &embedded, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		r.Embedded = embedded == 1
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// DeleteSubtitle removes subtitle records matching file from the database.
func (s *SQLStore) DeleteSubtitle(file string) error {
	_, err := s.db.Exec(`DELETE FROM subtitles WHERE file = ?`, file)
	return err
}

// InsertSubtitle stores a new subtitle record with associated metadata.
func InsertSubtitle(db *sql.DB, file, video, lang, service, release string, embedded bool) error {
	_, err := db.Exec(`INSERT INTO subtitles (file, video_file, release, language, service, embedded, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		file, video, release, lang, service, boolToInt(embedded), time.Now())
	return err
}

func ListSubtitles(db *sql.DB) ([]SubtitleRecord, error) {
	rows, err := db.Query(`SELECT id, file, video_file, release, language, service, embedded, created_at FROM subtitles ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recs []SubtitleRecord
	for rows.Next() {
		var r SubtitleRecord
		var embedded int
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Release, &r.Language, &r.Service, &embedded, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		r.Embedded = embedded == 1
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// ListSubtitlesByVideo retrieves subtitle records for a specific video file using a raw *sql.DB.
func ListSubtitlesByVideo(db *sql.DB, video string) ([]SubtitleRecord, error) {
	rows, err := db.Query(`SELECT id, file, video_file, release, language, service, embedded, created_at FROM subtitles WHERE video_file = ? ORDER BY id DESC`, video)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []SubtitleRecord
	for rows.Next() {
		var r SubtitleRecord
		var embedded int
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Release, &r.Language, &r.Service, &embedded, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		r.Embedded = embedded == 1
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// DeleteSubtitle removes subtitle records matching file from the database.
func DeleteSubtitle(db *sql.DB, file string) error {
	_, err := db.Exec(`DELETE FROM subtitles WHERE file = ?`, file)
	return err
}

// InsertMediaItem stores a media item using a raw *sql.DB.
func InsertMediaItem(db *sql.DB, path, title string, season, episode int) error {
	_, err := db.Exec(`INSERT INTO media_items (path, title, season, episode, release_group, alt_titles, field_locks, created_at) VALUES (?, ?, ?, ?, '', '', '', ?)`,
		path, title, season, episode, time.Now())
	return err
}

// ListMediaItems retrieves media items using a raw *sql.DB.
func ListMediaItems(db *sql.DB) ([]MediaItem, error) {
	rows, err := db.Query(`SELECT id, path, title, season, episode, release_group, alt_titles, field_locks, created_at FROM media_items ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MediaItem
	for rows.Next() {
		var it MediaItem
		var id int64
		if err := rows.Scan(&id, &it.Path, &it.Title, &it.Season, &it.Episode, &it.ReleaseGroup, &it.AltTitles, &it.FieldLocks, &it.CreatedAt); err != nil {
			return nil, err
		}
		it.ID = strconv.FormatInt(id, 10)
		items = append(items, it)
	}
	return items, rows.Err()
}

// DeleteMediaItem removes media items with the given path using a raw *sql.DB.
func DeleteMediaItem(db *sql.DB, path string) error {
	_, err := db.Exec(`DELETE FROM media_items WHERE path = ?`, path)
	return err
}

// SetMediaReleaseGroup updates the release group for a media item using a raw database handle.
func SetMediaReleaseGroup(db *sql.DB, path, group string) error {
	_, err := db.Exec(`UPDATE media_items SET release_group = ? WHERE path = ?`, group, path)
	return err
}

// SetMediaAltTitles updates alternate titles for a media item using a raw database handle.
func SetMediaAltTitles(db *sql.DB, path string, titles []string) error {
	data, err := json.Marshal(titles)
	if err != nil {
		return err
	}
	_, err = db.Exec(`UPDATE media_items SET alt_titles = ? WHERE path = ?`, string(data), path)
	return err
}

// SetMediaFieldLocks updates field locks for a media item using a raw database handle.
func SetMediaFieldLocks(db *sql.DB, path, locks string) error {
	_, err := db.Exec(`UPDATE media_items SET field_locks = ? WHERE path = ?`, locks, path)
	return err
}

// SetMediaTitle updates the title for a media item using a raw database handle.
func SetMediaTitle(db *sql.DB, path, title string) error {
	_, err := db.Exec(`UPDATE media_items SET title = ? WHERE path = ?`, title, path)
	return err
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

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// InsertTag adds a new tag to the database with enhanced metadata support.
func (s *SQLStore) InsertTag(name string) error {
	_, err := s.db.Exec(`INSERT INTO tags (name, type, entity_type, created_at) VALUES (?, 'user', 'all', ?)`, name, time.Now())
	return err
}

// InsertTagWithMetadata adds a new tag with full metadata support.
func (s *SQLStore) InsertTagWithMetadata(name, tagType, entityType, color, description string) error {
	_, err := s.db.Exec(`INSERT INTO tags (name, type, entity_type, color, description, created_at) VALUES (?, ?, ?, ?, ?, ?)`,
		name, tagType, entityType, color, description, time.Now())
	return err
}

// UpdateTag renames an existing tag and optionally updates metadata.
func (s *SQLStore) UpdateTag(id int64, name string) error {
	_, err := s.db.Exec(`UPDATE tags SET name = ? WHERE id = ?`, name, id)
	return err
}

// UpdateTagWithMetadata updates a tag with enhanced metadata.
func (s *SQLStore) UpdateTagWithMetadata(id int64, name, tagType, entityType, color, description string) error {
	query := `UPDATE tags SET name = ?, type = ?, entity_type = ?, color = ?, description = ? WHERE id = ?`
	_, err := s.db.Exec(query, name, tagType, entityType, color, description, id)
	return err
}

// ListTags returns all defined tags ordered by ID with full metadata.
func (s *SQLStore) ListTags() ([]Tag, error) {
	rows, err := s.db.Query(`SELECT id, name, type, entity_type, color, description, created_at FROM tags ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Tag
	for rows.Next() {
		var t Tag
		var id int64
		var color, description sql.NullString
		if err := rows.Scan(&id, &t.Name, &t.Type, &t.EntityType, &color, &description, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		if color.Valid {
			t.Color = color.String
		}
		if description.Valid {
			t.Description = description.String
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// DeleteTag removes a tag by ID.
func (s *SQLStore) DeleteTag(id int64) error {
	_, err := s.db.Exec(`DELETE FROM tags WHERE id = ?`, id)
	return err
}

// Universal tag association methods

// AssignTagToEntity creates a universal tag association with any entity type.
func (s *SQLStore) AssignTagToEntity(tagID int64, entityType, entityID string) error {
	_, err := s.db.Exec(`INSERT OR IGNORE INTO tag_associations (tag_id, entity_type, entity_id, created_at) VALUES (?, ?, ?, ?)`,
		tagID, entityType, entityID, time.Now())
	return err
}

// RemoveTagFromEntity removes a universal tag association.
func (s *SQLStore) RemoveTagFromEntity(tagID int64, entityType, entityID string) error {
	_, err := s.db.Exec(`DELETE FROM tag_associations WHERE tag_id = ? AND entity_type = ? AND entity_id = ?`,
		tagID, entityType, entityID)
	return err
}

// ListTagsForEntity returns all tags associated with a specific entity.
func (s *SQLStore) ListTagsForEntity(entityType, entityID string) ([]Tag, error) {
	rows, err := s.db.Query(`
		SELECT t.id, t.name, t.type, t.entity_type, t.color, t.description, t.created_at
		FROM tags t
		JOIN tag_associations ta ON t.id = ta.tag_id
		WHERE ta.entity_type = ? AND ta.entity_id = ?
		ORDER BY t.id`, entityType, entityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Tag
	for rows.Next() {
		var t Tag
		var id int64
		var color, description sql.NullString
		if err := rows.Scan(&id, &t.Name, &t.Type, &t.EntityType, &color, &description, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		if color.Valid {
			t.Color = color.String
		}
		if description.Valid {
			t.Description = description.String
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// ListEntitiesWithTag returns all entity IDs of a specific type that have the given tag.
func (s *SQLStore) ListEntitiesWithTag(tagID int64, entityType string) ([]string, error) {
	rows, err := s.db.Query(`SELECT entity_id FROM tag_associations WHERE tag_id = ? AND entity_type = ? ORDER BY entity_id`,
		tagID, entityType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var entityID string
		if err := rows.Scan(&entityID); err != nil {
			return nil, err
		}
		out = append(out, entityID)
	}
	return out, rows.Err()
}

// BulkAssignTags assigns multiple tags to multiple entities in a single transaction.
func (s *SQLStore) BulkAssignTags(tagIDs []int64, entities []TagAssociation) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO tag_associations (tag_id, entity_type, entity_id, created_at) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now()
	for _, tagID := range tagIDs {
		for _, entity := range entities {
			if _, err := stmt.Exec(tagID, entity.EntityType, entity.EntityID, now); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

// Legacy compatibility methods (will be deprecated)

// AssignTagToUser associates a tag with a user (legacy method).
func (s *SQLStore) AssignTagToUser(userID, tagID int64) error {
	// Use both old and new systems for compatibility
	if err := s.AssignTagToEntity(tagID, "user", strconv.FormatInt(userID, 10)); err != nil {
		return err
	}
	_, err := s.db.Exec(`INSERT OR IGNORE INTO user_tags (user_id, tag_id) VALUES (?, ?)`, userID, tagID)
	return err
}

// RemoveTagFromUser deletes a tag association for a user (legacy method).
func (s *SQLStore) RemoveTagFromUser(userID, tagID int64) error {
	// Remove from both old and new systems for compatibility
	_ = s.RemoveTagFromEntity(tagID, "user", strconv.FormatInt(userID, 10))
	_, err := s.db.Exec(`DELETE FROM user_tags WHERE user_id = ? AND tag_id = ?`, userID, tagID)
	return err
}

// ListTagsForUser returns tags assigned to a user (enhanced legacy method).
func (s *SQLStore) ListTagsForUser(userID int64) ([]Tag, error) {
	// Prefer new system, fallback to old system for compatibility
	tags, err := s.ListTagsForEntity("user", strconv.FormatInt(userID, 10))
	if err == nil && len(tags) > 0 {
		return tags, nil
	}

	// Fallback to legacy table
	rows, err := s.db.Query(`
		SELECT t.id, t.name, t.type, t.entity_type, t.color, t.description, t.created_at
		FROM tags t
		JOIN user_tags u ON t.id = u.tag_id
		WHERE u.user_id = ?
		ORDER BY t.id`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Tag
	for rows.Next() {
		var t Tag
		var id int64
		var color, description sql.NullString
		if err := rows.Scan(&id, &t.Name, &t.Type, &t.EntityType, &color, &description, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		if color.Valid {
			t.Color = color.String
		}
		if description.Valid {
			t.Description = description.String
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// AssignTagToMedia associates a tag with a media item.
func (s *SQLStore) AssignTagToMedia(mediaID, tagID int64) error {
	_, err := s.db.Exec(`INSERT OR IGNORE INTO media_tags (media_id, tag_id) VALUES (?, ?)`, mediaID, tagID)
	return err
}

// RemoveTagFromMedia deletes a tag association for a media item.
func (s *SQLStore) RemoveTagFromMedia(mediaID, tagID int64) error {
	_, err := s.db.Exec(`DELETE FROM media_tags WHERE media_id = ? AND tag_id = ?`, mediaID, tagID)
	return err
}

// ListTagsForMedia returns tags assigned to a media item.
func (s *SQLStore) ListTagsForMedia(mediaID int64) ([]Tag, error) {
	rows, err := s.db.Query(`SELECT t.id, t.name, t.created_at FROM tags t JOIN media_tags m ON t.id = m.tag_id WHERE m.media_id = ? ORDER BY t.id`, mediaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Tag
	for rows.Next() {
		var t Tag
		var id int64
		if err := rows.Scan(&id, &t.Name, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		out = append(out, t)
	}
	return out, rows.Err()
}

// InsertTag adds a tag using a raw *sql.DB (legacy compatibility).
func InsertTag(db *sql.DB, name string) error {
	_, err := db.Exec(`INSERT INTO tags (name, type, entity_type, created_at) VALUES (?, 'user', 'all', ?)`, name, time.Now())
	return err
}

// UpdateTag updates the name for an existing tag using a raw *sql.DB.
func UpdateTag(db *sql.DB, id int64, name string) error {
	_, err := db.Exec(`UPDATE tags SET name = ? WHERE id = ?`, name, id)
	return err
}

// ListTags retrieves tags using a raw *sql.DB with enhanced metadata.
func ListTags(db *sql.DB) ([]Tag, error) {
	rows, err := db.Query(`SELECT id, name, type, entity_type, color, description, created_at FROM tags ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Tag
	for rows.Next() {
		var t Tag
		var id int64
		var color, description sql.NullString
		if err := rows.Scan(&id, &t.Name, &t.Type, &t.EntityType, &color, &description, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		if color.Valid {
			t.Color = color.String
		}
		if description.Valid {
			t.Description = description.String
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// DeleteTag removes a tag by ID using a raw *sql.DB.
func DeleteTag(db *sql.DB, id int64) error {
	_, err := db.Exec(`DELETE FROM tags WHERE id = ?`, id)
	return err
}

// AssignTagToUser associates a tag with a user using a raw *sql.DB.
func AssignTagToUser(db *sql.DB, userID, tagID int64) error {
	_, err := db.Exec(`INSERT OR IGNORE INTO user_tags (user_id, tag_id) VALUES (?, ?)`, userID, tagID)
	return err
}

// RemoveTagFromUser removes a tag from a user using a raw *sql.DB.
func RemoveTagFromUser(db *sql.DB, userID, tagID int64) error {
	_, err := db.Exec(`DELETE FROM user_tags WHERE user_id = ? AND tag_id = ?`, userID, tagID)
	return err
}

// ListTagsForUser retrieves tags for a user using a raw *sql.DB.
func ListTagsForUser(db *sql.DB, userID int64) ([]Tag, error) {
	rows, err := db.Query(`SELECT t.id, t.name, t.created_at FROM tags t JOIN user_tags u ON t.id = u.tag_id WHERE u.user_id = ? ORDER BY t.id`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Tag
	for rows.Next() {
		var t Tag
		var id int64
		if err := rows.Scan(&id, &t.Name, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		out = append(out, t)
	}
	return out, rows.Err()
}

// AssignTagToMedia associates a tag with a media item using a raw *sql.DB.
func AssignTagToMedia(db *sql.DB, mediaID, tagID int64) error {
	_, err := db.Exec(`INSERT OR IGNORE INTO media_tags (media_id, tag_id) VALUES (?, ?)`, mediaID, tagID)
	return err
}

// RemoveTagFromMedia deletes a tag from a media item using a raw *sql.DB.
func RemoveTagFromMedia(db *sql.DB, mediaID, tagID int64) error {
	_, err := db.Exec(`DELETE FROM media_tags WHERE media_id = ? AND tag_id = ?`, mediaID, tagID)
	return err
}

// ListTagsForMedia retrieves tags for a media item using a raw *sql.DB.
func ListTagsForMedia(db *sql.DB, mediaID int64) ([]Tag, error) {
	rows, err := db.Query(`SELECT t.id, t.name, t.created_at FROM tags t JOIN media_tags m ON t.id = m.tag_id WHERE m.media_id = ? ORDER BY t.id`, mediaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Tag
	for rows.Next() {
		var t Tag
		var id int64
		if err := rows.Scan(&id, &t.Name, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		out = append(out, t)
	}
	return out, rows.Err()
}

// Universal tagging methods for raw *sql.DB access

// AssignTagToEntity creates a universal tag association using raw *sql.DB.
func AssignTagToEntity(db *sql.DB, tagID int64, entityType, entityID string) error {
	_, err := db.Exec(`INSERT OR IGNORE INTO tag_associations (tag_id, entity_type, entity_id, created_at) VALUES (?, ?, ?, ?)`,
		tagID, entityType, entityID, time.Now())
	return err
}

// RemoveTagFromEntity removes a universal tag association using raw *sql.DB.
func RemoveTagFromEntity(db *sql.DB, tagID int64, entityType, entityID string) error {
	_, err := db.Exec(`DELETE FROM tag_associations WHERE tag_id = ? AND entity_type = ? AND entity_id = ?`,
		tagID, entityType, entityID)
	return err
}

// ListTagsForEntity returns all tags for a specific entity using raw *sql.DB.
func ListTagsForEntity(db *sql.DB, entityType, entityID string) ([]Tag, error) {
	rows, err := db.Query(`
		SELECT t.id, t.name, t.type, t.entity_type, t.color, t.description, t.created_at
		FROM tags t
		JOIN tag_associations ta ON t.id = ta.tag_id
		WHERE ta.entity_type = ? AND ta.entity_id = ?
		ORDER BY t.id`, entityType, entityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Tag
	for rows.Next() {
		var t Tag
		var id int64
		var color, description sql.NullString
		if err := rows.Scan(&id, &t.Name, &t.Type, &t.EntityType, &color, &description, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		if color.Valid {
			t.Color = color.String
		}
		if description.Valid {
			t.Description = description.String
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// GetMediaIDByPath returns the media item ID for the given file path.
// It returns sql.ErrNoRows if the path has not been indexed.
func GetMediaIDByPath(db *sql.DB, path string) (int64, error) {
	var id int64
	err := db.QueryRow(`SELECT id FROM media_items WHERE path = ?`, path).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// EnsureMediaItem retrieves the ID for path, creating a new record if needed.
func EnsureMediaItem(db *sql.DB, path string) (int64, error) {
	id, err := GetMediaIDByPath(db, path)
	if err == nil {
		return id, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	res, err := db.Exec(`INSERT INTO media_items (path, title, season, episode, release_group, alt_titles, field_locks, created_at) VALUES (?, ?, 0, 0, '', '', '', ?)`,
		path, filepath.Base(path), time.Now())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// SetDashboardLayout stores the widget layout JSON for a user.
//
// Parameters:
//
//	db - active database handle
//	userID - identifier for the user
//	layout - JSON string describing widget placement
//
// Returns any error encountered while saving.
func SetDashboardLayout(db *sql.DB, userID int64, layout string) error {
	_, err := db.Exec(`INSERT INTO dashboard_prefs (user_id, layout, updated_at) VALUES (?, ?, ?)
        ON CONFLICT(user_id) DO UPDATE SET layout=excluded.layout, updated_at=excluded.updated_at`,
		userID, layout, time.Now())
	return err
}

// GetDashboardLayout retrieves the widget layout JSON for a user.
//
// Parameters:
//
//	db - active database handle
//	userID - identifier for the user
//
// Returns the stored layout JSON string and any error encountered. If no layout
// exists, the string will be empty and error will be nil.
func GetDashboardLayout(db *sql.DB, userID int64) (string, error) {
	row := db.QueryRow(`SELECT layout FROM dashboard_prefs WHERE user_id = ?`, userID)
	var layout string
	if err := row.Scan(&layout); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}
	return layout, nil
}
