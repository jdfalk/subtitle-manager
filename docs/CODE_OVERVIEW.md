# file: docs/CODE_OVERVIEW.md

# Code Overview

This document provides a high level summary of the main source files and how
they work together.

## Directory Structure

- **cmd/** – Cobra commands implementing the CLI interface.
  - `root.go` bootstraps configuration and logging.
  - `convert.go`, `merge.go`, `translate.go` and others expose subtitle
    operations.
  - `scan.go`, `watch.go` and `autoscan.go` provide library management tools.
- **pkg/** – Core application packages used by both CLI and web server.
  - `auth/` handles password, session and OAuth2 logic.
  - `database/` contains SQLite, PebbleDB and PostgreSQL backends.
  - `providers/` implements the 40+ subtitle provider clients.
  - `scheduler/` runs cron based tasks and automated scans.
  - `webserver/` exposes REST endpoints and serves the React UI.
- **webui/** – React front‑end built with Vite.
  - `src/` includes page components such as `Dashboard.jsx`, `Settings.jsx` and
    `MediaLibrary.jsx`.
  - The UI is embedded into the Go binary via `go generate` in `webui/`.

## Interaction Flow

\``` +-------------+ +---------------+ +-------------+ | CLI (cmd) +------>+
Core Packages +------>+ Storage | +-------------+ +---------------+
+-------------+ | ^ ^ | | | v | +-----------+ +------------+ | | Web UI | |
Webserver +---------------+---------------->+ (React) | +------------+
+-----------+ \```

CLI commands and HTTP handlers both rely on the packages in `pkg/` for all
business logic. The `webserver` package exposes the same functionality over REST
which is consumed by the React application.

## Key Files and Responsibilities

| File/Directory      | Purpose                                                                |
| ------------------- | ---------------------------------------------------------------------- |
| `cmd/root.go`       | Entry point for CLI commands. Loads configuration and sets up logging. |
| `cmd/convert.go`    | Converts subtitles between formats using go-astisub.                   |
| `cmd/translate.go`  | Translates subtitles via Google or ChatGPT APIs.                       |
| `cmd/scan.go`       | Scans media libraries for missing subtitles.                           |
| `cmd/watch.go`      | Watches directories and triggers scans.                                |
| `cmd/transcribe.go` | Uses Whisper to transcribe audio tracks into subtitles.                |
| `pkg/auth/`         | Authentication and session management including OAuth2.                |
| `pkg/database/`     | Database backends: SQLite, PebbleDB, PostgreSQL.                       |
| `pkg/providers/`    | 40+ subtitle provider clients.                                         |
| `pkg/scheduler/`    | Cron based task runner used by CLI and web server.                     |
| `pkg/webserver/`    | REST API handlers and static file serving.                             |
| `webui/`            | React front end with pages under `src/`.                               |
| `Dockerfile`        | Multi-stage build producing the container image.                       |

These components interact as shown in the chart above. The CLI and web server
call into the packages under `pkg/`, which manage database access and provider
logic. The React UI communicates with the REST API served by `pkg/webserver/`.
