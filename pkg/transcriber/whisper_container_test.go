// file: pkg/transcriber/whisper_container_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174003

package transcriber

import (
	"testing"

	"github.com/spf13/viper"
)

func TestValidateModel(t *testing.T) {
	testCases := []struct {
		model    string
		expected bool
	}{
		{"tiny", true},
		{"base", true},
		{"small", true},
		{"medium", true},
		{"large", true},
		{"invalid", false},
		{"", false},
		{"TINY", false}, // case sensitive
	}

	for _, tc := range testCases {
		result := ValidateModel(tc.model)
		if result != tc.expected {
			t.Errorf("ValidateModel(%q) = %v, expected %v", tc.model, result, tc.expected)
		}
	}
}

func TestSetDefaultConfig(t *testing.T) {
	// Clear any existing config
	viper.Reset()

	// Set defaults
	SetDefaultConfig()

	// Check if defaults are set correctly
	if viper.GetString("whisper.container_name") != DefaultConfig.ContainerName {
		t.Errorf("Expected container_name %q, got %q", DefaultConfig.ContainerName, viper.GetString("whisper.container_name"))
	}

	if viper.GetString("whisper.image") != DefaultConfig.Image {
		t.Errorf("Expected image %q, got %q", DefaultConfig.Image, viper.GetString("whisper.image"))
	}

	if viper.GetString("whisper.port") != DefaultConfig.Port {
		t.Errorf("Expected port %q, got %q", DefaultConfig.Port, viper.GetString("whisper.port"))
	}

	if viper.GetString("whisper.model") != DefaultConfig.Model {
		t.Errorf("Expected model %q, got %q", DefaultConfig.Model, viper.GetString("whisper.model"))
	}

	if viper.GetString("whisper.device") != DefaultConfig.Device {
		t.Errorf("Expected device %q, got %q", DefaultConfig.Device, viper.GetString("whisper.device"))
	}

	if viper.GetBool("whisper.use_gpu") != DefaultConfig.UseGPU {
		t.Errorf("Expected use_gpu %v, got %v", DefaultConfig.UseGPU, viper.GetBool("whisper.use_gpu"))
	}
}

func TestLoadContainerConfig(t *testing.T) {
	// Clear any existing config
	viper.Reset()

	// Set test values
	viper.Set("whisper.container_name", "test-container")
	viper.Set("whisper.image", "test-image:latest")
	viper.Set("whisper.port", "8080")
	viper.Set("whisper.model", "small")
	viper.Set("whisper.device", "cpu")
	viper.Set("whisper.use_gpu", false)

	config := loadContainerConfig()

	if config.ContainerName != "test-container" {
		t.Errorf("Expected container_name %q, got %q", "test-container", config.ContainerName)
	}

	if config.Image != "test-image:latest" {
		t.Errorf("Expected image %q, got %q", "test-image:latest", config.Image)
	}

	if config.Port != "8080" {
		t.Errorf("Expected port %q, got %q", "8080", config.Port)
	}

	if config.Model != "small" {
		t.Errorf("Expected model %q, got %q", "small", config.Model)
	}

	if config.Device != "cpu" {
		t.Errorf("Expected device %q, got %q", "cpu", config.Device)
	}

	if config.UseGPU != false {
		t.Errorf("Expected use_gpu %v, got %v", false, config.UseGPU)
	}
}
