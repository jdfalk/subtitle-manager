// file: pkg/database/scoring_test.go
// version: 1.0.0
// guid: f6g7h8i9-j0k1-2345-fgbc-789012345cde

package database

import (
	"testing"
)

// TestDefaultScoringWeights verifies the default scoring weights are reasonable.
func TestDefaultScoringWeights(t *testing.T) {
	weights := DefaultScoringWeights()

	// Check individual weights are in valid range
	if weights.LanguageMatch < 0 || weights.LanguageMatch > 1 {
		t.Errorf("LanguageMatch weight out of range: %f", weights.LanguageMatch)
	}

	if weights.ProviderRank < 0 || weights.ProviderRank > 1 {
		t.Errorf("ProviderRank weight out of range: %f", weights.ProviderRank)
	}

	// Check that weights sum to approximately 1.0
	total := weights.LanguageMatch + weights.ProviderRank + weights.ReleaseMatch + weights.FormatMatch + weights.UserRating
	if total < 0.99 || total > 1.01 {
		t.Errorf("weights should sum to ~1.0, got %f", total)
	}

	// Language match should have the highest weight (it's most important)
	if weights.LanguageMatch < weights.ProviderRank ||
		weights.LanguageMatch < weights.ReleaseMatch ||
		weights.LanguageMatch < weights.FormatMatch ||
		weights.LanguageMatch < weights.UserRating {
		t.Error("LanguageMatch should have the highest weight")
	}
}

// TestScoreCalculator verifies score calculation logic.
func TestScoreCalculator(t *testing.T) {
	calculator := NewScoreCalculator(DefaultScoringWeights())

	if calculator.version != "1.0" {
		t.Errorf("expected version '1.0', got %s", calculator.version)
	}

	// Test language matching
	tests := []struct {
		requested, provided string
		expectedScore       float64
	}{
		{"en", "en", 1.0},
		{"en", "eng", 0.95},
		{"eng", "en", 0.95},
		{"en", "es", 0.0},
		{"fr", "fr-FR", 0.8}, // Partial match
		{"", "", 0.0},
	}

	for _, test := range tests {
		score := calculator.CalculateLanguageMatch(test.requested, test.provided)
		if score != test.expectedScore {
			t.Errorf("CalculateLanguageMatch(%q, %q) = %f, expected %f",
				test.requested, test.provided, score, test.expectedScore)
		}
	}
}

// TestProviderRanking verifies provider ranking logic.
func TestProviderRanking(t *testing.T) {
	calculator := NewScoreCalculator(DefaultScoringWeights())

	tests := []struct {
		provider      string
		expectedScore float64
	}{
		{"opensubtitles", 0.9},
		{"addic7ed", 0.95},
		{"manual", 1.0},
		{"whisper", 0.7},
		{"unknown_provider", 0.5},
	}

	for _, test := range tests {
		score := calculator.CalculateProviderRank(test.provider)
		if score != test.expectedScore {
			t.Errorf("CalculateProviderRank(%q) = %f, expected %f",
				test.provider, score, test.expectedScore)
		}
	}
}

// TestReleaseMatching verifies release name matching logic.
func TestReleaseMatching(t *testing.T) {
	calculator := NewScoreCalculator(DefaultScoringWeights())

	tests := []struct {
		media, subtitle string
		expectedMin     float64 // Minimum expected score
	}{
		{"Movie.2023.BluRay.x264-GROUP", "Movie.2023.BluRay.x264-GROUP", 1.0},
		{"Movie.2023.BluRay.x264-GROUP", "Movie.2023.WEB-DL.x264-OTHER", 0.5},
		{"Movie.2023.BluRay.x264-GROUP", "Different.Movie.2023", 0.1},
		{"", "Some.Release", 0.4}, // No media release info
		{"Some.Release", "", 0.4}, // No subtitle release info
	}

	for _, test := range tests {
		score := calculator.CalculateReleaseMatch(test.media, test.subtitle)
		if score < test.expectedMin {
			t.Errorf("CalculateReleaseMatch(%q, %q) = %f, expected >= %f",
				test.media, test.subtitle, score, test.expectedMin)
		}
	}
}

// TestFormatMatching verifies format quality scoring.
func TestFormatMatching(t *testing.T) {
	calculator := NewScoreCalculator(DefaultScoringWeights())

	tests := []struct {
		format        string
		expectedScore float64
	}{
		{"srt", 1.0},
		{"SRT", 1.0}, // Case insensitive
		{"ass", 0.9},
		{"vtt", 0.8},
		{"unknown", 0.5},
	}

	for _, test := range tests {
		score := calculator.CalculateFormatMatch(test.format)
		if score != test.expectedScore {
			t.Errorf("CalculateFormatMatch(%q) = %f, expected %f",
				test.format, score, test.expectedScore)
		}
	}
}

// TestCompleteScoreCalculation verifies the complete scoring calculation.
func TestCompleteScoreCalculation(t *testing.T) {
	calculator := NewScoreCalculator(DefaultScoringWeights())

	// Test perfect match
	score := calculator.CalculateSubtitleScore(
		"en", "en", // Perfect language match
		"opensubtitles",                          // Good provider
		"Movie.2023.BluRay", "Movie.2023.BluRay", // Perfect release match
		"srt", // Best format
		1.0,   // Perfect user rating
	)

	if score.TotalScore < 0.9 {
		t.Errorf("perfect match should have high score, got %f", score.TotalScore)
	}

	if score.LanguageMatch != 1.0 {
		t.Errorf("expected perfect language match (1.0), got %f", score.LanguageMatch)
	}

	if score.ProviderName != "opensubtitles" {
		t.Errorf("expected provider 'opensubtitles', got %s", score.ProviderName)
	}

	// Test poor match
	poorScore := calculator.CalculateSubtitleScore(
		"en", "es", // No language match
		"unknown",                          // Unknown provider
		"Movie.2023.BluRay", "Other.Movie", // No release match
		"unknown", // Unknown format
		0.0,       // No user rating
	)

	if poorScore.TotalScore > 0.6 {
		t.Errorf("poor match should have low score, got %f", poorScore.TotalScore)
	}
}

// TestSubtitleScoreCalculateScore verifies the SubtitleScore.CalculateScore method.
func TestSubtitleScoreCalculateScore(t *testing.T) {
	score := &SubtitleScore{
		LanguageMatch: 1.0,
		ProviderRank:  0.9,
		ReleaseMatch:  0.8,
		FormatMatch:   1.0,
		UserRating:    0.9,
	}

	weights := DefaultScoringWeights()
	total := score.CalculateScore(weights)

	// Manually calculate expected score
	expected := 1.0*weights.LanguageMatch +
		0.9*weights.ProviderRank +
		0.8*weights.ReleaseMatch +
		1.0*weights.FormatMatch +
		0.9*weights.UserRating

	if total != expected {
		t.Errorf("CalculateScore() = %f, expected %f", total, expected)
	}

	if score.TotalScore != total {
		t.Errorf("TotalScore not updated correctly: %f != %f", score.TotalScore, total)
	}
}

// TestSubtitleScoringSQLiteIntegration tests score operations with SQLite.
func TestSubtitleScoringSQLiteIntegration(t *testing.T) {
	if !HasSQLite() {
		t.Skip("SQLite not available")
	}

	store, err := OpenSQLStore(":memory:")
	if err != nil {
		t.Fatalf("failed to open SQLite store: %v", err)
	}
	defer store.Close()

	// First create a subtitle record
	subtitle := &SubtitleRecord{
		File:      "test.srt",
		VideoFile: "test.mkv",
		Language:  "en",
		Service:   "opensubtitles",
	}
	err = store.InsertSubtitle(subtitle)
	if err != nil {
		t.Fatalf("failed to insert subtitle: %v", err)
	}

	// Get the subtitle ID
	subtitles, err := store.ListSubtitles()
	if err != nil {
		t.Fatalf("failed to list subtitles: %v", err)
	}
	if len(subtitles) == 0 {
		t.Fatal("no subtitles found")
	}
	subtitleID := subtitles[0].ID

	// Create and insert a score
	score := &SubtitleScore{
		SubtitleID:    subtitleID,
		ProviderName:  "opensubtitles",
		LanguageMatch: 1.0,
		ProviderRank:  0.9,
		ReleaseMatch:  0.8,
		FormatMatch:   1.0,
		UserRating:    0.9,
		TotalScore:    0.92,
		ScoreVersion:  "1.0",
		Metadata:      map[string]interface{}{"test": "data"},
	}

	err = store.InsertSubtitleScore(score)
	if err != nil {
		t.Fatalf("failed to insert subtitle score: %v", err)
	}

	// Test retrieving score by subtitle ID
	retrievedScore, err := store.GetSubtitleScoreBySubtitleID(subtitleID)
	if err != nil {
		t.Fatalf("failed to get subtitle score: %v", err)
	}

	if retrievedScore.SubtitleID != subtitleID {
		t.Errorf("expected subtitle ID %s, got %s", subtitleID, retrievedScore.SubtitleID)
	}

	if retrievedScore.ProviderName != "opensubtitles" {
		t.Errorf("expected provider 'opensubtitles', got %s", retrievedScore.ProviderName)
	}

	if retrievedScore.TotalScore != 0.92 {
		t.Errorf("expected total score 0.92, got %f", retrievedScore.TotalScore)
	}

	// Test listing all scores
	scores, err := store.ListSubtitleScores()
	if err != nil {
		t.Fatalf("failed to list subtitle scores: %v", err)
	}

	if len(scores) != 1 {
		t.Fatalf("expected 1 score, got %d", len(scores))
	}

	// Test updating user rating
	err = store.UpdateUserRating(retrievedScore.ID, 0.95)
	if err != nil {
		t.Fatalf("failed to update user rating: %v", err)
	}

	// Verify rating was updated
	updatedScore, err := store.GetSubtitleScore(retrievedScore.ID)
	if err != nil {
		t.Fatalf("failed to get updated score: %v", err)
	}

	if updatedScore.UserRating != 0.95 {
		t.Errorf("expected user rating 0.95, got %f", updatedScore.UserRating)
	}

	// Test incrementing download count
	originalCount := updatedScore.DownloadCount
	err = store.IncrementDownloadCount(retrievedScore.ID)
	if err != nil {
		t.Fatalf("failed to increment download count: %v", err)
	}

	// Verify download count was incremented
	finalScore, err := store.GetSubtitleScore(retrievedScore.ID)
	if err != nil {
		t.Fatalf("failed to get final score: %v", err)
	}

	if finalScore.DownloadCount != originalCount+1 {
		t.Errorf("expected download count %d, got %d", originalCount+1, finalScore.DownloadCount)
	}
}
