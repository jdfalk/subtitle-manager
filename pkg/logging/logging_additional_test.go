// file: pkg/logging/logging_additional_test.go
// version: 1.0.0
// guid: 9b76c2c9-8c2f-4f63-8d54-65aa8b80fbad

package logging

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func resetLoggingState() {
	mu.Lock()
	loggers = map[string]*logrus.Logger{}
	mu.Unlock()
	output = os.Stdout
	logrus.SetOutput(output)
	viper.Reset()
}

func TestConfigureUsesStdoutWhenLogFileEmpty(t *testing.T) {
	// Arrange
	resetLoggingState()
	t.Cleanup(resetLoggingState)

	// Act
	Configure()

	// Assert
	if output != os.Stdout {
		t.Fatalf("expected stdout output when no log_file is set")
	}
}

func TestConfigureFallsBackWhenLogFileInvalid(t *testing.T) {
	// Arrange
	resetLoggingState()
	t.Cleanup(resetLoggingState)
	viper.Set("log_file", "relative/path.log")
	output = &bytes.Buffer{}

	// Act
	Configure()

	// Assert
	if output != os.Stdout {
		t.Fatalf("expected stdout output when log_file is invalid")
	}
}

func TestConfigureFallsBackWhenLogDirectoryCreationFails(t *testing.T) {
	// Arrange
	resetLoggingState()
	t.Cleanup(resetLoggingState)
	dir := t.TempDir()
	blockingFile := filepath.Join(dir, "blocking-file")
	if err := os.WriteFile(blockingFile, []byte("data"), 0o600); err != nil {
		t.Fatalf("write blocking file: %v", err)
	}
	viper.Set("log_file", filepath.Join(blockingFile, "log.txt"))
	output = &bytes.Buffer{}

	// Act
	Configure()

	// Assert
	if output != os.Stdout {
		t.Fatalf("expected stdout output when log directory cannot be created")
	}
}

func TestGetLoggerUsesComponentLevelAndCachesLogger(t *testing.T) {
	// Arrange
	resetLoggingState()
	t.Cleanup(resetLoggingState)
	viper.Set("log_levels.component", "debug")

	// Act
	first := GetLogger("component")
	second := GetLogger("component")

	// Assert
	if first.Logger.Level != logrus.DebugLevel {
		t.Fatalf("expected debug level, got %v", first.Logger.Level)
	}
	if first.Logger != second.Logger {
		t.Fatalf("expected cached logger instance")
	}
	if first.Data["component"] != "component" {
		t.Fatalf("expected component field to be set")
	}
}

func TestGetLoggerUsesInfoLevelOnInvalidComponentLevel(t *testing.T) {
	// Arrange
	resetLoggingState()
	t.Cleanup(resetLoggingState)
	viper.Set("log_levels.component", "not-a-level")
	viper.Set("log-level", "warn")

	// Act
	entry := GetLogger("component")

	// Assert
	if entry.Logger.Level != logrus.InfoLevel {
		t.Fatalf("expected info level on invalid component setting, got %v", entry.Logger.Level)
	}
}
