# TODO

This file tracks planned work, architectural decisions, and implementation status for Subtitle Manager.

## Roadmap

1. **Feature Parity with Bazarr**
   - Monitor media libraries for new subtitles.
   - Support multiple subtitle providers.
   - Download, manage and upgrade subtitles automatically.
   - Integrate with media servers (e.g. Plex, Emby).

2. **Configuration with Cobra & Viper**
   - Centralise configuration using Viper.
   - Provide CLI commands using Cobra.

3. **Logging Improvements**
   - Component based logging with adjustable levels via flags or config.

4. **Subtitle Processing**
   - Merge two subtitles in different languages into one.
   - Extract subtitles from various container formats and convert them to SRT.
   - Translate subtitles through Google Translate or ChatGPT.

5. **Database Schema**
   - Design an efficient schema to store subtitle metadata and history.

6. **Quality Assurance**
   - Unit tests for all packages.
   - Continuous integration workflow (future work).

## Implementation Plan

1. **Command Structure**
   - Use Cobra to define subcommands such as `convert`, `merge`, `translate` and `history`.
   - Load configuration with Viper including per-component log levels.

2. **Subtitle Operations**
   - Build utilities using the `go-astisub` library to read and manipulate subtitles.
   - Implement merging and translation logic in commands.
   - Implement extraction of subtitles from media containers.

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
   - Expose translation via a gRPC server and client.
   - Document protobuf messages and regeneration steps.

7. **Future Enhancements**
   - Replace manual HTTP calls with provider SDKs where available.
   - Consider asynchronous processing for bulk translations.
   - Evaluate performance of subtitle merging and translation.
   - Add optional web interface for managing subtitles.

## Additional Documentation

For a detailed description of the planned file layout, key functions and
Protobuf definitions, see `docs/TECHNICAL_DESIGN.md`.
The file `docs/BAZARR_FEATURES.md` enumerates all Bazarr features to ensure parity.
