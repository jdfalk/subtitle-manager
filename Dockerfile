# file: Dockerfile
# Build the subtitle-manager Go binary and package it in a container.

FROM golang:1.23 AS builder

# Build arguments for cross-compilation
ARG TARGETOS=linux
ARG TARGETARCH=amd64

WORKDIR /src

# Install build dependencies including cross-compilation tools
RUN apt-get update && apt-get install -y \
    nodejs npm \
    gcc-aarch64-linux-gnu \
    gcc-x86-64-linux-gnu \
    libc6-dev-arm64-cross \
    libc6-dev-amd64-cross \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go generate ./webui

# Build the Go application with proper CGO support
RUN case ${TARGETARCH} in \
    "arm64") \
    export CC=aarch64-linux-gnu-gcc && \
    export CXX=aarch64-linux-gnu-g++ && \
    export CGO_ENABLED=1 && \
    export GOOS=${TARGETOS} && \
    export GOARCH=${TARGETARCH} && \
    go build -o subtitle-manager ./ ;; \
    "amd64") \
    export CC=x86_64-linux-gnu-gcc && \
    export CXX=x86_64-linux-gnu-g++ && \
    export CGO_ENABLED=1 && \
    export GOOS=${TARGETOS} && \
    export GOARCH=${TARGETARCH} && \
    go build -o subtitle-manager ./ ;; \
    *) \
    echo "Unsupported architecture: ${TARGETARCH}" && exit 1 ;; \
    esac

FROM debian:bookworm-slim

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
    ffmpeg \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Create directories for config and media
RUN mkdir -p /config /media

COPY --from=builder /src/subtitle-manager /subtitle-manager

# Set default environment variables
ENV SM_CONFIG_FILE=/config/subtitle-manager.yaml
ENV SM_DB_PATH=/config/db
ENV SM_DB_BACKEND=pebble
ENV SM_SQLITE3_FILENAME=subtitle-manager.db
ENV SM_LOG_LEVEL=info
# Optional: Set these to automatically create admin user on first run
# ENV SM_ADMIN_USER=admin
# ENV SM_ADMIN_PASSWORD=changeme

# Expose the web interface port
EXPOSE 8080

# Default command starts the web interface
ENTRYPOINT ["/subtitle-manager"]
CMD ["web", "--addr", ":8080"]
