package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SubtitleRecord struct {
	ID        int64
	File      string
	Language  string
	Service   string
	CreatedAt time.Time
}

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err := initSchema(db); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func initSchema(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS subtitles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        file TEXT NOT NULL,
        language TEXT NOT NULL,
        service TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL
    )`)
	return err
}

func InsertSubtitle(db *sql.DB, file, lang, service string) error {
	_, err := db.Exec(`INSERT INTO subtitles (file, language, service, created_at) VALUES (?, ?, ?, ?)`,
		file, lang, service, time.Now())
	return err
}

func ListSubtitles(db *sql.DB) ([]SubtitleRecord, error) {
	rows, err := db.Query(`SELECT id, file, language, service, created_at FROM subtitles ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recs []SubtitleRecord
	for rows.Next() {
		var r SubtitleRecord
		if err := rows.Scan(&r.ID, &r.File, &r.Language, &r.Service, &r.CreatedAt); err != nil {
			return nil, err
		}
		recs = append(recs, r)
	}
	return recs, rows.Err()
}

// DeleteSubtitle removes subtitle records matching file from the database.
func DeleteSubtitle(db *sql.DB, file string) error {
	_, err := db.Exec(`DELETE FROM subtitles WHERE file = ?`, file)
	return err
}
