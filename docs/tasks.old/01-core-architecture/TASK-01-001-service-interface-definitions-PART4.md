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
// file: pkg/services/interfaces.go
// version: 1.0.0
// guid: srv00001-1111-2222-3333-444444444444

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
