# TASK-01-001: Replace configpb with gcommon config

<!-- file: docs/tasks/01-gcommon-migration/TASK-01-001-replace-configpb.md -->
<!-- version: 1.0.0 -->
<!-- guid: b1c2d3e4-f5g6-7890-abcd-ef1234567891 -->

## 🎯 Objective

Replace all usage of the local `pkg/configpb` package with the gcommon config
package (`github.com/jdfalk/gcommon/sdks/go/v1/config`).

## 📋 Acceptance Criteria

- [ ] All imports of `github.com/jdfalk/subtitle-manager/pkg/configpb` are
      replaced
- [ ] All config-related types use gcommon protobuf types with opaque API
- [ ] All getter/setter methods use the correct opaque API pattern
- [ ] No compilation errors after migration
- [ ] All tests pass with new config types
- [ ] Local `pkg/configpb` directory can be safely removed

## 🔍 Current State Analysis

### Files Currently Using configpb

Based on codebase analysis, these files import `configpb`:

1. `pkg/translatorpb/translator.pb.go`
2. `pkg/translatorpb/translator_protoopaque.pb.go`
3. `pkg/subtitle/translator/v1/translator.pb.go`
4. `pkg/subtitle/translator/v1/translator_protoopaque.pb.go`
5. `pkg/gcommon/config/config.go`

### Current configpb Types to Replace

From `pkg/configpb/config.pb.go`:

- `LogLevel` enum → `github.com/jdfalk/gcommon/sdks/go/v1/common.LogLevel`
- `SubtitleManagerConfig` →
  `github.com/jdfalk/gcommon/sdks/go/v1/config.ApplicationConfig`
- `DatabaseConfig` →
  `github.com/jdfalk/gcommon/sdks/go/v1/database.DatabaseConfig`
- Other config types as needed

## 🔧 Implementation Steps

### Step 1: Analyze gcommon config package structure

```bash
# Generate documentation for gcommon config package
gomarkdoc --output docs/gcommon-api/config.md github.com/jdfalk/gcommon/sdks/go/v1/config
```

### Step 2: Create migration mapping

Create a mapping file `docs/tasks/01-gcommon-migration/config-migration-map.md`:

```markdown
| Local Type                     | gcommon Type             | Import Path                                 |
| ------------------------------ | ------------------------ | ------------------------------------------- |
| configpb.LogLevel              | common.LogLevel          | github.com/jdfalk/gcommon/sdks/go/v1/common |
| configpb.SubtitleManagerConfig | config.ApplicationConfig | github.com/jdfalk/gcommon/sdks/go/v1/config |
```

### Step 3: Update import statements

For each file in the list:

```go
// OLD
import (
    configpb "github.com/jdfalk/subtitle-manager/pkg/configpb"
)

// NEW
import (
    "github.com/jdfalk/gcommon/sdks/go/v1/config"
    "github.com/jdfalk/gcommon/sdks/go/v1/common"
)
```

### Step 4: Update type usage with opaque API

Replace direct field access with setter/getter methods:

```go
// OLD - Direct field access
config := &configpb.SubtitleManagerConfig{
    LogLevel: configpb.LogLevel_LOG_LEVEL_INFO,
    Database: &configpb.DatabaseConfig{
        Type: "sqlite",
    },
}

// NEW - Opaque API with setters
config := &config.ApplicationConfig{}
config.SetLogLevel(common.LogLevel_LOG_LEVEL_INFO)
dbConfig := &database.DatabaseConfig{}
dbConfig.SetType("sqlite")
config.SetDatabase(dbConfig)
```

### Step 5: Update all usages

For each file:

1. **pkg/gcommon/config/config.go**:
   - Replace `configpb.SubtitleManagerConfig` with `config.ApplicationConfig`
   - Update all field access to use getters/setters
   - Update function signatures

2. **Generated protobuf files**:
   - Regenerate these files after updating the proto definitions
   - Update proto files to import from gcommon instead of local configpb

### Step 6: Update proto files

If any `.proto` files reference the local configpb, update them:

```protobuf
// OLD
import "configpb/config.proto";

// NEW
import "gcommon/v1/config/application_config.proto";
```

### Step 7: Regenerate protobuf files

```bash
# Use VS Code task for protobuf generation
# This will ensure proper logging and error handling
```

Use the VS Code task: `Buf Generate with Output`

### Step 8: Update tests

Update any test files that use configpb types:

```go
// OLD
func TestConfig(t *testing.T) {
    config := &configpb.SubtitleManagerConfig{
        LogLevel: configpb.LogLevel_LOG_LEVEL_DEBUG,
    }
    assert.Equal(t, configpb.LogLevel_LOG_LEVEL_DEBUG, config.LogLevel)
}

// NEW
func TestConfig(t *testing.T) {
    config := &config.ApplicationConfig{}
    config.SetLogLevel(common.LogLevel_LOG_LEVEL_DEBUG)
    assert.Equal(t, common.LogLevel_LOG_LEVEL_DEBUG, config.GetLogLevel())
}
```

## 📚 Required Documentation

### Coding Instructions Reference

**CRITICAL**: Follow these instructions precisely:

```markdown
From .github/instructions/general-coding.instructions.md:

## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS

**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction
of any kind.**

## 🚨 CRITICAL: Use VS Code Tasks First

**MANDATORY RULE: Always attempt to use VS Code tasks before manual commands.**

## Version Update Requirements

**When modifying any file with a version header, ALWAYS update the version
number:**

- Patch version (x.y.Z): Bug fixes, typos, minor formatting changes
- Minor version (x.Y.z): New features, significant content additions
- Major version (X.y.z): Breaking changes, structural overhauls
```

### File Header Requirements

All modified files must have proper headers:

```go
// file: path/to/file.go
// version: 1.1.0  // INCREMENT THIS
// guid: existing-guid-unchanged
```

## 🧪 Testing Requirements

### Unit Tests

Create tests for config migration:

```go
// file: pkg/config/migration_test.go
func TestConfigMigration(t *testing.T) {
    // Test that old config format can be converted to new format
    // Test that all required fields are properly set using opaque API
    // Test that getters return expected values
}
```

### Integration Tests

```go
// file: pkg/config/integration_test.go
func TestConfigIntegration(t *testing.T) {
    // Test loading config from file using new types
    // Test that application boots with new config types
    // Test that all subsystems can access config via new API
}
```

## 🎯 Success Metrics

- [ ] All configpb imports removed
- [ ] `go build ./...` completes successfully
- [ ] All existing tests pass
- [ ] New tests added with 80%+ coverage
- [ ] No direct field access to protobuf fields (all via getters/setters)
- [ ] Performance benchmarks show no regression

## 🚨 Common Pitfalls

1. **Opaque API Confusion**: Remember to use `SetField()` and `GetField()`
   methods instead of direct field access
2. **Type Mismatches**: Ensure enum values match between old and new packages
3. **Nil Pointer Issues**: Always check for nil before calling methods on
   protobuf messages
4. **Import Cycles**: Be careful not to create import cycles when updating
   imports

## 📖 Additional Resources

- [gcommon config documentation](../../gcommon-api/config.md)
- [Protobuf Go Generated Code Guide](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
- [General Coding Instructions](../../../.github/instructions/general-coding.instructions.md)

## 🔄 Related Tasks

- **TASK-01-002**: Replace databasepb (may share some config types)
- **TASK-01-005**: Migrate protobuf message types (will need updated config
  types)
- **TASK-01-006**: Update import statements (will verify this task's imports)

## 📝 Notes for AI Agent

- This task is completely independent and can be executed without waiting for
  other tasks
- Use VS Code tasks for all build operations: `Buf Generate with Output`,
  `Git Add All`, etc.
- Follow the version update rules strictly - increment version numbers in all
  modified files
- The opaque API is critical - never access protobuf fields directly, always use
  getters/setters
- If any step fails, document the error and continue with remaining steps where
  possible
