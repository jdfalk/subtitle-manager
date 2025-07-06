// file: pkg/database/metadata_helpers_test.go
// version: 1.0.0
// guid: 9e8d7c6b-5a49-3821-9f8e-7d6c5b4a3210

package database

import (
	"strings"
	"testing"
)

func TestProviderMetadata(t *testing.T) {
	metadata := &ProviderMetadata{
		Quality:    "good",
		Uploader:   "user123",
		Rating:     4.5,
		Downloads:  100,
		Format:     "srt",
		Encoding:   "utf-8",
		FileSize:   51200,
		Language:   "en",
		Release:    "BluRay.x264",
		SourceID:   "123456",
		SourceName: "Movie Subtitle",
	}

	// Test JSON serialization
	jsonStr, err := metadata.ToJSON()
	if err != nil {
		t.Fatalf("failed to serialize metadata: %v", err)
	}

	if !strings.Contains(jsonStr, "\"quality\":\"good\"") {
		t.Errorf("JSON should contain quality field")
	}
	if !strings.Contains(jsonStr, "\"rating\":4.5") {
		t.Errorf("JSON should contain rating field")
	}

	// Test JSON deserialization
	var newMetadata ProviderMetadata
	if err := newMetadata.FromJSON(jsonStr); err != nil {
		t.Fatalf("failed to deserialize metadata: %v", err)
	}

	if newMetadata.Quality != "good" {
		t.Errorf("expected quality 'good', got '%s'", newMetadata.Quality)
	}
	if newMetadata.Rating != 4.5 {
		t.Errorf("expected rating 4.5, got %f", newMetadata.Rating)
	}
	if newMetadata.FileSize != 51200 {
		t.Errorf("expected file size 51200, got %d", newMetadata.FileSize)
	}
}

func TestCalculateSubtitleHash(t *testing.T) {
	content := []byte("1\n00:00:01,000 --> 00:00:02,000\nHello World\n")
	hash := CalculateSubtitleHash(content)

	if !strings.HasPrefix(hash, "sha256:") {
		t.Errorf("hash should start with 'sha256:', got '%s'", hash)
	}

	// Same content should produce same hash
	hash2 := CalculateSubtitleHash(content)
	if hash != hash2 {
		t.Errorf("same content should produce same hash")
	}

	// Different content should produce different hash
	content2 := []byte("2\n00:00:02,000 --> 00:00:03,000\nGoodbye World\n")
	hash3 := CalculateSubtitleHash(content2)
	if hash == hash3 {
		t.Errorf("different content should produce different hash")
	}
}

func TestCalculateSubtitleHashFromReader(t *testing.T) {
	content := "1\n00:00:01,000 --> 00:00:02,000\nHello World\n"
	reader := strings.NewReader(content)

	hash, err := CalculateSubtitleHashFromReader(reader)
	if err != nil {
		t.Fatalf("failed to calculate hash from reader: %v", err)
	}

	if !strings.HasPrefix(hash, "sha256:") {
		t.Errorf("hash should start with 'sha256:', got '%s'", hash)
	}

	// Compare with direct hash calculation
	expectedHash := CalculateSubtitleHash([]byte(content))
	if hash != expectedHash {
		t.Errorf("reader hash should match direct hash")
	}
}

func TestCreateSubtitleRecord(t *testing.T) {
	metadata := &ProviderMetadata{
		Quality:  "excellent",
		Uploader: "trusted_user",
		Rating:   4.8,
	}

	rec, err := CreateSubtitleRecord("test.srt", "video.mkv", "en", "opensubtitles", metadata)
	if err != nil {
		t.Fatalf("failed to create subtitle record: %v", err)
	}

	if rec.File != "test.srt" {
		t.Errorf("expected file 'test.srt', got '%s'", rec.File)
	}
	if rec.Language != "en" {
		t.Errorf("expected language 'en', got '%s'", rec.Language)
	}
	if rec.ModificationType != ModificationTypeOriginal {
		t.Errorf("expected modification type '%s', got '%s'", ModificationTypeOriginal, rec.ModificationType)
	}

	// Check that metadata was serialized
	if rec.ProviderMetadata == "" {
		t.Error("expected provider metadata to be set")
	}
	if !strings.Contains(rec.ProviderMetadata, "\"quality\":\"excellent\"") {
		t.Error("provider metadata should contain quality field")
	}
}

func TestCreateDownloadRecord(t *testing.T) {
	rec := CreateDownloadRecord("download.srt", "movie.mkv", "subscene", "es", "The Matrix 1999")

	if rec.File != "download.srt" {
		t.Errorf("expected file 'download.srt', got '%s'", rec.File)
	}
	if rec.SearchQuery != "The Matrix 1999" {
		t.Errorf("expected search query 'The Matrix 1999', got '%s'", rec.SearchQuery)
	}
	if rec.DownloadAttempts != 1 {
		t.Errorf("expected download attempts 1, got %d", rec.DownloadAttempts)
	}
}

func TestCreateSubtitleSource(t *testing.T) {
	metadata := &ProviderMetadata{
		SourceName: "Movie Subtitle",
		Release:    "BluRay.x264",
		FileSize:   51200,
	}

	src, err := CreateSubtitleSource("sha256:abc123", "https://provider.com/123", "opensubtitles", metadata)
	if err != nil {
		t.Fatalf("failed to create subtitle source: %v", err)
	}

	if src.SourceHash != "sha256:abc123" {
		t.Errorf("expected source hash 'sha256:abc123', got '%s'", src.SourceHash)
	}
	if src.Provider != "opensubtitles" {
		t.Errorf("expected provider 'opensubtitles', got '%s'", src.Provider)
	}
	if src.Title != "Movie Subtitle" {
		t.Errorf("expected title 'Movie Subtitle', got '%s'", src.Title)
	}
	if src.FileSize == nil || *src.FileSize != 51200 {
		t.Errorf("expected file size 51200, got %v", src.FileSize)
	}
}

func TestTrackSubtitleRelationship(t *testing.T) {
	parentRec := &SubtitleRecord{
		ID:              "parent-123",
		File:            "parent.srt",
		VideoFile:       "video.mkv",
		Language:        "en",
		Service:         "opensubtitles",
		SourceURL:       "https://example.com/123",
		ConfidenceScore: func() *float64 { score := 0.95; return &score }(),
	}

	childRec := TrackSubtitleRelationship(parentRec, "synced.srt", ModificationTypeSync)

	if childRec.File != "synced.srt" {
		t.Errorf("expected file 'synced.srt', got '%s'", childRec.File)
	}
	if childRec.ParentID == nil || *childRec.ParentID != "parent-123" {
		t.Errorf("expected parent ID 'parent-123', got %v", childRec.ParentID)
	}
	if childRec.ModificationType != ModificationTypeSync {
		t.Errorf("expected modification type '%s', got '%s'", ModificationTypeSync, childRec.ModificationType)
	}
	if childRec.SourceURL != parentRec.SourceURL {
		t.Errorf("expected source URL to be inherited from parent")
	}
	if childRec.ConfidenceScore == nil || *childRec.ConfidenceScore != 0.95 {
		t.Errorf("expected confidence score to be inherited from parent")
	}
}

func TestValidateConfidenceScore(t *testing.T) {
	// Valid scores
	validScores := []*float64{
		nil, // nil is valid
		func() *float64 { score := 0.0; return &score }(), // 0.0 is valid
		func() *float64 { score := 0.5; return &score }(), // 0.5 is valid
		func() *float64 { score := 1.0; return &score }(), // 1.0 is valid
	}

	for i, score := range validScores {
		if err := ValidateConfidenceScore(score); err != nil {
			t.Errorf("valid score %d should not return error: %v", i, err)
		}
	}

	// Invalid scores
	invalidScores := []*float64{
		func() *float64 { score := -0.1; return &score }(), // negative
		func() *float64 { score := 1.1; return &score }(),  // greater than 1
		func() *float64 { score := 2.0; return &score }(),  // much greater than 1
	}

	for i, score := range invalidScores {
		if err := ValidateConfidenceScore(score); err == nil {
			t.Errorf("invalid score %d should return error", i)
		}
	}
}

func TestValidateMatchScore(t *testing.T) {
	// Valid scores
	validScores := []*float64{
		nil, // nil is valid
		func() *float64 { score := 0.0; return &score }(),  // 0.0 is valid
		func() *float64 { score := 0.85; return &score }(), // 0.85 is valid
		func() *float64 { score := 1.0; return &score }(),  // 1.0 is valid
	}

	for i, score := range validScores {
		if err := ValidateMatchScore(score); err != nil {
			t.Errorf("valid score %d should not return error: %v", i, err)
		}
	}

	// Invalid scores
	invalidScores := []*float64{
		func() *float64 { score := -0.5; return &score }(), // negative
		func() *float64 { score := 1.5; return &score }(),  // greater than 1
	}

	for i, score := range invalidScores {
		if err := ValidateMatchScore(score); err == nil {
			t.Errorf("invalid score %d should return error", i)
		}
	}
}

func TestGetProviderPerformanceStats(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Create test subtitle sources
	rating1 := 4.0
	rating2 := 4.5
	fileSize := 51200

	sources := []*SubtitleSource{
		{
			SourceHash:    "hash1",
			OriginalURL:   "url1",
			Provider:      "test_provider",
			DownloadCount: 10,
			SuccessCount:  8,
			AvgRating:     &rating1,
			FileSize:      &fileSize,
		},
		{
			SourceHash:    "hash2",
			OriginalURL:   "url2",
			Provider:      "test_provider",
			DownloadCount: 5,
			SuccessCount:  4,
			AvgRating:     &rating2,
		},
	}

	for _, src := range sources {
		if err := store.InsertSubtitleSource(src); err != nil {
			t.Fatalf("failed to insert subtitle source: %v", err)
		}
	}

	stats, err := GetProviderPerformanceStats(store, "test_provider")
	if err != nil {
		t.Fatalf("failed to get provider stats: %v", err)
	}

	if stats.Provider != "test_provider" {
		t.Errorf("expected provider 'test_provider', got '%s'", stats.Provider)
	}
	if stats.TotalSources != 2 {
		t.Errorf("expected 2 sources, got %d", stats.TotalSources)
	}
	if stats.TotalDownloads != 15 {
		t.Errorf("expected 15 total downloads, got %d", stats.TotalDownloads)
	}
	if stats.TotalSuccesses != 12 {
		t.Errorf("expected 12 total successes, got %d", stats.TotalSuccesses)
	}
	if stats.SuccessRate != 0.8 {
		t.Errorf("expected success rate 0.8, got %f", stats.SuccessRate)
	}
	if stats.AvgRating == nil || *stats.AvgRating != 4.25 {
		t.Errorf("expected avg rating 4.25, got %v", stats.AvgRating)
	}
}

func TestGetSubtitleHistoryHierarchy(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Insert base subtitle
	base := &SubtitleRecord{
		File:      "base.srt",
		VideoFile: "video.mkv",
		Language:  "en",
		Service:   "test",
	}
	if err := store.InsertSubtitle(base); err != nil {
		t.Fatalf("insert base subtitle: %v", err)
	}

	recs, err := store.ListSubtitles()
	if err != nil {
		t.Fatalf("list subtitles: %v", err)
	}
	if len(recs) == 0 {
		t.Fatal("no subtitles returned")
	}
	baseStored := recs[0]

	// Create a synced subtitle linked to the base
	syncRec := TrackSubtitleRelationship(&baseStored, "sync.srt", ModificationTypeSync)
	if err := store.InsertSubtitle(syncRec); err != nil {
		t.Fatalf("insert sync subtitle: %v", err)
	}

	syncStored, err := store.ListSubtitlesByVideo("video.mkv")
	if err != nil {
		t.Fatalf("list subtitles by video: %v", err)
	}
	// syncStored[0] will be newest due to DESC ordering, so baseStored at end
	// Determine ID of the sync subtitle
	var storedSync SubtitleRecord
	for _, r := range syncStored {
		if r.File == "sync.srt" {
			storedSync = r
			break
		}
	}

	// Add a translated subtitle linked to the synced subtitle
	transRec := TrackSubtitleRelationship(&storedSync, "trans.srt", ModificationTypeTranslate)
	if err := store.InsertSubtitle(transRec); err != nil {
		t.Fatalf("insert translated subtitle: %v", err)
	}

	history, err := GetSubtitleHistory(store, "video.mkv")
	if err != nil {
		t.Fatalf("GetSubtitleHistory failed: %v", err)
	}

	if len(history) != 3 {
		t.Fatalf("expected 3 records, got %d", len(history))
	}

	if history[0].File != "base.srt" || history[1].File != "sync.srt" || history[2].File != "trans.srt" {
		t.Fatalf("unexpected order: %+v", history)
	}
}
