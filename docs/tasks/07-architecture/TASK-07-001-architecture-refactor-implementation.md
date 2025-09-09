# TASK-07-001: Architecture Refactor Implementation (gcommon Edition)

<!-- file: docs/tasks/07-architecture/TASK-07-001-architecture-refactor-implementation.md -->
<!-- version: 2.0.0 -->
<!-- guid: arch07001-aaaa-bbbb-cccc-dddddddddddd -->

## Overview

Refactor the subtitle-manager architecture to implement a clean, microservices-based design with full gcommon integration, improved separation of concerns, and enhanced maintainability.

## Implementation Plan

### Step 1: Core Architecture Framework

**Create `pkg/architecture/framework.go`**:

```go
// file: pkg/architecture/framework.go
// version: 2.0.0
// guid: arch-framework-1111-2222-3333-444444444444

package architecture

import (
    "context"
    "fmt"
    "sync"

    "go.uber.org/zap"

    // gcommon imports
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/common"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/config"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/health"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/metrics"
)

// ServiceFramework provides a unified framework for all microservices
type ServiceFramework struct {
    // Core components
    logger        *zap.Logger
    config        *common.Config
    healthChecker *health.Checker
    metrics       *metrics.Collector

    // Service registry
    services      map[string]Service
    dependencies  map[string][]string

    // Lifecycle management
    startOrder    []string
    stopOrder     []string
    running       bool
    mu            sync.RWMutex
}

// Service interface for all microservices
type Service interface {
    // Lifecycle
    Start(ctx context.Context) error
    Stop() error
    IsHealthy() bool

    // Metadata
    Name() string
    Version() string
    Dependencies() []string

    // Metrics
    GetMetrics() map[string]interface{}

    // Configuration
    GetConfig() interface{}
    UpdateConfig(config interface{}) error
}

// ServiceConfig defines common service configuration
type ServiceConfig struct {
    Name         string                  `yaml:"name" json:"name"`
    Version      string                  `yaml:"version" json:"version"`
    Description  string                  `yaml:"description" json:"description"`
    Server       *common.ServerConfig    `yaml:"server" json:"server"`
    Health       *health.Config          `yaml:"health" json:"health"`
    Metrics      *metrics.Config         `yaml:"metrics" json:"metrics"`
    Dependencies []string                `yaml:"dependencies" json:"dependencies"`
    Features     map[string]interface{}  `yaml:"features" json:"features"`
}

// NewServiceFramework creates a new service framework
func NewServiceFramework(logger *zap.Logger, config *common.Config) (*ServiceFramework, error) {
    if logger == nil {
        return nil, fmt.Errorf("logger is required")
    }
    if config == nil {
        return nil, fmt.Errorf("config is required")
    }

    healthChecker := health.NewChecker(logger)
    metricsCollector := metrics.NewCollector(logger, "service-framework")

    return &ServiceFramework{
        logger:        logger,
        config:        config,
        healthChecker: healthChecker,
        metrics:       metricsCollector,
        services:      make(map[string]Service),
        dependencies:  make(map[string][]string),
    }, nil
}

// RegisterService registers a service with the framework
func (sf *ServiceFramework) RegisterService(service Service) error {
    sf.mu.Lock()
    defer sf.mu.Unlock()

    name := service.Name()
    if _, exists := sf.services[name]; exists {
        return fmt.Errorf("service already registered: %s", name)
    }

    sf.services[name] = service
    sf.dependencies[name] = service.Dependencies()

    // Update health checker
    sf.healthChecker.RegisterComponent(name, service)

    sf.logger.Info("Service registered",
        zap.String("name", name),
        zap.String("version", service.Version()),
        zap.Strings("dependencies", service.Dependencies()))

    return nil
}

// Start starts all services in dependency order
func (sf *ServiceFramework) Start(ctx context.Context) error {
    sf.mu.Lock()
    defer sf.mu.Unlock()

    if sf.running {
        return fmt.Errorf("service framework is already running")
    }

    // Calculate start order based on dependencies
    startOrder, err := sf.calculateStartOrder()
    if err != nil {
        return fmt.Errorf("failed to calculate start order: %w", err)
    }

    sf.startOrder = startOrder
    sf.stopOrder = sf.reverseOrder(startOrder)

    // Start services in order
    for _, serviceName := range sf.startOrder {
        service := sf.services[serviceName]
        if err := service.Start(ctx); err != nil {
            // Rollback: stop already started services
            sf.stopStartedServices(serviceName)
            return fmt.Errorf("failed to start service %s: %w", serviceName, err)
        }

        sf.logger.Info("Service started", zap.String("name", serviceName))
    }

    sf.running = true
    sf.logger.Info("Service framework started", zap.Strings("start_order", sf.startOrder))

    return nil
}

// Stop stops all services in reverse dependency order
func (sf *ServiceFramework) Stop() error {
    sf.mu.Lock()
    defer sf.mu.Unlock()

    if !sf.running {
        return nil
    }

    // Stop services in reverse order
    for _, serviceName := range sf.stopOrder {
        service := sf.services[serviceName]
        if err := service.Stop(); err != nil {
            sf.logger.Error("Failed to stop service",
                zap.String("name", serviceName),
                zap.Error(err))
        } else {
            sf.logger.Info("Service stopped", zap.String("name", serviceName))
        }
    }

    sf.running = false
    sf.logger.Info("Service framework stopped")

    return nil
}

// IsHealthy checks if all services are healthy
func (sf *ServiceFramework) IsHealthy() bool {
    sf.mu.RLock()
    defer sf.mu.RUnlock()

    return sf.running && sf.healthChecker.IsHealthy()
}

// GetMetrics returns framework and service metrics
func (sf *ServiceFramework) GetMetrics() map[string]interface{} {
    sf.mu.RLock()
    defer sf.mu.RUnlock()

    result := make(map[string]interface{})
    result["framework"] = sf.metrics.GetMetrics()
    result["total_services"] = len(sf.services)
    result["running"] = sf.running

    // Collect service metrics
    services := make(map[string]interface{})
    for name, service := range sf.services {
        services[name] = service.GetMetrics()
    }
    result["services"] = services

    return result
}

// calculateStartOrder calculates the order to start services based on dependencies
func (sf *ServiceFramework) calculateStartOrder() ([]string, error) {
    // Topological sort implementation
    visited := make(map[string]bool)
    visiting := make(map[string]bool)
    result := make([]string, 0)

    var visit func(string) error
    visit = func(serviceName string) error {
        if visiting[serviceName] {
            return fmt.Errorf("circular dependency detected involving service: %s", serviceName)
        }
        if visited[serviceName] {
            return nil
        }

        visiting[serviceName] = true

        // Visit dependencies first
        for _, dep := range sf.dependencies[serviceName] {
            if _, exists := sf.services[dep]; !exists {
                return fmt.Errorf("dependency not found: %s required by %s", dep, serviceName)
            }
            if err := visit(dep); err != nil {
                return err
            }
        }

        visiting[serviceName] = false
        visited[serviceName] = true
        result = append(result, serviceName)

        return nil
    }

    // Visit all services
    for serviceName := range sf.services {
        if err := visit(serviceName); err != nil {
            return nil, err
        }
    }

    return result, nil
}

// reverseOrder creates a reverse order for stopping services
func (sf *ServiceFramework) reverseOrder(order []string) []string {
    result := make([]string, len(order))
    for i, serviceName := range order {
        result[len(order)-1-i] = serviceName
    }
    return result
}

// stopStartedServices stops services that were already started during a failed startup
func (sf *ServiceFramework) stopStartedServices(failedService string) {
    for _, serviceName := range sf.startOrder {
        if serviceName == failedService {
            break
        }
        service := sf.services[serviceName]
        if err := service.Stop(); err != nil {
            sf.logger.Error("Failed to stop service during rollback",
                zap.String("name", serviceName),
                zap.Error(err))
        }
    }
}

// GetService returns a registered service
func (sf *ServiceFramework) GetService(name string) (Service, error) {
    sf.mu.RLock()
    defer sf.mu.RUnlock()

    service, exists := sf.services[name]
    if !exists {
        return nil, fmt.Errorf("service not found: %s", name)
    }

    return service, nil
}

// ListServices returns all registered services
func (sf *ServiceFramework) ListServices() []string {
    sf.mu.RLock()
    defer sf.mu.RUnlock()

    services := make([]string, 0, len(sf.services))
    for name := range sf.services {
        services = append(services, name)
    }

    return services
}
```

### Step 2: Service Base Implementation

**Create `pkg/architecture/service_base.go`**:

```go
// file: pkg/architecture/service_base.go
// version: 2.0.0
// guid: arch-servicebase-2222-3333-4444-555555555555

package architecture

import (
    "context"
    "fmt"
    "sync"

    "go.uber.org/zap"
    "google.golang.org/grpc"

    // gcommon imports
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/common"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/health"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/metrics"
)

// BaseService provides common functionality for all services
type BaseService struct {
    // Service metadata
    name         string
    version      string
    description  string
    dependencies []string

    // Core components
    logger        *zap.Logger
    config        *ServiceConfig
    healthChecker *health.Checker
    metrics       *metrics.Collector

    // gRPC server
    server   *grpc.Server
    listener net.Listener

    // Lifecycle
    running   bool
    mu        sync.RWMutex
    cancelCtx context.CancelFunc

    // Extension points
    startHooks []StartHook
    stopHooks  []StopHook
}

// StartHook defines a hook that runs during service startup
type StartHook func(ctx context.Context, service *BaseService) error

// StopHook defines a hook that runs during service shutdown
type StopHook func(service *BaseService) error

// NewBaseService creates a new base service
func NewBaseService(logger *zap.Logger, config *ServiceConfig) (*BaseService, error) {
    if logger == nil {
        return nil, fmt.Errorf("logger is required")
    }
    if config == nil {
        return nil, fmt.Errorf("config is required")
    }

    healthChecker := health.NewChecker(logger)
    metricsCollector := metrics.NewCollector(logger, config.Name)

    return &BaseService{
        name:          config.Name,
        version:       config.Version,
        description:   config.Description,
        dependencies:  config.Dependencies,
        logger:        logger,
        config:        config,
        healthChecker: healthChecker,
        metrics:       metricsCollector,
        startHooks:    make([]StartHook, 0),
        stopHooks:     make([]StopHook, 0),
    }, nil
}

// Start starts the base service
func (bs *BaseService) Start(ctx context.Context) error {
    bs.mu.Lock()
    defer bs.mu.Unlock()

    if bs.running {
        return fmt.Errorf("service %s is already running", bs.name)
    }

    // Create cancellable context
    ctx, bs.cancelCtx = context.WithCancel(ctx)

    // Initialize gRPC server if configured
    if bs.config.Server != nil {
        if err := bs.initializeGRPCServer(); err != nil {
            return fmt.Errorf("failed to initialize gRPC server: %w", err)
        }
    }

    // Run start hooks
    for _, hook := range bs.startHooks {
        if err := hook(ctx, bs); err != nil {
            bs.cleanup()
            return fmt.Errorf("start hook failed: %w", err)
        }
    }

    // Start gRPC server
    if bs.server != nil {
        go func() {
            if err := bs.server.Serve(bs.listener); err != nil {
                bs.logger.Error("gRPC server failed", zap.Error(err))
            }
        }()
    }

    bs.running = true
    bs.logger.Info("Service started",
        zap.String("name", bs.name),
        zap.String("version", bs.version))

    return nil
}

// Stop stops the base service
func (bs *BaseService) Stop() error {
    bs.mu.Lock()
    defer bs.mu.Unlock()

    if !bs.running {
        return nil
    }

    // Cancel context
    if bs.cancelCtx != nil {
        bs.cancelCtx()
    }

    // Run stop hooks
    for _, hook := range bs.stopHooks {
        if err := hook(bs); err != nil {
            bs.logger.Error("Stop hook failed", zap.Error(err))
        }
    }

    // Stop gRPC server
    if bs.server != nil {
        bs.server.GracefulStop()
        bs.server = nil
    }

    // Close listener
    if bs.listener != nil {
        bs.listener.Close()
        bs.listener = nil
    }

    bs.running = false
    bs.logger.Info("Service stopped", zap.String("name", bs.name))

    return nil
}

// IsHealthy checks if the service is healthy
func (bs *BaseService) IsHealthy() bool {
    bs.mu.RLock()
    defer bs.mu.RUnlock()

    return bs.running && bs.healthChecker.IsHealthy()
}

// Name returns the service name
func (bs *BaseService) Name() string {
    return bs.name
}

// Version returns the service version
func (bs *BaseService) Version() string {
    return bs.version
}

// Dependencies returns the service dependencies
func (bs *BaseService) Dependencies() []string {
    return bs.dependencies
}

// GetMetrics returns service metrics
func (bs *BaseService) GetMetrics() map[string]interface{} {
    bs.mu.RLock()
    defer bs.mu.RUnlock()

    result := bs.metrics.GetMetrics()
    result["running"] = bs.running
    result["health_status"] = bs.healthChecker.GetStatus()

    return result
}

// GetConfig returns the service configuration
func (bs *BaseService) GetConfig() interface{} {
    return bs.config
}

// UpdateConfig updates the service configuration
func (bs *BaseService) UpdateConfig(config interface{}) error {
    serviceConfig, ok := config.(*ServiceConfig)
    if !ok {
        return fmt.Errorf("invalid config type")
    }

    bs.mu.Lock()
    defer bs.mu.Unlock()

    bs.config = serviceConfig
    bs.logger.Info("Service configuration updated", zap.String("name", bs.name))

    return nil
}

// AddStartHook adds a start hook
func (bs *BaseService) AddStartHook(hook StartHook) {
    bs.startHooks = append(bs.startHooks, hook)
}

// AddStopHook adds a stop hook
func (bs *BaseService) AddStopHook(hook StopHook) {
    bs.stopHooks = append(bs.stopHooks, hook)
}

// GetLogger returns the service logger
func (bs *BaseService) GetLogger() *zap.Logger {
    return bs.logger
}

// GetHealthChecker returns the health checker
func (bs *BaseService) GetHealthChecker() *health.Checker {
    return bs.healthChecker
}

// GetMetricsCollector returns the metrics collector
func (bs *BaseService) GetMetricsCollector() *metrics.Collector {
    return bs.metrics
}

// GetGRPCServer returns the gRPC server
func (bs *BaseService) GetGRPCServer() *grpc.Server {
    return bs.server
}

// initializeGRPCServer initializes the gRPC server
func (bs *BaseService) initializeGRPCServer() error {
    var err error

    bs.listener, err = common.CreateListener(bs.config.Server)
    if err != nil {
        return fmt.Errorf("failed to create listener: %w", err)
    }

    bs.server = grpc.NewServer()

    return nil
}

// cleanup performs cleanup during failed startup
func (bs *BaseService) cleanup() {
    if bs.cancelCtx != nil {
        bs.cancelCtx()
    }

    if bs.server != nil {
        bs.server.Stop()
        bs.server = nil
    }

    if bs.listener != nil {
        bs.listener.Close()
        bs.listener = nil
    }
}
```

### Step 3: Dependency Injection System

**Create `pkg/architecture/dependency_injection.go`**:

```go
// file: pkg/architecture/dependency_injection.go
// version: 2.0.0
// guid: arch-di-3333-4444-5555-666666666666

package architecture

import (
    "fmt"
    "reflect"
    "sync"

    "go.uber.org/zap"
)

// Container provides dependency injection functionality
type Container struct {
    logger     *zap.Logger
    providers  map[string]Provider
    instances  map[string]interface{}
    singletons map[string]bool
    mu         sync.RWMutex
}

// Provider defines a dependency provider function
type Provider func(container *Container) (interface{}, error)

// Injectable interface for services that can receive dependencies
type Injectable interface {
    Inject(dependencies map[string]interface{}) error
}

// NewContainer creates a new dependency injection container
func NewContainer(logger *zap.Logger) *Container {
    return &Container{
        logger:     logger,
        providers:  make(map[string]Provider),
        instances:  make(map[string]interface{}),
        singletons: make(map[string]bool),
    }
}

// Register registers a provider for a dependency
func (c *Container) Register(name string, provider Provider, singleton bool) error {
    c.mu.Lock()
    defer c.mu.Unlock()

    if _, exists := c.providers[name]; exists {
        return fmt.Errorf("provider already registered: %s", name)
    }

    c.providers[name] = provider
    c.singletons[name] = singleton

    c.logger.Debug("Provider registered",
        zap.String("name", name),
        zap.Bool("singleton", singleton))

    return nil
}

// RegisterInstance registers a specific instance
func (c *Container) RegisterInstance(name string, instance interface{}) error {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.instances[name] = instance
    c.singletons[name] = true

    c.logger.Debug("Instance registered", zap.String("name", name))

    return nil
}

// Get retrieves a dependency by name
func (c *Container) Get(name string) (interface{}, error) {
    c.mu.RLock()

    // Check for existing instance
    if instance, exists := c.instances[name]; exists {
        c.mu.RUnlock()
        return instance, nil
    }

    // Check for provider
    provider, exists := c.providers[name]
    if !exists {
        c.mu.RUnlock()
        return nil, fmt.Errorf("dependency not found: %s", name)
    }

    isSingleton := c.singletons[name]
    c.mu.RUnlock()

    // Create instance
    instance, err := provider(c)
    if err != nil {
        return nil, fmt.Errorf("failed to create instance for %s: %w", name, err)
    }

    // Cache singleton instances
    if isSingleton {
        c.mu.Lock()
        c.instances[name] = instance
        c.mu.Unlock()
    }

    return instance, nil
}

// InjectDependencies injects dependencies into a service
func (c *Container) InjectDependencies(service Injectable, dependencies []string) error {
    deps := make(map[string]interface{})

    for _, depName := range dependencies {
        dep, err := c.Get(depName)
        if err != nil {
            return fmt.Errorf("failed to get dependency %s: %w", depName, err)
        }
        deps[depName] = dep
    }

    return service.Inject(deps)
}

// AutoInject automatically injects dependencies based on struct tags
func (c *Container) AutoInject(target interface{}) error {
    v := reflect.ValueOf(target)
    if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
        return fmt.Errorf("target must be a pointer to struct")
    }

    v = v.Elem()
    t := v.Type()

    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldType := t.Field(i)

        // Check for inject tag
        injectTag := fieldType.Tag.Get("inject")
        if injectTag == "" {
            continue
        }

        // Skip if field is not settable
        if !field.CanSet() {
            continue
        }

        // Get dependency
        dep, err := c.Get(injectTag)
        if err != nil {
            return fmt.Errorf("failed to inject %s: %w", injectTag, err)
        }

        // Set field value
        depValue := reflect.ValueOf(dep)
        if !depValue.Type().AssignableTo(field.Type()) {
            return fmt.Errorf("dependency %s is not assignable to field %s", injectTag, fieldType.Name)
        }

        field.Set(depValue)
        c.logger.Debug("Dependency injected",
            zap.String("field", fieldType.Name),
            zap.String("dependency", injectTag))
    }

    return nil
}

// BuildServiceGraph builds a dependency graph for services
func (c *Container) BuildServiceGraph(services []Service) (*ServiceGraph, error) {
    graph := NewServiceGraph()

    // Add all services to graph
    for _, service := range services {
        graph.AddService(service.Name(), service.Dependencies())
    }

    // Validate graph (check for cycles)
    if err := graph.Validate(); err != nil {
        return nil, fmt.Errorf("invalid service graph: %w", err)
    }

    return graph, nil
}

// ServiceGraph represents the dependency graph of services
type ServiceGraph struct {
    services     map[string][]string
    dependents   map[string][]string
    mu           sync.RWMutex
}

// NewServiceGraph creates a new service graph
func NewServiceGraph() *ServiceGraph {
    return &ServiceGraph{
        services:   make(map[string][]string),
        dependents: make(map[string][]string),
    }
}

// AddService adds a service and its dependencies to the graph
func (sg *ServiceGraph) AddService(name string, dependencies []string) {
    sg.mu.Lock()
    defer sg.mu.Unlock()

    sg.services[name] = dependencies

    // Build reverse dependency map
    for _, dep := range dependencies {
        sg.dependents[dep] = append(sg.dependents[dep], name)
    }
}

// Validate validates the service graph for circular dependencies
func (sg *ServiceGraph) Validate() error {
    sg.mu.RLock()
    defer sg.mu.RUnlock()

    visited := make(map[string]bool)
    visiting := make(map[string]bool)

    var visit func(string) error
    visit = func(service string) error {
        if visiting[service] {
            return fmt.Errorf("circular dependency detected: %s", service)
        }
        if visited[service] {
            return nil
        }

        visiting[service] = true

        for _, dep := range sg.services[service] {
            if err := visit(dep); err != nil {
                return err
            }
        }

        visiting[service] = false
        visited[service] = true

        return nil
    }

    for service := range sg.services {
        if err := visit(service); err != nil {
            return err
        }
    }

    return nil
}

// GetStartOrder returns the order in which services should be started
func (sg *ServiceGraph) GetStartOrder() []string {
    sg.mu.RLock()
    defer sg.mu.RUnlock()

    visited := make(map[string]bool)
    result := make([]string, 0)

    var visit func(string)
    visit = func(service string) {
        if visited[service] {
            return
        }

        // Visit dependencies first
        for _, dep := range sg.services[service] {
            visit(dep)
        }

        visited[service] = true
        result = append(result, service)
    }

    for service := range sg.services {
        visit(service)
    }

    return result
}

// GetStopOrder returns the order in which services should be stopped
func (sg *ServiceGraph) GetStopOrder() []string {
    startOrder := sg.GetStartOrder()
    stopOrder := make([]string, len(startOrder))

    for i, service := range startOrder {
        stopOrder[len(startOrder)-1-i] = service
    }

    return stopOrder
}
```

This architecture refactor provides a comprehensive framework for microservices with dependency injection, service lifecycle management, and clean separation of concerns. The implementation includes full gcommon integration and supports complex dependency graphs with validation.
