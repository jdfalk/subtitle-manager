# Complete Technical Design

This document maps the current implementation of Subtitle Manager with a focus on how each package
and its functions contribute to the overall workflow. The goal is to highlight how the
application achieves feature parity with Bazarr while providing a Go based tool chain.

## Directory Overview

- `cmd/` – Cobra commands and entry points.
- `pkg/` – Core application packages used by the CLI and web UI.
- `webui/` – React front end embedded with `go generate`.

## Workflow Summary

1. CLI commands in `cmd/` parse flags and environment variables.
2. Commands invoke helper packages in `pkg/` to perform actions such as subtitle
   extraction, translation and synchronization.
3. Results are persisted via `pkg/database` and surfaced through the web server in
   `pkg/webserver`.
4. Scheduled tasks run from `pkg/tasks` and `pkg/scheduler` to keep subtitles current.

## Key Packages and Functions

Below is a high level listing of important packages and the exported functions
they expose. Detailed code comments in each file provide parameter and return
value information.

### `pkg/audio`

- `SetFFmpegPath(path string)` – override ffmpeg binary.
- `ExtractTrack(mediaPath string, track int) (string, error)` – extract a single
  audio track to a temporary WAV file.
- `ExtractTrackWithDuration(mediaPath string, track int, offset, duration time.Duration) (string, error)` –
  extract a portion of an audio track.
- `GetAudioTracks(mediaPath string) ([]map[string]string, error)` – list available
  audio tracks. **Uses placeholder parsing logic.**
- `splitLines(s string) []string` – helper used by `GetAudioTracks`.

### `pkg/syncer`

- `Sync(mediaPath, subPath string, opts Options) ([]*astisub.Item, error)` – core
  synchronization routine.
- `Shift(items []*astisub.Item, offset time.Duration) []*astisub.Item` – apply a
  time offset.
- `Translate(items []*astisub.Item, lang, service, googleKey, gptKey, grpcAddr string) ([]*astisub.Item, error)` –
  helper called by `Sync` when translation is requested.
- `computeOffset(ref, target []*astisub.Item) time.Duration` – internal offset
  calculation.

### `pkg/subtitles`

- `ConvertToSRT(data []byte) ([]byte, error)` – convert any subtitle format to SRT.
- `ExtractTrack(mediaPath string, track int) ([]*astisub.Item, error)` – extract
  embedded subtitle tracks using ffmpeg.
- `Merge(itemsA, itemsB []*astisub.Item) []*astisub.Item` – merge two subtitle lists.

### `pkg/translator`

- `Translate(service, text, lang, googleKey, gptKey, grpcAddr string) (string, error)` –
  unified translation entry point.
- `GoogleTranslate(text, lang, key string) (string, error)` – Google Cloud client.
- `GPTTranslate(text, lang, key string) (string, error)` – ChatGPT API.

### `pkg/webserver`

- `New(cfg Config) *Server` – create HTTP server instance.
- `(*Server) Start() error` – launch REST API and web UI.

### Commands in `cmd/`

- `convert`, `merge`, `translate`, `sync`, `scan`, `watch`, `autoscan`, `web` –
  user facing commands that orchestrate operations by calling the packages above.

## Bazarr Parity

Each module implements functionality found in Bazarr:

- Provider integrations in `pkg/providers` mirror the subtitle download logic.
- `pkg/watchers` and `pkg/tasks` replicate Bazarr's automatic library scanning.
- The synchronization features in `pkg/syncer` replace Bazarr's sync engine with
  additional translation support.

## Future Enhancements

Planned improvements are tracked in `TODO.md`. Notably the audio
synchronization code will evolve to support CPU vs accuracy tuning and advanced
multi-language alignment.
