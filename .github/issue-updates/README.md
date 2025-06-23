# Issue Updates Directory

This directory contains individual issue update files that are processed by the unified issue management workflow.

## Format

Each file should be named with a UUID (e.g., `bc03b7dc-eba7-4b95-9a90-a0224b274633.json`) and contain a single issue operation.

### File Structure

```json
{
  "action": "create|update|comment|close|delete",
  "title": "Issue title (required for create)",
  "body": "Issue body or comment text",
  "number": 123,
  "labels": ["label1", "label2"],
  "assignees": ["username1", "username2"],
  "milestone": 5,
  "state": "open|closed",
  "state_reason": "completed|not_planned|reopened",
  "guid": "unique-identifier-for-deduplication"
}
```

### Examples

#### Create Issue

```json
{
  "action": "create",
  "title": "Add dark mode support",
  "body": "Implement dark mode toggle for better user experience",
  "labels": ["enhancement", "ui"],
  "guid": "create-dark-mode-2025-06-21"
}
```

#### Update Issue

```json
{
  "action": "update",
  "number": 42,
  "title": "Updated: Performance optimization complete",
  "labels": ["performance", "completed"],
  "guid": "update-issue-42-2025-06-21"
}
```

#### Comment on Issue

```json
{
  "action": "comment",
  "number": 42,
  "body": "Testing completed successfully!",
  "guid": "comment-42-testing-2025-06-21"
}
```

#### Close Issue

```json
{
  "action": "close",
  "number": 42,
  "state_reason": "completed",
  "guid": "close-issue-42-2025-06-21"
}
```

## Benefits

- **No merge conflicts**: Each update is in its own file
- **Parallel development**: Multiple developers can create updates simultaneously
- **Atomic operations**: Each file represents a single issue operation
- **Easy tracking**: Files can be linked to specific features or pull requests
- **UUID naming**: Prevents filename conflicts

## Workflow Integration

The workflow automatically:

1. Scans all `.json` files in this directory
2. Processes them in order
3. Moves processed files to a `processed/` subdirectory (optional)
4. Continues to support the legacy `issue_updates.json` format

## Creating New Updates

1. Generate a UUID: `uuidgen` (on macOS/Linux) or use online UUID generator
2. Create a new file: `{uuid}.json`
3. Add your issue operation following the format above
4. Commit and push - the workflow will process it automatically

## Legacy Support

The workflow still processes the root `issue_updates.json` file for backward compatibility, but new updates should use this directory structure.
