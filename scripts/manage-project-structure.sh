#!/bin/bash
# file: scripts/manage-project-structure.sh
# version: 1.0.0
# guid: 1a2b3c4d-5e6f-7890-abcd-ef1234567890

set -euo pipefail

# Manage GitHub Project structure and repository linking
# Yes, projects can be linked to multiple repositories!

ORG="${ORG:-jdfalk}"

# Project structure - projects that should be linked to multiple repos
declare -A multi_repo_projects=(
    ["gcommon Refactor"]="gcommon subtitle-manager"
    ["Security & Logging"]="gcommon subtitle-manager ghcommon"
)

# Project structure - single repo projects
declare -A single_repo_projects=(
    ["Metadata Editor"]="subtitle-manager"
    ["Whisper Container Integration"]="subtitle-manager"
    ["ghcommon Cleanup"]="ghcommon"
    ["ghcommon Core Improvements"]="ghcommon"
    ["ghcommon Testing & Quality"]="ghcommon"
)

# Function to link a repository to a project
link_repo_to_project() {
    local project_number="$1"
    local repo_name="$2"
    local project_title="$3"

    echo "Linking $repo_name to project #$project_number ($project_title)..."
    if gh project link "$project_number" --owner "$ORG" --repo "$repo_name" >/dev/null 2>&1; then
        echo "✅ Successfully linked $repo_name"
    else
        echo "⚠️ Failed to link $repo_name (might already be linked)"
    fi
}

# Function to get project number by title
get_project_number() {
    local title="$1"
    gh project list --owner "$ORG" --format json | jq -r ".projects[] | select(.title == \"$title\") | .number"
}

# Function to setup project structure
setup_project_structure() {
    echo "🔗 Setting up project repository links..."

    # Handle multi-repo projects
    for project_title in "${!multi_repo_projects[@]}"; do
        repos="${multi_repo_projects[$project_title]}"
        project_num=$(get_project_number "$project_title")

        if [[ -n "$project_num" && "$project_num" != "null" ]]; then
            echo ""
            echo "📋 Project: $project_title (#$project_num)"
            echo "   Repos: $repos"

            for repo in $repos; do
                link_repo_to_project "$project_num" "$repo" "$project_title"
            done
        else
            echo "⚠️ Project not found: $project_title"
        fi
    done

    # Handle single-repo projects
    for project_title in "${!single_repo_projects[@]}"; do
        repo="${single_repo_projects[$project_title]}"
        project_num=$(get_project_number "$project_title")

        if [[ -n "$project_num" && "$project_num" != "null" ]]; then
            echo ""
            echo "📋 Project: $project_title (#$project_num)"
            echo "   Repo: $repo"
            link_repo_to_project "$project_num" "$repo" "$project_title"
        else
            echo "⚠️ Project not found: $project_title"
        fi
    done
}

# Function to show current project structure
show_project_structure() {
    echo "📊 Current Project Structure:"
    echo ""

    # Get all projects
    gh project list --owner "$ORG" --format json | jq -r '.projects[] | "\(.number):\(.title)"' | while read -r line; do
        IFS=':' read -r number title <<< "$line"
        echo "📋 #$number: $title"

        # Try to get linked repositories (this might require GraphQL API)
        # For now, just show based on our expected structure
        if [[ -n "${multi_repo_projects[$title]:-}" ]]; then
            echo "   🔗 Multi-repo: ${multi_repo_projects[$title]}"
        elif [[ -n "${single_repo_projects[$title]:-}" ]]; then
            echo "   🔗 Single repo: ${single_repo_projects[$title]}"
        else
            echo "   🔗 Unknown structure"
        fi
        echo ""
    done
}

# Main function
main() {
    echo "🚀 GitHub Project Structure Manager"
    echo ""

    case "${1:-setup}" in
        "setup")
            setup_project_structure
            ;;
        "show")
            show_project_structure
            ;;
        *)
            echo "Usage: $0 [setup|show]"
            echo "  setup - Link repositories to projects"
            echo "  show  - Display current project structure"
            exit 1
            ;;
    esac

    echo ""
    echo "🎉 Operation completed!"
}

main "$@"
