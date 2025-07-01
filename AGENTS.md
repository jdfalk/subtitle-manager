<!-- file: AGENTS.md -->
<!-- version: 2.0.0 -->
<!-- guid: 2e7c1a4b-5d3f-4b8c-9e1f-7a6b2c3d4e5f -->

# AI Agent Instructions

This file provides comprehensive instructions for AI agents working on this repository.

---

## Core Instructions

All AI agents MUST follow the comprehensive guidelines in these files (in order of precedence):

1. **[.github/copilot-instructions.md](.github/copilot-instructions.md)** - Complete coding guidelines, styles, and workflows
2. **[.github/commit-messages.md](.github/commit-messages.md)** - Conventional commits with required file change documentation
3. **[.github/pull-request-descriptions.md](.github/pull-request-descriptions.md)** - PR templates and guidelines
4. **[.github/review-selection.md](.github/review-selection.md)** - Code review focus areas and process
5. **[.github/test-generation.md](.github/test-generation.md)** - Testing standards and patterns

---

## Repository-Specific Requirements

### 1. File Headers

Every file MUST begin with a standard header:

```markdown
<!-- file: relative/path/to/file -->
<!-- version: 1.0.0 -->
<!-- guid: unique-identifier -->
```

### 2. Commit Message Format (REQUIRED)

```text
<type>[scope]: <description>

<body with motivation and changes>

Files changed:
- Summary of change: [file/path.ext](file/path.ext)
- Another change: [other/file.go](other/file.go)
```

### 3. Documentation Organization

- **README.md**: Setup, usage, critical info
- **TODO.md**: Roadmap, planning, architecture
- **CHANGELOG.md**: Versions, releases, technical docs
- **docs/**: Comprehensive guides (no scattered summaries)

### 4. Task Completion

- Complete all terminal commands before finishing
- Leave Git worktree in clean state
- Run formatters and linters
- Document all changes in commit body

---

## GitHub Issue Management

Use `scripts/create-issue-update.sh` for all GitHub operations:

```bash
# Create new issue
./scripts/create-issue-update.sh create "Title" "Body" "label1,label2"

# Update existing issue
./scripts/create-issue-update.sh update 123 "New body" "updated-label"

# Add comment
./scripts/create-issue-update.sh comment 123 "Comment text"

# Close issue
./scripts/create-issue-update.sh close 123 "completed"
```

Do NOT manually edit `.github/issue-updates/*.json` files.

---

## Agent Workflow

1. Read task requirements
2. Follow file-specific guidelines from copilot-instructions.md
3. Use conventional commits with file change documentation
4. Run formatters/linters before completion
5. Ensure clean Git state
6. Add proper citations in final response

---

## Priority Override

Direct user instructions always take precedence over this file. When conflicts arise, ask for clarification.

---

## Citations Format

For file references: `【F:<file_path>†L<line_start>(-L<line_end>)?】`

For terminal output: `【<chunk_id>†L<line_start>(-L<line_end>)?】`
