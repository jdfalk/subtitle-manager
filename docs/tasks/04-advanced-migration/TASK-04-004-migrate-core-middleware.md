# TASK-04-004: Migrate Core Middleware and Error Types

<!-- file: docs/tasks/04-advanced-migration/TASK-04-004-migrate-core-middleware.md -->
<!-- version: 1.0.0 -->
<!-- guid: f4g5h6i7-8901-2345-6789-012345678901 -->

## 🎯 Task Overview

**Primary Objective**: Migrate all middleware components and error handling to
use gcommon types with opaque API

**Task Type**: Core Middleware Migration

**Estimated Effort**: 1-2 hours

**Dependencies**:

- TASK-04-001 (User type migration) completed
- TASK-04-002 (Session type migration) completed
- TASK-04-003 (Auth flow updates) completed

## 📋 Acceptance Criteria

- [ ] All middleware uses gcommon types with opaque API
- [ ] Error handling updated to gcommon error types
- [ ] Logging middleware works with new types
- [ ] CORS middleware integrated
- [ ] Request validation middleware updated
- [ ] Context management properly handles gcommon types
- [ ] Middleware chain functions correctly

## 🔄 Dependencies

**Input Requirements**:

- Authentication flow fully migrated
- User and Session types updated
- Database stores compatible with gcommon types

**External Dependencies**:

- gcommon/sdks/go/v1/common package
- HTTP middleware patterns
- Logging framework

## 📝 Implementation Steps

### Step 1: Update middleware types and interfaces

**Create `pkg/middleware/types.go`**:

```go
package middleware

import (
    "context"
    "net/http"

    "github.com/jdfalk/gcommon/sdks/go/v1/common"
)

// Context keys for middleware
type contextKey string

const (
    UserContextKey    contextKey = "user"
    SessionContextKey contextKey = "session"
    RequestIDKey      contextKey = "request_id"
    StartTimeKey      contextKey = "start_time"
)

// Middleware interface
type Middleware interface {
    Handler(next http.Handler) http.Handler
}

// Helper functions for context management
func SetUserInContext(ctx context.Context, user *common.User) context.Context {
    return context.WithValue(ctx, UserContextKey, user)
}

func GetUserFromContext(ctx context.Context) (*common.User, bool) {
    user, ok := ctx.Value(UserContextKey).(*common.User)
    return user, ok
}

func SetSessionInContext(ctx context.Context, session *common.Session) context.Context {
    return context.WithValue(ctx, SessionContextKey, session)
}

func GetSessionFromContext(ctx context.Context) (*common.Session, bool) {
    session, ok := ctx.Value(SessionContextKey).(*common.Session)
    return session, ok
}

func SetRequestIDInContext(ctx context.Context, requestID string) context.Context {
    return context.WithValue(ctx, RequestIDKey, requestID)
}

func GetRequestIDFromContext(ctx context.Context) (string, bool) {
    requestID, ok := ctx.Value(RequestIDKey).(string)
    return requestID, ok
}
```

### Step 2: Update authentication middleware

**Create `pkg/middleware/auth.go`**:

```go
package middleware

import (
    "context"
    "crypto/rand"
    "encoding/hex"
    "log"
    "net/http"
    "time"

    "github.com/jdfalk/subtitle-manager/pkg/auth"
)

type AuthMiddleware struct {
    authService auth.AuthService
    logger      *log.Logger
}

func NewAuthMiddleware(authService auth.AuthService, logger *log.Logger) *AuthMiddleware {
    return &AuthMiddleware{
        authService: authService,
        logger:      logger,
    }
}

func (a *AuthMiddleware) Handler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Add request ID for tracing
        requestID := generateRequestID()
        ctx := SetRequestIDInContext(r.Context(), requestID)
        ctx = context.WithValue(ctx, StartTimeKey, time.Now())
        r = r.WithContext(ctx)

        // Check for session cookie
        cookie, err := r.Cookie("session_id")
        if err != nil {
            // No session cookie, redirect to login
            a.redirectToLogin(w, r)
            return
        }

        // Validate session
        user, session, err := a.authService.ValidateSession(r.Context(), cookie.Value)
        if err != nil {
            a.logger.Printf("Session validation failed for request %s: %v", requestID, err)
            a.redirectToLogin(w, r)
            return
        }

        // Add user and session to context
        ctx = SetUserInContext(r.Context(), user)
        ctx = SetSessionInContext(ctx, session)

        // Log successful authentication
        a.logger.Printf("User %s authenticated for request %s (session: %s)",
            user.GetUsername(), requestID, session.GetId())

        // Continue to next handler
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func (a *AuthMiddleware) redirectToLogin(w http.ResponseWriter, r *http.Request) {
    // Clear any existing session cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    "",
        Path:     "/",
        Expires:  time.Unix(0, 0),
        HttpOnly: true,
    })

    // Redirect to login page
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func generateRequestID() string {
    bytes := make([]byte, 8)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}
```

### Step 3: Create API authentication middleware

**Create `pkg/middleware/api_auth.go`**:

```go
package middleware

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/jdfalk/subtitle-manager/pkg/auth"
)

type APIAuthMiddleware struct {
    authService auth.AuthService
    logger      *log.Logger
}

func NewAPIAuthMiddleware(authService auth.AuthService, logger *log.Logger) *APIAuthMiddleware {
    return &APIAuthMiddleware{
        authService: authService,
        logger:      logger,
    }
}

func (a *APIAuthMiddleware) Handler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Add request ID
        requestID := generateRequestID()
        ctx := SetRequestIDInContext(r.Context(), requestID)
        r = r.WithContext(ctx)

        // Get Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            a.sendAPIError(w, "Authorization header required", http.StatusUnauthorized)
            return
        }

        // Parse Bearer token
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            a.sendAPIError(w, "Invalid authorization format. Use: Bearer <token>", http.StatusUnauthorized)
            return
        }

        token := parts[1]
        if token == "" {
            a.sendAPIError(w, "Token cannot be empty", http.StatusUnauthorized)
            return
        }

        // Validate token
        user, err := a.authService.ValidateAuthToken(r.Context(), token)
        if err != nil {
            a.logger.Printf("Token validation failed for request %s: %v", requestID, err)
            a.sendAPIError(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        // Add user to context (API doesn't use sessions)
        ctx = SetUserInContext(r.Context(), user)

        // Log successful API authentication
        a.logger.Printf("API user %s authenticated for request %s", user.GetUsername(), requestID)

        // Continue to next handler
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func (a *APIAuthMiddleware) sendAPIError(w http.ResponseWriter, message string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    response := struct {
        Success bool   `json:"success"`
        Error   string `json:"error"`
    }{
        Success: false,
        Error:   message,
    }

    json.NewEncoder(w).Encode(response)
}
```

### Step 4: Create logging middleware

**Create `pkg/middleware/logging.go`**:

```go
package middleware

import (
    "log"
    "net/http"
    "time"
)

type LoggingMiddleware struct {
    logger *log.Logger
}

func NewLoggingMiddleware(logger *log.Logger) *LoggingMiddleware {
    return &LoggingMiddleware{logger: logger}
}

func (l *LoggingMiddleware) Handler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Get request ID from context
        requestID, _ := GetRequestIDFromContext(r.Context())
        if requestID == "" {
            requestID = generateRequestID()
            ctx := SetRequestIDInContext(r.Context(), requestID)
            r = r.WithContext(ctx)
        }

        // Log request start
        l.logger.Printf("Started %s %s [%s] from %s",
            r.Method, r.URL.Path, requestID, r.RemoteAddr)

        // Get user info if available
        if user, ok := GetUserFromContext(r.Context()); ok {
            l.logger.Printf("Request %s authenticated as user: %s (ID: %s)",
                requestID, user.GetUsername(), user.GetId())
        }

        // Wrap response writer to capture status code
        wrapped := &responseWriter{
            ResponseWriter: w,
            statusCode:     http.StatusOK,
        }

        // Call next handler
        next.ServeHTTP(wrapped, r)

        // Log completion
        duration := time.Since(start)
        l.logger.Printf("Completed %s %s [%s] %d in %v",
            r.Method, r.URL.Path, requestID, wrapped.statusCode, duration)
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
```

### Step 5: Create CORS middleware

**Create `pkg/middleware/cors.go`**:

```go
package middleware

import (
    "net/http"
    "strings"
)

type CORSMiddleware struct {
    allowedOrigins []string
    allowedMethods []string
    allowedHeaders []string
    allowedCredentials bool
}

func NewCORSMiddleware(origins, methods, headers []string, credentials bool) *CORSMiddleware {
    return &CORSMiddleware{
        allowedOrigins:     origins,
        allowedMethods:     methods,
        allowedHeaders:     headers,
        allowedCredentials: credentials,
    }
}

func (c *CORSMiddleware) Handler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        origin := r.Header.Get("Origin")

        // Check if origin is allowed
        if c.isOriginAllowed(origin) {
            w.Header().Set("Access-Control-Allow-Origin", origin)
        }

        // Set other CORS headers
        if len(c.allowedMethods) > 0 {
            w.Header().Set("Access-Control-Allow-Methods", strings.Join(c.allowedMethods, ", "))
        }

        if len(c.allowedHeaders) > 0 {
            w.Header().Set("Access-Control-Allow-Headers", strings.Join(c.allowedHeaders, ", "))
        }

        if c.allowedCredentials {
            w.Header().Set("Access-Control-Allow-Credentials", "true")
        }

        // Handle preflight requests
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func (c *CORSMiddleware) isOriginAllowed(origin string) bool {
    for _, allowed := range c.allowedOrigins {
        if allowed == "*" || allowed == origin {
            return true
        }
    }
    return false
}
```

### Step 6: Create request validation middleware

**Create `pkg/middleware/validation.go`**:

```go
package middleware

import (
    "encoding/json"
    "net/http"
    "strings"
)

type ValidationMiddleware struct {
    maxRequestSize int64
    allowedPaths   []string
}

func NewValidationMiddleware(maxRequestSize int64, allowedPaths []string) *ValidationMiddleware {
    return &ValidationMiddleware{
        maxRequestSize: maxRequestSize,
        allowedPaths:   allowedPaths,
    }
}

func (v *ValidationMiddleware) Handler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Validate request size
        if r.ContentLength > v.maxRequestSize {
            v.sendValidationError(w, "Request body too large", http.StatusRequestEntityTooLarge)
            return
        }

        // Validate path
        if !v.isPathAllowed(r.URL.Path) {
            v.sendValidationError(w, "Path not allowed", http.StatusForbidden)
            return
        }

        // Validate content type for POST/PUT requests
        if r.Method == http.MethodPost || r.Method == http.MethodPut {
            contentType := r.Header.Get("Content-Type")
            if contentType == "" {
                v.sendValidationError(w, "Content-Type header required", http.StatusBadRequest)
                return
            }

            // Allow form data and JSON
            if !strings.Contains(contentType, "application/json") &&
               !strings.Contains(contentType, "application/x-www-form-urlencoded") &&
               !strings.Contains(contentType, "multipart/form-data") {
                v.sendValidationError(w, "Unsupported content type", http.StatusUnsupportedMediaType)
                return
            }
        }

        next.ServeHTTP(w, r)
    })
}

func (v *ValidationMiddleware) isPathAllowed(path string) bool {
    // Allow all paths if none specified
    if len(v.allowedPaths) == 0 {
        return true
    }

    for _, allowed := range v.allowedPaths {
        if strings.HasPrefix(path, allowed) {
            return true
        }
    }
    return false
}

func (v *ValidationMiddleware) sendValidationError(w http.ResponseWriter, message string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    response := struct {
        Success bool   `json:"success"`
        Error   string `json:"error"`
    }{
        Success: false,
        Error:   message,
    }

    json.NewEncoder(w).Encode(response)
}
```

### Step 7: Create middleware chain manager

**Create `pkg/middleware/chain.go`**:

```go
package middleware

import (
    "log"
    "net/http"

    "github.com/jdfalk/subtitle-manager/pkg/auth"
)

type MiddlewareChain struct {
    middlewares []Middleware
}

func NewMiddlewareChain() *MiddlewareChain {
    return &MiddlewareChain{
        middlewares: make([]Middleware, 0),
    }
}

func (mc *MiddlewareChain) Add(middleware Middleware) *MiddlewareChain {
    mc.middlewares = append(mc.middlewares, middleware)
    return mc
}

func (mc *MiddlewareChain) Then(handler http.Handler) http.Handler {
    // Build chain in reverse order
    for i := len(mc.middlewares) - 1; i >= 0; i-- {
        handler = mc.middlewares[i].Handler(handler)
    }
    return handler
}

// Predefined middleware chains
func WebMiddlewareChain(authService auth.AuthService, logger *log.Logger) *MiddlewareChain {
    return NewMiddlewareChain().
        Add(NewLoggingMiddleware(logger)).
        Add(NewValidationMiddleware(10*1024*1024, nil)). // 10MB max request
        Add(NewCORSMiddleware(
            []string{"http://localhost:3000", "http://localhost:8080"}, // Dev origins
            []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
            []string{"Content-Type", "Authorization", "X-Requested-With"},
            true,
        )).
        Add(NewAuthMiddleware(authService, logger))
}

func APIMiddlewareChain(authService auth.AuthService, logger *log.Logger) *MiddlewareChain {
    return NewMiddlewareChain().
        Add(NewLoggingMiddleware(logger)).
        Add(NewValidationMiddleware(5*1024*1024, nil)). // 5MB max for API
        Add(NewCORSMiddleware(
            []string{"*"}, // Allow all origins for API (configure as needed)
            []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
            []string{"Content-Type", "Authorization", "X-Requested-With"},
            false, // No credentials for API CORS
        )).
        Add(NewAPIAuthMiddleware(authService, logger))
}

func PublicMiddlewareChain(logger *log.Logger) *MiddlewareChain {
    return NewMiddlewareChain().
        Add(NewLoggingMiddleware(logger)).
        Add(NewValidationMiddleware(10*1024*1024, nil))
}
```

### Step 8: Update WebServer to use new middleware

**Update `pkg/webserver/server.go`**:

```go
package webserver

import (
    "log"
    "net/http"
    "os"

    "github.com/jdfalk/subtitle-manager/pkg/auth"
    "github.com/jdfalk/subtitle-manager/pkg/middleware"
)

type WebServer struct {
    server      *http.Server
    authService auth.AuthService
    logger      *log.Logger
    config      *Config
}

func NewWebServer(config *Config, authService auth.AuthService) *WebServer {
    logger := log.New(os.Stdout, "[WEB] ", log.LstdFlags|log.Lshortfile)

    ws := &WebServer{
        config:      config,
        authService: authService,
        logger:      logger,
    }

    ws.setupRoutes()
    return ws
}

func (s *WebServer) setupRoutes() {
    mux := http.NewServeMux()

    // Create middleware chains
    webChain := middleware.WebMiddlewareChain(s.authService, s.logger)
    apiChain := middleware.APIMiddlewareChain(s.authService, s.logger)
    publicChain := middleware.PublicMiddlewareChain(s.logger)

    // Public routes (no authentication required)
    mux.Handle("/login", publicChain.Then(http.HandlerFunc(s.handleLogin)))
    mux.Handle("/register", publicChain.Then(http.HandlerFunc(s.handleRegister)))
    mux.Handle("/health", publicChain.Then(http.HandlerFunc(s.handleHealth)))
    mux.Handle("/static/", publicChain.Then(http.FileServer(http.Dir("./web/"))))

    // Protected web routes
    mux.Handle("/", webChain.Then(http.HandlerFunc(s.handleHome)))
    mux.Handle("/profile", webChain.Then(http.HandlerFunc(s.handleProfile)))
    mux.Handle("/settings", webChain.Then(http.HandlerFunc(s.handleSettings)))
    mux.Handle("/logout", webChain.Then(http.HandlerFunc(s.handleLogout)))

    // API routes (token authentication)
    mux.Handle("/api/login", publicChain.Then(http.HandlerFunc(s.handleAPILogin)))
    mux.Handle("/api/register", publicChain.Then(http.HandlerFunc(s.handleAPIRegister)))
    mux.Handle("/api/subtitles", apiChain.Then(http.HandlerFunc(s.handleAPISubtitles)))
    mux.Handle("/api/movies", apiChain.Then(http.HandlerFunc(s.handleAPIMovies)))
    mux.Handle("/api/user/profile", apiChain.Then(http.HandlerFunc(s.handleAPIUserProfile)))

    s.server = &http.Server{
        Addr:    s.config.Address,
        Handler: mux,
    }
}

func (s *WebServer) handleHome(w http.ResponseWriter, r *http.Request) {
    user, ok := middleware.GetUserFromContext(r.Context())
    if !ok {
        http.Error(w, "User not found in context", http.StatusInternalServerError)
        return
    }

    data := struct {
        User      *common.User
        RequestID string
    }{
        User:      user,
        RequestID: getRequestID(r.Context()),
    }

    s.renderTemplate(w, "home.html", data)
}

func getRequestID(ctx context.Context) string {
    if id, ok := middleware.GetRequestIDFromContext(ctx); ok {
        return id
    }
    return "unknown"
}
```

## 🧪 Testing Requirements

### Unit Tests for Middleware

```go
func TestAuthMiddleware(t *testing.T) {
    authService := setupMockAuthService()
    logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)
    authMW := middleware.NewAuthMiddleware(authService, logger)

    // Test successful authentication
    req := httptest.NewRequest("GET", "/protected", nil)
    req.AddCookie(&http.Cookie{Name: "session_id", Value: "valid_session"})

    rr := httptest.NewRecorder()
    handler := authMW.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, ok := middleware.GetUserFromContext(r.Context())
        assert.True(t, ok)
        assert.Equal(t, "testuser", user.GetUsername())
        w.WriteHeader(http.StatusOK)
    }))

    handler.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAPIAuthMiddleware(t *testing.T) {
    authService := setupMockAuthService()
    logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)
    apiMW := middleware.NewAPIAuthMiddleware(authService, logger)

    // Test successful API authentication
    req := httptest.NewRequest("GET", "/api/test", nil)
    req.Header.Set("Authorization", "Bearer valid_token")

    rr := httptest.NewRecorder()
    handler := apiMW.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, ok := middleware.GetUserFromContext(r.Context())
        assert.True(t, ok)
        assert.Equal(t, "apiuser", user.GetUsername())
        w.WriteHeader(http.StatusOK)
    }))

    handler.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusOK, rr.Code)
}
```

### Integration Tests for Middleware Chain

```go
func TestMiddlewareChain(t *testing.T) {
    authService := setupTestAuthService()
    logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)

    chain := middleware.WebMiddlewareChain(authService, logger)

    handler := chain.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verify all context values are set
        user, hasUser := middleware.GetUserFromContext(r.Context())
        session, hasSession := middleware.GetSessionFromContext(r.Context())
        requestID, hasRequestID := middleware.GetRequestIDFromContext(r.Context())

        assert.True(t, hasUser)
        assert.True(t, hasSession)
        assert.True(t, hasRequestID)
        assert.NotEmpty(t, requestID)

        w.WriteHeader(http.StatusOK)
    }))

    // Create request with valid session
    req := httptest.NewRequest("GET", "/test", nil)
    req.AddCookie(&http.Cookie{Name: "session_id", Value: "valid_session"})

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
}
```

## 📚 Required Documentation

**Embedded from .github/instructions/general-coding.instructions.md:**

### Critical Guidelines

```markdown
## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS

**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction
of any kind.**

## Required File Header (File Identification)

All source, script, and documentation files MUST begin with a standard header.

## Version Update Requirements

**When modifying any file with a version header, ALWAYS update the version
number**
```

### Middleware Security Best Practices

1. **Context Safety**: Always type-assert context values safely
2. **Request Limits**: Implement reasonable request size limits
3. **CORS Configuration**: Configure CORS appropriately for environment
4. **Error Handling**: Don't leak sensitive information in errors
5. **Logging**: Log security events but not sensitive data

## 🎯 Success Metrics

- [ ] All middleware uses gcommon types with opaque API
- [ ] Authentication middleware works for web and API routes
- [ ] Logging middleware captures request flow
- [ ] CORS middleware properly configured
- [ ] Request validation prevents malformed requests
- [ ] Middleware chain functions correctly
- [ ] All middleware tests pass

## 🚨 Common Pitfalls

1. **Context Type Assertions**: Unsafe type assertions from context
2. **Middleware Order**: Incorrect middleware execution order
3. **CORS Misconfiguration**: Overly permissive or restrictive CORS
4. **Request Size Limits**: Missing or inappropriate size limits
5. **Error Information Leakage**: Exposing internals in error messages
6. **Logging Sensitive Data**: Accidentally logging passwords or tokens
7. **Chain Building**: Building middleware chains in wrong order

## 📖 Additional Resources

- [Go HTTP Middleware Patterns](https://www.alexedwards.net/blog/making-and-using-middleware)
- [CORS Security Guide](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
- [Request Validation Best Practices](https://cheatsheetseries.owasp.org/cheatsheets/Input_Validation_Cheat_Sheet.html)

## 🔄 Related Tasks

**Must complete before this task**:

- **TASK-04-001**: Migrate User Type
- **TASK-04-002**: Migrate Session Type
- **TASK-04-003**: Update Authentication Flow

**Enables these tasks**:

- **TASK-05-001**: Performance optimization
- **TASK-06-001**: Integration testing
- **TASK-06-002**: Load testing middleware

## 📝 Notes for AI Agent

- Use opaque API for all gcommon types in middleware
- Implement safe context value handling with type assertions
- Configure CORS appropriately for development and production
- Add comprehensive request validation
- Implement proper logging without sensitive data
- Test middleware chains thoroughly
- Ensure proper error handling and user feedback
- Verify middleware order for optimal performance

## 🔚 Completion Verification

```bash
# Verify middleware compilation
echo "Checking middleware compilation..."
go build ./pkg/middleware && echo "✅ Middleware compiles"

# Verify opaque API usage in middleware
echo "Checking middleware opaque API usage..."
grep -r "GetUsername()\|GetId()" --include="*.go" ./pkg/middleware && echo "✅ Middleware uses opaque API"

# Run middleware tests
echo "Running middleware tests..."
go test -run TestMiddleware ./... && echo "✅ Middleware tests pass"

# Test web server with new middleware
echo "Testing web server compilation..."
go build ./pkg/webserver && echo "✅ Web server with new middleware compiles"
```
