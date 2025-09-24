<!-- file: TEST_ANALYSIS.md -->
<!-- version: 1.0.0 -->
<!-- guid: a1b2c3d4-e5f6-7890-abcd-ef1234567890 -->

# Comprehensive Test Analysis and Test Implementation Plan

## Overview

This document provides a complete analysis of the current test status in the subtitle-manager repository, identifying failing tests and files that need test coverage. This serves as a comprehensive instruction document for implementing missing tests and fixing failing ones.

**IMPORTANT: Repository Status Confirmed**
- ✅ All required protobuf packages are generated and present in `pkg/engine/v1/`, `pkg/file/v1/`, `pkg/subtitle/translator/v1/`, `pkg/web/v1/`
- ✅ The Go build system works correctly: `go build ./...` succeeds
- ✅ Tests can be executed: `go test ./...` runs (with expected failures)
- ✅ `buf generate` is working and protobuf files are up to date
- 🎯 **Ready for comprehensive test implementation work**

## Quick Verification Commands

To verify the repository is in working order before starting test implementation:

```bash
# Verify Go build works
go build ./...

# Verify protobuf generation works  
buf generate

# Verify tests can run (will show expected failures)
go test ./pkg/authserver -v
go test ./pkg/cache -v  
go test ./pkg/grpcserver -v
go test ./pkg/providers -v

# Check test coverage for working packages
go test ./pkg/webserver -cover
```

## Current Test Status Summary

### Test Statistics
- **Total Go source files**: 271
- **Total test files**: 172
- **Test coverage ratio**: ~63.5%

### Package Test Results

#### ✅ PASSING PACKAGES (70 packages)
```
ok      github.com/jdfalk/subtitle-manager/cmd
ok      github.com/jdfalk/subtitle-manager/pkg/audio
ok      github.com/jdfalk/subtitle-manager/pkg/backups
ok      github.com/jdfalk/subtitle-manager/pkg/bazarr
ok      github.com/jdfalk/subtitle-manager/pkg/captcha
ok      github.com/jdfalk/subtitle-manager/pkg/cli
ok      github.com/jdfalk/subtitle-manager/pkg/config
ok      github.com/jdfalk/subtitle-manager/pkg/database
ok      github.com/jdfalk/subtitle-manager/pkg/errors
ok      github.com/jdfalk/subtitle-manager/pkg/events
ok      github.com/jdfalk/subtitle-manager/pkg/gcommonauth
ok      github.com/jdfalk/subtitle-manager/pkg/gcommonlog
ok      github.com/jdfalk/subtitle-manager/pkg/i18n
ok      github.com/jdfalk/subtitle-manager/pkg/logging
ok      github.com/jdfalk/subtitle-manager/pkg/maintenance
ok      github.com/jdfalk/subtitle-manager/pkg/metadata
ok      github.com/jdfalk/subtitle-manager/pkg/metrics
ok      github.com/jdfalk/subtitle-manager/pkg/monitoring
ok      github.com/jdfalk/subtitle-manager/pkg/notifications
ok      github.com/jdfalk/subtitle-manager/pkg/plex
ok      github.com/jdfalk/subtitle-manager/pkg/profiles
ok      github.com/jdfalk/subtitle-manager/pkg/queue
ok      github.com/jdfalk/subtitle-manager/pkg/radarr
ok      github.com/jdfalk/subtitle-manager/pkg/renamer
ok      github.com/jdfalk/subtitle-manager/pkg/scanner
ok      github.com/jdfalk/subtitle-manager/pkg/scheduler
ok      github.com/jdfalk/subtitle-manager/pkg/scoring
ok      github.com/jdfalk/subtitle-manager/pkg/security
ok      github.com/jdfalk/subtitle-manager/pkg/selftest
ok      github.com/jdfalk/subtitle-manager/pkg/services
ok      github.com/jdfalk/subtitle-manager/pkg/sonarr
ok      github.com/jdfalk/subtitle-manager/pkg/storage
ok      github.com/jdfalk/subtitle-manager/pkg/subtitles
ok      github.com/jdfalk/subtitle-manager/pkg/syncer
ok      github.com/jdfalk/subtitle-manager/pkg/tagging
ok      github.com/jdfalk/subtitle-manager/pkg/tasks
ok      github.com/jdfalk/subtitle-manager/pkg/testutil
ok      github.com/jdfalk/subtitle-manager/pkg/transcriber
ok      github.com/jdfalk/subtitle-manager/pkg/translator
ok      github.com/jdfalk/subtitle-manager/pkg/updater
ok      github.com/jdfalk/subtitle-manager/pkg/watcher
ok      github.com/jdfalk/subtitle-manager/pkg/webhooks
ok      github.com/jdfalk/subtitle-manager/pkg/webserver
```

Plus 28 individual provider packages under `pkg/providers/*` (all passing).

#### ❌ FAILING PACKAGES (4 packages)

**High Priority Failures - Need Immediate Attention:**

1. **pkg/authserver** - OAuth authentication flow tests
2. **pkg/cache** - Cache functionality tests
3. **pkg/grpcserver** - gRPC server tests
4. **pkg/providers** - Main provider interface tests

---

## Detailed Failure Analysis

### 1. pkg/authserver - OAuth Flow Test Failures

**Issues:**
- Mock expectations not properly set up
- `handleToken` method calls not mocked
- SQLite dependency issues causing skipped tests

**Failing Tests:**
- `TestOAuthTokenExchangeError`
- `TestOAuthUserInfoRetrievalError`

**Error Pattern:**
```
assert: mock: I don't know what to return because the method call was unexpected.
Either do Mock.On("handleToken").Return(...) first, or remove the handleToken() call.
```

**Root Cause:** The mock OAuth provider is not properly configured with expected method calls. The test is making calls to `handleToken` but the mock doesn't have expectations set up for this method.

**Solution Strategy:**
- Fix mock setup in `oauth_flow_test.go`
- Add proper `Mock.On("handleToken").Return(...)` expectations
- Ensure SQLite dependency is properly handled or mocked

### 2. pkg/cache - Cache System Test Failures

**Issues:**
- Long-running tests (10.201s duration suggests timeout or hang)
- Likely cache invalidation or cleanup issues
- Possible race conditions in concurrent cache operations

**Solution Strategy:**
- Investigate cache cleanup between tests
- Add proper test isolation
- Check for goroutine leaks or blocking operations
- Implement timeout controls for cache operations

### 3. pkg/grpcserver - gRPC Server Test Failures

**Issues:**
- Server lifecycle management issues
- Port binding conflicts
- gRPC service registration problems

**Solution Strategy:**
- Implement proper server start/stop in tests
- Use random ports or test-specific ports
- Mock gRPC dependencies where appropriate
- Add proper cleanup in test teardown

### 4. pkg/providers - Provider Interface Test Failures

**Issues:**
- Provider registry conflicts
- Configuration setup issues
- Mock provider implementations

**Solution Strategy:**
- Fix provider registration in tests
- Implement proper provider isolation
- Add comprehensive provider interface tests
- Mock external dependencies

---

## Files Without Test Coverage

The following Go source files do not have corresponding `*_test.go` files and need test coverage:

### Critical Files Needing Tests (High Priority)

```
pkg/gcommonauth/rbac.go
pkg/metrics/metrics_gcommon.go
pkg/database/store.go
pkg/database/store_factory.go
pkg/gcommon/config/config.go
pkg/gcommon/subtitle_format.go
pkg/video/video.go
pkg/types/types.go
```

### Generated/Mock Files (Lower Priority)

```
pkg/database/mocks/taggedentity_mock.go
pkg/database/mocks/subtitlestore_mock.go
pkg/cache/mocks/cache_mock.go
pkg/cache/mocks/statsprovider_mock.go
pkg/proto/config_protoopaque.pb.go
pkg/proto/config.pb.go
```

### Database Driver Files (Platform-Specific)

```
pkg/database/sqlite_disabled.go
pkg/database/drivers_nosqlite.go
pkg/database/drivers_sqlite.go
pkg/database/sqlite_support.go
pkg/database/sqlite_no_support.go
pkg/database/sqlite_enabled.go
```

---

## Go Coding Instructions (Embedded)

<!-- Including complete Go coding instructions for reference -->

## Core Principles

- Clarity over cleverness: Code should be clear and readable
- Simplicity: Prefer simple solutions over complex ones
- Consistency: Follow established patterns within the codebase
- Readability: Code is written for humans to read

## Version Requirements

- **MANDATORY**: All Go projects must use Go 1.23.0 or higher
- **NO EXCEPTIONS**: Do not use older Go versions in any repository
- Update `go.mod` files to specify `go 1.23` minimum version
- Update `go.work` files to specify `go 1.23` minimum version
- All Go file headers must use version 1.23.0 or higher
- Use `go version` to verify your installation meets requirements

## Testing Best Practices

### Test Structure
- Use table-driven tests for multiple scenarios
- Follow Arrange-Act-Assert pattern
- Test file names end with `_test.go`
- Test function names start with `Test`
- Use descriptive test names that explain what is being tested

### Test Organization
```go
func TestFunctionName(t *testing.T) {
    // Arrange
    input := "test data"
    expected := "expected result"

    // Act
    result := FunctionName(input)

    // Assert
    assert.Equal(t, expected, result)
}
```

### Table-Driven Tests
```go
func TestMultipleScenarios(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    "test",
            expected: "TEST",
            wantErr:  false,
        },
        {
            name:     "empty input",
            input:    "",
            expected: "",
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := ProcessString(tt.input)

            if tt.wantErr {
                assert.Error(t, err)
                return
            }

            assert.NoError(t, err)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

### Mock Usage
- Use testify/mock for complex mocking needs
- Set up proper expectations with `Mock.On().Return()`
- Verify expectations with `mock.AssertExpectations(t)`
- Clean up mocks in test teardown

### Test Isolation
- Each test should be independent
- Use setup/teardown functions for resource management
- Reset global state between tests
- Use t.Cleanup() for automatic cleanup

## Error Handling in Tests

- Always check for unexpected errors
- Use `assert.NoError(t, err)` for operations that should succeed
- Use `assert.Error(t, err)` for operations that should fail
- Test specific error types with `assert.ErrorIs(t, err, expectedErr)`

## Required File Header

All Go files must begin with a standard header:

```go
// file: path/to/file.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174000
```

## Naming Conventions

- Use short, concise, evocative package names (lowercase, no underscores)
- Use camelCase for unexported names, PascalCase for exported names
- Single-method interfaces should end in "-er" (e.g., Reader, Writer)

## Code Organization

- Use `goimports` to format imports automatically
- Group imports: standard library, third-party, local
- Keep functions short and focused
- Order: receiver, name, parameters, return values

---

## Implementation Priority

### Phase 1: Critical Failures (Immediate)
1. **Fix pkg/authserver OAuth tests**
   - Repair mock setup in `oauth_flow_test.go`
   - Add missing `handleToken` mock expectations
   - Handle SQLite dependencies properly

2. **Fix pkg/cache test hangs**
   - Add timeout controls
   - Fix cleanup between tests
   - Investigate race conditions

3. **Fix pkg/grpcserver lifecycle**
   - Implement proper server start/stop
   - Use test-specific ports
   - Add cleanup in teardown

4. **Fix pkg/providers interface tests**
   - Repair provider registration
   - Add provider isolation
   - Mock external dependencies

### Phase 2: Missing Test Coverage (High Priority)
1. **Core Business Logic**
   - `pkg/gcommonauth/rbac.go` - RBAC authorization logic
   - `pkg/gcommon/config/config.go` - Configuration management
   - `pkg/database/store.go` - Core data store operations
   - `pkg/database/store_factory.go` - Store factory pattern

2. **Data Types and Models**
   - `pkg/types/types.go` - Core type definitions
   - `pkg/gcommon/subtitle_format.go` - Subtitle format handling
   - `pkg/video/video.go` - Video file operations

3. **Metrics and Monitoring**
   - `pkg/metrics/metrics_gcommon.go` - gcommon metrics integration

### Phase 3: Platform-Specific Tests (Medium Priority)
- Database driver tests for SQLite variants
- Platform-specific code paths

### Phase 4: Generated Code Tests (Low Priority)
- Mock file testing (if needed)
- Protocol buffer generated code testing

---

## Success Criteria

### Test Execution Goals
- All test packages should pass: `go test ./...` returns exit code 0
- No skipped tests due to improper setup
- Test execution time under 30 seconds total
- No race conditions detected with `go test -race ./...`

### Coverage Goals
- Core business logic: 90%+ test coverage
- Public APIs: 100% test coverage
- Error paths: 80%+ test coverage
- Integration points: 95%+ test coverage

### Quality Goals
- All tests follow Arrange-Act-Assert pattern
- Table-driven tests for multiple scenarios
- Proper mock usage with expectation verification
- Clean test isolation and cleanup
- Descriptive test names and documentation

---

## Tools and Commands

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with race detection
go test -race ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./pkg/authserver -v

# Run specific test
go test -run TestOAuthTokenExchangeError ./pkg/authserver -v
```

### Test Development
```bash
# Generate mocks (if using mockery)
mockery --all

# Format code
gofmt -w .
goimports -w .

# Lint code
golangci-lint run
```

---

## Notes for Implementation

### Development Approach
- Fix failing tests first (they block CI/CD)
- Add missing tests for critical paths second
- Use TDD approach for new functionality
- Maintain test isolation and independence

### Common Patterns in Codebase
- Heavy use of testify/assert and testify/mock
- Table-driven tests are preferred
- Provider pattern with interface testing
- Database abstraction with SQLite/non-SQLite variants
- gRPC service testing patterns

### Integration Considerations
- Tests must work in CI environment
- No external dependencies without mocks
- Proper resource cleanup to prevent flaky tests
- Cross-platform compatibility required

This document serves as the complete specification for implementing comprehensive test coverage and fixing all failing tests in the subtitle-manager repository.
