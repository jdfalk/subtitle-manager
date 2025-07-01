// file: pkg/errors/integration_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174007

package errors

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestErrorTrackerIntegration(t *testing.T) {
	// Reset the global tracker for clean test
	GlobalTracker.Reset()

	// Create some test errors to track
	testErrors := []*AppError{
		NewAppError(CodeProviderTimeout, "Provider timeout", "Service unavailable", nil),
		NewAppError(CodeNetworkTimeout, "Network timeout", "Connection timed out", nil),
		NewAppError(CodeProviderTimeout, "Provider timeout again", "Service unavailable", nil),
		NewAppError(CodeValidationInput, "Invalid input", "Please check your input", nil),
	}

	// Track the errors
	for _, err := range testErrors {
		Track(err)
		time.Sleep(1 * time.Millisecond) // Small delay to ensure different timestamps
	}

	// Test stats endpoint
	dashboard := GetDashboard()

	// Test StatsHandler
	req := httptest.NewRequest("GET", "/api/errors/stats", nil)
	w := httptest.NewRecorder()
	dashboard.StatsHandler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var stats map[ErrorCode]*ErrorStats
	if err := json.NewDecoder(w.Body).Decode(&stats); err != nil {
		t.Fatalf("Failed to decode stats response: %v", err)
	}

	// Check that we have the expected error types
	if len(stats) != 3 {
		t.Errorf("Expected 3 error types, got %d", len(stats))
	}

	// Check provider timeout occurred twice
	if providerStats, exists := stats[CodeProviderTimeout]; !exists {
		t.Error("Expected provider timeout stats")
	} else if providerStats.Count != 2 {
		t.Errorf("Expected provider timeout count of 2, got %d", providerStats.Count)
	}

	// Test RecentHandler
	req = httptest.NewRequest("GET", "/api/errors/recent", nil)
	w = httptest.NewRecorder()
	dashboard.RecentHandler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var events []ErrorEvent
	if err := json.NewDecoder(w.Body).Decode(&events); err != nil {
		t.Fatalf("Failed to decode events response: %v", err)
	}

	if len(events) != 4 {
		t.Errorf("Expected 4 events, got %d", len(events))
	}

	// Events should be in reverse chronological order (newest first)
	for i := 1; i < len(events); i++ {
		if events[i].Timestamp.After(events[i-1].Timestamp) {
			t.Error("Events are not in reverse chronological order")
		}
	}

	// Test TopErrorsHandler
	req = httptest.NewRequest("GET", "/api/errors/top", nil)
	w = httptest.NewRecorder()
	dashboard.TopErrorsHandler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var topErrors []*ErrorStats
	if err := json.NewDecoder(w.Body).Decode(&topErrors); err != nil {
		t.Fatalf("Failed to decode top errors response: %v", err)
	}

	// Should be sorted by count descending
	if len(topErrors) > 0 && topErrors[0].Code != CodeProviderTimeout {
		t.Error("Expected provider timeout to be the top error")
	}

	// Test HealthHandler
	req = httptest.NewRequest("GET", "/api/errors/health", nil)
	w = httptest.NewRecorder()
	dashboard.HealthHandler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var health map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&health); err != nil {
		t.Fatalf("Failed to decode health response: %v", err)
	}

	// Check required fields
	requiredFields := []string{"total_errors", "critical_errors_recent", "unique_error_types", "status", "timestamp"}
	for _, field := range requiredFields {
		if _, exists := health[field]; !exists {
			t.Errorf("Expected health response to contain field: %s", field)
		}
	}

	if totalErrors, ok := health["total_errors"].(float64); !ok || totalErrors != 4 {
		t.Errorf("Expected total_errors to be 4, got %v", health["total_errors"])
	}
}

func TestErrorHandlerIntegration(t *testing.T) {
	// Test that the error handler properly tracks errors
	GlobalTracker.Reset()

	handler := NewDefaultErrorHandler("integration-test")
	ctx := context.Background()

	// Handle some errors
	testErr := NewAppError(CodeProviderTimeout, "Test timeout", "Test message", nil)
	response := handler.Handle(ctx, testErr)

	if response.Error == nil {
		t.Fatal("Expected error in response")
	}

	// Check that the error was tracked
	stats := GlobalTracker.GetStats()
	if len(stats) != 1 {
		t.Errorf("Expected 1 tracked error, got %d", len(stats))
	}

	if providerStats, exists := stats[CodeProviderTimeout]; !exists {
		t.Error("Expected provider timeout to be tracked")
	} else if providerStats.Count != 1 {
		t.Errorf("Expected count of 1, got %d", providerStats.Count)
	}
}

func TestRetryIntegration(t *testing.T) {
	// Test retry mechanism with AppError
	retrier := NewRetrier(DefaultRetryConfig())

	attempts := 0
	operation := func(ctx context.Context) (any, error) {
		attempts++
		if attempts < 3 {
			return nil, NewAppError(CodeProviderTimeout, "Timeout", "Try again", nil)
		}
		return "success", nil
	}

	ctx := context.Background()
	result := retrier.Retry(ctx, "test-operation", operation)

	if !result.Succeeded {
		t.Errorf("Expected operation to succeed, but it failed: %v", result.Err)
	}

	if result.Attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", result.Attempts)
	}

	if result.Value != "success" {
		t.Errorf("Expected success result, got %v", result.Value)
	}
}

func TestCircuitBreakerIntegration(t *testing.T) {
	// Test circuit breaker with error handling
	cb := NewCircuitBreaker(2, 1*time.Second)

	// Function that always fails
	failingOp := func(ctx context.Context) (any, error) {
		return nil, NewAppError(CodeProviderUnavailable, "Service down", "Service unavailable", nil)
	}

	ctx := context.Background()

	// First failure - circuit should remain closed
	_, err := cb.Execute(ctx, "test", failingOp)
	if err == nil {
		t.Error("Expected error from failing operation")
	}

	if cb.GetState() != CircuitClosed {
		t.Error("Expected circuit to be closed after 1 failure")
	}

	// Second failure should open the circuit
	_, err = cb.Execute(ctx, "test", failingOp)
	if err == nil {
		t.Error("Expected error from failing operation")
	}

	if cb.GetState() != CircuitOpen {
		t.Error("Expected circuit to be open after 2 failures")
	}

	// Now the circuit should fail fast
	_, err = cb.Execute(ctx, "test", failingOp)
	if err == nil {
		t.Error("Expected error from open circuit")
	}

	// Check that it's an AppError with the right code
	if appErr, ok := err.(*AppError); !ok {
		t.Error("Expected AppError from circuit breaker")
	} else if appErr.Code != CodeProviderUnavailable {
		t.Errorf("Expected provider unavailable error, got %s", appErr.Code)
	}
}
