# Automatic Episode Monitoring System

The subtitle-manager now includes a comprehensive automatic episode monitoring
system that can monitor your Sonarr/Radarr library for new episodes and
automatically download subtitles when they become available.

## Features

- **Automatic Detection**: Monitors Sonarr/Radarr for new episodes and movies
- **Multi-language Support**: Track subtitle availability for multiple languages
  simultaneously
- **Provider Integration**: Uses existing provider priority system for downloads
- **Retry Logic**: Configurable retry attempts with exponential backoff
- **Blacklist Management**: Automatic and manual blacklisting of problematic
  items
- **Quality Upgrades**: Monitor for better quality subtitles
- **Statistics**: Detailed monitoring statistics and progress reporting
- **Scheduler Integration**: Periodic monitoring with configurable intervals

## Quick Start

### 1. Sync Media Library

First, sync your media library from Sonarr/Radarr to populate the monitoring
database:

```bash
# Sync from both Sonarr and Radarr
subtitle-manager monitor sync --languages=en,es --max-retries=3

# Sync only from Sonarr
subtitle-manager monitor sync --source=sonarr --languages=en

# Force refresh existing items
subtitle-manager monitor sync --force-refresh
```

### 2. Check Monitoring Status

View current monitoring statistics:

```bash
subtitle-manager monitor status
```

Example output:

```
Monitoring Status:
  Total items:    150
  Pending:        45
  Monitoring:     80
  Found:          20
  Failed:         3
  Blacklisted:    2
```

### 3. Start Monitoring Daemon

Start the monitoring daemon to automatically check for subtitles:

```bash
# Monitor every hour with quality upgrades enabled
subtitle-manager monitor start --interval=1h --quality-check

# Monitor every 30 minutes with custom retry limit
subtitle-manager monitor start --interval=30m --max-retries=5
```

### 4. List Monitored Items

View all items currently being monitored:

```bash
subtitle-manager monitor list
```

### 5. Blacklist Management

List blacklisted items:

```bash
subtitle-manager monitor blacklist list
```

Remove an item from blacklist:

```bash
subtitle-manager monitor blacklist remove <item-id>
```

## Configuration

The monitoring system integrates with your existing Sonarr/Radarr configuration.
Make sure you have the following configured in your config file:

```yaml
# Sonarr configuration
sonarr_url: 'http://localhost:8989'
sonarr_api_key: 'your-sonarr-api-key'

# Radarr configuration
radarr_url: 'http://localhost:7878'
radarr_api_key: 'your-radarr-api-key'

# Database configuration
db_backend: 'pebble' # or "sqlite", "postgres"
db_path: '/path/to/database'
```

## Monitoring Workflow

1. **Sync Phase**: The system syncs your media library from Sonarr/Radarr
2. **Detection Phase**: New episodes/movies are added to the monitoring queue
3. **Processing Phase**: The daemon periodically checks for available subtitles
4. **Download Phase**: When subtitles are found, they're automatically
   downloaded
5. **Retry Phase**: Failed attempts are retried with exponential backoff
6. **Blacklist Phase**: Items exceeding retry limits are automatically
   blacklisted

## Provider Integration

The monitoring system leverages the existing provider system:

- Uses provider priority ordering
- Respects provider backoff timers
- Supports all configured subtitle providers
- Integrates with provider API rate limiting

## Quality Upgrades

When quality checking is enabled, the system will:

- Compare new subtitle file sizes with existing ones
- Download larger (presumably better quality) subtitles
- Replace existing subtitles only when improvements are found

## Blacklist Management

Items can be blacklisted for various reasons:

- **Maximum Retries Exceeded**: Automatic blacklisting after retry limit
- **Manual Blacklist**: User-initiated blacklisting
- **No Subtitles Found**: Long-term unavailability
- **Quality Issues**: Poor quality or corrupted subtitles
- **Provider Errors**: Persistent provider failures

Blacklisted items can be manually restored to monitoring at any time.

## Statistics and Reporting

The monitoring system provides detailed statistics:

- **Total Items**: Number of items being monitored
- **Pending**: Items awaiting first check
- **Monitoring**: Items actively being checked
- **Found**: Items with subtitles successfully downloaded
- **Failed**: Items that have exceeded retry limits
- **Blacklisted**: Items excluded from monitoring

## Performance Considerations

- Monitoring interval should balance responsiveness with system load
- Consider provider rate limits when setting check frequencies
- Database backend choice affects performance at scale
- Jitter is automatically applied to prevent synchronized spikes

## Troubleshooting

### No items found during sync

- Verify Sonarr/Radarr API connectivity
- Check API keys and URLs in configuration
- Ensure media files exist on disk

### Monitoring daemon stops

- Check logs for provider errors
- Verify database connectivity
- Monitor system resources

### Subtitles not downloading

- Check provider configuration and API keys
- Verify subtitle providers are enabled
- Review blacklist status of items

### High retry rates

- Check provider reliability
- Consider adjusting retry limits
- Review provider rate limiting
