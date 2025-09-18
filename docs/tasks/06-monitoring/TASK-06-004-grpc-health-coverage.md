<!-- file: docs/tasks/06-monitoring/TASK-06-004-grpc-health-coverage.md -->
<!-- version: 1.0.0 -->
<!-- guid: 9c2f5e1a-1b3c-4d5e-8f0a-6a2b4c7d9e12 -->

# TASK-06-004: gRPC health coverage using gcommon

Increase test coverage for health endpoints by wiring and testing gRPC health service consistent with gcommon/health. Related to #1771.

## Goals

- Confirm HealthProvider initialization flows through gRPC registration path
- Add minimal client check using bufconn to hit the gRPC health Check

## Acceptance criteria

- New tests validate gRPC health responds with SERVING when initialized
- HTTP health remains covered (status "up")
- No external network access; all in-memory

## Implementation notes

- If a minimal gRPC server bootstrap isn’t present, add a small helper exclusively for tests
- Keep surface area small; avoid changing production paths unless necessary
