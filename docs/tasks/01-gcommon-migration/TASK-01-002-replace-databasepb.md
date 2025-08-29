# TASK-01-002: Replace databasepb with gcommon database

<!-- file: docs/tasks/01-gcommon-migration/TASK-01-002-replace-databasepb.md -->
<!-- version: 1.0.0 -->
<!-- guid: c2d3e4f5-g6h7-8901-bcde-f23456789012 -->

## 🎯 Objective

Replace all usage of the local `pkg/databasepb` package with the gcommon database package (`github.com/jdfalk/gcommon/sdks/go/v1/database`).

## 📋 Acceptance Criteria

- [ ] All imports of `github.com/jdfalk/subtitle-manager/pkg/databasepb` are replaced
- [ ] All database-related types use gcommon protobuf types with opaque API
- [ ] All getter/setter methods use the correct opaque API pattern
- [ ] No compilation errors after migration
- [ ] All tests pass with new database types
- [ ] Local `pkg/databasepb` directory can be safely removed

## 🔍 Current State Analysis

### Files Currently Using databasepb

Based on codebase analysis, these files import `databasepb`:

1. `pkg/database/pb_conversions.go`

### Current databasepb Types to Replace

From `pkg/databasepb/databasepb.go`:
- `SubtitleRecord` → `github.com/jdfalk/gcommon/sdks/go/v1/database.SubtitleRecord`
- `DownloadRecord` → `github.com/jdfalk/gcommon/sdks/go/v1/database.DownloadRecord`
- `MediaMetadata` → `github.com/jdfalk/gcommon/sdks/go/v1/database.MediaMetadata`

## 🔧 Implementation Steps

### Step 1: Analyze gcommon database package structure

```bash
# Generate documentation for gcommon database package
gomarkdoc --output docs/gcommon-api/database.md github.com/jdfalk/gcommon/sdks/go/v1/database
```

### Step 2: Create migration mapping

Create a mapping file `docs/tasks/01-gcommon-migration/database-migration-map.md`:

```markdown
| Local Type | gcommon Type | Import Path |
|------------|--------------|-------------|
| databasepb.SubtitleRecord | database.SubtitleRecord | github.com/jdfalk/gcommon/sdks/go/v1/database |
| databasepb.DownloadRecord | database.DownloadRecord | github.com/jdfalk/gcommon/sdks/go/v1/database |
| databasepb.MediaMetadata | database.MediaMetadata | github.com/jdfalk/gcommon/sdks/go/v1/database |
```

### Step 3: Update import statements

For `pkg/database/pb_conversions.go`:

```go
// OLD
import (
    "github.com/jdfalk/subtitle-manager/pkg/databasepb"
)

// NEW
import (
    "github.com/jdfalk/gcommon/sdks/go/v1/database"
    "google.golang.org/protobuf/types/known/timestamppb"
)
```

### Step 4: Update type usage with opaque API

Replace direct field access with setter/getter methods:

```go
// OLD - Direct field access
record := &databasepb.SubtitleRecord{
    Id:        "123",
    File:      "subtitle.srt",
    VideoFile: "movie.mkv",
    Language:  "en",
    CreatedAt: timestamppb.Now(),
}

// NEW - Opaque API with setters
record := &database.SubtitleRecord{}
record.SetId("123")
record.SetFile("subtitle.srt")
record.SetVideoFile("movie.mkv")
record.SetLanguage("en")
record.SetCreatedAt(timestamppb.Now())
```

### Step 5: Update conversion functions

Update `pkg/database/pb_conversions.go`:

```go
// OLD
func ToSubtitleRecord(data map[string]interface{}) *databasepb.SubtitleRecord {
    record := &databasepb.SubtitleRecord{
        Id:        getString(data, "id"),
        File:      getString(data, "file"),
        VideoFile: getString(data, "video_file"),
        // ... other fields
    }
    return record
}

// NEW
func ToSubtitleRecord(data map[string]interface{}) *database.SubtitleRecord {
    record := &database.SubtitleRecord{}
    record.SetId(getString(data, "id"))
    record.SetFile(getString(data, "file"))
    record.SetVideoFile(getString(data, "video_file"))
    // ... other fields with setters
    return record
}
```

### Step 6: Update database store interface

Update `pkg/database/store.go` if it references databasepb types:

```go
// OLD
func (s *Store) CreateSubtitleRecord(record *databasepb.SubtitleRecord) error {
    // implementation
}

// NEW
func (s *Store) CreateSubtitleRecord(record *database.SubtitleRecord) error {
    // implementation with getter methods
    id := record.GetId()
    file := record.GetFile()
    // ... use getters throughout
}
```

### Step 7: Update tests

Update any test files that use databasepb types:

```go
// OLD
func TestSubtitleRecord(t *testing.T) {
    record := &databasepb.SubtitleRecord{
        Id:   "test-id",
        File: "test.srt",
    }
    assert.Equal(t, "test-id", record.Id)
    assert.Equal(t, "test.srt", record.File)
}

// NEW
func TestSubtitleRecord(t *testing.T) {
    record := &database.SubtitleRecord{}
    record.SetId("test-id")
    record.SetFile("test.srt")
    assert.Equal(t, "test-id", record.GetId())
    assert.Equal(t, "test.srt", record.GetFile())
}
```

### Step 8: Build and test

```bash
# Build to check for compilation errors
go build ./...

# Run tests to ensure everything works
go test ./pkg/database/...

# Run full test suite
go test ./...
```

## 📚 Required Documentation

### Coding Instructions Reference

**CRITICAL**: Follow these instructions precisely:

```markdown
From .github/instructions/general-coding.instructions.md:

## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS
**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction of any kind.**

## Version Update Requirements
**When modifying any file with a version header, ALWAYS update the version number:**
- Patch version (x.y.Z): Bug fixes, typos, minor formatting changes
- Minor version (x.Y.z): New features, significant content additions
- Major version (X.y.z): Breaking changes, structural overhauls
```

### File Header Requirements

All modified files must have proper headers:

```go
// file: path/to/file.go
// version: 1.1.0  // INCREMENT THIS
// guid: existing-guid-unchanged
```

## 🧪 Testing Requirements

### Unit Tests

Create tests for database migration:

```go
// file: pkg/database/migration_test.go
func TestDatabaseMigration(t *testing.T) {
    // Test that old database types can be converted to new format
    // Test that all required fields are properly set using opaque API
    // Test that getters return expected values
}
```

### Integration Tests

```go
// file: pkg/database/integration_test.go
func TestDatabaseIntegration(t *testing.T) {
    // Test storing and retrieving records using new types
    // Test that all database operations work with new API
    // Test database migrations from old to new format
}
```

## 🎯 Success Metrics

- [ ] All databasepb imports removed
- [ ] `go build ./...` completes successfully
- [ ] All existing tests pass
- [ ] New tests added with 80%+ coverage
- [ ] No direct field access to protobuf fields (all via getters/setters)
- [ ] Performance benchmarks show no regression

## 🚨 Common Pitfalls

1. **Opaque API Confusion**: Remember to use `SetField()` and `GetField()` methods instead of direct field access
2. **Nil Pointer Issues**: Always check for nil before calling methods on protobuf messages
3. **Type Conversions**: Ensure proper type conversions when moving from local types to gcommon types
4. **Database Schema**: Verify that database schema still matches the new protobuf field names

## 📖 Additional Resources

- [gcommon database documentation](../../gcommon-api/database.md)
- [Protobuf Go Generated Code Guide](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
- [General Coding Instructions](../../../.github/instructions/general-coding.instructions.md)

## 🔄 Related Tasks

- **TASK-01-001**: Replace configpb (may share some database config types)
- **TASK-01-005**: Migrate protobuf message types (will need updated database types)
- **TASK-01-006**: Update import statements (will verify this task's imports)

## 📝 Notes for AI Agent

- This task is independent but should be done after TASK-01-001 if there are shared config types
- Standard command-line operations throughout - no VS Code specific requirements
- Follow the version update rules strictly - increment version numbers in all modified files
- The opaque API is critical - never access protobuf fields directly, always use getters/setters
- Pay special attention to database field mappings to ensure data integrity
- If any step fails, document the error and continue with remaining steps where possible
