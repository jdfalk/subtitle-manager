// file: pkg/cache/proto_test.go
// version: 1.0.0
// guid: d0768c0c-f2fd-44b8-8e0f-596b6f30625f

package cache

import (
	"testing"
	"time"
)

// TestTTLConfigProtoRoundTrip verifies TTLConfig to proto conversion and back.
func TestTTLConfigProtoRoundTrip(t *testing.T) {
	original := TTLConfig{
		ProviderSearchResults: 5 * time.Minute,
		SearchResults:         10 * time.Minute,
		TMDBMetadata:          24 * time.Hour,
		TranslationResults:    0,
		UserSessions:          2 * time.Hour,
		APIResponses:          15 * time.Minute,
	}

	protoMap := TTLConfigToProto(original)
	result := TTLConfigFromProto(protoMap)

	if original != result {
		t.Errorf("expected %v, got %v", original, result)
	}
}
