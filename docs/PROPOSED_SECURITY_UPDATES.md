# file: PROPOSED_SECURITY_UPDATES.md

# Proposed Security Updates for Subtitle Manager

This document outlines a comprehensive plan to secure the JavaScript components and Go backend of the Subtitle Manager project. We propose a layered security approach that includes code sanitization, permission models, third-party libraries, and thorough testing. The following sections describe common exploits, how they apply to this project, and code examples demonstrating mitigations.

---

## Table of Contents

1. [Overview of Potential Exploits](#overview-of-potential-exploits)
2. [Sanitizing User Input](#sanitizing-user-input)
3. [Cross-Site Scripting (XSS)](#cross-site-scripting-xss)
4. [SQL Injection and Database Safety](#sql-injection-and-database-safety)
5. [Path Traversal and File Access](#path-traversal-and-file-access)
6. [Server-Side Request Forgery (SSRF)](#server-side-request-forgery-ssrf)
7. [Cross-Site Request Forgery (CSRF)](#cross-site-request-forgery-csrf)
8. [Authentication and Authorization Improvements](#authentication-and-authorization-improvements)
9. [Node.js 24 Permissions](#nodejs-24-permissions)
10. [Code Scanning and Dependency Management](#code-scanning-and-dependency-management)
11. [Library Recommendations](#library-recommendations)
12. [Example Implementation Steps](#example-implementation-steps)
13. [Testing Strategy](#testing-strategy)
14. [Long-Term Maintenance](#long-term-maintenance)

---

## Overview of Potential Exploits

The Subtitle Manager application combines a Go backend with a React front-end. Because it accepts user input from uploaded subtitle files, API requests, and web forms, it is susceptible to multiple classes of vulnerabilities. Below are key exploit categories we must consider.

### 1. Improper Input Handling

- Malicious data in HTTP request parameters or JSON payloads could be passed directly to database queries or filesystem operations.
- Without sanitization, attackers might inject arbitrary commands or cause crashes.

### 2. Cross-Site Scripting (XSS)

- Users could upload or submit text that includes `<script>` tags or other HTML/JavaScript payloads.
- If the React UI or API responses echo this data without sanitization, it can execute in another user's browser.

### 3. SQL Injection

- Even though Go's standard database libraries help prevent injection, mistakes in query building could expose SQL injection vectors.
- Input that is not sanitized might be combined with raw SQL statements.

### 4. Path Traversal

- Previous security fixes addressed path traversal, but we must ensure new functionality does not reintroduce it.
- Attackers might manipulate file paths to access sensitive files outside the intended directories.

### 5. SSRF

- The application communicates with external translation services. If URLs are accepted from users, there is a risk of SSRF.

### 6. CSRF

- The web interface includes forms for user management and settings. Without CSRF protection, attackers could forge requests from an authenticated user.

### 7. Authentication Weaknesses

- Token handling, session expiration, and permission checks must be consistent across the API.

### 8. Dependencies with Known Vulnerabilities

- Outdated packages can introduce vulnerabilities. Continuous scanning is necessary.

---

## Sanitizing User Input

### Why Sanitization Matters

Untrusted input should always be sanitized before further processing. This prevents injection attacks and ensures data conforms to expected formats.

### Recommended Libraries

For JavaScript, we recommend using the following:

- **DOMPurify**: A proven library that sanitizes HTML and prevents XSS by removing dangerous content.
- **express-validator** (for API endpoints implemented in Node): Provides declarative validation of request parameters and bodies.

For Go backend parameters:

- Use strict type conversions and explicit validation functions.
- Avoid directly inserting user input into shell commands or SQL queries.

### Example – Sanitizing Form Input in React

```jsx
import DOMPurify from "dompurify";

function safeContent(userInput) {
  return { __html: DOMPurify.sanitize(userInput) };
}

// Usage within a React component
export default function Comment({ text }) {
  return <div dangerouslySetInnerHTML={safeContent(text)} />;
}
```

In this example, `DOMPurify.sanitize` removes any unexpected or dangerous HTML from `text` before it is inserted into the DOM.

### Example – Validating API Parameters with express-validator

```javascript
import { body, validationResult } from "express-validator";

export const createSubtitleValidator = [
  body("title").isString().trim().isLength({ min: 1, max: 100 }),
  body("language").isISO31661Alpha2(),
];

export function handleValidation(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }
  next();
}
```

Any route that creates or updates subtitles should use this validator to ensure only proper input is accepted.

---

## Cross-Site Scripting (XSS)

### Exploit Scenario

1. A malicious user uploads a subtitle file containing a string like `<script>alert('XSS')</script>`.
2. If the server or UI displays the subtitle text directly without sanitization, another user's browser will execute the script, compromising their session.

### Mitigation Strategy

- Sanitize all user-supplied text before rendering it in the browser. Use DOMPurify as shown earlier.
- Avoid using `dangerouslySetInnerHTML` unless sanitized.
- Configure Content Security Policy (CSP) headers from the Go backend to restrict script execution sources.

### Example – Setting CSP Headers in Go

```go
func setSecurityHeaders(w http.ResponseWriter) {
    w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.Header().Set("X-Frame-Options", "DENY")
}
```

This function can be used in middleware to ensure all responses include strong security headers, reducing XSS impact.

---

## SQL Injection and Database Safety

### Exploit Scenario

Consider a backend endpoint that performs a raw SQL query using user input:

```go
query := fmt.Sprintf("SELECT * FROM subtitles WHERE user = '%s'", userInput)
rows, err := db.Query(query)
```

If `userInput` contains `' OR 1=1 --`, an attacker could retrieve every record.

### Mitigation Strategy

- Always use prepared statements with parameter placeholders.
- Validate and sanitize user inputs before using them in queries.

### Secure Example

```go
stmt, err := db.Prepare("SELECT * FROM subtitles WHERE user = $1")
if err != nil {
    return err
}
rows, err := stmt.Query(userInput)
```

Prepared statements ensure user input is treated as data, not as part of the SQL command.

### Additional Precautions

- Enforce least-privilege on database roles. Application accounts should only have necessary permissions.
- Enable parameterized queries in all database access code. Review all existing Go files for potential string concatenation in SQL queries.

---

## Path Traversal and File Access

### Exploit Scenario

A user submits a filename like `../../../../etc/passwd` via a web interface. If the backend uses this path directly with `os.Open` or `os.Stat`, it might leak sensitive files.

### Mitigation Strategy

- Continue using the `validateAndSanitizePath` function (now in `pkg/security/security.go`).
- Normalize paths using `filepath.Clean` and restrict them to allowed directories.

### Example Validation Function

```go
func validateAndSanitizePath(baseDir, inputPath string) (string, error) {
    cleaned := filepath.Clean(inputPath)
    fullPath := filepath.Join(baseDir, cleaned)
    if !strings.HasPrefix(fullPath, filepath.Clean(baseDir)) {
        return "", errors.New("path traversal detected")
    }
    return fullPath, nil
}
```

This ensures that even if the user tries to traverse directories, the resulting path stays within `baseDir`.

### Additional Measures

- Consider running the application with limited filesystem permissions to further isolate it from sensitive files.
- Store user-uploaded files in a separate directory with restricted access.

---

## Server-Side Request Forgery (SSRF)

### Exploit Scenario

The translation functionality might allow users to specify external URLs for custom translation providers. If unchecked, an attacker could trick the server into accessing internal services (`http://localhost:3000/metadata`), exposing data.

### Mitigation Strategy

- Implement URL validation to ensure only allowed domains are reachable.
- Use allowlists for external services.
- When possible, avoid letting users provide arbitrary URLs.

### Example URL Validation

```go
func validateExternalURL(rawURL string) (string, error) {
    u, err := url.ParseRequestURI(rawURL)
    if err != nil {
        return "", err
    }
    allowedHosts := []string{"translate.google.com", "api.example.com"}
    for _, host := range allowedHosts {
        if strings.EqualFold(u.Host, host) {
            return u.String(), nil
        }
    }
    return "", errors.New("host not allowed")
}
```

This code ensures that only approved hosts can be contacted, preventing SSRF via internal addresses.

---

## Cross-Site Request Forgery (CSRF)

### Exploit Scenario

Without CSRF protection, if a logged-in user visits a malicious site, that site could submit a hidden form to the Subtitle Manager API, performing actions on behalf of the user.

### Mitigation Strategy

- Use secure cookies with the `SameSite` attribute set to `Strict` or `Lax`.
- Implement CSRF tokens in all state-changing forms and API calls.

### CSRF Token Example in Express

```javascript
import csrf from "csurf";
import cookieParser from "cookie-parser";

const csrfProtection = csrf({ cookie: true });
app.use(cookieParser());
app.use(csrfProtection);

app.get("/form", (req, res) => {
  res.render("form", { csrfToken: req.csrfToken() });
});

app.post("/process", (req, res) => {
  res.send("data is being processed");
});
```

For the Go backend, use middleware packages like `gorilla/csrf`.

---

## Authentication and Authorization Improvements

### Current State

Subtitle Manager already features user accounts, API keys, and role-based access control (RBAC). However, we should re-evaluate our token handling and permission checks.

### Key Points

- Ensure token expiration times are enforced and rotated.
- Protect admin-only endpoints with explicit role checks.
- Audit logs should record authentication and authorization failures.

### Example – Role Middleware in Go

```go
func requireRole(role string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user := getUserFromContext(r.Context())
        if user == nil || !user.HasRole(role) {
            http.Error(w, "forbidden", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

This middleware ensures that certain handlers can only be accessed by users with the required role.

---

## Node.js 24 Permissions

### Overview

Node.js version 24 introduces an experimental permissions model that restricts filesystem, network, and child process access for your scripts. If we upgrade our build tools or any server-side JavaScript to Node 24, we can leverage this feature for stronger isolation.

### Potential Benefits

- Prevent unauthorized file access by specifying allowed directories.
- Limit network calls to predetermined domains.
- Block spawning new processes unless explicitly permitted.

### Example Node 24 Permission Usage

```bash
node --allow-fs-read=./data --allow-net=api.example.com server.js
```

This command restricts the script to read from `./data` and only make outbound requests to `api.example.com`.

### Migration Considerations

- Updating to Node 24 may require dependency updates. We should test thoroughly in a staging environment.
- Some packages may not yet support Node 24 features, so consider a gradual rollout.

---

## Code Scanning and Dependency Management

### Current Tools

GitHub CodeQL scanning is already enabled for the repository. However, we can bolster security by adding additional scanners and automation.

### Proposed Enhancements

- Integrate [`npm audit`](https://docs.npmjs.com/cli/v10/commands/npm-audit) to check for vulnerable front-end packages.
- Use [`govulncheck`](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) for Go dependencies.
- Enable Dependabot for automatic dependency update PRs.
- Schedule weekly or daily scans, with notifications for critical vulnerabilities.

### Example – Adding govulncheck to CI

```yaml
# .github/workflows/security.yml
jobs:
  govulncheck:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Install Go
        uses: actions/setup-go@v5
        with:
          go-version: 1.22
      - name: Run govulncheck
        run: govulncheck ./...
```

This workflow ensures Go dependencies are continuously monitored for known issues.

---

## Library Recommendations

Below are libraries and utilities we can adopt to streamline security.

| Area              | Library                              | Purpose                                     |
| ----------------- | ------------------------------------ | ------------------------------------------- |
| Input Validation  | `express-validator`                  | Request validation in Node.js/Express       |
| HTML Sanitization | `DOMPurify`                          | Cleans user-provided HTML                   |
| CSRF Protection   | `csurf` (Node) / `gorilla/csrf` (Go) | Mitigates CSRF attacks                      |
| Rate Limiting     | `express-rate-limit`                 | Prevents brute-force attacks on endpoints   |
| Security Headers  | `helmet`                             | Sets standard security-related HTTP headers |
| Dependency Scans  | `npm audit` / `govulncheck`          | Detects vulnerable packages                 |

### Example – Integrating Helmet for Express

```javascript
import helmet from "helmet";
app.use(helmet());
```

Even though our primary backend is in Go, we may have Node.js scripts or microservices that can benefit from this setup.

---

## Example Implementation Steps

1. **Audit Existing Code** – Review every route handler in `pkg/webserver`, verifying parameter parsing and sanitization.
2. **Integrate DOMPurify** – Add DOMPurify to `webui/src` and sanitize all user-facing data before rendering.
3. **Implement express-validator** – If any server-side JavaScript (such as CLI helpers or Node-based microservices) accepts input, enforce strict validation.
4. **Add Security Headers** – Create middleware in Go that sets CSP, `X-Frame-Options`, `X-Content-Type-Options`, and `Strict-Transport-Security` headers.
5. **CSRF Tokens** – Integrate CSRF tokens in forms for user settings, login, and any stateful operations. Consider `gorilla/csrf` for Go or `csurf` for Node.
6. **Rate Limiting** – Protect authentication endpoints with rate limiting to mitigate brute-force attempts.
7. **Enable Node 24 Permissions** – Where Node.js is used, run scripts with restricted permissions to limit the impact of any exploit.
8. **Continuous Scanning** – Set up CI jobs for `govulncheck` and `npm audit`, and consider enabling dependabot.
9. **Update Documentation** – Document security best practices in the README and add contributor guidelines for writing secure code.

---

## Testing Strategy

Security features must be validated with automated tests. Below is an outline of how we can ensure our security measures are effective.

### Unit Tests

- **Input Validation** – Test that invalid data is rejected by our validation functions.
- **Sanitization** – Verify that dangerous HTML is removed from user input.
- **Permission Checks** – Ensure middleware denies access when roles do not match.

### Integration Tests

- Simulate API requests containing potential XSS payloads to confirm the output is sanitized.
- Run queries with attempted SQL injection strings and verify that they do not alter results or cause errors.

### End-to-End Tests

- Use Playwright to automate UI interactions, ensuring that CSRF tokens are required for all sensitive actions.
- Confirm that rate limiting triggers after repeated failed logins.

### Example – Vitest for React Input Sanitization

```jsx
import { render, screen } from "@testing-library/react";
import Comment from "../Comment";

test("sanitizes dangerous HTML", () => {
  render(<Comment text="<img src=x onerror=alert(1)>" />);
  const img = screen.queryByRole("img");
  expect(img).toBeNull();
});
```

This test ensures that an image tag with an `onerror` handler is removed by DOMPurify.

### CI Integration

- All tests should run in GitHub Actions on each pull request.
- Add security-specific test suites for new middleware and permission logic.

---

## Long-Term Maintenance

### Regular Reviews

- Schedule quarterly security reviews to revisit dependencies and audit new code for vulnerabilities.
- Monitor CVE databases and mailing lists for updates affecting our dependencies.

### Developer Education

- Provide team training on secure coding practices, with a focus on input validation, sanitization, and error handling.
- Maintain documentation in `SECURITY_FIXES.md` and this proposal to keep developers informed of best practices.

### Incident Response

- Define a clear incident response plan, including contact information and steps to contain any future breach.
- Store logs in a central, tamper-evident location for analysis.

---

## Conclusion

By adopting the steps outlined in this document, Subtitle Manager will be hardened against common web exploits and benefit from modern security practices. Leveraging libraries such as DOMPurify and express-validator, enforcing Node 24 permissions, and incorporating continuous scanning will drastically reduce the application's attack surface.

Security is an ongoing process. This proposal provides a roadmap for immediate improvements and long-term maintenance, ensuring that Subtitle Manager remains resilient as it evolves.

---

```
End of Proposed Security Updates
```
