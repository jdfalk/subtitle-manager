package logging

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

// TestConfigure verifies that logs are written to the configured file.
func TestConfigure(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.log")
	viper.Set("log_file", file)

	Configure()
	logger := GetLogger("test")
	logger.Info("message")

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	if !strings.Contains(string(data), "message") {
		t.Fatalf("log entry missing")
	}
}
