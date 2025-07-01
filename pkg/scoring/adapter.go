// file: pkg/scoring/adapter.go
// version: 1.0.0
// guid: abcd1234-5678-9def-0123-456789abcdef
package scoring

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/providers/opensubtitles"
)

// FromOpenSubtitlesResult converts an OpenSubtitles search result to a Subtitle for scoring.
func FromOpenSubtitlesResult(result opensubtitles.SearchResult, providerName string) Subtitle {
	subtitle := Subtitle{
		ProviderName:      providerName,
		IsTrusted:         result.Attributes.FromTrusted,
		Release:           result.Attributes.Release,
		Format:            getFormatFromURL(result.Attributes.URL),
		HearingImpaired:   result.Attributes.HearingImpaired,
		ForcedSubtitle:    result.Attributes.ForeignPartsOnly,
		DownloadCount:     result.Attributes.DownloadCount,
		Rating:            result.Attributes.Ratings,
		Votes:             result.Attributes.Votes,
		AutoTranslated:    result.Attributes.AutoTranslated,
		MachineTranslated: result.Attributes.MachineTranslated,
		HD:                result.Attributes.HD,
	}

	// Parse upload date
	if result.Attributes.UploadDate != "" {
		if t, err := time.Parse("2006-01-02T15:04:05.000Z", result.Attributes.UploadDate); err == nil {
			subtitle.UploadDate = t
		} else if t, err := time.Parse("2006-01-02", result.Attributes.UploadDate); err == nil {
			subtitle.UploadDate = t
		}
	}

	// Use first file name if available
	if len(result.Attributes.Files) > 0 {
		subtitle.FileName = result.Attributes.Files[0].FileName
	}

	return subtitle
}

// FromMediaPath extracts media information from a file path for scoring context.
func FromMediaPath(mediaPath string) MediaItem {
	filename := filepath.Base(mediaPath)
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))

	media := MediaItem{
		Title: filename,
	}

	// Parse common release group patterns
	parts := strings.Split(filename, ".")
	for _, part := range parts {
		partLower := strings.ToLower(part)

		// Look for resolution indicators
		if strings.Contains(partLower, "1080p") || strings.Contains(partLower, "1080i") {
			media.Resolution = "1080p"
		} else if strings.Contains(partLower, "720p") {
			media.Resolution = "720p"
		} else if strings.Contains(partLower, "2160p") || strings.Contains(partLower, "4k") {
			media.Resolution = "2160p"
		} else if strings.Contains(partLower, "480p") {
			media.Resolution = "480p"
		}

		// Look for source indicators
		if strings.Contains(partLower, "bluray") || strings.Contains(partLower, "blu-ray") ||
			strings.Contains(partLower, "brrip") || strings.Contains(partLower, "bdrip") {
			media.Source = "bluray"
		} else if strings.Contains(partLower, "web-dl") || strings.Contains(partLower, "webdl") {
			media.Source = "web-dl"
		} else if strings.Contains(partLower, "webrip") {
			media.Source = "webrip"
		} else if strings.Contains(partLower, "hdtv") {
			media.Source = "hdtv"
		} else if strings.Contains(partLower, "dvdrip") {
			media.Source = "dvdrip"
		}

		// Look for codec indicators
		if strings.Contains(partLower, "x264") || strings.Contains(partLower, "h264") {
			media.Codec = "x264"
		} else if strings.Contains(partLower, "x265") || strings.Contains(partLower, "h265") ||
			strings.Contains(partLower, "hevc") {
			media.Codec = "x265"
		} else if strings.Contains(partLower, "xvid") {
			media.Codec = "xvid"
		}
	}

	// Try to extract season/episode from filename
	if season, episode := parseSeasonEpisode(filename); season > 0 {
		media.Season = season
		media.Episode = episode
	}

	// Look for release group (usually last part, often after dash)
	// Release groups are typically at the end and not resolution/source/codec indicators
	if len(parts) > 1 {
		lastPart := parts[len(parts)-1]
		if len(lastPart) > 2 && !isResolution(lastPart) && !isSource(lastPart) && !isCodec(lastPart) {
			media.ReleaseGroup = lastPart
		}
		// Also check for dash-separated release groups
		for _, part := range parts {
			if strings.Contains(part, "-") && len(part) > 3 {
				dashParts := strings.Split(part, "-")
				if len(dashParts) > 1 {
					candidate := dashParts[len(dashParts)-1]
					if len(candidate) > 2 && !isResolution(candidate) && !isSource(candidate) && !isCodec(candidate) {
						media.ReleaseGroup = candidate
						break
					}
				}
			}
		}
	} else if len(parts) == 1 {
		// Single part filename (no dots) - use as release group if it's not obvious metadata
		part := parts[0]
		if len(part) > 2 && !isResolution(part) && !isSource(part) && !isCodec(part) {
			media.ReleaseGroup = part
		}
	}

	return media
}

// getFormatFromURL tries to determine subtitle format from download URL or filename.
func getFormatFromURL(url string) string {
	url = strings.ToLower(url)

	if strings.Contains(url, ".srt") {
		return "srt"
	} else if strings.Contains(url, ".ass") {
		return "ass"
	} else if strings.Contains(url, ".ssa") {
		return "ssa"
	} else if strings.Contains(url, ".vtt") {
		return "vtt"
	} else if strings.Contains(url, ".sub") {
		return "sub"
	} else if strings.Contains(url, ".idx") {
		return "idx"
	}

	return "srt" // Default assumption
}

// isResolution checks if a string part represents a video resolution.
func isResolution(part string) bool {
	partLower := strings.ToLower(part)
	resolutions := []string{"1080p", "1080i", "720p", "2160p", "4k", "480p", "360p"}
	for _, res := range resolutions {
		if strings.Contains(partLower, res) {
			return true
		}
	}
	return false
}

// isSource checks if a string part represents a video source.
func isSource(part string) bool {
	partLower := strings.ToLower(part)
	sources := []string{"bluray", "blu-ray", "web-dl", "webdl", "webrip", "hdtv", "dvdrip", "brrip", "bdrip"}
	for _, src := range sources {
		if strings.Contains(partLower, src) {
			return true
		}
	}
	return false
}

// isCodec checks if a string part represents a video codec.
func isCodec(part string) bool {
	partLower := strings.ToLower(part)
	codecs := []string{"x264", "h264", "x265", "h265", "hevc", "xvid"}
	for _, codec := range codecs {
		if strings.Contains(partLower, codec) {
			return true
		}
	}
	return false
}

// parseSeasonEpisode attempts to extract season and episode numbers from filename.
func parseSeasonEpisode(filename string) (season, episode int) {
	filename = strings.ToLower(filename)

	// Pattern: S01E05, s01e05, etc.
	// Look for 's' followed immediately by digits
	for i := 0; i < len(filename)-1; i++ {
		if filename[i] == 's' && i+1 < len(filename) && filename[i+1] >= '0' && filename[i+1] <= '9' {
			// Extract season number (can be 1-2 digits)
			seasonStart := i + 1
			seasonEnd := seasonStart
			for seasonEnd < len(filename) && filename[seasonEnd] >= '0' && filename[seasonEnd] <= '9' {
				seasonEnd++
			}

			if seasonEnd > seasonStart && seasonEnd < len(filename) && filename[seasonEnd] == 'e' {
				seasonStr := filename[seasonStart:seasonEnd]
				episodeStart := seasonEnd + 1
				episodeEnd := episodeStart
				for episodeEnd < len(filename) && filename[episodeEnd] >= '0' && filename[episodeEnd] <= '9' {
					episodeEnd++
				}

				if episodeEnd > episodeStart {
					episodeStr := filename[episodeStart:episodeEnd]
					season, _ = strconv.Atoi(seasonStr)
					episode, _ = strconv.Atoi(episodeStr)
					if season > 0 {
						return
					}
				}
			}
		}
	}

	// Pattern: 1x05, 01x05, etc.
	for i := 0; i < len(filename)-2; i++ {
		if filename[i] == 'x' && i > 0 && i < len(filename)-2 {
			// Look for digits before and after 'x'
			seasonStart := i - 1
			for seasonStart > 0 && filename[seasonStart-1] >= '0' && filename[seasonStart-1] <= '9' {
				seasonStart--
			}
			episodeEnd := i + 1
			for episodeEnd < len(filename) && filename[episodeEnd] >= '0' && filename[episodeEnd] <= '9' {
				episodeEnd++
			}

			if seasonStart < i && episodeEnd > i+1 {
				seasonStr := filename[seasonStart:i]
				episodeStr := filename[i+1 : episodeEnd]
				season, _ = strconv.Atoi(seasonStr)
				episode, _ = strconv.Atoi(episodeStr)
				if season > 0 {
					return
				}
			}
		}
	}

	return 0, 0
}

// isDigits checks if a string contains only digits.
func isDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(s) > 0
}
