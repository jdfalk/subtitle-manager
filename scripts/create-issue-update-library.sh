#!/bin/bash
# file: scripts/create-issue-update-library.sh
# version: 1.0.0
# guid: 1a2b3c4d-5e6f-7890-abcd-ef1234567890

# Combined library for issue update creation with robust UUID generation

generate_uuid() {
  if command -v uuidgen >/dev/null 2>&1; then
    uuidgen | tr '[:upper:]' '[:lower:]'
  elif [ -r /proc/sys/kernel/random/uuid ]; then
    cat /proc/sys/kernel/random/uuid
  else
    python3 - <<'EOF'
import uuid
print(str(uuid.uuid4()))
EOF
  fi
}

generate_legacy_guid() {
  local action="$1"
  local title="$2"
  if [ -n "$title" ]; then
    # Use title-based GUID for backwards compatibility
    printf "%s-%s-%s" "$action" "$(echo "$title" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/--*/-/g' | sed 's/^\-|-$//g')" "$(date +%Y-%m-%d)"
  else
    printf "%s-%s" "$action" "$(date +%Y%m%d%H%M%S)"
  fi
}

create_issue_file() {
  local action="$1"
  local uuid="$2"
  local title="$3"
  local body="$4"
  local labels="$5"
  local file_path=".github/issue-updates/${uuid}.json"
  local legacy_guid="$(generate_legacy_guid "$action" "$title")"

  # Build label JSON array
  local label_json=""
  if [ -n "$labels" ]; then
    IFS=',' read -ra arr <<< "$labels"
    for l in "${arr[@]}"; do
      [ -n "$l" ] && label_json+="\"$l\", "
    done
    label_json="${label_json%, }"
  fi

  mkdir -p .github/issue-updates
  cat > "$file_path" <<JSON
{
  "action": "$action",
  "title": "$title",
  "body": "$body",
  "labels": [${label_json}],
  "guid": "$uuid",
  "legacy_guid": "$legacy_guid"
}
JSON
  echo "✅ Created: $file_path"
}

run_issue_update() {
  local action="$1"
  shift
  case "$action" in
    create)
      local title="$1"
      local body="$2"
      local labels="$3"
      local uuid="$(generate_uuid)"
      create_issue_file "$action" "$uuid" "$title" "$body" "$labels"
      ;;
    *)
      echo "unsupported action" >&2
      return 1
      ;;
  esac
}
