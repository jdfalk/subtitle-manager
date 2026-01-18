// file: pkg/queue/jobs_test.go
// version: 1.0.0
// guid: 2b1c4a4e-7f2d-4d31-93d2-1b8f1f19f8c0
package queue

import (
	"context"
	"strings"
	"testing"

	jobpb "github.com/jdfalk/subtitle-manager/pkg/jobpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestSingleFileJob_MetadataAndQueueMessage_PopulatesFields(t *testing.T) {
	// Arrange
	job := &SingleFileJob{
		JobID:      "job-123",
		InputPath:  "/tmp/input.srt",
		OutputPath: "/tmp/output.srt",
		Language:   "es",
		Service:    "grpc",
		GoogleKey:  "g-key",
		GPTKey:     "gpt-key",
		GRPCAddr:   "grpc:1234",
	}

	// Exercise
	queueMsg, err := job.QueueMessage()

	// Verify
	if err != nil {
		t.Fatalf("queue message: %v", err)
	}
	if job.ID() != "job-123" {
		t.Fatalf("expected id job-123, got %s", job.ID())
	}
	if job.Type() != JobTypeSingleFile {
		t.Fatalf("expected type %s, got %s", JobTypeSingleFile, job.Type())
	}
	if job.Description() == "" {
		t.Fatal("expected description to be populated")
	}

	payload := &jobpb.TranslationJob{}
	if err := anypb.UnmarshalTo(queueMsg.GetBody(), payload, proto.UnmarshalOptions{}); err != nil {
		t.Fatalf("unmarshal body: %v", err)
	}
	if queueMsg.GetId() != "job-123" {
		t.Fatalf("expected queue message id job-123, got %s", queueMsg.GetId())
	}
	if got := payload.GetInputPaths(); len(got) != 1 || got[0] != "/tmp/input.srt" {
		t.Fatalf("expected input paths [/tmp/input.srt], got %v", got)
	}
	if payload.GetOutputPath() != "/tmp/output.srt" {
		t.Fatalf("expected output path /tmp/output.srt, got %s", payload.GetOutputPath())
	}
	if payload.GetLanguage() != "es" {
		t.Fatalf("expected language es, got %s", payload.GetLanguage())
	}
	if payload.GetService() != "grpc" {
		t.Fatalf("expected service grpc, got %s", payload.GetService())
	}
	if payload.GetGoogleKey() != "g-key" {
		t.Fatalf("expected google key g-key, got %s", payload.GetGoogleKey())
	}
	if payload.GetGptKey() != "gpt-key" {
		t.Fatalf("expected gpt key gpt-key, got %s", payload.GetGptKey())
	}
	if payload.GetGrpcAddr() != "grpc:1234" {
		t.Fatalf("expected grpc addr grpc:1234, got %s", payload.GetGrpcAddr())
	}
	if payload.GetWorkers() != 1 {
		t.Fatalf("expected workers 1, got %d", payload.GetWorkers())
	}
}

func TestSingleFileJob_Execute_ReturnsPathValidationError(t *testing.T) {
	// Arrange
	job := &SingleFileJob{
		JobID:      "job-err",
		InputPath:  "relative/input.srt",
		OutputPath: "relative/output.srt",
		Language:   "es",
		Service:    "grpc",
	}

	// Exercise
	err := job.Execute(context.Background())

	// Verify
	if err == nil {
		t.Fatal("expected error for relative path, got nil")
	}
	if !strings.Contains(err.Error(), "path must be absolute") {
		t.Fatalf("expected path validation error, got %v", err)
	}
}

func TestBatchFilesJob_MetadataAndQueueMessage_PopulatesFields(t *testing.T) {
	// Arrange
	job := &BatchFilesJob{
		JobID:      "job-456",
		InputPaths: []string{"/tmp/one.srt", "/tmp/two.srt"},
		Language:   "fr",
		Service:    "google",
		GoogleKey:  "g-key",
		GPTKey:     "gpt-key",
		GRPCAddr:   "grpc:1234",
		Workers:    4,
	}

	// Exercise
	queueMsg, err := job.QueueMessage()

	// Verify
	if err != nil {
		t.Fatalf("queue message: %v", err)
	}
	if job.ID() != "job-456" {
		t.Fatalf("expected id job-456, got %s", job.ID())
	}
	if job.Type() != JobTypeBatchFiles {
		t.Fatalf("expected type %s, got %s", JobTypeBatchFiles, job.Type())
	}
	if job.Description() == "" {
		t.Fatal("expected description to be populated")
	}

	payload := &jobpb.TranslationJob{}
	if err := anypb.UnmarshalTo(queueMsg.GetBody(), payload, proto.UnmarshalOptions{}); err != nil {
		t.Fatalf("unmarshal body: %v", err)
	}
	if queueMsg.GetId() != "job-456" {
		t.Fatalf("expected queue message id job-456, got %s", queueMsg.GetId())
	}
	if got := payload.GetInputPaths(); len(got) != 2 || got[0] != "/tmp/one.srt" || got[1] != "/tmp/two.srt" {
		t.Fatalf("expected input paths [/tmp/one.srt /tmp/two.srt], got %v", got)
	}
	if payload.GetLanguage() != "fr" {
		t.Fatalf("expected language fr, got %s", payload.GetLanguage())
	}
	if payload.GetService() != "google" {
		t.Fatalf("expected service google, got %s", payload.GetService())
	}
	if payload.GetGoogleKey() != "g-key" {
		t.Fatalf("expected google key g-key, got %s", payload.GetGoogleKey())
	}
	if payload.GetGptKey() != "gpt-key" {
		t.Fatalf("expected gpt key gpt-key, got %s", payload.GetGptKey())
	}
	if payload.GetGrpcAddr() != "grpc:1234" {
		t.Fatalf("expected grpc addr grpc:1234, got %s", payload.GetGrpcAddr())
	}
	if payload.GetWorkers() != 4 {
		t.Fatalf("expected workers 4, got %d", payload.GetWorkers())
	}
}

func TestBatchFilesJob_Execute_ReturnsPathValidationError(t *testing.T) {
	// Arrange
	job := &BatchFilesJob{
		JobID:      "job-err",
		InputPaths: []string{"relative/one.srt", "relative/two.srt"},
		Language:   "fr",
		Service:    "google",
		Workers:    2,
	}

	// Exercise
	err := job.Execute(context.Background())

	// Verify
	if err == nil {
		t.Fatal("expected error for relative path, got nil")
	}
	if !strings.Contains(err.Error(), "path must be absolute") {
		t.Fatalf("expected path validation error, got %v", err)
	}
}
