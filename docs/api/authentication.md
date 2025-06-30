# file: docs/api/authentication.md
# version: 1.0.0
# guid: 550e8400-e29b-41d4-a716-446655440003

# Authentication Guide

Subtitle Manager API supports multiple authentication methods to accommodate different use cases. This guide covers all available authentication options and their appropriate usage scenarios.

## Overview

The API implements a three-tier authorization system:

- **`read`**: View-only access to resources (history, system info, logs)
- **`basic`**: Standard user operations (download, convert, translate, scan)
- **`admin`**: Administrative operations (user management, OAuth config, system settings)

## Authentication Methods

### 1. Session Cookie Authentication

Session authentication is primarily used by the web UI and is the most secure method for browser-based interactions.

#### How it works:
1. User authenticates via `/api/login` endpoint
2. Server creates a session and sets an HTTP-only cookie
3. Subsequent requests include the session cookie automatically
4. Sessions expire after 24 hours of inactivity

#### Login Example:

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "your_password"
  }' \
  -c cookies.txt
```

#### Using the session:

```bash
curl -X GET http://localhost:8080/api/history \
  -b cookies.txt
```

#### Logout:

```bash
curl -X POST http://localhost:8080/api/logout \
  -b cookies.txt
```

### 2. API Key Authentication

API keys are ideal for programmatic access and automation. They don't expire and provide persistent access.

#### Generating an API Key:

API keys are generated through the web UI or can be created by administrators. Each user can have multiple API keys for different applications.

#### Using API Keys:

```bash
curl -X GET http://localhost:8080/api/system \
  -H "X-API-Key: your_api_key_here"
```

#### Example with curl:

```bash
# Download subtitles using API key
curl -X POST http://localhost:8080/api/download \
  -H "X-API-Key: your_api_key_here" \
  -H "Content-Type: application/json" \
  -d '{
    "path": "/movies/example.mkv",
    "language": "en"
  }'
```

#### API Key Best Practices:

- Store API keys securely (environment variables, secret managers)
- Use different API keys for different applications
- Rotate API keys regularly
- Never commit API keys to version control
- Use the minimum required permission level

### 3. OAuth2 Authentication (GitHub)

OAuth2 provides secure authentication without sharing passwords. Currently supports GitHub as the OAuth2 provider.

#### Setting up GitHub OAuth2:

1. **Admin Setup** (requires admin privileges):
   ```bash
   # Generate OAuth2 credentials
   curl -X POST http://localhost:8080/api/oauth/github/generate \
     -H "X-API-Key: admin_api_key"
   ```

2. **Configure GitHub App**:
   - Create a GitHub OAuth App in your GitHub settings
   - Set the authorization callback URL to: `https://your-domain.com/api/oauth/github/callback`
   - Use the generated client_id and client_secret

#### OAuth2 Flow:

1. **Initiate OAuth2 Flow**:
   ```
   GET https://your-domain.com/api/oauth/github/login
   ```

2. **User is redirected to GitHub for authorization**

3. **GitHub redirects back with authorization code**:
   ```
   GET https://your-domain.com/api/oauth/github/callback?code=auth_code
   ```

4. **User is authenticated and session is created**

#### OAuth2 Management (Admin Only):

```bash
# Regenerate client secret
curl -X POST http://localhost:8080/api/oauth/github/regenerate \
  -H "X-API-Key: admin_api_key"

# Reset OAuth2 configuration
curl -X POST http://localhost:8080/api/oauth/github/reset \
  -H "X-API-Key: admin_api_key"
```

## Permission Levels

### Read Permission
Allows access to:
- View operation history (`/api/history`)
- View system information (`/api/system`)
- View application logs (`/api/logs`)

### Basic Permission
Includes read permissions plus:
- Download subtitles (`/api/download`)
- Convert subtitle files (`/api/convert`)
- Translate subtitles (`/api/translate`)
- Extract embedded subtitles (`/api/extract`)
- Start library scans (`/api/scan`)
- View scan status (`/api/scan/status`)

### Admin Permission
Includes basic permissions plus:
- Manage OAuth2 configuration (`/api/oauth/github/*`)
- Manage user accounts (`/api/users/*`)
- System configuration (`/api/config`)
- Database operations (`/api/database/*`)

## Error Handling

Authentication errors are returned with appropriate HTTP status codes:

### 401 Unauthorized
No valid authentication provided:
```json
{
  "error": "unauthorized",
  "message": "Authentication required"
}
```

### 403 Forbidden
Valid authentication but insufficient permissions:
```json
{
  "error": "forbidden", 
  "message": "Admin access required"
}
```

## Security Best Practices

### For Web Applications:
- Use session authentication with HTTPS
- Implement proper CSRF protection
- Set secure cookie attributes (HttpOnly, Secure, SameSite)

### For API Clients:
- Use API keys stored securely
- Implement exponential backoff for rate limiting
- Validate SSL certificates
- Use the minimum required permission level

### For Server Administrators:
- Enable HTTPS in production
- Regularly rotate API keys
- Monitor authentication logs
- Use strong passwords for user accounts
- Keep GitHub OAuth2 credentials secure

## Rate Limiting

API requests are subject to rate limiting based on authentication method and user role:

- **Read operations**: 1000 requests/hour
- **Basic operations**: 500 requests/hour  
- **Admin operations**: 200 requests/hour

Rate limit headers are included in responses:
```
X-RateLimit-Limit: 500
X-RateLimit-Remaining: 487
X-RateLimit-Reset: 1640995200
```

## Code Examples

### Python with requests:

```python
import requests

# Using API key
headers = {"X-API-Key": "your_api_key"}
response = requests.get("http://localhost:8080/api/system", headers=headers)

# Using session authentication
session = requests.Session()
login_data = {"username": "admin", "password": "password"}
session.post("http://localhost:8080/api/login", json=login_data)
response = session.get("http://localhost:8080/api/history")
```

### JavaScript/Node.js:

```javascript
// Using fetch with API key
const response = await fetch('http://localhost:8080/api/system', {
  headers: {
    'X-API-Key': 'your_api_key'
  }
});

// Using session authentication
const loginResponse = await fetch('http://localhost:8080/api/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ username: 'admin', password: 'password' }),
  credentials: 'include'
});

const historyResponse = await fetch('http://localhost:8080/api/history', {
  credentials: 'include'
});
```

### Go:

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    // Using API key
    client := &http.Client{}
    req, _ := http.NewRequest("GET", "http://localhost:8080/api/system", nil)
    req.Header.Set("X-API-Key", "your_api_key")
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    
    // Using session authentication
    loginData := map[string]string{
        "username": "admin",
        "password": "password",
    }
    jsonData, _ := json.Marshal(loginData)
    resp, _ = http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(jsonData))
    // Handle cookies from response for subsequent requests
}
```

## Troubleshooting

### Common Issues:

1. **401 Unauthorized**: Check API key format and validity
2. **403 Forbidden**: Verify user has required permission level
3. **Session expired**: Re-authenticate if using session cookies
4. **Rate limited**: Implement backoff and respect rate limits
5. **CORS issues**: Ensure proper CORS configuration for web clients

### Debug Tips:

- Check response headers for detailed error information
- Monitor rate limiting headers
- Verify SSL certificate validity
- Test authentication with curl first
- Check server logs for authentication events

## Migration Guide

### From Basic Auth to API Keys:
Replace HTTP Basic Authentication with API key headers:

```diff
- Authorization: Basic base64(username:password)
+ X-API-Key: your_api_key
```

### From Cookie Auth to API Keys:
For programmatic access, replace session cookies with API keys:

```diff
- Cookie: session=session_value
+ X-API-Key: your_api_key
```