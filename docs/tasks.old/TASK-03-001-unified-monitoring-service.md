<!-- file: docs/tasks/TASK-03-001-unified-monitoring-service-CONSOLIDATED.md -->
<!-- version: 1.0.0 -->
<!-- guid: 1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d -->

# TASK-03-001: Unified Monitoring + Service Renaming (CONSOLIDATED)

## Priority: HIGH

## Status: PENDING

## Estimated Time: 6-8 hours

## Overview

**CONSOLIDATED TASK**: This combines unified monitoring service creation with service renaming for efficiency. Instead of doing these as separate tasks, we'll rename services while refactoring monitoring functionality.

## Part A: Unified Monitoring Service

### Current Problem

Currently we have three confusing, overlapping commands:

- `autoscan` - Periodic scanning with time intervals
- `watch` - Real-time file system event monitoring
- `monitor` - Sonarr/Radarr integration with sub-commands

This creates confusion and requires users to understand the differences and run multiple services.

### Unified Monitoring Solution

Create single `subtitle-manager monitor` command that:

- **Default Behavior**: Enable all three monitoring modes simultaneously
- **Selective Control**: Allow individual modes to be disabled via flags
- **Unified Configuration**: Single configuration for all monitoring types

#### Command Interface

```bash
# Default: Enable all monitoring modes
subtitle-manager monitor /path/to/media --languages eng,spa

# Selective modes
subtitle-manager monitor /path/to/media --languages eng --no-filesystem --no-periodic
subtitle-manager monitor /path/to/media --languages eng --no-sonarr --no-radarr
```

#### Configuration Flags

- `--periodic-interval` (default: 1h) - Time-based scanning
- `--no-periodic` - Disable periodic scanning
- `--no-filesystem` - Disable real-time file watching
- `--no-sonarr` - Disable Sonarr integration
- `--no-radarr` - Disable Radarr integration
- `--recursive` - Watch directories recursively
- `--languages` - Comma-separated language codes

## Part B: Service Renaming

### Service Naming Problem

Current service names are confusing:

- `grpc-server` - Should describe function (translation), not protocol
- `grpc-set-config` - Implementation detail in name
- Mixed protocol and functional naming

### Service Renaming Plan

| Current Name      | New Name     | Reason                           |
| ----------------- | ------------ | -------------------------------- |
| `grpc-server`     | `translator` | Describes function, not protocol |
| `grpc-set-config` | `config set` | Organized under config namespace |
| `web`             | `web`        | Keep - clear and functional      |
| `monitor`         | `monitor`    | Keep - clear after unification   |

### New Command Structure

```bash
# Translation service
subtitle-manager translator --port 50051

# Configuration management
subtitle-manager config set google_api_key "key123"
subtitle-manager config get google_api_key
subtitle-manager config list

# Unified monitoring
subtitle-manager monitor /path/to/media --languages eng

# All-in-one mode (for next task)
subtitle-manager serve
```

## Implementation Steps

### Step 1: Create Unified Monitor Command (3-4 hours)

1. **Create new unified monitor command**
   - Replace existing `cmd/monitor.go` structure
   - Implement unified flag parsing
   - Create service orchestrator

2. **Integrate existing functionality**
   - Move `autoscan` periodic logic into monitor
   - Move `watch` filesystem logic into monitor
   - Keep existing Sonarr/Radarr integration
   - Ensure no functionality is lost

### Step 2: Rename Services Simultaneously (2-3 hours)

1. **Create new command structure**
   - Add `cmd/translator.go` (replaces grpc-server)
   - Create `cmd/config/` package for config commands
   - Add aliases for old commands with deprecation warnings

2. **Update service interfaces**
   - Rename internal service structures
   - Update gRPC service registration
   - Maintain protocol compatibility

### Step 3: Update CLI and Documentation (1 hour)

1. **Update CLI interface**
   - Remove old `autoscan` and `watch` commands
   - Add deprecation warnings for old service names
   - Update help text and documentation

2. **Testing**
   - Test all monitoring modes individually
   - Test combined mode operation
   - Verify no performance degradation
   - Test backward compatibility

## Files to Create/Modify

### New Files

- `cmd/translator.go` - New translation service command
- `cmd/config/` - Config command package
- `pkg/monitoring/unified.go` - Unified monitoring service

### Modified Files

- `cmd/monitor.go` - Complete rewrite for unified functionality
- `cmd/autoscan.go` - Remove or add deprecation
- `cmd/watch.go` - Remove or add deprecation
- `cmd/grpcserver_cmd.go` - Add deprecation, alias to translator
- `cmd/grpcsetconfig.go` - Move to config package
- `cmd/root.go` - Update command registration
- `pkg/scheduler/` - Integration updates
- `pkg/watcher/` - Integration updates

## Success Criteria

**Unified Monitoring:**

- ✅ Single `monitor` command replaces three separate commands
- ✅ All previous functionality preserved
- ✅ Default behavior enables all monitoring modes
- ✅ Individual modes can be disabled via flags
- ✅ Performance is equal or better than separate commands

**Service Renaming:**

- ✅ Clear, semantic service names
- ✅ Organized command structure
- ✅ Backward compatibility maintained
- ✅ Deprecation warnings for old commands
- ✅ Updated documentation

## Dependencies

- None - standalone refactor

## Related Tasks

- TASK-03-002 (All-in-one mode) - Requires this task
- TASK-03-004 (Distributed architecture) - Uses renamed services

## Benefits of Consolidation

1. **Efficiency**: Do renaming while refactoring instead of separate pass
2. **Consistency**: Ensure all service names are updated together
3. **Testing**: Single test cycle for both changes
4. **Documentation**: Update all docs once instead of multiple times
5. **User Experience**: Single update for users instead of multiple changes
