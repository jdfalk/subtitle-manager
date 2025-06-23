package database

import "testing"

func TestPebbleInsertAndList(t *testing.T) {
	db, err := OpenPebble(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rec := &SubtitleRecord{File: "f.srt", VideoFile: "v.mkv", Language: "es", Service: "test"}
	if err := db.InsertSubtitle(rec); err != nil {
		t.Fatalf("insert: %v", err)
	}

	recs, err := db.ListSubtitles()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 1 {
		t.Fatalf("expected 1 record, got %d", len(recs))
	}
	if recs[0].File != "f.srt" || recs[0].Language != "es" {
		t.Fatalf("unexpected record %+v", recs[0])
	}
}

func TestPebbleDeleteSubtitle(t *testing.T) {
	db, err := OpenPebble(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	rec := &SubtitleRecord{File: "f.srt", VideoFile: "v.mkv", Language: "es", Service: "test"}
	if err := db.InsertSubtitle(rec); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := db.DeleteSubtitle("f.srt"); err != nil {
		t.Fatalf("delete: %v", err)
	}
	recs, err := db.ListSubtitles()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0 records, got %d", len(recs))
	}
}

func TestPebbleDownloadRecords(t *testing.T) {
	db, err := OpenPebble(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rec := &DownloadRecord{File: "f.srt", VideoFile: "v.mkv", Language: "es", Provider: "test"}
	if err := db.InsertDownload(rec); err != nil {
		t.Fatalf("insert download: %v", err)
	}

	recs, err := db.ListDownloads()
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}
	if len(recs) != 1 {
		t.Fatalf("expected 1 record, got %d", len(recs))
	}
	if recs[0].Provider != "test" || recs[0].Language != "es" {
		t.Fatalf("unexpected record %+v", recs[0])
	}

	if err := db.DeleteDownload("f.srt"); err != nil {
		t.Fatalf("delete download: %v", err)
	}
	recs, err = db.ListDownloads()
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0 records, got %d", len(recs))
	}
}

func TestPebbleMediaItemUpdates(t *testing.T) {
	db, err := OpenPebble(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	item := &MediaItem{Path: "video.mkv", Title: "Old"}
	if err := db.InsertMediaItem(item); err != nil {
		t.Fatalf("insert: %v", err)
	}

	if err := db.SetMediaReleaseGroup("video.mkv", "GRP"); err != nil {
		t.Fatalf("set group: %v", err)
	}
	if err := db.SetMediaAltTitles("video.mkv", []string{"Alt"}); err != nil {
		t.Fatalf("set alt: %v", err)
	}
	if err := db.SetMediaFieldLocks("video.mkv", "title"); err != nil {
		t.Fatalf("set locks: %v", err)
	}
	if err := db.SetMediaTitle("video.mkv", "New"); err != nil {
		t.Fatalf("set title: %v", err)
	}

	items, err := db.ListMediaItems()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(items) != 1 || items[0].Title != "New" || items[0].ReleaseGroup != "GRP" {
		t.Fatalf("unexpected items %+v", items)
	}
}
