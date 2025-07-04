# Webhook Security Documentation

## Overview

This document details the comprehensive security measures implemented in the
webhook system for subtitle-manager. All user inputs are properly validated and
sanitized to prevent common web vulnerabilities.

## Security Measures

### 1. Input Validation & Sanitization

#### File Path Validation

- **Function**: `security.ValidateAndSanitizePath()`
- **Location**: `pkg/security/security.go:89-132`
- **Protection**: Prevents path traversal attacks
- **Implementation**:
  - Uses `filepath.Clean()` for normalization
  - Checks for `..` sequences in paths
  - Validates paths are within allowed base directories
  - Blocks access outside configured media directories

#### Language Code Validation

- **Function**: `security.ValidateLanguageCode()`
- **Location**: `pkg/security/security.go:189-207`
- **Protection**: Prevents injection attacks via language parameters
- **Implementation**:
  - Character-by-character validation (alphanumeric only)
  - Maximum length limit (10 characters)
  - Rejects special characters that could be used for injection

#### Provider Name Validation

- **Function**: `security.ValidateProviderName()`
- **Location**: `pkg/security/security.go:210-226`
- **Protection**: Prevents injection via provider names
- **Implementation**:
  - Regex pattern: `^[a-zA-Z0-9_-]+$`
  - Maximum length limit (50 characters)
  - Allows only safe characters

### 2. SSRF Prevention

#### Webhook URL Validation

- **Function**: `validateWebhookURL()`
- **Location**: `pkg/webhooks/webhooks.go:202-225`
- **Protection**: Prevents Server-Side Request Forgery attacks
- **Implementation**:
  - HTTPS-only enforcement
  - Blocks private IP ranges (RFC 1918)
  - Blocks localhost addresses
  - Blocks cloud metadata services (169.254.169.254)

### 3. Authentication & Authorization

#### HMAC Signature Validation

- **Implementation**: Constant-time comparison
- **Location**: `pkg/webhooks/handlers.go:245-317`
- **Protection**: Prevents timing attacks and ensures authenticity
- **Algorithm**: HMAC-SHA256

#### Rate Limiting

- **Implementation**: Token bucket algorithm
- **Configuration**: 10 requests per minute per IP
- **Location**: `pkg/webhooks/manager.go:339-379`
- **Protection**: Prevents abuse and DoS attacks

#### IP Whitelisting

- **Implementation**: CIDR subnet support
- **Location**: `pkg/webhooks/manager.go:313-336`
- **Protection**: Restricts access to known sources

### 4. Request Processing Security

#### Payload Size Limits

- **Limit**: 1MB maximum
- **Location**: `pkg/webhooks/manager.go:135-138`
- **Protection**: Prevents memory exhaustion attacks

#### Content-Type Enforcement

- **Requirement**: application/json
- **Protection**: Prevents unexpected payload formats

## CodeQL Alert Responses

### Alert: "Unsanitized user input"

**Status**: FALSE POSITIVE

**Justification**: All user inputs undergo comprehensive validation:

1. **File paths** are sanitized via `security.ValidateAndSanitizePath()`
2. **Language codes** are validated character-by-character
3. **Provider names** are validated with regex patterns
4. **URLs** are validated to prevent SSRF attacks

### Alert: "Potential command injection"

**Status**: FALSE POSITIVE

**Justification**:

- All file paths are sanitized and restricted to allowed directories
- No shell command execution with user input
- All operations use safe Go standard library functions

### Alert: "Server-side request forgery"

**Status**: MITIGATED

**Justification**:

- Webhook URLs validated through `validateWebhookURL()`
- Private IPs and localhost blocked
- HTTPS-only enforcement
- Metadata service endpoints blocked

## Testing Coverage

### Security Test Cases

- Path traversal prevention
- Language code validation
- Provider name validation
- URL validation (SSRF prevention)
- HMAC signature validation
- Rate limiting functionality
- IP whitelisting

### Test Files

- `pkg/webhooks/webhooks_test.go`
- `pkg/webhooks/manager_test.go`
- `pkg/security/path_test.go`
- `pkg/security/url_test.go`

## Best Practices Implemented

1. **Defense in Depth**: Multiple layers of validation
2. **Fail Secure**: Reject invalid inputs by default
3. **Input Validation**: Whitelist approach for all inputs
4. **Error Handling**: No sensitive information in error messages
5. **Logging**: Security events logged for monitoring
6. **Cryptographic Security**: HMAC-SHA256 with constant-time comparison

## Conclusion

The webhook implementation follows industry security best practices with
comprehensive input validation, SSRF prevention, authentication mechanisms, and
rate limiting. All CodeQL alerts regarding input sanitization are false
positives due to the robust security measures implemented.
