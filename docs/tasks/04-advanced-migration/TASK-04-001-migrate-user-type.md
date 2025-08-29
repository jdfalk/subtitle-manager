# TASK-04-001: Migrate User Type from gcommonauth to gcommon/common

<!-- file: docs/tasks/04-advanced-migration/TASK-04-001-migrate-user-type.md -->
<!-- version: 1.0.0 -->
<!-- guid: c1d2e3f4-5678-9012-3456-789012345678 -->

## 🎯 Task Overview

**Primary Objective**: Replace all local User types with gcommon User type using opaque API pattern

**Task Type**: Core Type Migration - User Entity

**Estimated Effort**: 3-4 hours

**Dependencies**: 
- TASK-01-003 (Replace gcommonauth package)
- gcommon v1.8.0 properly configured

## 📋 Acceptance Criteria

- [ ] All imports updated from `pkg/gcommonauth` to `gcommon/sdks/go/v1/common`
- [ ] All User field access converted to opaque API (Set*/Get* methods)
- [ ] Return types corrected to use `[]*common.User` instead of `[]common.User`
- [ ] All authentication methods implemented in SQLStore and PostgresStore
- [ ] Tests pass with new User type implementation
- [ ] No direct field access remains in codebase
- [ ] Type safety verified through compilation

## 🔄 Dependencies

**Input Requirements**:
- gcommon v1.8.0 available and configured
- Understanding of opaque API pattern for protobuf access

**External Dependencies**:
- gcommon/sdks/go/v1/common package
- Database interface implementations
- Authentication system components

## 📝 Implementation Steps

### Step 1: Update imports across all user-related files

**Files to Modify**:
- `pkg/database/store.go`
- `pkg/database/pebble.go` 
- `pkg/database/database.go`
- `pkg/database/postgres.go`
- `pkg/webserver/auth.go`
- `pkg/webserver/users.go`
- `pkg/authserver/server.go`
- All test files using User type

**Import Replacement**:
```go
// OLD - Remove these imports
import auth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
import "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"

// NEW - Add this import
import "github.com/jdfalk/gcommon/sdks/go/v1/common"
```

### Step 2: Convert User field access to opaque API

**Critical Pattern Changes**:

```go
// OLD - Direct field access (REMOVE)
user := auth.User{
    Id:       "user123",
    Username: "john_doe",
    Email:    "john@example.com",
    IsAdmin:  true,
}
userID := user.Id
username := user.Username

// NEW - Opaque API pattern (IMPLEMENT)
user := &common.User{}
user.SetId("user123")
user.SetUsername("john_doe") 
user.SetEmail("john@example.com")
user.SetIsAdmin(true)

userID := user.GetId()
username := user.GetUsername()
```

**Common Field Mappings**:
- `user.Id` → `user.GetId()` / `user.SetId("value")`
- `user.Username` → `user.GetUsername()` / `user.SetUsername("value")`
- `user.Email` → `user.GetEmail()` / `user.SetEmail("value")`
- `user.PasswordHash` → `user.GetPasswordHash()` / `user.SetPasswordHash("value")`
- `user.IsAdmin` → `user.GetIsAdmin()` / `user.SetIsAdmin(true)`
- `user.CreatedAt` → `user.GetCreatedAt()` / `user.SetCreatedAt(time.Now())`
- `user.UpdatedAt` → `user.GetUpdatedAt()` / `user.SetUpdatedAt(time.Now())`

### Step 3: Fix return type mismatches

**Critical Type Updates**:

```go
// OLD - Slice of values (INCORRECT)
func GetAllUsers() []common.User { ... }
func ListUsers() []auth.User { ... }

// NEW - Slice of pointers (CORRECT)
func GetAllUsers() []*common.User { ... }
func ListUsers() []*common.User { ... }
```

**Method Signature Updates**:
```go
// OLD
func CreateUser(user auth.User) error
func UpdateUser(user auth.User) error
func GetUserByID(id string) (auth.User, error)

// NEW  
func CreateUser(user *common.User) error
func UpdateUser(user *common.User) error
func GetUserByID(id string) (*common.User, error)
```

### Step 4: Implement missing authentication methods

**Required Method Implementations for SQLStore and PostgresStore**:

```go
// 1. CreateOneTimeToken
func (s *SQLStore) CreateOneTimeToken(ctx context.Context, userID string, token string, expiresAt time.Time) error {
    _, err := s.db.ExecContext(ctx, 
        "INSERT INTO one_time_tokens (user_id, token, expires_at) VALUES (?, ?, ?)",
        userID, token, expiresAt)
    return err
}

// 2. CleanupExpiredSessions 
func (s *SQLStore) CleanupExpiredSessions(ctx context.Context) error {
    _, err := s.db.ExecContext(ctx,
        "DELETE FROM sessions WHERE expires_at < ?", time.Now())
    return err
}

// 3. GetUserByID
func (s *SQLStore) GetUserByID(ctx context.Context, id string) (*common.User, error) {
    row := s.db.QueryRowContext(ctx,
        "SELECT id, username, email, password_hash, is_admin, created_at, updated_at FROM users WHERE id = ?", id)
    
    user := &common.User{}
    var createdAt, updatedAt time.Time
    var id, username, email, passwordHash string
    var isAdmin bool
    
    err := row.Scan(&id, &username, &email, &passwordHash, &isAdmin, &createdAt, &updatedAt)
    if err != nil {
        return nil, err
    }
    
    // Use opaque API to populate user
    user.SetId(id)
    user.SetUsername(username)
    user.SetEmail(email)
    user.SetPasswordHash(passwordHash)
    user.SetIsAdmin(isAdmin)
    user.SetCreatedAt(timestamppb.New(createdAt))
    user.SetUpdatedAt(timestamppb.New(updatedAt))
    
    return user, nil
}

// 4. GetUserByEmail
func (s *SQLStore) GetUserByEmail(ctx context.Context, email string) (*common.User, error) {
    row := s.db.QueryRowContext(ctx,
        "SELECT id, username, email, password_hash, is_admin, created_at, updated_at FROM users WHERE email = ?", email)
    
    // Similar implementation to GetUserByID...
    return s.scanUser(row)
}

// 5. GetUserByUsername  
func (s *SQLStore) GetUserByUsername(ctx context.Context, username string) (*common.User, error) {
    row := s.db.QueryRowContext(ctx,
        "SELECT id, username, email, password_hash, is_admin, created_at, updated_at FROM users WHERE username = ?", username)
    
    return s.scanUser(row)
}

// 6. ListUsers
func (s *SQLStore) ListUsers(ctx context.Context) ([]*common.User, error) {
    rows, err := s.db.QueryContext(ctx,
        "SELECT id, username, email, password_hash, is_admin, created_at, updated_at FROM users ORDER BY username")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []*common.User
    for rows.Next() {
        user, err := s.scanUser(rows)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    return users, rows.Err()
}
```

**Helper Method for User Scanning**:
```go
func (s *SQLStore) scanUser(scanner interface{}) (*common.User, error) {
    user := &common.User{}
    var createdAt, updatedAt time.Time
    var id, username, email, passwordHash string
    var isAdmin bool
    
    var err error
    switch s := scanner.(type) {
    case *sql.Row:
        err = s.Scan(&id, &username, &email, &passwordHash, &isAdmin, &createdAt, &updatedAt)
    case *sql.Rows:
        err = s.Scan(&id, &username, &email, &passwordHash, &isAdmin, &createdAt, &updatedAt)
    default:
        return nil, fmt.Errorf("unsupported scanner type")
    }
    
    if err != nil {
        return nil, err
    }
    
    // Use opaque API
    user.SetId(id)
    user.SetUsername(username)
    user.SetEmail(email)
    user.SetPasswordHash(passwordHash)
    user.SetIsAdmin(isAdmin)
    user.SetCreatedAt(timestamppb.New(createdAt))
    user.SetUpdatedAt(timestamppb.New(updatedAt))
    
    return user, nil
}
```

### Step 5: Update PostgreSQL implementations

**PostgresStore Method Implementations**:

```go
// PostgreSQL versions use $1, $2, etc. instead of ?
func (p *PostgresStore) GetUserByID(ctx context.Context, id string) (*common.User, error) {
    row := p.db.QueryRowContext(ctx,
        "SELECT id, username, email, password_hash, is_admin, created_at, updated_at FROM users WHERE id = $1", id)
    
    return p.scanUser(row)
}

func (p *PostgresStore) CreateOneTimeToken(ctx context.Context, userID string, token string, expiresAt time.Time) error {
    _, err := p.db.ExecContext(ctx,
        "INSERT INTO one_time_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)",
        userID, token, expiresAt)
    return err
}

// Continue pattern for all other methods...
```

### Step 6: Update authentication server components

**Auth Server Updates in `pkg/authserver/server.go`**:

```go
// Update method signatures
func (s *AuthServer) CreateUser(ctx context.Context, req *common.CreateUserRequest) (*common.User, error) {
    user := &common.User{}
    user.SetId(generateUserID())
    user.SetUsername(req.GetUsername())
    user.SetEmail(req.GetEmail())
    user.SetPasswordHash(hashPassword(req.GetPassword()))
    user.SetIsAdmin(req.GetIsAdmin())
    user.SetCreatedAt(timestamppb.Now())
    user.SetUpdatedAt(timestamppb.Now())
    
    err := s.store.CreateUser(ctx, user)
    return user, err
}

func (s *AuthServer) GetUser(ctx context.Context, req *common.GetUserRequest) (*common.User, error) {
    return s.store.GetUserByID(ctx, req.GetId())
}
```

### Step 7: Update web server authentication

**Web Server Updates in `pkg/webserver/auth.go` and `pkg/webserver/users.go`**:

```go
// Update user handling in HTTP handlers
func (s *WebServer) handleLogin(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    
    user, err := s.store.GetUserByUsername(r.Context(), username)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    
    // Use opaque API
    if !verifyPassword(password, user.GetPasswordHash()) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    
    // Create session using user.GetId()
    sessionID := generateSessionID()
    err = s.store.CreateSession(r.Context(), sessionID, user.GetId(), time.Now().Add(24*time.Hour))
    if err != nil {
        http.Error(w, "Failed to create session", http.StatusInternalServerError)
        return
    }
    
    // Set session cookie and redirect
    http.SetCookie(w, &http.Cookie{
        Name:  "session_id",
        Value: sessionID,
        Path:  "/",
    })
    
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
```

## 🧪 Testing Requirements

### Unit Tests

Create comprehensive tests for User type migration:

```go
func TestUserTypeMigration(t *testing.T) {
    // Test opaque API usage
    user := &common.User{}
    user.SetId("test123")
    user.SetUsername("testuser")
    user.SetEmail("test@example.com")
    
    assert.Equal(t, "test123", user.GetId())
    assert.Equal(t, "testuser", user.GetUsername())
    assert.Equal(t, "test@example.com", user.GetEmail())
}

func TestUserDatabaseOperations(t *testing.T) {
    store := setupTestStore(t)
    
    // Test user creation
    user := &common.User{}
    user.SetId("user123")
    user.SetUsername("john")
    user.SetEmail("john@example.com")
    
    err := store.CreateUser(context.Background(), user)
    assert.NoError(t, err)
    
    // Test user retrieval
    retrieved, err := store.GetUserByID(context.Background(), "user123")
    assert.NoError(t, err)
    assert.Equal(t, "john", retrieved.GetUsername())
    assert.Equal(t, "john@example.com", retrieved.GetEmail())
}
```

### Integration Tests

```go
func TestAuthenticationFlow(t *testing.T) {
    // Test complete authentication with new User type
    server := setupTestServer(t)
    
    // Create user
    user := &common.User{}
    user.SetUsername("integration_test")
    user.SetEmail("test@integration.com")
    user.SetPasswordHash(hashPassword("password123"))
    
    err := server.store.CreateUser(context.Background(), user)
    assert.NoError(t, err)
    
    // Test login flow
    // Test session management
    // Test user operations
}
```

## 📚 Required Documentation

**Embedded from .github/instructions/general-coding.instructions.md:**

### Critical Guidelines

```markdown
## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS

**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction of any kind.**

## Required File Header (File Identification)

All source, script, and documentation files MUST begin with a standard header containing:
- The exact relative file path from the repository root
- The file's semantic version 
- The file's GUID

## Version Update Requirements

**When modifying any file with a version header, ALWAYS update the version number**
```

### Opaque API Pattern Requirements

**CRITICAL**: gcommon v1.8.0 uses opaque API pattern for all protobuf types:

```go
// ❌ NEVER do this - Direct field access
user.Id = "123"
username := user.Username

// ✅ ALWAYS do this - Use setter/getter methods
user.SetId("123") 
username := user.GetUsername()
```

## 🎯 Success Metrics

- [ ] Zero compilation errors after migration
- [ ] All User field access uses opaque API (verified by grep)
- [ ] All authentication methods implemented and tested
- [ ] Database operations work with new User type
- [ ] Web authentication flow functional
- [ ] All tests pass with new implementation
- [ ] No direct protobuf field access remains

## 🚨 Common Pitfalls

1. **Direct Field Access**: Using `user.Id` instead of `user.GetId()`
2. **Wrong Return Types**: Returning `[]User` instead of `[]*User`
3. **Nil Pointer Issues**: Not initializing with `&common.User{}`
4. **Time Conversion**: Forgetting `timestamppb.New()` for time fields
5. **SQL Parameter Style**: Using `?` in PostgreSQL instead of `$1`
6. **Context Handling**: Not passing context through all database calls
7. **Error Wrapping**: Not preserving error context in database operations

## 📖 Additional Resources

- [gcommon User API Documentation](../../gcommon-api/common.md#user)
- [Opaque API Pattern Guide](../../gcommon-api/README.md#opaque-api)
- [Database Interface Requirements](../../gcommon-api/database.md)
- [Authentication Flow Documentation](../../gcommon-api/common.md#sessions)

## 🔄 Related Tasks

**Must complete before this task**:
- **TASK-01-003**: Replace gcommonauth package (provides foundation)

**Enables these tasks**:
- **TASK-04-002**: Migrate Session type
- **TASK-04-003**: Update authentication middleware
- **TASK-06-001**: Integration testing with new User type

## 📝 Notes for AI Agent

- Focus on one file at a time to avoid overwhelming changes
- Always use opaque API (`Get*()`, `Set*()`) - never direct field access
- Test each database backend separately (SQLite, PostgreSQL)
- Verify return types are `[]*common.User`, not `[]common.User`
- Use context for all database operations
- Pay attention to SQL parameter styles (`?` vs `$1`, `$2`)
- Include proper error handling and validation
- Update both method signatures and implementations
- Don't forget to update test files that use User type

## 🔚 Completion Verification

```bash
# Verify no direct field access remains
echo "Checking for direct field access..."
grep -r "\.Id\|\.Username\|\.Email" --include="*.go" . || echo "✅ No direct field access found"

# Verify opaque API usage
echo "Checking for opaque API usage..."
grep -r "GetId()\|SetId(\|GetUsername()" --include="*.go" . && echo "✅ Opaque API in use"

# Test compilation
echo "Testing compilation..."
go build ./... && echo "✅ Compilation successful"

# Run tests
echo "Running tests..."
go test ./... && echo "✅ All tests pass"
```
