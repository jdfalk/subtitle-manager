# Changelog

All notable changes to this project will be documented in this file.

## [0.1.0] - 2023-11-20
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
- Helper functions `TranslateFileToSRT` and `TranslateFilesToSRT` in `pkg/subtitles`.
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
- Comprehensive subtitle provider list from Bazarr documented in README and TODO.
- Implemented Addic7ed, BetaSeries, BSplayer, Podnapisi, TVSubtitles, Titlovi,
  LegendasDivx and GreekSubs providers.

## [0.1.9] - 2025-06-14
### Added
- Implemented the remaining subtitle providers from Bazarr's list.
- Unified provider selection via a registry.
