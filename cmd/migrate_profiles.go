// file: cmd/migrate_profiles.go
// version: 1.0.0
// guid: 3d2e1f0a-9b8c-4d5e-7f6a-0b1c2d3e4f5a

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/profiles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateProfilesCmd represents the migrate-profiles command
var migrateProfilesCmd = &cobra.Command{
	Use:   "migrate-profiles",
	Short: "Migrate existing language settings to language profiles",
	Long: `Migrate existing language configuration to the new language profiles system.

This command helps migrate from simple language settings to the new language profiles
system introduced in the Language Profiles feature. It will:

1. Check for existing language configurations in environment variables or config files
2. Create a default language profile based on those settings
3. Preserve existing preferences while enabling the new profile system

Examples:
  subtitle-manager migrate-profiles
  subtitle-manager migrate-profiles --languages en,es,fr
  subtitle-manager migrate-profiles --cutoff-score 80`,
	RunE: migrateProfiles,
}

var (
	languagesList  string
	cutoffScore    int
	forceMigration bool
)

func init() {
	rootCmd.AddCommand(migrateProfilesCmd)

	migrateProfilesCmd.Flags().StringVar(&languagesList, "languages", "", "Comma-separated list of language codes (e.g., en,es,fr)")
	migrateProfilesCmd.Flags().IntVar(&cutoffScore, "cutoff-score", 75, "Default cutoff score for the profile (0-100)")
	migrateProfilesCmd.Flags().BoolVar(&forceMigration, "force", false, "Force migration even if profiles already exist")
}

// migrateProfiles performs the migration from simple language settings to language profiles.
func migrateProfiles(cmd *cobra.Command, args []string) error {
	// Open database store
	store, err := database.OpenSQLStore(database.GetDatabasePath())
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer store.Close()

	// Check if profiles already exist (unless forcing)
	if !forceMigration {
		existingProfiles, err := store.ListLanguageProfiles()
		if err != nil {
			return fmt.Errorf("failed to check existing profiles: %w", err)
		}

		if len(existingProfiles) > 0 {
			fmt.Printf("Found %d existing language profiles. Use --force to migrate anyway.\n", len(existingProfiles))
			fmt.Println("Existing profiles:")
			for _, profile := range existingProfiles {
				fmt.Printf("  - %s (ID: %s, Default: %t)\n", profile.Name, profile.ID, profile.IsDefault)
			}
			return nil
		}
	}

	// Determine languages to migrate
	var languages []string
	if languagesList != "" {
		languages = strings.Split(languagesList, ",")
		for i, lang := range languages {
			languages[i] = strings.TrimSpace(lang)
		}
	} else {
		// Try to detect languages from various configuration sources
		languages = detectExistingLanguages()
		if len(languages) == 0 {
			languages = []string{"en"} // Default to English
		}
	}

	// Validate cutoff score
	if cutoffScore < 0 || cutoffScore > 100 {
		return fmt.Errorf("cutoff score must be between 0 and 100, got %d", cutoffScore)
	}

	// Create language configurations
	var languageConfigs []profiles.LanguageConfig
	for i, lang := range languages {
		if lang == "" {
			continue
		}
		languageConfigs = append(languageConfigs, profiles.LanguageConfig{
			Language: lang,
			Priority: i + 1,
			Forced:   false,
			HI:       false,
		})
	}

	if len(languageConfigs) == 0 {
		return fmt.Errorf("no valid languages specified")
	}

	// Create the migration profile
	migrationProfile := &profiles.LanguageProfile{
		ID:          uuid.NewString(),
		Name:        "Migrated Default Profile",
		Languages:   languageConfigs,
		CutoffScore: cutoffScore,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Validate the profile
	if err := migrationProfile.Validate(); err != nil {
		return fmt.Errorf("migration profile validation failed: %w", err)
	}

	// If forcing migration, clear existing default flags first
	if forceMigration {
		existingProfiles, err := store.ListLanguageProfiles()
		if err != nil {
			return fmt.Errorf("failed to list existing profiles: %w", err)
		}

		for _, profile := range existingProfiles {
			if profile.IsDefault {
				profile.IsDefault = false
				profile.UpdatedAt = time.Now()
				if err := store.UpdateLanguageProfile(&profile); err != nil {
					return fmt.Errorf("failed to update existing default profile: %w", err)
				}
			}
		}
	}

	// Create the profile
	if err := store.CreateLanguageProfile(migrationProfile); err != nil {
		return fmt.Errorf("failed to create migration profile: %w", err)
	}

	fmt.Println("✅ Language profile migration completed successfully!")
	fmt.Printf("Created profile: %s (ID: %s)\n", migrationProfile.Name, migrationProfile.ID)
	fmt.Printf("Languages: %s\n", formatLanguages(migrationProfile.Languages))
	fmt.Printf("Cutoff Score: %d%%\n", migrationProfile.CutoffScore)
	fmt.Printf("Default Profile: %t\n", migrationProfile.IsDefault)

	fmt.Println("\nYou can now manage language profiles through:")
	fmt.Println("  - Web UI: Settings → Language Profiles")
	fmt.Println("  - API: /api/profiles endpoints")

	return nil
}

// detectExistingLanguages attempts to detect existing language settings from various sources.
func detectExistingLanguages() []string {
	var languages []string

	// Check common environment variables
	envVars := []string{
		"SUBTITLE_LANGUAGES",
		"LANG_CODES",
		"LANGUAGES",
		"DEFAULT_LANGUAGES",
	}

	for _, envVar := range envVars {
		if val := viper.GetString(envVar); val != "" {
			detected := strings.Split(val, ",")
			for _, lang := range detected {
				lang = strings.TrimSpace(lang)
				if lang != "" && len(lang) == 2 { // Basic validation for 2-letter codes
					languages = append(languages, lang)
				}
			}
			if len(languages) > 0 {
				fmt.Printf("Detected languages from %s: %s\n", envVar, strings.Join(languages, ", "))
				break
			}
		}
	}

	// Check configuration file settings
	if len(languages) == 0 {
		configKeys := []string{
			"subtitle.languages",
			"languages",
			"default_languages",
		}

		for _, key := range configKeys {
			if viper.IsSet(key) {
				val := viper.GetString(key)
				if val != "" {
					detected := strings.Split(val, ",")
					for _, lang := range detected {
						lang = strings.TrimSpace(lang)
						if lang != "" && len(lang) == 2 {
							languages = append(languages, lang)
						}
					}
					if len(languages) > 0 {
						fmt.Printf("Detected languages from config %s: %s\n", key, strings.Join(languages, ", "))
						break
					}
				}
			}
		}
	}

	// Remove duplicates and validate
	seen := make(map[string]bool)
	var unique []string
	for _, lang := range languages {
		if !seen[lang] && len(lang) == 2 {
			seen[lang] = true
			unique = append(unique, lang)
		}
	}

	return unique
}

// formatLanguages formats language configurations for display.
func formatLanguages(languages []profiles.LanguageConfig) string {
	var parts []string
	for _, lang := range languages {
		part := fmt.Sprintf("%s (priority: %d)", lang.Language, lang.Priority)
		if lang.Forced {
			part += " [forced]"
		}
		if lang.HI {
			part += " [HI]"
		}
		parts = append(parts, part)
	}
	return strings.Join(parts, ", ")
}
