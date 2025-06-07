package scanner

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

type fakeProvider struct{ data []byte }

func (f fakeProvider) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
	return f.data, nil
}

func TestScanDirectory(t *testing.T) {
	dir := t.TempDir()
	vid := filepath.Join(dir, "movie.mkv")
	if err := os.WriteFile(vid, []byte("x"), 0644); err != nil {
		t.Fatalf("create video: %v", err)
	}
	// first scan creates subtitle
	if err := ScanDirectory(context.Background(), dir, "en", fakeProvider{[]byte("a")}, false, 2); err != nil {
		t.Fatalf("scan: %v", err)
	}
	sub := filepath.Join(dir, "movie.en.srt")
	data, err := os.ReadFile(sub)
	if err != nil {
		t.Fatalf("read subtitle: %v", err)
	}
	if string(data) != "a" {
		t.Fatalf("unexpected subtitle %q", data)
	}
	// second scan without upgrade should keep existing subtitle
	if err := ScanDirectory(context.Background(), dir, "en", fakeProvider{[]byte("b")}, false, 2); err != nil {
		t.Fatalf("scan 2: %v", err)
	}
	data, _ = os.ReadFile(sub)
	if string(data) != "a" {
		t.Fatalf("subtitle overwritten without upgrade")
	}
	// scan with upgrade should replace subtitle
	if err := ScanDirectory(context.Background(), dir, "en", fakeProvider{[]byte("c")}, true, 2); err != nil {
		t.Fatalf("scan upgrade: %v", err)
	}
	data, _ = os.ReadFile(sub)
	if string(data) != "c" {
		t.Fatalf("subtitle not upgraded: %q", data)
	}
}
