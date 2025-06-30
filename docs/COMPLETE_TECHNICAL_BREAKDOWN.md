# Subtitle Manager - Complete Technical Breakdown

## Table of Contents

1. [Command Layer Analysis](#command-layer-analysis)
2. [Package Layer Analysis](#package-layer-analysis)
3. [Web UI Analysis](#web-ui-analysis)
4. [API Endpoints](#api-endpoints)
5. [Database Schema](#database-schema)
6. [Provider Implementations](#provider-implementations)
7. [Integration Points](#integration-points)

## Command Layer Analysis

### Root Command (`cmd/root.go`)

**Purpose**: Initialize application configuration and logging

**Key Methods**:

- `init()`: Register flags and bind to Viper
- `initConfig()`: Load configuration from file and environment
- `GetVersionInfo()`: Return version metadata

**Configuration Hierarchy**:

1. CLI flags
2. Environment variables (SM\_\* prefix)
3. Config file (YAML)
4. Defaults

### Web Command (`cmd/web.go`)

**Purpose**: Launch HTTP server with embedded UI

**Key Methods**:

- `RunE`: Start web server on specified address
- Integrates with `webserver.StartServer()`

**Configuration**:

- `--addr`: Listen address (default: :8080)

### Scan Command (`cmd/scan.go`)

**Purpose**: Scan directory and download subtitles

**Key Methods**:

- `RunE`: Execute scan with progress tracking
- Input validation via `security.ValidateAndSanitizePath()`
- Provider selection from registry

**Features**:

- Parallel processing with configurable workers
- Upgrade mode for quality improvements
- Database storage of download history

### Media Library Commands

#### Scan Library (`cmd/scanlib.go`)

- Indexes video files with metadata
- Fetches from TMDB/OMDb APIs
- Stores in media_items table

#### Watch Directory (`cmd/watch.go`)

- File system monitoring with fsnotify
- Automatic subtitle download on new files
- Recursive directory support

### Subtitle Processing Commands

#### Convert (`cmd/convert.go`)

- Uses astisub library for format detection
- Supports: SRT, ASS, SSA, VTT, others
- Preserves timing and styling where possible

#### Translate (`cmd/translate.go`)

- Multi-service support (Google, OpenAI, gRPC)
- In-memory caching for duplicate lines
- Batch processing for multiple files

#### Sync (`cmd/sync.go`)

- Audio-based synchronization
- Embedded subtitle extraction
- Weighted alignment algorithm

#### Extract (`cmd/extract.go`)

- FFmpeg wrapper for subtitle extraction
- Multi-track support
- Format auto-detection

## Package Layer Analysis

### Authentication Package (`pkg/auth/`)

#### Core Components

**auth.go**

- `AuthenticateUser()`: Validate credentials, return user ID
- `GenerateAPIKey()`: Create secure API keys
- `HashPassword()`: bcrypt with cost 10

**session.go**

- In-memory session store
- 24-hour default expiration
- Secure cookie generation

**oauth.go**

- GitHub OAuth2 flow
- Automatic user creation
- Profile synchronization

**rbac.go**

- Permission levels: read, basic, all
- Role-based access control
- Database-backed permissions

### Database Package (`pkg/database/`)

#### Store Interface

\`\`\`go type SubtitleStore interface { // Subtitle operations
InsertSubtitle(\*SubtitleRecord) error ListSubtitles() ([]SubtitleRecord, error)
DeleteSubtitle(path string) error

    // Download tracking
    InsertDownload(*DownloadRecord) error
    ListDownloads() ([]DownloadRecord, error)

    // Media management
    InsertMediaItem(*MediaItem) error
    UpdateMediaItem(*MediaItem) error
    ListMediaItems() ([]MediaItem, error)

} \`\`\`

#### Implementations

**SQLite (`pkg/database/sqlite_enabled.go`)**

- Conditional compilation with `sqlite` tag
- Automatic migrations
- Full-text search support

**PostgreSQL (`pkg/database/postgres.go`)**

- Connection pooling
- Prepared statements
- JSONB for metadata

**PebbleDB (`pkg/database/pebble.go`)**

- Pure Go implementation
- Key-value store
- JSON serialization

### Provider Package (`pkg/providers`)

#### Provider Interface

\`\`\`go type Provider interface { Fetch(ctx context.Context, mediaPath, lang
string) ([]byte, error) }

type Searcher interface { Search(ctx context.Context, mediaPath, lang string)
([]string, error) } \`\`\`

#### Registry System

- Dynamic provider registration
- Factory pattern for instantiation
- Configuration injection

#### Instance Management

- Priority-based selection
- Tag-based filtering
- Backoff tracking

### Web Server Package (`pkg/webserver`)

#### HTTP Handlers

**API Structure**:

- `/api/config`: Configuration management
- `/api/scan`: Directory scanning
- `/api/search`: Subtitle search
- `/api/translate`: Translation service
- `/api/providers`: Provider management
- `/api/history`: Download history
- `/api/media/*`: Media library
- `/api/system/*`: System information

**Middleware Stack**:

1. Security headers
2. Authentication
3. CORS
4. Logging
5. Recovery

**WebSocket Support**:

- Real-time task updates
- Progress streaming
- Event notifications

### Media & Metadata Package (`pkg/metadata`)

#### Parsing Engine

\`\`\`go type MediaInfo struct { Title string Year int Season int Episode int
Type MediaType }

func ParseFileName(filename string) (\*MediaInfo, error) \`\`\`

#### API Clients

**TMDB Client**:

- Movie search and details
- TV show information
- Rate limiting

**OMDb Client**:

- IMDB ID lookup
- Rating information
- Episode data

### Task Management (`pkg/tasks`)

#### Task System

\`\`\`go type Task struct { ID string Status Status Progress int Error error }

func Start(ctx context.Context, id string, fn TaskFunc) \*Task \`\`\`

#### Features

- Concurrent execution
- Progress reporting
- Error tracking
- WebSocket broadcasting

### Scheduler Package (`pkg/scheduler`)

#### Scheduling Options

- Interval-based execution
- Cron expressions
- Jitter support
- Max run limits

#### Built-in Jobs

- Database cleanup
- Metadata refresh
- Library scanning
- Subtitle upgrades

### Translation Services (`pkg/translator`)

#### Google Translate

- REST API client
- Batch translation
- Language detection

#### OpenAI Integration

- GPT-3.5/4 support
- Context preservation
- Custom prompts

#### gRPC Service

- Protocol buffer definitions
- Streaming support
- Load balancing

## Web UI Analysis

### Architecture

#### Technology Stack

- React 18 with hooks
- Material-UI components
- Vite build system
- React Router for navigation

#### Component Structure

\`\`\` App.jsx ├── Dashboard.jsx ├── MediaLibrary.jsx ├── Settings.jsx │ ├──
GeneralSettings.jsx │ ├── ProviderSettings.jsx │ └── NotificationSettings.jsx
├── System.jsx └── Setup.jsx \`\`\`

#### State Management

- Local component state
- API service layer
- Real-time WebSocket updates

#### Key Features

**Dashboard**:

- Scan progress monitoring
- Recent activity
- System statistics

**Media Library**:

- File browser interface
- Subtitle management
- Metadata editing

**Settings**:

- Provider configuration
- User management
- System preferences

## API Endpoints

### Authentication

\`\`\` POST /api/login - User login POST /api/logout - User logout GET
/api/user - Current user info POST /api/users - Create user GET /api/users -
List users PUT /api/users/:id - Update user DELETE /api/users/:id - Delete user
\`\`\`

### Subtitles

\`\`\` POST /api/scan - Start directory scan GET /api/scan/status - Scan
progress POST /api/search - Search subtitles POST /api/download - Download
subtitle POST /api/translate - Translate subtitle POST /api/sync - Synchronize
subtitle POST /api/extract - Extract embedded GET /api/history - Download
history \`\`\`

### Media Library

\`\`\` GET /api/media - List media items GET /api/media/:id - Get media details
PUT /api/media/:id - Update media POST /api/media/scan - Scan library GET
/api/media/browse - File browser \`\`\`

### Providers

\`\`\` GET /api/providers - List providers GET /api/providers/:id - Provider
details PUT /api/providers/:id - Update provider POST /api/providers/test - Test
provider \`\`\`

### System

\`\`\` GET /api/system/info - System information GET /api/system/logs -
Application logs GET /api/system/tasks - Running tasks POST /api/system/backup -
Create backup POST /api/system/restore - Restore backup \`\`\`

## Database Schema

### Core Tables

#### users

\`\`\`sql CREATE TABLE users ( id INTEGER PRIMARY KEY, username TEXT UNIQUE,
password TEXT, email TEXT, role TEXT, created_at TIMESTAMP ); \`\`\`

#### subtitles

\`\`\`sql CREATE TABLE subtitles ( id INTEGER PRIMARY KEY, file TEXT, video_file
TEXT, language TEXT, service TEXT, created_at TIMESTAMP ); \`\`\`

#### media_items

\`\`\`sql CREATE TABLE media_items ( id INTEGER PRIMARY KEY, path TEXT UNIQUE,
title TEXT, year INTEGER, media_type TEXT, tmdb_id TEXT, imdb_id TEXT, metadata
JSONB ); \`\`\`

#### download_history

\`\`\`sql CREATE TABLE download_history ( id INTEGER PRIMARY KEY, video_path
TEXT, subtitle_path TEXT, language TEXT, provider TEXT, score INTEGER,
downloaded_at TIMESTAMP ); \`\`\`

### Indexes

- media_items(path)
- media_items(tmdb_id)
- subtitles(video_file)
- download_history(video_path, language)

## Provider Implementations

### OpenSubtitles

- REST API v1 support
- Hash-based search
- User agent requirements

### Subscene

- HTML scraping
- CloudFlare bypass
- Multi-language support

### YIFY Subtitles

- Movie-focused
- Direct download links
- Quality indicators

### Addic7ed

- TV show specialization
- Version matching
- Registration required

### Internal Providers

- Embedded provider for local files
- Folder provider for network shares
- Database provider for cached results

## Integration Points

### Sonarr Integration

- Episode file monitoring
- Webhook support
- API synchronization

### Radarr Integration

- Movie file monitoring
- Webhook support
- API synchronization

### Plex Integration

- Library synchronization
- Metadata refresh triggers
- Subtitle availability updates

### Notification Systems

- Discord webhooks
- Telegram bot API
- SMTP email
- Custom webhooks

## Performance Characteristics

### Bottlenecks

1. Metadata API rate limits
2. Provider API throttling
3. File system I/O for large libraries
4. Database queries without indexes

### Optimizations

1. Concurrent provider queries
2. In-memory caching
3. Batch database operations
4. Connection pooling

### Scalability

- Horizontal scaling via multiple instances
- PostgreSQL for larger deployments
- Redis for distributed caching
- Load balancer ready

## Security Model

### Authentication Layers

1. Session cookies (HTTP only)
2. API keys (header-based)
3. OAuth2 (GitHub)

### Authorization

- RBAC with three permission levels
- Per-endpoint access control
- Resource-level permissions

### Security Headers

- Content Security Policy
- X-Frame-Options
- X-Content-Type-Options
- Strict-Transport-Security

### Input Validation

- Path traversal prevention
- SQL injection protection
- XSS mitigation
- CSRF tokens

## Error Handling

### Error Types

\`\`\`go type SubtitleError struct { Code string Message string Err error }
\`\`\`

### Recovery Mechanisms

- Panic recovery in HTTP handlers
- Graceful provider fallback
- Transaction rollback
- Error logging and alerting

## Testing Infrastructure

### Unit Tests

- 70% code coverage
- Table-driven tests
- Mock interfaces
- Parallel execution

### Integration Tests

- Database operations
- API endpoints
- Provider functionality
- File operations

### E2E Tests

- Playwright for UI
- Full workflow coverage
- Cross-browser support

## Build & Deployment

### Build Process

1. Generate web assets
2. Embed static files
3. Compile Go binary
4. Create Docker image

### Deployment Options

- Standalone binary
- Docker container
- Kubernetes deployment
- Systemd service

### Configuration Management

- Environment variables
- Configuration file
- Command-line flags
- Runtime updates

## Monitoring & Observability

### Logging

- Structured logging with logrus
- Component-based log levels
- File and console output
- Log rotation

### Metrics

- Download statistics
- Provider performance
- API latency
- Error rates

### Health Checks

- Database connectivity
- Provider availability
- Disk space
- Memory usage

## Maintenance Operations

### Database Maintenance

- Automatic cleanup of old sessions
- Orphaned subtitle removal
- Index optimization
- Backup scheduling

### Cache Management

- TTL-based expiration
- Memory limits
- Cache warming
- Invalidation strategies

This completes the comprehensive technical breakdown of every component in the
Subtitle Manager system.
