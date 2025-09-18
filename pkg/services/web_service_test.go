// file: pkg/services/web_service_test.go
// version: 1.0.1
// guid: 0b8f0a9e-2c2a-4c3f-9f2e-8ef7d6a5b4c3

package services

import (
    "context"
    "testing"

    webv1 "github.com/jdfalk/subtitle-manager/pkg/web/v1"
    "github.com/stretchr/testify/require"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func TestWebService_LogoutUser_Success(t *testing.T) {
    svc := NewWebService()
    resp, err := svc.LogoutUser(context.Background(), &webv1.LogoutUserRequest{})
    require.NoError(t, err)
    require.NotNil(t, resp)
    require.True(t, resp.GetSuccess())
}

func TestWebService_HealthCheck_Success(t *testing.T) {
    svc := NewWebService()
    resp, err := svc.HealthCheck(context.Background(), &webv1.HealthCheckRequest{})
    require.NoError(t, err)
    require.NotNil(t, resp)
}

func TestWebService_UnimplementedMethods(t *testing.T) {
    t.Parallel()
    svc := NewWebService()

    tests := []struct {
        name string
        call func() error
    }{
        {"AuthenticateUser", func() error { _, err := svc.AuthenticateUser(context.Background(), &webv1.AuthenticateUserRequest{}); return err }},
        {"GetUser", func() error { _, err := svc.GetUser(context.Background(), &webv1.GetUserRequest{}); return err }},
        {"UpdateUser", func() error { _, err := svc.UpdateUser(context.Background(), &webv1.UpdateUserRequest{}); return err }},
        {"UpdateUserPreferences", func() error { _, err := svc.UpdateUserPreferences(context.Background(), &webv1.UpdateUserPreferencesRequest{}); return err }},
        {"UploadSubtitle", func() error { _, err := svc.UploadSubtitle(context.Background(), &webv1.UploadSubtitleRequest{}); return err }},
        {"DownloadSubtitle", func() error { _, err := svc.DownloadSubtitle(context.Background(), &webv1.DownloadSubtitleRequest{}); return err }},
        {"SearchSubtitles", func() error { _, err := svc.SearchSubtitles(context.Background(), &webv1.SearchSubtitlesRequest{}); return err }},
        {"TranslateSubtitle", func() error { _, err := svc.TranslateSubtitle(context.Background(), &webv1.TranslateSubtitleRequest{}); return err }},
        {"GetTranslationStatus", func() error { _, err := svc.GetTranslationStatus(context.Background(), &webv1.GetTranslationStatusRequest{}); return err }},
        {"CancelTranslation", func() error { _, err := svc.CancelTranslation(context.Background(), &webv1.CancelTranslationRequest{}); return err }},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            err := tc.call()
            require.Error(t, err)
            require.Equal(t, codes.Unimplemented, status.Code(err))
        })
    }
}
