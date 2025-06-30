// file: pkg/backups/scheduler.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174009

package backups

import (
	"context"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/scheduler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ScheduledBackupService manages scheduled backup operations.
type ScheduledBackupService struct {
	service *Service
	logger  *logrus.Entry
}

// NewScheduledBackupService creates a new scheduled backup service.
func NewScheduledBackupService(service *Service) *ScheduledBackupService {
	return &ScheduledBackupService{
		service: service,
		logger:  logging.GetLogger("backup-scheduler"),
	}
}

// StartScheduledBackups starts the backup scheduler based on configuration.
func (s *ScheduledBackupService) StartScheduledBackups(ctx context.Context) error {
	// Get backup schedule from configuration
	backupFrequency := viper.GetString("backup_frequency")
	if backupFrequency == "" {
		backupFrequency = "daily" // default
	}

	// Parse frequency
	interval, err := parseFrequency(backupFrequency)
	if err != nil {
		s.logger.Warnf("Invalid backup frequency '%s', using daily: %v", backupFrequency, err)
		interval = 24 * time.Hour
	}

	if interval == 0 {
		s.logger.Info("Backup scheduling disabled")
		return nil
	}

	s.logger.Infof("Starting backup scheduler with frequency: %s", backupFrequency)

	// Start backup scheduler
	go func() {
		err := scheduler.Run(ctx, interval, s.performScheduledBackup)
		if err != nil && ctx.Err() == nil {
			s.logger.Errorf("Backup scheduler error: %v", err)
		}
	}()

	// Start backup rotation scheduler (daily)
	go func() {
		err := scheduler.Run(ctx, 24*time.Hour, s.performBackupRotation)
		if err != nil && ctx.Err() == nil {
			s.logger.Errorf("Backup rotation scheduler error: %v", err)
		}
	}()

	return nil
}

// performScheduledBackup performs a scheduled backup operation.
func (s *ScheduledBackupService) performScheduledBackup(ctx context.Context) error {
	s.logger.Info("Starting scheduled backup")

	backupType := viper.GetString("backup_type")
	if backupType == "" {
		backupType = "full"
	}

	var backup *Backup
	var err error

	switch backupType {
	case "database":
		backup, err = s.service.CreateDatabaseBackup(ctx)
	case "configuration":
		backup, err = s.service.CreateConfigBackup(ctx)
	case "full":
		backup, err = s.service.CreateFullBackup(ctx)
	default:
		s.logger.Warnf("Unknown backup type '%s', using full backup", backupType)
		backup, err = s.service.CreateFullBackup(ctx)
	}

	if err != nil {
		s.logger.Errorf("Scheduled backup failed: %v", err)
		return err
	}

	s.logger.Infof("Scheduled backup completed successfully: %s", backup.ID)
	return nil
}

// performBackupRotation performs backup cleanup based on retention policy.
func (s *ScheduledBackupService) performBackupRotation(ctx context.Context) error {
	s.logger.Info("Starting backup rotation")

	// Get rotation settings from configuration
	maxAgeDays := viper.GetInt("backup_retention_days")
	if maxAgeDays <= 0 {
		maxAgeDays = 30 // default
	}

	maxCount := viper.GetInt("backup_retention_count")
	if maxCount <= 0 {
		maxCount = 10 // default
	}

	maxAge := time.Duration(maxAgeDays) * 24 * time.Hour
	err := s.service.RotateBackups(ctx, maxAge, maxCount)
	if err != nil {
		s.logger.Errorf("Backup rotation failed: %v", err)
		return err
	}

	s.logger.Info("Backup rotation completed successfully")
	return nil
}

// parseFrequency converts a frequency string to a time.Duration.
func parseFrequency(frequency string) (time.Duration, error) {
	switch frequency {
	case "disabled", "never":
		return 0, nil
	case "hourly":
		return time.Hour, nil
	case "daily":
		return 24 * time.Hour, nil
	case "weekly":
		return 7 * 24 * time.Hour, nil
	case "monthly":
		return 30 * 24 * time.Hour, nil
	default:
		// Try to parse as duration
		return time.ParseDuration(frequency)
	}
}