<!-- file: docs/tasks/10-testing-quality/TASK-10-001-testing-roadmap.md -->
<!-- version: 1.0.1 -->
<!-- guid: a2b4c1d2-3e4f-5a6b-7c8d-9e0f1a2b3c4d -->

# TASK-10-001: Testing & Quality Roadmap

This document tracks the concrete, bite-sized testing tasks to sustain reliability and raise coverage.

## Goals

- Keep CI green across gcommon integration (opaque types, new protos)
- Increase meaningful coverage around webserver, services, and monitoring
- Make tests faster, more hermetic, and better isolated

## Task List

- [ ] Webserver: API key header handling
  - Ensure all tests derive header token via `apiKey.GetId()` and `session.GetId()`
  - Validate 401/403 paths for missing/invalid API keys

- [ ] Health endpoints: consistency
  - Align HTTP JSON responses with `HealthProvider` results (status strings: `up`/`down`)
  - Add negative-path tests (dependency errors, cache failures)

- [ ] Monitoring unification
  - Add unit tests for `monitor_unified` command/options
  - Mock providers and assert event/metrics emission

- [ ] Services (gRPC): registry wiring
  - Expand `registry_test.go` to verify registration of all services and reflection
  - Add bufconn E2E test for `WebService` happy-path and error-path

- [ ] Queue: adopt gcommon QueueMessage
  - Assert ID casing (`ID`) and payload mapping
  - Include a boundary test for large payloads or missing fields

- [ ] Security: file system safety
  - Ensure tests use `t.TempDir()` for any path creation to be CI-safe

- [ ] Coverage reporting
  - Add a CI step to output coverage HTML as artifact for quick inspection

## References

- `docs/tasks/GCOMMON-CORRECTION.md`
- `pkg/webserver/*_test.go`
- `pkg/services/*_test.go`
- `pkg/monitoring/*`
