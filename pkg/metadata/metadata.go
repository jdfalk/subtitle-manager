// Package metadata provides tools for extracting, parsing, and managing media metadata.
// It supports TMDB lookups, filename parsing, and metadata enrichment for subtitle management.
//
// This package is used by subtitle-manager to associate subtitles with the correct media files.
//
// See: https://www.themoviedb.org/
// File: pkg/metadata/metadata.go
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

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/security"
)

var tmdbAPIBase = "https://api.themoviedb.org/3"
var omdbAPIBase = "https://www.omdbapi.com"

// SetTMDBAPIBase overrides the default TMDB API base URL. Primarily used for testing.
func SetTMDBAPIBase(u string) { tmdbAPIBase = u }

// SetOMDBAPIBase overrides the default OMDb API base URL. Primarily used for testing.
func SetOMDBAPIBase(u string) { omdbAPIBase = u }

// Supported video file extensions used for library scanning.
var videoExts = map[string]bool{
	".mp4": true, ".mkv": true, ".avi": true, ".mov": true,
	".wmv": true, ".flv": true, ".webm": true, ".m4v": true,
}

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
	Languages    []string  // spoken languages
	Rating       float64   // IMDB rating
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

// FetchMovieMetadata retrieves movie details from TMDB and enriches them with
// language and rating information from OMDb.
// tmdbKey is the TMDB API key, omdbKey the OMDb API key.
func FetchMovieMetadata(ctx context.Context, title string, year int, tmdbKey, omdbKey string) (*MediaInfo, error) {
	info, err := QueryMovie(ctx, title, year, tmdbKey)
	if err != nil {
		return nil, err
	}
	langs, rating, err := fetchOMDBInfo(ctx, url.Values{
		"t":      []string{title},
		"apikey": []string{omdbKey},
		"y":      []string{strconv.Itoa(year)},
	})
	if err == nil {
		info.Languages = langs
		info.Rating = rating
	}
	return info, nil
}

// FetchEpisodeMetadata retrieves episode details from TMDB and enriches them
// with language and rating information from OMDb.
func FetchEpisodeMetadata(ctx context.Context, show string, season, episode int, tmdbKey, omdbKey string) (*MediaInfo, error) {
	info, err := QueryEpisode(ctx, show, season, episode, tmdbKey)
	if err != nil {
		return nil, err
	}
	langs, rating, err := fetchOMDBInfo(ctx, url.Values{
		"t":       []string{show},
		"Season":  []string{strconv.Itoa(season)},
		"Episode": []string{strconv.Itoa(episode)},
		"apikey":  []string{omdbKey},
	})
	if err == nil {
		info.Languages = langs
		info.Rating = rating
	}
	return info, nil
}

// fetchOMDBInfo queries the OMDb API and returns language and rating fields. It
// returns an error only if the HTTP request fails or the response cannot be
// decoded. API errors are ignored.
func fetchOMDBInfo(ctx context.Context, q url.Values) ([]string, float64, error) {
	if q.Get("apikey") == "" {
		return nil, 0, fmt.Errorf("missing api key")
	}
	u := fmt.Sprintf("%s/?%s", omdbAPIBase, q.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("status %d", resp.StatusCode)
	}
	var r struct {
		Response   string `json:"Response"`
		Language   string `json:"Language"`
		IMDBRating string `json:"imdbRating"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, 0, err
	}
	if strings.ToLower(r.Response) != "true" {
		return nil, 0, fmt.Errorf("not found")
	}
	var langs []string
	for _, l := range strings.Split(r.Language, ",") {
		l = strings.TrimSpace(l)
		if l != "" {
			langs = append(langs, l)
		}
	}
	rating, _ := strconv.ParseFloat(r.IMDBRating, 64)
	return langs, rating, nil
}

// ScanLibrary walks a directory tree and inserts video files into the media database.
// It parses filenames to extract metadata and stores the results using the provided store.
func ScanLibrary(ctx context.Context, dir string, store database.SubtitleStore) error {
	return scanLibrary(ctx, dir, store, nil)
}

func scanLibrary(ctx context.Context, dir string, store database.SubtitleStore, cb ProgressFunc) error {
	sanitizedDir, err := security.ValidateAndSanitizePath(dir)
	if err != nil {
		return err
	}

	return filepath.Walk(sanitizedDir, func(path string, info os.FileInfo, err error) error {
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
		if err := store.InsertMediaItem(item); err != nil {
			return err
		}

		if cb != nil {
			cb(path)
		}

		return nil
	})
}

// ProgressFunc is called with each processed video file path during scanning.
type ProgressFunc func(file string)

// ScanLibraryProgress performs ScanLibrary and invokes cb for each processed file.
func ScanLibraryProgress(ctx context.Context, dir string, store database.SubtitleStore, cb ProgressFunc) error {
	return scanLibrary(ctx, dir, store, cb)
}
