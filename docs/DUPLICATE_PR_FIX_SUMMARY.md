# Fix for Duplicate Archive PRs - Solution Summary

## Problem Identified

The unified issue management workflow was creating too many pull requests for
archiving processed files (e.g., PR #1093, #1094, #1096, #1097). Every time the
workflow ran on main branch merges, it created a new PR instead of updating the
existing one.

## Root Cause

1. **Unique branch names**: The workflow used `${{ github.run_id }}` in branch
   names, making them unique for each run
2. **No existing PR check**: No logic to detect if a similar PR already existed
3. **Auto-delete branches**: Set to `delete-branch: true`, removing branches
   after merge

## Solution Implemented

### 1. **Consistent Branch Names**

Changed from dynamic branch names to static ones:

- `archive-processed-files-${{ github.run_id }}` → `archive-processed-files`
- `update-issue-permalinks-${{ github.run_id }}` → `update-issue-permalinks`

### 2. **Existing PR Detection**

Added a new step `check-prs` that:

- Uses GitHub CLI to search for existing open PRs on the same branches
- Stores found PR numbers in outputs for later use
- Provides clear logging about existing PRs

```yaml
- name: Check for existing archive PRs
  run: |
    existing_processed=$(gh pr list --base "${{ github.event.repository.default_branch }}" --head "archive-processed-files" --state open --json number,headRefName --jq '.[0].number // empty')
    existing_legacy=$(gh pr list --base "${{ github.event.repository.default_branch }}" --head "update-issue-permalinks" --state open --json number,headRefName --jq '.[0].number // empty')
```

### 3. **Rebase and Update Logic**

The `peter-evans/create-pull-request` action automatically handles:

- **If PR exists**: Updates the existing PR with new commits and refreshed body
- **If no PR exists**: Creates a new PR

### 4. **Enhanced PR Bodies**

Updated PR descriptions to:

- Show "Latest Workflow run" with links
- Include timestamps for when last updated
- Indicate whether it's updating an existing PR or creating new
- Provide better context about accumulated changes

### 5. **Improved Branch Management**

- Changed `delete-branch: false` to keep branches persistent
- Branches now serve as staging areas for accumulated changes
- PRs get updated with new commits rather than creating new branches

### 6. **Better Reporting**

Enhanced workflow summaries to show:

- Whether PRs were updated or newly created
- Links to existing PRs even when no changes were made
- Clear distinction between different types of archive operations

## Benefits

### ✅ **Reduced PR Clutter**

- One PR per archive type instead of multiple PRs per workflow run
- Cleaner PR list and easier repository management

### ✅ **Better Change Tracking**

- Accumulated changes in single PRs show the full scope of recent updates
- Easier to review all processed files together

### ✅ **Improved Workflow Efficiency**

- Reuses existing branches and PRs
- Faster execution since no new branch creation needed
- Less GitHub API usage

### ✅ **Enhanced Visibility**

- PR bodies show latest run information and timestamps
- Clear indication of whether PR was updated vs newly created
- Links to specific workflow runs that made changes

## How It Works Now

### First Run (No existing PRs)

1. Check for existing PRs → None found
2. Create new PRs with static branch names
3. Archive processed files
4. Report: "Created new PR"

### Subsequent Runs (PRs exist)

1. Check for existing PRs → Found existing
2. Update existing PRs with new commits
3. Refresh PR body with latest run info
4. Report: "Updated existing PR"

### When No Changes Needed

1. Check for existing PRs → Found existing
2. No new commits needed
3. Report: "Existing PR available (no new changes)"

## Example Output

### Before (Problem)

```
📦 Archive Files PR: #1093 (run 123)
📦 Archive Files PR: #1094 (run 124)
📦 Archive Files PR: #1096 (run 125)
📦 Archive Files PR: #1097 (run 126)
```

### After (Solution)

```
📦 Archive Files PR: #1093 (Updated existing - latest run 126)
```

## Additional Fix Applied

### **Timestamp Issue Resolution**

Fixed an issue where the `$(date -u '+%Y-%m-%d %H:%M:%S UTC')` command in PR
body templates was not being properly executed:

- **Problem**: The date command was being treated as literal text instead of
  being executed
- **Solution**: Added a dedicated `timestamp` step that generates the current
  UTC timestamp
- **Implementation**: Uses `${{ steps.timestamp.outputs.current }}` in PR body
  templates
- **Result**: PR bodies now show accurate "Updated" timestamps

### **Example PR Body Output**

```
## Summary
Automatically archiving processed distributed issue update files to the processed/ directory.

Latest Operation: update-issues
Latest Workflow run: 15799193636
Event: schedule
Triggered by: jdfalk
Updated: 2025-06-21 14:30:15 UTC
```

## Configuration Changes Made

1. **Branch names**: Static instead of run-specific
2. **PR detection**: Added check step before creation
3. **Branch persistence**: `delete-branch: false`
4. **Enhanced messaging**: Better PR bodies and workflow summaries
5. **Conditional reporting**: Shows update vs create status

This solution maintains all the functionality while eliminating the duplicate PR
problem, providing a much cleaner and more manageable workflow experience.
