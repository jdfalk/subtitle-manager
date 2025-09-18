// file: pkg/services/registry_test.go
// version: 1.0.0
// guid: 5b6c7d8e-9f0a-4b3c-8d7e-6f5e4d3c2b1a

package services

import (
	"context"
	"errors"
	"testing"

	enginev1 "github.com/jdfalk/subtitle-manager/pkg/engine/v1"
	filev1 "github.com/jdfalk/subtitle-manager/pkg/file/v1"
	webv1 "github.com/jdfalk/subtitle-manager/pkg/web/v1"
	"github.com/stretchr/testify/require"
)

// minimal stubs to simulate health responses
type healthyWeb struct{ WebServiceImpl }
type healthyEngine struct{ EngineServiceImpl }
type healthyFile struct{ FileServiceImpl }

func (h healthyWeb) HealthCheck(ctx context.Context, _ *webv1.HealthCheckRequest) (*webv1.HealthCheckResponse, error) {
	return &webv1.HealthCheckResponse{}, nil
}
func (h healthyEngine) HealthCheck(ctx context.Context, _ *enginev1.HealthCheckRequest) (*enginev1.HealthCheckResponse, error) {
	return &enginev1.HealthCheckResponse{}, nil
}
func (h healthyFile) HealthCheck(ctx context.Context, _ *filev1.HealthCheckRequest) (*filev1.HealthCheckResponse, error) {
	return &filev1.HealthCheckResponse{}, nil
}

type failingWeb struct{ WebServiceImpl }

func (f failingWeb) HealthCheck(ctx context.Context, _ *webv1.HealthCheckRequest) (*webv1.HealthCheckResponse, error) {
	return nil, errors.New("web unhealthy")
}

func TestServiceRegistry_HealthCheck_AllHealthy(t *testing.T) {
	reg := NewServiceRegistry(&healthyWeb{}, &healthyEngine{}, &healthyFile{})
	err := reg.HealthCheck(context.Background())
	require.NoError(t, err)
}

func TestServiceRegistry_HealthCheck_WebFails(t *testing.T) {
	reg := NewServiceRegistry(&failingWeb{}, &healthyEngine{}, &healthyFile{})
	err := reg.HealthCheck(context.Background())
	require.Error(t, err)
}
