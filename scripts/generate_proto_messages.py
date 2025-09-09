#!/usr/bin/env python3
# file: scripts/generate_proto_messages.py
# version: 2.0.0
# guid: script-proto-gen-555555555555

"""
Generate basic protobuf message files for subtitle-manager.
Creates simple message definitions that will use gcommon types in Go implementation.
"""

import os
from pathlib import Path


def create_proto_file(
    file_path: str, package: str, message_name: str, fields: list, guid: str
):
    """Create a basic protobuf message file"""
    content = f"""// file: {file_path}
// version: 2.0.0
// guid: {guid}

edition = "2023";

package {package};

option go_package = "github.com/jdfalk/subtitle-manager/pkg/proto/{package.replace(".", "/")}";

import "google/protobuf/go_features.proto";

option features.(pb.go).api_level = API_OPAQUE;

// {message_name} - Go implementation will use gcommon types internally
message {message_name} {{
{fields}
}}
"""

    # Create directory if it doesn't exist
    Path(file_path).parent.mkdir(parents=True, exist_ok=True)

    with open(file_path, "w") as f:
        f.write(content)

    print(f"Created {file_path}")


def main():
    # Web service messages
    web_messages = [
        # User management
        {
            "file": "proto/web/v1/get_user_response.proto",
            "message": "GetUserResponse",
            "fields": """  // User information (Go implementation populates from gcommon.User)
  string user_id = 1;
  string username = 2;
  string email = 3;
  bool success = 4;
  string error_message = 5;""",
            "guid": "web10000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/update_user_request.proto",
            "message": "UpdateUserRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // User ID to update
  string user_id = 2;

  // Fields to update
  string username = 3;
  string email = 4;
  map<string, string> metadata = 5;""",
            "guid": "web11000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/update_user_response.proto",
            "message": "UpdateUserResponse",
            "fields": """  // Updated user info (populated from gcommon.User)
  string user_id = 1;
  string username = 2;
  string email = 3;
  bool success = 4;
  string error_message = 5;""",
            "guid": "web12000-2222-3333-4444-555555555555",
        },
        # User preferences
        {
            "file": "proto/web/v1/update_user_preferences_request.proto",
            "message": "UpdateUserPreferencesRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // User ID
  string user_id = 2;

  // Preferences to update
  string language = 3;
  string theme = 4;
  string timezone = 5;
  map<string, string> custom_settings = 6;""",
            "guid": "web13000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/update_user_preferences_response.proto",
            "message": "UpdateUserPreferencesResponse",
            "fields": """  // Updated preferences
  string language = 1;
  string theme = 2;
  string timezone = 3;
  map<string, string> custom_settings = 4;
  bool success = 5;
  string error_message = 6;""",
            "guid": "web14000-2222-3333-4444-555555555555",
        },
        # File operations
        {
            "file": "proto/web/v1/upload_subtitle_request.proto",
            "message": "UploadSubtitleRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // File information
  string filename = 2;
  bytes content = 3;
  string content_type = 4;
  map<string, string> metadata = 5;""",
            "guid": "web15000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/upload_subtitle_response.proto",
            "message": "UploadSubtitleResponse",
            "fields": """  // Upload result
  string file_id = 1;
  string filename = 2;
  int64 size = 3;
  bool success = 4;
  string error_message = 5;""",
            "guid": "web16000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/download_subtitle_request.proto",
            "message": "DownloadSubtitleRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // File to download
  string file_id = 2;
  string format = 3;""",
            "guid": "web17000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/download_subtitle_response.proto",
            "message": "DownloadSubtitleResponse",
            "fields": """  // Download result
  string filename = 1;
  bytes content = 2;
  string content_type = 3;
  bool success = 4;
  string error_message = 5;""",
            "guid": "web18000-2222-3333-4444-555555555555",
        },
        # Search
        {
            "file": "proto/web/v1/search_subtitles_request.proto",
            "message": "SearchSubtitlesRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // Search criteria
  string query = 2;
  string language = 3;
  int32 limit = 4;
  int32 offset = 5;""",
            "guid": "web19000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/search_subtitles_response.proto",
            "message": "SearchSubtitlesResponse",
            "fields": """  // Search results
  repeated string file_ids = 1;
  repeated string filenames = 2;
  int32 total_count = 3;
  bool success = 4;
  string error_message = 5;""",
            "guid": "web20000-2222-3333-4444-555555555555",
        },
        # Translation
        {
            "file": "proto/web/v1/translate_subtitle_request.proto",
            "message": "TranslateSubtitleRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // File to translate
  string file_id = 2;
  string source_language = 3;
  string target_language = 4;
  map<string, string> options = 5;""",
            "guid": "web21000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/translate_subtitle_response.proto",
            "message": "TranslateSubtitleResponse",
            "fields": """  // Translation job
  string job_id = 1;
  string status = 2;
  string result_file_id = 3;
  bool success = 4;
  string error_message = 5;""",
            "guid": "web22000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/get_translation_status_request.proto",
            "message": "GetTranslationStatusRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // Job ID to check
  string job_id = 2;""",
            "guid": "web23000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/get_translation_status_response.proto",
            "message": "GetTranslationStatusResponse",
            "fields": """  // Job status
  string job_id = 1;
  string status = 2;
  float progress = 3;
  string result_file_id = 4;
  bool success = 5;
  string error_message = 6;""",
            "guid": "web24000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/cancel_translation_request.proto",
            "message": "CancelTranslationRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // Job ID to cancel
  string job_id = 2;""",
            "guid": "web25000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/cancel_translation_response.proto",
            "message": "CancelTranslationResponse",
            "fields": """  // Cancellation result
  string job_id = 1;
  bool cancelled = 2;
  bool success = 3;
  string error_message = 4;""",
            "guid": "web26000-2222-3333-4444-555555555555",
        },
        # Streaming
        {
            "file": "proto/web/v1/upload_file_request.proto",
            "message": "UploadFileRequest",
            "fields": """  // File chunk or metadata
  oneof data {
    FileMetadata metadata = 1;
    bytes chunk = 2;
  }

  message FileMetadata {
    string filename = 1;
    string content_type = 2;
    int64 total_size = 3;
  }""",
            "guid": "web27000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/upload_file_response.proto",
            "message": "UploadFileResponse",
            "fields": """  // Upload result
  string file_id = 1;
  string filename = 2;
  int64 size = 3;
  bool success = 4;
  string error_message = 5;""",
            "guid": "web28000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/download_file_request.proto",
            "message": "DownloadFileRequest",
            "fields": """  // Request ID for tracking
  string request_id = 1;

  // File to download
  string file_id = 2;""",
            "guid": "web29000-2222-3333-4444-555555555555",
        },
        {
            "file": "proto/web/v1/download_file_response.proto",
            "message": "DownloadFileResponse",
            "fields": """  // File chunk or metadata
  oneof data {
    FileMetadata metadata = 1;
    bytes chunk = 2;
  }

  message FileMetadata {
    string filename = 1;
    string content_type = 2;
    int64 total_size = 3;
  }""",
            "guid": "web30000-2222-3333-4444-555555555555",
        },
    ]

    # Create all web service messages
    for msg in web_messages:
        create_proto_file(
            file_path=msg["file"],
            package="subtitle_manager.web.v1",
            message_name=msg["message"],
            fields=msg["fields"],
            guid=msg["guid"],
        )

    print(f"Generated {len(web_messages)} web service message files")


if __name__ == "__main__":
    main()
