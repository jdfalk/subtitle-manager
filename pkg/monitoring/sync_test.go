// file: pkg/monitoring/sync_test.go
// version: 1.0.1
// guid: bebb472c-0502-46ec-8a5d-58ff63261068

package monitoring

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

func createTempMediaFile(t *testing.T) string {
	t.Helper()

	tempDir := t.TempDir()
	mediaPath := filepath.Join(tempDir, "video.mkv")
	if err := os.WriteFile(mediaPath, []byte("data"), 0600); err != nil {
		t.Fatalf("write media file: %v", err)
	}

	return mediaPath
}

func sameStringSet(expected, actual []string) bool {
	if len(expected) != len(actual) {
		return false
	}

	counts := make(map[string]int, len(expected))
	for _, value := range expected {
		counts[value]++
	}

	for _, value := range actual {
		count, ok := counts[value]
		if !ok || count == 0 {
			return false
		}
		counts[value]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func TestAddToMonitoring(t *testing.T) {
	t.Run("adds new item", func(t *testing.T) {
		mediaPath := createTempMediaFile(t)
		store := &MockSubtitleStore{}
		monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}
		opts := SyncOptions{Languages: []string{"en", "es"}, MaxRetries: 3}

		store.On("ListMonitoredItems").Return([]database.MonitoredItem{}, nil)
		store.On("InsertMonitoredItem", mock.MatchedBy(func(item *database.MonitoredItem) bool {
			if item.Path != mediaPath || item.MediaID != "media-1" || item.MaxRetries != opts.MaxRetries {
				return false
			}
			if item.Status != "pending" || item.RetryCount != 0 {
				return false
			}
			var languages []string
			if err := json.Unmarshal([]byte(item.Languages), &languages); err != nil {
				return false
			}
			return sameStringSet(opts.Languages, languages)
		})).Return(nil)

		err := monitor.addToMonitoring(database.MediaItem{ID: "media-1", Path: mediaPath}, opts)

		assert.NoError(t, err)
		store.AssertExpectations(t)
	})

	t.Run("updates existing item", func(t *testing.T) {
		mediaPath := createTempMediaFile(t)
		store := &MockSubtitleStore{}
		monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}
		opts := SyncOptions{Languages: []string{"en", "es"}, MaxRetries: 2}

		existingLanguages, err := json.Marshal([]string{"en"})
		if err != nil {
			t.Fatalf("marshal languages: %v", err)
		}
		existing := database.MonitoredItem{ID: "item-1", Path: mediaPath, Languages: string(existingLanguages), MaxRetries: 1}

		store.On("ListMonitoredItems").Return([]database.MonitoredItem{existing}, nil)
		store.On("UpdateMonitoredItem", mock.MatchedBy(func(item *database.MonitoredItem) bool {
			var languages []string
			if err := json.Unmarshal([]byte(item.Languages), &languages); err != nil {
				return false
			}
			return item.MaxRetries == opts.MaxRetries && sameStringSet(opts.Languages, languages)
		})).Return(nil)

		err = monitor.addToMonitoring(database.MediaItem{ID: "media-2", Path: mediaPath}, opts)

		assert.NoError(t, err)
		store.AssertExpectations(t)
	})

	t.Run("skips update when unchanged", func(t *testing.T) {
		mediaPath := createTempMediaFile(t)
		store := &MockSubtitleStore{}
		monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}
		opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 2}

		existingLanguages, err := json.Marshal([]string{"en"})
		if err != nil {
			t.Fatalf("marshal languages: %v", err)
		}
		existing := database.MonitoredItem{ID: "item-2", Path: mediaPath, Languages: string(existingLanguages), MaxRetries: 2}

		store.On("ListMonitoredItems").Return([]database.MonitoredItem{existing}, nil)

		err = monitor.addToMonitoring(database.MediaItem{ID: "media-3", Path: mediaPath}, opts)

		assert.NoError(t, err)
		store.AssertNotCalled(t, "UpdateMonitoredItem", mock.Anything)
		store.AssertNotCalled(t, "InsertMonitoredItem", mock.Anything)
		store.AssertExpectations(t)
	})

	t.Run("forces refresh", func(t *testing.T) {
		mediaPath := createTempMediaFile(t)
		store := &MockSubtitleStore{}
		monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}
		opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 2, ForceRefresh: true}

		existingLanguages, err := json.Marshal([]string{"en"})
		if err != nil {
			t.Fatalf("marshal languages: %v", err)
		}
		existing := database.MonitoredItem{ID: "item-3", Path: mediaPath, Languages: string(existingLanguages), MaxRetries: 2}

		store.On("ListMonitoredItems").Return([]database.MonitoredItem{existing}, nil)
		store.On("UpdateMonitoredItem", mock.MatchedBy(func(item *database.MonitoredItem) bool {
			return item.Path == mediaPath && item.MaxRetries == opts.MaxRetries
		})).Return(nil)

		err = monitor.addToMonitoring(database.MediaItem{ID: "media-4", Path: mediaPath}, opts)

		assert.NoError(t, err)
		store.AssertExpectations(t)
	})
}

func TestAddToMonitoringSkipsMissingFile(t *testing.T) {
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}
	opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 1}

	err := monitor.addToMonitoring(database.MediaItem{ID: "media-5", Path: "/missing/file.mkv"}, opts)

	assert.NoError(t, err)
	store.AssertNotCalled(t, "ListMonitoredItems")
	store.AssertExpectations(t)
}

func TestAddToMonitoringSkipsInvalidLanguageJSON(t *testing.T) {
	mediaPath := createTempMediaFile(t)
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}
	opts := SyncOptions{Languages: []string{"en"}, MaxRetries: 2}

	badItem := database.MonitoredItem{ID: "item-4", Path: mediaPath, Languages: "{bad", MaxRetries: 2}

	store.On("ListMonitoredItems").Return([]database.MonitoredItem{badItem}, nil)

	err := monitor.addToMonitoring(database.MediaItem{ID: "media-6", Path: mediaPath}, opts)

	assert.NoError(t, err)
	store.AssertNotCalled(t, "UpdateMonitoredItem", mock.Anything)
	store.AssertNotCalled(t, "InsertMonitoredItem", mock.Anything)
	store.AssertExpectations(t)
}

func TestAddToMonitoringReturnsMarshalError(t *testing.T) {
	mediaPath := createTempMediaFile(t)
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}

	store.On("ListMonitoredItems").Return([]database.MonitoredItem{}, nil)

	bad := SyncOptions{Languages: []string{string([]byte{0xff})}, MaxRetries: 1}

	err := monitor.addToMonitoring(database.MediaItem{ID: "media-7", Path: mediaPath}, bad)

	assert.Error(t, err)
	store.AssertNotCalled(t, "InsertMonitoredItem", mock.Anything)
	store.AssertExpectations(t)
}

func TestRemoveFromMonitoring(t *testing.T) {
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}

	items := []database.MonitoredItem{
		{ID: "item-1", Path: "/path/a.mkv"},
		{ID: "item-2", Path: "/path/b.mkv"},
	}

	store.On("ListMonitoredItems").Return(items, nil)
	store.On("DeleteMonitoredItem", "item-2").Return(nil)

	err := monitor.RemoveFromMonitoring("/path/b.mkv")

	assert.NoError(t, err)
	store.AssertExpectations(t)
}

func TestRemoveFromMonitoringNoMatch(t *testing.T) {
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}

	store.On("ListMonitoredItems").Return([]database.MonitoredItem{{ID: "item-1", Path: "/path/a.mkv"}}, nil)

	err := monitor.RemoveFromMonitoring("/path/unknown.mkv")

	assert.NoError(t, err)
	store.AssertNotCalled(t, "DeleteMonitoredItem", mock.Anything)
	store.AssertExpectations(t)
}

func TestGetMonitoringStats(t *testing.T) {
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logrus.NewEntry(logrus.New())}

	items := []database.MonitoredItem{
		{Status: "pending"},
		{Status: "monitoring"},
		{Status: "found"},
		{Status: "failed"},
		{Status: "blacklisted"},
		{Status: "pending"},
	}

	store.On("ListMonitoredItems").Return(items, nil)

	stats, err := monitor.GetMonitoringStats()

	assert.NoError(t, err)
	assert.Equal(t, 6, stats.Total)
	assert.Equal(t, 2, stats.Pending)
	assert.Equal(t, 1, stats.Monitoring)
	assert.Equal(t, 1, stats.Found)
	assert.Equal(t, 1, stats.Failed)
	assert.Equal(t, 1, stats.Blacklisted)
	store.AssertExpectations(t)
}

func TestContainsString(t *testing.T) {
	assert.True(t, containsString([]string{"en", "es"}, "es"))
	assert.False(t, containsString([]string{"en"}, "fr"))
	assert.False(t, containsString(nil, "en"))
	assert.False(t, containsString([]string{}, "en"))
}
