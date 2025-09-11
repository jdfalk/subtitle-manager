# TASK-06-001: Monitoring and Observability Implementation (gcommon Edition)

<!-- file: docs/tasks/06-monitoring/TASK-06-001-monitoring-implementation.md -->
<!-- version: 2.0.0 -->
<!-- guid: monitor06001-aaaa-bbbb-cccc-dddddddddddd -->

## Overview

Implement comprehensive monitoring and observability for the subtitle-manager
ecosystem with full gcommon integration, providing metrics, logging, tracing,
and alerting capabilities.

## Implementation Plan

### Step 1: Monitoring Service Structure

**Create `pkg/services/monitoring/service.go`**:

```go
// file: pkg/services/monitoring/service.go
// version: 2.0.0
// guid: monitor-service-1111-2222-3333-444444444444

package monitoring

import (
    "context"
    "fmt"
    "sync"
    "time"

    "go.uber.org/zap"
    "google.golang.org/grpc"

    // gcommon imports
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/common"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/config"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/health"
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/metrics"

    // Generated protobuf types
    monitoringpb "github.com/jdfalk/subtitle-manager/pkg/proto/monitoring/v1"
)

// MonitoringService implements comprehensive monitoring and observability
type MonitoringService struct {
    monitoringpb.UnimplementedMonitoringServiceServer

    // Core components
    logger        *zap.Logger
    config        *MonitoringConfig
    healthChecker *health.Checker
    metrics       *metrics.Collector

    // Monitoring components
    metricsCollector *MetricsCollector
    alertManager     *AlertManager
    traceCollector   *TraceCollector
    logAggregator    *LogAggregator
    dashboardServer  *DashboardServer

    // Service lifecycle
    server    *grpc.Server
    running   bool
    mu        sync.RWMutex
    cancelCtx context.CancelFunc
}

// MonitoringConfig defines monitoring service configuration
type MonitoringConfig struct {
    Server    *common.ServerConfig    `yaml:"server" json:"server"`
    Metrics   *MetricsConfig          `yaml:"metrics" json:"metrics"`
    Alerts    *AlertsConfig           `yaml:"alerts" json:"alerts"`
    Tracing   *TracingConfig          `yaml:"tracing" json:"tracing"`
    Logging   *LoggingConfig          `yaml:"logging" json:"logging"`
    Dashboard *DashboardConfig        `yaml:"dashboard" json:"dashboard"`
}

// MetricsConfig defines metrics collection configuration
type MetricsConfig struct {
    Enabled            bool                    `yaml:"enabled" json:"enabled"`
    CollectionInterval time.Duration          `yaml:"collection_interval" json:"collection_interval"`
    RetentionPeriod    time.Duration          `yaml:"retention_period" json:"retention_period"`
    StorageBackend     string                  `yaml:"storage_backend" json:"storage_backend"`
    Exporters          map[string]*config.Config `yaml:"exporters" json:"exporters"`
    CustomMetrics      []*CustomMetricConfig   `yaml:"custom_metrics" json:"custom_metrics"`
}

// AlertsConfig defines alerting configuration
type AlertsConfig struct {
    Enabled      bool                       `yaml:"enabled" json:"enabled"`
    Rules        []*AlertRule               `yaml:"rules" json:"rules"`
    Channels     map[string]*AlertChannel   `yaml:"channels" json:"channels"`
    Escalation   *EscalationConfig          `yaml:"escalation" json:"escalation"`
}

// TracingConfig defines distributed tracing configuration
type TracingConfig struct {
    Enabled     bool           `yaml:"enabled" json:"enabled"`
    Sampler     *SamplerConfig `yaml:"sampler" json:"sampler"`
    Exporters   []string       `yaml:"exporters" json:"exporters"`
    Attributes  map[string]string `yaml:"attributes" json:"attributes"`
}

// LoggingConfig defines log aggregation configuration
type LoggingConfig struct {
    Enabled        bool                    `yaml:"enabled" json:"enabled"`
    Level          string                  `yaml:"level" json:"level"`
    Formatters     map[string]*config.Config `yaml:"formatters" json:"formatters"`
    Destinations   map[string]*config.Config `yaml:"destinations" json:"destinations"`
    Retention      *LogRetentionConfig     `yaml:"retention" json:"retention"`
}

// DashboardConfig defines dashboard server configuration
type DashboardConfig struct {
    Enabled      bool   `yaml:"enabled" json:"enabled"`
    Host         string `yaml:"host" json:"host"`
    Port         int    `yaml:"port" json:"port"`
    TLSEnabled   bool   `yaml:"tls_enabled" json:"tls_enabled"`
    AuthEnabled  bool   `yaml:"auth_enabled" json:"auth_enabled"`
}

// NewMonitoringService creates a new monitoring service
func NewMonitoringService(logger *zap.Logger, config *MonitoringConfig) (*MonitoringService, error) {
    if logger == nil {
        return nil, fmt.Errorf("logger is required")
    }
    if config == nil {
        return nil, fmt.Errorf("config is required")
    }

    // Initialize health checker
    healthChecker := health.NewChecker(logger)

    // Initialize metrics collector
    metricsCollector := metrics.NewCollector(logger, config.Server.ServiceName)

    // Initialize monitoring components
    ms := &MonitoringService{
        logger:        logger,
        config:        config,
        healthChecker: healthChecker,
        metrics:       metricsCollector,
    }

    // Initialize components based on configuration
    if err := ms.initializeComponents(); err != nil {
        return nil, fmt.Errorf("failed to initialize components: %w", err)
    }

    return ms, nil
}

// Start starts the monitoring service
func (ms *MonitoringService) Start(ctx context.Context) error {
    ms.mu.Lock()
    defer ms.mu.Unlock()

    if ms.running {
        return fmt.Errorf("monitoring service is already running")
    }

    // Create cancellable context
    ctx, ms.cancelCtx = context.WithCancel(ctx)

    // Start gRPC server
    if err := ms.startGRPCServer(ctx); err != nil {
        return fmt.Errorf("failed to start gRPC server: %w", err)
    }

    // Start monitoring components
    if err := ms.startComponents(ctx); err != nil {
        ms.stopGRPCServer()
        return fmt.Errorf("failed to start components: %w", err)
    }

    ms.running = true
    ms.logger.Info("Monitoring service started",
        zap.String("address", fmt.Sprintf("%s:%d", ms.config.Server.Host, ms.config.Server.Port)))

    return nil
}

// Stop stops the monitoring service
func (ms *MonitoringService) Stop() error {
    ms.mu.Lock()
    defer ms.mu.Unlock()

    if !ms.running {
        return nil
    }

    // Cancel context to stop all background operations
    if ms.cancelCtx != nil {
        ms.cancelCtx()
    }

    // Stop components
    ms.stopComponents()

    // Stop gRPC server
    ms.stopGRPCServer()

    ms.running = false
    ms.logger.Info("Monitoring service stopped")

    return nil
}

// IsHealthy checks if the monitoring service is healthy
func (ms *MonitoringService) IsHealthy() bool {
    ms.mu.RLock()
    defer ms.mu.RUnlock()

    if !ms.running {
        return false
    }

    // Check component health
    return ms.healthChecker.IsHealthy()
}

// GetMetrics returns service metrics
func (ms *MonitoringService) GetMetrics() map[string]interface{} {
    ms.mu.RLock()
    defer ms.mu.RUnlock()

    return ms.metrics.GetMetrics()
}

// initializeComponents initializes monitoring components based on configuration
func (ms *MonitoringService) initializeComponents() error {
    var err error

    // Initialize metrics collector
    if ms.config.Metrics.Enabled {
        ms.metricsCollector, err = NewMetricsCollector(ms.logger, ms.config.Metrics)
        if err != nil {
            return fmt.Errorf("failed to initialize metrics collector: %w", err)
        }
    }

    // Initialize alert manager
    if ms.config.Alerts.Enabled {
        ms.alertManager, err = NewAlertManager(ms.logger, ms.config.Alerts)
        if err != nil {
            return fmt.Errorf("failed to initialize alert manager: %w", err)
        }
    }

    // Initialize trace collector
    if ms.config.Tracing.Enabled {
        ms.traceCollector, err = NewTraceCollector(ms.logger, ms.config.Tracing)
        if err != nil {
            return fmt.Errorf("failed to initialize trace collector: %w", err)
        }
    }

    // Initialize log aggregator
    if ms.config.Logging.Enabled {
        ms.logAggregator, err = NewLogAggregator(ms.logger, ms.config.Logging)
        if err != nil {
            return fmt.Errorf("failed to initialize log aggregator: %w", err)
        }
    }

    // Initialize dashboard server
    if ms.config.Dashboard.Enabled {
        ms.dashboardServer, err = NewDashboardServer(ms.logger, ms.config.Dashboard)
        if err != nil {
            return fmt.Errorf("failed to initialize dashboard server: %w", err)
        }
    }

    return nil
}

// startGRPCServer starts the gRPC server
func (ms *MonitoringService) startGRPCServer(ctx context.Context) error {
    ms.server = grpc.NewServer()
    monitoringpb.RegisterMonitoringServiceServer(ms.server, ms)

    listener, err := common.CreateListener(ms.config.Server)
    if err != nil {
        return fmt.Errorf("failed to create listener: %w", err)
    }

    go func() {
        if err := ms.server.Serve(listener); err != nil {
            ms.logger.Error("gRPC server failed", zap.Error(err))
        }
    }()

    return nil
}

// stopGRPCServer stops the gRPC server
func (ms *MonitoringService) stopGRPCServer() {
    if ms.server != nil {
        ms.server.GracefulStop()
        ms.server = nil
    }
}

// startComponents starts all monitoring components
func (ms *MonitoringService) startComponents(ctx context.Context) error {
    components := []Component{
        ms.metricsCollector,
        ms.alertManager,
        ms.traceCollector,
        ms.logAggregator,
        ms.dashboardServer,
    }

    for _, component := range components {
        if component != nil {
            if err := component.Start(ctx); err != nil {
                return fmt.Errorf("failed to start component: %w", err)
            }
        }
    }

    return nil
}

// stopComponents stops all monitoring components
func (ms *MonitoringService) stopComponents() {
    components := []Component{
        ms.dashboardServer,
        ms.logAggregator,
        ms.traceCollector,
        ms.alertManager,
        ms.metricsCollector,
    }

    for _, component := range components {
        if component != nil {
            component.Stop()
        }
    }
}

// Component interface for monitoring components
type Component interface {
    Start(ctx context.Context) error
    Stop() error
    IsHealthy() bool
    GetMetrics() map[string]interface{}
}
```

### Step 2: Metrics Collection Implementation

**Create `pkg/services/monitoring/metrics_collector.go`**:

```go
// file: pkg/services/monitoring/metrics_collector.go
// version: 2.0.0
// guid: monitor-metrics-2222-3333-4444-555555555555

package monitoring

import (
    "context"
    "fmt"
    "sync"
    "time"

    "go.uber.org/zap"

    // gcommon imports
    "github.com/jdfalk/subtitle-manager/pkg/gcommon/metrics"

    // Generated protobuf types
    monitoringpb "github.com/jdfalk/subtitle-manager/pkg/proto/monitoring/v1"
)

// MetricsCollector collects and aggregates metrics from various sources
type MetricsCollector struct {
    logger    *zap.Logger
    config    *MetricsConfig
    collector *metrics.Collector

    // Metric storage
    metrics    map[string]*monitoringpb.Metric
    mu         sync.RWMutex

    // Background workers
    workers    []Worker
    running    bool
    cancelCtx  context.CancelFunc

    // Custom metrics
    customMetrics map[string]*CustomMetric
}

// CustomMetricConfig defines configuration for custom metrics
type CustomMetricConfig struct {
    Name        string            `yaml:"name" json:"name"`
    Type        string            `yaml:"type" json:"type"`
    Description string            `yaml:"description" json:"description"`
    Labels      []string          `yaml:"labels" json:"labels"`
    Query       string            `yaml:"query" json:"query"`
    Interval    time.Duration     `yaml:"interval" json:"interval"`
    Thresholds  map[string]float64 `yaml:"thresholds" json:"thresholds"`
}

// CustomMetric represents a custom metric implementation
type CustomMetric struct {
    Config   *CustomMetricConfig
    Executor MetricExecutor
    LastValue float64
    LastRun   time.Time
}

// MetricExecutor interface for custom metric execution
type MetricExecutor interface {
    Execute(ctx context.Context) (float64, error)
}

// Worker interface for background workers
type Worker interface {
    Start(ctx context.Context) error
    Stop() error
    Name() string
}

// NewMetricsCollector creates a new metrics collector
func NewMetricsCollector(logger *zap.Logger, config *MetricsConfig) (*MetricsCollector, error) {
    if logger == nil {
        return nil, fmt.Errorf("logger is required")
    }
    if config == nil {
        return nil, fmt.Errorf("config is required")
    }

    collector := metrics.NewCollector(logger, "monitoring-service")

    mc := &MetricsCollector{
        logger:        logger,
        config:        config,
        collector:     collector,
        metrics:       make(map[string]*monitoringpb.Metric),
        customMetrics: make(map[string]*CustomMetric),
    }

    // Initialize custom metrics
    if err := mc.initializeCustomMetrics(); err != nil {
        return nil, fmt.Errorf("failed to initialize custom metrics: %w", err)
    }

    // Initialize workers
    mc.initializeWorkers()

    return mc, nil
}

// Start starts the metrics collector
func (mc *MetricsCollector) Start(ctx context.Context) error {
    mc.mu.Lock()
    defer mc.mu.Unlock()

    if mc.running {
        return fmt.Errorf("metrics collector is already running")
    }

    // Create cancellable context
    ctx, mc.cancelCtx = context.WithCancel(ctx)

    // Start workers
    for _, worker := range mc.workers {
        if err := worker.Start(ctx); err != nil {
            return fmt.Errorf("failed to start worker %s: %w", worker.Name(), err)
        }
    }

    mc.running = true
    mc.logger.Info("Metrics collector started")

    return nil
}

// Stop stops the metrics collector
func (mc *MetricsCollector) Stop() error {
    mc.mu.Lock()
    defer mc.mu.Unlock()

    if !mc.running {
        return nil
    }

    // Cancel context
    if mc.cancelCtx != nil {
        mc.cancelCtx()
    }

    // Stop workers
    for _, worker := range mc.workers {
        worker.Stop()
    }

    mc.running = false
    mc.logger.Info("Metrics collector stopped")

    return nil
}

// IsHealthy checks if the metrics collector is healthy
func (mc *MetricsCollector) IsHealthy() bool {
    mc.mu.RLock()
    defer mc.mu.RUnlock()

    return mc.running
}

// GetMetrics returns collected metrics
func (mc *MetricsCollector) GetMetrics() map[string]interface{} {
    mc.mu.RLock()
    defer mc.mu.RUnlock()

    result := make(map[string]interface{})
    result["total_metrics"] = len(mc.metrics)
    result["custom_metrics"] = len(mc.customMetrics)
    result["collection_enabled"] = mc.config.Enabled

    return result
}

// RecordMetric records a new metric value
func (mc *MetricsCollector) RecordMetric(ctx context.Context, metric *monitoringpb.Metric) error {
    mc.mu.Lock()
    defer mc.mu.Unlock()

    mc.metrics[metric.Name] = metric

    // Update collector metrics
    mc.collector.IncrementCounter("metrics_recorded_total", map[string]string{
        "metric_name": metric.Name,
        "metric_type": metric.Type,
    })

    mc.logger.Debug("Metric recorded",
        zap.String("name", metric.Name),
        zap.String("type", metric.Type),
        zap.Float64("value", metric.Value))

    return nil
}

// GetMetric retrieves a specific metric
func (mc *MetricsCollector) GetMetric(ctx context.Context, name string) (*monitoringpb.Metric, error) {
    mc.mu.RLock()
    defer mc.mu.RUnlock()

    metric, exists := mc.metrics[name]
    if !exists {
        return nil, fmt.Errorf("metric not found: %s", name)
    }

    return metric, nil
}

// ListMetrics returns all collected metrics
func (mc *MetricsCollector) ListMetrics(ctx context.Context, filters map[string]string) ([]*monitoringpb.Metric, error) {
    mc.mu.RLock()
    defer mc.mu.RUnlock()

    metrics := make([]*monitoringpb.Metric, 0)

    for _, metric := range mc.metrics {
        if mc.metricMatchesFilters(metric, filters) {
            metrics = append(metrics, metric)
        }
    }

    return metrics, nil
}

// initializeCustomMetrics initializes custom metrics based on configuration
func (mc *MetricsCollector) initializeCustomMetrics() error {
    for _, config := range mc.config.CustomMetrics {
        executor, err := mc.createMetricExecutor(config)
        if err != nil {
            return fmt.Errorf("failed to create executor for metric %s: %w", config.Name, err)
        }

        mc.customMetrics[config.Name] = &CustomMetric{
            Config:   config,
            Executor: executor,
        }
    }

    return nil
}

// createMetricExecutor creates a metric executor based on configuration
func (mc *MetricsCollector) createMetricExecutor(config *CustomMetricConfig) (MetricExecutor, error) {
    switch config.Type {
    case "query":
        return NewQueryExecutor(config.Query), nil
    case "system":
        return NewSystemExecutor(config.Name), nil
    case "service":
        return NewServiceExecutor(config.Name), nil
    default:
        return nil, fmt.Errorf("unknown metric type: %s", config.Type)
    }
}

// initializeWorkers initializes background workers
func (mc *MetricsCollector) initializeWorkers() {
    mc.workers = []Worker{
        NewMetricsCollectionWorker(mc),
        NewMetricsCleanupWorker(mc),
        NewCustomMetricsWorker(mc),
    }
}

// metricMatchesFilters checks if a metric matches the given filters
func (mc *MetricsCollector) metricMatchesFilters(metric *monitoringpb.Metric, filters map[string]string) bool {
    for key, value := range filters {
        switch key {
        case "type":
            if metric.Type != value {
                return false
            }
        case "service":
            if service, exists := metric.Labels["service"]; !exists || service != value {
                return false
            }
        case "name_prefix":
            if !strings.HasPrefix(metric.Name, value) {
                return false
            }
        }
    }
    return true
}
```

### Step 3: Alert Management Implementation

**Create `pkg/services/monitoring/alert_manager.go`**:

```go
// file: pkg/services/monitoring/alert_manager.go
// version: 2.0.0
// guid: monitor-alerts-3333-4444-5555-666666666666

package monitoring

import (
    "context"
    "fmt"
    "sync"
    "time"

    "go.uber.org/zap"

    // Generated protobuf types
    monitoringpb "github.com/jdfalk/subtitle-manager/pkg/proto/monitoring/v1"
)

// AlertManager manages alert rules and notifications
type AlertManager struct {
    logger       *zap.Logger
    config       *AlertsConfig

    // Alert state
    rules        map[string]*AlertRule
    activeAlerts map[string]*monitoringpb.Alert
    channels     map[string]AlertChannel

    // Background processing
    evaluator    *AlertEvaluator
    notifier     *AlertNotifier
    running      bool
    mu           sync.RWMutex
    cancelCtx    context.CancelFunc
}

// AlertRule defines an alert rule
type AlertRule struct {
    Name        string            `yaml:"name" json:"name"`
    Description string            `yaml:"description" json:"description"`
    Expression  string            `yaml:"expression" json:"expression"`
    Severity    string            `yaml:"severity" json:"severity"`
    Duration    time.Duration     `yaml:"duration" json:"duration"`
    Labels      map[string]string `yaml:"labels" json:"labels"`
    Annotations map[string]string `yaml:"annotations" json:"annotations"`
    Enabled     bool              `yaml:"enabled" json:"enabled"`
}

// AlertChannel defines an alert notification channel
type AlertChannel interface {
    Send(ctx context.Context, alert *monitoringpb.Alert) error
    Name() string
    IsEnabled() bool
}

// EscalationConfig defines alert escalation configuration
type EscalationConfig struct {
    Enabled    bool                     `yaml:"enabled" json:"enabled"`
    Levels     []*EscalationLevel       `yaml:"levels" json:"levels"`
    MaxRetries int                      `yaml:"max_retries" json:"max_retries"`
}

// EscalationLevel defines an escalation level
type EscalationLevel struct {
    Duration time.Duration `yaml:"duration" json:"duration"`
    Channels []string      `yaml:"channels" json:"channels"`
    Severity string        `yaml:"severity" json:"severity"`
}

// NewAlertManager creates a new alert manager
func NewAlertManager(logger *zap.Logger, config *AlertsConfig) (*AlertManager, error) {
    if logger == nil {
        return nil, fmt.Errorf("logger is required")
    }
    if config == nil {
        return nil, fmt.Errorf("config is required")
    }

    am := &AlertManager{
        logger:       logger,
        config:       config,
        rules:        make(map[string]*AlertRule),
        activeAlerts: make(map[string]*monitoringpb.Alert),
        channels:     make(map[string]AlertChannel),
    }

    // Initialize alert rules
    for _, rule := range config.Rules {
        am.rules[rule.Name] = rule
    }

    // Initialize alert channels
    if err := am.initializeChannels(); err != nil {
        return nil, fmt.Errorf("failed to initialize channels: %w", err)
    }

    // Initialize components
    am.evaluator = NewAlertEvaluator(logger, am.rules)
    am.notifier = NewAlertNotifier(logger, am.channels, config.Escalation)

    return am, nil
}

// Start starts the alert manager
func (am *AlertManager) Start(ctx context.Context) error {
    am.mu.Lock()
    defer am.mu.Unlock()

    if am.running {
        return fmt.Errorf("alert manager is already running")
    }

    // Create cancellable context
    ctx, am.cancelCtx = context.WithCancel(ctx)

    // Start evaluator
    if err := am.evaluator.Start(ctx); err != nil {
        return fmt.Errorf("failed to start evaluator: %w", err)
    }

    // Start notifier
    if err := am.notifier.Start(ctx); err != nil {
        am.evaluator.Stop()
        return fmt.Errorf("failed to start notifier: %w", err)
    }

    am.running = true
    am.logger.Info("Alert manager started")

    return nil
}

// Stop stops the alert manager
func (am *AlertManager) Stop() error {
    am.mu.Lock()
    defer am.mu.Unlock()

    if !am.running {
        return nil
    }

    // Cancel context
    if am.cancelCtx != nil {
        am.cancelCtx()
    }

    // Stop components
    am.notifier.Stop()
    am.evaluator.Stop()

    am.running = false
    am.logger.Info("Alert manager stopped")

    return nil
}

// IsHealthy checks if the alert manager is healthy
func (am *AlertManager) IsHealthy() bool {
    am.mu.RLock()
    defer am.mu.RUnlock()

    return am.running && am.evaluator.IsHealthy() && am.notifier.IsHealthy()
}

// GetMetrics returns alert manager metrics
func (am *AlertManager) GetMetrics() map[string]interface{} {
    am.mu.RLock()
    defer am.mu.RUnlock()

    return map[string]interface{}{
        "total_rules":     len(am.rules),
        "active_alerts":   len(am.activeAlerts),
        "total_channels":  len(am.channels),
        "evaluator_healthy": am.evaluator.IsHealthy(),
        "notifier_healthy":  am.notifier.IsHealthy(),
    }
}

// CreateAlert creates a new alert
func (am *AlertManager) CreateAlert(ctx context.Context, alert *monitoringpb.Alert) error {
    am.mu.Lock()
    defer am.mu.Unlock()

    am.activeAlerts[alert.Id] = alert

    // Send to notifier
    go am.notifier.ProcessAlert(ctx, alert)

    am.logger.Info("Alert created",
        zap.String("alert_id", alert.Id),
        zap.String("rule", alert.RuleName),
        zap.String("severity", alert.Severity))

    return nil
}

// ResolveAlert resolves an active alert
func (am *AlertManager) ResolveAlert(ctx context.Context, alertID string) error {
    am.mu.Lock()
    defer am.mu.Unlock()

    alert, exists := am.activeAlerts[alertID]
    if !exists {
        return fmt.Errorf("alert not found: %s", alertID)
    }

    // Update alert status
    alert.Status = "resolved"
    alert.ResolvedAt = timestamppb.Now()

    // Send resolution notification
    go am.notifier.ProcessAlert(ctx, alert)

    // Remove from active alerts
    delete(am.activeAlerts, alertID)

    am.logger.Info("Alert resolved",
        zap.String("alert_id", alertID))

    return nil
}

// ListAlerts returns active alerts
func (am *AlertManager) ListAlerts(ctx context.Context, filters map[string]string) ([]*monitoringpb.Alert, error) {
    am.mu.RLock()
    defer am.mu.RUnlock()

    alerts := make([]*monitoringpb.Alert, 0)

    for _, alert := range am.activeAlerts {
        if am.alertMatchesFilters(alert, filters) {
            alerts = append(alerts, alert)
        }
    }

    return alerts, nil
}

// initializeChannels initializes alert notification channels
func (am *AlertManager) initializeChannels() error {
    for name, channelConfig := range am.config.Channels {
        channel, err := am.createChannel(name, channelConfig)
        if err != nil {
            return fmt.Errorf("failed to create channel %s: %w", name, err)
        }
        am.channels[name] = channel
    }

    return nil
}

// createChannel creates an alert channel based on configuration
func (am *AlertManager) createChannel(name string, config *AlertChannel) (AlertChannel, error) {
    // Implementation would create different channel types
    // (email, slack, webhook, etc.) based on configuration
    return NewWebhookChannel(name, config), nil
}

// alertMatchesFilters checks if an alert matches the given filters
func (am *AlertManager) alertMatchesFilters(alert *monitoringpb.Alert, filters map[string]string) bool {
    for key, value := range filters {
        switch key {
        case "severity":
            if alert.Severity != value {
                return false
            }
        case "status":
            if alert.Status != value {
                return false
            }
        case "rule":
            if alert.RuleName != value {
                return false
            }
        }
    }
    return true
}
```

This monitoring implementation provides comprehensive observability with metrics
collection, alerting, and full gcommon integration. The system includes
customizable metrics, alert rules, multiple notification channels, and dashboard
capabilities.
