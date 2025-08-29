<!-- file: docs/tasks/05-package-replacements/TASK-05-005-metrics-migration.md -->
<!-- version: 1.0.0 -->
<!-- guid: e5f6a7b8-c9d0-1e2f-3a4b-5c6d7e8f9a0b -->

# TASK-05-005: Metrics Migration

## Overview

**Objective**: Replace current metrics collection with gcommon metrics system and update monitoring dashboards.

**Phase**: 3 (Package Replacements)
**Priority**: High
**Estimated Effort**: 6-8 hours
**Prerequisites**: TASK-05-004 (media package integration) and gcommon package foundation completed

## Required Reading

**CRITICAL**: Read these documents before starting:

- `docs/gcommon-api/metrics.md` - gcommon metrics type specifications and patterns
- Current metrics implementation in `pkg/metrics/` directory
- Monitoring dashboard configurations
- Alerting rule definitions
- `docs/MIGRATION_INVENTORY.md` - Metrics usage inventory

## Problem Statement

The subtitle-manager currently uses custom metrics collection for monitoring application performance and health. This needs to be replaced with gcommon metrics types to:

1. **Standardize Metrics**: Use gcommon MetricRegistry, Counter, Gauge, Histogram types
2. **Improve Monitoring**: Leverage gcommon's comprehensive monitoring capabilities
3. **Enhanced Dashboards**: Use standardized dashboard configurations
4. **Better Alerting**: Enable consistent alerting across gcommon-based services

### Current Metrics Implementation

```go
// Current implementation (to be replaced)
type MetricsCollector struct {
    downloadCount    int64
    processingTime   time.Duration
    errorCount       int64
    activeConnections int64
}

type ApplicationMetrics struct {
    RequestCount    map[string]int64
    ResponseTimes   map[string][]time.Duration
    ErrorRates      map[string]float64
    SystemMetrics   SystemStats
}
```

### Target gcommon Metrics Types

```go
// New implementation using gcommon
import "github.com/jdfalk/gcommon/sdks/go/v1/metrics"

metricRegistry := metrics.NewRegistry()
downloadCounter := metricRegistry.Counter("downloads_total")
processingHistogram := metricRegistry.Histogram("processing_duration_seconds")
```

## Technical Approach

### Metrics Type Mapping Strategy

1. **Registry Migration**: Convert custom metrics to gcommon MetricRegistry
2. **Type Conversion**: Map counters, gauges, histograms to gcommon types
3. **Collection Points**: Update all metric collection points
4. **Dashboard Updates**: Migrate Grafana/Prometheus configurations

### Key Mappings

```go
// Metrics type mappings
type MetricsMapping struct {
    // OLD custom types -> NEW gcommon/metrics
    MetricsCollector -> metrics.Registry + individual metrics
    Counter          -> metrics.Counter
    Gauge            -> metrics.Gauge
    Timer            -> metrics.Histogram
    SystemMetrics    -> metrics.SystemMetrics
}
```

## Implementation Steps

### Step 1: Create Metrics Registry Service

```go
// File: pkg/metrics/registry.go
package metrics

import (
    "fmt"
    "time"

    "github.com/jdfalk/gcommon/sdks/go/v1/metrics"
    "github.com/jdfalk/subtitle-manager/pkg/config"
)

// MetricsService manages application metrics using gcommon
type MetricsService struct {
    registry    *metrics.Registry
    config      *config.ApplicationConfig

    // Application-specific metrics
    downloadCounter      *metrics.Counter
    processingHistogram  *metrics.Histogram
    errorCounter         *metrics.Counter
    activeGauge          *metrics.Gauge
    systemMetrics        *metrics.SystemMetrics
}

// NewMetricsService creates a new metrics service
func NewMetricsService(config *config.ApplicationConfig) *MetricsService {
    registry := metrics.NewRegistry()

    ms := &MetricsService{
        registry: registry,
        config:   config,
    }

    // Initialize application metrics
    ms.initializeMetrics()

    return ms
}

// initializeMetrics sets up all application metrics
func (ms *MetricsService) initializeMetrics() {
    // Download metrics
    ms.downloadCounter = ms.registry.Counter("subtitle_downloads_total",
        "Total number of subtitle downloads",
        []string{"status", "language", "format"})

    // Processing metrics
    ms.processingHistogram = ms.registry.Histogram("subtitle_processing_duration_seconds",
        "Time spent processing subtitle files",
        []string{"operation", "format"},
        []float64{0.001, 0.01, 0.1, 1.0, 10.0, 60.0})

    // Error metrics
    ms.errorCounter = ms.registry.Counter("subtitle_errors_total",
        "Total number of errors by type",
        []string{"error_type", "component"})

    // System metrics
    ms.activeGauge = ms.registry.Gauge("subtitle_active_operations",
        "Number of currently active operations",
        []string{"operation_type"})

    // System metrics using gcommon
    ms.systemMetrics = ms.registry.SystemMetrics("subtitle_manager")
}

// GetRegistry returns the metrics registry
func (ms *MetricsService) GetRegistry() *metrics.Registry {
    return ms.registry
}

// RecordDownload records a subtitle download
func (ms *MetricsService) RecordDownload(status, language, format string) {
    ms.downloadCounter.With(map[string]string{
        "status":   status,
        "language": language,
        "format":   format,
    }).Inc()
}

// RecordProcessingTime records subtitle processing duration
func (ms *MetricsService) RecordProcessingTime(operation, format string, duration time.Duration) {
    ms.processingHistogram.With(map[string]string{
        "operation": operation,
        "format":    format,
    }).Observe(duration.Seconds())
}

// RecordError records an application error
func (ms *MetricsService) RecordError(errorType, component string) {
    ms.errorCounter.With(map[string]string{
        "error_type": errorType,
        "component":  component,
    }).Inc()
}

// SetActiveOperations sets the number of active operations
func (ms *MetricsService) SetActiveOperations(operationType string, count int) {
    ms.activeGauge.With(map[string]string{
        "operation_type": operationType,
    }).Set(float64(count))
}

// IncrementActiveOperations increments active operations
func (ms *MetricsService) IncrementActiveOperations(operationType string) {
    ms.activeGauge.With(map[string]string{
        "operation_type": operationType,
    }).Inc()
}

// DecrementActiveOperations decrements active operations
func (ms *MetricsService) DecrementActiveOperations(operationType string) {
    ms.activeGauge.With(map[string]string{
        "operation_type": operationType,
    }).Dec()
}

// UpdateSystemMetrics updates system resource metrics
func (ms *MetricsService) UpdateSystemMetrics() error {
    return ms.systemMetrics.Update()
}

// GetMetricValue gets current value of a metric (for testing/debugging)
func (ms *MetricsService) GetMetricValue(metricName string, labels map[string]string) interface{} {
    metric := ms.registry.GetMetric(metricName)
    if metric == nil {
        return nil
    }

    switch m := metric.(type) {
    case *metrics.Counter:
        return m.With(labels).Value()
    case *metrics.Gauge:
        return m.With(labels).Value()
    case *metrics.Histogram:
        return m.With(labels).Snapshot()
    default:
        return nil
    }
}

// CollectCustomMetrics collects application-specific metrics
func (ms *MetricsService) CollectCustomMetrics() (*CustomMetrics, error) {
    custom := &CustomMetrics{
        Timestamp: time.Now(),
        Metrics:   make(map[string]interface{}),
    }

    // Collect download statistics
    downloadStats := ms.collectDownloadStats()
    custom.Metrics["downloads"] = downloadStats

    // Collect processing statistics
    processingStats := ms.collectProcessingStats()
    custom.Metrics["processing"] = processingStats

    // Collect error statistics
    errorStats := ms.collectErrorStats()
    custom.Metrics["errors"] = errorStats

    // Collect system statistics
    systemStats, err := ms.collectSystemStats()
    if err != nil {
        return nil, fmt.Errorf("failed to collect system stats: %v", err)
    }
    custom.Metrics["system"] = systemStats

    return custom, nil
}

// CustomMetrics holds collected metrics data
type CustomMetrics struct {
    Timestamp time.Time              `json:"timestamp"`
    Metrics   map[string]interface{} `json:"metrics"`
}

// collectDownloadStats collects download-related metrics
func (ms *MetricsService) collectDownloadStats() map[string]interface{} {
    stats := make(map[string]interface{})

    // Get download counts by status
    stats["total"] = ms.GetMetricValue("subtitle_downloads_total", map[string]string{})
    stats["successful"] = ms.GetMetricValue("subtitle_downloads_total",
        map[string]string{"status": "success"})
    stats["failed"] = ms.GetMetricValue("subtitle_downloads_total",
        map[string]string{"status": "error"})

    // Get download counts by language
    languages := []string{"en", "es", "fr", "de", "it"}
    languageStats := make(map[string]interface{})
    for _, lang := range languages {
        languageStats[lang] = ms.GetMetricValue("subtitle_downloads_total",
            map[string]string{"language": lang})
    }
    stats["by_language"] = languageStats

    return stats
}

// collectProcessingStats collects processing-related metrics
func (ms *MetricsService) collectProcessingStats() map[string]interface{} {
    stats := make(map[string]interface{})

    // Get processing time statistics
    operations := []string{"download", "convert", "validate", "store"}
    for _, operation := range operations {
        histogram := ms.GetMetricValue("subtitle_processing_duration_seconds",
            map[string]string{"operation": operation})
        if snapshot, ok := histogram.(*metrics.HistogramSnapshot); ok {
            stats[operation] = map[string]interface{}{
                "count":      snapshot.Count(),
                "mean":       snapshot.Mean(),
                "median":     snapshot.Percentile(0.5),
                "p95":        snapshot.Percentile(0.95),
                "p99":        snapshot.Percentile(0.99),
            }
        }
    }

    return stats
}

// collectErrorStats collects error-related metrics
func (ms *MetricsService) collectErrorStats() map[string]interface{} {
    stats := make(map[string]interface{})

    // Get error counts by type
    errorTypes := []string{"download_failed", "conversion_failed", "validation_failed", "storage_failed"}
    for _, errorType := range errorTypes {
        stats[errorType] = ms.GetMetricValue("subtitle_errors_total",
            map[string]string{"error_type": errorType})
    }

    // Get error counts by component
    components := []string{"downloader", "converter", "validator", "storage"}
    componentStats := make(map[string]interface{})
    for _, component := range components {
        componentStats[component] = ms.GetMetricValue("subtitle_errors_total",
            map[string]string{"component": component})
    }
    stats["by_component"] = componentStats

    return stats
}

// collectSystemStats collects system resource metrics
func (ms *MetricsService) collectSystemStats() (map[string]interface{}, error) {
    // Update system metrics first
    if err := ms.UpdateSystemMetrics(); err != nil {
        return nil, err
    }

    // Collect current system stats from gcommon
    systemStats := ms.systemMetrics.GetCurrentStats()

    stats := map[string]interface{}{
        "cpu_usage":     systemStats.CPUUsage,
        "memory_usage":  systemStats.MemoryUsage,
        "disk_usage":    systemStats.DiskUsage,
        "network_io":    systemStats.NetworkIO,
        "file_handles":  systemStats.FileHandles,
        "goroutines":    systemStats.Goroutines,
    }

    return stats, nil
}
```

### Step 2: Create Metrics Collection Middleware

```go
// File: pkg/metrics/middleware.go
package metrics

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
)

// HTTPMetricsMiddleware creates middleware for HTTP request metrics
func (ms *MetricsService) HTTPMetricsMiddleware() gin.HandlerFunc {
    // Create HTTP-specific metrics
    requestCounter := ms.registry.Counter("http_requests_total",
        "Total number of HTTP requests",
        []string{"method", "path", "status"})

    requestDuration := ms.registry.Histogram("http_request_duration_seconds",
        "HTTP request duration in seconds",
        []string{"method", "path"},
        []float64{0.001, 0.01, 0.1, 1.0, 10.0})

    requestSize := ms.registry.Histogram("http_request_size_bytes",
        "HTTP request size in bytes",
        []string{"method", "path"},
        []float64{100, 1000, 10000, 100000, 1000000})

    responseSize := ms.registry.Histogram("http_response_size_bytes",
        "HTTP response size in bytes",
        []string{"method", "path"},
        []float64{100, 1000, 10000, 100000, 1000000})

    return func(c *gin.Context) {
        start := time.Now()

        // Process request
        c.Next()

        // Record metrics
        duration := time.Since(start)
        status := strconv.Itoa(c.Writer.Status())

        labels := map[string]string{
            "method": c.Request.Method,
            "path":   c.FullPath(),
        }

        labelsWithStatus := map[string]string{
            "method": c.Request.Method,
            "path":   c.FullPath(),
            "status": status,
        }

        requestCounter.With(labelsWithStatus).Inc()
        requestDuration.With(labels).Observe(duration.Seconds())

        if c.Request.ContentLength > 0 {
            requestSize.With(labels).Observe(float64(c.Request.ContentLength))
        }

        responseSize.With(labels).Observe(float64(c.Writer.Size()))
    }
}

// DatabaseMetricsMiddleware creates middleware for database operation metrics
func (ms *MetricsService) DatabaseMetricsMiddleware() DatabaseMiddleware {
    dbOperationCounter := ms.registry.Counter("database_operations_total",
        "Total number of database operations",
        []string{"operation", "table", "status"})

    dbOperationDuration := ms.registry.Histogram("database_operation_duration_seconds",
        "Database operation duration in seconds",
        []string{"operation", "table"},
        []float64{0.001, 0.01, 0.1, 1.0, 10.0})

    return DatabaseMiddleware{
        operationCounter:  dbOperationCounter,
        operationDuration: dbOperationDuration,
    }
}

// DatabaseMiddleware wraps database operations with metrics
type DatabaseMiddleware struct {
    operationCounter  *metrics.Counter
    operationDuration *metrics.Histogram
}

// WrapDatabaseOperation wraps a database operation with metrics collection
func (dm *DatabaseMiddleware) WrapDatabaseOperation(operation, table string, fn func() error) error {
    start := time.Now()

    err := fn()

    duration := time.Since(start)
    status := "success"
    if err != nil {
        status = "error"
    }

    labels := map[string]string{
        "operation": operation,
        "table":     table,
    }

    labelsWithStatus := map[string]string{
        "operation": operation,
        "table":     table,
        "status":    status,
    }

    dm.operationCounter.With(labelsWithStatus).Inc()
    dm.operationDuration.With(labels).Observe(duration.Seconds())

    return err
}

// BackgroundTaskMetrics creates metrics for background tasks
func (ms *MetricsService) BackgroundTaskMetrics() *BackgroundMetrics {
    taskCounter := ms.registry.Counter("background_tasks_total",
        "Total number of background tasks",
        []string{"task_type", "status"})

    taskDuration := ms.registry.Histogram("background_task_duration_seconds",
        "Background task duration in seconds",
        []string{"task_type"},
        []float64{1.0, 10.0, 60.0, 300.0, 1800.0, 3600.0})

    taskQueue := ms.registry.Gauge("background_task_queue_size",
        "Current size of background task queue",
        []string{"task_type"})

    return &BackgroundMetrics{
        taskCounter:  taskCounter,
        taskDuration: taskDuration,
        taskQueue:    taskQueue,
    }
}

// BackgroundMetrics holds background task metrics
type BackgroundMetrics struct {
    taskCounter  *metrics.Counter
    taskDuration *metrics.Histogram
    taskQueue    *metrics.Gauge
}

// RecordTaskStart records the start of a background task
func (bm *BackgroundMetrics) RecordTaskStart(taskType string) {
    bm.taskQueue.With(map[string]string{
        "task_type": taskType,
    }).Inc()
}

// RecordTaskComplete records the completion of a background task
func (bm *BackgroundMetrics) RecordTaskComplete(taskType string, duration time.Duration, success bool) {
    status := "success"
    if !success {
        status = "error"
    }

    bm.taskCounter.With(map[string]string{
        "task_type": taskType,
        "status":    status,
    }).Inc()

    bm.taskDuration.With(map[string]string{
        "task_type": taskType,
    }).Observe(duration.Seconds())

    bm.taskQueue.With(map[string]string{
        "task_type": taskType,
    }).Dec()
}
```

### Step 3: Create Dashboard Configuration

```yaml
# File: monitoring/dashboards/subtitle-manager.yml
apiVersion: v1
kind: ConfigMap
metadata:
  name: subtitle-manager-dashboard
  labels:
    grafana_dashboard: "1"
data:
  subtitle-manager.json: |
    {
      "dashboard": {
        "id": null,
        "title": "Subtitle Manager Metrics",
        "tags": ["gcommon", "subtitle-manager"],
        "style": "dark",
        "timezone": "browser",
        "panels": [
          {
            "id": 1,
            "title": "Download Statistics",
            "type": "stat",
            "targets": [
              {
                "expr": "rate(subtitle_downloads_total[5m])",
                "legendFormat": "Downloads/sec"
              }
            ],
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "displayMode": "list",
                  "orientation": "horizontal"
                },
                "mappings": [],
                "thresholds": {
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 10
                    }
                  ]
                }
              }
            },
            "gridPos": {
              "h": 8,
              "w": 12,
              "x": 0,
              "y": 0
            }
          },
          {
            "id": 2,
            "title": "Processing Time",
            "type": "graph",
            "targets": [
              {
                "expr": "histogram_quantile(0.95, rate(subtitle_processing_duration_seconds_bucket[5m]))",
                "legendFormat": "95th percentile"
              },
              {
                "expr": "histogram_quantile(0.50, rate(subtitle_processing_duration_seconds_bucket[5m]))",
                "legendFormat": "50th percentile"
              }
            ],
            "yAxes": [
              {
                "label": "Duration (seconds)",
                "min": 0
              }
            ],
            "gridPos": {
              "h": 8,
              "w": 12,
              "x": 12,
              "y": 0
            }
          },
          {
            "id": 3,
            "title": "Error Rate",
            "type": "graph",
            "targets": [
              {
                "expr": "rate(subtitle_errors_total[5m])",
                "legendFormat": "{{error_type}}"
              }
            ],
            "yAxes": [
              {
                "label": "Errors/sec",
                "min": 0
              }
            ],
            "gridPos": {
              "h": 8,
              "w": 12,
              "x": 0,
              "y": 8
            }
          },
          {
            "id": 4,
            "title": "System Resources",
            "type": "graph",
            "targets": [
              {
                "expr": "subtitle_manager_cpu_usage",
                "legendFormat": "CPU Usage %"
              },
              {
                "expr": "subtitle_manager_memory_usage",
                "legendFormat": "Memory Usage %"
              }
            ],
            "yAxes": [
              {
                "label": "Percentage",
                "min": 0,
                "max": 100
              }
            ],
            "gridPos": {
              "h": 8,
              "w": 12,
              "x": 12,
              "y": 8
            }
          }
        ],
        "time": {
          "from": "now-1h",
          "to": "now"
        },
        "refresh": "30s"
      }
    }
```

### Step 4: Create Prometheus Alerting Rules

```yaml
# File: monitoring/alerts/subtitle-manager.yml
groups:
  - name: subtitle-manager
    rules:
      - alert: HighDownloadErrorRate
        expr: rate(subtitle_errors_total{error_type="download_failed"}[5m]) > 0.1
        for: 2m
        labels:
          severity: warning
          service: subtitle-manager
        annotations:
          summary: "High download error rate detected"
          description: "Download error rate is {{ $value }} errors per second"

      - alert: SlowProcessingTime
        expr: histogram_quantile(0.95, rate(subtitle_processing_duration_seconds_bucket[5m])) > 30
        for: 5m
        labels:
          severity: warning
          service: subtitle-manager
        annotations:
          summary: "Slow subtitle processing detected"
          description: "95th percentile processing time is {{ $value }} seconds"

      - alert: HighMemoryUsage
        expr: subtitle_manager_memory_usage > 90
        for: 5m
        labels:
          severity: critical
          service: subtitle-manager
        annotations:
          summary: "High memory usage detected"
          description: "Memory usage is {{ $value }}%"

      - alert: ServiceDown
        expr: up{job="subtitle-manager"} == 0
        for: 1m
        labels:
          severity: critical
          service: subtitle-manager
        annotations:
          summary: "Subtitle Manager service is down"
          description: "The subtitle-manager service has been down for more than 1 minute"

      - alert: HighActiveOperations
        expr: subtitle_active_operations > 100
        for: 5m
        labels:
          severity: warning
          service: subtitle-manager
        annotations:
          summary: "High number of active operations"
          description: "{{ $value }} active operations detected"
```

### Step 5: Create Metrics HTTP Endpoint

```go
// File: pkg/api/metrics.go
package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jdfalk/subtitle-manager/pkg/metrics"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsHandler provides metrics-related HTTP endpoints
type MetricsHandler struct {
    metricsService *metrics.MetricsService
}

// NewMetricsHandler creates a new metrics handler
func NewMetricsHandler(metricsService *metrics.MetricsService) *MetricsHandler {
    return &MetricsHandler{
        metricsService: metricsService,
    }
}

// RegisterRoutes registers metrics routes
func (mh *MetricsHandler) RegisterRoutes(router *gin.Engine) {
    metricsGroup := router.Group("/metrics")
    {
        // Prometheus metrics endpoint
        metricsGroup.GET("/prometheus", mh.PrometheusMetrics)

        // Custom metrics endpoint
        metricsGroup.GET("/custom", mh.CustomMetrics)

        // Health metrics endpoint
        metricsGroup.GET("/health", mh.HealthMetrics)
    }
}

// PrometheusMetrics serves Prometheus-format metrics
func (mh *MetricsHandler) PrometheusMetrics(c *gin.Context) {
    // Use gcommon metrics registry for Prometheus export
    handler := promhttp.HandlerFor(
        mh.metricsService.GetRegistry().GetPrometheusRegistry(),
        promhttp.HandlerOpts{})

    handler.ServeHTTP(c.Writer, c.Request)
}

// CustomMetrics returns application-specific metrics in JSON format
func (mh *MetricsHandler) CustomMetrics(c *gin.Context) {
    customMetrics, err := mh.metricsService.CollectCustomMetrics()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to collect metrics: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, customMetrics)
}

// HealthMetrics returns health-related metrics
func (mh *MetricsHandler) HealthMetrics(c *gin.Context) {
    // Update system metrics
    if err := mh.metricsService.UpdateSystemMetrics(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to update system metrics: " + err.Error(),
        })
        return
    }

    healthMetrics := map[string]interface{}{
        "timestamp": time.Now(),
        "status":    "healthy",
        "metrics": map[string]interface{}{
            "active_operations": mh.metricsService.GetMetricValue("subtitle_active_operations", nil),
            "total_downloads":   mh.metricsService.GetMetricValue("subtitle_downloads_total", nil),
            "total_errors":      mh.metricsService.GetMetricValue("subtitle_errors_total", nil),
        },
    }

    c.JSON(http.StatusOK, healthMetrics)
}
```

## Testing Requirements

### Metrics Service Tests

```go
// File: pkg/metrics/service_test.go
package metrics

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestMetricsService(t *testing.T) {
    mockConfig := &MockConfig{}
    metricsService := NewMetricsService(mockConfig)

    t.Run("Download metrics", func(t *testing.T) {
        // Record download
        metricsService.RecordDownload("success", "en", "srt")

        // Verify counter incremented
        value := metricsService.GetMetricValue("subtitle_downloads_total",
            map[string]string{
                "status":   "success",
                "language": "en",
                "format":   "srt",
            })

        assert.Equal(t, float64(1), value)
    })

    t.Run("Processing time metrics", func(t *testing.T) {
        duration := 5 * time.Second
        metricsService.RecordProcessingTime("download", "srt", duration)

        // Verify histogram recorded
        value := metricsService.GetMetricValue("subtitle_processing_duration_seconds",
            map[string]string{
                "operation": "download",
                "format":    "srt",
            })

        assert.NotNil(t, value)
    })

    t.Run("Error metrics", func(t *testing.T) {
        metricsService.RecordError("download_failed", "downloader")

        value := metricsService.GetMetricValue("subtitle_errors_total",
            map[string]string{
                "error_type": "download_failed",
                "component":  "downloader",
            })

        assert.Equal(t, float64(1), value)
    })
}

func TestCustomMetrics(t *testing.T) {
    metricsService := NewMetricsService(&MockConfig{})

    // Record some test data
    metricsService.RecordDownload("success", "en", "srt")
    metricsService.RecordError("download_failed", "downloader")

    custom, err := metricsService.CollectCustomMetrics()
    require.NoError(t, err)

    assert.NotNil(t, custom.Metrics["downloads"])
    assert.NotNil(t, custom.Metrics["errors"])
    assert.NotEmpty(t, custom.Timestamp)
}
```

## Success Metrics

### Functional Requirements

- [ ] All metrics use gcommon metrics types
- [ ] Dashboard displays correct data from gcommon metrics
- [ ] Alerting rules work with new metric names
- [ ] HTTP endpoints serve both Prometheus and JSON metrics
- [ ] System metrics integration functional

### Technical Requirements

- [ ] Performance impact minimal (< 5% overhead)
- [ ] Metric accuracy maintained during migration
- [ ] Dashboard refresh rates acceptable
- [ ] Alert notification delivery working
- [ ] Custom metrics collection comprehensive

### Integration Requirements

- [ ] Prometheus scraping configuration updated
- [ ] Grafana dashboard import successful
- [ ] Alert manager rule deployment working
- [ ] Log correlation with metrics maintained

## Common Pitfalls

1. **Metric Name Changes**: Ensure dashboard queries updated for new metric names
2. **Label Consistency**: Maintain consistent labeling across all metrics
3. **Performance Impact**: Monitor overhead of metrics collection
4. **Alert Threshold**: Verify alert thresholds still appropriate with new metrics
5. **Dashboard Compatibility**: Test dashboard functionality with new metric structure

## Dependencies

- **Requires**: TASK-05-004 (media package integration) for media-related metrics
- **Requires**: gcommon metrics SDK properly installed
- **Requires**: Prometheus and Grafana configuration access
- **Enables**: Comprehensive monitoring with gcommon ecosystem
- **Blocks**: Production monitoring until migration complete

This comprehensive metrics migration replaces custom metrics collection with gcommon standardized metrics while maintaining all monitoring capabilities and improving dashboard consistency across the gcommon ecosystem.
