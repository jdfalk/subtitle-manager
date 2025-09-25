// file: pkg/types/types_test.go
// version: 1.0.0
// guid: 550e8400-e29b-41d4-a716-446655440027

package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestAPIError tests the APIError type and its methods
func TestAPIError(t *testing.T) {
	tests := []struct {
		name       string
		apiError   *APIError
		wantError  string
		wantAuth   bool
		wantAuthz  bool
		wantNotFnd bool
		wantRate   bool
	}{
		{
			name: "authentication error",
			apiError: &APIError{
				StatusCode: 401,
				Err:        "Unauthorized",
				Message:    "Invalid credentials",
			},
			wantError: "API error (401): Unauthorized - Invalid credentials",
			wantAuth:  true,
		},
		{
			name: "authorization error",
			apiError: &APIError{
				StatusCode: 403,
				Err:        "Forbidden",
				Message:    "Insufficient permissions",
			},
			wantError: "API error (403): Forbidden - Insufficient permissions",
			wantAuthz: true,
		},
		{
			name: "not found error",
			apiError: &APIError{
				StatusCode: 404,
				Err:        "Not Found",
				Message:    "Resource not found",
			},
			wantError:  "API error (404): Not Found - Resource not found",
			wantNotFnd: true,
		},
		{
			name: "rate limit error",
			apiError: &APIError{
				StatusCode: 429,
				Err:        "Too Many Requests",
				Message:    "Rate limit exceeded",
			},
			wantError: "API error (429): Too Many Requests - Rate limit exceeded",
			wantRate:  true,
		},
		{
			name: "internal server error",
			apiError: &APIError{
				StatusCode: 500,
				Err:        "Internal Server Error",
				Message:    "Something went wrong",
			},
			wantError: "API error (500): Internal Server Error - Something went wrong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantError, tt.apiError.Error())
			assert.Equal(t, tt.wantAuth, tt.apiError.IsAuthenticationError())
			assert.Equal(t, tt.wantAuthz, tt.apiError.IsAuthorizationError())
			assert.Equal(t, tt.wantNotFnd, tt.apiError.IsNotFoundError())
			assert.Equal(t, tt.wantRate, tt.apiError.IsRateLimitError())
		})
	}
}

// TestLoginResponse tests the LoginResponse type and its role checking methods
func TestLoginResponse(t *testing.T) {
	tests := []struct {
		name          string
		loginResponse *LoginResponse
		wantAdmin     bool
		wantBasicAcc  bool
		wantReadAcc   bool
	}{
		{
			name: "admin user",
			loginResponse: &LoginResponse{
				UserID:   1,
				Username: "admin",
				Role:     RoleAdmin,
			},
			wantAdmin:    true,
			wantBasicAcc: true,
			wantReadAcc:  true,
		},
		{
			name: "basic user",
			loginResponse: &LoginResponse{
				UserID:   2,
				Username: "basicuser",
				Role:     RoleBasic,
			},
			wantAdmin:    false,
			wantBasicAcc: true,
			wantReadAcc:  true,
		},
		{
			name: "read-only user",
			loginResponse: &LoginResponse{
				UserID:   3,
				Username: "readuser",
				Role:     RoleRead,
			},
			wantAdmin:    false,
			wantBasicAcc: false,
			wantReadAcc:  true,
		},
		{
			name: "unknown role",
			loginResponse: &LoginResponse{
				UserID:   4,
				Username: "unknownuser",
				Role:     "unknown",
			},
			wantAdmin:    false,
			wantBasicAcc: false,
			wantReadAcc:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantAdmin, tt.loginResponse.IsAdmin())
			assert.Equal(t, tt.wantBasicAcc, tt.loginResponse.HasBasicAccess())
			assert.Equal(t, tt.wantReadAcc, tt.loginResponse.HasReadAccess())
		})
	}
}

// TestSystemInfo tests the SystemInfo type and its calculations
func TestSystemInfo(t *testing.T) {
	tests := []struct {
		name         string
		systemInfo   *SystemInfo
		wantDiskUsed float64
	}{
		{
			name: "normal disk usage",
			systemInfo: &SystemInfo{
				DiskFree:  300 * 1024 * 1024 * 1024, // 300GB free
				DiskTotal: 500 * 1024 * 1024 * 1024, // 500GB total
			},
			wantDiskUsed: 40.0, // 200GB used / 500GB total = 40%
		},
		{
			name: "full disk",
			systemInfo: &SystemInfo{
				DiskFree:  0,
				DiskTotal: 100 * 1024 * 1024 * 1024, // 100GB total
			},
			wantDiskUsed: 100.0,
		},
		{
			name: "empty disk",
			systemInfo: &SystemInfo{
				DiskFree:  100 * 1024 * 1024 * 1024, // 100GB free
				DiskTotal: 100 * 1024 * 1024 * 1024, // 100GB total
			},
			wantDiskUsed: 0.0,
		},
		{
			name: "zero total disk",
			systemInfo: &SystemInfo{
				DiskFree:  0,
				DiskTotal: 0,
			},
			wantDiskUsed: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantDiskUsed, tt.systemInfo.DiskUsagePercent())
		})
	}
}

// TestLogEntry tests the LogEntry type and its level checking methods
func TestLogEntry(t *testing.T) {
	timestamp := time.Now()
	fields := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	tests := []struct {
		name      string
		logEntry  *LogEntry
		wantError bool
		wantWarn  bool
	}{
		{
			name: "error log entry",
			logEntry: &LogEntry{
				Timestamp: timestamp,
				Level:     LogLevelError,
				Component: "auth",
				Message:   "Authentication failed",
				Fields:    fields,
			},
			wantError: true,
			wantWarn:  false,
		},
		{
			name: "warning log entry",
			logEntry: &LogEntry{
				Timestamp: timestamp,
				Level:     LogLevelWarn,
				Component: "scanner",
				Message:   "File not found",
				Fields:    fields,
			},
			wantError: false,
			wantWarn:  true,
		},
		{
			name: "info log entry",
			logEntry: &LogEntry{
				Timestamp: timestamp,
				Level:     LogLevelInfo,
				Component: "server",
				Message:   "Server started",
				Fields:    fields,
			},
			wantError: false,
			wantWarn:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantError, tt.logEntry.IsError())
			assert.Equal(t, tt.wantWarn, tt.logEntry.IsWarning())
		})
	}
}

// TestHistoryItem tests the HistoryItem type and its status checking methods
func TestHistoryItem(t *testing.T) {
	createdAt := time.Now()
	subtitlePath := "/path/to/subtitle.srt"
	language := "en"
	provider := "opensubtitles"
	errorMessage := "Failed to download"

	tests := []struct {
		name        string
		historyItem *HistoryItem
		wantSuccess bool
		wantFailed  bool
		wantPending bool
	}{
		{
			name: "successful operation",
			historyItem: &HistoryItem{
				ID:           1,
				Type:         OperationTypeDownload,
				FilePath:     "/path/to/movie.mp4",
				SubtitlePath: &subtitlePath,
				Language:     &language,
				Provider:     &provider,
				Status:       OperationStatusSuccess,
				CreatedAt:    createdAt,
				UserID:       123,
			},
			wantSuccess: true,
		},
		{
			name: "failed operation",
			historyItem: &HistoryItem{
				ID:           2,
				Type:         OperationTypeDownload,
				FilePath:     "/path/to/movie.mp4",
				Language:     &language,
				Provider:     &provider,
				Status:       OperationStatusFailed,
				CreatedAt:    createdAt,
				UserID:       123,
				ErrorMessage: &errorMessage,
			},
			wantFailed: true,
		},
		{
			name: "pending operation",
			historyItem: &HistoryItem{
				ID:        3,
				Type:      OperationTypeTranslate,
				FilePath:  "/path/to/movie.mp4",
				Status:    OperationStatusPending,
				CreatedAt: createdAt,
				UserID:    123,
			},
			wantPending: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantSuccess, tt.historyItem.IsSuccess())
			assert.Equal(t, tt.wantFailed, tt.historyItem.IsFailed())
			assert.Equal(t, tt.wantPending, tt.historyItem.IsPending())
		})
	}
}

// TestHistoryResponse tests the HistoryResponse type and its pagination methods
func TestHistoryResponse(t *testing.T) {
	items := []HistoryItem{
		{ID: 1, Type: OperationTypeDownload},
		{ID: 2, Type: OperationTypeConvert},
	}

	tests := []struct {
		name            string
		historyResponse *HistoryResponse
		wantHasNext     bool
		wantHasPrev     bool
		wantTotalPages  int
	}{
		{
			name: "first page with more pages",
			historyResponse: &HistoryResponse{
				Items: items,
				Total: 50,
				Page:  1,
				Limit: 10,
			},
			wantHasNext:    true,
			wantHasPrev:    false,
			wantTotalPages: 5,
		},
		{
			name: "middle page",
			historyResponse: &HistoryResponse{
				Items: items,
				Total: 50,
				Page:  3,
				Limit: 10,
			},
			wantHasNext:    true,
			wantHasPrev:    true,
			wantTotalPages: 5,
		},
		{
			name: "last page",
			historyResponse: &HistoryResponse{
				Items: items,
				Total: 25,
				Page:  5,
				Limit: 5,
			},
			wantHasNext:    false,
			wantHasPrev:    true,
			wantTotalPages: 5,
		},
		{
			name: "single page",
			historyResponse: &HistoryResponse{
				Items: items,
				Total: 2,
				Page:  1,
				Limit: 10,
			},
			wantHasNext:    false,
			wantHasPrev:    false,
			wantTotalPages: 1,
		},
		{
			name: "zero limit",
			historyResponse: &HistoryResponse{
				Items: items,
				Total: 50,
				Page:  1,
				Limit: 0,
			},
			wantHasNext:    true, // With zero limit, Page*Limit (1*0=0) < Total (50) = true
			wantHasPrev:    false,
			wantTotalPages: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantHasNext, tt.historyResponse.HasNextPage())
			assert.Equal(t, tt.wantHasPrev, tt.historyResponse.HasPreviousPage())
			assert.Equal(t, tt.wantTotalPages, tt.historyResponse.TotalPages())
		})
	}
}

// TestScanStatus tests the ScanStatus type and its calculation methods
func TestScanStatus(t *testing.T) {
	tests := []struct {
		name            string
		scanStatus      *ScanStatus
		wantProgressPct float64
		wantRemaining   int
	}{
		{
			name: "scan in progress",
			scanStatus: &ScanStatus{
				Scanning:       true,
				Progress:       0.45,
				FilesProcessed: func() *int { v := 45; return &v }(),
				FilesTotal:     func() *int { v := 100; return &v }(),
			},
			wantProgressPct: 45.0,
			wantRemaining:   55,
		},
		{
			name: "scan completed",
			scanStatus: &ScanStatus{
				Scanning:       false,
				Progress:       1.0,
				FilesProcessed: func() *int { v := 100; return &v }(),
				FilesTotal:     func() *int { v := 100; return &v }(),
			},
			wantProgressPct: 100.0,
			wantRemaining:   0,
		},
		{
			name: "scan starting",
			scanStatus: &ScanStatus{
				Scanning:       true,
				Progress:       0.0,
				FilesProcessed: func() *int { v := 0; return &v }(),
				FilesTotal:     func() *int { v := 50; return &v }(),
			},
			wantProgressPct: 0.0,
			wantRemaining:   50,
		},
		{
			name: "no file counts",
			scanStatus: &ScanStatus{
				Scanning: true,
				Progress: 0.3,
			},
			wantProgressPct: 30.0,
			wantRemaining:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantProgressPct, tt.scanStatus.ProgressPercent())
			assert.Equal(t, tt.wantRemaining, tt.scanStatus.RemainingFiles())
		})
	}
}

// TestConstants tests that all constants are defined correctly
func TestConstants(t *testing.T) {
	// Test operation type constants
	assert.Equal(t, "download", OperationTypeDownload)
	assert.Equal(t, "convert", OperationTypeConvert)
	assert.Equal(t, "translate", OperationTypeTranslate)
	assert.Equal(t, "extract", OperationTypeExtract)

	// Test operation status constants
	assert.Equal(t, "success", OperationStatusSuccess)
	assert.Equal(t, "failed", OperationStatusFailed)
	assert.Equal(t, "pending", OperationStatusPending)

	// Test log level constants
	assert.Equal(t, "debug", LogLevelDebug)
	assert.Equal(t, "info", LogLevelInfo)
	assert.Equal(t, "warn", LogLevelWarn)
	assert.Equal(t, "error", LogLevelError)

	// Test user role constants
	assert.Equal(t, "read", RoleRead)
	assert.Equal(t, "basic", RoleBasic)
	assert.Equal(t, "admin", RoleAdmin)

	// Test provider constants
	assert.Equal(t, "google", ProviderGoogle)
	assert.Equal(t, "openai", ProviderOpenAI)
}

// TestStructureIntegrity tests that structures have expected fields and types
func TestStructureIntegrity(t *testing.T) {
	// Test LoginRequest structure
	loginReq := LoginRequest{
		Username: "testuser",
		Password: "testpass",
	}
	assert.Equal(t, "testuser", loginReq.Username)
	assert.Equal(t, "testpass", loginReq.Password)

	// Test DownloadRequest structure
	providers := []string{"opensubtitles", "subscene"}
	downloadReq := DownloadRequest{
		Path:      "/path/to/movie.mp4",
		Language:  "en",
		Providers: providers,
	}
	assert.Equal(t, "/path/to/movie.mp4", downloadReq.Path)
	assert.Equal(t, "en", downloadReq.Language)
	assert.Equal(t, providers, downloadReq.Providers)

	// Test DownloadResult structure
	subtitlePath := "/path/to/subtitle.srt"
	provider := "opensubtitles"
	downloadResult := DownloadResult{
		Success:      true,
		SubtitlePath: &subtitlePath,
		Provider:     &provider,
	}
	assert.True(t, downloadResult.Success)
	assert.NotNil(t, downloadResult.SubtitlePath)
	assert.Equal(t, subtitlePath, *downloadResult.SubtitlePath)
	assert.NotNil(t, downloadResult.Provider)
	assert.Equal(t, provider, *downloadResult.Provider)

	// Test ScanRequest structure
	scanPath := "/media/movies"
	scanReq := ScanRequest{
		Path:  &scanPath,
		Force: true,
	}
	assert.NotNil(t, scanReq.Path)
	assert.Equal(t, scanPath, *scanReq.Path)
	assert.True(t, scanReq.Force)

	// Test ScanResult structure
	scanResult := ScanResult{
		ScanID: "scan123",
	}
	assert.Equal(t, "scan123", scanResult.ScanID)

	// Test OAuthCredentials structure
	redirectURL := "http://localhost:8080/callback"
	oauthCreds := OAuthCredentials{
		ClientID:     "client123",
		ClientSecret: "secret456",
		RedirectURL:  &redirectURL,
	}
	assert.Equal(t, "client123", oauthCreds.ClientID)
	assert.Equal(t, "secret456", oauthCreds.ClientSecret)
	assert.NotNil(t, oauthCreds.RedirectURL)
	assert.Equal(t, redirectURL, *oauthCreds.RedirectURL)

	// Test LogParams structure
	logParams := LogParams{
		Level: LogLevelError,
		Limit: 100,
	}
	assert.Equal(t, LogLevelError, logParams.Level)
	assert.Equal(t, 100, logParams.Limit)

	// Test HistoryParams structure
	startDate := time.Now().AddDate(0, -1, 0) // 1 month ago
	endDate := time.Now()
	historyParams := HistoryParams{
		Page:      1,
		Limit:     50,
		Type:      OperationTypeDownload,
		StartDate: startDate,
		EndDate:   endDate,
	}
	assert.Equal(t, 1, historyParams.Page)
	assert.Equal(t, 50, historyParams.Limit)
	assert.Equal(t, OperationTypeDownload, historyParams.Type)
	assert.Equal(t, startDate, historyParams.StartDate)
	assert.Equal(t, endDate, historyParams.EndDate)
}
