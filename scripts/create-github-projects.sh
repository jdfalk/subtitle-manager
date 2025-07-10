#!/bin/bash
# file: scripts/create-github-projects.sh
# version: 2.0.3
# guid: 81edeeb0-a3d2-4393-9ace-8da101bb8f7d

set -euo pipefail

# GitHub Project Setup Script
#
# This script creates GitHub Projects for Subtitle Manager and assigns
# relevant issues. It uses the GitHub CLI and automatically handles
# authentication checks.
#
# Usage:
#   ./scripts/create-github-projects.sh
#
# Environment variables:
#   ORG  - GitHub organization or user (default: derived from git remote)
#   REPO - Repository name (default: derived from git remote)

# Check for required tools
for tool in gh jq; do
    if ! command -v "$tool" >/dev/null 2>&1; then
        echo "❌ Required tool '$tool' is not installed" >&2
        exit 1
    fi
done

# Extract org and repo from git remote URL
_git_remote_url=$(git config --get remote.origin.url 2>/dev/null || echo "")
if [[ "$_git_remote_url" =~ github\.com[:/]([^/]+)/([^/]+)(\.git)?$ ]]; then
    ORG="${ORG:-${BASH_REMATCH[1]}}"
    REPO_RAW="${BASH_REMATCH[2]}"
    # Remove .git suffix if present
    REPO="${REPO:-${REPO_RAW%.git}}"
else
    ORG="${ORG:-jdfalk}"
    REPO="${REPO:-subtitle-manager}"
fi

# Check GitHub CLI authentication and project scope
setup_auth() {
    echo "Checking GitHub CLI authentication..."

    if ! gh auth token >/dev/null 2>&1; then
        echo "❌ GitHub CLI is not authenticated."
        echo "Please run: gh auth login"
        echo "Make sure to select 'repo' and 'project' scopes when prompted."
        exit 1
    fi

    echo "Checking project permissions..."
    if ! gh project list --owner "$ORG" >/dev/null 2>&1; then
        echo "⚠️ Missing required project scopes. Refreshing authentication..."
        if gh auth refresh -s project,read:project; then
            echo "✅ Authentication refreshed with project scopes"
        else
            echo "❌ Failed to refresh authentication with project scopes"
            echo "Please run manually: gh auth refresh -s project,read:project"
            exit 1
        fi
    fi

    export GH_TOKEN=$(gh auth token)
    echo "✅ GitHub CLI authenticated successfully with project access"
}

# Create a project and add a description
create_project() {
    local title="$1"
    local description="$2"

    echo "Creating project: $title" >&2
    local project_data
    project_data=$(gh project create --owner "$ORG" --title "$title" --format json 2>/dev/null)
    local project_number
    project_number=$(echo "$project_data" | jq -r '.number')

    if [[ -n "$project_number" && "$project_number" != "null" ]]; then
        gh project edit "$project_number" --owner "$ORG" --description "$description" >/dev/null 2>&1
        echo "$project_number"
    else
        echo "❌ Failed to create project: $title" >&2
        return 1
    fi
}

# Retrieve a project number by title
get_project_number() {
    local title="$1"
    gh project list --owner "$ORG" --format json | jq -r ".projects[] | select(.title == \"$title\") | .number"
}

# Link the repository to a project
link_repository() {
    local project_number="$1"
    echo "Linking repository $REPO to project #$project_number..." >&2
    if gh project link "$project_number" --owner "$ORG" --repo "$REPO" >/dev/null 2>&1; then
        echo "✅ Successfully linked repository" >&2
    else
        echo "⚠️ Failed to link repository $REPO to project #${project_number}" >&2
    fi
}

# Add an issue to a project
add_issue() {
    local project_number="$1"
    local issue_number="$2"

    # First check if the issue exists
    echo "Checking if issue #$issue_number exists..." >&2
    if ! gh issue view "$issue_number" --repo "$ORG/$REPO" >/dev/null 2>&1; then
        echo "⚠️ Issue #$issue_number does not exist, skipping..." >&2
        return 0
    fi

    echo "Adding issue #$issue_number to project #$project_number..." >&2
    if gh project item-add "$project_number" --owner "$ORG" --url "https://github.com/$ORG/$REPO/issues/$issue_number" >/dev/null 2>&1; then
        echo "✅ Added issue #$issue_number" >&2
    else
        echo "⚠️ Failed to add issue #$issue_number to project $project_number" >&2
        echo "   Debug: trying to add https://github.com/$ORG/$REPO/issues/$issue_number" >&2
    fi
}

# Main entry point
main() {
    echo "🚀 Starting GitHub Projects creation for $ORG/$REPO"

    setup_auth

    declare -a project_numbers=()

    echo ""
    echo "Creating projects..."

    # Check for existing projects first, create only if they don't exist
    project_titles=("gcommon Refactor" "Metadata Editor" "Whisper Container Integration" "Security & Logging")
    project_descriptions=(
        "Track migration to gcommon modules and protobuf types"
        "Manual metadata editing and search improvements"
        "Container-based Whisper ASR service"
        "Security enhancements and logging improvements"
    )

    for i in "${!project_titles[@]}"; do
        title="${project_titles[$i]}"
        description="${project_descriptions[$i]}"

        # Check if project already exists
        existing_num=$(get_project_number "$title")
        if [[ -n "$existing_num" && "$existing_num" != "null" ]]; then
            project_numbers+=("$existing_num")
            echo "✅ Found existing project: $title (#$existing_num)"
        else
            # Create new project
            if project_num=$(create_project "$title" "$description"); then
                project_numbers+=("$project_num")
                echo "✅ Created new project: $title (#$project_num)"
            else
                echo "❌ Failed to create project: $title"
            fi
        fi
    done

    echo ""
    echo "Linking repositories and adding issues..."

    if [[ ${#project_numbers[@]} -gt 0 ]]; then
        link_repository "${project_numbers[0]}"
    fi

    # Add issues to projects - using direct array indexing
    echo "Adding issues to projects..."

    # Project 0: gcommon Refactor
    if [[ ${#project_numbers[@]} -gt 0 ]]; then
        for issue_num in 1255 891; do
            add_issue "${project_numbers[0]}" "$issue_num"
        done
    fi

    # Project 1: Metadata Editor
    if [[ ${#project_numbers[@]} -gt 1 ]]; then
        for issue_num in 1135 1330; do
            add_issue "${project_numbers[1]}" "$issue_num"
        done
    fi

    # Project 2: Whisper Container Integration
    if [[ ${#project_numbers[@]} -gt 2 ]]; then
        add_issue "${project_numbers[2]}" "1132"
    fi

    # Project 3: Security & Logging
    if [[ ${#project_numbers[@]} -gt 3 ]]; then
        add_issue "${project_numbers[3]}" "545"
    fi

    echo ""
    echo "🎉 Project creation completed!"
    echo "Created ${#project_numbers[@]} projects for $ORG/$REPO"

    echo ""
    echo "📋 Project summary:"
    gh project list --owner "$ORG" --format json | jq -r '.projects[] | "  • #\(.number): \(.title)"'
}

main "$@"
