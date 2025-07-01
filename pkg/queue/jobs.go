// file: pkg/queue/jobs.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174001
package queue

import (
	"context"
	"fmt"

	"github.com/jdfalk/subtitle-manager/pkg/subtitles"
)

// JobType represents the type of translation job.
type JobType string

const (
	// JobTypeSingleFile represents a single file translation job.
	JobTypeSingleFile JobType = "single_file"
	// JobTypeBatchFiles represents a batch file translation job.
	JobTypeBatchFiles JobType = "batch_files"
)

// Job represents a translation job that can be queued for asynchronous processing.
type Job interface {
	// ID returns the unique identifier for this job.
	ID() string
	// Type returns the type of job.
	Type() JobType
	// Execute runs the job and returns an error if it fails.
	Execute(ctx context.Context) error
	// Description returns a human-readable description of the job.
	Description() string
}

// SingleFileJob represents a job to translate a single subtitle file.
type SingleFileJob struct {
	JobID     string
	InputPath string
	OutputPath string
	Language  string
	Service   string
	GoogleKey string
	GPTKey    string
	GRPCAddr  string
}

// ID returns the job identifier.
func (j *SingleFileJob) ID() string {
	return j.JobID
}

// Type returns the job type.
func (j *SingleFileJob) Type() JobType {
	return JobTypeSingleFile
}

// Execute performs the translation.
func (j *SingleFileJob) Execute(ctx context.Context) error {
	return subtitles.TranslateFileToSRT(
		j.InputPath,
		j.OutputPath,
		j.Language,
		j.Service,
		j.GoogleKey,
		j.GPTKey,
		j.GRPCAddr,
	)
}

// Description returns a description of the job.
func (j *SingleFileJob) Description() string {
	return fmt.Sprintf("Translate %s to %s (%s)", j.InputPath, j.Language, j.OutputPath)
}

// BatchFilesJob represents a job to translate multiple subtitle files.
type BatchFilesJob struct {
	JobID     string
	InputPaths []string
	Language  string
	Service   string
	GoogleKey string
	GPTKey    string
	GRPCAddr  string
	Workers   int
}

// ID returns the job identifier.
func (j *BatchFilesJob) ID() string {
	return j.JobID
}

// Type returns the job type.
func (j *BatchFilesJob) Type() JobType {
	return JobTypeBatchFiles
}

// Execute performs the batch translation.
func (j *BatchFilesJob) Execute(ctx context.Context) error {
	return subtitles.TranslateFilesToSRT(
		j.InputPaths,
		j.Language,
		j.Service,
		j.GoogleKey,
		j.GPTKey,
		j.GRPCAddr,
		j.Workers,
	)
}

// Description returns a description of the job.
func (j *BatchFilesJob) Description() string {
	return fmt.Sprintf("Translate %d files to %s", len(j.InputPaths), j.Language)
}