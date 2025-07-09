# file: docs/api/webhooks.md

# version: 1.0.0

# guid: 550e8400-e29b-41d4-a716-446655440021

# Webhook Integration Guide

Subtitle Manager supports webhook integrations with popular media management
systems like Sonarr and Radarr, as well as custom webhook endpoints for
automation and integration with external systems.

## Overview

Webhooks allow external systems to notify Subtitle Manager about media events,
triggering automatic subtitle downloads and processing. This enables seamless
automation of subtitle management within your media pipeline.

### Supported Webhook Types

- **Sonarr**: TV series management integration
- **Radarr**: Movie management integration
- **Custom**: Generic webhook endpoint for custom integrations

## Webhook Endpoints

### Sonarr Integration

**Endpoint**: `POST /api/webhooks/sonarr`

Handles notifications from Sonarr when episodes are downloaded or upgraded.

#### Configuration in Sonarr

1. Go to **Settings → Connect** in Sonarr
2. Add a new **Webhook** connection
3. Configure the following settings:
   - **Name**: Subtitle Manager
   - **URL**: `http://your-subtitle-manager:8080/api/webhooks/sonarr`
   - **Method**: POST
   - **Triggers**: Check "On Download" and "On Upgrade"

#### Webhook Payload

Sonarr sends the following payload structure:

```json
{
  "eventType": "Download",
  "series": {
    "id": 1,
    "title": "Example Series",
    "path": "/tv/Example Series",
    "tvdbId": 12345
  },
  "episodes": [
    {
      "id": 123,
      "episodeNumber": 1,
      "seasonNumber": 1,
      "title": "Pilot",
      "airDate": "2024-01-01",
      "episodeFile": {
        "id": 456,
        "path": "/tv/Example Series/Season 01/S01E01 - Pilot.mkv",
        "quality": "Bluray-1080p",
        "size": 2147483648
      }
    }
  ],
  "isUpgrade": false
}
```

#### Processing Logic

When a Sonarr webhook is received:

1. **Extract Episode Information**: Parse series title, season, episode number
2. **Locate Media File**: Use the episode file path to identify the media
3. **Download Subtitles**: Attempt to download subtitles for configured
   languages
4. **Store Results**: Log the operation in the history database

### Radarr Integration

**Endpoint**: `POST /api/webhooks/radarr`

Handles notifications from Radarr when movies are downloaded or upgraded.

#### Configuration in Radarr

1. Go to **Settings → Connect** in Radarr
2. Add a new **Webhook** connection
3. Configure the following settings:
   - **Name**: Subtitle Manager
   - **URL**: `http://your-subtitle-manager:8080/api/webhooks/radarr`
   - **Method**: POST
   - **Triggers**: Check "On Download" and "On Upgrade"

#### Webhook Payload

Radarr sends the following payload structure:

```json
{
  "eventType": "Download",
  "movie": {
    "id": 1,
    "title": "Example Movie",
    "year": 2024,
    "path": "/movies/Example Movie (2024)",
    "tmdbId": 12345,
    "imdbId": "tt1234567"
  },
  "movieFile": {
    "id": 456,
    "path": "/movies/Example Movie (2024)/Example Movie (2024) Bluray-1080p.mkv",
    "quality": "Bluray-1080p",
    "size": 8589934592
  },
  "isUpgrade": false
}
```

#### Processing Logic

When a Radarr webhook is received:

1. **Extract Movie Information**: Parse movie title, year, and metadata
2. **Locate Media File**: Use the movie file path to identify the media
3. **Download Subtitles**: Attempt to download subtitles for configured
   languages
4. **Store Results**: Log the operation in the history database

### Custom Webhooks

**Endpoint**: `POST /api/webhooks/custom`

Accepts custom webhook payloads for integration with other systems.

#### Custom Payload Format

```json
{
  "event": "media_added",
  "media": {
    "type": "movie" | "episode",
    "title": "Media Title",
    "year": 2024,
    "season": 1,      // For episodes only
    "episode": 1,     // For episodes only
    "file_path": "/path/to/media/file.mkv",
    "languages": ["en", "es", "fr"],
    "metadata": {
      "imdb_id": "tt1234567",
      "tmdb_id": 12345,
      "tvdb_id": 67890
    }
  }
}
```

#### Custom Integration Examples

**Plex Integration** (via custom script):

```bash
#!/bin/bash
# Plex post-processing script

MEDIA_PATH="$1"
MEDIA_TYPE="$2"  # movie or episode
TITLE="$3"

curl -X POST http://subtitle-manager:8080/api/webhooks/custom \
  -H "Content-Type: application/json" \
  -d '{
    "event": "media_added",
    "media": {
      "type": "'$MEDIA_TYPE'",
      "title": "'$TITLE'",
      "file_path": "'$MEDIA_PATH'",
      "languages": ["en"]
    }
  }'
```

**Jellyfin Integration** (via plugin):

```javascript
// Jellyfin webhook plugin
const webhookUrl = 'http://subtitle-manager:8080/api/webhooks/custom';

function onMediaAdded(mediaItem) {
  const payload = {
    event: 'media_added',
    media: {
      type: mediaItem.Type === 'Movie' ? 'movie' : 'episode',
      title: mediaItem.Name,
      year: mediaItem.ProductionYear,
      file_path: mediaItem.Path,
      languages: ['en', 'es'],
      metadata: {
        imdb_id: mediaItem.ProviderIds?.Imdb,
        tmdb_id: mediaItem.ProviderIds?.Tmdb,
      },
    },
  };

  fetch(webhookUrl, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
}
```

## Configuration

### Global Webhook Settings

Configure webhook behavior in the application settings:

```yaml
# config.yaml
webhooks:
  enabled: true
  default_languages: ['en']
  auto_download: true
  provider_priority: ['opensubtitles', 'subscene', 'addic7ed']
  quality_filter: ['bluray', 'web', 'hdtv']

sonarr:
  enabled: true
  languages: ['en', 'es']
  series_mapping:
    'Breaking Bad': ['en', 'es', 'fr']
    'Game of Thrones': ['en']

radarr:
  enabled: true
  languages: ['en']
  quality_profiles:
    'Ultra-HD': ['en', 'es', 'fr']
    'HD-1080p': ['en', 'es']
    'HD-720p': ['en']
```

### Environment Variables

```bash
# Webhook configuration
WEBHOOK_ENABLED=true
WEBHOOK_DEFAULT_LANGUAGES=en,es
WEBHOOK_AUTO_DOWNLOAD=true

# Provider settings
SUBTITLE_PROVIDERS=opensubtitles,subscene,addic7ed
PROVIDER_TIMEOUT=30s
MAX_DOWNLOAD_ATTEMPTS=3
```

## Security

### Authentication

Webhooks support optional authentication for security:

#### API Key Authentication

```bash
curl -X POST http://subtitle-manager:8080/api/webhooks/sonarr \
  -H "X-API-Key: your-api-key" \
  -H "Content-Type: application/json" \
  -d @webhook-payload.json
```

#### IP Whitelisting

Configure allowed IP addresses in settings:

```yaml
webhooks:
  security:
    ip_whitelist:
      - '192.168.1.100' # Sonarr server
      - '192.168.1.101' # Radarr server
      - '10.0.0.0/8' # Local network
    require_api_key: false
```

#### Webhook Signatures

For enhanced security, validate webhook signatures:

```bash
# Generate signature (sender)
PAYLOAD='{"event":"media_added"}'
SECRET="webhook-secret-key"
SIGNATURE=$(echo -n "$PAYLOAD" | openssl dgst -sha256 -hmac "$SECRET" -binary | base64)

# Send with signature
curl -X POST http://subtitle-manager:8080/api/webhooks/custom \
  -H "X-Webhook-Signature: sha256=$SIGNATURE" \
  -H "Content-Type: application/json" \
  -d "$PAYLOAD"
```

### Input Validation

All webhook payloads are validated for:

- Required fields presence
- Data type validation
- Path traversal prevention
- Maximum payload size limits
- Rate limiting per IP

## Error Handling

### HTTP Status Codes

- **200**: Webhook processed successfully
- **400**: Invalid payload format
- **401**: Authentication required
- **403**: IP not whitelisted or invalid API key
- **413**: Payload too large
- **429**: Rate limit exceeded
- **500**: Internal processing error

### Error Response Format

```json
{
  "error": "invalid_payload",
  "message": "Missing required field: media.file_path",
  "details": {
    "field": "media.file_path",
    "expected": "string",
    "received": "null"
  }
}
```

### Retry Logic

Failed webhooks should implement exponential backoff:

```python
import time
import requests

def send_webhook_with_retry(url, payload, max_retries=3):
    for attempt in range(max_retries):
        try:
            response = requests.post(url, json=payload, timeout=30)
            if response.status_code == 200:
                return response
            elif response.status_code in [429, 500, 502, 503]:
                # Retry on rate limit or server errors
                delay = 2 ** attempt  # Exponential backoff
                time.sleep(delay)
                continue
            else:
                # Don't retry on client errors
                response.raise_for_status()
        except requests.RequestException as e:
            if attempt == max_retries - 1:
                raise
            time.sleep(2 ** attempt)

    raise Exception("Max retries exceeded")
```

## Testing Webhooks

### Manual Testing

Test webhook endpoints using curl:

```bash
# Test Sonarr webhook
curl -X POST http://localhost:8080/api/webhooks/sonarr \
  -H "Content-Type: application/json" \
  -d '{
    "eventType": "Download",
    "series": {
      "title": "Test Series",
      "path": "/tv/test"
    },
    "episodes": [{
      "seasonNumber": 1,
      "episodeNumber": 1,
      "episodeFile": {
        "path": "/tv/test/S01E01.mkv"
      }
    }]
  }'

# Test custom webhook
curl -X POST http://localhost:8080/api/webhooks/custom \
  -H "Content-Type: application/json" \
  -d '{
    "event": "media_added",
    "media": {
      "type": "movie",
      "title": "Test Movie",
      "file_path": "/movies/test.mkv",
      "languages": ["en"]
    }
  }'
```

### Integration Testing

Create test scripts for your media server integration:

```python
#!/usr/bin/env python3
# webhook_test.py

import requests
import json

def test_sonarr_webhook():
    payload = {
        "eventType": "Download",
        "series": {
            "title": "Breaking Bad",
            "path": "/tv/Breaking Bad"
        },
        "episodes": [{
            "seasonNumber": 1,
            "episodeNumber": 1,
            "episodeFile": {
                "path": "/tv/Breaking Bad/Season 01/S01E01.mkv"
            }
        }]
    }

    response = requests.post(
        "http://localhost:8080/api/webhooks/sonarr",
        json=payload,
        headers={"X-API-Key": "your-api-key"}
    )

    print(f"Status: {response.status_code}")
    print(f"Response: {response.text}")

    assert response.status_code == 200

if __name__ == "__main__":
    test_sonarr_webhook()
    print("Webhook test passed!")
```

## Monitoring and Logging

### Webhook Logs

Monitor webhook activity through the logs API:

```bash
# Get webhook-related logs
curl -X GET "http://localhost:8080/api/logs?component=webhook&level=info&limit=100" \
  -H "X-API-Key: your-api-key"
```

### Metrics and Statistics

Track webhook performance:

```bash
# Get webhook statistics
curl -X GET "http://localhost:8080/api/system/webhooks/stats" \
  -H "X-API-Key: your-api-key"
```

Response:

```json
{
  "total_webhooks": 1250,
  "successful_webhooks": 1180,
  "failed_webhooks": 70,
  "average_processing_time": 245,
  "webhooks_by_type": {
    "sonarr": 800,
    "radarr": 350,
    "custom": 100
  },
  "recent_errors": [
    {
      "timestamp": "2024-01-01T12:00:00Z",
      "type": "sonarr",
      "error": "file_not_found",
      "details": "Media file not accessible"
    }
  ]
}
```

### Real-time Monitoring

Use WebSocket connections for real-time webhook monitoring:

```javascript
const ws = new WebSocket('ws://localhost:8080/ws/webhooks');

ws.onmessage = event => {
  const data = JSON.parse(event.data);
  console.log('Webhook received:', data);

  if (data.type === 'webhook_processed') {
    console.log(
      `Processed ${data.webhook_type} webhook for ${data.media_title}`
    );
    console.log(
      `Status: ${data.status}, Subtitles: ${data.subtitles_downloaded}`
    );
  }
};
```

## Advanced Usage

### Conditional Processing

Configure conditional subtitle downloads based on media properties:

```yaml
webhooks:
  rules:
    - name: 'High quality movies only'
      condition:
        media_type: 'movie'
        quality: ['Bluray-2160p', 'Bluray-1080p']
      action:
        languages: ['en', 'es', 'fr']
        providers: ['opensubtitles', 'subscene']

    - name: 'Popular TV series'
      condition:
        media_type: 'episode'
        series_title: ['Breaking Bad', 'Game of Thrones', 'The Office']
      action:
        languages: ['en', 'es', 'fr', 'de']
        providers: ['all']

    - name: 'Default rule'
      condition: {}
      action:
        languages: ['en']
        providers: ['opensubtitles']
```

### Batch Processing

Handle multiple media files in a single webhook:

```json
{
  "event": "batch_added",
  "media_batch": [
    {
      "type": "movie",
      "title": "Movie 1",
      "file_path": "/movies/movie1.mkv"
    },
    {
      "type": "movie",
      "title": "Movie 2",
      "file_path": "/movies/movie2.mkv"
    }
  ],
  "languages": ["en", "es"]
}
```

### Custom Processing Pipelines

Implement custom processing workflows:

```python
# custom_webhook_processor.py

import requests
from typing import Dict, List

class CustomWebhookProcessor:
    def __init__(self, subtitle_manager_url: str, api_key: str):
        self.base_url = subtitle_manager_url
        self.api_key = api_key
        self.headers = {"X-API-Key": api_key}

    def process_media_batch(self, media_files: List[Dict]) -> Dict:
        results = []

        for media in media_files:
            # Step 1: Download subtitles
            download_result = self.download_subtitles(
                media["file_path"],
                media.get("languages", ["en"])
            )

            # Step 2: If download fails, try extraction
            if not download_result["success"]:
                extract_result = self.extract_subtitles(media["file_path"])
                if extract_result["success"]:
                    # Step 3: Translate extracted subtitles
                    translate_result = self.translate_subtitles(
                        extract_result["subtitle_path"],
                        media.get("target_language", "en")
                    )
                    results.append(translate_result)
                else:
                    results.append(extract_result)
            else:
                results.append(download_result)

        return {"processed": len(results), "results": results}

    def download_subtitles(self, file_path: str, languages: List[str]) -> Dict:
        # Implementation using Subtitle Manager API
        pass

    def extract_subtitles(self, file_path: str) -> Dict:
        # Implementation using Subtitle Manager API
        pass

    def translate_subtitles(self, subtitle_path: str, target_language: str) -> Dict:
        # Implementation using Subtitle Manager API
        pass
```

## Troubleshooting

### Common Issues

1. **Webhook not triggered**
   - Check media server webhook configuration
   - Verify URL and network connectivity
   - Check firewall settings

2. **Authentication failures**
   - Verify API key is correct
   - Check IP whitelist configuration
   - Ensure proper headers are sent

3. **File not found errors**
   - Verify file paths are accessible to Subtitle Manager
   - Check file permissions
   - Ensure network mounts are properly configured

4. **Subtitle download failures**
   - Check provider API status
   - Verify media metadata (IMDB ID, etc.)
   - Check language availability

### Debug Mode

Enable debug logging for detailed webhook processing information:

```yaml
logging:
  level: debug
  components:
    webhook: debug
    download: debug
    provider: debug
```

### Health Checks

Monitor webhook endpoint health:

```bash
# Check webhook endpoint availability
curl -X GET http://localhost:8080/api/webhooks/health \
  -H "X-API-Key: your-api-key"
```

Response:

```json
{
  "status": "healthy",
  "endpoints": {
    "sonarr": "active",
    "radarr": "active",
    "custom": "active"
  },
  "recent_activity": {
    "last_webhook": "2024-01-01T12:05:00Z",
    "webhooks_last_hour": 15,
    "average_response_time": 234
  }
}
```

## Best Practices

1. **Use HTTPS**: Always use HTTPS for webhook endpoints in production
2. **Implement Retry Logic**: Handle temporary failures with exponential backoff
3. **Validate Inputs**: Always validate webhook payloads before processing
4. **Monitor Performance**: Track webhook processing times and success rates
5. **Use Queues**: For high-volume environments, consider using message queues
6. **Test Thoroughly**: Test webhook integrations with realistic payloads
7. **Document Custom Integrations**: Maintain documentation for custom webhook
   implementations
8. **Regular Health Checks**: Implement health checks for webhook endpoints
9. **Rate Limiting**: Implement rate limiting to prevent abuse
10. **Security First**: Use authentication, IP whitelisting, and signature
    validation

This comprehensive webhook integration enables seamless automation of subtitle
management within your media infrastructure, ensuring subtitles are
automatically downloaded and processed as new content becomes available.
