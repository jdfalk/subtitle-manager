// file: pkg/webserver/health.go
// version: 1.3.0
// guid: 9e8d7c6b-5a4f-3d2c-1b0a-9f8e7d6c5b4a

package webserver

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/errors"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

// Simple local health provider replacement
type HealthStatus string

const (
	HealthStatusUp       HealthStatus = "up"
	HealthStatusDown     HealthStatus = "down"
	HealthStatusDegraded HealthStatus = "degraded"
)

type HealthResult struct {
	Status  HealthStatus           `json:"status"`
	Details map[string]interface{} `json:"details,omitempty"`
	Error   error                  `json:"error,omitempty"`
}

type HealthCheck struct {
	Name   string
	Status HealthStatus
	Error  error
}

type SimpleHealthProvider struct {
	mu     sync.RWMutex
	checks map[string]func(context.Context) HealthResult
}

func (p *SimpleHealthProvider) Status() HealthStatus {
	return HealthStatusUp
}

func (p *SimpleHealthProvider) CheckAll(ctx context.Context) (HealthResult, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	overallStatus := HealthStatusUp
	details := make(map[string]interface{})

	for name, checkFunc := range p.checks {
		result := checkFunc(ctx)
		details[name] = map[string]interface{}{
			"status":  result.Status,
			"details": result.Details,
		}

		if result.Status == HealthStatusDown {
			overallStatus = HealthStatusDown
		} else if result.Status == HealthStatusDegraded && overallStatus != HealthStatusDown {
			overallStatus = HealthStatusDegraded
		}
	}

	return HealthResult{
		Status:  overallStatus,
		Details: details,
	}, nil
}

func (p *SimpleHealthProvider) AddCheck(name string, checkFunc func(context.Context) HealthResult) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.checks[name] = checkFunc
}

var HealthProvider *SimpleHealthProvider

// InitializeHealth sets up the global health provider with built-in checks.
func InitializeHealth(endpoint string) error {
	if HealthProvider != nil {
		return nil
	}

	HealthProvider = &SimpleHealthProvider{
		checks: make(map[string]func(context.Context) HealthResult),
	}

	// Add error health check
	HealthProvider.AddCheck("errors", errorHealthCheck)

	// Add cache health check
	HealthProvider.AddCheck("cache", cacheHealthCheck)

	return nil
}

func errorHealthCheck(ctx context.Context) HealthResult {
	stats := errors.GlobalTracker.GetStats()
	var total, critical int
	recent := time.Now().Add(-5 * time.Minute)
	for _, stat := range stats {
		total += stat.Count
		if stat.LastOccurred.After(recent) && !stat.Retryable {
			critical += stat.Count
		}
	}

	status := HealthStatusUp
	if critical > 10 {
		status = HealthStatusDown
	} else if critical > 5 {
		status = HealthStatusDegraded
	}

	return HealthResult{
		Status: status,
		Details: map[string]interface{}{
			"total_errors":           total,
			"critical_errors_recent": critical,
			"unique_error_types":     len(stats),
		},
	}
}

func cacheHealthCheck(ctx context.Context) HealthResult {
	logger := logging.GetLogger("webserver.cache")
	if cacheManager == nil {
		return HealthResult{
			Status: HealthStatusDown,
			Error:  fmt.Errorf("cache not initialized"),
		}
	}

	testKey := "health-check"
	testValue := []byte("ok")

	if err := cacheManager.SetAPIResponse(ctx, testKey, testValue); err != nil {
		logger.Errorf("cache health check failed (set): %v", err)
		return HealthResult{
			Status:  HealthStatusDown,
			Error:   err,
			Details: map[string]interface{}{"message": "Failed to write to cache"},
		}
	}

	value, err := cacheManager.GetAPIResponse(ctx, testKey)
	if err != nil {
		logger.Errorf("cache health check failed (get): %v", err)
		return HealthResult{
			Status:  HealthStatusDown,
			Error:   err,
			Details: map[string]interface{}{"message": "Failed to read from cache"},
		}
	}
	if string(value) != string(testValue) {
		logger.Error("cache health check failed: value mismatch")
		return HealthResult{
			Status:  HealthStatusDown,
			Details: map[string]interface{}{"message": "Cache value mismatch"},
		}
	}

	_ = cacheManager.Delete(ctx, "api:"+testKey)

	stats, err := cacheManager.Stats(ctx)
	if err != nil {
		logger.Warnf("failed to get stats for health check: %v", err)
	}

	details := map[string]interface{}{"message": "Cache is operational"}
	if stats != nil {
		details["stats"] = stats
	}

	return HealthResult{
		Status:  HealthStatusUp,
		Details: details,
	}
}

// GetHealthProvider returns the global health provider instance.
func GetHealthProvider() *SimpleHealthProvider {
	return HealthProvider
}
