# Subtitle Manager

Subtitle Manager is a command line application written in Go for converting, merging and translating subtitle files. It uses Cobra for its CLI interface and Viper for configuration management.

## Features

- Convert subtitles from many formats to SRT.
- Merge two subtitle tracks sorted by start time.
- Translate subtitles via Google Translate or ChatGPT APIs.
- Store translation history in an SQLite database.
- Per component logging with adjustable levels.

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
```

Configuration values are loaded from `$HOME/.subtitle-manager.yaml` by default. API keys may be specified via flags `--google-key` and `--openai-key` or in the configuration file. The SQLite database location defaults to `$HOME/.subtitle-manager.db` and can be overridden with `--db`.

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

## License

This project is licensed under the terms of the MIT license. See `LICENSE` for details.
