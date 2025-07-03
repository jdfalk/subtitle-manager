package database

import (
	"testing"
	"time"

	"github.com/spf13/viper"
)

// getTestStore returns a test store, preferring Pebble when SQLite is not available
func getTestStore(t *testing.T) SubtitleStore {
	// Always use Pebble for pure Go builds (no SQLite/CGO)
	store, err := OpenStore(t.TempDir(), "pebble")
	if err != nil {
		t.Fatalf("Failed to open Pebble store: %v", err)
	}
	return store
}

// TestGetDatabasePath tests the GetDatabasePath function with different backends.
func TestGetDatabasePath(t *testing.T) {
	// Save original viper settings
	originalBackend := viper.GetString("db_backend")
	originalPath := viper.GetString("db_path")
	originalFilename := viper.GetString("sqlite3_filename")
	defer func() {
		viper.Set("db_backend", originalBackend)
		viper.Set("db_path", originalPath)
		viper.Set("sqlite3_filename", originalFilename)
	}()

	tests := []struct {
		name            string
		backend         string
		dbPath          string
		sqlite3Filename string
		expectedPath    string
	}{
		{
			name:            "SQLite backend with filename",
			backend:         "sqlite",
			dbPath:          "/var/data",
			sqlite3Filename: "subtitles.db",
			expectedPath:    "/var/data/subtitles.db",
		},
		{
			name:            "SQLite backend with empty filename",
			backend:         "sqlite",
			dbPath:          "/tmp",
			sqlite3Filename: "",
			expectedPath:    "/tmp",
		},
		{
			name:         "Pebble backend",
			backend:      "pebble",
			dbPath:       "/var/pebble",
			expectedPath: "/var/pebble",
		},
		{
			name:         "Postgres backend",
			backend:      "postgres",
			dbPath:       "postgresql://user:pass@localhost/db",
			expectedPath: "postgresql://user:pass@localhost/db",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Set("db_backend", tt.backend)
			viper.Set("db_path", tt.dbPath)
			viper.Set("sqlite3_filename", tt.sqlite3Filename)

			result := GetDatabasePath()
			if result != tt.expectedPath {
				t.Errorf("expected path %q, got %q", tt.expectedPath, result)
			}
		})
	}
}

// TestGetDatabaseBackend tests the GetDatabaseBackend function.
func TestGetDatabaseBackend(t *testing.T) {
	// Save original viper setting
	originalBackend := viper.GetString("db_backend")
	defer func() {
		viper.Set("db_backend", originalBackend)
	}()

	tests := []struct {
		name     string
		backend  string
		expected string
	}{
		{
			name:     "SQLite backend",
			backend:  "sqlite",
			expected: "sqlite",
		},
		{
			name:     "Pebble backend",
			backend:  "pebble",
			expected: "pebble",
		},
		{
			name:     "Postgres backend",
			backend:  "postgres",
			expected: "postgres",
		},
		{
			name:     "Empty backend",
			backend:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Set("db_backend", tt.backend)

			result := GetDatabaseBackend()
			if result != tt.expected {
				t.Errorf("expected backend %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestInsertAndList(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	rec := &SubtitleRecord{
		File:      "file.srt",
		VideoFile: "video.mkv",
		Language:  "es",
		Service:   "google",
		Release:   "",
		Embedded:  false,
		CreatedAt: time.Now(),
	}

	if err := store.InsertSubtitle(rec); err != nil {
		t.Fatalf("insert: %v", err)
	}

	recs, err := store.ListSubtitles()
	if err != nil {
		t.Fatalf("list: %v", err)
	}

	if len(recs) != 1 {
		t.Fatalf("expected 1 record, got %d", len(recs))
	}
	r := recs[0]
	if r.File != "file.srt" || r.Language != "es" || r.Service != "google" {
		t.Fatalf("unexpected record %+v", r)
	}
}

func TestDeleteSubtitle(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	rec := &SubtitleRecord{
		File:      "file.srt",
		VideoFile: "video.mkv",
		Language:  "es",
		Service:   "google",
		Release:   "",
		Embedded:  false,
		CreatedAt: time.Now(),
	}

	if err := store.InsertSubtitle(rec); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := store.DeleteSubtitle("file.srt"); err != nil {
		t.Fatalf("delete: %v", err)
	}
	recs, err := store.ListSubtitles()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0 records, got %d", len(recs))
	}
}

func TestDownloadHistory(t *testing.T) {
	store, err := OpenStore(t.TempDir(), "pebble")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	rec := &DownloadRecord{
		File:      "file.srt",
		VideoFile: "video.mkv",
		Provider:  "opensubtitles",
		Language:  "en",
		CreatedAt: time.Now(),
	}
	if err := store.InsertDownload(rec); err != nil {
		t.Fatalf("insert download: %v", err)
	}
	recs, err := store.ListDownloads()
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}
	if len(recs) != 1 || recs[0].Provider != "opensubtitles" {
		t.Fatalf("unexpected records %+v", recs)
	}
	if err := store.DeleteDownload("file.srt"); err != nil {
		t.Fatalf("delete download: %v", err)
	}
	recs, err = store.ListDownloads()
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0 records, got %d", len(recs))
	}
}

func TestHistoryByVideo(t *testing.T) {
	store, err := OpenStore(t.TempDir(), "pebble")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	rec1 := &SubtitleRecord{
		File:      "a.srt",
		VideoFile: "a.mkv",
		Language:  "en",
		Service:   "g",
		CreatedAt: time.Now(),
	}
	rec2 := &SubtitleRecord{
		File:      "b.srt",
		VideoFile: "b.mkv",
		Language:  "en",
		Service:   "g",
		CreatedAt: time.Now(),
	}
	_ = store.InsertSubtitle(rec1)
	_ = store.InsertSubtitle(rec2)

	recs, err := store.ListSubtitlesByVideo("b.mkv")
	if err != nil {
		t.Fatalf("list by video: %v", err)
	}
	if len(recs) != 1 || recs[0].VideoFile != "b.mkv" {
		t.Fatalf("unexpected result %+v", recs)
	}
}

func TestMediaItems(t *testing.T) {
	store, err := OpenStore(t.TempDir(), "pebble")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	rec := &MediaItem{
		Path:      "video.mkv",
		Title:     "Show",
		Season:    1,
		Episode:   2,
		CreatedAt: time.Now(),
	}
	if err := store.InsertMediaItem(rec); err != nil {
		t.Fatalf("insert media item: %v", err)
	}
	if err := store.SetMediaReleaseGroup("video.mkv", "GROUP"); err != nil {
		t.Fatalf("set release group: %v", err)
	}
	if err := store.SetMediaAltTitles("video.mkv", []string{"Alt"}); err != nil {
		t.Fatalf("set alt titles: %v", err)
	}
	if err := store.SetMediaFieldLocks("video.mkv", "title"); err != nil {
		t.Fatalf("set locks: %v", err)
	}

	items, err := store.ListMediaItems()
	if err != nil {
		t.Fatalf("list media items: %v", err)
	}
	if len(items) != 1 || items[0].ReleaseGroup != "GROUP" {
		t.Fatalf("unexpected items %+v", items)
	}

	if err := store.DeleteMediaItem("video.mkv"); err != nil {
		t.Fatalf("delete media item: %v", err)
	}
	items, err = store.ListMediaItems()
	if err != nil {
		t.Fatalf("list media items: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected 0 items, got %d", len(items))
	}
}

// TestCountFunctions tests count operations for different entity types.
func TestCountFunctions(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Test counting subtitles
	initialSubCount, err := store.CountSubtitles()
	if err != nil {
		t.Fatalf("CountSubtitles failed: %v", err)
	}

	// Add a subtitle and check count
	rec := &SubtitleRecord{
		File:      "test.srt",
		VideoFile: "test.mkv",
		Language:  "en",
		Service:   "test",
		CreatedAt: time.Now(),
	}
	if err := store.InsertSubtitle(rec); err != nil {
		t.Fatalf("InsertSubtitle failed: %v", err)
	}

	subCount, err := store.CountSubtitles()
	if err != nil {
		t.Fatalf("CountSubtitles failed: %v", err)
	}
	if subCount != initialSubCount+1 {
		t.Errorf("expected subtitle count %d, got %d", initialSubCount+1, subCount)
	}

	// Test counting downloads
	initialDownloadCount, err := store.CountDownloads()
	if err != nil {
		t.Fatalf("CountDownloads failed: %v", err)
	}

	// Add a download and check count
	downloadRec := &DownloadRecord{
		File:      "download.srt",
		VideoFile: "download.mkv",
		Provider:  "test-provider",
		Language:  "es",
		CreatedAt: time.Now(),
	}
	if err := store.InsertDownload(downloadRec); err != nil {
		t.Fatalf("InsertDownload failed: %v", err)
	}

	downloadCount, err := store.CountDownloads()
	if err != nil {
		t.Fatalf("CountDownloads failed: %v", err)
	}
	if downloadCount != initialDownloadCount+1 {
		t.Errorf("expected download count %d, got %d", initialDownloadCount+1, downloadCount)
	}

	// Test counting media items
	initialMediaCount, err := store.CountMediaItems()
	if err != nil {
		t.Fatalf("CountMediaItems failed: %v", err)
	}

	// Add a media item and check count
	mediaItem := &MediaItem{
		Path:      "media.mkv",
		Title:     "Test Movie",
		Season:    1,
		Episode:   1,
		CreatedAt: time.Now(),
	}
	if err := store.InsertMediaItem(mediaItem); err != nil {
		t.Fatalf("InsertMediaItem failed: %v", err)
	}

	mediaCount, err := store.CountMediaItems()
	if err != nil {
		t.Fatalf("CountMediaItems failed: %v", err)
	}
	if mediaCount != initialMediaCount+1 {
		t.Errorf("expected media count %d, got %d", initialMediaCount+1, mediaCount)
	}
}

// TestListSubtitlesByVideo tests filtering subtitles by video file.
func TestListSubtitlesByVideo(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Insert subtitles for different videos
	subtitles := []*SubtitleRecord{
		{
			File:      "video1_en.srt",
			VideoFile: "video1.mkv",
			Language:  "en",
			Service:   "test",
			CreatedAt: time.Now(),
		},
		{
			File:      "video1_es.srt",
			VideoFile: "video1.mkv",
			Language:  "es",
			Service:   "test",
			CreatedAt: time.Now(),
		},
		{
			File:      "video2_en.srt",
			VideoFile: "video2.mkv",
			Language:  "en",
			Service:   "test",
			CreatedAt: time.Now(),
		},
	}

	for _, sub := range subtitles {
		if err := store.InsertSubtitle(sub); err != nil {
			t.Fatalf("InsertSubtitle failed: %v", err)
		}
	}

	// Test filtering by video file
	video1Subs, err := store.ListSubtitlesByVideo("video1.mkv")
	if err != nil {
		t.Fatalf("ListSubtitlesByVideo failed: %v", err)
	}
	if len(video1Subs) != 2 {
		t.Errorf("expected 2 subtitles for video1.mkv, got %d", len(video1Subs))
	}

	video2Subs, err := store.ListSubtitlesByVideo("video2.mkv")
	if err != nil {
		t.Fatalf("ListSubtitlesByVideo failed: %v", err)
	}
	if len(video2Subs) != 1 {
		t.Errorf("expected 1 subtitle for video2.mkv, got %d", len(video2Subs))
	}

	// Test with non-existent video
	nonExistentSubs, err := store.ListSubtitlesByVideo("nonexistent.mkv")
	if err != nil {
		t.Fatalf("ListSubtitlesByVideo failed: %v", err)
	}
	if len(nonExistentSubs) != 0 {
		t.Errorf("expected 0 subtitles for nonexistent video, got %d", len(nonExistentSubs))
	}
}

// TestListDownloadsByVideo tests filtering downloads by video file.
func TestListDownloadsByVideo(t *testing.T) {
	store := getTestStore(t)
	defer store.Close()

	// Insert downloads for different videos
	downloads := []*DownloadRecord{
		{
			File:      "video1_en.srt",
			VideoFile: "video1.mkv",
			Provider:  "provider1",
			Language:  "en",
			CreatedAt: time.Now(),
		},
		{
			File:      "video1_es.srt",
			VideoFile: "video1.mkv",
			Provider:  "provider2",
			Language:  "es",
			CreatedAt: time.Now(),
		},
		{
			File:      "video2_en.srt",
			VideoFile: "video2.mkv",
			Provider:  "provider1",
			Language:  "en",
			CreatedAt: time.Now(),
		},
	}

	for _, download := range downloads {
		if err := store.InsertDownload(download); err != nil {
			t.Fatalf("InsertDownload failed: %v", err)
		}
	}

	// Test filtering by video file
	video1Downloads, err := store.ListDownloadsByVideo("video1.mkv")
	if err != nil {
		t.Fatalf("ListDownloadsByVideo failed: %v", err)
	}
	if len(video1Downloads) != 2 {
		t.Errorf("expected 2 downloads for video1.mkv, got %d", len(video1Downloads))
	}

	video2Downloads, err := store.ListDownloadsByVideo("video2.mkv")
	if err != nil {
		t.Fatalf("ListDownloadsByVideo failed: %v", err)
	}
	if len(video2Downloads) != 1 {
		t.Errorf("expected 1 download for video2.mkv, got %d", len(video2Downloads))
	}
}
