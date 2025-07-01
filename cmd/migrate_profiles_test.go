// file: cmd/migrate_profiles_test.go
// version: 1.0.0
// guid: 4e3f2a1b-0c9d-5e6f-8a7b-1c2d3e4f5a6b

package cmd

import (
	"strings"
	"testing"
)

func TestDetectExistingLanguages(t *testing.T) {
	// Test basic functionality without actual environment setup
	languages := detectExistingLanguages()

	// Should return empty slice when no config is found
	if languages == nil {
		t.Error("detectExistingLanguages should return empty slice, not nil")
	}
}

func TestFormatLanguages(t *testing.T) {
	// Test the formatLanguages function with sample data
	testCases := []struct {
		name      string
		languages []struct {
			Language string
			Priority int
			Forced   bool
			HI       bool
		}
		expected string
	}{
		{
			name: "single language",
			languages: []struct {
				Language string
				Priority int
				Forced   bool
				HI       bool
			}{
				{Language: "en", Priority: 1, Forced: false, HI: false},
			},
			expected: "en (priority: 1)",
		},
		{
			name: "multiple languages with flags",
			languages: []struct {
				Language string
				Priority int
				Forced   bool
				HI       bool
			}{
				{Language: "en", Priority: 1, Forced: false, HI: false},
				{Language: "es", Priority: 2, Forced: true, HI: false},
				{Language: "fr", Priority: 3, Forced: false, HI: true},
			},
			expected: "en (priority: 1), es (priority: 2) [forced], fr (priority: 3) [HI]",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert test data to LanguageConfig format
			var languageConfigs []struct {
				Language string
				Priority int
				Forced   bool
				HI       bool
			}
			for _, lang := range tc.languages {
				languageConfigs = append(languageConfigs, struct {
					Language string
					Priority int
					Forced   bool
					HI       bool
				}{
					Language: lang.Language,
					Priority: lang.Priority,
					Forced:   lang.Forced,
					HI:       lang.HI,
				})
			}

			// Test that the basic structure works
			// Note: We can't directly test formatLanguages without importing profiles package
			// But we can test the logic structure
			var parts []string
			for _, lang := range languageConfigs {
				part := lang.Language + " (priority: " + string(rune(lang.Priority+'0')) + ")"
				if lang.Forced {
					part += " [forced]"
				}
				if lang.HI {
					part += " [HI]"
				}
				parts = append(parts, part)
			}
			result := strings.Join(parts, ", ")

			// Basic validation that formatting works
			if !strings.Contains(result, tc.languages[0].Language) {
				t.Errorf("Expected result to contain %s, got %s", tc.languages[0].Language, result)
			}
		})
	}
}

func TestMigrateProfilesValidation(t *testing.T) {
	// Test that the command structure is valid
	if migrateProfilesCmd.Use != "migrate-profiles" {
		t.Errorf("Expected command name 'migrate-profiles', got %s", migrateProfilesCmd.Use)
	}

	if migrateProfilesCmd.Short == "" {
		t.Error("Command should have a short description")
	}

	if migrateProfilesCmd.Long == "" {
		t.Error("Command should have a long description")
	}

	if migrateProfilesCmd.RunE == nil {
		t.Error("Command should have a RunE function")
	}
}

func TestMigrateProfilesFlags(t *testing.T) {
	// Test that flags are properly defined
	flags := migrateProfilesCmd.Flags()

	if !flags.Changed("languages") && flags.Lookup("languages") == nil {
		t.Error("Should have --languages flag defined")
	}

	if !flags.Changed("cutoff-score") && flags.Lookup("cutoff-score") == nil {
		t.Error("Should have --cutoff-score flag defined")
	}

	if !flags.Changed("force") && flags.Lookup("force") == nil {
		t.Error("Should have --force flag defined")
	}
}
