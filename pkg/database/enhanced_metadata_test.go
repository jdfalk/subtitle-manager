// file: pkg/database/enhanced_metadata_test.go
// version: 1.0.0
// guid: 2c8f6e9a-3d7b-4a1c-8e5f-6d7a2b4c9e1f

package database

import (
	"testing"
	"time"
)

func TestEnhancedSubtitleMetadata(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Test enhanced subtitle record with new metadata fields
	score := 0.85
	parentID := "parent-123"
	rec := &SubtitleRecord{
		File:             "enhanced.srt",
		VideoFile:        "video.mkv",
		Language:         "en",
		Service:          "opensubtitles",
		Release:          "HDTV.x264",
		Embedded:         false,
		SourceURL:        "https://opensubtitles.org/download/123456",
		ProviderMetadata: `{"quality":"good","uploader":"user123","rating":4.5}`,
		ConfidenceScore:  &score,
		ParentID:         &parentID,
		ModificationType: "sync",
		CreatedAt:        time.Now(),
	}

	if err := store.InsertSubtitle(rec); err != nil {
		t.Fatalf("insert enhanced subtitle: %v", err)
	}

	recs, err := store.ListSubtitles()
	if err != nil {
		t.Fatalf("list subtitles: %v", err)
	}

	if len(recs) != 1 {
		t.Fatalf("expected 1 record, got %d", len(recs))
	}

	retrieved := recs[0]
	if retrieved.SourceURL != "https://opensubtitles.org/download/123456" {
		t.Errorf("expected source URL 'https://opensubtitles.org/download/123456', got '%s'", retrieved.SourceURL)
	}
	if retrieved.ProviderMetadata != `{"quality":"good","uploader":"user123","rating":4.5}` {
		t.Errorf("expected provider metadata, got '%s'", retrieved.ProviderMetadata)
	}
	if retrieved.ConfidenceScore == nil || *retrieved.ConfidenceScore != 0.85 {
		t.Errorf("expected confidence score 0.85, got %v", retrieved.ConfidenceScore)
	}
	if retrieved.ParentID == nil || *retrieved.ParentID != "parent-123" {
		t.Errorf("expected parent ID 'parent-123', got %v", retrieved.ParentID)
	}
	if retrieved.ModificationType != "sync" {
		t.Errorf("expected modification type 'sync', got '%s'", retrieved.ModificationType)
	}
}

func TestEnhancedDownloadMetadata(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Test enhanced download record with search and performance metadata
	matchScore := 0.92
	responseTime := 1500
	rec := &DownloadRecord{
		File:             "download.srt",
		VideoFile:        "movie.mkv",
		Provider:         "subscene",
		Language:         "es",
		SearchQuery:      "The Matrix 1999",
		MatchScore:       &matchScore,
		DownloadAttempts: 2,
		ErrorMessage:     "",
		ResponseTimeMs:   &responseTime,
		CreatedAt:        time.Now(),
	}

	if err := store.InsertDownload(rec); err != nil {
		t.Fatalf("insert enhanced download: %v", err)
	}

	recs, err := store.ListDownloads()
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}

	if len(recs) != 1 {
		t.Fatalf("expected 1 record, got %d", len(recs))
	}

	retrieved := recs[0]
	if retrieved.SearchQuery != "The Matrix 1999" {
		t.Errorf("expected search query 'The Matrix 1999', got '%s'", retrieved.SearchQuery)
	}
	if retrieved.MatchScore == nil || *retrieved.MatchScore != 0.92 {
		t.Errorf("expected match score 0.92, got %v", retrieved.MatchScore)
	}
	if retrieved.DownloadAttempts != 2 {
		t.Errorf("expected download attempts 2, got %d", retrieved.DownloadAttempts)
	}
	if retrieved.ResponseTimeMs == nil || *retrieved.ResponseTimeMs != 1500 {
		t.Errorf("expected response time 1500ms, got %v", retrieved.ResponseTimeMs)
	}
}

func TestSubtitleSourceTracking(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Test subtitle source tracking
	rating := 4.2
	fileSize := 51200
	src := &SubtitleSource{
		SourceHash:    "sha256:abc123def456",
		OriginalURL:   "https://provider.com/subtitle/789",
		Provider:      "opensubtitles",
		Title:         "Movie Subtitle",
		ReleaseInfo:   "BluRay.x264-GROUP",
		FileSize:      &fileSize,
		DownloadCount: 10,
		SuccessCount:  8,
		AvgRating:     &rating,
		LastSeen:      time.Now(),
		Metadata:      `{"language":"en","format":"srt","encoding":"utf-8"}`,
		CreatedAt:     time.Now(),
	}

	if err := store.InsertSubtitleSource(src); err != nil {
		t.Fatalf("insert subtitle source: %v", err)
	}

	// Test retrieval by hash
	retrieved, err := store.GetSubtitleSource("sha256:abc123def456")
	if err != nil {
		t.Fatalf("get subtitle source: %v", err)
	}

	if retrieved.Provider != "opensubtitles" {
		t.Errorf("expected provider 'opensubtitles', got '%s'", retrieved.Provider)
	}
	if retrieved.Title != "Movie Subtitle" {
		t.Errorf("expected title 'Movie Subtitle', got '%s'", retrieved.Title)
	}
	if retrieved.FileSize == nil || *retrieved.FileSize != 51200 {
		t.Errorf("expected file size 51200, got %v", retrieved.FileSize)
	}
	if retrieved.AvgRating == nil || *retrieved.AvgRating != 4.2 {
		t.Errorf("expected avg rating 4.2, got %v", retrieved.AvgRating)
	}

	// Test listing sources by provider
	sources, err := store.ListSubtitleSources("opensubtitles", 10)
	if err != nil {
		t.Fatalf("list subtitle sources: %v", err)
	}

	if len(sources) != 1 {
		t.Fatalf("expected 1 source, got %d", len(sources))
	}

	// Test updating stats
	newRating := 4.5
	if err := store.UpdateSubtitleSourceStats("sha256:abc123def456", 15, 12, &newRating); err != nil {
		t.Fatalf("update subtitle source stats: %v", err)
	}

	updated, err := store.GetSubtitleSource("sha256:abc123def456")
	if err != nil {
		t.Fatalf("get updated subtitle source: %v", err)
	}

	if updated.DownloadCount != 15 {
		t.Errorf("expected download count 15, got %d", updated.DownloadCount)
	}
	if updated.SuccessCount != 12 {
		t.Errorf("expected success count 12, got %d", updated.SuccessCount)
	}
	if updated.AvgRating == nil || *updated.AvgRating != 4.5 {
		t.Errorf("expected avg rating 4.5, got %v", updated.AvgRating)
	}

	// Test deletion
	if err := store.DeleteSubtitleSource("sha256:abc123def456"); err != nil {
		t.Fatalf("delete subtitle source: %v", err)
	}

	_, err = store.GetSubtitleSource("sha256:abc123def456")
	if err == nil {
		t.Error("expected error when getting deleted subtitle source")
	}
}

func TestBackwardCompatibility(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Test that old-style subtitle records still work
	rec := &SubtitleRecord{
		File:      "legacy.srt",
		VideoFile: "legacy.mkv",
		Language:  "en",
		Service:   "legacy",
		Release:   "Legacy.Release",
		Embedded:  false,
		CreatedAt: time.Now(),
		// New fields left as defaults (empty/nil)
	}

	if err := store.InsertSubtitle(rec); err != nil {
		t.Fatalf("insert legacy subtitle: %v", err)
	}

	recs, err := store.ListSubtitles()
	if err != nil {
		t.Fatalf("list subtitles: %v", err)
	}

	if len(recs) != 1 {
		t.Fatalf("expected 1 record, got %d", len(recs))
	}

	retrieved := recs[0]
	if retrieved.File != "legacy.srt" {
		t.Errorf("expected file 'legacy.srt', got '%s'", retrieved.File)
	}
	if retrieved.SourceURL != "" {
		t.Errorf("expected empty source URL for legacy record, got '%s'", retrieved.SourceURL)
	}
	if retrieved.ConfidenceScore != nil {
		t.Errorf("expected nil confidence score for legacy record, got %v", retrieved.ConfidenceScore)
	}

	// Test that old-style download records still work
	downloadRec := &DownloadRecord{
		File:      "legacy_download.srt",
		VideoFile: "legacy_video.mkv",
		Provider:  "legacy_provider",
		Language:  "en",
		CreatedAt: time.Now(),
		// New fields left as defaults
		DownloadAttempts: 1, // Set to default value
	}

	if err := store.InsertDownload(downloadRec); err != nil {
		t.Fatalf("insert legacy download: %v", err)
	}

	downloads, err := store.ListDownloads()
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}

	if len(downloads) != 1 {
		t.Fatalf("expected 1 download record, got %d", len(downloads))
	}

	retrievedDownload := downloads[0]
	if retrievedDownload.File != "legacy_download.srt" {
		t.Errorf("expected file 'legacy_download.srt', got '%s'", retrievedDownload.File)
	}
	if retrievedDownload.SearchQuery != "" {
		t.Errorf("expected empty search query for legacy record, got '%s'", retrievedDownload.SearchQuery)
	}
	if retrievedDownload.MatchScore != nil {
		t.Errorf("expected nil match score for legacy record, got %v", retrievedDownload.MatchScore)
	}
}