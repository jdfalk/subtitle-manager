<!-- file: docs/tasks/01-core-architecture/TASK-01-001-service-interface-definitions-PART1.md -->
<!-- version: 1.0.0 -->
<!-- guid: 01001001-1111-2222-3333-444444444444 -->

# TASK-01-001: Service Interface Definitions (PART 1)

## 🎯 Task Overview

**Primary Objective**: Design and implement the core service interfaces for the
3-service active-active architecture with gRPC communication protocols.

**Task Type**: Core Architecture - Interface Design **Estimated Effort**: 12-16
hours **Prerequisites**: Architecture overview understanding **Part**: 1 of 4
(Interface Design)

## 📋 Acceptance Criteria

- [ ] Complete gRPC service definitions for Web, Engine, and File services
- [ ] Protocol buffer message definitions for all service communications
- [ ] Service interface abstractions in Go
- [ ] Error handling protocols defined
- [ ] Health checking interfaces implemented
- [ ] Load balancing considerations documented
- [ ] Security authentication protocols defined

## 🔍 Current State Analysis

### Existing gRPC Infrastructure

The subtitle-manager currently has basic gRPC infrastructure:

**Existing Files to Analyze:**

- `proto/config.proto` - Configuration service definitions
- `pkg/grpcserver/server.go` - Current gRPC server implementation
- `pkg/grpcserver/` - gRPC service implementations

**Current Service Structure (TO BE REPLACED):**

```go
// Current monolithic approach
type TranslatorServer struct {
    translator *translator.Translator
    webserver  *webserver.WebServer
}
```

**Target Service Structure (NEW):**

```go
// New distributed approach
type WebService interface {
    // Client API methods
}

type EngineService interface {
    // Translation and coordination methods
}

type FileService interface {
    // File operation methods
}
```

## 📝 Implementation Steps

### Step 1: Create Core Service Directory Structure

Create the new service interface package structure:

```bash
mkdir -p pkg/services/interfaces
mkdir -p pkg/services/web
mkdir -p pkg/services/engine
mkdir -p pkg/services/file
mkdir -p proto/services/web/v1
mkdir -p proto/services/engine/v1
mkdir -p proto/services/file/v1
mkdir -p proto/common/v1
```

### Step 2: Define Common Types and Messages

**Create `proto/common/v1/common.proto`**:

```protobuf
// file: proto/common/v1/common.proto
// version: 1.0.0
// guid: common01-1111-2222-3333-444444444444

syntax = "proto3";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/timestamp.proto";
import "google/protobuf/duration.proto";

// Common request/response patterns
message RequestMetadata {
  string request_id = 1;
  string user_id = 2;
  google.protobuf.Timestamp timestamp = 3;
  map<string, string> headers = 4;
}

message ResponseMetadata {
  string request_id = 1;
  google.protobuf.Timestamp timestamp = 2;
  google.protobuf.Duration processing_time = 3;
  string service_instance = 4;
}

// Error handling
message Error {
  enum Code {
    UNKNOWN = 0;
    INVALID_ARGUMENT = 1;
    NOT_FOUND = 2;
    PERMISSION_DENIED = 3;
    RESOURCE_EXHAUSTED = 4;
    INTERNAL = 5;
    UNAVAILABLE = 6;
    DEADLINE_EXCEEDED = 7;
  }

  Code code = 1;
  string message = 2;
  string details = 3;
  map<string, string> metadata = 4;
}

// Health checking
message HealthCheckRequest {
  string service = 1;
}

message HealthCheckResponse {
  enum Status {
    UNKNOWN = 0;
    SERVING = 1;
    NOT_SERVING = 2;
    SERVICE_UNKNOWN = 3;
  }

  Status status = 1;
  string message = 2;
  map<string, string> details = 3;
}

// File information
message FileInfo {
  string path = 1;
  string name = 2;
  int64 size = 3;
  google.protobuf.Timestamp modified_time = 4;
  string mime_type = 5;
  string checksum = 6;
  map<string, string> metadata = 7;
}

// Language specification
message Language {
  string code = 1;      // ISO 639-1 code (e.g., "en")
  string name = 2;      // Human readable name (e.g., "English")
  string region = 3;    // Optional region (e.g., "US")
}

// Translation job
message TranslationJob {
  string id = 1;
  string source_text = 2;
  Language source_language = 3;
  Language target_language = 4;
  string provider = 5;  // "google", "openai", "whisper"
  map<string, string> options = 6;
  google.protobuf.Timestamp created_at = 7;
}

// Translation result
message TranslationResult {
  string job_id = 1;
  string translated_text = 2;
  float confidence = 3;
  google.protobuf.Duration processing_time = 4;
  string provider_used = 5;
  map<string, string> metadata = 6;
}
```

### Step 3: Define Web Service Interface

**Create `proto/services/web/v1/web.proto`**:

```protobuf
// file: proto/services/web/v1/web.proto
// version: 1.0.0
// guid: web00001-1111-2222-3333-444444444444

syntax = "proto3";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "proto/common/v1/common.proto";
import "google/protobuf/empty.proto";

// Web Service handles all client interactions
service WebService {
  // Authentication
  rpc Authenticate(AuthenticateRequest) returns (AuthenticateResponse);
  rpc ValidateSession(ValidateSessionRequest) returns (ValidateSessionResponse);
  rpc Logout(LogoutRequest) returns (google.protobuf.Empty);

  // File upload/download
  rpc UploadFile(stream UploadFileRequest) returns (UploadFileResponse);
  rpc DownloadFile(DownloadFileRequest) returns (stream DownloadFileResponse);
  rpc GetFileInfo(GetFileInfoRequest) returns (GetFileInfoResponse);

  // Client API endpoints
  rpc GetStatus(google.protobuf.Empty) returns (StatusResponse);
  rpc ListMediaFiles(ListMediaFilesRequest) returns (ListMediaFilesResponse);
  rpc SearchSubtitles(SearchSubtitlesRequest) returns (SearchSubtitlesResponse);

  // Real-time updates
  rpc Subscribe(SubscribeRequest) returns (stream EventResponse);

  // Health checking
  rpc Health(common.v1.HealthCheckRequest) returns (common.v1.HealthCheckResponse);
}

// Authentication messages
message AuthenticateRequest {
  string username = 1;
  string password = 2;
  string provider = 3; // "local", "oauth", etc.
}

message AuthenticateResponse {
  string token = 1;
  string session_id = 2;
  google.protobuf.Timestamp expires_at = 3;
  UserInfo user = 4;
}

message ValidateSessionRequest {
  string session_id = 1;
  string token = 2;
}

message ValidateSessionResponse {
  bool valid = 1;
  UserInfo user = 2;
  google.protobuf.Timestamp expires_at = 3;
}

message LogoutRequest {
  string session_id = 1;
}

message UserInfo {
  string id = 1;
  string username = 2;
  string email = 3;
  repeated string roles = 4;
  map<string, string> preferences = 5;
}

// File upload/download messages
message UploadFileRequest {
  oneof data {
    UploadMetadata metadata = 1;
    bytes chunk = 2;
  }
}

message UploadMetadata {
  string filename = 1;
  string content_type = 2;
  int64 total_size = 3;
  string destination_path = 4;
  bool overwrite = 5;
}

message UploadFileResponse {
  common.v1.FileInfo file_info = 1;
  string upload_id = 2;
  bool success = 3;
  common.v1.Error error = 4;
}

message DownloadFileRequest {
  string file_path = 1;
  int64 offset = 2;
  int64 limit = 3;
}

message DownloadFileResponse {
  bytes chunk = 1;
  int64 total_size = 2;
  string content_type = 3;
}

message GetFileInfoRequest {
  string file_path = 1;
}

message GetFileInfoResponse {
  common.v1.FileInfo file_info = 1;
  common.v1.Error error = 2;
}

// Client API messages
message StatusResponse {
  string version = 1;
  google.protobuf.Timestamp uptime = 2;
  map<string, string> services = 3; // service_name -> status
  int32 active_jobs = 4;
  int64 total_files = 5;
}

message ListMediaFilesRequest {
  string path = 1;
  bool recursive = 2;
  repeated string file_types = 3; // "video", "subtitle"
  int32 page_size = 4;
  string page_token = 5;
}

message ListMediaFilesResponse {
  repeated MediaFile files = 1;
  string next_page_token = 2;
  int64 total_count = 3;
}

message MediaFile {
  common.v1.FileInfo file_info = 1;
  MediaMetadata metadata = 2;
  repeated SubtitleFile subtitles = 3;
}

message MediaMetadata {
  string title = 1;
  google.protobuf.Duration duration = 2;
  string resolution = 3;
  repeated common.v1.Language audio_tracks = 4;
  repeated common.v1.Language subtitle_tracks = 5;
}

message SubtitleFile {
  common.v1.FileInfo file_info = 1;
  common.v1.Language language = 2;
  string format = 3; // "srt", "vtt", "ass"
  bool embedded = 4;
}

message SearchSubtitlesRequest {
  string query = 1;
  repeated common.v1.Language languages = 2;
  repeated string providers = 3;
  int32 limit = 4;
}

message SearchSubtitlesResponse {
  repeated SubtitleSearchResult results = 1;
  string search_id = 2;
}

message SubtitleSearchResult {
  string id = 1;
  string title = 2;
  common.v1.Language language = 3;
  string provider = 4;
  float score = 5;
  string download_url = 6;
  map<string, string> metadata = 7;
}

// Real-time events
message SubscribeRequest {
  repeated string event_types = 1; // "translation", "file_change", "job_status"
  map<string, string> filters = 2;
}

message EventResponse {
  string event_type = 1;
  google.protobuf.Timestamp timestamp = 2;
  oneof data {
    TranslationEvent translation = 3;
    FileChangeEvent file_change = 4;
    JobStatusEvent job_status = 5;
  }
}

message TranslationEvent {
  string job_id = 1;
  string status = 2; // "started", "completed", "failed"
  float progress = 3;
  common.v1.TranslationResult result = 4;
}

message FileChangeEvent {
  string file_path = 1;
  string change_type = 2; // "created", "modified", "deleted"
  common.v1.FileInfo file_info = 3;
}

message JobStatusEvent {
  string job_id = 1;
  string job_type = 2;
  string status = 3;
  float progress = 4;
  google.protobuf.Timestamp started_at = 5;
  google.protobuf.Timestamp updated_at = 6;
}
```

### Step 4: Define Engine Service Interface

**Create `proto/services/engine/v1/engine.proto`**:

```protobuf
// file: proto/services/engine/v1/engine.proto
// version: 1.0.0
// guid: engine01-1111-2222-3333-444444444444

syntax = "proto3";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "proto/common/v1/common.proto";
import "google/protobuf/empty.proto";
import "google/protobuf/timestamp.proto";
import "google/protobuf/duration.proto";

// Engine Service handles translation, monitoring, and coordination
service EngineService {
  // Translation services (Active-Active)
  rpc TranslateText(TranslateTextRequest) returns (TranslateTextResponse);
  rpc TranslateSubtitle(TranslateSubtitleRequest) returns (TranslateSubtitleResponse);
  rpc BatchTranslate(BatchTranslateRequest) returns (stream BatchTranslateResponse);
  rpc GetTranslationStatus(GetTranslationStatusRequest) returns (GetTranslationStatusResponse);
  rpc CancelTranslation(CancelTranslationRequest) returns (google.protobuf.Empty);

  // Job queue services (Active-Active)
  rpc SubmitJob(SubmitJobRequest) returns (SubmitJobResponse);
  rpc GetJobStatus(GetJobStatusRequest) returns (GetJobStatusResponse);
  rpc ListJobs(ListJobsRequest) returns (ListJobsResponse);
  rpc CancelJob(CancelJobRequest) returns (google.protobuf.Empty);

  // Monitoring services (Active-Active)
  rpc StartMonitoring(StartMonitoringRequest) returns (StartMonitoringResponse);
  rpc StopMonitoring(StopMonitoringRequest) returns (google.protobuf.Empty);
  rpc GetMonitoringStatus(GetMonitoringStatusRequest) returns (GetMonitoringStatusResponse);
  rpc ProcessFileEvent(ProcessFileEventRequest) returns (ProcessFileEventResponse);

  // Coordination services (Leader Only)
  rpc DistributeWork(DistributeWorkRequest) returns (DistributeWorkResponse);
  rpc CoordinateMonitoring(CoordinateMonitoringRequest) returns (CoordinateMonitoringResponse);
  rpc GetSystemStatus(GetSystemStatusRequest) returns (GetSystemStatusResponse);
  rpc ScheduleTask(ScheduleTaskRequest) returns (ScheduleTaskResponse);

  // Leader election services
  rpc RequestLeadership(RequestLeadershipRequest) returns (RequestLeadershipResponse);
  rpc HeartbeatLeader(HeartbeatLeaderRequest) returns (HeartbeatLeaderResponse);
  rpc GetLeaderInfo(GetLeaderInfoRequest) returns (GetLeaderInfoResponse);
  rpc TransferLeadership(TransferLeadershipRequest) returns (TransferLeadershipResponse);

  // Health and discovery
  rpc Health(common.v1.HealthCheckRequest) returns (common.v1.HealthCheckResponse);
  rpc GetServiceInfo(GetServiceInfoRequest) returns (GetServiceInfoResponse);
  rpc RegisterService(RegisterServiceRequest) returns (RegisterServiceResponse);
}

// Translation request messages
message TranslateTextRequest {
  string text = 1;
  common.v1.Language source_language = 2;
  common.v1.Language target_language = 3;
  string provider = 4; // "google", "openai", "whisper"
  map<string, string> options = 5;
  string priority = 6; // "low", "normal", "high"
}

message TranslateTextResponse {
  string translated_text = 1;
  float confidence = 2;
  google.protobuf.Duration processing_time = 3;
  string provider_used = 4;
  string job_id = 5;
  common.v1.Error error = 6;
}

message TranslateSubtitleRequest {
  string subtitle_content = 1;
  string format = 2; // "srt", "vtt", "ass"
  common.v1.Language source_language = 3;
  common.v1.Language target_language = 4;
  string provider = 5;
  map<string, string> options = 6;
  bool preserve_timing = 7;
  bool preserve_formatting = 8;
}

message TranslateSubtitleResponse {
  string translated_content = 1;
  string format = 2;
  float confidence = 3;
  google.protobuf.Duration processing_time = 4;
  string provider_used = 5;
  string job_id = 6;
  repeated TranslationSegment segments = 7;
  common.v1.Error error = 8;
}

message TranslationSegment {
  int32 sequence = 1;
  google.protobuf.Duration start_time = 2;
  google.protobuf.Duration end_time = 3;
  string original_text = 4;
  string translated_text = 5;
  float confidence = 6;
}

message BatchTranslateRequest {
  repeated TranslateTextRequest requests = 1;
  string batch_id = 2;
  int32 max_concurrent = 3;
  string priority = 4;
}

message BatchTranslateResponse {
  string batch_id = 1;
  int32 total_items = 2;
  int32 completed_items = 3;
  int32 failed_items = 4;
  repeated TranslateTextResponse results = 5;
  bool finished = 6;
  common.v1.Error error = 7;
}

message GetTranslationStatusRequest {
  string job_id = 1;
}

message GetTranslationStatusResponse {
  string job_id = 1;
  string status = 2; // "pending", "processing", "completed", "failed"
  float progress = 3;
  google.protobuf.Timestamp started_at = 4;
  google.protobuf.Timestamp updated_at = 5;
  google.protobuf.Duration estimated_remaining = 6;
  TranslateTextResponse result = 7;
  common.v1.Error error = 8;
}

message CancelTranslationRequest {
  string job_id = 1;
  string reason = 2;
}

// Job queue messages
message SubmitJobRequest {
  string job_type = 1; // "translation", "extraction", "monitoring"
  bytes job_data = 2;
  string priority = 3;
  google.protobuf.Timestamp schedule_at = 4;
  map<string, string> metadata = 5;
}

message SubmitJobResponse {
  string job_id = 1;
  string status = 2;
  google.protobuf.Timestamp created_at = 3;
  common.v1.Error error = 4;
}

message GetJobStatusRequest {
  string job_id = 1;
}

message GetJobStatusResponse {
  string job_id = 1;
  string job_type = 2;
  string status = 3;
  float progress = 4;
  google.protobuf.Timestamp created_at = 5;
  google.protobuf.Timestamp started_at = 6;
  google.protobuf.Timestamp updated_at = 7;
  google.protobuf.Timestamp completed_at = 8;
  bytes result_data = 9;
  common.v1.Error error = 10;
  string assigned_worker = 11;
  int32 retry_count = 12;
}

message ListJobsRequest {
  repeated string job_types = 1;
  repeated string statuses = 2;
  google.protobuf.Timestamp start_time = 3;
  google.protobuf.Timestamp end_time = 4;
  int32 page_size = 5;
  string page_token = 6;
}

message ListJobsResponse {
  repeated GetJobStatusResponse jobs = 1;
  string next_page_token = 2;
  int64 total_count = 3;
}

message CancelJobRequest {
  string job_id = 1;
  string reason = 2;
}

// Monitoring messages
message StartMonitoringRequest {
  repeated string paths = 1;
  repeated string file_patterns = 2;
  bool recursive = 3;
  google.protobuf.Duration scan_interval = 4;
  map<string, string> options = 5;
}

message StartMonitoringResponse {
  string monitor_id = 1;
  string status = 2;
  google.protobuf.Timestamp started_at = 3;
  common.v1.Error error = 4;
}

message StopMonitoringRequest {
  string monitor_id = 1;
}

message GetMonitoringStatusRequest {
  string monitor_id = 1;
}

message GetMonitoringStatusResponse {
  string monitor_id = 1;
  string status = 2;
  repeated string monitored_paths = 3;
  int64 files_watched = 4;
  int64 events_processed = 5;
  google.protobuf.Timestamp last_scan = 6;
  google.protobuf.Timestamp started_at = 7;
}

message ProcessFileEventRequest {
  string event_type = 1; // "created", "modified", "deleted", "moved"
  string file_path = 2;
  string old_path = 3; // for move events
  common.v1.FileInfo file_info = 4;
  google.protobuf.Timestamp event_time = 5;
}

message ProcessFileEventResponse {
  bool processed = 1;
  repeated string actions_taken = 2;
  string job_id = 3; // if work was scheduled
  common.v1.Error error = 4;
}

// Coordination messages (Leader Only)
message DistributeWorkRequest {
  string work_type = 1;
  bytes work_data = 2;
  repeated string preferred_workers = 3;
  string priority = 4;
  map<string, string> requirements = 5;
}

message DistributeWorkResponse {
  string work_id = 1;
  string assigned_worker = 2;
  google.protobuf.Timestamp assigned_at = 3;
  common.v1.Error error = 4;
}

message CoordinateMonitoringRequest {
  repeated MonitoringAssignment assignments = 1;
}

message MonitoringAssignment {
  string worker_id = 1;
  repeated string paths = 2;
  map<string, string> configuration = 3;
}

message CoordinateMonitoringResponse {
  repeated MonitoringAssignmentResult results = 1;
  common.v1.Error error = 2;
}

message MonitoringAssignmentResult {
  string worker_id = 1;
  bool success = 2;
  string error_message = 3;
  google.protobuf.Timestamp assigned_at = 4;
}

message GetSystemStatusRequest {
  bool include_details = 1;
}

message GetSystemStatusResponse {
  string overall_status = 1; // "healthy", "degraded", "unhealthy"
  int32 total_workers = 2;
  int32 active_workers = 3;
  int32 pending_jobs = 4;
  int32 running_jobs = 5;
  repeated WorkerStatus workers = 6;
  repeated ServiceHealth services = 7;
  SystemMetrics metrics = 8;
}

message WorkerStatus {
  string worker_id = 1;
  string status = 2;
  string worker_type = 3; // "engine", "file"
  google.protobuf.Timestamp last_heartbeat = 4;
  int32 active_jobs = 5;
  map<string, string> capabilities = 6;
  ResourceUsage resource_usage = 7;
}

message ServiceHealth {
  string service_name = 1;
  string status = 2;
  google.protobuf.Timestamp last_check = 3;
  string endpoint = 4;
  google.protobuf.Duration response_time = 5;
}

message SystemMetrics {
  int64 total_translations = 1;
  int64 successful_translations = 2;
  int64 failed_translations = 3;
  google.protobuf.Duration average_processing_time = 4;
  int64 files_monitored = 5;
  int64 events_processed = 6;
  google.protobuf.Timestamp metrics_time = 7;
}

message ResourceUsage {
  float cpu_percent = 1;
  int64 memory_bytes = 2;
  int64 disk_bytes = 3;
  int32 open_files = 4;
  int32 network_connections = 5;
}

message ScheduleTaskRequest {
  string task_type = 1;
  bytes task_data = 2;
  google.protobuf.Timestamp schedule_at = 3;
  google.protobuf.Duration repeat_interval = 4;
  string target_worker = 5;
  map<string, string> options = 6;
}

message ScheduleTaskResponse {
  string task_id = 1;
  google.protobuf.Timestamp scheduled_at = 2;
  string assigned_worker = 3;
  common.v1.Error error = 4;
}

// Leader election messages
message RequestLeadershipRequest {
  string candidate_id = 1;
  string candidate_endpoint = 2;
  int64 term = 3;
  google.protobuf.Timestamp candidate_start_time = 4;
  map<string, string> candidate_metadata = 5;
}

message RequestLeadershipResponse {
  bool granted = 1;
  int64 term = 2;
  string current_leader = 3;
  google.protobuf.Duration lease_duration = 4;
  common.v1.Error error = 5;
}

message HeartbeatLeaderRequest {
  string leader_id = 1;
  int64 term = 2;
  google.protobuf.Timestamp timestamp = 3;
  SystemMetrics metrics = 4;
}

message HeartbeatLeaderResponse {
  bool acknowledged = 1;
  int64 term = 2;
  google.protobuf.Duration next_heartbeat = 3;
  common.v1.Error error = 4;
}

message GetLeaderInfoRequest {
}

message GetLeaderInfoResponse {
  string leader_id = 1;
  string leader_endpoint = 2;
  int64 term = 3;
  google.protobuf.Timestamp elected_at = 4;
  google.protobuf.Timestamp last_heartbeat = 5;
  bool is_healthy = 6;
}

message TransferLeadershipRequest {
  string current_leader = 1;
  string new_leader = 2;
  string reason = 3;
}

message TransferLeadershipResponse {
  bool success = 1;
  string new_leader = 2;
  int64 new_term = 3;
  common.v1.Error error = 4;
}

// Service discovery messages
message GetServiceInfoRequest {
}

message GetServiceInfoResponse {
  string service_id = 1;
  string service_type = 2; // "engine", "web", "file"
  string endpoint = 3;
  string version = 4;
  google.protobuf.Timestamp started_at = 5;
  repeated string capabilities = 6;
  map<string, string> metadata = 7;
  ResourceUsage resource_usage = 8;
}

message RegisterServiceRequest {
  string service_id = 1;
  string service_type = 2;
  string endpoint = 3;
  repeated string capabilities = 4;
  map<string, string> metadata = 5;
  google.protobuf.Duration ttl = 6;
}

message RegisterServiceResponse {
  bool success = 1;
  google.protobuf.Timestamp registered_at = 2;
  google.protobuf.Duration heartbeat_interval = 3;
  common.v1.Error error = 4;
}
```

This completes Part 2 with the comprehensive Engine Service interface. Part 3
will cover the File Service interface, and Part 4 will provide the Go
implementations.

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

syntax = "proto3";

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

<!-- file: docs/tasks/01-core-architecture/TASK-01-001-service-interface-definitions-PART4.md -->
<!-- version: 1.0.0 -->
<!-- guid: 01001004-1111-2222-3333-444444444444 -->

# TASK-01-001: Service Interface Definitions (PART 4)

## Go Interface Implementations

**Part**: 4 of 4 (Go Interfaces and Implementation Guidelines) **Focus**:
Concrete Go interfaces, service abstractions, and implementation patterns

### Step 6: Define Go Service Interfaces

**Create `pkg/services/interfaces.go`**:

```go

package services

import (
 "context"
 "io"
 "time"

 commonv1 "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1"
 enginev1 "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1"
 filev1 "github.com/jdfalk/subtitle-manager/pkg/proto/file/v1"
 webv1 "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1"
)

// ServiceManager coordinates all services and handles lifecycle
type ServiceManager interface {
 // Service lifecycle
 StartServices(ctx context.Context) error
 StopServices(ctx context.Context) error
 RestartService(ctx context.Context, serviceName string) error
 GetServiceStatus(serviceName string) (*commonv1.ServiceStatus, error)

 // Service discovery
 RegisterService(service Service) error
 UnregisterService(serviceID string) error
 DiscoverServices(serviceType string) ([]Service, error)

 // Health monitoring
 HealthCheck(ctx context.Context) (*commonv1.HealthCheckResponse, error)
 GetAllServiceHealth(ctx context.Context) (map[string]*commonv1.HealthCheckResponse, error)
}

// Service represents the common interface for all services
type Service interface {
 // Basic service operations
 ID() string
 Type() string
 Start(ctx context.Context) error
 Stop(ctx context.Context) error
 Health(ctx context.Context) (*commonv1.HealthCheckResponse, error)

 // Configuration
 Configure(config interface{}) error
 GetConfig() interface{}

 // Metrics and monitoring
 GetMetrics() (*commonv1.ServiceMetrics, error)
 GetEvents(since time.Time) ([]*commonv1.ServiceEvent, error)
}

// WebService handles all web-facing operations
type WebService interface {
 Service

 // User management
 AuthenticateUser(ctx context.Context, req *webv1.AuthenticateUserRequest) (*webv1.AuthenticateUserResponse, error)
 GetUser(ctx context.Context, req *webv1.GetUserRequest) (*webv1.GetUserResponse, error)
 UpdateUserPreferences(ctx context.Context, req *webv1.UpdateUserPreferencesRequest) (*webv1.UpdateUserPreferencesResponse, error)

 // Subtitle management
 UploadSubtitle(ctx context.Context, req *webv1.UploadSubtitleRequest) (*webv1.UploadSubtitleResponse, error)
 DownloadSubtitle(ctx context.Context, req *webv1.DownloadSubtitleRequest) (*webv1.DownloadSubtitleResponse, error)
 SearchSubtitles(ctx context.Context, req *webv1.SearchSubtitlesRequest) (*webv1.SearchSubtitlesResponse, error)
 GetSubtitleMetadata(ctx context.Context, req *webv1.GetSubtitleMetadataRequest) (*webv1.GetSubtitleMetadataResponse, error)

 // Translation operations
 TranslateSubtitle(ctx context.Context, req *webv1.TranslateSubtitleRequest) (*webv1.TranslateSubtitleResponse, error)
 GetTranslationStatus(ctx context.Context, req *webv1.GetTranslationStatusRequest) (*webv1.GetTranslationStatusResponse, error)

 // File operations
 UploadFile(stream webv1.WebService_UploadFileServer) error
 DownloadFile(req *webv1.DownloadFileRequest, stream webv1.WebService_DownloadFileServer) error
 GetFileInfo(ctx context.Context, req *webv1.GetFileInfoRequest) (*webv1.GetFileInfoResponse, error)

 // Library management
 GetLibraries(ctx context.Context, req *webv1.GetLibrariesRequest) (*webv1.GetLibrariesResponse, error)
 ScanLibrary(ctx context.Context, req *webv1.ScanLibraryRequest) (*webv1.ScanLibraryResponse, error)

 // System operations
 GetSystemStatus(ctx context.Context, req *webv1.GetSystemStatusRequest) (*webv1.GetSystemStatusResponse, error)
 UpdateConfiguration(ctx context.Context, req *webv1.UpdateConfigurationRequest) (*webv1.UpdateConfigurationResponse, error)
}

// EngineService handles core processing, translation, and coordination
type EngineService interface {
 Service

 // Translation operations
 ProcessTranslation(ctx context.Context, req *enginev1.ProcessTranslationRequest) (*enginev1.ProcessTranslationResponse, error)
 GetTranslationProgress(ctx context.Context, req *enginev1.GetTranslationProgressRequest) (*enginev1.GetTranslationProgressResponse, error)
 CancelTranslation(ctx context.Context, req *enginev1.CancelTranslationRequest) (*enginev1.CancelTranslationResponse, error)

 // Monitoring operations
 StartMonitoring(ctx context.Context, req *enginev1.StartMonitoringRequest) (*enginev1.StartMonitoringResponse, error)
 StopMonitoring(ctx context.Context, req *enginev1.StopMonitoringRequest) (*enginev1.StopMonitoringResponse, error)
 GetMonitoringStatus(ctx context.Context, req *enginev1.GetMonitoringStatusRequest) (*enginev1.GetMonitoringStatusResponse, error)
 GetMonitoringEvents(req *enginev1.GetMonitoringEventsRequest, stream enginev1.EngineService_GetMonitoringEventsServer) error

 // Processing operations
 ProcessMedia(ctx context.Context, req *enginev1.ProcessMediaRequest) (*enginev1.ProcessMediaResponse, error)
 ExtractSubtitles(ctx context.Context, req *enginev1.ExtractSubtitlesRequest) (*enginev1.ExtractSubtitlesResponse, error)
 ConvertFormat(ctx context.Context, req *enginev1.ConvertFormatRequest) (*enginev1.ConvertFormatResponse, error)

 // Coordination operations
 RegisterWorker(ctx context.Context, req *enginev1.RegisterWorkerRequest) (*enginev1.RegisterWorkerResponse, error)
 HeartbeatWorker(ctx context.Context, req *enginev1.HeartbeatWorkerRequest) (*enginev1.HeartbeatWorkerResponse, error)
 AssignTask(ctx context.Context, req *enginev1.AssignTaskRequest) (*enginev1.AssignTaskResponse, error)
 CompleteTask(ctx context.Context, req *enginev1.CompleteTaskRequest) (*enginev1.CompleteTaskResponse, error)

 // Leader election operations
 RequestLeadership(ctx context.Context, req *enginev1.RequestLeadershipRequest) (*enginev1.RequestLeadershipResponse, error)
 ReleaseLeadership(ctx context.Context, req *enginev1.ReleaseLeadershipRequest) (*enginev1.ReleaseLeadershipResponse, error)
 GetLeaderInfo(ctx context.Context, req *enginev1.GetLeaderInfoRequest) (*enginev1.GetLeaderInfoResponse, error)
 HeartbeatLeader(ctx context.Context, req *enginev1.HeartbeatLeaderRequest) (*enginev1.HeartbeatLeaderResponse, error)
}

// FileService handles all file system operations
type FileService interface {
 Service

 // Basic file operations
 ReadFile(ctx context.Context, req *filev1.ReadFileRequest) (*filev1.ReadFileResponse, error)
 WriteFile(ctx context.Context, req *filev1.WriteFileRequest) (*filev1.WriteFileResponse, error)
 DeleteFile(ctx context.Context, req *filev1.DeleteFileRequest) error
 MoveFile(ctx context.Context, req *filev1.MoveFileRequest) (*filev1.MoveFileResponse, error)
 CopyFile(ctx context.Context, req *filev1.CopyFileRequest) (*filev1.CopyFileResponse, error)

 // Streaming operations
 StreamRead(req *filev1.StreamReadRequest, stream filev1.FileService_StreamReadServer) error
 StreamWrite(stream filev1.FileService_StreamWriteServer) error

 // Directory operations
 CreateDirectory(ctx context.Context, req *filev1.CreateDirectoryRequest) (*filev1.CreateDirectoryResponse, error)
 DeleteDirectory(ctx context.Context, req *filev1.DeleteDirectoryRequest) error
 ScanDirectory(req *filev1.ScanDirectoryRequest, stream filev1.FileService_ScanDirectoryServer) error

 // File watching
 StartWatching(ctx context.Context, req *filev1.StartWatchingRequest) (*filev1.StartWatchingResponse, error)
 StopWatching(ctx context.Context, req *filev1.StopWatchingRequest) error
 GetFileEvents(req *filev1.GetFileEventsRequest, stream filev1.FileService_GetFileEventsServer) error

 // Media operations
 ExtractSubtitles(ctx context.Context, req *filev1.ExtractSubtitlesRequest) (*filev1.ExtractSubtitlesResponse, error)
 GetMediaMetadata(ctx context.Context, req *filev1.GetMediaMetadataRequest) (*filev1.GetMediaMetadataResponse, error)
 ConvertSubtitleFormat(ctx context.Context, req *filev1.ConvertSubtitleFormatRequest) (*filev1.ConvertSubtitleFormatResponse, error)

 // Storage management
 GetStorageInfo(ctx context.Context, req *filev1.GetStorageInfoRequest) (*filev1.GetStorageInfoResponse, error)
 CleanupFiles(ctx context.Context, req *filev1.CleanupFilesRequest) (*filev1.CleanupFilesResponse, error)
 ValidateFiles(req *filev1.ValidateFilesRequest, stream filev1.FileService_ValidateFilesServer) error
}

// ServiceClient provides client interfaces for inter-service communication
type ServiceClient interface {
 // Service discovery
 ConnectToService(ctx context.Context, serviceName, serviceType string) error
 DisconnectFromService(serviceName string) error

 // Client instances
 GetWebClient() (webv1.WebServiceClient, error)
 GetEngineClient() (enginev1.EngineServiceClient, error)
 GetFileClient() (filev1.FileServiceClient, error)

 // Health monitoring
 PingService(ctx context.Context, serviceName string) error
 WatchServiceHealth(ctx context.Context, serviceName string) (<-chan *commonv1.HealthCheckResponse, error)
}

// Configuration interfaces
type ServiceConfig interface {
 Validate() error
 GetServiceName() string
 GetServiceType() string
 GetListenAddress() string
 GetDiscoveryConfig() *DiscoveryConfig
}

type DiscoveryConfig struct {
 Method   string            // "static", "consul", "etcd", "kubernetes"
 Endpoints []string         // service endpoints
 Options  map[string]string // discovery-specific options
}

// Worker and task management interfaces
type Worker interface {
 ID() string
 Type() string
 Capabilities() []string
 ProcessTask(ctx context.Context, task *enginev1.Task) (*enginev1.TaskResult, error)
 GetStatus() *enginev1.WorkerStatus
 Shutdown(ctx context.Context) error
}

type TaskManager interface {
 SubmitTask(ctx context.Context, task *enginev1.Task) (*enginev1.TaskSubmitResponse, error)
 GetTaskStatus(ctx context.Context, taskID string) (*enginev1.TaskStatus, error)
 CancelTask(ctx context.Context, taskID string) error
 ListTasks(ctx context.Context, filter *enginev1.TaskFilter) ([]*enginev1.Task, error)
 GetTaskResult(ctx context.Context, taskID string) (*enginev1.TaskResult, error)
}

// Leader election interface
type LeaderElection interface {
 Start(ctx context.Context) error
 Stop() error
 IsLeader() bool
 GetLeader() (string, error)
 OnBecomeLeader(callback func(ctx context.Context))
 OnLoseLeadership(callback func())
}

// File watcher interface
type FileWatcher interface {
 Start(ctx context.Context) error
 Stop() error
 AddPath(path string, recursive bool) error
 RemovePath(path string) error
 GetEvents() <-chan *filev1.FileEventResponse
 GetStatus() *filev1.GetWatchStatusResponse
}

// Media processor interface
type MediaProcessor interface {
 ExtractSubtitles(ctx context.Context, mediaPath string, options *MediaProcessingOptions) ([]*ExtractedSubtitle, error)
 GetMetadata(ctx context.Context, mediaPath string) (*MediaMetadata, error)
 ConvertFormat(ctx context.Context, inputPath, outputPath, format string) error
 ValidateMedia(ctx context.Context, mediaPath string) (*MediaValidationResult, error)
}

type MediaProcessingOptions struct {
 OutputDir      string
 Languages      []string
 Formats        []string
 TrackIndices   []int
 PreserveAspect bool
 Quality        string
}

type ExtractedSubtitle struct {
 FilePath     string
 TrackIndex   int
 Language     string
 Format       string
 SubtitleCount int64
 Duration     time.Duration
}

type MediaMetadata struct {
 FilePath   string
 Format     string
 Duration   time.Duration
 Bitrate    int64
 Size       int64
 Title      string
 Tags       map[string]string
 Streams    []*MediaStream
 Chapters   []*MediaChapter
}

type MediaStream struct {
 Index      int
 Type       string // "video", "audio", "subtitle"
 Codec      string
 Language   string
 Title      string
 Metadata   map[string]string
 Width      int32
 Height     int32
 FrameRate  float32
 Channels   int32
 SampleRate int32
}

type MediaChapter struct {
 Index     int
 StartTime time.Duration
 EndTime   time.Duration
 Title     string
 Metadata  map[string]string
}

type MediaValidationResult struct {
 Valid   bool
 Issues  []ValidationIssue
 Details map[string]interface{}
}

type ValidationIssue struct {
 Type        string
 Description string
 Severity    string
}

// Translation service interface
type TranslationService interface {
 Translate(ctx context.Context, req *TranslationRequest) (*TranslationResponse, error)
 GetSupportedLanguages() ([]string, error)
 EstimateTranslationTime(ctx context.Context, req *TranslationRequest) (time.Duration, error)
 GetTranslationQuality(ctx context.Context, sourceText, translatedText string, fromLang, toLang string) (float64, error)
}

type TranslationRequest struct {
 Text         string
 FromLanguage string
 ToLanguage   string
 Context      string
 Options      map[string]interface{}
}

type TranslationResponse struct {
 TranslatedText string
 Confidence     float64
 ProcessingTime time.Duration
 Metadata       map[string]interface{}
}

// Storage interface
type StorageManager interface {
 GetStorageInfo(ctx context.Context, paths []string) ([]*StorageInfo, error)
 CleanupStorage(ctx context.Context, rules []*CleanupRule) (*CleanupResult, error)
 BackupFiles(ctx context.Context, req *BackupRequest) (*BackupResult, error)
 RestoreFiles(ctx context.Context, req *RestoreRequest) (*RestoreResult, error)
 ValidateIntegrity(ctx context.Context, paths []string) ([]*ValidationResult, error)
}

type StorageInfo struct {
 Path           string
 TotalSpace     int64
 FreeSpace      int64
 UsedSpace      int64
 UsagePercent   float64
 FilesystemType string
 ReadOnly       bool
 FileCount      int64
 DirectoryCount int64
}

type CleanupRule struct {
 Pattern       string
 OlderThan     time.Duration
 LargerThan    int64
 SmallerThan   int64
 ExcludePatterns []string
}

type CleanupResult struct {
 FilesProcessed int64
 FilesDeleted   int64
 SpaceFreed     int64
 ProcessingTime time.Duration
 Results        []*FileCleanupResult
}

type FileCleanupResult struct {
 FilePath string
 Action   string
 Reason   string
 Size     int64
 Error    error
}

type BackupRequest struct {
 SourcePaths     []string
 DestinationPath string
 Incremental     bool
 Compress        bool
 VerifyBackup    bool
 ExcludePatterns []string
}

type BackupResult struct {
 BackupID        string
 DestinationPath string
 FilesCopied     int64
 TotalSize       int64
 ProcessingTime  time.Duration
 VerificationPassed bool
 Errors          []string
}

type RestoreRequest struct {
 BackupID        string
 DestinationPath string
 SelectivePaths  []string
 OverwriteExisting bool
}

type RestoreResult struct {
 FilesRestored   int64
 TotalSize       int64
 ProcessingTime  time.Duration
 VerificationPassed bool
 Errors          []string
}

type ValidationResult struct {
 FilePath string
 Valid    bool
 Issues   []*ValidationIssue
 FileInfo *FileInfo
}

type FileInfo struct {
 Path         string
 Size         int64
 ModTime      time.Time
 IsDirectory  bool
 Permissions  string
 Checksum     string
 ContentType  string
}
```

### Step 7: Implementation Guidelines

**Create `docs/implementation-guidelines.md`**:

```markdown
# Service Implementation Guidelines

## Service Architecture Patterns

### 1. Service Structure

Each service should follow this directory structure:
```

pkg/services/ ├── web/ │ ├── server.go # Web service implementation │ ├──
handlers.go # HTTP/gRPC handlers │ ├── middleware.go # Authentication, logging,
etc. │ └── config.go # Service configuration ├── engine/ │ ├── server.go #
Engine service implementation │ ├── translation/ # Translation workers │ ├──
monitoring/ # Monitoring workers │ ├── coordination/ # Task coordination │ └──
leader/ # Leader election └── file/ ├── server.go # File service implementation
├── operations.go # File operations ├── watcher.go # File watching └──
media.go # Media processing

````

### 2. Service Initialization
```go
type ServiceOptions struct {
    Config     ServiceConfig
    Logger     *zap.Logger
    Metrics    prometheus.Registerer
    Tracer     trace.TracerProvider
    Discovery  ServiceDiscovery
}

func NewService(opts ServiceOptions) (Service, error) {
    // Validate configuration
    // Initialize dependencies
    // Setup monitoring
    // Register with discovery
}
````

### 3. Error Handling

- Use structured errors with context
- Implement proper error wrapping
- Return gRPC status codes correctly
- Log errors with appropriate levels

### 4. Configuration Management

- Use environment variables for secrets
- Support configuration hot-reloading
- Validate configurations on startup
- Provide sensible defaults

### 5. Monitoring and Observability

- Implement structured logging
- Export Prometheus metrics
- Add distributed tracing
- Health check endpoints

### 6. Testing Strategy

- Unit tests for business logic
- Integration tests for service interactions
- Contract tests for gRPC interfaces
- End-to-end tests for workflows

## Implementation Priority

1. Core interfaces and common types
2. File service (foundational)
3. Engine service (processing)
4. Web service (user interface)
5. Service discovery and coordination
6. Monitoring and observability
7. Integration testing

```

This completes the comprehensive service interface definitions. All four parts can now be merged into the final TASK-01-001 document.
```
