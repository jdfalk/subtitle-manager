// file: pkg/services/validation.go
// version: 1.0.0
// guid: 1234abcd-5678-90ef-1234-567890abcdef

package services

import (
	// Generated protobuf packages
	enginev1 "github.com/jdfalk/subtitle-manager/pkg/engine/v1"
	filev1 "github.com/jdfalk/subtitle-manager/pkg/file/v1"
	webv1 "github.com/jdfalk/subtitle-manager/pkg/web/v1"
)

// Interface compatibility validation
// These variables ensure our service interfaces are compatible with the generated gRPC server interfaces

var (
	// Validate WebServiceInterface is compatible with WebServiceServer
	_ webv1.WebServiceServer = (*WebServiceImpl)(nil)

	// Validate EngineServiceInterface is compatible with EngineServiceServer
	_ enginev1.EngineServiceServer = (*EngineServiceImpl)(nil)

	// Validate FileServiceInterface is compatible with FileServiceServer
	_ filev1.FileServiceServer = (*FileServiceImpl)(nil)
)

// Additional interface compatibility checks
var (
	// Validate our custom interfaces can be assigned to our implementations
	_ WebServiceInterface    = (*WebServiceImpl)(nil)
	_ EngineServiceInterface = (*EngineServiceImpl)(nil)
	_ FileServiceInterface   = (*FileServiceImpl)(nil)

	// Validate service registry compatibility
	_ ServiceRegistry = (*ServiceRegistryImpl)(nil)
)
