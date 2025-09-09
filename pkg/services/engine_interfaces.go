// file: pkg/services/engine_interfaces.go
// version: 1.0.0
// guid: e65cc565-9d1b-460c-9063-db7a2a0b8c9b

package services

import (
	"context"
	"time"

	"github.com/jdfalk/gcommon/sdks/go/v1/common"
	enginev1 "github.com/jdfalk/subtitle-manager/pkg/engine/v1"
)

// EngineServiceInterface defines the interface for the engine service layer
// This bridges gRPC protobuf messages with gcommon SDK types for business logic
type EngineServiceInterface interface {
	// Health check
	HealthCheck(ctx context.Context, req *enginev1.HealthCheckRequest) (*enginev1.HealthCheckResponse, error)

	// Add specific methods based on service type here
	// This interface will be expanded with actual service operations
}