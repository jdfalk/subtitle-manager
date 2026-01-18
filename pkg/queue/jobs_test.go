// file: pkg/queue/jobs_test.go
// version: 1.23.0
// guid: cf3906ee-32bb-4288-ba78-d2b0417369f8
package queue

import (
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/jobpb"
)

func TestSingleFileJobQueueMessage(t *testing.T) {
	// Arrange: build a single-file job with all fields populated.
	job := &SingleFileJob{
		JobID:      "job-123",
		InputPath:  "/media/show/episode.srt",
		OutputPath: "/media/show/episode.en.srt",
		Language:   "en",
		Service:    "google",
		GoogleKey:  "google-key",
		GPTKey:     "gpt-key",
		GRPCAddr:   "localhost:9090",
	}

	// Act: convert the job to a queue message and decode the payload.
	queueMessage, err := job.QueueMessage()
	if err != nil {
		t.Fatalf("QueueMessage() error = %v", err)
	}
	if queueMessage.GetBody() == nil {
		t.Fatalf("QueueMessage() body was nil")
	}
	payload := &jobpb.TranslationJob{}
	if err := queueMessage.GetBody().UnmarshalTo(payload); err != nil {
		t.Fatalf("UnmarshalTo() error = %v", err)
	}

	// Assert: verify the queue metadata and payload contents.
	if queueMessage.GetId() != job.JobID {
		t.Errorf("QueueMessage() ID = %q, want %q", queueMessage.GetId(), job.JobID)
	}
	if got := payload.GetInputPaths(); len(got) != 1 || got[0] != job.InputPath {
		t.Errorf("payload input paths = %v, want [%s]", got, job.InputPath)
	}
	if payload.GetOutputPath() != job.OutputPath {
		t.Errorf("payload output path = %q, want %q", payload.GetOutputPath(), job.OutputPath)
	}
	if payload.GetLanguage() != job.Language {
		t.Errorf("payload language = %q, want %q", payload.GetLanguage(), job.Language)
	}
	if payload.GetService() != job.Service {
		t.Errorf("payload service = %q, want %q", payload.GetService(), job.Service)
	}
	if payload.GetGoogleKey() != job.GoogleKey {
		t.Errorf("payload google key = %q, want %q", payload.GetGoogleKey(), job.GoogleKey)
	}
	if payload.GetGptKey() != job.GPTKey {
		t.Errorf("payload gpt key = %q, want %q", payload.GetGptKey(), job.GPTKey)
	}
	if payload.GetGrpcAddr() != job.GRPCAddr {
		t.Errorf("payload grpc addr = %q, want %q", payload.GetGrpcAddr(), job.GRPCAddr)
	}
	if payload.GetWorkers() != 1 {
		t.Errorf("payload workers = %d, want 1", payload.GetWorkers())
	}
}

func TestBatchFilesJobQueueMessage(t *testing.T) {
	// Arrange: build a batch job with multiple inputs and worker count.
	job := &BatchFilesJob{
		JobID:      "job-456",
		InputPaths: []string{"/media/a.srt", "/media/b.srt"},
		Language:   "es",
		Service:    "grpc",
		GoogleKey:  "unused",
		GPTKey:     "unused",
		GRPCAddr:   "localhost:7070",
		Workers:    3,
	}

	// Act: convert the job to a queue message and decode the payload.
	queueMessage, err := job.QueueMessage()
	if err != nil {
		t.Fatalf("QueueMessage() error = %v", err)
	}
	if queueMessage.GetBody() == nil {
		t.Fatalf("QueueMessage() body was nil")
	}
	payload := &jobpb.TranslationJob{}
	if err := queueMessage.GetBody().UnmarshalTo(payload); err != nil {
		t.Fatalf("UnmarshalTo() error = %v", err)
	}

	// Assert: verify the queue metadata and payload contents.
	if queueMessage.GetId() != job.JobID {
		t.Errorf("QueueMessage() ID = %q, want %q", queueMessage.GetId(), job.JobID)
	}
	if got := payload.GetInputPaths(); len(got) != len(job.InputPaths) {
		t.Errorf("payload input paths = %v, want %v", got, job.InputPaths)
	} else {
		for i, path := range job.InputPaths {
			if got[i] != path {
				t.Errorf("payload input path[%d] = %q, want %q", i, got[i], path)
			}
		}
	}
	if payload.GetLanguage() != job.Language {
		t.Errorf("payload language = %q, want %q", payload.GetLanguage(), job.Language)
	}
	if payload.GetService() != job.Service {
		t.Errorf("payload service = %q, want %q", payload.GetService(), job.Service)
	}
	if payload.GetGoogleKey() != job.GoogleKey {
		t.Errorf("payload google key = %q, want %q", payload.GetGoogleKey(), job.GoogleKey)
	}
	if payload.GetGptKey() != job.GPTKey {
		t.Errorf("payload gpt key = %q, want %q", payload.GetGptKey(), job.GPTKey)
	}
	if payload.GetGrpcAddr() != job.GRPCAddr {
		t.Errorf("payload grpc addr = %q, want %q", payload.GetGrpcAddr(), job.GRPCAddr)
	}
	if payload.GetWorkers() != int32(job.Workers) {
		t.Errorf("payload workers = %d, want %d", payload.GetWorkers(), job.Workers)
	}
}

func TestJobDescriptionsAndTypes(t *testing.T) {
	// Arrange: prepare single and batch jobs.
	single := &SingleFileJob{
		JobID:      "job-789",
		InputPath:  "/input.srt",
		OutputPath: "/output.srt",
		Language:   "fr",
	}
	batch := &BatchFilesJob{
		JobID:      "job-012",
		InputPaths: []string{"/input1.srt", "/input2.srt"},
		Language:   "de",
	}

	// Act: collect descriptions and types.
	singleDescription := single.Description()
	batchDescription := batch.Description()

	// Assert: ensure IDs, types, and descriptions are coherent.
	if single.ID() != "job-789" {
		t.Errorf("SingleFileJob ID = %q, want %q", single.ID(), "job-789")
	}
	if single.Type() != JobTypeSingleFile {
		t.Errorf("SingleFileJob Type = %q, want %q", single.Type(), JobTypeSingleFile)
	}
	if batch.ID() != "job-012" {
		t.Errorf("BatchFilesJob ID = %q, want %q", batch.ID(), "job-012")
	}
	if batch.Type() != JobTypeBatchFiles {
		t.Errorf("BatchFilesJob Type = %q, want %q", batch.Type(), JobTypeBatchFiles)
	}
	if singleDescription == "" {
		t.Error("SingleFileJob Description() returned empty string")
	}
	if batchDescription == "" {
		t.Error("BatchFilesJob Description() returned empty string")
	}
}
