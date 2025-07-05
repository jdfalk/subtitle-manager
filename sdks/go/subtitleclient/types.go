// file: sdks/go/subtitleclient/types.go
package subtitleclient

import "github.com/jdfalk/subtitle-manager/pkg/types"

// Re-export shared types from the central package.
type (
	APIError         = types.APIError
	LoginRequest     = types.LoginRequest
	LoginResponse    = types.LoginResponse
	SystemInfo       = types.SystemInfo
	LogEntry         = types.LogEntry
	LogParams        = types.LogParams
	HistoryItem      = types.HistoryItem
	HistoryParams    = types.HistoryParams
	HistoryResponse  = types.HistoryResponse
	DownloadRequest  = types.DownloadRequest
	DownloadResult   = types.DownloadResult
	ScanRequest      = types.ScanRequest
	ScanResult       = types.ScanResult
	ScanStatus       = types.ScanStatus
	OAuthCredentials = types.OAuthCredentials
)

const (
	OperationTypeDownload  = types.OperationTypeDownload
	OperationTypeConvert   = types.OperationTypeConvert
	OperationTypeTranslate = types.OperationTypeTranslate
	OperationTypeExtract   = types.OperationTypeExtract

	OperationStatusSuccess = types.OperationStatusSuccess
	OperationStatusFailed  = types.OperationStatusFailed
	OperationStatusPending = types.OperationStatusPending

	LogLevelDebug = types.LogLevelDebug
	LogLevelInfo  = types.LogLevelInfo
	LogLevelWarn  = types.LogLevelWarn
	LogLevelError = types.LogLevelError

	RoleRead  = types.RoleRead
	RoleBasic = types.RoleBasic
	RoleAdmin = types.RoleAdmin

	ProviderGoogle = types.ProviderGoogle
	ProviderOpenAI = types.ProviderOpenAI
)
