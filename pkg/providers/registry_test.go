package providers

import (
	"sort"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/providers/mocks"
)

func TestRegisterFactoryAndGet(t *testing.T) {
	p := mocks.NewMockProvider(t)
	RegisterFactory("mockreg", func() Provider { return p })
	t.Cleanup(func() { delete(factories, "mockreg") })

	got, err := Get("mockreg", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != p {
		t.Fatalf("unexpected provider %+v", got)
	}
}

func TestGetUnknownProvider(t *testing.T) {
	if _, err := Get("unknown-provider", ""); err == nil {
		t.Fatal("expected error for unknown provider")
	}
}

func TestAllIncludesRegisteredSorted(t *testing.T) {
	base := All()
	RegisterFactory("mocka", func() Provider { return nil })
	RegisterFactory("mockz", func() Provider { return nil })
	t.Cleanup(func() { delete(factories, "mocka"); delete(factories, "mockz") })

	names := All()
	if len(names) != len(base)+2 {
		t.Fatalf("expected %d names, got %d", len(base)+2, len(names))
	}
	if !sort.StringsAreSorted(names) {
		t.Fatal("names not sorted")
	}
	foundA := false
	foundZ := false
	for _, n := range names {
		if n == "mocka" {
			foundA = true
		}
		if n == "mockz" {
			foundZ = true
		}
	}
	if !foundA || !foundZ {
		t.Fatalf("registered names missing: %v", names)
	}
}
