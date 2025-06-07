# TODO

This file tracks planned work, architectural decisions, and implementation status for Subtitle Manager.

## Roadmap

1. **Feature Parity with Bazarr**
   - Monitor media libraries for new subtitles. *(watch command implemented)*
   - Support multiple subtitle providers. *(Full Bazarr provider list implemented)*
   - Download, manage and upgrade subtitles automatically.
   - Integrate with media servers (e.g. Plex, Emby, Sonarr, Radarr).

2. **Configuration with Cobra & Viper**
   - Centralise configuration using Viper.
   - Provide CLI commands using Cobra.

3. **Logging Improvements**
   - Component based logging with adjustable levels via flags or config.

4. **Subtitle Processing**
   - Merge two subtitles in different languages into one.
   - Extract subtitles from various container formats and convert them to SRT.
   - Translate subtitles through Google Translate or ChatGPT.
   - Allow configuring the `ffmpeg` binary path.
   - Delete external subtitles directly from disk.

5. **Database Schema**
   - Design an efficient schema to store subtitle metadata and history.

6. **Quality Assurance**
   - Unit tests for all packages.
   - Continuous integration workflow using GitHub Actions.

## Implementation Plan

1. **Command Structure**
   - Use Cobra to define subcommands such as `convert`, `merge`, `translate` and `history`.
   - Load configuration with Viper including per-component log levels.

2. **Subtitle Operations**
   - Build utilities using the `go-astisub` library to read and manipulate subtitles.
   - Implement merging and translation logic in commands.
  - Implement extraction of subtitles from media containers. *(implemented in v0.1.3)*
   - Register providers via a registry for unified selection. *(implemented in v0.1.9)*

3. **Translation Services**
   - Provide Google Translate and ChatGPT implementations behind a common interface.
   - Allow users to choose the service via config or CLI flags.

4. **Database Layer**
   - Use SQLite to store records of translations and operations.
   - Plan schema upgrades to support richer metadata in the future.

5. **Testing Strategy**
   - Write tests for database interactions and translation providers.
   - Test command behaviour with edge cases.

6. **Remote Services**
   - Expose translation via a gRPC server and client. *(client and server implemented)*
   - Document protobuf messages and regeneration steps.

7. **Media Library Monitoring**
   - Implement filesystem watchers to detect new video files.
   - Automatically fetch subtitles when media appears. *(recursive watching implemented)*

8. **Future Enhancements**
   - Replace manual HTTP calls with provider SDKs where available.
   - Asynchronous processing for bulk translations implemented via the `batch` command.
   - Evaluate performance of subtitle merging and translation.
   - Add optional web interface for managing subtitles.
   - Distribute official Docker image via GitHub Actions.

## Additional Documentation

For a detailed description of the planned file layout, key functions and
Protobuf definitions, see `docs/TECHNICAL_DESIGN.md`.
The file `docs/BAZARR_FEATURES.md` enumerates all Bazarr features to ensure parity.

## Web Front End Plan

The Bazarr project exposes many pages in its web UI. Pages identified in the repository include:
- Authentication
- Blacklist (Movies, Series)
- Episodes
- History (Movies, Series, Statistics)
- Movies (Details, Editor)
- Series (Editor)
- Settings (General, Languages, Notifications, Plex, Providers, Radarr, Scheduler, Sonarr, Subtitles, UI)
- System (Announcements, Backups, Logs, Providers, Releases, Status, Tasks)
- Wanted (Movies, Series)
- Various error pages and utility views

Subtitle Manager will initially implement a simplified set of pages organised for faster navigation:
- **Dashboard** – landing page summarising recent activity.
- **History** – combined view of translated subtitles.
- **Settings** – grouped configuration panels similar to Bazarr's settings sections.
- **System** – log viewer and task status.
- **Wanted** – missing subtitle search page.

Additional pages such as blacklist management or per-movie editors can be added once core functionality matches Bazarr.

The front end is built with React and Vite under `webui/`. `go generate ./webui` builds the single page application which is embedded into the binary and served by the `web` command.
