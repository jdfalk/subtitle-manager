package database

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// PostgresStore implements SubtitleStore backed by PostgreSQL.
type PostgresStore struct {
	db *sql.DB
}

// OpenPostgresStore opens a connection to a PostgreSQL database using dsn.
// The database schema is created if it does not already exist.
func OpenPostgresStore(dsn string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := initPostgresSchema(db); err != nil {
		db.Close()
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func initPostgresSchema(db *sql.DB) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS subtitles (
            id SERIAL PRIMARY KEY,
            file TEXT NOT NULL,
            video_file TEXT,
            release TEXT,
            language TEXT NOT NULL,
            service TEXT NOT NULL,
            embedded BOOLEAN NOT NULL DEFAULT FALSE,
            created_at TIMESTAMP NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS downloads (
            id SERIAL PRIMARY KEY,
            file TEXT NOT NULL,
            video_file TEXT NOT NULL,
            provider TEXT NOT NULL,
            language TEXT NOT NULL,
            created_at TIMESTAMP NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS media_items (
            id SERIAL PRIMARY KEY,
            path TEXT NOT NULL,
            title TEXT NOT NULL,
            season INTEGER,
            episode INTEGER,
            created_at TIMESTAMP NOT NULL
        )`,
	}
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}
	return nil
}

// Close closes the underlying PostgreSQL connection.
func (p *PostgresStore) Close() error { return p.db.Close() }

// InsertSubtitle stores a subtitle record.
func (p *PostgresStore) InsertSubtitle(rec *SubtitleRecord) error {
	_, err := p.db.Exec(`INSERT INTO subtitles (file, video_file, release, language, service, embedded, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		rec.File, rec.VideoFile, rec.Release, rec.Language, rec.Service, rec.Embedded, time.Now())
	return err
}

// ListSubtitles retrieves stored subtitle records ordered by most recent.
func (p *PostgresStore) ListSubtitles() ([]SubtitleRecord, error) {
	rows, err := p.db.Query(`SELECT id, file, video_file, release, language, service, embedded, created_at FROM subtitles ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var recs []SubtitleRecord
	for rows.Next() {
		var r SubtitleRecord
		var id int64
		if err := rows.Scan(&id, &r.File, &r.VideoFile, &r.Release, &r.Language, &r.Service, &r.Embedded, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.ID = strconv.FormatInt(id, 10)
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// DeleteSubtitle removes subtitle records matching file.
func (p *PostgresStore) DeleteSubtitle(file string) error {
	_, err := p.db.Exec(`DELETE FROM subtitles WHERE file = $1`, file)
	return err
}

// InsertDownload stores a download record.
func (p *PostgresStore) InsertDownload(rec *DownloadRecord) error {
	_, err := p.db.Exec(`INSERT INTO downloads (file, video_file, provider, language, created_at) VALUES ($1,$2,$3,$4,$5)`,
		rec.File, rec.VideoFile, rec.Provider, rec.Language, time.Now())
	return err
}

// ListDownloads retrieves download records ordered by most recent.
func (p *PostgresStore) ListDownloads() ([]DownloadRecord, error) {
	rows, err := p.db.Query(`SELECT id, file, video_file, provider, language, created_at FROM downloads ORDER BY id DESC`)
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

// DeleteDownload removes download records matching file.
func (p *PostgresStore) DeleteDownload(file string) error {
	_, err := p.db.Exec(`DELETE FROM downloads WHERE file = $1`, file)
	return err
}

// InsertMediaItem stores a media library record.
func (p *PostgresStore) InsertMediaItem(rec *MediaItem) error {
	_, err := p.db.Exec(`INSERT INTO media_items (path, title, season, episode, created_at) VALUES ($1,$2,$3,$4,$5)`,
		rec.Path, rec.Title, rec.Season, rec.Episode, time.Now())
	return err
}

// ListMediaItems retrieves media items ordered by most recent.
func (p *PostgresStore) ListMediaItems() ([]MediaItem, error) {
	rows, err := p.db.Query(`SELECT id, path, title, season, episode, created_at FROM media_items ORDER BY id DESC`)
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

// DeleteMediaItem removes records matching path.
func (p *PostgresStore) DeleteMediaItem(path string) error {
	_, err := p.db.Exec(`DELETE FROM media_items WHERE path = $1`, path)
	return err
}
