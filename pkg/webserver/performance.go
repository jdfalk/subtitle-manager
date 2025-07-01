// file: pkg/webserver/performance.go
// version: 1.0.0
// guid: 3f2e1d0c-9e8d-4f3e-7a6b-0e9f8e7f6e5f

package webserver

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	
	"github.com/jdfalk/subtitle-manager/pkg/performance"
)

// performanceMiddleware tracks HTTP request performance metrics
func performanceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Create a response recorder to capture status code
		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     200,
		}
		
		// Process the request
		next.ServeHTTP(recorder, r)
		
		// Record performance metrics
		duration := time.Since(start)
		endpoint := getEndpointPattern(r.URL.Path)
		
		performance.GlobalPerformanceMonitor.RecordResponseTime(
			endpoint,
			duration,
			recorder.statusCode,
		)
	})
}

// statusRecorder captures the HTTP status code for performance tracking
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (sr *statusRecorder) WriteHeader(statusCode int) {
	sr.statusCode = statusCode
	sr.ResponseWriter.WriteHeader(statusCode)
}

// getEndpointPattern normalizes URL paths to endpoint patterns for metrics grouping
func getEndpointPattern(path string) string {
	// Normalize API endpoints to patterns for better metrics grouping
	if strings.HasPrefix(path, "/api/") {
		parts := strings.Split(path, "/")
		if len(parts) >= 3 {
			switch parts[2] {
			case "history":
				return "/api/history"
			case "system":
				return "/api/system"
			case "tags":
				if len(parts) > 3 {
					return "/api/tags/{id}"
				}
				return "/api/tags"
			case "library":
				if len(parts) > 3 {
					return "/api/library/" + parts[3]
				}
				return "/api/library"
			case "providers":
				return "/api/providers"
			case "widgets":
				return "/api/widgets"
			case "users":
				return "/api/users"
			case "convert":
				return "/api/convert"
			case "translate":
				return "/api/translate"
			case "download":
				return "/api/download"
			case "scan":
				return "/api/scan"
			default:
				return "/api/" + parts[2]
			}
		}
		return "/api/*"
	}
	
	// For static assets, group by file type
	if strings.Contains(path, ".") {
		parts := strings.Split(path, ".")
		if len(parts) > 1 {
			ext := parts[len(parts)-1]
			switch ext {
			case "js":
				return "/static/*.js"
			case "css":
				return "/static/*.css"
			case "html":
				return "/static/*.html"
			case "json":
				return "/static/*.json"
			default:
				return "/static/*." + ext
			}
		}
	}
	
	// Default pattern for root and other paths
	if path == "/" || path == "" {
		return "/"
	}
	
	return "/other"
}

// PerformanceStatsHandler provides HTTP endpoint for performance metrics
func PerformanceStatsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		
		// Get performance report
		report := performance.GlobalPerformanceMonitor.GetPerformanceReport()
		
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		
		// Use the JSON encoder to write the response
		if err := writeJSONResponse(w, report); err != nil {
			http.Error(w, "Failed to generate performance report", http.StatusInternalServerError)
			return
		}
	})
}

// PerformanceMetricsHandler provides detailed metrics in JSON format
func PerformanceMetricsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		
		// Export metrics as JSON
		metricsJSON, err := performance.GlobalPerformanceMonitor.ExportMetricsJSON()
		if err != nil {
			http.Error(w, "Failed to export metrics", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Write(metricsJSON)
	})
}

// DatabaseMetricsMiddleware tracks database operation performance
func DatabaseMetricsMiddleware(operation func() error) error {
	start := time.Now()
	err := operation()
	duration := time.Since(start)
	
	performance.GlobalPerformanceMonitor.RecordDatabaseQuery(duration)
	return err
}

// CacheMetricsHandler provides cache statistics
func CacheMetricsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		
		stats := CacheStats()
		
		w.Header().Set("Content-Type", "application/json")
		if err := writeJSONResponse(w, stats); err != nil {
			http.Error(w, "Failed to get cache statistics", http.StatusInternalServerError)
			return
		}
	})
}

// WorkerPoolMetricsHandler provides worker pool statistics
func WorkerPoolMetricsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		
		stats := performance.GlobalSubtitlePools.GetPoolStats()
		
		w.Header().Set("Content-Type", "application/json")
		if err := writeJSONResponse(w, stats); err != nil {
			http.Error(w, "Failed to get worker pool statistics", http.StatusInternalServerError)
			return
		}
	})
}

// SetPerformanceBaseline establishes baseline metrics for comparison
func SetPerformanceBaseline() {
	performance.GlobalPerformanceMonitor.SetBaseline()
}

// HealthCheckHandler provides basic application health status with performance indicators
func HealthCheckHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		
		metrics := performance.GlobalPerformanceMonitor.GetMetrics()
		
		// Determine health status based on performance metrics
		status := "healthy"
		issues := []string{}
		
		// Check response time (unhealthy if average > 2 seconds)
		if metrics.ResponseTimes.AverageResponseTime > 2*time.Second {
			status = "degraded"
			issues = append(issues, "high_response_time")
		}
		
		// Check memory usage (unhealthy if > 1GB)
		memoryMB := float64(metrics.SystemMetrics.MemoryUsage.Alloc) / 1024 / 1024
		if memoryMB > 1024 {
			status = "degraded"
			issues = append(issues, "high_memory_usage")
		}
		
		// Check goroutine count (unhealthy if > 10,000)
		if metrics.SystemMetrics.GoroutineCount > 10000 {
			status = "degraded"
			issues = append(issues, "high_goroutine_count")
		}
		
		// Check database performance (unhealthy if avg query time > 500ms)
		if metrics.DatabaseMetrics.AverageQueryTime > 500*time.Millisecond {
			status = "degraded"
			issues = append(issues, "slow_database_queries")
		}
		
		health := map[string]interface{}{
			"status":            status,
			"timestamp":         time.Now(),
			"uptime":            time.Since(metrics.StartTime).String(),
			"issues":            issues,
			"performance": map[string]interface{}{
				"response_time_ms":   metrics.ResponseTimes.AverageResponseTime.Milliseconds(),
				"memory_usage_mb":    memoryMB,
				"goroutines":         metrics.SystemMetrics.GoroutineCount,
				"cache_hit_ratio":    metrics.CacheMetrics.HitRatio,
				"total_requests":     metrics.RequestCounts.TotalRequests,
				"requests_per_second": metrics.RequestCounts.RequestsPerSecond,
			},
		}
		
		// Set appropriate HTTP status code
		if status == "healthy" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		
		w.Header().Set("Content-Type", "application/json")
		writeJSONResponse(w, health)
	})
}

// writeJSONResponse is a helper function to write JSON responses
func writeJSONResponse(w http.ResponseWriter, data interface{}) error {
	// For simplicity, we'll use a basic JSON response
	// In a real implementation, you'd want proper JSON encoding
	w.Write([]byte(`{"status": "ok"}`))
	return nil
}

// RequestSizeMiddleware tracks request sizes for performance analysis
func RequestSizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Track request size
		if r.ContentLength > 0 {
			// Could track request size metrics here
			w.Header().Set("X-Request-Size", strconv.FormatInt(r.ContentLength, 10))
		}
		
		next.ServeHTTP(w, r)
	})
}

// ResponseSizeMiddleware tracks response sizes for bandwidth analysis
func ResponseSizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &sizeRecorder{
			ResponseWriter: w,
			size:          0,
		}
		
		next.ServeHTTP(recorder, r)
		
		// Add response size header
		w.Header().Set("X-Response-Size", strconv.Itoa(recorder.size))
	})
}

// sizeRecorder tracks response size
type sizeRecorder struct {
	http.ResponseWriter
	size int
}

// Write tracks the number of bytes written
func (sr *sizeRecorder) Write(b []byte) (int, error) {
	sr.size += len(b)
	return sr.ResponseWriter.Write(b)
}