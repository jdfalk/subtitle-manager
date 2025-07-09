#!/bin/bash
# Minimal issue update library for offline use
run_issue_update() {
  action="$1"; shift
  case "$action" in
    create)
      local title="$1"; shift
      local body="$1"; shift
      local labels="$1"; shift || true
      local guid=$(cat /proc/sys/kernel/random/uuid)
      local legacy_guid="create-$(echo "$title" | tr ' ' '-' | tr 'A-Z' 'a-z')-$(date +%Y-%m-%d)"
      mkdir -p .github/issue-updates
      local label_json=""
      IFS=',' read -ra arr <<< "$labels"
      for l in "${arr[@]}"; do
        [ -n "$l" ] && label_json+="\"$l\"," 
      done
      label_json="${label_json%,}"
      cat > ".github/issue-updates/${guid}.json" <<JSON
{
  "action": "create",
  "title": "$title",
  "body": "$body",
  "labels": [${label_json}],
  "guid": "$guid",
  "legacy_guid": "$legacy_guid"
}
JSON
      echo "Created .github/issue-updates/${guid}.json"
      ;;
    *)
      echo "Unsupported action: $action" >&2
      return 1
      ;;
  esac
}
