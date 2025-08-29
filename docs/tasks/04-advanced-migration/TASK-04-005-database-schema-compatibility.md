# file: docs/tasks/04-advanced-migration/TASK-04-005-database-schema-compatibility.md

# version: 1.0.0

# guid: e8f9a0b1-2c3d-4e5f-6a7b-8c9d0e1f2a3b

# TASK-04-005: Database Schema Compatibility

## Overview

**Objective**: Ensure database schemas work with gcommon types and add any
missing fields required for full gcommon compatibility.

**Phase**: 2 (Core Type Migration) **Priority**: High **Estimated Effort**: 4-6
hours **Prerequisites**: TASK-04-001 (User migration) and TASK-04-002 (Session
migration)

## Required Reading

**CRITICAL**: Read these documents before starting:

- `docs/gcommon-api/database.md` - Database type specifications and schema
  requirements
- `docs/gcommon-api/common.md` - Common type field requirements
- `pkg/database/schema.go` - Current database schema implementation
- `pkg/database/migrations/` - Existing migration files
- `docs/MIGRATION_INVENTORY.md` - Field mappings and compatibility requirements

## Problem Statement

The current database schema was designed for local protobuf types and may lack
fields required by gcommon types. After migrating to gcommon User and Session
types, the database schema must support all fields and constraints required by
these types across multiple database backends.

### Current Issues

1. **Missing Fields**: gcommon types may have fields not present in current
   schema
2. **Type Mismatches**: Field types may not match gcommon requirements
3. **Constraint Differences**: Foreign keys, indexes, and constraints may be
   incompatible
4. **Multi-Backend Support**: Schema must work with SQLite, PebbleDB, and
   PostgreSQL

## Technical Approach

### Database Schema Analysis

```go
// Current User table structure (example)
type UserTable struct {
    ID        string `db:"id"`
    Username  string `db:"username"`
    Email     string `db:"email"`
    CreatedAt time.Time `db:"created_at"`
}

// Required fields for gcommon User (check actual requirements)
// May need: role, permissions, metadata, auth_provider, etc.
```

### Migration Strategy

1. **Schema Comparison**: Compare current schema with gcommon requirements
2. **Additive Migrations**: Add missing fields without breaking existing data
3. **Type Compatibility**: Ensure field types match gcommon expectations
4. **Index Optimization**: Add indexes for gcommon query patterns

## Implementation Steps

### Step 1: Schema Analysis

```bash
# Analyze current schema
sqlite3 subtitle_manager.db ".schema users"
sqlite3 subtitle_manager.db ".schema sessions"

# Compare with gcommon requirements
# Document differences in schema_analysis.md
```

### Step 2: Create Migration Scripts

```go
// Example migration: add gcommon User fields
// File: pkg/database/migrations/003_add_gcommon_user_fields.go
package migrations

import (
    "database/sql"
    "fmt"
)

func AddGcommonUserFields(db *sql.DB, backend string) error {
    migrations := map[string][]string{
        "sqlite": {
            "ALTER TABLE users ADD COLUMN role TEXT DEFAULT 'user'",
            "ALTER TABLE users ADD COLUMN auth_provider TEXT DEFAULT 'local'",
            "ALTER TABLE users ADD COLUMN metadata_json TEXT DEFAULT '{}'",
            "ALTER TABLE users ADD COLUMN permissions_json TEXT DEFAULT '[]'",
        },
        "postgresql": {
            "ALTER TABLE users ADD COLUMN role VARCHAR(50) DEFAULT 'user'",
            "ALTER TABLE users ADD COLUMN auth_provider VARCHAR(100) DEFAULT 'local'",
            "ALTER TABLE users ADD COLUMN metadata_json JSONB DEFAULT '{}'",
            "ALTER TABLE users ADD COLUMN permissions_json JSONB DEFAULT '[]'",
        },
    }

    statements := migrations[backend]
    for _, stmt := range statements {
        if _, err := db.Exec(stmt); err != nil {
            return fmt.Errorf("migration failed: %v", err)
        }
    }
    return nil
}
```

### Step 3: Update Schema Definitions

```go
// File: pkg/database/schema.go
package database

// UpdatedUserSchema represents the schema after gcommon migration
type UpdatedUserSchema struct {
    ID           string    `db:"id" json:"id"`
    Username     string    `db:"username" json:"username"`
    Email        string    `db:"email" json:"email"`
    Role         string    `db:"role" json:"role"`
    AuthProvider string    `db:"auth_provider" json:"auth_provider"`
    MetadataJSON string    `db:"metadata_json" json:"metadata_json"`
    PermissionsJSON string `db:"permissions_json" json:"permissions_json"`
    CreatedAt    time.Time `db:"created_at" json:"created_at"`
    UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// Conversion methods for gcommon compatibility
func (u *UpdatedUserSchema) ToGcommonUser() (*common.User, error) {
    user := &common.User{}
    user.SetId(u.ID)
    user.SetUsername(u.Username)
    user.SetEmail(u.Email)
    user.SetRole(u.Role)

    // Parse metadata from JSON
    var metadata map[string]interface{}
    if err := json.Unmarshal([]byte(u.MetadataJSON), &metadata); err != nil {
        return nil, fmt.Errorf("failed to parse metadata: %v", err)
    }

    for key, value := range metadata {
        user.SetField(key, value)
    }

    return user, nil
}

func (u *UpdatedUserSchema) FromGcommonUser(user *common.User) error {
    u.ID = user.GetId()
    u.Username = user.GetUsername()
    u.Email = user.GetEmail()
    u.Role = user.GetRole()

    // Serialize metadata to JSON
    metadata := make(map[string]interface{})
    // Extract custom fields from gcommon User
    // (Implementation depends on gcommon API)

    metadataBytes, err := json.Marshal(metadata)
    if err != nil {
        return fmt.Errorf("failed to serialize metadata: %v", err)
    }
    u.MetadataJSON = string(metadataBytes)

    return nil
}
```

### Step 4: Database Backend Updates

```go
// File: pkg/database/sqlite.go
func (s *SQLiteDB) CreateUserTables() error {
    createSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        username TEXT UNIQUE NOT NULL,
        email TEXT UNIQUE NOT NULL,
        role TEXT DEFAULT 'user',
        auth_provider TEXT DEFAULT 'local',
        metadata_json TEXT DEFAULT '{}',
        permissions_json TEXT DEFAULT '[]',
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );

    CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
    CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
    CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);
    `

    _, err := s.db.Exec(createSQL)
    return err
}

// File: pkg/database/pebbledb.go
func (p *PebbleDB) EnsureUserSchema() error {
    // PebbleDB is key-value, so ensure consistent serialization
    // Test with sample gcommon User to verify compatibility
    testUser := &common.User{}
    testUser.SetId("test")
    testUser.SetUsername("test")
    testUser.SetEmail("test@example.com")

    // Serialize and deserialize to test compatibility
    data, err := p.SerializeUser(testUser)
    if err != nil {
        return fmt.Errorf("serialization test failed: %v", err)
    }

    _, err = p.DeserializeUser(data)
    if err != nil {
        return fmt.Errorf("deserialization test failed: %v", err)
    }

    return nil
}
```

### Step 5: Rollback Procedures

```go
// File: pkg/database/migrations/rollback_003.go
func RollbackGcommonUserFields(db *sql.DB, backend string) error {
    rollbackSQL := map[string][]string{
        "sqlite": {
            // SQLite doesn't support DROP COLUMN, so create backup table
            "CREATE TABLE users_backup AS SELECT id, username, email, created_at FROM users",
            "DROP TABLE users",
            "ALTER TABLE users_backup RENAME TO users",
        },
        "postgresql": {
            "ALTER TABLE users DROP COLUMN role",
            "ALTER TABLE users DROP COLUMN auth_provider",
            "ALTER TABLE users DROP COLUMN metadata_json",
            "ALTER TABLE users DROP COLUMN permissions_json",
        },
    }

    statements := rollbackSQL[backend]
    for _, stmt := range statements {
        if _, err := db.Exec(stmt); err != nil {
            return fmt.Errorf("rollback failed: %v", err)
        }
    }
    return nil
}
```

## Testing Requirements

### Database Migration Tests

```go
// File: pkg/database/migrations_test.go
func TestSchemaCompatibility(t *testing.T) {
    tests := []struct {
        name    string
        backend string
        setup   func() *sql.DB
    }{
        {"SQLite", "sqlite", setupSQLiteTest},
        {"PostgreSQL", "postgresql", setupPostgreSQLTest},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            db := tt.setup()
            defer db.Close()

            // Test forward migration
            err := AddGcommonUserFields(db, tt.backend)
            require.NoError(t, err)

            // Test data compatibility
            testUser := createTestGcommonUser()
            err = storeUserInDB(db, testUser)
            require.NoError(t, err)

            retrievedUser, err := loadUserFromDB(db, testUser.GetId())
            require.NoError(t, err)

            // Verify all fields preserved
            assert.Equal(t, testUser.GetId(), retrievedUser.GetId())
            assert.Equal(t, testUser.GetUsername(), retrievedUser.GetUsername())
            assert.Equal(t, testUser.GetEmail(), retrievedUser.GetEmail())

            // Test rollback
            err = RollbackGcommonUserFields(db, tt.backend)
            require.NoError(t, err)
        })
    }
}

func TestPebbleDBCompatibility(t *testing.T) {
    pdb := setupPebbleDBTest()
    defer pdb.Close()

    // Test gcommon User serialization
    testUser := createTestGcommonUser()

    err := pdb.StoreUser(testUser)
    require.NoError(t, err)

    retrievedUser, err := pdb.GetUser(testUser.GetId())
    require.NoError(t, err)

    // Verify opaque API fields preserved
    assert.Equal(t, testUser.GetId(), retrievedUser.GetId())
    assert.Equal(t, testUser.GetUsername(), retrievedUser.GetUsername())

    // Test custom field preservation
    testUser.SetField("custom_field", "test_value")
    err = pdb.StoreUser(testUser)
    require.NoError(t, err)

    retrievedUser, err = pdb.GetUser(testUser.GetId())
    require.NoError(t, err)

    customValue, exists := retrievedUser.GetField("custom_field")
    assert.True(t, exists)
    assert.Equal(t, "test_value", customValue)
}
```

### Performance Tests

```go
func TestMigrationPerformance(t *testing.T) {
    // Test migration performance with large datasets
    db := setupSQLiteTest()
    defer db.Close()

    // Create 10,000 test users
    for i := 0; i < 10000; i++ {
        createTestUser(db, fmt.Sprintf("user%d", i))
    }

    start := time.Now()
    err := AddGcommonUserFields(db, "sqlite")
    require.NoError(t, err)

    migrationTime := time.Since(start)
    t.Logf("Migration of 10,000 users took: %v", migrationTime)

    // Migration should complete in reasonable time (< 30s)
    assert.Less(t, migrationTime, 30*time.Second)
}
```

## Validation Scripts

### Schema Validation

```bash
#!/bin/bash
# File: scripts/validate_schema.sh

echo "Validating database schema compatibility..."

# Check SQLite schema
echo "=== SQLite Schema ==="
sqlite3 test.db ".schema users"
sqlite3 test.db ".schema sessions"

# Check PostgreSQL schema (if available)
if command -v psql >/dev/null 2>&1; then
    echo "=== PostgreSQL Schema ==="
    psql -d test_db -c "\d users"
    psql -d test_db -c "\d sessions"
fi

# Validate migration files
echo "=== Migration Files ==="
ls -la pkg/database/migrations/
```

### Data Integrity Check

```go
// File: scripts/check_data_integrity.go
package main

import (
    "fmt"
    "log"

    "github.com/jdfalk/subtitle-manager/pkg/database"
    "github.com/jdfalk/gcommon/sdks/go/v1/common"
)

func main() {
    db, err := database.NewSQLiteDB("subtitle_manager.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Verify all users can be loaded as gcommon types
    users, err := db.GetAllUsers()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Checking %d users for gcommon compatibility...\n", len(users))

    for i, user := range users {
        // Verify each user meets gcommon requirements
        if user.GetId() == "" {
            fmt.Printf("User %d: missing ID\n", i)
        }
        if user.GetUsername() == "" {
            fmt.Printf("User %d: missing username\n", i)
        }
        if user.GetEmail() == "" {
            fmt.Printf("User %d: missing email\n", i)
        }

        // Test opaque API access
        role, exists := user.GetField("role")
        if !exists {
            fmt.Printf("User %d: missing role field\n", i)
        } else {
            fmt.Printf("User %d: role = %v\n", i, role)
        }
    }

    fmt.Println("Data integrity check complete")
}
```

## Success Metrics

### Functional Requirements

- [ ] Schema supports all gcommon User fields without data loss
- [ ] Schema supports all gcommon Session fields without data loss
- [ ] Migration scripts work on SQLite, PebbleDB, and PostgreSQL
- [ ] Rollback procedures restore original schema correctly
- [ ] No performance degradation after migration

### Technical Requirements

- [ ] All existing data preserved during migration
- [ ] New fields have appropriate defaults and constraints
- [ ] Indexes optimized for gcommon query patterns
- [ ] Foreign key constraints maintained
- [ ] Database size increase < 20%

### Validation Requirements

- [ ] All database tests pass with new schema
- [ ] Performance tests show acceptable migration time
- [ ] Data integrity verification script passes
- [ ] Schema validation script confirms compatibility
- [ ] Load testing with gcommon types successful

## Common Pitfalls

1. **Data Loss During Migration**: Always backup before schema changes
2. **Type Conversion Errors**: Test with real data, not just synthetic
3. **Performance Impact**: Monitor query performance after adding fields
4. **Cross-Platform Issues**: Different SQL dialects have different constraints
5. **Rollback Failures**: SQLite DROP COLUMN limitations require special
   handling

## Dependencies

- **Requires**: TASK-04-001 (User type migration) completed
- **Requires**: TASK-04-002 (Session type migration) completed
- **Enables**: All subsequent database operations with gcommon types
- **Blocks**: Cannot complete package replacements without schema compatibility

## Embedded Documentation

### gcommon Database Types Reference

```go
// Key gcommon types that require schema support
type User interface {
    GetId() string
    SetId(string)
    GetUsername() string
    SetUsername(string)
    GetEmail() string
    SetEmail(string)
    GetRole() string
    SetRole(string)
    GetField(string) (interface{}, bool)
    SetField(string, interface{})
}

type Session interface {
    GetId() string
    SetId(string)
    GetUserId() string
    SetUserId(string)
    GetExpiresAt() time.Time
    SetExpiresAt(time.Time)
    GetField(string) (interface{}, bool)
    SetField(string, interface{})
}
```

### Migration Best Practices

1. **Always Backup**: Create database backup before migration
2. **Test Locally**: Run migration on development database first
3. **Monitor Performance**: Check query performance after schema changes
4. **Validate Data**: Verify data integrity after migration
5. **Plan Rollback**: Have tested rollback procedure ready

### Database Backend Considerations

- **SQLite**: Limited ALTER TABLE support, use backup/restore for complex
  changes
- **PostgreSQL**: Full ALTER TABLE support, JSONB for metadata storage
- **PebbleDB**: Key-value store, ensure consistent serialization format

This comprehensive task ensures the database schema fully supports gcommon types
while maintaining data integrity and performance across all database backends.
