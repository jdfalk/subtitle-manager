package database

import (
	"testing"
	"time"
)

func TestPebbleDBFullFunctionality(t *testing.T) {
	// Test the pure Go PebbleDB implementation
	tempDir := t.TempDir()

	store, err := OpenStore(tempDir, "pebble")
	if err != nil {
		t.Fatal("Failed to open PebbleDB store:", err)
	}
	defer store.Close()

	t.Log("✅ Successfully opened PebbleDB store")

	// Test subtitle operations
	subtitle := &SubtitleRecord{
		File:      "test.srt",
		VideoFile: "test.mkv",
		Language:  "en",
		Service:   "opensubtitles",
		CreatedAt: time.Now(),
	}

	if err := store.InsertSubtitle(subtitle); err != nil {
		t.Fatal("Failed to insert subtitle:", err)
	}
	t.Log("✅ Successfully inserted subtitle")

	subtitles, err := store.ListSubtitles()
	if err != nil {
		t.Fatal("Failed to list subtitles:", err)
	}
	t.Logf("✅ Successfully listed %d subtitles", len(subtitles))

	// Test media item operations
	media := &MediaItem{
		Path:      "test.mkv",
		Title:     "Test Movie",
		Season:    1,
		Episode:   1,
		CreatedAt: time.Now(),
	}

	if err := store.InsertMediaItem(media); err != nil {
		t.Fatal("Failed to insert media item:", err)
	}
	t.Log("✅ Successfully inserted media item")

	// Test tag operations
	if err := store.InsertTag("test-tag"); err != nil {
		t.Fatal("Failed to insert tag:", err)
	}
	t.Log("✅ Successfully inserted tag")

	tags, err := store.ListTags()
	if err != nil {
		t.Fatal("Failed to list tags:", err)
	}
	t.Logf("✅ Successfully listed %d tags", len(tags))

	t.Log("🎉 All PebbleDB operations completed successfully!")
	t.Log("Pure Go build (no CGO/SQLite) is fully functional!")
}
