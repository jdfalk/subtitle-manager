# Rebase Script Fixes for CI Environments

## Problem

The rebase scripts (`codex-rebase.sh` and `smart-rebase.sh`) were failing in CI environments because they assumed the `origin` remote would always be configured. In CI environments like GitHub Actions, the repository might be cloned without a remote configured.

## Error

```
fatal: 'origin' does not appear to be a git repository
fatal: Could not read from remote repository.
```

## Solution

Added automatic remote configuration to both scripts with the following features:

### 1. Remote Setup Function

- Check if `origin` remote exists using `git remote get-url origin`
- If not found, automatically configure it for `https://github.com/jdfalk/subtitle-manager.git`
- Log the configuration process for transparency

### 2. Enhanced Error Handling

- Better error messages for fetch and push operations
- Display the remote URL when operations fail
- Explain common failure reasons (authentication, network, branch protection)

## Files Modified

### scripts/codex-rebase.sh

- Added `setup_remote()` function
- Enhanced error handling for `git fetch origin` and `git push --force-with-lease origin`
- Call `setup_remote()` at script start

### scripts/smart-rebase.sh

- Added `setup_remote()` function
- Enhanced error handling for `git push --force-with-lease origin`
- Call `setup_remote()` after prerequisites check

## Benefits

1. **CI Compatibility**: Scripts now work in CI environments without manual remote configuration
2. **Better Debugging**: Clear error messages help identify authentication and network issues
3. **Automatic Recovery**: Scripts auto-configure missing remotes instead of failing
4. **Transparency**: All remote operations are logged for troubleshooting

## Testing

Both scripts should now work in:

- Local development environments (existing remotes preserved)
- CI/CD pipelines (remotes auto-configured)
- Fresh repository clones (remotes auto-configured)

The scripts will continue to work normally if `origin` is already configured, and will automatically set it up if missing.
