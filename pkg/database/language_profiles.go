// file: pkg/database/language_profiles.go
// version: 1.0.0
// guid: c3d4e5f6-g7h8-9012-cdef-345678901234

package database

import (
	"encoding/json"
	"strconv"
	"time"
)

// LanguageProfile represents a language profile configuration for media.
// Language profiles define preferred languages, subtitle providers, and scoring criteria
// for different types of media content.
type LanguageProfile struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	Languages     []string               `json:"languages"`      // Ordered list of preferred languages
	Providers     []string               `json:"providers"`      // Ordered list of preferred providers
	SubtitleTypes []string               `json:"subtitle_types"` // e.g., "forced", "sdh", "normal"
	ScoreWeights  map[string]float64     `json:"score_weights"`  // Weights for scoring algorithm
	MinScore      float64                `json:"min_score"`      // Minimum acceptable score
	IsDefault     bool                   `json:"is_default"`     // Whether this is the default profile
	MediaTypes    []string               `json:"media_types"`    // e.g., "movie", "series", "anime"
	Metadata      map[string]interface{} `json:"metadata"`       // Additional profile metadata
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// LanguageProfileAssignment represents the assignment of a language profile to media.
type LanguageProfileAssignment struct {
	ID             string    `json:"id"`
	ProfileID      string    `json:"profile_id"`
	MediaID        string    `json:"media_id"`
	MediaType      string    `json:"media_type"`      // "movie", "series", "episode"
	MediaPath      string    `json:"media_path"`      // Path to media file
	AssignmentType string    `json:"assignment_type"` // "manual", "automatic", "rule-based"
	Priority       int       `json:"priority"`        // Higher priority assignments override lower ones
	CreatedAt      time.Time `json:"created_at"`
}

// LanguageProfileRule represents automatic assignment rules for language profiles.
type LanguageProfileRule struct {
	ID          string                 `json:"id"`
	ProfileID   string                 `json:"profile_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Conditions  map[string]interface{} `json:"conditions"` // JSON conditions for rule matching
	Priority    int                    `json:"priority"`   // Rule evaluation priority
	Enabled     bool                   `json:"enabled"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

// DefaultLanguageProfile returns a sensible default language profile.
func DefaultLanguageProfile() *LanguageProfile {
	return &LanguageProfile{
		Name:          "Default",
		Description:   "Default language profile for subtitle downloads",
		Languages:     []string{"en", "eng"},
		Providers:     []string{"opensubtitles", "subscene"},
		SubtitleTypes: []string{"normal", "sdh"},
		ScoreWeights: map[string]float64{
			"language_match": 0.4,
			"provider_rank":  0.2,
			"release_match":  0.2,
			"format_match":   0.1,
			"user_rating":    0.1,
		},
		MinScore:   0.7,
		IsDefault:  true,
		MediaTypes: []string{"movie", "series"},
		Metadata:   make(map[string]interface{}),
	}
}

// AddLanguageProfileMethods extends SubtitleStore interface for language profiles
type LanguageProfileStore interface {
	// Language Profile CRUD operations
	InsertLanguageProfile(profile *LanguageProfile) error
	GetLanguageProfile(id string) (*LanguageProfile, error)
	GetLanguageProfileByName(name string) (*LanguageProfile, error)
	ListLanguageProfiles() ([]LanguageProfile, error)
	UpdateLanguageProfile(profile *LanguageProfile) error
	DeleteLanguageProfile(id string) error
	GetDefaultLanguageProfile() (*LanguageProfile, error)
	SetDefaultLanguageProfile(id string) error

	// Profile Assignment operations
	AssignProfileToMedia(assignment *LanguageProfileAssignment) error
	GetProfileAssignmentForMedia(mediaPath string) (*LanguageProfileAssignment, error)
	ListProfileAssignments() ([]LanguageProfileAssignment, error)
	RemoveProfileAssignment(id string) error

	// Profile Rule operations
	InsertProfileRule(rule *LanguageProfileRule) error
	GetProfileRule(id string) (*LanguageProfileRule, error)
	ListProfileRules(profileID string) ([]LanguageProfileRule, error)
	UpdateProfileRule(rule *LanguageProfileRule) error
	DeleteProfileRule(id string) error
	ListEnabledProfileRules() ([]LanguageProfileRule, error)
}

// Language Profile methods for SQLStore

// InsertLanguageProfile stores a new language profile.
func (s *SQLStore) InsertLanguageProfile(profile *LanguageProfile) error {
	languagesJSON, err := json.Marshal(profile.Languages)
	if err != nil {
		return err
	}
	providersJSON, err := json.Marshal(profile.Providers)
	if err != nil {
		return err
	}
	subtitleTypesJSON, err := json.Marshal(profile.SubtitleTypes)
	if err != nil {
		return err
	}
	scoreWeightsJSON, err := json.Marshal(profile.ScoreWeights)
	if err != nil {
		return err
	}
	mediaTypesJSON, err := json.Marshal(profile.MediaTypes)
	if err != nil {
		return err
	}
	metadataJSON, err := json.Marshal(profile.Metadata)
	if err != nil {
		return err
	}

	now := time.Now()
	_, err = s.db.Exec(`
		INSERT INTO language_profiles 
		(name, description, languages, providers, subtitle_types, score_weights, min_score, is_default, media_types, metadata, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		profile.Name, profile.Description, string(languagesJSON), string(providersJSON),
		string(subtitleTypesJSON), string(scoreWeightsJSON), profile.MinScore, boolToInt(profile.IsDefault),
		string(mediaTypesJSON), string(metadataJSON), now, now)
	return err
}

// GetLanguageProfile retrieves a language profile by ID.
func (s *SQLStore) GetLanguageProfile(id string) (*LanguageProfile, error) {
	row := s.db.QueryRow(`
		SELECT id, name, description, languages, providers, subtitle_types, score_weights, min_score, is_default, media_types, metadata, created_at, updated_at
		FROM language_profiles WHERE id = ?`, id)

	return s.scanLanguageProfile(row)
}

// GetLanguageProfileByName retrieves a language profile by name.
func (s *SQLStore) GetLanguageProfileByName(name string) (*LanguageProfile, error) {
	row := s.db.QueryRow(`
		SELECT id, name, description, languages, providers, subtitle_types, score_weights, min_score, is_default, media_types, metadata, created_at, updated_at
		FROM language_profiles WHERE name = ?`, name)

	return s.scanLanguageProfile(row)
}

// ListLanguageProfiles retrieves all language profiles.
func (s *SQLStore) ListLanguageProfiles() ([]LanguageProfile, error) {
	rows, err := s.db.Query(`
		SELECT id, name, description, languages, providers, subtitle_types, score_weights, min_score, is_default, media_types, metadata, created_at, updated_at
		FROM language_profiles ORDER BY is_default DESC, name ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []LanguageProfile
	for rows.Next() {
		profile, err := s.scanLanguageProfile(rows)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, *profile)
	}
	return profiles, rows.Err()
}

// UpdateLanguageProfile updates an existing language profile.
func (s *SQLStore) UpdateLanguageProfile(profile *LanguageProfile) error {
	languagesJSON, err := json.Marshal(profile.Languages)
	if err != nil {
		return err
	}
	providersJSON, err := json.Marshal(profile.Providers)
	if err != nil {
		return err
	}
	subtitleTypesJSON, err := json.Marshal(profile.SubtitleTypes)
	if err != nil {
		return err
	}
	scoreWeightsJSON, err := json.Marshal(profile.ScoreWeights)
	if err != nil {
		return err
	}
	mediaTypesJSON, err := json.Marshal(profile.MediaTypes)
	if err != nil {
		return err
	}
	metadataJSON, err := json.Marshal(profile.Metadata)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`
		UPDATE language_profiles SET 
		name = ?, description = ?, languages = ?, providers = ?, subtitle_types = ?, 
		score_weights = ?, min_score = ?, is_default = ?, media_types = ?, metadata = ?, updated_at = ?
		WHERE id = ?`,
		profile.Name, profile.Description, string(languagesJSON), string(providersJSON),
		string(subtitleTypesJSON), string(scoreWeightsJSON), profile.MinScore, boolToInt(profile.IsDefault),
		string(mediaTypesJSON), string(metadataJSON), time.Now(), profile.ID)
	return err
}

// DeleteLanguageProfile removes a language profile.
func (s *SQLStore) DeleteLanguageProfile(id string) error {
	_, err := s.db.Exec(`DELETE FROM language_profiles WHERE id = ?`, id)
	return err
}

// GetDefaultLanguageProfile retrieves the default language profile.
func (s *SQLStore) GetDefaultLanguageProfile() (*LanguageProfile, error) {
	row := s.db.QueryRow(`
		SELECT id, name, description, languages, providers, subtitle_types, score_weights, min_score, is_default, media_types, metadata, created_at, updated_at
		FROM language_profiles WHERE is_default = 1 LIMIT 1`)

	return s.scanLanguageProfile(row)
}

// SetDefaultLanguageProfile sets a profile as the default.
func (s *SQLStore) SetDefaultLanguageProfile(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Clear existing default
	_, err = tx.Exec(`UPDATE language_profiles SET is_default = 0`)
	if err != nil {
		return err
	}

	// Set new default
	_, err = tx.Exec(`UPDATE language_profiles SET is_default = 1 WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// scanLanguageProfile scans a row into a LanguageProfile struct.
func (s *SQLStore) scanLanguageProfile(scanner interface {
	Scan(dest ...interface{}) error
}) (*LanguageProfile, error) {
	var profile LanguageProfile
	var id int64
	var isDefault int
	var languagesJSON, providersJSON, subtitleTypesJSON, scoreWeightsJSON, mediaTypesJSON, metadataJSON string

	err := scanner.Scan(&id, &profile.Name, &profile.Description, &languagesJSON, &providersJSON,
		&subtitleTypesJSON, &scoreWeightsJSON, &profile.MinScore, &isDefault,
		&mediaTypesJSON, &metadataJSON, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		return nil, err
	}

	profile.ID = strconv.FormatInt(id, 10)
	profile.IsDefault = isDefault == 1

	// Unmarshal JSON fields
	if err := json.Unmarshal([]byte(languagesJSON), &profile.Languages); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(providersJSON), &profile.Providers); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(subtitleTypesJSON), &profile.SubtitleTypes); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(scoreWeightsJSON), &profile.ScoreWeights); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(mediaTypesJSON), &profile.MediaTypes); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(metadataJSON), &profile.Metadata); err != nil {
		return nil, err
	}

	return &profile, nil
}

// Profile Assignment methods

// AssignProfileToMedia assigns a language profile to a media item.
func (s *SQLStore) AssignProfileToMedia(assignment *LanguageProfileAssignment) error {
	_, err := s.db.Exec(`
		INSERT INTO language_profile_assignments 
		(profile_id, media_id, media_type, media_path, assignment_type, priority, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		assignment.ProfileID, assignment.MediaID, assignment.MediaType, assignment.MediaPath,
		assignment.AssignmentType, assignment.Priority, time.Now())
	return err
}

// GetProfileAssignmentForMedia retrieves the profile assignment for a media item.
func (s *SQLStore) GetProfileAssignmentForMedia(mediaPath string) (*LanguageProfileAssignment, error) {
	row := s.db.QueryRow(`
		SELECT id, profile_id, media_id, media_type, media_path, assignment_type, priority, created_at
		FROM language_profile_assignments WHERE media_path = ? ORDER BY priority DESC LIMIT 1`, mediaPath)

	var assignment LanguageProfileAssignment
	var id int64
	err := row.Scan(&id, &assignment.ProfileID, &assignment.MediaID, &assignment.MediaType,
		&assignment.MediaPath, &assignment.AssignmentType, &assignment.Priority, &assignment.CreatedAt)
	if err != nil {
		return nil, err
	}
	assignment.ID = strconv.FormatInt(id, 10)
	return &assignment, nil
}

// ListProfileAssignments retrieves all profile assignments.
func (s *SQLStore) ListProfileAssignments() ([]LanguageProfileAssignment, error) {
	rows, err := s.db.Query(`
		SELECT id, profile_id, media_id, media_type, media_path, assignment_type, priority, created_at
		FROM language_profile_assignments ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assignments []LanguageProfileAssignment
	for rows.Next() {
		var assignment LanguageProfileAssignment
		var id int64
		err := rows.Scan(&id, &assignment.ProfileID, &assignment.MediaID, &assignment.MediaType,
			&assignment.MediaPath, &assignment.AssignmentType, &assignment.Priority, &assignment.CreatedAt)
		if err != nil {
			return nil, err
		}
		assignment.ID = strconv.FormatInt(id, 10)
		assignments = append(assignments, assignment)
	}
	return assignments, rows.Err()
}

// RemoveProfileAssignment removes a profile assignment.
func (s *SQLStore) RemoveProfileAssignment(id string) error {
	_, err := s.db.Exec(`DELETE FROM language_profile_assignments WHERE id = ?`, id)
	return err
}

// Profile Rule methods

// InsertProfileRule stores a new profile rule.
func (s *SQLStore) InsertProfileRule(rule *LanguageProfileRule) error {
	conditionsJSON, err := json.Marshal(rule.Conditions)
	if err != nil {
		return err
	}

	now := time.Now()
	_, err = s.db.Exec(`
		INSERT INTO language_profile_rules 
		(profile_id, name, description, conditions, priority, enabled, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		rule.ProfileID, rule.Name, rule.Description, string(conditionsJSON),
		rule.Priority, boolToInt(rule.Enabled), now, now)
	return err
}

// GetProfileRule retrieves a profile rule by ID.
func (s *SQLStore) GetProfileRule(id string) (*LanguageProfileRule, error) {
	row := s.db.QueryRow(`
		SELECT id, profile_id, name, description, conditions, priority, enabled, created_at, updated_at
		FROM language_profile_rules WHERE id = ?`, id)

	return s.scanProfileRule(row)
}

// ListProfileRules retrieves all rules for a profile.
func (s *SQLStore) ListProfileRules(profileID string) ([]LanguageProfileRule, error) {
	rows, err := s.db.Query(`
		SELECT id, profile_id, name, description, conditions, priority, enabled, created_at, updated_at
		FROM language_profile_rules WHERE profile_id = ? ORDER BY priority DESC`, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []LanguageProfileRule
	for rows.Next() {
		rule, err := s.scanProfileRule(rows)
		if err != nil {
			return nil, err
		}
		rules = append(rules, *rule)
	}
	return rules, rows.Err()
}

// UpdateProfileRule updates an existing profile rule.
func (s *SQLStore) UpdateProfileRule(rule *LanguageProfileRule) error {
	conditionsJSON, err := json.Marshal(rule.Conditions)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`
		UPDATE language_profile_rules SET 
		name = ?, description = ?, conditions = ?, priority = ?, enabled = ?, updated_at = ?
		WHERE id = ?`,
		rule.Name, rule.Description, string(conditionsJSON), rule.Priority,
		boolToInt(rule.Enabled), time.Now(), rule.ID)
	return err
}

// DeleteProfileRule removes a profile rule.
func (s *SQLStore) DeleteProfileRule(id string) error {
	_, err := s.db.Exec(`DELETE FROM language_profile_rules WHERE id = ?`, id)
	return err
}

// ListEnabledProfileRules retrieves all enabled profile rules.
func (s *SQLStore) ListEnabledProfileRules() ([]LanguageProfileRule, error) {
	rows, err := s.db.Query(`
		SELECT id, profile_id, name, description, conditions, priority, enabled, created_at, updated_at
		FROM language_profile_rules WHERE enabled = 1 ORDER BY priority DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []LanguageProfileRule
	for rows.Next() {
		rule, err := s.scanProfileRule(rows)
		if err != nil {
			return nil, err
		}
		rules = append(rules, *rule)
	}
	return rules, rows.Err()
}

// scanProfileRule scans a row into a LanguageProfileRule struct.
func (s *SQLStore) scanProfileRule(scanner interface {
	Scan(dest ...interface{}) error
}) (*LanguageProfileRule, error) {
	var rule LanguageProfileRule
	var id int64
	var enabled int
	var conditionsJSON string

	err := scanner.Scan(&id, &rule.ProfileID, &rule.Name, &rule.Description,
		&conditionsJSON, &rule.Priority, &enabled, &rule.CreatedAt, &rule.UpdatedAt)
	if err != nil {
		return nil, err
	}

	rule.ID = strconv.FormatInt(id, 10)
	rule.Enabled = enabled == 1

	if err := json.Unmarshal([]byte(conditionsJSON), &rule.Conditions); err != nil {
		return nil, err
	}

	return &rule, nil
}
