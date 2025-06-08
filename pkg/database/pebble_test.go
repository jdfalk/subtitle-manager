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
