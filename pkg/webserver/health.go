// file: pkg/webserver/health.go
// version: 1.0.0
// guid: 9e8d7c6b-5a4f-3d2c-1b0a-9f8e7d6c5b4a

package webserver

import (
	"context"
	"fmt"
	"time"

	ghealth "github.com/jdfalk/gcommon/pkg/health"
	"github.com/jdfalk/subtitle-manager/pkg/errors"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
)

var HealthProvider ghealth.Provider

// initializeHealth sets up the global health provider with built-in checks.
func initializeHealth(endpoint string) error {
	if HealthProvider != nil {
		return nil
	}

	cfg := ghealth.DefaultConfig()
	cfg.Endpoint = endpoint
	cfg.EnableLivenessEndpoint = true
	cfg.EnableReadinessEndpoint = true
	cfg.MetricsEnabled = true

	provider, err := ghealth.NewProvider(cfg)
	if err != nil {
		return err
	}

	provider.Register("errors", ghealth.CheckFunc(errorHealthCheck), ghealth.WithType(ghealth.TypeComponent))
	provider.Register("cache", ghealth.CheckFunc(cacheHealthCheck), ghealth.WithType(ghealth.TypeComponent))

	HealthProvider = provider
	return nil
}

func errorHealthCheck(ctx context.Context) (ghealth.Result, error) {
	stats := errors.GlobalTracker.GetStats()
	var total, critical int
	recent := time.Now().Add(-5 * time.Minute)
	for _, stat := range stats {
		total += stat.Count
		if stat.LastOccurred.After(recent) && !stat.Retryable {
			critical += stat.Count
		}
	}

	status := ghealth.StatusUp
	if critical > 10 {
		status = ghealth.StatusDown
	} else if critical > 5 {
		status = ghealth.StatusDegraded
	}

	res := ghealth.NewResult(status).WithDetails(map[string]any{
		"total_errors":           total,
		"critical_errors_recent": critical,
		"unique_error_types":     len(stats),
	})
	return res, nil
}

func cacheHealthCheck(ctx context.Context) (ghealth.Result, error) {
	logger := logging.GetLogger("webserver.cache")
	if cacheManager == nil {
		return ghealth.NewResult(ghealth.StatusDown).WithError(fmt.Errorf("cache not initialized")), nil
	}

	testKey := "health-check"
	testValue := []byte("ok")

	if err := cacheManager.SetAPIResponse(ctx, testKey, testValue); err != nil {
		logger.Errorf("cache health check failed (set): %v", err)
		return ghealth.NewResult(ghealth.StatusDown).
			WithError(err).
			WithDetails(map[string]any{"message": "Failed to write to cache"}), nil
	}

	value, err := cacheManager.GetAPIResponse(ctx, testKey)
	if err != nil {
		logger.Errorf("cache health check failed (get): %v", err)
		return ghealth.NewResult(ghealth.StatusDown).
			WithError(err).
			WithDetails(map[string]any{"message": "Failed to read from cache"}), nil
	}
	if string(value) != string(testValue) {
		logger.Error("cache health check failed: value mismatch")
		return ghealth.NewResult(ghealth.StatusDown).
			WithDetails(map[string]any{"message": "Cache value mismatch"}), nil
	}

	_ = cacheManager.Delete(ctx, "api:"+testKey)

	stats, err := cacheManager.Stats(ctx)
	if err != nil {
		logger.Warnf("failed to get stats for health check: %v", err)
	}

	res := ghealth.NewResult(ghealth.StatusUp).WithDetails(map[string]any{"message": "Cache is operational"})
	if stats != nil {
		res.WithDetails(map[string]any{"stats": stats})
	}
	return res, nil
}
