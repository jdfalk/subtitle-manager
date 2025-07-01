// file: pkg/scoring/scorer.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174000
package scoring

import (
	"strings"
	"time"
)

// SubtitleScore represents the calculated quality score for a subtitle.
type SubtitleScore struct {
	Total         int `json:"total"`
	ProviderScore int `json:"providerScore"`
	ReleaseScore  int `json:"releaseScore"`
	FormatScore   int `json:"formatScore"`
	MetadataScore int `json:"metadataScore"`
}

// Subtitle represents a subtitle candidate with metadata for scoring.
type Subtitle struct {
	// Provider information
	ProviderName string `json:"providerName"`
	IsTrusted    bool   `json:"isTrusted"`

	// Release information
	Release  string `json:"release"`
	FileName string `json:"fileName"`

	// Quality metadata
	Format          string    `json:"format"`
	HearingImpaired bool      `json:"hearingImpaired"`
	ForcedSubtitle  bool      `json:"forcedSubtitle"`
	UploadDate      time.Time `json:"uploadDate"`
	DownloadCount   int       `json:"downloadCount"`
	Rating          float64   `json:"rating"`
	Votes           int       `json:"votes"`
	FileSize        int64     `json:"fileSize"`

	// Additional flags
	AutoTranslated    bool `json:"autoTranslated"`
	MachineTranslated bool `json:"machineTranslated"`
	HD                bool `json:"hd"`
}

// MediaItem represents the media file for context in scoring.
type MediaItem struct {
	Title        string `json:"title"`
	Season       int    `json:"season"`
	Episode      int    `json:"episode"`
	ReleaseGroup string `json:"releaseGroup"`
	Resolution   string `json:"resolution"`
	Source       string `json:"source"` // bluray, web-dl, hdtv, etc.
	Codec        string `json:"codec"`
	FileSize     int64  `json:"fileSize"`
}

// Profile represents user preferences for subtitle scoring.
type Profile struct {
	// Scoring weights (0.0 to 1.0)
	ProviderWeight float64 `json:"providerWeight"`
	ReleaseWeight  float64 `json:"releaseWeight"`
	FormatWeight   float64 `json:"formatWeight"`
	MetadataWeight float64 `json:"metadataWeight"`

	// Preferences
	PreferredFormats []string `json:"preferredFormats"`
	AllowHI          bool     `json:"allowHI"`
	PreferHI         bool     `json:"preferHI"`
	AllowForced      bool     `json:"allowForced"`
	PreferForced     bool     `json:"preferForced"`

	// Thresholds
	MinScore int           `json:"minScore"`
	MaxAge   time.Duration `json:"maxAge"`
}

// DefaultProfile returns a profile with sensible default scoring weights.
func DefaultProfile() Profile {
	return Profile{
		ProviderWeight:   0.25,
		ReleaseWeight:    0.35,
		FormatWeight:     0.15,
		MetadataWeight:   0.25,
		PreferredFormats: []string{"srt", "ass", "ssa"},
		AllowHI:          true,
		PreferHI:         false,
		AllowForced:      true,
		PreferForced:     false,
		MinScore:         50,
		MaxAge:           365 * 24 * time.Hour, // 1 year
	}
}

// CalculateScore evaluates subtitle quality based on various criteria.
func CalculateScore(subtitle Subtitle, media MediaItem, profile Profile) SubtitleScore {
	providerScore := calculateProviderScore(subtitle, profile)
	releaseScore := calculateReleaseScore(subtitle, media, profile)
	formatScore := calculateFormatScore(subtitle, profile)
	metadataScore := calculateMetadataScore(subtitle, media, profile)

	// Apply weights and calculate total
	total := int(
		float64(providerScore)*profile.ProviderWeight +
			float64(releaseScore)*profile.ReleaseWeight +
			float64(formatScore)*profile.FormatWeight +
			float64(metadataScore)*profile.MetadataWeight,
	)

	return SubtitleScore{
		Total:         total,
		ProviderScore: providerScore,
		ReleaseScore:  releaseScore,
		FormatScore:   formatScore,
		MetadataScore: metadataScore,
	}
}

// calculateProviderScore evaluates provider reliability and trustworthiness.
func calculateProviderScore(subtitle Subtitle, profile Profile) int {
	score := 50 // Base score

	// Trusted providers get a bonus
	if subtitle.IsTrusted {
		score += 30
	}

	// Provider-specific scoring
	switch strings.ToLower(subtitle.ProviderName) {
	case "opensubtitles", "opensubtitlescom":
		score += 20
	case "subscene", "addic7ed":
		score += 15
	case "podnapisi", "tvsubtitles":
		score += 10
	default:
		score += 5
	}

	// Penalty for machine-translated content
	if subtitle.MachineTranslated {
		score -= 20
	}
	if subtitle.AutoTranslated {
		score -= 10
	}

	return clampScore(score)
}

// calculateReleaseScore evaluates how well the subtitle matches the media release.
func calculateReleaseScore(subtitle Subtitle, media MediaItem, profile Profile) int {
	score := 50 // Base score

	subtitleRelease := strings.ToLower(subtitle.Release)
	mediaSource := strings.ToLower(media.Source)
	mediaGroup := strings.ToLower(media.ReleaseGroup)

	// Perfect release group match
	if mediaGroup != "" && strings.Contains(subtitleRelease, mediaGroup) {
		score += 40
		return clampScore(score)
	}

	// Source quality matching
	if mediaSource != "" {
		if strings.Contains(subtitleRelease, mediaSource) {
			score += 30
		} else {
			// Partial matches based on quality hierarchy
			switch mediaSource {
			case "bluray", "blu-ray":
				if strings.Contains(subtitleRelease, "remux") ||
					strings.Contains(subtitleRelease, "bdremux") {
					score += 25
				} else if strings.Contains(subtitleRelease, "bdrip") ||
					strings.Contains(subtitleRelease, "brrip") {
					score += 20
				}
			case "web-dl", "webdl":
				if strings.Contains(subtitleRelease, "web") {
					score += 20
				}
			case "hdtv":
				if strings.Contains(subtitleRelease, "hdtv") {
					score += 20
				}
			}
		}
	}

	// Resolution matching
	if media.Resolution != "" {
		if strings.Contains(subtitleRelease, media.Resolution) {
			score += 15
		}
	}

	// Codec matching
	if media.Codec != "" {
		codecLower := strings.ToLower(media.Codec)
		if strings.Contains(subtitleRelease, codecLower) {
			score += 10
		}
	}

	// Penalty for mismatched quality indicators
	if mediaSource == "bluray" && strings.Contains(subtitleRelease, "cam") {
		score -= 30
	}
	if mediaSource == "web-dl" && strings.Contains(subtitleRelease, "ts") {
		score -= 25
	}

	return clampScore(score)
}

// calculateFormatScore evaluates subtitle format preferences.
func calculateFormatScore(subtitle Subtitle, profile Profile) int {
	score := 50 // Base score

	format := strings.ToLower(subtitle.Format)

	// Check preferred formats
	for i, preferred := range profile.PreferredFormats {
		if format == strings.ToLower(preferred) {
			// Higher score for earlier preferences
			score += 30 - (i * 5)
			break
		}
	}

	// Default format scoring if not in preferences
	if len(profile.PreferredFormats) == 0 || !contains(profile.PreferredFormats, format) {
		switch format {
		case "srt":
			score += 25
		case "ass", "ssa":
			score += 20
		case "vtt":
			score += 15
		case "sub", "idx":
			score += 10
		default:
			score += 5
		}
	}

	return clampScore(score)
}

// calculateMetadataScore evaluates additional quality indicators.
func calculateMetadataScore(subtitle Subtitle, media MediaItem, profile Profile) int {
	score := 50 // Base score

	// Age scoring (newer is better, but with diminishing returns)
	if !subtitle.UploadDate.IsZero() {
		age := time.Since(subtitle.UploadDate)
		if age <= 7*24*time.Hour { // 1 week
			score += 20
		} else if age <= 30*24*time.Hour { // 1 month
			score += 15
		} else if age <= 90*24*time.Hour { // 3 months
			score += 10
		} else if age <= 365*24*time.Hour { // 1 year
			score += 5
		} else if age > profile.MaxAge {
			score -= 20 // Penalty for very old subtitles
		}
	}

	// Download popularity (logarithmic scaling)
	if subtitle.DownloadCount > 0 {
		// Convert to score: log base 10 scaled
		downloadScore := int(float64(subtitle.DownloadCount) / 100.0)
		if downloadScore > 25 {
			downloadScore = 25
		}
		score += downloadScore
	}

	// Rating scoring
	if subtitle.Rating > 0 && subtitle.Votes > 0 {
		// Weight rating by number of votes
		ratingScore := int(subtitle.Rating * 5) // Scale 0-10 rating to 0-50
		voteWeight := 1.0
		if subtitle.Votes >= 10 {
			voteWeight = 1.2
		}
		if subtitle.Votes >= 50 {
			voteWeight = 1.4
		}
		score += int(float64(ratingScore) * voteWeight)
		if score > 100 {
			score = 100
		}
	}

	// Hearing Impaired preferences
	if profile.PreferHI && subtitle.HearingImpaired {
		score += 15
	} else if !profile.AllowHI && subtitle.HearingImpaired {
		score -= 25
	}

	// Forced subtitle preferences
	if profile.PreferForced && subtitle.ForcedSubtitle {
		score += 10
	} else if !profile.AllowForced && subtitle.ForcedSubtitle {
		score -= 15
	}

	// HD content bonus
	if subtitle.HD {
		score += 10
	}

	// File size correlation (if media size is known)
	if media.FileSize > 0 && subtitle.FileSize > 0 {
		// Expect subtitle to be roughly 0.01% to 0.1% of video size
		ratio := float64(subtitle.FileSize) / float64(media.FileSize)
		if ratio >= 0.0001 && ratio <= 0.001 {
			score += 5
		} else if ratio > 0.001 {
			score -= 5 // Too large
		}
	}

	return clampScore(score)
}

// clampScore ensures scores stay within 0-100 range.
func clampScore(score int) int {
	if score < 0 {
		return 0
	}
	if score > 100 {
		return 100
	}
	return score
}

// contains checks if a slice contains a string (case-insensitive).
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}

// ScoredSubtitle represents a subtitle with its calculated score.
type ScoredSubtitle struct {
	Subtitle Subtitle
	Score    SubtitleScore
}

// ScoreSubtitles sorts a slice of subtitles by their calculated scores.
func ScoreSubtitles(subtitles []Subtitle, media MediaItem, profile Profile) []ScoredSubtitle {
	scored := make([]ScoredSubtitle, len(subtitles))
	for i, sub := range subtitles {
		scored[i] = ScoredSubtitle{
			Subtitle: sub,
			Score:    CalculateScore(sub, media, profile),
		}
	}

	// Sort by total score (descending)
	for i := 0; i < len(scored)-1; i++ {
		for j := i + 1; j < len(scored); j++ {
			if scored[i].Score.Total < scored[j].Score.Total {
				scored[i], scored[j] = scored[j], scored[i]
			}
		}
	}

	return scored
}

// SelectBest returns the highest scoring subtitle that meets minimum requirements.
func SelectBest(subtitles []Subtitle, media MediaItem, profile Profile) (*Subtitle, *SubtitleScore) {
	if len(subtitles) == 0 {
		return nil, nil
	}

	scored := ScoreSubtitles(subtitles, media, profile)

	// Find first subtitle that meets minimum score
	for _, s := range scored {
		if s.Score.Total >= profile.MinScore {
			return &s.Subtitle, &s.Score
		}
	}

	// If none meet minimum, return highest scoring
	return &scored[0].Subtitle, &scored[0].Score
}
