# file: docs/tasks/05-package-replacements/TASK-05-002-replace-databasepb.md
# version: 1.0.0
# guid: b2c3d4e5-f6a7-8b9c-0d1e-2f3a4b5c6d7e

# TASK-05-002: Replace databasepb with gcommon/database

## Overview

**Objective**: Replace local databasepb types with gcommon database types throughout the subtitle-manager application.

**Phase**: 3 (Package Replacements)
**Priority**: High
**Estimated Effort**: 6-8 hours
**Prerequisites**: TASK-05-001 (configpb replacement) and Phase 2 Core Type Migration completed

## Required Reading

**CRITICAL**: Read these documents before starting:

- `docs/gcommon-api/database.md` - gcommon database type specifications and patterns
- `pkg/databasepb/databasepb.go` - Current databasepb implementation to be replaced
- `pkg/database/pb_conversions.go` - Current protobuf conversion utilities
- `docs/MIGRATION_INVENTORY.md` - Database type usage inventory
- Database schema documentation and current storage implementations

## Problem Statement

The subtitle-manager currently uses a local `databasepb` package for database record types (SubtitleRecord, DownloadRecord, etc.). This package needs to be completely replaced with gcommon database types to:

1. **Eliminate Local Database Types**: Remove custom protobuf database record types
2. **Standardize Data Models**: Use gcommon standard database patterns
3. **Improve Serialization**: Leverage gcommon optimized serialization
4. **Enable Data Interoperability**: Share data models with other gcommon-based services

### Current Database Types

```go
// Current implementation (to be replaced)
import databasepb "github.com/jdfalk/subtitle-manager/pkg/databasepb"

type SubtitleRecord struct {
    ID           string
    VideoFile    string
    SubtitleFile string
    Language     string
    Format       string
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

type DownloadRecord struct {
    ID        string
    URL       string
    Status    string
    CreatedAt time.Time
    FilePath  string
}
```

### Target gcommon Database Types

```go
// New implementation using gcommon
import "github.com/jdfalk/gcommon/sdks/go/v1/database"

// Use gcommon database types with opaque API
subtitleRecord := &database.Record{}
subtitleRecord.SetId("subtitle_123")
subtitleRecord.SetField("video_file", "/path/to/video.mp4")
subtitleRecord.SetField("subtitle_file", "/path/to/subtitle.srt")
```

## Technical Approach

### Type Mapping Strategy

1. **Record Type Mapping**: Map local record types to gcommon database records
2. **Field Migration**: Convert fields to gcommon opaque API patterns
3. **Serialization Update**: Use gcommon serialization mechanisms
4. **Query Integration**: Update database queries to work with gcommon types

### Key Mappings

```go
// Database type mappings
type DatabaseTypeMapping struct {
    // OLD databasepb -> NEW gcommon/database
    SubtitleRecord  -> database.Record (with type="subtitle")
    DownloadRecord  -> database.Record (with type="download")
    UserRecord      -> database.Record (with type="user") // if exists
    SessionRecord   -> database.Record (with type="session") // if exists
}
```

## Implementation Steps

### Step 1: Analyze Current Database Usage

```bash
# Find all databasepb imports and usages
grep -r "databasepb" pkg/ --include="*.go"
grep -r "SubtitleRecord\|DownloadRecord" pkg/ --include="*.go"

# Document current database operations
# Create mapping document: database_type_mapping.md
```

### Step 2: Create Database Type Wrappers

```go
// File: pkg/database/gcommon_types.go
package database

import (
    "fmt"
    "time"

    "github.com/jdfalk/gcommon/sdks/go/v1/database"
)

// SubtitleRecord wraps gcommon database.Record for subtitle-specific operations
type SubtitleRecord struct {
    record *database.Record
}

// NewSubtitleRecord creates a new subtitle record
func NewSubtitleRecord() *SubtitleRecord {
    record := &database.Record{}
    record.SetField("record_type", "subtitle")
    record.SetField("created_at", time.Now())

    return &SubtitleRecord{
        record: record,
    }
}

// Subtitle-specific field accessors
func (s *SubtitleRecord) GetID() string {
    return s.record.GetId()
}

func (s *SubtitleRecord) SetID(id string) {
    s.record.SetId(id)
}

func (s *SubtitleRecord) GetVideoFile() string {
    if videoFile, exists := s.record.GetField("video_file"); exists {
        if fileStr, ok := videoFile.(string); ok {
            return fileStr
        }
    }
    return ""
}

func (s *SubtitleRecord) SetVideoFile(videoFile string) {
    s.record.SetField("video_file", videoFile)
}

func (s *SubtitleRecord) GetSubtitleFile() string {
    if subtitleFile, exists := s.record.GetField("subtitle_file"); exists {
        if fileStr, ok := subtitleFile.(string); ok {
            return fileStr
        }
    }
    return ""
}

func (s *SubtitleRecord) SetSubtitleFile(subtitleFile string) {
    s.record.SetField("subtitle_file", subtitleFile)
}

func (s *SubtitleRecord) GetLanguage() string {
    if language, exists := s.record.GetField("language"); exists {
        if langStr, ok := language.(string); ok {
            return langStr
        }
    }
    return ""
}

func (s *SubtitleRecord) SetLanguage(language string) {
    s.record.SetField("language", language)
}

func (s *SubtitleRecord) GetFormat() string {
    if format, exists := s.record.GetField("format"); exists {
        if formatStr, ok := format.(string); ok {
            return formatStr
        }
    }
    return ""
}

func (s *SubtitleRecord) SetFormat(format string) {
    s.record.SetField("format", format)
}

func (s *SubtitleRecord) GetCreatedAt() time.Time {
    if createdAt, exists := s.record.GetField("created_at"); exists {
        if timeVal, ok := createdAt.(time.Time); ok {
            return timeVal
        }
    }
    return time.Time{}
}

func (s *SubtitleRecord) SetCreatedAt(createdAt time.Time) {
    s.record.SetField("created_at", createdAt)
}

func (s *SubtitleRecord) GetUpdatedAt() time.Time {
    if updatedAt, exists := s.record.GetField("updated_at"); exists {
        if timeVal, ok := updatedAt.(time.Time); ok {
            return timeVal
        }
    }
    return time.Time{}
}

func (s *SubtitleRecord) SetUpdatedAt(updatedAt time.Time) {
    s.record.SetField("updated_at", updatedAt)
}

// Get underlying gcommon record for advanced operations
func (s *SubtitleRecord) GetRecord() *database.Record {
    return s.record
}

// Create from gcommon record
func NewSubtitleRecordFromGcommon(record *database.Record) *SubtitleRecord {
    return &SubtitleRecord{
        record: record,
    }
}

// DownloadRecord wraps gcommon database.Record for download-specific operations
type DownloadRecord struct {
    record *database.Record
}

// NewDownloadRecord creates a new download record
func NewDownloadRecord() *DownloadRecord {
    record := &database.Record{}
    record.SetField("record_type", "download")
    record.SetField("created_at", time.Now())

    return &DownloadRecord{
        record: record,
    }
}

// Download-specific field accessors
func (d *DownloadRecord) GetID() string {
    return d.record.GetId()
}

func (d *DownloadRecord) SetID(id string) {
    d.record.SetId(id)
}

func (d *DownloadRecord) GetURL() string {
    if url, exists := d.record.GetField("url"); exists {
        if urlStr, ok := url.(string); ok {
            return urlStr
        }
    }
    return ""
}

func (d *DownloadRecord) SetURL(url string) {
    d.record.SetField("url", url)
}

func (d *DownloadRecord) GetStatus() string {
    if status, exists := d.record.GetField("status"); exists {
        if statusStr, ok := status.(string); ok {
            return statusStr
        }
    }
    return ""
}

func (d *DownloadRecord) SetStatus(status string) {
    d.record.SetField("status", status)
}

func (d *DownloadRecord) GetFilePath() string {
    if filePath, exists := d.record.GetField("file_path"); exists {
        if pathStr, ok := filePath.(string); ok {
            return pathStr
        }
    }
    return ""
}

func (d *DownloadRecord) SetFilePath(filePath string) {
    d.record.SetField("file_path", filePath)
}

func (d *DownloadRecord) GetCreatedAt() time.Time {
    if createdAt, exists := d.record.GetField("created_at"); exists {
        if timeVal, ok := createdAt.(time.Time); ok {
            return timeVal
        }
    }
    return time.Time{}
}

func (d *DownloadRecord) SetCreatedAt(createdAt time.Time) {
    d.record.SetField("created_at", createdAt)
}

// Get underlying gcommon record for advanced operations
func (d *DownloadRecord) GetRecord() *database.Record {
    return d.record
}

// Create from gcommon record
func NewDownloadRecordFromGcommon(record *database.Record) *DownloadRecord {
    return &DownloadRecord{
        record: record,
    }
}
```

### Step 3: Update Database Operations

```go
// File: pkg/database/operations.go
package database

import (
    "fmt"

    "github.com/jdfalk/gcommon/sdks/go/v1/database"
)

// DatabaseManager handles gcommon database operations
type DatabaseManager struct {
    store database.Store // gcommon database interface
}

// NewDatabaseManager creates a new database manager
func NewDatabaseManager(store database.Store) *DatabaseManager {
    return &DatabaseManager{
        store: store,
    }
}

// Subtitle record operations
func (dm *DatabaseManager) CreateSubtitleRecord(record *SubtitleRecord) error {
    record.SetUpdatedAt(time.Now())
    return dm.store.Save(record.GetRecord())
}

func (dm *DatabaseManager) GetSubtitleRecord(id string) (*SubtitleRecord, error) {
    record, err := dm.store.Get(id)
    if err != nil {
        return nil, fmt.Errorf("failed to get subtitle record: %v", err)
    }

    // Verify this is a subtitle record
    if recordType, exists := record.GetField("record_type"); exists {
        if typeStr, ok := recordType.(string); ok && typeStr == "subtitle" {
            return NewSubtitleRecordFromGcommon(record), nil
        }
    }

    return nil, fmt.Errorf("record %s is not a subtitle record", id)
}

func (dm *DatabaseManager) UpdateSubtitleRecord(record *SubtitleRecord) error {
    record.SetUpdatedAt(time.Now())
    return dm.store.Update(record.GetRecord())
}

func (dm *DatabaseManager) DeleteSubtitleRecord(id string) error {
    return dm.store.Delete(id)
}

func (dm *DatabaseManager) ListSubtitleRecords(limit, offset int) ([]*SubtitleRecord, error) {
    // Create query for subtitle records
    query := &database.Query{}
    query.SetField("record_type", "subtitle")
    query.SetLimit(limit)
    query.SetOffset(offset)

    records, err := dm.store.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to query subtitle records: %v", err)
    }

    var subtitleRecords []*SubtitleRecord
    for _, record := range records {
        subtitleRecords = append(subtitleRecords, NewSubtitleRecordFromGcommon(record))
    }

    return subtitleRecords, nil
}

// Download record operations
func (dm *DatabaseManager) CreateDownloadRecord(record *DownloadRecord) error {
    return dm.store.Save(record.GetRecord())
}

func (dm *DatabaseManager) GetDownloadRecord(id string) (*DownloadRecord, error) {
    record, err := dm.store.Get(id)
    if err != nil {
        return nil, fmt.Errorf("failed to get download record: %v", err)
    }

    // Verify this is a download record
    if recordType, exists := record.GetField("record_type"); exists {
        if typeStr, ok := recordType.(string); ok && typeStr == "download" {
            return NewDownloadRecordFromGcommon(record), nil
        }
    }

    return nil, fmt.Errorf("record %s is not a download record", id)
}

func (dm *DatabaseManager) UpdateDownloadRecord(record *DownloadRecord) error {
    return dm.store.Update(record.GetRecord())
}

func (dm *DatabaseManager) DeleteDownloadRecord(id string) error {
    return dm.store.Delete(id)
}

func (dm *DatabaseManager) ListDownloadRecords(status string, limit, offset int) ([]*DownloadRecord, error) {
    // Create query for download records
    query := &database.Query{}
    query.SetField("record_type", "download")
    if status != "" {
        query.SetField("status", status)
    }
    query.SetLimit(limit)
    query.SetOffset(offset)

    records, err := dm.store.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to query download records: %v", err)
    }

    var downloadRecords []*DownloadRecord
    for _, record := range records {
        downloadRecords = append(downloadRecords, NewDownloadRecordFromGcommon(record))
    }

    return downloadRecords, nil
}

// Advanced query operations
func (dm *DatabaseManager) SearchSubtitlesByLanguage(language string) ([]*SubtitleRecord, error) {
    query := &database.Query{}
    query.SetField("record_type", "subtitle")
    query.SetField("language", language)

    records, err := dm.store.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to search subtitles by language: %v", err)
    }

    var subtitleRecords []*SubtitleRecord
    for _, record := range records {
        subtitleRecords = append(subtitleRecords, NewSubtitleRecordFromGcommon(record))
    }

    return subtitleRecords, nil
}

func (dm *DatabaseManager) GetActiveDownloads() ([]*DownloadRecord, error) {
    query := &database.Query{}
    query.SetField("record_type", "download")
    query.SetField("status", "in_progress")

    records, err := dm.store.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to get active downloads: %v", err)
    }

    var downloadRecords []*DownloadRecord
    for _, record := range records {
        downloadRecords = append(downloadRecords, NewDownloadRecordFromGcommon(record))
    }

    return downloadRecords, nil
}
```

### Step 4: Migrate Serialization/Deserialization

```go
// File: pkg/database/serialization.go
package database

import (
    "encoding/json"
    "fmt"

    "github.com/jdfalk/gcommon/sdks/go/v1/database"
)

// SerializationManager handles data serialization for gcommon types
type SerializationManager struct{}

// NewSerializationManager creates a new serialization manager
func NewSerializationManager() *SerializationManager {
    return &SerializationManager{}
}

// Serialize subtitle record to JSON (for compatibility)
func (sm *SerializationManager) SerializeSubtitleRecord(record *SubtitleRecord) ([]byte, error) {
    // Extract data from gcommon record using opaque API
    data := map[string]interface{}{
        "id":            record.GetID(),
        "video_file":    record.GetVideoFile(),
        "subtitle_file": record.GetSubtitleFile(),
        "language":      record.GetLanguage(),
        "format":        record.GetFormat(),
        "created_at":    record.GetCreatedAt(),
        "updated_at":    record.GetUpdatedAt(),
        "record_type":   "subtitle",
    }

    // Include any custom fields from the gcommon record
    gRecord := record.GetRecord()
    // Note: This assumes gcommon provides a way to iterate over fields
    // Implementation may vary based on actual gcommon API

    return json.Marshal(data)
}

// Deserialize subtitle record from JSON
func (sm *SerializationManager) DeserializeSubtitleRecord(data []byte) (*SubtitleRecord, error) {
    var recordData map[string]interface{}
    if err := json.Unmarshal(data, &recordData); err != nil {
        return nil, fmt.Errorf("failed to unmarshal subtitle record: %v", err)
    }

    record := NewSubtitleRecord()

    // Set basic fields
    if id, exists := recordData["id"]; exists {
        if idStr, ok := id.(string); ok {
            record.SetID(idStr)
        }
    }

    if videoFile, exists := recordData["video_file"]; exists {
        if fileStr, ok := videoFile.(string); ok {
            record.SetVideoFile(fileStr)
        }
    }

    if subtitleFile, exists := recordData["subtitle_file"]; exists {
        if fileStr, ok := subtitleFile.(string); ok {
            record.SetSubtitleFile(fileStr)
        }
    }

    if language, exists := recordData["language"]; exists {
        if langStr, ok := language.(string); ok {
            record.SetLanguage(langStr)
        }
    }

    if format, exists := recordData["format"]; exists {
        if formatStr, ok := format.(string); ok {
            record.SetFormat(formatStr)
        }
    }

    // Parse timestamps
    if createdAt, exists := recordData["created_at"]; exists {
        if timeStr, ok := createdAt.(string); ok {
            if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
                record.SetCreatedAt(t)
            }
        }
    }

    if updatedAt, exists := recordData["updated_at"]; exists {
        if timeStr, ok := updatedAt.(string); ok {
            if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
                record.SetUpdatedAt(t)
            }
        }
    }

    return record, nil
}

// Serialize download record to JSON
func (sm *SerializationManager) SerializeDownloadRecord(record *DownloadRecord) ([]byte, error) {
    data := map[string]interface{}{
        "id":          record.GetID(),
        "url":         record.GetURL(),
        "status":      record.GetStatus(),
        "file_path":   record.GetFilePath(),
        "created_at":  record.GetCreatedAt(),
        "record_type": "download",
    }

    return json.Marshal(data)
}

// Deserialize download record from JSON
func (sm *SerializationManager) DeserializeDownloadRecord(data []byte) (*DownloadRecord, error) {
    var recordData map[string]interface{}
    if err := json.Unmarshal(data, &recordData); err != nil {
        return nil, fmt.Errorf("failed to unmarshal download record: %v", err)
    }

    record := NewDownloadRecord()

    if id, exists := recordData["id"]; exists {
        if idStr, ok := id.(string); ok {
            record.SetID(idStr)
        }
    }

    if url, exists := recordData["url"]; exists {
        if urlStr, ok := url.(string); ok {
            record.SetURL(urlStr)
        }
    }

    if status, exists := recordData["status"]; exists {
        if statusStr, ok := status.(string); ok {
            record.SetStatus(statusStr)
        }
    }

    if filePath, exists := recordData["file_path"]; exists {
        if pathStr, ok := filePath.(string); ok {
            record.SetFilePath(pathStr)
        }
    }

    if createdAt, exists := recordData["created_at"]; exists {
        if timeStr, ok := createdAt.(string); ok {
            if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
                record.SetCreatedAt(t)
            }
        }
    }

    return record, nil
}

// Migration helpers for converting old databasepb types
func (sm *SerializationManager) MigrateFromOldSubtitleRecord(oldRecord interface{}) (*SubtitleRecord, error) {
    // This would handle migration from old databasepb.SubtitleRecord
    // Implementation depends on the actual structure of old records

    newRecord := NewSubtitleRecord()

    // Use reflection or type assertions to extract fields from old record
    // and populate new gcommon record

    return newRecord, nil
}

func (sm *SerializationManager) MigrateFromOldDownloadRecord(oldRecord interface{}) (*DownloadRecord, error) {
    // This would handle migration from old databasepb.DownloadRecord
    // Implementation depends on the actual structure of old records

    newRecord := NewDownloadRecord()

    // Use reflection or type assertions to extract fields from old record
    // and populate new gcommon record

    return newRecord, nil
}
```

### Step 5: Update All Database References

```go
// Example: Update subtitle service
// File: pkg/subtitle/service.go
package subtitle

import (
    "fmt"

    "github.com/jdfalk/subtitle-manager/pkg/database"
)

type SubtitleService struct {
    dbManager *database.DatabaseManager
}

func NewSubtitleService(dbManager *database.DatabaseManager) *SubtitleService {
    return &SubtitleService{
        dbManager: dbManager,
    }
}

func (s *SubtitleService) CreateSubtitle(videoFile, subtitleFile, language, format string) (*database.SubtitleRecord, error) {
    record := database.NewSubtitleRecord()
    record.SetID(generateID()) // Implement ID generation
    record.SetVideoFile(videoFile)
    record.SetSubtitleFile(subtitleFile)
    record.SetLanguage(language)
    record.SetFormat(format)

    if err := s.dbManager.CreateSubtitleRecord(record); err != nil {
        return nil, fmt.Errorf("failed to create subtitle record: %v", err)
    }

    return record, nil
}

func (s *SubtitleService) GetSubtitle(id string) (*database.SubtitleRecord, error) {
    return s.dbManager.GetSubtitleRecord(id)
}

func (s *SubtitleService) ListSubtitlesByLanguage(language string) ([]*database.SubtitleRecord, error) {
    return s.dbManager.SearchSubtitlesByLanguage(language)
}
```

### Step 6: Remove Old databasepb Package

```bash
# Remove old package files
rm -rf pkg/databasepb/

# Update go.mod
go mod tidy

# Verify no references remain
grep -r "databasepb" pkg/ --include="*.go" || echo "All databasepb references removed"
```

## Testing Requirements

### Database Type Tests

```go
// File: pkg/database/gcommon_types_test.go
package database

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestSubtitleRecord(t *testing.T) {
    record := NewSubtitleRecord()

    // Test setting and getting fields
    record.SetID("test_id")
    assert.Equal(t, "test_id", record.GetID())

    record.SetVideoFile("/path/to/video.mp4")
    assert.Equal(t, "/path/to/video.mp4", record.GetVideoFile())

    record.SetSubtitleFile("/path/to/subtitle.srt")
    assert.Equal(t, "/path/to/subtitle.srt", record.GetSubtitleFile())

    record.SetLanguage("en")
    assert.Equal(t, "en", record.GetLanguage())

    record.SetFormat("srt")
    assert.Equal(t, "srt", record.GetFormat())

    // Test timestamps
    now := time.Now()
    record.SetCreatedAt(now)
    assert.True(t, record.GetCreatedAt().Equal(now))

    // Test gcommon integration
    gRecord := record.GetRecord()
    assert.NotNil(t, gRecord)

    // Verify record type is set
    recordType, exists := gRecord.GetField("record_type")
    assert.True(t, exists)
    assert.Equal(t, "subtitle", recordType)
}

func TestDownloadRecord(t *testing.T) {
    record := NewDownloadRecord()

    // Test setting and getting fields
    record.SetID("download_123")
    assert.Equal(t, "download_123", record.GetID())

    record.SetURL("https://example.com/video.mp4")
    assert.Equal(t, "https://example.com/video.mp4", record.GetURL())

    record.SetStatus("in_progress")
    assert.Equal(t, "in_progress", record.GetStatus())

    record.SetFilePath("/downloads/video.mp4")
    assert.Equal(t, "/downloads/video.mp4", record.GetFilePath())

    // Test gcommon integration
    gRecord := record.GetRecord()
    assert.NotNil(t, gRecord)

    // Verify record type is set
    recordType, exists := gRecord.GetField("record_type")
    assert.True(t, exists)
    assert.Equal(t, "download", recordType)
}

func TestRecordFromGcommon(t *testing.T) {
    // Test creating records from gcommon records
    gRecord := &database.Record{}
    gRecord.SetId("test_123")
    gRecord.SetField("record_type", "subtitle")
    gRecord.SetField("video_file", "/test/video.mp4")
    gRecord.SetField("language", "es")

    subtitleRecord := NewSubtitleRecordFromGcommon(gRecord)
    assert.Equal(t, "test_123", subtitleRecord.GetID())
    assert.Equal(t, "/test/video.mp4", subtitleRecord.GetVideoFile())
    assert.Equal(t, "es", subtitleRecord.GetLanguage())
}
```

### Database Operations Tests

```go
// File: pkg/database/operations_test.go
func TestDatabaseOperations(t *testing.T) {
    // Create mock gcommon database store
    mockStore := &MockDatabaseStore{}
    dbManager := NewDatabaseManager(mockStore)

    // Test subtitle record operations
    t.Run("SubtitleRecord CRUD", func(t *testing.T) {
        record := NewSubtitleRecord()
        record.SetID("subtitle_test")
        record.SetVideoFile("/test/video.mp4")
        record.SetLanguage("en")

        // Test create
        err := dbManager.CreateSubtitleRecord(record)
        assert.NoError(t, err)

        // Test get
        retrieved, err := dbManager.GetSubtitleRecord("subtitle_test")
        assert.NoError(t, err)
        assert.Equal(t, "subtitle_test", retrieved.GetID())
        assert.Equal(t, "/test/video.mp4", retrieved.GetVideoFile())

        // Test update
        retrieved.SetLanguage("es")
        err = dbManager.UpdateSubtitleRecord(retrieved)
        assert.NoError(t, err)

        // Test delete
        err = dbManager.DeleteSubtitleRecord("subtitle_test")
        assert.NoError(t, err)
    })

    // Test download record operations
    t.Run("DownloadRecord CRUD", func(t *testing.T) {
        record := NewDownloadRecord()
        record.SetID("download_test")
        record.SetURL("https://example.com/video.mp4")
        record.SetStatus("pending")

        // Test create
        err := dbManager.CreateDownloadRecord(record)
        assert.NoError(t, err)

        // Test get
        retrieved, err := dbManager.GetDownloadRecord("download_test")
        assert.NoError(t, err)
        assert.Equal(t, "download_test", retrieved.GetID())
        assert.Equal(t, "pending", retrieved.GetStatus())
    })
}

// Mock implementation for testing
type MockDatabaseStore struct {
    records map[string]*database.Record
}

func (m *MockDatabaseStore) Save(record *database.Record) error {
    if m.records == nil {
        m.records = make(map[string]*database.Record)
    }
    m.records[record.GetId()] = record
    return nil
}

func (m *MockDatabaseStore) Get(id string) (*database.Record, error) {
    if record, exists := m.records[id]; exists {
        return record, nil
    }
    return nil, fmt.Errorf("record not found")
}

// ... other mock implementations
```

### Serialization Tests

```go
// File: pkg/database/serialization_test.go
func TestSerialization(t *testing.T) {
    serializer := NewSerializationManager()

    // Test subtitle record serialization
    record := NewSubtitleRecord()
    record.SetID("test_123")
    record.SetVideoFile("/test/video.mp4")
    record.SetLanguage("en")
    record.SetFormat("srt")

    // Serialize
    data, err := serializer.SerializeSubtitleRecord(record)
    require.NoError(t, err)
    assert.NotEmpty(t, data)

    // Deserialize
    deserialized, err := serializer.DeserializeSubtitleRecord(data)
    require.NoError(t, err)

    // Verify data preservation
    assert.Equal(t, record.GetID(), deserialized.GetID())
    assert.Equal(t, record.GetVideoFile(), deserialized.GetVideoFile())
    assert.Equal(t, record.GetLanguage(), deserialized.GetLanguage())
    assert.Equal(t, record.GetFormat(), deserialized.GetFormat())
}
```

## Migration Scripts

### Data Migration Script

```go
// File: scripts/migrate_database_types.go
package main

import (
    "fmt"
    "log"

    "github.com/jdfalk/subtitle-manager/pkg/database"
)

func main() {
    fmt.Println("Migrating database types to gcommon...")

    // Connect to existing database
    // Load old databasepb records
    // Convert to new gcommon records
    // Verify data integrity
    // Update database schema if needed

    fmt.Println("Database type migration complete")
}
```

## Success Metrics

### Functional Requirements

- [ ] All databasepb types completely replaced with gcommon types
- [ ] Database operations work correctly with new types
- [ ] Serialization/deserialization preserves all data
- [ ] Historical data compatibility maintained
- [ ] Query performance equivalent or better

### Technical Requirements

- [ ] pkg/databasepb directory completely removed
- [ ] All database operations use gcommon opaque API
- [ ] No data loss during migration
- [ ] Database schema supports gcommon types
- [ ] All database tests pass with new types

### Migration Requirements

- [ ] Data migration script successfully converts existing data
- [ ] Backwards compatibility maintained during transition
- [ ] No service interruption during migration
- [ ] Rollback procedure tested and available

## Common Pitfalls

1. **Type Assertion Failures**: gcommon opaque API returns interface{}, ensure proper type checking
2. **Field Name Changes**: Ensure field names are consistent between old and new systems
3. **Serialization Format Changes**: Maintain compatibility with existing serialized data
4. **Query Performance**: Monitor query performance with new gcommon types
5. **Data Migration Errors**: Test migration thoroughly with real data

## Dependencies

- **Requires**: TASK-05-001 (configpb replacement) completed
- **Requires**: Phase 2 Core Type Migration completed
- **Requires**: gcommon database SDK properly installed
- **Enables**: Standardized database operations across gcommon services
- **Blocks**: Advanced database features until migration complete

This comprehensive task ensures complete migration from databasepb to gcommon database types while maintaining data integrity and improving standardization across the gcommon ecosystem.
