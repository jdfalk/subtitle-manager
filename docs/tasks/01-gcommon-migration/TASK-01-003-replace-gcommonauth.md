# TASK-01-003: Replace gcommonauth with gcommon common

<!-- file: docs/tasks/01-gcommon-migration/TASK-01-003-replace-gcommonauth.md -->
<!-- version: 1.0.0 -->
<!-- guid: d3e4f5g6-h7i8-9012-cdef-345678901234 -->

## 🎯 Objective

Replace all usage of the local `pkg/gcommonauth` package with the gcommon common package authentication types (`github.com/jdfalk/gcommon/sdks/go/v1/common`).

## 📋 Acceptance Criteria

- [ ] All imports of `github.com/jdfalk/subtitle-manager/pkg/gcommonauth` are replaced
- [ ] All authentication-related types use gcommon protobuf types with opaque API
- [ ] All getter/setter methods use the correct opaque API pattern
- [ ] No compilation errors after migration
- [ ] All tests pass with new authentication types
- [ ] Local `pkg/gcommonauth` directory can be safely removed

## 🔍 Current State Analysis

### Files Currently Using gcommonauth

Based on codebase analysis, these files import `gcommonauth`:

1. `pkg/maintenance/maintenance.go`
2. `pkg/maintenance/maintenance_test.go`
3. `pkg/webserver/system_test.go`
4. `pkg/webserver/server_test.go`
5. `pkg/webserver/oauth_management_test.go`
6. `pkg/webserver/oauth.go`
7. `pkg/webserver/widgets_test.go`
8. `pkg/webserver/database_test.go`
9. `pkg/webserver/oauth_test.go`
10. `pkg/webserver/auth.go`
11. `pkg/webserver/server.go`
12. `pkg/webserver/users.go`
13. `pkg/authserver/server_test.go`
14. `pkg/authserver/server.go`
15. `cmd/login_test.go`

### Current gcommonauth Types to Replace

From `pkg/gcommonauth/`:

- `User` → `github.com/jdfalk/gcommon/sdks/go/v1/common.User`
- `Session` → `github.com/jdfalk/gcommon/sdks/go/v1/common.Session`
- `APIKey` → `github.com/jdfalk/gcommon/sdks/go/v1/common.APIKey`
- `Role` → `github.com/jdfalk/gcommon/sdks/go/v1/common.Role`
- `Permission` → `github.com/jdfalk/gcommon/sdks/go/v1/common.Permission`

## 🔧 Implementation Steps

### Step 1: Analyze gcommon common package structure

```bash
# Generate documentation for gcommon common package
gomarkdoc --output docs/gcommon-api/common.md github.com/jdfalk/gcommon/sdks/go/v1/common
```

### Step 2: Create migration mapping

Create a mapping file `docs/tasks/01-gcommon-migration/auth-migration-map.md`:

```markdown
| Local Type | gcommon Type | Import Path |
|------------|--------------|-------------|
| gcommonauth.User | common.User | github.com/jdfalk/gcommon/sdks/go/v1/common |
| gcommonauth.Session | common.Session | github.com/jdfalk/gcommon/sdks/go/v1/common |
| gcommonauth.APIKey | common.APIKey | github.com/jdfalk/gcommon/sdks/go/v1/common |
| gcommonauth.Role | common.Role | github.com/jdfalk/gcommon/sdks/go/v1/common |
```

### Step 3: Update import statements

For each file in the list:

```go
// OLD
import (
    auth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
)

// NEW
import (
    "github.com/jdfalk/gcommon/sdks/go/v1/common"
)
```

### Step 4: Update type usage with opaque API

Replace direct field access with setter/getter methods:

```go
// OLD - Direct field access
user := &auth.User{
    ID:       "123",
    Username: "john",
    Email:    "john@example.com",
    Role:     auth.RoleAdmin,
}

// NEW - Opaque API with setters
user := &common.User{}
user.SetId("123")
user.SetUsername("john")
user.SetEmail("john@example.com")
user.SetRole(common.Role_ROLE_ADMIN)
```

### Step 5: Update authentication middleware

Update `pkg/webserver/auth.go`:

```go
// OLD
func (s *Server) authenticateUser(token string) (*auth.User, error) {
    user, err := s.authService.ValidateToken(token)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// NEW
func (s *Server) authenticateUser(token string) (*common.User, error) {
    user, err := s.authService.ValidateToken(token)
    if err != nil {
        return nil, err
    }
    return user, nil
}
```

### Step 6: Update session management

Update session handling to use opaque API:

```go
// OLD
session := &auth.Session{
    ID:        sessionID,
    UserID:    user.ID,
    ExpiresAt: time.Now().Add(24 * time.Hour),
}

// NEW
session := &common.Session{}
session.SetId(sessionID)
session.SetUserId(user.GetId())
session.SetExpiresAt(timestamppb.New(time.Now().Add(24 * time.Hour)))
```

### Step 7: Update OAuth handlers

Update `pkg/webserver/oauth.go`:

```go
// OLD
func (s *Server) handleOAuthCallback(user *auth.User) error {
    if user.Role != auth.RoleAdmin {
        return errors.New("insufficient permissions")
    }
    return nil
}

// NEW
func (s *Server) handleOAuthCallback(user *common.User) error {
    if user.GetRole() != common.Role_ROLE_ADMIN {
        return errors.New("insufficient permissions")
    }
    return nil
}
```

### Step 8: Update tests

Update all test files to use new types:

```go
// OLD
func TestUserAuthentication(t *testing.T) {
    user := &auth.User{
        ID:       "test-user",
        Username: "testuser",
        Role:     auth.RoleUser,
    }
    assert.Equal(t, "test-user", user.ID)
    assert.Equal(t, auth.RoleUser, user.Role)
}

// NEW
func TestUserAuthentication(t *testing.T) {
    user := &common.User{}
    user.SetId("test-user")
    user.SetUsername("testuser")
    user.SetRole(common.Role_ROLE_USER)
    assert.Equal(t, "test-user", user.GetId())
    assert.Equal(t, common.Role_ROLE_USER, user.GetRole())
}
```

### Step 9: Update database layer

Ensure database layer works with new types:

```go
// Update pkg/database/store.go methods
func (s *Store) GetUserByID(id string) (*common.User, error) {
    // Use getters throughout implementation
}

func (s *Store) CreateUser(user *common.User) error {
    // Use getters to extract data for storage
    username := user.GetUsername()
    email := user.GetEmail()
    // ... implementation
}
```

### Step 10: Build and test

```bash
# Build to check for compilation errors
go build ./...

# Run authentication tests
go test ./pkg/webserver/... -v
go test ./pkg/authserver/... -v

# Run full test suite
go test ./...
```

## 📚 Required Documentation

### Coding Instructions Reference

**CRITICAL**: Follow these instructions precisely:

```markdown
From .github/instructions/general-coding.instructions.md:

## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS
**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction of any kind.**

## Version Update Requirements
**When modifying any file with a version header, ALWAYS update the version number:**
- Patch version (x.y.Z): Bug fixes, typos, minor formatting changes
- Minor version (x.Y.z): New features, significant content additions
- Major version (X.y.z): Breaking changes, structural overhauls
```

### File Header Requirements

All modified files must have proper headers:

```go
// file: path/to/file.go
// version: 1.1.0  // INCREMENT THIS
// guid: existing-guid-unchanged
```

## 🧪 Testing Requirements

### Unit Tests

Create tests for auth migration:

```go
// file: pkg/auth/migration_test.go
func TestAuthMigration(t *testing.T) {
    // Test that old auth types can be converted to new format
    // Test that all required fields are properly set using opaque API
    // Test that role mappings work correctly
}
```

### Integration Tests

```go
// file: pkg/auth/integration_test.go
func TestAuthIntegration(t *testing.T) {
    // Test full authentication flow with new types
    // Test OAuth flow with new user types
    // Test session management with new session types
}
```

## 🎯 Success Metrics

- [ ] All gcommonauth imports removed
- [ ] `go build ./...` completes successfully
- [ ] All existing tests pass
- [ ] New tests added with 80%+ coverage
- [ ] No direct field access to protobuf fields (all via getters/setters)
- [ ] Authentication flows work correctly
- [ ] OAuth integration still functional

## 🚨 Common Pitfalls

1. **Role Enum Mapping**: Ensure local role constants map correctly to gcommon role enums
2. **Opaque API Confusion**: Remember to use `SetField()` and `GetField()` methods
3. **Nil Pointer Issues**: Always check for nil before calling methods on protobuf messages
4. **Session Expiration**: Ensure timestamp handling works correctly with protobuf timestamps
5. **Database Compatibility**: Verify that user/session storage works with new types

## 📖 Additional Resources

- [gcommon common documentation](../../gcommon-api/common.md)
- [Protobuf Go Generated Code Guide](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
- [General Coding Instructions](../../../.github/instructions/general-coding.instructions.md)

## 🔄 Related Tasks

- **TASK-01-001**: Replace configpb (may share some authentication config)
- **TASK-01-002**: Replace databasepb (user storage may be affected)
- **TASK-01-005**: Migrate protobuf message types (authentication messages)

## 📝 Notes for AI Agent

- This task affects critical authentication flows - test thoroughly
- Standard command-line operations throughout - no VS Code specific requirements
- Follow the version update rules strictly - increment version numbers in all modified files
- The opaque API is critical - never access protobuf fields directly, always use getters/setters
- Pay special attention to role mappings and enum values
- Test OAuth flows carefully after migration
- Ensure database user operations still work correctly
- If any step fails, document the error and continue with remaining steps where possible
