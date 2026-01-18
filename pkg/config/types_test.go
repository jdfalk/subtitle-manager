// file: pkg/config/types_test.go
// version: 1.0.0
// guid: 9fdf1b0d-4f8b-4b30-8bd1-d8784d8b53ff

package config

import (
	"strings"
	"testing"
)

func TestLogLevelString_KnownValues_ReturnName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		level    LogLevel
		expected string
	}{
		{name: "unspecified", level: LogLevel_UNSPECIFIED, expected: "UNSPECIFIED"},
		{name: "trace", level: LogLevel_TRACE, expected: "TRACE"},
		{name: "debug", level: LogLevel_DEBUG, expected: "DEBUG"},
		{name: "info", level: LogLevel_INFO, expected: "INFO"},
		{name: "warn", level: LogLevel_WARN, expected: "WARN"},
		{name: "error", level: LogLevel_ERROR, expected: "ERROR"},
		{name: "fatal", level: LogLevel_FATAL, expected: "FATAL"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange: use the provided log level.
			level := tt.level

			// Act: get the string representation.
			result := LogLevelString(level)

			// Assert: verify the expected name is returned.
			if result != tt.expected {
				t.Fatalf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestLogLevelString_UnknownValue_ReturnsFormatted(t *testing.T) {
	t.Parallel()

	// Arrange: use a log level that is not in the map.
	level := LogLevel(99)

	// Act: convert to string.
	result := LogLevelString(level)

	// Assert: verify fallback formatting.
	if result != "LogLevel(99)" {
		t.Fatalf("expected fallback LogLevel(99), got %q", result)
	}
}

func TestParseLogLevel_KnownValues_ReturnLevels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected LogLevel
	}{
		{name: "upper", input: "INFO", expected: LogLevel_INFO},
		{name: "lower", input: "warn", expected: LogLevel_WARN},
		{name: "mixed", input: "DeBuG", expected: LogLevel_DEBUG},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange: use the provided input.
			input := tt.input

			// Act: parse the input.
			result, err := ParseLogLevel(input)

			// Assert: ensure parsing succeeded and returns expected level.
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseLogLevel_InvalidValue_ReturnsError(t *testing.T) {
	t.Parallel()

	// Arrange: use an invalid log level string.
	input := "not-a-level"

	// Act: parse the input.
	result, err := ParseLogLevel(input)

	// Assert: ensure error and fallback to UNSPECIFIED.
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "unknown log level") {
		t.Fatalf("expected unknown log level error, got %v", err)
	}
	if result != LogLevel_UNSPECIFIED {
		t.Fatalf("expected %v, got %v", LogLevel_UNSPECIFIED, result)
	}
}

func TestAppenderTypeString_KnownValues_ReturnName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    AppenderType
		expected string
	}{
		{name: "unspecified", value: AppenderType_UNSPECIFIED, expected: "UNSPECIFIED"},
		{name: "console", value: AppenderType_CONSOLE, expected: "CONSOLE"},
		{name: "file", value: AppenderType_FILE, expected: "FILE"},
		{name: "syslog", value: AppenderType_SYSLOG, expected: "SYSLOG"},
		{name: "network", value: AppenderType_NETWORK, expected: "NETWORK"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange: use the provided appender type.
			value := tt.value

			// Act: convert to string.
			result := value.String()

			// Assert: verify expected name.
			if result != tt.expected {
				t.Fatalf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestAppenderTypeString_UnknownValue_ReturnsFormatted(t *testing.T) {
	t.Parallel()

	// Arrange: use an unknown value.
	value := AppenderType(99)

	// Act: convert to string.
	result := value.String()

	// Assert: ensure fallback formatting.
	if result != "AppenderType(99)" {
		t.Fatalf("expected fallback AppenderType(99), got %q", result)
	}
}

func TestParseAppenderType_KnownValues_ReturnTypes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected AppenderType
	}{
		{name: "upper", input: "CONSOLE", expected: AppenderType_CONSOLE},
		{name: "lower", input: "file", expected: AppenderType_FILE},
		{name: "mixed", input: "SySlOg", expected: AppenderType_SYSLOG},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange: use the provided input.
			input := tt.input

			// Act: parse the input.
			result, err := ParseAppenderType(input)

			// Assert: ensure parsing succeeded and returned expected type.
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseAppenderType_InvalidValue_ReturnsError(t *testing.T) {
	t.Parallel()

	// Arrange: use invalid input.
	input := "unknown"

	// Act: parse the input.
	result, err := ParseAppenderType(input)

	// Assert: ensure error and fallback to UNSPECIFIED.
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "unknown appender type") {
		t.Fatalf("expected unknown appender type error, got %v", err)
	}
	if result != AppenderType_UNSPECIFIED {
		t.Fatalf("expected %v, got %v", AppenderType_UNSPECIFIED, result)
	}
}
