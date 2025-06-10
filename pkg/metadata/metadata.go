// file: pkg/metadata/metadata.go
package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"subtitle-manager/pkg/database"
)

var tmdbAPIBase = "https://api.themoviedb.org/3"

// SetTMDBAPIBase overrides the default TMDB API base URL. Primarily used for testing.
func SetTMDBAPIBase(u string) { tmdbAPIBase = u }

// MediaType differentiates between movie and TV episode metadata.
type MediaType int

const (
	// TypeMovie represents a movie file.
	TypeMovie MediaType = iota
	// TypeEpisode represents an episode of a TV show.
	TypeEpisode
)

// MediaInfo stores basic metadata for a movie or episode.
type MediaInfo struct {
	Type         MediaType // movie or episode
	Title        string    // movie title or show name
	Year         int       // release year for movies
	TMDBID       int       // TMDB identifier
	EpisodeTitle string    // title of the episode (if any)
	Season       int       // season number for episodes
	Episode      int       // episode number for episodes
}

var (
	tvRegex    = regexp.MustCompile(`(?i)^(.*?)[ ._-]+s(\d{1,2})e(\d{1,2})`)
	movieRegex = regexp.MustCompile(`(?i)^(.*?)[ ._\(](\d{4})`)
)

// ParseFileName attempts to derive metadata from a media file name.
// It recognises common patterns such as "Show.Name.S01E02" for episodes
// and "Movie Title (2020)" for movies.
func ParseFileName(path string) (*MediaInfo, error) {
	name := filepath.Base(path)
	name = strings.TrimSuffix(name, filepath.Ext(name))

	if m := tvRegex.FindStringSubmatch(name); m != nil {
		season, _ := strconv.Atoi(m[2])
		episode, _ := strconv.Atoi(m[3])
		title := cleanTitle(m[1])
		return &MediaInfo{Type: TypeEpisode, Title: title, Season: season, Episode: episode}, nil
	}
	if m := movieRegex.FindStringSubmatch(name); m != nil {
		year, _ := strconv.Atoi(m[2])
		title := cleanTitle(m[1])
		return &MediaInfo{Type: TypeMovie, Title: title, Year: year}, nil
	}
	return nil, fmt.Errorf("unrecognised file name")
}

func cleanTitle(s string) string {
	s = strings.ReplaceAll(s, ".", " ")
	s = strings.ReplaceAll(s, "_", " ")
	return strings.TrimSpace(s)
}

// QueryMovie searches TMDB for a movie matching the title and optional year.
// It returns the first result's metadata.
func QueryMovie(ctx context.Context, title string, year int, apiKey string) (*MediaInfo, error) {
	q := url.Values{}
	q.Set("api_key", apiKey)
	q.Set("query", title)
	if year > 0 {
		q.Set("year", strconv.Itoa(year))
	}
	u := fmt.Sprintf("%s/search/movie?%s", tmdbAPIBase, q.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	var mr struct {
		Results []struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Overview    string `json:"overview"`
			ReleaseDate string `json:"release_date"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&mr); err != nil {
		return nil, err
	}
	if len(mr.Results) == 0 {
		return nil, fmt.Errorf("movie not found")
	}
	r := mr.Results[0]
	y := year
	if y == 0 && len(r.ReleaseDate) >= 4 {
		y, _ = strconv.Atoi(r.ReleaseDate[:4])
	}
	return &MediaInfo{Type: TypeMovie, Title: r.Title, Year: y, TMDBID: r.ID}, nil
}

// QueryEpisode retrieves episode details from TMDB. The function searches for
// the show by title then requests the specific season and episode metadata.
func QueryEpisode(ctx context.Context, show string, season, episode int, apiKey string) (*MediaInfo, error) {
	q := url.Values{}
	q.Set("api_key", apiKey)
	q.Set("query", show)
	u := fmt.Sprintf("%s/search/tv?%s", tmdbAPIBase, q.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	var sr struct {
		Results []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return nil, err
	}
	if len(sr.Results) == 0 {
		return nil, fmt.Errorf("show not found")
	}
	sid := sr.Results[0].ID
	showName := sr.Results[0].Name

	u = fmt.Sprintf("%s/tv/%d/season/%d/episode/%d?api_key=%s", tmdbAPIBase, sid, season, episode, url.QueryEscape(apiKey))
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	var er struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		AirDate string `json:"air_date"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&er); err != nil {
		return nil, err
	}
	return &MediaInfo{Type: TypeEpisode, Title: showName, Season: season, Episode: episode, EpisodeTitle: er.Name, TMDBID: er.ID}, nil
}

// ScanLibrary walks a directory tree and inserts video files into the media database.
// It parses filenames to extract metadata and stores the results using the provided store.
func ScanLibrary(ctx context.Context, dir string, store interface{}) error {
	// Import the database types we need
	type MediaStore interface {
		InsertMediaItem(*database.MediaItem) error
	}

	mediaStore, ok := store.(MediaStore)
	if !ok {
		return fmt.Errorf("store does not implement MediaItem methods")
	}

	// Supported video extensions
	videoExts := map[string]bool{
		".mp4": true, ".mkv": true, ".avi": true, ".mov": true,
		".wmv": true, ".flv": true, ".webm": true, ".m4v": true,
	}

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-video files
		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !videoExts[ext] {
			return nil
		}

		// Parse the filename to extract metadata
		mediaInfo, err := ParseFileName(path)
		if err != nil {
			// If we can't parse, just use the base filename as title
			name := filepath.Base(path)
			name = strings.TrimSuffix(name, filepath.Ext(name))
			mediaInfo = &MediaInfo{
				Type:  TypeMovie, // Default to movie
				Title: cleanTitle(name),
			}
		}

		// Create database record
		item := &database.MediaItem{
			Path:    path,
			Title:   mediaInfo.Title,
			Season:  mediaInfo.Season,
			Episode: mediaInfo.Episode,
		}

		// Insert into database
		return mediaStore.InsertMediaItem(item)
	})
}
