// file: pkg/monitoring/sync.go
// version: 1.0.0
// guid: 12345678-1234-1234-1234-123456789013

package monitoring

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// SyncOptions contains configuration for media library synchronization.
type SyncOptions struct {
	Languages   []string // Languages to monitor for
	MaxRetries  int      // Maximum retry attempts per item
	ForceRefresh bool    // Whether to refresh existing items
}

// SyncFromSonarr synchronizes TV episodes from Sonarr and adds them to monitoring.
func (m *EpisodeMonitor) SyncFromSonarr(ctx context.Context, opts SyncOptions) error {
	if m.sonarr == nil {
		return nil // Skip if Sonarr is not configured
	}

	m.logger.Info("Syncing episodes from Sonarr")
	
	// Get episodes from Sonarr
	episodes, err := m.sonarr.Episodes(ctx)
	if err != nil {
		return err
	}

	m.logger.Infof("Found %d episodes from Sonarr", len(episodes))

	// Process each episode
	for _, episode := range episodes {
		if err := m.addToMonitoring(episode, opts); err != nil {
			m.logger.Warnf("Failed to add episode %s to monitoring: %v", episode.Path, err)
		}
	}

	return nil
}

// SyncFromRadarr synchronizes movies from Radarr and adds them to monitoring.
func (m *EpisodeMonitor) SyncFromRadarr(ctx context.Context, opts SyncOptions) error {
	if m.radarr == nil {
		return nil // Skip if Radarr is not configured
	}

	m.logger.Info("Syncing movies from Radarr")
	
	// Get movies from Radarr
	movies, err := m.radarr.Movies(ctx)
	if err != nil {
		return err
	}

	m.logger.Infof("Found %d movies from Radarr", len(movies))

	// Process each movie
	for _, movie := range movies {
		if err := m.addToMonitoring(movie, opts); err != nil {
			m.logger.Warnf("Failed to add movie %s to monitoring: %v", movie.Path, err)
		}
	}

	return nil
}

// addToMonitoring adds a media item to the monitoring system.
func (m *EpisodeMonitor) addToMonitoring(mediaItem database.MediaItem, opts SyncOptions) error {
	// Check if file exists
	if _, err := os.Stat(mediaItem.Path); os.IsNotExist(err) {
		m.logger.Debugf("Skipping non-existent file: %s", mediaItem.Path)
		return nil
	}

	// Check if already being monitored (unless force refresh)
	if !opts.ForceRefresh {
		existing, err := m.store.ListMonitoredItems()
		if err != nil {
			return err
		}
		for _, item := range existing {
			if item.Path == mediaItem.Path {
				m.logger.Debugf("Already monitoring: %s", mediaItem.Path)
				return nil
			}
		}
	}

	// Serialize languages
	languagesJSON, err := json.Marshal(opts.Languages)
	if err != nil {
		return err
	}

	// Create monitored item
	monitoredItem := &database.MonitoredItem{
		MediaID:     mediaItem.ID,
		Path:        mediaItem.Path,
		Languages:   string(languagesJSON),
		LastChecked: time.Time{}, // Will be checked on first monitoring cycle
		Status:      "pending",
		RetryCount:  0,
		MaxRetries:  opts.MaxRetries,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Insert into database
	if err := m.store.InsertMonitoredItem(monitoredItem); err != nil {
		return err
	}

	m.logger.Debugf("Added to monitoring: %s (languages: %v)", mediaItem.Path, opts.Languages)
	return nil
}

// RemoveFromMonitoring removes a media item from monitoring by path.
func (m *EpisodeMonitor) RemoveFromMonitoring(path string) error {
	items, err := m.store.ListMonitoredItems()
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.Path == path {
			if err := m.store.DeleteMonitoredItem(item.ID); err != nil {
				return err
			}
			m.logger.Infof("Removed from monitoring: %s", path)
			return nil
		}
	}

	return nil // Not found, no error
}

// GetMonitoringStats returns statistics about the monitoring system.
func (m *EpisodeMonitor) GetMonitoringStats() (*MonitoringStats, error) {
	items, err := m.store.ListMonitoredItems()
	if err != nil {
		return nil, err
	}

	stats := &MonitoringStats{
		Total:       len(items),
		Pending:     0,
		Monitoring:  0,
		Found:       0,
		Failed:      0,
		Blacklisted: 0,
	}

	for _, item := range items {
		switch item.Status {
		case "pending":
			stats.Pending++
		case "monitoring":
			stats.Monitoring++
		case "found":
			stats.Found++
		case "failed":
			stats.Failed++
		case "blacklisted":
			stats.Blacklisted++
		}
	}

	return stats, nil
}

// MonitoringStats contains statistics about the monitoring system.
type MonitoringStats struct {
	Total       int `json:"total"`
	Pending     int `json:"pending"`
	Monitoring  int `json:"monitoring"`
	Found       int `json:"found"`
	Failed      int `json:"failed"`
	Blacklisted int `json:"blacklisted"`
}