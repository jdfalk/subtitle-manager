# Standardized Error Handling Documentation

## Overview

The subtitle manager application now implements a comprehensive, standardized
error handling system that provides consistent error reporting, monitoring, and
recovery mechanisms across all components.

## Key Components

### 1. Error Types (`pkg/errors/types.go`)

#### AppError Structure

```go
type AppError struct {
    Code       ErrorCode                  // Standardized error code
    Message    string                     // Technical error message
    UserMsg    string                     // User-friendly message
    Err        error                      // Underlying error (if any)
    Retryable  bool                       // Whether error can be retried
    StatusCode int                        // HTTP status code
    Timestamp  time.Time                  // When error occurred
    Context    map[string]interface{}     // Additional context data
}
```

#### Error Categories

- **Provider errors** (retry-able): `PROVIDER_TIMEOUT`, `PROVIDER_UNAVAILABLE`,
  `PROVIDER_RATE_LIMIT`, `PROVIDER_AUTH`
- **Network errors** (retry-able): `NETWORK_TIMEOUT`, `NETWORK_UNREACHABLE`,
  `NETWORK_DNS`
- **Authentication errors**: `AUTH_INVALID`, `AUTH_EXPIRED`, `AUTH_FORBIDDEN`
- **Validation errors**: `VALIDATION_INPUT`, `VALIDATION_FORMAT`,
  `VALIDATION_MISSING`
- **System errors**: `SYSTEM_DATABASE`, `SYSTEM_FILE_IO`, `SYSTEM_MEMORY`,
  `SYSTEM_INTERNAL`
- **User errors**: `USER_NOT_FOUND`, `USER_CONFLICT`, `USER_QUOTA_EXCEEDED`

### 2. Error Handler (`pkg/errors/handler.go`)

The `ErrorHandler` interface provides standardized error processing:

```go
type ErrorHandler interface {
    Handle(ctx context.Context, err error) *Response
    Recover(ctx context.Context, recovered interface{}) *Response
}
```

#### Features

- **Automatic categorization** of unknown errors based on error messages
- **Structured logging** with contextual information
- **Panic recovery** with stack trace capture
- **Request context preservation** (user ID, request ID, etc.)
- **Consistent response formatting**

### 3. Retry Mechanisms (`pkg/errors/retry.go`)

#### Retrier

Provides exponential backoff retry logic:

```go
type RetryConfig struct {
    MaxAttempts   int           // Maximum retry attempts (default: 3)
    InitialDelay  time.Duration // Initial delay (default: 1s)
    MaxDelay      time.Duration // Maximum delay (default: 30s)
    BackoffFactor float64       // Backoff multiplier (default: 2.0)
}
```

#### Circuit Breaker

Prevents cascading failures:

- **Closed**: Normal operation
- **Open**: Failing fast after max failures
- **Half-Open**: Testing if service has recovered

### 4. Error Tracking (`pkg/errors/tracker.go`)

#### Monitoring Dashboard

Provides HTTP endpoints for error monitoring:

- `/api/errors/stats` - Error statistics by type
- `/api/errors/recent` - Recent error events
- `/api/errors/top` - Most frequent errors
- `/api/errors/health` - Overall error health status

#### Error Aggregation

- Tracks error counts by type
- Maintains recent error events (configurable limit)
- Provides health status based on error frequency

## Usage Examples

### Basic Error Handling

```go
import "github.com/jdfalk/subtitle-manager/pkg/errors"

// Create a new error
err := errors.NewAppError(
    errors.CodeProviderTimeout,
    "Provider request timed out",
    "The subtitle service is temporarily unavailable",
    originalErr,
)

// Handle an error with context
ctx := context.WithValue(context.Background(), "user_id", 123)
response := errors.Handle(ctx, err)

// Use the response (contains structured error data)
if response.Error != nil {
    log.Printf("Error: %s (retryable: %v)", response.Error.Message, response.Error.Retryable)
}
```

### Error Wrapping

```go
// Wrap an existing error
wrappedErr := errors.WrapError(
    originalErr,
    errors.CodeSystemDatabase,
    "Database operation failed",
    "Unable to save your changes",
)
```

### Retry with Backoff

```go
// Simple retry
result := errors.WithRetry(ctx, "fetch-subtitles", func(ctx context.Context) (any, error) {
    return fetchSubtitles(movieID)
})

if result.Succeeded {
    subtitles := result.Value.([]Subtitle)
    // Use subtitles
}
```

### Circuit Breaker

```go
cb := errors.NewCircuitBreaker(5, 30*time.Second) // 5 failures, 30s reset

result, err := cb.Execute(ctx, "api-call", func(ctx context.Context) (any, error) {
    return callExternalAPI()
})
```

### Panic Recovery in HTTP Handlers

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if recovered := recover(); recovered != nil {
            response := errors.Recover(r.Context(), recovered)
            writeErrorResponse(w, response)
        }
    }()

    // Handler logic that might panic
    riskyOperation()
}
```

## Integration with Existing Systems

### Logging

Error handling integrates with the existing `pkg/logging` system:

- All errors are automatically logged with structured data
- Log levels are chosen based on error severity (error/warn/info)
- Context information is preserved in log entries

### Web UI

Error information is available through admin endpoints:

- Admins can monitor error trends and health
- Recent errors are displayed with full context
- Error statistics help identify problem areas

### Provider System

Integrates with existing provider backoff mechanisms:

- Provider errors trigger automatic backoff
- Circuit breakers prevent provider overload
- Retry logic respects provider-specific timeouts

## Configuration

### Default Settings

- **Retry attempts**: 3
- **Initial delay**: 1 second
- **Max delay**: 30 seconds
- **Backoff factor**: 2.0
- **Circuit breaker failures**: 5
- **Circuit breaker reset**: 30 seconds

### Customization

Create custom configurations as needed:

```go
config := &errors.RetryConfig{
    MaxAttempts:   5,
    InitialDelay:  500 * time.Millisecond,
    MaxDelay:      10 * time.Second,
    BackoffFactor: 1.5,
}
retrier := errors.NewRetrier(config)
```

## Error Tracking Dashboard

### API Endpoints

All endpoints require admin authentication (`basic` level):

#### GET `/api/errors/stats`

Returns error statistics by type:

```json
{
  "PROVIDER_TIMEOUT": {
    "code": "PROVIDER_TIMEOUT",
    "count": 15,
    "last_occurred": "2025-06-30T23:00:00Z",
    "first_seen": "2025-06-30T20:00:00Z",
    "message": "Provider timeout",
    "retryable": true
  }
}
```

#### GET `/api/errors/recent?limit=50`

Returns recent error events (newest first):

```json
[
  {
    "timestamp": "2025-06-30T23:00:00Z",
    "code": "PROVIDER_TIMEOUT",
    "message": "Provider timeout",
    "user_message": "Service temporarily unavailable",
    "context": {
      "request_id": "abc123",
      "user_id": 42
    }
  }
]
```

#### GET `/api/errors/top?limit=10`

Returns most frequent errors:

```json
[
  {
    "code": "PROVIDER_TIMEOUT",
    "count": 15,
    "last_occurred": "2025-06-30T23:00:00Z"
  }
]
```

#### GET `/api/errors/health`

Returns overall error health status:

```json
{
  "status": "healthy",
  "total_errors": 25,
  "critical_errors_recent": 2,
  "unique_error_types": 5,
  "timestamp": "2025-06-30T23:00:00Z"
}
```

## Best Practices

### Error Creation

1. Use appropriate error codes for categorization
2. Provide both technical and user-friendly messages
3. Include relevant context information
4. Preserve underlying errors when wrapping

### Error Handling

1. Use the global error handler for consistency
2. Include request context (user ID, request ID) when available
3. Log errors at appropriate levels
4. Return structured error responses

### Retry Logic

1. Only retry on retryable errors
2. Use exponential backoff to avoid overwhelming services
3. Respect circuit breaker states
4. Set reasonable timeout limits

### Monitoring

1. Monitor error rates and trends
2. Set up alerts for critical error thresholds
3. Review top errors regularly
4. Use error context for debugging

## Migration Guide

### Existing Error Handling

Existing error handling continues to work unchanged. The new system provides
additional standardization and monitoring capabilities.

### Gradual Adoption

1. **Start with new code** - Use standardized error types in new features
2. **Wrap existing errors** - Use `WrapError()` to standardize existing error
   handling
3. **Add retry logic** - Implement retry mechanisms for external service calls
4. **Monitor and adjust** - Use error tracking to identify areas for improvement

### Examples of Migration

```go
// Before
if err != nil {
    log.Printf("Error: %v", err)
    return err
}

// After
if err != nil {
    appErr := errors.WrapError(err, errors.CodeSystemInternal, "Operation failed", "Please try again")
    response := errors.Handle(ctx, appErr)
    return response.Error
}
```

## Testing

The error handling system includes comprehensive tests:

- Unit tests for all error types and handlers
- Integration tests for tracking and monitoring
- Circuit breaker and retry mechanism tests

Run tests with:

```bash
go test ./pkg/errors/ -v
```

## Future Enhancements

Potential improvements for the error handling system:

- **Metrics integration** - Export error metrics to Prometheus
- **Alert configuration** - Configurable error rate alerts
- **Error aggregation** - Group similar errors automatically
- **Recovery suggestions** - Automated suggestions for error resolution
- **Error correlation** - Link related errors across requests
