package maintenance

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/metadata"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// TestCleanupDatabase ensures expired sessions are removed and VACUUM executes.
func TestCleanupDatabase(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	if err := auth.CreateUser(db, "user", "pass", "u@example.com", "basic"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	id, err := auth.AuthenticateUser(db, "user", "pass")
	if err != nil {
		t.Fatalf("auth: %v", err)
	}
	if _, err := auth.GenerateSession(db, id, -1*time.Minute); err != nil {
		t.Fatalf("expired session: %v", err)
	}
	if _, err := auth.GenerateSession(db, id, time.Hour); err != nil {
		t.Fatalf("valid session: %v", err)
	}

	if err := CleanupDatabase(context.Background(), db); err != nil {
		t.Fatalf("cleanup failed: %v", err)
	}

	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM sessions`).Scan(&count); err != nil {
		t.Fatalf("count sessions: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 session after cleanup, got %d", count)
	}
}

// TestDiskScan verifies the size calculation for a directory tree.
func TestDiskScan(t *testing.T) {
	dir := t.TempDir()
	file1 := filepath.Join(dir, "a.txt")
	if err := os.WriteFile(file1, []byte("hello"), 0644); err != nil {
		t.Fatalf("write file1: %v", err)
	}
	sub := filepath.Join(dir, "sub")
	if err := os.Mkdir(sub, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	file2 := filepath.Join(sub, "b.txt")
	if err := os.WriteFile(file2, []byte("world"), 0644); err != nil {
		t.Fatalf("write file2: %v", err)
	}

	size, err := DiskScan(context.Background(), dir)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if size != int64(len("hello")+len("world")) {
		t.Fatalf("unexpected size %d", size)
	}
}

// TestRefreshMetadata_RespectsFieldLocks ensures locked fields are not updated.
func TestRefreshMetadata_RespectsFieldLocks(t *testing.T) {
	store, err := database.OpenSQLStore(":memory:")
	if err != nil {
		t.Skip("SQLite support not available")
	}
	defer store.Close()

	item := &database.MediaItem{Path: "video.mkv", Title: "Old", FieldLocks: "title"}
	if err := store.InsertMediaItem(item); err != nil {
		t.Fatalf("insert: %v", err)
	}

	metadata.FetchMovieMetadataFunc = func(ctx context.Context, title string, year int, tmdb, omdb string) (*metadata.MediaInfo, error) {
		return &metadata.MediaInfo{Title: "New"}, nil
	}
	defer func() { metadata.FetchMovieMetadataFunc = metadata.FetchMovieMetadata }()

	if err := RefreshMetadata(context.Background(), store, "k", "o"); err != nil {
		t.Fatalf("refresh: %v", err)
	}

	items, err := store.ListMediaItems()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if items[0].Title != "Old" {
		t.Fatalf("title changed: %s", items[0].Title)
	}
}
