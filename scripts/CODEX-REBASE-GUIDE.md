# Codex Rebase Quick Reference

## When to Use

Use these scripts when Codex encounters rebase conflicts or gets stuck in rebase loops.

## Simple Auto-Rebase (Recommended for Codex)

```bash
./scripts/codex-rebase.sh
```

This script:

- ✅ Automatically resolves conflicts by keeping current changes
- ✅ Saves incoming changes as `.main.incoming` files
- ✅ Creates backup branch automatically
- ✅ Force pushes the result
- ✅ Generates summary report

## Advanced Rebase with Options

```bash
# Dry run to see what would happen
./scripts/smart-rebase.sh --dry-run main

# Full rebase with force push
./scripts/smart-rebase.sh -f main

# Verbose output for debugging
./scripts/smart-rebase.sh -v -f main
```

## VS Code Tasks

Use these task IDs in VS Code:

- `"Codex Auto Rebase"` - Simple automated rebase
- `"Smart Rebase Dry Run"` - Preview changes
- `"Smart Rebase onto Main"` - Full featured rebase

## Conflict Resolution Strategy

The scripts use intelligent strategies based on file types:

| File Type                              | What Happens                                              |
| -------------------------------------- | --------------------------------------------------------- |
| Documentation (`.md`)                  | Accept incoming from main                                 |
| Build files (`.github/`, `Dockerfile`) | Accept incoming from main                                 |
| Source code (`.go`, `.js`, etc.)       | Keep current + save incoming as `.filename.main.incoming` |
| Config files (`.json`, `.yml`)         | Attempt smart merge                                       |

## Recovery if Something Goes Wrong

1. Find backup branch: `git branch | grep backup`
2. Restore: `git checkout backup-TIMESTAMP-BRANCH`
3. Recreate branch: `git checkout -b original-branch-name`
4. Force push: `git push --force-with-lease origin original-branch-name`

## Files Created

After running, you may see:

- `*.main.incoming` - Incoming versions of conflicted files
- `rebase-summary-TIMESTAMP.md` - Summary of what happened
- `backup-TIMESTAMP-BRANCH` - Backup branch (safe to delete after verification)

## Best Practices for Codex

1. **Always use `codex-rebase.sh` first** - it's designed for automation
2. **Check the summary file** for important information
3. **Review `.main.incoming` files** for important changes you might want to merge
4. **Test the rebased code** before continuing with other tasks
5. **Clean up backup branches** when done: `git branch -D backup-*`

## Environment Variables for Automation

```bash
# Set these in Codex environment if needed
export FORCE_PUSH=true
export DRY_RUN=false
export VERBOSE=true
```

## Common Error Solutions

- **"Not in git repository"**: Make sure you're in project root
- **"Target branch does not exist"**: Run `git fetch origin` first
- **"Force push failed"**: Someone else pushed changes, try `git pull --rebase` first
- **Script hangs**: Kill with Ctrl+C, then `git rebase --abort`
