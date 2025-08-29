# TASK-04-002: Migrate Session Type and Authentication Middleware

<!-- file: docs/tasks/04-advanced-migration/TASK-04-002-migrate-session-type.md -->
<!-- version: 1.0.0 -->
<!-- guid: d2e3f4g5-6789-0123-4567-890123456789 -->

## 🎯 Task Overview

**Primary Objective**: Replace local Session types with gcommon Session type and update authentication middleware

**Task Type**: Core Type Migration - Session Management

**Estimated Effort**: 2-3 hours

**Dependencies**: 
- TASK-04-001 (Migrate User Type) completed
- Authentication middleware understanding

## 📋 Acceptance Criteria

- [ ] All Session imports updated to gcommon/v1/common
- [ ] Session field access converted to opaque API
- [ ] Authentication middleware updated for new Session type
- [ ] Session management in web handlers updated
- [ ] Database session operations work with new type
- [ ] Session creation, validation, and cleanup functional
- [ ] All session-related tests pass

## 🔄 Dependencies

**Input Requirements**:
- TASK-04-001 completed (User type migrated)
- gcommon Session type understanding

**External Dependencies**:
- gcommon/sdks/go/v1/common package
- Updated authentication middleware
- Session storage backends

## 📝 Implementation Steps

### Step 1: Update Session type imports

**Files to Modify**:
- `pkg/webserver/auth.go`
- `pkg/webserver/middleware.go`
- `pkg/database/store.go`
- `pkg/database/pebble.go`
- `pkg/database/database.go`
- `pkg/database/postgres.go`
- `pkg/authserver/server.go`

**Import Updates**:
```go
// OLD - Remove local session imports
import "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"

// NEW - Use gcommon session
import "github.com/jdfalk/gcommon/sdks/go/v1/common"
```

### Step 2: Convert Session field access to opaque API

**Critical Session Field Mappings**:
```go
// OLD - Direct field access (REMOVE)
session := auth.Session{
    Id:        "sess123",
    UserId:    "user456", 
    ExpiresAt: time.Now().Add(24 * time.Hour),
    CreatedAt: time.Now(),
}
sessionID := session.Id
userID := session.UserId

// NEW - Opaque API pattern (IMPLEMENT)
session := &common.Session{}
session.SetId("sess123")
session.SetUserId("user456")
session.SetExpiresAt(timestamppb.New(time.Now().Add(24 * time.Hour)))
session.SetCreatedAt(timestamppb.Now())

sessionID := session.GetId()
userID := session.GetUserId()
```

**All Session Field Conversions**:
- `session.Id` → `session.GetId()` / `session.SetId("value")`
- `session.UserId` → `session.GetUserId()` / `session.SetUserId("value")`
- `session.ExpiresAt` → `session.GetExpiresAt()` / `session.SetExpiresAt(timestamppb.New(time))`
- `session.CreatedAt` → `session.GetCreatedAt()` / `session.SetCreatedAt(timestamppb.New(time))`

### Step 3: Update authentication middleware

**Update `pkg/webserver/middleware.go`**:

```go
func (s *WebServer) AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("session_id")
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Get session using new Session type
        session, err := s.store.GetSession(r.Context(), cookie.Value)
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Use opaque API to check expiration
        if session.GetExpiresAt().AsTime().Before(time.Now()) {
            // Clean up expired session
            s.store.DeleteSession(r.Context(), session.GetId())
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Get user associated with session
        user, err := s.store.GetUserByID(r.Context(), session.GetUserId())
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Add user and session to request context
        ctx := context.WithValue(r.Context(), "user", user)
        ctx = context.WithValue(ctx, "session", session)
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Helper function to get user from context
func GetUserFromContext(ctx context.Context) (*common.User, bool) {
    user, ok := ctx.Value("user").(*common.User)
    return user, ok
}

// Helper function to get session from context  
func GetSessionFromContext(ctx context.Context) (*common.Session, bool) {
    session, ok := ctx.Value("session").(*common.Session)
    return session, ok
}
```

### Step 4: Update session database operations

**SQLStore Session Methods**:

```go
func (s *SQLStore) CreateSession(ctx context.Context, sessionID, userID string, expiresAt time.Time) error {
    _, err := s.db.ExecContext(ctx,
        "INSERT INTO sessions (id, user_id, expires_at, created_at) VALUES (?, ?, ?, ?)",
        sessionID, userID, expiresAt, time.Now())
    return err
}

func (s *SQLStore) GetSession(ctx context.Context, sessionID string) (*common.Session, error) {
    row := s.db.QueryRowContext(ctx,
        "SELECT id, user_id, expires_at, created_at FROM sessions WHERE id = ?", sessionID)
    
    session := &common.Session{}
    var id, userID string
    var expiresAt, createdAt time.Time
    
    err := row.Scan(&id, &userID, &expiresAt, &createdAt)
    if err != nil {
        return nil, err
    }
    
    // Use opaque API to populate session
    session.SetId(id)
    session.SetUserId(userID)
    session.SetExpiresAt(timestamppb.New(expiresAt))
    session.SetCreatedAt(timestamppb.New(createdAt))
    
    return session, nil
}

func (s *SQLStore) DeleteSession(ctx context.Context, sessionID string) error {
    _, err := s.db.ExecContext(ctx,
        "DELETE FROM sessions WHERE id = ?", sessionID)
    return err
}

func (s *SQLStore) GetSessionsByUserID(ctx context.Context, userID string) ([]*common.Session, error) {
    rows, err := s.db.QueryContext(ctx,
        "SELECT id, user_id, expires_at, created_at FROM sessions WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var sessions []*common.Session
    for rows.Next() {
        session := &common.Session{}
        var id, userID string
        var expiresAt, createdAt time.Time
        
        err := rows.Scan(&id, &userID, &expiresAt, &createdAt)
        if err != nil {
            return nil, err
        }
        
        session.SetId(id)
        session.SetUserId(userID)
        session.SetExpiresAt(timestamppb.New(expiresAt))
        session.SetCreatedAt(timestamppb.New(createdAt))
        
        sessions = append(sessions, session)
    }
    
    return sessions, rows.Err()
}

func (s *SQLStore) CleanupExpiredSessions(ctx context.Context) error {
    _, err := s.db.ExecContext(ctx,
        "DELETE FROM sessions WHERE expires_at < ?", time.Now())
    return err
}
```

**PostgresStore Session Methods** (similar with $1, $2 parameters):

```go
func (p *PostgresStore) CreateSession(ctx context.Context, sessionID, userID string, expiresAt time.Time) error {
    _, err := p.db.ExecContext(ctx,
        "INSERT INTO sessions (id, user_id, expires_at, created_at) VALUES ($1, $2, $3, $4)",
        sessionID, userID, expiresAt, time.Now())
    return err
}

func (p *PostgresStore) GetSession(ctx context.Context, sessionID string) (*common.Session, error) {
    row := p.db.QueryRowContext(ctx,
        "SELECT id, user_id, expires_at, created_at FROM sessions WHERE id = $1", sessionID)
    
    return p.scanSession(row)
}

// Helper method for session scanning
func (p *PostgresStore) scanSession(scanner interface{}) (*common.Session, error) {
    session := &common.Session{}
    var id, userID string
    var expiresAt, createdAt time.Time
    
    var err error
    switch s := scanner.(type) {
    case *sql.Row:
        err = s.Scan(&id, &userID, &expiresAt, &createdAt)
    case *sql.Rows:
        err = s.Scan(&id, &userID, &expiresAt, &createdAt)
    default:
        return nil, fmt.Errorf("unsupported scanner type")
    }
    
    if err != nil {
        return nil, err
    }
    
    session.SetId(id)
    session.SetUserId(userID)
    session.SetExpiresAt(timestamppb.New(expiresAt))
    session.SetCreatedAt(timestamppb.New(createdAt))
    
    return session, nil
}
```

### Step 5: Update web handlers for session management

**Login Handler Update**:

```go
func (s *WebServer) handleLogin(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")

    // Get user (already updated in TASK-04-001)
    user, err := s.store.GetUserByUsername(r.Context(), username)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Verify password
    if !verifyPassword(password, user.GetPasswordHash()) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Create new session
    sessionID := generateSessionID()
    expiresAt := time.Now().Add(24 * time.Hour)
    
    err = s.store.CreateSession(r.Context(), sessionID, user.GetId(), expiresAt)
    if err != nil {
        http.Error(w, "Failed to create session", http.StatusInternalServerError)
        return
    }

    // Set session cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    sessionID,
        Path:     "/",
        Expires:  expiresAt,
        HttpOnly: true,
        Secure:   s.config.UseHTTPS,
        SameSite: http.SameSiteStrictMode,
    })

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
```

**Logout Handler Update**:

```go
func (s *WebServer) handleLogout(w http.ResponseWriter, r *http.Request) {
    // Get session from context (set by middleware)
    session, ok := GetSessionFromContext(r.Context())
    if ok {
        // Delete session from database
        s.store.DeleteSession(r.Context(), session.GetId())
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

### Step 6: Update user profile and session management pages

**User Profile Handler**:

```go
func (s *WebServer) handleUserProfile(w http.ResponseWriter, r *http.Request) {
    user, ok := GetUserFromContext(r.Context())
    if !ok {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Get user's active sessions
    sessions, err := s.store.GetSessionsByUserID(r.Context(), user.GetId())
    if err != nil {
        log.Printf("Failed to get user sessions: %v", err)
        sessions = []*common.Session{} // Empty slice on error
    }

    data := struct {
        User     *common.User
        Sessions []*common.Session
    }{
        User:     user,
        Sessions: sessions,
    }

    s.renderTemplate(w, "profile.html", data)
}
```

## 🧪 Testing Requirements

### Unit Tests for Session Migration

```go
func TestSessionTypeMigration(t *testing.T) {
    // Test opaque API usage
    session := &common.Session{}
    session.SetId("sess123")
    session.SetUserId("user456")
    session.SetExpiresAt(timestamppb.New(time.Now().Add(time.Hour)))
    session.SetCreatedAt(timestamppb.Now())
    
    assert.Equal(t, "sess123", session.GetId())
    assert.Equal(t, "user456", session.GetUserId())
    assert.True(t, session.GetExpiresAt().AsTime().After(time.Now()))
}

func TestSessionDatabaseOperations(t *testing.T) {
    store := setupTestStore(t)
    ctx := context.Background()
    
    // Create session
    sessionID := "test_session_123"
    userID := "test_user_456"
    expiresAt := time.Now().Add(time.Hour)
    
    err := store.CreateSession(ctx, sessionID, userID, expiresAt)
    assert.NoError(t, err)
    
    // Retrieve session
    session, err := store.GetSession(ctx, sessionID)
    assert.NoError(t, err)
    assert.Equal(t, sessionID, session.GetId())
    assert.Equal(t, userID, session.GetUserId())
    
    // Delete session
    err = store.DeleteSession(ctx, sessionID)
    assert.NoError(t, err)
    
    // Verify deletion
    _, err = store.GetSession(ctx, sessionID)
    assert.Error(t, err)
}
```

### Integration Tests for Authentication Flow

```go
func TestAuthenticationMiddleware(t *testing.T) {
    server := setupTestServer(t)
    
    // Create test user and session
    user := createTestUser(t, server.store)
    sessionID := "test_session"
    expiresAt := time.Now().Add(time.Hour)
    
    err := server.store.CreateSession(context.Background(), sessionID, user.GetId(), expiresAt)
    assert.NoError(t, err)
    
    // Test protected endpoint with valid session
    req := httptest.NewRequest("GET", "/protected", nil)
    req.AddCookie(&http.Cookie{Name: "session_id", Value: sessionID})
    
    rr := httptest.NewRecorder()
    server.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        contextUser, ok := GetUserFromContext(r.Context())
        assert.True(t, ok)
        assert.Equal(t, user.GetId(), contextUser.GetId())
        w.WriteHeader(http.StatusOK)
    })).ServeHTTP(rr, req)
    
    assert.Equal(t, http.StatusOK, rr.Code)
}

func TestExpiredSessionCleanup(t *testing.T) {
    store := setupTestStore(t)
    ctx := context.Background()
    
    // Create expired session
    sessionID := "expired_session"
    userID := "test_user"
    expiresAt := time.Now().Add(-time.Hour) // Already expired
    
    err := store.CreateSession(ctx, sessionID, userID, expiresAt)
    assert.NoError(t, err)
    
    // Run cleanup
    err = store.CleanupExpiredSessions(ctx)
    assert.NoError(t, err)
    
    // Verify expired session is gone
    _, err = store.GetSession(ctx, sessionID)
    assert.Error(t, err)
}
```

## 📚 Required Documentation

**Embedded from .github/instructions/general-coding.instructions.md:**

### Critical Guidelines

```markdown
## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS

**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction of any kind.**

## Required File Header (File Identification)

All source, script, and documentation files MUST begin with a standard header.

## Version Update Requirements

**When modifying any file with a version header, ALWAYS update the version number**
```

### Session Security Best Practices

1. **HttpOnly Cookies**: Always set `HttpOnly: true` for session cookies
2. **Secure Transmission**: Use `Secure: true` in HTTPS environments
3. **SameSite Protection**: Set appropriate `SameSite` policy
4. **Expiration Management**: Implement proper session timeout
5. **Cleanup**: Regular cleanup of expired sessions

## 🎯 Success Metrics

- [ ] All Session field access uses opaque API
- [ ] Authentication middleware works with new Session type
- [ ] Session creation, validation, and cleanup functional
- [ ] Database operations handle new Session type correctly
- [ ] Web handlers properly manage sessions
- [ ] All session-related tests pass
- [ ] Session security measures implemented

## 🚨 Common Pitfalls

1. **Time Conversion**: Forgetting `timestamppb.New()` for time fields
2. **Nil Session Handling**: Not checking for nil sessions in middleware
3. **Context Propagation**: Not passing session in request context
4. **Cookie Security**: Missing security flags on session cookies
5. **Expired Session Handling**: Not cleaning up expired sessions properly
6. **SQL Parameter Style**: Mixing `?` and `$1` parameter styles
7. **Type Assertions**: Unsafe type assertions from context

## 📖 Additional Resources

- [gcommon Session API Documentation](../../gcommon-api/common.md#session)
- [HTTP Cookie Security Best Practices](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies)
- [Session Management Security Guide](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html)

## 🔄 Related Tasks

**Must complete before this task**:
- **TASK-04-001**: Migrate User Type (provides updated User handling)

**Enables these tasks**:
- **TASK-04-003**: Update authentication flow end-to-end
- **TASK-04-004**: Implement session management UI
- **TASK-06-001**: Integration testing with authentication

## 📝 Notes for AI Agent

- Always use opaque API for Session fields
- Handle time conversions properly with timestamppb
- Implement proper session security measures
- Test both SQLite and PostgreSQL backends
- Ensure middleware properly propagates user/session context
- Clean up expired sessions regularly
- Use proper HTTP cookie security flags
- Test authentication flow end-to-end after changes

## 🔚 Completion Verification

```bash
# Verify Session opaque API usage
echo "Checking for Session opaque API usage..."
grep -r "GetId()\|SetId(\|GetUserId()\|SetUserId(" --include="*.go" . && echo "✅ Session opaque API in use"

# Verify no direct field access
echo "Checking for direct field access..."
grep -r "\.Id\|\.UserId\|\.ExpiresAt" --include="*.go" . || echo "✅ No direct field access found"

# Test compilation
echo "Testing compilation..."
go build ./... && echo "✅ Compilation successful"

# Run authentication tests
echo "Running authentication tests..."
go test -run TestAuth ./... && echo "✅ Authentication tests pass"
```
