<!-- file: docs/tasks/TASK-03-000-architecture-refactor-overview.md -->
<!-- version: 1.0.0 -->
<!-- guid: 0a1b2c3d-4e5f-6a7b-8c9d-0e1f2a3b4c5d -->

# TASK-03-000: Architecture Refactor Overview

## Priority: HIGH

## Status: PENDING

## Estimated Time: 30-40 hours total

## Overview

This is the master task that coordinates the complete architecture refactor of
subtitle-manager to address command line complexity, service organization, and
deployment flexibility issues.

## Problem Statement

Current subtitle-manager has several architectural issues:

1. **Command Line Chaos**: 40+ top-level commands creating confusion
2. **Service Fragmentation**: 3 separate monitoring commands doing similar
   things
3. **Deployment Complexity**: Must run 4+ separate services for full
   functionality
4. **Protocol Confusion**: Service names reference implementation (grpc-server)
   not function
5. **Poor Service Boundaries**: Functions scattered across services illogically

## Solution Architecture

### Deployment Modes

1. **All-in-One Mode** (Default)
   - Single binary, single process
   - In-process communication (no gRPC overhead)
   - Perfect for single-server deployments

2. **Distributed Mode** (Advanced)
   - Specialized service components
   - gRPC communication between services
   - Coordinator-based service discovery
   - Optimized for multi-server deployments

### Service Components

```
All-in-One Mode:
┌─────────────────────────────┐
│     subtitle-manager serve  │
│  ┌─────────┬──────────────┐ │
│  │   Web   │ Translation  │ │
│  │ Service │   Engine     │ │
│  ├─────────┼──────────────┤ │
│  │ File I/O│  Monitoring  │ │
│  │ Engine  │   Service    │ │
│  └─────────┴──────────────┘ │
└─────────────────────────────┘

Distributed Mode (DMZ + Backend):
┌─────────────┐    ┌─────────────┐
│   DMZ Web   │    │   DMZ Web   │
│   Server    │    │   Server    │
└──────┬──────┘    └──────┬──────┘
       │ gRPC             │ gRPC
       └──────────┬───────┘
                  │ Firewall
       ┌──────────▼───────────────┐
       │    Backend Network       │
       │ ┌─────────┬─────────────┐│
       │ │Service  │ File Server ││
       │ │Discovery│             ││
       │ └─────────┴─────────────┘│
       │ ┌─────────┬─────────────┐│
       │ │Translator│ File Server ││
       │ │ Service │             ││
       │ └─────────┴─────────────┘│
       └─────────────────────────┘
```

## Task Dependencies and Execution Order

### Phase 1: Core Service Refactoring (6-8 hours)

**TASK-03-001**: Unified Monitoring + Service Renaming

- Consolidate autoscan, watch, monitor into single service
- Rename grpc-server → translator (do renaming while refactoring)
- Create semantic service names
- Update CLI structure simultaneously

### Phase 2: All-in-One Mode (6-8 hours)

**TASK-03-002**: All-in-One Mode Implementation

- Requires: TASK-03-001 completed
- Create `subtitle-manager serve` command
- In-process service communication
- Foundation for distributed mode

### Phase 3: Distributed Architecture (12-16 hours)

**TASK-03-004**: Secure Distributed Architecture + Function Redistribution

- Requires: TASK-03-002 completed
- Design DMZ web servers + backend file servers
- Implement gRPC-only file servers (no web interface)
- Move functions to appropriate services simultaneously
- Use gRPC service discovery mechanisms

### Phase 4: Quality Assurance (12-16 hours)

**TASK-02-002-B**: Comprehensive Selenium Testing

- Validate all web functionality works
- Multi-platform testing with screenshots/video
- Performance validation

## Expected Outcomes

### User Experience Improvements

**Before Refactor:**

```bash
# Confusing - what's the difference?
subtitle-manager autoscan /media eng --interval 1h
subtitle-manager watch /media eng --recursive
subtitle-manager monitor start

# Must run 4 separate terminals
subtitle-manager web &
subtitle-manager grpc-server &
subtitle-manager autoscan /media eng &
subtitle-manager monitor start &
```

**After Refactor:**

```bash
# Simple - all monitoring unified
subtitle-manager monitor /media --languages eng

# Single command for everything
subtitle-manager serve --monitor-paths /media --languages eng

# Secure distributed deployment
# DMZ Web Server
subtitle-manager web --discovery backend:50050

# Backend Services (behind firewall)
subtitle-manager file-service --paths /media/movies
subtitle-manager translator --workers 4
```

### Technical Improvements

- **50% reduction** in required tasks (6 → 3 consolidated tasks)
- **Security by design** with DMZ web servers and backend file servers
- **75% reduction** in required terminals (4 → 1 for basic use)
- **Unified service names** that describe function, not protocol
- **gRPC-only file servers** for enhanced security
- **Flexible deployment** from single-server to enterprise scale

## Success Criteria

- ✅ Single `serve` command runs complete system
- ✅ Unified `monitor` command replaces 3 separate commands
- ✅ Organized command structure with logical grouping
- ✅ Semantic service names (translator, not grpc-server)
- ✅ Both all-in-one and distributed deployment modes
- ✅ All existing functionality preserved
- ✅ Comprehensive test coverage validates changes
- ✅ Performance equal or better than current implementation

## Risk Mitigation

1. **Backward Compatibility**: All tasks maintain aliases for old commands
2. **Incremental Changes**: Each task can be completed independently
3. **Testing**: Comprehensive selenium tests validate no regressions
4. **Documentation**: Migration guides for users upgrading

## Resource Requirements

- **Development Time**: 26-34 hours across 3 consolidated phases (down from 6
  tasks)
- **Testing Environment**: Docker setup for multi-service testing with DMZ
  simulation
- **Documentation Updates**: README, API docs, security guides

## Consolidated Task Benefits

1. **Efficiency**: 50% fewer tasks by combining related work
2. **Consistency**: Ensure renaming and refactoring happen together
3. **Security**: Proper DMZ/backend architecture from the start
4. **Testing**: Single test cycle for related changes
5. **User Experience**: Fewer disruptive updates

This refactor will transform subtitle-manager from a complex, fragmented tool
into a clean, intuitive, and flexible subtitle management system suitable for
both simple and enterprise deployments.
