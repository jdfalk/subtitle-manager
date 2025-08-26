// file: pkg/queue/job.go
// version: 1.0.0
// guid: abc12345-def6-7890-abcd-ef1234567890

package queue

import (
	"context"
	"fmt"
)

// Job represents a queued job
type Job interface {
	Execute(ctx context.Context) error
	ID() string
	Type() string
	Description() string
	QueueMessage() (interface{}, error)
}

// JobType defines the type of job
type JobType string

const (
	JobTypeSingleFile = "single_file"
	JobTypeBatchFiles = "batch_files"
)

// SingleFileJob represents a job for processing a single file
type SingleFileJob struct {
	JobID      string
	InputPath  string
	OutputPath string
	Language   string
	Service    string
	GoogleKey  string
	GPTKey     string
	GRPCAddr   string
}

// ID returns the job identifier.
func (j *SingleFileJob) ID() string {
	return j.JobID
}

// Type returns the job type.
func (j *SingleFileJob) Type() string {
	return JobTypeSingleFile
}

// Description returns a human-readable description of the job.
func (j *SingleFileJob) Description() string {
	return fmt.Sprintf("Translate %s to %s", j.InputPath, j.Language)
}

// QueueMessage returns the queue message for this job.
func (j *SingleFileJob) QueueMessage() (interface{}, error) {
	return map[string]interface{}{
		"job_id":     j.JobID,
		"input":      j.InputPath,
		"output":     j.OutputPath,
		"language":   j.Language,
		"service":    j.Service,
	}, nil
}

// Execute processes the single file job.
func (j *SingleFileJob) Execute(ctx context.Context) error {
	// Placeholder implementation
	return nil
}

// BatchFilesJob represents a job for processing multiple files
type BatchFilesJob struct {
	JobID      string
	InputPaths []string
	Language   string
	Service    string
	GoogleKey  string
	GPTKey     string
	GRPCAddr   string
	Workers    int
}

// ID returns the job identifier.
func (j *BatchFilesJob) ID() string {
	return j.JobID
}

// Type returns the job type.
func (j *BatchFilesJob) Type() string {
	return JobTypeBatchFiles
}

// Description returns a human-readable description of the job.
func (j *BatchFilesJob) Description() string {
	return fmt.Sprintf("Batch translate %d files to %s", len(j.InputPaths), j.Language)
}

// QueueMessage returns the queue message for this job.
func (j *BatchFilesJob) QueueMessage() (interface{}, error) {
	return map[string]interface{}{
		"job_id":   j.JobID,
		"inputs":   j.InputPaths,
		"language": j.Language,
		"service":  j.Service,
		"workers":  j.Workers,
	}, nil
}

// Execute processes the batch files job.
func (j *BatchFilesJob) Execute(ctx context.Context) error {
	// Placeholder implementation
	return nil
}