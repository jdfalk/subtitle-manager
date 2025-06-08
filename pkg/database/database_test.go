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
