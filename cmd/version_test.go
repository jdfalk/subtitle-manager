package cmd

import (
	"io"
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestSetGetVersionInfo(t *testing.T) {
	SetVersionInfo("1.2.3", "2024-01-02", "abcdef")
	v, _, _ := GetVersionInfo()
	if v != "1.2.3" {
		t.Fatalf("unexpected version %q", v)
	}
}

func TestPrintVersion(t *testing.T) {
	SetVersionInfo("1.2.3", "2024-01-02", "abcdef")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	orig := os.Stdout
	os.Stdout = w
	printVersion()
	w.Close()
	os.Stdout = orig
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	s := string(out)
	if !strings.Contains(s, "Subtitle Manager 1.2.3") {
		t.Errorf("version missing: %s", s)
	}
	if !strings.Contains(s, "Build Time: 2024-01-02") {
		t.Errorf("build time missing: %s", s)
	}
	if !strings.Contains(s, "Git Commit: abcdef") {
		t.Errorf("commit missing: %s", s)
	}
	if !strings.Contains(s, runtime.Version()) {
		t.Errorf("runtime missing: %s", s)
	}
}
