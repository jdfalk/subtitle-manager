<!-- file: docs/tasks/06-monitoring/TASK-06-003-unified-monitoring-tests.md -->
<!-- version: 1.0.0 -->
<!-- guid: e1c3b8f4-5d6a-7c8b-9a0b-1c2d3e4f5a6b -->

# TASK-06-003: Unified Monitoring Tests

Add tests for the unified monitoring flow, including the CLI command and provider integrations via mocks.

## Scope

- cmd/monitor_unified.go: CLI parsing and wiring
- pkg/monitoring/unified.go: orchestration and provider calls
- Mock provider interfaces to avoid external dependencies

## Done Criteria

- Unit tests cover: CLI flag parsing, provider errors, and normal run
- Mocks verify sequencing and error propagation
- No network or filesystem writes outside `t.TempDir()`

## Subtasks

- [ ] CLI tests
  - [ ] Validate flags and defaults
  - [ ] Invalid input triggers helpful error

- [ ] Orchestration tests
  - [ ] Successful provider run emits events/metrics
  - [ ] Provider error bubbles correctly and is logged

## References

- `pkg/monitoring/unified.go`
- `cmd/monitor_unified.go`
