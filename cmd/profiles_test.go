// file: cmd/profiles_test.go
// version: 1.0.0
// guid: 2f19b0dd-3cdf-4c0f-bf9f-c8c5848240b2

package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/profiles"
	"github.com/spf13/viper"
)

func configureProfileTestPath(t *testing.T) string {
	t.Helper()

	tempDir := t.TempDir()
	setViperString(t, "db_backend", "pebble")
	setViperString(t, "db_path", tempDir)
	setViperString(t, "media_directory", tempDir)
	return tempDir
}

func setViperString(t *testing.T, key, value string) {
	t.Helper()

	previous := viper.GetString(key)
	viper.Set(key, value)
	t.Cleanup(func() {
		viper.Set(key, previous)
	})
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}

	originalStdout := os.Stdout
	os.Stdout = writer

	t.Cleanup(func() {
		os.Stdout = originalStdout
	})

	fn()
	if err := writer.Close(); err != nil {
		t.Fatalf("failed to close stdout writer: %v", err)
	}

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}
	return string(output)
}

func TestCreateProfileCmd_CreatesProfileWithFlags(t *testing.T) {
	// Arrange: configure store and command flags.
	storePath := configureProfileTestPath(t)
	if err := createProfileCmd.Flags().Set("languages", "en,es:forced,fr:hi"); err != nil {
		t.Fatalf("failed to set languages flag: %v", err)
	}
	if err := createProfileCmd.Flags().Set("cutoff", "65"); err != nil {
		t.Fatalf("failed to set cutoff flag: %v", err)
	}
	if err := createProfileCmd.Flags().Set("default", "true"); err != nil {
		t.Fatalf("failed to set default flag: %v", err)
	}

	// Act: execute the create command.
	output := captureStdout(t, func() {
		if err := createProfileCmd.RunE(createProfileCmd, []string{"Cinema"}); err != nil {
			t.Fatalf("createProfileCmd.RunE returned error: %v", err)
		}
	})

	// Assert: profile exists with parsed language flags and default status.
	profileID := extractProfileID(t, output)
	store := openProfileStore(t, storePath)
	defer closeStore(t, store)
	profile, err := store.GetLanguageProfile(profileID)
	if err != nil {
		t.Fatalf("failed to load profile: %v", err)
	}
	if profile == nil {
		t.Fatalf("expected profile to be created")
	}
	if profile.Name != "Cinema" {
		t.Fatalf("expected profile name Cinema, got %s", profile.Name)
	}
	if profile.CutoffScore != 65 {
		t.Fatalf("expected cutoff score 65, got %d", profile.CutoffScore)
	}
	if !profile.IsDefault {
		t.Fatalf("expected profile to be default")
	}
	if len(profile.Languages) != 3 {
		t.Fatalf("expected 3 languages, got %d", len(profile.Languages))
	}
	if !profile.Languages[1].Forced {
		t.Fatalf("expected second language to be forced")
	}
	if !profile.Languages[2].HI {
		t.Fatalf("expected third language to be hearing impaired")
	}
}

func TestListProfilesCmd_PrintsTable(t *testing.T) {
	// Arrange: seed a profile in the store.
	storePath := configureProfileTestPath(t)
	store := openProfileStore(t, storePath)
	profile := &profiles.LanguageProfile{
		ID:   "abcd1234efgh",
		Name: "Series",
		Languages: []profiles.LanguageConfig{
			{Language: "en", Priority: 1, Forced: true, HI: true},
		},
		CutoffScore: 85,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := store.CreateLanguageProfile(profile); err != nil {
		t.Fatalf("failed to create profile: %v", err)
	}
	closeStore(t, store)

	// Act: execute the list command.
	output := captureStdout(t, func() {
		if err := listProfilesCmd.RunE(listProfilesCmd, nil); err != nil {
			t.Fatalf("listProfilesCmd.RunE returned error: %v", err)
		}
	})

	// Assert: output contains the expected table details.
	if !strings.Contains(output, "ID") || !strings.Contains(output, "Name") ||
		!strings.Contains(output, "Default") || !strings.Contains(output, "Languages") {
		t.Fatalf("expected table header in output, got %s", output)
	}
	if !strings.Contains(output, "abcd1234") {
		t.Fatalf("expected profile ID prefix in output, got %s", output)
	}
	if !strings.Contains(output, "Series") {
		t.Fatalf("expected profile name in output, got %s", output)
	}
	if !strings.Contains(output, "✓") {
		t.Fatalf("expected default marker in output, got %s", output)
	}
	if !strings.Contains(output, "85%") {
		t.Fatalf("expected cutoff score in output, got %s", output)
	}
	if !strings.Contains(output, "en(1)[F][HI]") {
		t.Fatalf("expected language flags in output, got %s", output)
	}
}

func TestAssignAndShowProfileCmd_AssignsCustomProfile(t *testing.T) {
	// Arrange: create a custom profile and target media path.
	tempDir := configureProfileTestPath(t)
	store := openProfileStore(t, tempDir)
	profile := &profiles.LanguageProfile{
		ID:   "custom1234",
		Name: "Custom",
		Languages: []profiles.LanguageConfig{
			{Language: "en", Priority: 1},
		},
		CutoffScore: 70,
		IsDefault:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := store.CreateLanguageProfile(profile); err != nil {
		t.Fatalf("failed to create profile: %v", err)
	}
	closeStore(t, store)
	mediaPath := filepath.Join(tempDir, "movie.mkv")

	// Act: assign and show the profile.
	assignOutput := captureStdout(t, func() {
		if err := assignProfileCmd.RunE(assignProfileCmd, []string{mediaPath, profile.ID}); err != nil {
			t.Fatalf("assignProfileCmd.RunE returned error: %v", err)
		}
	})
	showOutput := captureStdout(t, func() {
		if err := showMediaProfileCmd.RunE(showMediaProfileCmd, []string{mediaPath}); err != nil {
			t.Fatalf("showMediaProfileCmd.RunE returned error: %v", err)
		}
	})

	// Assert: assignment output and profile details are correct.
	if !strings.Contains(assignOutput, "Assigned profile") {
		t.Fatalf("expected assignment output, got %s", assignOutput)
	}
	if !strings.Contains(showOutput, fmt.Sprintf("Profile: %s (ID: %s)", profile.Name, profile.ID)) {
		t.Fatalf("expected profile info in output, got %s", showOutput)
	}
	if !strings.Contains(showOutput, "Type: Custom Profile") {
		t.Fatalf("expected custom profile type, got %s", showOutput)
	}
	if !strings.Contains(showOutput, "1. en") {
		t.Fatalf("expected language list in output, got %s", showOutput)
	}
}

func TestRemoveProfileCmd_RemovesAssignment(t *testing.T) {
	// Arrange: assign a profile to a media path.
	tempDir := configureProfileTestPath(t)
	store := openProfileStore(t, tempDir)
	profile := &profiles.LanguageProfile{
		ID:   "remove1234",
		Name: "Remove",
		Languages: []profiles.LanguageConfig{
			{Language: "en", Priority: 1},
		},
		CutoffScore: 60,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := store.CreateLanguageProfile(profile); err != nil {
		t.Fatalf("failed to create profile: %v", err)
	}
	if err := store.SetDefaultLanguageProfile(profile.ID); err != nil {
		t.Fatalf("failed to set default profile: %v", err)
	}
	mediaPath := filepath.Join(tempDir, "movie.mkv")
	if err := store.AssignProfileToMedia(mediaPath, profile.ID); err != nil {
		t.Fatalf("failed to assign profile: %v", err)
	}
	closeStore(t, store)

	// Act: remove the assignment.
	output := captureStdout(t, func() {
		if err := removeProfileCmd.RunE(removeProfileCmd, []string{mediaPath}); err != nil {
			t.Fatalf("removeProfileCmd.RunE returned error: %v", err)
		}
	})

	// Assert: assignment is removed and default profile is returned.
	if !strings.Contains(output, "Removed profile assignment") {
		t.Fatalf("expected removal output, got %s", output)
	}
	store = openProfileStore(t, tempDir)
	defer closeStore(t, store)
	assignedProfile, err := store.GetMediaProfile(mediaPath)
	if err != nil {
		t.Fatalf("failed to load media profile: %v", err)
	}
	if assignedProfile == nil || !assignedProfile.IsDefault {
		t.Fatalf("expected default profile after removal")
	}
	if assignedProfile.ID != profile.ID {
		t.Fatalf("expected default profile ID %s, got %s", profile.ID, assignedProfile.ID)
	}
}

func TestDeleteProfileCmd_DeletesCustomAndBlocksDefault(t *testing.T) {
	// Arrange: create a default and custom profile.
	storePath := configureProfileTestPath(t)
	store := openProfileStore(t, storePath)
	defaultProfile := &profiles.LanguageProfile{
		ID:   "default1234",
		Name: "DefaultProfile",
		Languages: []profiles.LanguageConfig{
			{Language: "en", Priority: 1},
		},
		CutoffScore: 80,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	customProfile := &profiles.LanguageProfile{
		ID:   "custom5678",
		Name: "CustomProfile",
		Languages: []profiles.LanguageConfig{
			{Language: "es", Priority: 1},
		},
		CutoffScore: 75,
		IsDefault:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := store.CreateLanguageProfile(defaultProfile); err != nil {
		t.Fatalf("failed to create default profile: %v", err)
	}
	if err := store.CreateLanguageProfile(customProfile); err != nil {
		t.Fatalf("failed to create custom profile: %v", err)
	}
	if err := store.SetDefaultLanguageProfile(defaultProfile.ID); err != nil {
		t.Fatalf("failed to set default profile: %v", err)
	}
	closeStore(t, store)

	// Act: delete the custom profile.
	deleteOutput := captureStdout(t, func() {
		if err := deleteProfileCmd.RunE(deleteProfileCmd, []string{customProfile.ID}); err != nil {
			t.Fatalf("deleteProfileCmd.RunE returned error: %v", err)
		}
	})

	// Assert: custom profile deleted and default profile deletion is blocked.
	if !strings.Contains(deleteOutput, "Deleted language profile") {
		t.Fatalf("expected delete output, got %s", deleteOutput)
	}
	store = openProfileStore(t, storePath)
	deletedProfile, err := store.GetLanguageProfile(customProfile.ID)
	if err != nil {
		t.Fatalf("failed to read deleted profile: %v", err)
	}
	if deletedProfile != nil {
		t.Fatalf("expected custom profile to be deleted")
	}
	closeStore(t, store)
	if err := deleteProfileCmd.RunE(deleteProfileCmd, []string{defaultProfile.ID}); err == nil {
		t.Fatalf("expected error when deleting default profile")
	}
}

func extractProfileID(t *testing.T, output string) string {
	t.Helper()

	start := strings.Index(output, "ID: ")
	if start == -1 {
		t.Fatalf("expected output to contain profile ID, got %s", output)
	}
	idPart := strings.TrimSpace(output[start+len("ID: "):])
	idPart = strings.TrimSuffix(idPart, ")")
	idPart = strings.TrimSpace(idPart)
	return idPart
}

func openProfileStore(t *testing.T, storePath string) database.SubtitleStore {
	t.Helper()

	store, err := database.OpenStore(storePath, "pebble")
	if err != nil {
		t.Fatalf("failed to open test store: %v", err)
	}
	return store
}

func closeStore(t *testing.T, store database.SubtitleStore) {
	t.Helper()

	if err := store.Close(); err != nil {
		t.Fatalf("failed to close test store: %v", err)
	}
}
