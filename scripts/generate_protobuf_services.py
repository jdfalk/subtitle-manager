#!/usr/bin/env python3
# file: scripts/generate_protobuf_services.py
# version: 1.1.0
# guid: 3c4d5e6f-7a8b-9c0d-1e2f-3a4b5c6d7e8f

"""
Generate protobuf service definitions for Engine and File services following the Web service pattern.
This script creates comprehensive service definitions with all message types in single files.

CRITICAL PROTOBUF FILE STRUCTURE ORDER:
1. header comments (file, version, guid)
2. edition = "2023";
3. package declaration
4. imports (if any)
5. options (go_package, etc.)
6. service and message definitions

DO NOT PUT IMPORTS AFTER OPTIONS - THIS IS WRONG!
"""

import sys
from pathlib import Path


def create_engine_service_proto() -> str:
    """Generate the Engine service protobuf definition."""
    return """// file: proto/engine/v1/engine_service.proto
// version: 1.0.0
// guid: 4d5e6f7a-8b9c-0d1e-2f3a-4b5c6d7e8f9a

edition = "2023";

package subtitle_manager.engine.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/engine/v1";

// Engine service handles subtitle processing, transcription, and translation
service EngineService {
  // Transcription operations
  rpc TranscribeAudio (TranscribeAudioRequest) returns (TranscribeAudioResponse);
  rpc GetTranscriptionStatus (GetTranscriptionStatusRequest) returns (GetTranscriptionStatusResponse);
  rpc CancelTranscription (CancelTranscriptionRequest) returns (CancelTranscriptionResponse);

  // Translation operations
  rpc TranslateSubtitle (TranslateSubtitleRequest) returns (TranslateSubtitleResponse);
  rpc GetTranslationProgress (GetTranslationProgressRequest) returns (GetTranslationProgressResponse);
  rpc CancelTranslation (CancelTranslationRequest) returns (CancelTranslationResponse);

  // Subtitle processing
  rpc ConvertSubtitle (ConvertSubtitleRequest) returns (ConvertSubtitleResponse);
  rpc ValidateSubtitle (ValidateSubtitleRequest) returns (ValidateSubtitleResponse);
  rpc MergeSubtitles (MergeSubtitlesRequest) returns (MergeSubtitlesResponse);

  // Engine management
  rpc GetEngineStatus (GetEngineStatusRequest) returns (GetEngineStatusResponse);
  rpc HealthCheck (HealthCheckRequest) returns (HealthCheckResponse);
}

// Transcription Messages
message TranscribeAudioRequest {
  string request_id = 1;
  string audio_file_id = 2;
  string source_language = 3;
  repeated string target_languages = 4;
  TranscriptionOptions options = 5;
}

message TranscriptionOptions {
  string model = 1;
  bool speaker_detection = 2;
  bool word_timestamps = 3;
  float confidence_threshold = 4;
}

message TranscribeAudioResponse {
  string job_id = 1;
  string status = 2;
  bool success = 3;
  string error_message = 4;
}

message GetTranscriptionStatusRequest {
  string request_id = 1;
  string job_id = 2;
}

message GetTranscriptionStatusResponse {
  string job_id = 1;
  string status = 2;
  float progress = 3;
  repeated string result_file_ids = 4;
  bool success = 5;
  string error_message = 6;
}

message CancelTranscriptionRequest {
  string request_id = 1;
  string job_id = 2;
}

message CancelTranscriptionResponse {
  string job_id = 1;
  bool cancelled = 2;
  bool success = 3;
  string error_message = 4;
}

// Translation Messages
message TranslateSubtitleRequest {
  string request_id = 1;
  string subtitle_file_id = 2;
  string source_language = 3;
  string target_language = 4;
  TranslationOptions options = 5;
}

message TranslationOptions {
  string model = 1;
  bool preserve_timing = 2;
  bool preserve_formatting = 3;
  map<string, string> custom_settings = 4;
}

message TranslateSubtitleResponse {
  string job_id = 1;
  string status = 2;
  bool success = 3;
  string error_message = 4;
}

message GetTranslationProgressRequest {
  string request_id = 1;
  string job_id = 2;
}

message GetTranslationProgressResponse {
  string job_id = 1;
  string status = 2;
  float progress = 3;
  string result_file_id = 4;
  bool success = 5;
  string error_message = 6;
}

message CancelTranslationRequest {
  string request_id = 1;
  string job_id = 2;
}

message CancelTranslationResponse {
  string job_id = 1;
  bool cancelled = 2;
  bool success = 3;
  string error_message = 4;
}

// Subtitle Processing Messages
message ConvertSubtitleRequest {
  string request_id = 1;
  string source_file_id = 2;
  string target_format = 3;
  ConversionOptions options = 4;
}

message ConversionOptions {
  string encoding = 1;
  int32 framerate = 2;
  bool preserve_metadata = 3;
}

message ConvertSubtitleResponse {
  string result_file_id = 1;
  string format = 2;
  bool success = 3;
  string error_message = 4;
}

message ValidateSubtitleRequest {
  string request_id = 1;
  string file_id = 2;
  ValidationOptions options = 3;
}

message ValidationOptions {
  bool check_timing = 1;
  bool check_encoding = 2;
  bool check_format = 3;
}

message ValidateSubtitleResponse {
  bool is_valid = 1;
  repeated ValidationError errors = 2;
  repeated ValidationWarning warnings = 3;
  bool success = 4;
  string error_message = 5;
}

message ValidationError {
  string code = 1;
  string message = 2;
  int32 line_number = 3;
}

message ValidationWarning {
  string code = 1;
  string message = 2;
  int32 line_number = 3;
}

message MergeSubtitlesRequest {
  string request_id = 1;
  repeated string file_ids = 2;
  MergeOptions options = 3;
}

message MergeOptions {
  string output_format = 1;
  bool preserve_timing = 2;
  string merge_strategy = 3;
}

message MergeSubtitlesResponse {
  string result_file_id = 1;
  bool success = 2;
  string error_message = 3;
}

// Engine Management Messages
message GetEngineStatusRequest {
  string service = 1;
}

message GetEngineStatusResponse {
  string status = 1;
  map<string, string> engines = 2;
  repeated JobStatus active_jobs = 3;
  bool success = 4;
  string error_message = 5;
}

message JobStatus {
  string job_id = 1;
  string type = 2;
  string status = 3;
  float progress = 4;
}

message HealthCheckRequest {
  string service = 1;
}

message HealthCheckResponse {
  string status = 1;
  string message = 2;
  bool success = 3;
}"""


def create_file_service_proto() -> str:
    """Generate the File service protobuf definition."""
    return """// file: proto/file/v1/file_service.proto
// version: 1.0.0
// guid: 5e6f7a8b-9c0d-1e2f-3a4b-5c6d7e8f9a0b

edition = "2023";

package subtitle_manager.file.v1;

option go_package = "github.com/jdfalk/subtitle-manager/pkg/file/v1";

// File service handles file storage, metadata, and operations
service FileService {
  // File management
  rpc UploadFile (stream UploadFileRequest) returns (UploadFileResponse);
  rpc DownloadFile (DownloadFileRequest) returns (stream DownloadFileResponse);
  rpc DeleteFile (DeleteFileRequest) returns (DeleteFileResponse);
  rpc GetFileInfo (GetFileInfoRequest) returns (GetFileInfoResponse);

  // File operations
  rpc CopyFile (CopyFileRequest) returns (CopyFileResponse);
  rpc MoveFile (MoveFileRequest) returns (MoveFileResponse);
  rpc ListFiles (ListFilesRequest) returns (ListFilesResponse);

  // Metadata operations
  rpc UpdateFileMetadata (UpdateFileMetadataRequest) returns (UpdateFileMetadataResponse);
  rpc SearchFiles (SearchFilesRequest) returns (SearchFilesResponse);

  // Storage management
  rpc GetStorageInfo (GetStorageInfoRequest) returns (GetStorageInfoResponse);
  rpc CleanupFiles (CleanupFilesRequest) returns (CleanupFilesResponse);

  // Health check
  rpc HealthCheck (HealthCheckRequest) returns (HealthCheckResponse);
}

// File Upload/Download Messages
message UploadFileRequest {
  oneof request {
    FileMetadata metadata = 1;
    FileChunk chunk = 2;
  }
}

message FileMetadata {
  string filename = 1;
  string content_type = 2;
  int64 total_size = 3;
  string user_id = 4;
  map<string, string> custom_metadata = 5;
}

message FileChunk {
  bytes data = 1;
  int64 offset = 2;
  bool is_last = 3;
}

message UploadFileResponse {
  string file_id = 1;
  string filename = 2;
  int64 size = 3;
  bool success = 4;
  string error_message = 5;
}

message DownloadFileRequest {
  string request_id = 1;
  string file_id = 2;
  int64 offset = 3;
  int64 length = 4;
}

message DownloadFileResponse {
  oneof response {
    FileInfo file_info = 1;
    FileChunk chunk = 2;
  }
}

message FileInfo {
  string filename = 1;
  string content_type = 2;
  int64 total_size = 3;
  map<string, string> metadata = 4;
}

// File Management Messages
message DeleteFileRequest {
  string request_id = 1;
  string file_id = 2;
  bool permanent = 3;
}

message DeleteFileResponse {
  string file_id = 1;
  bool deleted = 2;
  bool success = 3;
  string error_message = 4;
}

message GetFileInfoRequest {
  string request_id = 1;
  string file_id = 2;
  bool include_metadata = 3;
}

message GetFileInfoResponse {
  string file_id = 1;
  string filename = 2;
  string content_type = 3;
  int64 size = 4;
  string created_at = 5;
  string modified_at = 6;
  string user_id = 7;
  map<string, string> metadata = 8;
  bool success = 9;
  string error_message = 10;
}

// File Operations Messages
message CopyFileRequest {
  string request_id = 1;
  string source_file_id = 2;
  string destination_name = 3;
  string user_id = 4;
}

message CopyFileResponse {
  string new_file_id = 1;
  string filename = 2;
  bool success = 3;
  string error_message = 4;
}

message MoveFileRequest {
  string request_id = 1;
  string file_id = 2;
  string new_name = 3;
  string new_location = 4;
}

message MoveFileResponse {
  string file_id = 1;
  string new_filename = 2;
  bool success = 3;
  string error_message = 4;
}

message ListFilesRequest {
  string request_id = 1;
  string user_id = 2;
  string path = 3;
  int32 limit = 4;
  int32 offset = 5;
  ListOptions options = 6;
}

message ListOptions {
  bool include_metadata = 1;
  string sort_by = 2;
  string sort_order = 3;
  repeated string file_types = 4;
}

message ListFilesResponse {
  repeated FileEntry files = 1;
  int32 total_count = 2;
  bool has_more = 3;
  bool success = 4;
  string error_message = 5;
}

message FileEntry {
  string file_id = 1;
  string filename = 2;
  string content_type = 3;
  int64 size = 4;
  string created_at = 5;
  string modified_at = 6;
  map<string, string> metadata = 7;
}

// Metadata Operations Messages
message UpdateFileMetadataRequest {
  string request_id = 1;
  string file_id = 2;
  map<string, string> metadata = 3;
  bool replace_all = 4;
}

message UpdateFileMetadataResponse {
  string file_id = 1;
  map<string, string> updated_metadata = 2;
  bool success = 3;
  string error_message = 4;
}

message SearchFilesRequest {
  string request_id = 1;
  string user_id = 2;
  string query = 3;
  SearchOptions options = 4;
}

message SearchOptions {
  repeated string file_types = 1;
  string date_from = 2;
  string date_to = 3;
  int64 min_size = 4;
  int64 max_size = 5;
  map<string, string> metadata_filters = 6;
  int32 limit = 7;
  int32 offset = 8;
}

message SearchFilesResponse {
  repeated FileEntry files = 1;
  int32 total_count = 2;
  bool success = 3;
  string error_message = 4;
}

// Storage Management Messages
message GetStorageInfoRequest {
  string user_id = 1;
}

message GetStorageInfoResponse {
  int64 total_space = 1;
  int64 used_space = 2;
  int64 available_space = 3;
  int32 file_count = 4;
  map<string, int64> file_type_distribution = 5;
  bool success = 6;
  string error_message = 7;
}

message CleanupFilesRequest {
  string user_id = 1;
  CleanupOptions options = 2;
}

message CleanupOptions {
  int32 days_old = 1;
  bool remove_temporary = 2;
  bool remove_orphaned = 3;
  repeated string file_types = 4;
}

message CleanupFilesResponse {
  int32 files_deleted = 1;
  int64 space_freed = 2;
  bool success = 3;
  string error_message = 4;
}

// Health Check Messages
message HealthCheckRequest {
  string service = 1;
}

message HealthCheckResponse {
  string status = 1;
  string message = 2;
  StorageHealth storage_health = 3;
  bool success = 4;
}

message StorageHealth {
  int64 total_space = 1;
  int64 available_space = 2;
  float usage_percentage = 3;
  bool healthy = 4;
}"""


def create_go_interfaces(service_name: str, package_name: str) -> str:
    """Generate Go interface definitions for a service."""
    return f"""// file: pkg/services/{service_name}_interfaces.go
// version: 1.0.0
// guid: {generate_guid()}

package services

import (
	"context"
	"time"

	"github.com/jdfalk/gcommon/sdks/go/v1/common"
	{service_name}v1 "github.com/jdfalk/subtitle-manager/pkg/{service_name}/v1"
)

// {service_name.title()}ServiceInterface defines the interface for the {service_name} service layer
// This bridges gRPC protobuf messages with gcommon SDK types for business logic
type {service_name.title()}ServiceInterface interface {{
	// Health check
	HealthCheck(ctx context.Context, req *{service_name}v1.HealthCheckRequest) (*{service_name}v1.HealthCheckResponse, error)

	// Add specific methods based on service type here
	// This interface will be expanded with actual service operations
}}"""


def generate_guid() -> str:
    """Generate a simple GUID for demonstration."""
    import uuid

    return str(uuid.uuid4())


def main():
    """Main script function."""
    if len(sys.argv) < 2:
        print("Usage: python generate_protobuf_services.py <output_directory>")
        sys.exit(1)

    output_dir = Path(sys.argv[1])

    # Create Engine service files
    engine_proto_dir = output_dir / "proto" / "engine" / "v1"
    engine_proto_dir.mkdir(parents=True, exist_ok=True)

    with open(engine_proto_dir / "engine_service.proto", "w") as f:
        f.write(create_engine_service_proto())

    print(
        f"✅ Created Engine service protobuf: {engine_proto_dir / 'engine_service.proto'}"
    )

    # Create File service files
    file_proto_dir = output_dir / "proto" / "file" / "v1"
    file_proto_dir.mkdir(parents=True, exist_ok=True)

    with open(file_proto_dir / "file_service.proto", "w") as f:
        f.write(create_file_service_proto())

    print(f"✅ Created File service protobuf: {file_proto_dir / 'file_service.proto'}")

    # Create Go interface stubs
    services_dir = output_dir / "pkg" / "services"
    services_dir.mkdir(parents=True, exist_ok=True)

    with open(services_dir / "engine_interfaces.go", "w") as f:
        f.write(create_go_interfaces("engine", "engine"))

    with open(services_dir / "file_interfaces.go", "w") as f:
        f.write(create_go_interfaces("file", "file"))

    print(f"✅ Created Go interface stubs in: {services_dir}")

    print("\n🎯 Next steps:")
    print("1. Run 'buf generate' to create Go protobuf code")
    print("2. Update the Go interface files with actual service methods")
    print("3. Implement service business logic using gcommon SDK integration")


if __name__ == "__main__":
    main()
