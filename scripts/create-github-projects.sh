#!/bin/bash
# file: scripts/create-github-projects.sh
# version: 1.0.1
# guid: e8ac48cd-67d1-49a8-a543-d23656009be8

set -euo pipefail

# Create GitHub Projects for open feature issues.
# Requires GitHub CLI with project scopes.
# Environment variables:
#   ORG  - GitHub organization (default: jdfalk)
#   REPO - Repository name (default: subtitle-manager)

ORG="${ORG:-jdfalk}"
REPO="${REPO:-subtitle-manager}"

check_requirements() {
  if ! command -v gh >/dev/null 2>&1; then
    echo "❌ gh CLI is required" >&2
    exit 1
  fi

  if ! gh auth status >/dev/null 2>&1; then
    echo "❌ gh CLI is not authenticated" >&2
    exit 1
  fi

  if ! gh auth status --scopes | grep -q 'project'; then
    echo "❌ gh CLI lacks project scopes" >&2
    exit 1
  fi
}

# Check whether project already exists
project_exists() {
  local title="$1"
  gh project list --owner "$ORG" --limit 100 --json title | jq -e ".[] | select(.title==\"$title\")" >/dev/null 2>&1
}

create_project_for_issue() {
  local number="$1"
  local title="Feature: $2"

  if project_exists "$title"; then
    echo "⚠️  Project already exists: $title"
    return
  fi

  echo "Creating project: $title"
  local pj_json
  pj_json=$(gh project create --title "$title" --owner "$ORG" --format json)
  local pj_number
  pj_number=$(echo "$pj_json" | jq -r '.number')
  local issue_id
  issue_id=$(gh issue view "$number" --repo "$ORG/$REPO" --json id -q '.id')
  gh project item-add "$pj_number" --owner "$ORG" --content-id "$issue_id" >/dev/null
}

main() {
  check_requirements

  echo "📋 Fetching open feature issues from $ORG/$REPO"
  gh issue list --repo "$ORG/$REPO" --label feature --state open --json number,title | \
    jq -c '.[]' | while read -r issue; do
      number=$(echo "$issue" | jq -r '.number')
      title=$(echo "$issue" | jq -r '.title')
      create_project_for_issue "$number" "$title"
      sleep 1
    done
  echo "✅ Project creation complete"
}

main "$@"
