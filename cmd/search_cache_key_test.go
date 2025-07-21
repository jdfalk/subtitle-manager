// file: cmd/search_cache_key_test.go
// version: 1.0.0
// guid: fb4d1fb5-e79d-4a64-93a3-3d0e7e4d6c3b
package cmd

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

// computeCacheKey replicates the CLI cache key generation logic.
func computeCacheKey(names []string, media, lang string) string {
	sorted := make([]string, len(names))
	copy(sorted, names)
	sort.Strings(sorted)
	req := struct {
		Providers []string `json:"providers"`
		MediaPath string   `json:"mediaPath"`
		Language  string   `json:"language"`
	}{Providers: sorted, MediaPath: media, Language: lang}
	data, _ := json.Marshal(req)
	sum := sha1.Sum(data)
	return fmt.Sprintf("%x", sum)
}

// TestComputeCacheKeyOrderIndependence verifies provider order does not affect the cache key.
func TestComputeCacheKeyOrderIndependence(t *testing.T) {
	key1 := computeCacheKey([]string{"b", "a"}, "movie.mkv", "en")
	key2 := computeCacheKey([]string{"a", "b"}, "movie.mkv", "en")
	if key1 != key2 {
		t.Fatalf("cache key should be provider order independent: %s vs %s", key1, key2)
	}
}
