// file: pkg/performance/monitor.go
// version: 1.0.0
// guid: 2f1e0d9c-8d7c-3f2e-6a5b-9d8e7f6e5f4e

package performance

import (
	"encoding/json"
	"runtime"
	"sync"
	"time"
)

// PerformanceMetrics tracks comprehensive application performance data
type PerformanceMetrics struct {
	// HTTP Response metrics
	ResponseTimes  ResponseTimeMetrics `json:"response_times"`
	RequestCounts  RequestCountMetrics `json:"request_counts"`
	
	// Database metrics
	DatabaseMetrics DatabasePerformanceMetrics `json:"database"`
	
	// Memory and CPU metrics
	SystemMetrics SystemResourceMetrics `json:"system"`
	
	// Cache metrics
	CacheMetrics CachePerformanceMetrics `json:"cache"`
	
	// Worker pool metrics
	WorkerMetrics map[string]PoolMetrics `json:"workers"`
	
	// Application startup and runtime
	StartTime    time.Time `json:"start_time"`
	LastUpdated  time.Time `json:"last_updated"`
}

// ResponseTimeMetrics tracks HTTP response performance
type ResponseTimeMetrics struct {
	AverageResponseTime time.Duration            `json:"average_response_time"`
	P50ResponseTime     time.Duration            `json:"p50_response_time"`
	P95ResponseTime     time.Duration            `json:"p95_response_time"`
	P99ResponseTime     time.Duration            `json:"p99_response_time"`
	EndpointTimes       map[string]time.Duration `json:"endpoint_times"`
	TotalRequests       int64                    `json:"total_requests"`
	mu                  sync.RWMutex
	responseTimes       []time.Duration
}

// RequestCountMetrics tracks request volume and patterns
type RequestCountMetrics struct {
	TotalRequests     int64            `json:"total_requests"`
	RequestsPerSecond float64          `json:"requests_per_second"`
	EndpointCounts    map[string]int64 `json:"endpoint_counts"`
	StatusCodes       map[int]int64    `json:"status_codes"`
	mu                sync.RWMutex
}

// DatabasePerformanceMetrics tracks database operation performance
type DatabasePerformanceMetrics struct {
	QueryCount        int64             `json:"query_count"`
	AverageQueryTime  time.Duration     `json:"average_query_time"`
	SlowQueries       int64             `json:"slow_queries"`
	ConnectionPoolUse map[string]int    `json:"connection_pool_use"`
	IndexHitRatio     float64           `json:"index_hit_ratio"`
	mu                sync.RWMutex
}

// SystemResourceMetrics tracks system resource usage
type SystemResourceMetrics struct {
	MemoryUsage    MemoryUsage    `json:"memory"`
	CPUUsage       CPUUsage       `json:"cpu"`
	GoroutineCount int            `json:"goroutine_count"`
	LastGC         time.Time      `json:"last_gc"`
	GCPause        time.Duration  `json:"gc_pause"`
}

// MemoryUsage tracks memory consumption
type MemoryUsage struct {
	Alloc        uint64  `json:"alloc"`         // Currently allocated memory
	TotalAlloc   uint64  `json:"total_alloc"`   // Total allocated memory
	Sys          uint64  `json:"sys"`           // Memory from system
	NumGC        uint32  `json:"num_gc"`        // Number of GC runs
	HeapAlloc    uint64  `json:"heap_alloc"`    // Heap allocated memory
	HeapSys      uint64  `json:"heap_sys"`      // Heap system memory
	StackInuse   uint64  `json:"stack_inuse"`   // Stack memory in use
	ReductionPct float64 `json:"reduction_pct"` // Percentage reduction from baseline
}

// CPUUsage tracks CPU utilization
type CPUUsage struct {
	NumCPU      int     `json:"num_cpu"`
	GOMAXPROCS  int     `json:"gomaxprocs"`
	NumCgoCall  int64   `json:"num_cgo_call"`
	UsagePct    float64 `json:"usage_pct"`
}

// CachePerformanceMetrics tracks cache effectiveness
type CachePerformanceMetrics struct {
	HitCount    int64   `json:"hit_count"`
	MissCount   int64   `json:"miss_count"`
	HitRatio    float64 `json:"hit_ratio"`
	Size        int     `json:"size"`
	Evictions   int64   `json:"evictions"`
	LastCleanup time.Time `json:"last_cleanup"`
}

// PerformanceMonitor provides centralized performance tracking
type PerformanceMonitor struct {
	metrics      *PerformanceMetrics
	mu           sync.RWMutex
	baseline     *PerformanceMetrics // Baseline metrics for comparison
	startTime    time.Time
}

// NewPerformanceMonitor creates a new performance monitoring instance
func NewPerformanceMonitor() *PerformanceMonitor {
	now := time.Now()
	
	monitor := &PerformanceMonitor{
		metrics: &PerformanceMetrics{
			ResponseTimes: ResponseTimeMetrics{
				EndpointTimes:  make(map[string]time.Duration),
				responseTimes:  make([]time.Duration, 0, 1000), // Pre-allocate for 1000 samples
			},
			RequestCounts: RequestCountMetrics{
				EndpointCounts: make(map[string]int64),
				StatusCodes:    make(map[int]int64),
			},
			DatabaseMetrics: DatabasePerformanceMetrics{
				ConnectionPoolUse: make(map[string]int),
			},
			WorkerMetrics: make(map[string]PoolMetrics),
			StartTime:     now,
			LastUpdated:   now,
		},
		startTime: now,
	}
	
	// Start background monitoring
	go monitor.backgroundMonitoring()
	
	return monitor
}

// RecordResponseTime records an HTTP response time for performance tracking
func (pm *PerformanceMonitor) RecordResponseTime(endpoint string, duration time.Duration, statusCode int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	// Update response time metrics
	pm.metrics.ResponseTimes.mu.Lock()
	pm.metrics.ResponseTimes.responseTimes = append(pm.metrics.ResponseTimes.responseTimes, duration)
	pm.metrics.ResponseTimes.EndpointTimes[endpoint] = duration
	pm.metrics.ResponseTimes.TotalRequests++
	
	// Keep only recent response times (last 1000)
	if len(pm.metrics.ResponseTimes.responseTimes) > 1000 {
		pm.metrics.ResponseTimes.responseTimes = pm.metrics.ResponseTimes.responseTimes[100:]
	}
	
	pm.calculateResponseTimePercentiles()
	pm.metrics.ResponseTimes.mu.Unlock()
	
	// Update request count metrics
	pm.metrics.RequestCounts.mu.Lock()
	pm.metrics.RequestCounts.TotalRequests++
	pm.metrics.RequestCounts.EndpointCounts[endpoint]++
	pm.metrics.RequestCounts.StatusCodes[statusCode]++
	pm.calculateRequestsPerSecond()
	pm.metrics.RequestCounts.mu.Unlock()
	
	pm.metrics.LastUpdated = time.Now()
}

// RecordDatabaseQuery records database query performance
func (pm *PerformanceMonitor) RecordDatabaseQuery(duration time.Duration) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	pm.metrics.DatabaseMetrics.mu.Lock()
	pm.metrics.DatabaseMetrics.QueryCount++
	
	// Update average query time
	if pm.metrics.DatabaseMetrics.QueryCount == 1 {
		pm.metrics.DatabaseMetrics.AverageQueryTime = duration
	} else {
		// Calculate rolling average
		current := pm.metrics.DatabaseMetrics.AverageQueryTime
		count := pm.metrics.DatabaseMetrics.QueryCount
		pm.metrics.DatabaseMetrics.AverageQueryTime = (current*time.Duration(count-1) + duration) / time.Duration(count)
	}
	
	// Track slow queries (> 100ms)
	if duration > 100*time.Millisecond {
		pm.metrics.DatabaseMetrics.SlowQueries++
	}
	
	pm.metrics.DatabaseMetrics.mu.Unlock()
}

// UpdateSystemMetrics refreshes system resource usage metrics
func (pm *PerformanceMonitor) UpdateSystemMetrics() {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	// Calculate memory reduction percentage if we have a baseline
	var reductionPct float64
	if pm.baseline != nil && pm.baseline.SystemMetrics.MemoryUsage.Alloc > 0 {
		baseline := float64(pm.baseline.SystemMetrics.MemoryUsage.Alloc)
		current := float64(m.Alloc)
		reductionPct = (baseline - current) / baseline * 100
	}
	
	pm.metrics.SystemMetrics = SystemResourceMetrics{
		MemoryUsage: MemoryUsage{
			Alloc:        m.Alloc,
			TotalAlloc:   m.TotalAlloc,
			Sys:          m.Sys,
			NumGC:        m.NumGC,
			HeapAlloc:    m.HeapAlloc,
			HeapSys:      m.HeapSys,
			StackInuse:   m.StackInuse,
			ReductionPct: reductionPct,
		},
		CPUUsage: CPUUsage{
			NumCPU:     runtime.NumCPU(),
			GOMAXPROCS: runtime.GOMAXPROCS(0),
			NumCgoCall: runtime.NumCgoCall(),
		},
		GoroutineCount: runtime.NumGoroutine(),
	}
	
	if m.NumGC > 0 {
		pm.metrics.SystemMetrics.LastGC = time.Unix(0, int64(m.LastGC))
		pm.metrics.SystemMetrics.GCPause = time.Duration(m.PauseNs[(m.NumGC+255)%256])
	}
}

// UpdateCacheMetrics updates cache performance statistics
func (pm *PerformanceMonitor) UpdateCacheMetrics(hitCount, missCount int64, size int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	totalRequests := hitCount + missCount
	var hitRatio float64
	if totalRequests > 0 {
		hitRatio = float64(hitCount) / float64(totalRequests) * 100
	}
	
	pm.metrics.CacheMetrics = CachePerformanceMetrics{
		HitCount:    hitCount,
		MissCount:   missCount,
		HitRatio:    hitRatio,
		Size:        size,
		LastCleanup: time.Now(),
	}
}

// UpdateWorkerMetrics updates worker pool performance statistics
func (pm *PerformanceMonitor) UpdateWorkerMetrics(poolName string, metrics PoolMetrics) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	pm.metrics.WorkerMetrics[poolName] = metrics
}

// GetMetrics returns a copy of current performance metrics
func (pm *PerformanceMonitor) GetMetrics() PerformanceMetrics {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	pm.UpdateSystemMetrics()
	
	// Return a copy to avoid race conditions
	metricsCopy := *pm.metrics
	metricsCopy.LastUpdated = time.Now()
	
	return metricsCopy
}

// SetBaseline sets the baseline metrics for comparison
func (pm *PerformanceMonitor) SetBaseline() {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	pm.UpdateSystemMetrics()
	baseline := *pm.metrics
	pm.baseline = &baseline
}

// GetPerformanceReport generates a comprehensive performance report
func (pm *PerformanceMonitor) GetPerformanceReport() map[string]interface{} {
	metrics := pm.GetMetrics()
	
	report := map[string]interface{}{
		"summary": map[string]interface{}{
			"uptime":                time.Since(pm.startTime).String(),
			"total_requests":        metrics.RequestCounts.TotalRequests,
			"requests_per_second":   metrics.RequestCounts.RequestsPerSecond,
			"average_response_time": metrics.ResponseTimes.AverageResponseTime.String(),
			"memory_usage_mb":       float64(metrics.SystemMetrics.MemoryUsage.Alloc) / 1024 / 1024,
			"memory_reduction_pct":  metrics.SystemMetrics.MemoryUsage.ReductionPct,
			"cache_hit_ratio":       metrics.CacheMetrics.HitRatio,
			"database_queries":      metrics.DatabaseMetrics.QueryCount,
			"slow_queries":          metrics.DatabaseMetrics.SlowQueries,
		},
		"response_times": map[string]interface{}{
			"average": metrics.ResponseTimes.AverageResponseTime.String(),
			"p50":     metrics.ResponseTimes.P50ResponseTime.String(),
			"p95":     metrics.ResponseTimes.P95ResponseTime.String(),
			"p99":     metrics.ResponseTimes.P99ResponseTime.String(),
		},
		"system_resources": map[string]interface{}{
			"memory_mb":       float64(metrics.SystemMetrics.MemoryUsage.Alloc) / 1024 / 1024,
			"heap_mb":         float64(metrics.SystemMetrics.MemoryUsage.HeapAlloc) / 1024 / 1024,
			"goroutines":      metrics.SystemMetrics.GoroutineCount,
			"gc_runs":         metrics.SystemMetrics.MemoryUsage.NumGC,
			"gc_pause_ms":     metrics.SystemMetrics.GCPause.Nanoseconds() / 1e6,
		},
		"cache_performance": map[string]interface{}{
			"hit_ratio":   metrics.CacheMetrics.HitRatio,
			"cache_size":  metrics.CacheMetrics.Size,
			"hit_count":   metrics.CacheMetrics.HitCount,
			"miss_count":  metrics.CacheMetrics.MissCount,
		},
		"worker_pools": metrics.WorkerMetrics,
		"database": map[string]interface{}{
			"total_queries":      metrics.DatabaseMetrics.QueryCount,
			"average_query_time": metrics.DatabaseMetrics.AverageQueryTime.String(),
			"slow_queries":       metrics.DatabaseMetrics.SlowQueries,
		},
	}
	
	return report
}

// ExportMetricsJSON exports metrics as JSON for external monitoring tools
func (pm *PerformanceMonitor) ExportMetricsJSON() ([]byte, error) {
	metrics := pm.GetMetrics()
	return json.Marshal(metrics)
}

// backgroundMonitoring runs periodic system metrics updates
func (pm *PerformanceMonitor) backgroundMonitoring() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		pm.UpdateSystemMetrics()
		
		// Update worker pool metrics
		if GlobalSubtitlePools != nil {
			workerStats := GlobalSubtitlePools.GetPoolStats()
			for name, metrics := range workerStats {
				pm.UpdateWorkerMetrics(name, metrics)
			}
		}
	}
}

// calculateResponseTimePercentiles calculates response time percentiles
func (pm *PerformanceMonitor) calculateResponseTimePercentiles() {
	times := pm.metrics.ResponseTimes.responseTimes
	if len(times) == 0 {
		return
	}
	
	// Simple percentile calculation (could be optimized with proper sorting)
	total := time.Duration(0)
	for _, t := range times {
		total += t
	}
	pm.metrics.ResponseTimes.AverageResponseTime = total / time.Duration(len(times))
	
	// For now, use average for all percentiles (could implement proper percentile calculation)
	pm.metrics.ResponseTimes.P50ResponseTime = pm.metrics.ResponseTimes.AverageResponseTime
	pm.metrics.ResponseTimes.P95ResponseTime = pm.metrics.ResponseTimes.AverageResponseTime
	pm.metrics.ResponseTimes.P99ResponseTime = pm.metrics.ResponseTimes.AverageResponseTime
}

// calculateRequestsPerSecond calculates current requests per second
func (pm *PerformanceMonitor) calculateRequestsPerSecond() {
	elapsed := time.Since(pm.startTime).Seconds()
	if elapsed > 0 {
		pm.metrics.RequestCounts.RequestsPerSecond = float64(pm.metrics.RequestCounts.TotalRequests) / elapsed
	}
}

// Global performance monitor instance
var GlobalPerformanceMonitor = NewPerformanceMonitor()