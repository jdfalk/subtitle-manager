# Cloud Storage for Subtitles

Subtitle Manager now supports optional cloud storage for subtitle files and history data. This allows you to store subtitles in cloud providers such as Amazon S3, Azure Blob Storage, and Google Cloud Storage.

## Supported Providers

- **Local filesystem** (default) - Store files locally
- **Amazon S3** - Store files in Amazon S3 or S3-compatible services
- **Azure Blob Storage** - Store files in Microsoft Azure (coming soon)
- **Google Cloud Storage** - Store files in Google Cloud (coming soon)

## Configuration

### Command Line Flags

```bash
# Basic storage provider selection
--storage-provider string           # Storage provider: local, s3, azure, gcs (default "local")
--storage-local-path string         # Local storage path (default "subtitles")

# S3 configuration
--s3-region string                  # S3 region
--s3-bucket string                  # S3 bucket name
--s3-endpoint string                # S3 endpoint URL (for S3-compatible services)
--s3-access-key string              # S3 access key
--s3-secret-key string              # S3 secret key

# Azure configuration (coming soon)
--azure-account string              # Azure storage account name
--azure-key string                  # Azure storage account key
--azure-container string            # Azure blob container name

# GCS configuration (coming soon)
--gcs-bucket string                 # Google Cloud Storage bucket name
--gcs-credentials string            # Google Cloud credentials JSON file path

# Storage options
--storage-enable-backup             # Enable cloud backup of subtitle files
--storage-backup-history            # Enable cloud backup of history data
```

### Environment Variables

All flags can also be set via environment variables with the `SM_` prefix:

```bash
export SM_STORAGE_PROVIDER=s3
export SM_S3_REGION=us-west-2
export SM_S3_BUCKET=my-subtitles-bucket
export SM_S3_ACCESS_KEY=your-access-key
export SM_S3_SECRET_KEY=your-secret-key
```

### Configuration File

You can also configure cloud storage in your `~/.subtitle-manager.yaml` file:

```yaml
storage:
  provider: s3
  local_path: subtitles
  s3_region: us-west-2
  s3_bucket: my-subtitles-bucket
  s3_access_key: your-access-key
  s3_secret_key: your-secret-key
  enable_backup: true
  backup_history: false
```

## Usage

### Testing Storage Connection

Test your cloud storage configuration:

```bash
subtitle-manager storage test
```

### Managing Files

Upload a subtitle file:
```bash
subtitle-manager storage upload local-file.srt remote/path/file.srt
```

Download a subtitle file:
```bash
subtitle-manager storage download remote/path/file.srt local-file.srt
```

List files in storage:
```bash
subtitle-manager storage list
subtitle-manager storage list movies/  # List files with prefix
```

## S3 Configuration Examples

### Amazon S3

```bash
subtitle-manager --storage-provider s3 \
  --s3-region us-west-2 \
  --s3-bucket my-subtitles \
  --s3-access-key AKIA... \
  --s3-secret-key secret... \
  storage test
```

### MinIO (Self-hosted S3-compatible)

```bash
subtitle-manager --storage-provider s3 \
  --s3-endpoint http://localhost:9000 \
  --s3-region us-east-1 \
  --s3-bucket subtitles \
  --s3-access-key minioadmin \
  --s3-secret-key minioadmin \
  storage test
```

### DigitalOcean Spaces

```bash
subtitle-manager --storage-provider s3 \
  --s3-endpoint https://nyc3.digitaloceanspaces.com \
  --s3-region nyc3 \
  --s3-bucket my-spaces-bucket \
  --s3-access-key your-access-key \
  --s3-secret-key your-secret-key \
  storage test
```

## Security Considerations

- Store credentials securely using environment variables or configuration files with proper permissions
- Use IAM roles and least-privilege access when possible
- Consider encryption at rest and in transit
- Regularly rotate access keys
- Monitor access logs for unusual activity

## Integration with Existing Features

When cloud storage is configured, subtitle-manager will:

1. **Automatically store downloaded subtitles** in the cloud provider
2. **Maintain local copies** if backup is enabled
3. **Sync history data** to cloud storage if backup_history is enabled
4. **Provide URLs** for direct access to subtitle files (where supported)

## Error Handling

The storage system includes robust error handling:

- Connection failures are logged and retried
- Invalid file paths are rejected for security
- Missing files return appropriate error messages
- Network timeouts are handled gracefully

## Future Enhancements

Planned features include:

- Full Azure Blob Storage implementation
- Full Google Cloud Storage implementation
- Compression support for stored files
- Automatic synchronization between local and cloud storage
- Webhook notifications for storage events
- Storage usage metrics and monitoring

## Troubleshooting

### Common Issues

1. **Connection failed**: Check your credentials and network connectivity
2. **Bucket not found**: Ensure the bucket exists and you have access
3. **Permission denied**: Verify your access keys have the required permissions
4. **Invalid key**: Check that file paths don't contain invalid characters

### Debug Mode

Enable debug logging to troubleshoot issues:

```bash
subtitle-manager --log-level debug storage test
```

This will show detailed information about the storage operations and any errors encountered.