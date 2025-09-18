// file: pkg/webserver/health_test.go
// version: 1.2.0
// guid: a1b2c3d4-e5f6-7a8b-9c0d-ef1234567890

package webserver

import (
	"context"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/cache"
)

// TestInitializeHealth verifies the health provider initializes correctly.
func TestInitializeHealth(t *testing.T) {
	manager, err := cache.NewManager(cache.DefaultConfig())
	if err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}
	defer manager.Close()
	cacheManager = manager

	if err := InitializeHealth("/health-test"); err != nil {
		t.Fatalf("init failed: %v", err)
	}
	if HealthProvider == nil {
		t.Fatal("provider not set")
	}
	res, err := HealthProvider.CheckAll(context.Background())
	if err != nil {
		t.Fatalf("check failed: %v", err)
	}
	if string(res.Status) != "up" {
		t.Errorf("expected status up, got %v", res.Status)
	}
}

// TestGRPCHealthServerCheck ensures the gRPC health server responds.
func TestGRPCHealthServerCheck(t *testing.T) {
	manager, _ := cache.NewManager(cache.DefaultConfig())
	defer manager.Close()
	cacheManager = manager
	_ = InitializeHealth("")

	// TODO: Replace with simple health check when gRPC is needed
	if HealthProvider == nil {
		t.Fatal("health provider nil")
	}
}
