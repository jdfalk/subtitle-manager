# file: docs/api/README.md
# version: 1.0.0
# guid: 550e8400-e29b-41d4-a716-446655440029

# Subtitle Manager API Documentation

Complete documentation for the Subtitle Manager REST API, including interactive documentation, client SDKs, and integration guides.

## 📚 Documentation Overview

### Core API Documentation
- **[📖 Interactive API Documentation](index.html)** - Complete OpenAPI specification with Swagger UI
- **[🔐 Authentication Guide](authentication.md)** - Comprehensive guide to all authentication methods
- **🔗 [Webhook Integration Guide](webhooks.md)** - Integration with Sonarr, Radarr, and custom webhooks

### Client SDKs

| Language | Status | Documentation | Features |
|----------|--------|---------------|----------|
| **Python** | ✅ Complete | [Python SDK](../../sdks/python/README.md) | Type safety, async support, automatic retry |
| **JavaScript/TypeScript** | ✅ Complete | [JS/TS SDK](../../sdks/javascript/README.md) | Full TypeScript, React hooks, Node.js support |
| **Go** | ✅ Complete | [Go SDK](../../sdks/go/README.md) | Context support, rate limiting, concurrency safe |
| **C#** | 🚧 Planned | Coming soon | Async/await, .NET Core support |
| **Java** | 🚧 Planned | Coming soon | Modern patterns, Spring Boot integration |

### Integration Examples
- **[🐍 Python Integration Examples](examples/python-integration.py)** - Production-ready Python integration patterns
- **[🟨 JavaScript Integration Examples](examples/javascript-integration.js)** - React, Node.js, and Express.js examples

## 🚀 Quick Start

### 1. Interactive API Explorer

Visit the [Interactive API Documentation](index.html) to:
- Explore all 58+ API endpoints
- Test API calls directly in your browser  
- View detailed request/response schemas
- Copy code examples for your language

### 2. Choose Your SDK

#### Python
```bash
pip install subtitle-manager-sdk
```

```python
from subtitle_manager_sdk import SubtitleManagerClient

client = SubtitleManagerClient("http://localhost:8080", api_key="your-key")
system_info = client.get_system_info()
```

#### JavaScript/TypeScript
```bash
npm install subtitle-manager-sdk
```

```javascript
import { SubtitleManagerClient } from 'subtitle-manager-sdk';

const client = new SubtitleManagerClient({
  baseURL: 'http://localhost:8080',
  apiKey: 'your-key'
});

const systemInfo = await client.getSystemInfo();
```

#### Go
```bash
go get github.com/jdfalk/subtitle-manager/sdks/go
```

```go
import "github.com/jdfalk/subtitle-manager/sdks/go/subtitleclient"

client := subtitleclient.NewDefaultClient("http://localhost:8080", "your-key")
systemInfo, err := client.GetSystemInfo(ctx)
```

### 3. Authentication Setup

The API supports multiple authentication methods:

- **API Keys**: For programmatic access
- **Session Cookies**: For web UI integration  
- **OAuth2**: For secure user authentication

See the [Authentication Guide](authentication.md) for detailed setup instructions.

## 📋 API Overview

### Core Endpoints

| Category | Endpoints | Description |
|----------|-----------|-------------|
| **Authentication** | `/api/login`, `/api/logout` | User authentication and session management |
| **Subtitle Operations** | `/api/convert`, `/api/translate`, `/api/extract` | Core subtitle processing |
| **Downloads** | `/api/download` | Subtitle downloads from providers |
| **Library** | `/api/scan`, `/api/scan/status` | Media library management |
| **History** | `/api/history` | Operation history and audit logs |
| **System** | `/api/system`, `/api/logs` | System monitoring and information |
| **Webhooks** | `/api/webhooks/*` | Integration with media servers |

### Supported Operations

- **Convert**: Transform subtitle files between formats (VTT → SRT, ASS → SRT, etc.)
- **Translate**: Translate subtitles using Google Translate or OpenAI
- **Extract**: Extract embedded subtitles from video files
- **Download**: Fetch subtitles from providers (OpenSubtitles, Subscene, etc.)
- **Library Scan**: Index media files and download missing subtitles
- **Webhook Integration**: Automatic processing via Sonarr/Radarr

## 🔐 Security & Authentication

### Authentication Methods

1. **API Key Authentication** (Recommended for automation)
   ```bash
   curl -H "X-API-Key: your-api-key" http://localhost:8080/api/system
   ```

2. **Session Cookie Authentication** (For web applications)
   ```bash
   curl -X POST -d '{"username":"admin","password":"pass"}' \
        -H "Content-Type: application/json" \
        http://localhost:8080/api/login
   ```

3. **OAuth2 Authentication** (For user-facing applications)
   - GitHub OAuth2 integration
   - Automatic user creation and profile sync

### Permission Levels

- **`read`**: View-only access (history, logs, system info)
- **`basic`**: Standard operations (convert, translate, download, scan)
- **`admin`**: Full administrative access (user management, OAuth config)

### Rate Limiting

API requests are rate limited by user role:
- Read operations: 1000 requests/hour
- Basic operations: 500 requests/hour
- Admin operations: 200 requests/hour

Rate limit headers are included in all responses:
```
X-RateLimit-Limit: 500
X-RateLimit-Remaining: 487
X-RateLimit-Reset: 1640995200
```

## 🔌 Integration Patterns

### Media Server Integration

#### Sonarr Integration
Configure Sonarr to automatically download subtitles when episodes are added:

1. Add webhook in Sonarr: **Settings → Connect → Webhook**
2. URL: `http://subtitle-manager:8080/api/webhooks/sonarr`
3. Triggers: "On Download" and "On Upgrade"

#### Radarr Integration  
Similar setup for automatic movie subtitle downloads:

1. Add webhook in Radarr: **Settings → Connect → Webhook**
2. URL: `http://subtitle-manager:8080/api/webhooks/radarr`
3. Triggers: "On Download" and "On Upgrade"

See the [Webhook Integration Guide](webhooks.md) for complete setup instructions.

### Custom Integrations

#### Plex Integration
```bash
#!/bin/bash
# Plex post-processing script
curl -X POST http://subtitle-manager:8080/api/webhooks/custom \
  -H "Content-Type: application/json" \
  -d '{
    "event": "media_added",
    "media": {
      "type": "movie",
      "title": "'$TITLE'",
      "file_path": "'$MEDIA_PATH'",
      "languages": ["en", "es"]
    }
  }'
```

#### Jellyfin Integration
```javascript
// Jellyfin webhook plugin
function onMediaAdded(mediaItem) {
  fetch('http://subtitle-manager:8080/api/webhooks/custom', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      event: 'media_added',
      media: {
        type: mediaItem.Type === 'Movie' ? 'movie' : 'episode',
        title: mediaItem.Name,
        file_path: mediaItem.Path,
        languages: ['en']
      }
    })
  });
}
```

## 📊 Monitoring & Observability

### Health Monitoring

Check API health programmatically:

```python
# Python
healthy = client.health_check()

# JavaScript
const healthy = await client.healthCheck();

# Go
healthy := client.HealthCheck(ctx)

# cURL
curl http://localhost:8080/api/system
```

### Metrics & Analytics

Get operational insights:

```python
# Get system information
system_info = client.get_system_info()
print(f"Disk usage: {system_info.disk_usage_percent:.1f}%")

# Get recent operation history
history = client.get_history(limit=100)
success_rate = sum(1 for item in history.items if item.is_success) / len(history.items)
print(f"Success rate: {success_rate:.1%}")

# Get error logs
error_logs = client.get_logs(level="error", limit=50)
for log in error_logs:
    print(f"{log.timestamp}: {log.component} - {log.message}")
```

### Real-time Updates

Monitor operations in real-time using WebSocket connections:

```javascript
const ws = new WebSocket('ws://localhost:8080/ws/tasks');
ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log('Real-time update:', data);
};
```

## 🛠️ Development & Testing

### API Testing

Use the interactive documentation to test API endpoints:

1. Visit [Interactive API Documentation](index.html)
2. Click "Authorize" and enter your API key
3. Expand any endpoint and click "Try it out"
4. Fill in parameters and click "Execute"

### SDK Development

#### Running Tests

```bash
# Python SDK
cd sdks/python
python -m pytest tests/ -v

# JavaScript SDK  
cd sdks/javascript
npm test

# Go SDK
cd sdks/go
go test ./... -v
```

#### Building SDKs

```bash
# JavaScript SDK
npm run build

# Go SDK (no build needed - source distribution)
go mod tidy
```

### Integration Testing

Test webhook integrations:

```bash
# Test Sonarr webhook
curl -X POST http://localhost:8080/api/webhooks/sonarr \
  -H "Content-Type: application/json" \
  -d '{
    "eventType": "Download",
    "series": {"title": "Test Series"},
    "episodes": [{"episodeFile": {"path": "/tv/test.mkv"}}]
  }'
```

## 🐛 Troubleshooting

### Common Issues

#### Authentication Errors
- **401 Unauthorized**: Check API key validity
- **403 Forbidden**: Verify user has required permission level
- **Session expired**: Re-authenticate if using session cookies

#### Rate Limiting
- **429 Too Many Requests**: Implement exponential backoff
- Check rate limit headers in responses
- Consider using different permission levels

#### File Operations
- **400 Bad Request**: Verify file format and size limits
- **404 Not Found**: Check file paths and accessibility
- **500 Internal Server Error**: Check server logs for details

### Debug Mode

Enable debug logging in SDKs:

```python
# Python - enable logging
import logging
logging.getLogger('subtitle_manager_sdk').setLevel(logging.DEBUG)

# JavaScript - enable debug mode
const client = new SubtitleManagerClient({
  baseURL: 'http://localhost:8080',
  debug: true
});

# Go - use context with logging
ctx := context.WithValue(context.Background(), "debug", true)
```

### Support Resources

- **API Documentation**: This documentation
- **GitHub Issues**: [Report bugs and request features](https://github.com/jdfalk/subtitle-manager/issues)
- **Source Code**: [GitHub Repository](https://github.com/jdfalk/subtitle-manager)
- **Examples**: See integration examples in each SDK documentation

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](../../LICENSE) file for details.

---

## 📚 Related Documentation

- [Project README](../../README.md) - Main project documentation
- [Technical Design](../TECHNICAL_DESIGN.md) - Architecture and design decisions  
- [Deployment Guide](../DEPLOYMENT.md) - Installation and deployment instructions
- [Contributing Guide](../../CONTRIBUTING.md) - How to contribute to the project