<!-- file: AGENTS.md -->
<!-- version: 2.3.0 -->
<!-- guid: 2e7c1a4b-5d3f-4b8c-9e1f-7a6b2c3d4e5f -->

# AGENTS.md

This file consolidates the canonical Copilot/AI agent instructions for this
repository. It replaces the old pointer-only content with a practical, single
source of operational guidance pulled from `.github/instructions/*.md` and the
linked instruction files.

## Repository Context

- This repository is a workflow infrastructure hub for reusable GitHub Actions,
  scripts, and configuration shared across multiple repositories.
- Key areas: reusable workflows (`.github/workflows/reusable-*.yml`), script
  library (`scripts/`), instruction system (`.github/instructions/`), and
  protobuf tooling (`tools/`, `scripts/`).
- Focus on workflow reliability, protobuf tooling, and cross-repo sync.

## Operating Priorities (MANDATORY)

1. **Check current state first**: Before creating files or running operations,
   verify the state and make work idempotent.
2. **Use VS Code tasks first** for non-git operations. Check `logs/` after
   task runs for details.
3. **Git operations**: Prefer MCP GitHub tools. If unavailable, use
   `copilot-agent-util` / `safe-ai-util` utilities. Use native git only as a
   last resort.
4. **Prefer Python for scripts** unless the task is trivial; shell scripts are
   only for simple operations. Convert complex shell logic to Python.
5. **No legacy doc-update scripts**: Do not use
   `create-doc-update.sh`, `doc_update_manager.py`, or `.github/doc-updates/`.

## File Headers and Versioning (CRITICAL)

- All files with headers must include `file`, `version`, and `guid`.
- **Always bump the version** when editing a file with a version header:
  - Patch: typo/formatting/fix.
  - Minor: new features/sections.
  - Major: breaking changes/structural overhaul.

## Git Commit Standards (MANDATORY)

- Conventional commit header required: `type(scope): description`.
- Only include issue numbers when working on a real issue.
- Body must document changes and list **all modified files** with descriptions.
- For multiple logical changes, include multiple conventional headers in the
  body under a **Changes Made** section.

## Pull Request Description Standards (MANDATORY)

Use this structure:

```markdown
## Summary

## Changes Made

### type(scope): description
**Description:** ...
**Files Modified:**
- [`path/to/file`](./path/to/file) - ... | [[diff]](...) [[repo]](...)

## Testing

## Breaking Changes

## Additional Notes

## Related Issues
```

- Each distinct change gets its own subsection with a conventional header.
- Include file links, diff links, and repo links for each file.

## Security Requirements

- **Least privilege** permissions for workflows.
- **Never** hardcode secrets; use GitHub Secrets and prefer OIDC.
- Validate inputs, environment variables, file paths, and URLs.
- Pin action versions and dependencies; avoid `@main`/`@master`.
- Use secure triggers and environment protections.
- Apply container hardening: minimal base images, non-root users, scanning,
  signing, attestations.
- Shell scripts must use `set -euo pipefail` and quote variables.

## Testing Requirements

- Follow Arrange-Act-Assert.
- Use descriptive test names (e.g., `testUnitOfWork_State_Expected`).
- Test one behavior per test; cover edge cases.
- Prefer mocks/stubs for external dependencies.
- Keep tests deterministic and isolated.

## GitHub Actions Workflow Rules

- Use `.github/workflows/` and descriptive, lowercase, hyphenated filenames.
- 2 spaces indentation, reasonable line length, and clear comments.
- Jobs and steps require human-readable names.
- Pin actions to versions (SHA for critical workflows).
- Prefer `run: |` for multi-line scripts.
- Use `env` at the appropriate scope.

## Protobuf Rules (CRITICAL)

- **Edition 2023 required**: `edition = "2023";` as first non-comment line.
- **1-1-1 pattern required**: one message/enum/service per file.
- Use module prefixes for all messages (e.g., `AuthUserInfo`).
- Filenames are `lower_snake_case.proto`, and services use
  `{service_name}_service.proto`.

## Language-Specific Rules (High-Level)

### Markdown
- One H1 per file, consistent heading hierarchy.
- Use code fences with language identifiers.
- Line length 80–100, no trailing whitespace.
- Required Markdown header (file/version/guid).

### Go
- **Go 1.23.0+ only** (update `go.mod`/`go.work` accordingly).
- Follow Google Go Style Guide.
- Prefer clarity and simplicity; required file header.

### Python
- **Python 3.13.0+ only** (update `requirements.txt` / `pyproject.toml`).
- Use absolute imports; 4 spaces; 80-char lines.
- Google-style docstrings; type annotations when applicable.
- Required file header after shebang.

### JavaScript
- 2 spaces; semicolons; single quotes; 80-char lines.
- Use ES modules, `const`/`let`, and JSDoc where needed.
- Required file header.

### TypeScript
- Use ES imports, explicit types, and interfaces for object shapes.
- Prefer optional fields over `|undefined`.
- Avoid `any` and unsafe type assertions.
- Required file header.

### Rust
- Use `rustfmt`/`clippy`; 4-space indentation; 100-char lines.
- Document public items with `///`; prefer `Result` over panics.
- Required file header.

### Shell
- Use bash/sh with shebang and header.
- `set -euo pipefail`, quote variables, avoid heredocs unless last resort.

### R
- snake_case, 2-space indent, explicit `return()`.
- Avoid `attach()` and right-hand assignment.
- Required file header after shebang for scripts.

### HTML/CSS
- HTML5 doctype, semantic markup, UTF-8 meta tag.
- 2 spaces indentation; lowercase tags/attributes; double quotes.
- Required file header.

### JSON/JSONC
- `.json` must be valid JSON (no comments); camelCase keys; 2 spaces.
- `.jsonc` requires the standard header; `.json` does not.

## Utilities and Tooling

- Prefer `copilot-agent-util` or `safe-ai-util` for operations when tasks/MCP
  tools are unavailable.
- Use `copilot-util-args` or `safe-ai-util-args` for shared configuration.

