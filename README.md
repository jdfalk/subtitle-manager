# Subtitle Manager

Subtitle Manager is a comprehensive subtitle management application written in Go that provides both CLI and web interfaces for converting, translating, and managing subtitle files. The project aims to reach feature parity with Bazarr. Some advanced features such as scheduling and webhooks are still in progress.

## ✨ Key Highlights

- 🎯 **Production Ready**: Complete with authentication, RBAC, and 40+ subtitle providers
- 🌐 **Full Web UI**: React-based interface with real-time scanning and configuration
- 🔐 **Enterprise Auth**: Password, OAuth2, API keys, and role-based access control
- 🚀 **High Performance**: Concurrent processing with worker pools and gRPC support
- 📦 **Container Ready**: Docker images published to GitHub Container Registry

## Features

- Convert subtitles from many formats to SRT.
- Merge two subtitle tracks sorted by start time.
- Translate subtitles via Google Translate or ChatGPT APIs.
- Store translation history in an SQLite database or optional PebbleDB store. Retrieve history via the `history` command or `/api/history` endpoint.
- Per component logging with adjustable levels.
- Extract subtitles from media containers using ffmpeg.
- Convert uploaded subtitle files to SRT via `/api/convert`.
- Transcribe audio tracks to subtitles via Whisper.
- Download subtitles from a comprehensive list of providers based on Bazarr,
  including Addic7ed, AnimeKalesi, Animetosho, Assrt, Avistaz, BetaSeries,
  BSplayer, GreekSubs, Podnapisi, Subscene, TVSubtitles, Titlovi, LegendasDivx
  and many more.
- Batch translate multiple files concurrently.
- Monitor directories and automatically download subtitles.
- Scan existing libraries and fetch missing or upgraded subtitles.
- Schedule periodic scans with the `autoscan` command.
- Parse file names and retrieve movie or episode details from TheMovieDB.
- High performance scanning using concurrent workers.
- Recursive directory watching with -r flag.
- Integrate with Sonarr, Radarr and Plex using dedicated commands.
- Run a translation gRPC server.
- Delete subtitle files and remove history records.
- Track subtitle download history and list with `downloads` command or `/api/history`.
- GitHub OAuth2 login enabled via `/api/oauth/github` endpoints.
- Manually search for subtitles with `search` command.
- Provider registry simplifies adding new sources.
- Dockerfile and workflow for container builds.
- Prebuilt images published to GitHub Container Registry.
- Integrated authentication system with password, token, OAuth2 and API key support.
- Generate one time login tokens using `user token` and authenticate with `login-token`.
- Minimal React web UI with login page.
- Role based access control with sensible defaults and session storage in the database.
- Manage accounts with `user add`, `user role`, `user token` and `user list` commands.

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

subtitle-manager downloads
github_client_id: yourClientID
github_client_secret: yourClientSecret
github_redirect_url: http://localhost:8080/api/oauth/github/callback
## Installation

```bash
# clone repository
$ git clone <this repository>
$ cd subtitle-manager
# install dependencies, build web UI and compile
$ go generate ./webui
$ go build
```

## Usage

```
subtitle-manager convert [input] [output]
subtitle-manager merge [sub1] [sub2] [output]
subtitle-manager translate [input] [output] [lang]
subtitle-manager history
subtitle-manager extract [media] [output]
subtitle-manager transcribe [media] [output] [lang]
subtitle-manager fetch opensubtitles [media] [lang] [output]
subtitle-manager fetch subscene [media] [lang] [output]
subtitle-manager search opensubtitles [media] [lang]
subtitle-manager batch [lang] [files...]
subtitle-manager scan opensubtitles [directory] [lang] [-u]
subtitle-manager scan subscene [directory] [lang] [-u]
subtitle-manager autoscan [provider] [directory] [lang] [-i duration] [-u]
subtitle-manager scanlib [directory]
subtitle-manager watch opensubtitles [directory] [lang] [-r]
subtitle-manager watch subscene [directory] [lang] [-r]
subtitle-manager grpc-server --addr :50051
subtitle-manager delete [file]
subtitle-manager downloads
subtitle-manager login [username] [password]
subtitle-manager login-token [token]
subtitle-manager user add [username] [email] [password]
subtitle-manager user apikey [username]
subtitle-manager user token [email]
subtitle-manager user role [username] [role]
subtitle-manager user list
```

The `extract` command accepts `--ffmpeg` to specify a custom ffmpeg binary.

### Web UI

Run `subtitle-manager web` to start the embedded React interface on `:8080`. The SPA is built via `go generate` in the `webui` directory and embedded using Go 1.16's `embed` package.

The interface now exposes a **Settings** page that mirrors Bazarr's options. Configuration values are retrieved from `/api/config` and saved back via a `POST` to the same endpoint. Any changes are written to the active YAML configuration file.

Configuration values are loaded from `$HOME/.subtitle-manager.yaml` by default. Environment variables prefixed with `SM_` override configuration values. API keys may be specified via flags `--google-key`, `--openai-key` and `--opensubtitles-key` or in the configuration file. Additional flags include `--ffmpeg-path` for a custom ffmpeg binary, `--batch-workers` and `--scan-workers` to control concurrency. The SQLite database location defaults to `$HOME/.subtitle-manager.db` and can be overridden with `--db`. Use `--db-backend` to switch between the `sqlite` and `pebble` engines. When using PebbleDB a directory path may be supplied instead of a file. Translation can be delegated to a remote gRPC server using the `--grpc` flag and providing an address such as `localhost:50051`. Generic provider options may also be set with variables like `SM_PROVIDERS_GENERIC_API_URL`.
Run `subtitle-manager migrate old.db newdir` to copy existing subtitle history from SQLite to PebbleDB.

Example configuration:

```
log-level: info
log_levels:
  translate: debug
translate_service: google
google_api_key: your-google-api-key
openai_api_key: your-openai-api-key
google_api_url: https://translation.googleapis.com/language/translate/v2
openai_model: gpt-3.5-turbo
ffmpeg_path: /usr/bin/ffmpeg
batch_workers: 4
scan_workers: 4
opensubtitles:
  api_key: your-os-key
  api_url: https://rest.opensubtitles.org
  user_agent: subtitle-manager/0.1
providers:
  generic:
    api_url: https://example.com/subtitles
    username: myuser
    password: secret
    api_key: token123
github_client_id: yourClientID
github_client_secret: yourClientSecret
github_redirect_url: http://localhost:8080/api/oauth/github/callback
```

### Docker

Build a container image using the provided `Dockerfile`:

```bash
$ docker build -t subtitle-manager .
```

The Docker build runs `go generate ./webui` so the final image contains the latest
compiled web assets.

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
Web UI unit tests live in `webui/src/__tests__` and are executed with `npm test` from the `webui` directory.
End-to-end tests use Playwright and run with `npm run test:e2e` once browsers are installed via `npx playwright install`.
Continuous integration is provided via a GitHub Actions workflow that verifies formatting, vets code and runs the test suite on each push.

### Issue updates

Pushing an `issue_updates.json` file to the repository root allows the `update-issues` workflow to create, update or delete GitHub issues using the repository `GITHUB_TOKEN`. The file contains a JSON array where each object specifies an `action` and any supported issue fields. Duplicate issues are avoided by checking for an existing issue with the same title before creation. Example:

```json
[
  { "action": "create", "title": "New issue", "body": "Details" },
  { "action": "update", "number": 2, "state": "closed" }
]
```

The workflow runs on every push to `main` and processes the listed operations. After all entries are handled the file is removed on a new branch and a pull request is opened so that the cleanup can be merged back to `main`.

### Duplicate ticket cleanup

The `close-duplicates` workflow runs daily and on demand to detect open issues
with the same title. The script chooses the lowest numbered ticket as the
canonical reference and automatically closes the rest with a comment noting the
duplicate. This keeps the issue tracker focused on a single discussion for each
problem.

The project aims to eventually reach feature parity with [Bazarr](https://github.com/morpheus65535/bazarr) while offering improved configuration and logging. See `TODO.md` for the full roadmap and implementation plan.
Extensive architectural details and design decisions are documented in
`docs/TECHNICAL_DESIGN.md`. New contributors should review this document to
understand package responsibilities and future plans.
For a detailed list of Bazarr features used as the parity target, see [docs/BAZARR_FEATURES.md](docs/BAZARR_FEATURES.md).
Instructions for importing an existing Bazarr configuration are documented in [docs/BAZARR_SETTINGS_SYNC.md](docs/BAZARR_SETTINGS_SYNC.md).

## License

This project is licensed under the terms of the MIT license. See `LICENSE` for details.
