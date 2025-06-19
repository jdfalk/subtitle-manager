package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/auth"
	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// TestLoginCmd verifies that the login command authenticates a user and stores a session token.
func TestLoginCmd(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")
	viper.Set("db_path", dbPath)
	viper.Set("db_backend", "sqlite")
	defer viper.Reset()

	db, err := database.Open(dbPath)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := auth.CreateUser(db, "u", "p", "e@example.com", "user"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	db.Close()

	buf := &bytes.Buffer{}
	loginCmd.SetOut(buf)
	if err := loginCmd.RunE(loginCmd, []string{"u", "p"}); err != nil {
		t.Fatalf("run: %v", err)
	}
	if !strings.Contains(buf.String(), "Session Token:") {
		t.Fatalf("token output missing: %s", buf.String())
	}

	db, err = database.Open(dbPath)
	if err != nil {
		t.Fatalf("reopen db: %v", err)
	}
	defer db.Close()
	var count int
	row := db.QueryRow(`SELECT COUNT(*) FROM sessions`)
	if err := row.Scan(&count); err != nil {
		t.Fatalf("scan: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 session, got %d", count)
	}
}

// TestLoginTokenCmd verifies authentication using a one-time token.
func TestLoginTokenCmd(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")
	viper.Set("db_path", dbPath)
	viper.Set("db_backend", "sqlite")
	defer viper.Reset()

	db, err := database.Open(dbPath)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := auth.CreateUser(db, "u", "p", "e@example.com", "user"); err != nil {
		t.Fatalf("create user: %v", err)
	}
	var id int64
	err = db.QueryRow(`SELECT id FROM users WHERE username = ?`, "u").Scan(&id)
	if err != nil {
		t.Fatalf("scan user ID: %v", err)
	}
	token, err := auth.GenerateOneTimeToken(db, id, time.Hour)
	if err != nil {
		t.Fatalf("token: %v", err)
	}
	db.Close()

	buf := &bytes.Buffer{}
	loginTokenCmd.SetOut(buf)
	if err := loginTokenCmd.RunE(loginTokenCmd, []string{token}); err != nil {
		t.Fatalf("run: %v", err)
	}
	if !strings.Contains(buf.String(), "Session Token:") {
		t.Fatalf("token output missing: %s", buf.String())
	}

	db, err = database.Open(dbPath)
	if err != nil {
		t.Fatalf("reopen db: %v", err)
	}
	defer db.Close()
	var count int
	row := db.QueryRow(`SELECT COUNT(*) FROM sessions`)
	if err := row.Scan(&count); err != nil {
		t.Fatalf("scan: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 session, got %d", count)
	}
}
