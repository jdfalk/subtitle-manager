// file: pkg/database/scoring.go
// version: 1.0.0  
// guid: d4e5f6g7-h8i9-0123-def4-56789012345a

package database

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// SubtitleScore represents the scoring information for a subtitle.
type SubtitleScore struct {
	ID              string                 `json:"id"`
	SubtitleID      string                 `json:"subtitle_id"`      // Reference to subtitle record
	ProviderName    string                 `json:"provider_name"`    // Provider that supplied the subtitle
	LanguageMatch   float64                `json:"language_match"`   // Language matching score (0-1)
	ProviderRank    float64                `json:"provider_rank"`    // Provider reliability score (0-1) 
	ReleaseMatch    float64                `json:"release_match"`    // Release group/name matching score (0-1)
	FormatMatch     float64                `json:"format_match"`     // Format quality score (0-1)
	UserRating      float64                `json:"user_rating"`      // User-provided rating (0-1)
	DownloadCount   int                    `json:"download_count"`   // Number of downloads
	TotalScore      float64                `json:"total_score"`      // Calculated total score (0-1)
	ScoreVersion    string                 `json:"score_version"`    // Version of scoring algorithm used
	Metadata        map[string]interface{} `json:"metadata"`         // Additional scoring metadata
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

// ScoringWeights defines the weights used in score calculation.
type ScoringWeights struct {
	LanguageMatch float64 `json:"language_match"`
	ProviderRank  float64 `json:"provider_rank"`
	ReleaseMatch  float64 `json:"release_match"`
	FormatMatch   float64 `json:"format_match"`
	UserRating    float64 `json:"user_rating"`
}

// DefaultScoringWeights returns the default weights for subtitle scoring.
func DefaultScoringWeights() ScoringWeights {
	return ScoringWeights{
		LanguageMatch: 0.4,
		ProviderRank:  0.2,
		ReleaseMatch:  0.2,
		FormatMatch:   0.1,
		UserRating:    0.1,
	}
}

// CalculateScore computes the total score using the provided weights.
func (s *SubtitleScore) CalculateScore(weights ScoringWeights) float64 {
	total := s.LanguageMatch*weights.LanguageMatch +
		s.ProviderRank*weights.ProviderRank +
		s.ReleaseMatch*weights.ReleaseMatch +
		s.FormatMatch*weights.FormatMatch +
		s.UserRating*weights.UserRating
	
	s.TotalScore = total
	return total
}

// ScoreCalculator provides methods for calculating subtitle scores.
type ScoreCalculator struct {
	weights ScoringWeights
	version string
}

// NewScoreCalculator creates a new score calculator with the given weights.
func NewScoreCalculator(weights ScoringWeights) *ScoreCalculator {
	return &ScoreCalculator{
		weights: weights,
		version: "1.0", // Current scoring algorithm version
	}
}

// CalculateLanguageMatch calculates language matching score.
func (sc *ScoreCalculator) CalculateLanguageMatch(requested, provided string) float64 {
	if requested == provided {
		return 1.0
	}
	
	// Handle language code variants (e.g., "en" vs "eng")
	if (requested == "en" && provided == "eng") || (requested == "eng" && provided == "en") {
		return 0.95
	}
	
	// Partial matches for language families
	if len(requested) >= 2 && len(provided) >= 2 {
		if requested[:2] == provided[:2] {
			return 0.8
		}
	}
	
	return 0.0
}

// CalculateProviderRank calculates provider reliability score.
func (sc *ScoreCalculator) CalculateProviderRank(providerName string) float64 {
	// Provider rankings based on reliability and quality
	providerRanks := map[string]float64{
		"opensubtitles": 0.9,
		"subscene":      0.85,
		"addic7ed":      0.95,
		"tvsubtitles":   0.8,
		"podnapisi":     0.85,
		"whisper":       0.7, // Lower for AI-generated
		"manual":        1.0, // Highest for manual uploads
	}
	
	if rank, exists := providerRanks[providerName]; exists {
		return rank
	}
	
	return 0.5 // Default for unknown providers
}

// CalculateReleaseMatch calculates release name matching score.
func (sc *ScoreCalculator) CalculateReleaseMatch(mediaRelease, subtitleRelease string) float64 {
	if mediaRelease == "" || subtitleRelease == "" {
		return 0.5 // Neutral if no release info
	}
	
	if mediaRelease == subtitleRelease {
		return 1.0
	}
	
	// Fuzzy matching for release names
	// This is a simplified implementation - could be enhanced with proper fuzzy matching
	mediaLower := strings.ToLower(mediaRelease)
	subtitleLower := strings.ToLower(subtitleRelease)
	
	if strings.Contains(mediaLower, subtitleLower) || strings.Contains(subtitleLower, mediaLower) {
		return 0.8
	}
	
	// Check for common patterns
	if sc.containsCommonReleaseGroup(mediaLower, subtitleLower) {
		return 0.6
	}
	
	return 0.2
}

// containsCommonReleaseGroup checks for common release group patterns.
func (sc *ScoreCalculator) containsCommonReleaseGroup(media, subtitle string) bool {
	commonGroups := []string{"bluray", "web", "hdtv", "dvdrip", "webrip", "x264", "x265", "h264", "h265"}
	
	mediaWords := strings.Fields(media)
	subtitleWords := strings.Fields(subtitle)
	
	for _, group := range commonGroups {
		mediaHas := false
		subtitleHas := false
		
		for _, word := range mediaWords {
			if strings.Contains(word, group) {
				mediaHas = true
				break
			}
		}
		
		for _, word := range subtitleWords {
			if strings.Contains(word, group) {
				subtitleHas = true
				break
			}
		}
		
		if mediaHas && subtitleHas {
			return true
		}
	}
	
	return false
}

// CalculateFormatMatch calculates format quality score.
func (sc *ScoreCalculator) CalculateFormatMatch(format string) float64 {
	// Format quality rankings
	formatScores := map[string]float64{
		"srt": 1.0,   // Most compatible
		"ass": 0.9,   // Advanced features
		"ssa": 0.85,  // Similar to ASS
		"vtt": 0.8,   // Web format
		"sub": 0.7,   // Basic format
		"idx": 0.6,   // Image-based
		"sup": 0.6,   // Image-based
	}
	
	if score, exists := formatScores[strings.ToLower(format)]; exists {
		return score
	}
	
	return 0.5 // Default for unknown formats
}

// CalculateSubtitleScore calculates the complete score for a subtitle.
func (sc *ScoreCalculator) CalculateSubtitleScore(
	requestedLang, providedLang, providerName, mediaRelease, subtitleRelease, format string,
	userRating float64) *SubtitleScore {
	
	score := &SubtitleScore{
		ProviderName:  providerName,
		LanguageMatch: sc.CalculateLanguageMatch(requestedLang, providedLang),
		ProviderRank:  sc.CalculateProviderRank(providerName),
		ReleaseMatch:  sc.CalculateReleaseMatch(mediaRelease, subtitleRelease),
		FormatMatch:   sc.CalculateFormatMatch(format),
		UserRating:    userRating,
		ScoreVersion:  sc.version,
		Metadata:      make(map[string]interface{}),
	}
	
	score.TotalScore = score.CalculateScore(sc.weights)
	return score
}

// Scoring methods for SQLStore

// InsertSubtitleScore stores a subtitle score record.
func (s *SQLStore) InsertSubtitleScore(score *SubtitleScore) error {
	metadataJSON, err := json.Marshal(score.Metadata)
	if err != nil {
		return err
	}

	now := time.Now()
	_, err = s.db.Exec(`
		INSERT INTO subtitle_scores 
		(subtitle_id, provider_name, language_match, provider_rank, release_match, format_match, 
		 user_rating, download_count, total_score, score_version, metadata, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		score.SubtitleID, score.ProviderName, score.LanguageMatch, score.ProviderRank,
		score.ReleaseMatch, score.FormatMatch, score.UserRating, score.DownloadCount,
		score.TotalScore, score.ScoreVersion, string(metadataJSON), now, now)
	return err
}

// GetSubtitleScore retrieves a subtitle score by ID.
func (s *SQLStore) GetSubtitleScore(id string) (*SubtitleScore, error) {
	row := s.db.QueryRow(`
		SELECT id, subtitle_id, provider_name, language_match, provider_rank, release_match, 
		       format_match, user_rating, download_count, total_score, score_version, metadata, 
		       created_at, updated_at
		FROM subtitle_scores WHERE id = ?`, id)
	
	return s.scanSubtitleScore(row)
}

// GetSubtitleScoreBySubtitleID retrieves the score for a subtitle.
func (s *SQLStore) GetSubtitleScoreBySubtitleID(subtitleID string) (*SubtitleScore, error) {
	row := s.db.QueryRow(`
		SELECT id, subtitle_id, provider_name, language_match, provider_rank, release_match, 
		       format_match, user_rating, download_count, total_score, score_version, metadata, 
		       created_at, updated_at
		FROM subtitle_scores WHERE subtitle_id = ?`, subtitleID)
	
	return s.scanSubtitleScore(row)
}

// ListSubtitleScores retrieves all subtitle scores, ordered by total score descending.
func (s *SQLStore) ListSubtitleScores() ([]SubtitleScore, error) {
	rows, err := s.db.Query(`
		SELECT id, subtitle_id, provider_name, language_match, provider_rank, release_match, 
		       format_match, user_rating, download_count, total_score, score_version, metadata, 
		       created_at, updated_at
		FROM subtitle_scores ORDER BY total_score DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []SubtitleScore
	for rows.Next() {
		score, err := s.scanSubtitleScore(rows)
		if err != nil {
			return nil, err
		}
		scores = append(scores, *score)
	}
	return scores, rows.Err()
}

// UpdateSubtitleScore updates an existing subtitle score.
func (s *SQLStore) UpdateSubtitleScore(score *SubtitleScore) error {
	metadataJSON, err := json.Marshal(score.Metadata)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`
		UPDATE subtitle_scores SET 
		provider_name = ?, language_match = ?, provider_rank = ?, release_match = ?, 
		format_match = ?, user_rating = ?, download_count = ?, total_score = ?, 
		score_version = ?, metadata = ?, updated_at = ?
		WHERE id = ?`,
		score.ProviderName, score.LanguageMatch, score.ProviderRank, score.ReleaseMatch,
		score.FormatMatch, score.UserRating, score.DownloadCount, score.TotalScore,
		score.ScoreVersion, string(metadataJSON), time.Now(), score.ID)
	return err
}

// DeleteSubtitleScore removes a subtitle score record.
func (s *SQLStore) DeleteSubtitleScore(id string) error {
	_, err := s.db.Exec(`DELETE FROM subtitle_scores WHERE id = ?`, id)
	return err
}

// GetTopScoredSubtitles retrieves the highest-scored subtitles for a video file.
func (s *SQLStore) GetTopScoredSubtitles(videoFile string, limit int) ([]SubtitleScore, error) {
	rows, err := s.db.Query(`
		SELECT ss.id, ss.subtitle_id, ss.provider_name, ss.language_match, ss.provider_rank, 
		       ss.release_match, ss.format_match, ss.user_rating, ss.download_count, 
		       ss.total_score, ss.score_version, ss.metadata, ss.created_at, ss.updated_at
		FROM subtitle_scores ss
		JOIN subtitles s ON ss.subtitle_id = s.id
		WHERE s.video_file = ?
		ORDER BY ss.total_score DESC
		LIMIT ?`, videoFile, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []SubtitleScore
	for rows.Next() {
		score, err := s.scanSubtitleScore(rows)
		if err != nil {
			return nil, err
		}
		scores = append(scores, *score)
	}
	return scores, rows.Err()
}

// UpdateUserRating updates the user rating for a subtitle score.
func (s *SQLStore) UpdateUserRating(scoreID string, rating float64) error {
	_, err := s.db.Exec(`
		UPDATE subtitle_scores SET user_rating = ?, updated_at = ? WHERE id = ?`,
		rating, time.Now(), scoreID)
	return err
}

// IncrementDownloadCount increments the download count for a subtitle score.
func (s *SQLStore) IncrementDownloadCount(scoreID string) error {
	_, err := s.db.Exec(`
		UPDATE subtitle_scores SET download_count = download_count + 1, updated_at = ? WHERE id = ?`,
		time.Now(), scoreID)
	return err
}

// scanSubtitleScore scans a row into a SubtitleScore struct.
func (s *SQLStore) scanSubtitleScore(scanner interface {
	Scan(dest ...interface{}) error
}) (*SubtitleScore, error) {
	var score SubtitleScore
	var id int64
	var metadataJSON string

	err := scanner.Scan(&id, &score.SubtitleID, &score.ProviderName, &score.LanguageMatch,
		&score.ProviderRank, &score.ReleaseMatch, &score.FormatMatch, &score.UserRating,
		&score.DownloadCount, &score.TotalScore, &score.ScoreVersion, &metadataJSON,
		&score.CreatedAt, &score.UpdatedAt)
	if err != nil {
		return nil, err
	}

	score.ID = strconv.FormatInt(id, 10)

	if err := json.Unmarshal([]byte(metadataJSON), &score.Metadata); err != nil {
		return nil, err
	}

	return &score, nil
}