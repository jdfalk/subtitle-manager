// file: pkg/errors/types.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174001

// Package errors provides standardized error handling, reporting, and recovery
// mechanisms for the subtitle manager application.
package errors

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ErrorCode represents a standardized error code for classification.
type ErrorCode string

// Standard error codes for different categories of errors.
const (
	// Provider errors (retry-able)
	CodeProviderTimeout     ErrorCode = "PROVIDER_TIMEOUT"
	CodeProviderUnavailable ErrorCode = "PROVIDER_UNAVAILABLE"
	CodeProviderRateLimit   ErrorCode = "PROVIDER_RATE_LIMIT"
	CodeProviderAuth        ErrorCode = "PROVIDER_AUTH"

	// Network errors (retry-able)
	CodeNetworkTimeout     ErrorCode = "NETWORK_TIMEOUT"
	CodeNetworkUnreachable ErrorCode = "NETWORK_UNREACHABLE"
	CodeNetworkDNS         ErrorCode = "NETWORK_DNS"

	// Authentication errors
	CodeAuthInvalid   ErrorCode = "AUTH_INVALID"
	CodeAuthExpired   ErrorCode = "AUTH_EXPIRED"
	CodeAuthForbidden ErrorCode = "AUTH_FORBIDDEN"

	// Validation errors
	CodeValidationInput   ErrorCode = "VALIDATION_INPUT"
	CodeValidationFormat  ErrorCode = "VALIDATION_FORMAT"
	CodeValidationMissing ErrorCode = "VALIDATION_MISSING"

	// System errors
	CodeSystemDatabase ErrorCode = "SYSTEM_DATABASE"
	CodeSystemFileIO   ErrorCode = "SYSTEM_FILE_IO"
	CodeSystemMemory   ErrorCode = "SYSTEM_MEMORY"
	CodeSystemInternal ErrorCode = "SYSTEM_INTERNAL"

	// User errors
	CodeUserNotFound      ErrorCode = "USER_NOT_FOUND"
	CodeUserConflict      ErrorCode = "USER_CONFLICT"
	CodeUserQuotaExceeded ErrorCode = "USER_QUOTA_EXCEEDED"
)

// AppError represents a standardized application error with context and metadata.
type AppError struct {
	Code       ErrorCode              `json:"code"`
	Message    string                 `json:"message"`
	UserMsg    string                 `json:"user_message"`
	Err        error                  `json:"-"`
	Retryable  bool                   `json:"retryable"`
	StatusCode int                    `json:"status_code"`
	Timestamp  time.Time              `json:"timestamp"`
	Context    map[string]interface{} `json:"context,omitempty"`
}

// Error implements the error interface.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap allows errors.Is and errors.As to work with wrapped errors.
func (e *AppError) Unwrap() error {
	return e.Err
}

// IsRetryable returns whether this error should be retried.
func (e *AppError) IsRetryable() bool {
	return e.Retryable
}

// WithContext adds contextual information to the error.
func (e *AppError) WithContext(key string, value interface{}) *AppError {
	if e.Context == nil {
		e.Context = make(map[string]interface{})
	}
	e.Context[key] = value
	return e
}

// NewAppError creates a new standardized application error.
func NewAppError(code ErrorCode, message, userMsg string, err error) *AppError {
	appErr := &AppError{
		Code:      code,
		Message:   message,
		UserMsg:   userMsg,
		Err:       err,
		Timestamp: time.Now(),
	}

	// Set retryable and status code based on error category
	switch code {
	case CodeProviderTimeout, CodeProviderUnavailable, CodeProviderRateLimit:
		appErr.Retryable = true
		appErr.StatusCode = http.StatusServiceUnavailable

	case CodeNetworkTimeout, CodeNetworkUnreachable, CodeNetworkDNS:
		appErr.Retryable = true
		appErr.StatusCode = http.StatusServiceUnavailable

	case CodeProviderAuth, CodeAuthInvalid, CodeAuthExpired, CodeAuthForbidden:
		appErr.Retryable = false
		appErr.StatusCode = http.StatusUnauthorized

	case CodeValidationInput, CodeValidationFormat, CodeValidationMissing:
		appErr.Retryable = false
		appErr.StatusCode = http.StatusBadRequest

	case CodeSystemDatabase, CodeSystemFileIO, CodeSystemMemory, CodeSystemInternal:
		appErr.Retryable = false
		appErr.StatusCode = http.StatusInternalServerError

	case CodeUserNotFound:
		appErr.Retryable = false
		appErr.StatusCode = http.StatusNotFound

	case CodeUserConflict:
		appErr.Retryable = false
		appErr.StatusCode = http.StatusConflict

	case CodeUserQuotaExceeded:
		appErr.Retryable = false
		appErr.StatusCode = http.StatusTooManyRequests

	default:
		appErr.Retryable = false
		appErr.StatusCode = http.StatusInternalServerError
	}

	return appErr
}

// WrapError wraps an existing error as an AppError if it isn't already one.
func WrapError(err error, code ErrorCode, message, userMsg string) *AppError {
	if err == nil {
		return nil
	}

	// If it's already an AppError, return it
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}

	return NewAppError(code, message, userMsg, err)
}

// Response represents the structure returned by error handlers.
type Response struct {
	Error      *AppError              `json:"error,omitempty"`
	Data       interface{}            `json:"data,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	StatusCode int                    `json:"-"`
}

// NewErrorResponse creates a new error response.
func NewErrorResponse(err *AppError) *Response {
	return &Response{
		Error:      err,
		StatusCode: err.StatusCode,
	}
}

// NewSuccessResponse creates a new success response.
func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Data:       data,
		StatusCode: http.StatusOK,
	}
}

// ErrorHandler defines the interface for handling application errors.
type ErrorHandler interface {
	// Handle processes an error and returns an appropriate response.
	Handle(ctx context.Context, err error) *Response

	// Recover handles panic recovery and returns an appropriate response.
	Recover(ctx context.Context, recovered interface{}) *Response
}

// RetryConfig defines configuration for retry mechanisms.
type RetryConfig struct {
	MaxAttempts   int           `json:"max_attempts"`
	InitialDelay  time.Duration `json:"initial_delay"`
	MaxDelay      time.Duration `json:"max_delay"`
	BackoffFactor float64       `json:"backoff_factor"`
}

// DefaultRetryConfig returns sensible defaults for retry configuration.
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxAttempts:   3,
		InitialDelay:  time.Second,
		MaxDelay:      30 * time.Second,
		BackoffFactor: 2.0,
	}
}
