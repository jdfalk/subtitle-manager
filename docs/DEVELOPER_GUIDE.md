<!-- file: docs/DEVELOPER_GUIDE.md -->
# Developer Guide

This guide helps new contributors get started with Subtitle Manager.

## Prerequisites

- Go 1.22 or later
- Node.js 20 for the React web UI
- Docker for optional container builds

## Setup

1. Clone the repository and install dependencies:
   \```bash
   git clone https://github.com/jdfalk/subtitle-manager.git
   cd subtitle-manager
   go mod download
   npm --prefix webui install
   \```
2. Build the web assets and compile the binary:
   \```bash
   go generate ./webui
   go build
   \```
3. Run tests to verify your environment:
   \```bash
   make test-all
   \```

## Contributing

- Follow the commit message guidelines in `AGENTS.md`.
- Document new functions and packages with Go comments.
- Update relevant design documents when changing architecture.
- Use the issue update scripts to reference your work.

## Development Workflow

- Branch from `main` and keep your branch up to date using the rebase command in `AGENTS.md`.
- Run `go fmt` and `goimports` before committing, or use the provided pre-commit hook:
  \```bash
  ./scripts/install-pre-commit-hooks.sh
  \```
- Submit pull requests with concise descriptions and link related issues.

