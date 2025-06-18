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

// InsertTag adds a tag to the database.
func (p *PostgresStore) InsertTag(name string) error {
	_, err := p.db.Exec(`INSERT INTO tags (name, created_at) VALUES ($1, $2)`, name, time.Now())
	return err
}

// ListTags returns all tags.
func (p *PostgresStore) ListTags() ([]Tag, error) {
	rows, err := p.db.Query(`SELECT id, name, created_at FROM tags ORDER BY id`)
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

// DeleteTag removes a tag.
func (p *PostgresStore) DeleteTag(id int64) error {
	_, err := p.db.Exec(`DELETE FROM tags WHERE id = $1`, id)
	return err
}

// AssignTagToUser associates a tag with a user.
func (p *PostgresStore) AssignTagToUser(userID, tagID int64) error {
	_, err := p.db.Exec(`INSERT INTO user_tags (user_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, userID, tagID)
	return err
}

// RemoveTagFromUser removes a tag from a user.
func (p *PostgresStore) RemoveTagFromUser(userID, tagID int64) error {
	_, err := p.db.Exec(`DELETE FROM user_tags WHERE user_id = $1 AND tag_id = $2`, userID, tagID)
	return err
}

// ListTagsForUser returns tags associated with a user.
func (p *PostgresStore) ListTagsForUser(userID int64) ([]Tag, error) {
	rows, err := p.db.Query(`SELECT t.id, t.name, t.created_at FROM tags t JOIN user_tags u ON t.id = u.tag_id WHERE u.user_id = $1 ORDER BY t.id`, userID)
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
func (p *PostgresStore) AssignTagToMedia(mediaID, tagID int64) error {
	_, err := p.db.Exec(`INSERT INTO media_tags (media_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, mediaID, tagID)
	return err
}

// RemoveTagFromMedia removes a tag from a media item.
func (p *PostgresStore) RemoveTagFromMedia(mediaID, tagID int64) error {
	_, err := p.db.Exec(`DELETE FROM media_tags WHERE media_id = $1 AND tag_id = $2`, mediaID, tagID)
	return err
}

// ListTagsForMedia returns tags associated with a media item.
func (p *PostgresStore) ListTagsForMedia(mediaID int64) ([]Tag, error) {
	rows, err := p.db.Query(`SELECT t.id, t.name, t.created_at FROM tags t JOIN media_tags m ON t.id = m.tag_id WHERE m.media_id = $1 ORDER BY t.id`, mediaID)
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
