# TASK CORRECTION: Use gcommon Protobuf Definitions

Based on the analysis, our current task implementations are incorrectly defining custom protobuf types instead of leveraging the extensive gcommon protobuf library. This document outlines the necessary corrections.

## Issues Identified

### 1. Custom Types Instead of gcommon
- **Current**: Defining `proto/common/v1/error.proto`, `proto/common/v1/timestamp.proto`, etc.
- **Should Use**: `github.com/jdfalk/gcommon/proto/common/v1/error.proto`, etc.

### 2. Missing gcommon Imports
- **Current**: Creating our own common types
- **Should Use**: Import and leverage existing gcommon types

### 3. Configuration Management
- **Current**: Custom `Config` structs
- **Should Use**: `gcommon/config` package types

## Available gcommon Types We Should Use

Based on the gcommon documentation:

### From gcommon/common:
- `User` - User management
- `Session` - Session handling
- `Error` - Standardized error responses
- `Timestamp` - Time handling
- `Metadata` - Request/response metadata
- `APIKey` - API key authentication
- `Role` & `Permission` - Authorization

### From gcommon/config:
- `ApplicationConfig` - Application configuration
- `ServerConfig` - Server settings
- `DatabaseConfig` - Database configuration
- Various environment and deployment configs

### From gcommon/health:
- `HealthCheck` - Health monitoring
- `HealthStatus` - Status reporting

### From gcommon/media:
- Media processing types for subtitle/video handling

## Required Corrections

### 1. Update TASK-01-001 Service Interface Definitions

**Remove custom common types and use gcommon imports:**

```protobuf
// Instead of defining custom types, import gcommon
import "gcommon/v1/common/user.proto";
import "gcommon/v1/common/session.proto";
import "gcommon/v1/common/error.proto";
import "gcommon/v1/config/application_config.proto";
```

### 2. Update TASK-02-001 Web Service Implementation

**Use gcommon types for:**
- Authentication: `gcommon.User`, `gcommon.Session`
- Configuration: `gcommon.config.ApplicationConfig`
- Error handling: `gcommon.Error`
- Health checks: `gcommon.health.HealthCheck`

### 3. Update All Service Definitions

**Replace custom types with gcommon equivalents:**
- User management → `gcommon/common`
- Configuration → `gcommon/config`
- Health monitoring → `gcommon/health`
- Database operations → `gcommon/database`

## Implementation Strategy

1. **Audit Current Tasks**: Review all existing task files for custom protobuf definitions
2. **Map to gcommon**: Create mapping of custom types to gcommon equivalents
3. **Update Imports**: Replace custom proto imports with gcommon imports
4. **Update Go Code**: Use gcommon SDKs with opaque API patterns
5. **Validate**: Ensure all services use consistent gcommon types

## Next Steps

1. **Immediate**: Update TASK-01-001 to use gcommon imports instead of custom types
2. **Follow-up**: Update TASK-02-001 and subsequent tasks to leverage gcommon
3. **Validate**: Ensure all protobuf definitions follow gcommon patterns
4. **Test**: Verify opaque API usage is consistent throughout

This correction ensures we're properly leveraging the extensive gcommon protobuf library instead of duplicating functionality.
