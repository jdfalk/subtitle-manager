// file: pkg/storage/config_test.go
// version: 1.0.0
// guid: e7f6c5d4-3b2a-1c0f-9e8d-7c6b5a4f3e2d

package storage

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetConfigFromViper_WithDefaults(t *testing.T) {
	// Arrange
	viper.Reset()
	t.Cleanup(viper.Reset)

	// Act
	config := GetConfigFromViper()

	// Assert
	if config.Provider != "" {
		t.Fatalf("expected empty provider, got %q", config.Provider)
	}
	if config.LocalPath != "" {
		t.Fatalf("expected empty local path, got %q", config.LocalPath)
	}
	if config.EnableBackup {
		t.Fatalf("expected EnableBackup to be false by default")
	}
	if config.BackupHistory {
		t.Fatalf("expected BackupHistory to be false by default")
	}
	if config.CompressionType != "" {
		t.Fatalf("expected empty compression type, got %q", config.CompressionType)
	}
}

func TestGetConfigFromViper_WithConfiguredValues(t *testing.T) {
	// Arrange
	viper.Reset()
	t.Cleanup(viper.Reset)

	viper.Set("storage.provider", "s3")
	viper.Set("storage.local_path", "/tmp/subtitles")
	viper.Set("storage.s3_region", "us-west-2")
	viper.Set("storage.s3_bucket", "subtitle-bucket")
	viper.Set("storage.s3_endpoint", "https://s3.example.com")
	viper.Set("storage.s3_access_key", "access-key")
	viper.Set("storage.s3_secret_key", "secret-key")
	viper.Set("storage.azure_account", "azure-account")
	viper.Set("storage.azure_key", "azure-key")
	viper.Set("storage.azure_container", "azure-container")
	viper.Set("storage.gcs_bucket", "gcs-bucket")
	viper.Set("storage.gcs_credentials", "gcs-creds")
	viper.Set("storage.enable_backup", true)
	viper.Set("storage.backup_history", true)
	viper.Set("storage.compression", "gzip")

	// Act
	config := GetConfigFromViper()

	// Assert
	if config.Provider != "s3" {
		t.Fatalf("expected provider %q, got %q", "s3", config.Provider)
	}
	if config.LocalPath != "/tmp/subtitles" {
		t.Fatalf("expected local path %q, got %q", "/tmp/subtitles", config.LocalPath)
	}
	if config.S3Region != "us-west-2" {
		t.Fatalf("expected s3 region %q, got %q", "us-west-2", config.S3Region)
	}
	if config.S3Bucket != "subtitle-bucket" {
		t.Fatalf("expected s3 bucket %q, got %q", "subtitle-bucket", config.S3Bucket)
	}
	if config.S3Endpoint != "https://s3.example.com" {
		t.Fatalf("expected s3 endpoint %q, got %q", "https://s3.example.com", config.S3Endpoint)
	}
	if config.S3AccessKey != "access-key" {
		t.Fatalf("expected s3 access key %q, got %q", "access-key", config.S3AccessKey)
	}
	if config.S3SecretKey != "secret-key" {
		t.Fatalf("expected s3 secret key %q, got %q", "secret-key", config.S3SecretKey)
	}
	if config.AzureAccount != "azure-account" {
		t.Fatalf("expected azure account %q, got %q", "azure-account", config.AzureAccount)
	}
	if config.AzureKey != "azure-key" {
		t.Fatalf("expected azure key %q, got %q", "azure-key", config.AzureKey)
	}
	if config.AzureContainer != "azure-container" {
		t.Fatalf("expected azure container %q, got %q", "azure-container", config.AzureContainer)
	}
	if config.GCSBucket != "gcs-bucket" {
		t.Fatalf("expected gcs bucket %q, got %q", "gcs-bucket", config.GCSBucket)
	}
	if config.GCSCredentials != "gcs-creds" {
		t.Fatalf("expected gcs credentials %q, got %q", "gcs-creds", config.GCSCredentials)
	}
	if !config.EnableBackup {
		t.Fatalf("expected EnableBackup to be true")
	}
	if !config.BackupHistory {
		t.Fatalf("expected BackupHistory to be true")
	}
	if config.CompressionType != "gzip" {
		t.Fatalf("expected compression type %q, got %q", "gzip", config.CompressionType)
	}
}
