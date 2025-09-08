<!-- file: docs/tasks/01-core-architecture/TASK-01-001-service-interface-definitions-PART1.md -->
<!-- version: 1.0.0 -->
<!-- guid: 01001001-1111-2222-3333-444444444444 -->

# TASK-01-001: Service Interface Definitions (PART 1)

## 🎯 Task Overview

**Primary Objective**: Design and implement the core service interfaces for the 3-service active-active architecture with gRPC communication protocols.

**Task Type**: Core Architecture - Interface Design
**Estimated Effort**: 12-16 hours
**Prerequisites**: Architecture overview understanding
**Part**: 1 of 4 (Interface Design)

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

This completes Part 1 of the service interface definitions. The next parts will cover:
- Part 2: Engine Service Interface
- Part 3: File Service Interface  
- Part 4: Go Interface Implementations

Each part will be detailed with complete code examples, testing procedures, and integration guidelines.
