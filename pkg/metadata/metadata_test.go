package metadata

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseFileName(t *testing.T) {
	m, err := ParseFileName("Show.Name.S01E02.mkv")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if m.Type != TypeEpisode || m.Title != "Show Name" || m.Season != 1 || m.Episode != 2 {
		t.Fatalf("unexpected result: %+v", m)
	}
	m, err = ParseFileName("Movie Title (2020).mp4")
	if err != nil {
		t.Fatalf("parse movie: %v", err)
	}
	if m.Type != TypeMovie || m.Title != "Movie Title" || m.Year != 2020 {
		t.Fatalf("unexpected movie: %+v", m)
	}
}

func TestQueryMovie(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/search/movie" {
			fmt.Fprint(w, `{"results":[{"id":1,"title":"Test","release_date":"2020-01-01"}]}`)
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()
	SetTMDBAPIBase(srv.URL)

	m, err := QueryMovie(context.Background(), "Test", 0, "key")
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	if m.TMDBID != 1 || m.Title != "Test" || m.Year != 2020 {
		t.Fatalf("unexpected movie: %+v", m)
	}
}

func TestQueryEpisode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/search/tv":
			fmt.Fprint(w, `{"results":[{"id":2,"name":"Show"}]}`)
		case "/tv/2/season/1/episode/3":
			fmt.Fprint(w, `{"id":5,"name":"Episode"}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	SetTMDBAPIBase(srv.URL)

	m, err := QueryEpisode(context.Background(), "Show", 1, 3, "key")
	if err != nil {
		t.Fatalf("query ep: %v", err)
	}
	if m.TMDBID != 5 || m.Title != "Show" || m.EpisodeTitle != "Episode" || m.Season != 1 || m.Episode != 3 {
		t.Fatalf("unexpected episode: %+v", m)
	}
}
