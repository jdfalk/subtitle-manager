// file: pkg/monitoring/sync_test.go
// version: 1.0.1
// guid: 6e0da6f6-6e9b-4cf2-9b89-5f4e7e1f0f66

package monitoring

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/database/mocks"
	"github.com/jdfalk/subtitle-manager/pkg/radarr"
	"github.com/jdfalk/subtitle-manager/pkg/sonarr"
)

func TestAddToMonitoringSkipsMissingFile(t *testing.T) {
	store := mocks.NewMockSubtitleStore(t)
	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)

	opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 1}
	item := database.MediaItem{ID: "media-1", Path: "does-not-exist.mkv"}

	err := monitor.addToMonitoring(item, opts)

	if err != nil {
		t.Fatalf("expected no error for missing file, got %v", err)
	}
	store.AssertNotCalled(t, "ListMonitoredItems")
	store.AssertNotCalled(t, "InsertMonitoredItem", mock.Anything)
}

func TestAddToMonitoringInsertsNewItem(t *testing.T) {
	tempFile, err := os.CreateTemp(t.TempDir(), "media-*.mkv")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("close temp file: %v", err)
	}

	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{}, nil)
	store.EXPECT().InsertMonitoredItem(mock.MatchedBy(func(item *database.MonitoredItem) bool {
		if item == nil {
			return false
		}
		if item.MediaID != "media-2" || item.Path != tempFile.Name() || item.MaxRetries != 3 {
			return false
		}
		return item.Status == "pending" && item.RetryCount == 0
	})).Return(nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)
	opts := SyncOptions{Languages: []string{"en", "es"}, MaxRetries: 3}
	item := database.MediaItem{ID: "media-2", Path: tempFile.Name()}

	if err := monitor.addToMonitoring(item, opts); err != nil {
		t.Fatalf("add to monitoring: %v", err)
	}
}

func TestAddToMonitoringUpdatesExistingItem(t *testing.T) {
	tempFile, err := os.CreateTemp(t.TempDir(), "media-*.mkv")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("close temp file: %v", err)
	}

	store := mocks.NewMockSubtitleStore(t)
	languages, err := json.Marshal([]string{"en"})
	if err != nil {
		t.Fatalf("marshal languages: %v", err)
	}
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{
		{
			ID:         "mon-1",
			MediaID:    "media-3",
			Path:       tempFile.Name(),
			Languages:  string(languages),
			MaxRetries: 1,
		},
	}, nil)
	store.EXPECT().UpdateMonitoredItem(mock.MatchedBy(func(item *database.MonitoredItem) bool {
		if item == nil {
			return false
		}
		if item.ID != "mon-1" || item.Path != tempFile.Name() || item.MaxRetries != 2 {
			return false
		}
		var updated []string
		if err := json.Unmarshal([]byte(item.Languages), &updated); err != nil {
			return false
		}
		return containsString(updated, "en") && containsString(updated, "fr")
	})).Return(nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)
	opts := SyncOptions{Languages: []string{"en", "fr"}, MaxRetries: 2}
	item := database.MediaItem{ID: "media-3", Path: tempFile.Name()}

	if err := monitor.addToMonitoring(item, opts); err != nil {
		t.Fatalf("add to monitoring: %v", err)
	}
}

func TestAddToMonitoringLeavesExistingItemAlone(t *testing.T) {
	tempFile, err := os.CreateTemp(t.TempDir(), "media-*.mkv")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("close temp file: %v", err)
	}

	store := mocks.NewMockSubtitleStore(t)
	languages, err := json.Marshal([]string{"en", "fr"})
	if err != nil {
		t.Fatalf("marshal languages: %v", err)
	}
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{
		{
			ID:         "mon-2",
			MediaID:    "media-4",
			Path:       tempFile.Name(),
			Languages:  string(languages),
			MaxRetries: 2,
		},
	}, nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)
	opts := SyncOptions{Languages: []string{"en", "fr"}, MaxRetries: 2}
	item := database.MediaItem{ID: "media-4", Path: tempFile.Name()}

	if err := monitor.addToMonitoring(item, opts); err != nil {
		t.Fatalf("add to monitoring: %v", err)
	}

	store.AssertNotCalled(t, "UpdateMonitoredItem", mock.Anything)
	store.AssertNotCalled(t, "InsertMonitoredItem", mock.Anything)
}

func TestAddToMonitoringUpdatesOnForceRefresh(t *testing.T) {
	// Arrange
	tempFile, err := os.CreateTemp(t.TempDir(), "media-*.mkv")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("close temp file: %v", err)
	}

	store := mocks.NewMockSubtitleStore(t)
	languages, err := json.Marshal([]string{"en"})
	if err != nil {
		t.Fatalf("marshal languages: %v", err)
	}
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{
		{
			ID:         "mon-force",
			MediaID:    "media-force",
			Path:       tempFile.Name(),
			Languages:  string(languages),
			MaxRetries: 2,
		},
	}, nil)
	store.EXPECT().UpdateMonitoredItem(mock.MatchedBy(func(item *database.MonitoredItem) bool {
		return item != nil && item.ID == "mon-force" && item.Path == tempFile.Name() && item.MaxRetries == 2
	})).Return(nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)
	opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 2, ForceRefresh: true}
	item := database.MediaItem{ID: "media-force", Path: tempFile.Name()}

	// Act
	if err := monitor.addToMonitoring(item, opts); err != nil {
		t.Fatalf("add to monitoring: %v", err)
	}
}

func TestAddToMonitoringSkipsInvalidLanguageJSON(t *testing.T) {
	// Arrange
	tempFile, err := os.CreateTemp(t.TempDir(), "media-*.mkv")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("close temp file: %v", err)
	}

	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{
		{
			ID:        "mon-invalid-json",
			MediaID:   "media-invalid-json",
			Path:      tempFile.Name(),
			Languages: "not-json",
		},
	}, nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)
	opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 1}
	item := database.MediaItem{ID: "media-invalid-json", Path: tempFile.Name()}

	// Act
	if err := monitor.addToMonitoring(item, opts); err != nil {
		t.Fatalf("add to monitoring: %v", err)
	}

	// Assert
	store.AssertNotCalled(t, "UpdateMonitoredItem", mock.Anything)
	store.AssertNotCalled(t, "InsertMonitoredItem", mock.Anything)
}

func TestRemoveFromMonitoringDeletesMatch(t *testing.T) {
	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{
		{ID: "mon-3", Path: "video-1.mkv"},
	}, nil)
	store.EXPECT().DeleteMonitoredItem("mon-3").Return(nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)

	if err := monitor.RemoveFromMonitoring("video-1.mkv"); err != nil {
		t.Fatalf("remove from monitoring: %v", err)
	}
}

func TestRemoveFromMonitoringSkipsMissingItem(t *testing.T) {
	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{
		{ID: "mon-4", Path: "video-2.mkv"},
	}, nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)

	if err := monitor.RemoveFromMonitoring("video-3.mkv"); err != nil {
		t.Fatalf("remove from monitoring: %v", err)
	}
	store.AssertNotCalled(t, "DeleteMonitoredItem", mock.Anything)
}

func TestSyncFromSonarrAddsEpisodes(t *testing.T) {
	// Arrange
	tempFile, err := os.CreateTemp(t.TempDir(), "episode-*.mkv")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("close temp file: %v", err)
	}

	sonarrServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/episode" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"series":{"title":"Test Show"},"episodeFile":{"path":"` + tempFile.Name() + `"},"seasonNumber":1,"episodeNumber":2}]`))
	}))
	defer sonarrServer.Close()

	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{}, nil)
	store.EXPECT().InsertMonitoredItem(mock.Anything).Return(nil)

	monitor := NewEpisodeMonitor(time.Minute, sonarr.NewClient(sonarrServer.URL, "api-key"), nil, store, 2, false)
	opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 1}

	// Act
	if err := monitor.SyncFromSonarr(context.Background(), opts); err != nil {
		t.Fatalf("sync from sonarr: %v", err)
	}
}

func TestSyncFromSonarrReturnsError(t *testing.T) {
	// Arrange
	sonarrServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer sonarrServer.Close()

	store := mocks.NewMockSubtitleStore(t)
	monitor := NewEpisodeMonitor(time.Minute, sonarr.NewClient(sonarrServer.URL, "api-key"), nil, store, 2, false)

	// Act
	if err := monitor.SyncFromSonarr(context.Background(), SyncOptions{}); err == nil {
		t.Fatalf("expected error from sonarr sync")
	}
}

func TestSyncFromRadarrAddsMovies(t *testing.T) {
	// Arrange
	tempFile, err := os.CreateTemp(t.TempDir(), "movie-*.mkv")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("close temp file: %v", err)
	}

	radarrServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/movie" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"title":"Test Movie","movieFile":{"path":"` + tempFile.Name() + `"}}]`))
	}))
	defer radarrServer.Close()

	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{}, nil)
	store.EXPECT().InsertMonitoredItem(mock.Anything).Return(nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, radarr.NewClient(radarrServer.URL, "api-key"), store, 2, false)
	opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 1}

	// Act
	if err := monitor.SyncFromRadarr(context.Background(), opts); err != nil {
		t.Fatalf("sync from radarr: %v", err)
	}
}

func TestSyncFromRadarrReturnsError(t *testing.T) {
	// Arrange
	radarrServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer radarrServer.Close()

	store := mocks.NewMockSubtitleStore(t)
	monitor := NewEpisodeMonitor(time.Minute, nil, radarr.NewClient(radarrServer.URL, "api-key"), store, 2, false)

	// Act
	if err := monitor.SyncFromRadarr(context.Background(), SyncOptions{}); err == nil {
		t.Fatalf("expected error from radarr sync")
	}
}

func TestGetMonitoringStatsCountsStatuses(t *testing.T) {
	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMonitoredItems().Return([]database.MonitoredItem{
		{Status: "pending"},
		{Status: "monitoring"},
		{Status: "found"},
		{Status: "failed"},
		{Status: "blacklisted"},
		{Status: "pending"},
	}, nil)

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)

	stats, err := monitor.GetMonitoringStats()
	if err != nil {
		t.Fatalf("get monitoring stats: %v", err)
	}
	if stats.Total != 6 {
		t.Fatalf("expected total 6, got %d", stats.Total)
	}
	if stats.Pending != 2 || stats.Monitoring != 1 || stats.Found != 1 || stats.Failed != 1 || stats.Blacklisted != 1 {
		t.Fatalf("unexpected stats counts: %+v", stats)
	}
}

func TestSyncFromSonarrSkipsWhenNil(t *testing.T) {
	store := mocks.NewMockSubtitleStore(t)
	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)

	if err := monitor.SyncFromSonarr(context.Background(), SyncOptions{}); err != nil {
		t.Fatalf("expected nil error for nil sonarr, got %v", err)
	}
	store.AssertNotCalled(t, "ListMonitoredItems")
}

func TestSyncFromRadarrSkipsWhenNil(t *testing.T) {
	store := mocks.NewMockSubtitleStore(t)
	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)

	if err := monitor.SyncFromRadarr(context.Background(), SyncOptions{}); err != nil {
		t.Fatalf("expected nil error for nil radarr, got %v", err)
	}
	store.AssertNotCalled(t, "ListMonitoredItems")
}

func TestAddToMonitoringReturnsListError(t *testing.T) {
	tempFile, err := os.CreateTemp(t.TempDir(), "media-*.mkv")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("close temp file: %v", err)
	}

	store := mocks.NewMockSubtitleStore(t)
	store.EXPECT().ListMonitoredItems().Return(nil, errors.New("list failed"))

	monitor := NewEpisodeMonitor(time.Minute, nil, nil, store, 2, false)
	opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 1}
	item := database.MediaItem{ID: "media-5", Path: tempFile.Name()}

	if err := monitor.addToMonitoring(item, opts); err == nil {
		t.Fatalf("expected error")
	}
	store.AssertNotCalled(t, "InsertMonitoredItem", mock.Anything)
}
