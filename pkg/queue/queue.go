// file: pkg/queue/queue.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174002
package queue

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/tasks"
)

// Queue manages asynchronous translation jobs using a worker pool.
type Queue struct {
	mu      sync.RWMutex
	jobs    chan Job
	workers int
	running bool
	logger  *logrus.Entry
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
	stopped bool // Track if queue has been stopped
}

// NewQueue creates a new translation queue with the specified number of workers.
func NewQueue(workers int) *Queue {
	ctx, cancel := context.WithCancel(context.Background())
	return &Queue{
		jobs:    make(chan Job, 100), // Buffer of 100 jobs
		workers: workers,
		logger:  logging.GetLogger("queue"),
		ctx:     ctx,
		cancel:  cancel,
	}
}

// Start begins processing jobs in the queue with the configured number of workers.
func (q *Queue) Start() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.running {
		return fmt.Errorf("queue is already running")
	}

	// Reinitialize context and jobs channel if they were closed from a previous Stop()
	if q.stopped {
		q.ctx, q.cancel = context.WithCancel(context.Background())
		q.jobs = make(chan Job, 100) // Buffer of 100 jobs
		q.stopped = false
	}

	q.logger.Infof("Starting translation queue with %d workers", q.workers)
	q.running = true

	// Start worker goroutines
	for i := 0; i < q.workers; i++ {
		q.wg.Add(1)
		go q.worker(i)
	}

	return nil
}

// Stop gracefully shuts down the queue, waiting for current jobs to complete.
func (q *Queue) Stop() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if !q.running {
		return fmt.Errorf("queue is not running")
	}

	q.logger.Info("Stopping translation queue")
	q.running = false
	q.stopped = true
	q.cancel()
	close(q.jobs)
	q.wg.Wait()

	q.logger.Info("Translation queue stopped")
	return nil
}

// IsRunning returns true if the queue is currently processing jobs.
func (q *Queue) IsRunning() bool {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.running
}

// Add queues a job for asynchronous processing and returns the task ID for tracking.
func (q *Queue) Add(job Job) (string, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if !q.running {
		return "", fmt.Errorf("queue is not running")
	}

	select {
	case q.jobs <- job:
		q.logger.Infof("Queued job %s: %s", job.ID(), job.Description())
		return job.ID(), nil
	default:
		return "", fmt.Errorf("queue is full")
	}
}

// worker processes jobs from the queue.
func (q *Queue) worker(id int) {
	defer q.wg.Done()
	workerLogger := q.logger.WithField("worker", id)
	workerLogger.Debug("Worker started")

	for {
		select {
		case job, ok := <-q.jobs:
			if !ok {
				workerLogger.Debug("Worker stopping: jobs channel closed")
				return
			}
			q.processJob(workerLogger, job)

		case <-q.ctx.Done():
			workerLogger.Debug("Worker stopping: context cancelled")
			return
		}
	}
}

// processJob executes a single job and tracks its progress.
func (q *Queue) processJob(logger *logrus.Entry, job Job) {
	jobID := job.ID()
	logger = logger.WithField("job", jobID)
	logger.Infof("Processing job: %s", job.Description())

	// Start tracking the job
	task := tasks.Start(q.ctx, jobID, func(ctx context.Context) error {
		return job.Execute(ctx)
	})

	// Monitor job progress
	startTime := time.Now()
	for {
		status := task.GetStatus()
		if status == "completed" {
			logger.Infof("Job completed successfully in %v", time.Since(startTime))
			return
		} else if status == "failed" {
			logger.Errorf("Job failed: %s", task.GetError())
			return
		}

		// Check if context is cancelled
		select {
		case <-q.ctx.Done():
			logger.Warn("Job cancelled due to queue shutdown")
			return
		case <-time.After(100 * time.Millisecond):
			// Continue monitoring
		}
	}
}

// QueueStatus provides information about the current state of the queue.
type QueueStatus struct {
	Running     bool `json:"running"`
	Workers     int  `json:"workers"`
	QueueLength int  `json:"queue_length"`
}

// Status returns the current status of the queue.
func (q *Queue) Status() QueueStatus {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return QueueStatus{
		Running:     q.running,
		Workers:     q.workers,
		QueueLength: len(q.jobs),
	}
}

// Global queue instance
var (
	globalQueue *Queue
	globalMu    sync.RWMutex
)

// GetQueue returns the global queue instance, creating it if necessary.
func GetQueue() *Queue {
	globalMu.Lock()
	defer globalMu.Unlock()

	if globalQueue == nil {
		globalQueue = NewQueue(3) // Default to 3 workers
	}
	return globalQueue
}

// SetQueue sets the global queue instance (useful for testing).
func SetQueue(q *Queue) {
	globalMu.Lock()
	defer globalMu.Unlock()
	globalQueue = q
}

// Helper functions for creating jobs

// NewSingleFileJob creates a new single file translation job with a generated ID.
func NewSingleFileJob(inputPath, outputPath, language, service, googleKey, gptKey, grpcAddr string) *SingleFileJob {
	return &SingleFileJob{
		JobID:      uuid.New().String(),
		InputPath:  inputPath,
		OutputPath: outputPath,
		Language:   language,
		Service:    service,
		GoogleKey:  googleKey,
		GPTKey:     gptKey,
		GRPCAddr:   grpcAddr,
	}
}

// NewBatchFilesJob creates a new batch files translation job with a generated ID.
func NewBatchFilesJob(inputPaths []string, language, service, googleKey, gptKey, grpcAddr string, workers int) *BatchFilesJob {
	return &BatchFilesJob{
		JobID:      uuid.New().String(),
		InputPaths: inputPaths,
		Language:   language,
		Service:    service,
		GoogleKey:  googleKey,
		GPTKey:     gptKey,
		GRPCAddr:   grpcAddr,
		Workers:    workers,
	}
}
