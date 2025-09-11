<!-- file: docs/tasks/01-core-architecture/TASK-01-001-service-interface-definitions.md -->
<!-- version: 2.0.0 -->
<!-- guid: 01001000-1111-2222-3333-444444444444 -->

# TASK-01-001: Service Interface Definitions (gcommon Edition)

## Overview

Define complete service interface definitions for the 3-service architecture
using Edition 2023 protobuf with comprehensive gcommon integration. This
implementation leverages the extensive gcommon protobuf library for common
types, configuration, authentication, health monitoring, and media processing
instead of defining custom types.

## Requirements

### Core Technology Requirements

- **Edition 2023 Protobuf**: Latest protobuf edition with enhanced features
- **Opaque API**: All protobuf access via getters/setters (no direct field
  access)
- **gcommon Integration**: Extensive use of gcommon protobuf types
- **1-1-1 Pattern**: One top-level entity per protobuf file
- **3-Service Architecture**: Web, Engine, File services with clear boundaries

### gcommon Dependencies

This implementation leverages these gcommon packages:

- **gcommon/common**: User, Session, Error, Metadata, Authentication types
- **gcommon/config**: Configuration management and application settings
- **gcommon/health**: Health monitoring and status reporting
- **gcommon/media**: Media processing and subtitle/video handling
- **gcommon/metrics**: Performance monitoring and observability

## Implementation Steps

### Step 1: Define Web Service Interface

**Create `proto/web/v1/web_service.proto`**:

```protobuf
// file: proto/web/v1/web_service.proto
// version: 2.0.0
// guid: web01000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";
import "google/protobuf/empty.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Import gcommon types instead of defining custom ones
import "gcommon/v1/common/user.proto";
import "gcommon/v1/common/session.proto";
import "gcommon/v1/common/error.proto";
import "gcommon/v1/common/metadata.proto";
import "gcommon/v1/health/health_check.proto";

// Web Service - handles all client-facing operations
service WebService {
  // Authentication operations using gcommon types
  rpc AuthenticateUser(AuthenticateUserRequest) returns (AuthenticateUserResponse);
  rpc RefreshToken(RefreshTokenRequest) returns (RefreshTokenResponse);
  rpc LogoutUser(LogoutUserRequest) returns (google.protobuf.Empty);

  // User management using gcommon User types
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse);
  rpc UpdateUserPreferences(UpdateUserPreferencesRequest) returns (UpdateUserPreferencesResponse);

  // File operations
  rpc UploadSubtitle(UploadSubtitleRequest) returns (UploadSubtitleResponse);
  rpc DownloadSubtitle(DownloadSubtitleRequest) returns (DownloadSubtitleResponse);
  rpc SearchSubtitles(SearchSubtitlesRequest) returns (SearchSubtitlesResponse);

  // Translation operations
  rpc TranslateSubtitle(TranslateSubtitleRequest) returns (TranslateSubtitleResponse);
  rpc GetTranslationStatus(GetTranslationStatusRequest) returns (GetTranslationStatusResponse);
  rpc CancelTranslation(CancelTranslationRequest) returns (CancelTranslationResponse);

  // Streaming operations
  rpc UploadFile(stream UploadFileRequest) returns (UploadFileResponse);
  rpc DownloadFile(DownloadFileRequest) returns (stream DownloadFileResponse);

  // Health check using gcommon health types
  rpc HealthCheck(gcommon.v1.health.HealthCheckRequest) returns (gcommon.v1.health.HealthCheckResponse);
}
```

**Create `proto/web/v1/authenticate_user_request.proto`**:

```protobuf
// file: proto/web/v1/authenticate_user_request.proto
// version: 2.0.0
// guid: web02000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Import gcommon metadata instead of defining custom types
import "gcommon/v1/common/metadata.proto";

// Authentication request using gcommon patterns
message AuthenticateUserRequest {
  // Request metadata using gcommon types
  gcommon.v1.common.Metadata metadata = 1;

  // Authentication credentials
  string username = 2;
  string password = 3;

  // Optional: Remember this session
  bool remember_me = 4;

  // Optional: Two-factor authentication code
  string two_factor_code = 5;
}
```

**Create `proto/web/v1/authenticate_user_response.proto`**:

```protobuf
// file: proto/web/v1/authenticate_user_response.proto
// version: 2.0.0
// guid: web03000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Import gcommon types for standardized responses
import "gcommon/v1/common/user.proto";
import "gcommon/v1/common/session.proto";
import "gcommon/v1/common/error.proto";

// Authentication response using gcommon User and Session types
message AuthenticateUserResponse {
  // Success case: user and session from gcommon
  gcommon.v1.common.User user = 1;
  gcommon.v1.common.Session session = 2;

  // Authentication tokens
  string access_token = 3;
  string refresh_token = 4;

  // Token expiration info
  int64 expires_in = 5;

  // Error case: standardized error from gcommon
  gcommon.v1.common.Error error = 6;
}
```

**Create `proto/web/v1/get_user_request.proto`**:

```protobuf
// file: proto/web/v1/get_user_request.proto
// version: 2.0.0
// guid: web04000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

import "gcommon/v1/common/metadata.proto";

// Get user request with gcommon metadata
message GetUserRequest {
  // Request metadata from gcommon
  gcommon.v1.common.Metadata metadata = 1;

  // User ID to retrieve (empty for current user)
  string user_id = 2;

  // Include user preferences in response
  bool include_preferences = 3;
}
```

**Create `proto/web/v1/get_user_response.proto`**:

```protobuf
// file: proto/web/v1/get_user_response.proto
// version: 2.0.0
// guid: web05000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.web.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Import gcommon types for user management
import "gcommon/v1/common/user.proto";
import "gcommon/v1/common/error.proto";

// Get user response using gcommon User type
message GetUserResponse {
  // User data from gcommon (includes all standard fields)
  gcommon.v1.common.User user = 1;

  // Optional: User preferences (if requested)
  UserPreferences preferences = 2;

  // Error case using gcommon Error
  gcommon.v1.common.Error error = 3;
}

// User preferences (service-specific, not in gcommon)
message UserPreferences {
  // UI preferences
  string language = 1;
  string theme = 2;
  string timezone = 3;

  // Subtitle preferences
  string default_subtitle_language = 4;
  string preferred_subtitle_format = 5;
  bool auto_download_subtitles = 6;

  // Notification preferences
  bool email_notifications = 7;
  bool push_notifications = 8;

  // Advanced preferences
  map<string, string> custom_settings = 9;
}
```

### Step 2: Define Engine Service Interface

**Create `proto/engine/v1/engine_service.proto`**:

```protobuf
// file: proto/engine/v1/engine_service.proto
// version: 2.0.0
// guid: engine01000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";
import "google/protobuf/empty.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Import gcommon types for consistent patterns
import "gcommon/v1/common/metadata.proto";
import "gcommon/v1/common/error.proto";
import "gcommon/v1/media/media_file.proto";
import "gcommon/v1/health/health_check.proto";

// Engine Service - handles translation and processing operations
service EngineService {
  // Translation operations using gcommon media types
  rpc ProcessTranslation(ProcessTranslationRequest) returns (ProcessTranslationResponse);
  rpc GetTranslationProgress(GetTranslationProgressRequest) returns (GetTranslationProgressResponse);
  rpc CancelTranslation(CancelTranslationRequest) returns (CancelTranslationResponse);
  rpc ListActiveTranslations(ListActiveTranslationsRequest) returns (ListActiveTranslationsResponse);

  // Transcription operations using gcommon media types
  rpc ProcessTranscription(ProcessTranscriptionRequest) returns (ProcessTranscriptionResponse);
  rpc GetTranscriptionProgress(GetTranscriptionProgressRequest) returns (GetTranscriptionProgressResponse);

  // Worker management
  rpc GetWorkerStatus(GetWorkerStatusRequest) returns (GetWorkerStatusResponse);
  rpc ScaleWorkers(ScaleWorkersRequest) returns (ScaleWorkersResponse);

  // Health check using gcommon health types
  rpc HealthCheck(gcommon.v1.health.HealthCheckRequest) returns (gcommon.v1.health.HealthCheckResponse);
}
```

**Create `proto/engine/v1/process_translation_request.proto`**:

```protobuf
// file: proto/engine/v1/process_translation_request.proto
// version: 2.0.0
// guid: engine02000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Import gcommon types for media processing
import "gcommon/v1/common/metadata.proto";
import "gcommon/v1/media/media_file.proto";
import "gcommon/v1/media/language.proto";

// Translation processing request using gcommon media types
message ProcessTranslationRequest {
  // Request metadata from gcommon
  gcommon.v1.common.Metadata metadata = 1;

  // Unique request identifier
  string request_id = 2;

  // Source file using gcommon media types
  gcommon.v1.media.MediaFile source_file = 3;

  // Language settings using gcommon media types
  gcommon.v1.media.Language source_language = 4;
  gcommon.v1.media.Language target_language = 5;

  // Processing options
  TranslationOptions options = 6;

  // Priority level
  string priority = 7;

  // User ID for tracking
  string user_id = 8;
}

// Translation options (service-specific)
message TranslationOptions {
  // Translation engine to use
  string engine = 1;

  // Quality level
  string quality = 2;

  // Custom model or settings
  string model = 3;

  // Post-processing options
  bool apply_formatting = 4;
  bool preserve_timing = 5;
  bool auto_correct = 6;

  // Additional parameters
  map<string, string> custom_parameters = 7;
}
```

### Step 3: Define File Service Interface

**Create `proto/file/v1/file_service.proto`**:

```protobuf
// file: proto/file/v1/file_service.proto
// version: 2.0.0
// guid: file01000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.file.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/file/v1";

import "google/protobuf/go_features.proto";
import "google/protobuf/empty.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Import gcommon types for file operations
import "gcommon/v1/common/metadata.proto";
import "gcommon/v1/common/error.proto";
import "gcommon/v1/media/media_file.proto";
import "gcommon/v1/health/health_check.proto";

// File Service - handles all file system operations
service FileService {
  // Basic file operations using gcommon media types
  rpc WriteFile(WriteFileRequest) returns (WriteFileResponse);
  rpc ReadFile(ReadFileRequest) returns (ReadFileResponse);
  rpc DeleteFile(DeleteFileRequest) returns (DeleteFileResponse);
  rpc ListFiles(ListFilesRequest) returns (ListFilesResponse);
  rpc GetFileInfo(GetFileInfoRequest) returns (GetFileInfoResponse);

  // Directory operations
  rpc CreateDirectory(CreateDirectoryRequest) returns (CreateDirectoryResponse);
  rpc DeleteDirectory(DeleteDirectoryRequest) returns (DeleteDirectoryResponse);

  // File watching and monitoring
  rpc WatchFiles(WatchFilesRequest) returns (stream FileWatchEvent);
  rpc StopWatching(StopWatchingRequest) returns (google.protobuf.Empty);

  // Batch operations
  rpc BatchFileOperations(BatchFileOperationsRequest) returns (BatchFileOperationsResponse);

  // Health check using gcommon health types
  rpc HealthCheck(gcommon.v1.health.HealthCheckRequest) returns (gcommon.v1.health.HealthCheckResponse);
}
```

**Create `proto/file/v1/write_file_request.proto`**:

```protobuf
// file: proto/file/v1/write_file_request.proto
// version: 2.0.0
// guid: file02000-2222-3333-4444-555555555555

edition = "2023";

package subtitle_manager.file.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/file/v1";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// Import gcommon types for file operations
import "gcommon/v1/common/metadata.proto";
import "gcommon/v1/media/media_file.proto";

// Write file request using gcommon patterns
message WriteFileRequest {
  // Request metadata from gcommon
  gcommon.v1.common.Metadata metadata = 1;

  // File path (absolute)
  string file_path = 2;

  // File content
  bytes content = 3;

  // Write options
  bool create_directories = 4;
  bool overwrite = 5;

  // File metadata using gcommon media types when applicable
  map<string, string> file_metadata = 6;

  // Permissions (Unix-style)
  string permissions = 7;
}
```

### Step 4: Create Go Interface Definitions

**Create `pkg/services/interfaces.go`**:

```go
// file: pkg/services/interfaces.go
// version: 2.0.0
// guid: interfaces-2222-3333-4444-555555555555

package services

import (
    "context"
    "io"

    // Import gcommon types instead of custom ones
    "github.com/jdfalk/gcommon/sdks/go/v1/common"
    "github.com/jdfalk/gcommon/sdks/go/v1/config"
    "github.com/jdfalk/gcommon/sdks/go/v1/health"
    "github.com/jdfalk/gcommon/sdks/go/v1/media"

    // Our generated protobuf types
    webv1 "github.com/jdfalk/subtitle-manager/pkg/proto/web/v1"
    enginev1 "github.com/jdfalk/subtitle-manager/pkg/proto/engine/v1"
    filev1 "github.com/jdfalk/subtitle-manager/pkg/proto/file/v1"
)

// WebServiceInterface defines the Web Service contract
type WebServiceInterface interface {
    // Authentication using gcommon User and Session types
    AuthenticateUser(ctx context.Context, username, password string) (*common.User, *common.Session, error)
    ValidateSession(ctx context.Context, sessionToken string) (*common.Session, error)
    RefreshToken(ctx context.Context, refreshToken string) (*common.Session, error)
    LogoutUser(ctx context.Context, sessionToken string) error

    // User management using gcommon User types
    GetUser(ctx context.Context, userID string) (*common.User, error)
    UpdateUser(ctx context.Context, userID string, updates *webv1.UpdateUserRequest) (*common.User, error)
    UpdateUserPreferences(ctx context.Context, userID string, prefs *webv1.UserPreferences) (*webv1.UserPreferences, error)

    // File operations
    UploadSubtitle(ctx context.Context, filename string, content []byte, metadata map[string]string) (*webv1.UploadSubtitleResponse, error)
    DownloadSubtitle(ctx context.Context, subtitleID string) (*webv1.DownloadSubtitleResponse, error)

    // Translation operations
    RequestTranslation(ctx context.Context, req *webv1.TranslateSubtitleRequest) (*webv1.TranslateSubtitleResponse, error)
    GetTranslationStatus(ctx context.Context, jobID string) (*webv1.GetTranslationStatusResponse, error)

    // Health check using gcommon health types
    HealthCheck(ctx context.Context) (*health.HealthCheckResponse, error)
}

// EngineServiceInterface defines the Engine Service contract
type EngineServiceInterface interface {
    // Translation processing using gcommon media types
    ProcessTranslation(ctx context.Context, req *enginev1.ProcessTranslationRequest) (*enginev1.ProcessTranslationResponse, error)
    GetTranslationProgress(ctx context.Context, jobID string) (*enginev1.GetTranslationProgressResponse, error)
    CancelTranslation(ctx context.Context, jobID string) error

    // Transcription using gcommon media types
    ProcessTranscription(ctx context.Context, mediaFile *media.MediaFile, options *enginev1.TranscriptionOptions) (*enginev1.ProcessTranscriptionResponse, error)

    // Worker management
    GetWorkerStatus(ctx context.Context) (*enginev1.GetWorkerStatusResponse, error)
    ScaleWorkers(ctx context.Context, targetCount int32) error

    // Health check using gcommon health types
    HealthCheck(ctx context.Context) (*health.HealthCheckResponse, error)
}

// FileServiceInterface defines the File Service contract
type FileServiceInterface interface {
    // File operations using gcommon media types where applicable
    WriteFile(ctx context.Context, path string, content []byte, options *filev1.WriteOptions) (*media.MediaFile, error)
    ReadFile(ctx context.Context, path string) ([]byte, *media.MediaFile, error)
    DeleteFile(ctx context.Context, path string) error

    // Directory operations
    CreateDirectory(ctx context.Context, path string, permissions string) error
    ListFiles(ctx context.Context, directory string, filter *filev1.FileFilter) ([]*media.MediaFile, error)

    // File monitoring
    WatchFiles(ctx context.Context, paths []string) (<-chan *filev1.FileWatchEvent, error)
    StopWatching(ctx context.Context, watchID string) error

    // Health check using gcommon health types
    HealthCheck(ctx context.Context) (*health.HealthCheckResponse, error)
}

// ConfigurationManager handles application configuration using gcommon config types
type ConfigurationManager interface {
    // Configuration using gcommon config types
    GetApplicationConfig(ctx context.Context) (*config.ApplicationConfig, error)
    UpdateApplicationConfig(ctx context.Context, cfg *config.ApplicationConfig) error

    // Service-specific configuration
    GetWebServiceConfig(ctx context.Context) (*WebServiceConfig, error)
    GetEngineServiceConfig(ctx context.Context) (*EngineServiceConfig, error)
    GetFileServiceConfig(ctx context.Context) (*FileServiceConfig, error)

    // Configuration validation
    ValidateConfig(ctx context.Context, cfg interface{}) error

    // Configuration watching
    WatchConfigChanges(ctx context.Context) (<-chan *config.ConfigChangeEvent, error)
}

// AuthenticationManager handles user authentication using gcommon common types
type AuthenticationManager interface {
    // Authentication using gcommon User and Session types
    Authenticate(ctx context.Context, credentials *AuthCredentials) (*common.User, *common.Session, error)
    ValidateToken(ctx context.Context, token string) (*common.User, error)
    RefreshSession(ctx context.Context, refreshToken string) (*common.Session, error)
    RevokeSession(ctx context.Context, sessionToken string) error

    // User management using gcommon User types
    CreateUser(ctx context.Context, user *common.User, password string) (*common.User, error)
    GetUser(ctx context.Context, userID string) (*common.User, error)
    UpdateUser(ctx context.Context, userID string, updates *common.User) (*common.User, error)
    DeleteUser(ctx context.Context, userID string) error

    // Session management using gcommon Session types
    ListUserSessions(ctx context.Context, userID string) ([]*common.Session, error)
    RevokeUserSessions(ctx context.Context, userID string) error
}

// MediaProcessor handles media processing using gcommon media types
type MediaProcessor interface {
    // Media processing using gcommon media types
    ProcessMediaFile(ctx context.Context, file *media.MediaFile, options *ProcessingOptions) (*media.MediaFile, error)
    ExtractSubtitles(ctx context.Context, videoFile *media.MediaFile) ([]*media.MediaFile, error)
    ConvertSubtitleFormat(ctx context.Context, subtitle *media.MediaFile, targetFormat string) (*media.MediaFile, error)

    // Language detection using gcommon media types
    DetectLanguage(ctx context.Context, subtitle *media.MediaFile) (*media.Language, float32, error)

    // Translation using gcommon media types
    TranslateSubtitle(ctx context.Context, subtitle *media.MediaFile, sourceLanguage, targetLanguage *media.Language, options *TranslationOptions) (*media.MediaFile, error)
}

// Supporting types that extend gcommon types

// AuthCredentials extends gcommon authentication patterns
type AuthCredentials struct {
    Username        string
    Password        string
    TwoFactorCode   string
    RememberMe      bool
}

// WebServiceConfig extends gcommon config patterns
type WebServiceConfig struct {
    // Embed gcommon application config
    *config.ApplicationConfig

    // Service-specific settings
    Server          *ServerConfig
    Authentication  *AuthConfig
    RateLimit       *RateLimitConfig
    CORS           *CORSConfig
}

// ServerConfig for web service
type ServerConfig struct {
    GRPCPort     int
    HTTPPort     int
    TLSEnabled   bool
    CertFile     string
    KeyFile      string
}

// AuthConfig for authentication
type AuthConfig struct {
    JWTSecret      string
    JWTExpiration  int64
    SessionTimeout int64
}

// EngineServiceConfig for engine service
type EngineServiceConfig struct {
    // Embed gcommon application config
    *config.ApplicationConfig

    // Engine-specific settings
    Workers         *WorkerConfig
    Translation     *TranslationConfig
    Transcription   *TranscriptionConfig
}

// FileServiceConfig for file service
type FileServiceConfig struct {
    // Embed gcommon application config
    *config.ApplicationConfig

    // File service settings
    Storage         *StorageConfig
    Monitoring      *MonitoringConfig
    Backup          *BackupConfig
}

// ProcessingOptions for media processing
type ProcessingOptions struct {
    Quality         string
    Format          string
    CustomSettings  map[string]string
}

// TranslationOptions for translation processing
type TranslationOptions struct {
    Engine          string
    Quality         string
    Model           string
    PreserveFormat  bool
    AutoCorrect     bool
}
```

### Step 5: Implementation Summary

This service interface definition provides:

1. **Complete gcommon Integration**: Leverages gcommon types for User, Session,
   Error, Media, Config, and Health
2. **Edition 2023 Protobuf**: Modern protobuf with opaque API support
3. **3-Service Architecture**: Clear separation of concerns with gcommon
   consistency
4. **Comprehensive Interfaces**: Go interfaces that use gcommon types throughout
5. **Configuration Management**: Uses gcommon config types with service-specific
   extensions

**Key gcommon Leveraging**:

- **Authentication**: `gcommon.User`, `gcommon.Session` instead of custom types
- **Error Handling**: `gcommon.Error` for standardized error responses
- **Media Processing**: `gcommon.media.MediaFile`, `gcommon.media.Language` for
  media operations
- **Configuration**: `gcommon.config.ApplicationConfig` as base for all service
  configs
- **Health Monitoring**: `gcommon.health.HealthCheck` for service health
- **Metadata**: `gcommon.common.Metadata` for request/response metadata

**Next Implementation**:

- TASK-02-001: Web Service Implementation (using these gcommon-based interfaces)
- TASK-03-001: Engine Service Implementation (with gcommon media processing)
- TASK-04-001: File Service Implementation (with gcommon file handling)

This foundation ensures all services use consistent gcommon types and patterns
while maintaining the opaque API throughout the system.
