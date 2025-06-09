// file: pkg/scanner/progress_test.go
package scanner

import (
    "context"
    "os"
    "path/filepath"
    "testing"
)

// fakeProvider implements providers.Provider for testing.
type fakeProvider2 struct{ data []byte }
func (f fakeProvider2) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
    return f.data, nil
}

// TestScanDirectoryProgress verifies that the callback is invoked for each file.
func TestScanDirectoryProgress(t *testing.T) {
    dir := t.TempDir()
    vid := filepath.Join(dir, "movie.mkv")
    if err := os.WriteFile(vid, []byte("x"), 0644); err != nil {
        t.Fatalf("create video: %v", err)
    }
    var called int
    cb := func(string) { called++ }
    err := ScanDirectoryProgress(context.Background(), dir, "en", "test", fakeProvider2{[]byte("a")}, false, 1, nil, cb)
    if err != nil {
        t.Fatalf("scan: %v", err)
    }
    if called != 1 {
        t.Fatalf("callback not called")
    }
}
