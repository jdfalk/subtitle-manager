<!-- file: .github/copilot-instructions.md -->
<!-- version: 4.0.0 -->
<!-- guid: 4d5e6f7a-8b9c-0d1e-2f3a-4b5c6d7e8f9a -->

# Subtitle Manager - AI Agent Guide

This is a comprehensive subtitle management application written in Go with a React web interface. The project provides both CLI and web interfaces for converting, translating, and managing subtitle files with enterprise-grade features.

## 🚨 CRITICAL: Documentation Update Protocol

**NEVER edit markdown files directly. ALWAYS use the documentation update system:**

1. **Create GitHub Issue First** (if none exists):
   ```bash
   ./scripts/create-issue-update.sh "Update [filename] - [description]" "Detailed description of what needs to be updated"
   ```

2. **Create Documentation Update**:
   ```bash
   ./scripts/create-doc-update.sh [filename] "[content]" [mode] --issue [issue-number]
   ```

3. **Link to Issue**: Every documentation change MUST reference a GitHub issue for tracking and context.

## 🏗️ Architecture Overview

**Hybrid CLI + Web Application**: The main binary serves dual purposes - CLI tool via Cobra commands and web server with React UI embedded via `webui/embed.go`.

**Key Components:**
- **CLI Layer** (`cmd/`): Cobra-based commands for all operations (convert, fetch, sync, etc.)
- **Core Engine** (`pkg/engine/`): Business logic and service orchestration
- **Web Server** (`pkg/webserver/`): REST API and embedded React UI serving
- **Database Layer** (`pkg/database/`): Multi-backend support (SQLite, PebbleDB, PostgreSQL)
- **Provider System** (`pkg/providers/`): 40+ subtitle provider integrations (Bazarr-compatible)
- **React Frontend** (`webui/`): Material-UI based single-page application

## 🔧 Essential Development Workflows

### Build System (Use Makefile targets)
```bash
make help                    # Show all available targets
make build                   # Build Go binary with version info
make build-web              # Build React frontend only
make build-all              # Build both backend and frontend
make docker                 # Build Docker image
make test                   # Run Go tests
make test-web               # Run frontend tests
```

### Protocol Buffer Generation
```bash
make generate               # Generate Go code from .proto files
buf generate                # Direct buf command (after changes to proto/)
```

### Frontend Development
```bash
cd webui && npm run dev     # Start Vite dev server on :5173
npm run build              # Build for production (outputs to webui/dist/)
```

### VS Code Tasks (Preferred)
Use VS Code tasks instead of manual commands when available:
- `Git Add All`, `Git Commit`, `Git Push` for version control
- `Buf Generate with Output` for protobuf compilation

## 🎯 Project-Specific Patterns

### Command Structure
All CLI commands follow this pattern in `cmd/`:
```go
// Each command file exports its cobra.Command
var convertCmd = &cobra.Command{
    Use:   "convert [input] [output]",
    Short: "Convert subtitle formats",
    RunE:  runConvert,  // Actual implementation
}

func init() {
    rootCmd.AddCommand(convertCmd)  // Register with root
}
```

### Database Abstraction
Multi-backend database support via `pkg/database/`:
```go
// All database operations use this interface
type DB interface {
    Store(key []byte, value []byte) error
    Get(key []byte) ([]byte, error)
    // ... other methods
}
// Implementations: SQLiteDB, PebbleDB, PostgresDB
```

### Provider Registration
Subtitle providers auto-register via init() functions:
```go
func init() {
    providers.Register("opensubtitles", &OpenSubtitlesProvider{})
}
```

### Configuration Management
Uses Viper for configuration with multiple sources:
- Command-line flags (highest priority)
- Environment variables (`SUBTITLE_MANAGER_*`)
- Config file (`~/.subtitle-manager.yaml`)
- Built-in defaults (lowest priority)

## 🔌 Critical Integration Points

### gcommon Dependency
This project imports shared protocol buffers from `github.com/jdfalk/gcommon`:
- Common types in `pkg/gcommon/` (users, sessions, config)
- Generated Go code maps gcommon protos via `buf.gen.yaml` import mappings
- Updates require coordination between repositories

### Embedded Web UI
React app is embedded in Go binary:
```go
//go:embed webui/dist/*
var webUIFiles embed.FS
```
Build process: `npm run build` → `webui/dist/` → Go embed → single binary

### External Service Integrations
- **Translation**: Google Translate, OpenAI ChatGPT APIs
- **Transcription**: OpenAI Whisper integration
- **Media Integration**: Sonarr, Radarr, Plex APIs
- **Storage**: S3, Azure Blob, Google Cloud Storage
- **Authentication**: OAuth2 (GitHub), password, API keys

### Database Schemas
Key entities: Users, Media, Subtitles, Translation History, Provider Configs
- Migration system in `cmd/migrate.go`
- Schema evolution tracked in `pkg/database/migrations/`

## 🛠️ Development Environment Setup

### Prerequisites
- Go 1.24+ (specified in go.mod)
- Node.js 24+ for frontend
- buf CLI for protobuf generation
- Docker for containerized builds

### Quick Start
```bash
# 1. Install dependencies
go mod download
cd webui && npm install

# 2. Build frontend
make build-web

# 3. Build application
make build

# 4. Run with embedded UI
./bin/subtitle-manager web --port 8080
```

### Testing Strategy
- **Go tests**: `*_test.go` files alongside source
- **Integration tests**: `test/` directory with real provider tests
- **Frontend tests**: Vitest + Playwright in `webui/tests/`
- **E2E tests**: End-to-end workflows in `tests/`

## 📝 Code Style Notes

### Package Organization
- `cmd/`: CLI command implementations (one file per command)
- `pkg/`: Core library packages (importable by external code)
- `webui/`: Complete React application
- `proto/`: Protocol buffer definitions
- `sdks/`: Generated code for multiple languages

### Error Handling
Consistent error wrapping with context:
```go
if err != nil {
    return fmt.Errorf("failed to process subtitle file %s: %w", filename, err)
}
```

### Logging
Per-component loggers via `pkg/logging`:
```go
logger := logging.GetLogger("subtitle.converter")
logger.WithField("file", filename).Info("Converting subtitle")
```

## 🔧 Provider System Architecture

### Factory Registry Pattern
All 48+ subtitle providers use centralized factory registration:
```go
// pkg/providers/registry.go - Factory-based provider creation
var factories = map[string]func() Provider{
    "opensubtitles":  func() Provider { return opensubtitles.New() },
    "addic7ed":       func() Provider { return addic7ed.New() },
    "subscene":       func() Provider { return subscene.New() },
    // ... 45+ more providers
}

// Dynamic provider lookup
func Get(name, _ string) (Provider, error) {
    if f, ok := factories[name]; ok {
        return f(), nil
    }
    return nil, fmt.Errorf("unknown provider %s", name)
}
```

### Provider Interface Standards
All providers implement consistent interfaces:
```go
// Core provider interface (required)
type Provider interface {
    Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error)
}

// Optional search interface (many providers)
type Searcher interface {
    Search(ctx context.Context, mediaPath, lang string) ([]string, error)
}
```

### Provider Configuration Patterns
Provider-specific configuration through Viper with fallbacks:
```go
// Each provider gets namespaced config keys
viper.GetString("providers.opensubtitles.api_key")
viper.GetString("providers.addic7ed.username")
viper.GetBool("providers.embedded.enabled")
```

### Frontend Provider Management
React frontend uses Bazarr-style provider cards with dynamic loading:
```jsx
// Provider metadata structure
{
  name: "opensubtitles",
  displayName: "OpenSubtitles",
  description: "Community-driven subtitle database",
  supportedLanguages: ["en", "es", "fr", ...],
  configFields: [
    { name: "apiKey", type: "password", required: true }
  ]
}
```

## 🛠️ Development & Debugging Workflows

### Testing Patterns
Use consistent test patterns across the codebase:
```go
// Test helper utilities
func TestProviderFetch(t *testing.T) {
    skipIfNoSQLite(t)  // Skip tests requiring database
    db := testutil.GetTestDB(t)
    defer db.Close()

    key := setupTestUser(t, db)  // Create test user with API key
    testutil.MustNoError(t, "operation", err)  // Standard error checking
}
```

### gRPC Development
Use the test client for gRPC server development:
```bash
# Test gRPC functionality with included client
go run test_grpc_client.go
```

### Configuration Debugging
Debug configuration issues using Viper's built-in tools:
```go
// Check configuration sources
fmt.Println("Using config file:", viper.ConfigFileUsed())
fmt.Printf("All settings: %+v\n", viper.AllSettings())

// Environment variable mapping
// SM_OPENAI_API_KEY maps to openai_api_key
// SM_PROVIDERS_EMBEDDED_ENABLED maps to providers.embedded.enabled
```

### Database Backend Switching
Support multiple database backends through environment configuration:
```bash
# Use SQLite for development
export SM_DB_BACKEND=sqlite
export SM_SQLITE3_FILENAME=dev.db

# Use PebbleDB for production (default)
export SM_DB_BACKEND=pebble

# Use PostgreSQL for enterprise
export SM_DB_BACKEND=postgres
export SM_POSTGRES_URL=postgres://user:pass@host/db
```

### Cache System Development
Configure caching for different environments:
```yaml
# Development - memory cache
cache:
  backend: memory
  memory:
    max_entries: 1000

# Production - Redis cache
cache:
  backend: redis
  redis:
    address: redis:6379
    password: secret
```

### Provider Development Workflow
1. **Add Provider**: Create new provider in `pkg/providers/newprovider/`
2. **Register**: Add to factory map in `registry.go`
3. **Configure**: Add default config values in `cmd/root.go`
4. **Test**: Create tests following existing patterns
5. **Frontend**: Add to provider metadata in frontend components

---

For complete coding standards, see `.github/instructions/` directory. This guide focuses on architectural understanding essential for productive development in this codebase.
