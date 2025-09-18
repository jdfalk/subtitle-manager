<!-- file: docs/tasks/01-core-architecture/TASK-01-002-gcommon-adoption.md -->
<!-- version: 1.0.0 -->
<!-- guid: 7b2a4f9e-0e58-41e4-9a58-0a2d1c3f4b7a -->

# TASK-01-002: Complete gcommon Adoption

Adopt gcommon types and patterns consistently across services, replacing custom types and ensuring opaque ID usage.

## Scope

- Protobuf imports use `gcommon/v1/*` exclusively
- Go code uses generated SDKs from gcommon
- Opaque types respected: use getters (e.g., `GetId()`) not direct fields

## Done Criteria

- buf lint/generate passes with no custom duplicates
- All references updated to gcommon types in code and tests
- CI passes; integration tests align to gcommon

## Subtasks

- [ ] Proto imports
  - [ ] Replace any custom `proto/common/v1/*` with `gcommon/v1/common/*`
  - [ ] Health/config imports align with gcommon locations

- [ ] Go code updates
  - [ ] Replace direct struct access with getters on opaque types
  - [ ] Update service interfaces and handlers to gcommon messages

- [ ] Tests
  - [ ] Update helpers to extract IDs via getters
  - [ ] Validate error/happy paths with new message shapes

## References

- `docs/tasks/GCOMMON-CORRECTION.md`
- gcommon repo: proto and sdks structure
