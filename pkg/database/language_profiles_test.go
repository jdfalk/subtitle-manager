// file: pkg/database/language_profiles_test.go
// version: 1.0.0
// guid: e5f6g7h8-i9j0-1234-efab-6789012345bc

package database

import (
	"testing"
)

// TestDefaultLanguageProfile verifies the default profile has sensible defaults.
func TestDefaultLanguageProfile(t *testing.T) {
	profile := DefaultLanguageProfile()

	if profile.Name != "Default" {
		t.Errorf("expected name 'Default', got %s", profile.Name)
	}

	if !profile.IsDefault {
		t.Error("expected IsDefault to be true")
	}

	if len(profile.Languages) == 0 {
		t.Error("expected at least one language")
	}

	if len(profile.Providers) == 0 {
		t.Error("expected at least one provider")
	}

	if profile.MinScore <= 0 || profile.MinScore > 1 {
		t.Errorf("expected MinScore between 0 and 1, got %f", profile.MinScore)
	}

	if len(profile.ScoreWeights) == 0 {
		t.Error("expected score weights to be defined")
	}

	// Verify score weights sum to 1.0 (approximately)
	var total float64
	for _, weight := range profile.ScoreWeights {
		total += weight
	}
	if total < 0.99 || total > 1.01 {
		t.Errorf("expected score weights to sum to ~1.0, got %f", total)
	}
}

// TestLanguageProfileSQLiteIntegration tests language profile operations with SQLite.
func TestLanguageProfileSQLiteIntegration(t *testing.T) {
	if !hasSQLite() {
		t.Skip("SQLite not available")
	}

	store, err := OpenSQLStore(":memory:")
	if err != nil {
		t.Fatalf("failed to open SQLite store: %v", err)
	}
	defer store.Close()

	// Test inserting a language profile
	profile := DefaultLanguageProfile()
	profile.Name = "Test Profile"
	profile.Description = "Test language profile"

	err = store.InsertLanguageProfile(profile)
	if err != nil {
		t.Fatalf("failed to insert language profile: %v", err)
	}

	// Test listing profiles
	profiles, err := store.ListLanguageProfiles()
	if err != nil {
		t.Fatalf("failed to list language profiles: %v", err)
	}

	if len(profiles) != 1 {
		t.Fatalf("expected 1 profile, got %d", len(profiles))
	}

	retrieved := profiles[0]
	if retrieved.Name != "Test Profile" {
		t.Errorf("expected name 'Test Profile', got %s", retrieved.Name)
	}

	if retrieved.Description != "Test language profile" {
		t.Errorf("expected description 'Test language profile', got %s", retrieved.Description)
	}

	// Test getting by name
	profileByName, err := store.GetLanguageProfileByName("Test Profile")
	if err != nil {
		t.Fatalf("failed to get profile by name: %v", err)
	}

	if profileByName.ID != retrieved.ID {
		t.Error("profile by name should match listed profile")
	}

	// Test getting default profile
	defaultProfile, err := store.GetDefaultLanguageProfile()
	if err != nil {
		t.Fatalf("failed to get default profile: %v", err)
	}

	if defaultProfile.ID != retrieved.ID {
		t.Error("default profile should match the inserted profile")
	}

	// Test profile assignment
	assignment := &LanguageProfileAssignment{
		ProfileID:      retrieved.ID,
		MediaPath:      "/test/movie.mkv",
		MediaType:      "movie",
		AssignmentType: "manual",
		Priority:       1,
	}

	err = store.AssignProfileToMedia(assignment)
	if err != nil {
		t.Fatalf("failed to assign profile to media: %v", err)
	}

	// Test getting assignment
	retrievedAssignment, err := store.GetProfileAssignmentForMedia("/test/movie.mkv")
	if err != nil {
		t.Fatalf("failed to get profile assignment: %v", err)
	}

	if retrievedAssignment.ProfileID != retrieved.ID {
		t.Error("assignment should match the assigned profile")
	}

	if retrievedAssignment.MediaPath != "/test/movie.mkv" {
		t.Errorf("expected media path '/test/movie.mkv', got %s", retrievedAssignment.MediaPath)
	}
}

// TestLanguageProfileRules tests profile rule functionality.
func TestLanguageProfileRules(t *testing.T) {
	if !hasSQLite() {
		t.Skip("SQLite not available")
	}

	store, err := OpenSQLStore(":memory:")
	if err != nil {
		t.Fatalf("failed to open SQLite store: %v", err)
	}
	defer store.Close()

	// Create a profile first
	profile := DefaultLanguageProfile()
	profile.Name = "Test Profile"
	err = store.InsertLanguageProfile(profile)
	if err != nil {
		t.Fatalf("failed to insert language profile: %v", err)
	}

	profiles, err := store.ListLanguageProfiles()
	if err != nil {
		t.Fatalf("failed to list profiles: %v", err)
	}
	profileID := profiles[0].ID

	// Create a rule
	rule := &LanguageProfileRule{
		ProfileID:   profileID,
		Name:        "Anime Rule",
		Description: "Apply to anime content",
		Conditions: map[string]interface{}{
			"path_contains": []string{"anime", "Animation"},
			"genre":         "Animation",
		},
		Priority: 10,
		Enabled:  true,
	}

	err = store.InsertProfileRule(rule)
	if err != nil {
		t.Fatalf("failed to insert profile rule: %v", err)
	}

	// Test listing rules for profile
	rules, err := store.ListProfileRules(profileID)
	if err != nil {
		t.Fatalf("failed to list profile rules: %v", err)
	}

	if len(rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(rules))
	}

	retrievedRule := rules[0]
	if retrievedRule.Name != "Anime Rule" {
		t.Errorf("expected name 'Anime Rule', got %s", retrievedRule.Name)
	}

	if retrievedRule.Priority != 10 {
		t.Errorf("expected priority 10, got %d", retrievedRule.Priority)
	}

	if !retrievedRule.Enabled {
		t.Error("expected rule to be enabled")
	}

	// Test getting enabled rules
	enabledRules, err := store.ListEnabledProfileRules()
	if err != nil {
		t.Fatalf("failed to list enabled rules: %v", err)
	}

	if len(enabledRules) != 1 {
		t.Fatalf("expected 1 enabled rule, got %d", len(enabledRules))
	}

	// Test updating rule
	retrievedRule.Enabled = false
	err = store.UpdateProfileRule(&retrievedRule)
	if err != nil {
		t.Fatalf("failed to update profile rule: %v", err)
	}

	// Verify enabled rules list is now empty
	enabledRules, err = store.ListEnabledProfileRules()
	if err != nil {
		t.Fatalf("failed to list enabled rules: %v", err)
	}

	if len(enabledRules) != 0 {
		t.Fatalf("expected 0 enabled rules, got %d", len(enabledRules))
	}
}

// hasSQLite checks if SQLite support is available
func hasSQLite() bool {
	store, err := OpenSQLStore(":memory:")
	if err != nil {
		return false
	}
	store.Close()
	return true
}
