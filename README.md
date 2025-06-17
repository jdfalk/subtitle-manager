# Subtitle Manager

Subtitle Manager is a comprehensive subtitle management application written in Go that provides both CLI and web interfaces for converting, translating, and managing subtitle files. **The project has achieved ~99% backend completion with full Bazarr feature parity** including 40+ subtitle providers, complete authentication system, PostgreSQL support, webhook system, anti-captcha integration, and a React-based web interface that requires UI/UX enhancements.

## ✨ Key Highlights

- 🎯 **Production Ready**: Complete with authentication, RBAC, and 40+ subtitle providers
- 🌐 **Full Web UI**: React-based interface with complete dashboard, settings, history, and system pages
- 🔐 **Enterprise Auth**: Password, OAuth2, API keys, and role-based access control
- 🚀 **High Performance**: Concurrent processing with worker pools and gRPC support
- 📦 **Container Ready**: Docker images published to GitHub Container Registry
- ✅ **Bazarr Parity**: Full feature compatibility with all major subtitle providers
- 🔄 **Enterprise Features**: PostgreSQL, webhooks, notifications, anti-captcha, and advanced scheduling

## Features

- Convert subtitles from many formats to SRT.
- Merge two subtitle tracks sorted by start time.
- Translate subtitles via Google Translate or ChatGPT APIs.
- Store translation history in SQLite, PebbleDB or PostgreSQL databases. Retrieve history via the `history` command or `/api/history` endpoint.
- Per component logging with adjustable levels.
- Extract subtitles from media containers using ffmpeg.
- Convert uploaded subtitle files to SRT via `/api/convert`.
- Transcribe audio tracks to subtitles via Whisper.
- Automatic subtitle synchronization using audio transcription and embedded tracks with advanced options for track selection, weighted averaging, and translation integration.
- Synchronize and translate subtitles to any language during the sync process.
- Download subtitles from a comprehensive list of providers based on Bazarr,
  including Addic7ed, AnimeKalesi, Animetosho, Assrt, Avistaz, BetaSeries,
  BSplayer, GreekSubs, Podnapisi, Subscene, TVSubtitles, Titlovi, LegendasDivx
  and many more.
- Batch translate multiple files concurrently.
- Monitor directories and automatically download subtitles.
- Scan existing libraries and fetch missing or upgraded subtitles.
- Download individual subtitles through the web API at `/api/download`.
- Schedule scans with the `autoscan` command using intervals or cron expressions.
- Parse file names and retrieve movie or episode details from TheMovieDB.
- High performance scanning using concurrent workers.
- Recursive directory watching with -r flag.
- Integrate with Sonarr, Radarr and Plex using dedicated commands.
- Run a translation gRPC server.
- Translate uploaded subtitles through `/api/translate` endpoint.
- Delete subtitle files and remove history records.
- Track subtitle download history and list with `downloads` command or `/api/history`.
- GitHub OAuth2 login enabled via `/api/oauth/github` endpoints.
- Manually search for subtitles with `search` command.
- Provider registry simplifies adding new sources.
- Dockerfile and workflow for container builds.
- Prebuilt images published to GitHub Container Registry.
- Integrated authentication system with password, token, OAuth2 and API key support.
- Generate one time login tokens using `user token` and authenticate with `login-token`.
- Minimal React web UI with login page and file upload forms for conversion and translation.
- Role based access control with sensible defaults and session storage in the database.
- Manage accounts with `user add`, `user role`, `user token` and `user list` commands.

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
- Whisper (requires external web service)
- Wizdom
- XSubs
- Yavka.net
- YIFY Subtitles
- Zimuku

## Current Status

**Subtitle Manager backend is ~100% complete** with full production readiness, but the frontend requires significant UI/UX improvements for optimal user experience.

### ✅ Completed (Production Ready Backend)

- **Core Functionality**: All CLI commands fully implemented
- **Web Interface**: Complete React UI with all major pages (Dashboard, Settings, Extract, History, System, Wanted)
- **Authentication**: Full enterprise-grade auth with OAuth2, API keys, RBAC
- **Providers**: 40+ subtitle providers with full Bazarr parity
- **APIs**: Complete REST API coverage for all operations
- **Infrastructure**: Docker support, CI/CD, automated testing
- **Database**: SQLite, PebbleDB and PostgreSQL backends
- **Enterprise Features**: Webhooks, notifications, anti-captcha, advanced scheduling

### 🔄 High Priority UI/UX Improvements Needed

- **Navigation Issues**: Fix user management display, implement back button, add sidebar pinning
- **Settings Enhancements**: Bazarr-compatible general settings, improved database management, card-based authentication
- **Provider System**: Fix configuration modals, implement global language settings
- **User Experience**: Reorganize navigation, create Tools section, add Languages page

### 🎯 Optional Remaining Features

- Advanced reverse proxy base URL support
- ~~Automatic subtitle synchronization using audio and embedded tracks with selectable audio and subtitle stream indices~~ ✅ **COMPLETED**

The backend provides full production functionality with feature parity to Bazarr for all core subtitle management operations, plus additional enterprise features. The frontend requires UI/UX improvements to match the backend's quality and completeness.

## Installation

### Using Make (Recommended)

```bash
# Clone repository
git clone <this repository>
cd subtitle-manager

# Build everything with one command
make build

# Or for development
make quick-build

# See all available targets
make help
```

### Manual Installation

```bash
# Clone repository
git clone <this repository>
cd subtitle-manager

# Install dependencies, build web UI and compile
go generate ./webui
go build
```

### Build Automation

The project includes a comprehensive Makefile that automates all build, test, and deployment tasks:

- **`make build`** - Complete build including web UI and Go binary
- **`make quick-build`** - Fast build for development
- **`make test-all`** - Run all tests (Go + Web UI)
- **`make docker`** - Build Docker image
- **`make fix-webui`** - Fix web UI dependency conflicts
- **`make dev`** - Build and run in development mode
- **`make help`** - Show all available targets

The Makefile handles dependency resolution, web UI building, Go compilation, testing, Docker builds, and more.

## Usage

```
subtitle-manager convert [input] [output]
subtitle-manager merge [sub1] [sub2] [output]
subtitle-manager translate [input] [output] [lang]
subtitle-manager sync [media] [subtitle] [output] [--use-audio] [--use-embedded] [--translate]
subtitle-manager history
subtitle-manager extract [media] [output]
subtitle-manager transcribe [media] [output] [lang]
subtitle-manager fetch opensubtitles [media] [lang] [output]
subtitle-manager fetch subscene [media] [lang] [output]
subtitle-manager search opensubtitles [media] [lang]
subtitle-manager batch [lang] [files...]
subtitle-manager scan opensubtitles [directory] [lang] [-u]
subtitle-manager scan subscene [directory] [lang] [-u]
subtitle-manager autoscan [provider] [directory] [lang] [-i duration] [-s cron] [-u]
subtitle-manager scanlib [directory]
subtitle-manager watch opensubtitles [directory] [lang] [-r]
subtitle-manager watch subscene [directory] [lang] [-r]
subtitle-manager grpc-server --addr :50051
subtitle-manager delete [file]
subtitle-manager downloads
subtitle-manager login [username] [password]
subtitle-manager login-token [token]
subtitle-manager user add [username] [email] [password]
subtitle-manager user apikey [username]
subtitle-manager user token [email]
subtitle-manager user role [username] [role]
subtitle-manager user list
```

The `extract` command accepts `--ffmpeg` to specify a custom ffmpeg binary.

### Web UI

Run `subtitle-manager web` to start the embedded React interface on `:8080`. The SPA is built via `go generate` in the `webui` directory and embedded using Go 1.16's `embed` package.

The interface provides a complete user experience with:

- **Dashboard**: Library scanning with progress tracking and provider selection
- **Settings**: Configuration management that mirrors Bazarr's options, with values retrieved from `/api/config` and saved via POST to the same endpoint
- **Extract**: Subtitle extraction from media files using ffmpeg
- **History**: Complete translation and download history with language filtering
- **System**: Real-time logs, task progress, and system information
- **Wanted**: Search interface for missing subtitles with provider selection

Configuration values are loaded from `$HOME/.subtitle-manager.yaml` by default. Environment variables prefixed with `SM_` override configuration values. API keys may be specified via flags `--google-key`, `--openai-key` and `--opensubtitles-key` or in the configuration file. Additional flags include `--ffmpeg-path` for a custom ffmpeg binary, `--batch-workers` and `--scan-workers` to control concurrency. The SQLite database location defaults to `$HOME/.subtitle-manager.db` and can be overridden with `--db`. Use `--db-backend` to switch between the `sqlite` and `pebble` engines. When using PebbleDB a directory path may be supplied instead of a file. Translation can be delegated to a remote gRPC server using the `--grpc` flag and providing an address such as `localhost:50051`. Generic provider options may also be set with variables like `SM_PROVIDERS_GENERIC_API_URL`.
Run `subtitle-manager migrate old.db newdir` to copy existing subtitle history from SQLite to PebbleDB.

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
- `POST /api/extract` - Extract subtitles from media files using ffmpeg
- `POST /api/download` - Download subtitles for specific media files

#### Library Management

- `POST /api/scan` - Start library scanning with provider selection
- `GET /api/scan/status` - Check scan progress and status
- `GET /api/search` - Search for subtitles with provider and language filters
- `GET /api/wanted` - Retrieve wanted subtitles list
- `POST /api/wanted` - Add subtitles to wanted list

#### History and Monitoring

- `GET /api/history` - Retrieve translation and download history
- `GET /api/logs` - Get recent log entries
- `GET /api/system` - System information (Go version, OS, architecture, goroutines)
- `GET /api/tasks` - Current task status and progress
- `GET /api/users` - List all users (admin only)
- `POST /api/users/{id}/reset` - Reset a user's password and email credentials

All endpoints require authentication via session cookies or API keys using the `X-API-Key` header. Role-based access control is enforced with three permission levels: `read`, `basic`, and `admin`.

Example configuration:

```
log-level: info
log_levels:
  translate: debug
translate_service: google
google_api_key: your-google-api-key
openai_api_key: your-openai-api-key
google_api_url: https://translation.googleapis.com/language/translate/v2
openai_model: gpt-3.5-turbo
ffmpeg_path: /usr/bin/ffmpeg
batch_workers: 4
scan_workers: 4
opensubtitles:
  api_key: your-os-key
  api_url: https://rest.opensubtitles.org
  user_agent: subtitle-manager/0.1
providers:
  generic:
    api_url: https://example.com/subtitles
    username: myuser
    password: secret
    api_key: token123
github_client_id: yourClientID
github_client_secret: yourClientSecret
github_redirect_url: http://localhost:8080/api/oauth/github/callback
```

### Docker

Subtitle Manager provides official Docker images with the web interface enabled by default.
These images include `ffmpeg` so subtitle extraction works without additional
dependencies. The binary is located at `/usr/bin/ffmpeg` and the container sets
`SM_FFMPEG_PATH` accordingly.

#### Quick Start

```bash
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
```

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

**Performance Tuning:**

- `SM_BATCH_WORKERS` - Number of concurrent translation workers - Default: `4`
- `SM_SCAN_WORKERS` - Number of concurrent scanning workers - Default: `4`
- `SM_FFMPEG_PATH` - Path to ffmpeg binary - Default: `/usr/bin/ffmpeg`

**Provider Configuration:**

- `SM_PROVIDERS_GENERIC_API_URL` - Generic provider API URL
- `SM_PROVIDERS_GENERIC_USERNAME` - Generic provider username
- `SM_PROVIDERS_GENERIC_PASSWORD` - Generic provider password
- `SM_PROVIDERS_GENERIC_API_KEY` - Generic provider API key

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

```bash
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
```

**Sample docker-compose.yml:**

```yaml
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
```

#### Production Deployment (Docker Swarm)

For production deployments with Docker Swarm:

```bash
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
```

#### Building Custom Images

Build a container image with your customizations:

```bash
# Clone the repository
git clone https://github.com/jdfalk/subtitle-manager.git
cd subtitle-manager

# Build the image
docker build -t subtitle-manager-custom .

# Run your custom image
docker run -d -p 8080:8080 subtitle-manager-custom
```

#### Initial Setup

On first run, Subtitle Manager will start with the web interface where you can:

1. Complete initial setup through the web UI at `http://localhost:8080`
2. Create admin user account
3. Configure API keys and provider settings
4. Set up library paths

All configuration will be persisted in the `/config` volume.

## Development

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

For developers who prefer to format code locally before pushing, you can install pre-commit hooks:

```bash
# Install the auto-formatting pre-commit hook
./scripts/install-pre-commit-hooks.sh

# Or install the legacy quality-check pre-commit hook
./scripts/install-hooks.sh
```

The **auto-formatting hook** (`install-pre-commit-hooks.sh`) will:

- Format Go files with `gofmt -s` and `goimports`
- Format frontend files with `prettier`
- Automatically stage the formatted files

The **legacy quality-check hook** (`install-hooks.sh`) will:

- Check Go file formatting with `gofmt -s`
- Run `go vet` for static analysis
- Prevent commits that don't pass these checks

To bypass any hook temporarily, use `git commit --no-verify`.

### Issue updates

Pushing an `issue_updates.json` file to the repository root allows the `update-issues` workflow to create, update or delete GitHub issues using the repository `GITHUB_TOKEN`. The file contains a JSON array where each object specifies an `action` and any supported issue fields. Duplicate issues are avoided by checking for an existing issue with the same title before creation. Example:

```json
[
  { "action": "create", "title": "New issue", "body": "Details" },
  { "action": "update", "number": 2, "state": "closed" }
]
```

The workflow runs on every push to `main` and processes the listed operations. After all entries are handled the file is removed on a new branch and a pull request is opened so that the cleanup can be merged back to `main`.

### Duplicate ticket cleanup

The `close-duplicates` workflow runs daily and on demand to detect open issues
with the same title. The script chooses the lowest numbered ticket as the
canonical reference and automatically closes the rest with a comment noting the
duplicate. This keeps the issue tracker focused on a single discussion for each
problem.

The project has achieved **~99% completion** with full Bazarr feature parity for core operations and nearly all optional enterprise features. See `TODO.md` for the remaining optional advanced features and implementation plan.
Extensive architectural details and design decisions are documented in
`docs/TECHNICAL_DESIGN.md`. New contributors should review this document to
understand package responsibilities and completed features.
For a detailed list of Bazarr features used as the parity target, see [docs/BAZARR_FEATURES.md](docs/BAZARR_FEATURES.md).
Instructions for importing an existing Bazarr configuration are documented in [docs/BAZARR_SETTINGS_SYNC.md](docs/BAZARR_SETTINGS_SYNC.md).

## License

This project is licensed under the terms of the MIT license. See `LICENSE` for details.
