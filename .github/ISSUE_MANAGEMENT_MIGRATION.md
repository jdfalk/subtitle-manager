# file: .github/ISSUE_MANAGEMENT_MIGRATION.md

# Issue Management Migration Guide

This document outlines the migration from separate issue management scripts to a
unified Python-based system.

## Overview

The repository is transitioning from multiple separate workflows and scripts to
a unified issue management system:

### Before (Legacy)

- `update-issues.yml` - Bash-based issue updates from JSON
- `copilot-tickets.yml` - Python script for Copilot review tickets
- `close-duplicates.py` - Standalone Python script for duplicate closure
- No CodeQL alert ticket generation

### After (Unified)

- `issue_manager.py` - Single Python script handling all operations
- `unified-issue-management.yml` - Comprehensive workflow
- `codeql-alert-tickets.yml` - Dedicated CodeQL alert handling
- Enhanced error handling, logging, and API consistency

## Migration Status

### ✅ Completed

1. **Unified Python Script** - `issue_manager.py` created with all functionality
2. **CodeQL Alert Tickets** - New feature to auto-generate security tickets
3. **Enhanced API Client** - Consistent GitHub API usage with proper error
   handling
4. **Workflow Updates** - Existing workflows updated to use new script

### 🔄 In Progress (Parallel Operation)

- Both old and new systems running simultaneously
- Legacy workflows updated to use new script while maintaining compatibility
- Migration notices added to existing workflows

### 📋 TODO

1. **Validation Period** - Monitor both systems for 1-2 weeks
2. **Feature Parity Verification** - Ensure all original functionality works
3. **Performance Comparison** - Compare execution times and reliability
4. **Legacy Cleanup** - Remove old workflows after successful validation

## New Features

### 1. CodeQL Alert Ticket Generation

- Automatically creates GitHub issues for CodeQL security alerts
- Includes detailed alert information (rule, severity, location)
- Prevents duplicate tickets for the same alert
- Labels tickets with "security" for easy filtering

### 2. Enhanced Duplicate Detection

- Better title matching and grouping
- Dry-run mode for testing
- Improved logging and reporting

### 3. Unified CLI Interface

```bash
# Process issue updates
python issue_manager.py update-issues

# Handle Copilot review tickets
python issue_manager.py copilot-tickets

# Close duplicate issues
python issue_manager.py close-duplicates --dry-run

# Generate CodeQL alert tickets
python issue_manager.py codeql-alerts

# Handle webhook events
python issue_manager.py event-handler
```

### 4. Better Error Handling

- Comprehensive try-catch blocks
- Detailed error logging
- Graceful failure handling
- API rate limit awareness

## Configuration

### Environment Variables

- `GH_TOKEN` - GitHub personal access token
- `REPO` - Repository in owner/name format
- `GITHUB_EVENT_NAME` - Event type for webhook processing
- `GITHUB_EVENT_PATH` - Path to event payload

### Labels Used

- `copilot-review` - Copilot review comment tickets
- `security` - CodeQL alert tickets
- `duplicate-check` - Issues being checked for duplicates

### Permissions Required

- `issues: write` - Create, update, close issues
- `pull-requests: read` - Read PR information for Copilot tickets
- `contents: read` - Read repository content
- `security-events: read` - Read CodeQL alerts

## Workflow Triggers

### Unified Issue Management (`unified-issue-management.yml`)

- **Push to main** - Issue updates, CodeQL alerts, push analysis
- **PR events** - Copilot review ticket management
- **Schedule** - Duplicate cleanup (daily), CodeQL checks (twice daily)
- **Manual** - All operations with configurable options

### Legacy Workflows (During Migration)

- Continue running but use new Python script
- Will be removed after validation period

## Testing

### Manual Testing

```bash
# Test duplicate detection (dry run)
python .github/scripts/issue_manager.py close-duplicates --dry-run

# Test CodeQL alert processing
python .github/scripts/issue_manager.py codeql-alerts

# Test issue updates (requires issue_updates.json)
python .github/scripts/issue_manager.py update-issues
```

### Workflow Testing

- Use `workflow_dispatch` triggers to test individual components
- Monitor scheduled runs for automated operations
- Check issue creation and updates in repository

## Rollback Plan

If issues are discovered during migration:

1. **Immediate** - Disable new workflows
2. **Restore** - Re-enable original workflow logic
3. **Investigate** - Debug issues in new system
4. **Re-deploy** - Deploy fixes and re-enable

## Success Criteria

- [ ] All issue update operations work correctly
- [ ] Copilot ticket management maintains functionality
- [ ] Duplicate closure works as expected
- [ ] CodeQL alerts generate appropriate tickets
- [ ] No data loss or corruption
- [ ] Performance equal or better than legacy system
- [ ] Error rates reduced
- [ ] Code maintainability improved

## Post-Migration

After successful validation:

1. Remove legacy workflow files
2. Archive old Python scripts
3. Update documentation references
4. Add new features only available in unified system
5. Consider additional integrations (security tools, etc.)

## Support

For issues during migration:

- Check workflow run logs for detailed error information
- Review issue management activity in repository
- Compare behavior with legacy system outputs
- Report any discrepancies or unexpected behavior
