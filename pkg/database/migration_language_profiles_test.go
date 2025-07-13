// file: pkg/database/migration_language_profiles_test.go
// version: 1.0.0
// guid: 20250713-abcdef01-2345-6789-abcdefabcdef

package database

import (
	"os"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/profiles"
)

func TestMigration_LanguageProfileAssignments_SQLite(t *testing.T) {
	path := "test_migration_sqlite.db"
	_ = os.Remove(path)
	store, err := OpenSQLStore(path)
	if err != nil {
		t.Fatalf("failed to open SQLite store: %v", err)
	}
	defer func() { store.Close(); os.Remove(path) }()

	profile := profiles.LanguageProfile{
		ID:          "test-profile",
		Name:        "Test Profile",
		Languages:   []profiles.LanguageConfig{{Language: "en", Priority: 1}},
		CutoffScore: 80,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := store.CreateLanguageProfile(&profile); err != nil {
		t.Fatalf("failed to create language profile: %v", err)
	}

	// Assign profile to media
	mediaID := "media-123"
	if err := store.AssignProfileToMedia(mediaID, profile.ID); err != nil {
		t.Fatalf("failed to assign profile to media: %v", err)
	}

	// Verify assignment
	assigned, err := store.GetMediaProfile(mediaID)
	if err != nil {
		t.Fatalf("failed to get assigned profile: %v", err)
	}
	if assigned == nil || assigned.ID != profile.ID {
		t.Errorf("expected assigned profile %s, got %+v", profile.ID, assigned)
	}
}

func TestMigration_LanguageProfileAssignments_Pebble(t *testing.T) {
	path := "test_migration_pebble.db"
	_ = os.RemoveAll(path)
	store, err := OpenPebble(path)
	if err != nil {
		t.Fatalf("failed to open Pebble store: %v", err)
	}
	defer func() { store.Close(); os.RemoveAll(path) }()

	profile := &LanguageProfile{
		ID:          "test-profile",
		Name:        "Test Profile",
		Languages:   []profiles.LanguageConfig{{Language: "en", Priority: 1}},
		CutoffScore: 80,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := store.CreateLanguageProfile(profile); err != nil {
		t.Fatalf("failed to create language profile: %v", err)
	}

	mediaID := "media-123"
	if err := store.AssignProfileToMedia(mediaID, profile.ID); err != nil {
		t.Fatalf("failed to assign profile to media: %v", err)
	}

	assigned, err := store.GetMediaProfile(mediaID)
	if err != nil {
		t.Fatalf("failed to get assigned profile: %v", err)
	}
	if assigned == nil || assigned.ID != profile.ID {
		t.Errorf("expected assigned profile %s, got %+v", profile.ID, assigned)
	}
}

// Add similar test for Postgres if DSN/test infra is available
