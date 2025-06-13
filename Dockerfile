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
    export CGO_ENABLED=1 && \
    export GOOS=${TARGETOS} && \
    export GOARCH=${TARGETARCH} && \
    go build -o subtitle-manager ./ ;; \
    *) \
    echo "Unsupported architecture: ${TARGETARCH}" && exit 1 ;; \
    esac

FROM debian:bookworm-slim
COPY --from=builder /src/subtitle-manager /subtitle-manager
ENTRYPOINT ["/subtitle-manager"]
