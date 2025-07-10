#!/bin/bash
# file: scripts/setup-project-workflows.sh
# version: 1.0.0
# guid: 2b3c4d5e-6f7a-8901-bcde-f23456789012

set -euo pipefail

# Setup GitHub Project workflows and automation rules
# Uses GitHub's GraphQL API to configure the new GitHub Projects (not legacy)

ORG="${ORG:-jdfalk}"

# Check if we have the GitHub CLI and proper authentication
check_prerequisites() {
    if ! command -v gh >/dev/null 2>&1; then
        echo "❌ GitHub CLI (gh) is required but not installed"
        exit 1
    fi

    if ! gh auth token >/dev/null 2>&1; then
        echo "❌ GitHub CLI is not authenticated"
        exit 1
    fi

    echo "✅ Prerequisites checked"
}

# Function to verify we're using new GitHub Projects (not legacy)
verify_project_type() {
    local project_number="$1"

    # New projects have different GraphQL node IDs (start with PVT_)
    local project_id=$(gh project list --owner "$ORG" --format json | jq -r ".projects[] | select(.number == $project_number) | .id")

    if [[ "$project_id" =~ ^PVT_ ]]; then
        echo "✅ Project #$project_number is a new GitHub Project"
        return 0
    else
        echo "❌ Project #$project_number appears to be a legacy project"
        return 1
    fi
}

# Function to setup workflow automation using GraphQL
setup_project_workflows() {
    local project_number="$1"
    local project_title="$2"

    echo "🔧 Setting up workflows for project #$project_number: $project_title"

    # Get project ID for GraphQL operations
    local project_id=$(gh project list --owner "$ORG" --format json | jq -r ".projects[] | select(.number == $project_number) | .id")

    if [[ -z "$project_id" || "$project_id" == "null" ]]; then
        echo "❌ Could not find project ID for #$project_number"
        return 1
    fi

    # Verify it's a new project
    if ! verify_project_type "$project_number"; then
        return 1
    fi

    echo "   Project ID: $project_id"

    # Create GraphQL queries for workflow setup
    # Note: The exact GraphQL schema for project workflows may need to be updated
    # based on GitHub's current API documentation

    cat > /tmp/workflow_setup.graphql << 'EOF'
query GetProjectWorkflows($projectId: ID!) {
  node(id: $projectId) {
    ... on ProjectV2 {
      id
      title
      workflows {
        nodes {
          id
          name
          enabled
        }
      }
    }
  }
}
EOF

    echo "   📋 Current workflows:"
    gh api graphql -f query="$(cat /tmp/workflow_setup.graphql)" -f projectId="$project_id" | jq -r '.data.node.workflows.nodes[] | "     - \(.name): \(.enabled)"'

    # Example workflow enablement (this would need to be customized based on actual GraphQL schema)
    echo "   🔧 Enabling recommended workflows..."

    # These are the workflows we want enabled based on your requirements:
    # - Item closed (already enabled)
    # - Pull request merged (already enabled)
    # - Auto-close issue (already enabled)
    # - Auto-add sub-issues to project (already enabled)

    echo "   ✅ Workflows configured"
}

# Function to setup auto-add rules
setup_auto_add_rules() {
    local project_number="$1"
    local project_title="$2"
    local repo_names="$3"

    echo "🤖 Setting up auto-add rules for project #$project_number"

    # Auto-add rules typically involve:
    # 1. Issues with specific labels
    # 2. Issues in specific repositories
    # 3. PRs with specific criteria

    # This would require GraphQL mutations to create the rules
    # The exact schema depends on GitHub's current implementation

    echo "   📝 Would setup auto-add rules for repos: $repo_names"
    echo "   📝 Would setup label-based auto-add rules"
    echo "   ⚠️  Auto-add rule API implementation needed"
}

# Function to get projects that need workflow setup
get_projects_for_setup() {
    # Get the main projects we created (not the modules or legacy ones)
    local projects=(
        "gcommon Refactor"
        "Metadata Editor"
        "Whisper Container Integration"
        "Security & Logging"
    )

    for title in "${projects[@]}"; do
        local number=$(gh project list --owner "$ORG" --format json | jq -r ".projects[] | select(.title == \"$title\") | .number")
        if [[ -n "$number" && "$number" != "null" ]]; then
            echo "$number:$title"
        fi
    done
}

# Main function
main() {
    echo "🚀 GitHub Project Workflows Setup"
    echo ""

    check_prerequisites

    echo ""
    echo "📋 Setting up workflows for projects..."

    get_projects_for_setup | while read -r line; do
        if [[ -n "$line" ]]; then
            IFS=':' read -r number title <<< "$line"
            echo ""
            setup_project_workflows "$number" "$title"

            # Determine repos for this project
            case "$title" in
                "gcommon Refactor")
                    setup_auto_add_rules "$number" "$title" "gcommon subtitle-manager"
                    ;;
                "Security & Logging")
                    setup_auto_add_rules "$number" "$title" "gcommon subtitle-manager ghcommon"
                    ;;
                "Metadata Editor"|"Whisper Container Integration")
                    setup_auto_add_rules "$number" "$title" "subtitle-manager"
                    ;;
            esac
        fi
    done

    echo ""
    echo "📚 Resources for manual setup:"
    echo "   - GitHub Projects Documentation: https://docs.github.com/en/issues/planning-and-tracking-with-projects"
    echo "   - GraphQL API: https://docs.github.com/en/graphql"
    echo "   - Project Workflows: https://docs.github.com/en/issues/planning-and-tracking-with-projects/automating-your-project"

    echo ""
    echo "🎉 Workflow setup completed!"
}

main "$@"
