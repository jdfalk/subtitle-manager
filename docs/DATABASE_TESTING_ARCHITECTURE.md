// file: docs/DATABASE_TESTING_ARCHITECTURE.md
// version: 1.0.0
// guid: 2b3c4d5e-6f78-9012-3456-789012345678

# Database Testing Architecture: CGO vs Pure Go Builds

## Overview

The subtitle-manager project supports multiple database backends with different build requirements:
- **SQLite**: Requires CGO compilation (build tag: `sqlite`)
- **Pebble**: Pure Go implementation (always available)
- **PostgreSQL**: Requires external database connection

This document describes the comprehensive testing architecture implemented to handle both CGO-enabled and pure Go builds.

## Build Configuration Detection

### SQLite Availability

The system uses build tags to conditionally compile SQLite support:

```go
// +build sqlite
func HasSQLite() bool { return true }

// +build !sqlite
func HasSQLite() bool { return false }
```

### Backend Selection Logic

```go
func OpenStore(path, backend string) (SubtitleStore, error) {
    switch backend {
    case "pebble":
        return OpenPebble(path)
    case "postgres":
        return OpenPostgresStore(path)
    default:
        return OpenSQLStore(path) // Falls back to error if SQLite unavailable
    }
}
```

## Test Architecture

### 1. Backend Selection Tests (`TestBackendSelectionAndCGOSupport`)

**Purpose**: Validates proper backend selection based on build configuration

**Features**:
- Detects SQLite availability using `HasSQLite()`
- Tests all available backends dynamically
- Provides clear build instruction guidance
- Handles expected failures gracefully

**Test Matrix**:
- **CGO Build** (`go build -tags sqlite`): Tests SQLite + Pebble
- **Pure Go Build** (`go build`): Tests only Pebble, expects SQLite failure

### 2. Interface Compatibility Tests (`TestSubtitleStoreInterface`)

**Purpose**: Ensures all backends implement the `SubtitleStore` interface correctly

**Features**:
- Automatically selects best available backend (SQLite if available, otherwise Pebble)
- Tests all interface methods comprehensively
- Handles implementation-specific limitations (e.g., authentication field population)

### 3. Factory Function Tests (`TestOpenStore`)

**Purpose**: Tests the store factory with various backend configurations

**Features**:
- Dynamic test cases based on SQLite availability
- Proper error handling for unavailable backends
- Path handling specific to each backend type

## Implementation Details

### CGO-Enabled Testing

When building with CGO support:
```bash
go test -tags sqlite ./pkg/database -v
```

**Expected behavior**:
- SQLite backend works with `:memory:` or file paths
- Pebble backend works normally
- PostgreSQL fails with connection errors (expected)

### Pure Go Testing

When building without CGO:
```bash
go test ./pkg/database -v
```

**Expected behavior**:
- SQLite backend fails with clear error message
- Pebble backend works normally as fallback
- Error messages guide users to use Pebble or enable CGO

## Test Coverage Achievements

### ✅ Comprehensive Interface Testing
- All 50+ `SubtitleStore` methods tested
- Covers subtitles, downloads, media, tags, monitoring, and authentication
- Implementation-agnostic testing with backend auto-selection

### ✅ Factory Pattern Testing
- Tests `OpenStore()` and `OpenStoreWithConfig()` functions
- Validates backend selection logic
- Proper error handling for invalid configurations

### ✅ Build Configuration Testing
- Automated detection of SQLite availability
- Clear logging of build configuration and capabilities
- Guidance for enabling missing features

### ✅ Resource Management Testing
- Store cleanup and resource deallocation
- Graceful handling of double-close scenarios
- Memory leak prevention

## Usage Examples

### Testing with SQLite Support
```bash
# Enable CGO and SQLite driver
go build -tags sqlite
go test -tags sqlite ./pkg/database -v -run TestBackendSelection
```

### Testing Pure Go Build
```bash
# Disable CGO for pure Go compatibility
CGO_ENABLED=0 go build
go test ./pkg/database -v -run TestBackendSelection
```

### Continuous Integration
```bash
# Test both configurations in CI
go test ./pkg/database -v                    # Pure Go
go test -tags sqlite ./pkg/database -v      # CGO enabled
```

## Key Benefits

1. **Build Flexibility**: Supports both CGO and pure Go deployments
2. **Clear Error Messages**: Users understand why SQLite isn't available
3. **Automatic Fallback**: Tests use best available backend automatically
4. **Comprehensive Coverage**: All interface methods tested across backends
5. **CI-Friendly**: Tests work in both local and CI environments

## Known Limitations

### Authentication Field Population
Some authentication operations may have implementation-specific behavior across backends. Tests accommodate these differences by making authentication assertions lenient where appropriate.

### Double-Close Handling
Different backends handle double-close differently:
- SQLite: Returns error
- Pebble: Panics (by design)
- Tests use recovery mechanisms to handle both approaches

## Future Enhancements

1. **Cross-Backend Migration Testing**: Validate data migration between SQLite and Pebble
2. **Performance Benchmarking**: Compare backend performance characteristics
3. **Concurrent Access Testing**: Multi-threaded access patterns
4. **Data Consistency Validation**: Ensure identical behavior across backends
