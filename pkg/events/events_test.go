// file: pkg/events/events_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174009

package events

import (
	"context"
	"testing"
	"time"
)

// TestEventPublisher is a test implementation of EventPublisher.
type TestEventPublisher struct {
	SubtitleDownloadedCalled bool
	SubtitleUpgradedCalled   bool
	SubtitleFailedCalled     bool
	SearchFailedCalled       bool
	LastDownloadData         SubtitleDownloadedData
	LastUpgradeData          SubtitleUpgradedData
	LastFailedData           SubtitleFailedData
	LastSearchFailedData     SearchFailedData
}

func (t *TestEventPublisher) PublishSubtitleDownloaded(ctx context.Context, data SubtitleDownloadedData) {
	t.SubtitleDownloadedCalled = true
	t.LastDownloadData = data
}

func (t *TestEventPublisher) PublishSubtitleUpgraded(ctx context.Context, data SubtitleUpgradedData) {
	t.SubtitleUpgradedCalled = true
	t.LastUpgradeData = data
}

func (t *TestEventPublisher) PublishSubtitleFailed(ctx context.Context, data SubtitleFailedData) {
	t.SubtitleFailedCalled = true
	t.LastFailedData = data
}

func (t *TestEventPublisher) PublishSearchFailed(ctx context.Context, data SearchFailedData) {
	t.SearchFailedCalled = true
	t.LastSearchFailedData = data
}

// TestEventPublishing verifies that events are published through the global publisher.
func TestEventPublishing(t *testing.T) {
	// Set up test publisher
	testPublisher := &TestEventPublisher{}
	SetGlobalPublisher(testPublisher)

	ctx := context.Background()

	// Test subtitle downloaded event
	downloadData := SubtitleDownloadedData{
		FilePath:     "/test/movie.mkv",
		SubtitlePath: "/test/movie.en.srt",
		Language:     "en",
		Provider:     "opensubtitles",
		Score:        0.95,
		Size:         1024,
		Timestamp:    time.Now(),
	}
	PublishSubtitleDownloaded(ctx, downloadData)

	if !testPublisher.SubtitleDownloadedCalled {
		t.Error("Expected subtitle downloaded event to be published")
	}

	if testPublisher.LastDownloadData.FilePath != downloadData.FilePath {
		t.Errorf("Expected file path %s, got %s", downloadData.FilePath, testPublisher.LastDownloadData.FilePath)
	}

	// Test subtitle upgrade event
	upgradeData := SubtitleUpgradedData{
		FilePath:        "/test/movie.mkv",
		NewSubtitlePath: "/test/movie.en.srt",
		Language:        "en",
		NewProvider:     "subscene",
		NewScore:        0.98,
		Timestamp:       time.Now(),
	}
	PublishSubtitleUpgraded(ctx, upgradeData)

	if !testPublisher.SubtitleUpgradedCalled {
		t.Error("Expected subtitle upgraded event to be published")
	}

	// Test subtitle failed event
	failedData := SubtitleFailedData{
		FilePath:  "/test/movie.mkv",
		Language:  "en",
		Provider:  "opensubtitles",
		Error:     "Connection timeout",
		Timestamp: time.Now(),
	}
	PublishSubtitleFailed(ctx, failedData)

	if !testPublisher.SubtitleFailedCalled {
		t.Error("Expected subtitle failed event to be published")
	}

	// Test search failed event
	searchFailedData := SearchFailedData{
		Query:     "movie search query",
		Language:  "en",
		Error:     "No results found",
		Timestamp: time.Now(),
	}
	PublishSearchFailed(ctx, searchFailedData)

	if !testPublisher.SearchFailedCalled {
		t.Error("Expected search failed event to be published")
	}

	// Clean up
	SetGlobalPublisher(nil)
}

// TestNoPublisher verifies that events are handled gracefully when no publisher is set.
func TestNoPublisher(t *testing.T) {
	// Clear global publisher
	SetGlobalPublisher(nil)

	ctx := context.Background()

	// These should not panic
	PublishSubtitleDownloaded(ctx, SubtitleDownloadedData{})
	PublishSubtitleUpgraded(ctx, SubtitleUpgradedData{})
	PublishSubtitleFailed(ctx, SubtitleFailedData{})
	PublishSearchFailed(ctx, SearchFailedData{})

	// If we reach here, the test passed
}
