// file: pkg/scoring/adapter_test.go
// version: 1.0.0
// guid: 456789ab-cdef-0123-4567-89abcdef0123
package scoring

import (
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/providers/opensubtitles"
)

func TestFromOpenSubtitlesResult(t *testing.T) {
	result := opensubtitles.SearchResult{
		ID:   "test-id",
		Type: "subtitle",
		Attributes: struct {
			SubtitleID        string  `json:"subtitle_id"`
			Language          string  `json:"language"`
			DownloadCount     int     `json:"download_count"`
			NewDownloadCount  int     `json:"new_download_count"`
			HearingImpaired   bool    `json:"hearing_impaired"`
			HD                bool    `json:"hd"`
			FPS               float64 `json:"fps"`
			Votes             int     `json:"votes"`
			Ratings           float64 `json:"ratings"`
			FromTrusted       bool    `json:"from_trusted"`
			ForeignPartsOnly  bool    `json:"foreign_parts_only"`
			AutoTranslated    bool    `json:"auto_translated"`
			MachineTranslated bool    `json:"machine_translated"`
			UploadDate        string  `json:"upload_date"`
			Release           string  `json:"release"`
			Comments          string  `json:"comments"`
			LegacySubtitleID  int     `json:"legacy_subtitle_id"`
			Uploader          struct {
				UploaderID int    `json:"uploader_id"`
				Name       string `json:"name"`
				Rank       string `json:"rank"`
			} `json:"uploader"`
			FeatureDetails struct {
				FeatureID   int    `json:"feature_id"`
				FeatureType string `json:"feature_type"`
				Year        int    `json:"year"`
				Title       string `json:"title"`
				MovieName   string `json:"movie_name"`
				ImdbID      int    `json:"imdb_id"`
				TmdbID      int    `json:"tmdb_id"`
			} `json:"feature_details"`
			URL          string `json:"url"`
			RelatedLinks struct {
				Label  string `json:"label"`
				URL    string `json:"url"`
				ImgURL string `json:"img_url"`
			} `json:"related_links"`
			Files []struct {
				FileID   int    `json:"file_id"`
				CDNumber int    `json:"cd_number"`
				FileName string `json:"file_name"`
			} `json:"files"`
		}{
			SubtitleID:        "12345",
			Language:          "en",
			DownloadCount:     1000,
			HearingImpaired:   false,
			HD:                true,
			Votes:             50,
			Ratings:           8.5,
			FromTrusted:       true,
			ForeignPartsOnly:  false,
			AutoTranslated:    false,
			MachineTranslated: false,
			UploadDate:        "2023-06-15T10:30:00.000Z",
			Release:           "Movie.2023.1080p.BluRay.x264-GROUP",
			URL:               "https://example.com/subtitle.srt",
			Files: []struct {
				FileID   int    `json:"file_id"`
				CDNumber int    `json:"cd_number"`
				FileName string `json:"file_name"`
			}{
				{
					FileID:   1,
					CDNumber: 1,
					FileName: "movie.srt",
				},
			},
		},
	}

	subtitle := FromOpenSubtitlesResult(result, "opensubtitles")

	// Check basic mapping
	if subtitle.ProviderName != "opensubtitles" {
		t.Errorf("Expected provider name 'opensubtitles', got '%s'", subtitle.ProviderName)
	}

	if !subtitle.IsTrusted {
		t.Error("Expected IsTrusted to be true")
	}

	if subtitle.Release != "Movie.2023.1080p.BluRay.x264-GROUP" {
		t.Errorf("Expected release 'Movie.2023.1080p.BluRay.x264-GROUP', got '%s'", subtitle.Release)
	}

	if subtitle.Format != "srt" {
		t.Errorf("Expected format 'srt', got '%s'", subtitle.Format)
	}

	if subtitle.DownloadCount != 1000 {
		t.Errorf("Expected download count 1000, got %d", subtitle.DownloadCount)
	}

	if subtitle.Rating != 8.5 {
		t.Errorf("Expected rating 8.5, got %f", subtitle.Rating)
	}

	if subtitle.Votes != 50 {
		t.Errorf("Expected votes 50, got %d", subtitle.Votes)
	}

	if !subtitle.HD {
		t.Error("Expected HD to be true")
	}

	if subtitle.FileName != "movie.srt" {
		t.Errorf("Expected filename 'movie.srt', got '%s'", subtitle.FileName)
	}

	// Check date parsing
	expectedDate := time.Date(2023, 6, 15, 10, 30, 0, 0, time.UTC)
	if !subtitle.UploadDate.Equal(expectedDate) {
		t.Errorf("Expected upload date %v, got %v", expectedDate, subtitle.UploadDate)
	}
}

func TestFromMediaPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected MediaItem
	}{
		{
			name: "full quality BluRay release",
			path: "/movies/The.Movie.2023.1080p.BluRay.x264-GROUP.mkv",
			expected: MediaItem{
				Title:        "The.Movie.2023.1080p.BluRay.x264-GROUP",
				Resolution:   "1080p",
				Source:       "bluray",
				Codec:        "x264",
				ReleaseGroup: "GROUP",
			},
		},
		{
			name: "web-dl release",
			path: "/series/Show.S01E05.720p.WEB-DL.x264-TEAM.mp4",
			expected: MediaItem{
				Title:        "Show.S01E05.720p.WEB-DL.x264-TEAM",
				Season:       1,
				Episode:      5,
				Resolution:   "720p",
				Source:       "web-dl",
				Codec:        "x264",
				ReleaseGroup: "TEAM",
			},
		},
		{
			name: "simple TV episode with 1x05 format",
			path: "/tv/Show.1x05.HDTV.XviD.avi",
			expected: MediaItem{
				Title:   "Show.1x05.HDTV.XviD",
				Season:  1,
				Episode: 5,
				Source:  "hdtv",
				Codec:   "xvid",
			},
		},
		{
			name: "4K release",
			path: "/movies/Movie.2023.2160p.UHD.BluRay.x265-RELEASE.mkv",
			expected: MediaItem{
				Title:        "Movie.2023.2160p.UHD.BluRay.x265-RELEASE",
				Resolution:   "2160p",
				Source:       "bluray",
				Codec:        "x265",
				ReleaseGroup: "RELEASE",
			},
		},
		{
			name: "basic filename without extension",
			path: "simple_movie",
			expected: MediaItem{
				Title: "simple_movie",
				// For simple filenames without dots, the whole name becomes the release group
				ReleaseGroup: "simple_movie",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromMediaPath(tt.path)

			if result.Title != tt.expected.Title {
				t.Errorf("Title: expected '%s', got '%s'", tt.expected.Title, result.Title)
			}

			if result.Season != tt.expected.Season {
				t.Errorf("Season: expected %d, got %d", tt.expected.Season, result.Season)
			}

			if result.Episode != tt.expected.Episode {
				t.Errorf("Episode: expected %d, got %d", tt.expected.Episode, result.Episode)
			}

			if result.Resolution != tt.expected.Resolution {
				t.Errorf("Resolution: expected '%s', got '%s'", tt.expected.Resolution, result.Resolution)
			}

			if result.Source != tt.expected.Source {
				t.Errorf("Source: expected '%s', got '%s'", tt.expected.Source, result.Source)
			}

			if result.Codec != tt.expected.Codec {
				t.Errorf("Codec: expected '%s', got '%s'", tt.expected.Codec, result.Codec)
			}

			if result.ReleaseGroup != tt.expected.ReleaseGroup {
				t.Errorf("ReleaseGroup: expected '%s', got '%s'", tt.expected.ReleaseGroup, result.ReleaseGroup)
			}
		})
	}
}

func TestGetFormatFromURL(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"https://example.com/subtitle.srt", "srt"},
		{"https://example.com/subtitle.ass", "ass"},
		{"https://example.com/subtitle.ssa", "ssa"},
		{"https://example.com/subtitle.vtt", "vtt"},
		{"https://example.com/subtitle.sub", "sub"},
		{"https://example.com/subtitle.idx", "idx"},
		{"https://example.com/download?format=SRT", "srt"},
		{"https://example.com/download", "srt"}, // default
	}

	for _, tt := range tests {
		result := getFormatFromURL(tt.url)
		if result != tt.expected {
			t.Errorf("getFormatFromURL('%s') = '%s', want '%s'", tt.url, result, tt.expected)
		}
	}
}

func TestParseSeasonEpisode(t *testing.T) {
	tests := []struct {
		filename        string
		expectedSeason  int
		expectedEpisode int
	}{
		{"Show.S01E05.720p", 1, 5},
		{"show.s02e10.hdtv", 2, 10},
		{"Series.S1E1.web-dl", 1, 1},
		{"Show.1x05.HDTV", 1, 5},
		{"Series.2x10.BluRay", 2, 10},
		{"Show.10x01.480p", 10, 1},
		{"Movie.2023.1080p", 0, 0}, // no season/episode
		{"Random.File.Name", 0, 0}, // no season/episode
	}

	for _, tt := range tests {
		season, episode := parseSeasonEpisode(tt.filename)
		if season != tt.expectedSeason || episode != tt.expectedEpisode {
			t.Errorf("parseSeasonEpisode('%s') = (%d, %d), want (%d, %d)",
				tt.filename, season, episode, tt.expectedSeason, tt.expectedEpisode)
		}
	}
}

func TestIsDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"01", true},
		{"0", true},
		{"abc", false},
		{"12a", false},
		{"", false},
		{"1.5", false},
	}

	for _, tt := range tests {
		result := isDigits(tt.input)
		if result != tt.expected {
			t.Errorf("isDigits('%s') = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsResolution(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"1080p", true},
		{"720p", true},
		{"2160p", true},
		{"4k", true},
		{"random", false},
		{"", false},
	}

	for _, tt := range tests {
		result := isResolution(tt.input)
		if result != tt.expected {
			t.Errorf("isResolution('%s') = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsSource(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"bluray", true},
		{"web-dl", true},
		{"hdtv", true},
		{"webrip", true},
		{"random", false},
		{"", false},
	}

	for _, tt := range tests {
		result := isSource(tt.input)
		if result != tt.expected {
			t.Errorf("isSource('%s') = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsCodec(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"x264", true},
		{"h264", true},
		{"x265", true},
		{"hevc", true},
		{"xvid", true},
		{"random", false},
		{"", false},
	}

	for _, tt := range tests {
		result := isCodec(tt.input)
		if result != tt.expected {
			t.Errorf("isCodec('%s') = %v, want %v", tt.input, result, tt.expected)
		}
	}
}
