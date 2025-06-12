# file: Dockerfile
# Build the subtitle-manager Go binary and package it in a container.

FROM golang:1.23 AS builder
WORKDIR /src
RUN apt-get update && apt-get install -y nodejs npm
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go generate ./webui
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o subtitle-manager ./

FROM gcr.io/distroless/static
COPY --from=builder /src/subtitle-manager /subtitle-manager
ENTRYPOINT ["/subtitle-manager"]
