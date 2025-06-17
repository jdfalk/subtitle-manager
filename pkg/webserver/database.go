package webserver

import (
	"archive/tar"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/backups"
	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// dirSize returns the total size of files under path.
func dirSize(path string) int64 {
	var size int64
	_ = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size
}

// databaseInfoHandler returns basic information about the configured database.
func databaseInfoHandler(db *sql.DB) http.Handler {
	type info struct {
		Type      string `json:"type"`
		Version   string `json:"version"`
		Size      int64  `json:"size"`
		Path      string `json:"path"`
		Connected bool   `json:"connected"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backend := database.GetDatabaseBackend()
		path := database.GetDatabasePath()
		out := info{Type: backend, Path: path, Connected: true}
		switch backend {
		case "sqlite":
			row := db.QueryRow(`SELECT sqlite_version()`)
			_ = row.Scan(&out.Version)
			if fi, err := os.Stat(path); err == nil {
				out.Size = fi.Size()
			}
		case "pebble":
			out.Version = "pebble"
			out.Size = dirSize(path)
		case "postgres":
			pdb, err := sql.Open("postgres", path)
			if err != nil {
				out.Connected = false
				break
			}
			defer pdb.Close()
			row := pdb.QueryRow(`SHOW server_version`)
			_ = row.Scan(&out.Version)
			row = pdb.QueryRow(`SELECT pg_database_size(current_database())`)
			_ = row.Scan(&out.Size)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(out)
	})
}

// databaseStatsHandler returns summary statistics about stored data.
func databaseStatsHandler(db *sql.DB) http.Handler {
	type stats struct {
		TotalRecords int       `json:"totalRecords"`
		Users        int       `json:"users"`
		Downloads    int       `json:"downloads"`
		MediaItems   int       `json:"mediaItems"`
		LastBackup   time.Time `json:"lastBackup,omitempty"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backend := database.GetDatabaseBackend()
		path := database.GetDatabasePath()
		var store database.SubtitleStore
		var err error
		switch backend {
		case "pebble":
			store, err = database.OpenPebble(path)
		case "postgres":
			store, err = database.OpenPostgresStore(path)
		default:
			store, err = database.OpenSQLStore(path)
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer store.Close()
		subs, _ := store.ListSubtitles()
		downloads, _ := store.ListDownloads()
		media, _ := store.ListMediaItems()
		var users int
		row := db.QueryRow(`SELECT COUNT(1) FROM users`)
		_ = row.Scan(&users)
		st := stats{
			TotalRecords: len(subs),
			Users:        users,
			Downloads:    len(downloads),
			MediaItems:   len(media),
		}
		if hist := backups.List(); len(hist) > 0 {
			st.LastBackup = hist[len(hist)-1].CreatedAt
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(st)
	})
}

// databaseBackupHandler streams a tar.gz of the database files to the client.
func databaseBackupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		backend := database.GetDatabaseBackend()
		path := database.GetDatabasePath()
		name := "db-backup-" + time.Now().Format("20060102-150405") + ".tar.gz"
		w.Header().Set("Content-Type", "application/gzip")
		w.Header().Set("Content-Disposition", "attachment; filename="+name)
		gw := gzip.NewWriter(w)
		defer gw.Close()
		tw := tar.NewWriter(gw)
		defer tw.Close()
		add := func(p, name string) error {
			fi, err := os.Stat(p)
			if err != nil {
				return err
			}
			hdr, err := tar.FileInfoHeader(fi, "")
			if err != nil {
				return err
			}
			hdr.Name = name
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			f, err := os.Open(p)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(tw, f)
			return err
		}
		if backend == "sqlite" {
			if err := add(path, filepath.Base(path)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		if err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				rel, _ := filepath.Rel(filepath.Dir(path), p)
				if err := add(p, rel); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

// databaseOptimizeHandler performs simple optimization operations.
func databaseOptimizeHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if database.GetDatabaseBackend() == "sqlite" {
			_, _ = db.Exec("VACUUM")
		}
		w.WriteHeader(http.StatusOK)
	})
}
