<!-- file: docs/api/rate-limiting.md -->
<!-- version: 1.0.0 -->
<!-- guid: 550e8400-e29b-41d4-a716-446655440030 -->

# Rate Limiting Guide

Subtitle Manager implements comprehensive rate limiting to ensure fair usage and
system stability. This guide covers rate limit policies, headers, and best
practices for handling rate limits in your applications.

## Overview

Rate limiting is enforced at multiple levels:

- **Global limits**: Overall API usage limits
- **User-based limits**: Limits based on user permission level
- **Endpoint-specific limits**: Special limits for resource-intensive operations
- **IP-based limits**: Protection against abuse from specific IP addresses

## Rate Limit Policies

### Standard Rate Limits (Per User)

| Permission Level | Requests/Hour | Requests/Minute | Burst Limit |
| ---------------- | ------------- | --------------- | ----------- |
| **Read**         | 1,000         | 60              | 10          |
| **Basic**        | 500           | 30              | 5           |
| **Admin**        | 200           | 20              | 3           |

**Why different limits?**

- Read operations are less resource-intensive
- Basic operations include file processing which uses more resources
- Admin operations often affect system-wide settings

### Endpoint-Specific Limits

#### File Processing Operations

| Endpoint         | Limit   | Window | Reason                                 |
| ---------------- | ------- | ------ | -------------------------------------- |
| `/api/convert`   | 50/hour | 1 hour | CPU-intensive subtitle conversion      |
| `/api/translate` | 30/hour | 1 hour | External API costs and processing time |
| `/api/extract`   | 20/hour | 1 hour | Memory-intensive video file processing |

#### Provider Operations

| Endpoint        | Limit    | Window | Reason                              |
| --------------- | -------- | ------ | ----------------------------------- |
| `/api/download` | 100/hour | 1 hour | External provider API limits        |
| `/api/scan`     | 10/hour  | 1 hour | Resource-intensive library scanning |

#### System Operations

| Endpoint                     | Limit  | Window | Reason                       |
| ---------------------------- | ------ | ------ | ---------------------------- |
| `/api/oauth/github/generate` | 5/hour | 1 hour | Security-sensitive operation |
| `/api/database/backup`       | 3/hour | 1 hour | Heavy I/O operation          |

### IP-Based Rate Limits

| Category                     | Limit     | Window | Action               |
| ---------------------------- | --------- | ------ | -------------------- |
| **Unauthenticated requests** | 100/hour  | 1 hour | Block IP for 1 hour  |
| **Failed login attempts**    | 10/hour   | 1 hour | Block IP for 2 hours |
| **Webhook requests**         | 1000/hour | 1 hour | Return 429 status    |

## Rate Limit Headers

All API responses include rate limiting information in headers:

```http
X-RateLimit-Limit: 500
X-RateLimit-Remaining: 487
X-RateLimit-Reset: 1640995200
X-RateLimit-Window: 3600
X-RateLimit-Type: user
```

### Header Descriptions

- **`X-RateLimit-Limit`**: Maximum requests allowed in the current window
- **`X-RateLimit-Remaining`**: Number of requests remaining in current window
- **`X-RateLimit-Reset`**: Unix timestamp when the current window resets
- **`X-RateLimit-Window`**: Window duration in seconds
- **`X-RateLimit-Type`**: Type of rate limit applied (`user`, `ip`, `endpoint`)

### Additional Headers (When Rate Limited)

When a rate limit is exceeded (HTTP 429), additional headers are included:

```http
Retry-After: 3600
X-RateLimit-Exceeded: user
X-RateLimit-Reset-Time: 2024-01-01T15:00:00Z
```

- **`Retry-After`**: Seconds until the rate limit resets
- **`X-RateLimit-Exceeded`**: Which rate limit was exceeded
- **`X-RateLimit-Reset-Time`**: ISO 8601 timestamp of reset time

## Handling Rate Limits

### Detecting Rate Limits

#### HTTP Status Code 429

```http
HTTP/1.1 429 Too Many Requests
Content-Type: application/json
Retry-After: 3600
X-RateLimit-Limit: 500
X-RateLimit-Remaining: 0
X-RateLimit-Reset: 1640995200

{
  "error": "rate_limit_exceeded",
  "message": "API rate limit exceeded. Try again in 3600 seconds.",
  "details": {
    "limit": 500,
    "window": "1 hour",
    "reset_time": "2024-01-01T15:00:00Z"
  }
}
```

### Best Practices for Rate Limit Handling

#### 1. Respect Rate Limit Headers

Always check rate limit headers before making requests:

```python
# Python example
import time
import requests

def make_request_with_rate_limit_check(url, headers):
    # Check rate limit before making request
    response = requests.get(url, headers=headers)

    remaining = int(response.headers.get('X-RateLimit-Remaining', 0))
    if remaining < 10:  # Buffer for safety
        reset_time = int(response.headers.get('X-RateLimit-Reset', 0))
        sleep_time = max(0, reset_time - time.time())
        print(f"Rate limit low, sleeping for {sleep_time} seconds")
        time.sleep(sleep_time)

    return response
```

#### 2. Implement Exponential Backoff

When rate limited, use exponential backoff with jitter:

```javascript
// JavaScript example
async function requestWithBackoff(url, options, maxRetries = 3) {
  for (let attempt = 0; attempt < maxRetries; attempt++) {
    try {
      const response = await fetch(url, options);

      if (response.status === 429) {
        const retryAfter = parseInt(
          response.headers.get('Retry-After') || '60'
        );
        const jitter = Math.random() * 1000; // Add jitter
        const delay = retryAfter * 1000 + jitter;

        console.log(
          `Rate limited, waiting ${delay}ms before retry ${attempt + 1}`
        );
        await new Promise(resolve => setTimeout(resolve, delay));
        continue;
      }

      return response;
    } catch (error) {
      if (attempt === maxRetries - 1) throw error;

      const delay = Math.pow(2, attempt) * 1000 + Math.random() * 1000;
      await new Promise(resolve => setTimeout(resolve, delay));
    }
  }
}
```

#### 3. Use Request Queuing

Implement request queuing to stay within rate limits:

```go
// Go example
package main

import (
    "context"
    "sync"
    "time"
    "golang.org/x/time/rate"
)

type RateLimitedClient struct {
    limiter *rate.Limiter
    client  *http.Client
}

func NewRateLimitedClient(requestsPerSecond rate.Limit) *RateLimitedClient {
    return &RateLimitedClient{
        limiter: rate.NewLimiter(requestsPerSecond, 1), // Burst of 1
        client:  &http.Client{},
    }
}

func (c *RateLimitedClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
    // Wait for rate limiter
    if err := c.limiter.Wait(ctx); err != nil {
        return nil, err
    }

    return c.client.Do(req)
}
```

#### 4. Batch Operations

Group multiple operations into single requests when possible:

```python
# Instead of multiple individual downloads
for movie in movies:
    client.download_subtitles(movie.path, "en")

# Use batch operations
batch_requests = [
    {"path": movie.path, "language": "en"}
    for movie in movies
]
results = client.batch_download_subtitles(batch_requests)
```

### SDK Rate Limit Handling

#### Python SDK

```python
from subtitle_manager_sdk import SubtitleManagerClient
from subtitle_manager_sdk.exceptions import RateLimitError
import time

client = SubtitleManagerClient("http://localhost:8080", api_key="your-key")

def download_with_rate_limit_handling(path, language):
    max_retries = 3
    for attempt in range(max_retries):
        try:
            return client.download_subtitles(path, language)
        except RateLimitError as e:
            if attempt < max_retries - 1:
                wait_time = e.retry_after or (2 ** attempt * 60)
                print(f"Rate limited, waiting {wait_time} seconds")
                time.sleep(wait_time)
                continue
            raise
```

#### JavaScript SDK

```javascript
import { SubtitleManagerClient, RateLimitError } from 'subtitle-manager-sdk';

const client = new SubtitleManagerClient({
  baseURL: 'http://localhost:8080',
  apiKey: 'your-key',
  // Built-in rate limit handling
  maxRetries: 3,
  retryDelay: 2000,
});

// SDK automatically handles rate limits with exponential backoff
try {
  const result = await client.downloadSubtitles('/movies/example.mkv', 'en');
} catch (error) {
  if (error instanceof RateLimitError) {
    console.log(`Rate limited, retry after ${error.retryAfter} seconds`);
  }
}
```

#### Go SDK

```go
import (
    "context"
    "time"
    "golang.org/x/time/rate"
    "github.com/jdfalk/subtitle-manager/sdks/go/subtitleclient"
)

// Create client with built-in rate limiting
client := subtitleclient.NewClient(subtitleclient.Config{
    BaseURL:   "http://localhost:8080",
    APIKey:    "your-key",
    RateLimit: rate.Limit(5), // 5 requests per second
})

// SDK automatically respects rate limits
ctx := context.Background()
result, err := client.DownloadSubtitles(ctx, subtitleclient.DownloadRequest{
    Path:     "/movies/example.mkv",
    Language: "en",
})
```

## Monitoring Rate Limits

### Check Current Usage

Get your current rate limit status:

```bash
curl -H "X-API-Key: your-key" \
     -I http://localhost:8080/api/system
```

Response headers will show current usage:

```
X-RateLimit-Limit: 500
X-RateLimit-Remaining: 243
X-RateLimit-Reset: 1640995200
```

### Rate Limit Metrics API

Get detailed rate limit information:

```bash
curl -H "X-API-Key: your-key" \
     http://localhost:8080/api/system/rate-limits
```

```json
{
  "user": {
    "limit": 500,
    "remaining": 243,
    "reset_time": "2024-01-01T15:00:00Z",
    "window": "1 hour"
  },
  "endpoints": {
    "/api/convert": {
      "limit": 50,
      "remaining": 45,
      "reset_time": "2024-01-01T15:00:00Z"
    },
    "/api/translate": {
      "limit": 30,
      "remaining": 28,
      "reset_time": "2024-01-01T15:00:00Z"
    }
  },
  "history": {
    "last_hour": 257,
    "last_24_hours": 1432,
    "peak_hour": 489
  }
}
```

### Webhooks for Rate Limit Events

Configure webhooks to be notified of rate limit events:

```json
{
  "event": "rate_limit_exceeded",
  "user": {
    "id": 123,
    "username": "api_user",
    "permission": "basic"
  },
  "limit": {
    "type": "user",
    "limit": 500,
    "window": 3600,
    "exceeded_at": "2024-01-01T14:30:00Z"
  },
  "request": {
    "endpoint": "/api/download",
    "ip": "192.168.1.100",
    "user_agent": "subtitle-manager-python-sdk/1.0.0"
  }
}
```

## Rate Limit Bypass Options

### Increased Limits for Production

Contact support for increased limits in production environments:

- **Volume discounts**: Reduced per-request costs for high-volume usage
- **Dedicated limits**: Custom limits for enterprise deployments
- **Priority queues**: Faster processing for critical operations

### Self-Hosted Deployments

Self-hosted instances can customize rate limits:

```yaml
# config.yaml
rate_limits:
  enabled: true

  # User-based limits
  user_limits:
    read: 2000 # requests per hour
    basic: 1000
    admin: 500

  # Endpoint-specific limits
  endpoint_limits:
    '/api/convert': 100
    '/api/translate': 60
    '/api/extract': 40

  # IP-based limits
  ip_limits:
    unauthenticated: 200
    failed_auth: 20

  # Window settings
  window_size: 3600 # 1 hour in seconds
  burst_size: 10 # Burst allowance
```

### Rate Limit Exemptions

Certain operations may be exempt from rate limits:

- **Health checks**: System monitoring endpoints
- **Webhooks**: Incoming webhook requests (with separate limits)
- **Emergency operations**: Database backups during maintenance
- **System tasks**: Internal background processes

## Troubleshooting Rate Limits

### Common Issues

#### Sudden Rate Limit Exceeded

**Symptoms**: Unexpected 429 errors despite normal usage **Causes**:

- Clock skew between client and server
- Multiple API clients sharing the same API key
- Background processes consuming rate limit quota

**Solutions**:

- Check system time synchronization
- Use separate API keys for different applications
- Monitor and audit API key usage

#### Rate Limits Not Resetting

**Symptoms**: Rate limits remain at 0 even after window expires **Causes**:

- Server-side cache issues
- Timezone configuration problems
- Database connectivity issues

**Solutions**:

- Check server logs for errors
- Verify timezone settings
- Restart rate limiting service if self-hosted

#### Inconsistent Rate Limits

**Symptoms**: Different rate limits reported by different endpoints **Causes**:

- Load balancer with sticky sessions disabled
- Multiple server instances with different configurations
- Cache inconsistency

**Solutions**:

- Enable session affinity in load balancer
- Ensure consistent configuration across instances
- Use centralized rate limiting store (Redis)

### Debug Rate Limits

Enable debug logging for rate limiting:

```bash
# Check rate limit logs
curl -H "X-API-Key: your-admin-key" \
     "http://localhost:8080/api/logs?component=ratelimit&level=debug"
```

Example debug output:

```json
[
  {
    "timestamp": "2024-01-01T14:30:00Z",
    "level": "debug",
    "component": "ratelimit",
    "message": "Rate limit check",
    "fields": {
      "user_id": 123,
      "endpoint": "/api/download",
      "current_usage": 487,
      "limit": 500,
      "window_start": "2024-01-01T14:00:00Z",
      "ip": "192.168.1.100"
    }
  }
]
```

### Rate Limit Testing

Test rate limit behavior:

```python
# Python script to test rate limits
import time
import requests

def test_rate_limits():
    headers = {"X-API-Key": "test-api-key"}
    base_url = "http://localhost:8080"

    for i in range(600):  # Test beyond limit
        response = requests.get(f"{base_url}/api/system", headers=headers)

        print(f"Request {i+1}: {response.status_code}")
        print(f"Remaining: {response.headers.get('X-RateLimit-Remaining')}")

        if response.status_code == 429:
            retry_after = int(response.headers.get('Retry-After', 60))
            print(f"Rate limited! Retry after {retry_after} seconds")
            break

        time.sleep(0.1)  # Small delay between requests

if __name__ == "__main__":
    test_rate_limits()
```

## Best Practices Summary

1. **Always check rate limit headers** before making requests
2. **Implement exponential backoff** with jitter for retries
3. **Use request queuing** to stay within limits
4. **Batch operations** when possible to reduce request count
5. **Monitor usage patterns** and adjust accordingly
6. **Use separate API keys** for different applications
7. **Cache responses** when appropriate to reduce API calls
8. **Implement circuit breakers** to handle sustained rate limiting
9. **Log rate limit events** for monitoring and debugging
10. **Plan for growth** by designing scalable request patterns

## Related Documentation

- [Authentication Guide](authentication.md) - API key management and
  authentication methods
- [SDK Documentation](../../sdks/) - Rate limit handling in each SDK
- [Monitoring Guide](monitoring.md) - System monitoring and observability
- [Troubleshooting Guide](troubleshooting.md) - Common issues and solutions
