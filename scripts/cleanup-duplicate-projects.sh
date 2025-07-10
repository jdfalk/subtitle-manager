#!/bin/bash
# file: scripts/cleanup-duplicate-projects.sh
# version: 1.0.0
# guid: 9a1b2c3d-4e5f-6789-abcd-ef0123456789

set -euo pipefail

# Clean up duplicate GitHub projects, keeping only the most recent ones
ORG="${ORG:-jdfalk}"

# Function to delete a project
delete_project() {
    local project_number="$1"
    local title="$2"
    echo "Deleting project #$project_number: $title"
    if gh project delete "$project_number" --owner "$ORG" >/dev/null 2>&1; then
        echo "✅ Deleted project #$project_number"
    else
        echo "❌ Failed to delete project #$project_number"
    fi
}

echo "🧹 Cleaning up duplicate GitHub projects for $ORG"

# Based on the list, keep the highest numbered (most recent) projects and delete the rest
# Keep: #32, #31, #30, #29 (most recent of each type)
# Delete: #28, #27, #26, #25, #24, #23, #22, #15, #14, #13, #12

declare -a projects_to_delete=(
    "28:Whisper Container Integration"
    "27:Metadata Editor"
    "26:gcommon Refactor"
    "25:Security & Logging"
    "24:Whisper Container Integration"
    "23:Metadata Editor"
    "22:gcommon Refactor"
    "15:Security & Logging"
    "14:Whisper Container Integration"
    "13:Metadata Editor"
    "12:gcommon Refactor"
)

echo "Will delete the following duplicate projects:"
for item in "${projects_to_delete[@]}"; do
    IFS=':' read -r number title <<< "$item"
    echo "  - #$number: $title"
done

echo ""
read -p "Continue with deletion? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    for item in "${projects_to_delete[@]}"; do
        IFS=':' read -r number title <<< "$item"
        delete_project "$number" "$title"
        sleep 1  # Rate limiting
    done
    echo "🎉 Cleanup completed!"
else
    echo "❌ Cleanup cancelled"
fi
