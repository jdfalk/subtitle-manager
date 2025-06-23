<!-- file: docs/API_DESIGN.md -->

# API Design

This document outlines the design philosophy of the REST and gRPC APIs provided by Subtitle Manager.

## REST API Principles

- **Authentication**: All endpoints require either session cookies or API keys sent via `X-API-Key` header.
- **Versioning**: APIs are prefixed with `/api` and may introduce `/v{n}` segments for future breaking changes.
- **Error Handling**: Errors are returned as JSON objects with `error` and `message` fields. HTTP status codes follow standard semantics.
- **Pagination**: List endpoints accept `page` and `limit` query parameters. Responses include `total` and `items` fields.

### Common Endpoints

- `POST /api/login` – authenticate user and start session.
- `GET /api/config` – fetch current configuration.
- `POST /api/convert` – convert uploaded subtitle files to SRT.
- `POST /api/translate` – translate uploaded subtitles to another language.
- `POST /api/extract` – extract embedded subtitles from media files.
- `GET /api/history` – list translation and download history.
- `POST /api/library/scan` – index media files and generate metadata.
- `GET /api/library/scan/status` – check progress for the last library scan.
- `GET /api/wanted` – retrieve wanted subtitles list.

Example error response:

\```json
{
"error": "invalid_request",
"message": "missing language parameter"
}
\```

## gRPC Service

The optional gRPC service defined in `proto/translator.proto` exposes translation operations.

\```
service Translator {
rpc Translate(TranslateRequest) returns (TranslateResponse);
}
\```

- **Authentication** is provided through token metadata headers.
- **Streaming** endpoints may be added in future revisions for batch translation.

## API Development Guidelines

1. Document new endpoints in this file and update the README summary.
2. Maintain backward compatibility when possible. Introduce new versions for breaking changes.
3. Keep request and response schemas concise. Optional fields should have `omitempty` JSON tags.
4. Write integration tests for each handler using the `webserver` package.
