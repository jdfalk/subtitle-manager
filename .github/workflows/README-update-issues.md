# file: .github/workflows/README-update-issues.md

## Issue Updates Workflow

The unified issue management workflow processes bulk issue operations from
`issue_updates.json` using a new grouped format for better organization and
GUID-based duplicate prevention.

## New Grouped Format

The `issue_updates.json` file now uses a grouped structure for better
organization:

```json
{
  "create": [
    {
      "title": "Issue Title",
      "body": "Issue description",
      "labels": ["bug", "enhancement"],
      "guid": "create-unique-id-2025-06-20"
    }
  ],
  "update": [
    {
      "number": 123,
      "title": "New Title",
      "labels": ["updated-label"],
      "guid": "update-issue-123-2025-06-20"
    }
  ],
  "comment": [
    {
      "number": 123,
      "body": "Comment text",
      "guid": "comment-123-specific-purpose"
    }
  ],
  "close": [
    {
      "number": 456,
      "state_reason": "completed",
      "guid": "close-issue-456-2025-06-20"
    }
  ],
  "delete": [
    {
      "number": 999,
      "guid": "delete-issue-999-2025-06-20"
    }
  ]
}
```

## GUID-Based Duplicate Prevention

All actions now support optional `guid` fields that prevent duplicate
operations:

- **Comments**: GUIDs are stored as HTML comments in issue comments
- **Creates**: GUIDs are embedded in issue body to prevent duplicate creation
- **Updates/Close/Delete**: GUIDs provide operation tracking

### GUID Format Recommendations

Use descriptive, unique identifiers:

- `"create-feature-request-2025-06-20"`
- `"update-issue-123-labels-2025-06-20"`
- `"comment-456-status-update"`
- `"close-issue-789-completed-2025-06-20"`
- `"delete-spam-issue-999-2025-06-20"`

## Processing Order

Operations are processed in a specific order to maintain logical sequencing:

1. **Create** - Establish new issues first
2. **Update** - Modify existing issues
3. **Comment** - Add comments before any state changes
4. **Close** - Close issues after comments are added
5. **Delete** - Process deletions last

## Supported Actions

### Create Issue

```json
{
  "title": "Issue Title",
  "body": "Issue description",
  "labels": ["bug", "enhancement"],
  "assignees": ["username"],
  "milestone": 1,
  "guid": "create-unique-id-2025-06-20"
}
```

### Update Issue

```json
{
  "number": 123,
  "title": "New Title",
  "body": "Updated description",
  "labels": ["updated-label"],
  "state": "open",
  "guid": "update-issue-123-2025-06-20"
}
```

### Add Comment

```json
{
  "number": 123,
  "body": "Comment text to add to the issue",
  "guid": "comment-123-unique-identifier"
}
```

### Close Issue

```json
{
  "number": 123,
  "state_reason": "completed",
  "guid": "close-issue-123-2025-06-20"
}
```

Valid `state_reason` values:

- `completed` - Issue was completed
- `not_planned` - Issue will not be worked on

### Delete Issue

```json
{
  "number": 123,
  "guid": "delete-issue-123-2025-06-20"
}
```

**Note**: GitHub API doesn't support issue deletion. Issues are closed with
"not_planned" reason and marked for deletion instead.

## Legacy Format Support

The workflow continues to support the old flat array format for backward
compatibility:

```json
[
  {
    "action": "create",
    "title": "Issue Title",
    "body": "Description",
    "labels": ["bug"]
  }
]
```

## Migration from Legacy Format

To migrate from the old format to the new grouped format:

1. Group actions by type (`create`, `update`, `comment`, `close`, `delete`)
2. Remove the `action` field from individual items
3. Add `guid` fields for duplicate prevention
4. Place items in appropriate sections

## Example Usage

Create an `issue_updates.json` file in the repository root:

```json
{
  "create": [
    {
      "title": "New Feature Request",
      "body": "Description of the new feature",
      "labels": ["enhancement"],
      "guid": "create-feature-xyz-2025-06-20"
    }
  ],
  "comment": [
    {
      "number": 456,
      "body": "Adding final update before closing",
      "guid": "closing-comment-456-2025-06-20"
    }
  ],
  "close": [
    {
      "number": 456,
      "state_reason": "completed",
      "guid": "close-completed-456-2025-06-20"
    }
  ]
}
```

The unified workflow will:

1. Create the new issue
2. Add the comment to issue #456 (only if no comment with that GUID exists)
3. Close issue #456 with reason "completed"
4. Remove the processed `issue_updates.json` file
5. Create a PR with the cleanup

## Workflow Integration

The issue updates are processed by the unified issue management workflow
(`unified-issue-management.yml`) which handles:

- **Matrix parallelization** for efficient processing
- **Comprehensive logging** with operation summaries
- **Error handling** with detailed error messages
- **Cleanup automation** with PR generation for file removal
