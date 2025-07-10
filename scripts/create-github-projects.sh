#!/bin/bash
# file: scripts/create-github-projects.sh
# version: 2.0.0
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

ORG="${ORG:-$(git config --get remote.origin.url | sed -n 's#.*github.com[:/]\(.*\)/.*#\1#p' | cut -d/ -f1)}"
REPO="${REPO:-$(git config --get remote.origin.url | sed -n 's#.*/\(.*\)\.git#\1#p')}"

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

    echo "Creating project: $title"
    local project_data
    project_data=$(gh project create --owner "$ORG" --title "$title" --format json)
    local project_number
    project_number=$(echo "$project_data" | jq -r '.number')

    if [[ -n "$project_number" && "$project_number" != "null" ]]; then
        gh project edit "$project_number" --owner "$ORG" --description "$description"
        echo "$project_number"
    else
        echo "❌ Failed to create project: $title"
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
    gh project link --owner "$ORG" --repo "$REPO" "$project_number" || {
        echo "⚠️ Failed to link repository $REPO to project #${project_number}"
    }
}

# Add an issue to a project
add_issue() {
    local project_number="$1"
    local issue_number="$2"
    gh project item-add "$project_number" --owner "$ORG" --repo "$REPO" --issue "$issue_number" || {
        echo "⚠️ Failed to add issue #$issue_number to project $project_number"
    }
}

# Main entry point
main() {
    echo "🚀 Starting GitHub Projects creation for $ORG/$REPO"

    setup_auth

    declare -a project_numbers=()

    echo ""
    echo "Creating projects..."

    if project_num=$(create_project "gcommon Refactor" "Track migration to gcommon modules and protobuf types"); then
        project_numbers+=("$project_num")
    fi

    if project_num=$(create_project "Metadata Editor" "Manual metadata editing and search improvements"); then
        project_numbers+=("$project_num")
    fi

    if project_num=$(create_project "Whisper Container Integration" "Container-based Whisper ASR service"); then
        project_numbers+=("$project_num")
    fi

    if project_num=$(create_project "Security & Logging" "Security enhancements and logging improvements"); then
        project_numbers+=("$project_num")
    fi

    if [[ ${#project_numbers[@]} -gt 0 ]]; then
        link_repository "${project_numbers[0]}"
    fi

    # Example issue assignments (update issue numbers as needed)
    add_issue "${project_numbers[0]}" 1255
    add_issue "${project_numbers[0]}" 891
    add_issue "${project_numbers[1]}" 1135
    add_issue "${project_numbers[1]}" 1330
    add_issue "${project_numbers[2]}" 1132
    add_issue "${project_numbers[3]}" 545

    echo ""
    echo "🎉 Project creation completed!"
    echo "Created ${#project_numbers[@]} projects for $ORG/$REPO"

    echo ""
    echo "📋 Project summary:"
    gh project list --owner "$ORG" --format json | jq -r '.projects[] | "  • #\(.number): \(.title)"'
}

main "$@"
