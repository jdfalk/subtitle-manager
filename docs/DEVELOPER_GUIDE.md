<!-- file: docs/DEVELOPER_GUIDE.md -->

# Developer Guide

This guide helps new contributors get started with Subtitle Manager.

## Prerequisites

### Option 1: Local Development
- Go 1.22 or later
- Node.js 20 for the React web UI
- Docker for optional container builds

### Option 2: Dev Container (Recommended)
- [Visual Studio Code](https://code.visualstudio.com/)
- [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
- [Docker Desktop](https://www.docker.com/products/docker-desktop) or Docker Engine

## Setup

### Using Dev Container (Recommended)

The easiest way to get started is using the provided development container:

1. Clone the repository
2. Open the repository in VS Code
3. When prompted, click "Reopen in Container" or run the command palette: `Dev Containers: Reopen in Container`
4. Wait for the container to build and the post-create script to run
5. Start developing!

The dev container includes:
- Go 1.24+ with development tools
- Node.js 18.19 for React development
- FFmpeg for subtitle processing
- SQLite with CGO support enabled
- Pre-configured VS Code extensions
- All necessary build tools

See [.devcontainer/README.md](.devcontainer/README.md) for detailed documentation.

### Local Development Setup

### Local Development Setup

1. Clone the repository and install dependencies: \```bash git clone
   https://github.com/jdfalk/subtitle-manager.git cd subtitle-manager go mod
   download npm --prefix webui install \```
2. Build the web assets and compile the binary: \```bash go generate ./webui go
   build \```
3. Run tests to verify your environment: \```bash make test-all \```

## Contributing

- Follow the commit message guidelines in `AGENTS.md`.
- Document new functions and packages with Go comments.
- Update relevant design documents when changing architecture.
- Use the issue update scripts to reference your work.

## Development Workflow

- Branch from `main` and keep your branch up to date using the rebase command in
  `AGENTS.md`.
- Run `go fmt` and `goimports` before committing, or use the provided pre-commit
  hook: \```bash ./scripts/install-pre-commit-hooks.sh \```
- Submit pull requests with concise descriptions and link related issues.
