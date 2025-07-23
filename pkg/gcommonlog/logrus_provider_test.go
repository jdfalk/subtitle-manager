// file: pkg/gcommonlog/logrus_provider_test.go
// version: 1.0.0
// guid: 3456b7c8-9d0e-4f12-8b34-5a6c7d8e9f01
package gcommonlog

import (
	"bytes"
	"strings"
	"testing"

	gclog "github.com/jdfalk/gcommon/pkg/log"
)

// TestLogrusProvider verifies that the provider writes log messages and fields.
func TestLogrusProvider(t *testing.T) {
	buf := &bytes.Buffer{}
	p, err := gclog.NewProvider(gclog.Config{Provider: "logrus", Level: "debug"})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}
	p.SetOutput(buf)
	p.Info("hello", gclog.Field{Key: "k", Value: "v"})

	out := buf.String()
	if !strings.Contains(out, "hello") || !strings.Contains(out, "k=v") {
		t.Errorf("unexpected output: %s", out)
	}
}

// TestWithField ensures With adds structured fields.
func TestWithField(t *testing.T) {
	buf := &bytes.Buffer{}
	p, err := gclog.NewProvider(gclog.Config{Provider: "logrus", Level: "info"})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}
	p.SetOutput(buf)

	logger := p.With(gclog.Field{Key: "module", Value: "test"})
	logger.Info("msg")

	out := buf.String()
	if !strings.Contains(out, "module=test") {
		t.Errorf("expected field not found: %s", out)
	}
}
