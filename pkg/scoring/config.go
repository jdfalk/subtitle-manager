// file: pkg/scoring/config.go
// version: 1.0.0
// guid: 11111111-2222-3333-4444-555555555555
package scoring

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// LoadProfileFromConfig creates a scoring profile from viper configuration.
func LoadProfileFromConfig() Profile {
	profile := DefaultProfile()

	// Load scoring weights
	if viper.IsSet("scoring.provider_weight") {
		profile.ProviderWeight = viper.GetFloat64("scoring.provider_weight")
	}
	if viper.IsSet("scoring.release_weight") {
		profile.ReleaseWeight = viper.GetFloat64("scoring.release_weight")
	}
	if viper.IsSet("scoring.format_weight") {
		profile.FormatWeight = viper.GetFloat64("scoring.format_weight")
	}
	if viper.IsSet("scoring.metadata_weight") {
		profile.MetadataWeight = viper.GetFloat64("scoring.metadata_weight")
	}

	// Load format preferences
	if viper.IsSet("scoring.preferred_formats") {
		profile.PreferredFormats = viper.GetStringSlice("scoring.preferred_formats")
	}

	// Load hearing impaired preferences
	if viper.IsSet("scoring.allow_hi") {
		profile.AllowHI = viper.GetBool("scoring.allow_hi")
	}
	if viper.IsSet("scoring.prefer_hi") {
		profile.PreferHI = viper.GetBool("scoring.prefer_hi")
	}

	// Load forced subtitle preferences
	if viper.IsSet("scoring.allow_forced") {
		profile.AllowForced = viper.GetBool("scoring.allow_forced")
	}
	if viper.IsSet("scoring.prefer_forced") {
		profile.PreferForced = viper.GetBool("scoring.prefer_forced")
	}

	// Load thresholds
	if viper.IsSet("scoring.min_score") {
		profile.MinScore = viper.GetInt("scoring.min_score")
	}
	if viper.IsSet("scoring.max_age_days") {
		days := viper.GetInt("scoring.max_age_days")
		profile.MaxAge = time.Duration(days) * 24 * time.Hour
	}

	return profile
}

// SaveProfileToConfig saves a scoring profile to viper configuration.
func SaveProfileToConfig(profile Profile) {
	viper.Set("scoring.provider_weight", profile.ProviderWeight)
	viper.Set("scoring.release_weight", profile.ReleaseWeight)
	viper.Set("scoring.format_weight", profile.FormatWeight)
	viper.Set("scoring.metadata_weight", profile.MetadataWeight)
	
	viper.Set("scoring.preferred_formats", profile.PreferredFormats)
	viper.Set("scoring.allow_hi", profile.AllowHI)
	viper.Set("scoring.prefer_hi", profile.PreferHI)
	viper.Set("scoring.allow_forced", profile.AllowForced)
	viper.Set("scoring.prefer_forced", profile.PreferForced)
	
	viper.Set("scoring.min_score", profile.MinScore)
	viper.Set("scoring.max_age_days", int(profile.MaxAge.Hours()/24))
}

// ValidateProfile ensures a profile has valid settings.
func ValidateProfile(profile Profile) error {
	// Check weights sum to approximately 1.0
	weightSum := profile.ProviderWeight + profile.ReleaseWeight + profile.FormatWeight + profile.MetadataWeight
	if weightSum < 0.9 || weightSum > 1.1 {
		return fmt.Errorf("scoring weights must sum to approximately 1.0, got %f", weightSum)
	}

	// Check individual weights are in valid range
	if profile.ProviderWeight < 0 || profile.ProviderWeight > 1 {
		return fmt.Errorf("provider_weight must be between 0 and 1, got %f", profile.ProviderWeight)
	}
	if profile.ReleaseWeight < 0 || profile.ReleaseWeight > 1 {
		return fmt.Errorf("release_weight must be between 0 and 1, got %f", profile.ReleaseWeight)
	}
	if profile.FormatWeight < 0 || profile.FormatWeight > 1 {
		return fmt.Errorf("format_weight must be between 0 and 1, got %f", profile.FormatWeight)
	}
	if profile.MetadataWeight < 0 || profile.MetadataWeight > 1 {
		return fmt.Errorf("metadata_weight must be between 0 and 1, got %f", profile.MetadataWeight)
	}

	// Check score threshold
	if profile.MinScore < 0 || profile.MinScore > 100 {
		return fmt.Errorf("min_score must be between 0 and 100, got %d", profile.MinScore)
	}

	return nil
}