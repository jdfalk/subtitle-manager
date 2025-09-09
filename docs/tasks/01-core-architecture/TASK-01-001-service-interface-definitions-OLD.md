<!-- file: docs/tasks/01-core-architecture/TASK-01-001-service-interface-definitions.md -->
<!-- version: 1.0.0 -->
<!-- guid: 01001000-1111-2222-3333-444444444444 -->

# TASK-01-001: Service Interface Definitions

## Overview

Define comprehensive gRPC service interfaces using Protocol Buffers Edition 2023 with the 1-1-1 pattern (one top-level entity per file) and opaque API support. This task establishes the foundation for all inter-service communication in the 3-service active-active architecture.

## Requirements

### Core Technology Requirements
- **Protocol Buffers Edition 2023**: Use latest protobuf edition
- **1-1-1 Pattern**: One top-level entity (message/service) per .proto file
- **Opaque API**: Use opaque API with getters/setters instead of direct field access
- **gRPC Services**: All inter-service communication via gRPC
- **Comprehensive Interfaces**: Cover all service operations and data types

### Service Interface Requirements
- **Web Service**: User management, API gateway, file upload/download
- **Engine Service**: Translation, monitoring, coordination, leader election
- **File Service**: File operations, media processing, file watching
- **Common Types**: Shared data structures and error handling

## Implementation Steps

### Step 1: Define Common Protobuf Types

**Create `proto/common/v1/timestamp.proto`**:

```protobuf
// file: proto/common/v1/timestamp.proto
// version: 1.0.0
// guid: common01-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "google/protobuf/timestamp.proto";

// Timestamp wrapper for consistent time handling
message Timestamp {
  google.protobuf.Timestamp value = 1;
}
```

**Create `proto/common/v1/duration.proto`**:

```protobuf
// file: proto/common/v1/duration.proto
// version: 1.0.0
// guid: common02-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "google/protobuf/duration.proto";

// Duration wrapper for consistent duration handling
message Duration {
  google.protobuf.Duration value = 1;
}
```

**Create `proto/common/v1/language.proto`**:

```protobuf
// file: proto/common/v1/language.proto
// version: 1.0.0
// guid: common03-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Language representation with ISO codes
message Language {
  // ISO 639-1 language code (e.g., "en", "fr", "es")
  string code = 1;
  // ISO 639-2 three-letter code (e.g., "eng", "fre", "spa")
  string code_3 = 2;
  // Human-readable language name (e.g., "English", "French", "Spanish")
  string name = 3;
  // Native language name (e.g., "English", "Français", "Español")
  string native_name = 4;
  // Writing direction: "ltr" or "rtl"
  string direction = 5;
}
```

**Create `proto/common/v1/error.proto`**:

```protobuf
// file: proto/common/v1/error.proto
// version: 1.0.0
// guid: common04-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";

// Structured error information
message Error {
  // Error code (e.g., "FILE_NOT_FOUND", "PERMISSION_DENIED")
  string code = 1;
  // Human-readable error message
  string message = 2;
  // Error details for debugging
  string details = 3;
  // When the error occurred
  Timestamp occurred_at = 4;
  // Component that generated the error
  string component = 5;
  // Request ID for tracing
  string request_id = 6;
  // Additional context
  map<string, string> metadata = 7;
}
```

**Create `proto/common/v1/file_info.proto`**:

```protobuf
// file: proto/common/v1/file_info.proto
// version: 1.0.0
// guid: common05-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";

// File system information
message FileInfo {
  // Absolute file path
  string path = 1;
  // File size in bytes
  int64 size = 2;
  // Last modification time
  Timestamp modified_at = 3;
  // Creation time
  Timestamp created_at = 4;
  // File permissions (Unix-style)
  string permissions = 5;
  // File type/MIME type
  string content_type = 6;
  // File checksum (SHA256)
  string checksum = 7;
  // Whether it's a directory
  bool is_directory = 8;
  // Whether it's readable
  bool is_readable = 9;
  // Whether it's writable
  bool is_writable = 10;
  // Additional metadata
  map<string, string> metadata = 11;
}
```

**Create `proto/common/v1/health_check_request.proto`**:

```protobuf
// file: proto/common/v1/health_check_request.proto
// version: 1.0.0
// guid: common06-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/duration.proto";

// Health check request
message HealthCheckRequest {
  // Service name to check (empty for self)
  string service_name = 1;
  // Include detailed status information
  bool include_details = 2;
  // Timeout for health check
  Duration timeout = 3;
}
```

**Create `proto/common/v1/health_check_response.proto`**:

```protobuf
// file: proto/common/v1/health_check_response.proto
// version: 1.0.0
// guid: common07-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";

// Health check response
message HealthCheckResponse {
  // Service status: "SERVING", "NOT_SERVING", "UNKNOWN"
  string status = 1;
  // Service name
  string service_name = 2;
  // Service version
  string version = 3;
  // Uptime since last restart
  int64 uptime_seconds = 4;
  // Last health check time
  Timestamp checked_at = 5;
  // Additional status details
  map<string, string> details = 6;
  // Health check message
  string message = 7;
}
```

**Create `proto/common/v1/service_status.proto`**:

```protobuf
// file: proto/common/v1/service_status.proto
// version: 1.0.0
// guid: common08-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";

// Service status information
message ServiceStatus {
  // Service identifier
  string service_id = 1;
  // Service type: "web", "engine", "file"
  string service_type = 2;
  // Current status: "starting", "running", "stopping", "stopped", "error"
  string status = 3;
  // Service endpoint address
  string address = 4;
  // Service port
  int32 port = 5;
  // When service was started
  Timestamp started_at = 6;
  // Last status update
  Timestamp updated_at = 7;
  // Service capabilities
  repeated string capabilities = 8;
  // Service metadata
  map<string, string> metadata = 9;
}
```

**Create `proto/common/v1/service_metrics.proto`**:

```protobuf
// file: proto/common/v1/service_metrics.proto
// version: 1.0.0
// guid: common09-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";

// Service metrics information
message ServiceMetrics {
  // Service identifier
  string service_id = 1;
  // Metrics collection time
  Timestamp collected_at = 2;
  // CPU usage percentage (0-100)
  double cpu_usage_percent = 3;
  // Memory usage in bytes
  int64 memory_usage_bytes = 4;
  // Number of active connections
  int64 active_connections = 5;
  // Total requests processed
  int64 total_requests = 6;
  // Failed requests count
  int64 failed_requests = 7;
  // Average response time in milliseconds
  double avg_response_time_ms = 8;
  // Custom metrics
  map<string, double> custom_metrics = 9;
}
```

**Create `proto/common/v1/service_event.proto`**:

```protobuf
// file: proto/common/v1/service_event.proto
// version: 1.0.0
// guid: common10-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.common.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/common/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";

// Service event information
message ServiceEvent {
  // Event identifier
  string event_id = 1;
  // Service that generated the event
  string service_id = 2;
  // Event type: "started", "stopped", "error", "warning", "info"
  string event_type = 3;
  // Event severity: "critical", "error", "warning", "info", "debug"
  string severity = 4;
  // Event message
  string message = 5;
  // When the event occurred
  Timestamp occurred_at = 6;
  // Event source component
  string component = 7;
  // Request ID if applicable
  string request_id = 8;
  // Additional event data
  map<string, string> data = 9;
}
```

### Step 2: Define Web Service Interfaces

**Create `proto/services/web/v1/web_service.proto`**:

```protobuf
// file: proto/services/web/v1/web_service.proto
// version: 1.0.0
// guid: web00001-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/health_check_request.proto";
import "proto/common/v1/health_check_response.proto";
import "google/protobuf/empty.proto";

// Web Service handles all web-facing operations
service WebService {
  // User management
  rpc AuthenticateUser(AuthenticateUserRequest) returns (AuthenticateUserResponse);
  rpc RefreshToken(RefreshTokenRequest) returns (RefreshTokenResponse);
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse);
  rpc UpdateUserPreferences(UpdateUserPreferencesRequest) returns (UpdateUserPreferencesResponse);
  rpc LogoutUser(LogoutUserRequest) returns (google.protobuf.Empty);

  // Subtitle management
  rpc UploadSubtitle(UploadSubtitleRequest) returns (UploadSubtitleResponse);
  rpc DownloadSubtitle(DownloadSubtitleRequest) returns (DownloadSubtitleResponse);
  rpc SearchSubtitles(SearchSubtitlesRequest) returns (SearchSubtitlesResponse);
  rpc GetSubtitleMetadata(GetSubtitleMetadataRequest) returns (GetSubtitleMetadataResponse);
  rpc UpdateSubtitleMetadata(UpdateSubtitleMetadataRequest) returns (UpdateSubtitleMetadataResponse);
  rpc DeleteSubtitle(DeleteSubtitleRequest) returns (google.protobuf.Empty);

  // Translation operations
  rpc TranslateSubtitle(TranslateSubtitleRequest) returns (TranslateSubtitleResponse);
  rpc GetTranslationStatus(GetTranslationStatusRequest) returns (GetTranslationStatusResponse);
  rpc CancelTranslation(CancelTranslationRequest) returns (CancelTranslationResponse);
  rpc GetTranslationHistory(GetTranslationHistoryRequest) returns (GetTranslationHistoryResponse);

  // File operations
  rpc UploadFile(stream UploadFileRequest) returns (UploadFileResponse);
  rpc DownloadFile(DownloadFileRequest) returns (stream DownloadFileResponse);
  rpc GetFileInfo(GetFileInfoRequest) returns (GetFileInfoResponse);
  rpc ListFiles(ListFilesRequest) returns (ListFilesResponse);

  // Library management
  rpc GetLibraries(GetLibrariesRequest) returns (GetLibrariesResponse);
  rpc CreateLibrary(CreateLibraryRequest) returns (CreateLibraryResponse);
  rpc UpdateLibrary(UpdateLibraryRequest) returns (UpdateLibraryResponse);
  rpc DeleteLibrary(DeleteLibraryRequest) returns (google.protobuf.Empty);
  rpc ScanLibrary(ScanLibraryRequest) returns (ScanLibraryResponse);
  rpc GetScanStatus(GetScanStatusRequest) returns (GetScanStatusResponse);

  // System operations
  rpc GetSystemStatus(GetSystemStatusRequest) returns (GetSystemStatusResponse);
  rpc UpdateConfiguration(UpdateConfigurationRequest) returns (UpdateConfigurationResponse);
  rpc GetConfiguration(GetConfigurationRequest) returns (GetConfigurationResponse);
  rpc GetLogs(GetLogsRequest) returns (stream GetLogsResponse);

  // Health check
  rpc Health(common.v1.HealthCheckRequest) returns (common.v1.HealthCheckResponse);
}
```
**Create `proto/services/web/v1/authenticate_user_request.proto`**:

```protobuf
// file: proto/services/web/v1/authenticate_user_request.proto
// version: 1.0.0
// guid: web00002-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// AuthenticateUserRequest for user login
message AuthenticateUserRequest {
  // Authentication method
  oneof auth_method {
    UserPasswordAuth password_auth = 1;
    TokenAuth token_auth = 2;
    ApiKeyAuth api_key_auth = 3;
  }
  // Remember login session
  bool remember_me = 4;
  // Client information
  string client_info = 5;
  // IP address for security logging
  string client_ip = 6;
  // Session duration in seconds (0 for default)
  int64 session_duration_seconds = 7;
}
```

**Create `proto/services/web/v1/user_password_auth.proto`**:

```protobuf
// file: proto/services/web/v1/user_password_auth.proto
// version: 1.0.0
// guid: web00003-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Username and password authentication
message UserPasswordAuth {
  string username = 1;
  string password = 2;
}
```

**Create `proto/services/web/v1/token_auth.proto`**:

```protobuf
// file: proto/services/web/v1/token_auth.proto
// version: 1.0.0
// guid: web00004-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Token-based authentication
message TokenAuth {
  string token = 1;
  string token_type = 2; // "jwt", "session", "api"
}
```

**Create `proto/services/web/v1/api_key_auth.proto`**:

```protobuf
// file: proto/services/web/v1/api_key_auth.proto
// version: 1.0.0
// guid: web00005-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// API key authentication
message ApiKeyAuth {
  string api_key = 1;
  string api_secret = 2;
}
```

**Create `proto/services/web/v1/authenticate_user_response.proto`**:

```protobuf
// file: proto/services/web/v1/authenticate_user_response.proto
// version: 1.0.0
// guid: web00006-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";
import "proto/common/v1/error.proto";

// AuthenticateUserResponse with tokens and user info
message AuthenticateUserResponse {
  // Authentication success
  bool success = 1;
  // Access token (JWT)
  string access_token = 2;
  // Refresh token
  string refresh_token = 3;
  // Token expiration time
  common.v1.Timestamp expires_at = 4;
  // User information
  User user = 5;
  // Session ID
  string session_id = 6;
  // User permissions
  repeated string permissions = 7;
  // Authentication error if failed
  common.v1.Error error = 8;
}
```

**Create `proto/services/web/v1/user.proto`**:

```protobuf
// file: proto/services/web/v1/user.proto
// version: 1.0.0
// guid: web00007-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";
import "proto/common/v1/language.proto";

// User account information
message User {
  // User unique identifier
  string user_id = 1;
  // Username for login
  string username = 2;
  // Display name
  string display_name = 3;
  // Email address
  string email = 4;
  // User role: "admin", "user", "viewer"
  string role = 5;
  // Account status: "active", "inactive", "suspended"
  string status = 6;
  // Account creation time
  common.v1.Timestamp created_at = 7;
  // Last login time
  common.v1.Timestamp last_login_at = 8;
  // User preferences
  UserPreferences preferences = 9;
  // User avatar URL
  string avatar_url = 10;
  // Account metadata
  map<string, string> metadata = 11;
}
```

**Create `proto/services/web/v1/user_preferences.proto`**:

```protobuf
// file: proto/services/web/v1/user_preferences.proto
// version: 1.0.0
// guid: web00008-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/language.proto";

// User preferences and settings
message UserPreferences {
  // Preferred UI language
  common.v1.Language ui_language = 1;
  // Default source language for translation
  common.v1.Language default_source_language = 2;
  // Default target language for translation
  common.v1.Language default_target_language = 3;
  // Theme preference: "light", "dark", "auto"
  string theme = 4;
  // Timezone
  string timezone = 5;
  // Date format preference
  string date_format = 6;
  // Time format preference: "12h", "24h"
  string time_format = 7;
  // Email notifications enabled
  bool email_notifications = 8;
  // Desktop notifications enabled
  bool desktop_notifications = 9;
  // Auto-download subtitles
  bool auto_download = 10;
  // Default subtitle format: "srt", "vtt", "ass"
  string default_subtitle_format = 11;
  // Quality preference for downloads: "best", "good", "any"
  string download_quality = 12;
  // Custom preferences
  map<string, string> custom_settings = 13;
}
```

**Create `proto/services/web/v1/refresh_token_request.proto`**:

```protobuf
// file: proto/services/web/v1/refresh_token_request.proto
// version: 1.0.0
// guid: web00009-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// RefreshTokenRequest to get new access token
message RefreshTokenRequest {
  // Refresh token
  string refresh_token = 1;
  // Client information
  string client_info = 2;
}
```

**Create `proto/services/web/v1/refresh_token_response.proto`**:

```protobuf
// file: proto/services/web/v1/refresh_token_response.proto
// version: 1.0.0
// guid: web00010-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";
import "proto/common/v1/error.proto";

// RefreshTokenResponse with new tokens
message RefreshTokenResponse {
  // New access token
  string access_token = 1;
  // New refresh token (optional)
  string refresh_token = 2;
  // Token expiration time
  common.v1.Timestamp expires_at = 3;
  // Token type: "Bearer"
  string token_type = 4;
  // Error if refresh failed
  common.v1.Error error = 5;
}
```

**Create `proto/services/web/v1/get_user_request.proto`**:

```protobuf
// file: proto/services/web/v1/get_user_request.proto
// version: 1.0.0
// guid: web00011-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// GetUserRequest to retrieve user information
message GetUserRequest {
  // User ID to retrieve (empty for current user)
  string user_id = 1;
  // Include user preferences
  bool include_preferences = 2;
  // Include user metadata
  bool include_metadata = 3;
}
```

**Create `proto/services/web/v1/get_user_response.proto`**:

```protobuf
// file: proto/services/web/v1/get_user_response.proto
// version: 1.0.0
// guid: web00012-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/error.proto";

// GetUserResponse with user information
message GetUserResponse {
  // User information
  User user = 1;
  // Error if retrieval failed
  common.v1.Error error = 2;
}
```

**Create `proto/services/web/v1/update_user_request.proto`**:

```protobuf
// file: proto/services/web/v1/update_user_request.proto
// version: 1.0.0
// guid: web00013-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// UpdateUserRequest to modify user information
message UpdateUserRequest {
  // User ID to update (empty for current user)
  string user_id = 1;
  // Updated display name
  string display_name = 2;
  // Updated email address
  string email = 3;
  // Updated avatar URL
  string avatar_url = 4;
  // Fields to update (field mask)
  repeated string update_fields = 5;
  // Updated metadata
  map<string, string> metadata = 6;
}
```

**Create `proto/services/web/v1/update_user_response.proto`**:

```protobuf
// file: proto/services/web/v1/update_user_response.proto
// version: 1.0.0
// guid: web00014-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/error.proto";

// UpdateUserResponse with updated user info
message UpdateUserResponse {
  // Updated user information
  User user = 1;
  // Error if update failed
  common.v1.Error error = 2;
}
```

**Create `proto/services/web/v1/update_user_preferences_request.proto`**:

```protobuf
// file: proto/services/web/v1/update_user_preferences_request.proto
// version: 1.0.0
// guid: web00015-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// UpdateUserPreferencesRequest to modify user preferences
message UpdateUserPreferencesRequest {
  // User ID (empty for current user)
  string user_id = 1;
  // Updated preferences
  UserPreferences preferences = 2;
  // Fields to update (field mask)
  repeated string update_fields = 3;
}
```

**Create `proto/services/web/v1/update_user_preferences_response.proto`**:

```protobuf
// file: proto/services/web/v1/update_user_preferences_response.proto
// version: 1.0.0
// guid: web00016-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/error.proto";

// UpdateUserPreferencesResponse with updated preferences
message UpdateUserPreferencesResponse {
  // Updated user preferences
  UserPreferences preferences = 1;
  // Error if update failed
  common.v1.Error error = 2;
}
```

**Create `proto/services/web/v1/logout_user_request.proto`**:

```protobuf
// file: proto/services/web/v1/logout_user_request.proto
// version: 1.0.0
// guid: web00017-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// LogoutUserRequest to end user session
message LogoutUserRequest {
  // Session ID to invalidate
  string session_id = 1;
  // Invalidate all sessions for user
  bool invalidate_all_sessions = 2;
  // Refresh token to invalidate
  string refresh_token = 3;
}
```
### Step 3: Define Engine Service Interfaces

**Create `proto/services/engine/v1/engine_service.proto`**:

```protobuf
// file: proto/services/engine/v1/engine_service.proto
// version: 1.0.0
// guid: engine01-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/health_check_request.proto";
import "proto/common/v1/health_check_response.proto";
import "google/protobuf/empty.proto";

// Engine Service handles core processing, translation, and coordination
service EngineService {
  // Translation operations
  rpc ProcessTranslation(ProcessTranslationRequest) returns (ProcessTranslationResponse);
  rpc GetTranslationProgress(GetTranslationProgressRequest) returns (GetTranslationProgressResponse);
  rpc CancelTranslation(CancelTranslationRequest) returns (CancelTranslationResponse);
  rpc ListActiveTranslations(ListActiveTranslationsRequest) returns (ListActiveTranslationsResponse);

  // Monitoring operations
  rpc StartMonitoring(StartMonitoringRequest) returns (StartMonitoringResponse);
  rpc StopMonitoring(StopMonitoringRequest) returns (google.protobuf.Empty);
  rpc GetMonitoringStatus(GetMonitoringStatusRequest) returns (GetMonitoringStatusResponse);
  rpc GetMonitoringEvents(GetMonitoringEventsRequest) returns (stream GetMonitoringEventsResponse);
  rpc UpdateMonitoringConfig(UpdateMonitoringConfigRequest) returns (UpdateMonitoringConfigResponse);

  // Processing operations
  rpc ProcessMedia(ProcessMediaRequest) returns (ProcessMediaResponse);
  rpc ExtractSubtitles(ExtractSubtitlesRequest) returns (ExtractSubtitlesResponse);
  rpc ConvertFormat(ConvertFormatRequest) returns (ConvertFormatResponse);
  rpc ValidateSubtitle(ValidateSubtitleRequest) returns (ValidateSubtitleResponse);

  // Coordination operations
  rpc RegisterWorker(RegisterWorkerRequest) returns (RegisterWorkerResponse);
  rpc UnregisterWorker(UnregisterWorkerRequest) returns (google.protobuf.Empty);
  rpc HeartbeatWorker(HeartbeatWorkerRequest) returns (HeartbeatWorkerResponse);
  rpc GetWorkerStatus(GetWorkerStatusRequest) returns (GetWorkerStatusResponse);
  rpc ListWorkers(ListWorkersRequest) returns (ListWorkersResponse);
  rpc AssignTask(AssignTaskRequest) returns (AssignTaskResponse);
  rpc CompleteTask(CompleteTaskRequest) returns (CompleteTaskResponse);
  rpc GetTaskStatus(GetTaskStatusRequest) returns (GetTaskStatusResponse);
  rpc ListTasks(ListTasksRequest) returns (ListTasksResponse);

  // Leader election operations
  rpc RequestLeadership(RequestLeadershipRequest) returns (RequestLeadershipResponse);
  rpc ReleaseLeadership(ReleaseLeadershipRequest) returns (google.protobuf.Empty);
  rpc GetLeaderInfo(GetLeaderInfoRequest) returns (GetLeaderInfoResponse);
  rpc HeartbeatLeader(HeartbeatLeaderRequest) returns (HeartbeatLeaderResponse);
  rpc GetLeadershipStatus(GetLeadershipStatusRequest) returns (GetLeadershipStatusResponse);

  // Health check
  rpc Health(common.v1.HealthCheckRequest) returns (common.v1.HealthCheckResponse);
}
```

**Create `proto/services/engine/v1/process_translation_request.proto`**:

```protobuf
// file: proto/services/engine/v1/process_translation_request.proto
// version: 1.0.0
// guid: engine02-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/language.proto";

// ProcessTranslationRequest to start translation job
message ProcessTranslationRequest {
  // Unique request ID for tracking
  string request_id = 1;
  // Source subtitle file path
  string source_file_path = 2;
  // Target output file path
  string target_file_path = 3;
  // Source language
  common.v1.Language source_language = 4;
  // Target language
  common.v1.Language target_language = 5;
  // Translation engine: "google", "openai", "azure", "aws"
  string translation_engine = 6;
  // Translation quality: "fast", "balanced", "best"
  string quality = 7;
  // Priority: "low", "normal", "high", "urgent"
  string priority = 8;
  // User ID requesting translation
  string user_id = 9;
  // Translation options
  map<string, string> options = 10;
  // Callback URL for completion notification
  string callback_url = 11;
  // Whether to preserve formatting
  bool preserve_formatting = 12;
  // Maximum translation time in seconds
  int64 timeout_seconds = 13;
}
```

**Create `proto/services/engine/v1/process_translation_response.proto`**:

```protobuf
// file: proto/services/engine/v1/process_translation_response.proto
// version: 1.0.0
// guid: engine03-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";
import "proto/common/v1/error.proto";

// ProcessTranslationResponse with job information
message ProcessTranslationResponse {
  // Translation job ID
  string job_id = 1;
  // Request ID that was processed
  string request_id = 2;
  // Job status: "queued", "processing", "completed", "failed", "cancelled"
  string status = 3;
  // Estimated completion time
  common.v1.Timestamp estimated_completion = 4;
  // Assigned worker ID
  string worker_id = 5;
  // Job priority
  string priority = 6;
  // Job creation time
  common.v1.Timestamp created_at = 7;
  // Progress percentage (0-100)
  float progress = 8;
  // Error if job creation failed
  common.v1.Error error = 9;
}
```

**Create `proto/services/engine/v1/translation_progress.proto`**:

```protobuf
// file: proto/services/engine/v1/translation_progress.proto
// version: 1.0.0
// guid: engine04-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";
import "proto/common/v1/duration.proto";

// TranslationProgress tracking information
message TranslationProgress {
  // Translation job ID
  string job_id = 1;
  // Current status
  string status = 2;
  // Progress percentage (0-100)
  float progress = 3;
  // Subtitles processed so far
  int64 subtitles_processed = 4;
  // Total subtitles to process
  int64 total_subtitles = 5;
  // Current step description
  string current_step = 6;
  // Processing start time
  common.v1.Timestamp started_at = 7;
  // Last update time
  common.v1.Timestamp updated_at = 8;
  // Estimated remaining time
  common.v1.Duration estimated_remaining = 9;
  // Worker processing the job
  string worker_id = 10;
  // Progress details
  map<string, string> details = 11;
}
```

**Create `proto/services/engine/v1/worker.proto`**:

```protobuf
// file: proto/services/engine/v1/worker.proto
// version: 1.0.0
// guid: engine05-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";

// Worker registration and status information
message Worker {
  // Unique worker identifier
  string worker_id = 1;
  // Worker type: "translation", "monitoring", "processing"
  string worker_type = 2;
  // Worker status: "idle", "busy", "offline", "error"
  string status = 3;
  // Worker capabilities
  repeated string capabilities = 4;
  // Supported languages for translation workers
  repeated string supported_languages = 5;
  // Maximum concurrent tasks
  int32 max_concurrent_tasks = 6;
  // Currently assigned tasks
  int32 current_tasks = 7;
  // Worker endpoint address
  string endpoint = 8;
  // Worker version
  string version = 9;
  // Registration time
  common.v1.Timestamp registered_at = 10;
  // Last heartbeat time
  common.v1.Timestamp last_heartbeat = 11;
  // Worker metadata
  map<string, string> metadata = 12;
  // Performance metrics
  WorkerMetrics metrics = 13;
}
```

**Create `proto/services/engine/v1/worker_metrics.proto`**:

```protobuf
// file: proto/services/engine/v1/worker_metrics.proto
// version: 1.0.0
// guid: engine06-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/duration.proto";

// Worker performance metrics
message WorkerMetrics {
  // Total tasks completed
  int64 tasks_completed = 1;
  // Total tasks failed
  int64 tasks_failed = 2;
  // Average task completion time
  common.v1.Duration avg_completion_time = 3;
  // CPU usage percentage
  double cpu_usage = 4;
  // Memory usage in bytes
  int64 memory_usage = 5;
  // Network throughput in bytes/second
  double network_throughput = 6;
  // Success rate (0.0 to 1.0)
  double success_rate = 7;
  // Queue length
  int32 queue_length = 8;
}
```

**Create `proto/services/engine/v1/task.proto`**:

```protobuf
// file: proto/services/engine/v1/task.proto
// version: 1.0.0
// guid: engine07-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";
import "proto/common/v1/duration.proto";

// Task definition for worker assignment
message Task {
  // Unique task identifier
  string task_id = 1;
  // Task type: "translation", "monitoring", "processing", "extraction"
  string task_type = 2;
  // Task status: "pending", "assigned", "running", "completed", "failed", "cancelled"
  string status = 3;
  // Task priority: "low", "normal", "high", "urgent"
  string priority = 4;
  // Task data payload
  bytes task_data = 5;
  // Required worker capabilities
  repeated string required_capabilities = 6;
  // Assigned worker ID
  string assigned_worker = 7;
  // Task creation time
  common.v1.Timestamp created_at = 8;
  // Task assignment time
  common.v1.Timestamp assigned_at = 9;
  // Task start time
  common.v1.Timestamp started_at = 10;
  // Task completion time
  common.v1.Timestamp completed_at = 11;
  // Maximum execution time
  common.v1.Duration timeout = 12;
  // Number of retry attempts
  int32 retry_count = 13;
  // Maximum retries allowed
  int32 max_retries = 14;
  // Task metadata
  map<string, string> metadata = 15;
  // User who submitted the task
  string user_id = 16;
}
```

**Create `proto/services/engine/v1/task_result.proto`**:

```protobuf
// file: proto/services/engine/v1/task_result.proto
// version: 1.0.0
// guid: engine08-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";
import "proto/common/v1/duration.proto";
import "proto/common/v1/error.proto";

// Task execution result
message TaskResult {
  // Task identifier
  string task_id = 1;
  // Worker that executed the task
  string worker_id = 2;
  // Execution status: "success", "failure", "timeout", "cancelled"
  string status = 3;
  // Result data payload
  bytes result_data = 4;
  // Execution start time
  common.v1.Timestamp started_at = 5;
  // Execution completion time
  common.v1.Timestamp completed_at = 6;
  // Total execution time
  common.v1.Duration execution_time = 7;
  // Error information if failed
  common.v1.Error error = 8;
  // Result metadata
  map<string, string> metadata = 9;
  // Output files created
  repeated string output_files = 10;
  // Log messages
  repeated string logs = 11;
}
```

**Create `proto/services/engine/v1/leader_info.proto`**:

```protobuf
// file: proto/services/engine/v1/leader_info.proto
// version: 1.0.0
// guid: engine09-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/timestamp.proto";
import "proto/common/v1/duration.proto";

// Leader election information
message LeaderInfo {
  // Current leader ID
  string leader_id = 1;
  // Leader endpoint
  string leader_endpoint = 2;
  // Leadership term
  int64 term = 3;
  // When leadership was acquired
  common.v1.Timestamp elected_at = 4;
  // Leadership lease duration
  common.v1.Duration lease_duration = 5;
  // Last heartbeat time
  common.v1.Timestamp last_heartbeat = 6;
  // Leader status: "active", "inactive", "transitioning"
  string status = 7;
  // Number of workers coordinated
  int32 coordinated_workers = 8;
  // Number of active tasks
  int32 active_tasks = 9;
  // Leadership metadata
  map<string, string> metadata = 10;
}
```

### Step 4: Define File Service Interfaces

**Create `proto/services/file/v1/file_service.proto`**:

```protobuf
// file: proto/services/file/v1/file_service.proto
// version: 1.0.0
// guid: file0001-1111-2222-3333-444444444444

edition = "2023";

package subtitle_manager.file.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/file/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "proto/common/v1/health_check_request.proto";
import "proto/common/v1/health_check_response.proto";
import "google/protobuf/empty.proto";

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
```
### Step 5: Go Service Interface Implementation

**Create `pkg/services/interfaces.go`**:

```go
// file: pkg/services/interfaces.go
// version: 1.0.0
// guid: srv00001-1111-2222-3333-444444444444

package services

import (
    "context"
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

    // User management (using opaque API getters/setters)
    AuthenticateUser(ctx context.Context, req *webv1.AuthenticateUserRequest) (*webv1.AuthenticateUserResponse, error)
    GetUser(ctx context.Context, req *webv1.GetUserRequest) (*webv1.GetUserResponse, error)
    UpdateUserPreferences(ctx context.Context, req *webv1.UpdateUserPreferencesRequest) (*webv1.UpdateUserPreferencesResponse, error)

    // Subtitle management (using opaque API)
    UploadSubtitle(ctx context.Context, req *webv1.UploadSubtitleRequest) (*webv1.UploadSubtitleResponse, error)
    SearchSubtitles(ctx context.Context, req *webv1.SearchSubtitlesRequest) (*webv1.SearchSubtitlesResponse, error)

    // Translation operations (using opaque API)
    TranslateSubtitle(ctx context.Context, req *webv1.TranslateSubtitleRequest) (*webv1.TranslateSubtitleResponse, error)
    GetTranslationStatus(ctx context.Context, req *webv1.GetTranslationStatusRequest) (*webv1.GetTranslationStatusResponse, error)
}

// EngineService handles core processing, translation, and coordination
type EngineService interface {
    Service

    // Translation operations (using opaque API getters/setters)
    ProcessTranslation(ctx context.Context, req *enginev1.ProcessTranslationRequest) (*enginev1.ProcessTranslationResponse, error)
    GetTranslationProgress(ctx context.Context, req *enginev1.GetTranslationProgressRequest) (*enginev1.GetTranslationProgressResponse, error)

    // Worker coordination (using opaque API)
    RegisterWorker(ctx context.Context, req *enginev1.RegisterWorkerRequest) (*enginev1.RegisterWorkerResponse, error)
    AssignTask(ctx context.Context, req *enginev1.AssignTaskRequest) (*enginev1.AssignTaskResponse, error)

    // Leader election (using opaque API)
    RequestLeadership(ctx context.Context, req *enginev1.RequestLeadershipRequest) (*enginev1.RequestLeadershipResponse, error)
    GetLeaderInfo(ctx context.Context, req *enginev1.GetLeaderInfoRequest) (*enginev1.GetLeaderInfoResponse, error)
}

// FileService handles all file system operations
type FileService interface {
    Service

    // Basic file operations (using opaque API getters/setters)
    ReadFile(ctx context.Context, req *filev1.ReadFileRequest) (*filev1.ReadFileResponse, error)
    WriteFile(ctx context.Context, req *filev1.WriteFileRequest) (*filev1.WriteFileResponse, error)

    // Media operations (using opaque API)
    ExtractSubtitles(ctx context.Context, req *filev1.ExtractSubtitlesRequest) (*filev1.ExtractSubtitlesResponse, error)
    GetMediaMetadata(ctx context.Context, req *filev1.GetMediaMetadataRequest) (*filev1.GetMediaMetadataResponse, error)
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
```

### Step 6: Opaque API Usage Examples

When implementing services with the opaque API, use getters and setters instead of direct field access:

```go
// Example: Processing a translation request with opaque API
func (s *engineService) ProcessTranslation(ctx context.Context, req *enginev1.ProcessTranslationRequest) (*enginev1.ProcessTranslationResponse, error) {
    // Use getters to access request fields
    requestID := req.GetRequestId()
    sourceFile := req.GetSourceFilePath()
    targetFile := req.GetTargetFilePath()
    sourceLang := req.GetSourceLanguage()
    targetLang := req.GetTargetLanguage()
    engine := req.GetTranslationEngine()

    // Process translation...
    jobID := generateJobID()

    // Use builder pattern for response
    resp := &enginev1.ProcessTranslationResponse{}
    resp.SetJobId(jobID)
    resp.SetRequestId(requestID)
    resp.SetStatus("queued")
    resp.SetProgress(0.0)

    // Set timestamp using common timestamp wrapper
    now := &commonv1.Timestamp{}
    now.GetValue().SetSeconds(time.Now().Unix())
    resp.SetCreatedAt(now)

    return resp, nil
}

// Example: File operations with opaque API
func (s *fileService) ReadFile(ctx context.Context, req *filev1.ReadFileRequest) (*filev1.ReadFileResponse, error) {
    // Use getters
    filePath := req.GetFilePath()
    offset := req.GetOffset()
    length := req.GetLength()
    encoding := req.GetEncoding()

    // Read file content...
    content, err := s.readFileContent(filePath, offset, length)
    if err != nil {
        // Create error response
        resp := &filev1.ReadFileResponse{}
        errorInfo := &commonv1.Error{}
        errorInfo.SetCode("FILE_READ_ERROR")
        errorInfo.SetMessage(err.Error())
        resp.SetError(errorInfo)
        return resp, nil
    }

    // Create successful response using setters
    resp := &filev1.ReadFileResponse{}
    resp.SetContent(content)
    resp.SetBytesRead(int64(len(content)))
    resp.SetEncoding(encoding)

    // Set file info
    fileInfo := &commonv1.FileInfo{}
    fileInfo.SetPath(filePath)
    fileInfo.SetSize(int64(len(content)))
    resp.SetFileInfo(fileInfo)

    return resp, nil
}

// Example: User authentication with opaque API
func (s *webService) AuthenticateUser(ctx context.Context, req *webv1.AuthenticateUserRequest) (*webv1.AuthenticateUserResponse, error) {
    // Access auth method using oneof getter
    switch authMethod := req.GetAuthMethod().(type) {
    case *webv1.AuthenticateUserRequest_PasswordAuth:
        username := authMethod.PasswordAuth.GetUsername()
        password := authMethod.PasswordAuth.GetPassword()
        // Authenticate with username/password...

    case *webv1.AuthenticateUserRequest_TokenAuth:
        token := authMethod.TokenAuth.GetToken()
        tokenType := authMethod.TokenAuth.GetTokenType()
        // Authenticate with token...

    case *webv1.AuthenticateUserRequest_ApiKeyAuth:
        apiKey := authMethod.ApiKeyAuth.GetApiKey()
        apiSecret := authMethod.ApiKeyAuth.GetApiSecret()
        // Authenticate with API key...
    }

    // Create response using setters
    resp := &webv1.AuthenticateUserResponse{}
    resp.SetSuccess(true)
    resp.SetAccessToken(generatedToken)
    resp.SetRefreshToken(generatedRefreshToken)

    // Set user info
    user := &webv1.User{}
    user.SetUserId(userID)
    user.SetUsername(username)
    user.SetDisplayName(displayName)
    resp.SetUser(user)

    return resp, nil
}
```

### Step 7: Implementation Guidelines

**Service Architecture Patterns**:

1. **Service Structure**: Each service follows this directory structure:
   ```ascii
   pkg/services/
   ├── web/
   │   ├── server.go       # Web service implementation
   │   ├── handlers.go     # HTTP/gRPC handlers
   │   ├── middleware.go   # Authentication, logging, etc.
   │   └── config.go       # Service configuration
   ├── engine/
   │   ├── server.go       # Engine service implementation
   │   ├── translation/    # Translation workers
   │   ├── monitoring/     # Monitoring workers
   │   ├── coordination/   # Task coordination
   │   └── leader/         # Leader election
   └── file/
       ├── server.go       # File service implementation
       ├── operations.go   # File operations
       ├── watcher.go      # File watching
       └── media.go        # Media processing
   ```

2. **Opaque API Best Practices**:
   - Always use getters to access message fields
   - Use setters to populate response messages
   - Handle nil checks when accessing nested messages
   - Use builder patterns for complex message construction
   - Validate required fields using getter methods

3. **Error Handling**:
   - Use structured errors with the common.v1.Error message type
   - Set appropriate error codes and messages using setters
   - Include context and request IDs in error responses
   - Log errors with appropriate levels using structured logging

4. **Configuration Management**:
   - Use environment variables for secrets
   - Support configuration hot-reloading
   - Validate configurations on startup using getter methods
   - Provide sensible defaults for all configuration options

5. **Testing Strategy**:
   - Unit tests for business logic with mock protobuf messages
   - Integration tests for service interactions using real gRPC calls
   - Contract tests for gRPC interfaces with opaque API validation
   - End-to-end tests for complete workflows

### Step 8: Buf Configuration

**Update `buf.gen.yaml`** for Edition 2023 and opaque API:

```yaml
version: v2
plugins:
  - remote: buf.build/protocolbuffers/go
    out: pkg/proto
    opt:
      - paths=source_relative
      - features=pb.go.opaque_api
  - remote: buf.build/grpc/go
    out: pkg/proto
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
```

**Update `buf.yaml`** for Edition 2023:

```yaml
version: v2
modules:
  - path: proto
lint:
  use:
    - STANDARD
  except:
    - PACKAGE_DIRECTORY_MATCH
breaking:
  use:
    - FILE
deps:
  - buf.build/googleapis/googleapis
  - buf.build/protocolbuffers/wellknowntypes
```

This completes the comprehensive service interface definitions with proper Edition 2023 protobuf, 1-1-1 pattern, and opaque API support. All services are designed for the 3-service active-active architecture with proper separation of concerns.
