// file: pkg/maintenance/maintenance_additional_test.go
// version: 1.0.0
// guid: 7f3f65cb-02b3-4af1-a1b7-4d809055a2f3

package maintenance

import (
	"context"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/database/mocks"
	"github.com/jdfalk/subtitle-manager/pkg/metadata"
)

// TestFrequencyToDuration_ParsesNamedAndCustomDurations verifies supported
// frequency labels, custom durations, and fallback behavior.
func TestFrequencyToDuration_ParsesNamedAndCustomDurations(t *testing.T) {
	// Arrange: define input/output pairs covering label, custom, and fallback.
	testCases := []struct {
		name     string
		input    string
		expected time.Duration
	}{
		{
			name:     "hourly",
			input:    "hourly",
			expected: time.Hour,
		},
		{
			name:     "daily",
			input:    "daily",
			expected: 24 * time.Hour,
		},
		{
			name:     "weekly",
			input:    "weekly",
			expected: 7 * 24 * time.Hour,
		},
		{
			name:     "monthly",
			input:    "monthly",
			expected: 30 * 24 * time.Hour,
		},
		{
			name:     "custom",
			input:    "2h",
			expected: 2 * time.Hour,
		},
		{
			name:     "invalid",
			input:    "not-a-duration",
			expected: 24 * time.Hour,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			// Act: parse the frequency into a duration.
			duration := frequencyToDuration(testCase.input)

			// Assert: confirm the conversion matches expected values.
			if duration != testCase.expected {
				t.Fatalf("expected %v, got %v", testCase.expected, duration)
			}
		})
	}
}

// TestParseLocks_TrimsAndSkipsEmpty verifies lock parsing trims and ignores
// empty values.
func TestParseLocks_TrimsAndSkipsEmpty(t *testing.T) {
	// Arrange: define a comma-delimited lock string with padding and blanks.
	locks := " title , ,release_group,,"

	// Act: parse the lock string into a set.
	parsed := parseLocks(locks)

	// Assert: confirm expected keys are present and blanks are skipped.
	if !parsed["title"] {
		t.Fatalf("expected title lock to be set")
	}
	if !parsed["release_group"] {
		t.Fatalf("expected release_group lock to be set")
	}
	if parsed[""] {
		t.Fatalf("expected empty lock to be omitted")
	}
}

// TestRefreshMetadata_UpdatesUnlockedTitles ensures only unlocked items are
// updated using fetched metadata.
func TestRefreshMetadata_UpdatesUnlockedTitles(t *testing.T) {
	// Arrange: create a mock store with multiple items and metadata handlers.
	store := mocks.NewMockSubtitleStore(t)
	items := []database.MediaItem{
		{
			Path:       "movie.mkv",
			Title:      "Old Movie",
			FieldLocks: "",
			Season:     0,
			Episode:    0,
		},
		{
			Path:       "locked.mkv",
			Title:      "Locked",
			FieldLocks: "title",
			Season:     0,
			Episode:    0,
		},
		{
			Path:       "episode.mkv",
			Title:      "Old Episode",
			FieldLocks: "",
			Season:     1,
			Episode:    2,
		},
		{
			Path:       "missing.mkv",
			Title:      "Missing",
			FieldLocks: "",
			Season:     0,
			Episode:    0,
		},
	}
	store.EXPECT().ListMediaItems().Return(items, nil)
	store.EXPECT().SetMediaTitle("movie.mkv", "New Movie").Return(nil)
	store.EXPECT().SetMediaTitle("episode.mkv", "New Episode").Return(nil)

	originalMovie := metadata.FetchMovieMetadataFunc
	originalEpisode := metadata.FetchEpisodeMetadataFunc
	metadata.FetchMovieMetadataFunc = func(ctx context.Context, title string, year int, tmdbKey, omdbKey string) (*metadata.MediaInfo, error) {
		switch title {
		case "Old Movie":
			return &metadata.MediaInfo{Title: "New Movie"}, nil
		case "Locked":
			return &metadata.MediaInfo{Title: "Locked New"}, nil
		case "Missing":
			return nil, nil
		default:
			return &metadata.MediaInfo{Title: title}, nil
		}
	}
	metadata.FetchEpisodeMetadataFunc = func(ctx context.Context, title string, season int, episode int, tmdbKey, omdbKey string) (*metadata.MediaInfo, error) {
		if title == "Old Episode" {
			return &metadata.MediaInfo{Title: "New Episode"}, nil
		}
		return nil, nil
	}
	t.Cleanup(func() {
		metadata.FetchMovieMetadataFunc = originalMovie
		metadata.FetchEpisodeMetadataFunc = originalEpisode
	})

	// Act: refresh metadata for all items.
	err := RefreshMetadata(context.Background(), store, "tmdb", "omdb")

	// Assert: verify no error is returned and updates happened for unlocked items.
	if err != nil {
		t.Fatalf("refresh metadata failed: %v", err)
	}
}

// TestRefreshMetadata_StopsOnContextCancel ensures cancellation stops iteration.
func TestRefreshMetadata_StopsOnContextCancel(t *testing.T) {
	// Arrange: cancel the context before refresh begins.
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMediaItems().Return([]database.MediaItem{{
		Path:       "movie.mkv",
		Title:      "Old Movie",
		FieldLocks: "",
		Season:     0,
		Episode:    0,
	}}, nil)

	originalMovie := metadata.FetchMovieMetadataFunc
	originalEpisode := metadata.FetchEpisodeMetadataFunc
	metadata.FetchMovieMetadataFunc = func(ctx context.Context, title string, year int, tmdbKey, omdbKey string) (*metadata.MediaInfo, error) {
		t.Fatalf("unexpected movie metadata fetch")
		return nil, nil
	}
	metadata.FetchEpisodeMetadataFunc = func(ctx context.Context, title string, season int, episode int, tmdbKey, omdbKey string) (*metadata.MediaInfo, error) {
		t.Fatalf("unexpected episode metadata fetch")
		return nil, nil
	}
	t.Cleanup(func() {
		metadata.FetchMovieMetadataFunc = originalMovie
		metadata.FetchEpisodeMetadataFunc = originalEpisode
	})

	// Act: refresh metadata with the canceled context.
	err := RefreshMetadata(ctx, store, "tmdb", "omdb")

	// Assert: ensure the cancellation error is returned.
	if err == nil {
		t.Fatalf("expected context cancellation error")
	}
	if err != context.Canceled {
		t.Fatalf("expected %v, got %v", context.Canceled, err)
	}
}
