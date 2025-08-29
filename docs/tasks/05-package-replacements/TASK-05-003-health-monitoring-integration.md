# file: docs/tasks/05-package-replacements/TASK-05-003-health-monitoring-integration.md
# version: 1.0.0
# guid: c3d4e5f6-a7b8-9c0d-1e2f-3a4b5c6d7e8f

# TASK-05-003: Health Monitoring Integration

## Overview

**Objective**: Integrate gcommon health monitoring throughout the subtitle-manager application.

**Phase**: 3 (Package Replacements)
**Priority**: Medium
**Estimated Effort**: 4-6 hours
**Prerequisites**: TASK-05-002 (databasepb replacement) and gcommon database integration

## Required Reading

**CRITICAL**: Read these documents before starting:

- `docs/gcommon-api/health.md` - gcommon health monitoring specifications and patterns
- Current health checking implementations in the application
- `pkg/webserver/` directory for existing health endpoints
- `docs/MIGRATION_INVENTORY.md` - Health monitoring usage inventory
- Application monitoring and uptime requirements

## Problem Statement

The subtitle-manager currently lacks comprehensive health monitoring or uses basic health check implementations. This needs to be upgraded to use gcommon health monitoring to:

1. **Standardize Health Checks**: Use gcommon HealthStatus types for consistent monitoring
2. **Comprehensive Monitoring**: Monitor database, external services, and application health
3. **Integration Ready**: Prepare for integration with monitoring systems (Prometheus, etc.)
4. **Debugging Support**: Provide detailed health information for troubleshooting

### Current State

```go
// Current implementation (basic or missing)
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}
```

### Target gcommon Health Monitoring

```go
// New implementation using gcommon
import "github.com/jdfalk/gcommon/sdks/go/v1/health"

healthService := health.NewService()
healthService.AddCheck("database", dbHealthCheck)
healthService.AddCheck("external_api", apiHealthCheck)

status := healthService.GetOverallHealth()
// Returns comprehensive health status with individual component states
```

## Technical Approach

### Health Monitoring Strategy

1. **Service Registration**: Register all critical services for health monitoring
2. **Check Implementation**: Implement comprehensive health checks for each service
3. **Status Aggregation**: Use gcommon health status aggregation
4. **Endpoint Integration**: Expose health status via HTTP endpoints

### Key Components

```go
// Health monitoring architecture
type HealthComponents struct {
    Database         health.Check // Database connectivity and performance
    FileSystem       health.Check // Storage availability and space
    ExternalServices health.Check // API connectivity and response times
    Memory           health.Check // Memory usage and availability
    Configuration    health.Check // Configuration validity
}
```

## Implementation Steps

### Step 1: Create Health Monitoring Package

```go
// File: pkg/health/service.go
package health

import (
    "context"
    "fmt"
    "time"

    "github.com/jdfalk/gcommon/sdks/go/v1/health"
    "github.com/jdfalk/subtitle-manager/pkg/database"
    "github.com/jdfalk/subtitle-manager/pkg/config"
)

// HealthService manages application health monitoring
type HealthService struct {
    healthService  health.Service
    dbManager     *database.DatabaseManager
    config        *config.ApplicationConfig
}

// NewHealthService creates a new health monitoring service
func NewHealthService(dbManager *database.DatabaseManager, config *config.ApplicationConfig) *HealthService {
    service := &HealthService{
        healthService: health.NewService(),
        dbManager:     dbManager,
        config:        config,
    }

    // Register all health checks
    service.registerHealthChecks()

    return service
}

// registerHealthChecks sets up all application health checks
func (hs *HealthService) registerHealthChecks() {
    // Database health check
    hs.healthService.AddCheck("database", hs.databaseHealthCheck)

    // File system health check
    hs.healthService.AddCheck("filesystem", hs.fileSystemHealthCheck)

    // Configuration health check
    hs.healthService.AddCheck("configuration", hs.configurationHealthCheck)

    // Memory health check
    hs.healthService.AddCheck("memory", hs.memoryHealthCheck)

    // External services health check (if any)
    hs.healthService.AddCheck("external_services", hs.externalServicesHealthCheck)
}

// Database health check implementation
func (hs *HealthService) databaseHealthCheck(ctx context.Context) *health.CheckResult {
    result := health.NewCheckResult()

    // Test database connectivity
    testRecord := database.NewSubtitleRecord()
    testRecord.SetID("health_check_test")

    start := time.Now()

    // Perform a simple database operation
    err := hs.dbManager.CreateSubtitleRecord(testRecord)
    if err != nil {
        result.SetStatus(health.StatusUnhealthy)
        result.SetMessage(fmt.Sprintf("Database create failed: %v", err))
        return result
    }

    // Clean up test record
    hs.dbManager.DeleteSubtitleRecord("health_check_test")

    latency := time.Since(start)

    // Check if database response time is acceptable
    if latency > 5*time.Second {
        result.SetStatus(health.StatusDegraded)
        result.SetMessage(fmt.Sprintf("Database slow response: %v", latency))
    } else {
        result.SetStatus(health.StatusHealthy)
        result.SetMessage("Database operational")
    }

    // Add performance metrics
    result.SetMetric("response_time_ms", float64(latency.Milliseconds()))

    return result
}

// File system health check implementation
func (hs *HealthService) fileSystemHealthCheck(ctx context.Context) *health.CheckResult {
    result := health.NewCheckResult()

    // Check data directory accessibility
    dataDir := hs.config.GetDataDirectory()
    if dataDir == "" {
        result.SetStatus(health.StatusUnhealthy)
        result.SetMessage("Data directory not configured")
        return result
    }

    // Check directory exists and is writable
    testFile := filepath.Join(dataDir, "health_check_test.tmp")
    err := ioutil.WriteFile(testFile, []byte("test"), 0644)
    if err != nil {
        result.SetStatus(health.StatusUnhealthy)
        result.SetMessage(fmt.Sprintf("Data directory not writable: %v", err))
        return result
    }

    // Clean up test file
    os.Remove(testFile)

    // Check available disk space
    stat, err := os.Stat(dataDir)
    if err != nil {
        result.SetStatus(health.StatusDegraded)
        result.SetMessage(fmt.Sprintf("Cannot check disk space: %v", err))
        return result
    }

    // Check disk space (implementation depends on OS)
    // This is a simplified example
    result.SetStatus(health.StatusHealthy)
    result.SetMessage("File system operational")

    return result
}

// Configuration health check implementation
func (hs *HealthService) configurationHealthCheck(ctx context.Context) *health.CheckResult {
    result := health.NewCheckResult()

    // Validate configuration completeness
    if hs.config.GetServerPort() == 0 {
        result.SetStatus(health.StatusUnhealthy)
        result.SetMessage("Server port not configured")
        return result
    }

    if hs.config.GetDataDirectory() == "" {
        result.SetStatus(health.StatusUnhealthy)
        result.SetMessage("Data directory not configured")
        return result
    }

    // Check for required configuration values
    requiredSettings := []string{
        "database_path",
        "log_level",
        "bind_address",
    }

    for _, setting := range requiredSettings {
        if value := hs.config.GetStringSetting(setting); value == "" {
            result.SetStatus(health.StatusDegraded)
            result.SetMessage(fmt.Sprintf("Configuration missing: %s", setting))
            return result
        }
    }

    result.SetStatus(health.StatusHealthy)
    result.SetMessage("Configuration valid")

    return result
}

// Memory health check implementation
func (hs *HealthService) memoryHealthCheck(ctx context.Context) *health.CheckResult {
    result := health.NewCheckResult()

    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)

    // Convert bytes to MB
    allocMB := float64(memStats.Alloc) / 1024 / 1024
    sysMB := float64(memStats.Sys) / 1024 / 1024

    // Set memory usage metrics
    result.SetMetric("allocated_mb", allocMB)
    result.SetMetric("system_mb", sysMB)
    result.SetMetric("gc_cycles", float64(memStats.NumGC))

    // Check memory thresholds
    if allocMB > 512 { // 512MB threshold
        result.SetStatus(health.StatusDegraded)
        result.SetMessage(fmt.Sprintf("High memory usage: %.2f MB", allocMB))
    } else {
        result.SetStatus(health.StatusHealthy)
        result.SetMessage(fmt.Sprintf("Memory usage normal: %.2f MB", allocMB))
    }

    return result
}

// External services health check implementation
func (hs *HealthService) externalServicesHealthCheck(ctx context.Context) *health.CheckResult {
    result := health.NewCheckResult()

    // Check external services if any are configured
    // For subtitle-manager, this might include:
    // - Subtitle databases/APIs
    // - Download services
    // - Authentication services

    // Example: Check internet connectivity for download features
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    resp, err := client.Get("https://httpbin.org/status/200")
    if err != nil {
        result.SetStatus(health.StatusDegraded)
        result.SetMessage(fmt.Sprintf("External connectivity issue: %v", err))
        return result
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        result.SetStatus(health.StatusDegraded)
        result.SetMessage(fmt.Sprintf("External service responded with: %d", resp.StatusCode))
        return result
    }

    result.SetStatus(health.StatusHealthy)
    result.SetMessage("External services accessible")

    return result
}

// GetOverallHealth returns the overall application health status
func (hs *HealthService) GetOverallHealth() *health.OverallStatus {
    return hs.healthService.GetOverallHealth()
}

// GetDetailedHealth returns detailed health information for all components
func (hs *HealthService) GetDetailedHealth() *health.DetailedStatus {
    return hs.healthService.GetDetailedStatus()
}

// GetComponentHealth returns health status for a specific component
func (hs *HealthService) GetComponentHealth(component string) *health.CheckResult {
    return hs.healthService.GetComponentHealth(component)
}
```

### Step 2: Create Health HTTP Endpoints

```go
// File: pkg/webserver/health.go
package webserver

import (
    "encoding/json"
    "net/http"

    "github.com/jdfalk/subtitle-manager/pkg/health"
    "github.com/gorilla/mux"
)

// HealthHandler manages health-related HTTP endpoints
type HealthHandler struct {
    healthService *health.HealthService
}

// NewHealthHandler creates a new health HTTP handler
func NewHealthHandler(healthService *health.HealthService) *HealthHandler {
    return &HealthHandler{
        healthService: healthService,
    }
}

// RegisterRoutes registers health-related routes
func (hh *HealthHandler) RegisterRoutes(router *mux.Router) {
    // Basic health check endpoint
    router.HandleFunc("/health", hh.basicHealthCheck).Methods("GET")

    // Detailed health status endpoint
    router.HandleFunc("/health/detailed", hh.detailedHealthCheck).Methods("GET")

    // Component-specific health endpoint
    router.HandleFunc("/health/{component}", hh.componentHealthCheck).Methods("GET")

    // Readiness probe endpoint (for Kubernetes)
    router.HandleFunc("/ready", hh.readinessCheck).Methods("GET")

    // Liveness probe endpoint (for Kubernetes)
    router.HandleFunc("/live", hh.livenessCheck).Methods("GET")
}

// basicHealthCheck provides a simple health status
func (hh *HealthHandler) basicHealthCheck(w http.ResponseWriter, r *http.Request) {
    overallHealth := hh.healthService.GetOverallHealth()

    // Set HTTP status based on health
    switch overallHealth.GetStatus() {
    case health.StatusHealthy:
        w.WriteHeader(http.StatusOK)
    case health.StatusDegraded:
        w.WriteHeader(http.StatusOK) // Still serving traffic
    case health.StatusUnhealthy:
        w.WriteHeader(http.StatusServiceUnavailable)
    default:
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")

    response := map[string]interface{}{
        "status":     overallHealth.GetStatus().String(),
        "message":    overallHealth.GetMessage(),
        "timestamp":  overallHealth.GetTimestamp(),
    }

    json.NewEncoder(w).Encode(response)
}

// detailedHealthCheck provides comprehensive health information
func (hh *HealthHandler) detailedHealthCheck(w http.ResponseWriter, r *http.Request) {
    detailedHealth := hh.healthService.GetDetailedHealth()

    // Set HTTP status based on overall health
    switch detailedHealth.GetOverallStatus() {
    case health.StatusHealthy:
        w.WriteHeader(http.StatusOK)
    case health.StatusDegraded:
        w.WriteHeader(http.StatusOK)
    case health.StatusUnhealthy:
        w.WriteHeader(http.StatusServiceUnavailable)
    default:
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")

    // Convert to JSON-friendly format
    response := map[string]interface{}{
        "overall_status": detailedHealth.GetOverallStatus().String(),
        "overall_message": detailedHealth.GetOverallMessage(),
        "timestamp": detailedHealth.GetTimestamp(),
        "components": make(map[string]interface{}),
    }

    // Add component details
    for componentName, checkResult := range detailedHealth.GetComponents() {
        response["components"].(map[string]interface{})[componentName] = map[string]interface{}{
            "status":     checkResult.GetStatus().String(),
            "message":    checkResult.GetMessage(),
            "metrics":    checkResult.GetMetrics(),
            "timestamp":  checkResult.GetTimestamp(),
        }
    }

    json.NewEncoder(w).Encode(response)
}

// componentHealthCheck provides health status for a specific component
func (hh *HealthHandler) componentHealthCheck(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    component := vars["component"]

    componentHealth := hh.healthService.GetComponentHealth(component)
    if componentHealth == nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{
            "error": "Component not found",
        })
        return
    }

    // Set HTTP status based on component health
    switch componentHealth.GetStatus() {
    case health.StatusHealthy:
        w.WriteHeader(http.StatusOK)
    case health.StatusDegraded:
        w.WriteHeader(http.StatusOK)
    case health.StatusUnhealthy:
        w.WriteHeader(http.StatusServiceUnavailable)
    default:
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")

    response := map[string]interface{}{
        "component":  component,
        "status":     componentHealth.GetStatus().String(),
        "message":    componentHealth.GetMessage(),
        "metrics":    componentHealth.GetMetrics(),
        "timestamp":  componentHealth.GetTimestamp(),
    }

    json.NewEncoder(w).Encode(response)
}

// readinessCheck determines if the application is ready to serve traffic
func (hh *HealthHandler) readinessCheck(w http.ResponseWriter, r *http.Request) {
    // Check critical components for readiness
    dbHealth := hh.healthService.GetComponentHealth("database")
    configHealth := hh.healthService.GetComponentHealth("configuration")

    if dbHealth.GetStatus() == health.StatusUnhealthy ||
       configHealth.GetStatus() == health.StatusUnhealthy {
        w.WriteHeader(http.StatusServiceUnavailable)
        json.NewEncoder(w).Encode(map[string]string{
            "status": "not_ready",
            "reason": "Critical components unhealthy",
        })
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status": "ready",
    })
}

// livenessCheck determines if the application is alive
func (hh *HealthHandler) livenessCheck(w http.ResponseWriter, r *http.Request) {
    // Basic liveness check - if we can respond, we're alive
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status": "alive",
    })
}
```

### Step 3: Create Health Management Command

```go
// File: cmd/health.go
package cmd

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/jdfalk/subtitle-manager/pkg/health"
    "github.com/jdfalk/subtitle-manager/pkg/database"
    "github.com/jdfalk/subtitle-manager/pkg/config"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
    Use:   "health",
    Short: "Check application health status",
    Long: `Check the health status of the subtitle-manager application.

This command performs comprehensive health checks on all application components
including database, file system, configuration, and external services.`,
    Run: runHealthCheck,
}

// healthCheckOptions holds command line options for health checks
type healthCheckOptions struct {
    component string
    format    string
    verbose   bool
}

var healthOpts healthCheckOptions

func init() {
    rootCmd.AddCommand(healthCmd)

    healthCmd.Flags().StringVarP(&healthOpts.component, "component", "c", "", "Check specific component only")
    healthCmd.Flags().StringVarP(&healthOpts.format, "format", "f", "text", "Output format (text, json)")
    healthCmd.Flags().BoolVarP(&healthOpts.verbose, "verbose", "v", false, "Verbose output with metrics")
}

// runHealthCheck executes the health check
func runHealthCheck(cmd *cobra.Command, args []string) {
    // Initialize dependencies
    cfg, err := config.LoadConfig()
    if err != nil {
        fmt.Printf("Failed to load configuration: %v\n", err)
        os.Exit(1)
    }

    dbManager, err := database.NewDatabaseManager(cfg.GetDatabasePath())
    if err != nil {
        fmt.Printf("Failed to initialize database: %v\n", err)
        os.Exit(1)
    }

    healthService := health.NewHealthService(dbManager, cfg)

    // Perform health checks
    if healthOpts.component != "" {
        checkSingleComponent(healthService, healthOpts.component)
    } else {
        checkAllComponents(healthService)
    }
}

// checkSingleComponent checks a specific component
func checkSingleComponent(healthService *health.HealthService, component string) {
    componentHealth := healthService.GetComponentHealth(component)
    if componentHealth == nil {
        fmt.Printf("Component '%s' not found\n", component)
        os.Exit(1)
    }

    if healthOpts.format == "json" {
        outputJSON(map[string]interface{}{
            "component": component,
            "status":    componentHealth.GetStatus().String(),
            "message":   componentHealth.GetMessage(),
            "metrics":   componentHealth.GetMetrics(),
            "timestamp": componentHealth.GetTimestamp(),
        })
    } else {
        outputTextComponent(component, componentHealth)
    }

    // Exit with non-zero code if unhealthy
    if componentHealth.GetStatus() == health.StatusUnhealthy {
        os.Exit(1)
    }
}

// checkAllComponents checks all application components
func checkAllComponents(healthService *health.HealthService) {
    if healthOpts.format == "json" {
        detailedHealth := healthService.GetDetailedHealth()
        outputJSON(detailedHealth)
    } else {
        outputTextOverall(healthService)
    }

    // Exit with non-zero code if overall health is unhealthy
    overallHealth := healthService.GetOverallHealth()
    if overallHealth.GetStatus() == health.StatusUnhealthy {
        os.Exit(1)
    }
}

// outputJSON outputs health status in JSON format
func outputJSON(data interface{}) {
    encoder := json.NewEncoder(os.Stdout)
    encoder.SetIndent("", "  ")
    encoder.Encode(data)
}

// outputTextComponent outputs single component health in text format
func outputTextComponent(component string, checkResult *health.CheckResult) {
    fmt.Printf("Component: %s\n", component)
    fmt.Printf("Status: %s\n", checkResult.GetStatus().String())
    fmt.Printf("Message: %s\n", checkResult.GetMessage())

    if healthOpts.verbose && len(checkResult.GetMetrics()) > 0 {
        fmt.Printf("Metrics:\n")
        for name, value := range checkResult.GetMetrics() {
            fmt.Printf("  %s: %v\n", name, value)
        }
    }

    fmt.Printf("Timestamp: %s\n", checkResult.GetTimestamp())
}

// outputTextOverall outputs overall health status in text format
func outputTextOverall(healthService *health.HealthService) {
    overallHealth := healthService.GetOverallHealth()
    detailedHealth := healthService.GetDetailedHealth()

    fmt.Printf("Overall Status: %s\n", overallHealth.GetStatus().String())
    fmt.Printf("Overall Message: %s\n", overallHealth.GetMessage())
    fmt.Printf("Timestamp: %s\n\n", overallHealth.GetTimestamp())

    fmt.Printf("Component Health:\n")
    fmt.Printf("=================\n")

    for componentName, checkResult := range detailedHealth.GetComponents() {
        fmt.Printf("\n%s:\n", componentName)
        fmt.Printf("  Status: %s\n", checkResult.GetStatus().String())
        fmt.Printf("  Message: %s\n", checkResult.GetMessage())

        if healthOpts.verbose && len(checkResult.GetMetrics()) > 0 {
            fmt.Printf("  Metrics:\n")
            for name, value := range checkResult.GetMetrics() {
                fmt.Printf("    %s: %v\n", name, value)
            }
        }
    }
}
```

### Step 4: Integration with Application Startup

```go
// File: cmd/serve.go (modify existing server command)
func runServer(cmd *cobra.Command, args []string) {
    // ... existing initialization code ...

    // Initialize health service
    healthService := health.NewHealthService(dbManager, cfg)

    // Perform startup health checks
    overallHealth := healthService.GetOverallHealth()
    if overallHealth.GetStatus() == health.StatusUnhealthy {
        log.Fatalf("Application startup failed - unhealthy components: %s",
                   overallHealth.GetMessage())
    }

    // Register health endpoints
    healthHandler := webserver.NewHealthHandler(healthService)
    healthHandler.RegisterRoutes(router)

    // ... rest of server startup ...
}
```

## Testing Requirements

### Health Service Tests

```go
// File: pkg/health/service_test.go
package health

import (
    "context"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "github.com/jdfalk/gcommon/sdks/go/v1/health"
)

func TestHealthService(t *testing.T) {
    // Create mock dependencies
    mockDB := &MockDatabaseManager{}
    mockConfig := &MockConfig{}

    healthService := NewHealthService(mockDB, mockConfig)

    t.Run("Overall health check", func(t *testing.T) {
        overallHealth := healthService.GetOverallHealth()
        assert.NotNil(t, overallHealth)
        assert.NotEmpty(t, overallHealth.GetStatus())
    })

    t.Run("Component health checks", func(t *testing.T) {
        components := []string{"database", "filesystem", "configuration", "memory"}

        for _, component := range components {
            componentHealth := healthService.GetComponentHealth(component)
            assert.NotNil(t, componentHealth, "Component %s should have health status", component)
        }
    })

    t.Run("Database health check", func(t *testing.T) {
        // Test healthy database
        mockDB.SetHealthy(true)
        dbHealth := healthService.GetComponentHealth("database")
        assert.Equal(t, health.StatusHealthy, dbHealth.GetStatus())

        // Test unhealthy database
        mockDB.SetHealthy(false)
        dbHealth = healthService.GetComponentHealth("database")
        assert.Equal(t, health.StatusUnhealthy, dbHealth.GetStatus())
    })
}

// Mock implementations for testing
type MockDatabaseManager struct {
    healthy bool
}

func (m *MockDatabaseManager) SetHealthy(healthy bool) {
    m.healthy = healthy
}

func (m *MockDatabaseManager) CreateSubtitleRecord(record *SubtitleRecord) error {
    if !m.healthy {
        return errors.New("database unavailable")
    }
    return nil
}

func (m *MockDatabaseManager) DeleteSubtitleRecord(id string) error {
    if !m.healthy {
        return errors.New("database unavailable")
    }
    return nil
}

type MockConfig struct{}

func (m *MockConfig) GetDataDirectory() string {
    return "/tmp/test"
}

func (m *MockConfig) GetServerPort() int {
    return 8080
}

func (m *MockConfig) GetStringSetting(key string) string {
    settings := map[string]string{
        "database_path": "/tmp/test.db",
        "log_level":     "info",
        "bind_address":  "0.0.0.0",
    }
    return settings[key]
}
```

### HTTP Endpoint Tests

```go
// File: pkg/webserver/health_test.go
func TestHealthEndpoints(t *testing.T) {
    mockHealthService := &MockHealthService{}
    handler := NewHealthHandler(mockHealthService)

    router := mux.NewRouter()
    handler.RegisterRoutes(router)

    t.Run("Basic health endpoint", func(t *testing.T) {
        req := httptest.NewRequest("GET", "/health", nil)
        w := httptest.NewRecorder()

        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)
        assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
    })

    t.Run("Detailed health endpoint", func(t *testing.T) {
        req := httptest.NewRequest("GET", "/health/detailed", nil)
        w := httptest.NewRecorder()

        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)

        var response map[string]interface{}
        err := json.Unmarshal(w.Body.Bytes(), &response)
        assert.NoError(t, err)
        assert.Contains(t, response, "overall_status")
        assert.Contains(t, response, "components")
    })
}
```

## Success Metrics

### Functional Requirements

- [ ] Health service monitors all critical application components
- [ ] HTTP endpoints provide comprehensive health information
- [ ] Command-line health checking tool works correctly
- [ ] Health checks complete within acceptable timeframes
- [ ] Unhealthy components properly reported and handled

### Technical Requirements

- [ ] gcommon health monitoring fully integrated
- [ ] Health status uses standardized gcommon types
- [ ] Performance metrics included in health reports
- [ ] Health checks are non-intrusive to application performance
- [ ] Proper error handling and graceful degradation

### Monitoring Requirements

- [ ] Ready for integration with monitoring systems (Prometheus, etc.)
- [ ] Kubernetes readiness and liveness probes supported
- [ ] Health metrics suitable for alerting
- [ ] Component-specific monitoring enables targeted troubleshooting

## Common Pitfalls

1. **Performance Impact**: Health checks should be lightweight and non-blocking
2. **False Positives**: Ensure health checks accurately reflect component state
3. **Timeout Handling**: Implement proper timeouts for external service checks
4. **Resource Leaks**: Clean up test resources created during health checks
5. **Security**: Avoid exposing sensitive information in health endpoints

## Dependencies

- **Requires**: TASK-05-002 (databasepb replacement) for database health checks
- **Requires**: gcommon health SDK properly installed
- **Requires**: Database and configuration services initialized
- **Enables**: Comprehensive application monitoring and troubleshooting
- **Enables**: Integration with external monitoring systems

This comprehensive health monitoring integration provides robust application monitoring using gcommon standardized health patterns while enabling effective troubleshooting and system reliability monitoring.
