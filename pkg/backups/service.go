// file: pkg/backups/service.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174006

package backups

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Service provides comprehensive backup and restore operations.
type Service struct {
	manager           *BackupManager
	dbBackupper       *DatabaseBackupper
	configBackupper   *ConfigBackupper
	subtitleBackupper *SubtitleBackupper
	logger            *logrus.Entry
}

// ServiceConfig holds configuration for the backup service.
type ServiceConfig struct {
	BackupPath         string
	EnableEncryption   bool
	EncryptionKey      []byte
	EnableCompression  bool
	DatabaseStore      database.SubtitleStore
	EnableCloudStorage bool
	CloudStorageType   string // "s3", "gcs", "multi"
}

// NewService creates a new backup service with the provided configuration.
func NewService(config ServiceConfig) (*Service, error) {
	logger := logging.GetLogger("backup")

	// Create storage
	var storage Storage
	if config.EnableCloudStorage && config.CloudStorageType != "" {
		// Try to create cloud storage
		cloudStorage, err := NewCloudStorageFromConfig()
		if err != nil {
			logger.Warnf("Failed to create cloud storage, falling back to local: %v", err)
			storage = NewLocalStorage(config.BackupPath)
		} else {
			// Combine local and cloud storage
			localStorage := NewLocalStorage(config.BackupPath)
			storage = NewMultiStorage(localStorage, cloudStorage)
			logger.Info("Using multi-storage (local + cloud)")
		}
	} else {
		storage = NewLocalStorage(config.BackupPath)
	}

	// Create compression if enabled
	var compression Compression
	if config.EnableCompression {
		compression = NewGzipCompression()
	}

	// Create encryption if enabled
	var encryption Encryption
	if config.EnableEncryption && len(config.EncryptionKey) > 0 {
		var err error
		encryption, err = NewAESEncryption(config.EncryptionKey)
		if err != nil {
			return nil, fmt.Errorf("failed to create encryption: %w", err)
		}
	}

	// Create backup manager
	manager := NewBackupManager(storage, compression, encryption)

	// Create specialized backuppers
	dbBackupper := NewDatabaseBackupper(config.DatabaseStore)
	configBackupper := NewConfigBackupper()
	subtitleBackupper := NewSubtitleBackupper(config.DatabaseStore)

	return &Service{
		manager:           manager,
		dbBackupper:       dbBackupper,
		configBackupper:   configBackupper,
		subtitleBackupper: subtitleBackupper,
		logger:            logger,
	}, nil
}

// CreateFullBackup creates a comprehensive backup of all system data.
func (s *Service) CreateFullBackup(ctx context.Context) (*Backup, error) {
	s.logger.Info("Starting full backup")

	backupData := make(map[string][]byte)
	contents := []string{}

	// Backup database
	s.logger.Info("Backing up database")
	dbData, err := s.dbBackupper.CreateDatabaseBackup(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create database backup: %w", err)
	}
	backupData["database.json"] = dbData
	contents = append(contents, "database")

	// Backup configuration
	s.logger.Info("Backing up configuration")
	configData, err := s.configBackupper.CreateConfigBackup(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create config backup: %w", err)
	}
	backupData["config.json"] = configData
	contents = append(contents, "configuration")

	// Backup config file if it exists
	if configFile := viper.ConfigFileUsed(); configFile != "" {
		s.logger.Infof("Backing up config file: %s", configFile)
		configFileData, err := s.configBackupper.BackupConfigFile(ctx, configFile)
		if err != nil {
			return nil, fmt.Errorf("failed to backup config file: %w", err)
		}
		if len(configFileData) > 0 {
			ext := filepath.Ext(configFile)
			backupData["config-file"+ext] = configFileData
			contents = append(contents, "config-file")
		}
	}

	// Backup subtitle files if enabled
	subtitlePaths := GetSubtitleBackupPaths()
	if len(subtitlePaths) > 0 {
		s.logger.Infof("Backing up subtitle files from %d paths", len(subtitlePaths))
		subtitleData, err := s.subtitleBackupper.CreateSubtitleBackup(ctx, subtitlePaths)
		if err != nil {
			return nil, fmt.Errorf("failed to backup subtitle files: %w", err)
		}
		backupData["subtitles.json"] = subtitleData
		contents = append(contents, "subtitles")
	}

	// Combine all backup data into a single archive
	combinedData, err := s.createArchive(backupData)
	if err != nil {
		return nil, fmt.Errorf("failed to create backup archive: %w", err)
	}

	// Create backup using manager
	backup, err := s.manager.CreateBackup(ctx, BackupTypeFull, contents, combinedData)
	if err != nil {
		return nil, fmt.Errorf("failed to store backup: %w", err)
	}

	s.logger.Infof("Full backup created successfully: %s", backup.ID)
	return backup, nil
}

// CreateDatabaseBackup creates a backup of only the database.
func (s *Service) CreateDatabaseBackup(ctx context.Context) (*Backup, error) {
	s.logger.Info("Starting database backup")

	dbData, err := s.dbBackupper.CreateDatabaseBackup(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create database backup: %w", err)
	}

	backup, err := s.manager.CreateBackup(ctx, BackupTypeDatabase, []string{"database"}, dbData)
	if err != nil {
		return nil, fmt.Errorf("failed to store database backup: %w", err)
	}

	s.logger.Infof("Database backup created successfully: %s", backup.ID)
	return backup, nil
}

// CreateConfigBackup creates a backup of only the configuration.
func (s *Service) CreateConfigBackup(ctx context.Context) (*Backup, error) {
	s.logger.Info("Starting configuration backup")

	configData, err := s.configBackupper.CreateConfigBackup(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create config backup: %w", err)
	}

	backup, err := s.manager.CreateBackup(ctx, BackupTypeConfiguration, []string{"configuration"}, configData)
	if err != nil {
		return nil, fmt.Errorf("failed to store config backup: %w", err)
	}

	s.logger.Infof("Configuration backup created successfully: %s", backup.ID)
	return backup, nil
}

// CreateSubtitleBackup creates a backup of only subtitle files.
func (s *Service) CreateSubtitleBackup(ctx context.Context) (*Backup, error) {
	s.logger.Info("Starting subtitle backup")

	subtitlePaths := GetSubtitleBackupPaths()
	if len(subtitlePaths) == 0 {
		return nil, fmt.Errorf("no subtitle paths configured for backup")
	}

	subtitleData, err := s.subtitleBackupper.CreateSubtitleBackup(ctx, subtitlePaths)
	if err != nil {
		return nil, fmt.Errorf("failed to create subtitle backup: %w", err)
	}

	backup, err := s.manager.CreateBackup(ctx, BackupTypeSubtitles, []string{"subtitles"}, subtitleData)
	if err != nil {
		return nil, fmt.Errorf("failed to store subtitle backup: %w", err)
	}

	s.logger.Infof("Subtitle backup created successfully: %s", backup.ID)
	return backup, nil
}

// RestoreBackup restores data from a backup.
func (s *Service) RestoreBackup(ctx context.Context, backupID string) error {
	s.logger.Infof("Starting restore from backup: %s", backupID)

	backup, exists := s.manager.GetBackup(backupID)
	if !exists {
		return fmt.Errorf("backup not found: %s", backupID)
	}

	data, err := s.manager.RestoreBackup(ctx, backupID)
	if err != nil {
		return fmt.Errorf("failed to retrieve backup data: %w", err)
	}

	switch backup.Type {
	case BackupTypeFull:
		return s.restoreFullBackup(ctx, data)
	case BackupTypeDatabase:
		return s.dbBackupper.RestoreDatabaseBackup(ctx, data)
	case BackupTypeConfiguration:
		return s.configBackupper.RestoreConfigBackup(ctx, data)
	case BackupTypeSubtitles:
		restorePath := viper.GetString("backup_subtitle_restore_path")
		if restorePath == "" {
			restorePath = "/config/restored-subtitles"
		}
		return s.subtitleBackupper.RestoreSubtitleBackup(ctx, data, restorePath)
	default:
		return fmt.Errorf("unsupported backup type: %s", backup.Type)
	}
}

// ListBackups returns all available backups.
func (s *Service) ListBackups() []*Backup {
	return s.manager.ListBackups()
}

// DeleteBackup removes a backup.
func (s *Service) DeleteBackup(ctx context.Context, backupID string) error {
	s.logger.Infof("Deleting backup: %s", backupID)
	return s.manager.DeleteBackup(ctx, backupID)
}

// RotateBackups removes old backups based on retention policy.
func (s *Service) RotateBackups(ctx context.Context, maxAge time.Duration, maxCount int) error {
	s.logger.Info("Starting backup rotation")

	backups := s.manager.ListBackups()

	// Sort backups by creation time (newest first)
	for i := 0; i < len(backups)-1; i++ {
		for j := i + 1; j < len(backups); j++ {
			if backups[i].CreatedAt.Before(backups[j].CreatedAt) {
				backups[i], backups[j] = backups[j], backups[i]
			}
		}
	}

	deleted := 0
	cutoff := time.Now().Add(-maxAge)

	for i, backup := range backups {
		shouldDelete := false

		// Delete if too old
		if backup.CreatedAt.Before(cutoff) {
			shouldDelete = true
		}

		// Delete if exceeds count limit (keep newest)
		if i >= maxCount {
			shouldDelete = true
		}

		if shouldDelete {
			if err := s.manager.DeleteBackup(ctx, backup.ID); err != nil {
				s.logger.Warnf("Failed to delete backup %s: %v", backup.ID, err)
			} else {
				deleted++
			}
		}
	}

	s.logger.Infof("Backup rotation completed, deleted %d backups", deleted)
	return nil
}

// Helper methods

func (s *Service) createArchive(data map[string][]byte) ([]byte, error) {
	// For now, just return the combined JSON data
	// In a more sophisticated implementation, this could create a tar/zip archive
	archive := make(map[string]interface{})
	for name, content := range data {
		archive[name] = string(content)
	}

	result, err := json.Marshal(archive)
	if err != nil {
		return nil, fmt.Errorf("failed to create archive: %w", err)
	}

	return result, nil
}

func (s *Service) restoreFullBackup(ctx context.Context, data []byte) error {
	var archive map[string]interface{}
	if err := json.Unmarshal(data, &archive); err != nil {
		return fmt.Errorf("failed to parse backup archive: %w", err)
	}

	// Restore database if present
	if dbDataStr, ok := archive["database.json"].(string); ok {
		if err := s.dbBackupper.RestoreDatabaseBackup(ctx, []byte(dbDataStr)); err != nil {
			return fmt.Errorf("failed to restore database: %w", err)
		}
		s.logger.Info("Database restored successfully")
	}

	// Restore configuration if present
	if configDataStr, ok := archive["config.json"].(string); ok {
		if err := s.configBackupper.RestoreConfigBackup(ctx, []byte(configDataStr)); err != nil {
			return fmt.Errorf("failed to restore configuration: %w", err)
		}
		s.logger.Info("Configuration restored successfully")
	}

	// Restore config file if present
	for name, content := range archive {
		if name == "config-file.yaml" || name == "config-file.yml" {
			if contentStr, ok := content.(string); ok {
				configFile := viper.ConfigFileUsed()
				if configFile == "" {
					// Use default config file location
					configFile = filepath.Join(os.Getenv("HOME"), ".subtitle-manager.yaml")
				}
				if err := s.configBackupper.RestoreConfigFile(ctx, configFile, []byte(contentStr)); err != nil {
					return fmt.Errorf("failed to restore config file: %w", err)
				}
				s.logger.Info("Config file restored successfully")
			}
		}
	}

	// Restore subtitles if present
	if subtitleDataStr, ok := archive["subtitles.json"].(string); ok {
		restorePath := viper.GetString("backup_subtitle_restore_path")
		if restorePath == "" {
			restorePath = "/config/restored-subtitles"
		}
		if err := s.subtitleBackupper.RestoreSubtitleBackup(ctx, []byte(subtitleDataStr), restorePath); err != nil {
			return fmt.Errorf("failed to restore subtitles: %w", err)
		}
		s.logger.Info("Subtitles restored successfully")
	}

	return nil
}
