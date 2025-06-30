# Subtitle Manager - Complete Technical Analysis and Implementation Roadmap

## Executive Summary

Subtitle Manager is a Go-based subtitle management application that aims for
feature parity with Bazarr. The project is approximately 85% complete with a
robust backend, 40+ subtitle providers, and a React-based web UI. This document
provides a complete technical breakdown and the shortest path to full
operational status with Bazarr parity.

## Current State Analysis

### Core Architecture

#### Backend (Go)

- **Framework**: Cobra CLI with Viper configuration
- **Web Server**: Custom HTTP server with gorilla/websocket support
- **Database**: Multi-backend support (SQLite, PostgreSQL, PebbleDB)
- **Authentication**: Session-based auth with RBAC, OAuth2 (GitHub), API keys
- **Providers**: 40+ subtitle providers with plugin architecture
- **Translation**: Google Translate, OpenAI GPT, gRPC service support
- **Media Integration**: Sonarr, Radarr, Plex support

#### Frontend (React)

- **Framework**: React 18 with Material-UI
- **Build**: Vite-based build system
- **Features**: Complete dashboard, media library, settings, system monitoring
- **State Management**: React hooks with API service layer
- **Testing**: Playwright E2E tests, Vitest unit tests

### Technical Component Breakdown

#### 1. Command Layer (`cmd/`)

##### Core Commands

- `root.go`: Base command initialization, config loading via Viper
  - Initializes logging, database, and global configuration
  - Environment variable mapping (SM\_\* prefix)
- `web.go`: Web server launcher
  - Starts HTTP server on configurable port
  - Integrates embedded React UI

##### Media Operations

- `scan.go`: Directory scanning for subtitle downloads
  - Multi-provider support with fallback
  - Progress tracking and parallel processing
- `scanlib.go`: Media library indexing

  - TMDB/OMDb metadata fetching
  - Database storage of media items

- `watch.go`: File system monitoring
  - Real-time subtitle downloading
  - Recursive directory support

##### Subtitle Operations

- `convert.go`: Format conversion (any → SRT)
- `translate.go`: Multi-service translation
- `sync.go`: Audio-based synchronization
- `extract.go`: Embedded subtitle extraction
- `merge.go`: Combine multiple subtitle tracks

##### Integration Commands

- `sonarr.go`, `radarr.go`: \*arr integration
- `plex.go`: Plex library sync and refresh
- `import.go`: Bazarr settings import

#### 2. Core Packages (`pkg/`)

##### Authentication & Security (`auth/`)

- User management with bcrypt passwords
- Session handling with Redis-like in-memory store
- RBAC with configurable permissions
- OAuth2 GitHub integration
- API key generation and validation

##### Database Layer (`database/`)

- Abstract `SubtitleStore` interface
- SQLite implementation with migrations
- PostgreSQL support with full feature parity
- PebbleDB for embedded pure-Go option
- Migration utilities between backends

##### Provider System (`providers/`)

- Provider interface with context support
- Registry pattern for dynamic loading
- Instance-based configuration with priorities
- Tag-based provider selection
- Built-in providers for major services

##### Web Server (`webserver/`)

- RESTful API endpoints
- WebSocket support for real-time updates
- Static file serving with SPA support
- Security headers and CORS handling
- Comprehensive middleware stack

##### Media & Metadata (`metadata/`)

- Filename parsing for TV/Movie detection
- TMDB API integration
- OMDb API integration
- Library scanning with progress callbacks

##### Task Management (`tasks/`)

- Concurrent task execution
- Progress tracking and reporting
- WebSocket broadcast for UI updates

##### Scheduler (`scheduler/`)

- Cron-based task scheduling
- Interval-based execution
- Jitter and max run support

##### Notifications (`notifications/`)

- Discord webhook support
- Telegram bot integration
- SMTP email notifications
- Extensible notifier interface

### Current Limitations & Missing Features

#### Critical Missing Features (Bazarr Parity)

1. **Whisper Integration**: No embedded Whisper ASR support
2. **Advanced Provider Features**: Missing some provider-specific options
3. **Subtitle Scoring**: No quality-based subtitle selection
4. **Language Profiles**: Basic language support, no profiles
5. **Episode Monitoring**: Limited automatic monitoring
6. **Backup/Restore**: Basic backup, no automated restore
7. **Webhook System**: Partial implementation
8. **Manual Search UI**: Backend ready, UI incomplete

#### Technical Debt

1. **Test Coverage**: ~70% coverage, needs improvement
2. **Error Handling**: Inconsistent error propagation
3. **Configuration**: Some hardcoded values remain
4. **Documentation**: API documentation incomplete
5. **Performance**: No caching layer for API calls

## Shortest Path to Full Operational Status

### Phase 1: Critical Features (2-3 weeks)

#### 1.1 Whisper Integration (Issue #1132)

\`\`\`go // pkg/transcriber/whisper_container.go type WhisperContainer struct {
client \*docker.Client containerID string }

func (w *WhisperContainer) Start() error func (w *WhisperContainer)
Transcribe(mediaPath string) ([]byte, error) \`\`\`

#### 1.2 Advanced Provider Configuration

- Implement provider-specific settings UI
- Add quality scoring for subtitle selection
- Enable provider chaining and fallback

#### 1.3 Language Profiles

\`\`\`go // pkg/profiles/language.go type LanguageProfile struct { ID string
Name string Languages []LanguageConfig CutoffScore int } \`\`\`

### Phase 2: Enhanced Monitoring (1-2 weeks)

#### 2.1 Episode Monitoring

- Implement missing episode detection
- Add automatic subtitle upgrade
- Create monitoring dashboard

#### 2.2 Webhook System Completion

\`\`\`go // pkg/webhooks/manager.go type WebhookManager struct { endpoints
map[string]WebhookEndpoint } \`\`\`

### Phase 3: UI Enhancements (2-3 weeks)

#### 3.1 Manual Search Interface

- Implement provider search UI
- Add subtitle preview
- Enable manual selection and download

#### 3.2 Settings Migration

- Complete Bazarr settings import
- Add export functionality
- Create settings backup/restore

### Phase 4: Performance & Reliability (1-2 weeks)

#### 4.1 Caching Layer

\`\`\`go // pkg/cache/redis.go type CacheLayer interface { Get(key string)
(interface{}, error) Set(key string, value interface{}, ttl time.Duration) error
} \`\`\`

#### 4.2 Error Handling Standardization

- Implement consistent error types
- Add error recovery mechanisms
- Improve logging and debugging

### Phase 5: Testing & Documentation (1 week)

#### 5.1 Test Coverage

- Achieve 90%+ test coverage
- Add integration test suite
- Implement load testing

#### 5.2 Documentation

- Complete API documentation
- Add user guide
- Create developer documentation

## Implementation Priority Matrix

| Feature                | Impact | Effort | Priority |
| ---------------------- | ------ | ------ | -------- |
| Whisper Integration    | High   | High   | 1        |
| Language Profiles      | High   | Medium | 2        |
| Manual Search UI       | High   | Medium | 3        |
| Provider Scoring       | Medium | Low    | 4        |
| Episode Monitoring     | High   | Medium | 5        |
| Webhook Completion     | Medium | Low    | 6        |
| Caching Layer          | Medium | Medium | 7        |
| Settings Import/Export | Low    | Low    | 8        |

## Technical Recommendations

### Architecture Improvements

1. **Implement Service Layer** \`\`\`go // pkg/services/subtitle_service.go type
   SubtitleService struct { store database.SubtitleStore providers
   []providers.Provider cache cache.CacheLayer } \`\`\`

2. **Add Event Bus** \`\`\`go // pkg/events/bus.go type EventBus interface {
   Publish(event Event) error Subscribe(eventType string, handler EventHandler)
   error } \`\`\`

3. **Standardize API Responses** \`\`\`go // pkg/api/response.go type
   APIResponse struct { Success bool \`json:"success"\` Data interface{}
   \`json:"data,omitempty"\` Error \*APIError \`json:"error,omitempty"\` }
   \`\`\`

### Performance Optimizations

1. **Database Indexing** \`\`\`sql CREATE INDEX idx_media_items_path ON
   media_items(path); CREATE INDEX idx_subtitles_video_file ON
   subtitles(video_file); \`\`\`

2. **Connection Pooling** \`\`\`go db.SetMaxOpenConns(25) db.SetMaxIdleConns(5)
   db.SetConnMaxLifetime(5 \* time.Minute) \`\`\`

3. **Concurrent Processing** \`\`\`go // Use worker pools for batch operations
   pool := pond.New(10, 1000) for \_, item := range items { pool.Submit(func() {
   processItem(item) }) } \`\`\`

## Migration Strategy from Bazarr

### Step 1: Database Migration

\`\`\`bash

# Export Bazarr database

sqlite3 bazarr.db .dump > bazarr_dump.sql

# Import to Subtitle Manager

subtitle-manager migrate import-bazarr bazarr_dump.sql \`\`\`

### Step 2: Configuration Migration

\`\`\`bash

# Import Bazarr settings

subtitle-manager import-bazarr http://bazarr:6767 <api-key> \`\`\`

### Step 3: Provider Migration

- Map Bazarr providers to Subtitle Manager providers
- Migrate API keys and credentials
- Set up provider priorities

## Monitoring & Maintenance

### Health Checks

\`\`\`go // pkg/health/checks.go type HealthChecker interface { CheckDatabase()
error CheckProviders() error CheckDiskSpace() error } \`\`\`

### Metrics Collection

\`\`\`go // pkg/metrics/collector.go type MetricsCollector struct {
downloadCount prometheus.Counter apiLatency prometheus.Histogram } \`\`\`

## Security Considerations

1. **API Security**

   - Implement rate limiting
   - Add request signing for webhooks
   - Enable CORS configuration

2. **Data Protection**
   - Encrypt sensitive configuration
   - Implement audit logging
   - Add session timeout

## Conclusion

Subtitle Manager is a well-architected application that's close to achieving
full Bazarr parity. The shortest path involves focusing on the critical missing
features (Whisper integration, language profiles, manual search UI) while
maintaining the existing robust architecture. With the outlined implementation
plan, full operational status can be achieved in 8-10 weeks of focused
development.

The modular architecture and clean separation of concerns make it
straightforward to add the remaining features without major refactoring. The
existing test infrastructure and CI/CD pipeline ensure that new features can be
added safely and reliably.
