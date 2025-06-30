// file: pkg/errors/retry.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174003

package errors

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/sirupsen/logrus"
)

// RetryFunc represents a function that can be retried.
type RetryFunc func(ctx context.Context) (any, error)

// RetryResult represents the result of a retry operation.
type RetryResult struct {
	Value     any
	Err       error
	Attempts  int
	Duration  time.Duration
	Succeeded bool
}

// Retrier provides retry functionality with exponential backoff.
type Retrier struct {
	config *RetryConfig
	logger *logrus.Entry
}

// NewRetrier creates a new retrier with the given configuration.
func NewRetrier(config *RetryConfig) *Retrier {
	if config == nil {
		config = DefaultRetryConfig()
	}
	
	return &Retrier{
		config: config,
		logger: logging.GetLogger("retry"),
	}
}

// Retry executes the given function with retry logic and exponential backoff.
func (r *Retrier) Retry(ctx context.Context, operation string, fn RetryFunc) *RetryResult {
	startTime := time.Now()
	var lastErr error
	var result any
	
	for attempt := 1; attempt <= r.config.MaxAttempts; attempt++ {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return &RetryResult{
				Value:     result,
				Err:       fmt.Errorf("retry cancelled: %w", ctx.Err()),
				Attempts:  attempt - 1,
				Duration:  time.Since(startTime),
				Succeeded: false,
			}
		default:
		}
		
		// Execute the function
		value, err := fn(ctx)
		if err == nil {
			// Success
			r.logger.WithFields(map[string]interface{}{
				"operation": operation,
				"attempts":  attempt,
				"duration":  time.Since(startTime),
			}).Info("Operation succeeded after retry")
			
			return &RetryResult{
				Value:     value,
				Err:       nil,
				Attempts:  attempt,
				Duration:  time.Since(startTime),
				Succeeded: true,
			}
		}
		
		lastErr = err
		
		// Check if error is retryable
		if !r.shouldRetry(err) {
			r.logger.WithFields(map[string]interface{}{
				"operation": operation,
				"attempts":  attempt,
				"error":     err.Error(),
				"duration":  time.Since(startTime),
			}).Warn("Operation failed with non-retryable error")
			
			return &RetryResult{
				Value:     result,
				Err:       err,
				Attempts:  attempt,
				Duration:  time.Since(startTime),
				Succeeded: false,
			}
		}
		
		// If this was the last attempt, don't wait
		if attempt == r.config.MaxAttempts {
			break
		}
		
		// Calculate backoff delay
		delay := r.calculateDelay(attempt)
		
		r.logger.WithFields(map[string]interface{}{
			"operation": operation,
			"attempt":   attempt,
			"error":     err.Error(),
			"delay":     delay,
		}).Warn("Operation failed, retrying")
		
		// Wait before retrying
		select {
		case <-ctx.Done():
			return &RetryResult{
				Value:     result,
				Err:       fmt.Errorf("retry cancelled during backoff: %w", ctx.Err()),
				Attempts:  attempt,
				Duration:  time.Since(startTime),
				Succeeded: false,
			}
		case <-time.After(delay):
		}
	}
	
	// All attempts failed
	r.logger.WithFields(map[string]interface{}{
		"operation":     operation,
		"total_attempts": r.config.MaxAttempts,
		"final_error":   lastErr.Error(),
		"duration":      time.Since(startTime),
	}).Error("Operation failed after all retry attempts")
	
	return &RetryResult{
		Value:     result,
		Err:       fmt.Errorf("operation failed after %d attempts: %w", r.config.MaxAttempts, lastErr),
		Attempts:  r.config.MaxAttempts,
		Duration:  time.Since(startTime),
		Succeeded: false,
	}
}

// shouldRetry determines if an error should trigger a retry.
func (r *Retrier) shouldRetry(err error) bool {
	// Check if it's an AppError
	if appErr, ok := err.(*AppError); ok {
		return appErr.IsRetryable()
	}
	
	// For non-AppErrors, be conservative and don't retry by default
	return false
}

// calculateDelay calculates the delay for the given attempt using exponential backoff.
func (r *Retrier) calculateDelay(attempt int) time.Duration {
	// Calculate exponential backoff: initialDelay * (backoffFactor ^ (attempt - 1))
	delay := float64(r.config.InitialDelay) * math.Pow(r.config.BackoffFactor, float64(attempt-1))
	
	// Cap at maximum delay
	if delay > float64(r.config.MaxDelay) {
		delay = float64(r.config.MaxDelay)
	}
	
	return time.Duration(delay)
}

// CircuitBreaker provides circuit breaker functionality to prevent cascading failures.
type CircuitBreaker struct {
	maxFailures    int
	resetTimeout   time.Duration
	failureCount   int
	lastFailureTime time.Time
	state          CircuitState
	logger         *logrus.Entry
}

// CircuitState represents the state of a circuit breaker.
type CircuitState int

const (
	// CircuitClosed - normal operation
	CircuitClosed CircuitState = iota
	// CircuitOpen - circuit is open, requests are failing fast
	CircuitOpen
	// CircuitHalfOpen - testing if service has recovered
	CircuitHalfOpen
)

// String returns the string representation of the circuit state.
func (s CircuitState) String() string {
	switch s {
	case CircuitClosed:
		return "closed"
	case CircuitOpen:
		return "open"
	case CircuitHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

// NewCircuitBreaker creates a new circuit breaker.
func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
		state:        CircuitClosed,
		logger:       logging.GetLogger("circuit-breaker"),
	}
}

// Execute runs the given function through the circuit breaker.
func (cb *CircuitBreaker) Execute(ctx context.Context, operation string, fn RetryFunc) (any, error) {
	// Check circuit state
	switch cb.state {
	case CircuitOpen:
		if time.Since(cb.lastFailureTime) > cb.resetTimeout {
			// Time to try half-open
			cb.state = CircuitHalfOpen
			cb.logger.WithFields(map[string]interface{}{
				"operation": operation,
				"state":     cb.state.String(),
			}).Info("Circuit breaker transitioning to half-open")
		} else {
			// Fail fast
			return nil, NewAppError(
				CodeProviderUnavailable,
				"Circuit breaker is open",
				"Service is temporarily unavailable",
				nil,
			)
		}
		
	case CircuitHalfOpen:
		// Allow one request through
		
	case CircuitClosed:
		// Normal operation
	}
	
	// Execute the function
	result, err := fn(ctx)
	
	if err != nil {
		cb.recordFailure(operation)
		return nil, err
	}
	
	cb.recordSuccess(operation)
	return result, nil
}

// recordFailure records a failure and potentially opens the circuit.
func (cb *CircuitBreaker) recordFailure(operation string) {
	cb.failureCount++
	cb.lastFailureTime = time.Now()
	
	if cb.failureCount >= cb.maxFailures && cb.state == CircuitClosed {
		cb.state = CircuitOpen
		cb.logger.WithFields(map[string]interface{}{
			"operation":     operation,
			"failure_count": cb.failureCount,
			"state":         cb.state.String(),
		}).Warn("Circuit breaker opened due to failures")
	}
}

// recordSuccess records a success and potentially closes the circuit.
func (cb *CircuitBreaker) recordSuccess(operation string) {
	if cb.state == CircuitHalfOpen {
		cb.state = CircuitClosed
		cb.failureCount = 0
		cb.logger.WithFields(map[string]interface{}{
			"operation": operation,
			"state":     cb.state.String(),
		}).Info("Circuit breaker closed after successful test")
	} else if cb.state == CircuitClosed {
		// Reset failure count on success
		cb.failureCount = 0
	}
}

// GetState returns the current state of the circuit breaker.
func (cb *CircuitBreaker) GetState() CircuitState {
	return cb.state
}

// DefaultRetrier is a global retrier instance with default configuration.
var DefaultRetrier = NewRetrier(DefaultRetryConfig())

// WithRetry is a convenience function that uses the default retrier.
func WithRetry(ctx context.Context, operation string, fn RetryFunc) *RetryResult {
	return DefaultRetrier.Retry(ctx, operation, fn)
}