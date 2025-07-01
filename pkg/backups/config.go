// file: pkg/backups/config.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174005

package backups

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// ConfigBackupper provides configuration backup and restore functionality.
type ConfigBackupper struct{}

// NewConfigBackupper creates a new configuration backup instance.
func NewConfigBackupper() *ConfigBackupper {
	return &ConfigBackupper{}
}

// CreateConfigBackup creates a backup of the current configuration.
func (cb *ConfigBackupper) CreateConfigBackup(ctx context.Context) ([]byte, error) {
	// Get all configuration settings from viper
	settings := viper.AllSettings()

	// Serialize to JSON
	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to serialize config data: %w", err)
	}

	return data, nil
}

// RestoreConfigBackup restores configuration from backup data.
func (cb *ConfigBackupper) RestoreConfigBackup(ctx context.Context, data []byte) error {
	var settings map[string]interface{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return fmt.Errorf("failed to deserialize config data: %w", err)
	}

	// Restore all settings to viper
	for key, value := range settings {
		viper.Set(key, value)
	}

	return nil
}

// BackupConfigFile creates a backup of the configuration file itself.
func (cb *ConfigBackupper) BackupConfigFile(ctx context.Context, configFilePath string) ([]byte, error) {
	if configFilePath == "" {
		// If no config file path provided, return empty backup
		return []byte{}, nil
	}

	// Check if config file exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// Config file doesn't exist, return empty backup
		return []byte{}, nil
	}

	// Read the config file
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return data, nil
}

// RestoreConfigFile restores a configuration file from backup data.
func (cb *ConfigBackupper) RestoreConfigFile(ctx context.Context, configFilePath string, data []byte) error {
	if configFilePath == "" {
		return fmt.Errorf("config file path not provided")
	}

	// Create directory if it doesn't exist
	if dir := filepath.Dir(configFilePath); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	// Write the config file
	file, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
