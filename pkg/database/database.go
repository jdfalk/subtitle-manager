package database

import (
	"database/sql"
	"path/filepath"
	"strconv"
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
	ID        string
	Path      string
	Title     string
	Season    int
	Episode   int
	CreatedAt time.Time
}

// Tag represents a user or media tag used for language and provider preferences.
// Name is the unique tag identifier and CreatedAt records when the tag was added.
type Tag struct {
	ID        string
	Name      string
	CreatedAt time.Time
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
        created_at TIMESTAMP NOT NULL
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

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tags (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT UNIQUE NOT NULL,
        created_at TIMESTAMP NOT NULL
    )`); err != nil {
		return err
	}

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

// DeleteDownload removes download records matching file using a raw *sql.DB.
func DeleteDownload(db *sql.DB, file string) error {
	_, err := db.Exec(`DELETE FROM downloads WHERE file = ?`, file)
	return err
}

// InsertMediaItem stores a media library record.
func (s *SQLStore) InsertMediaItem(rec *MediaItem) error {
	_, err := s.db.Exec(`INSERT INTO media_items (path, title, season, episode, created_at) VALUES (?, ?, ?, ?, ?)`,
		rec.Path, rec.Title, rec.Season, rec.Episode, time.Now())
	return err
}

// ListMediaItems retrieves all media items sorted by creation time.
func (s *SQLStore) ListMediaItems() ([]MediaItem, error) {
	rows, err := s.db.Query(`SELECT id, path, title, season, episode, created_at FROM media_items ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []MediaItem
	for rows.Next() {
		var r MediaItem
		var id int64
		if err := rows.Scan(&id, &r.Path, &r.Title, &r.Season, &r.Episode, &r.CreatedAt); err != nil {
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

// Close closes the underlying SQLite database.
func (s *SQLStore) Close() error { return s.db.Close() }

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

// DeleteSubtitle removes subtitle records matching file from the database.
func DeleteSubtitle(db *sql.DB, file string) error {
	_, err := db.Exec(`DELETE FROM subtitles WHERE file = ?`, file)
	return err
}

// InsertMediaItem stores a media item using a raw *sql.DB.
func InsertMediaItem(db *sql.DB, path, title string, season, episode int) error {
	_, err := db.Exec(`INSERT INTO media_items (path, title, season, episode, created_at) VALUES (?, ?, ?, ?, ?)`,
		path, title, season, episode, time.Now())
	return err
}

// ListMediaItems retrieves media items using a raw *sql.DB.
func ListMediaItems(db *sql.DB) ([]MediaItem, error) {
	rows, err := db.Query(`SELECT id, path, title, season, episode, created_at FROM media_items ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MediaItem
	for rows.Next() {
		var it MediaItem
		var id int64
		if err := rows.Scan(&id, &it.Path, &it.Title, &it.Season, &it.Episode, &it.CreatedAt); err != nil {
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

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// InsertTag adds a new tag to the database.
func (s *SQLStore) InsertTag(name string) error {
	_, err := s.db.Exec(`INSERT INTO tags (name, created_at) VALUES (?, ?)`, name, time.Now())
	return err
}

// UpdateTag renames an existing tag.
func (s *SQLStore) UpdateTag(id int64, name string) error {
	_, err := s.db.Exec(`UPDATE tags SET name = ? WHERE id = ?`, name, id)
	return err
}

// ListTags returns all defined tags ordered by ID.
func (s *SQLStore) ListTags() ([]Tag, error) {
	rows, err := s.db.Query(`SELECT id, name, created_at FROM tags ORDER BY id`)
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

// DeleteTag removes a tag by ID.
func (s *SQLStore) DeleteTag(id int64) error {
	_, err := s.db.Exec(`DELETE FROM tags WHERE id = ?`, id)
	return err
}

// AssignTagToUser associates a tag with a user.
func (s *SQLStore) AssignTagToUser(userID, tagID int64) error {
	_, err := s.db.Exec(`INSERT OR IGNORE INTO user_tags (user_id, tag_id) VALUES (?, ?)`, userID, tagID)
	return err
}

// RemoveTagFromUser deletes a tag association for a user.
func (s *SQLStore) RemoveTagFromUser(userID, tagID int64) error {
	_, err := s.db.Exec(`DELETE FROM user_tags WHERE user_id = ? AND tag_id = ?`, userID, tagID)
	return err
}

// ListTagsForUser returns tags assigned to a user.
func (s *SQLStore) ListTagsForUser(userID int64) ([]Tag, error) {
	rows, err := s.db.Query(`SELECT t.id, t.name, t.created_at FROM tags t JOIN user_tags u ON t.id = u.tag_id WHERE u.user_id = ? ORDER BY t.id`, userID)
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

// InsertTag adds a tag using a raw *sql.DB.
func InsertTag(db *sql.DB, name string) error {
	_, err := db.Exec(`INSERT INTO tags (name, created_at) VALUES (?, ?)`, name, time.Now())
	return err
}

// UpdateTag updates the name for an existing tag using a raw *sql.DB.
func UpdateTag(db *sql.DB, id int64, name string) error {
	_, err := db.Exec(`UPDATE tags SET name = ? WHERE id = ?`, name, id)
	return err
}

// ListTags retrieves tags using a raw *sql.DB.
func ListTags(db *sql.DB) ([]Tag, error) {
	rows, err := db.Query(`SELECT id, name, created_at FROM tags ORDER BY id`)
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
