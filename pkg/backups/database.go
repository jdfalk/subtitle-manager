// file: pkg/backups/database.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174004

package backups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// DatabaseBackupData represents the complete database backup structure.
type DatabaseBackupData struct {
	Subtitles    []database.SubtitleRecord `json:"subtitles"`
	Downloads    []database.DownloadRecord `json:"downloads"`
	MediaItems   []database.MediaItem      `json:"media_items"`
	Tags         []database.Tag            `json:"tags"`
	// Add other tables as needed
}

// DatabaseBackupper provides database backup and restore functionality.
type DatabaseBackupper struct {
	store database.SubtitleStore
}

// NewDatabaseBackupper creates a new database backup instance.
func NewDatabaseBackupper(store database.SubtitleStore) *DatabaseBackupper {
	return &DatabaseBackupper{
		store: store,
	}
}

// CreateDatabaseBackup creates a complete backup of the database.
func (db *DatabaseBackupper) CreateDatabaseBackup(ctx context.Context) ([]byte, error) {
	backupData := &DatabaseBackupData{}

	// Backup subtitles
	subtitles, err := db.store.ListSubtitles()
	if err != nil {
		return nil, fmt.Errorf("failed to backup subtitles: %w", err)
	}
	backupData.Subtitles = subtitles

	// Backup downloads
	downloads, err := db.store.ListDownloads()
	if err != nil {
		return nil, fmt.Errorf("failed to backup downloads: %w", err)
	}
	backupData.Downloads = downloads

	// Backup media items
	mediaItems, err := db.store.ListMediaItems()
	if err != nil {
		return nil, fmt.Errorf("failed to backup media items: %w", err)
	}
	backupData.MediaItems = mediaItems

	// Backup tags
	tags, err := db.store.ListTags()
	if err != nil {
		return nil, fmt.Errorf("failed to backup tags: %w", err)
	}
	backupData.Tags = tags

	// Serialize to JSON
	data, err := json.MarshalIndent(backupData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to serialize backup data: %w", err)
	}

	return data, nil
}

// RestoreDatabaseBackup restores a database from backup data.
func (db *DatabaseBackupper) RestoreDatabaseBackup(ctx context.Context, data []byte) error {
	var backupData DatabaseBackupData
	if err := json.Unmarshal(data, &backupData); err != nil {
		return fmt.Errorf("failed to deserialize backup data: %w", err)
	}

	// Restore subtitles
	for _, subtitle := range backupData.Subtitles {
		if err := db.store.InsertSubtitle(&subtitle); err != nil {
			return fmt.Errorf("failed to restore subtitle: %w", err)
		}
	}

	// Restore downloads
	for _, download := range backupData.Downloads {
		if err := db.store.InsertDownload(&download); err != nil {
			return fmt.Errorf("failed to restore download: %w", err)
		}
	}

	// Restore media items
	for _, mediaItem := range backupData.MediaItems {
		if err := db.store.InsertMediaItem(&mediaItem); err != nil {
			return fmt.Errorf("failed to restore media item: %w", err)
		}
	}

	// Restore tags
	for _, tag := range backupData.Tags {
		if err := db.store.InsertTag(tag.Name); err != nil {
			return fmt.Errorf("failed to restore tag: %w", err)
		}
	}

	return nil
}