<!-- file: AGENTS.md -->
<!-- version: 1.1.0 -->
<!-- guid: 2e7c1a4b-5d3f-4b8c-9e1f-7a6b2c3d4e5f -->

# AGENTS.md

---

## 0. Creating GitHub Issues via Script (v1.1.0)

- **All new GitHub issues, comments, updates, and closes must be created by
  generating a JSON file using the `scripts/create-issue-update.sh` script.**
- Do **not** create or edit `.github/issue-updates/*.json` files by hand. Always
  use the script to ensure GUID uniqueness and correct formatting.
- The script generates a random GUID and legacy GUID, and places the JSON file
  in `.github/issue-updates/`.
- **Usage examples:**

  ```bash
  # Create a new issue
  ./scripts/create-issue-update.sh create "Issue Title" "Issue body" "label1,label2"

  # Update an existing issue
  ./scripts/create-issue-update.sh update 123 "Updated body" "label1,label2"

  # Add a comment to an issue
  ./scripts/create-issue-update.sh comment 123 "Comment text"

  # Close an issue
  ./scripts/create-issue-update.sh close 123 "completed"
  ```

- The script will:
  - Ensure the `.github/issue-updates/` directory exists
  - Generate a unique GUID and legacy GUID for each file
  - Check for GUID collisions
  - Output a JSON file with the correct structure for the unified issue
    management workflow
- **Best practices:**
  - Use clear, concise titles and bodies
  - Use comma-separated labels (no spaces)
  - Always check the script output for success and file path
  - Do not manually edit or move generated files
  - Let the workflow process and archive files as needed
- For more details, see the script header in
  [`scripts/create-issue-update.sh`](scripts/create-issue-update.sh)

---

## 1. General Agent Workflow

- The user will provide a task.
- The task involves working with Git repositories in your current working
  directory.
- Wait for all terminal commands to be completed (or terminate them) before
  finishing.

### File Header Requirements

Every file MUST begin with a standard header using the appropriate comment
syntax for the file type:

```bash
# Shell scripts (.sh, .bash)
#!/bin/bash
# file: scripts/example.sh
# version: 1.0.0
# guid: unique-identifier
```

```go
// Go files (.go)
// file: cmd/example.go
// version: 1.0.0
// guid: unique-identifier
```

```python
# Python files (.py)
#!/usr/bin/env python3
# file: scripts/example.py
# version: 1.0.0
# guid: unique-identifier
```

```html
<!-- HTML/Markdown files (.html, .md) -->
<!-- file: docs/example.md -->
<!-- version: 1.0.0 -->
<!-- guid: unique-identifier -->
```

```css
/* CSS files (.css) */
/* file: styles/example.css */
/* version: 1.0.0 */
/* guid: unique-identifier */
```

```sql
-- SQL files (.sql)
-- file: database/example.sql
-- version: 1.0.0
-- guid: unique-identifier
```

**Important Notes:**

- For shell scripts and Python files, the shebang line must be first, followed
  immediately by the file header
- Use the appropriate comment syntax for each language (`#`, `//`, `/* */`,
  `<!-- -->`, `--`, etc.)
- The header must include: file path, version, and GUID
- This header helps with file identification and change tracking

### Git Instructions

- Use git to commit your changes
- If pre-commit fails, fix issues and retry
- Check git status to confirm your commit. You must leave your worktree in a
  clean state

### Commit Message Format (REQUIRED)

```text
<type>[scope]: <description>

<body with motivation and changes>

Files changed:
- Summary of change: [file/path.ext](file/path.ext)
- Another change: [other/file.go](other/file.go)
```

### Documentation Organization

- **README.md**: Setup, usage, critical info
- **TODO.md**: Roadmap, planning, architecture
- **CHANGELOG.md**: Versions, releases, technical docs
- **docs/**: Comprehensive guides (no scattered summaries)

### Task Completion

- Complete all terminal commands before finishing
- Leave Git worktree in clean state
- Run formatters and linters
- Document all changes in commit body

---

## 2. AGENTS.md File Policy

- AGENTS.md files may appear anywhere in the container or repo. Their scope is
  the entire directory tree rooted at their location.
- For every file you touch, you must obey instructions in any AGENTS.md file
  whose scope includes that file.
- More deeply nested AGENTS.md files take precedence in case of conflicts.
- Direct user/system instructions always take precedence over AGENTS.md.
- If AGENTS.md includes programmatic checks, you MUST run them and validate that
  they pass after all code changes.

---

## 3. Citations Policy

- If you browse files or use terminal commands, you must add citations to the
  final response (not the PR body) where relevant.
- **Citations reference file paths and terminal outputs:**
  1. `【F:<file_path>†L<line_start>(-L<line_end>)?】` for file paths
  2. `【<chunk_id>†L<line_start>(-L<line_end>)?】` for terminal output

- Ensure line numbers are correct and directly relevant.
- Do not cite empty lines or previous PR diffs/comments.
- Prefer file citations unless terminal output is directly relevant.
- For PR creation, use file citations in the summary and terminal citations in
  the testing section.

---

## 4. Detailed Guidelines Reference

**All detailed guidelines are maintained in separate files to avoid duplication.
Agents MUST consult these files:**

### 4.1 Code Style Guidelines

- **Go**: [.github/code-style-go.md](.github/code-style-go.md)
- **Python**: [.github/code-style-python.md](.github/code-style-python.md)
- **TypeScript**:
  [.github/code-style-typescript.md](.github/code-style-typescript.md)
- **JavaScript**:
  [.github/code-style-javascript.md](.github/code-style-javascript.md)
- **Markdown**: [.github/code-style-markdown.md](.github/code-style-markdown.md)
- **HTML/CSS**: [.github/code-style-html-css.md](.github/code-style-html-css.md)
- **Shell Scripts**: [.github/code-style-shell.md](.github/code-style-shell.md)
- **GitHub Actions**:
  [.github/code-style-github-actions.md](.github/code-style-github-actions.md)
- **Protocol Buffers**:
  [.github/code-style-protobuf.md](.github/code-style-protobuf.md)
- **Additional Languages**: C++
  ([.github/code-style-cpp.md](.github/code-style-cpp.md)), C#
  ([.github/code-style-csharp.md](.github/code-style-csharp.md)), Swift
  ([.github/code-style-swift.md](.github/code-style-swift.md)), R
  ([.github/code-style-r.md](.github/code-style-r.md)), Angular
  ([.github/code-style-angular.md](.github/code-style-angular.md)), JSON
  ([.github/code-style-json.md](.github/code-style-json.md))

### 4.2 Development Process Guidelines

- **Commit Messages**:
  [.github/commit-messages.md](.github/commit-messages.md) - **REQUIRED**
  conventional commit format with file documentation
- **Pull Request Descriptions**:
  [.github/pull-request-descriptions.md](.github/pull-request-descriptions.md)
- **Code Review Guidelines**:
  [.github/review-selection.md](.github/review-selection.md)
- **Test Generation**: [.github/test-generation.md](.github/test-generation.md)
- **Security Guidelines**:
  [.github/security-guidelines.md](.github/security-guidelines.md)

### 4.3 Repository and Workflow Management

- **Complete Copilot/AI Instructions**:
  [.github/copilot-instructions.md](.github/copilot-instructions.md)
- **Repository Setup**:
  [.github/repository-setup.md](.github/repository-setup.md)
- **Workflow Usage**: [.github/workflow-usage.md](.github/workflow-usage.md)

### 4.4 Documentation Organization

- **README.md**: Repository introduction, setup instructions, basic usage,
  critical information
- **TODO.md**: Project roadmap, planning, implementation status, architectural
  decisions
- **CHANGELOG.md**: Version information, release notes, breaking changes,
  feature additions, bug fixes
- **docs/**: Comprehensive technical guides and consolidated documentation

---

## 5. Security & Best Practices Summary

- Avoid hardcoding sensitive information (API keys, passwords, tokens)
- Use proper error handling with meaningful messages
- Validate inputs appropriately (sanitize file paths, validate formats)
- Consider performance implications of code changes
- Look for injection vulnerabilities (command injection, path traversal)
- Use secure defaults for file permissions and temporary files
- Review file handling security (path traversal, file size limits)
- Verify secure handling of sensitive data

**See [.github/security-guidelines.md](.github/security-guidelines.md) for
complete security guidelines.**

---

## 6. Project-Specific Guidelines

- Import from project modules rather than duplicating functionality
- Respect the established architecture patterns
- Before suggesting one-off commands, check for a defined task in tasks.json
- Use meaningful variable names that indicate purpose
- Keep functions small and focused on a single responsibility
- Use the existing CLI framework patterns for new commands
- Follow established patterns for web UI handlers and templates
- Use appropriate logging levels and structured logging

---

## 8. Complete Reference List

**For comprehensive details, always consult these authoritative files:**

### Core Agent Instructions

- [.github/copilot-instructions.md](.github/copilot-instructions.md) - Complete
  AI agent instructions
- [.github/commit-messages.md](.github/commit-messages.md) - **REQUIRED**
  conventional commit format
- [.github/pull-request-descriptions.md](.github/pull-request-descriptions.md)
- [.github/review-selection.md](.github/review-selection.md)
- [.github/test-generation.md](.github/test-generation.md)
- [.github/security-guidelines.md](.github/security-guidelines.md)

### Code Style Guides (All Languages)

- [.github/code-style-go.md](.github/code-style-go.md) (Primary language)
- [.github/code-style-python.md](.github/code-style-python.md)
- [.github/code-style-typescript.md](.github/code-style-typescript.md)
- [.github/code-style-javascript.md](.github/code-style-javascript.md)
- [.github/code-style-markdown.md](.github/code-style-markdown.md)
- [.github/code-style-html-css.md](.github/code-style-html-css.md)
- [.github/code-style-shell.md](.github/code-style-shell.md)
- [.github/code-style-github-actions.md](.github/code-style-github-actions.md)
- [.github/code-style-protobuf.md](.github/code-style-protobuf.md)
- [.github/code-style-cpp.md](.github/code-style-cpp.md)
- [.github/code-style-csharp.md](.github/code-style-csharp.md)
- [.github/code-style-swift.md](.github/code-style-swift.md)
- [.github/code-style-r.md](.github/code-style-r.md)
- [.github/code-style-angular.md](.github/code-style-angular.md)
- [.github/code-style-json.md](.github/code-style-json.md)

### Repository Management

- [.github/repository-setup.md](.github/repository-setup.md)
- [.github/workflow-usage.md](.github/workflow-usage.md)

**Note**: This AGENTS.md file provides essential workflow guidelines. For
detailed implementation guidelines, always reference the specific files listed
above to ensure current and complete information.

---
