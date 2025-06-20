package scanner

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	providersmocks "github.com/jdfalk/subtitle-manager/pkg/providers/mocks"
	"github.com/stretchr/testify/mock"
)

func TestScanDirectory(t *testing.T) {
	dir := t.TempDir()
	vid := filepath.Join(dir, "movie.mkv")
	if err := os.WriteFile(vid, []byte("x"), 0644); err != nil {
		t.Fatalf("create video: %v", err)
	}
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
	m3.On("Fetch", mock.Anything, mock.Anything, "en").Return([]byte("c"), nil)
	if err := ScanDirectory(context.Background(), dir, "en", "test", m3, true, 2, nil); err != nil {
		t.Fatalf("scan upgrade: %v", err)
	}
	m3.AssertExpectations(t)
	data, _ = os.ReadFile(sub)
	if string(data) != "c" {
		t.Fatalf("subtitle not upgraded: %q", data)
	}
}
