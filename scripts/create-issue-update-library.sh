#!/bin/bash
# file: scripts/create-issue-update-library.sh
# version: 1.0.0
# guid: 1a2b3c4d-5e6f-7890-abcd-ef1234567890

# Minimal issue update library for offline use
run_issue_update() {
    local action="$1"; shift
    case "$action" in
        create)
            local title="$1"; shift
            local body="$1"; shift
            local labels="$1"; shift || true
            
            # Generate UUID - try multiple methods for compatibility
            local guid
            if command -v uuidgen >/dev/null 2>&1; then
                guid=$(uuidgen | tr '[:upper:]' '[:lower:]')
            elif [ -f /proc/sys/kernel/random/uuid ]; then
                guid=$(cat /proc/sys/kernel/random/uuid)
            else
                guid=$(python3 - <<'EOF'
import uuid
print(str(uuid.uuid4()))
EOF
)
            fi
            
            local legacy_guid="create-$(echo "$title" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/--*/-/g' | sed 's/^\-|-$//g')-$(date +%Y-%m-%d)"
            mkdir -p .github/issue-updates
            
            # Build label JSON array
            local label_json=""
            if [ -n "$labels" ]; then
                IFS=',' read -ra arr <<< "$labels"
                for l in "${arr[@]}"; do
                    [ -n "$l" ] && label_json+="\"$l\", " 
                done
                label_json="${label_json%, }"
            fi
            
            local file=".github/issue-updates/${guid}.json"
            cat > "$file" <<JSON
{
  "action": "create",
  "title": "$title",
  "body": "$body",
  "labels": [${label_json}],
  "guid": "$guid",
  "legacy_guid": "$legacy_guid"
}
JSON
            echo "✅ Created: $file"
            ;;
        *)
            echo "Unsupported action: $action" >&2
            return 1
            ;;
    esac
}
