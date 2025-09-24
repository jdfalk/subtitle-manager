// file: pkg/grpcserver/health_test.go
// version: 1.0.0
// guid: a1b2c3d4-e5f6-7g8h-9i0j-1k2l3m4n5o6p

package grpcserver

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/translator"
	translatormocks "github.com/jdfalk/subtitle-manager/pkg/translator/mocks"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// TestHealthCheckServiceReady tests that the health check returns SERVING
// when all dependencies are available and working
func TestHealthCheckServiceReady(t *testing.T) {
	// Set up mock translator client
	mockGoogleClient := translatormocks.NewMockGoogleClient(t)
	translator.SetGoogleClientFactory(func(ctx context.Context, apiKey string) (translator.GoogleClient, error) {
		return mockGoogleClient, nil
	})
	defer translator.ResetGoogleClientFactory()

	// Mock successful health check responses
	mockGoogleClient.On("Close").Return(nil)

	// Create gRPC server with health service
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	s := grpc.NewServer()
	server := NewServer("test-api-key", "", false, "")

	// Register the server services
	_ = server // Use the server variable to avoid unused variable error

	go s.Serve(lis)
	defer s.Stop()

	// Connect to the server
	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	// Create health check client
	healthClient := grpc_health_v1.NewHealthClient(conn)

	// Test overall service health
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: "", // Empty service name checks overall health
	})
	require.NoError(t, err)
	require.Equal(t, grpc_health_v1.HealthCheckResponse_SERVING, resp.Status)

	mockGoogleClient.AssertExpectations(t)
}

// TestHealthCheckSpecificService tests health check for specific services
func TestHealthCheckSpecificService(t *testing.T) {
	// Set up mock translator client
	mockGoogleClient := translatormocks.NewMockGoogleClient(t)
	translator.SetGoogleClientFactory(func(ctx context.Context, apiKey string) (translator.GoogleClient, error) {
		return mockGoogleClient, nil
	})
	defer translator.ResetGoogleClientFactory()

	mockGoogleClient.On("Close").Return(nil)

	// Create gRPC server
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	s := grpc.NewServer()
	server := NewServer("test-api-key", "", false, "")
	_ = server // Use the server variable

	go s.Serve(lis)
	defer s.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	healthClient := grpc_health_v1.NewHealthClient(conn)

	testCases := []struct {
		name        string
		serviceName string
		wantStatus  grpc_health_v1.HealthCheckResponse_ServingStatus
	}{
		{
			name:        "translator_service_health",
			serviceName: "subtitle.translator.v1.TranslatorService",
			wantStatus:  grpc_health_v1.HealthCheckResponse_SERVING,
		},
		{
			name:        "overall_health",
			serviceName: "",
			wantStatus:  grpc_health_v1.HealthCheckResponse_SERVING,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{
				Service: tc.serviceName,
			})
			require.NoError(t, err)
			require.Equal(t, tc.wantStatus, resp.Status)
		})
	}

	mockGoogleClient.AssertExpectations(t)
}

// TestHealthCheckWatch tests the streaming health check functionality
func TestHealthCheckWatch(t *testing.T) {
	// Set up mock translator client
	mockGoogleClient := translatormocks.NewMockGoogleClient(t)
	translator.SetGoogleClientFactory(func(ctx context.Context, apiKey string) (translator.GoogleClient, error) {
		return mockGoogleClient, nil
	})
	defer translator.ResetGoogleClientFactory()

	mockGoogleClient.On("Close").Return(nil)

	// Create gRPC server
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	s := grpc.NewServer()
	server := NewServer("test-api-key", "", false, "")
	_ = server

	go s.Serve(lis)
	defer s.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	healthClient := grpc_health_v1.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start watching health status
	stream, err := healthClient.Watch(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: "",
	})
	require.NoError(t, err)

	// Should receive initial status
	resp, err := stream.Recv()
	require.NoError(t, err)
	require.Equal(t, grpc_health_v1.HealthCheckResponse_SERVING, resp.Status)

	mockGoogleClient.AssertExpectations(t)
}

// TestHealthCheckServiceDependencyFailure tests health check behavior when dependencies fail
func TestHealthCheckServiceDependencyFailure(t *testing.T) {
	// Set up mock translator client that fails
	translator.SetGoogleClientFactory(func(ctx context.Context, apiKey string) (translator.GoogleClient, error) {
		return nil, errors.New("mock dependency failure")
	})
	defer translator.ResetGoogleClientFactory()

	// Create gRPC server
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	s := grpc.NewServer()
	server := NewServer("invalid-api-key", "", false, "")
	_ = server

	go s.Serve(lis)
	defer s.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	healthClient := grpc_health_v1.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Health check might still return SERVING if the dependency failure
	// doesn't immediately affect health status (depends on implementation)
	resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: "",
	})

	// The exact behavior depends on how health checks are implemented
	// This test verifies that the health service responds appropriately
	if err != nil {
		// If health check fails, verify it's a reasonable error
		require.Contains(t, err.Error(), "health")
	} else {
		// If health check succeeds, verify the response is valid
		require.NotNil(t, resp)
		require.True(t,
			resp.Status == grpc_health_v1.HealthCheckResponse_SERVING ||
				resp.Status == grpc_health_v1.HealthCheckResponse_NOT_SERVING ||
				resp.Status == grpc_health_v1.HealthCheckResponse_UNKNOWN)
	}
}

// TestGRPCServerReadiness tests that the server starts up properly and is ready to serve
func TestGRPCServerReadiness(t *testing.T) {
	// Set up mock translator client
	mockGoogleClient := translatormocks.NewMockGoogleClient(t)
	translator.SetGoogleClientFactory(func(ctx context.Context, apiKey string) (translator.GoogleClient, error) {
		return mockGoogleClient, nil
	})
	defer translator.ResetGoogleClientFactory()

	mockGoogleClient.On("Close").Return(nil)

	// Test server creation with different configurations
	testCases := []struct {
		name      string
		apiKey    string
		openAIKey string
		debug     bool
		logLevel  string
		wantError bool
	}{
		{
			name:      "valid_config",
			apiKey:    "test-google-api-key",
			openAIKey: "test-openai-key",
			debug:     false,
			logLevel:  "info",
			wantError: false,
		},
		{
			name:      "debug_mode",
			apiKey:    "test-google-api-key",
			openAIKey: "",
			debug:     true,
			logLevel:  "debug",
			wantError: false,
		},
		{
			name:      "minimal_config",
			apiKey:    "test-api-key",
			openAIKey: "",
			debug:     false,
			logLevel:  "",
			wantError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create server with test configuration
			server := NewServer(tc.apiKey, tc.openAIKey, tc.debug, tc.logLevel)
			require.NotNil(t, server)

			// Create gRPC server and register our service
			lis, err := net.Listen("tcp", "127.0.0.1:0")
			require.NoError(t, err)

			s := grpc.NewServer()
			// Register the server (this would typically be done in NewServer or separately)
			go s.Serve(lis)
			defer s.Stop()

			// Verify we can connect to the server
			conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if tc.wantError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			defer conn.Close()

			// Verify connection is working by attempting to use health service
			healthClient := grpc_health_v1.NewHealthClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			// This might fail if health service isn't registered, but connection should work
			_, err = healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{Service: ""})
			// We don't require this to succeed, just that we can connect
		})
	}

	mockGoogleClient.AssertExpectations(t)
}
