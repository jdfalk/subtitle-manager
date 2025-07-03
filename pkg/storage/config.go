// file: pkg/storage/config.go
// version: 1.0.0
// guid: 78901234-56f1-789a-bcde-f123456789ab

package storage

import (
	"github.com/spf13/viper"
)

// GetConfigFromViper extracts storage configuration from viper settings.
func GetConfigFromViper() StorageConfig {
	return StorageConfig{
		Provider:        viper.GetString("storage.provider"),
		LocalPath:       viper.GetString("storage.local_path"),
		S3Region:        viper.GetString("storage.s3_region"),
		S3Bucket:        viper.GetString("storage.s3_bucket"),
		S3Endpoint:      viper.GetString("storage.s3_endpoint"),
		S3AccessKey:     viper.GetString("storage.s3_access_key"),
		S3SecretKey:     viper.GetString("storage.s3_secret_key"),
		AzureAccount:    viper.GetString("storage.azure_account"),
		AzureKey:        viper.GetString("storage.azure_key"),
		AzureContainer:  viper.GetString("storage.azure_container"),
		GCSBucket:       viper.GetString("storage.gcs_bucket"),
		GCSCredentials:  viper.GetString("storage.gcs_credentials"),
		EnableBackup:    viper.GetBool("storage.enable_backup"),
		BackupHistory:   viper.GetBool("storage.backup_history"),
		CompressionType: viper.GetString("storage.compression"),
	}
}
