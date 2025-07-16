#!/bin/bash
# file: scripts/codex-rebase.sh
# version: 1.1.0
# guid: 38643ccd-055a-45ea-b990-4c12defc6421

# Codex-specific rebase automation script
# This script is designed to be used by AI agents with minimal interaction

set -euo pipefail

# Configuration for Codex usage
FORCE_PUSH=true
AUTO_COMMIT=true
BACKUP_ENABLED=true
CONFLICT_STRATEGY="auto-resolve"
STASHED=false

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log() {
    echo -e "${BLUE}[CODEX-REBASE]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

check_dependencies() {
    if ! command -v git >/dev/null 2>&1; then
        error "git command not found"
        exit 1
    fi
}

check_clean_state() {
    if [[ -d .git/rebase-merge || -d .git/rebase-apply || -d .git/MERGE_HEAD ]]; then
        error "Another Git operation is in progress"
        exit 1
    fi
}

stash_changes() {
    if ! git diff-index --quiet HEAD --; then
        log "Stashing uncommitted changes"
        git stash push -m "codex-pre-rebase-$(date +%Y%m%d-%H%M%S)" >/dev/null
        STASHED=true
    else
        STASHED=false
    fi
}

restore_stash() {
    if [[ "$STASHED" == true ]]; then
        if git stash list | grep -q "codex-pre-rebase"; then
            log "Restoring stashed changes"
            git stash pop >/dev/null || true
        fi
    fi
}
trap restore_stash EXIT

# Get target branch (default to main)
TARGET_BRANCH="${1:-main}"
CURRENT_BRANCH=$(git branch --show-current)

# Set up remote if it doesn't exist (for CI environments)
setup_remote() {
    if ! git remote get-url origin >/dev/null 2>&1; then
        log "No 'origin' remote found, configuring for jdfalk/subtitle-manager"
        git remote add origin https://github.com/jdfalk/subtitle-manager.git
        log "Added origin remote: https://github.com/jdfalk/subtitle-manager.git"
    else
        log "Origin remote already configured: $(git remote get-url origin)"
    fi
}

log "Starting Codex rebase automation"
log "Current branch: $CURRENT_BRANCH"
log "Target branch: $TARGET_BRANCH"

# Setup remote if needed
setup_remote

check_dependencies
check_clean_state

# Pre-flight checks
if [[ "$CURRENT_BRANCH" == "$TARGET_BRANCH" ]]; then
    error "Cannot rebase branch onto itself"
    exit 1
fi

# Stash any uncommitted changes
stash_changes

# Create backup branch
BACKUP_BRANCH="codex-backup-$(date +%Y%m%d-%H%M%S)-$CURRENT_BRANCH"
git branch "$BACKUP_BRANCH"
log "Created backup branch: $BACKUP_BRANCH"

# Fetch latest changes if 'origin' remote exists
if git remote get-url origin >/dev/null 2>&1; then
    log "Fetching latest changes"
    if ! git fetch origin; then
        error "Failed to fetch from origin. Check network connectivity and authentication."
        error "Remote URL: $(git remote get-url origin 2>/dev/null || echo 'Not configured')"
        exit 1
    fi
else
    log "Remote 'origin' not found, skipping fetch"
fi

# Resolve a single conflict using Codex or ChatGPT CLI if available
resolve_conflict_ai() {
    local file="$1"

    if command -v codex >/dev/null 2>&1; then
        log "Using Codex CLI to resolve $file"
        git diff --ours --theirs "$file" | codex patch --stdin --filename "$file" && git add "$file" && return 0
    elif command -v chatgpt >/dev/null 2>&1; then
        log "Using ChatGPT CLI to resolve $file"
        git diff --ours --theirs "$file" | chatgpt patch --stdin --filename "$file" && git add "$file" && return 0
    fi

    return 1
}

# Function to auto-resolve conflicts with Codex-friendly strategies
auto_resolve_conflicts() {
    local conflicted_files
    conflicted_files=$(git diff --name-only --diff-filter=U)

    if [[ -z "$conflicted_files" ]]; then
        return 0
    fi

    log "Auto-resolving conflicts in: $conflicted_files"

    echo "$conflicted_files" | while read -r file; do
        if [[ -n "$file" ]]; then
            if resolve_conflict_ai "$file"; then
                log "Resolved $file using AI CLI"
                continue
            fi

            # Save incoming version with .main.incoming suffix
            local base_name="${file%.*}"
            local extension="${file##*.}"
            local incoming_file="${base_name}.${extension}.main.incoming"

            # Extract and save incoming version
            git show :3:"$file" > "$incoming_file" 2>/dev/null || {
                log "Warning: Could not extract incoming version of $file"
            }

            # Keep current version (our changes)
            git checkout --ours "$file"
            git add "$file"

            log "Resolved $file: kept current, saved incoming as $incoming_file"
        fi
    done

    # Continue the rebase
    git rebase --continue
}

# Perform the rebase with auto-conflict resolution
log "Starting rebase onto $TARGET_BRANCH"

# Set up git config for automated commits
git config user.email "codex@subtitle-manager.local" 2>/dev/null || true
git config user.name "Codex Auto-Rebase" 2>/dev/null || true

# Start rebase and handle conflicts automatically
while true; do
    if git rebase "$TARGET_BRANCH"; then
        success "Rebase completed successfully"
        break
    else
        # Check if we're in a rebase state with conflicts
        if git status --porcelain | grep -q "^UU\|^AA\|^DD"; then
            log "Conflicts detected, auto-resolving..."
            auto_resolve_conflicts
        else
            # Some other error occurred
            error "Rebase failed for unknown reason"
            git rebase --abort
            exit 1
        fi
    fi
done

# Commit any remaining changes (e.g., resolved files)
if ! git diff-index --quiet HEAD --; then
    log "Committing remaining changes"
    git add -A
    git commit -m "chore: finalize auto-rebase" || true
fi

# Force push the rebased branch
log "Force pushing rebased branch"
if git push --force-with-lease origin "$CURRENT_BRANCH"; then
    success "Force push completed"
else
    error "Force push failed. This might be due to:"
    error "1. Authentication issues (no push access)"
    error "2. Network connectivity problems"
    error "3. Branch protection rules"
    error "Remote URL: $(git remote get-url origin 2>/dev/null || echo 'Not configured')"
    exit 1
fi

# Create a summary of what happened
SUMMARY_FILE="rebase-summary-$(date +%Y%m%d-%H%M%S).md"
cat > "$SUMMARY_FILE" << EOF
# Codex Rebase Summary

**Date:** $(date)
**Current Branch:** $CURRENT_BRANCH
**Target Branch:** $TARGET_BRANCH
**Backup Branch:** $BACKUP_BRANCH

## Changes Made
- Rebased $CURRENT_BRANCH onto $TARGET_BRANCH
- Auto-resolved conflicts using AI CLI when available
- Force pushed rebased branch to origin

## Conflict Resolution
$(if ls *.main.incoming 2>/dev/null; then
    echo "The following files had conflicts and incoming versions were saved:"
    ls *.main.incoming 2>/dev/null | while read f; do
        echo "- $f"
    done
else
    echo "No conflicts detected during rebase."
fi)

## Recovery Instructions
If you need to undo this rebase:
\`\`\`bash
git checkout $BACKUP_BRANCH
git branch -D $CURRENT_BRANCH
git checkout -b $CURRENT_BRANCH
git push --force-with-lease origin $CURRENT_BRANCH
\`\`\`

## Cleanup
To clean up the backup branch after confirming everything is working:
\`\`\`bash
git branch -D $BACKUP_BRANCH
\`\`\`
EOF

log "Rebase summary saved to: $SUMMARY_FILE"
success "Codex rebase automation completed successfully!"

# Optionally clean up .main.incoming files if they're identical to resolved files
log "Checking for redundant .main.incoming files..."
shopt -s nullglob
for incoming_file in *.main.incoming; do
    if [[ -f "$incoming_file" ]]; then
        original_file="${incoming_file%.main.incoming}"
        if [[ -f "$original_file" ]] && cmp -s "$incoming_file" "$original_file"; then
            log "Removing redundant $incoming_file (identical to $original_file)"
            rm "$incoming_file"
        fi
    fi
done
shopt -u nullglob

log "Rebase automation complete. Check $SUMMARY_FILE for details."
restore_stash
