<!-- file: docs/tasks/01-core-architecture/TASK-01-001-service-interface-definitions-PART2.md -->
<!-- version: 1.0.0 -->
<!-- guid: 01001002-1111-2222-3333-444444444444 -->

# TASK-01-001: Service Interface Definitions (PART 2)

## Engine Service Interface Design

**Part**: 2 of 4 (Engine Service Interface) **Focus**: Translation, Monitoring,
and Coordination with Leader Election

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
