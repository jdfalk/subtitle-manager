// file: pkg/database/language_profiles.go
// version: 1.0.0
// guid: c3d4e5f6-g7h8-9012-cdef-345678901234

package database

import (
	"encoding/json"
	"time"
)

// LanguageProfile represents a collection of language preferences with priority ordering and quality thresholds.
// Users can define multiple profiles for different content types (movies, TV shows, etc.)
type LanguageProfile struct {
	ID          string           `json:"id" db:"id"`
	Name        string           `json:"name" db:"name"`
	Languages   []LanguageConfig `json:"languages" db:"config"`
	Providers   []string         `json:"providers" db:"providers"`
	CutoffScore int              `json:"cutoff_score" db:"cutoff_score"`
	IsDefault   bool             `json:"is_default" db:"is_default"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
}

// LanguageConfig defines the configuration for a specific language within a profile.
// Priority determines the order of preference (lower number = higher priority).
type LanguageConfig struct {
	Language string `json:"language"`
	Priority int    `json:"priority"`
	Forced   bool   `json:"forced"` // Whether forced subtitles are preferred
	HI       bool   `json:"hi"`     // Whether hearing impaired subtitles are preferred
}

// MediaProfileAssignment represents the assignment of a language profile to a media item.
type MediaProfileAssignment struct {
	MediaID   string    `json:"media_id" db:"media_id"`
	ProfileID string    `json:"profile_id" db:"profile_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// MarshalConfig serializes the Languages slice to JSON for database storage.
func (lp *LanguageProfile) MarshalConfig() ([]byte, error) {
	return json.Marshal(lp.Languages)
}

// UnmarshalConfig deserializes the Languages slice from JSON database storage.
func (lp *LanguageProfile) UnmarshalConfig(data []byte) error {
	return json.Unmarshal(data, &lp.Languages)
}

// DefaultLanguageProfile returns a sensible default language profile.
func DefaultLanguageProfile() *LanguageProfile {
	return &LanguageProfile{
		ID:          "default",
		Name:        "Default",
		Languages:   []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}},
		Providers:   []string{},
		CutoffScore: 80,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
