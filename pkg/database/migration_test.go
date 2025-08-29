// file: pkg/database/migration_test.go
// version: 1.0.0
// guid: 7b4d4003-442f-4395-b208-c375162f6fea

package database

import (
	"testing"

	"github.com/jdfalk/gcommon/sdks/go/v1/database"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TestDatabaseMigration ensures gcommon types use the opaque API correctly.
func TestDatabaseMigration(t *testing.T) {
	rec := &database.SubtitleRecord{}
	rec.SetId("test-id")
	rec.SetFile("test.srt")
	rec.SetCreatedAt(timestamppb.Now())

	if rec.GetId() != "test-id" {
		t.Fatalf("expected id to be set")
	}
	if rec.GetFile() != "test.srt" {
		t.Fatalf("expected file to be set")
	}
	if rec.GetCreatedAt() == nil {
		t.Fatalf("expected created_at to be set")
	}
}
