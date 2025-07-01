// file: pkg/errors/retry_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174007

package errors

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRetrier(t *testing.T) {
	tests := []struct {
		name   string
		config *RetryConfig
	}{
		{
			name:   "with nil config uses default",
			config: nil,
		},
		{
			name: "with custom config",
			config: &RetryConfig{
				MaxAttempts:   5,
				InitialDelay:  100 * time.Millisecond,
				MaxDelay:      5 * time.Second,
				BackoffFactor: 2.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retrier := NewRetrier(tt.config)

			assert.NotNil(t, retrier)
			assert.NotNil(t, retrier.config)
			assert.NotNil(t, retrier.logger)

			if tt.config == nil {
				// Should use default config
				assert.NotNil(t, retrier.config)
			} else {
				assert.Equal(t, tt.config, retrier.config)
			}
		})
	}
}

func TestRetrier_Retry_Success(t *testing.T) {
	retrier := NewRetrier(&RetryConfig{
		MaxAttempts:   3,
		InitialDelay:  10 * time.Millisecond,
		MaxDelay:      1 * time.Second,
		BackoffFactor: 2.0,
	})

	// Function that succeeds on first try
	successFn := func(ctx context.Context) (any, error) {
		return "success", nil
	}

	result := retrier.Retry(context.Background(), "test_operation", successFn)

	assert.True(t, result.Succeeded)
	assert.NoError(t, result.Err)
	assert.Equal(t, "success", result.Value)
	assert.Equal(t, 1, result.Attempts)
	assert.Greater(t, result.Duration, time.Duration(0))
}

func TestRetrier_Retry_SuccessAfterRetries(t *testing.T) {
	retrier := NewRetrier(&RetryConfig{
		MaxAttempts:   3,
		InitialDelay:  1 * time.Millisecond,
		MaxDelay:      1 * time.Second,
		BackoffFactor: 2.0,
	})

	attempts := 0
	// Function that succeeds on third try
	retryFn := func(ctx context.Context) (any, error) {
		attempts++
		if attempts < 3 {
			return nil, NewAppError(CodeProviderTimeout, "temporary failure", "Try again", nil)
		}
		return "success_after_retries", nil
	}

	result := retrier.Retry(context.Background(), "test_operation", retryFn)

	assert.True(t, result.Succeeded)
	assert.NoError(t, result.Err)
	assert.Equal(t, "success_after_retries", result.Value)
	assert.Equal(t, 3, result.Attempts)
	assert.Equal(t, 3, attempts)
}

func TestRetrier_Retry_NonRetryableError(t *testing.T) {
	retrier := NewRetrier(&RetryConfig{
		MaxAttempts:   3,
		InitialDelay:  1 * time.Millisecond,
		MaxDelay:      1 * time.Second,
		BackoffFactor: 2.0,
	})

	attempts := 0
	// Function that fails with non-retryable error
	nonRetryableFn := func(ctx context.Context) (any, error) {
		attempts++
		return nil, NewAppError(CodeValidationInput, "validation error", "Bad input", nil)
	}

	result := retrier.Retry(context.Background(), "test_operation", nonRetryableFn)

	assert.False(t, result.Succeeded)
	assert.Error(t, result.Err)
	assert.Nil(t, result.Value)
	assert.Equal(t, 1, result.Attempts)
	assert.Equal(t, 1, attempts) // Should not retry
}

func TestRetrier_Retry_MaxAttemptsExceeded(t *testing.T) {
	retrier := NewRetrier(&RetryConfig{
		MaxAttempts:   3,
		InitialDelay:  1 * time.Millisecond,
		MaxDelay:      1 * time.Second,
		BackoffFactor: 2.0,
	})

	attempts := 0
	// Function that always fails with retryable error
	alwaysFailFn := func(ctx context.Context) (any, error) {
		attempts++
		return nil, NewAppError(CodeProviderTimeout, "always fails", "Service unavailable", nil)
	}

	result := retrier.Retry(context.Background(), "test_operation", alwaysFailFn)

	assert.False(t, result.Succeeded)
	assert.Error(t, result.Err)
	assert.Nil(t, result.Value)
	assert.Equal(t, 3, result.Attempts)
	assert.Equal(t, 3, attempts)
	assert.Contains(t, result.Err.Error(), "operation failed after 3 attempts")
}

func TestRetrier_shouldRetry(t *testing.T) {
	retrier := NewRetrier(nil)

	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "retryable app error",
			err:      NewAppError(CodeProviderTimeout, "timeout", "Try again", nil),
			expected: true,
		},
		{
			name:     "non-retryable app error",
			err:      NewAppError(CodeValidationInput, "validation", "Bad input", nil),
			expected: false,
		},
		{
			name:     "regular error",
			err:      errors.New("regular error"),
			expected: false,
		},
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := retrier.shouldRetry(tt.err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRetrier_calculateDelay(t *testing.T) {
	retrier := NewRetrier(&RetryConfig{
		MaxAttempts:   5,
		InitialDelay:  100 * time.Millisecond,
		MaxDelay:      5 * time.Second,
		BackoffFactor: 2.0,
	})

	tests := []struct {
		name            string
		attempt         int
		expectedMinimum time.Duration
		expectedMaximum time.Duration
	}{
		{
			name:            "first retry",
			attempt:         1,
			expectedMinimum: 100 * time.Millisecond,
			expectedMaximum: 100 * time.Millisecond,
		},
		{
			name:            "second retry",
			attempt:         2,
			expectedMinimum: 200 * time.Millisecond,
			expectedMaximum: 200 * time.Millisecond,
		},
		{
			name:            "third retry",
			attempt:         3,
			expectedMinimum: 400 * time.Millisecond,
			expectedMaximum: 400 * time.Millisecond,
		},
		{
			name:            "large attempt (should cap at max)",
			attempt:         10,
			expectedMinimum: 5 * time.Second,
			expectedMaximum: 5 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delay := retrier.calculateDelay(tt.attempt)
			assert.GreaterOrEqual(t, delay, tt.expectedMinimum)
			assert.LessOrEqual(t, delay, tt.expectedMaximum)
		})
	}
}

func TestRetryConfig_DefaultValues(t *testing.T) {
	config := DefaultRetryConfig()

	require.NotNil(t, config)
	assert.Greater(t, config.MaxAttempts, 0)
	assert.Greater(t, config.InitialDelay, time.Duration(0))
	assert.Greater(t, config.MaxDelay, config.InitialDelay)
	assert.Greater(t, config.BackoffFactor, 1.0)
}
