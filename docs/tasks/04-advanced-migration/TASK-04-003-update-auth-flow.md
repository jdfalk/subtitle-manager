# TASK-04-003: Update Authentication Flow End-to-End

<!-- file: docs/tasks/04-advanced-migration/TASK-04-003-update-auth-flow.md -->
<!-- version: 1.0.0 -->
<!-- guid: e3f4g5h6-7890-1234-5678-901234567890 -->

## 🎯 Task Overview

**Primary Objective**: Update complete authentication flow to use gcommon types
with opaque API patterns

**Task Type**: Authentication System Migration

**Estimated Effort**: 2-3 hours

**Dependencies**:

- TASK-04-001 (Migrate User Type) completed
- TASK-04-002 (Migrate Session Type) completed

## 📋 Acceptance Criteria

- [ ] Login flow uses gcommon User and Session types
- [ ] Registration process creates gcommon User types
- [ ] Password validation and hashing updated
- [ ] Token-based authentication implemented
- [ ] API authentication endpoints functional
- [ ] Web authentication handlers working
- [ ] Authentication middleware properly integrated

## 🔄 Dependencies

**Input Requirements**:

- TASK-04-001 completed (User type migrated)
- TASK-04-002 completed (Session type migrated)

**External Dependencies**:

- gcommon/sdks/go/v1/common package
- Updated database stores
- Authentication middleware

## 📝 Implementation Steps

### Step 1: Update authentication service interface

**Create `pkg/auth/service.go`**:

```go
package auth

import (
    "context"
    "crypto/rand"
    "crypto/subtle"
    "encoding/hex"
    "fmt"
    "time"

    "github.com/jdfalk/gcommon/sdks/go/v1/common"
    "golang.org/x/crypto/bcrypt"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService interface {
    // User management
    CreateUser(ctx context.Context, username, email, password string) (*common.User, error)
    AuthenticateUser(ctx context.Context, username, password string) (*common.User, error)
    GetUserByID(ctx context.Context, userID string) (*common.User, error)
    GetUserByUsername(ctx context.Context, username string) (*common.User, error)
    UpdateUserPassword(ctx context.Context, userID, newPassword string) error

    // Session management
    CreateSession(ctx context.Context, userID string) (*common.Session, error)
    GetSession(ctx context.Context, sessionID string) (*common.Session, error)
    ValidateSession(ctx context.Context, sessionID string) (*common.User, *common.Session, error)
    DeleteSession(ctx context.Context, sessionID string) error
    CleanupExpiredSessions(ctx context.Context) error

    // Token management
    GenerateAuthToken(ctx context.Context, userID string) (string, error)
    ValidateAuthToken(ctx context.Context, token string) (*common.User, error)
}

type authService struct {
    store DataStore
}

type DataStore interface {
    // User operations
    CreateUser(ctx context.Context, user *common.User) error
    GetUserByID(ctx context.Context, userID string) (*common.User, error)
    GetUserByUsername(ctx context.Context, username string) (*common.User, error)
    GetUserByEmail(ctx context.Context, email string) (*common.User, error)
    UpdateUser(ctx context.Context, user *common.User) error

    // Session operations
    CreateSession(ctx context.Context, sessionID, userID string, expiresAt time.Time) error
    GetSession(ctx context.Context, sessionID string) (*common.Session, error)
    DeleteSession(ctx context.Context, sessionID string) error
    GetSessionsByUserID(ctx context.Context, userID string) ([]*common.Session, error)
    CleanupExpiredSessions(ctx context.Context) error
}

func NewAuthService(store DataStore) AuthService {
    return &authService{store: store}
}
```

### Step 2: Implement user authentication methods

**Continue `pkg/auth/service.go`**:

```go
func (a *authService) CreateUser(ctx context.Context, username, email, password string) (*common.User, error) {
    // Check if username exists
    if _, err := a.store.GetUserByUsername(ctx, username); err == nil {
        return nil, fmt.Errorf("username already exists")
    }

    // Check if email exists
    if _, err := a.store.GetUserByEmail(ctx, email); err == nil {
        return nil, fmt.Errorf("email already exists")
    }

    // Hash password
    hashedPassword, err := hashPassword(password)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password: %w", err)
    }

    // Create user with opaque API
    user := &common.User{}
    userID := generateUserID()
    user.SetId(userID)
    user.SetUsername(username)
    user.SetEmail(email)
    user.SetPasswordHash(hashedPassword)
    user.SetCreatedAt(timestamppb.Now())
    user.SetUpdatedAt(timestamppb.Now())
    user.SetIsActive(true)

    err = a.store.CreateUser(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    return user, nil
}

func (a *authService) AuthenticateUser(ctx context.Context, username, password string) (*common.User, error) {
    user, err := a.store.GetUserByUsername(ctx, username)
    if err != nil {
        return nil, fmt.Errorf("invalid credentials")
    }

    if !user.GetIsActive() {
        return nil, fmt.Errorf("account is inactive")
    }

    if !verifyPassword(password, user.GetPasswordHash()) {
        return nil, fmt.Errorf("invalid credentials")
    }

    // Update last login
    user.SetLastLogin(timestamppb.Now())
    user.SetUpdatedAt(timestamppb.Now())

    err = a.store.UpdateUser(ctx, user)
    if err != nil {
        // Log error but don't fail authentication
        // The user is valid, login timestamp update is not critical
    }

    return user, nil
}

func (a *authService) GetUserByID(ctx context.Context, userID string) (*common.User, error) {
    return a.store.GetUserByID(ctx, userID)
}

func (a *authService) GetUserByUsername(ctx context.Context, username string) (*common.User, error) {
    return a.store.GetUserByUsername(ctx, username)
}

func (a *authService) UpdateUserPassword(ctx context.Context, userID, newPassword string) error {
    user, err := a.store.GetUserByID(ctx, userID)
    if err != nil {
        return fmt.Errorf("user not found: %w", err)
    }

    hashedPassword, err := hashPassword(newPassword)
    if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }

    user.SetPasswordHash(hashedPassword)
    user.SetUpdatedAt(timestamppb.Now())

    return a.store.UpdateUser(ctx, user)
}
```

### Step 3: Implement session management methods

**Continue `pkg/auth/service.go`**:

```go
func (a *authService) CreateSession(ctx context.Context, userID string) (*common.Session, error) {
    sessionID := generateSessionID()
    expiresAt := time.Now().Add(24 * time.Hour) // 24 hour session

    err := a.store.CreateSession(ctx, sessionID, userID, expiresAt)
    if err != nil {
        return nil, fmt.Errorf("failed to create session: %w", err)
    }

    // Return session object
    session := &common.Session{}
    session.SetId(sessionID)
    session.SetUserId(userID)
    session.SetExpiresAt(timestamppb.New(expiresAt))
    session.SetCreatedAt(timestamppb.Now())

    return session, nil
}

func (a *authService) GetSession(ctx context.Context, sessionID string) (*common.Session, error) {
    return a.store.GetSession(ctx, sessionID)
}

func (a *authService) ValidateSession(ctx context.Context, sessionID string) (*common.User, *common.Session, error) {
    session, err := a.store.GetSession(ctx, sessionID)
    if err != nil {
        return nil, nil, fmt.Errorf("session not found: %w", err)
    }

    // Check if session is expired
    if session.GetExpiresAt().AsTime().Before(time.Now()) {
        // Clean up expired session
        a.store.DeleteSession(ctx, sessionID)
        return nil, nil, fmt.Errorf("session expired")
    }

    // Get user associated with session
    user, err := a.store.GetUserByID(ctx, session.GetUserId())
    if err != nil {
        return nil, nil, fmt.Errorf("user not found: %w", err)
    }

    if !user.GetIsActive() {
        return nil, nil, fmt.Errorf("user account is inactive")
    }

    return user, session, nil
}

func (a *authService) DeleteSession(ctx context.Context, sessionID string) error {
    return a.store.DeleteSession(ctx, sessionID)
}

func (a *authService) CleanupExpiredSessions(ctx context.Context) error {
    return a.store.CleanupExpiredSessions(ctx)
}
```

### Step 4: Implement utility functions

**Continue `pkg/auth/service.go`**:

```go
// Password hashing and verification
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func verifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// ID generation
func generateUserID() string {
    return "user_" + generateRandomID()
}

func generateSessionID() string {
    return "sess_" + generateRandomID()
}

func generateRandomID() string {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
        // Fallback to timestamp-based ID if crypto/rand fails
        return fmt.Sprintf("%d", time.Now().UnixNano())
    }
    return hex.EncodeToString(bytes)
}

// Token-based authentication (for API endpoints)
func (a *authService) GenerateAuthToken(ctx context.Context, userID string) (string, error) {
    token := generateRandomID()
    // Store token with expiration (could be in database or cache)
    // For simplicity, using session table with longer expiration
    expiresAt := time.Now().Add(30 * 24 * time.Hour) // 30 days

    err := a.store.CreateSession(ctx, "token_"+token, userID, expiresAt)
    if err != nil {
        return "", fmt.Errorf("failed to create auth token: %w", err)
    }

    return token, nil
}

func (a *authService) ValidateAuthToken(ctx context.Context, token string) (*common.User, error) {
    session, err := a.store.GetSession(ctx, "token_"+token)
    if err != nil {
        return nil, fmt.Errorf("invalid token")
    }

    if session.GetExpiresAt().AsTime().Before(time.Now()) {
        a.store.DeleteSession(ctx, "token_"+token)
        return nil, fmt.Errorf("token expired")
    }

    user, err := a.store.GetUserByID(ctx, session.GetUserId())
    if err != nil {
        return nil, fmt.Errorf("user not found")
    }

    if !user.GetIsActive() {
        return nil, fmt.Errorf("user account inactive")
    }

    return user, nil
}
```

### Step 5: Update web authentication handlers

**Update `pkg/webserver/auth_handlers.go`**:

```go
package webserver

import (
    "html/template"
    "net/http"
    "time"

    "github.com/jdfalk/subtitle-manager/pkg/auth"
)

func (s *WebServer) handleLogin(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        s.renderLoginPage(w, "")
    case http.MethodPost:
        s.processLogin(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (s *WebServer) renderLoginPage(w http.ResponseWriter, errorMsg string) {
    data := struct {
        Error string
    }{
        Error: errorMsg,
    }

    tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Login - Subtitle Manager</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 400px; margin: 100px auto; }
        .form-group { margin-bottom: 15px; }
        label { display: block; margin-bottom: 5px; }
        input { width: 100%; padding: 8px; border: 1px solid #ddd; border-radius: 4px; }
        button { width: 100%; padding: 10px; background: #007cba; color: white; border: none; border-radius: 4px; }
        .error { color: red; margin-bottom: 15px; }
    </style>
</head>
<body>
    <h2>Login</h2>
    {{if .Error}}<div class="error">{{.Error}}</div>{{end}}
    <form method="post">
        <div class="form-group">
            <label for="username">Username:</label>
            <input type="text" id="username" name="username" required>
        </div>
        <div class="form-group">
            <label for="password">Password:</label>
            <input type="password" id="password" name="password" required>
        </div>
        <button type="submit">Login</button>
    </form>
    <p><a href="/register">Create new account</a></p>
</body>
</html>`

    t, _ := template.New("login").Parse(tmpl)
    t.Execute(w, data)
}

func (s *WebServer) processLogin(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")

    if username == "" || password == "" {
        s.renderLoginPage(w, "Username and password are required")
        return
    }

    // Authenticate user using auth service
    user, err := s.authService.AuthenticateUser(r.Context(), username, password)
    if err != nil {
        s.renderLoginPage(w, "Invalid username or password")
        return
    }

    // Create session
    session, err := s.authService.CreateSession(r.Context(), user.GetId())
    if err != nil {
        s.renderLoginPage(w, "Failed to create session")
        return
    }

    // Set session cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    session.GetId(),
        Path:     "/",
        Expires:  session.GetExpiresAt().AsTime(),
        HttpOnly: true,
        Secure:   s.config.UseHTTPS,
        SameSite: http.SameSiteStrictMode,
    })

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *WebServer) handleRegister(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        s.renderRegisterPage(w, "")
    case http.MethodPost:
        s.processRegister(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (s *WebServer) processRegister(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")
    confirmPassword := r.FormValue("confirm_password")

    // Validation
    if username == "" || email == "" || password == "" {
        s.renderRegisterPage(w, "All fields are required")
        return
    }

    if password != confirmPassword {
        s.renderRegisterPage(w, "Passwords do not match")
        return
    }

    if len(password) < 8 {
        s.renderRegisterPage(w, "Password must be at least 8 characters")
        return
    }

    // Create user using auth service
    user, err := s.authService.CreateUser(r.Context(), username, email, password)
    if err != nil {
        s.renderRegisterPage(w, err.Error())
        return
    }

    // Auto-login after registration
    session, err := s.authService.CreateSession(r.Context(), user.GetId())
    if err != nil {
        // Redirect to login page if session creation fails
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Set session cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    session.GetId(),
        Path:     "/",
        Expires:  session.GetExpiresAt().AsTime(),
        HttpOnly: true,
        Secure:   s.config.UseHTTPS,
        SameSite: http.SameSiteStrictMode,
    })

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *WebServer) handleLogout(w http.ResponseWriter, r *http.Request) {
    // Get session from context
    if session, ok := GetSessionFromContext(r.Context()); ok {
        s.authService.DeleteSession(r.Context(), session.GetId())
    }

    // Clear session cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    "",
        Path:     "/",
        Expires:  time.Unix(0, 0),
        HttpOnly: true,
    })

    http.Redirect(w, r, "/login", http.StatusSeeOther)
}
```

### Step 6: Update API authentication endpoints

**Create `pkg/webserver/api_auth.go`**:

```go
package webserver

import (
    "encoding/json"
    "net/http"
    "strings"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Success bool   `json:"success"`
    Token   string `json:"token,omitempty"`
    UserID  string `json:"user_id,omitempty"`
    Message string `json:"message,omitempty"`
}

type RegisterRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (s *WebServer) handleAPILogin(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        s.sendAPIError(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    user, err := s.authService.AuthenticateUser(r.Context(), req.Username, req.Password)
    if err != nil {
        s.sendAPIError(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := s.authService.GenerateAuthToken(r.Context(), user.GetId())
    if err != nil {
        s.sendAPIError(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    response := LoginResponse{
        Success: true,
        Token:   token,
        UserID:  user.GetId(),
    }

    s.sendAPIResponse(w, response, http.StatusOK)
}

func (s *WebServer) handleAPIRegister(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        s.sendAPIError(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    user, err := s.authService.CreateUser(r.Context(), req.Username, req.Email, req.Password)
    if err != nil {
        s.sendAPIError(w, err.Error(), http.StatusBadRequest)
        return
    }

    token, err := s.authService.GenerateAuthToken(r.Context(), user.GetId())
    if err != nil {
        s.sendAPIError(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    response := LoginResponse{
        Success: true,
        Token:   token,
        UserID:  user.GetId(),
    }

    s.sendAPIResponse(w, response, http.StatusCreated)
}

// API Authentication Middleware
func (s *WebServer) APIAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            s.sendAPIError(w, "Authorization header required", http.StatusUnauthorized)
            return
        }

        // Expected format: "Bearer <token>"
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            s.sendAPIError(w, "Invalid authorization format", http.StatusUnauthorized)
            return
        }

        token := parts[1]
        user, err := s.authService.ValidateAuthToken(r.Context(), token)
        if err != nil {
            s.sendAPIError(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        // Add user to request context
        ctx := context.WithValue(r.Context(), "api_user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Helper functions
func (s *WebServer) sendAPIResponse(w http.ResponseWriter, data interface{}, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

func (s *WebServer) sendAPIError(w http.ResponseWriter, message string, statusCode int) {
    response := struct {
        Success bool   `json:"success"`
        Message string `json:"message"`
    }{
        Success: false,
        Message: message,
    }
    s.sendAPIResponse(w, response, statusCode)
}

// Helper to get API user from context
func GetAPIUserFromContext(ctx context.Context) (*common.User, bool) {
    user, ok := ctx.Value("api_user").(*common.User)
    return user, ok
}
```

### Step 7: Update WebServer to use AuthService

**Update `pkg/webserver/server.go`**:

```go
type WebServer struct {
    server      *http.Server
    store       database.Store
    authService auth.AuthService  // Add auth service
    config      *Config
    templates   map[string]*template.Template
}

func NewWebServer(config *Config, store database.Store) *WebServer {
    authService := auth.NewAuthService(store)  // Create auth service

    ws := &WebServer{
        config:      config,
        store:       store,
        authService: authService,
        templates:   make(map[string]*template.Template),
    }

    ws.setupRoutes()
    return ws
}

func (s *WebServer) setupRoutes() {
    mux := http.NewServeMux()

    // Public routes
    mux.HandleFunc("/login", s.handleLogin)
    mux.HandleFunc("/register", s.handleRegister)
    mux.HandleFunc("/logout", s.handleLogout)

    // API routes
    mux.HandleFunc("/api/login", s.handleAPILogin)
    mux.HandleFunc("/api/register", s.handleAPIRegister)

    // Protected web routes
    mux.Handle("/", s.AuthMiddleware(http.HandlerFunc(s.handleHome)))
    mux.Handle("/profile", s.AuthMiddleware(http.HandlerFunc(s.handleProfile)))
    mux.Handle("/settings", s.AuthMiddleware(http.HandlerFunc(s.handleSettings)))

    // Protected API routes
    apiMux := http.NewServeMux()
    apiMux.HandleFunc("/subtitles", s.handleAPISubtitles)
    apiMux.HandleFunc("/movies", s.handleAPIMovies)
    mux.Handle("/api/", s.APIAuthMiddleware(apiMux))

    s.server = &http.Server{
        Addr:    s.config.Address,
        Handler: mux,
    }
}
```

## 🧪 Testing Requirements

### Unit Tests for Authentication Service

```go
func TestAuthService(t *testing.T) {
    store := setupTestStore(t)
    authService := auth.NewAuthService(store)
    ctx := context.Background()

    // Test user creation
    user, err := authService.CreateUser(ctx, "testuser", "test@example.com", "password123")
    assert.NoError(t, err)
    assert.Equal(t, "testuser", user.GetUsername())
    assert.Equal(t, "test@example.com", user.GetEmail())
    assert.True(t, user.GetIsActive())

    // Test authentication
    authUser, err := authService.AuthenticateUser(ctx, "testuser", "password123")
    assert.NoError(t, err)
    assert.Equal(t, user.GetId(), authUser.GetId())

    // Test session creation
    session, err := authService.CreateSession(ctx, user.GetId())
    assert.NoError(t, err)
    assert.Equal(t, user.GetId(), session.GetUserId())

    // Test session validation
    validUser, validSession, err := authService.ValidateSession(ctx, session.GetId())
    assert.NoError(t, err)
    assert.Equal(t, user.GetId(), validUser.GetId())
    assert.Equal(t, session.GetId(), validSession.GetId())
}
```

### Integration Tests for Web Authentication

```go
func TestWebAuthentication(t *testing.T) {
    server := setupTestWebServer(t)

    // Test registration
    regData := url.Values{
        "username":         {"testuser"},
        "email":           {"test@example.com"},
        "password":        {"password123"},
        "confirm_password": {"password123"},
    }

    req := httptest.NewRequest("POST", "/register", strings.NewReader(regData.Encode()))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    rr := httptest.NewRecorder()

    server.handleRegister(rr, req)
    assert.Equal(t, http.StatusSeeOther, rr.Code)

    // Extract session cookie
    cookies := rr.Result().Cookies()
    var sessionCookie *http.Cookie
    for _, cookie := range cookies {
        if cookie.Name == "session_id" {
            sessionCookie = cookie
            break
        }
    }
    assert.NotNil(t, sessionCookie)

    // Test protected endpoint with session
    req = httptest.NewRequest("GET", "/profile", nil)
    req.AddCookie(sessionCookie)
    rr = httptest.NewRecorder()

    server.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, ok := GetUserFromContext(r.Context())
        assert.True(t, ok)
        assert.Equal(t, "testuser", user.GetUsername())
        w.WriteHeader(http.StatusOK)
    })).ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
}
```

### API Authentication Tests

```go
func TestAPIAuthentication(t *testing.T) {
    server := setupTestWebServer(t)

    // Test API registration
    regReq := LoginRequest{
        Username: "apiuser",
        Email:    "api@example.com",
        Password: "password123",
    }

    reqBody, _ := json.Marshal(regReq)
    req := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    server.handleAPIRegister(rr, req)
    assert.Equal(t, http.StatusCreated, rr.Code)

    var response LoginResponse
    json.NewDecoder(rr.Body).Decode(&response)
    assert.True(t, response.Success)
    assert.NotEmpty(t, response.Token)

    // Test protected API endpoint
    req = httptest.NewRequest("GET", "/api/subtitles", nil)
    req.Header.Set("Authorization", "Bearer "+response.Token)
    rr = httptest.NewRecorder()

    server.APIAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, ok := GetAPIUserFromContext(r.Context())
        assert.True(t, ok)
        assert.Equal(t, "apiuser", user.GetUsername())
        w.WriteHeader(http.StatusOK)
    })).ServeHTTP(rr, req)

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

### Authentication Security Best Practices

1. **Password Security**: Use bcrypt for password hashing
2. **Session Security**: HttpOnly, Secure, and SameSite cookie attributes
3. **Token Security**: Proper Bearer token format validation
4. **Input Validation**: Sanitize and validate all user inputs
5. **Error Handling**: Don't leak sensitive information in error messages

## 🎯 Success Metrics

- [ ] Complete authentication flow using gcommon types
- [ ] Web and API authentication working independently
- [ ] Password hashing and verification secure
- [ ] Session and token management functional
- [ ] All authentication tests passing
- [ ] Security best practices implemented

## 🚨 Common Pitfalls

1. **Password Leakage**: Returning password hashes in API responses
2. **Timing Attacks**: Different response times for valid/invalid users
3. **Session Fixation**: Not regenerating session IDs after login
4. **Token Exposure**: Logging or exposing authentication tokens
5. **CORS Issues**: Missing CORS headers for API endpoints
6. **Context Propagation**: Not properly passing user context
7. **Error Information**: Revealing too much in authentication error messages

## 📖 Additional Resources

- [OWASP Authentication Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
- [JWT Security Best Practices](https://auth0.com/blog/a-look-at-the-latest-draft-for-jwt-bcp/)
- [Go Security Guidelines](https://github.com/OWASP/Go-SCP)

## 🔄 Related Tasks

**Must complete before this task**:

- **TASK-04-001**: Migrate User Type
- **TASK-04-002**: Migrate Session Type

**Enables these tasks**:

- **TASK-04-004**: Implement session management UI
- **TASK-05-001**: Performance optimization for authentication
- **TASK-06-001**: Integration testing with full auth flow

## 📝 Notes for AI Agent

- Always use opaque API for User and Session fields
- Implement both web and API authentication patterns
- Use secure password hashing with bcrypt
- Implement proper session and token management
- Add comprehensive input validation
- Follow security best practices for authentication
- Test both successful and failure authentication scenarios
- Ensure proper context propagation in middleware

## 🔚 Completion Verification

```bash
# Verify authentication service compilation
echo "Checking authentication service..."
go build ./pkg/auth && echo "✅ Auth service compiles"

# Verify opaque API usage
echo "Checking for opaque API usage..."
grep -r "GetUsername()\|SetUsername(\|GetPasswordHash()" --include="*.go" ./pkg/auth && echo "✅ Auth service uses opaque API"

# Run authentication tests
echo "Running authentication tests..."
go test -run TestAuth ./... && echo "✅ Authentication tests pass"

# Test web server compilation
echo "Testing web server compilation..."
go build ./pkg/webserver && echo "✅ Web server compiles"
```
