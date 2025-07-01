# Security Policy

## Supported Versions

We provide security updates for the following versions of Subtitle Manager:

| Version  | Supported          |
| -------- | ------------------ |
| Latest   | :white_check_mark: |
| < Latest | :x:                |

We recommend always using the latest version to ensure you have the most recent
security fixes.

## Reporting a Vulnerability

We take security vulnerabilities seriously and appreciate your help in keeping
Subtitle Manager secure.

### How to Report

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, please report security vulnerabilities by:

1. **Email**: Send details to the project maintainers
2. **GitHub Security Advisories**: Use GitHub's private vulnerability reporting
   feature
3. **Direct Contact**: Reach out through the repository's communication channels

### What to Include

When reporting a vulnerability, please include:

- **Description**: A clear description of the vulnerability
- **Impact**: What could an attacker accomplish by exploiting this?
- **Reproduction**: Steps to reproduce the vulnerability
- **Environment**: Version numbers, operating system, configuration details
- **Suggested Fix**: If you have ideas for how to fix the issue

### Response Timeline

We aim to respond to security reports according to the following timeline:

- **Initial Response**: Within 48 hours of receiving the report
- **Assessment**: Within 5 business days, we'll provide an initial assessment
- **Updates**: Regular updates every 5 business days until resolution
- **Resolution**: Security fixes will be prioritized and released as soon as
  possible

### Security Fix Process

1. **Validation**: We'll validate and reproduce the reported vulnerability
2. **Assessment**: Determine the severity and impact of the vulnerability
3. **Fix Development**: Develop and test a security fix
4. **Release**: Deploy the fix in a new release
5. **Disclosure**: Coordinate responsible disclosure with the reporter

## Security Measures

Subtitle Manager implements several security measures:

### Input Validation

- All user inputs are validated and sanitized
- File paths are checked to prevent directory traversal attacks
- Subtitle file content is processed safely to prevent injection attacks

### Authentication & Authorization

- Role-based access control (RBAC) with configurable permissions
- Secure session management with proper expiration
- Support for OAuth2, API keys, and password authentication
- Rate limiting to prevent brute force attacks

### Data Protection

- Sensitive configuration data can be encrypted
- API keys and credentials are never logged
- Secure HTTP headers including CSRF protection
- TLS enforcement for external communications

### Infrastructure Security

- Docker containers run with non-root users
- Minimal container images to reduce attack surface
- Regular dependency updates and vulnerability scanning
- Proper file permissions and temporary file handling

## Security Best Practices

### For Users

- Keep Subtitle Manager updated to the latest version
- Use strong, unique passwords for user accounts
- Enable OAuth2 authentication when possible
- Regularly review user permissions and API keys
- Monitor logs for suspicious activity
- Use HTTPS for web interface access
- Properly configure firewalls and access controls

### For Developers

- Follow secure coding practices outlined in our security documentation
- Validate all inputs and sanitize outputs
- Use parameterized queries to prevent SQL injection
- Implement proper error handling without information leakage
- Regularly update dependencies and monitor for vulnerabilities
- Conduct security reviews for new features
- Test security controls before deployment

## Additional Security Resources

- [Proposed Security Updates](docs/PROPOSED_SECURITY_UPDATES.md) - Comprehensive
  security enhancement plan
- [Technical Design Security](docs/TECHNICAL_DESIGN.md#security-considerations) -
  Architecture security details
- [GitHub Security Guidelines](.github/security-guidelines.md) - Development
  security practices

## Security Considerations by Feature

### Web Interface

- Input sanitization using DOMPurify
- CSRF token validation for state-changing operations
- Secure session management
- XSS prevention through proper output encoding

### File Processing

- Path traversal prevention for subtitle files
- File size limits to prevent DoS attacks
- Safe parsing of subtitle file formats
- Temporary file cleanup

### API Endpoints

- Authentication required for all operations
- Rate limiting on sensitive endpoints
- Input validation for all parameters
- Proper error handling without information disclosure

### External Integrations

- Secure API key management for translation services
- HTTPS enforcement for external requests
- Request timeout limits to prevent hanging connections
- Validation of external service responses

## Acknowledgments

We appreciate the security research community's efforts to improve software
security. Security researchers who responsibly disclose vulnerabilities may be
acknowledged in our security advisories (with their permission).

## Contact

For security-related questions or concerns that are not vulnerabilities, please
open a regular GitHub issue or contact the maintainers through the repository's
established communication channels.

---

This security policy is subject to change. Please check back regularly for
updates.
