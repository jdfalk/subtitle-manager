package database

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/google/uuid"
)

func createTestDB(t *testing.T) string {
	// Check if PostgreSQL is available
	if _, err := exec.LookPath("createdb"); err != nil {
		t.Skip("PostgreSQL not available: createdb command not found")
	}

	// Check if we can connect as postgres user
	checkCmd := exec.Command("sudo", "-u", "postgres", "psql", "-c", "SELECT 1;")
	if err := checkCmd.Run(); err != nil {
		t.Skip("PostgreSQL not available or cannot connect as postgres user")
	}

	name := "test" + uuid.New().String()[:8]
	cmd := exec.Command("sudo", "-u", "postgres", "createdb", name)
	if err := cmd.Run(); err != nil {
		t.Skipf("create db: %v (PostgreSQL may not be running)", err)
	}
	t.Cleanup(func() {
		exec.Command("sudo", "-u", "postgres", "dropdb", name).Run()
	})
	return name
}

func openTestStore(t *testing.T) *PostgresStore {
	dbName := createTestDB(t)
	dsn := fmt.Sprintf("user=postgres dbname=%s sslmode=disable", dbName)
	s, err := OpenPostgresStore(dsn)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { s.Close() })
	return s
}

func TestPostgresInsertAndList(t *testing.T) {
	db := openTestStore(t)
	rec := &SubtitleRecord{File: "f.srt", VideoFile: "v.mkv", Language: "es", Service: "test"}
	if err := db.InsertSubtitle(rec); err != nil {
		t.Fatalf("insert: %v", err)
	}

	recs, err := db.ListSubtitles()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 1 || recs[0].File != "f.srt" {
		t.Fatalf("unexpected records %+v", recs)
	}
}

func TestPostgresDeleteSubtitle(t *testing.T) {
	db := openTestStore(t)
	rec := &SubtitleRecord{File: "d.srt", VideoFile: "v.mkv", Language: "es", Service: "test"}
	if err := db.InsertSubtitle(rec); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := db.DeleteSubtitle("d.srt"); err != nil {
		t.Fatalf("delete: %v", err)
	}
	recs, err := db.ListSubtitles()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0, got %d", len(recs))
	}
}

func TestPostgresDownloads(t *testing.T) {
	db := openTestStore(t)
	rec := &DownloadRecord{File: "x.srt", VideoFile: "v.mkv", Provider: "p", Language: "en"}
	if err := db.InsertDownload(rec); err != nil {
		t.Fatalf("insert: %v", err)
	}
	recs, err := db.ListDownloads()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 1 || recs[0].Provider != "p" {
		t.Fatalf("unexpected downloads %+v", recs)
	}
	if err := db.DeleteDownload("x.srt"); err != nil {
		t.Fatalf("delete: %v", err)
	}
	recs, err = db.ListDownloads()
	if err != nil {
		t.Fatalf("list2: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0, got %d", len(recs))
	}
}
