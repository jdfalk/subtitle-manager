package plex

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/library/sections/all" || r.URL.Query().Get("X-Plex-Token") != "tok" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"MediaContainer":{"Metadata":[{"type":"movie","title":"Movie","ratingKey":"1","Media":[{"Part":[{"file":"/m.mov"}]}]},{"type":"episode","grandparentTitle":"Show","parentIndex":1,"index":2,"ratingKey":"2","Media":[{"Part":[{"file":"/s.mkv"}]}]}]}}`)
	}))
	defer srv.Close()
	c := NewClient(srv.URL, "tok")
	items, err := c.AllItems(context.Background())
	if err != nil {
		t.Fatalf("items: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if items[0].Title != "Movie" || items[1].Season != 1 || items[1].Episode != 2 {
		t.Fatalf("unexpected items: %+v", items)
	}
}

func TestRefresh(t *testing.T) {
	calledAll := false
	calledItem := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/library/sections/all/refresh":
			calledAll = true
		case "/library/metadata/5/refresh":
			calledItem = true
		default:
			t.Fatalf("unexpected path %s", r.URL.Path)
		}
	}))
	defer srv.Close()

	c := NewClient(srv.URL, "")
	if err := c.RefreshLibrary(context.Background()); err != nil {
		t.Fatalf("refresh all: %v", err)
	}
	if !calledAll {
		t.Fatal("all refresh not called")
	}
	if err := c.RefreshItem(context.Background(), "5"); err != nil {
		t.Fatalf("refresh item: %v", err)
	}
	if !calledItem {
		t.Fatal("item refresh not called")
	}
}
