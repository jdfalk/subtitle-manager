// file: pkg/performance/performance_test.go
// version: 1.0.0
// guid: 6f5e4d3c-2f1e-7f6e-0a9b-3f2e1f0e9f8e

package performance

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetOptimalWorkerConfig(t *testing.T) {
	tests := []struct {
		name         string
		workloadType string
		expectedMin  int
		expectedMax  int
	}{
		{
			name:         "CPU intensive should use CPU cores",
			workloadType: "cpu_intensive",
			expectedMin:  runtime.NumCPU(),
			expectedMax:  runtime.NumCPU(),
		},
		{
			name:         "IO intensive should use more than CPU cores",
			workloadType: "io_intensive",
			expectedMin:  runtime.NumCPU() * 2,
			expectedMax:  runtime.NumCPU() * 2,
		},
		{
			name:         "Network intensive should be capped at 20",
			workloadType: "network_intensive",
			expectedMin:  runtime.NumCPU(),
			expectedMax:  20,
		},
		{
			name:         "Mixed workload should be balanced",
			workloadType: "mixed",
			expectedMin:  runtime.NumCPU(),
			expectedMax:  runtime.NumCPU() * 2,
		},
		{
			name:         "Unknown workload should use defaults",
			workloadType: "unknown",
			expectedMin:  runtime.NumCPU(),
			expectedMax:  runtime.NumCPU(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := GetOptimalWorkerConfig(tt.workloadType)
			
			if config.MaxWorkers < tt.expectedMin {
				t.Errorf("MaxWorkers %d is less than expected minimum %d", config.MaxWorkers, tt.expectedMin)
			}
			
			if config.MaxWorkers > tt.expectedMax {
				t.Errorf("MaxWorkers %d is greater than expected maximum %d", config.MaxWorkers, tt.expectedMax)
			}
			
			if config.QueueSize <= 0 {
				t.Error("QueueSize should be greater than 0")
			}
			
			if config.TaskTimeout <= 0 {
				t.Error("TaskTimeout should be greater than 0")
			}
		})
	}
}

func TestOptimizedPool(t *testing.T) {
	config := WorkerPoolConfig{
		MaxWorkers:  2,
		QueueSize:   4,
		TaskTimeout: 1 * time.Second,
	}
	
	pool := NewOptimizedPool(config)
	
	// Test basic functionality
	if pool == nil {
		t.Fatal("Failed to create optimized pool")
	}
	
	// Test metrics initialization
	metrics := pool.GetMetrics()
	if metrics.TasksSubmitted != 0 {
		t.Errorf("Expected 0 tasks submitted, got %d", metrics.TasksSubmitted)
	}
	
	// Test task submission
	taskCompleted := make(chan bool, 1)
	
	pool.Submit(func() error {
		time.Sleep(10 * time.Millisecond)
		taskCompleted <- true
		return nil
	})
	
	// Wait for task completion
	select {
	case <-taskCompleted:
		// Task completed successfully
	case <-time.After(500 * time.Millisecond):
		t.Error("Task did not complete within timeout")
	}
	
	pool.Wait()
	
	// Check metrics were updated
	finalMetrics := pool.GetMetrics()
	if finalMetrics.TasksSubmitted == 0 {
		t.Error("Expected tasks submitted to be greater than 0")
	}
}

func TestSubtitleProcessingPool(t *testing.T) {
	pool := NewSubtitleProcessingPool()
	
	if pool.ConversionPool == nil {
		t.Error("ConversionPool should not be nil")
	}
	
	if pool.DownloadPool == nil {
		t.Error("DownloadPool should not be nil")
	}
	
	if pool.TranslationPool == nil {
		t.Error("TranslationPool should not be nil")
	}
	
	if pool.ScanningPool == nil {
		t.Error("ScanningPool should not be nil")
	}
	
	// Test processing batch
	files := []string{"file1.srt", "file2.srt", "file3.srt"}
	processedFiles := make(map[string]bool)
	
	processor := func(file string) error {
		processedFiles[file] = true
		return nil
	}
	
	pool.ProcessSubtitlesBatch(files, "convert", processor)
	
	// Check all files were processed
	for _, file := range files {
		if !processedFiles[file] {
			t.Errorf("File %s was not processed", file)
		}
	}
	
	// Test pool stats
	stats := pool.GetPoolStats()
	if len(stats) == 0 {
		t.Error("Expected pool stats to be available")
	}
}

// TestBatchProcessor creates a simple test without global state
func TestBatchProcessor(t *testing.T) {
	// Create a new batch processor instead of using the global one
	processor := &BatchProcessor{
		pools:      NewSubtitleProcessingPool(),
		batchSize:  2, // Small batch size for testing
		maxRetries: 3,
	}
	
	if processor.batchSize <= 0 {
		t.Error("Batch size should be greater than 0")
	}
	
	// Test processing in batches with a smaller set
	items := []string{"a", "b", "c", "d"}
	
	processedItems := make(map[string]bool)
	var mu sync.Mutex
	
	itemProcessor := func(item string) error {
		mu.Lock()
		processedItems[item] = true
		mu.Unlock()
		return nil
	}
	
	processor.ProcessInBatches(items, "test", itemProcessor)
	
	// Check all items were processed
	mu.Lock()
	for _, item := range items {
		if !processedItems[item] {
			t.Errorf("Item %s was not processed", item)
		}
	}
	mu.Unlock()
}

func BenchmarkWorkerPool(b *testing.B) {
	config := GetOptimalWorkerConfig("cpu_intensive")
	pool := NewOptimizedPool(config)
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		pool.Submit(func() error {
			// Simulate some work
			sum := 0
			for j := 0; j < 1000; j++ {
				sum += j
			}
			return nil
		})
	}
	
	pool.Wait()
}

func BenchmarkBatchProcessing(b *testing.B) {
	processor := NewBatchProcessor()
	items := make([]string, 100)
	for i := range items {
		items[i] = string(rune('a' + i%26))
	}
	
	itemProcessor := func(item string) error {
		// Simulate work
		time.Sleep(1 * time.Microsecond)
		return nil
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		processor.ProcessInBatches(items, "benchmark", itemProcessor)
	}
}

// TestPerformanceMonitorBasics tests basic performance monitoring functionality
func TestPerformanceMonitorBasics(t *testing.T) {
	monitor := NewPerformanceMonitor()
	
	if monitor == nil {
		t.Fatal("Failed to create performance monitor")
	}
	
	// Test recording response time
	monitor.RecordResponseTime("/api/test", 100*time.Millisecond, 200)
	
	// Test recording database query
	monitor.RecordDatabaseQuery(50 * time.Millisecond)
	
	// Test getting metrics
	metrics := monitor.GetMetrics()
	
	if metrics.RequestCounts.TotalRequests == 0 {
		t.Error("Expected total requests to be greater than 0")
	}
	
	if metrics.DatabaseMetrics.QueryCount == 0 {
		t.Error("Expected query count to be greater than 0")
	}
	
	// Test performance report generation
	report := monitor.GetPerformanceReport()
	
	if report == nil {
		t.Error("Expected performance report to be generated")
	}
	
	if summary, ok := report["summary"]; !ok {
		t.Error("Expected summary section in performance report")
	} else if summaryMap, ok := summary.(map[string]interface{}); !ok {
		t.Error("Expected summary to be a map")
	} else if totalRequests, ok := summaryMap["total_requests"]; !ok {
		t.Error("Expected total_requests in summary")
	} else if totalRequests.(int64) == 0 {
		t.Error("Expected total_requests to be greater than 0")
	}
}