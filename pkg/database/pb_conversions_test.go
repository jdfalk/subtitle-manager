// file: pkg/database/pb_conversions_test.go
// version: 1.0.0
// guid: 126b63da-4b50-4d5d-8299-1e4aec6ed6dd

package database

import (
	"reflect"
	"testing"
	"time"
)

func TestSubtitleRecordConversions(t *testing.T) {
	now := time.Now().UTC()
	score := 0.8
	parent := "abc"
	rec := &SubtitleRecord{
		ID:               "1",
		File:             "a.srt",
		VideoFile:        "a.mkv",
		Release:          "rel",
		Language:         "en",
		Service:          "svc",
		Embedded:         true,
		SourceURL:        "url",
		ProviderMetadata: "meta",
		ConfidenceScore:  &score,
		ParentID:         &parent,
		ModificationType: "sync",
		CreatedAt:        now,
	}
	pb := rec.ToProto()
	if pb.Id != rec.ID || pb.File != rec.File || pb.VideoFile != rec.VideoFile {
		t.Fatalf("basic fields mismatch")
	}
	back := SubtitleRecordFromProto(pb)
	if !reflect.DeepEqual(rec, back) {
		t.Errorf("round trip mismatch: %#v != %#v", rec, back)
	}
}

func TestDownloadRecordConversions(t *testing.T) {
	now := time.Now().UTC()
	score := 0.5
	rt := 120
	rec := &DownloadRecord{
		ID:               "2",
		File:             "b.srt",
		VideoFile:        "b.mkv",
		Provider:         "p",
		Language:         "fr",
		SearchQuery:      "query",
		MatchScore:       &score,
		DownloadAttempts: 3,
		ErrorMessage:     "",
		ResponseTimeMs:   &rt,
		CreatedAt:        now,
	}
	pb := rec.ToProto()
	if pb.Id != rec.ID || pb.Provider != rec.Provider {
		t.Fatalf("basic fields mismatch")
	}
	back := DownloadRecordFromProto(pb)
	if !reflect.DeepEqual(rec, back) {
		t.Errorf("round trip mismatch")
	}
}
