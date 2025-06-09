package database

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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
	ID      string
	Path    string
	Title   string
	Season  int
	Episode int
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
        episode INTEGER
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

// InsertMediaItem adds a library entry to the media_items table.
func (s *SQLStore) InsertMediaItem(item *MediaItem) error {
	_, err := s.db.Exec(`INSERT INTO media_items (path, title, season, episode) VALUES (?, ?, ?, ?)`,
		item.Path, item.Title, item.Season, item.Episode)
	return err
}

// ListMediaItems retrieves all stored media items ordered by id descending.
func (s *SQLStore) ListMediaItems() ([]MediaItem, error) {
	rows, err := s.db.Query(`SELECT id, path, title, season, episode FROM media_items ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MediaItem
	for rows.Next() {
		var it MediaItem
		var id int64
		if err := rows.Scan(&id, &it.Path, &it.Title, &it.Season, &it.Episode); err != nil {
			return nil, err
		}
		it.ID = strconv.FormatInt(id, 10)
		items = append(items, it)
	}
	return items, rows.Err()
}

// DeleteMediaItem removes all records matching path from the database.
func (s *SQLStore) DeleteMediaItem(path string) error {
	_, err := s.db.Exec(`DELETE FROM media_items WHERE path = ?`, path)
	return err
}

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

// InsertMediaItem stores a library item using a raw *sql.DB.
func InsertMediaItem(db *sql.DB, path, title string, season, episode int) error {
	_, err := db.Exec(`INSERT INTO media_items (path, title, season, episode) VALUES (?, ?, ?, ?)`,
		path, title, season, episode)
	return err
}

// ListMediaItems retrieves media items using a raw *sql.DB.
func ListMediaItems(db *sql.DB) ([]MediaItem, error) {
	rows, err := db.Query(`SELECT id, path, title, season, episode FROM media_items ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MediaItem
	for rows.Next() {
		var it MediaItem
		var id int64
		if err := rows.Scan(&id, &it.Path, &it.Title, &it.Season, &it.Episode); err != nil {
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
