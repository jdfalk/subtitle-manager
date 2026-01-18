// file: pkg/monitoring/blacklist_test.go
// version: 1.0.0
// guid: 8c3b7a26-02bd-4f6c-95a7-62b21aa1c2d2

package monitoring

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

func TestEpisodeMonitorAddToBlacklistUpdatesItem(t *testing.T) {
	// Arrange
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logging.GetLogger("test")}
	itemID := "item-1"
	store.On("ListMonitoredItems").Return([]database.MonitoredItem{{
		ID:     itemID,
		Path:   "/media/show.mkv",
		Status: string(StatusPending),
	}}, nil)
	store.On("UpdateMonitoredItem", mock.MatchedBy(func(item *database.MonitoredItem) bool {
		return item.ID == itemID &&
			item.Status == string(StatusBlacklisted) &&
			!item.UpdatedAt.IsZero()
	})).Return(nil)
	blacklistDuration := time.Hour

	// Act
	err := monitor.AddToBlacklist(itemID, "/media/show.mkv", "en", ReasonManualBlacklist, "manual", &blacklistDuration)

	// Assert
	assert.NoError(t, err)
	store.AssertExpectations(t)
}

func TestEpisodeMonitorAddToBlacklistSkipsUnknownItem(t *testing.T) {
	// Arrange
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logging.GetLogger("test")}
	store.On("ListMonitoredItems").Return([]database.MonitoredItem{{
		ID:     "another",
		Path:   "/media/show.mkv",
		Status: string(StatusPending),
	}}, nil)

	// Act
	err := monitor.AddToBlacklist("missing", "/media/show.mkv", "en", ReasonManualBlacklist, "manual", nil)

	// Assert
	assert.NoError(t, err)
	store.AssertNotCalled(t, "UpdateMonitoredItem", mock.Anything)
	store.AssertExpectations(t)
}

func TestEpisodeMonitorRemoveFromBlacklistResetsStatus(t *testing.T) {
	// Arrange
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logging.GetLogger("test")}
	itemID := "item-2"
	store.On("ListMonitoredItems").Return([]database.MonitoredItem{{
		ID:         itemID,
		Path:       "/media/movie.mkv",
		Status:     string(StatusBlacklisted),
		RetryCount: 4,
	}}, nil)
	store.On("UpdateMonitoredItem", mock.MatchedBy(func(item *database.MonitoredItem) bool {
		return item.ID == itemID &&
			item.Status == string(StatusPending) &&
			item.RetryCount == 0 &&
			!item.UpdatedAt.IsZero()
	})).Return(nil)

	// Act
	err := monitor.RemoveFromBlacklist(itemID)

	// Assert
	assert.NoError(t, err)
	store.AssertExpectations(t)
}

func TestEpisodeMonitorIsBlacklisted(t *testing.T) {
	// Arrange
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logging.GetLogger("test")}
	store.On("ListMonitoredItems").Return([]database.MonitoredItem{{
		ID:     "item-3",
		Status: string(StatusBlacklisted),
	}, {
		ID:     "item-4",
		Status: string(StatusPending),
	}}, nil)

	// Act
	blacklisted := monitor.IsBlacklisted("item-3", "en")
	notBlacklisted := monitor.IsBlacklisted("item-4", "en")

	// Assert
	assert.True(t, blacklisted)
	assert.False(t, notBlacklisted)
	store.AssertExpectations(t)
}

func TestEpisodeMonitorIsBlacklistedReturnsFalseOnError(t *testing.T) {
	// Arrange
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logging.GetLogger("test")}
	store.On("ListMonitoredItems").Return([]database.MonitoredItem{}, errors.New("boom"))

	// Act
	blacklisted := monitor.IsBlacklisted("item-5", "en")

	// Assert
	assert.False(t, blacklisted)
	store.AssertExpectations(t)
}

func TestEpisodeMonitorGetBlacklistedItemsFiltersStatus(t *testing.T) {
	// Arrange
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logging.GetLogger("test")}
	store.On("ListMonitoredItems").Return([]database.MonitoredItem{{
		ID:     "item-6",
		Status: string(StatusBlacklisted),
	}, {
		ID:     "item-7",
		Status: string(StatusPending),
	}}, nil)

	// Act
	items, err := monitor.GetBlacklistedItems()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "item-6", items[0].ID)
	store.AssertExpectations(t)
}

func TestEpisodeMonitorAutoBlacklistOnFailure(t *testing.T) {
	// Arrange
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logging.GetLogger("test")}
	itemID := "item-8"
	store.On("ListMonitoredItems").Return([]database.MonitoredItem{{
		ID:     itemID,
		Path:   "/media/show.mkv",
		Status: string(StatusPending),
	}}, nil)
	store.On("UpdateMonitoredItem", mock.MatchedBy(func(item *database.MonitoredItem) bool {
		return item.ID == itemID && item.Status == string(StatusBlacklisted)
	})).Return(nil)
	item := &MonitoredItem{
		ID:         itemID,
		Path:       "/media/show.mkv",
		RetryCount: 3,
		MaxRetries: 3,
	}

	// Act
	err := monitor.AutoBlacklistOnFailure(item)

	// Assert
	assert.NoError(t, err)
	store.AssertExpectations(t)
}

func TestEpisodeMonitorAutoBlacklistOnFailureNoop(t *testing.T) {
	// Arrange
	store := &MockSubtitleStore{}
	monitor := &EpisodeMonitor{store: store, logger: logging.GetLogger("test")}
	item := &MonitoredItem{
		ID:         "item-9",
		Path:       "/media/show.mkv",
		RetryCount: 1,
		MaxRetries: 3,
	}

	// Act
	err := monitor.AutoBlacklistOnFailure(item)

	// Assert
	assert.NoError(t, err)
	store.AssertNotCalled(t, "ListMonitoredItems")
	store.AssertNotCalled(t, "UpdateMonitoredItem", mock.Anything)
}
