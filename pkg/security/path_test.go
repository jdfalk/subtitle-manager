// file: pkg/security/path_test.go
package security

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestValidateAndSanitizePath(t *testing.T) {
	dir := t.TempDir()
	viper.Set("media_directory", dir)
	defer viper.Reset()

	inside := filepath.Join(dir, "sub")
	if err := os.Mkdir(inside, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	p, err := ValidateAndSanitizePath(inside)
	if err != nil || p != inside {
		t.Fatalf("expected valid path, got %v %v", p, err)
	}

	if _, err := ValidateAndSanitizePath(filepath.Join(dir, "..", "other")); err == nil {
		t.Fatalf("expected traversal error")
	}

	if _, err := ValidateAndSanitizePath("relative/path"); err == nil {
		t.Fatalf("expected absolute path error")
	}

	if _, err := ValidateAndSanitizePath("/etc"); err == nil {
		t.Fatalf("expected disallowed path error")
	}
}
