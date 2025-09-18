<!-- file: docs/tasks/02-web-service/TASK-02-003-cached-provider-integration-tests.md -->
<!-- version: 1.0.0 -->
<!-- guid: 7a1c9e2f-8d3b-4b5f-8d2b-2f4c6b8e1a23 -->

# TASK-02-003: Integration tests for cached provider results

Implement integration tests validating manual search cache behavior across providers. Tracks issue #1884.

## Scope

- Verify cache hit on identical query with same provider set
- Verify cache miss after relevant setting/provider changes
- Validate TTL behavior if configured
- Exercise API endpoints (/api/search) and CLI where feasible

## Acceptance criteria

- Tests pass reliably in CI with seeded providers/mocks
- Coverage includes at least two providers with distinct result shapes
- Documentation updated (testing guide) with how to run and interpret results

## Test strategy

- Use in-memory cache manager with controlled TTL
- Mock provider responses; first call populates cache, subsequent call reads from cache
- Change provider configuration mid-test to trigger invalidation and assert miss

## Notes

- Keep provider mocks deterministic; avoid hitting external services
- Use t.TempDir() for any file I/O; avoid global state leakage
