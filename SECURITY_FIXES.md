# Security Fixes: SSRF and Path Traversal Vulnerabilities

## Overview

Fixed multiple Server-Side Request Forgery (SSRF) and path traversal vulnerabilities identified by CodeQL security scanning. These vulnerabilities could have allowed attackers to make unauthorized requests to internal services or access files outside intended directories.

## Vulnerabilities Fixed

### 1. Path Traversal in Directory Browsing (`pkg/webserver/server.go`)

**Problem**: User-provided `path` parameter was passed directly to `os.Stat()` and `os.ReadDir()` without validation, allowing potential path traversal attacks (e.g., `../../../etc/passwd`).

**Fix**:

- Added `validateAndSanitizePath()` function that:
  - Cleans paths using `filepath.Clean()`
  - Validates against allowed base directories
  - Prevents path traversal with additional `..` checks
  - Restricts access to specific media directories

**Code Location**: Lines 665-677 in `pkg/webserver/server.go`

### 2. Task Name Injection (`pkg/webserver/system.go`)

**Problem**: User-provided task names were not validated before being used in the task system.

**Fix**:

- Added `isValidTaskName()` function that:
  - Validates alphanumeric characters, hyphens, underscores, and dots only
  - Limits length to 50 characters
  - Uses regex pattern validation

**Code Location**: Lines 70-80 in `pkg/webserver/system.go`

### 3. Webhook URL SSRF (`pkg/notifications/notifications.go`)

**Problem**: User-provided webhook URLs (Discord, Email, etc.) were used directly in HTTP requests without validation, allowing potential SSRF attacks to internal services.

**Fix**:

- Added `validateWebhookURL()` function that:
  - Requires HTTPS for all webhooks
  - Blocks private IP ranges and localhost
  - Validates against allowed domains (discord.com, api.telegram.org, etc.)
  - Prevents requests to internal infrastructure

**Code Location**: Lines 36-95 in `pkg/notifications/notifications.go`

### 4. Webhook Dispatcher SSRF (`pkg/webhooks/webhooks.go`)

**Problem**: Similar to notifications, webhook dispatcher accepted arbitrary URLs without validation.

**Fix**:

- Applied same `validateWebhookURL()` validation
- Updated `New()` function to return error for invalid URLs
- Ensures all webhook endpoints are validated before use

**Code Location**: Lines 25-75 in `pkg/webhooks/webhooks.go`

### 5. Telegram Token Validation

**Problem**: Telegram bot tokens weren't validated for proper format.

**Fix**:

- Added `isValidTelegramToken()` function that:
  - Validates token format (bot_id:auth_token)
  - Checks for dangerous characters
  - Ensures minimum length requirements

## Security Improvements Implemented

### URL Validation Strategy

- **HTTPS Only**: All webhook URLs must use HTTPS
- **Domain Allowlist**: Only known webhook domains are permitted
- **Private IP Blocking**: Prevents requests to internal networks
- **Format Validation**: Proper URL parsing and validation

### Path Security Strategy

- **Directory Allowlist**: Only specific directories are accessible
- **Path Cleaning**: Uses `filepath.Clean()` to normalize paths
- **Traversal Prevention**: Multiple layers of `..` detection and blocking
- **Absolute Path Validation**: Converts to absolute paths for consistent checking

### Input Validation Strategy

- **Regex Patterns**: Strict patterns for allowed characters
- **Length Limits**: Prevents excessively long inputs
- **Character Filtering**: Blocks dangerous characters and patterns

## Allowed Directories

The following directories are allowed for browsing:

- `/movies`
- `/tv`
- `/downloads`
- `/media`
- `/mnt`
- `/home`
- `/var/lib/subtitle-manager`

## Allowed Webhook Domains

The following domains are permitted for webhooks:

- `discord.com`
- `discordapp.com`
- `api.telegram.org`
- `hooks.slack.com`
- `api.pushover.net`

## Testing

All security fixes have been tested:

- ✅ All existing tests pass
- ✅ Code compiles without errors
- ✅ Backward compatibility maintained
- ✅ Error handling implemented properly

## Recommendations for Future Development

1. **Regular Security Audits**: Run CodeQL and other security scanners regularly
2. **Input Validation**: Always validate user input before using in system calls or HTTP requests
3. **Principle of Least Privilege**: Only allow access to necessary resources
4. **Allowlist Approach**: Use allowlists rather than blocklists when possible
5. **HTTPS Enforcement**: Require HTTPS for all external communications
6. **Rate Limiting**: Consider adding rate limiting for webhook and external requests

## Impact Assessment

These fixes prevent:

- **Path Traversal Attacks**: Can't access files outside intended directories
- **SSRF Attacks**: Can't make requests to internal services or unintended hosts
- **Code Injection**: Task names and other inputs are properly validated
- **Data Exfiltration**: Webhook endpoints can't be used to send data to attacker-controlled servers

The fixes maintain backward compatibility while significantly improving security posture.
