package database

import "testing"

func TestMigrateToPebble(t *testing.T) {
	sqlitePath := t.TempDir() + "/test.db"
	pebblePath := t.TempDir()

	s, err := OpenSQLStore(sqlitePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.InsertSubtitle(&SubtitleRecord{File: "a.srt", VideoFile: "a.mkv", Language: "en", Service: "g"}); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := s.InsertDownload(&DownloadRecord{File: "b.srt", VideoFile: "b.mkv", Provider: "p", Language: "en"}); err != nil {
		t.Fatalf("insert download: %v", err)
	}

	// Verify data in SQLite before migration
	dRecs, err := s.ListDownloads()
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}
	if len(dRecs) != 1 || dRecs[0].File != "b.srt" {
		t.Fatalf("unexpected download records %+v", dRecs)
	}

	s.Close()

	if err := MigrateToPebble(sqlitePath, pebblePath); err != nil {
		t.Fatalf("migrate: %v", err)
	}

	p, err := OpenPebble(pebblePath)
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()

	recs, err := p.ListSubtitles()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(recs) != 1 || recs[0].File != "a.srt" {
		t.Fatalf("unexpected records %+v", recs)
	}
	dRecs, err = p.ListDownloads()
	if err != nil {
		t.Fatalf("list downloads: %v", err)
	}
	if len(dRecs) != 1 || dRecs[0].File != "b.srt" {
		t.Fatalf("unexpected download records %+v", dRecs)
	}
}
