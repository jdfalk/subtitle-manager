// file: pkg/scoring/config_test.go
// version: 1.0.0
// guid: aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee
package scoring

import (
	"testing"
	"time"

	"github.com/spf13/viper"
)

func TestLoadProfileFromConfig(t *testing.T) {
	// Clear any existing config
	viper.Reset()

	// Set some config values
	viper.Set("scoring.provider_weight", 0.3)
	viper.Set("scoring.release_weight", 0.4)
	viper.Set("scoring.format_weight", 0.1)
	viper.Set("scoring.metadata_weight", 0.2)
	viper.Set("scoring.preferred_formats", []string{"srt", "vtt"})
	viper.Set("scoring.allow_hi", false)
	viper.Set("scoring.prefer_hi", true)
	viper.Set("scoring.min_score", 75)
	viper.Set("scoring.max_age_days", 180)

	profile := LoadProfileFromConfig()

	// Check weights
	if profile.ProviderWeight != 0.3 {
		t.Errorf("Expected ProviderWeight 0.3, got %f", profile.ProviderWeight)
	}
	if profile.ReleaseWeight != 0.4 {
		t.Errorf("Expected ReleaseWeight 0.4, got %f", profile.ReleaseWeight)
	}
	if profile.FormatWeight != 0.1 {
		t.Errorf("Expected FormatWeight 0.1, got %f", profile.FormatWeight)
	}
	if profile.MetadataWeight != 0.2 {
		t.Errorf("Expected MetadataWeight 0.2, got %f", profile.MetadataWeight)
	}

	// Check preferences
	if len(profile.PreferredFormats) != 2 || profile.PreferredFormats[0] != "srt" || profile.PreferredFormats[1] != "vtt" {
		t.Errorf("Expected PreferredFormats [srt, vtt], got %v", profile.PreferredFormats)
	}
	if profile.AllowHI {
		t.Error("Expected AllowHI false, got true")
	}
	if !profile.PreferHI {
		t.Error("Expected PreferHI true, got false")
	}

	// Check thresholds
	if profile.MinScore != 75 {
		t.Errorf("Expected MinScore 75, got %d", profile.MinScore)
	}
	expectedAge := 180 * 24 * time.Hour
	if profile.MaxAge != expectedAge {
		t.Errorf("Expected MaxAge %v, got %v", expectedAge, profile.MaxAge)
	}
}

func TestLoadProfileFromConfig_Defaults(t *testing.T) {
	// Clear any existing config
	viper.Reset()

	profile := LoadProfileFromConfig()
	defaultProfile := DefaultProfile()

	// Should match default profile when no config is set
	if profile.ProviderWeight != defaultProfile.ProviderWeight {
		t.Errorf("Expected default ProviderWeight %f, got %f", defaultProfile.ProviderWeight, profile.ProviderWeight)
	}
	if profile.MinScore != defaultProfile.MinScore {
		t.Errorf("Expected default MinScore %d, got %d", defaultProfile.MinScore, profile.MinScore)
	}
}

func TestSaveProfileToConfig(t *testing.T) {
	// Clear any existing config
	viper.Reset()

	profile := Profile{
		ProviderWeight:   0.2,
		ReleaseWeight:    0.5,
		FormatWeight:     0.1,
		MetadataWeight:   0.2,
		PreferredFormats: []string{"ass", "srt"},
		AllowHI:          false,
		PreferHI:         true,
		MinScore:         80,
		MaxAge:           90 * 24 * time.Hour,
	}

	SaveProfileToConfig(profile)

	// Check that values were saved to viper
	if viper.GetFloat64("scoring.provider_weight") != 0.2 {
		t.Errorf("Expected saved ProviderWeight 0.2, got %f", viper.GetFloat64("scoring.provider_weight"))
	}
	if viper.GetInt("scoring.min_score") != 80 {
		t.Errorf("Expected saved MinScore 80, got %d", viper.GetInt("scoring.min_score"))
	}
	if viper.GetInt("scoring.max_age_days") != 90 {
		t.Errorf("Expected saved MaxAgeDays 90, got %d", viper.GetInt("scoring.max_age_days"))
	}

	formats := viper.GetStringSlice("scoring.preferred_formats")
	if len(formats) != 2 || formats[0] != "ass" || formats[1] != "srt" {
		t.Errorf("Expected saved PreferredFormats [ass, srt], got %v", formats)
	}
}

func TestValidateProfile(t *testing.T) {
	tests := []struct {
		name    string
		profile Profile
		wantErr bool
	}{
		{
			name:    "valid default profile",
			profile: DefaultProfile(),
			wantErr: false,
		},
		{
			name: "valid custom profile",
			profile: Profile{
				ProviderWeight:  0.2,
				ReleaseWeight:   0.3,
				FormatWeight:    0.2,
				MetadataWeight:  0.3,
				MinScore:        60,
			},
			wantErr: false,
		},
		{
			name: "weights don't sum to 1.0",
			profile: Profile{
				ProviderWeight:  0.1,
				ReleaseWeight:   0.1,
				FormatWeight:    0.1,
				MetadataWeight:  0.1, // Sum = 0.4
				MinScore:        50,
			},
			wantErr: true,
		},
		{
			name: "negative weight",
			profile: Profile{
				ProviderWeight:  -0.1,
				ReleaseWeight:   0.4,
				FormatWeight:    0.3,
				MetadataWeight:  0.4,
				MinScore:        50,
			},
			wantErr: true,
		},
		{
			name: "weight too large",
			profile: Profile{
				ProviderWeight:  1.5,
				ReleaseWeight:   0.0,
				FormatWeight:    0.0,
				MetadataWeight:  0.0,
				MinScore:        50,
			},
			wantErr: true,
		},
		{
			name: "invalid min score",
			profile: Profile{
				ProviderWeight:  0.25,
				ReleaseWeight:   0.25,
				FormatWeight:    0.25,
				MetadataWeight:  0.25,
				MinScore:        150, // > 100
			},
			wantErr: true,
		},
		{
			name: "negative min score",
			profile: Profile{
				ProviderWeight:  0.25,
				ReleaseWeight:   0.25,
				FormatWeight:    0.25,
				MetadataWeight:  0.25,
				MinScore:        -10,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProfile(tt.profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}