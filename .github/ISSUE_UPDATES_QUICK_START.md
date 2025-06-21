# 🚀 New Issue Updates System

## What Changed?

We've upgraded our issue management system to eliminate merge conflicts! Instead of everyone editing the same `issue_updates.json` file, we now use individual files.

## Quick Start

### Option 1: Use the Helper Script (Recommended)

```bash
# Create a new issue
./scripts/create-issue-update.sh create "Your Issue Title" "Issue description" "label1,label2"

# Update an existing issue
./scripts/create-issue-update.sh update 123 "Updated description" "label1,label2"

# Add a comment
./scripts/create-issue-update.sh comment 123 "Your comment text"

# Close an issue
./scripts/create-issue-update.sh close 123 "completed"
```

### Option 2: Manual File Creation

1. Generate a UUID: `uuidgen`
2. Create file: `.github/issue-updates/{uuid}.json`
3. Use this format:

```json
{
  "action": "create",
  "title": "Your Issue Title",
  "body": "Your issue description",
  "labels": ["enhancement", "feature"],
  "guid": "create-your-issue-title-2025-06-21"
}
```

## File Locations

- **New format**: `.github/issue-updates/{uuid}.json` ← Use this!
- **Legacy format**: `issue_updates.json` ← Still works but avoid for new updates

## Examples in the Repo

Check out these example files:
- `bc03b7dc-eba7-4b95-9a90-a0224b274633.json` - Create issue
- `ff8d269a-ce03-4985-b653-0e0af2d363e2.json` - Update issue
- `a8f2c4e6-b1d3-4a7e-9c5b-2f8e1a4d6c9b.json` - Add comment

## Benefits

✅ **No more merge conflicts** - Everyone works on separate files
✅ **Parallel development** - Multiple people can create updates simultaneously
✅ **Cleaner git history** - Each update is tracked individually
✅ **Easier reviews** - Changes are isolated and clear

## Actions Supported

| Action    | Purpose               | Required Fields           |
| --------- | --------------------- | ------------------------- |
| `create`  | Create new issue      | `title`, `body`, `labels` |
| `update`  | Update existing issue | `number`, `body`/`labels` |
| `comment` | Add comment to issue  | `number`, `body`          |
| `close`   | Close an issue        | `number`, `state_reason`  |
| `delete`  | Delete an issue       | `number`                  |

## Need Help?

- 📖 Full migration guide: `.github/ISSUE_UPDATES_MIGRATION.md`
- 📋 Format documentation: `.github/issue-updates/README.md`
- 🛠️ Helper script: `scripts/create-issue-update.sh`
- 💡 Examples: Check `.github/issue-updates/` directory

## Testing

Use dry-run mode to test your changes:
1. Create your update file
2. Go to Actions → Unified Issue Management
3. Click "Run workflow"
4. Check "Run in dry-run mode"
5. Review the output before making real changes

---

**Questions?** Check the documentation or ask the team! 🚀
