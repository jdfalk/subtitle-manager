# Subtitle Manager Project Status
<!-- file: PROJECT_STATUS.md -->
<!-- version: 1.0.0 -->
<!-- guid: a1b2c3d4-e5f6-7890-abcd-ef1234567890 -->

## Current Status: CRITICAL ISSUES BLOCKING COMPLETION

**Last Updated:** September 1, 2025
**Overall Status:** 🔴 **BROKEN - Multiple P0 Critical Issues**

## Executive Summary

The subtitle-manager project is currently in a **non-functional state** with multiple critical P0 issues that prevent compilation and basic operation. While significant infrastructure and architecture work has been completed, several breaking changes in the protobuf migration and configuration system have left the codebase in an unstable state.

## Critical P0 Issues (MUST FIX IMMEDIATELY)

### 1. **Build System Broken**
- **Issue:** Unused imports in auth middleware (`strconv`, `github.com/jdfalk/gcommon/sdks/go/v1/common`)
- **Impact:** Go compilation fails, server cannot build
- **Location:** Look for auth middleware files in `pkg/webserver/` or similar
- **Fix Required:** Remove unused imports or implement their usage

### 2. **Protobuf Schema Mismatch**
- **Issue:** Generated client code still references deleted `configpb.SubtitleManagerConfig`
- **Impact:** Code does not compile, gRPC services unusable
- **Location:** Generated protobuf files and any code importing `configpb`
- **Fix Required:** Regenerate protobuf files to match new schema

### 3. **Opaque Protobuf Types Outdated**
- **Issue:** Generated opaque types still expose `xxx_hidden_Config` fields for deleted config structure
- **Impact:** Compilation failures, config system broken
- **Location:** Generated opaque protobuf files
- **Fix Required:** Regenerate opaque protobuf files

## High Priority P1 Issues (DATA LOSS RISK)

### 4. **SubtitleRow Data Loss**
- **Issue:** Conversion only populates ID, file, video file, language - drops release, provider metadata, service, embedded flag, modification type
- **Impact:** Critical data lost in gRPC round-trips, monitoring/history broken
- **Location:** Data conversion functions for SubtitleRow
- **Fix Required:** Update conversion to preserve all fields

### 5. **DownloadRow Data Loss**
- **Issue:** Conversion omits VideoFile, Language, SearchQuery, MatchScore, DownloadAttempts, error messages, response time
- **Impact:** Download history and monitoring features broken
- **Location:** Data conversion functions for DownloadRow
- **Fix Required:** Complete conversion implementation

### 6. **SQLite Configuration Lost**
- **Issue:** `sqlite3_filename` key dropped from config migration
- **Impact:** Existing deployments cannot load SQLite file configuration
- **Location:** Configuration migration code
- **Fix Required:** Add sqlite3_filename to config round-trip handling

### 7. **Invalid Protobuf Any Messages**
- **Issue:** Raw string bytes written to `anypb.Any` instead of valid protobuf messages
- **Impact:** Other components cannot unpack timestamp data
- **Location:** Timestamp serialization in data conversion
- **Fix Required:** Properly serialize timestamps as protobuf messages

## Additional Development Environment Issues

### 8. **Issue Analysis Scripts Broken**
- **Issue:** `NameError: name 'IssueAnalyzer' is not defined` in issue automation scripts
- **Impact:** Development tooling not functional
- **Location:** `scripts/issue_automation_wrapper.py`
- **Fix Required:** Fix import issues in Python scripts

## Detailed Action Plan for Coding Agent

### Step 1: Assess Current Build State (30 minutes)
```bash
# Test if project compiles
cd /Users/jdfalk/repos/github.com/jdfalk/subtitle-manager
go build ./...
go test ./... -v

# Check specific package compilation
go build ./pkg/webserver/...
go build ./cmd/...
```

### Step 2: Fix Critical Build Issues (2-4 hours)

#### 2.1 Fix Import Issues
1. Search for unused imports:
   ```bash
   grep -r "strconv" pkg/webserver/
   grep -r "github.com/jdfalk/gcommon/sdks/go/v1/common" pkg/webserver/
   ```
2. Remove unused imports or implement their usage
3. Test compilation after each fix

#### 2.2 Regenerate Protobuf Files
```bash
# Use the existing VS Code task
buf generate

# Or manually if task fails
buf generate --template buf.gen.yaml
```

#### 2.3 Remove configpb References
1. Search for all configpb references:
   ```bash
   grep -r "configpb" . --exclude-dir=.git
   grep -r "SubtitleManagerConfig" . --exclude-dir=.git
   ```
2. Remove or update all references to use new config schema

### Step 3: Fix Data Conversion Issues (4-6 hours)

#### 3.1 Locate Conversion Functions
```bash
# Find data conversion code
find . -name "*.go" -exec grep -l "SubtitleRow\|DownloadRow" {} \;
grep -r "ToProto\|FromProto" pkg/
```

#### 3.2 Fix SubtitleRow Conversion
Fields that MUST be preserved:
- ID ✓ (currently preserved)
- File ✓ (currently preserved)
- VideoFile ✓ (currently preserved)
- Language ✓ (currently preserved)
- **Release** ❌ (MISSING - add this)
- **Provider metadata** ❌ (MISSING - add this)
- **Service** ❌ (MISSING - add this)
- **Embedded flag** ❌ (MISSING - add this)
- **Modification type** ❌ (MISSING - add this)

#### 3.3 Fix DownloadRow Conversion
Fields that MUST be preserved:
- ID ✓ (currently preserved)
- File ✓ (currently preserved)
- Provider ✓ (currently preserved)
- CreatedAt ✓ (currently preserved)
- **VideoFile** ❌ (MISSING - add this)
- **Language** ❌ (MISSING - add this)
- **SearchQuery** ❌ (MISSING - add this)
- **MatchScore** ❌ (MISSING - add this)
- **DownloadAttempts** ❌ (MISSING - add this)
- **Error message** ❌ (MISSING - add this)
- **Response time** ❌ (MISSING - add this)

#### 3.4 Fix Any Message Serialization
Current issue: Raw string bytes written to `anypb.Any`
Required fix: Properly serialize timestamps as protobuf messages

```go
// WRONG (current):
any.Value = []byte(timestampString)

// CORRECT (needed):
timestampProto := &timestamppb.Timestamp{...}
any, _ := anypb.New(timestampProto)
```

### Step 4: Fix Configuration System (2-3 hours)

#### 4.1 Add sqlite3_filename Back
1. Find config migration code:
   ```bash
   grep -r "db_path\|db_backend" pkg/
   grep -r "ToProto\|ApplyProto" pkg/
   ```
2. Add `sqlite3_filename` to both ToProto and ApplyProto functions
3. Test configuration round-trip

### Step 5: Integration Testing (2-3 hours)

#### 5.1 Test Basic Functionality
```bash
# Test compilation
go build ./...
go test ./... -v

# Test gRPC services
go run main.go grpcserver --help

# Test web server
go run main.go web --help
```

#### 5.2 Test Data Integrity
Create test cases that:
1. Create SubtitleRow with all fields
2. Convert to protobuf and back
3. Verify all fields are preserved
4. Repeat for DownloadRow

#### 5.3 Test Configuration System
1. Create config with sqlite3_filename
2. Convert to protobuf
3. Convert back to config
4. Verify sqlite3_filename is preserved

### Step 6: Fix Development Tooling (1 hour)

#### 6.1 Fix Issue Analysis Scripts
```bash
cd scripts/
python3 -c "import issue_analyzer; print('Import works')"
```
Fix any import issues in the Python scripts.

### Step 7: Final Validation (1 hour)

#### 7.1 Full Test Suite
```bash
go test ./... -v -race
go build ./...
```

#### 7.2 Docker Build Test
```bash
docker build -t subtitle-manager .
```

#### 7.3 Integration Test
```bash
# Start services and test basic operations
go run main.go web &
curl http://localhost:8080/health
```

## File Locations Reference

### Key Files to Examine/Fix:
- `pkg/webserver/` - Auth middleware with import issues
- `pkg/grpc/` or similar - gRPC service implementations
- `pkg/database/` or similar - Data conversion functions
- `pkg/config/` or similar - Configuration handling
- `proto/` - Protocol buffer definitions
- Generated protobuf files (likely in `pkg/` subdirectories)

### Build/Proto Files:
- `buf.yaml` - Buf configuration
- `buf.gen.yaml` - Buf generation template
- `go.mod` - Go module dependencies

### Documentation:
- `docs/` - Technical documentation
- `TODO.md` - Current task list
- This file (`PROJECT_STATUS.md`) - Current status

## Success Criteria

The project is considered "finished" when:

1. ✅ **Project builds cleanly** - `go build ./...` succeeds
2. ✅ **All tests pass** - `go test ./... -v` succeeds
3. ✅ **gRPC services functional** - Can start grpcserver without errors
4. ✅ **No data loss in conversions** - All SubtitleRow/DownloadRow fields preserved
5. ✅ **Configuration system works** - sqlite3_filename round-trips correctly
6. ✅ **Web UI connects** - Web server starts and serves pages
7. ✅ **Docker builds** - Container builds successfully
8. ✅ **Development tools work** - Issue analysis scripts function

## Timeline Estimate

**Total: 12-20 hours of focused development work**

- Step 1 (Assessment): 0.5 hours
- Step 2 (Build fixes): 2-4 hours
- Step 3 (Data conversion): 4-6 hours
- Step 4 (Config system): 2-3 hours
- Step 5 (Integration testing): 2-3 hours
- Step 6 (Dev tooling): 1 hour
- Step 7 (Final validation): 1 hour

**Spread over 3-5 days depending on complexity of data conversion fixes.**

## Notes for Coding Agent

1. **Start with Step 1** - Always assess current state before making changes
2. **Fix build issues first** - Nothing else matters if the code doesn't compile
3. **Test incrementally** - Run `go build` after each major change
4. **Focus on data integrity** - The conversion functions are critical for not losing user data
5. **Use existing tasks** - The project has VS Code tasks set up for common operations
6. **Document changes** - Update this file with progress and any discoveries

**The architecture and infrastructure are solid - this is primarily about fixing the data conversion layer and protobuf integration that broke during recent migrations.**
