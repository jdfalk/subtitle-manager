// file: sdks/go/subtitleclient/types.go
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440026

package subtitleclient

import (
	"fmt"
	"time"
)

// Error types

// APIError represents an error response from the API.
type APIError struct {
	StatusCode int    `json:"-"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (%d): %s - %s", e.StatusCode, e.Error, e.Message)
}

// IsAuthenticationError returns true if the error is an authentication error (401).
func (e *APIError) IsAuthenticationError() bool {
	return e.StatusCode == 401
}

// IsAuthorizationError returns true if the error is an authorization error (403).
func (e *APIError) IsAuthorizationError() bool {
	return e.StatusCode == 403
}

// IsNotFoundError returns true if the error is a not found error (404).
func (e *APIError) IsNotFoundError() bool {
	return e.StatusCode == 404
}

// IsRateLimitError returns true if the error is a rate limit error (429).
func (e *APIError) IsRateLimitError() bool {
	return e.StatusCode == 429
}

// Authentication types

// LoginRequest represents a login request.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents a login response.
type LoginResponse struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// IsAdmin returns true if the user has admin role.
func (l *LoginResponse) IsAdmin() bool {
	return l.Role == "admin"
}

// HasBasicAccess returns true if the user has basic or admin access.
func (l *LoginResponse) HasBasicAccess() bool {
	return l.Role == "basic" || l.Role == "admin"
}

// HasReadAccess returns true if the user has read, basic, or admin access.
func (l *LoginResponse) HasReadAccess() bool {
	return l.Role == "read" || l.Role == "basic" || l.Role == "admin"
}

// System information types

// SystemInfo represents system information.
type SystemInfo struct {
	GoVersion   string  `json:"go_version"`
	OS          string  `json:"os"`
	Arch        string  `json:"arch"`
	Goroutines  int     `json:"goroutines"`
	DiskFree    uint64  `json:"disk_free"`
	DiskTotal   uint64  `json:"disk_total"`
	MemoryUsage *uint64 `json:"memory_usage,omitempty"`
	Uptime      *string `json:"uptime,omitempty"`
	Version     *string `json:"version,omitempty"`
}

// DiskUsagePercent returns the disk usage percentage.
func (s *SystemInfo) DiskUsagePercent() float64 {
	if s.DiskTotal == 0 {
		return 0
	}
	return float64(s.DiskTotal-s.DiskFree) / float64(s.DiskTotal) * 100
}

// LogEntry represents a log entry.
type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     string                 `json:"level"`
	Component string                 `json:"component"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields"`
}

// IsError returns true if the log entry is an error.
func (l *LogEntry) IsError() bool {
	return l.Level == "error"
}

// IsWarning returns true if the log entry is a warning.
func (l *LogEntry) IsWarning() bool {
	return l.Level == "warn"
}

// LogParams represents parameters for getting logs.
type LogParams struct {
	Level string
	Limit int
}

// Operation types

// HistoryItem represents an operation history item.
type HistoryItem struct {
	ID           int64     `json:"id"`
	Type         string    `json:"type"`
	FilePath     string    `json:"file_path"`
	SubtitlePath *string   `json:"subtitle_path,omitempty"`
	Language     *string   `json:"language,omitempty"`
	Provider     *string   `json:"provider,omitempty"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UserID       int64     `json:"user_id"`
	ErrorMessage *string   `json:"error_message,omitempty"`
}

// IsSuccess returns true if the operation was successful.
func (h *HistoryItem) IsSuccess() bool {
	return h.Status == "success"
}

// IsFailed returns true if the operation failed.
func (h *HistoryItem) IsFailed() bool {
	return h.Status == "failed"
}

// IsPending returns true if the operation is pending.
func (h *HistoryItem) IsPending() bool {
	return h.Status == "pending"
}

// HistoryParams represents parameters for getting history.
type HistoryParams struct {
	Page      int
	Limit     int
	Type      string
	StartDate time.Time
	EndDate   time.Time
}

// HistoryResponse represents a paginated history response.
type HistoryResponse struct {
	Items []HistoryItem `json:"items"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

// HasNextPage returns true if there are more pages available.
func (h *HistoryResponse) HasNextPage() bool {
	return h.Page*h.Limit < h.Total
}

// HasPreviousPage returns true if there are previous pages.
func (h *HistoryResponse) HasPreviousPage() bool {
	return h.Page > 1
}

// TotalPages returns the total number of pages.
func (h *HistoryResponse) TotalPages() int {
	if h.Limit == 0 {
		return 0
	}
	return (h.Total + h.Limit - 1) / h.Limit
}

// Download types

// DownloadRequest represents a request to download subtitles.
type DownloadRequest struct {
	Path      string   `json:"path"`
	Language  string   `json:"language"`
	Providers []string `json:"providers,omitempty"`
}

// DownloadResult represents the result of a subtitle download.
type DownloadResult struct {
	Success      bool    `json:"success"`
	SubtitlePath *string `json:"subtitle_path,omitempty"`
	Provider     *string `json:"provider,omitempty"`
}

// Scan types

// ScanRequest represents a request to start a library scan.
type ScanRequest struct {
	Path  *string `json:"path,omitempty"`
	Force bool    `json:"force,omitempty"`
}

// ScanResult represents the result of starting a scan.
type ScanResult struct {
	ScanID string `json:"scan_id"`
}

// ScanStatus represents the status of a library scan.
type ScanStatus struct {
	Scanning             bool       `json:"scanning"`
	Progress             float64    `json:"progress"`
	CurrentPath          *string    `json:"current_path,omitempty"`
	FilesProcessed       *int       `json:"files_processed,omitempty"`
	FilesTotal           *int       `json:"files_total,omitempty"`
	StartTime            *time.Time `json:"start_time,omitempty"`
	EstimatedCompletion  *time.Time `json:"estimated_completion,omitempty"`
}

// ProgressPercent returns the progress as a percentage (0-100).
func (s *ScanStatus) ProgressPercent() float64 {
	return s.Progress * 100
}

// RemainingFiles returns the number of files remaining to process.
func (s *ScanStatus) RemainingFiles() int {
	if s.FilesTotal != nil && s.FilesProcessed != nil {
		return *s.FilesTotal - *s.FilesProcessed
	}
	return 0
}

// OAuth types

// OAuthCredentials represents OAuth2 credentials.
type OAuthCredentials struct {
	ClientID     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	RedirectURL  *string `json:"redirect_url,omitempty"`
}

// Constants for operation types
const (
	OperationTypeDownload  = "download"
	OperationTypeConvert   = "convert"
	OperationTypeTranslate = "translate"
	OperationTypeExtract   = "extract"
)

// Constants for operation statuses
const (
	OperationStatusSuccess = "success"
	OperationStatusFailed  = "failed"
	OperationStatusPending = "pending"
)

// Constants for log levels
const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
)

// Constants for user roles
const (
	RoleRead  = "read"
	RoleBasic = "basic"
	RoleAdmin = "admin"
)

// Constants for translation providers
const (
	ProviderGoogle = "google"
	ProviderOpenAI = "openai"
)