// file: pkg/database/metadata_helpers.go
// version: 1.1.0
// guid: 8f7e6d5c-4b3a-2198-7e6f-5d4c3b2a1098

package database

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"time"
)

// ProviderMetadata represents structured metadata from subtitle providers.
type ProviderMetadata struct {
	Quality    string  `json:"quality,omitempty"`
	Uploader   string  `json:"uploader,omitempty"`
	Rating     float64 `json:"rating,omitempty"`
	Downloads  int     `json:"downloads,omitempty"`
	Format     string  `json:"format,omitempty"`
	Encoding   string  `json:"encoding,omitempty"`
	FileSize   int     `json:"file_size,omitempty"`
	Language   string  `json:"language,omitempty"`
	Release    string  `json:"release,omitempty"`
	SourceID   string  `json:"source_id,omitempty"`
	SourceName string  `json:"source_name,omitempty"`
}

// ToJSON serializes the provider metadata to JSON string.
func (pm *ProviderMetadata) ToJSON() (string, error) {
	data, err := json.Marshal(pm)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSON deserializes provider metadata from JSON string.
func (pm *ProviderMetadata) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), pm)
}

// ModificationTypes contains constants for subtitle modification types.
const (
	ModificationTypeOriginal      = "original"
	ModificationTypeSync          = "sync"
	ModificationTypeTranslate     = "translate"
	ModificationTypeManualEdit    = "manual_edit"
	ModificationTypeAutoCorrect   = "auto_correct"
	ModificationTypeFormatConvert = "format_convert"
)

// CalculateSubtitleHash generates a SHA256 hash for subtitle content.
// This is useful for detecting duplicate subtitles and tracking sources.
func CalculateSubtitleHash(content []byte) string {
	hash := sha256.Sum256(content)
	return "sha256:" + hex.EncodeToString(hash[:])
}

// CalculateSubtitleHashFromReader generates a SHA256 hash from a reader.
func CalculateSubtitleHashFromReader(reader io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}
	return "sha256:" + hex.EncodeToString(hash.Sum(nil)), nil
}

// CreateSubtitleRecord creates a new subtitle record with enhanced metadata.
func CreateSubtitleRecord(file, videoFile, language, service string, metadata *ProviderMetadata) (*SubtitleRecord, error) {
	rec := &SubtitleRecord{
		File:             file,
		VideoFile:        videoFile,
		Language:         language,
		Service:          service,
		Embedded:         false,
		ModificationType: ModificationTypeOriginal,
		CreatedAt:        time.Now(),
	}

	if metadata != nil {
		metadataJSON, err := metadata.ToJSON()
		if err != nil {
			return nil, fmt.Errorf("failed to serialize provider metadata: %w", err)
		}
		rec.ProviderMetadata = metadataJSON
	}

	return rec, nil
}

// CreateDownloadRecord creates a new download record with enhanced metadata.
func CreateDownloadRecord(file, videoFile, provider, language, searchQuery string) *DownloadRecord {
	return &DownloadRecord{
		File:             file,
		VideoFile:        videoFile,
		Provider:         provider,
		Language:         language,
		SearchQuery:      searchQuery,
		DownloadAttempts: 1,
		CreatedAt:        time.Now(),
	}
}

// CreateSubtitleSource creates a new subtitle source record.
func CreateSubtitleSource(sourceHash, originalURL, provider string, metadata *ProviderMetadata) (*SubtitleSource, error) {
	src := &SubtitleSource{
		SourceHash:    sourceHash,
		OriginalURL:   originalURL,
		Provider:      provider,
		DownloadCount: 0,
		SuccessCount:  0,
		LastSeen:      time.Now(),
		CreatedAt:     time.Now(),
	}

	if metadata != nil {
		metadataJSON, err := metadata.ToJSON()
		if err != nil {
			return nil, fmt.Errorf("failed to serialize provider metadata: %w", err)
		}
		src.Metadata = metadataJSON
		src.Title = metadata.SourceName
		src.ReleaseInfo = metadata.Release
		if metadata.FileSize > 0 {
			src.FileSize = &metadata.FileSize
		}
	}

	return src, nil
}

// UpdateDownloadWithResult updates a download record with the result of the download attempt.
//
// Because the store currently lacks an explicit update method, this function
// creates a new download record based on the original and stores the updated
// information. The original record is left unchanged.
func UpdateDownloadWithResult(store SubtitleStore, downloadID string, success bool, errorMsg string, responseTimeMs int) error {
	downloads, err := store.ListDownloads()
	if err != nil {
		return fmt.Errorf("failed to list downloads: %w", err)
	}

	var base *DownloadRecord
	for i := range downloads {
		if downloads[i].ID == downloadID {
			base = &downloads[i]
			break
		}
	}
	if base == nil {
		return fmt.Errorf("download record %s not found", downloadID)
	}

	newRec := &DownloadRecord{
		File:             base.File,
		VideoFile:        base.VideoFile,
		Provider:         base.Provider,
		Language:         base.Language,
		SearchQuery:      base.SearchQuery,
		MatchScore:       base.MatchScore,
		DownloadAttempts: base.DownloadAttempts + 1,
		ErrorMessage:     errorMsg,
		CreatedAt:        time.Now(),
	}
	if responseTimeMs > 0 {
		newRec.ResponseTimeMs = &responseTimeMs
	}

	if err := store.InsertDownload(newRec); err != nil {
		return fmt.Errorf("failed to insert updated download record: %w", err)
	}

	return nil
}

// GetProviderPerformanceStats calculates performance statistics for a provider.
func GetProviderPerformanceStats(store SubtitleStore, provider string) (*ProviderStats, error) {
	sources, err := store.ListSubtitleSources(provider, 0) // Get all sources for provider
	if err != nil {
		return nil, fmt.Errorf("failed to list subtitle sources: %w", err)
	}

	stats := &ProviderStats{
		Provider:       provider,
		TotalSources:   len(sources),
		TotalDownloads: 0,
		TotalSuccesses: 0,
		LastSeen:       time.Time{},
	}

	var ratingSum float64
	var ratingCount int

	for _, source := range sources {
		stats.TotalDownloads += source.DownloadCount
		stats.TotalSuccesses += source.SuccessCount

		if source.LastSeen.After(stats.LastSeen) {
			stats.LastSeen = source.LastSeen
		}

		if source.AvgRating != nil {
			ratingSum += *source.AvgRating
			ratingCount++
		}
	}

	if stats.TotalDownloads > 0 {
		stats.SuccessRate = float64(stats.TotalSuccesses) / float64(stats.TotalDownloads)
	}

	if ratingCount > 0 {
		avgRating := ratingSum / float64(ratingCount)
		stats.AvgRating = &avgRating
	}

	return stats, nil
}

// ProviderStats represents performance statistics for a subtitle provider.
type ProviderStats struct {
	Provider       string    `json:"provider"`
	TotalSources   int       `json:"total_sources"`
	TotalDownloads int       `json:"total_downloads"`
	TotalSuccesses int       `json:"total_successes"`
	SuccessRate    float64   `json:"success_rate"`
	AvgRating      *float64  `json:"avg_rating,omitempty"`
	LastSeen       time.Time `json:"last_seen"`
}

// TrackSubtitleRelationship creates a subtitle record that tracks its relationship to a parent.
func TrackSubtitleRelationship(parentRec *SubtitleRecord, newFile, modificationType string) *SubtitleRecord {
	return &SubtitleRecord{
		File:             newFile,
		VideoFile:        parentRec.VideoFile,
		Language:         parentRec.Language,
		Service:          parentRec.Service,
		Release:          parentRec.Release,
		Embedded:         false,
		SourceURL:        parentRec.SourceURL,
		ProviderMetadata: parentRec.ProviderMetadata,
		ConfidenceScore:  parentRec.ConfidenceScore,
		ParentID:         &parentRec.ID,
		ModificationType: modificationType,
		CreatedAt:        time.Now(),
	}
}

// GetSubtitleHistory retrieves the complete history of a subtitle including all modifications.
func GetSubtitleHistory(store SubtitleStore, videoFile string) ([]SubtitleRecord, error) {
	records, err := store.ListSubtitlesByVideo(videoFile)
	if err != nil {
		return nil, fmt.Errorf("failed to list subtitles for video: %w", err)
	}

	// Organize records by parent/child relationships to show
	// modification hierarchy in chronological order.
	byID := make(map[string]SubtitleRecord)
	children := make(map[string][]SubtitleRecord)
	var roots []SubtitleRecord

	for _, r := range records {
		byID[r.ID] = r
		if r.ParentID == nil || *r.ParentID == "" {
			roots = append(roots, r)
			continue
		}
		parent := *r.ParentID
		children[parent] = append(children[parent], r)
	}

	sort.Slice(roots, func(i, j int) bool { return roots[i].CreatedAt.Before(roots[j].CreatedAt) })
	for id := range children {
		sort.Slice(children[id], func(i, j int) bool {
			return children[id][i].CreatedAt.Before(children[id][j].CreatedAt)
		})
	}

	var ordered []SubtitleRecord
	var walk func(SubtitleRecord)
	walk = func(rec SubtitleRecord) {
		ordered = append(ordered, rec)
		for _, child := range children[rec.ID] {
			walk(child)
		}
	}

	for _, r := range roots {
		walk(r)
	}

	return ordered, nil
}

// ValidateConfidenceScore ensures confidence scores are within valid range (0-1).
func ValidateConfidenceScore(score *float64) error {
	if score == nil {
		return nil
	}
	if *score < 0 || *score > 1 {
		return fmt.Errorf("confidence score must be between 0 and 1, got %f", *score)
	}
	return nil
}

// ValidateMatchScore ensures match scores are within valid range (0-1).
func ValidateMatchScore(score *float64) error {
	if score == nil {
		return nil
	}
	if *score < 0 || *score > 1 {
		return fmt.Errorf("match score must be between 0 and 1, got %f", *score)
	}
	return nil
}
