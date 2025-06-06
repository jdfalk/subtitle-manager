# Technical Design

This document describes the overall architecture, data structures and
interfaces used by **Subtitle Manager**. The goal is to outline the files,
packages and Protobuf messages required to achieve feature parity with
[Bazarr](https://github.com/morpheus65535/bazarr) while remaining easy to
extend.

## Directory Structure

```text
cmd/                // CLI commands implemented with Cobra
pkg/
  database/         // SQLite database layer
  logging/          // Component based logging utilities
  providers/        // Subtitle provider integrations (future work)
  subtitles/        // Subtitle operations (convert, merge, extract)
  translator/       // Translation services (Google, ChatGPT)
```

## Key Files and Packages

### `cmd/`

- `root.go` – application entry point, loads configuration with Viper and sets
  global flags.
- `convert.go` – converts subtitles to SRT format.
- `merge.go` – merges two subtitle streams by start time.
- `translate.go` – translates subtitles via the configured service.
- `history.go` – displays stored translation history.

### `pkg/database`

- `database.go`
  - `Open(path string) (*sql.DB, error)` – opens or creates the SQLite database
    and initialises the schema.
  - `InsertSubtitle(db *sql.DB, file, lang, service string) error` – stores a
    record of a translation operation.
  - `ListSubtitles(db *sql.DB) ([]SubtitleRecord, error)` – returns translation
    history.
  - `SubtitleRecord` – struct representing a history row.

### `pkg/logging`

- `logging.go`
  - `GetLogger(component string) *logrus.Entry` – obtains a logger for the given
    component, honouring per component log levels from configuration.

### `pkg/subtitles` (planned)

- `convert.go` – utility functions for opening subtitle files in any supported
  format and writing them as SRT.
- `merge.go` – helpers to combine subtitle tracks.
- `extract.go` – routines for extracting subtitle streams from video containers.

### `pkg/providers` (planned)

- provider specific subpackages for downloading subtitles from online services.
- functions follow the common signature
  `FetchSubtitle(ctx context.Context, mediaPath string) (io.Reader, error)`.

### `pkg/translator`

- `translator.go`
  - `GoogleTranslate(text, targetLang, apiKey string) (string, error)` – uses the
    Google Translate API.
  - `GPTTranslate(text, targetLang, apiKey string) (string, error)` – uses the
    ChatGPT API.
  - `TranslateFunc` – function type utilised by translation commands.
  - `SetGoogleAPIURL(u string)` – test helper to override the Google endpoint.

## Database Schema

The database currently contains a single table `subtitles` used to track
translations:

```sql
CREATE TABLE IF NOT EXISTS subtitles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file TEXT NOT NULL,
    language TEXT NOT NULL,
    service TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
```

Future versions will extend the schema with tables for subtitle providers and
media library metadata. Schema migrations will be handled via `golang-migrate`.

## Protobuf Service Definition (Planned)

To enable remote translation or subtitle processing, a gRPC service can be used.
Below is an example service definition. The file would live in `proto/translator.proto`.

```protobuf
syntax = "proto3";
package translator;

service Translator {
  rpc Translate (TranslateRequest) returns (TranslateResponse);
}

message TranslateRequest {
  string text = 1;
  string language = 2;
}

message TranslateResponse {
  string translated_text = 1;
}
```

The generated Go code will reside in `pkg/translatorpb`. Commands would use the
gRPC client when configured to do so.

## Future Work

- Implement additional CLI commands for monitoring media libraries and
  downloading subtitles.
- Introduce workers and queues for asynchronous translation.
- Add integration tests covering provider interactions and command behaviours.
- Provide container images for easier deployment.

