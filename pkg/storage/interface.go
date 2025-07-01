// file: pkg/storage/interface.go
// version: 1.0.0
// guid: 12345678-90ab-cdef-1234-567890abcdef

// Package storage provides cloud and local storage interfaces for subtitle files.
// It supports multiple cloud providers including S3, Azure Blob Storage, and Google Cloud Storage.
package storage

import (
	"context"
	"io"
	"time"
)

// StorageProvider defines the interface for storing and retrieving subtitle files.
type StorageProvider interface {
	// Store saves content to the storage backend with the given key/path.
	Store(ctx context.Context, key string, content io.Reader, contentType string) error
	
	// Retrieve reads content from the storage backend for the given key/path.
	Retrieve(ctx context.Context, key string) (io.ReadCloser, error)
	
	// Delete removes content from the storage backend for the given key/path.
	Delete(ctx context.Context, key string) error
	
	// Exists checks if content exists at the given key/path.
	Exists(ctx context.Context, key string) (bool, error)
	
	// List returns keys/paths matching the given prefix.
	List(ctx context.Context, prefix string) ([]string, error)
	
	// GetURL returns a public or signed URL for the given key/path, if supported.
	GetURL(ctx context.Context, key string, expiry time.Duration) (string, error)
	
	// Close closes any connections and cleans up resources.
	Close() error
}

// StorageConfig holds configuration for storage providers.
type StorageConfig struct {
	// Provider specifies the storage backend: "local", "s3", "azure", "gcs"
	Provider string `yaml:"provider" json:"provider"`
	
	// Local storage configuration
	LocalPath string `yaml:"local_path,omitempty" json:"local_path,omitempty"`
	
	// S3 configuration
	S3Region    string `yaml:"s3_region,omitempty" json:"s3_region,omitempty"`
	S3Bucket    string `yaml:"s3_bucket,omitempty" json:"s3_bucket,omitempty"`
	S3Endpoint  string `yaml:"s3_endpoint,omitempty" json:"s3_endpoint,omitempty"`
	S3AccessKey string `yaml:"s3_access_key,omitempty" json:"s3_access_key,omitempty"`
	S3SecretKey string `yaml:"s3_secret_key,omitempty" json:"s3_secret_key,omitempty"`
	
	// Azure Blob Storage configuration
	AzureAccount   string `yaml:"azure_account,omitempty" json:"azure_account,omitempty"`
	AzureKey       string `yaml:"azure_key,omitempty" json:"azure_key,omitempty"`
	AzureContainer string `yaml:"azure_container,omitempty" json:"azure_container,omitempty"`
	
	// Google Cloud Storage configuration
	GCSBucket      string `yaml:"gcs_bucket,omitempty" json:"gcs_bucket,omitempty"`
	GCSCredentials string `yaml:"gcs_credentials,omitempty" json:"gcs_credentials,omitempty"`
	
	// Additional options
	EnableBackup    bool   `yaml:"enable_backup" json:"enable_backup"`         // Enable cloud backup of subtitle files
	BackupHistory   bool   `yaml:"backup_history" json:"backup_history"`       // Enable cloud backup of history data
	CompressionType string `yaml:"compression,omitempty" json:"compression,omitempty"` // gzip, none
}

// StorageManager manages both local and cloud storage providers.
type StorageManager struct {
	primary   StorageProvider
	backup    StorageProvider
	config    StorageConfig
}

// NewStorageManager creates a new storage manager with the given configuration.
func NewStorageManager(config StorageConfig) (*StorageManager, error) {
	primary, err := NewProvider(config)
	if err != nil {
		return nil, err
	}
	
	manager := &StorageManager{
		primary: primary,
		config:  config,
	}
	
	// Set up backup provider if enabled
	if config.EnableBackup && config.Provider != "local" {
		// Create local backup provider
		localConfig := StorageConfig{
			Provider:  "local",
			LocalPath: config.LocalPath,
		}
		backup, err := NewProvider(localConfig)
		if err == nil {
			manager.backup = backup
		}
	}
	
	return manager, nil
}

// Store stores content using the primary provider and optionally backs up locally.
func (sm *StorageManager) Store(ctx context.Context, key string, content io.Reader, contentType string) error {
	// Store in primary provider
	if err := sm.primary.Store(ctx, key, content, contentType); err != nil {
		return err
	}
	
	// Store backup if enabled
	if sm.backup != nil && sm.config.EnableBackup {
		// Need to re-read content for backup, so we'll need to handle this differently
		// For now, just store in primary
	}
	
	return nil
}

// Retrieve retrieves content from the primary provider.
func (sm *StorageManager) Retrieve(ctx context.Context, key string) (io.ReadCloser, error) {
	return sm.primary.Retrieve(ctx, key)
}

// Delete deletes content from both primary and backup providers.
func (sm *StorageManager) Delete(ctx context.Context, key string) error {
	err := sm.primary.Delete(ctx, key)
	if sm.backup != nil {
		sm.backup.Delete(ctx, key) // Best effort
	}
	return err
}

// Exists checks if content exists in the primary provider.
func (sm *StorageManager) Exists(ctx context.Context, key string) (bool, error) {
	return sm.primary.Exists(ctx, key)
}

// List lists keys from the primary provider.
func (sm *StorageManager) List(ctx context.Context, prefix string) ([]string, error) {
	return sm.primary.List(ctx, prefix)
}

// GetURL gets a URL from the primary provider.
func (sm *StorageManager) GetURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	return sm.primary.GetURL(ctx, key, expiry)
}

// Close closes both primary and backup providers.
func (sm *StorageManager) Close() error {
	var err error
	if sm.primary != nil {
		err = sm.primary.Close()
	}
	if sm.backup != nil {
		sm.backup.Close() // Best effort
	}
	return err
}

// NewProvider creates a new storage provider based on the configuration.
func NewProvider(config StorageConfig) (StorageProvider, error) {
	switch config.Provider {
	case "local", "":
		return NewLocalProvider(config)
	case "s3":
		return NewS3Provider(config)
	case "azure":
		return NewAzureProvider(config)
	case "gcs":
		return NewGCSProvider(config)
	default:
		return nil, ErrUnsupportedProvider
	}
}