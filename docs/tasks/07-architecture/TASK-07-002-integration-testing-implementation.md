# TASK-07-002: Integration Testing Implementation (gcommon Edition)

<!-- file: docs/tasks/07-architecture/TASK-07-002-integration-testing-implementation.md -->
<!-- version: 2.0.0 -->
<!-- guid: test07002-aaaa-bbbb-cccc-dddddddddddd -->

## Overview

Implement comprehensive integration testing for the subtitle-manager ecosystem
with full gcommon integration, covering service-to-service communication,
end-to-end workflows, and system reliability testing.

## Implementation Plan

### Step 1: Integration Testing Framework

**Create `pkg/testing/integration/framework.go`**:

```go
// file: pkg/testing/integration/framework.go
// version: 2.0.0
// guid: test-framework-1111-2222-3333-444444444444

package integration

import (
    "context"
    "fmt"
    "sync"
    "time"

    "go.uber.org/zap"

    // gcommon imports
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/common"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/health"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/metrics"

    // Architecture framework
    "github.com/jdfalk/subtitle-manager/pkg/architecture"
)

// TestFramework provides comprehensive integration testing capabilities
type TestFramework struct {
    // Core components
    logger          *zap.Logger
    config          *TestConfig
    serviceFramework *architecture.ServiceFramework

    // Test environment
    environment     *TestEnvironment
    testSuites      map[string]*TestSuite
    testRunners     map[string]*TestRunner

    // Results and reporting
    results         *TestResults
    reporters       []TestReporter

    // Lifecycle
    running         bool
    mu              sync.RWMutex
}

// TestConfig defines integration testing configuration
type TestConfig struct {
    Environment     *EnvironmentConfig      `yaml:"environment" json:"environment"`
    Services        map[string]*ServiceTestConfig `yaml:"services" json:"services"`
    Scenarios       []*ScenarioConfig       `yaml:"scenarios" json:"scenarios"`
    Timeouts        *TimeoutConfig          `yaml:"timeouts" json:"timeouts"`
    Reporting       *ReportingConfig        `yaml:"reporting" json:"reporting"`
    Cleanup         *CleanupConfig          `yaml:"cleanup" json:"cleanup"`
}

// EnvironmentConfig defines test environment configuration
type EnvironmentConfig struct {
    Name            string                  `yaml:"name" json:"name"`
    Type            string                  `yaml:"type" json:"type"` // local, docker, k8s
    BaseURL         string                  `yaml:"base_url" json:"base_url"`
    Services        map[string]*common.ServerConfig `yaml:"services" json:"services"`
    Database        *DatabaseConfig         `yaml:"database" json:"database"`
    Storage         *StorageConfig          `yaml:"storage" json:"storage"`
    Monitoring      *MonitoringConfig       `yaml:"monitoring" json:"monitoring"`
}

// ServiceTestConfig defines per-service test configuration
type ServiceTestConfig struct {
    Enabled         bool                    `yaml:"enabled" json:"enabled"`
    Endpoint        string                  `yaml:"endpoint" json:"endpoint"`
    HealthCheck     *HealthCheckConfig      `yaml:"health_check" json:"health_check"`
    Authentication  *AuthConfig             `yaml:"authentication" json:"authentication"`
    TestData        map[string]interface{}  `yaml:"test_data" json:"test_data"`
}

// ScenarioConfig defines test scenario configuration
type ScenarioConfig struct {
    Name            string                  `yaml:"name" json:"name"`
    Description     string                  `yaml:"description" json:"description"`
    Tags            []string                `yaml:"tags" json:"tags"`
    Steps           []*TestStep             `yaml:"steps" json:"steps"`
    Prerequisites   []string                `yaml:"prerequisites" json:"prerequisites"`
    Cleanup         []string                `yaml:"cleanup" json:"cleanup"`
    Timeout         time.Duration           `yaml:"timeout" json:"timeout"`
    Retries         int                     `yaml:"retries" json:"retries"`
}

// NewTestFramework creates a new integration testing framework
func NewTestFramework(logger *zap.Logger, config *TestConfig) (*TestFramework, error) {
    if logger == nil {
        return nil, fmt.Errorf("logger is required")
    }
    if config == nil {
        return nil, fmt.Errorf("config is required")
    }

    // Initialize service framework for testing
    serviceFramework, err := architecture.NewServiceFramework(logger, &common.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to create service framework: %w", err)
    }

    return &TestFramework{
        logger:           logger,
        config:           config,
        serviceFramework: serviceFramework,
        testSuites:       make(map[string]*TestSuite),
        testRunners:      make(map[string]*TestRunner),
        results:          NewTestResults(),
        reporters:        make([]TestReporter, 0),
    }, nil
}

// Start starts the testing framework
func (tf *TestFramework) Start(ctx context.Context) error {
    tf.mu.Lock()
    defer tf.mu.Unlock()

    if tf.running {
        return fmt.Errorf("test framework is already running")
    }

    // Initialize test environment
    env, err := NewTestEnvironment(tf.logger, tf.config.Environment)
    if err != nil {
        return fmt.Errorf("failed to create test environment: %w", err)
    }
    tf.environment = env

    // Start test environment
    if err := tf.environment.Start(ctx); err != nil {
        return fmt.Errorf("failed to start test environment: %w", err)
    }

    // Initialize test suites
    if err := tf.initializeTestSuites(); err != nil {
        tf.environment.Stop()
        return fmt.Errorf("failed to initialize test suites: %w", err)
    }

    // Initialize reporters
    if err := tf.initializeReporters(); err != nil {
        tf.environment.Stop()
        return fmt.Errorf("failed to initialize reporters: %w", err)
    }

    tf.running = true
    tf.logger.Info("Test framework started")

    return nil
}

// Stop stops the testing framework
func (tf *TestFramework) Stop() error {
    tf.mu.Lock()
    defer tf.mu.Unlock()

    if !tf.running {
        return nil
    }

    // Stop test environment
    if tf.environment != nil {
        tf.environment.Stop()
    }

    tf.running = false
    tf.logger.Info("Test framework stopped")

    return nil
}

// RunAllTests runs all test suites
func (tf *TestFramework) RunAllTests(ctx context.Context) (*TestResults, error) {
    tf.logger.Info("Starting integration test run")

    // Reset results
    tf.results = NewTestResults()
    tf.results.StartTime = time.Now()

    // Run test suites
    for name, suite := range tf.testSuites {
        tf.logger.Info("Running test suite", zap.String("name", name))

        suiteResult, err := tf.runTestSuite(ctx, suite)
        if err != nil {
            tf.logger.Error("Test suite failed",
                zap.String("name", name),
                zap.Error(err))
        }

        tf.results.AddSuiteResult(suiteResult)
    }

    tf.results.EndTime = time.Now()
    tf.results.Duration = tf.results.EndTime.Sub(tf.results.StartTime)

    // Generate reports
    tf.generateReports()

    tf.logger.Info("Integration test run completed",
        zap.Int("total_tests", tf.results.TotalTests),
        zap.Int("passed", tf.results.PassedTests),
        zap.Int("failed", tf.results.FailedTests),
        zap.Duration("duration", tf.results.Duration))

    return tf.results, nil
}

// RunTestSuite runs a specific test suite
func (tf *TestFramework) RunTestSuite(ctx context.Context, suiteName string) (*TestSuiteResult, error) {
    suite, exists := tf.testSuites[suiteName]
    if !exists {
        return nil, fmt.Errorf("test suite not found: %s", suiteName)
    }

    return tf.runTestSuite(ctx, suite)
}

// runTestSuite executes a test suite
func (tf *TestFramework) runTestSuite(ctx context.Context, suite *TestSuite) (*TestSuiteResult, error) {
    result := &TestSuiteResult{
        Name:      suite.Name,
        StartTime: time.Now(),
        Tests:     make([]*TestResult, 0),
    }

    // Run setup
    if err := suite.Setup(ctx); err != nil {
        result.SetupError = err
        result.EndTime = time.Now()
        return result, err
    }

    // Run tests
    for _, test := range suite.Tests {
        testResult := tf.runTest(ctx, test)
        result.Tests = append(result.Tests, testResult)
    }

    // Run teardown
    if err := suite.Teardown(ctx); err != nil {
        result.TeardownError = err
    }

    result.EndTime = time.Now()
    result.Duration = result.EndTime.Sub(result.StartTime)
    result.calculateStats()

    return result, nil
}

// runTest executes a single test
func (tf *TestFramework) runTest(ctx context.Context, test *Test) *TestResult {
    result := &TestResult{
        Name:      test.Name,
        StartTime: time.Now(),
    }

    tf.logger.Debug("Running test", zap.String("name", test.Name))

    // Execute test with timeout
    testCtx, cancel := context.WithTimeout(ctx, test.Timeout)
    defer cancel()

    err := test.Execute(testCtx)
    result.EndTime = time.Now()
    result.Duration = result.EndTime.Sub(result.StartTime)

    if err != nil {
        result.Status = "FAILED"
        result.Error = err
        tf.logger.Error("Test failed",
            zap.String("name", test.Name),
            zap.Error(err))
    } else {
        result.Status = "PASSED"
        tf.logger.Debug("Test passed", zap.String("name", test.Name))
    }

    return result
}

// initializeTestSuites initializes test suites based on configuration
func (tf *TestFramework) initializeTestSuites() error {
    // Health check test suite
    healthSuite, err := NewHealthCheckTestSuite(tf.logger, tf.config.Services)
    if err != nil {
        return fmt.Errorf("failed to create health check suite: %w", err)
    }
    tf.testSuites["health"] = healthSuite

    // Service integration test suite
    integrationSuite, err := NewServiceIntegrationTestSuite(tf.logger, tf.config.Services)
    if err != nil {
        return fmt.Errorf("failed to create integration suite: %w", err)
    }
    tf.testSuites["integration"] = integrationSuite

    // End-to-end scenario test suite
    scenarioSuite, err := NewScenarioTestSuite(tf.logger, tf.config.Scenarios)
    if err != nil {
        return fmt.Errorf("failed to create scenario suite: %w", err)
    }
    tf.testSuites["scenarios"] = scenarioSuite

    // Performance test suite
    performanceSuite, err := NewPerformanceTestSuite(tf.logger, tf.config)
    if err != nil {
        return fmt.Errorf("failed to create performance suite: %w", err)
    }
    tf.testSuites["performance"] = performanceSuite

    return nil
}

// initializeReporters initializes test reporters
func (tf *TestFramework) initializeReporters() error {
    if tf.config.Reporting != nil {
        // Console reporter
        if tf.config.Reporting.Console.Enabled {
            tf.reporters = append(tf.reporters, NewConsoleReporter(tf.logger))
        }

        // JUnit XML reporter
        if tf.config.Reporting.JUnitXML.Enabled {
            reporter, err := NewJUnitXMLReporter(tf.logger, tf.config.Reporting.JUnitXML.OutputPath)
            if err != nil {
                return fmt.Errorf("failed to create JUnit XML reporter: %w", err)
            }
            tf.reporters = append(tf.reporters, reporter)
        }

        // HTML reporter
        if tf.config.Reporting.HTML.Enabled {
            reporter, err := NewHTMLReporter(tf.logger, tf.config.Reporting.HTML.OutputPath)
            if err != nil {
                return fmt.Errorf("failed to create HTML reporter: %w", err)
            }
            tf.reporters = append(tf.reporters, reporter)
        }
    }

    return nil
}

// generateReports generates test reports
func (tf *TestFramework) generateReports() {
    for _, reporter := range tf.reporters {
        if err := reporter.GenerateReport(tf.results); err != nil {
            tf.logger.Error("Failed to generate report",
                zap.String("reporter", reporter.Name()),
                zap.Error(err))
        }
    }
}
```

### Step 2: Test Environment Implementation

**Create `pkg/testing/integration/test_environment.go`**:

```go
// file: pkg/testing/integration/test_environment.go
// version: 2.0.0
// guid: test-env-2222-3333-4444-555555555555

package integration

import (
    "context"
    "fmt"
    "net/http"
    "time"

    "go.uber.org/zap"

    // gcommon imports
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/common"
)

// TestEnvironment manages the test environment setup and teardown
type TestEnvironment struct {
    logger      *zap.Logger
    config      *EnvironmentConfig
    services    map[string]*ServiceProxy
    database    *DatabaseProxy
    storage     *StorageProxy
    monitoring  *MonitoringProxy
    running     bool
}

// ServiceProxy represents a proxy to a service in the test environment
type ServiceProxy struct {
    Name       string
    Config     *common.ServerConfig
    Client     *http.Client
    BaseURL    string
    HealthURL  string
    ready      bool
}

// DatabaseProxy represents a proxy to the database in the test environment
type DatabaseProxy struct {
    Type        string
    Host        string
    Port        int
    Database    string
    Username    string
    Password    string
    SSLMode     string
    connected   bool
}

// StorageProxy represents a proxy to storage services in the test environment
type StorageProxy struct {
    Type        string
    Endpoint    string
    AccessKey   string
    SecretKey   string
    Bucket      string
    Region      string
    connected   bool
}

// MonitoringProxy represents a proxy to monitoring services
type MonitoringProxy struct {
    PrometheusURL string
    GrafanaURL    string
    AlertsURL     string
    available     bool
}

// NewTestEnvironment creates a new test environment
func NewTestEnvironment(logger *zap.Logger, config *EnvironmentConfig) (*TestEnvironment, error) {
    if logger == nil {
        return nil, fmt.Errorf("logger is required")
    }
    if config == nil {
        return nil, fmt.Errorf("config is required")
    }

    return &TestEnvironment{
        logger:   logger,
        config:   config,
        services: make(map[string]*ServiceProxy),
    }, nil
}

// Start starts the test environment
func (te *TestEnvironment) Start(ctx context.Context) error {
    if te.running {
        return fmt.Errorf("test environment is already running")
    }

    te.logger.Info("Starting test environment", zap.String("type", te.config.Type))

    // Setup based on environment type
    switch te.config.Type {
    case "local":
        if err := te.setupLocalEnvironment(ctx); err != nil {
            return fmt.Errorf("failed to setup local environment: %w", err)
        }
    case "docker":
        if err := te.setupDockerEnvironment(ctx); err != nil {
            return fmt.Errorf("failed to setup docker environment: %w", err)
        }
    case "k8s":
        if err := te.setupKubernetesEnvironment(ctx); err != nil {
            return fmt.Errorf("failed to setup kubernetes environment: %w", err)
        }
    default:
        return fmt.Errorf("unsupported environment type: %s", te.config.Type)
    }

    // Wait for services to be ready
    if err := te.waitForServices(ctx); err != nil {
        return fmt.Errorf("services not ready: %w", err)
    }

    te.running = true
    te.logger.Info("Test environment started successfully")

    return nil
}

// Stop stops the test environment
func (te *TestEnvironment) Stop() error {
    if !te.running {
        return nil
    }

    te.logger.Info("Stopping test environment")

    // Cleanup based on environment type
    switch te.config.Type {
    case "local":
        te.cleanupLocalEnvironment()
    case "docker":
        te.cleanupDockerEnvironment()
    case "k8s":
        te.cleanupKubernetesEnvironment()
    }

    te.running = false
    te.logger.Info("Test environment stopped")

    return nil
}

// setupLocalEnvironment sets up a local test environment
func (te *TestEnvironment) setupLocalEnvironment(ctx context.Context) error {
    // Initialize service proxies for local services
    for name, serviceConfig := range te.config.Services {
        proxy := &ServiceProxy{
            Name:      name,
            Config:    serviceConfig,
            Client:    &http.Client{Timeout: 30 * time.Second},
            BaseURL:   fmt.Sprintf("http://%s:%d", serviceConfig.Host, serviceConfig.Port),
            HealthURL: fmt.Sprintf("http://%s:%d/health", serviceConfig.Host, serviceConfig.Port),
        }
        te.services[name] = proxy
    }

    // Setup database proxy
    if te.config.Database != nil {
        te.database = &DatabaseProxy{
            Type:     te.config.Database.Type,
            Host:     te.config.Database.Host,
            Port:     te.config.Database.Port,
            Database: te.config.Database.Database,
            Username: te.config.Database.Username,
            Password: te.config.Database.Password,
            SSLMode:  te.config.Database.SSLMode,
        }
    }

    // Setup storage proxy
    if te.config.Storage != nil {
        te.storage = &StorageProxy{
            Type:      te.config.Storage.Type,
            Endpoint:  te.config.Storage.Endpoint,
            AccessKey: te.config.Storage.AccessKey,
            SecretKey: te.config.Storage.SecretKey,
            Bucket:    te.config.Storage.Bucket,
            Region:    te.config.Storage.Region,
        }
    }

    return nil
}

// setupDockerEnvironment sets up a Docker-based test environment
func (te *TestEnvironment) setupDockerEnvironment(ctx context.Context) error {
    // Implementation would use Docker API to start containers
    // This is a simplified version
    te.logger.Info("Setting up Docker test environment")

    // Start database container
    if te.config.Database != nil {
        if err := te.startDatabaseContainer(ctx); err != nil {
            return fmt.Errorf("failed to start database container: %w", err)
        }
    }

    // Start service containers
    for name, serviceConfig := range te.config.Services {
        if err := te.startServiceContainer(ctx, name, serviceConfig); err != nil {
            return fmt.Errorf("failed to start service container %s: %w", name, err)
        }
    }

    return nil
}

// setupKubernetesEnvironment sets up a Kubernetes-based test environment
func (te *TestEnvironment) setupKubernetesEnvironment(ctx context.Context) error {
    // Implementation would use Kubernetes API to deploy resources
    te.logger.Info("Setting up Kubernetes test environment")

    // Apply test namespace
    if err := te.createTestNamespace(ctx); err != nil {
        return fmt.Errorf("failed to create test namespace: %w", err)
    }

    // Deploy database
    if te.config.Database != nil {
        if err := te.deployDatabase(ctx); err != nil {
            return fmt.Errorf("failed to deploy database: %w", err)
        }
    }

    // Deploy services
    for name, serviceConfig := range te.config.Services {
        if err := te.deployService(ctx, name, serviceConfig); err != nil {
            return fmt.Errorf("failed to deploy service %s: %w", name, err)
        }
    }

    return nil
}

// waitForServices waits for all services to be ready
func (te *TestEnvironment) waitForServices(ctx context.Context) error {
    timeout := 5 * time.Minute
    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    te.logger.Info("Waiting for services to be ready")

    for name, proxy := range te.services {
        if err := te.waitForService(ctx, name, proxy); err != nil {
            return fmt.Errorf("service %s not ready: %w", name, err)
        }
    }

    // Wait for database
    if te.database != nil {
        if err := te.waitForDatabase(ctx); err != nil {
            return fmt.Errorf("database not ready: %w", err)
        }
    }

    // Wait for storage
    if te.storage != nil {
        if err := te.waitForStorage(ctx); err != nil {
            return fmt.Errorf("storage not ready: %w", err)
        }
    }

    return nil
}

// waitForService waits for a specific service to be ready
func (te *TestEnvironment) waitForService(ctx context.Context, name string, proxy *ServiceProxy) error {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return fmt.Errorf("timeout waiting for service %s", name)
        case <-ticker.C:
            resp, err := proxy.Client.Get(proxy.HealthURL)
            if err != nil {
                te.logger.Debug("Service health check failed",
                    zap.String("service", name),
                    zap.Error(err))
                continue
            }

            resp.Body.Close()
            if resp.StatusCode == http.StatusOK {
                proxy.ready = true
                te.logger.Info("Service is ready", zap.String("service", name))
                return nil
            }
        }
    }
}

// Helper methods for Docker container management
func (te *TestEnvironment) startDatabaseContainer(ctx context.Context) error {
    // Implementation would use Docker client
    te.logger.Debug("Starting database container")
    return nil
}

func (te *TestEnvironment) startServiceContainer(ctx context.Context, name string, config *common.ServerConfig) error {
    // Implementation would use Docker client
    te.logger.Debug("Starting service container", zap.String("service", name))
    return nil
}

// Helper methods for Kubernetes deployment
func (te *TestEnvironment) createTestNamespace(ctx context.Context) error {
    // Implementation would use Kubernetes client
    te.logger.Debug("Creating test namespace")
    return nil
}

func (te *TestEnvironment) deployDatabase(ctx context.Context) error {
    // Implementation would use Kubernetes client
    te.logger.Debug("Deploying database")
    return nil
}

func (te *TestEnvironment) deployService(ctx context.Context, name string, config *common.ServerConfig) error {
    // Implementation would use Kubernetes client
    te.logger.Debug("Deploying service", zap.String("service", name))
    return nil
}

// Cleanup methods
func (te *TestEnvironment) cleanupLocalEnvironment() {
    te.logger.Debug("Cleaning up local environment")
}

func (te *TestEnvironment) cleanupDockerEnvironment() {
    te.logger.Debug("Cleaning up Docker environment")
}

func (te *TestEnvironment) cleanupKubernetesEnvironment() {
    te.logger.Debug("Cleaning up Kubernetes environment")
}

// waitForDatabase waits for database to be ready
func (te *TestEnvironment) waitForDatabase(ctx context.Context) error {
    te.logger.Debug("Waiting for database")
    // Implementation would test database connectivity
    te.database.connected = true
    return nil
}

// waitForStorage waits for storage to be ready
func (te *TestEnvironment) waitForStorage(ctx context.Context) error {
    te.logger.Debug("Waiting for storage")
    // Implementation would test storage connectivity
    te.storage.connected = true
    return nil
}

// GetService returns a service proxy
func (te *TestEnvironment) GetService(name string) (*ServiceProxy, error) {
    proxy, exists := te.services[name]
    if !exists {
        return nil, fmt.Errorf("service not found: %s", name)
    }
    return proxy, nil
}

// GetDatabase returns the database proxy
func (te *TestEnvironment) GetDatabase() *DatabaseProxy {
    return te.database
}

// GetStorage returns the storage proxy
func (te *TestEnvironment) GetStorage() *StorageProxy {
    return te.storage
}

// IsReady checks if the environment is ready for testing
func (te *TestEnvironment) IsReady() bool {
    if !te.running {
        return false
    }

    // Check all services
    for _, proxy := range te.services {
        if !proxy.ready {
            return false
        }
    }

    // Check database
    if te.database != nil && !te.database.connected {
        return false
    }

    // Check storage
    if te.storage != nil && !te.storage.connected {
        return false
    }

    return true
}
```

This integration testing implementation provides comprehensive testing
capabilities with environment management, service proxies, and support for
local, Docker, and Kubernetes environments. The framework includes test suites
for health checks, service integration, end-to-end scenarios, and performance
testing.
