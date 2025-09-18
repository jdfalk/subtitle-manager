<!-- file: docs/tasks/10-testing-quality/TASK-10-002-webserver-coverage.md -->
<!-- version: 1.0.0 -->
<!-- guid: 3f9a7b2c-1d4e-4a9f-8c6b-4a2f7d1e9c3a -->

# TASK-10-002: Raise Webserver Test Coverage

Focus: systematically increase coverage for pkg/webserver by testing happy/error paths and auth edge cases.

## Scope

- Handlers: auth, cache, system, widgets, oauth, oauth_management
- Middlewares: auth extraction, request logging (if present)
- Config: ensure CI-safe temp dirs and deterministic behavior

## Done Criteria

- Coverage for pkg/webserver >= 80%
- All external calls mocked or isolated; no network dependence
- Error paths validated (invalid headers, missing params, JSON errors)

## Subtasks

- [ ] Auth headers
  - [ ] Ensure API key is passed via `key.GetId()` and session via `session.GetId()`
  - [ ] Missing/invalid token returns 401/403 consistently

- [ ] Cache endpoints
  - [ ] Health returns `{"status":"up"|"down"}` and reflects failures
  - [ ] Typed clear operates on targeted namespaces only

- [ ] System endpoints
  - [ ] Version/build info present and JSON parseable
  - [ ] Negative path for corrupt config handled gracefully

- [ ] Widgets endpoints
  - [ ] Validate JSON structure and content types
  - [ ] Error path for invalid query parameters

## Notes

- Use `t.TempDir()` for any filesystem interactions
- Table-driven tests for parameterized cases
- Keep response fixtures minimal and assert only what matters
