// file: pkg/monitoring/monitor_test.go
// version: 1.0.0
// guid: 12345678-1234-1234-1234-123456789015

package monitoring

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// MockSubtitleStore implements database.SubtitleStore for testing
type MockSubtitleStore struct {
	mock.Mock
}

// Implement all required methods for database.SubtitleStore interface
func (m *MockSubtitleStore) InsertSubtitle(rec *database.SubtitleRecord) error {
	args := m.Called(rec)
	return args.Error(0)
}

func (m *MockSubtitleStore) ListSubtitles() ([]database.SubtitleRecord, error) {
	args := m.Called()
	return args.Get(0).([]database.SubtitleRecord), args.Error(1)
}

func (m *MockSubtitleStore) ListSubtitlesByVideo(video string) ([]database.SubtitleRecord, error) {
	args := m.Called(video)
	return args.Get(0).([]database.SubtitleRecord), args.Error(1)
}

func (m *MockSubtitleStore) CountSubtitles() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockSubtitleStore) DeleteSubtitle(file string) error {
	args := m.Called(file)
	return args.Error(0)
}

func (m *MockSubtitleStore) InsertDownload(rec *database.DownloadRecord) error {
	args := m.Called(rec)
	return args.Error(0)
}

func (m *MockSubtitleStore) ListDownloads() ([]database.DownloadRecord, error) {
	args := m.Called()
	return args.Get(0).([]database.DownloadRecord), args.Error(1)
}

func (m *MockSubtitleStore) ListDownloadsByVideo(video string) ([]database.DownloadRecord, error) {
	args := m.Called(video)
	return args.Get(0).([]database.DownloadRecord), args.Error(1)
}

func (m *MockSubtitleStore) CountDownloads() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockSubtitleStore) DeleteDownload(file string) error {
	args := m.Called(file)
	return args.Error(0)
}

func (m *MockSubtitleStore) InsertMediaItem(rec *database.MediaItem) error {
	args := m.Called(rec)
	return args.Error(0)
}

func (m *MockSubtitleStore) ListMediaItems() ([]database.MediaItem, error) {
	args := m.Called()
	return args.Get(0).([]database.MediaItem), args.Error(1)
}

func (m *MockSubtitleStore) CountMediaItems() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockSubtitleStore) DeleteMediaItem(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *MockSubtitleStore) InsertTag(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockSubtitleStore) ListTags() ([]database.Tag, error) {
	args := m.Called()
	return args.Get(0).([]database.Tag), args.Error(1)
}

func (m *MockSubtitleStore) DeleteTag(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSubtitleStore) UpdateTag(id int64, name string) error {
	args := m.Called(id, name)
	return args.Error(0)
}

func (m *MockSubtitleStore) AssignTagToUser(userID, tagID int64) error {
	args := m.Called(userID, tagID)
	return args.Error(0)
}

func (m *MockSubtitleStore) RemoveTagFromUser(userID, tagID int64) error {
	args := m.Called(userID, tagID)
	return args.Error(0)
}

func (m *MockSubtitleStore) ListTagsForUser(userID int64) ([]database.Tag, error) {
	args := m.Called(userID)
	return args.Get(0).([]database.Tag), args.Error(1)
}

func (m *MockSubtitleStore) AssignTagToMedia(mediaID, tagID int64) error {
	args := m.Called(mediaID, tagID)
	return args.Error(0)
}

func (m *MockSubtitleStore) RemoveTagFromMedia(mediaID, tagID int64) error {
	args := m.Called(mediaID, tagID)
	return args.Error(0)
}

func (m *MockSubtitleStore) ListTagsForMedia(mediaID int64) ([]database.Tag, error) {
	args := m.Called(mediaID)
	return args.Get(0).([]database.Tag), args.Error(1)
}

func (m *MockSubtitleStore) SetMediaReleaseGroup(path, group string) error {
	args := m.Called(path, group)
	return args.Error(0)
}

func (m *MockSubtitleStore) SetMediaAltTitles(path string, titles []string) error {
	args := m.Called(path, titles)
	return args.Error(0)
}

func (m *MockSubtitleStore) SetMediaFieldLocks(path, locks string) error {
	args := m.Called(path, locks)
	return args.Error(0)
}

func (m *MockSubtitleStore) SetMediaTitle(path, title string) error {
	args := m.Called(path, title)
	return args.Error(0)
}

func (m *MockSubtitleStore) InsertMonitoredItem(rec *database.MonitoredItem) error {
	args := m.Called(rec)
	return args.Error(0)
}

func (m *MockSubtitleStore) ListMonitoredItems() ([]database.MonitoredItem, error) {
	args := m.Called()
	return args.Get(0).([]database.MonitoredItem), args.Error(1)
}

func (m *MockSubtitleStore) UpdateMonitoredItem(rec *database.MonitoredItem) error {
	args := m.Called(rec)
	return args.Error(0)
}

func (m *MockSubtitleStore) DeleteMonitoredItem(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSubtitleStore) GetMonitoredItemsToCheck(interval time.Duration) ([]database.MonitoredItem, error) {
	args := m.Called(interval)
	return args.Get(0).([]database.MonitoredItem), args.Error(1)
}

func (m *MockSubtitleStore) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestEpisodeMonitor_NewEpisodeMonitor(t *testing.T) {
	store := &MockSubtitleStore{}
	
	monitor := NewEpisodeMonitor(
		time.Hour,
		nil, // sonarr client
		nil, // radarr client
		store,
		3,    // maxRetries
		true, // qualityCheck
	)

	assert.NotNil(t, monitor)
	assert.Equal(t, time.Hour, monitor.interval)
	assert.Equal(t, 3, monitor.maxRetries)
	assert.True(t, monitor.qualityCheck)
}

func TestEpisodeMonitor_GetMonitoringStats(t *testing.T) {
	store := &MockSubtitleStore{}
	
	// Mock data
	mockItems := []database.MonitoredItem{
		{ID: "1", Status: "pending"},
		{ID: "2", Status: "monitoring"},
		{ID: "3", Status: "found"},
		{ID: "4", Status: "failed"},
		{ID: "5", Status: "blacklisted"},
	}
	
	store.On("ListMonitoredItems").Return(mockItems, nil)

	monitor := NewEpisodeMonitor(
		time.Hour,
		nil, nil,
		store,
		3, true,
	)

	stats, err := monitor.GetMonitoringStats()
	
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, 5, stats.Total)
	assert.Equal(t, 1, stats.Pending)
	assert.Equal(t, 1, stats.Monitoring)
	assert.Equal(t, 1, stats.Found)
	assert.Equal(t, 1, stats.Failed)
	assert.Equal(t, 1, stats.Blacklisted)
	
	store.AssertExpectations(t)
}

func TestEpisodeMonitor_CheckForSubtitles(t *testing.T) {
	store := &MockSubtitleStore{}
	
	// Mock empty items to check
	store.On("GetMonitoredItemsToCheck", mock.Anything).Return([]database.MonitoredItem{}, nil)

	monitor := NewEpisodeMonitor(
		time.Minute,
		nil, nil,
		store,
		3, true,
	)

	err := monitor.checkForSubtitles(context.Background())
	
	assert.NoError(t, err)
	store.AssertExpectations(t)
}