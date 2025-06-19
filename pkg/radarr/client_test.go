// file: pkg/radarr/client_test.go
package radarr

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// TestMovies verifies that Movies fetches media items from the Radarr API.
func TestMovies(t *testing.T) {
	var reqPath, apiKey string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqPath = r.URL.String()
		apiKey = r.Header.Get("X-Api-Key")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"title":"Movie","movieFile":{"path":"/m.mkv"}},{"title":"Empty","movieFile":{"path":""}}]`)
	}))
	defer srv.Close()

	c := NewClient(srv.URL, "key")
	items, err := c.Movies(context.Background())
	if err != nil {
		t.Fatalf("movies: %v", err)
	}
	if reqPath != "/api/v3/movie?includeMovieFile=true" {
		t.Fatalf("unexpected path %s", reqPath)
	}
	if apiKey != "key" {
		t.Fatalf("expected header key, got %s", apiKey)
	}
	if len(items) != 1 || items[0].Path != "/m.mkv" || items[0].Title != "Movie" {
		t.Fatalf("unexpected items %+v", items)
	}
}

// TestSync verifies that Sync stores new media items only once.
func TestSync(t *testing.T) {
	store, err := database.OpenSQLStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	// Pre-existing item should not be inserted again.
	store.InsertMediaItem(&database.MediaItem{Path: "/m1.mkv", Title: "Old"})

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"title":"Old","movieFile":{"path":"/m1.mkv"}},{"title":"New","movieFile":{"path":"/m2.mkv"}}]`)
	}))
	defer srv.Close()

	c := NewClient(srv.URL, "")
	if err := Sync(context.Background(), c, store); err != nil {
		t.Fatalf("sync: %v", err)
	}
	items, err := store.ListMediaItems()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	var found bool
	for _, it := range items {
		if it.Path == "/m2.mkv" {
			found = true
		}
	}
	if !found {
		t.Fatal("new item not stored")
	}
}
