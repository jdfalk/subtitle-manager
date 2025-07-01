# Technical Design

This document provides an in depth description of the architecture, data flows
and components that make up **Subtitle Manager**. It consolidates design
decisions and implementation details to help contributors understand the
codebase. The design aims to meet the project goals of feature parity with
[Bazarr](https://github.com/morpheus65535/bazarr), robust configuration via
Cobra and Viper, pluggable translation services, and an extensible database
model.

For a summary of Bazarr's capabilities used as a target feature set, see
[docs/BAZARR_FEATURES.md](BAZARR_FEATURES.md).

## 1. Directory Structure

The repository is organised using a standard Go project layout. Top level
directories include command implementations, reusable packages and
documentation.

\```text cmd/ // CLI command definitions pkg/ // Core application packages
database/ // Database layer built on SQLite logging/ // Logging helpers
providers/ // Subtitle provider integrations subtitles/ // Subtitle processing
(convert, merge, extract) translator/ // Google Translate and ChatGPT clients
proto/ // gRPC service definitions (generated code in pkg/translatorpb)
internal/ // Internal utilities not intended for external use scripts/ // Helper
scripts for CI/CD (future work) configs/ // Example configuration files \```

Every package contains a README giving a quick overview of its responsibilities.
Functions are documented using Go doc comments following the style guidelines
laid out in `AGENTS.md`.

## 2. Command Line Interface

The CLI is composed of several Cobra commands. The root command initialises
configuration and logging. Subcommands perform operations on subtitle files and
manage history.

### 2.1 Root Command

`cmd/root.go` sets persistent flags for configuration paths, database locations
and global log level. It loads configuration using Viper before executing the
selected subcommand. Errors are logged and returned with `os.Exit(1)` where
appropriate.

\``` subtitle-manager [global flags] <command> [command flags] \```

Common flags:

- `--config` path to configuration file (defaults to
  `$HOME/.subtitle-manager.yaml`).
- `--db` path to the SQLite database.
- `--log-level` global log level (debug, info, warn, error).
- `--log-levels` comma separated list of per-component levels (e.g.
  `translate=debug`).

### 2.2 Subcommands

#### convert

\``` subtitle-manager convert <input> <output> \```

Converts any supported subtitle format to SRT using utilities in
`pkg/subtitles`. The command reads the input file, detects the format with
`go-astisub`, and writes SRT to the specified output. Errors in the conversion
process are returned to the caller.

#### merge

\``` subtitle-manager merge <sub1> <sub2> <output> \```

Merges two subtitle streams by time code. The command calls
`subtitles.MergeTracks` which returns a combined set of subtitles sorted by
start time. The resulting subtitles are saved to the output path in SRT format.

#### translate

\``` subtitle-manager translate <input> <output> <language> \```

Translates a subtitle file to the target language. The translator command uses
`translator.Translate` which selects the translation backend (Google or ChatGPT)
based on configuration. Translation results are stored in the database via
`database.InsertSubtitle`.

#### batch

\``` subtitle-manager batch <language> <files...> \```

Translates multiple subtitle files concurrently. The command invokes
`subtitles.TranslateFilesToSRT` which utilises the `conc` package to limit the
number of worker goroutines.

#### history

Displays previously translated files. `database.ListSubtitles` returns rows
which are printed in a tabular format.

## 3. Configuration Management

Configuration is handled by Viper with support for YAML files, environment
variables and command line flags. The default configuration file location is
`$HOME/.subtitle-manager.yaml` and may contain options shown below.

\```yaml log-level: info log_levels: translate: debug translator: google
translator_api_keys: google: "<API key>" chatgpt: "<API key>" database:
"~/.subtitle-manager.db" \```

Environment variables are automatically mapped to configuration keys using
Viper's `AutomaticEnv`. For example `SUBTITLE_MANAGER_TRANSLATOR=chatgpt`
overrides the translator option.

## 4. Logging

The `pkg/logging` package wraps `logrus` to provide component based logging.
`GetLogger(component string)` returns a logger configured with the level defined
in the configuration file. Log messages include timestamps and the component
name for filtering.

Example usage:

\```go log := logging.GetLogger("translate") log.Debug("calling Google API")
\```

Logging levels can be updated at runtime through configuration reload (future
work) or via CLI flags.

## 5. Database Layer

The application stores metadata in an SQLite database managed by `pkg/database`.
A PebbleDB implementation provides a drop-in replacement through the
`SubtitleStore` interface. The current SQLite schema is defined below and future
revisions will add additional tables for subtitle providers and media libraries.

\```sql CREATE TABLE IF NOT EXISTS subtitles ( id INTEGER PRIMARY KEY
AUTOINCREMENT, file TEXT NOT NULL, video_file TEXT, release TEXT, language TEXT
NOT NULL, service TEXT NOT NULL, embedded INTEGER NOT NULL DEFAULT 0, created_at
TIMESTAMP NOT NULL );

CREATE INDEX IF NOT EXISTS subtitles_file_idx ON subtitles(file); \```

The library scan feature stores discovered videos in a dedicated table:

\```sql CREATE TABLE IF NOT EXISTS media_items ( id INTEGER PRIMARY KEY
AUTOINCREMENT, path TEXT NOT NULL, title TEXT NOT NULL, season INTEGER, episode
INTEGER ); \```

### 5.1 Functions

- `Open(path string) (*sql.DB, error)` – opens or creates the database and runs
  migrations.
- `OpenSQLStore(path string) (*SQLStore, error)` – opens an SQLite backed
  `SubtitleStore`.
- `OpenStore(path, backend string) (SubtitleStore, error)` – returns either the
  SQLite or Pebble implementation based on `backend`.
- `InsertSubtitle(db *sql.DB, file, video, lang, service, release string, embedded bool) error`
  – inserts a translation record with metadata.
- `ListSubtitles(db *sql.DB) ([]SubtitleRecord, error)` – retrieves history
  sorted by most recent.
- `InsertMediaItem(db *sql.DB, path, title string, season, episode int) error` –
  records a media file in the library table.
- `ListMediaItems(db *sql.DB) ([]MediaItem, error)` – retrieves stored library
  entries.

Each function uses context-aware queries and prepares statements for efficiency.

## 6. Subtitle Processing

The `pkg/subtitles` package is responsible for reading, writing and manipulating
subtitle files.

### 6.1 Converters

`ConvertToSRT(path string) ([]byte, error)` reads a subtitle file in any
supported format (SSA/ASS, VTT, TTML, etc.) using `go-astisub`. It converts the
internal representation to SRT and returns the bytes for writing.

### 6.2 Merging

`MergeTracks(subA, subB []astisub.Item) ([]astisub.Item, error)` merges two
subtitle tracks by start time while preserving style cues if present. Collisions
are resolved by time code with configurable offsets.

### 6.3 Extraction

`ExtractFromMedia(mediaPath string) ([]astisub.Item, error)` utilises `astits`
to extract subtitle streams from containers such as MKV or MP4. Extracted tracks
are converted to SRT using `ConvertToSRT`.

## 7. Translation Services

`pkg/translator` defines a common interface for translation providers.

\```go // TranslateFunc represents a translation function. type TranslateFunc
func(ctx context.Context, text, lang string) (string, error) \```

### 7.1 Google Translate

`GoogleTranslate(text, lang, apiKey string) (string, error)` uses the official
Google Cloud client library. A client is created with `option.WithAPIKey` and
the endpoint can be overridden for tests. The translation result is returned
from the SDK response.

### 7.2 ChatGPT

`GPTTranslate(text, lang, apiKey string) (string, error)` sends the text to
OpenAI's ChatGPT API. Request and response structures are defined in
`translator.go` and include appropriate timeouts and error handling.

The translation command selects the implementation using a map of provider names
to functions:

\```go providers := map[string]translator.TranslateFunc{ "google":
translator.GoogleTranslate, "chatgpt": translator.GPTTranslate, } \```

## 8. gRPC Service (Optional)

To enable remote processing or integration with other services, Subtitle Manager
defines a gRPC service in `proto/translator.proto`. Generated code is committed
to `pkg/translatorpb`.

\```protobuf syntax = "proto3"; package translator; service Translator { rpc
Translate(TranslateRequest) returns (TranslateResponse); } message
TranslateRequest { string text = 1; string language = 2; } message
TranslateResponse { string translated_text = 1; } \```

A server implementation in `cmd/grpcserver` exposes translation functionality
over the network. Clients may configure the CLI to use a remote gRPC server via
the `--grpc` flag.

## 9. Provider Integrations

The `pkg/providers` directory will host modules for fetching subtitles from
online services. Each provider implements the following interface:

\```go // Provider downloads subtitles for a given media item. type Provider
interface { Fetch(ctx context.Context, mediaPath string, lang string) ([]byte,
error) } \```

Providers share common configuration such as API keys or user credentials via
Viper. Future work includes supporting services used by Bazarr (OpenSubtitles,
Addic7ed, Subscene, etc.). An initial provider based on the OpenSubtitles REST
API has been implemented under `pkg/providers/opensubtitles` and exposed through
the `fetch` command.

## 10. Concurrency Model

Operations like batch translation and subtitle downloading benefit from
concurrency. The project uses `conc` from Sourcegraph to limit the number of
parallel workers.

Example workflow for translating multiple files:

\```go pool := conc.NewPool(5) for \_, file := range files { f := file
pool.Go(func() error { return translateFile(ctx, f) }) } if err := pool.Wait();
err != nil { log.WithError(err).Error("batch translation failed") } \```

Workers share a single database connection pool and logger. Rate limiting is
built into each provider implementation where required by the external service.

## 10.1 Asynchronous Translation Queue

The `pkg/queue` package provides an asynchronous job queue for heavy translation 
tasks. It leverages the existing `pkg/tasks` infrastructure for progress tracking
and integrates with the established concurrency patterns.

### Queue Architecture

The queue supports different job types:

- **SingleFileJob** – Translates a single subtitle file
- **BatchFilesJob** – Translates multiple files concurrently

Jobs implement the `Job` interface:

\```go
type Job interface {
    ID() string
    Type() JobType
    Execute(ctx context.Context) error
    Description() string
}
\```

### Queue Management

The queue is managed through worker goroutines:

\```go
// Create and start queue with 3 workers
q := queue.NewQueue(3)
q.Start()

// Add translation job
job := queue.NewSingleFileJob(inPath, outPath, lang, service, apiKey, "", "")
taskID, err := q.Add(job)
\```

### CLI Integration

Translation commands support asynchronous processing via the `--async` flag:

\```bash
# Synchronous translation (default)
subtitle-manager translate input.srt output.srt en

# Asynchronous translation
subtitle-manager translate --async input.srt output.srt en
\```

Queue status and management:

\```bash
# Check queue status and active tasks
subtitle-manager queue status

# Start queue for testing (auto-managed in production)
subtitle-manager queue start
\```

The queue integrates with the existing task tracking system, allowing progress
monitoring through the `tasks` package and web interface.

## 11. Testing Strategy

Automated tests reside alongside packages. Unit tests mock external dependencies
such as HTTP APIs and the SQLite database. Integration tests validate CLI
behaviour using the `os/exec` package.

Key areas with test coverage:

- Database CRUD functions with an in-memory SQLite database.
- Subtitle conversion and merging functions with sample subtitle files.
- Translation provider clients using mocked HTTP servers.
- Command layer tests verifying flag parsing and error cases.

CI will run `go vet`, `golint` and `go test ./...` to ensure quality.

## 12. Security Considerations

Sensitive information such as API keys is never logged. Configuration files may
be encrypted in the future using `age` or `sops`. HTTP clients enforce TLS and
set reasonable timeouts. Input validation is performed when reading subtitle
files and parsing command line arguments.

## 13. Performance Notes

Subtitle processing can be CPU intensive for large files. The `go-astisub`
library is efficient for typical subtitle lengths. Batch translation leverages
concurrency while respecting provider rate limits. Database access is local and
uses indexes to minimise query latency.

## 14. Implementation Plan

The following steps outline the order of implementation to achieve the project
goals.

1. **Command Framework** – set up Cobra commands and global configuration.
2. **Logging Layer** – implement component based loggers and integrate with
   commands.
3. **Subtitle Utilities** – create conversion, merging and extraction helpers.
4. **Database Schema** – define SQLite schema and implement CRUD functions.
5. **Translation Providers** – implement Google and ChatGPT backends with
   interfaces for future providers.
6. **History Command** – allow users to view past translation actions.
7. **Provider Integrations** – add subtitle download modules for popular
   services.
8. **Media Library Monitoring** – implement watchers to detect new media and
   automatically download subtitles.
9. **Library Scanning** – add command to scan existing directories and fetch
   missing or improved subtitles.
10. **gRPC API (Optional)** – provide remote translation capabilities.
11. **Testing and CI** – expand unit tests and add CI workflows.

This plan aligns with the tasks listed in `TODO.md`.

## 15. Glossary

- **Subtitle** – Textual representation of dialog in a video.
- **SRT** – SubRip Subtitle format, widely supported.
- **Translation Provider** – External API used to translate text.
- **Viper** – Go configuration management library.
- **Cobra** – Library for building CLI applications.
- **SQLite** – Lightweight file based database used for storing history.
- **gRPC** – Remote procedure call system based on Protocol Buffers.

## 16. Conclusion

This design document details the intended architecture for Subtitle Manager. By
following the implementation plan, the project will deliver a Go based CLI tool
matching Bazarr's functionality while offering a modern configuration system,
flexible logging and extensible translation services.

## 17. Sample Configuration File

An example configuration demonstrating all available options:

\```yaml

# Global log level

log-level: info

# Component specific overrides

log_levels: database: warn translate: debug

# Database file location

# The path can be absolute or relative to the configuration file

# by default this uses $HOME/.subtitle-manager.db

database: /var/lib/subtitle-manager/app.db

# Translator selection and API keys

translate_service: chatgpt google_api_key: YOUR_GOOGLE_API_KEY openai_api_key:
YOUR_CHATGPT_API_KEY google_api_url:
https://translation.googleapis.com/language/translate/v2 openai_model:
gpt-3.5-turbo ffmpeg_path: /usr/bin/ffmpeg batch_workers: 4 scan_workers: 4
opensubtitles: api_key: YOUR_OS_API_KEY api_url: https://rest.opensubtitles.org
user_agent: subtitle-manager/0.1

providers: generic: api_url: https://example.com/subtitles username: myuser
password: secret api_key: token \```

The CLI supports environment variable overrides using the prefix
`SUBTITLE_MANAGER_`. For example `SUBTITLE_MANAGER_DATABASE=/tmp/test.db` will
force the database location.

## 18. Command Flow Examples

### 18.1 Convert Flow

\```text User -> convert command -> subtitles.ConvertToSRT -> write output \```

1. The user invokes `subtitle-manager convert spanish.ssa spanish.srt`.
2. `cmd/convert.go` loads configuration and obtains a logger for the `convert`
   component.
3. `subtitles.ConvertToSRT` opens `spanish.ssa`, reads all events, and
   serialises them to SRT.
4. The resulting bytes are written to `spanish.srt`. Any warnings during parsing
   are logged.

### 18.2 Translate Flow

\```text User -> translate command -> subtitles.ConvertToSRT ->
translator.Translate -> database.InsertSubtitle \```

1. The user invokes `subtitle-manager translate movie.srt translated.srt fr`.
2. The command converts `movie.srt` to a plain text block and calls
   `translator.Translate` with the target language `fr`.
3. The translator selects the provider configured in Viper and calls
   `GoogleTranslate` or `GPTTranslate` accordingly.
4. The translated text is reassembled into subtitle items with the same timing
   information and written to `translated.srt`.
5. A record of the translation is inserted into the database.

### 18.3 Merge Flow

\```text sub1 + sub2 -> subtitles.MergeTracks -> subtitles.WriteSRT \```

1. The user runs `subtitle-manager merge eng.srt fre.srt dual.srt`.
2. `subtitles.ReadFile` parses both input files.
3. `MergeTracks` interleaves subtitle items ordered by start time.
4. The combined list is saved to `dual.srt`.

## 19. Database Schema Evolution

To support additional features such as subtitle provider history and media
library monitoring, new tables will be added. Migrations are managed by
`golang-migrate` with files stored in `migrations/`.

### 19.1 Planned Tables

- `providers` – list of configured subtitle providers, credentials and status.
- `media_items` – records of movies or episodes by unique hash. _Implemented in
  v0.3.8_
- `downloads` – history of downloaded subtitles with provider references.
  _Implemented in v0.3.2_

Each migration file is numbered sequentially and includes both `up` and `down`
SQL scripts.

## 20. Error Handling Strategy

Errors are wrapped with context using the `%w` verb in `fmt.Errorf`. The top
level CLI catches errors and exits with non-zero status. Common error types
include:

- `ErrInvalidSubtitle` when input subtitle parsing fails.
- `ErrTranslateFailed` when an external API returns an error.
- `ErrDatabase` for database connectivity or SQL issues.

## 21. Style Guide

The project follows the Go community style guide in addition to rules in
`AGENTS.md`. Key points:

- Use descriptive names and keep functions short.
- Document all exported identifiers.
- Group related functions into files by topic.
- Keep error messages consistent and actionable.

## 22. Future Enhancements

-Additional features under consideration:

- **Web Interface** – Provide a lightweight web UI for managing translations and
  history. Implementation uses a React app built with Vite in `webui`.
  `go generate` builds the frontend and embeds it using the `embed` package for
  serving via the `web` command. Run `go generate ./webui` before compiling the
  binary to ensure the latest assets are included.
- **Asynchronous Queue** – ✅ **IMPLEMENTED** Use a worker queue for heavy translation tasks.
  Available via `pkg/queue` with CLI support through `--async` flag on translate
  commands and `queue` management commands.
- **Cloud Storage** – Allow storing subtitles and history in cloud buckets.
- **Internationalisation** – Localise CLI messages and documentation.

These ideas will be evaluated after core functionality is stable.

## 23. Frequently Asked Questions

**Q:** _How do I contribute new subtitle providers?_

A: Create a package under `pkg/providers/<name>` that implements the `Provider`
interface. Add configuration options to Viper and document them in README.

**Q:** _Can I run the translator API on a remote server?_

A: Yes. Build the gRPC server (`cmd/grpcserver`) and start it on the desired
host. Configure the CLI with `--grpc <address>` to use the remote server for
translations.

**Q:** _Where can I find example subtitles for testing?_

A: Sample subtitle files are located in `testdata/`. Tests reference these files
to verify correct behaviour.

## 24. Reference Material

- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Cobra Documentation](https://github.com/spf13/cobra)
- [Viper Documentation](https://github.com/spf13/viper)
- [go-astisub](https://github.com/asticode/go-astisub)
- [logrus](https://github.com/sirupsen/logrus)
- [Sourcegraph conc](https://github.com/sourcegraph/conc)

These resources provide background on the tools and libraries used in the
project.

## 25. ASCII Component Diagram

Below is a simplified view of the major components and how they interact during
a translation operation:

\```text +-------------+ +---------------+ +---------------+ | CLI Command | -->
| Subtitle Utils| --> | Translator API| +-------------+ +---------------+
+---------------+ | | | | v | | +-------------+ | | | Database | <-----------+ |
+-------------+ \```

The command reads the input, delegates parsing and formatting to the subtitle
utilities, which then call the translator API. Results are stored in the
database and returned to the command for output.

## 26. File Overview

A summary of key files and their responsibilities:

- `cmd/root.go` – CLI entry point and configuration loader.
- `cmd/convert.go` – Implements the `convert` command.
- `cmd/merge.go` – Implements the `merge` command.
- `cmd/translate.go` – Implements the `translate` command.
- `cmd/history.go` – Implements the `history` command.
- `pkg/database/database.go` – SQLite connection management and CRUD helpers.
- `pkg/logging/logging.go` – Component based logger retrieval.
- `pkg/subtitles/convert.go` – Subtitle format detection and conversion
  functions.
- `pkg/subtitles/merge.go` – Track merging utilities.
- `pkg/subtitles/extract.go` – Subtitle extraction from media containers.
- `pkg/translator/translator.go` – Common translation interface and provider
  implementations.
- `proto/translator.proto` – gRPC service definitions.

Keeping these files small and focused allows new contributors to quickly
understand the responsibilities of each package.

## 27. Bazarr Feature Reference

The file [BAZARR_FEATURES.md](BAZARR_FEATURES.md) lists the important functions,
features and subtitle providers implemented by Bazarr. Subtitle Manager aims to
implement equivalent capabilities. Use that document as a checklist when
evaluating progress toward full feature parity.

## 28. Subtitle Synchronization (WIP)

Package `pkg/syncer` provides the foundation for aligning external subtitle
files with media. The initial version simply loads subtitles and supports
shifting by a constant offset. Future iterations will analyze audio tracks and
embedded subtitles to calculate precise timing adjustments. The package can also
translate subtitles to a target language during the sync process using the
existing translation providers.

## 29. Related Documentation

For additional context, see the following documents:

- [API Design](API_DESIGN.md)
- [Test Design](TEST_DESIGN.md)
- [Process Workflows](PROCESS_WORKFLOWS.md)
- [Developer Guide](DEVELOPER_GUIDE.md)
