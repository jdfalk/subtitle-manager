# Subtitle Manager

Subtitle Manager is a command line application written in Go for converting, merging and translating subtitle files. It uses Cobra for its CLI interface and Viper for configuration management.

## Features

- Convert subtitles from many formats to SRT.
- Merge two subtitle tracks sorted by start time.
- Translate subtitles via Google Translate or ChatGPT APIs.
- Store translation history in an SQLite database.
- Per component logging with adjustable levels.
- Extract subtitles from media containers using ffmpeg.
- Download subtitles from OpenSubtitles.
- Batch translate multiple files concurrently.

## Installation

```bash
# clone repository
$ git clone <this repository>
$ cd subtitle-manager
# install dependencies and build
$ go build
```

## Usage

```
subtitle-manager convert [input] [output]
subtitle-manager merge [sub1] [sub2] [output]
subtitle-manager translate [input] [output] [lang]
subtitle-manager history
subtitle-manager extract [media] [output]
subtitle-manager fetch opensubtitles [media] [lang] [output]
subtitle-manager batch [lang] [files...]
```

### Web UI

Run `subtitle-manager web` to start the embedded React interface on `:8080`. The SPA is built via `go generate` in the `webui` directory and embedded using Go 1.16's `embed` package.

Configuration values are loaded from `$HOME/.subtitle-manager.yaml` by default. API keys may be specified via flags `--google-key` and `--openai-key` or in the configuration file. The SQLite database location defaults to `$HOME/.subtitle-manager.db` and can be overridden with `--db`.  Translation can be delegated to a remote gRPC server using the `--grpc` flag and providing an address such as `localhost:50051`.

Example configuration:

```
log-level: info
log_levels:
  translate: debug
translate_service: google
```

## Development

Tests can be run with `go test ./...`.

The project aims to eventually reach feature parity with [Bazarr](https://github.com/morpheus65535/bazarr) while offering improved configuration and logging. See `TODO.md` for the full roadmap and implementation plan.
Extensive architectural details and design decisions are documented in
`docs/TECHNICAL_DESIGN.md`. New contributors should review this document to
understand package responsibilities and future plans.
For a detailed list of Bazarr features used as the parity target, see [docs/BAZARR_FEATURES.md](docs/BAZARR_FEATURES.md).

## License

This project is licensed under the terms of the MIT license. See `LICENSE` for details.
