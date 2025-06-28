<!-- file: docs/TEST_DESIGN.md -->

# Test Design

Subtitle Manager uses multiple test layers to ensure code reliability.

## Unit Tests

- Located alongside packages using the `_test.go` naming convention.
- External services are mocked with the `httptest` package or custom mock
  implementations.
- Database tests use in-memory SQLite or Docker-based PostgreSQL when available.
- Coverage focuses on edge cases and error handling paths.

## Integration Tests

- CLI behaviour is tested by invoking commands via `os/exec` and inspecting
  output files.
- gRPC and REST endpoints are verified using a local web server started in
  tests.
- Integration tests run in CI and require no external network access.

## Web UI Tests

- React components have Jest and React Testing Library unit tests under
  `webui/src/__tests__`.
- End-to-end tests are written with Playwright and reside in `webui/tests`.
- E2E tests cover typical user workflows such as login, scanning, and settings
  management.

## Performance Tests

Benchmark tests measure subtitle merging and translation latency. Run with
`go test -bench . ./...`.

## Test Automation

The `Makefile` provides `make test-all` to run Go and web UI tests. CI workflows
execute the same command and upload coverage reports when possible.
