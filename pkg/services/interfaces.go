// file: pkg/services/interfaces.go
// version: 1.8.0
// guid: 789e0123-e45b-12d3-a456-426614174000

package services

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	enginev1 "github.com/jdfalk/subtitle-manager/pkg/engine/v1"
	filev1 "github.com/jdfalk/subtitle-manager/pkg/file/v1"
	webv1 "github.com/jdfalk/subtitle-manager/pkg/web/v1"
)

// WebServiceInterface - basic interface for web service operations
type WebServiceInterface interface {
	AuthenticateUser(ctx context.Context, req *webv1.AuthenticateUserRequest) (*webv1.AuthenticateUserResponse, error)
	LogoutUser(ctx context.Context, req *webv1.LogoutUserRequest) (*emptypb.Empty, error)
	GetUser(ctx context.Context, req *webv1.GetUserRequest) (*webv1.GetUserResponse, error)
	UpdateUser(ctx context.Context, req *webv1.UpdateUserRequest) (*webv1.UpdateUserResponse, error)
	UpdateUserPreferences(ctx context.Context, req *webv1.UpdateUserPreferencesRequest) (*webv1.UpdateUserPreferencesResponse, error)
	UploadSubtitle(ctx context.Context, req *webv1.UploadSubtitleRequest) (*webv1.UploadSubtitleResponse, error)
	DownloadSubtitle(ctx context.Context, req *webv1.DownloadSubtitleRequest) (*webv1.DownloadSubtitleResponse, error)
	SearchSubtitles(ctx context.Context, req *webv1.SearchSubtitlesRequest) (*webv1.SearchSubtitlesResponse, error)
	TranslateSubtitle(ctx context.Context, req *webv1.TranslateSubtitleRequest) (*webv1.TranslateSubtitleResponse, error)
	GetTranslationStatus(ctx context.Context, req *webv1.GetTranslationStatusRequest) (*webv1.GetTranslationStatusResponse, error)
	CancelTranslation(ctx context.Context, req *webv1.CancelTranslationRequest) (*webv1.CancelTranslationResponse, error)
	UploadFile(stream webv1.WebService_UploadFileServer) error
	DownloadFile(req *webv1.DownloadFileRequest, stream webv1.WebService_DownloadFileServer) error
	HealthCheck(ctx context.Context, req *webv1.HealthCheckRequest) (*webv1.HealthCheckResponse, error)
}

// EngineServiceInterface - basic interface for engine service operations
type EngineServiceInterface interface {
	TranscribeAudio(ctx context.Context, req *enginev1.TranscribeAudioRequest) (*enginev1.TranscribeAudioResponse, error)
	GetTranscriptionStatus(ctx context.Context, req *enginev1.GetTranscriptionStatusRequest) (*enginev1.GetTranscriptionStatusResponse, error)
	CancelTranscription(ctx context.Context, req *enginev1.CancelTranscriptionRequest) (*enginev1.CancelTranscriptionResponse, error)
	TranslateSubtitle(ctx context.Context, req *enginev1.TranslateSubtitleRequest) (*enginev1.TranslateSubtitleResponse, error)
	GetTranslationProgress(ctx context.Context, req *enginev1.GetTranslationProgressRequest) (*enginev1.GetTranslationProgressResponse, error)
	CancelTranslation(ctx context.Context, req *enginev1.CancelTranslationRequest) (*enginev1.CancelTranslationResponse, error)
	ConvertSubtitle(ctx context.Context, req *enginev1.ConvertSubtitleRequest) (*enginev1.ConvertSubtitleResponse, error)
	ValidateSubtitle(ctx context.Context, req *enginev1.ValidateSubtitleRequest) (*enginev1.ValidateSubtitleResponse, error)
	MergeSubtitles(ctx context.Context, req *enginev1.MergeSubtitlesRequest) (*enginev1.MergeSubtitlesResponse, error)
	GetEngineStatus(ctx context.Context, req *enginev1.GetEngineStatusRequest) (*enginev1.GetEngineStatusResponse, error)
	HealthCheck(ctx context.Context, req *enginev1.HealthCheckRequest) (*enginev1.HealthCheckResponse, error)
}

// FileServiceInterface - basic interface for file service operations
type FileServiceInterface interface {
	UploadFile(stream filev1.FileService_UploadFileServer) error
	DownloadFile(req *filev1.DownloadFileRequest, stream filev1.FileService_DownloadFileServer) error
	DeleteFile(ctx context.Context, req *filev1.DeleteFileRequest) (*filev1.DeleteFileResponse, error)
	GetFileInfo(ctx context.Context, req *filev1.GetFileInfoRequest) (*filev1.GetFileInfoResponse, error)
	CopyFile(ctx context.Context, req *filev1.CopyFileRequest) (*filev1.CopyFileResponse, error)
	MoveFile(ctx context.Context, req *filev1.MoveFileRequest) (*filev1.MoveFileResponse, error)
	ListFiles(ctx context.Context, req *filev1.ListFilesRequest) (*filev1.ListFilesResponse, error)
	UpdateFileMetadata(ctx context.Context, req *filev1.UpdateFileMetadataRequest) (*filev1.UpdateFileMetadataResponse, error)
	SearchFiles(ctx context.Context, req *filev1.SearchFilesRequest) (*filev1.SearchFilesResponse, error)
	GetStorageInfo(ctx context.Context, req *filev1.GetStorageInfoRequest) (*filev1.GetStorageInfoResponse, error)
	CleanupFiles(ctx context.Context, req *filev1.CleanupFilesRequest) (*filev1.CleanupFilesResponse, error)
	HealthCheck(ctx context.Context, req *filev1.HealthCheckRequest) (*filev1.HealthCheckResponse, error)
}

// ServiceRegistry - registry for all services
type ServiceRegistry interface {
	WebService() WebServiceInterface
	EngineService() EngineServiceInterface
	FileService() FileServiceInterface
	HealthCheck(ctx context.Context) error
}

