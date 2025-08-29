// file: pkg/database/integration_test.go
// version: 1.0.0
// guid: 1eaa50aa-d59a-4477-897f-3f73ff4fdf41

package database

import "testing"

// TestDatabaseIntegration verifies round-trip conversion with gcommon types.
func TestDatabaseIntegration(t *testing.T) {
	rec := &SubtitleRecord{ID: "it1", File: "file.srt"}
	pb := rec.ToProto()
	back := SubtitleRecordFromProto(pb)
	if back.ID != rec.ID || back.File != rec.File {
		t.Fatalf("round trip mismatch")
	}
}
