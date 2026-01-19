// file: pkg/webhooks/adapter_test.go
// version: 1.0.0
// guid: 9f0f582a-0bde-4a66-9f89-49d9dcd7e9f4

package webhooks

import (
	"context"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/events"
)

func TestWebhookEventPublisher_PublishEvents_RecordsEventHistory(t *testing.T) {
	// Arrange: create a webhook manager and publisher for recording events.
	manager := NewWebhookManager()
	publisher := NewWebhookEventPublisher(manager)
	ctx := context.Background()

	// Act: publish each supported event type.
	downloaded := events.SubtitleDownloadedData{
		FilePath:     "/media/movie.mkv",
		SubtitlePath: "/media/movie.en.srt",
		Language:     "en",
		Provider:     "opensubtitles",
		Score:        95.5,
		Size:         12345,
		Timestamp:    time.Date(2024, 3, 1, 10, 0, 0, 0, time.UTC),
	}
	publisher.PublishSubtitleDownloaded(ctx, downloaded)

	upgraded := events.SubtitleUpgradedData{
		FilePath:        "/media/series.mkv",
		OldSubtitlePath: "/media/series.en.srt",
		NewSubtitlePath: "/media/series.en-hi.srt",
		Language:        "en",
		OldProvider:     "provider-a",
		NewProvider:     "provider-b",
		OldScore:        88.1,
		NewScore:        99.2,
		Timestamp:       time.Date(2024, 3, 2, 11, 0, 0, 0, time.UTC),
	}
	publisher.PublishSubtitleUpgraded(ctx, upgraded)

	failed := events.SubtitleFailedData{
		FilePath:  "/media/failed.mkv",
		Language:  "fr",
		Provider:  "provider-x",
		Error:     "subtitle not found",
		Timestamp: time.Date(2024, 3, 3, 12, 0, 0, 0, time.UTC),
	}
	publisher.PublishSubtitleFailed(ctx, failed)

	searchFailed := events.SearchFailedData{
		Query:     "media:/media/missing.mkv",
		Language:  "es",
		Provider:  "provider-y",
		Error:     "timeout",
		Timestamp: time.Date(2024, 3, 4, 13, 0, 0, 0, time.UTC),
	}
	publisher.PublishSearchFailed(ctx, searchFailed)

	// Assert: validate each event is recorded with the correct type and data.
	history := manager.GetEventHistory(4)
	if len(history) != 4 {
		t.Fatalf("expected 4 events in history, got %d", len(history))
	}

	assertEvent := func(t *testing.T, event WebhookEvent, expectedType string, validateData func(interface{})) {
		t.Helper()
		if event.Type != expectedType {
			t.Errorf("event type = %q, want %q", event.Type, expectedType)
		}
		if event.Source != "subtitle-manager" {
			t.Errorf("event source = %q, want %q", event.Source, "subtitle-manager")
		}
		if event.ID == "" {
			t.Error("event ID should be set")
		}
		validateData(event.Data)
	}

	assertEvent(t, history[0], EventSubtitleDownloaded, func(data interface{}) {
		downloadedData, ok := data.(SubtitleDownloadedData)
		if !ok {
			t.Fatalf("expected SubtitleDownloadedData, got %T", data)
		}
		if downloadedData.FilePath != downloaded.FilePath {
			t.Errorf("downloaded file path = %q, want %q", downloadedData.FilePath, downloaded.FilePath)
		}
		if downloadedData.SubtitlePath != downloaded.SubtitlePath {
			t.Errorf("downloaded subtitle path = %q, want %q", downloadedData.SubtitlePath, downloaded.SubtitlePath)
		}
		if downloadedData.Language != downloaded.Language {
			t.Errorf("downloaded language = %q, want %q", downloadedData.Language, downloaded.Language)
		}
		if downloadedData.Provider != downloaded.Provider {
			t.Errorf("downloaded provider = %q, want %q", downloadedData.Provider, downloaded.Provider)
		}
		if downloadedData.Score != downloaded.Score {
			t.Errorf("downloaded score = %v, want %v", downloadedData.Score, downloaded.Score)
		}
		if downloadedData.Size != downloaded.Size {
			t.Errorf("downloaded size = %d, want %d", downloadedData.Size, downloaded.Size)
		}
		if !downloadedData.Timestamp.Equal(downloaded.Timestamp) {
			t.Errorf("downloaded timestamp = %v, want %v", downloadedData.Timestamp, downloaded.Timestamp)
		}
	})

	assertEvent(t, history[1], EventSubtitleUpgraded, func(data interface{}) {
		upgradedData, ok := data.(SubtitleUpgradedData)
		if !ok {
			t.Fatalf("expected SubtitleUpgradedData, got %T", data)
		}
		if upgradedData.FilePath != upgraded.FilePath {
			t.Errorf("upgraded file path = %q, want %q", upgradedData.FilePath, upgraded.FilePath)
		}
		if upgradedData.OldSubtitlePath != upgraded.OldSubtitlePath {
			t.Errorf("upgraded old subtitle path = %q, want %q", upgradedData.OldSubtitlePath, upgraded.OldSubtitlePath)
		}
		if upgradedData.NewSubtitlePath != upgraded.NewSubtitlePath {
			t.Errorf("upgraded new subtitle path = %q, want %q", upgradedData.NewSubtitlePath, upgraded.NewSubtitlePath)
		}
		if upgradedData.Language != upgraded.Language {
			t.Errorf("upgraded language = %q, want %q", upgradedData.Language, upgraded.Language)
		}
		if upgradedData.OldProvider != upgraded.OldProvider {
			t.Errorf("upgraded old provider = %q, want %q", upgradedData.OldProvider, upgraded.OldProvider)
		}
		if upgradedData.NewProvider != upgraded.NewProvider {
			t.Errorf("upgraded new provider = %q, want %q", upgradedData.NewProvider, upgraded.NewProvider)
		}
		if upgradedData.OldScore != upgraded.OldScore {
			t.Errorf("upgraded old score = %v, want %v", upgradedData.OldScore, upgraded.OldScore)
		}
		if upgradedData.NewScore != upgraded.NewScore {
			t.Errorf("upgraded new score = %v, want %v", upgradedData.NewScore, upgraded.NewScore)
		}
		if !upgradedData.Timestamp.Equal(upgraded.Timestamp) {
			t.Errorf("upgraded timestamp = %v, want %v", upgradedData.Timestamp, upgraded.Timestamp)
		}
	})

	assertEvent(t, history[2], EventSubtitleFailed, func(data interface{}) {
		failedData, ok := data.(SubtitleFailedData)
		if !ok {
			t.Fatalf("expected SubtitleFailedData, got %T", data)
		}
		if failedData.FilePath != failed.FilePath {
			t.Errorf("failed file path = %q, want %q", failedData.FilePath, failed.FilePath)
		}
		if failedData.Language != failed.Language {
			t.Errorf("failed language = %q, want %q", failedData.Language, failed.Language)
		}
		if failedData.Provider != failed.Provider {
			t.Errorf("failed provider = %q, want %q", failedData.Provider, failed.Provider)
		}
		if failedData.Error != failed.Error {
			t.Errorf("failed error = %q, want %q", failedData.Error, failed.Error)
		}
		if !failedData.Timestamp.Equal(failed.Timestamp) {
			t.Errorf("failed timestamp = %v, want %v", failedData.Timestamp, failed.Timestamp)
		}
	})

	assertEvent(t, history[3], EventSearchFailed, func(data interface{}) {
		searchFailedData, ok := data.(SearchFailedData)
		if !ok {
			t.Fatalf("expected SearchFailedData, got %T", data)
		}
		if searchFailedData.Query != searchFailed.Query {
			t.Errorf("search failed query = %q, want %q", searchFailedData.Query, searchFailed.Query)
		}
		if searchFailedData.Language != searchFailed.Language {
			t.Errorf("search failed language = %q, want %q", searchFailedData.Language, searchFailed.Language)
		}
		if searchFailedData.Provider != searchFailed.Provider {
			t.Errorf("search failed provider = %q, want %q", searchFailedData.Provider, searchFailed.Provider)
		}
		if searchFailedData.Error != searchFailed.Error {
			t.Errorf("search failed error = %q, want %q", searchFailedData.Error, searchFailed.Error)
		}
		if !searchFailedData.Timestamp.Equal(searchFailed.Timestamp) {
			t.Errorf("search failed timestamp = %v, want %v", searchFailedData.Timestamp, searchFailed.Timestamp)
		}
	})
}
