# Subtitle Manager

Subtitle Manager is a comprehensive subtitle management application written in
Go that provides both CLI and web interfaces for converting, translating, and
managing subtitle files. **The project is nearing feature completion with full
Bazarr parity planned. It already includes 40+ subtitle providers, enterprise
authentication, PostgreSQL support, webhook and notification systems,
anti-captcha integration, and a polished React web interface.**

## ✨ Key Highlights

- 🎯 **Production Ready**: Complete with authentication, RBAC, and 40+ subtitle
  providers
- 🌐 **Full Web UI**: React-based interface with complete dashboard, settings,
  history, and system pages
- 🔐 **Enterprise Auth**: Password, OAuth2, API keys, and role-based access
  control
- 🚀 **High Performance**: Concurrent processing with worker pools and gRPC
  support
- 📦 **Container Ready**: Docker images published to GitHub Container Registry
- ✅ **Bazarr Parity**: Full feature compatibility with all major subtitle
  providers
- 🔄 **Enterprise Features**: PostgreSQL, webhooks, notifications, anti-captcha,
  and advanced scheduling
- 🗄️ **Automated Issue Management**: Workflow archives processed files and opens
  pull requests.

## Features

- Convert subtitles from many formats to SRT.
- Merge two subtitle tracks sorted by start time.
- Translate subtitles via Google Translate or ChatGPT APIs.
- Store translation history in SQLite, PebbleDB or PostgreSQL databases.
  Retrieve history via the `history` command or `/api/history` endpoint with
  optional `lang` and `video` filters.
- **Optional cloud storage for subtitles and history** in Amazon S3, Azure Blob
  Storage, or Google Cloud Storage with local backup support.
- Per component logging with adjustable levels.
- Extract subtitles from media containers using ffmpeg.
- Convert uploaded subtitle files to SRT via `/api/convert`.
- Transcribe audio tracks to subtitles via Whisper.
- Automatic subtitle synchronization using audio transcription and embedded
  tracks with advanced options for track selection, weighted averaging, and
  translation integration.
- World-class synchronization algorithm with adjustable CPU vs accuracy slider
  and optional longer processing for improved precision.
- Advanced dual-subtitle alignment for languages with different grammar
  structures.
- Experimental minimum display time mode that delays subsequent subtitles and
  catches up during silent gaps.
- Synchronize and translate subtitles to any language during the sync process.
- Download subtitles from a comprehensive list of providers based on Bazarr,
  including Addic7ed, AnimeKalesi, Animetosho, Assrt, Avistaz, BetaSeries,
  BSplayer, GreekSubs, Podnapisi, Subscene, TVSubtitles, Titlovi, LegendasDivx
  and many more.
- Batch translate multiple files concurrently.
- Mass synchronize subtitles across entire libraries.
- Monitor directories and automatically download subtitles.
- Scan existing libraries and fetch missing or upgraded subtitles, replacing
  only when the new file is larger.
- Download individual subtitles through the web API at `/api/download`.
- Schedule scans with the `autoscan` command using intervals or cron
  expressions.
- Parse file names and retrieve movie or episode details from TheMovieDB with
  language and rating data from OMDb.
- High performance scanning using concurrent workers.
- Recursive directory watching with -r flag.
- Integrate with Sonarr, Radarr and Plex using dedicated commands.
- Run a translation gRPC server.
- Translate uploaded subtitles through `/api/translate` endpoint.
- Delete subtitle files and remove history records.
- Track subtitle download history and list with `downloads` command or
  `/api/history` using `lang` and `video` filters.
- GitHub OAuth2 login enabled via `/api/oauth/github` endpoints.
- Manually search for subtitles with `search` command.
- Provider registry simplifies adding new sources.
- Dockerfile and workflow for container builds.
- Prebuilt images published to GitHub Container Registry.
- Integrated authentication system with password, token, OAuth2 and API key
  support.
- Generate one time login tokens using `user token` and authenticate with
  `login-token`.
- Minimal React web UI with login page and file upload forms for conversion and
  translation.
- Role based access control with sensible defaults and session storage in the
  database.
- Manage accounts with `user add`, `user role`, `user token` and `user list`
  commands.
- **Universal Tagging System**: Unified tagging interface supporting all entity
  types (media, users, providers, language profiles) with consistent API and
  advanced filtering capabilities.

### Universal Tagging System

Subtitle Manager features a comprehensive tagging system that provides unified
organization and filtering capabilities across all entity types:

#### Supported Entity Types

- **Media Items**: Tag movies, TV series, seasons, and episodes for content
  organization
- **Users**: Apply preference tags for personalized subtitle selection
- **Providers**: Tag provider instances for selection logic and priority
  management
- **Language Profiles**: Create tagged language preference groups
- **Media Profiles**: Content-specific tagging for automated workflows

#### Key Features

- **Polymorphic Design**: Single interface for all entity types with consistent
  API endpoints
- **Advanced Filtering**: Search and filter any entity by tags with complex
  query support
- **Bulk Operations**: Apply or remove tags from multiple entities
  simultaneously
- **Tag Hierarchies**: Support for tag types (system, user, custom) with
  optional color coding
- **Provider Integration**: Tag-based provider selection and priority weighting
- **Migration Support**: Seamless migration from legacy tagging implementations

#### API Endpoints

The tagging system exposes consistent REST endpoints for all entity types:

\``` GET /api/{entityType}/{id}/tags - List tags for entity POST
/api/{entityType}/{id}/tags - Add tag to entity DELETE
/api/{entityType}/{id}/tags/{tagId} - Remove tag from entity GET
/api/tags?entity_type={type} - List tags by entity type POST /api/tags/bulk -
Bulk tag operations \```

#### Usage Examples

\```bash

# Tag a movie for family content

POST /api/media/12345/tags {"tag_id": "family"}

# List all users with premium tags

GET /api/tags?entity_type=user&name=premium

# Bulk tag multiple providers

POST /api/tags/bulk {"tag_id": "reliable", "entities": [{"type": "provider",
"id": "opensubtitles"}, {"type": "provider", "id": "subscene"}]} \```

### Language Profiles System

Language Profiles provide Bazarr-compatible language preference management,
allowing fine-grained control over subtitle language selection for different
media items. This system replaces simple language strings with sophisticated
profile-based configurations.

#### Key Features

- **Priority-Based Language Selection**: Define multiple languages with priority
  ordering (lower numbers = higher priority)
- **Profile Assignment**: Assign different language profiles to specific media
  items or use global defaults
- **Forced Subtitles Support**: Configure preference for forced/hearing-impaired
  subtitles per language
- **Quality Thresholds**: Set cutoff scores for subtitle quality acceptance
- **Default Profile Management**: Automatic fallback to default profiles when no
  specific assignment exists
- **Database Integration**: Full CRUD operations with SQLite and PostgreSQL
  support

#### Profile Structure

```json
{
  "id": "profile-uuid",
  "name": "English + Spanish",
  "languages": [
    {
      "language": "en",
      "priority": 1,
      "forced": false,
      "hi": false
    },
    {
      "language": "es",
      "priority": 2,
      "forced": false,
      "hi": true
    }
  ],
  "cutoff_score": 80,
  "is_default": false
}
```

#### API Endpoints

- `GET /api/language-profiles` - List all language profiles
- `POST /api/language-profiles` - Create a new language profile
- `GET /api/language-profiles/{id}` - Get specific language profile
- `PUT /api/language-profiles/{id}` - Update language profile
- `DELETE /api/language-profiles/{id}` - Delete language profile
- `GET /api/language-profiles/default` - Get default language profile
- `POST /api/media/{id}/profile` - Assign profile to media item
- `GET /api/media/{id}/profile` - Get media's assigned profile

#### CLI Integration

```bash
# Fetch subtitles using language profile preferences
subtitle-manager fetch --profile /path/to/movie.mkv dummy output.srt

# Scan directory with profile-based language selection
subtitle-manager scan --profiles /media/movies
```

#### Automatic Profile Usage

The system automatically uses language profiles in these scenarios:

1. **Media Assignment**: When a media item has an assigned profile, that
   profile's language preferences are used for subtitle searches
2. **Default Fallback**: Media without assigned profiles use the default profile
3. **Priority Processing**: Languages are tried in priority order until suitable
   subtitles are found
4. **Provider Integration**: All 40+ subtitle providers support profile-based
   language selection

### Supported Subtitle Providers

Subtitle Manager now supports the full provider list from Bazarr. The following
services are available:

- Addic7ed
- AnimeKalesi
- Animetosho
- Assrt
- AvistaZ / CinemaZ
- BetaSeries
- BSplayer
- Embedded Subtitles
- Gestdown.info
- GreekSubs
- GreekSubtitles
- HDBits.org
- Hosszupuska
- Karagarga.in
- Ktuvit
- LegendasDivx
- Legendas.net
- Napiprojekt
- Napisy24
- Nekur
- OpenSubtitles.com
- OpenSubtitles.org (VIP)
- Podnapisi
- RegieLive
- Sous-Titres.eu
- Subdivx
- subf2m.co
- Subs.sab.bz
- Subs4Free
- Subs4Series
- Subscene
- Subscenter
- Subsunacs.net
- SubSynchro
- Subtitrari-noi.ro
- subtitri.id.lv
- Subtitulamos.tv
- Supersubtitles
- Titlovi
- Titrari.ro
- Titulky.com
- Turkcealtyazi.org
- TuSubtitulo
- TVSubtitles
- Whisper (optional local service)
- Wizdom
- XSubs
- Yavka.net
- YIFY Subtitles
- Zimuku

## Current Status

**Subtitle Manager backend is mostly complete** with full production readiness
achieved. Tagging enhancements are ongoing, and Whisper transcription can now
run via an optional local service.

### ✅ Completed (Production Ready Backend)

- **Core Functionality**: All CLI commands fully implemented
- **Web Interface**: Complete React UI with all major pages (Dashboard,
  Settings, Extract, History, System, Wanted)
- **Authentication**: Full enterprise-grade auth with OAuth2, API keys, RBAC
- **Providers**: 40+ subtitle providers with full Bazarr parity
- **APIs**: Complete REST API coverage for all operations
- **Infrastructure**: Docker support, CI/CD, automated testing
- **Database**: SQLite, PebbleDB and PostgreSQL backends
- **Enterprise Features**: Webhooks, notifications, anti-captcha, advanced
  scheduling

### 🔄 High Priority UI/UX Improvements Needed

- **Navigation Issues**: Fix user management display, implement back button, add
  sidebar pinning
- **Settings Enhancements**: Bazarr-compatible general settings, improved
  database management, card-based authentication
- **Provider System**: Fix configuration modals, implement global language
  settings
- **User Experience**: Reorganize navigation, create Tools section, add
  Languages page

### 🎯 Optional Remaining Features

- Advanced reverse proxy base URL support
- ~~Automatic subtitle synchronization using audio and embedded tracks with
  selectable audio and subtitle stream indices~~ ✅ **COMPLETED**

The backend provides full production functionality with feature parity to Bazarr
for all core subtitle management operations, plus additional enterprise
features. The frontend requires UI/UX improvements to match the backend's
quality and completeness.

## Installation

### Using Make (Recommended)

\```bash

# Clone repository

git clone <this repository> cd subtitle-manager

# Build everything with one command

make build

# Or for development

make quick-build

# See all available targets

make help \```

### Manual Installation

\```bash

# Clone repository

git clone <this repository> cd subtitle-manager

# Install dependencies, build web UI and compile

go generate ./webui go build \```

### Build Automation

The project includes a comprehensive Makefile that automates all build, test,
and deployment tasks:

- **`make build`** - Complete build including web UI and Go binary
- **`make quick-build`** - Fast build for development
- **`make test-all`** - Run all tests (Go + Web UI)
- **`make docker`** - Build Docker image
- **`make fix-webui`** - Fix web UI dependency conflicts
- **`make dev`** - Build and run in development mode
- **`make help`** - Show all available targets

The Makefile handles dependency resolution, web UI building, Go compilation,
testing, Docker builds, and more.

### Database Backend Build Options

Subtitle Manager supports multiple database backends that can be selected at
build time:

#### Pure Go Build (Default)

```bash
# Build without CGO dependencies (uses PebbleDB)
go build -tags nosqlite .
```

- **No CGO required** - Fully portable binary
- **PebbleDB backend** - High-performance embedded key-value store
- **All features supported** - Authentication, sessions, tags, permissions, etc.
- **Smaller binary size** - No SQLite dependencies

#### SQLite Build (CGO Required)

```bash
# Build with SQLite support (requires CGO)
go build -tags sqlite .
```

- **CGO required** - Needs C compiler and SQLite libraries
- **SQLite backend** - Traditional SQL database
- **Full SQL querying** - Standard SQL operations available
- **Migration support** - Can migrate from existing SQLite databases

#### Build Tag Summary

- **No tags or `-tags nosqlite`**: Pure Go build with PebbleDB (recommended)
- **`-tags sqlite`**: CGO build with SQLite support
- **Tests**: Run `go test -tags nosqlite` or `go test -tags sqlite` respectively

Both backends provide identical functionality for all user-facing features
including authentication, session management, API keys, dashboard preferences,
tags, permissions, and subtitle history.

## Usage

\``` subtitle-manager convert [input] [output] subtitle-manager merge [sub1]
[sub2] [output] subtitle-manager translate [input] [output] [lang]
subtitle-manager sync [media] [subtitle] [output] [--use-audio] [--use-embedded]
[--translate] subtitle-manager history subtitle-manager extract [media] [output]
subtitle-manager transcribe [media] [output] [lang] subtitle-manager fetch
[media] [lang] [output] subtitle-manager fetch --tags tag1,tag2 [media] [lang]
[output] subtitle-manager search [media] [lang] subtitle-manager batch [lang]
[files...] subtitle-manager syncbatch -config file.json

# syncbatch expects a JSON file describing media and subtitle pairs

subtitle-manager scan [directory] [lang] [-u] subtitle-manager autoscan
[directory] [lang] [-i duration] [-s cron] [-u] subtitle-manager scanlib
[directory] subtitle-manager watch [directory] [lang] [-r] subtitle-manager
grpc-server --addr :50051 subtitle-manager grpc-set-config --addr :50051 --key
google_api_key --value NEWKEY subtitle-manager metadata search [query]
subtitle-manager metadata update [file] [--title T] [--release-group G] [--alt
"A,B"] [--lock fields] subtitle-manager delete [file] subtitle-manager rename
[video] [lang] subtitle-manager downloads subtitle-manager login [username]
[password] subtitle-manager login-token [token] subtitle-manager user add
[username] [email] [password] subtitle-manager user apikey [username]
subtitle-manager user token [email] subtitle-manager user role [username] [role]
subtitle-manager user list \```

The `extract` command accepts `--ffmpeg` to specify a custom ffmpeg binary.

### Web UI

Run `subtitle-manager web` to start the embedded React interface on `:8080`. The
SPA is built via `go generate` in the `webui` directory and embedded using Go
1.16's `embed` package.

The interface provides a complete user experience with:

- **Dashboard**: Library scanning with progress tracking and provider selection
- **Settings**: Configuration management that mirrors Bazarr's options, with
  values retrieved from `/api/config` and saved via POST to the same endpoint
- **Extract**: Subtitle extraction from media files using ffmpeg
- **History**: Complete translation and download history with language filtering
- **System**: Real-time logs, task progress, and system information
- **Wanted**: Search interface for missing subtitles with provider selection
- **Tags**: Edit and organize custom tags for language preferences
- **Config File**: Edit the YAML configuration directly in the browser

Configuration values are loaded from `$HOME/.subtitle-manager.yaml` by default.
Environment variables prefixed with `SM_` override configuration values. API
keys may be specified via flags `--google-key`, `--openai-key`,
`--opensubtitles-key`, `--tmdb-key`, and `--omdb-key` or in the configuration
file. Additional flags include `--ffmpeg-path` for a custom ffmpeg binary,
`--batch-workers` and `--scan-workers` to control concurrency.

#### Database Configuration

The database location and backend can be configured with:

- **`--db`** - Database file path (SQLite) or directory path (PebbleDB)
  - Default: `$HOME/.subtitle-manager.db` (SQLite) or
    `$HOME/.subtitle-manager-data/` (PebbleDB)
- **`--db-backend`** - Choose between `sqlite`, `pebble`, or `postgres` engines
  - Default: `pebble` (pure Go builds), `sqlite` (CGO builds)

**Database Backend Selection:**

- **PebbleDB** (`pebble`): Default for pure Go builds, no CGO required, high
  performance
- **SQLite** (`sqlite`): Traditional SQL database, requires CGO, migration
  support
- **PostgreSQL** (`postgres`): External database server, requires connection
  string

When using PebbleDB, a directory path should be supplied instead of a file. To
migrate existing SQLite databases to PebbleDB, run:
`subtitle-manager migrate old.db newdir`

Translation can be delegated to a remote gRPC server using the `--grpc` flag and
providing an address such as `localhost:50051`. Generic provider options may
also be set with variables like `SM_PROVIDERS_GENERIC_API_URL`. For WebSocket
security, use `--allowed-websocket-origins` or `SM_ALLOWED_WEBSOCKET_ORIGINS` to
specify comma-separated allowed origins (by default only localhost and
same-origin connections are allowed). Run
`subtitle-manager migrate old.db newdir` to copy existing subtitle history from
SQLite to PebbleDB.

### REST API

The web server exposes a comprehensive REST API for all subtitle operations:

#### Authentication Endpoints

- `POST /api/login` - User authentication with username/password
- `POST /api/setup` - Initial system configuration and admin user creation
- `GET /api/setup/status` - Check if initial setup is required
- `POST /api/oauth/github/login` - GitHub OAuth2 login initiation
- `GET /api/oauth/github/callback` - OAuth2 callback handler

#### Configuration Management

- `GET /api/config` - Retrieve current configuration as JSON
- `POST /api/config` - Update configuration values

#### Subtitle Operations

- `POST /api/convert` - Convert uploaded subtitle files to SRT format
- `POST /api/translate` - Translate uploaded subtitle files
- `POST /api/sync/batch` - Synchronize multiple subtitle files in one request
- `POST /api/extract` - Extract subtitles from media files using ffmpeg
- `POST /api/download` - Download subtitles for specific media files

-#### Library Management

- `POST /api/library/scan` - Scan directories and generate metadata
- `GET /api/library/scan/status` - Check library scan progress
- `GET /api/search` - Search for subtitles with provider and language filters
- `GET /api/wanted` - Retrieve wanted subtitles list
- `POST /api/wanted` - Add subtitles to wanted list

#### History and Monitoring

- `GET /api/history` - Retrieve translation and download history. Supports
  `lang` and `video` query parameters for filtering.
- `GET /api/logs` - Get recent log entries
- `GET /api/system` - System information (Go version, OS, architecture,
  goroutines)
- `GET /api/tasks` - Current task status and progress
- `GET /api/users` - List all users (admin only)
- `POST /api/users/{id}/reset` - Reset a user's password and email credentials

All endpoints require authentication via session cookies or API keys using the
`X-API-Key` header. Role-based access control is enforced with three permission
levels: `read`, `basic`, and `admin`.

Example configuration:

\``` log-level: info log_levels: translate: debug translate_service: google
google_api_key: your-google-api-key openai_api_key: your-openai-api-key
google_api_url: https://translation.googleapis.com/language/translate/v2
openai_model: gpt-3.5-turbo ffmpeg_path: /usr/bin/ffmpeg batch_workers: 4
scan_workers: 4 opensubtitles: api_key: your-os-key api_url:
https://rest.opensubtitles.org user_agent: subtitle-manager/0.1 providers:
generic: api_url: https://example.com/subtitles username: myuser password:
secret api_key: token123 github_client_id: yourClientID github_client_secret:
yourClientSecret github_redirect_url:
http://localhost:8080/api/oauth/github/callback \```

### Docker

Subtitle Manager provides official Docker images with the web interface enabled
by default. These images include `ffmpeg` so subtitle extraction works without
additional dependencies. The binary is located at `/usr/bin/ffmpeg` and the
container sets `SM_FFMPEG_PATH` accordingly.

#### Quick Start

\```bash

# Simple start - web interface only

docker run -p 8080:8080 ghcr.io/jdfalk/subtitle-manager:latest

# Complete setup with persistent storage and media access

docker run -d \
 --name subtitle-manager \
 --restart unless-stopped \
 -p 8080:8080 \
 -v $(pwd)/config:/config \
 -v /path/to/your/movies:/media/movies:ro \
 -v /path/to/your/tv:/media/tv:ro \
 -v $(pwd)/subtitles:/subtitles \
 -e SM_LOG_LEVEL=info \
 ghcr.io/jdfalk/subtitle-manager:latest

# User-requested format with variable substitution

docker run -d \
 --name subtitle-manager \
 --restart unless-stopped \
 -p 8080:8080 \
 -v ${path_on_host_system_to_configdir}:/config \
 -v ${path_on_host_system_to_media}:/media:ro \
 ghcr.io/jdfalk/subtitle-manager:latest

# Example with actual paths (update paths to match your system)

docker run -d \
 --name subtitle-manager \
 --restart unless-stopped \
 -p 8080:8080 \
 -v /home/user/subtitle-manager/config:/config \
 -v /media/movies:/media/movies:ro \
 -v /media/tv:/media/tv:ro \
 -v /home/user/subtitle-manager/subtitles:/subtitles \
 ghcr.io/jdfalk/subtitle-manager:latest

# With local Whisper transcription service (requires Docker socket)

docker run -d \
 --name subtitle-manager \
 --restart unless-stopped \
 -p 8080:8080 \
 -v /home/user/subtitle-manager/config:/config \
 -v /media/movies:/media/movies:ro \
 -v /media/tv:/media/tv:ro \
 -v /var/run/docker.sock:/var/run/docker.sock \
 -e ENABLE_WHISPER=1 \
 ghcr.io/jdfalk/subtitle-manager:latest

<!-- Customize WHISPER_MODEL=base or WHISPER_DEVICE=cpu to disable GPU -->
<!-- SM_OPENAI_API_URL will be set automatically to http://localhost:9000/v1 -->
<!-- Add --gpus all if you want GPU acceleration (requires NVIDIA Container Toolkit) -->

````

Access the web interface at `http://localhost:8080`

#### Environment Variables

Configure Subtitle Manager using environment variables with the `SM_` prefix:

**Basic Configuration:**

- `SM_LOG_LEVEL` - Log level (debug, info, warn, error) - Default: `info`
- `SM_LOG_FILE` - Path to log file - Default: `/config/logs/subtitle-manager.log`
- `SM_CONFIG_FILE` - Path to configuration file - Default: `/config/subtitle-manager.yaml`
- `SM_DB_PATH` - Database file path - Default: `/config/subtitle-manager.db`
- `SM_DB_BACKEND` - Database backend (sqlite, pebble, postgres) - Default: `sqlite`

**API Keys:**

- `SM_GOOGLE_API_KEY` - Google Translate API key
- `SM_OPENAI_API_KEY` - OpenAI/ChatGPT API key
- `SM_OPENSUBTITLES_API_KEY` - OpenSubtitles API key
- `SM_TMDB_API_KEY` - TheMovieDB API key
- `SM_OMDB_API_KEY` - OMDb API key

**Performance Tuning:**

- `SM_BATCH_WORKERS` - Number of concurrent translation workers - Default: `4`
- `SM_SCAN_WORKERS` - Number of concurrent scanning workers - Default: `4`
- `SM_FFMPEG_PATH` - Path to ffmpeg binary - Default: `/usr/bin/ffmpeg`

**Provider Configuration:**

- `SM_PROVIDERS_GENERIC_API_URL` - Generic provider API URL
- `SM_PROVIDERS_GENERIC_USERNAME` - Generic provider username
- `SM_PROVIDERS_GENERIC_PASSWORD` - Generic provider password
- `SM_PROVIDERS_GENERIC_API_KEY` - Generic provider API key
- `ENABLE_WHISPER` - Launch local Whisper service when set to `1`
- `SM_PROVIDERS_WHISPER_API_URL` - Override Whisper service URL (default `http://localhost:9000`)
- `SM_OPENAI_API_URL` - Override OpenAI/Whisper API base URL (default `https://api.openai.com/v1`)

#### Whisper Service Requirements
When `ENABLE_WHISPER=1` is set, the container launches `onerahmet/openai-whisper-asr-webservice` with automatic retry logic and health checks for reliable startup. Mount the Docker socket and install the NVIDIA Container Toolkit if GPU acceleration is desired. Customize the service with these optional variables:
- `WHISPER_CONTAINER_NAME` - Container name (default `whisper-asr-service`)
- `WHISPER_IMAGE` - Docker image to use
- `WHISPER_PORT` - Port to expose (default `9000`)
- `WHISPER_MODEL` - Whisper model (base, small, medium, large)
- `WHISPER_DEVICE` - `cuda` or `cpu` to toggle GPU usage
- `WHISPER_HEALTH_TIMEOUT` - Health check timeout in seconds (default `10`, set to `0` to skip)

**GitHub OAuth (Optional):**

- `SM_GITHUB_CLIENT_ID` - GitHub OAuth client ID
- `SM_GITHUB_CLIENT_SECRET` - GitHub OAuth client secret
- `SM_GITHUB_REDIRECT_URL` - OAuth redirect URL

#### Volume Mounts

**Required Volumes:**

- `/config` - Configuration and database storage
- `/media` - Media library access (read-only recommended)

**Optional Volumes:**

- `/subtitles` - Custom subtitle storage location

#### Docker Compose

For easier management, use the provided Docker Compose configuration:

\```bash
# Download the compose file
curl -O https://raw.githubusercontent.com/jdfalk/subtitle-manager/main/docker-compose.yml

# Edit the volume paths in docker-compose.yml to match your setup
# Update /path/to/your/movies and /path/to/your/tv

# Start the service
docker-compose up -d

# View logs
docker-compose logs -f
# Log file stored in ./config/logs/subtitle-manager.log

# Stop the service
docker-compose down
\```

**Sample docker-compose.yml:**

\```yaml
version: "3.8"
services:
  subtitle-manager:
    image: ghcr.io/jdfalk/subtitle-manager:latest
    container_name: subtitle-manager
    restart: unless-stopped
    ports:
      - "8080:8080"
    volumes:
      - ./config:/config
      - /path/to/your/movies:/media/movies:ro
      - /path/to/your/tv:/media/tv:ro
    environment:
      - SM_LOG_LEVEL=info
      - SM_GOOGLE_API_KEY=your_api_key_here
\```

#### Production Deployment (Docker Swarm)

For production deployments with Docker Swarm:

\```bash
# Download the stack file
curl -O https://raw.githubusercontent.com/jdfalk/subtitle-manager/main/docker-stack.yml

# Edit volume paths and resource limits as needed

# Deploy the stack
docker stack deploy -c docker-stack.yml subtitle-manager

# Check service status
docker service ls
docker service logs subtitle-manager_subtitle-manager

# Update the service
docker service update --image ghcr.io/jdfalk/subtitle-manager:latest subtitle-manager_subtitle-manager

# Remove the stack
docker stack rm subtitle-manager
\```

#### Building Custom Images

Build a container image with your customizations:

\```bash
# Clone the repository
git clone https://github.com/jdfalk/subtitle-manager.git
cd subtitle-manager

# Build the image
docker build -t subtitle-manager-custom .

# Run your custom image
docker run -d -p 8080:8080 subtitle-manager-custom
\```

#### Initial Setup

On first run, Subtitle Manager will start with the web interface where you can:

1. Complete initial setup through the web UI at `http://localhost:8080`
2. Create admin user account
3. Configure API keys and provider settings
4. Set up library paths

All configuration will be persisted in the `/config` volume.

### Linux Service (systemd)

For Linux servers, Subtitle Manager can run as a native systemd service. This provides better system integration and resource management compared to Docker for production deployments.

```bash
# Quick start - see systemd/README.md for detailed instructions

# 1. Create system user and directories
sudo useradd --system --home-dir /var/lib/subtitle-manager --create-home --shell /bin/false subtitle-manager
sudo mkdir -p /etc/subtitle-manager /var/log/subtitle-manager

# 2. Install binary (from release or build from source)
sudo cp subtitle-manager /usr/local/bin/
sudo chmod +x /usr/local/bin/subtitle-manager

# 3. Install service files
sudo cp systemd/subtitle-manager.service /etc/systemd/system/
sudo cp systemd/subtitle-manager.env /etc/subtitle-manager/

# 4. Configure and start service
sudo systemctl daemon-reload
sudo systemctl enable --now subtitle-manager

# 5. Check status and logs
sudo systemctl status subtitle-manager
sudo journalctl -u subtitle-manager -f
```

**Key Benefits:**
- Native Linux service integration
- Automatic startup on boot
- System resource management
- Security hardening with systemd features
- Centralized logging with journald

See [systemd/README.md](systemd/README.md) for complete installation instructions, configuration options, and troubleshooting guide.

## Development

### Quick Start with Dev Container

The easiest way to start developing is using the provided VS Code development container:

1. Install [VS Code](https://code.visualstudio.com/) and the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
2. Clone the repository and open it in VS Code
3. When prompted, click "Reopen in Container"
4. The container will automatically set up Go, Node.js, FFmpeg, SQLite, and all development tools

See [.devcontainer/README.md](.devcontainer/README.md) for detailed dev container documentation.

### Local Development

Tests can be run with `go test ./...`.
PostgreSQL tests require a local PostgreSQL installation and will skip gracefully if unavailable.
Web UI unit tests live in `webui/src/__tests__` and are executed with `npm test` from the `webui` directory.
End-to-end tests use Playwright and run with `npm run test:e2e` once browsers are installed via `npx playwright install`.
Continuous integration is provided via a GitHub Actions workflow that verifies formatting, vets code and runs the test suite on each push.

### Git Hooks and Code Formatting

#### Automatic Formatting on GitHub

The repository includes an **automatic code formatting system** that runs on all pull requests:

- **Go code**: Automatically formatted with `gofmt -s -w .` and `goimports`
- **Frontend code**: Automatically formatted with `prettier --write`

When you open or update a pull request, GitHub Actions will:

1. Check if any code needs formatting
2. Apply formatting automatically if needed
3. Commit the changes back to your PR branch
4. Add a comment explaining what was formatted

This means **you don't need to worry about code formatting** - just focus on your code logic and let the automation handle the style!

#### Optional Local Pre-commit Hooks

For developers who prefer to format code locally before pushing, you have several options:

##### Modern Pre-commit Framework (Recommended)

The repository includes a `.pre-commit-config.yaml` configuration with Ruff for Python and Prettier for JavaScript/Markdown:

\```bash
# Install pre-commit and tools
pip install pre-commit ruff

# Install the pre-commit hooks
pre-commit install

# Run on all files (one-time setup)
pre-commit run --all-files
\```

This will automatically:
- **Python files**: Lint and format with Ruff (fast, modern replacement for flake8/black)
- **JS/MD/CSS files**: Format with Prettier using the existing webui configuration
- **All files**: Check for trailing whitespace, file endings, YAML/JSON syntax

##### Shell-based Pre-commit Hooks (Legacy)

You can also use the existing shell-based hooks:

\```bash
# Install the auto-formatting pre-commit hook
./scripts/install-pre-commit-hooks.sh

# Or install the legacy quality-check pre-commit hook
./scripts/install-hooks.sh
\```

The **auto-formatting hook** (`install-pre-commit-hooks.sh`) will:

- Format Go files with `gofmt -s` and `goimports`
- Format frontend files with `prettier`
- Automatically stage the formatted files

The **legacy quality-check hook** (`install-hooks.sh`) will:

- Check Go file formatting with `gofmt -s`
- Run `go vet` for static analysis
- Prevent commits that don't pass these checks

**Benefits of the modern pre-commit framework:**
- ✅ **Ruff for Python**: Much faster than traditional tools (10-100x faster than flake8)
- ✅ **Consistent formatting**: Uses the same Prettier configuration as CI
- ✅ **Extensible**: Easy to add new languages and tools
- ✅ **Smart**: Only runs on changed files by default

To bypass any hook temporarily, use `git commit --no-verify`.

### Issue updates

🚀 **New Distributed System**: We now use individual UUID-named files in `.github/issue-updates/` to eliminate merge conflicts! Use the helper script: `./scripts/create-issue-update.sh create "Title" "Body" "labels"`. See [Quick Start Guide](.github/ISSUE_UPDATES_QUICK_START.md) for details.

**Legacy Support**: Pushing an `issue_updates.json` file to the repository root still works for backward compatibility. The unified issue management workflow processes both formats.

Note: If the unified issue management workflow fails due to a missing `requirements.txt` file, update to the latest version of [ghcommon](https://github.com/jdfalk/ghcommon) which now provides this file. This resolves issues #1249 and #1251.

The new format uses individual files with this structure:
\```json
{
  "action": "create",
  "title": "Issue title",
  "body": "Issue description",
  "labels": ["enhancement"]
}
\```

Benefits of the new system:
- ✅ No merge conflicts - each update is in its own file
- ✅ Parallel development - multiple people can create updates simultaneously
- ✅ Atomic operations - each file represents a single issue action
- ✅ Better git history - changes are tracked individually

The workflow runs on every push to `main` and processes all updates from both the legacy file and the new directory structure.

### Duplicate ticket cleanup

The `close-duplicates` workflow runs daily and on demand to detect open issues
with the same title. The script chooses the lowest numbered ticket as the
canonical reference and automatically closes the rest with a comment noting the
duplicate. This keeps the issue tracker focused on a single discussion for each
problem.

### Regenerating Protobuf files

The gRPC service definitions are located in `proto/translator.proto`. If you
modify this file, regenerate the Go bindings before committing:

\```bash
# Install protobuf tools if missing
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate updated gRPC code
make proto-gen
\```

The generated files live in `pkg/translatorpb` and should be committed with your
changes.

The project is **mostly feature complete** with full Bazarr parity as the target. Remaining work focuses on Sonarr/Radarr sync improvements and the metadata editor. See `TODO.md` for details.
Extensive architectural details and design decisions are documented in `docs/TECHNICAL_DESIGN.md`. For a package-by-package function reference see `docs/COMPLETE_DESIGN.md`. New contributors should review these documents to understand package responsibilities and completed features.
For a detailed list of Bazarr features used as the parity target, see [docs/BAZARR_FEATURES.md](docs/BAZARR_FEATURES.md).
Instructions for importing an existing Bazarr configuration are documented in [docs/BAZARR_SETTINGS_SYNC.md](docs/BAZARR_SETTINGS_SYNC.md).
A high-level code overview is available in [docs/CODE_OVERVIEW.md](docs/CODE_OVERVIEW.md).
Protobuf regeneration steps are documented in [docs/PROTOBUF_REGEN.md](docs/PROTOBUF_REGEN.md).
Additional references include [docs/DEVELOPER_GUIDE.md](docs/DEVELOPER_GUIDE.md) for environment setup and [docs/API_DESIGN.md](docs/API_DESIGN.md) for REST and gRPC design notes.

## Security

Subtitle Manager applies strict security headers, including `Referrer-Policy: no-referrer`, and sanitizes HTML content in the web interface using DOMPurify. When contributing UI or API code, ensure all user-provided data is properly validated and sanitized.

## License

This project is licensed under the terms of the MIT license. See `LICENSE` for details.
````
