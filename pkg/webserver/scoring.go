// file: pkg/webserver/scoring.go
// version: 1.0.0
// guid: 12345678-abcd-efgh-ijkl-1234567890ab
package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/scoring"
)

// scoringConfigHandler handles GET/POST requests for scoring configuration.
func scoringConfigHandler() http.Handler {
	logger := logging.GetLogger("scoring-config")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			// Return current scoring configuration
			profile := scoring.LoadProfileFromConfig()
			if err := json.NewEncoder(w).Encode(profile); err != nil {
				logger.Errorf("failed to encode profile: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "failed to load configuration"})
			}

		case http.MethodPost:
			// Update scoring configuration
			var profile scoring.Profile
			if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
				logger.Errorf("failed to decode profile: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
				return
			}

			// Validate the profile
			if err := scoring.ValidateProfile(profile); err != nil {
				logger.Errorf("invalid profile: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}

			// Save the profile
			scoring.SaveProfileToConfig(profile)

			logger.Infof("scoring configuration updated")
			json.NewEncoder(w).Encode(map[string]string{"status": "success"})

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		}
	})
}

// scoringTestHandler provides a test endpoint for scoring subtitles.
func scoringTestHandler() http.Handler {
	logger := logging.GetLogger("scoring-test")

	type testRequest struct {
		MediaPath string `json:"mediaPath"`
		Language  string `json:"language"`
	}

	type scoredResult struct {
		Subtitle scoring.Subtitle      `json:"subtitle"`
		Score    scoring.SubtitleScore `json:"score"`
	}

	// LocalMediaItem to avoid import conflicts
	type LocalMediaItem struct {
		Title        string `json:"title"`
		Season       int    `json:"season"`
		Episode      int    `json:"episode"`
		ReleaseGroup string `json:"releaseGroup"`
		Resolution   string `json:"resolution"`
		Source       string `json:"source"`
		Codec        string `json:"codec"`
		FileSize     int64  `json:"fileSize"`
	}

	type testResponse struct {
		MediaItem LocalMediaItem  `json:"mediaItem"`
		Profile   scoring.Profile `json:"profile"`
		Results   []scoredResult  `json:"results"`
		Best      *scoredResult   `json:"best,omitempty"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var req testRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Errorf("failed to decode request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
			return
		}

		if req.MediaPath == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "mediaPath is required"})
			return
		}

		// Create sample subtitles for demonstration
		sampleSubtitles := []scoring.Subtitle{
			{
				ProviderName:    "opensubtitles",
				IsTrusted:       true,
				Release:         req.MediaPath + ".BluRay.x264-GROUP",
				Format:          "srt",
				HearingImpaired: false,
				DownloadCount:   5000,
				Rating:          8.5,
				Votes:           100,
			},
			{
				ProviderName:    "subscene",
				IsTrusted:       false,
				Release:         req.MediaPath + ".WEB-DL.x264-TEAM",
				Format:          "srt",
				HearingImpaired: true,
				DownloadCount:   1000,
				Rating:          7.0,
				Votes:           25,
			},
			{
				ProviderName:      "unknown",
				IsTrusted:         false,
				Release:           req.MediaPath + ".CAM.XviD",
				Format:            "sub",
				HearingImpaired:   true,
				DownloadCount:     10,
				Rating:            3.0,
				Votes:             5,
				MachineTranslated: true,
			},
		}

		// Extract media information
		mediaItem := scoring.FromMediaPath(req.MediaPath)

		// Convert to our response type
		responseMediaItem := LocalMediaItem{
			Title:        mediaItem.Title,
			Season:       mediaItem.Season,
			Episode:      mediaItem.Episode,
			ReleaseGroup: mediaItem.ReleaseGroup,
			Resolution:   mediaItem.Resolution,
			Source:       mediaItem.Source,
			Codec:        mediaItem.Codec,
			FileSize:     mediaItem.FileSize,
		}

		// Load scoring profile
		profile := scoring.LoadProfileFromConfig()

		// Score all subtitles
		var results []scoredResult
		for _, subtitle := range sampleSubtitles {
			score := scoring.CalculateScore(subtitle, mediaItem, profile)
			results = append(results, scoredResult{
				Subtitle: subtitle,
				Score:    score,
			})
		}

		// Find best subtitle
		best, bestScore := scoring.SelectBest(sampleSubtitles, mediaItem, profile)
		var bestResult *scoredResult
		if best != nil && bestScore != nil {
			bestResult = &scoredResult{
				Subtitle: *best,
				Score:    *bestScore,
			}
		}

		response := testResponse{
			MediaItem: responseMediaItem,
			Profile:   profile,
			Results:   results,
			Best:      bestResult,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Errorf("failed to encode response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

// scoringDefaultsHandler returns the default scoring profile.
func scoringDefaultsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		profile := scoring.DefaultProfile()
		json.NewEncoder(w).Encode(profile)
	})
}
