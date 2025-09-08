<!-- file: docs/tasks/TASK-03-002-all-in-one-mode.md -->
<!-- version: 1.0.0 -->
<!-- guid: 2b3c4d5e-6f7a-8b9c-0d1e-2f3a4b5c6d7e -->

# TASK-03-002: All-in-One Mode Implementation

## Priority: HIGH

## Status: PENDING

## Estimated Time: 6-8 hours

## Overview

Create a single binary mode that runs all services (web, monitoring, translation) in one process without requiring gRPC communication between components. This should be the default deployment mode for simple installations.

## Current Problem

Currently users must run multiple services separately:
- `subtitle-manager web` - Web UI and API
- `subtitle-manager grpc-server` - Translation service
- `subtitle-manager monitor` - File monitoring
- Complex service coordination required

## Requirements

### 1. New All-in-One Command

```bash
# Default all-in-one mode
subtitle-manager serve

# With custom configuration
subtitle-manager serve --web-port 8080 --monitor-paths /media --languages eng,spa
```

### 2. Integrated Services

- **Web Server**: HTTP API and UI on configurable port
- **Translation Engine**: In-process translation (no gRPC needed)
- **File Monitoring**: Unified monitoring service from TASK-03-001
- **Background Workers**: Queue processing, scheduling, etc.

### 3. Configuration Options

```bash
--web-port (default: 8080)          # Web interface port
--monitor-paths (required)          # Comma-separated paths to monitor
--languages (default: eng)          # Languages to download
--disable-web                       # Run without web interface
--disable-monitoring                # Run without file monitoring
--config-file                       # Configuration file path
```

### 4. Service Communication

- **In-Process**: Direct function calls instead of gRPC
- **Shared State**: Single database connection pool
- **Event Bus**: Internal event system for service coordination
- **No Network Overhead**: Eliminate inter-service network calls

## Implementation Steps

1. **Create serve command**
   - New `cmd/serve.go` command
   - Parse all-in-one configuration flags
   - Initialize service managers

2. **Refactor service interfaces**
   - Create service interface abstractions
   - Implement in-process variants
   - Maintain gRPC interfaces for distributed mode

3. **Shared resource management**
   - Single database connection pool
   - Shared configuration management
   - Unified logging and metrics

4. **Service lifecycle management**
   - Graceful startup/shutdown
   - Health checking
   - Error handling and recovery

## Architecture

```
┌─────────────────────────────────────┐
│           All-in-One Mode           │
├─────────────────────────────────────┤
│  Web Server (HTTP/WebSocket)       │
│  ├─ REST API                       │
│  ├─ WebUI Assets                   │
│  └─ WebSocket Events               │
├─────────────────────────────────────┤
│  Translation Engine                 │
│  ├─ Direct Function Calls          │
│  ├─ Google Translate API           │
│  └─ OpenAI API                     │
├─────────────────────────────────────┤
│  File Monitoring Service           │
│  ├─ Filesystem Watcher             │
│  ├─ Periodic Scanner               │
│  └─ Sonarr/Radarr Integration      │
├─────────────────────────────────────┤
│  Shared Resources                  │
│  ├─ Database Pool                  │
│  ├─ Configuration                  │
│  ├─ Logging                        │
│  └─ Metrics                        │
└─────────────────────────────────────┘
```

## Files to Create/Modify

- `cmd/serve.go` - New all-in-one command
- `pkg/services/` - Service interface abstractions
- `pkg/services/web/` - Web service implementation
- `pkg/services/translation/` - In-process translation
- `pkg/services/monitoring/` - Monitoring service
- `pkg/coordinator/` - Service coordination logic

## Success Criteria

- ✅ Single command starts all services
- ✅ No gRPC communication overhead in all-in-one mode
- ✅ Same functionality as distributed mode
- ✅ Graceful startup and shutdown
- ✅ Shared resource management
- ✅ Configuration through single file or flags
- ✅ Performance equal or better than distributed mode

## Dependencies

- TASK-03-001 (Unified monitoring service)

## Related Tasks

- TASK-03-003 (Service renaming)
- TASK-03-004 (Distributed mode architecture)
