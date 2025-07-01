// file: pkg/transcriber/docker.go
// version: 1.0.0
// guid: a1b2c3d4-e5f6-7890-abcd-ef1234567890

package transcriber

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

// DockerTranscriberConfig holds configuration for Docker-based Whisper transcription.
type DockerTranscriberConfig struct {
	// Image is the Docker image to use for Whisper transcription
	Image string `json:"image" yaml:"image"`
	// ContainerName is the name to assign to the container
	ContainerName string `json:"container_name" yaml:"container_name"`
	// WorkDir is the working directory inside the container
	WorkDir string `json:"work_dir" yaml:"work_dir"`
	// Model is the Whisper model to use (tiny, base, small, medium, large)
	Model string `json:"model" yaml:"model"`
	// Language is the language code for transcription (optional)
	Language string `json:"language" yaml:"language"`
	// Device specifies the device to use (cpu, cuda)
	Device string `json:"device" yaml:"device"`
	// Port is the port to expose for API access (if using server mode)
	Port string `json:"port" yaml:"port"`
	// Timeout is the maximum time to wait for transcription
	Timeout time.Duration `json:"timeout" yaml:"timeout"`
}

// DefaultDockerConfig returns default configuration for Docker-based transcription.
func DefaultDockerConfig() *DockerTranscriberConfig {
	return &DockerTranscriberConfig{
		Image:         "openai/whisper:latest",
		ContainerName: "subtitle-manager-whisper",
		WorkDir:       "/workspace",
		Model:         "base",
		Device:        "cpu",
		Port:          "9000",
		Timeout:       10 * time.Minute,
	}
}

// DockerTranscriber manages Docker-based Whisper transcription.
type DockerTranscriber struct {
	client *client.Client
	config *DockerTranscriberConfig
}

// NewDockerTranscriber creates a new Docker-based transcriber.
func NewDockerTranscriber(config *DockerTranscriberConfig) (*DockerTranscriber, error) {
	if config == nil {
		config = DefaultDockerConfig()
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}

	return &DockerTranscriber{
		client: cli,
		config: config,
	}, nil
}

// IsAvailable checks if Docker is available and accessible.
func (dt *DockerTranscriber) IsAvailable(ctx context.Context) bool {
	_, err := dt.client.Ping(ctx)
	return err == nil
}

// EnsureImage pulls the Whisper image if it doesn't exist locally.
func (dt *DockerTranscriber) EnsureImage(ctx context.Context) error {
	images, err := dt.client.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list images: %w", err)
	}

	// Check if image already exists
	for _, img := range images {
		for _, tag := range img.RepoTags {
			if tag == dt.config.Image {
				return nil // Image already exists
			}
		}
	}

	// Pull the image
	reader, err := dt.client.ImagePull(ctx, dt.config.Image, image.PullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image %s: %w", dt.config.Image, err)
	}
	defer reader.Close()

	// Wait for pull to complete
	_, err = io.Copy(io.Discard, reader)
	return err
}

// TranscribeFile transcribes an audio/video file using Docker container.
func (dt *DockerTranscriber) TranscribeFile(ctx context.Context, filePath string) ([]byte, error) {
	if err := dt.EnsureImage(ctx); err != nil {
		return nil, fmt.Errorf("failed to ensure image: %w", err)
	}

	// Get absolute path for mounting
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	fileDir := filepath.Dir(absPath)
	fileName := filepath.Base(absPath)
	outputFile := fileName + ".srt"

	// Create container configuration
	containerConfig := &container.Config{
		Image:      dt.config.Image,
		WorkingDir: dt.config.WorkDir,
		Cmd: []string{
			"whisper",
			fileName,
			"--model", dt.config.Model,
			"--output_format", "srt",
			"--device", dt.config.Device,
		},
	}

	// Add language if specified
	if dt.config.Language != "" {
		containerConfig.Cmd = append(containerConfig.Cmd, "--language", dt.config.Language)
	}

	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: fileDir,
				Target: dt.config.WorkDir,
			},
		},
		AutoRemove: true,
	}

	// Create container
	resp, err := dt.client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %w", err)
	}

	// Start container
	if err := dt.client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}

	// Wait for container to finish with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, dt.config.Timeout)
	defer cancel()

	statusCh, errCh := dt.client.ContainerWait(timeoutCtx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return nil, fmt.Errorf("error waiting for container: %w", err)
		}
	case status := <-statusCh:
		if status.StatusCode != 0 {
			// Get container logs for debugging
			logs, _ := dt.getContainerLogs(ctx, resp.ID)
			return nil, fmt.Errorf("container exited with non-zero status %d: %s", status.StatusCode, logs)
		}
	case <-timeoutCtx.Done():
		// Kill the container if it's still running
		dt.client.ContainerKill(ctx, resp.ID, "KILL")
		return nil, fmt.Errorf("transcription timed out after %v", dt.config.Timeout)
	}

	// Read the output file
	outputPath := filepath.Join(fileDir, outputFile)
	content, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read output file %s: %w", outputPath, err)
	}

	// Clean up output file
	os.Remove(outputPath)

	return content, nil
}

// getContainerLogs retrieves logs from a container for debugging.
func (dt *DockerTranscriber) getContainerLogs(ctx context.Context, containerID string) (string, error) {
	options := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	}

	reader, err := dt.client.ContainerLogs(ctx, containerID, options)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	logs, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(logs), nil
}

// Close closes the Docker client connection.
func (dt *DockerTranscriber) Close() error {
	if dt.client != nil {
		return dt.client.Close()
	}
	return nil
}

// StopContainer stops a running container by name.
func (dt *DockerTranscriber) StopContainer(ctx context.Context, containerName string) error {
	return dt.client.ContainerStop(ctx, containerName, container.StopOptions{})
}

// ListContainers returns a list of containers with the Whisper image.
func (dt *DockerTranscriber) ListContainers(ctx context.Context) ([]types.Container, error) {
	return dt.client.ContainerList(ctx, container.ListOptions{All: true})
}