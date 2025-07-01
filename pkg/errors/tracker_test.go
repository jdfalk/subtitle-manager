// file: pkg/errors/tracker_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174008

package errors

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewErrorTracker(t *testing.T) {
	tracker := NewErrorTracker(10)

	assert.NotNil(t, tracker)
	assert.NotNil(t, tracker.stats)
	assert.NotNil(t, tracker.recent)
	assert.Equal(t, 10, tracker.maxRecentEvents)
	assert.NotNil(t, tracker.logger)
	assert.Equal(t, 0, len(tracker.stats))
	assert.Equal(t, 0, len(tracker.recent))
}

func TestErrorTracker_TrackError(t *testing.T) {
	tracker := NewErrorTracker(5)

	// Create test errors
	err1 := NewAppError(CodeProviderTimeout, "timeout error", "Service unavailable", nil)
	err2 := NewAppError(CodeValidationInput, "validation error", "Invalid input", nil)

	// Track first error
	tracker.TrackError(err1)

	stats := tracker.GetStats()
	assert.Len(t, stats, 1)
	assert.Contains(t, stats, CodeProviderTimeout)
	assert.Equal(t, 1, stats[CodeProviderTimeout].Count)
	assert.Equal(t, "timeout error", stats[CodeProviderTimeout].Message)
	assert.True(t, stats[CodeProviderTimeout].Retryable)

	// Track same error again
	tracker.TrackError(err1)

	stats = tracker.GetStats()
	assert.Len(t, stats, 1)
	assert.Equal(t, 2, stats[CodeProviderTimeout].Count)

	// Track different error
	tracker.TrackError(err2)

	stats = tracker.GetStats()
	assert.Len(t, stats, 2)
	assert.Contains(t, stats, CodeValidationInput)
	assert.Equal(t, 1, stats[CodeValidationInput].Count)
	assert.False(t, stats[CodeValidationInput].Retryable)
}

func TestErrorTracker_GetRecentEvents(t *testing.T) {
	tracker := NewErrorTracker(3)

	// Track multiple errors
	err1 := NewAppError(CodeProviderTimeout, "error 1", "Service unavailable", nil)
	err2 := NewAppError(CodeValidationInput, "error 2", "Invalid input", nil)
	err3 := NewAppError(CodeSystemDatabase, "error 3", "Database error", nil)
	err4 := NewAppError(CodeProviderRateLimit, "error 4", "Rate limited", nil)

	time.Sleep(1 * time.Millisecond) // Ensure different timestamps
	tracker.TrackError(err1)

	time.Sleep(1 * time.Millisecond)
	tracker.TrackError(err2)

	time.Sleep(1 * time.Millisecond)
	tracker.TrackError(err3)

	time.Sleep(1 * time.Millisecond)
	tracker.TrackError(err4) // This should push out err1 due to limit of 3

	recent := tracker.GetRecentEvents(10)
	assert.Len(t, recent, 3) // Limited by maxRecentEvents

	// Should be in reverse chronological order (newest first)
	assert.Equal(t, CodeProviderRateLimit, recent[0].Code)
	assert.Equal(t, CodeSystemDatabase, recent[1].Code)
	assert.Equal(t, CodeValidationInput, recent[2].Code)

	// Test with limit
	recent = tracker.GetRecentEvents(2)
	assert.Len(t, recent, 2)
	assert.Equal(t, CodeProviderRateLimit, recent[0].Code)
	assert.Equal(t, CodeSystemDatabase, recent[1].Code)
}

func TestErrorTracker_GetTopErrors(t *testing.T) {
	tracker := NewErrorTracker(10)

	err1 := NewAppError(CodeProviderTimeout, "timeout", "Service unavailable", nil)
	err2 := NewAppError(CodeValidationInput, "validation", "Invalid input", nil)
	err3 := NewAppError(CodeSystemDatabase, "database", "Database error", nil)

	// Track errors with different frequencies
	tracker.TrackError(err1) // 1 time

	tracker.TrackError(err2) // 3 times
	tracker.TrackError(err2)
	tracker.TrackError(err2)

	tracker.TrackError(err3) // 2 times
	tracker.TrackError(err3)

	topErrors := tracker.GetTopErrors(10)
	require.Len(t, topErrors, 3)

	// Should be sorted by count (descending)
	assert.Equal(t, CodeValidationInput, topErrors[0].Code)
	assert.Equal(t, 3, topErrors[0].Count)

	assert.Equal(t, CodeSystemDatabase, topErrors[1].Code)
	assert.Equal(t, 2, topErrors[1].Count)

	assert.Equal(t, CodeProviderTimeout, topErrors[2].Code)
	assert.Equal(t, 1, topErrors[2].Count)

	// Test with limit
	topErrors = tracker.GetTopErrors(2)
	assert.Len(t, topErrors, 2)
	assert.Equal(t, CodeValidationInput, topErrors[0].Code)
	assert.Equal(t, CodeSystemDatabase, topErrors[1].Code)
}

func TestErrorDashboard_StatsHandler(t *testing.T) {
	tracker := NewErrorTracker(5)
	dashboard := NewErrorDashboard(tracker)

	// Track some errors
	err1 := NewAppError(CodeProviderTimeout, "timeout", "Service unavailable", nil)
	err2 := NewAppError(CodeValidationInput, "validation", "Invalid input", nil)

	tracker.TrackError(err1)
	tracker.TrackError(err2)
	tracker.TrackError(err1) // err1 again

	// Test the HTTP handler
	handler := dashboard.StatsHandler()

	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err)

	// Check response structure
	assert.Len(t, response, 2) // Two different error codes
}

func TestErrorDashboard_RecentHandler(t *testing.T) {
	tracker := NewErrorTracker(5)
	dashboard := NewErrorDashboard(tracker)

	// Track an error
	err1 := NewAppError(CodeProviderTimeout, "timeout", "Service unavailable", nil)
	tracker.TrackError(err1)

	handler := dashboard.RecentHandler()

	req := httptest.NewRequest(http.MethodGet, "/recent", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	var events []ErrorEvent
	err := json.Unmarshal(recorder.Body.Bytes(), &events)
	require.NoError(t, err)
	assert.Len(t, events, 1)
}

func TestErrorTracker_ConcurrentAccess(t *testing.T) {
	tracker := NewErrorTracker(10)

	// Test concurrent tracking
	done := make(chan bool)

	// Start multiple goroutines tracking errors
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer func() { done <- true }()

			for j := 0; j < 10; j++ {
				err := NewAppError(CodeProviderTimeout, "concurrent error", "Test error", nil)
				tracker.TrackError(err)
			}
		}(i)
	}

	// Start goroutines reading stats
	for i := 0; i < 3; i++ {
		go func() {
			defer func() { done <- true }()

			for j := 0; j < 20; j++ {
				tracker.GetStats()
				tracker.GetRecentEvents(5)
				tracker.GetTopErrors(5)
			}
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 8; i++ { // 5 writers + 3 readers
		<-done
	}

	// Verify final state
	stats := tracker.GetStats()
	assert.Len(t, stats, 1) // Only one error code (all goroutines use same error)

	for _, stat := range stats {
		assert.Equal(t, 50, stat.Count) // 5 goroutines * 10 times each
	}
}

func TestErrorEvent_JSONSerialization(t *testing.T) {
	event := ErrorEvent{
		Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
		Code:      CodeProviderTimeout,
		Message:   "test message",
		UserMsg:   "test user message",
		Context: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	var decoded ErrorEvent
	err = json.Unmarshal(data, &decoded)
	require.NoError(t, err)

	assert.Equal(t, event.Code, decoded.Code)
	assert.Equal(t, event.Message, decoded.Message)
	assert.Equal(t, event.UserMsg, decoded.UserMsg)
	assert.Equal(t, event.Context, decoded.Context)
}

func TestErrorStats_JSONSerialization(t *testing.T) {
	stats := ErrorStats{
		Code:         CodeProviderTimeout,
		Count:        5,
		LastOccurred: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
		FirstSeen:    time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
		Message:      "test message",
		Retryable:    true,
	}

	data, err := json.Marshal(stats)
	require.NoError(t, err)

	var decoded ErrorStats
	err = json.Unmarshal(data, &decoded)
	require.NoError(t, err)

	assert.Equal(t, stats.Code, decoded.Code)
	assert.Equal(t, stats.Count, decoded.Count)
	assert.Equal(t, stats.Message, decoded.Message)
	assert.Equal(t, stats.Retryable, decoded.Retryable)
}
