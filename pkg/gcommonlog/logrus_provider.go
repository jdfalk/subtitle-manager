// file: pkg/gcommonlog/logrus_provider.go
// version: 1.0.0
// guid: a1e4bca3-27a2-4e4a-9d4c-1c2f3b4d5e6f

package gcommonlog

import (
	"context"
	"io"

	gclog "github.com/jdfalk/gcommon/pkg/log"
	"github.com/sirupsen/logrus"
)

type logrusProvider struct {
	entry *logrus.Entry
}

// NewLogrusProvider registers a logrus based logging provider compatible with gcommon.
func NewLogrusProvider(cfg gclog.Config) (gclog.Provider, error) {
	l := logrus.New()
	lvl, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	l.SetLevel(lvl)
	return &logrusProvider{entry: logrus.NewEntry(l)}, nil
}

func (p *logrusProvider) Name() string               { return "logrus" }
func (p *logrusProvider) Close() error               { return nil }
func (p *logrusProvider) Sync() error                { return nil }
func (p *logrusProvider) SetOutput(w io.Writer)      { p.entry.Logger.SetOutput(w) }
func (p *logrusProvider) AddHook(h gclog.Hook) error { return nil }

func (p *logrusProvider) log(level logrus.Level, msg string, fields []gclog.Field) {
	if !p.entry.Logger.IsLevelEnabled(level) {
		return
	}
	p.entry.WithFields(convert(fields)).Log(level, msg)
}

func (p *logrusProvider) Debug(msg string, f ...gclog.Field) { p.log(logrus.DebugLevel, msg, f) }
func (p *logrusProvider) DebugContext(ctx context.Context, msg string, f ...gclog.Field) {
	p.WithContext(ctx).Debug(msg, f...)
}
func (p *logrusProvider) Info(msg string, f ...gclog.Field) { p.log(logrus.InfoLevel, msg, f) }
func (p *logrusProvider) InfoContext(ctx context.Context, msg string, f ...gclog.Field) {
	p.WithContext(ctx).Info(msg, f...)
}
func (p *logrusProvider) Warn(msg string, f ...gclog.Field) { p.log(logrus.WarnLevel, msg, f) }
func (p *logrusProvider) WarnContext(ctx context.Context, msg string, f ...gclog.Field) {
	p.WithContext(ctx).Warn(msg, f...)
}
func (p *logrusProvider) Error(msg string, f ...gclog.Field) { p.log(logrus.ErrorLevel, msg, f) }
func (p *logrusProvider) ErrorContext(ctx context.Context, msg string, f ...gclog.Field) {
	p.WithContext(ctx).Error(msg, f...)
}
func (p *logrusProvider) Fatal(msg string, f ...gclog.Field) { p.log(logrus.FatalLevel, msg, f) }
func (p *logrusProvider) FatalContext(ctx context.Context, msg string, f ...gclog.Field) {
	p.WithContext(ctx).Fatal(msg, f...)
}

func (p *logrusProvider) With(fields ...gclog.Field) gclog.Logger {
	return &logrusProvider{entry: p.entry.WithFields(convert(fields))}
}
func (p *logrusProvider) WithContext(ctx context.Context) gclog.Logger {
	return &logrusProvider{entry: p.entry.WithContext(ctx)}
}
func (p *logrusProvider) SetLevel(l gclog.Level) { p.entry.Logger.SetLevel(toLogrusLevel(l)) }
func (p *logrusProvider) GetLevel() gclog.Level  { return fromLogrusLevel(p.entry.Logger.GetLevel()) }

func convert(fields []gclog.Field) logrus.Fields {
	out := make(logrus.Fields, len(fields))
	for _, f := range fields {
		out[f.Key] = f.Value
	}
	return out
}

func toLogrusLevel(l gclog.Level) logrus.Level {
	switch l {
	case gclog.DebugLevel:
		return logrus.DebugLevel
	case gclog.InfoLevel:
		return logrus.InfoLevel
	case gclog.WarnLevel:
		return logrus.WarnLevel
	case gclog.ErrorLevel:
		return logrus.ErrorLevel
	case gclog.FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

func fromLogrusLevel(l logrus.Level) gclog.Level {
	switch l {
	case logrus.DebugLevel:
		return gclog.DebugLevel
	case logrus.InfoLevel:
		return gclog.InfoLevel
	case logrus.WarnLevel:
		return gclog.WarnLevel
	case logrus.ErrorLevel:
		return gclog.ErrorLevel
	case logrus.FatalLevel, logrus.PanicLevel:
		return gclog.FatalLevel
	default:
		return gclog.InfoLevel
	}
}

func init() {
	gclog.RegisterProvider("logrus", NewLogrusProvider)
}
