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

func TestScanDirectory(t *testing.T) {
	dir := t.TempDir()
	vid := filepath.Join(dir, "movie.mkv")
	if err := os.WriteFile(vid, []byte("x"), 0644); err != nil {
		t.Fatalf("create video: %v", err)
	}
	viper.Set("media_directory", dir)
	defer viper.Reset()
	// first scan creates subtitle
	m := providersmocks.NewProvider(t)
	m.On("Fetch", mock.Anything, mock.Anything, "en").Return([]byte("a"), nil)
	if err := ScanDirectory(context.Background(), dir, "en", "test", m, false, 2, nil); err != nil {
		t.Fatalf("scan: %v", err)
	}
	m.AssertExpectations(t)
	sub := filepath.Join(dir, "movie.en.srt")
	data, err := os.ReadFile(sub)
	if err != nil {
		t.Fatalf("read subtitle: %v", err)
	}
	if string(data) != "a" {
		t.Fatalf("unexpected subtitle %q", data)
	}
	// second scan without upgrade should keep existing subtitle
	m2 := providersmocks.NewProvider(t)
	if err := ScanDirectory(context.Background(), dir, "en", "test", m2, false, 2, nil); err != nil {
		t.Fatalf("scan 2: %v", err)
	}
	data, _ = os.ReadFile(sub)
	if string(data) != "a" {
		t.Fatalf("subtitle overwritten without upgrade")
	}
	// scan with upgrade should replace subtitle
	m3 := providersmocks.NewProvider(t)
	m3.On("Fetch", mock.Anything, mock.Anything, "en").Return([]byte("cc"), nil)
	if err := ScanDirectory(context.Background(), dir, "en", "test", m3, true, 2, nil); err != nil {
		t.Fatalf("scan upgrade: %v", err)
	}
	m3.AssertExpectations(t)
	data, _ = os.ReadFile(sub)
	if string(data) != "cc" {
		t.Fatalf("subtitle not upgraded: %q", data)
	}
}

func TestScanDirectoryInvalidPath(t *testing.T) {
	err := ScanDirectory(context.Background(), "../invalid", "en", "test", nil, false, 1, nil)
	if err == nil {
		t.Fatalf("expected error for invalid path")
	}
}

func TestProcessFileInvalidPath(t *testing.T) {
	err := ProcessFile(context.Background(), "../bad/movie.mkv", "en", "test", nil, false, nil)
	if err == nil {
		t.Fatalf("expected error for invalid path")
	}
}

// TestProcessFile_UpgradeQuality ensures a subtitle is replaced only when the
// new version is larger, implying better quality.
func TestProcessFile_UpgradeQuality(t *testing.T) {
	dir := t.TempDir()
	vid := filepath.Join(dir, "movie.mkv")
	if err := os.WriteFile(vid, []byte("x"), 0644); err != nil {
		t.Fatalf("create video: %v", err)
	}
	sub := filepath.Join(dir, "movie.en.srt")
	if err := os.WriteFile(sub, []byte("existing subtitle"), 0644); err != nil {
		t.Fatalf("create sub: %v", err)
	}
	m := providersmocks.NewProvider(t)
	m.On("Fetch", mock.Anything, mock.Anything, "en").Return([]byte("a"), nil)
	if err := ProcessFile(context.Background(), vid, "en", "test", m, true, nil); err != nil {
		t.Fatalf("process: %v", err)
	}
	m.AssertExpectations(t)
	data, err := os.ReadFile(sub)
	if err != nil {
		t.Fatalf("read subtitle: %v", err)
	}
	if string(data) != "existing subtitle" {
		t.Fatalf("subtitle replaced with lower quality: %q", data)
	}
}
