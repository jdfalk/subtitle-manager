<!-- file: docs/tasks/01-core-architecture/TASK-01-001-service-interface-definitions-PART3.md -->
<!-- version: 1.0.0 -->
<!-- guid: 01001003-1111-2222-3333-444444444444 -->

# TASK-01-001: Service Interface Definitions (PART 3)

## File Service Interface Design

**Part**: 3 of 4 (File Service Interface) **Focus**: File Operations, Media
Processing, and Storage Management

### Step 5: Define File Service Interface

**Create `proto/services/file/v1/file.proto`**:

```protobuf
// file: proto/services/file/v1/file.proto
// version: 1.0.0
// guid: file0001-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.file.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/file/v1";

import "proto/common/v1/common.proto";
import "google/protobuf/empty.proto";
import "google/protobuf/timestamp.proto";
import "google/protobuf/duration.proto";

// File Service handles all file system operations
service FileService {
  // File operations
  rpc ReadFile(ReadFileRequest) returns (ReadFileResponse);
  rpc WriteFile(WriteFileRequest) returns (WriteFileResponse);
  rpc DeleteFile(DeleteFileRequest) returns (google.protobuf.Empty);
  rpc MoveFile(MoveFileRequest) returns (MoveFileResponse);
  rpc CopyFile(CopyFileRequest) returns (CopyFileResponse);
  rpc GetFileInfo(GetFileInfoRequest) returns (GetFileInfoResponse);
  rpc ListFiles(ListFilesRequest) returns (ListFilesResponse);

  // Streaming file operations
  rpc StreamRead(StreamReadRequest) returns (stream StreamReadResponse);
  rpc StreamWrite(stream StreamWriteRequest) returns (StreamWriteResponse);

  // Directory operations
  rpc CreateDirectory(CreateDirectoryRequest) returns (CreateDirectoryResponse);
  rpc DeleteDirectory(DeleteDirectoryRequest) returns (google.protobuf.Empty);
  rpc ScanDirectory(ScanDirectoryRequest) returns (stream ScanDirectoryResponse);
  rpc GetDirectoryInfo(GetDirectoryInfoRequest) returns (GetDirectoryInfoResponse);

  // File watching and monitoring
  rpc StartWatching(StartWatchingRequest) returns (StartWatchingResponse);
  rpc StopWatching(StopWatchingRequest) returns (google.protobuf.Empty);
  rpc GetWatchStatus(GetWatchStatusRequest) returns (GetWatchStatusResponse);
  rpc GetFileEvents(GetFileEventsRequest) returns (stream FileEventResponse);

  // Media file operations
  rpc ExtractSubtitles(ExtractSubtitlesRequest) returns (ExtractSubtitlesResponse);
  rpc GetMediaMetadata(GetMediaMetadataRequest) returns (GetMediaMetadataResponse);
  rpc ConvertSubtitleFormat(ConvertSubtitleFormatRequest) returns (ConvertSubtitleFormatResponse);
  rpc EmbedSubtitles(EmbedSubtitlesRequest) returns (EmbedSubtitlesResponse);

  // Storage management
  rpc GetStorageInfo(GetStorageInfoRequest) returns (GetStorageInfoResponse);
  rpc CleanupFiles(CleanupFilesRequest) returns (CleanupFilesResponse);
  rpc ValidateFiles(ValidateFilesRequest) returns (stream ValidateFilesResponse);
  rpc BackupFiles(BackupFilesRequest) returns (stream BackupFilesResponse);

  // Health and service info
  rpc Health(common.v1.HealthCheckRequest) returns (common.v1.HealthCheckResponse);
  rpc GetServiceInfo(GetServiceInfoRequest) returns (GetServiceInfoResponse);
}

// Basic file operation messages
message ReadFileRequest {
  string file_path = 1;
  int64 offset = 2;
  int64 length = 3; // 0 means read entire file
  string encoding = 4; // "utf-8", "binary", etc.
}

message ReadFileResponse {
  bytes content = 1;
  int64 bytes_read = 2;
  string encoding = 3;
  common.v1.FileInfo file_info = 4;
  common.v1.Error error = 5;
}

message WriteFileRequest {
  string file_path = 1;
  bytes content = 2;
  bool create_directories = 3;
  bool overwrite = 4;
  string encoding = 5;
  map<string, string> metadata = 6;
}

message WriteFileResponse {
  int64 bytes_written = 1;
  common.v1.FileInfo file_info = 2;
  common.v1.Error error = 3;
}

message DeleteFileRequest {
  string file_path = 1;
  bool force = 2; // delete even if read-only
}

message MoveFileRequest {
  string source_path = 1;
  string destination_path = 2;
  bool create_directories = 3;
  bool overwrite = 4;
}

message MoveFileResponse {
  common.v1.FileInfo file_info = 1;
  common.v1.Error error = 2;
}

message CopyFileRequest {
  string source_path = 1;
  string destination_path = 2;
  bool create_directories = 3;
  bool overwrite = 4;
  bool preserve_metadata = 5;
}

message CopyFileResponse {
  common.v1.FileInfo file_info = 1;
  common.v1.Error error = 2;
}

message GetFileInfoRequest {
  string file_path = 1;
  bool include_checksum = 2;
}

message GetFileInfoResponse {
  common.v1.FileInfo file_info = 1;
  common.v1.Error error = 2;
}

message ListFilesRequest {
  string directory_path = 1;
  bool recursive = 2;
  repeated string file_patterns = 3; // "*.mp4", "*.srt"
  repeated string exclude_patterns = 4;
  bool include_hidden = 5;
  int32 page_size = 6;
  string page_token = 7;
  string sort_by = 8; // "name", "size", "modified"
  bool sort_ascending = 9;
}

message ListFilesResponse {
  repeated common.v1.FileInfo files = 1;
  string next_page_token = 2;
  int64 total_count = 3;
  int64 total_size = 4;
}

// Streaming file operations
message StreamReadRequest {
  string file_path = 1;
  int64 offset = 2;
  int32 chunk_size = 3; // bytes per chunk
  string encoding = 4;
}

message StreamReadResponse {
  bytes chunk = 1;
  int64 offset = 2;
  int64 total_size = 3;
  bool is_last_chunk = 4;
  common.v1.Error error = 5;
}

message StreamWriteRequest {
  oneof data {
    StreamWriteMetadata metadata = 1;
    bytes chunk = 2;
  }
}

message StreamWriteMetadata {
  string file_path = 1;
  int64 total_size = 2;
  bool create_directories = 3;
  bool overwrite = 4;
  string encoding = 5;
  map<string, string> metadata = 6;
}

message StreamWriteResponse {
  int64 bytes_written = 1;
  common.v1.FileInfo file_info = 2;
  common.v1.Error error = 3;
}

// Directory operations
message CreateDirectoryRequest {
  string directory_path = 1;
  bool create_parents = 2;
  uint32 mode = 3; // unix permissions
}

message CreateDirectoryResponse {
  string created_path = 1;
  google.protobuf.Timestamp created_at = 2;
  common.v1.Error error = 3;
}

message DeleteDirectoryRequest {
  string directory_path = 1;
  bool recursive = 2;
  bool force = 3;
}

message ScanDirectoryRequest {
  string directory_path = 1;
  bool recursive = 2;
  repeated string file_patterns = 3;
  repeated string exclude_patterns = 4;
  bool include_hidden = 5;
  bool calculate_checksums = 6;
}

message ScanDirectoryResponse {
  common.v1.FileInfo file_info = 1;
  float progress = 2; // 0.0 to 1.0
  int64 files_scanned = 3;
  int64 total_files = 4;
  common.v1.Error error = 5;
}

message GetDirectoryInfoRequest {
  string directory_path = 1;
  bool include_subdirectories = 2;
}

message GetDirectoryInfoResponse {
  string directory_path = 1;
  int64 file_count = 2;
  int64 directory_count = 3;
  int64 total_size = 4;
  google.protobuf.Timestamp last_modified = 5;
  repeated DirectoryEntry entries = 6;
}

message DirectoryEntry {
  string name = 1;
  bool is_directory = 2;
  int64 size = 3;
  google.protobuf.Timestamp modified = 4;
}

// File watching messages
message StartWatchingRequest {
  repeated string paths = 1;
  bool recursive = 2;
  repeated string file_patterns = 3;
  repeated string exclude_patterns = 4;
  repeated string event_types = 5; // "create", "modify", "delete", "move"
  google.protobuf.Duration debounce_duration = 6;
}

message StartWatchingResponse {
  string watch_id = 1;
  repeated string watching_paths = 2;
  google.protobuf.Timestamp started_at = 3;
  common.v1.Error error = 4;
}

message StopWatchingRequest {
  string watch_id = 1;
}

message GetWatchStatusRequest {
  string watch_id = 1;
}

message GetWatchStatusResponse {
  string watch_id = 1;
  string status = 2; // "active", "stopped", "error"
  repeated string watched_paths = 3;
  int64 events_processed = 4;
  google.protobuf.Timestamp started_at = 5;
  google.protobuf.Timestamp last_event = 6;
  common.v1.Error error = 7;
}

message GetFileEventsRequest {
  string watch_id = 1;
  google.protobuf.Timestamp since = 2;
  repeated string event_types = 3;
  repeated string file_patterns = 4;
}

message FileEventResponse {
  string watch_id = 1;
  string event_type = 2; // "create", "modify", "delete", "move"
  string file_path = 3;
  string old_path = 4; // for move events
  common.v1.FileInfo file_info = 5;
  google.protobuf.Timestamp event_time = 6;
  map<string, string> metadata = 7;
}

// Media file operations
message ExtractSubtitlesRequest {
  string media_file_path = 1;
  string output_directory = 2;
  repeated int32 track_indices = 3; // specific tracks to extract
  repeated common.v1.Language languages = 4; // language filter
  string output_format = 5; // "srt", "vtt", "ass"
  map<string, string> options = 6;
}

message ExtractSubtitlesResponse {
  repeated ExtractedSubtitle subtitles = 1;
  string extraction_log = 2;
  google.protobuf.Duration processing_time = 3;
  common.v1.Error error = 4;
}

message ExtractedSubtitle {
  string file_path = 1;
  int32 track_index = 2;
  common.v1.Language language = 3;
  string format = 4;
  int64 subtitle_count = 5;
  google.protobuf.Duration duration = 6;
  common.v1.FileInfo file_info = 7;
}

message GetMediaMetadataRequest {
  string media_file_path = 1;
  bool include_streams = 2;
  bool include_chapters = 3;
  bool include_thumbnails = 4;
}

message GetMediaMetadataResponse {
  MediaMetadata metadata = 1;
  common.v1.Error error = 2;
}

message MediaMetadata {
  string file_path = 1;
  string format = 2;
  google.protobuf.Duration duration = 3;
  int64 bitrate = 4;
  int64 file_size = 5;
  string title = 6;
  map<string, string> tags = 7;
  repeated MediaStream streams = 8;
  repeated MediaChapter chapters = 9;
  string thumbnail_path = 10;
}

message MediaStream {
  int32 index = 1;
  string type = 2; // "video", "audio", "subtitle"
  string codec = 3;
  common.v1.Language language = 4;
  string title = 5;
  map<string, string> metadata = 6;
  // Video specific
  int32 width = 7;
  int32 height = 8;
  float frame_rate = 9;
  // Audio specific
  int32 channels = 10;
  int32 sample_rate = 11;
}

message MediaChapter {
  int32 index = 1;
  google.protobuf.Duration start_time = 2;
  google.protobuf.Duration end_time = 3;
  string title = 4;
  map<string, string> metadata = 5;
}

message ConvertSubtitleFormatRequest {
  string source_file_path = 1;
  string target_format = 2; // "srt", "vtt", "ass", "ssa"
  string output_file_path = 3;
  map<string, string> conversion_options = 4;
  bool preserve_formatting = 5;
  bool preserve_timing = 6;
}

message ConvertSubtitleFormatResponse {
  string output_file_path = 1;
  string original_format = 2;
  string target_format = 3;
  int64 subtitle_count = 4;
  google.protobuf.Duration processing_time = 5;
  common.v1.FileInfo file_info = 6;
  common.v1.Error error = 7;
}

message EmbedSubtitlesRequest {
  string media_file_path = 1;
  repeated SubtitleTrack subtitle_tracks = 2;
  string output_file_path = 3;
  bool replace_existing = 4;
  map<string, string> encoding_options = 5;
}

message SubtitleTrack {
  string subtitle_file_path = 1;
  common.v1.Language language = 2;
  string title = 3;
  bool default_track = 4;
  bool forced_track = 5;
}

message EmbedSubtitlesResponse {
  string output_file_path = 1;
  int32 embedded_tracks = 2;
  google.protobuf.Duration processing_time = 3;
  common.v1.FileInfo file_info = 4;
  common.v1.Error error = 5;
}

// Storage management messages
message GetStorageInfoRequest {
  repeated string paths = 1;
  bool include_subdirectories = 2;
}

message GetStorageInfoResponse {
  repeated StorageInfo storage_info = 1;
}

message StorageInfo {
  string path = 1;
  int64 total_space = 2;
  int64 free_space = 3;
  int64 used_space = 4;
  float usage_percent = 5;
  string filesystem_type = 6;
  bool read_only = 7;
  int64 file_count = 8;
  int64 directory_count = 9;
}

message CleanupFilesRequest {
  repeated string paths = 1;
  repeated CleanupRule rules = 2;
  bool dry_run = 3;
  int64 max_files_to_delete = 4;
}

message CleanupRule {
  string pattern = 1; // file pattern to match
  google.protobuf.Duration older_than = 2;
  int64 larger_than = 3; // bytes
  int64 smaller_than = 4; // bytes
  repeated string exclude_patterns = 5;
}

message CleanupFilesResponse {
  repeated CleanupResult results = 1;
  int64 total_files_processed = 2;
  int64 total_files_deleted = 3;
  int64 total_space_freed = 4;
  google.protobuf.Duration processing_time = 5;
}

message CleanupResult {
  string file_path = 1;
  string action = 2; // "deleted", "skipped", "error"
  string reason = 3;
  int64 file_size = 4;
  common.v1.Error error = 5;
}

message ValidateFilesRequest {
  repeated string paths = 1;
  bool recursive = 2;
  repeated string file_patterns = 3;
  bool check_checksums = 4;
  bool check_permissions = 5;
  bool check_media_integrity = 6;
}

message ValidateFilesResponse {
  ValidationResult result = 1;
  float progress = 2;
  int64 files_validated = 3;
  int64 total_files = 4;
}

message ValidationResult {
  string file_path = 1;
  bool valid = 2;
  repeated ValidationIssue issues = 3;
  common.v1.FileInfo file_info = 4;
}

message ValidationIssue {
  string issue_type = 1; // "checksum_mismatch", "permission_denied", "corrupted"
  string description = 2;
  string severity = 3; // "error", "warning", "info"
}

message BackupFilesRequest {
  repeated string source_paths = 1;
  string destination_path = 2;
  bool incremental = 3;
  bool compress = 4;
  bool verify_backup = 5;
  repeated string exclude_patterns = 6;
}

message BackupFilesResponse {
  BackupResult result = 1;
  float progress = 2;
  int64 files_backed_up = 3;
  int64 total_files = 4;
}

message BackupResult {
  string backup_id = 1;
  string destination_path = 2;
  int64 files_copied = 3;
  int64 total_size = 4;
  google.protobuf.Duration processing_time = 5;
  bool verification_passed = 6;
  repeated string errors = 7;
}

// Service info messages
message GetServiceInfoRequest {
}

message GetServiceInfoResponse {
  string service_id = 1;
  string version = 2;
  google.protobuf.Timestamp started_at = 3;
  repeated string monitored_paths = 4;
  int64 files_monitored = 5;
  int64 operations_performed = 6;
  map<string, string> capabilities = 7;
  StorageInfo primary_storage = 8;
}
```

This completes Part 3 with the comprehensive File Service interface. Part 4 will
provide the Go implementations and integration guidelines.
