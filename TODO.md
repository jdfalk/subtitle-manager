# TODO

This file tracks remaining work and implementation status for Subtitle Manager. **Note: The project is ~85% complete with most core functionality implemented.**

## 🎯 Remaining High Priority Tasks

### 1. Complete Web UI (Final 15%)

- [ ] **History Page**: Display translation and download history with filtering
- [ ] **System Page**: Log viewer, task status, system information
- [x] **Wanted Page**: Search for missing subtitles, manage wanted list
- [ ] **File Upload**: Forms for converting/translating uploaded subtitle files

### 2. Missing REST API Endpoints

- [ ] **`/api/download`**: Download subtitles for specific media files
- [x] **`/api/convert`**: Convert uploaded subtitle files between formats
- [ ] **`/api/translate`**: Translate uploaded subtitle files
- [x] **`/api/history`**: Retrieve translation and download history as JSON

### 3. Documentation Updates

- [ ] Update README to reflect current implementation status
- [ ] Mark completed items in roadmap sections
- [ ] Document new REST endpoints and web UI pages

### 4. Remaining Bazarr Features
- [ ] PostgreSQL database backend
- [ ] Reverse proxy base URL support
- [ ] Webhook endpoint for Plex events
- [ ] Anti-captcha service integration
- [ ] Scheduler for Sonarr/Radarr sync and subtitle upgrades

### 5. Bazarr Configuration Import

- [ ] Implement `import-bazarr` command that fetches settings from `/api/system/settings`
  using the user's API key.
- [ ] Map Bazarr preferences for languages, providers and network options into
  the Viper configuration.
- [ ] Document the synchronization process in `docs/BAZARR_SETTINGS_SYNC.md` and
  expose it through the welcome workflow.

## ✅ Completed Major Features

### Core Functionality (100% Complete)

- ✅ All CLI commands: `convert`, `merge`, `translate`, `history`, `extract`, `fetch`, `search`, `batch`, `scan`, `watch`, `delete`, `downloads`
- ✅ Configuration with Cobra & Viper including environment variables
- ✅ Component-based logging with adjustable levels

### Authentication & Authorization (100% Complete)

- ✅ Password authentication with hashed credentials
- ✅ One time token generation for email logins *(v0.3.5)*
- ✅ OAuth2 GitHub integration *(v0.3.3)*
- ✅ API key management with multiple keys per user
- ✅ Role based access control (admin, user, viewer) *(v0.3.4)*
- ✅ Session management with database persistence
- ✅ User management commands: `user add`, `user list`, `user role`, `user token`, `user apikey`

### Subtitle Processing (100% Complete)

- ✅ Convert between subtitle formats using go-astisub
- ✅ Merge two subtitle tracks sorted by time
- ✅ Extract subtitles from media containers via ffmpeg
- ✅ Translate subtitles through Google Translate (Cloud SDK) and ChatGPT
- ✅ Delete subtitle files and history records

### Provider Integration (100% Complete - Bazarr Parity Achieved)

- ✅ **40+ subtitle providers** including all major services:
  Addic7ed, OpenSubtitles, Subscene, Podnapisi, TVSubtitles, Titlovi,
  LegendasDivx, GreekSubs, BetaSeries, BSplayer, and 30+ more
- ✅ Provider registry for unified selection *(v0.1.9)*
- ✅ Manual subtitle search with `search` command *(v0.3.6)*

### Database & Storage (100% Complete)

- ✅ SQLite backend with full schema
- ✅ PebbleDB backend with migration support *(v0.3.1)*
- ✅ Translation history storage and retrieval
- ✅ Download history tracking *(v0.3.2)*
- ✅ Media items table for library metadata *(v0.3.8)*

### Library Management (100% Complete)

- ✅ Monitor directories for new media files (`watch` command)
- ✅ Scan existing libraries (`scan` and `scanlib` commands)
- ✅ Concurrent directory scanning with worker pools *(v0.3.0)*
- ✅ Recursive directory watching
- ✅ Sonarr and Radarr integration commands *(v0.3.0)*
- ✅ Metadata parsing with TheMovieDB integration

### Infrastructure (100% Complete)

- ✅ gRPC server for remote translation *(v0.1.6)*
- ✅ Docker support with automated builds *(v0.1.10)*
- ✅ GitHub Actions CI/CD pipeline *(v0.1.7)*
- ✅ Prebuilt container images on GitHub Container Registry

### Web UI (70% Complete)

- ✅ React application with Vite build system
- ✅ Authentication flow with login page
- ✅ Dashboard with library scanning functionality
- ✅ Settings page for configuration management
- ✅ Extract page for subtitle extraction
- ✅ Responsive design and navigation

## Web Front End Plan

The current React UI includes:

- **Authentication** – Login page with username/password and OAuth2 support
- **Dashboard** – Library scanning with progress tracking and provider selection
- **Settings** – Configuration management with live updates to YAML files
- **Extract** – Subtitle extraction from media files

**Remaining pages to implement:**

- **History** – Combined view of translation and download history with filtering
- **System** – Log viewer, task status, and system information

Additional pages such as blacklist management or per-movie editors can be added once core functionality is complete.

The front end is built with React and Vite under `webui/`. Run `go generate ./webui` to build the single page application which is embedded into the binary and served by the `web` command.

## Additional Documentation

For detailed architecture and design decisions, see `docs/TECHNICAL_DESIGN.md`.
The file `docs/BAZARR_FEATURES.md` enumerates all Bazarr features - parity has been achieved for providers and core functionality.
