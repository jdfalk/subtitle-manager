// file: pkg/profiles/language.go
// version: 1.0.0
// guid: 8b7a6c5d-4e3f-9a8b-2c1d-5e4f6a9b8c7d

// Package profiles provides language profile management for subtitle preferences and quality thresholds.
// Similar to Bazarr's language profiles, this allows users to define preferred languages with priority ordering.
package profiles

import (
	"encoding/json"
	"fmt"
	"time"
)

// LanguageProfile represents a collection of language preferences with priority ordering and quality thresholds.
// Users can define multiple profiles for different content types (movies, TV shows, etc.)
type LanguageProfile struct {
	ID          string           `json:"id" db:"id"`
	Name        string           `json:"name" db:"name"`
	Languages   []LanguageConfig `json:"languages" db:"config"`
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

// GetEntityType implements TaggedEntity interface for universal tagging support.
func (lp *LanguageProfile) GetEntityType() string {
	return "language_profile"
}

// GetEntityID implements TaggedEntity interface for universal tagging support.
func (lp *LanguageProfile) GetEntityID() string {
	return lp.ID
}

// GetTags implements TaggedEntity interface (placeholder for future tag integration).
func (lp *LanguageProfile) GetTags() []string {
	return []string{}
}

// SetTags implements TaggedEntity interface (placeholder for future tag integration).
func (lp *LanguageProfile) SetTags(tags []string) {
	// Future implementation for tag assignment
}

// ValidateProfile checks if the language profile configuration is valid.
func (lp *LanguageProfile) ValidateProfile() error {
	if lp.Name == "" {
		return &ValidationError{Field: "name", Message: "Name cannot be empty"}
	}
	if len(lp.Languages) == 0 {
		return &ValidationError{Field: "languages", Message: "At least one language must be specified"}
	}
	if lp.CutoffScore < 0 || lp.CutoffScore > 100 {
		return &ValidationError{Field: "cutoff_score", Message: "Cutoff score must be between 0 and 100"}
	}
	priorities := make(map[int]bool)
	for i, lang := range lp.Languages {
		if lang.Language == "" {
			return &ValidationError{Field: "languages", Message: "Language code cannot be empty", Index: i}
		}
		if priorities[lang.Priority] {
			return &ValidationError{Field: "languages", Message: "Duplicate priority", Index: i}
		}
		priorities[lang.Priority] = true
	}
	return nil
}

// Validate is a backward-compatible wrapper for ValidateProfile.
func (lp *LanguageProfile) Validate() error {
	return lp.ValidateProfile()
}

// GetPrimaryLanguage returns the language with the highest priority (lowest priority number).
func (lp *LanguageProfile) GetPrimaryLanguage() *LanguageConfig {
	if len(lp.Languages) == 0 {
		return nil
	}
	primary := &lp.Languages[0]
	for i := 1; i < len(lp.Languages); i++ {
		if lp.Languages[i].Priority < primary.Priority {
			primary = &lp.Languages[i]
		}
	}
	return primary
}

// MarshalConfig serializes the languages slice to JSON for database storage.
func (lp *LanguageProfile) MarshalConfig() ([]byte, error) {
	return json.Marshal(lp.Languages)
}

// UnmarshalConfig deserializes the languages slice from JSON database storage.
func (lp *LanguageProfile) UnmarshalConfig(data []byte) error {
	return json.Unmarshal(data, &lp.Languages)
}

// ToJSON serializes the language profile to a JSON string.
func (lp *LanguageProfile) ToJSON() (string, error) {
	b, err := json.Marshal(lp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJSON deserializes the language profile from a JSON string.
func (lp *LanguageProfile) FromJSON(configJSON string) error {
	return json.Unmarshal([]byte(configJSON), lp)
}

// GetLanguageConfig returns the configuration for a specific language if present.
func (lp *LanguageProfile) GetLanguageConfig(languageCode string) *LanguageConfig {
	for i := range lp.Languages {
		if lp.Languages[i].Language == languageCode {
			return &lp.Languages[i]
		}
	}
	return nil
}

// ValidationError represents a validation error for language profiles.
type ValidationError struct {
	Field   string
	Message string
	Index   int // Optional index for array field errors
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	if e.Index >= 0 {
		return fmt.Sprintf("%s[%d]: %s", e.Field, e.Index, e.Message)
	}
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// DefaultProfile creates a default English language profile.
func DefaultProfile() *LanguageProfile {
	return &LanguageProfile{
		ID:          "default",
		Name:        "Default",
		Languages:   []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}},
		CutoffScore: 80,
		IsDefault:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// DefaultLanguageProfile is an alias for DefaultProfile for backward compatibility.
func DefaultLanguageProfile() *LanguageProfile {
	return DefaultProfile()
}
