package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// TestDeleteCmd ensures that subtitle files and database records are removed.
func TestDeleteCmd(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "sub.srt")
	if err := os.WriteFile(file, []byte("x"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	dbPath := filepath.Join(dir, "test.db")
	db, err := database.Open(dbPath)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := database.InsertSubtitle(db, file, file, "en", "test", "", false); err != nil {
		t.Fatalf("insert: %v", err)
	}
	db.Close()

	orig := viper.GetString("db_path")
	viper.Set("db_path", dbPath)
	defer viper.Set("db_path", orig)

	if err := deleteCmd.RunE(deleteCmd, []string{file}); err != nil {
		t.Fatalf("run: %v", err)
	}

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		t.Fatalf("file not deleted: %v", err)
	}

	db, err = database.Open(dbPath)
	if err != nil {
		t.Fatalf("reopen db: %v", err)
	}
	defer db.Close()
	recs, err := database.ListSubtitles(db)
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0 records, got %d", len(recs))
	}
}
