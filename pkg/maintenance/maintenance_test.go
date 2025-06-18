package maintenance

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// TestCleanupDatabase ensures expired sessions are removed and VACUUM executes.
func TestCleanupDatabase(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
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
