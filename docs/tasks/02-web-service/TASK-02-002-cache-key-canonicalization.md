<!-- file: docs/tasks/02-web-service/TASK-02-002-cache-key-canonicalization.md -->
<!-- version: 1.0.0 -->
<!-- guid: 4f6d2b3a-2b11-4a8c-9f7d-2a8a5c3d9e21 -->

# TASK-02-002: Canonicalize provider order in cache keys

Normalize provider order when generating search cache keys so that permutations of the same set produce identical keys. Apply to both API and CLI paths. Track with issue #1885; duplicates: #1685, #1547, #1690.

## Why

- Prevent redundant cache entries for the same logical query
- Improve hit ratio and reduce storage

## Scope

- API: Manual search handler cache key generation
- CLI: `search` command cache key generation
- Shared helper for canonicalization (sort, dedupe, case-normalize if needed)

## Acceptance criteria

- Given providers ["A", "B"] and ["B", "A"], cache keys are identical
- Duplicates are removed (e.g., ["A", "A", "B"] → ["A", "B"]) before keying
- Unit tests cover typical and edge cases (empty list, single provider, many providers)
- Backward-compatibility note added to docs if key format changes

## Test plan

- Add unit tests for helper function: permutations → equal keys
- API handler tests: two requests with swapped provider order → second is a cache hit
- CLI tests: repeated runs with swapped order → cache hit indicated

## Implementation notes

- Prefer a small pure function `CanonicalizeProviders([]string) []string`
- Ensure consistent joiner and version prefix in keys to allow future migrations
