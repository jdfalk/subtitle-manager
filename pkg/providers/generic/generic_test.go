package generic

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

// TestClientFetch verifies that the client downloads subtitles using query parameters.
func TestClientFetch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("file") != "movie.mkv" || r.URL.Query().Get("lang") != "en" {
			t.Fatalf("unexpected query: %s", r.URL.String())
		}
		fmt.Fprint(w, "data")
	}))
	defer srv.Close()

	viper.Set("providers.generic.api_url", srv.URL)
	viper.Set("providers.generic.username", "")
	viper.Set("providers.generic.password", "")
	viper.Set("providers.generic.api_key", "")
	defer viper.Reset()

	c := New()
	b, err := c.Fetch(context.Background(), "/path/movie.mkv", "en")
	if err != nil {
		t.Fatalf("fetch: %v", err)
	}
	if string(b) != "data" {
		t.Fatalf("unexpected body: %s", b)
	}
}
