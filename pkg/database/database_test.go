package database

import "testing"

func TestInsertAndList(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := InsertSubtitle(db, "file.srt", "video.mkv", "es", "google", "", false); err != nil {
		t.Fatalf("insert: %v", err)
	}

	recs, err := ListSubtitles(db)
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
	db, err := Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := InsertSubtitle(db, "file.srt", "video.mkv", "es", "google", "", false); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := DeleteSubtitle(db, "file.srt"); err != nil {
		t.Fatalf("delete: %v", err)
	}
	recs, err := ListSubtitles(db)
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0 records, got %d", len(recs))
	}
}

func TestDownloadHistory(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := InsertDownload(db, "file.srt", "video.mkv", "opensubtitles", "en"); err != nil {
		t.Fatalf("insert download: %v", err)
	}
	recs, err := ListDownloads(db)
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}
	if len(recs) != 1 || recs[0].Provider != "opensubtitles" {
		t.Fatalf("unexpected records %+v", recs)
	}
	if err := DeleteDownload(db, "file.srt"); err != nil {
		t.Fatalf("delete download: %v", err)
	}
	recs, err = ListDownloads(db)
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}
	if len(recs) != 0 {
		t.Fatalf("expected 0 records, got %d", len(recs))
	}
}

func TestHistoryByVideo(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	_ = InsertSubtitle(db, "a.srt", "a.mkv", "en", "g", "", false)
	_ = InsertSubtitle(db, "b.srt", "b.mkv", "en", "g", "", false)
	recs, err := ListSubtitlesByVideo(db, "b.mkv")
	if err != nil {
		t.Fatalf("list by video: %v", err)
	}
	if len(recs) != 1 || recs[0].VideoFile != "b.mkv" {
		t.Fatalf("unexpected result %+v", recs)
	}
}

func TestMediaItems(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := InsertMediaItem(db, "video.mkv", "Show", 1, 2); err != nil {
		t.Fatalf("insert media item: %v", err)
	}
	if err := SetMediaReleaseGroup(db, "video.mkv", "GROUP"); err != nil {
		t.Fatalf("set release group: %v", err)
	}
	if err := SetMediaAltTitles(db, "video.mkv", []string{"Alt"}); err != nil {
		t.Fatalf("set alt titles: %v", err)
	}
	if err := SetMediaFieldLocks(db, "video.mkv", "title"); err != nil {
		t.Fatalf("set locks: %v", err)
	}

	items, err := ListMediaItems(db)
	if err != nil {
		t.Fatalf("list media items: %v", err)
	}
	if len(items) != 1 || items[0].ReleaseGroup != "GROUP" {
		t.Fatalf("unexpected items %+v", items)
	}

	if err := DeleteMediaItem(db, "video.mkv"); err != nil {
		t.Fatalf("delete media item: %v", err)
	}
	items, err = ListMediaItems(db)
	if err != nil {
		t.Fatalf("list media items: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected 0 items, got %d", len(items))
	}
}
