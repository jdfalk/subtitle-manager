# Subtitle Manager

Subtitle Manager is a command line application written in Go for converting, merging and translating subtitle files. It uses Cobra for its CLI interface and Viper for configuration management.

## Features

- Convert subtitles from many formats to SRT.
- Merge two subtitle tracks sorted by start time.
- Translate subtitles via Google Translate or ChatGPT APIs.
- Store translation history in an SQLite database.
- Per component logging with adjustable levels.
- Extract subtitles from media containers using ffmpeg.
- Download subtitles from a comprehensive list of providers based on Bazarr,
  including Addic7ed, AnimeKalesi, Animetosho, Assrt, Avistaz, BetaSeries,
  BSplayer, GreekSubs, Podnapisi, Subscene, TVSubtitles, Titlovi, LegendasDivx
  and many more.
- Batch translate multiple files concurrently.
- Monitor directories and automatically download subtitles.
- Recursive directory watching with -r flag.
- Run a translation gRPC server.
- Delete subtitle files and remove history records.
- Provider registry simplifies adding new sources.
- Dockerfile and workflow for container builds.
- Prebuilt images published to GitHub Container Registry.

### Supported Subtitle Providers

Subtitle Manager now supports the full provider list from Bazarr. The following
services are available:

- Addic7ed
- AnimeKalesi
- Animetosho
- Assrt
- AvistaZ / CinemaZ
- BetaSeries
- BSplayer
- Embedded Subtitles
- Gestdown.info
- GreekSubs
- GreekSubtitles
- HDBits.org
- Hosszupuska
- Karagarga.in
- Ktuvit
- LegendasDivx
- Legendas.net
- Napiprojekt
- Napisy24
- Nekur
- OpenSubtitles.com
- OpenSubtitles.org (VIP)
- Podnapisi
- RegieLive
- Sous-Titres.eu
- Subdivx
- subf2m.co
- Subs.sab.bz
- Subs4Free
- Subs4Series
- Subscene
- Subscenter
- Subsunacs.net
- SubSynchro
- Subtitrari-noi.ro
- subtitri.id.lv
- Subtitulamos.tv
- Supersubtitles
- Titlovi
- Titrari.ro
- Titulky.com
- Turkcealtyazi.org
- TuSubtitulo
- TVSubtitles
- Whisper (requires external web service)
- Wizdom
- XSubs
- Yavka.net
- YIFY Subtitles
- Zimuku

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
subtitle-manager fetch subscene [media] [lang] [output]
subtitle-manager batch [lang] [files...]
subtitle-manager watch opensubtitles [directory] [lang] [-r]
subtitle-manager watch subscene [directory] [lang] [-r]
subtitle-manager grpc-server --addr :50051
subtitle-manager delete [file]
```

The `extract` command accepts `--ffmpeg` to specify a custom ffmpeg binary.

### Web UI

Run `subtitle-manager web` to start the embedded React interface on `:8080`. The SPA is built via `go generate` in the `webui` directory and embedded using Go 1.16's `embed` package.

Configuration values are loaded from `$HOME/.subtitle-manager.yaml` by default. Environment variables prefixed with `SM_` override configuration values. API keys may be specified via flags `--google-key` and `--openai-key` or in the configuration file. The SQLite database location defaults to `$HOME/.subtitle-manager.db` and can be overridden with `--db`.  Translation can be delegated to a remote gRPC server using the `--grpc` flag and providing an address such as `localhost:50051`.

Example configuration:

```
log-level: info
log_levels:
  translate: debug
translate_service: google
```

### Docker

Build a container image using the provided `Dockerfile`:

```bash
$ docker build -t subtitle-manager .
```

Run commands inside the container:

```bash
$ docker run --rm subtitle-manager [command]
```

Prebuilt images are published to the GitHub Container Registry:

```bash
$ docker pull ghcr.io/jdfalk/subtitle-manager:latest
```

## Development

Tests can be run with `go test ./...`.
Continuous integration is provided via a GitHub Actions workflow that verifies formatting, vets code and runs the test suite on each push.

The project aims to eventually reach feature parity with [Bazarr](https://github.com/morpheus65535/bazarr) while offering improved configuration and logging. See `TODO.md` for the full roadmap and implementation plan.
Extensive architectural details and design decisions are documented in
`docs/TECHNICAL_DESIGN.md`. New contributors should review this document to
understand package responsibilities and future plans.
For a detailed list of Bazarr features used as the parity target, see [docs/BAZARR_FEATURES.md](docs/BAZARR_FEATURES.md).

## License

This project is licensed under the terms of the MIT license. See `LICENSE` for details.
