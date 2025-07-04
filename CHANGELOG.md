# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Planned: Hybrid Protobuf + Go Types + gcommon Migration

- We are migrating to a hybrid model for all shared types and business logic:

  - **Protobufs** will define the canonical data models (e.g., LanguageProfile, MediaItem, etc.)
  - **Go types** will be generated from Protobufs for use in all Go projects
  - **gcommon** will contain all shared business logic, helpers, and interface implementations, importing the generated types
  - **Other languages** (Python, JS, etc.) can generate types from the same Protobufs as needed
  - **All work for this migration will be done on the `gcommon-refactor` branch**
  - **Main branches** will continue using Go types and type packages until the migration is complete

- This approach will:

  - Eliminate duplication and type drift
  - Enable cross-language compatibility
  - Centralize business logic and validation
  - Allow for incremental migration and testing

### Added

- **Automatic Episode Monitoring System**: Complete subtitle monitoring solution
  for TV episodes and movies
  - Sonarr/Radarr integration for automatic media library synchronization
  - Multi-language subtitle monitoring with configurable retry logic
  - Intelligent blacklist management with automatic and manual controls
  - Provider integration using existing priority and backoff systems
  - Quality upgrade monitoring for improved subtitle versions
  - Comprehensive CLI commands for monitoring management:
    - `subtitle-manager monitor sync` - Sync media from Sonarr/Radarr
    - `subtitle-manager monitor start` - Start monitoring daemon
    - `subtitle-manager monitor status` - View monitoring statistics
    - `subtitle-manager monitor list` - List all monitored items
    - `subtitle-manager monitor blacklist` - Manage blacklisted items
  - Scheduler integration for periodic monitoring with jitter
  - Database schema extensions for monitored items with retry tracking
  - Detailed statistics and progress reporting
  - Files modified: [pkg/monitoring/monitor.go](pkg/monitoring/monitor.go),
    [pkg/monitoring/sync.go](pkg/monitoring/sync.go),
    [pkg/monitoring/blacklist.go](pkg/monitoring/blacklist.go),
    [pkg/monitoring/scheduler.go](pkg/monitoring/scheduler.go),
    [cmd/monitor.go](cmd/monitor.go)
- **Universal Tagging System**: Complete unified tagging interface supporting
  all entity types
  - Enhanced Tag model with Type, EntityType, Color, and Description fields
  - Polymorphic tag associations table for consistent entity relationships
  - Standardized REST API endpoints for all entity types:
    `/api/{entityType}/{id}/tags`
  - Provider instance integration with tag-based selection and priority logic
  - Bulk tagging operations for efficient multi-entity management
  - Legacy migration path for existing user and media tag implementations
- **Provider Instance Management**: Enhanced provider system with priority and
  tagging support
  - Instance registration with configurable priority levels
  - Per-instance backoff logic for improved reliability
  - Tag-based provider selection and filtering
  - Factory registration system for tests and extensions
  - API endpoints supporting provider instance ID references
- **Enhanced Web UI Testing**: Minimal webui index for comprehensive test
  coverage
- Automated maintenance tasks with configurable scheduling
- Fetch languages, ratings, and episode data from OMDb via new metadata
  functions
- CLI flags and env vars for TMDB and OMDb API keys
- Advanced audio synchronization improvements with CPU vs accuracy slider,
  runtime tradeoff, and dual-language alignment.
- Experimental minimum display time mode that delays subtitles and catches up
  during silence.
- Automatic subtitle upgrade detection avoids replacing smaller files.
- Dashboard Widgets API exposing available widgets and layout endpoints.
- Security header `Referrer-Policy: no-referrer` to reduce referrer leakage.
- History API now filters by video file path via `video` query parameter.
- **Enhanced Issue Management Workflow**: Updated GitHub workflow configuration
  - Automatic triggering on merges to main branch when issue files change
  - Weekly scheduled maintenance on Sundays at 2 AM UTC
  - Manual workflow dispatch with comprehensive input options
  - Integration with ghcommon unified issue management system
  - Support for sub-issues with parent-child linking and automatic labeling
- **Issue Update File Management**: Comprehensive analysis and migration tools
  - Created `analyze-issue-state.py` for detailed state analysis and duplicate
    detection
  - Created `manage-issue-state.sh` for automated migration and cleanup
  - Intelligent duplicate prevention when migrating from old to new formats
  - Automatic conversion of legacy filename formats to GUID-based naming
  - Fixed JSON syntax errors in corrupted issue update files
  - All 38 issue update files now follow standardized GUID naming convention
  - Smart migration handled 55 actions from old format without creating
    duplicates
  - Complete coverage verification for ughi-fixed.sh script (11/11 issues
    covered)

### Changed

- **Tagging Architecture**: Migrated from separate tag systems to unified
  interface
  - User tags now use universal system with `entity_type='user'`
  - Media tags now use universal system with `entity_type='media'`
  - Provider tags integrated into instance management system
- **Database Schema**: Enhanced tag tables to support polymorphic relationships
- **API Structure**: Standardized tagging endpoints across all entity types

### Migration Notes

Existing tag data will be automatically migrated to the new unified system
during the next database schema update. No manual intervention required for most
installations.

## [0.9.0] - 2025-06-30

### Status Update

- **Project Nearing Completion**: Backend and frontend largely implemented but
  tagging, containerized Whisper, and maintenance tooling remain
- **Production Ready Backend**: Complete authentication, APIs, and provider
  support
- **Enterprise Features**: PostgreSQL, webhooks, notifications, anti-captcha,
  advanced scheduling complete
- **UI/UX Overhaul Completed**: All enhancements implemented according to the
  plan
- Code overview documentation added in docs/CODE_OVERVIEW.md

### Major Features Completed Since Last Release

- **PostgreSQL Database Backend**: Complete enterprise database support with
  full test coverage
- **Advanced Webhook System**: Sonarr/Radarr/custom webhook endpoints for
  library event integration
- **Anti-Captcha Integration**: Support for Anti-Captcha.com and 2captcha.com
  services
- **Notification Services**: Discord, Telegram, and SMTP notification providers
- **Advanced Scheduler**: Cron-based scheduling with full expression support
- **Bazarr Configuration Import**: Command-line tool for seamless Bazarr
  migration
- Complete Web UI implementation with all major pages
- Full REST API coverage for all operations
- History page with translation and download filtering
- System page with real-time logs and task monitoring
- Wanted page with subtitle search and management
- Comprehensive testing and documentation updates
- Production-ready authentication and authorization

### Planned (Optional Enhancement)

- Advanced reverse proxy base URL support for complex network setups
- Tagging system for language preferences
- Optional containerized Whisper ASR service with NVIDIA runtime
- Automated database cleanup and metadata refresh tasks

### Added

- **PostgreSQL database backend** with full enterprise support and graceful test
  skipping
- **Advanced webhook system** with `/api/webhooks/sonarr`,
  `/api/webhooks/radarr`, and `/api/webhooks/custom` endpoints
- **Anti-captcha integration** supporting Anti-Captcha.com and 2captcha.com
  services
- **Notification services** with Discord, Telegram, and SMTP providers
- **Advanced cron-based scheduler** with full expression support and granular
  controls
- **Bazarr configuration import** command for seamless migration from existing
  Bazarr installations
- REST endpoint `/api/convert` for subtitle file conversion
- REST endpoint `/api/translate` for translating uploaded subtitle files
- REST endpoint `/api/download` for on-demand subtitle fetching
- Build process now runs `go generate ./webui` to embed the latest web assets in
  the binary and container image.
- Automated workflow closes duplicate issues by title
- Embedded provider now enabled by default. Other providers remain hidden until
  explicitly added or imported.
- Enhanced General Settings page with Bazarr-compatible host, proxy, update,
  logging, backup and analytics options.

## [0.4.0] - 2025-06-12

### Major Milestone: Production Ready Release

This release marks ~95% project completion with full production readiness
achieved.

### Added

- Complete Web UI implementation with all major pages:
  - History page with translation and download filtering
  - System page with real-time logs and task monitoring
  - Wanted page with subtitle search and management
- Full REST API coverage for all subtitle operations
- Enhanced documentation reflecting current implementation status
- Production-ready status with comprehensive testing coverage

### Changed

- Updated project status documentation to reflect near-completion
- Improved README with current feature set and completion status
- Enhanced TODO.md to focus on remaining optional features only

### Notes

- **Bazarr Feature Parity**: Achieved full compatibility for core operations
- **Production Ready**: Complete authentication, authorization, and monitoring
- **Remaining Work**: Only optional advanced features (5% of project scope)

## [0.3.9] - 2025-06-26

### Changed

- `GoogleTranslate` now uses the official Google Cloud SDK instead of manual
  HTTP requests.

### Added

- Initial implementation of Subtitle Manager CLI.
- Commands: `convert`, `merge`, `translate`, and `history`.
- Google Translate and ChatGPT support.
- SQLite storage for translation history.
- Component based logging with adjustable levels.
- Documentation updates and initial technical design.

## [0.1.1] - 2025-06-06

### Added

- Expanded technical design document with detailed implementation plans.
- Updated README and TODO to reference the comprehensive documentation.

## [0.1.2] - 2025-06-07

### Added

- Documented Bazarr feature set in `docs/BAZARR_FEATURES.md`
- Linked feature reference from README, TODO and TECHNICAL_DESIGN

## [0.1.3] - 2025-06-08

### Added

- Subtitle extraction from media containers using `ffmpeg`.
- New `extract` CLI command.
- React based web UI built with Vite under `webui/`.
- `web` command to serve the embedded single page application.
- Technical design and TODO updated with web front end plan.

## [0.1.4] - 2025-06-09

### Added

- OpenSubtitles provider and `fetch` CLI command.
- Provider implementation documented in README and TODO.

## [0.1.5] - 2025-06-10

### Added

- Batch translation command for concurrent processing of multiple files.
- Helper functions `TranslateFileToSRT` and `TranslateFilesToSRT` in
  `pkg/subtitles`.
- Documentation updates for the new command and concurrency model.

## [0.1.6] - 2025-06-11

### Added

- Subscene provider with `fetch` and `watch` support.
- `grpc-server` command to expose translation service.
- Customisable ffmpeg path for subtitle extraction.
- Recursive directory watching.
- `delete` command and database deletion helper.

## [0.1.7] - 2025-06-12

### Added

- Environment variable configuration via `SM_` prefix.
- GitHub Actions workflow for continuous integration.

## [0.1.8] - 2025-06-13

### Added

- Comprehensive subtitle provider list from Bazarr documented in README and
  TODO.
- Implemented Addic7ed, BetaSeries, BSplayer, Podnapisi, TVSubtitles, Titlovi,
  LegendasDivx and GreekSubs providers.

## [0.1.9] - 2025-06-14

### Added

- Implemented the remaining subtitle providers from Bazarr's list.
- Unified provider selection via a registry.

## [0.1.10] - 2025-06-15

### Added

- Dockerfile and GitHub Actions workflow for container images.
- Container images published to GitHub Container Registry.
- Documentation updates describing provider registry and Docker usage.

## [0.2.0] - 2025-06-16

### Added

- Library scanning command to automatically download and upgrade subtitles.
- Updated README and TODO for new feature.

## [0.2.1] - 2025-06-17

### Added

- Authentication system supporting password login, one time tokens, OAuth2 and
  API keys.
- Simple user manager commands for creating users and generating API keys.
- RBAC with default roles and database backed session storage.

## [0.3.0] - 2025-06-18

### Added

- Concurrent directory scanning with worker pool.
- Sonarr and Radarr integration commands.
- Initial React web UI with login page.

## [0.3.1] - 2025-06-19

### Added

- PebbleDB backend with migration command.
- Configurable database backend via `--db-backend` flag.

## [0.3.2] - 2025-06-20

### Added

- Download history stored in database with new `downloads` command.

## [0.3.3] - 2025-06-21

### Added

- GitHub OAuth2 login support with new web server endpoints.

## [0.3.4] - 2025-06-22

### Added

- Role based access control enforced on web routes.
- `user role` command to modify user permissions.

## [0.3.5] - 2025-06-23

### Added

- One time login tokens with `user token` and `login-token` commands.

## [0.3.6] - 2025-06-24

### Added

- Manual subtitle search command with `search` functionality.

## [0.3.7] - 2025-06-24

### Added

- `user list` command to display existing accounts.

## [0.3.8] - 2025-06-25

### Added

- `media_items` table to store video metadata for library scanning.
- Library scan command `scanlib` populating the `media_items` table.
- REST endpoint `/api/extract` exposing subtitle extraction from media.
- `media_items` table to store video metadata for library scanning.
- Library scan command `scanlib` populating the `media_items` table.
- REST endpoint `/api/extract` exposing subtitle extraction from media.

## [0.3.9] - 2025-06-26

### Changed

- `GoogleTranslate` now uses the official Google Cloud SDK instead of manual
  HTTP requests.

## [0.3.10] - 2025-06-27

### Added

- Documentation update verifying PR workflow.
