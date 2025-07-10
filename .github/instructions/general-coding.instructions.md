## <!-- file: .github/instructions/general-coding.instructions.md -->

applyTo: "\*\*" description: | General coding, documentation, and workflow rules
for all Copilot/AI agents and VS Code Copilot customization. These rules apply
to all files and languages unless overridden by a more specific instructions
file. For details, see the main documentation in
`.github/copilot-instructions.md`.

---

# General Coding Instructions

These instructions are the canonical source for all Copilot/AI agent coding,
documentation, and workflow rules in this repository. They are referenced by
language- and task-specific instructions, and are always included by default in
Copilot customization.

- Follow the [commit message standards](../commit-messages.md) and
  [pull request description guidelines](../pull-request-descriptions.md).
- All language/framework-specific style and workflow rules are now found in
  `.github/instructions/*.instructions.md` files. These are the only canonical
  source for code style, documentation, and workflow rules for each language or
  framework.
- Document all code, classes, functions, and tests extensively, using the
  appropriate style for the language.
- Use the Arrange-Act-Assert pattern for tests, and follow the
  [test generation guidelines](../test-generation.md).
- For agent/AI-specific instructions, see [AGENTS.md](../AGENTS.md) and related
  files.
- Do not duplicate rules; reference this file from more specific instructions.
- For VS Code Copilot customization, this file is included via symlink in
  `.vscode/copilot/`.

For more details and the full system, see
[copilot-instructions.md](../copilot-instructions.md).

## Required File Header (File Identification)

All source, script, and documentation files MUST begin with a standard header
containing:

- The exact relative file path from the repository root (e.g.,
  `# file: path/to/file.py`)
- The file's semantic version (e.g., `# version: 1.0.0`)
- The file's GUID (e.g., `# guid: 123e4567-e89b-12d3-a456-426614174000`)

**Header format varies by language/file type:**

- **Markdown:**
  ```markdown
  <!-- file: path/to/file.md -->
  <!-- version: 1.0.0 -->
  <!-- guid: 123e4567-e89b-12d3-a456-426614174000 -->
  ```
- **Python:**
  ```python
  #!/usr/bin/env python3
  # file: path/to/file.py
  # version: 1.0.0
  # guid: 123e4567-e89b-12d3-a456-426614174000
  ```
- **Go:**
  ```go
  // file: path/to/file.go
  // version: 1.0.0
  // guid: 123e4567-e89b-12d3-a456-426614174000
  ```
- **JavaScript/TypeScript:**
  ```js
  // file: path/to/file.js
  // version: 1.0.0
  // guid: 123e4567-e89b-12d3-a456-426614174000
  ```
- **Shell (bash/sh):**
  ```bash
  #!/bin/bash
  # file: path/to/script.sh
  # version: 1.0.0
  # guid: 123e4567-e89b-12d3-a456-426614174000
  ```
  (Header must come after the shebang line)
- **Protobuf:**
  ```protobuf
  // file: path/to/file.proto
  // version: 1.0.0
  // guid: 123e4567-e89b-12d3-a456-426614174000
  ```
- **CSS:**
  ```css
  /* file: path/to/file.css */
  /* version: 1.0.0 */
  /* guid: 123e4567-e89b-12d3-a456-426614174000 */
  ```
- **R:**
  ```r
  # file: path/to/file.R
  # version: 1.0.0
  # guid: 123e4567-e89b-12d3-a456-426614174000
  ```
- **JSON:**
  ```jsonc
  // file: path/to/file.json
  // version: 1.0.0
  // guid: 123e4567-e89b-12d3-a456-426614174000
  ```
- **TOML:**
  ```toml
  [section]
  # file: path/to/file.toml
  # version: 1.0.0
  # guid: 123e4567-e89b-12d3-a456-426614174000
  ```
  (Header must be inside a section as TOML doesn't support top-level comments)

**All files must include this header in the correct format for their type.**

## Documentation Update System

When making documentation updates to `README.md`, `CHANGELOG.md`, `TODO.md`, or
other documentation files, use the automated documentation update system instead
of direct edits:

### Creating Documentation Updates

1. **Use the script**: Always use `scripts/create-doc-update.sh` to create
   documentation updates
2. **Available modes**:
   - `append` - Add content to end of file
   - `prepend` - Add content to beginning of file
   - `replace-section` - Replace specific section
   - `changelog-entry` - Add properly formatted changelog entry
   - `task-add` - Add task to TODO list
   - `task-complete` - Mark task as complete

### Examples

```bash
# Add a new changelog entry
./scripts/create-doc-update.sh --template changelog-feature "Added user authentication system"

# Add a TODO task with high priority
./scripts/create-doc-update.sh TODO.md "Implement OAuth2 integration" task-add --priority HIGH

# Update a specific section
./scripts/create-doc-update.sh README.md "Updated installation instructions" replace-section --section "Installation"

# Interactive mode for complex updates
./scripts/create-doc-update.sh --interactive
```

### Processing Updates

- Updates are stored as JSON files in `.github/doc-updates/`
- The workflow `docs-update.yml` automatically processes these files
- Processed files are moved to `.github/doc-updates/processed/`
- Changes can be made via direct commit or pull request

### Benefits

- **Consistency**: Standardized formatting across all documentation
- **Traceability**: Each update has a GUID and timestamp
- **Automation**: Reduces manual errors and ensures proper formatting
- **Conflict Resolution**: Multiple agents can create updates simultaneously

**Always use this system for documentation updates instead of direct file
edits.**

## Script Idempotency Requirements

**All scripts and automation tools MUST be idempotent** - they should produce
the same result when run multiple times without creating duplicates or causing
conflicts.

### Idempotency Guidelines:

- **Check before create**: Always verify if resources (files, projects, issues,
  etc.) already exist before attempting to create them
- **Use unique identifiers**: When possible, use unique identifiers (names, IDs,
  etc.) to detect existing resources
- **Graceful handling**: Handle existing resources gracefully - either skip
  creation with a message or update if needed
- **Atomic operations**: Group related operations so they can be safely retried
  as a unit
- **State validation**: Include validation steps to verify the system is in the
  expected state before proceeding
- **Clear feedback**: Provide clear output indicating whether resources were
  created, found existing, or updated

### Examples of Idempotent Patterns:

```bash
# Good: Check before create
if ! gh project list --owner "$ORG" --format json | jq -e ".projects[] | select(.title == \"$title\")" >/dev/null; then
    gh project create --owner "$ORG" --title "$title"
    echo "✅ Created project: $title"
else
    echo "✅ Found existing project: $title"
fi

# Good: Use conditional logic
mkdir -p directory_name  # -p flag makes it idempotent

# Good: Check file existence
if [[ ! -f "config.json" ]]; then
    create_config_file
fi
```

### Anti-patterns to Avoid:

- Creating resources without checking if they exist
- Assuming clean state on every run
- Failing when encountering existing resources
- Not providing status feedback about what was created vs. found

**This requirement applies to all scripts, GitHub Actions, and automation
tools.**
