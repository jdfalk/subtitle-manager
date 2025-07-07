# Smart Rebase Tool Documentation

## Overview

The Smart Rebase Tool is a comprehensive Python-based Git rebase automation tool
that provides intelligent conflict resolution, persistent state management, and
robust recovery capabilities. This tool addresses common frustrations with Git
rebases by:

- **Persistent State Management**: Saves progress and can resume from
  interruptions
- **Intelligent Conflict Resolution**: Automatically resolves conflicts based on
  file patterns
- **Backup Management**: Creates backup branches and provides recovery
  instructions
- **Comprehensive Logging**: Detailed logs and progress tracking
- **Multiple Operation Modes**: Interactive, automated, and smart modes

## Features

### Core Capabilities

1. **State Persistence**: The tool saves its state to `.rebase-state/` directory
   and can resume from any interruption
2. **Conflict Resolution**: Smart strategies for different file types
   (documentation, code, config files)
3. **Backup Management**: Automatic backup branch creation with timestamps
4. **Recovery Instructions**: Generates detailed recovery guides for manual
   intervention
5. **Progress Tracking**: Real-time progress updates and conflict resolution
   tracking

### File Pattern-Based Conflict Resolution

The tool uses intelligent pattern matching to determine conflict resolution
strategies:

- **Documentation Files** (`.md`, `.rst`, `.txt`, `docs/`): Prefer incoming
  changes
- **Build/CI Files** (`.github/`, `Dockerfile`, `.yml`, `Makefile`): Prefer
  incoming changes
- **Package Files** (`go.mod`, `package.json`, `requirements.txt`): Smart merge
- **Source Code** (`.go`, `.py`, `.js`, `.ts`): Auto-resolve with pattern
  matching
- **Test Files** (`*_test.go`, `test/`): Prefer current (keep local tests)
- **Configuration Files** (`.conf`, `.ini`, `.toml`): Manual review required

## Usage

### Basic Commands

```bash
# Run rebase onto main branch
python3 scripts/smart-rebase.py main

# Run with verbose output
python3 scripts/smart-rebase.py --verbose main

# Dry run to see what would happen
python3 scripts/smart-rebase.py --dry-run main

# Rebase with force push after success
python3 scripts/smart-rebase.py --force main
```

### State Management Commands

```bash
# Resume interrupted rebase
python3 scripts/smart-rebase.py --resume

# Check current rebase status
python3 scripts/smart-rebase.py --status

# Abort current rebase and restore backup
python3 scripts/smart-rebase.py --abort

# Clean up state files
python3 scripts/smart-rebase.py --cleanup
```

### VS Code Tasks

The tool integrates with VS Code through predefined tasks:

- **Smart Rebase: Run** - Standard rebase operation
- **Smart Rebase: Dry Run** - Preview mode
- **Smart Rebase: Resume** - Resume interrupted operation
- **Smart Rebase: Status** - Check current status
- **Smart Rebase: Abort** - Abort and restore backup
- **Smart Rebase: Force Push** - Rebase with force push

## State Management

### State Directory Structure

```text
.rebase-state/
├── rebase.state          # Main state file (JSON)
├── progress.json         # Progress tracking
├── rebase.log           # Detailed logs
├── conflicts.log        # Conflict resolution log
├── summary.md           # Operation summary
└── recovery_instructions.md  # Manual recovery guide
```

### Backup Directory

```text
.rebase-backup/
├── <filename>_backup_<timestamp>  # Individual file backups
└── ...
```

## Conflict Resolution Strategies

### Auto-Resolution Process

1. **File Pattern Analysis**: Determine strategy based on file path/extension
2. **Conflict Parsing**: Extract current and incoming sections
3. **Smart Merging**: Apply appropriate merge strategy
4. **Validation**: Ensure resolved content is valid
5. **Staging**: Automatically stage resolved files

### Smart Merge Strategies

- **Prefer Incoming**: Take changes from target branch
- **Prefer Current**: Keep local changes
- **Smart Merge**: Intelligently combine both versions
- **Auto Resolve**: Use pattern-based resolution for code files
- **Manual Review**: Flag for manual intervention

## Recovery and Error Handling

### Automatic Recovery

The tool handles common scenarios:

- **Interrupted Operations**: Resume from last known state
- **Conflict Resolution Failures**: Provide detailed error messages
- **Git Command Failures**: Capture and log error details
- **State Corruption**: Validate and recover from backup

### Manual Recovery

When automatic recovery isn't possible:

1. Check `recovery_instructions.md` for detailed steps
2. Review backup branches for manual restoration
3. Use `--abort` to safely exit and restore original state
4. Consult logs for troubleshooting information

## Advanced Features

### Logging Levels

- **INFO**: Standard operation messages
- **SUCCESS**: Successful operation confirmations
- **WARNING**: Non-critical issues
- **ERROR**: Operation failures
- **VERBOSE**: Detailed debugging information

### Progress Tracking

- **Commit Progress**: Tracks processed vs. total commits
- **Conflict Resolution**: Monitors resolved vs. total conflicts
- **Phase Tracking**: Shows current operation phase
- **Time Tracking**: Records operation duration

### Backup Management

- **Automatic Backup**: Creates timestamped backup branches
- **File Backups**: Individual file backups before resolution
- **Recovery Points**: Multiple restoration options
- **Cleanup**: Automatic cleanup of old backups

## Best Practices

### Before Running

1. **Commit Changes**: Ensure working directory is clean
2. **Fetch Updates**: Pull latest changes from remote
3. **Review Conflicts**: Understand potential conflict areas
4. **Test Dry Run**: Use `--dry-run` to preview changes

### During Operation

1. **Monitor Progress**: Use `--status` to check progress
2. **Review Logs**: Check detailed logs for issues
3. **Validate Resolutions**: Ensure auto-resolved conflicts are correct
4. **Manual Intervention**: Be prepared for manual conflict resolution

### After Operation

1. **Review Changes**: Validate all merged content
2. **Test Build**: Ensure code still compiles/runs
3. **Check Tests**: Run test suite to validate functionality
4. **Clean Up**: Remove state files after successful completion

## Troubleshooting

### Common Issues

1. **State File Corruption**: Use `--cleanup` and restart
2. **Conflict Resolution Failures**: Check file permissions and encoding
3. **Git Command Failures**: Verify Git configuration and repository state
4. **Interrupted Operations**: Use `--resume` to continue

### Error Messages

The tool provides detailed error messages with:

- **Root Cause**: What went wrong
- **Context**: When and where it happened
- **Recovery Steps**: How to fix the issue
- **Prevention**: How to avoid in the future

### Support Files

- **Logs**: Check `.rebase-state/rebase.log` for detailed information
- **State**: Review `.rebase-state/rebase.state` for current state
- **Recovery**: Follow `.rebase-state/recovery_instructions.md`
- **Summary**: Check `.rebase-state/summary.md` for operation overview

## Integration

### VS Code Integration

The tool integrates seamlessly with VS Code through:

- **Tasks**: Predefined tasks for common operations
- **Terminal**: Run directly from integrated terminal
- **Problem Matcher**: Git error detection and navigation
- **Output Panel**: Structured output display

### CI/CD Integration

Can be integrated into CI/CD pipelines:

- **Automated Rebases**: Scheduled rebase operations
- **Conflict Detection**: Early conflict identification
- **Status Reporting**: Integration with build systems
- **Recovery Automation**: Automated recovery procedures

## File Structure

```text
scripts/
├── smart-rebase.py      # Main rebase tool
├── rebase               # Symlink for easy access
├── enhanced-rebase.sh   # Legacy bash version
├── smart-rebase.sh      # Legacy bash version
└── codex-rebase.sh      # Legacy bash version
```

## Migration from Legacy Scripts

If you're migrating from the legacy bash scripts:

1. **Backup Current State**: Save any in-progress rebase state
2. **Clean Up**: Remove old state files
3. **Test New Tool**: Run dry-run operations first
4. **Update Workflows**: Switch to new commands and tasks
5. **Remove Legacy**: Clean up old scripts when confident

## Contributing

When contributing to the Smart Rebase Tool:

1. **Add Tests**: Include test cases for new features
2. **Update Documentation**: Keep this documentation current
3. **Follow Patterns**: Match existing code style and patterns
4. **Add Logging**: Include appropriate logging for debugging
5. **Test Recovery**: Ensure recovery scenarios work correctly

## Support

For issues or questions:

1. **Check Logs**: Review detailed logs in `.rebase-state/`
2. **Review State**: Check current state and progress
3. **Try Recovery**: Use built-in recovery mechanisms
4. **Manual Intervention**: Follow recovery instructions
5. **Report Issues**: Include logs and state files in reports
