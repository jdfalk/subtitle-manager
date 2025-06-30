// file: pkg/errors/handler_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174006

package errors

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestDefaultErrorHandler_Handle(t *testing.T) {
	handler := NewDefaultErrorHandler("test")

	tests := []struct {
		name           string
		err            error
		expectCode     ErrorCode
		expectStatus   int
		expectRetryable bool
	}{
		{
			name:           "nil error",
			err:            nil,
			expectCode:     "",
			expectStatus:   200,
			expectRetryable: false,
		},
		{
			name:           "existing AppError",
			err:            NewAppError(CodeProviderTimeout, "Test", "User test", nil),
			expectCode:     CodeProviderTimeout,
			expectStatus:   503,
			expectRetryable: true,
		},
		{
			name:           "timeout error",
			err:            errors.New("connection timeout"),
			expectCode:     CodeNetworkTimeout,
			expectStatus:   503,
			expectRetryable: true,
		},
		{
			name:           "connection refused error",
			err:            errors.New("connection refused"),
			expectCode:     CodeNetworkUnreachable,
			expectStatus:   503,
			expectRetryable: true,
		},
		{
			name:           "dns error",
			err:            errors.New("no such host"),
			expectCode:     CodeNetworkDNS,
			expectStatus:   503,
			expectRetryable: true,
		},
		{
			name:           "database error",
			err:            errors.New("database connection failed"),
			expectCode:     CodeSystemDatabase,
			expectStatus:   500,
			expectRetryable: false,
		},
		{
			name:           "file error",
			err:            errors.New("no such file or directory"),
			expectCode:     CodeSystemFileIO,
			expectStatus:   500,
			expectRetryable: false,
		},
		{
			name:           "unauthorized error",
			err:            errors.New("unauthorized access"),
			expectCode:     CodeAuthInvalid,
			expectStatus:   401,
			expectRetryable: false,
		},
		{
			name:           "rate limit error",
			err:            errors.New("rate limit exceeded"),
			expectCode:     CodeProviderRateLimit,
			expectStatus:   503,
			expectRetryable: true,
		},
		{
			name:           "unknown error",
			err:            errors.New("some unknown error"),
			expectCode:     CodeSystemInternal,
			expectStatus:   500,
			expectRetryable: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response := handler.Handle(ctx, tt.err)

			if tt.err == nil {
				// Success case
				if response.Error != nil {
					t.Error("Expected nil error for success case")
				}
				if response.StatusCode != tt.expectStatus {
					t.Errorf("Expected status code %d, got %d", tt.expectStatus, response.StatusCode)
				}
				return
			}

			// Error case
			if response.Error == nil {
				t.Fatal("Expected error in response")
			}

			if response.Error.Code != tt.expectCode {
				t.Errorf("Expected error code %s, got %s", tt.expectCode, response.Error.Code)
			}

			if response.Error.StatusCode != tt.expectStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectStatus, response.Error.StatusCode)
			}

			if response.Error.Retryable != tt.expectRetryable {
				t.Errorf("Expected retryable %v, got %v", tt.expectRetryable, response.Error.Retryable)
			}

			if response.StatusCode != tt.expectStatus {
				t.Errorf("Expected response status code %d, got %d", tt.expectStatus, response.StatusCode)
			}
		})
	}
}

func TestDefaultErrorHandler_HandleWithContext(t *testing.T) {
	handler := NewDefaultErrorHandler("test")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", "test-123")
	ctx = context.WithValue(ctx, "user_id", 42)

	err := errors.New("test error")
	response := handler.Handle(ctx, err)

	if response.Error == nil {
		t.Fatal("Expected error in response")
	}

	if response.Error.Context["request_id"] != "test-123" {
		t.Errorf("Expected request_id to be 'test-123', got %v", response.Error.Context["request_id"])
	}

	if response.Error.Context["user_id"] != 42 {
		t.Errorf("Expected user_id to be 42, got %v", response.Error.Context["user_id"])
	}
}

func TestDefaultErrorHandler_Recover(t *testing.T) {
	handler := NewDefaultErrorHandler("test")

	tests := []struct {
		name      string
		recovered interface{}
	}{
		{
			name:      "error panic",
			recovered: errors.New("panic error"),
		},
		{
			name:      "string panic",
			recovered: "panic string",
		},
		{
			name:      "nil panic",
			recovered: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response := handler.Recover(ctx, tt.recovered)

			if response.Error == nil {
				t.Fatal("Expected error in response")
			}

			if response.Error.Code != CodeSystemInternal {
				t.Errorf("Expected error code %s, got %s", CodeSystemInternal, response.Error.Code)
			}

			if response.Error.StatusCode != 500 {
				t.Errorf("Expected status code 500, got %d", response.Error.StatusCode)
			}

			if response.Error.Context["panic_value"] != tt.recovered {
				t.Errorf("Expected panic value %v, got %v", tt.recovered, response.Error.Context["panic_value"])
			}

			// Check that stack trace is included
			if response.Error.Context["stack_trace"] == nil {
				t.Error("Expected stack trace in context")
			}

			stackTrace, ok := response.Error.Context["stack_trace"].(string)
			if !ok || len(stackTrace) == 0 {
				t.Error("Expected non-empty stack trace string")
			}
		})
	}
}

func TestCategorizeError(t *testing.T) {
	handler := NewDefaultErrorHandler("test")

	tests := []struct {
		errorString string
		expectCode  ErrorCode
	}{
		{"timeout occurred", CodeNetworkTimeout},
		{"context deadline exceeded", CodeNetworkTimeout},
		{"connection refused", CodeNetworkUnreachable},
		{"network unreachable", CodeNetworkUnreachable},
		{"dns resolution failed", CodeNetworkDNS},
		{"no such host", CodeNetworkDNS},
		{"database error", CodeSystemDatabase},
		{"sql error", CodeSystemDatabase},
		{"no such file", CodeSystemFileIO},
		{"permission denied", CodeSystemFileIO},
		{"unauthorized", CodeAuthInvalid},
		{"forbidden", CodeAuthInvalid},
		{"rate limit", CodeProviderRateLimit},
		{"too many requests", CodeProviderRateLimit},
		{"unknown error", CodeSystemInternal},
	}

	for _, tt := range tests {
		t.Run(tt.errorString, func(t *testing.T) {
			err := errors.New(tt.errorString)
			appErr := handler.categorizeError(err)

			if appErr.Code != tt.expectCode {
				t.Errorf("Expected error code %s, got %s", tt.expectCode, appErr.Code)
			}

			if appErr.Err != err {
				t.Error("Expected underlying error to be preserved")
			}
		})
	}
}

func TestConvenienceFunctions(t *testing.T) {
	// Test the global convenience functions
	ctx := context.Background()
	
	// Test Handle function
	err := errors.New("test error")
	response := Handle(ctx, err)
	
	if response.Error == nil {
		t.Fatal("Expected error in response")
	}
	
	// Test Recover function
	recovered := "test panic"
	response = Recover(ctx, recovered)
	
	if response.Error == nil {
		t.Fatal("Expected error in response")
	}
	
	if response.Error.Code != CodeSystemInternal {
		t.Errorf("Expected error code %s, got %s", CodeSystemInternal, response.Error.Code)
	}
}

func TestErrorHandlerInterface(t *testing.T) {
	// Ensure DefaultErrorHandler implements ErrorHandler interface
	var _ ErrorHandler = &DefaultErrorHandler{}
	
	// Test that the interface methods work
	handler := NewDefaultErrorHandler("test")
	ctx := context.Background()
	
	// Test Handle method through interface
	var errorHandler ErrorHandler = handler
	response := errorHandler.Handle(ctx, errors.New("test"))
	if response.Error == nil {
		t.Error("Expected error in response")
	}
	
	// Test Recover method through interface
	response = errorHandler.Recover(ctx, "test panic")
	if response.Error == nil {
		t.Error("Expected error in response")
	}
}

func TestErrorMessageBuilding(t *testing.T) {
	// Test error message construction in categorizeError
	handler := NewDefaultErrorHandler("test")
	
	originalErr := errors.New("original error message")
	appErr := handler.categorizeError(originalErr)
	
	// Check that error message includes the original error
	errorStr := appErr.Error()
	if !strings.Contains(errorStr, "original error message") {
		t.Errorf("Expected error string to contain original message, got: %s", errorStr)
	}
	
	// Check that it follows the expected format
	expectedFormat := fmt.Sprintf("%s:", appErr.Code)
	if !strings.HasPrefix(errorStr, expectedFormat) {
		t.Errorf("Expected error string to start with '%s', got: %s", expectedFormat, errorStr)
	}
}