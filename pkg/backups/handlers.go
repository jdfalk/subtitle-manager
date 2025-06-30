// file: pkg/backups/handlers.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174008

package backups

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Handlers provides HTTP handlers for backup operations.
type Handlers struct {
	service *Service
	logger  *logrus.Entry
}

// NewHandlers creates new backup handlers with the provided service.
func NewHandlers(service *Service) *Handlers {
	return &Handlers{
		service: service,
		logger:  logging.GetLogger("backup-handlers"),
	}
}

// CreateBackupHandler creates a new backup based on the type specified.
func (h *Handlers) CreateBackupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		backupType := r.URL.Query().Get("type")
		if backupType == "" {
			backupType = "full"
		}

		var backup *Backup
		var err error

		switch backupType {
		case "database":
			backup, err = h.service.CreateDatabaseBackup(r.Context())
		case "configuration":
			backup, err = h.service.CreateConfigBackup(r.Context())
		case "full":
			backup, err = h.service.CreateFullBackup(r.Context())
		default:
			http.Error(w, "Invalid backup type", http.StatusBadRequest)
			return
		}

		if err != nil {
			h.logger.Errorf("Failed to create backup: %v", err)
			http.Error(w, "Failed to create backup", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(backup)
	})
}

// ListBackupsHandler returns all available backups.
func (h *Handlers) ListBackupsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backups := h.service.ListBackups()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(backups)
	})
}

// RestoreBackupHandler restores a specific backup.
func (h *Handlers) RestoreBackupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		backupID := r.URL.Query().Get("id")
		if backupID == "" {
			http.Error(w, "Backup ID required", http.StatusBadRequest)
			return
		}

		err := h.service.RestoreBackup(r.Context(), backupID)
		if err != nil {
			h.logger.Errorf("Failed to restore backup %s: %v", backupID, err)
			http.Error(w, "Failed to restore backup", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":   "success",
			"message":  "Backup restored successfully",
			"backupId": backupID,
		})
	})
}

// DeleteBackupHandler deletes a specific backup.
func (h *Handlers) DeleteBackupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		backupID := r.URL.Query().Get("id")
		if backupID == "" {
			http.Error(w, "Backup ID required", http.StatusBadRequest)
			return
		}

		err := h.service.DeleteBackup(r.Context(), backupID)
		if err != nil {
			h.logger.Errorf("Failed to delete backup %s: %v", backupID, err)
			http.Error(w, "Failed to delete backup", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":   "success",
			"message":  "Backup deleted successfully",
			"backupId": backupID,
		})
	})
}

// RotateBackupsHandler triggers backup rotation.
func (h *Handlers) RotateBackupsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Parse rotation parameters
		maxAgeStr := r.URL.Query().Get("max_age_days")
		maxCountStr := r.URL.Query().Get("max_count")

		maxAgeDays := 30 // default
		if maxAgeStr != "" {
			if days, err := strconv.Atoi(maxAgeStr); err == nil && days > 0 {
				maxAgeDays = days
			}
		}

		maxCount := 10 // default
		if maxCountStr != "" {
			if count, err := strconv.Atoi(maxCountStr); err == nil && count > 0 {
				maxCount = count
			}
		}

		maxAge := time.Duration(maxAgeDays) * 24 * time.Hour
		err := h.service.RotateBackups(r.Context(), maxAge, maxCount)
		if err != nil {
			h.logger.Errorf("Failed to rotate backups: %v", err)
			http.Error(w, "Failed to rotate backups", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":      "success",
			"message":     "Backup rotation completed",
			"maxAgeDays":  maxAgeDays,
			"maxCount":    maxCount,
		})
	})
}

// NewServiceFromConfig creates a backup service from the current configuration.
func NewServiceFromConfig(store database.SubtitleStore) (*Service, error) {
	// Get backup path from config
	backupPath := viper.GetString("backup_path")
	if backupPath == "" {
		// Default backup path
		dbPath := viper.GetString("db_path")
		if dbPath != "" {
			backupPath = filepath.Join(filepath.Dir(dbPath), "backups")
		} else {
			backupPath = "/config/backups"
		}
	}

	// Get encryption settings
	enableEncryption := viper.GetBool("backup_encryption_enabled")
	encryptionKey := viper.GetString("backup_encryption_key")

	var encryptionKeyBytes []byte
	if enableEncryption && encryptionKey != "" {
		// For demo purposes, use a simple key derivation
		// In production, use proper key derivation like PBKDF2
		key := []byte(encryptionKey)
		if len(key) < 32 {
			// Pad key to 32 bytes for AES-256
			padded := make([]byte, 32)
			copy(padded, key)
			encryptionKeyBytes = padded
		} else {
			encryptionKeyBytes = key[:32]
		}
	}

	config := ServiceConfig{
		BackupPath:        backupPath,
		EnableEncryption:  enableEncryption,
		EncryptionKey:     encryptionKeyBytes,
		EnableCompression: viper.GetBool("backup_compression_enabled"),
		DatabaseStore:     store,
	}

	return NewService(config)
}