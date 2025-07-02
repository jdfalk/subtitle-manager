// file: pkg/profiles/service.go
// version: 1.0.0
// guid: e3f4a5b6-7c8d-9e0f-1a2b-3c4d5e6f7a8b

package profiles

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Service provides language profile management operations.
type Service struct {
	db *sql.DB
}

// NewService creates a new language profiles service.
func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

// CreateProfile creates a new language profile.
func (s *Service) CreateProfile(profile *LanguageProfile) error {
	if err := profile.ValidateProfile(); err != nil {
		return err
	}

	if profile.ID == "" {
		profile.ID = uuid.New().String()
	}

	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()

	configJSON, err := profile.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to serialize profile config: %w", err)
	}

	// If this is marked as default, unset other defaults
	if profile.IsDefault {
		if err := s.clearDefaultProfiles(); err != nil {
			return fmt.Errorf("failed to clear existing default profiles: %w", err)
		}
	}

	_, err = s.db.Exec(`
		INSERT INTO language_profiles (id, name, config, cutoff_score, is_default, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		profile.ID, profile.Name, configJSON, profile.CutoffScore, profile.IsDefault,
		profile.CreatedAt, profile.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create language profile: %w", err)
	}

	return nil
}

// GetProfile retrieves a language profile by ID.
func (s *Service) GetProfile(id string) (*LanguageProfile, error) {
	var profile LanguageProfile
	var configJSON string

	err := s.db.QueryRow(`
		SELECT id, name, config, cutoff_score, is_default, created_at, updated_at
		FROM language_profiles WHERE id = ?`, id).Scan(
		&profile.ID, &profile.Name, &configJSON, &profile.CutoffScore,
		&profile.IsDefault, &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("language profile not found: %s", id)
		}
		return nil, fmt.Errorf("failed to get language profile: %w", err)
	}

	if err := profile.FromJSON(configJSON); err != nil {
		return nil, fmt.Errorf("failed to deserialize profile config: %w", err)
	}

	return &profile, nil
}

// ListProfiles retrieves all language profiles.
func (s *Service) ListProfiles() ([]*LanguageProfile, error) {
	rows, err := s.db.Query(`
		SELECT id, name, config, cutoff_score, is_default, created_at, updated_at
		FROM language_profiles ORDER BY is_default DESC, name ASC`)
	if err != nil {
		return nil, fmt.Errorf("failed to list language profiles: %w", err)
	}
	defer rows.Close()

	var profiles []*LanguageProfile
	for rows.Next() {
		var profile LanguageProfile
		var configJSON string

		err := rows.Scan(&profile.ID, &profile.Name, &configJSON,
			&profile.CutoffScore, &profile.IsDefault, &profile.CreatedAt, &profile.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan profile row: %w", err)
		}

		if err := profile.FromJSON(configJSON); err != nil {
			return nil, fmt.Errorf("failed to deserialize profile config for %s: %w", profile.ID, err)
		}

		profiles = append(profiles, &profile)
	}

	return profiles, rows.Err()
}

// UpdateProfile updates an existing language profile.
func (s *Service) UpdateProfile(profile *LanguageProfile) error {
	if err := profile.ValidateProfile(); err != nil {
		return err
	}

	profile.UpdatedAt = time.Now()

	configJSON, err := profile.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to serialize profile config: %w", err)
	}

	// If this is marked as default, unset other defaults
	if profile.IsDefault {
		if err := s.clearDefaultProfiles(); err != nil {
			return fmt.Errorf("failed to clear existing default profiles: %w", err)
		}
	}

	result, err := s.db.Exec(`
		UPDATE language_profiles 
		SET name = ?, config = ?, cutoff_score = ?, is_default = ?, updated_at = ?
		WHERE id = ?`,
		profile.Name, configJSON, profile.CutoffScore, profile.IsDefault,
		profile.UpdatedAt, profile.ID)

	if err != nil {
		return fmt.Errorf("failed to update language profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("language profile not found: %s", profile.ID)
	}

	return nil
}

// DeleteProfile removes a language profile.
func (s *Service) DeleteProfile(id string) error {
	// Check if this profile is assigned to any media
	var count int
	err := s.db.QueryRow(`SELECT COUNT(*) FROM media_profiles WHERE profile_id = ?`, id).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check media profile assignments: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("cannot delete profile %s: it is assigned to %d media items", id, count)
	}

	result, err := s.db.Exec(`DELETE FROM language_profiles WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete language profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("language profile not found: %s", id)
	}

	return nil
}

// GetDefaultProfile retrieves the default language profile.
func (s *Service) GetDefaultProfile() (*LanguageProfile, error) {
	var profile LanguageProfile
	var configJSON string

	err := s.db.QueryRow(`
		SELECT id, name, config, cutoff_score, is_default, created_at, updated_at
		FROM language_profiles WHERE is_default = TRUE LIMIT 1`).Scan(
		&profile.ID, &profile.Name, &configJSON, &profile.CutoffScore,
		&profile.IsDefault, &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			// Create and return default profile if none exists
			defaultProfile := DefaultProfile()
			if err := s.CreateProfile(defaultProfile); err != nil {
				return nil, fmt.Errorf("failed to create default profile: %w", err)
			}
			return defaultProfile, nil
		}
		return nil, fmt.Errorf("failed to get default language profile: %w", err)
	}

	if err := profile.FromJSON(configJSON); err != nil {
		return nil, fmt.Errorf("failed to deserialize profile config: %w", err)
	}

	return &profile, nil
}

// AssignProfileToMedia assigns a language profile to a media item.
func (s *Service) AssignProfileToMedia(mediaID, profileID string) error {
	// Verify profile exists
	_, err := s.GetProfile(profileID)
	if err != nil {
		return err
	}

	// Check if media item exists
	var exists bool
	err = s.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM media_items WHERE id = ?)`, mediaID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check media item existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("media item not found: %s", mediaID)
	}

	// Insert or update the assignment
	_, err = s.db.Exec(`
		INSERT INTO media_profiles (media_id, profile_id, created_at)
		VALUES (?, ?, ?)
		ON CONFLICT(media_id) DO UPDATE SET profile_id = excluded.profile_id, created_at = excluded.created_at`,
		mediaID, profileID, time.Now())

	if err != nil {
		return fmt.Errorf("failed to assign profile to media: %w", err)
	}

	return nil
}

// GetMediaProfile retrieves the language profile assigned to a media item.
func (s *Service) GetMediaProfile(mediaID string) (*LanguageProfile, error) {
	var profileID string
	err := s.db.QueryRow(`SELECT profile_id FROM media_profiles WHERE media_id = ?`, mediaID).Scan(&profileID)

	if err != nil {
		if err == sql.ErrNoRows {
			// Return default profile if no specific assignment
			return s.GetDefaultProfile()
		}
		return nil, fmt.Errorf("failed to get media profile assignment: %w", err)
	}

	return s.GetProfile(profileID)
}

// RemoveProfileFromMedia removes a language profile assignment from a media item.
func (s *Service) RemoveProfileFromMedia(mediaID string) error {
	result, err := s.db.Exec(`DELETE FROM media_profiles WHERE media_id = ?`, mediaID)
	if err != nil {
		return fmt.Errorf("failed to remove profile from media: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no profile assignment found for media: %s", mediaID)
	}

	return nil
}

// ListMediaWithProfile returns all media IDs that use a specific profile.
func (s *Service) ListMediaWithProfile(profileID string) ([]string, error) {
	rows, err := s.db.Query(`SELECT media_id FROM media_profiles WHERE profile_id = ?`, profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to list media with profile: %w", err)
	}
	defer rows.Close()

	var mediaIDs []string
	for rows.Next() {
		var mediaID string
		if err := rows.Scan(&mediaID); err != nil {
			return nil, fmt.Errorf("failed to scan media ID: %w", err)
		}
		mediaIDs = append(mediaIDs, mediaID)
	}

	return mediaIDs, rows.Err()
}

// clearDefaultProfiles removes the default flag from all profiles.
func (s *Service) clearDefaultProfiles() error {
	_, err := s.db.Exec(`UPDATE language_profiles SET is_default = FALSE WHERE is_default = TRUE`)
	return err
}

// GetMediaProfileByPath retrieves the language profile for a media item by file path.
func (s *Service) GetMediaProfileByPath(path string) (*LanguageProfile, error) {
	var mediaID int64
	err := s.db.QueryRow(`SELECT id FROM media_items WHERE path = ?`, path).Scan(&mediaID)
	if err != nil {
		if err == sql.ErrNoRows {
			// If media item not found, return default profile
			return s.GetDefaultProfile()
		}
		return nil, fmt.Errorf("failed to get media item by path: %w", err)
	}

	return s.GetMediaProfile(strconv.FormatInt(mediaID, 10))
}
