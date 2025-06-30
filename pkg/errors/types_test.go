// file: pkg/errors/types_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174005

package errors

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestNewAppError(t *testing.T) {
	tests := []struct {
		name           string
		code           ErrorCode
		message        string
		userMsg        string
		err            error
		expectRetryable bool
		expectStatus   int
	}{
		{
			name:           "provider timeout",
			code:           CodeProviderTimeout,
			message:        "Provider timeout",
			userMsg:        "Service temporarily unavailable",
			err:            errors.New("timeout"),
			expectRetryable: true,
			expectStatus:   http.StatusServiceUnavailable,
		},
		{
			name:           "validation error",
			code:           CodeValidationInput,
			message:        "Invalid input",
			userMsg:        "Please check your input",
			err:            nil,
			expectRetryable: false,
			expectStatus:   http.StatusBadRequest,
		},
		{
			name:           "system database error",
			code:           CodeSystemDatabase,
			message:        "Database connection failed",
			userMsg:        "Database error occurred",
			err:            errors.New("connection refused"),
			expectRetryable: false,
			expectStatus:   http.StatusInternalServerError,
		},
		{
			name:           "user not found",
			code:           CodeUserNotFound,
			message:        "User not found",
			userMsg:        "User does not exist",
			err:            nil,
			expectRetryable: false,
			expectStatus:   http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appErr := NewAppError(tt.code, tt.message, tt.userMsg, tt.err)

			if appErr.Code != tt.code {
				t.Errorf("Expected code %s, got %s", tt.code, appErr.Code)
			}

			if appErr.Message != tt.message {
				t.Errorf("Expected message %s, got %s", tt.message, appErr.Message)
			}

			if appErr.UserMsg != tt.userMsg {
				t.Errorf("Expected user message %s, got %s", tt.userMsg, appErr.UserMsg)
			}

			if appErr.Err != tt.err {
				t.Errorf("Expected underlying error %v, got %v", tt.err, appErr.Err)
			}

			if appErr.Retryable != tt.expectRetryable {
				t.Errorf("Expected retryable %v, got %v", tt.expectRetryable, appErr.Retryable)
			}

			if appErr.StatusCode != tt.expectStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectStatus, appErr.StatusCode)
			}

			if appErr.Timestamp.IsZero() {
				t.Error("Expected timestamp to be set")
			}
		})
	}
}

func TestAppError_Error(t *testing.T) {
	tests := []struct {
		name     string
		appErr   *AppError
		expected string
	}{
		{
			name: "with underlying error",
			appErr: &AppError{
				Code:    CodeProviderTimeout,
				Message: "Timeout occurred",
				Err:     errors.New("original error"),
			},
			expected: "PROVIDER_TIMEOUT: Timeout occurred (caused by: original error)",
		},
		{
			name: "without underlying error",
			appErr: &AppError{
				Code:    CodeValidationInput,
				Message: "Invalid input provided",
			},
			expected: "VALIDATION_INPUT: Invalid input provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.appErr.Error()
			if result != tt.expected {
				t.Errorf("Expected error string %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestAppError_WithContext(t *testing.T) {
	appErr := NewAppError(CodeProviderTimeout, "Test error", "User message", nil)
	
	appErr.WithContext("request_id", "12345")
	appErr.WithContext("user_id", 42)

	if appErr.Context["request_id"] != "12345" {
		t.Errorf("Expected request_id to be '12345', got %v", appErr.Context["request_id"])
	}

	if appErr.Context["user_id"] != 42 {
		t.Errorf("Expected user_id to be 42, got %v", appErr.Context["user_id"])
	}
}

func TestWrapError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		code     ErrorCode
		message  string
		userMsg  string
		expected *AppError
	}{
		{
			name:     "nil error",
			err:      nil,
			code:     CodeSystemInternal,
			message:  "Test",
			userMsg:  "Test user",
			expected: nil,
		},
		{
			name:    "existing AppError",
			err:     NewAppError(CodeProviderTimeout, "Original", "Original user", nil),
			code:    CodeSystemInternal,
			message: "Test",
			userMsg: "Test user",
			expected: &AppError{
				Code:    CodeProviderTimeout,
				Message: "Original",
				UserMsg: "Original user",
			},
		},
		{
			name:    "regular error",
			err:     errors.New("regular error"),
			code:    CodeSystemInternal,
			message: "Wrapped error",
			userMsg: "Something went wrong",
			expected: &AppError{
				Code:    CodeSystemInternal,
				Message: "Wrapped error",
				UserMsg: "Something went wrong",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WrapError(tt.err, tt.code, tt.message, tt.userMsg)
			
			if tt.expected == nil {
				if result != nil {
					t.Errorf("Expected nil result, got %v", result)
				}
				return
			}

			if result == nil {
				t.Fatal("Expected non-nil result")
			}

			if result.Code != tt.expected.Code {
				t.Errorf("Expected code %s, got %s", tt.expected.Code, result.Code)
			}

			if result.Message != tt.expected.Message {
				t.Errorf("Expected message %s, got %s", tt.expected.Message, result.Message)
			}

			if result.UserMsg != tt.expected.UserMsg {
				t.Errorf("Expected user message %s, got %s", tt.expected.UserMsg, result.UserMsg)
			}
		})
	}
}

func TestAppError_Unwrap(t *testing.T) {
	originalErr := errors.New("original error")
	appErr := NewAppError(CodeSystemInternal, "Wrapped", "User message", originalErr)

	unwrapped := appErr.Unwrap()
	if unwrapped != originalErr {
		t.Errorf("Expected unwrapped error to be %v, got %v", originalErr, unwrapped)
	}

	// Test with errors.Is
	if !errors.Is(appErr, originalErr) {
		t.Error("Expected errors.Is to work with wrapped error")
	}
}

func TestNewErrorResponse(t *testing.T) {
	appErr := NewAppError(CodeValidationInput, "Test error", "User error", nil)
	response := NewErrorResponse(appErr)

	if response.Error != appErr {
		t.Error("Expected response to contain the app error")
	}

	if response.StatusCode != appErr.StatusCode {
		t.Errorf("Expected status code %d, got %d", appErr.StatusCode, response.StatusCode)
	}

	if response.Data != nil {
		t.Error("Expected data to be nil for error response")
	}
}

func TestNewSuccessResponse(t *testing.T) {
	data := map[string]string{"key": "value"}
	response := NewSuccessResponse(data)

	if response.Error != nil {
		t.Error("Expected error to be nil for success response")
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	if response.Data == nil {
		t.Error("Expected response to contain data")
	}
	
	// Type check the data
	if dataMap, ok := response.Data.(map[string]string); !ok {
		t.Error("Expected data to be a map[string]string")
	} else if dataMap["key"] != "value" {
		t.Error("Expected data to contain the correct key-value pair")
	}
}

func TestDefaultRetryConfig(t *testing.T) {
	config := DefaultRetryConfig()

	if config.MaxAttempts != 3 {
		t.Errorf("Expected max attempts to be 3, got %d", config.MaxAttempts)
	}

	if config.InitialDelay != time.Second {
		t.Errorf("Expected initial delay to be 1s, got %v", config.InitialDelay)
	}

	if config.MaxDelay != 30*time.Second {
		t.Errorf("Expected max delay to be 30s, got %v", config.MaxDelay)
	}

	if config.BackoffFactor != 2.0 {
		t.Errorf("Expected backoff factor to be 2.0, got %f", config.BackoffFactor)
	}
}