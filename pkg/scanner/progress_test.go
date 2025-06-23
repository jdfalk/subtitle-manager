// file: pkg/scanner/progress_test.go
package scanner

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	providersmocks "github.com/jdfalk/subtitle-manager/pkg/providers/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
)

// TestScanDirectoryProgress verifies that the callback is invoked for each file.
func TestScanDirectoryProgress(t *testing.T) {
	dir := t.TempDir()
	vid := filepath.Join(dir, "movie.mkv")
	if err := os.WriteFile(vid, []byte("x"), 0644); err != nil {
		t.Fatalf("create video: %v", err)
	}
	viper.Set("media_directory", dir)
	defer viper.Reset()
	var called int
	cb := func(string) { called++ }
	m := providersmocks.NewProvider(t)
	m.On("Fetch", mock.Anything, mock.Anything, "en").Return([]byte("a"), nil)
	err := ScanDirectoryProgress(context.Background(), dir, "en", "test", m, false, 1, nil, cb)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	m.AssertExpectations(t)
	if called != 1 {
		t.Fatalf("callback not called")
	}
}
