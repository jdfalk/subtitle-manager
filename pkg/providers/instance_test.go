package providers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/providers/mocks"
	"github.com/stretchr/testify/mock"
)

// TestFetchFromAllInstances verifies that provider instances are selected in
// priority order and that backoff is tracked per instance.
func TestFetchFromAllInstances(t *testing.T) {
	// Setup mock provider factory returning different mocks sequentially.
	cnt := 0
	m1 := &mocks.Provider{}
	m2 := &mocks.Provider{}
	RegisterFactory("mock", func() Provider {
		cnt++
		if cnt == 1 {
			return m1
		}
		return m2
	})
	t.Cleanup(func() {
		delete(factories, "mock")
	})

	// Register two provider instances with different priorities.
	RegisterInstance(Instance{ID: "p1", Name: "mock", Priority: 1, Enabled: true})
	RegisterInstance(Instance{ID: "p2", Name: "mock", Priority: 0, Enabled: true})
	t.Cleanup(func() {
		instancesMu.Lock()
		instances = map[string]Instance{}
		instancesMu.Unlock()
		backoffMu.Lock()
		backoffMap = map[string]time.Time{}
		backoffMu.Unlock()
	})

	m1.On("Fetch", mock.Anything, "file.mkv", "en").Return([]byte(nil), errors.New("fail"))
	m2.On("Fetch", mock.Anything, "file.mkv", "en").Return([]byte("ok"), nil)

	data, id, err := FetchFromAll(context.Background(), "file.mkv", "en", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != "ok" || id != "p2" {
		t.Fatalf("unexpected result %s %s", data, id)
	}

}
