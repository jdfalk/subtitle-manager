// file: pkg/monitoring/integration_test.go
// version: 1.0.0  
// guid: 12345678-1234-1234-1234-123456789018

package monitoring

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// TestMonitoringWorkflow tests the complete monitoring workflow
func TestMonitoringWorkflow(t *testing.T) {
	store := &MockSubtitleStore{}
	
	// Test setup
	monitor := NewEpisodeMonitor(
		time.Minute,
		nil, // no sonarr for this test
		nil, // no radarr for this test
		store,
		3,     // max retries
		false, // no quality check
	)

	// Test adding an item to monitoring
	mediaItem := database.MediaItem{
		ID:    "1",
		Path:  "/media/test.mkv",
		Title: "Test Movie",
	}

	opts := SyncOptions{
		Languages:  []string{"en", "es"},
		MaxRetries: 3,
	}

	// Mock store expectations for addToMonitoring
	store.On("ListMonitoredItems").Return([]database.MonitoredItem{}, nil).Once()
	store.On("InsertMonitoredItem", mock.AnythingOfType("*database.MonitoredItem")).Return(nil).Once()

	// Test adding to monitoring
	err := monitor.addToMonitoring(mediaItem, opts)
	assert.NoError(t, err)

	// Test getting monitoring statistics with fresh expectations
	mockItems := []database.MonitoredItem{
		{ID: "1", Status: "pending"},
		{ID: "2", Status: "monitoring"},
		{ID: "3", Status: "found"},
	}
	
	// Clear previous expectations and set new ones
	store.ExpectedCalls = nil
	store.On("ListMonitoredItems").Return(mockItems, nil).Once()
	
	stats, err := monitor.GetMonitoringStats()
	assert.NoError(t, err)
	assert.Equal(t, 3, stats.Total)
	assert.Equal(t, 1, stats.Pending)
	assert.Equal(t, 1, stats.Monitoring)
	assert.Equal(t, 1, stats.Found)

	store.AssertExpectations(t)
}

// TestBlacklistManagement tests blacklist functionality
func TestBlacklistManagement(t *testing.T) {
	store := &MockSubtitleStore{}
	
	monitor := NewEpisodeMonitor(
		time.Minute,
		nil, nil,
		store,
		3, false,
	)

	// Test blacklisting an item
	mockItems := []database.MonitoredItem{
		{ID: "1", Status: "failed", RetryCount: 3, MaxRetries: 3},
	}
	
	// Set expectations for AddToBlacklist call
	store.On("ListMonitoredItems").Return(mockItems, nil).Once()
	store.On("UpdateMonitoredItem", mock.MatchedBy(func(item *database.MonitoredItem) bool {
		return item.ID == "1" && item.Status == "blacklisted"
	})).Return(nil).Once()

	// Test adding to blacklist
	err := monitor.AddToBlacklist("1", "/media/test.mkv", "en", ReasonMaxRetriesExceeded, "Test blacklist", nil)
	assert.NoError(t, err)

	// Test checking if blacklisted - need fresh expectations
	blacklistedItems := []database.MonitoredItem{
		{ID: "1", Status: "blacklisted", RetryCount: 3, MaxRetries: 3},
	}
	
	store.ExpectedCalls = nil // Clear previous expectations
	store.On("ListMonitoredItems").Return(blacklistedItems, nil).Once()
	
	blacklisted := monitor.IsBlacklisted("1", "en")
	assert.True(t, blacklisted)

	store.AssertExpectations(t)
}

// TestScheduledMonitor tests the scheduler integration
func TestScheduledMonitor(t *testing.T) {
	store := &MockSubtitleStore{}
	
	scheduledMonitor := NewScheduledMonitor(
		nil, // no sonarr
		nil, // no radarr
		store,
		3,     // max retries
		false, // no quality check
	)

	assert.NotNil(t, scheduledMonitor)
	assert.NotNil(t, scheduledMonitor.monitor)

	// Test that we can create sync options
	opts := SyncOptions{
		Languages:  []string{"en"},
		MaxRetries: 3,
	}

	// This would normally sync from Sonarr/Radarr but we have nil clients
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	
	err := scheduledMonitor.RunSyncTask(ctx, opts)
	assert.NoError(t, err) // Should not error even with nil clients
}