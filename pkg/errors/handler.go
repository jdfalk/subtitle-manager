// file: pkg/errors/handler.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174002

package errors

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/sirupsen/logrus"
)

// DefaultErrorHandler provides the default implementation of ErrorHandler.
type DefaultErrorHandler struct {
	logger *logrus.Entry
}

// NewDefaultErrorHandler creates a new default error handler.
func NewDefaultErrorHandler(component string) *DefaultErrorHandler {
	return &DefaultErrorHandler{
		logger: logging.GetLogger(component),
	}
}

// Handle processes an error and returns an appropriate response.
// It logs the error with structured data and converts it to a standardized response.
func (h *DefaultErrorHandler) Handle(ctx context.Context, err error) *Response {
	if err == nil {
		return NewSuccessResponse(nil)
	}

	// Convert to AppError if needed
	var appErr *AppError
	if existingAppErr, ok := err.(*AppError); ok {
		appErr = existingAppErr
	} else {
		// Try to categorize unknown errors
		appErr = h.categorizeError(err)
	}

	// Add request context if available
	if reqID := ctx.Value("request_id"); reqID != nil {
		appErr.WithContext("request_id", reqID)
	}
	if userID := ctx.Value("user_id"); userID != nil {
		appErr.WithContext("user_id", userID)
	}

	// Log the error with structured data
	h.logError(ctx, appErr)

	// Track the error for monitoring
	Track(appErr)

	return NewErrorResponse(appErr)
}

// Recover handles panic recovery and returns an appropriate response.
func (h *DefaultErrorHandler) Recover(ctx context.Context, recovered interface{}) *Response {
	// Create stack trace
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	stackTrace := string(buf[:n])

	// Create error from panic
	var err error
	if e, ok := recovered.(error); ok {
		err = e
	} else {
		err = fmt.Errorf("panic: %v", recovered)
	}

	appErr := NewAppError(
		CodeSystemInternal,
		fmt.Sprintf("Internal panic: %v", recovered),
		"An unexpected error occurred. Please try again.",
		err,
	)

	appErr.WithContext("stack_trace", stackTrace)
	appErr.WithContext("panic_value", recovered)

	// Add request context if available
	if reqID := ctx.Value("request_id"); reqID != nil {
		appErr.WithContext("request_id", reqID)
	}

	// Log the panic with critical level
	h.logger.WithFields(map[string]interface{}{
		"error_code":  appErr.Code,
		"error_msg":   appErr.Message,
		"user_msg":    appErr.UserMsg,
		"retryable":   appErr.Retryable,
		"status_code": appErr.StatusCode,
		"panic_value": recovered,
		"stack_trace": stackTrace,
		"context":     appErr.Context,
	}).Error("Application panic recovered")

	// Track the error for monitoring
	Track(appErr)

	return NewErrorResponse(appErr)
}

// categorizeError attempts to classify unknown errors into appropriate categories.
func (h *DefaultErrorHandler) categorizeError(err error) *AppError {
	errStr := strings.ToLower(err.Error())

	// Network/timeout related errors
	if strings.Contains(errStr, "timeout") || strings.Contains(errStr, "deadline exceeded") {
		return WrapError(err, CodeNetworkTimeout, "Request timeout", "The request took too long to complete. Please try again.")
	}

	if strings.Contains(errStr, "connection refused") || strings.Contains(errStr, "unreachable") {
		return WrapError(err, CodeNetworkUnreachable, "Network unreachable", "Unable to connect to the service. Please check your connection.")
	}

	if strings.Contains(errStr, "dns") || strings.Contains(errStr, "no such host") {
		return WrapError(err, CodeNetworkDNS, "DNS resolution failed", "Unable to resolve the hostname.")
	}

	// Database related errors
	if strings.Contains(errStr, "database") || strings.Contains(errStr, "sql") {
		return WrapError(err, CodeSystemDatabase, "Database error", "A database error occurred. Please try again.")
	}

	// File I/O errors
	if strings.Contains(errStr, "no such file") || strings.Contains(errStr, "permission denied") {
		return WrapError(err, CodeSystemFileIO, "File operation failed", "Unable to access the requested file.")
	}

	// Authentication/authorization errors
	if strings.Contains(errStr, "unauthorized") || strings.Contains(errStr, "forbidden") {
		return WrapError(err, CodeAuthInvalid, "Authentication failed", "Invalid credentials provided.")
	}

	// Rate limiting
	if strings.Contains(errStr, "rate limit") || strings.Contains(errStr, "too many requests") {
		return WrapError(err, CodeProviderRateLimit, "Rate limit exceeded", "Too many requests. Please wait before trying again.")
	}

	// Default to internal error
	return WrapError(err, CodeSystemInternal, "Internal error", "An unexpected error occurred. Please try again.")
}

// logError logs the error with appropriate level and structured data.
func (h *DefaultErrorHandler) logError(ctx context.Context, appErr *AppError) {
	fields := map[string]interface{}{
		"error_code":  appErr.Code,
		"error_msg":   appErr.Message,
		"user_msg":    appErr.UserMsg,
		"retryable":   appErr.Retryable,
		"status_code": appErr.StatusCode,
		"timestamp":   appErr.Timestamp,
	}

	// Add underlying error if present
	if appErr.Err != nil {
		fields["underlying_error"] = appErr.Err.Error()
	}

	// Add context information
	if appErr.Context != nil {
		for k, v := range appErr.Context {
			fields[fmt.Sprintf("ctx_%s", k)] = v
		}
	}

	// Choose log level based on error severity
	var logLevel string
	switch {
	case appErr.StatusCode >= 500:
		logLevel = "error"
	case appErr.StatusCode >= 400:
		logLevel = "warn"
	default:
		logLevel = "info"
	}

	// Log with appropriate level
	switch logLevel {
	case "error":
		h.logger.WithFields(fields).Error("Application error")
	case "warn":
		h.logger.WithFields(fields).Warn("Application warning")
	default:
		h.logger.WithFields(fields).Info("Application info")
	}
}

// Global default error handler instance
var defaultHandler *DefaultErrorHandler

// init initializes the default error handler
func init() {
	defaultHandler = NewDefaultErrorHandler("errors")
}

// Handle is a convenience function that uses the default error handler.
func Handle(ctx context.Context, err error) *Response {
	return defaultHandler.Handle(ctx, err)
}

// Recover is a convenience function that uses the default error handler.
func Recover(ctx context.Context, recovered interface{}) *Response {
	return defaultHandler.Recover(ctx, recovered)
}
