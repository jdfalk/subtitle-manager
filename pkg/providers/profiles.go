// file: pkg/providers/profiles.go
// version: 1.0.0
// guid: f1a2b3c4-d5e6-7f8a-9b0c-1d2e3f4a5b6c

package providers

import (
	"context"
	"database/sql"
	"fmt"
	"sort"

	"github.com/jdfalk/subtitle-manager/pkg/profiles"
)

// FetchWithProfile attempts to fetch subtitles using language profile preferences.
// It tries languages in priority order from the profile associated with the media file.
func FetchWithProfile(ctx context.Context, db *sql.DB, mediaPath, key string) ([]byte, string, string, error) {
	service := profiles.NewService(db)
	
	// Get the language profile for this media file
	profile, err := service.GetMediaProfileByPath(mediaPath)
	if err != nil {
		// Fallback to default fetch if no profile found
		data, provider, fetchErr := FetchFromAll(ctx, mediaPath, "en", key)
		if fetchErr != nil {
			return nil, "", "", fmt.Errorf("failed to get profile (%v) and fallback fetch failed: %w", err, fetchErr)
		}
		return data, provider, "en", nil
	}

	// Sort languages by priority (lower number = higher priority)
	languages := make([]profiles.LanguageConfig, len(profile.Languages))
	copy(languages, profile.Languages)
	sort.Slice(languages, func(i, j int) bool {
		return languages[i].Priority < languages[j].Priority
	})

	// Try each language in priority order
	var lastErr error
	for _, langConfig := range languages {
		data, provider, err := FetchFromAll(ctx, mediaPath, langConfig.Language, key)
		if err == nil {
			return data, provider, langConfig.Language, nil
		}
		lastErr = err
		
		// Check if context was cancelled
		if ctx.Err() != nil {
			return nil, "", "", ctx.Err()
		}
	}

	return nil, "", "", fmt.Errorf("no subtitles found for any language in profile '%s': %w", profile.Name, lastErr)
}

// FetchWithProfileTagged attempts to fetch subtitles using both profile preferences and provider tags.
func FetchWithProfileTagged(ctx context.Context, db *sql.DB, mediaPath, key string, tags []string, tm interface {
	FilterByTags(string, []string) ([]string, error)
}) ([]byte, string, string, error) {
	service := profiles.NewService(db)
	
	// Get the language profile for this media file
	profile, err := service.GetMediaProfileByPath(mediaPath)
	if err != nil {
		// Fallback to default tagged fetch if no profile found
		data, provider, fetchErr := FetchFromTagged(ctx, mediaPath, "en", key, tags, tm)
		if fetchErr != nil {
			return nil, "", "", fmt.Errorf("failed to get profile (%v) and fallback fetch failed: %w", err, fetchErr)
		}
		return data, provider, "en", nil
	}

	// Sort languages by priority (lower number = higher priority)
	languages := make([]profiles.LanguageConfig, len(profile.Languages))
	copy(languages, profile.Languages)
	sort.Slice(languages, func(i, j int) bool {
		return languages[i].Priority < languages[j].Priority
	})

	// Try each language in priority order with tagged providers
	var lastErr error
	for _, langConfig := range languages {
		data, provider, err := FetchFromTagged(ctx, mediaPath, langConfig.Language, key, tags, tm)
		if err == nil {
			return data, provider, langConfig.Language, nil
		}
		lastErr = err
		
		// Check if context was cancelled
		if ctx.Err() != nil {
			return nil, "", "", ctx.Err()
		}
	}

	return nil, "", "", fmt.Errorf("no subtitles found for any language in profile '%s' with tags %v: %w", profile.Name, tags, lastErr)
}

// GetLanguagesFromProfile extracts an ordered list of language codes from a profile.
func GetLanguagesFromProfile(ctx context.Context, db *sql.DB, mediaPath string) ([]string, error) {
	service := profiles.NewService(db)
	
	profile, err := service.GetMediaProfileByPath(mediaPath)
	if err != nil {
		return []string{"en"}, err // fallback to English
	}

	// Sort languages by priority
	languages := make([]profiles.LanguageConfig, len(profile.Languages))
	copy(languages, profile.Languages)
	sort.Slice(languages, func(i, j int) bool {
		return languages[i].Priority < languages[j].Priority
	})

	// Extract language codes
	codes := make([]string, len(languages))
	for i, lang := range languages {
		codes[i] = lang.Language
	}

	return codes, nil
}