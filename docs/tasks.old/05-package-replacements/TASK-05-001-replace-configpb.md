<!-- file: docs/tasks/05-package-replacements/TASK-05-001-replace-configpb.md -->
<!-- version: 1.0.0 -->
<!-- guid: a1b2c3d4-e5f6-7a8b-9c0d-1e2f3a4b5c6d -->

# TASK-05-001: Replace configpb with gcommon/config

## Overview

**Objective**: Completely replace local configpb package with gcommon config
types throughout the subtitle-manager application.

**Phase**: 3 (Package Replacements) **Priority**: High **Estimated Effort**: 6-8
hours **Prerequisites**: Phase 2 Core Type Migration completed (Tasks 6-10)

## Required Reading

**CRITICAL**: Read these documents before starting:

- `docs/gcommon-api/config.md` - gcommon configuration type specifications
- `docs/MIGRATION_INVENTORY.md` - Current package usage inventory
- `pkg/configpb/config.pb.go` - Current configpb implementation to be replaced
- `pkg/translatorpb/translator.pb.go` - Translation service configuration
- `pkg/subtitle/translator/v1/translator.pb.go` - Subtitle translator
  configuration

## Problem Statement

The subtitle-manager currently uses a local `configpb` package for configuration
management. This package needs to be completely replaced with gcommon
configuration types to:

1. **Eliminate Local Dependencies**: Remove custom protobuf config types
2. **Standardize Configuration**: Use gcommon standard configuration patterns
3. **Improve Interoperability**: Enable configuration sharing with other
   gcommon-based services
4. **Leverage Opaque API**: Use gcommon's flexible configuration system

### Current Configuration Structure

```go
// Current implementation (to be replaced)
import configpb "github.com/jdfalk/subtitle-manager/pkg/configpb"

type SubtitleManagerConfig struct {
    Server    *ServerConfig
    Database  *DatabaseConfig
    Auth      *AuthConfig
    Logging   *LoggingConfig
    Features  *FeatureConfig
}
```

### Target gcommon Structure

```go
// New implementation using gcommon
import "github.com/jdfalk/gcommon/sdks/go/v1/config"

// Use gcommon ApplicationConfig with opaque API
appConfig := &config.ApplicationConfig{}
appConfig.SetName("subtitle-manager")
appConfig.SetVersion("1.0.0")
```

## Technical Approach

### Configuration Mapping Strategy

1. **Identify Configuration Sections**: Map current config sections to gcommon
   equivalents
2. **Opaque API Migration**: Convert field access to Set*/Get* methods
3. **Validation Integration**: Use gcommon validation for configuration
4. **Environment Integration**: Leverage gcommon environment variable support

### Key Mappings

```go
// Configuration section mappings
type ConfigMapping struct {
    // OLD configpb -> NEW gcommon
    SubtitleManagerConfig -> config.ApplicationConfig
    ServerConfig         -> config.ServerConfig
    DatabaseConfig       -> config.DatabaseConfig
    AuthConfig          -> config.AuthConfig
    LoggingConfig       -> config.LoggingConfig
}
```

## Implementation Steps

### Step 1: Analyze Current Configuration Usage

```bash
# Find all configpb imports and usages
grep -r "configpb" pkg/ --include="*.go"
grep -r "SubtitleManagerConfig" pkg/ --include="*.go"

# Document current configuration fields and their usage
# Create mapping document: config_field_mapping.md
```

### Step 2: Create Configuration Wrapper

```go
// File: pkg/config/gcommon_config.go
package config

import (
    "fmt"
    "os"

    "github.com/jdfalk/gcommon/sdks/go/v1/config"
)

// SubtitleManagerConfig wraps gcommon ApplicationConfig
type SubtitleManagerConfig struct {
    appConfig *config.ApplicationConfig
}

// NewSubtitleManagerConfig creates a new configuration instance
func NewSubtitleManagerConfig() *SubtitleManagerConfig {
    appConfig := &config.ApplicationConfig{}
    appConfig.SetName("subtitle-manager")
    appConfig.SetVersion("1.0.0")

    return &SubtitleManagerConfig{
        appConfig: appConfig,
    }
}

// Server configuration methods
func (c *SubtitleManagerConfig) GetServerPort() int {
    if port, exists := c.appConfig.GetField("server.port"); exists {
        if portInt, ok := port.(int); ok {
            return portInt
        }
    }
    return 8080 // default
}

func (c *SubtitleManagerConfig) SetServerPort(port int) {
    c.appConfig.SetField("server.port", port)
}

func (c *SubtitleManagerConfig) GetServerHost() string {
    if host, exists := c.appConfig.GetField("server.host"); exists {
        if hostStr, ok := host.(string); ok {
            return hostStr
        }
    }
    return "localhost" // default
}

func (c *SubtitleManagerConfig) SetServerHost(host string) {
    c.appConfig.SetField("server.host", host)
}

// Database configuration methods
func (c *SubtitleManagerConfig) GetDatabaseType() string {
    if dbType, exists := c.appConfig.GetField("database.type"); exists {
        if typeStr, ok := dbType.(string); ok {
            return typeStr
        }
    }
    return "sqlite" // default
}

func (c *SubtitleManagerConfig) SetDatabaseType(dbType string) {
    c.appConfig.SetField("database.type", dbType)
}

func (c *SubtitleManagerConfig) GetDatabasePath() string {
    if path, exists := c.appConfig.GetField("database.path"); exists {
        if pathStr, ok := path.(string); ok {
            return pathStr
        }
    }
    return "./subtitle_manager.db" // default
}

func (c *SubtitleManagerConfig) SetDatabasePath(path string) {
    c.appConfig.SetField("database.path", path)
}

// Auth configuration methods
func (c *SubtitleManagerConfig) GetOAuth2ClientID() string {
    if clientID, exists := c.appConfig.GetField("auth.oauth2.client_id"); exists {
        if idStr, ok := clientID.(string); ok {
            return idStr
        }
    }
    return ""
}

func (c *SubtitleManagerConfig) SetOAuth2ClientID(clientID string) {
    c.appConfig.SetField("auth.oauth2.client_id", clientID)
}

func (c *SubtitleManagerConfig) GetJWTSecret() string {
    if secret, exists := c.appConfig.GetField("auth.jwt.secret"); exists {
        if secretStr, ok := secret.(string); ok {
            return secretStr
        }
    }
    return ""
}

func (c *SubtitleManagerConfig) SetJWTSecret(secret string) {
    c.appConfig.SetField("auth.jwt.secret", secret)
}

// Environment variable loading
func (c *SubtitleManagerConfig) LoadFromEnvironment() error {
    // Load configuration from environment variables
    if port := os.Getenv("SM_SERVER_PORT"); port != "" {
        if portInt, err := strconv.Atoi(port); err == nil {
            c.SetServerPort(portInt)
        }
    }

    if host := os.Getenv("SM_SERVER_HOST"); host != "" {
        c.SetServerHost(host)
    }

    if dbType := os.Getenv("SM_DATABASE_TYPE"); dbType != "" {
        c.SetDatabaseType(dbType)
    }

    if dbPath := os.Getenv("SM_DATABASE_PATH"); dbPath != "" {
        c.SetDatabasePath(dbPath)
    }

    if clientID := os.Getenv("SM_OAUTH2_CLIENT_ID"); clientID != "" {
        c.SetOAuth2ClientID(clientID)
    }

    if jwtSecret := os.Getenv("SM_JWT_SECRET"); jwtSecret != "" {
        c.SetJWTSecret(jwtSecret)
    }

    return nil
}

// Configuration validation
func (c *SubtitleManagerConfig) Validate() error {
    // Use gcommon validation if available
    if c.GetJWTSecret() == "" {
        return fmt.Errorf("JWT secret is required")
    }

    if c.GetServerPort() <= 0 || c.GetServerPort() > 65535 {
        return fmt.Errorf("server port must be between 1 and 65535")
    }

    return nil
}

// Get underlying gcommon config for advanced usage
func (c *SubtitleManagerConfig) GetApplicationConfig() *config.ApplicationConfig {
    return c.appConfig
}
```

### Step 3: Update Configuration Loading

```go
// File: pkg/config/loader.go
package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

// ConfigLoader handles loading configuration from various sources
type ConfigLoader struct {
    config *SubtitleManagerConfig
}

// NewConfigLoader creates a new configuration loader
func NewConfigLoader() *ConfigLoader {
    return &ConfigLoader{
        config: NewSubtitleManagerConfig(),
    }
}

// LoadFromFile loads configuration from a file (JSON or YAML)
func (l *ConfigLoader) LoadFromFile(configPath string) error {
    data, err := ioutil.ReadFile(configPath)
    if err != nil {
        return fmt.Errorf("failed to read config file: %v", err)
    }

    ext := filepath.Ext(configPath)
    switch ext {
    case ".json":
        return l.loadFromJSON(data)
    case ".yaml", ".yml":
        return l.loadFromYAML(data)
    default:
        return fmt.Errorf("unsupported config file format: %s", ext)
    }
}

func (l *ConfigLoader) loadFromJSON(data []byte) error {
    var configData map[string]interface{}
    if err := json.Unmarshal(data, &configData); err != nil {
        return fmt.Errorf("failed to parse JSON config: %v", err)
    }

    return l.applyConfigData(configData)
}

func (l *ConfigLoader) loadFromYAML(data []byte) error {
    var configData map[string]interface{}
    if err := yaml.Unmarshal(data, &configData); err != nil {
        return fmt.Errorf("failed to parse YAML config: %v", err)
    }

    return l.applyConfigData(configData)
}

func (l *ConfigLoader) applyConfigData(data map[string]interface{}) error {
    // Apply server configuration
    if server, exists := data["server"]; exists {
        if serverMap, ok := server.(map[string]interface{}); ok {
            if port, exists := serverMap["port"]; exists {
                if portFloat, ok := port.(float64); ok {
                    l.config.SetServerPort(int(portFloat))
                }
            }
            if host, exists := serverMap["host"]; exists {
                if hostStr, ok := host.(string); ok {
                    l.config.SetServerHost(hostStr)
                }
            }
        }
    }

    // Apply database configuration
    if database, exists := data["database"]; exists {
        if dbMap, ok := database.(map[string]interface{}); ok {
            if dbType, exists := dbMap["type"]; exists {
                if typeStr, ok := dbType.(string); ok {
                    l.config.SetDatabaseType(typeStr)
                }
            }
            if path, exists := dbMap["path"]; exists {
                if pathStr, ok := path.(string); ok {
                    l.config.SetDatabasePath(pathStr)
                }
            }
        }
    }

    // Apply auth configuration
    if auth, exists := data["auth"]; exists {
        if authMap, ok := auth.(map[string]interface{}); ok {
            if oauth2, exists := authMap["oauth2"]; exists {
                if oauth2Map, ok := oauth2.(map[string]interface{}); ok {
                    if clientID, exists := oauth2Map["client_id"]; exists {
                        if idStr, ok := clientID.(string); ok {
                            l.config.SetOAuth2ClientID(idStr)
                        }
                    }
                }
            }
            if jwt, exists := authMap["jwt"]; exists {
                if jwtMap, ok := jwt.(map[string]interface{}); ok {
                    if secret, exists := jwtMap["secret"]; exists {
                        if secretStr, ok := secret.(string); ok {
                            l.config.SetJWTSecret(secretStr)
                        }
                    }
                }
            }
        }
    }

    return nil
}

// GetConfig returns the loaded configuration
func (l *ConfigLoader) GetConfig() *SubtitleManagerConfig {
    return l.config
}

// LoadDefault loads default configuration with environment overrides
func LoadDefault() (*SubtitleManagerConfig, error) {
    loader := NewConfigLoader()

    // Load from environment variables
    if err := loader.config.LoadFromEnvironment(); err != nil {
        return nil, fmt.Errorf("failed to load environment config: %v", err)
    }

    // Try to load from config file if it exists
    configPaths := []string{
        "./config.yaml",
        "./config.yml",
        "./config.json",
        "/etc/subtitle-manager/config.yaml",
        "$HOME/.subtitle-manager/config.yaml",
    }

    for _, path := range configPaths {
        expandedPath := os.ExpandEnv(path)
        if _, err := os.Stat(expandedPath); err == nil {
            if err := loader.LoadFromFile(expandedPath); err != nil {
                return nil, fmt.Errorf("failed to load config from %s: %v", expandedPath, err)
            }
            break
        }
    }

    // Validate configuration
    if err := loader.config.Validate(); err != nil {
        return nil, fmt.Errorf("configuration validation failed: %v", err)
    }

    return loader.config, nil
}
```

### Step 4: Update Application Initialization

```go
// File: cmd/subtitle-manager/main.go
package main

import (
    "log"

    "github.com/jdfalk/subtitle-manager/pkg/config"
    "github.com/jdfalk/subtitle-manager/pkg/webserver"
)

func main() {
    // Load configuration using new gcommon-based system
    cfg, err := config.LoadDefault()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Start web server with new configuration
    server := webserver.NewServer(cfg)
    if err := server.Start(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
```

### Step 5: Update All Configuration References

```go
// Example: Update web server initialization
// File: pkg/webserver/server.go
package webserver

import (
    "fmt"
    "net/http"

    "github.com/jdfalk/subtitle-manager/pkg/config"
)

type Server struct {
    config *config.SubtitleManagerConfig
    router *http.ServeMux
}

func NewServer(cfg *config.SubtitleManagerConfig) *Server {
    return &Server{
        config: cfg,
        router: http.NewServeMux(),
    }
}

func (s *Server) Start() error {
    // Use configuration from gcommon-based config
    addr := fmt.Sprintf("%s:%d", s.config.GetServerHost(), s.config.GetServerPort())

    log.Printf("Starting server on %s", addr)
    return http.ListenAndServe(addr, s.router)
}
```

### Step 6: Remove Old configpb Package

```bash
# Remove old package files
rm -rf pkg/configpb/

# Update go.mod to remove old dependencies
go mod tidy

# Verify no references remain
grep -r "configpb" pkg/ --include="*.go" || echo "All configpb references removed"
```

## Testing Requirements

### Configuration Loading Tests

```go
// File: pkg/config/config_test.go
package config

import (
    "os"
    "testing"
    "io/ioutil"
    "path/filepath"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestSubtitleManagerConfig(t *testing.T) {
    config := NewSubtitleManagerConfig()

    // Test default values
    assert.Equal(t, "localhost", config.GetServerHost())
    assert.Equal(t, 8080, config.GetServerPort())
    assert.Equal(t, "sqlite", config.GetDatabaseType())
    assert.Equal(t, "./subtitle_manager.db", config.GetDatabasePath())

    // Test setting values
    config.SetServerHost("0.0.0.0")
    config.SetServerPort(9000)
    config.SetDatabaseType("postgresql")
    config.SetDatabasePath("/data/db.sqlite")

    assert.Equal(t, "0.0.0.0", config.GetServerHost())
    assert.Equal(t, 9000, config.GetServerPort())
    assert.Equal(t, "postgresql", config.GetDatabaseType())
    assert.Equal(t, "/data/db.sqlite", config.GetDatabasePath())
}

func TestConfigFromEnvironment(t *testing.T) {
    // Set environment variables
    os.Setenv("SM_SERVER_HOST", "test.example.com")
    os.Setenv("SM_SERVER_PORT", "3000")
    os.Setenv("SM_DATABASE_TYPE", "postgresql")
    defer func() {
        os.Unsetenv("SM_SERVER_HOST")
        os.Unsetenv("SM_SERVER_PORT")
        os.Unsetenv("SM_DATABASE_TYPE")
    }()

    config := NewSubtitleManagerConfig()
    err := config.LoadFromEnvironment()
    require.NoError(t, err)

    assert.Equal(t, "test.example.com", config.GetServerHost())
    assert.Equal(t, 3000, config.GetServerPort())
    assert.Equal(t, "postgresql", config.GetDatabaseType())
}

func TestConfigFromYAMLFile(t *testing.T) {
    yamlConfig := `
server:
  host: yaml.example.com
  port: 4000
database:
  type: postgresql
  path: /yaml/db.sqlite
auth:
  oauth2:
    client_id: yaml_client_id
  jwt:
    secret: yaml_secret
`

    // Create temporary config file
    tmpFile, err := ioutil.TempFile("", "config_*.yaml")
    require.NoError(t, err)
    defer os.Remove(tmpFile.Name())

    _, err = tmpFile.WriteString(yamlConfig)
    require.NoError(t, err)
    tmpFile.Close()

    // Load configuration
    loader := NewConfigLoader()
    err = loader.LoadFromFile(tmpFile.Name())
    require.NoError(t, err)

    config := loader.GetConfig()
    assert.Equal(t, "yaml.example.com", config.GetServerHost())
    assert.Equal(t, 4000, config.GetServerPort())
    assert.Equal(t, "postgresql", config.GetDatabaseType())
    assert.Equal(t, "/yaml/db.sqlite", config.GetDatabasePath())
    assert.Equal(t, "yaml_client_id", config.GetOAuth2ClientID())
    assert.Equal(t, "yaml_secret", config.GetJWTSecret())
}

func TestConfigValidation(t *testing.T) {
    config := NewSubtitleManagerConfig()

    // Invalid port
    config.SetServerPort(-1)
    err := config.Validate()
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "server port")

    // Missing JWT secret
    config.SetServerPort(8080)
    config.SetJWTSecret("")
    err = config.Validate()
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "JWT secret")

    // Valid configuration
    config.SetJWTSecret("valid_secret")
    err = config.Validate()
    assert.NoError(t, err)
}

func TestGcommonIntegration(t *testing.T) {
    config := NewSubtitleManagerConfig()

    // Test that we can access the underlying gcommon config
    appConfig := config.GetApplicationConfig()
    assert.NotNil(t, appConfig)

    // Test custom field setting through gcommon API
    appConfig.SetField("custom.feature.enabled", true)

    value, exists := appConfig.GetField("custom.feature.enabled")
    assert.True(t, exists)
    assert.Equal(t, true, value)
}
```

### Integration Tests

```go
// File: pkg/config/integration_test.go
func TestConfigIntegrationWithWebServer(t *testing.T) {
    config := NewSubtitleManagerConfig()
    config.SetServerHost("127.0.0.1")
    config.SetServerPort(0) // Let OS choose port
    config.SetJWTSecret("test_secret")

    // Test that web server can be initialized with new config
    server := webserver.NewServer(config)
    assert.NotNil(t, server)

    // Test configuration is properly passed through
    // (Implementation depends on server.go updates)
}

func TestConfigIntegrationWithDatabase(t *testing.T) {
    config := NewSubtitleManagerConfig()
    config.SetDatabaseType("sqlite")
    config.SetDatabasePath(":memory:")

    // Test that database can be initialized with new config
    // (Implementation depends on database package updates)
}
```

## Validation Scripts

### Configuration Migration Verification

```bash
#!/bin/bash
# File: scripts/verify_configpb_migration.sh

echo "Verifying configpb to gcommon migration..."

# Check that no configpb imports remain
echo "=== Checking for remaining configpb imports ==="
if grep -r "configpb" pkg/ --include="*.go"; then
    echo "ERROR: Found remaining configpb imports"
    exit 1
else
    echo "✓ No configpb imports found"
fi

# Check that gcommon config is being used
echo "=== Checking for gcommon config usage ==="
if grep -r "github.com/jdfalk/gcommon/sdks/go/v1/config" pkg/ --include="*.go"; then
    echo "✓ Found gcommon config imports"
else
    echo "ERROR: No gcommon config imports found"
    exit 1
fi

# Verify configuration loading works
echo "=== Testing configuration loading ==="
go run cmd/subtitle-manager/main.go --config-test 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✓ Configuration loading works"
else
    echo "ERROR: Configuration loading failed"
    exit 1
fi

echo "✓ configpb migration verification complete"
```

### Configuration Testing Script

```go
// File: scripts/test_config_migration.go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/jdfalk/subtitle-manager/pkg/config"
)

func main() {
    fmt.Println("Testing gcommon configuration migration...")

    // Test 1: Basic configuration creation
    cfg := config.NewSubtitleManagerConfig()
    if cfg == nil {
        log.Fatal("Failed to create configuration")
    }
    fmt.Println("✓ Configuration creation successful")

    // Test 2: Environment variable loading
    os.Setenv("SM_SERVER_PORT", "9999")
    if err := cfg.LoadFromEnvironment(); err != nil {
        log.Fatalf("Failed to load environment config: %v", err)
    }
    if cfg.GetServerPort() != 9999 {
        log.Fatal("Environment variable not loaded correctly")
    }
    fmt.Println("✓ Environment variable loading successful")

    // Test 3: Validation
    cfg.SetJWTSecret("test_secret")
    if err := cfg.Validate(); err != nil {
        log.Fatalf("Configuration validation failed: %v", err)
    }
    fmt.Println("✓ Configuration validation successful")

    // Test 4: gcommon integration
    appConfig := cfg.GetApplicationConfig()
    if appConfig == nil {
        log.Fatal("Failed to get gcommon ApplicationConfig")
    }
    fmt.Println("✓ gcommon integration successful")

    fmt.Println("All configuration migration tests passed!")
}
```

## Success Metrics

### Functional Requirements

- [ ] All configpb imports completely removed from codebase
- [ ] Configuration loading uses gcommon ApplicationConfig
- [ ] All configuration access uses opaque API (Set*/Get* methods)
- [ ] Environment variable loading works correctly
- [ ] Configuration file loading (YAML/JSON) works correctly
- [ ] Configuration validation properly implemented

### Technical Requirements

- [ ] No breaking changes to existing configuration behavior
- [ ] Performance impact negligible (< 5ms configuration load time)
- [ ] Memory usage does not increase significantly
- [ ] All configuration tests pass
- [ ] Integration tests with other components pass
- [ ] Configuration backwards compatibility maintained

### Migration Requirements

- [ ] pkg/configpb directory completely removed
- [ ] All files importing configpb updated to use new config system
- [ ] Configuration file formats remain compatible
- [ ] Environment variable names unchanged (if applicable)
- [ ] Default values preserved during migration

## Common Pitfalls

1. **Field Type Mismatches**: gcommon opaque API returns interface{}, ensure
   proper type assertions
2. **Default Value Handling**: Ensure default values are consistent with old
   configpb behavior
3. **Environment Variable Precedence**: Maintain same precedence order as before
4. **Configuration Validation**: Don't lose existing validation logic during
   migration
5. **Nested Configuration**: Handle complex nested config structures properly
   with opaque API

## Dependencies

- **Requires**: Phase 2 Core Type Migration completed
- **Requires**: gcommon SDK properly installed and available
- **Enables**: Standard configuration patterns across gcommon-based services
- **Blocks**: Any configuration-dependent features until migration complete

## Embedded Documentation

### gcommon Configuration Patterns

```go
// Standard gcommon configuration usage patterns
import "github.com/jdfalk/gcommon/sdks/go/v1/config"

// Create application configuration
appConfig := &config.ApplicationConfig{}
appConfig.SetName("subtitle-manager")
appConfig.SetVersion("1.0.0")

// Use opaque API for custom fields
appConfig.SetField("server.port", 8080)
appConfig.SetField("server.host", "localhost")
appConfig.SetField("database.type", "sqlite")

// Retrieve values with type checking
if port, exists := appConfig.GetField("server.port"); exists {
    if portInt, ok := port.(int); ok {
        // Use portInt
    }
}
```

### Configuration Best Practices

1. **Type Safety**: Always check types when using opaque API
2. **Default Values**: Provide sensible defaults for all configuration values
3. **Validation**: Validate configuration at application startup
4. **Environment Variables**: Support environment variable overrides
5. **Documentation**: Document all configuration options clearly

This comprehensive task ensures complete migration from configpb to gcommon
configuration types while maintaining functionality and improving
standardization across the gcommon ecosystem.
