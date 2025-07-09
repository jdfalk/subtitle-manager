#!/bin/bash
# file: scripts/create-issue-update-library.sh
# simplified library for issue updates

run_issue_update() {
    action=$1
    shift
    case "$action" in
        create)
            title="$1"
            body="$2"
            labels="$3"
            if command -v uuidgen >/dev/null 2>&1; then
                uuid=$(uuidgen | tr '[:upper:]' '[:lower:]')
            else
                uuid=$(python3 - <<'EOF'
import uuid
print(str(uuid.uuid4()))
EOF
)
            fi
            legacy_guid="create-$(echo "$title" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/--*/-/g' | sed 's/^\-|-$//g')-$(date +%Y-%m-%d)"
            mkdir -p .github/issue-updates
            file=".github/issue-updates/${uuid}.json"
            cat > "$file" <<EOF
{
  "action": "create",
  "title": "$title",
  "body": "$body",
  "labels": [$(echo "$labels" | sed 's/,/", "/g' | sed 's/^/"/;s/$/"/')],
  "guid": "$uuid",
  "legacy_guid": "$legacy_guid"
}
EOF
            echo "✅ Created: $file"
            ;;
        *)
            echo "unsupported action" >&2
            return 1
            ;;
    esac
}
