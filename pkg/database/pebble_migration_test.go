// file: pkg/database/pebble_migration_test.go
// version: 1.0.0
// guid: 20250713-pebble-migration-abcdef

package database

import (
	"os"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/profiles"
)

func TestPebble_Migration_Versioning(t *testing.T) {
	path := "test_pebble_migration.db"
	_ = os.RemoveAll(path)
	store, err := OpenPebble(path)
	if err != nil {
		t.Fatalf("failed to open Pebble store: %v", err)
	}
	defer func() { store.Close(); os.RemoveAll(path) }()

	// Simulate legacy profile (missing fields)
	legacyProfile := &LanguageProfile{
		ID:          "legacy-profile",
		Name:        "Legacy",
		Languages:   []profiles.LanguageConfig{{Language: "en", Priority: 1}},
		CutoffScore: 70,
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   time.Now().Add(-24 * time.Hour),
	}
	if err := store.CreateLanguageProfile(legacyProfile); err != nil {
		t.Fatalf("failed to create legacy profile: %v", err)
	}

	// Simulate new profile
	newProfile := &LanguageProfile{
		ID:          "new-profile",
		Name:        "New",
		Languages:   []profiles.LanguageConfig{{Language: "en", Priority: 1}, {Language: "es", Priority: 2}},
		CutoffScore: 90,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := store.CreateLanguageProfile(newProfile); err != nil {
		t.Fatalf("failed to create new profile: %v", err)
	}

	// Assign legacy profile to media
	mediaID := "media-legacy"
	if err := store.AssignProfileToMedia(mediaID, legacyProfile.ID); err != nil {
		t.Fatalf("failed to assign legacy profile: %v", err)
	}

	// Assign new profile to media
	mediaID2 := "media-new"
	if err := store.AssignProfileToMedia(mediaID2, newProfile.ID); err != nil {
		t.Fatalf("failed to assign new profile: %v", err)
	}

	// Simulate migration: update legacy profile to new format
	legacyProfile.IsDefault = false
	legacyProfile.UpdatedAt = time.Now()
	if err := store.UpdateLanguageProfile(legacyProfile); err != nil {
		t.Fatalf("failed to migrate legacy profile: %v", err)
	}

	// Verify assignments
	assignedLegacy, err := store.GetMediaProfile(mediaID)
	if err != nil {
		t.Fatalf("failed to get assigned legacy profile: %v", err)
	}
	if assignedLegacy == nil || assignedLegacy.ID != legacyProfile.ID {
		t.Errorf("expected legacy profile %s, got %+v", legacyProfile.ID, assignedLegacy)
	}

	assignedNew, err := store.GetMediaProfile(mediaID2)
	if err != nil {
		t.Fatalf("failed to get assigned new profile: %v", err)
	}
	if assignedNew == nil || assignedNew.ID != newProfile.ID {
		t.Errorf("expected new profile %s, got %+v", newProfile.ID, assignedNew)
	}

	// Remove legacy profile and verify assignment fallback
	if err := store.DeleteLanguageProfile(legacyProfile.ID); err != nil {
		t.Fatalf("failed to delete legacy profile: %v", err)
	}
	fallback, err := store.GetMediaProfile(mediaID)
	if err != nil {
		t.Fatalf("failed to get fallback profile: %v", err)
	}
	if fallback == nil || fallback.ID != newProfile.ID {
		t.Errorf("expected fallback to new profile %s, got %+v", newProfile.ID, fallback)
	}
}
