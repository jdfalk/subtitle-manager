// file: pkg/transcriber/docker_test.go
// version: 1.0.0
// guid: b2c3d4e5-f6g7-8901-bcde-f23456789012

package transcriber

import (
	"context"
	"testing"
	"time"
)

// TestDefaultDockerConfig verifies the default configuration is sensible.
func TestDefaultDockerConfig(t *testing.T) {
	config := DefaultDockerConfig()

	if config.Image == "" {
		t.Error("default image should not be empty")
	}
	if config.Model == "" {
		t.Error("default model should not be empty")
	}
	if config.Device == "" {
		t.Error("default device should not be empty")
	}
	if config.Timeout == 0 {
		t.Error("default timeout should not be zero")
	}

	// Verify expected defaults
	if config.Model != "base" {
		t.Errorf("expected default model 'base', got %s", config.Model)
	}
	if config.Device != "cpu" {
		t.Errorf("expected default device 'cpu', got %s", config.Device)
	}
	if config.Timeout != 10*time.Minute {
		t.Errorf("expected default timeout 10m, got %v", config.Timeout)
	}
}

// TestNewDockerTranscriber verifies transcriber creation.
func TestNewDockerTranscriber(t *testing.T) {
	// Test with default config
	transcriber, err := NewDockerTranscriber(nil)
	if err != nil {
		t.Skipf("Docker not available: %v", err)
	}
	defer transcriber.Close()

	if transcriber.config == nil {
		t.Error("config should not be nil")
	}
	if transcriber.client == nil {
		t.Error("client should not be nil")
	}
}

// TestDockerTranscriberIsAvailable checks Docker availability detection.
func TestDockerTranscriberIsAvailable(t *testing.T) {
	transcriber, err := NewDockerTranscriber(nil)
	if err != nil {
		t.Skipf("Docker client creation failed: %v", err)
	}
	defer transcriber.Close()

	ctx := context.Background()
	available := transcriber.IsAvailable(ctx)

	// This test just verifies the method works without panicking
	// The actual result depends on Docker being installed/running
	t.Logf("Docker available: %v", available)
}

// TestTranscribeWithMethod verifies the unified transcription interface.
func TestTranscribeWithMethod(t *testing.T) {
	ctx := context.Background()

	// Test unsupported method
	_, err := TranscribeWithMethod(ctx, "unsupported", "test.wav", "en", "", nil)
	if err == nil {
		t.Error("expected error for unsupported method")
	}

	// Test Docker method (will fail if Docker not available, which is expected)
	_, err = TranscribeWithMethod(ctx, MethodDocker, "test.wav", "en", "", nil)
	if err == nil {
		t.Error("expected error for Docker method without Docker")
	}
	t.Logf("Docker method error (expected): %v", err)

	// Test OpenAI method (will fail without API key, which is expected)
	_, err = TranscribeWithMethod(ctx, MethodOpenAI, "test.wav", "en", "", nil)
	if err == nil {
		t.Error("expected error for OpenAI method without API key")
	}
	t.Logf("OpenAI method error (expected): %v", err)
}

// TestTranscriptionMethods verifies method constants.
func TestTranscriptionMethods(t *testing.T) {
	if MethodOpenAI != "openai" {
		t.Errorf("expected MethodOpenAI to be 'openai', got %s", MethodOpenAI)
	}
	if MethodDocker != "docker" {
		t.Errorf("expected MethodDocker to be 'docker', got %s", MethodDocker)
	}
}
