# gcommon SDK Documentation

This directory contains comprehensive API documentation for all gcommon v1.8.0
packages that subtitle-manager will use during the refactoring process.

## Package Overview

| Package                              | Size  | Description                             | Primary Use Cases              |
| ------------------------------------ | ----- | --------------------------------------- | ------------------------------ |
| [common.md](./common.md)             | 4.9MB | Core common types (User, Session, etc.) | Authentication, base types     |
| [config.md](./config.md)             | 1.6MB | Configuration management                | App configuration, settings    |
| [database.md](./database.md)         | 1.6MB | Database operations                     | Data storage, queries          |
| [health.md](./health.md)             | 641KB | Health monitoring                       | System health checks           |
| [media.md](./media.md)               | 886KB | Media handling                          | Video/subtitle file processing |
| [metrics.md](./metrics.md)           | 2.9MB | Metrics and monitoring                  | Performance tracking           |
| [organization.md](./organization.md) | 1.6MB | Organization management                 | Multi-tenant features          |
| [queue.md](./queue.md)               | 4.7MB | Queue operations                        | Background job processing      |
| [web.md](./web.md)                   | 1.8MB | Web services                            | HTTP handlers, middleware      |

## Migration Priority

1. **Phase 1 - Core Types**: `common` package (User, Session, basic types)
2. **Phase 2 - Configuration**: `config` package (replace configpb)
3. **Phase 3 - Database**: `database` package (replace databasepb)
4. **Phase 4 - Advanced**: `health`, `media`, `metrics`, `queue`, `web` packages

## Important Notes

- All packages use the **opaque API pattern** - use setter/getter methods
  instead of direct field access
- Generated documentation includes both exported and unexported symbols (-u
  flag)
- Documentation generated from gcommon v1.8.0
- Each package provides protobuf-generated types with Go-friendly APIs

## Usage Examples

### Common Package (User Management)

```go
import "github.com/jdfalk/gcommon/sdks/go/v1/common"

// Create user with opaque API
user := &common.User{}
user.SetId("user-123")
user.SetUsername("john.doe")
user.SetEmail("john@example.com")
user.SetRole(common.UserRole_USER_ROLE_ADMIN)

// Get values using getters
id := user.GetId()
email := user.GetEmail()
```

### Config Package (Settings)

```go
import "github.com/jdfalk/gcommon/sdks/go/v1/config"

// Application configuration
appConfig := &config.ApplicationConfig{}
appConfig.SetName("subtitle-manager")
appConfig.SetVersion("2.0.0")
```

### Database Package (Operations)

```go
import "github.com/jdfalk/gcommon/sdks/go/v1/database"

// Database configuration
dbConfig := &database.DatabaseConfig{}
dbConfig.SetHost("localhost")
dbConfig.SetPort(5432)
dbConfig.SetName("subtitles")
```

This documentation provides complete reference material for the gcommon
refactoring process.
