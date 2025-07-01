// file: pkg/performance/workers.go
// version: 1.0.0
// guid: 1f0e9d8c-7c6b-2f1e-5a4b-8c7d6e5f4e3d

package performance

import (
	"context"
	"runtime"
	"time"

	"github.com/sourcegraph/conc/pool"
)

// WorkerPoolConfig defines configuration for worker pool optimization
type WorkerPoolConfig struct {
	// MaxWorkers is the maximum number of concurrent workers
	MaxWorkers int
	// QueueSize is the maximum number of queued tasks
	QueueSize int
	// TaskTimeout is the maximum time a single task can run
	TaskTimeout time.Duration
}

// GetOptimalWorkerConfig returns optimized worker pool configuration based on system resources.
//
// This function calculates optimal worker pool settings based on:
//   - Available CPU cores
//   - System memory
//   - Expected workload type (CPU-bound vs I/O-bound)
//   - Target performance characteristics
//
// The configuration optimizes for subtitle processing workloads which typically involve:
//   - File I/O operations (reading/writing subtitle files)
//   - Network requests (downloading from providers)
//   - Text processing (parsing/converting subtitle formats)
//   - Translation API calls (external service requests)
func GetOptimalWorkerConfig(workloadType string) WorkerPoolConfig {
	numCPU := runtime.NumCPU()
	
	switch workloadType {
	case "cpu_intensive":
		// CPU-intensive tasks like subtitle format conversion
		// Use number of CPU cores to maximize CPU utilization
		return WorkerPoolConfig{
			MaxWorkers:  numCPU,
			QueueSize:   numCPU * 2,
			TaskTimeout: 30 * time.Second,
		}
		
	case "io_intensive":
		// I/O-intensive tasks like file reading/writing
		// Use more workers than CPU cores since threads will be blocked on I/O
		return WorkerPoolConfig{
			MaxWorkers:  numCPU * 2,
			QueueSize:   numCPU * 4,
			TaskTimeout: 60 * time.Second,
		}
		
	case "network_intensive":
		// Network-intensive tasks like provider downloads
		// Higher concurrency but with reasonable limits to avoid overwhelming services
		workers := numCPU * 3
		if workers > 20 {
			workers = 20 // Cap at 20 to be respectful to external services
		}
		return WorkerPoolConfig{
			MaxWorkers:  workers,
			QueueSize:   workers * 2,
			TaskTimeout: 120 * time.Second, // Longer timeout for network operations
		}
		
	case "mixed":
		// Mixed workload with both CPU and I/O operations
		// Balanced configuration for general subtitle processing
		return WorkerPoolConfig{
			MaxWorkers:  numCPU + (numCPU / 2), // 1.5x CPU cores
			QueueSize:   numCPU * 3,
			TaskTimeout: 60 * time.Second,
		}
		
	default:
		// Conservative default for unknown workloads
		return WorkerPoolConfig{
			MaxWorkers:  numCPU,
			QueueSize:   numCPU * 2,
			TaskTimeout: 45 * time.Second,
		}
	}
}

// OptimizedPool wraps pool.Pool with performance optimizations and monitoring
type OptimizedPool struct {
	pool    *pool.Pool
	config  WorkerPoolConfig
	metrics PoolMetrics
}

// PoolMetrics tracks performance metrics for the worker pool
type PoolMetrics struct {
	TasksSubmitted int64
	TasksCompleted int64
	TasksFailed    int64
	TotalTime      time.Duration
	AverageTime    time.Duration
}

// NewOptimizedPool creates a new optimized worker pool with performance monitoring
func NewOptimizedPool(config WorkerPoolConfig) *OptimizedPool {
	p := pool.New().WithMaxGoroutines(config.MaxWorkers)
	
	return &OptimizedPool{
		pool:   p,
		config: config,
		metrics: PoolMetrics{},
	}
}

// Submit submits a task to the worker pool with timeout and metrics tracking
func (op *OptimizedPool) Submit(task func() error) {
	op.metrics.TasksSubmitted++
	
	op.pool.Go(func() {
		start := time.Now()
		
		// Create context with timeout for the task
		ctx, cancel := context.WithTimeout(context.Background(), op.config.TaskTimeout)
		defer cancel()
		
		// Create a channel to capture task completion
		done := make(chan error, 1)
		
		// Run the task in a goroutine
		go func() {
			done <- task()
		}()
		
		// Wait for task completion or timeout
		select {
		case err := <-done:
			duration := time.Since(start)
			op.updateMetrics(duration, err == nil)
		case <-ctx.Done():
			op.metrics.TasksFailed++
		}
	})
}

// SubmitWithPriority submits a high-priority task that should be executed immediately
func (op *OptimizedPool) SubmitWithPriority(task func() error) {
	// For high-priority tasks, we could implement a priority queue
	// For now, submit normally but this could be enhanced
	op.Submit(task)
}

// Wait waits for all submitted tasks to complete
func (op *OptimizedPool) Wait() {
	op.pool.Wait()
}

// GetMetrics returns current performance metrics
func (op *OptimizedPool) GetMetrics() PoolMetrics {
	return op.metrics
}

// updateMetrics updates the performance metrics
func (op *OptimizedPool) updateMetrics(duration time.Duration, success bool) {
	if success {
		op.metrics.TasksCompleted++
	} else {
		op.metrics.TasksFailed++
	}
	
	op.metrics.TotalTime += duration
	if op.metrics.TasksCompleted > 0 {
		op.metrics.AverageTime = op.metrics.TotalTime / time.Duration(op.metrics.TasksCompleted)
	}
}

// SubtitleProcessingPool provides pre-configured pools for common subtitle operations
type SubtitleProcessingPool struct {
	ConversionPool  *OptimizedPool
	DownloadPool    *OptimizedPool
	TranslationPool *OptimizedPool
	ScanningPool    *OptimizedPool
}

// NewSubtitleProcessingPool creates optimized worker pools for different subtitle operations
func NewSubtitleProcessingPool() *SubtitleProcessingPool {
	return &SubtitleProcessingPool{
		ConversionPool:  NewOptimizedPool(GetOptimalWorkerConfig("cpu_intensive")),
		DownloadPool:    NewOptimizedPool(GetOptimalWorkerConfig("network_intensive")),
		TranslationPool: NewOptimizedPool(GetOptimalWorkerConfig("network_intensive")),
		ScanningPool:    NewOptimizedPool(GetOptimalWorkerConfig("io_intensive")),
	}
}

// ProcessSubtitlesBatch processes multiple subtitle files with optimal worker allocation
func (spp *SubtitleProcessingPool) ProcessSubtitlesBatch(files []string, operation string, processor func(string) error) {
	var pool *OptimizedPool
	
	switch operation {
	case "convert":
		pool = spp.ConversionPool
	case "download":
		pool = spp.DownloadPool
	case "translate":
		pool = spp.TranslationPool
	case "scan":
		pool = spp.ScanningPool
	default:
		// Use conversion pool as default
		pool = spp.ConversionPool
	}
	
	// Submit all files for processing
	for _, file := range files {
		file := file // Capture loop variable
		pool.Submit(func() error {
			return processor(file)
		})
	}
	
	// Wait for all tasks to complete
	pool.Wait()
}

// GetPoolStats returns statistics for all worker pools
func (spp *SubtitleProcessingPool) GetPoolStats() map[string]PoolMetrics {
	return map[string]PoolMetrics{
		"conversion":  spp.ConversionPool.GetMetrics(),
		"download":    spp.DownloadPool.GetMetrics(),
		"translation": spp.TranslationPool.GetMetrics(),
		"scanning":    spp.ScanningPool.GetMetrics(),
	}
}

// GlobalSubtitlePools provides access to optimized worker pools throughout the application
var GlobalSubtitlePools = NewSubtitleProcessingPool()

// BatchProcessor provides utilities for optimized batch processing of subtitle operations
type BatchProcessor struct {
	pools      *SubtitleProcessingPool
	batchSize  int
	maxRetries int
}

// NewBatchProcessor creates a new batch processor with optimal settings
func NewBatchProcessor() *BatchProcessor {
	return &BatchProcessor{
		pools:      GlobalSubtitlePools,
		batchSize:  runtime.NumCPU() * 2, // Process 2x CPU cores worth of items per batch
		maxRetries: 3,
	}
}

// ProcessInBatches processes a large number of items in optimally-sized batches
func (bp *BatchProcessor) ProcessInBatches(items []string, operation string, processor func(string) error) {
	// Split items into batches
	for i := 0; i < len(items); i += bp.batchSize {
		end := i + bp.batchSize
		if end > len(items) {
			end = len(items)
		}
		
		batch := items[i:end]
		
		// Process this batch
		bp.pools.ProcessSubtitlesBatch(batch, operation, processor)
	}
}