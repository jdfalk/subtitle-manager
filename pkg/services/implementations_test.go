// file: pkg/services/implementations_test.go
// version: 1.0.0
// guid: 7c6b1f7f-2fa3-4b92-8ed8-6c77aa36a7c7

package services

import (
	"context"
	"strings"
	"testing"

	enginev1 "github.com/jdfalk/subtitle-manager/pkg/engine/v1"
	filev1 "github.com/jdfalk/subtitle-manager/pkg/file/v1"
	webv1 "github.com/jdfalk/subtitle-manager/pkg/web/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewWebService_ReturnsInstance(t *testing.T) {
	// Arrange & Act
	service := NewWebService()

	// Assert
	if service == nil {
		t.Fatalf("expected web service implementation to be non-nil")
	}
}

func TestWebServiceImpl_LogoutUser_ReturnsSuccess(t *testing.T) {
	// Arrange
	service := NewWebService()
	ctx := context.Background()
	request := &webv1.LogoutUserRequest{}

	// Act
	response, err := service.LogoutUser(ctx, request)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if response == nil {
		t.Fatalf("expected response, got nil")
	}
	if !response.GetSuccess() {
		t.Fatalf("expected success to be true")
	}
}

func TestWebServiceImpl_UnimplementedMethods_ReturnUnimplemented(t *testing.T) {
	service := NewWebService()
	ctx := context.Background()

	tests := []struct {
		name         string
		call         func() (any, error)
		wantContains string
	}{
		{
			name: "AuthenticateUser",
			call: func() (any, error) {
				return service.AuthenticateUser(ctx, &webv1.AuthenticateUserRequest{})
			},
			wantContains: "AuthenticateUser not implemented",
		},
		{
			name: "GetUser",
			call: func() (any, error) {
				return service.GetUser(ctx, &webv1.GetUserRequest{})
			},
			wantContains: "GetUser not implemented",
		},
		{
			name: "UpdateUser",
			call: func() (any, error) {
				return service.UpdateUser(ctx, &webv1.UpdateUserRequest{})
			},
			wantContains: "UpdateUser not implemented",
		},
		{
			name: "UpdateUserPreferences",
			call: func() (any, error) {
				return service.UpdateUserPreferences(ctx, &webv1.UpdateUserPreferencesRequest{})
			},
			wantContains: "UpdateUserPreferences not implemented",
		},
		{
			name: "UploadSubtitle",
			call: func() (any, error) {
				return service.UploadSubtitle(ctx, &webv1.UploadSubtitleRequest{})
			},
			wantContains: "UploadSubtitle not implemented",
		},
		{
			name: "DownloadSubtitle",
			call: func() (any, error) {
				return service.DownloadSubtitle(ctx, &webv1.DownloadSubtitleRequest{})
			},
			wantContains: "DownloadSubtitle not implemented",
		},
		{
			name: "SearchSubtitles",
			call: func() (any, error) {
				return service.SearchSubtitles(ctx, &webv1.SearchSubtitlesRequest{})
			},
			wantContains: "SearchSubtitles not implemented",
		},
		{
			name: "TranslateSubtitle",
			call: func() (any, error) {
				return service.TranslateSubtitle(ctx, &webv1.TranslateSubtitleRequest{})
			},
			wantContains: "TranslateSubtitle not implemented",
		},
		{
			name: "GetTranslationStatus",
			call: func() (any, error) {
				return service.GetTranslationStatus(ctx, &webv1.GetTranslationStatusRequest{})
			},
			wantContains: "GetTranslationStatus not implemented",
		},
		{
			name: "CancelTranslation",
			call: func() (any, error) {
				return service.CancelTranslation(ctx, &webv1.CancelTranslationRequest{})
			},
			wantContains: "CancelTranslation not implemented",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			response, err := tt.call()

			// Assert
			if response == nil {
				t.Fatalf("expected response to be non-nil")
			}
			assertUnimplemented(t, err, tt.wantContains)
		})
	}
}

func TestWebServiceImpl_StreamMethods_ReturnUnimplemented(t *testing.T) {
	service := NewWebService()

	tests := []struct {
		name         string
		call         func() error
		wantContains string
	}{
		{
			name: "UploadFile",
			call: func() error {
				var stream webv1.WebService_UploadFileServer
				return service.UploadFile(stream)
			},
			wantContains: "UploadFile streaming not implemented",
		},
		{
			name: "DownloadFile",
			call: func() error {
				var stream webv1.WebService_DownloadFileServer
				return service.DownloadFile(&webv1.DownloadFileRequest{}, stream)
			},
			wantContains: "DownloadFile streaming not implemented",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			err := tt.call()

			// Assert
			assertUnimplemented(t, err, tt.wantContains)
		})
	}
}

func TestWebServiceImpl_HealthCheck_ReturnsHealthyResponse(t *testing.T) {
	// Arrange
	service := NewWebService()
	ctx := context.Background()
	request := &webv1.HealthCheckRequest{}

	// Act
	response, err := service.HealthCheck(ctx, request)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if response == nil {
		t.Fatalf("expected response, got nil")
	}
}

func TestNewEngineService_ReturnsInstance(t *testing.T) {
	// Arrange & Act
	service := NewEngineService()

	// Assert
	if service == nil {
		t.Fatalf("expected engine service implementation to be non-nil")
	}
}

func TestEngineServiceImpl_UnimplementedMethods_ReturnUnimplemented(t *testing.T) {
	service := NewEngineService()
	ctx := context.Background()

	tests := []struct {
		name         string
		call         func() (any, error)
		wantContains string
	}{
		{
			name: "TranscribeAudio",
			call: func() (any, error) {
				return service.TranscribeAudio(ctx, &enginev1.TranscribeAudioRequest{})
			},
			wantContains: "TranscribeAudio not implemented",
		},
		{
			name: "GetTranscriptionStatus",
			call: func() (any, error) {
				return service.GetTranscriptionStatus(ctx, &enginev1.GetTranscriptionStatusRequest{})
			},
			wantContains: "GetTranscriptionStatus not implemented",
		},
		{
			name: "CancelTranscription",
			call: func() (any, error) {
				return service.CancelTranscription(ctx, &enginev1.CancelTranscriptionRequest{})
			},
			wantContains: "CancelTranscription not implemented",
		},
		{
			name: "TranslateSubtitle",
			call: func() (any, error) {
				return service.TranslateSubtitle(ctx, &enginev1.TranslateSubtitleRequest{})
			},
			wantContains: "TranslateSubtitle not implemented",
		},
		{
			name: "GetTranslationProgress",
			call: func() (any, error) {
				return service.GetTranslationProgress(ctx, &enginev1.GetTranslationProgressRequest{})
			},
			wantContains: "GetTranslationProgress not implemented",
		},
		{
			name: "CancelTranslation",
			call: func() (any, error) {
				return service.CancelTranslation(ctx, &enginev1.CancelTranslationRequest{})
			},
			wantContains: "CancelTranslation not implemented",
		},
		{
			name: "ConvertSubtitle",
			call: func() (any, error) {
				return service.ConvertSubtitle(ctx, &enginev1.ConvertSubtitleRequest{})
			},
			wantContains: "ConvertSubtitle not implemented",
		},
		{
			name: "ValidateSubtitle",
			call: func() (any, error) {
				return service.ValidateSubtitle(ctx, &enginev1.ValidateSubtitleRequest{})
			},
			wantContains: "ValidateSubtitle not implemented",
		},
		{
			name: "MergeSubtitles",
			call: func() (any, error) {
				return service.MergeSubtitles(ctx, &enginev1.MergeSubtitlesRequest{})
			},
			wantContains: "MergeSubtitles not implemented",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			response, err := tt.call()

			// Assert
			if response == nil {
				t.Fatalf("expected response to be non-nil")
			}
			assertUnimplemented(t, err, tt.wantContains)
		})
	}
}

func TestEngineServiceImpl_HealthChecks_ReturnHealthyResponses(t *testing.T) {
	// Arrange
	service := NewEngineService()
	ctx := context.Background()

	tests := []struct {
		name string
		call func() (any, error)
	}{
		{
			name: "GetEngineStatus",
			call: func() (any, error) {
				return service.GetEngineStatus(ctx, &enginev1.GetEngineStatusRequest{})
			},
		},
		{
			name: "HealthCheck",
			call: func() (any, error) {
				return service.HealthCheck(ctx, &enginev1.HealthCheckRequest{})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			response, err := tt.call()

			// Assert
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if response == nil {
				t.Fatalf("expected response, got nil")
			}
		})
	}
}

func TestNewFileService_ReturnsInstance(t *testing.T) {
	// Arrange & Act
	service := NewFileService()

	// Assert
	if service == nil {
		t.Fatalf("expected file service implementation to be non-nil")
	}
}

func TestFileServiceImpl_UnimplementedMethods_ReturnUnimplemented(t *testing.T) {
	service := NewFileService()
	ctx := context.Background()

	tests := []struct {
		name         string
		call         func() (any, error)
		wantContains string
	}{
		{
			name: "DeleteFile",
			call: func() (any, error) {
				return service.DeleteFile(ctx, &filev1.DeleteFileRequest{})
			},
			wantContains: "DeleteFile not implemented",
		},
		{
			name: "GetFileInfo",
			call: func() (any, error) {
				return service.GetFileInfo(ctx, &filev1.GetFileInfoRequest{})
			},
			wantContains: "GetFileInfo not implemented",
		},
		{
			name: "CopyFile",
			call: func() (any, error) {
				return service.CopyFile(ctx, &filev1.CopyFileRequest{})
			},
			wantContains: "CopyFile not implemented",
		},
		{
			name: "MoveFile",
			call: func() (any, error) {
				return service.MoveFile(ctx, &filev1.MoveFileRequest{})
			},
			wantContains: "MoveFile not implemented",
		},
		{
			name: "ListFiles",
			call: func() (any, error) {
				return service.ListFiles(ctx, &filev1.ListFilesRequest{})
			},
			wantContains: "ListFiles not implemented",
		},
		{
			name: "UpdateFileMetadata",
			call: func() (any, error) {
				return service.UpdateFileMetadata(ctx, &filev1.UpdateFileMetadataRequest{})
			},
			wantContains: "UpdateFileMetadata not implemented",
		},
		{
			name: "SearchFiles",
			call: func() (any, error) {
				return service.SearchFiles(ctx, &filev1.SearchFilesRequest{})
			},
			wantContains: "SearchFiles not implemented",
		},
		{
			name: "GetStorageInfo",
			call: func() (any, error) {
				return service.GetStorageInfo(ctx, &filev1.GetStorageInfoRequest{})
			},
			wantContains: "GetStorageInfo not implemented",
		},
		{
			name: "CleanupFiles",
			call: func() (any, error) {
				return service.CleanupFiles(ctx, &filev1.CleanupFilesRequest{})
			},
			wantContains: "CleanupFiles not implemented",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			response, err := tt.call()

			// Assert
			if response == nil {
				t.Fatalf("expected response to be non-nil")
			}
			assertUnimplemented(t, err, tt.wantContains)
		})
	}
}

func TestFileServiceImpl_StreamMethods_ReturnUnimplemented(t *testing.T) {
	service := NewFileService()

	tests := []struct {
		name         string
		call         func() error
		wantContains string
	}{
		{
			name: "UploadFile",
			call: func() error {
				var stream filev1.FileService_UploadFileServer
				return service.UploadFile(stream)
			},
			wantContains: "UploadFile streaming not implemented",
		},
		{
			name: "DownloadFile",
			call: func() error {
				var stream filev1.FileService_DownloadFileServer
				return service.DownloadFile(&filev1.DownloadFileRequest{}, stream)
			},
			wantContains: "DownloadFile streaming not implemented",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			err := tt.call()

			// Assert
			assertUnimplemented(t, err, tt.wantContains)
		})
	}
}

func TestFileServiceImpl_HealthCheck_ReturnsHealthyResponse(t *testing.T) {
	// Arrange
	service := NewFileService()
	ctx := context.Background()
	request := &filev1.HealthCheckRequest{}

	// Act
	response, err := service.HealthCheck(ctx, request)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if response == nil {
		t.Fatalf("expected response, got nil")
	}
}

func assertUnimplemented(t *testing.T, err error, wantContains string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	statusErr, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected grpc status error, got %T", err)
	}
	if statusErr.Code() != codes.Unimplemented {
		t.Fatalf("expected unimplemented code, got %s", statusErr.Code())
	}
	if !strings.Contains(statusErr.Message(), wantContains) {
		t.Fatalf("expected error message to contain %q, got %q", wantContains, statusErr.Message())
	}
}
