// file: pkg/services/file_interfaces.go
// version: 1.0.0
// guid: 1f58d9fb-2da4-47ac-b136-14c9f98684f3

package services

import (
	"context"
	"time"

	"github.com/jdfalk/gcommon/sdks/go/v1/common"
	filev1 "github.com/jdfalk/subtitle-manager/pkg/file/v1"
)

// FileServiceInterface defines the interface for the file service layer
// This bridges gRPC protobuf messages with gcommon SDK types for business logic
type FileServiceInterface interface {
	// Health check
	HealthCheck(ctx context.Context, req *filev1.HealthCheckRequest) (*filev1.HealthCheckResponse, error)

	// Add specific methods based on service type here
	// This interface will be expanded with actual service operations
}