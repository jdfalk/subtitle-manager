#!/bin/bash
# file: scripts/create-github-projects.sh
# version: 1.0.0
# guid: fc4bb6a5-97e2-46fb-a869-55be8644ba34

set -euo pipefail

usage() {
    cat <<'USAGE'
Usage: create-github-projects.sh [options]

Verifies GitHub CLI authentication scopes before calling the unified
project manager script. Set GH_PROJECT_SCRIPT to override the default
script location.
USAGE
}

REQUIRED_SCOPES=("repo" "project")

check_scopes() {
    if ! command -v gh >/dev/null 2>&1; then
        echo "❌ GitHub CLI (gh) is not installed." >&2
        exit 1
    fi

    if ! gh auth status >/dev/null 2>&1; then
        echo "❌ GitHub CLI is not authenticated. Run:\n   gh auth login --scopes repo,project --web" >&2
        exit 1
    fi

    local scopes
    scopes=$(gh auth status -t 2>/dev/null | grep -i 'scopes' | cut -d ':' -f2- | tr -d ' ')

    for scope in "${REQUIRED_SCOPES[@]}"; do
        if [[ "$scopes" != *"$scope"* ]]; then
            echo "❌ Missing '$scope' scope. Re-authenticate with:\n   gh auth login --scopes repo,project --web" >&2
            exit 1
        fi
    done
}

run_project_manager() {
    local script="${GH_PROJECT_SCRIPT:-../ghcommon/scripts/unified_github_project_manager_v2.py}"
    if [[ ! -f "$script" ]]; then
        echo "❌ Unified project manager script not found at $script" >&2
        exit 1
    fi
    python3 "$script" "$@"
}

if [[ "${1:-}" == "--help" ]]; then
    usage
    exit 0
fi

check_scopes
run_project_manager "$@"

