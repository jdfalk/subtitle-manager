package database

import (
	"testing"
	"time"
)

// getTestStore returns a test store, preferring Pebble when SQLite is not available
func getTestStore(t *testing.T) SubtitleStore {
	// Try SQLite first (in-memory), fall back to Pebble (temp directory)
	store, err := OpenStore(":memory:", "sqlite")
	if err != nil {
		// SQLite not available, use Pebble with temp directory
		store, err = OpenStore(t.TempDir(), "pebble")
		if err != nil {
			t.Fatalf("Failed to open both SQLite and Pebble stores: %v", err)
		}
	}
	return store
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
