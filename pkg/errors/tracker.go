// file: pkg/errors/tracker.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174004

package errors

import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/sirupsen/logrus"
)

// ErrorStats represents statistics for a specific error type.
type ErrorStats struct {
	Code         ErrorCode `json:"code"`
	Count        int       `json:"count"`
	LastOccurred time.Time `json:"last_occurred"`
	FirstSeen    time.Time `json:"first_seen"`
	Message      string    `json:"message"`
	Retryable    bool      `json:"retryable"`
}

// ErrorTracker tracks application errors for monitoring and alerting.
type ErrorTracker struct {
	mu     sync.RWMutex
	stats  map[ErrorCode]*ErrorStats
	recent []ErrorEvent
	maxRecentEvents int
	logger *logrus.Entry
}

// ErrorEvent represents a single error occurrence.
type ErrorEvent struct {
	Timestamp time.Time              `json:"timestamp"`
	Code      ErrorCode              `json:"code"`
	Message   string                 `json:"message"`
	UserMsg   string                 `json:"user_message"`
	Context   map[string]interface{} `json:"context,omitempty"`
}

// NewErrorTracker creates a new error tracker.
func NewErrorTracker(maxRecentEvents int) *ErrorTracker {
	return &ErrorTracker{
		stats:           make(map[ErrorCode]*ErrorStats),
		recent:          make([]ErrorEvent, 0, maxRecentEvents),
		maxRecentEvents: maxRecentEvents,
		logger:          logging.GetLogger("error-tracker"),
	}
}

// TrackError records an error occurrence.
func (et *ErrorTracker) TrackError(appErr *AppError) {
	et.mu.Lock()
	defer et.mu.Unlock()
	
	now := time.Now()
	
	// Update or create stats
	stats, exists := et.stats[appErr.Code]
	if !exists {
		stats = &ErrorStats{
			Code:      appErr.Code,
			FirstSeen: now,
			Message:   appErr.Message,
			Retryable: appErr.Retryable,
		}
		et.stats[appErr.Code] = stats
	}
	
	stats.Count++
	stats.LastOccurred = now
	
	// Add to recent events
	event := ErrorEvent{
		Timestamp: now,
		Code:      appErr.Code,
		Message:   appErr.Message,
		UserMsg:   appErr.UserMsg,
		Context:   appErr.Context,
	}
	
	et.recent = append(et.recent, event)
	
	// Keep only the most recent events
	if len(et.recent) > et.maxRecentEvents {
		// Remove oldest events
		copy(et.recent, et.recent[len(et.recent)-et.maxRecentEvents:])
		et.recent = et.recent[:et.maxRecentEvents]
	}
	
	// Log the tracking
	et.logger.WithFields(map[string]interface{}{
		"error_code":    appErr.Code,
		"total_count":   stats.Count,
		"last_occurred": stats.LastOccurred,
	}).Debug("Error tracked")
}

// GetStats returns current error statistics.
func (et *ErrorTracker) GetStats() map[ErrorCode]*ErrorStats {
	et.mu.RLock()
	defer et.mu.RUnlock()
	
	// Create a copy to avoid race conditions
	result := make(map[ErrorCode]*ErrorStats, len(et.stats))
	for code, stats := range et.stats {
		statsCopy := *stats
		result[code] = &statsCopy
	}
	
	return result
}

// GetRecentEvents returns the most recent error events.
func (et *ErrorTracker) GetRecentEvents(limit int) []ErrorEvent {
	et.mu.RLock()
	defer et.mu.RUnlock()
	
	if limit <= 0 || limit > len(et.recent) {
		limit = len(et.recent)
	}
	
	// Return the most recent events (from the end of the slice)
	start := len(et.recent) - limit
	if start < 0 {
		start = 0
	}
	
	result := make([]ErrorEvent, limit)
	copy(result, et.recent[start:])
	
	// Reverse to get newest first
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	
	return result
}

// GetTopErrors returns the most frequently occurring errors.
func (et *ErrorTracker) GetTopErrors(limit int) []*ErrorStats {
	et.mu.RLock()
	defer et.mu.RUnlock()
	
	// Convert to slice and sort by count
	statsList := make([]*ErrorStats, 0, len(et.stats))
	for _, stats := range et.stats {
		statsCopy := *stats
		statsList = append(statsList, &statsCopy)
	}
	
	sort.Slice(statsList, func(i, j int) bool {
		return statsList[i].Count > statsList[j].Count
	})
	
	if limit > 0 && limit < len(statsList) {
		statsList = statsList[:limit]
	}
	
	return statsList
}

// Reset clears all tracked error data.
func (et *ErrorTracker) Reset() {
	et.mu.Lock()
	defer et.mu.Unlock()
	
	et.stats = make(map[ErrorCode]*ErrorStats)
	et.recent = et.recent[:0]
	
	et.logger.Info("Error tracker reset")
}

// ErrorDashboard provides HTTP handlers for error monitoring.
type ErrorDashboard struct {
	tracker *ErrorTracker
}

// NewErrorDashboard creates a new error dashboard.
func NewErrorDashboard(tracker *ErrorTracker) *ErrorDashboard {
	return &ErrorDashboard{
		tracker: tracker,
	}
}

// StatsHandler returns error statistics as JSON.
func (ed *ErrorDashboard) StatsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stats := ed.tracker.GetStats()
		
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewEncoder(w).Encode(stats); err != nil {
			http.Error(w, "Failed to encode stats", http.StatusInternalServerError)
			return
		}
	})
}

// RecentHandler returns recent error events as JSON.
func (ed *ErrorDashboard) RecentHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limit := 50 // Default limit
		
		// Parse limit from query parameter if provided
		if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
			if parsed, err := time.ParseDuration(limitStr); err == nil {
				limit = int(parsed)
			}
		}
		
		events := ed.tracker.GetRecentEvents(limit)
		
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewEncoder(w).Encode(events); err != nil {
			http.Error(w, "Failed to encode events", http.StatusInternalServerError)
			return
		}
	})
}

// TopErrorsHandler returns the most frequent errors as JSON.
func (ed *ErrorDashboard) TopErrorsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limit := 10 // Default limit
		
		// Parse limit from query parameter if provided
		if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
			if parsed, err := time.ParseDuration(limitStr); err == nil {
				limit = int(parsed)
			}
		}
		
		topErrors := ed.tracker.GetTopErrors(limit)
		
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewEncoder(w).Encode(topErrors); err != nil {
			http.Error(w, "Failed to encode top errors", http.StatusInternalServerError)
			return
		}
	})
}

// HealthHandler returns overall error health status.
func (ed *ErrorDashboard) HealthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stats := ed.tracker.GetStats()
		
		// Calculate health metrics
		var totalErrors int
		var criticalErrors int
		recentThreshold := time.Now().Add(-5 * time.Minute)
		
		for _, stat := range stats {
			totalErrors += stat.Count
			
			// Count recent critical errors (5xx status codes)
			if stat.LastOccurred.After(recentThreshold) && !stat.Retryable {
				criticalErrors += stat.Count
			}
		}
		
		health := map[string]interface{}{
			"total_errors":            totalErrors,
			"critical_errors_recent":  criticalErrors,
			"unique_error_types":      len(stats),
			"timestamp":               time.Now(),
		}
		
		// Determine health status
		status := "healthy"
		if criticalErrors > 10 {
			status = "critical"
		} else if criticalErrors > 5 {
			status = "warning"
		}
		health["status"] = status
		
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewEncoder(w).Encode(health); err != nil {
			http.Error(w, "Failed to encode health", http.StatusInternalServerError)
			return
		}
	})
}

// Global error tracker instance
var GlobalTracker = NewErrorTracker(1000)

// Track is a convenience function that uses the global tracker.
func Track(appErr *AppError) {
	GlobalTracker.TrackError(appErr)
}

// GetDashboard returns a dashboard instance using the global tracker.
func GetDashboard() *ErrorDashboard {
	return NewErrorDashboard(GlobalTracker)
}