# file: .github/scripts/close-duplicates.sh
#!/bin/bash

# Close duplicate GitHub issues by title, keeping the lowest numbered issue open.
# Parameters are provided via environment variables:
#   GH_TOKEN - GitHub token with repo permissions
#   REPO     - owner/repo name, e.g. "user/project"
# The script fetches all open issues, groups them by title, and closes
# duplicates while commenting with a reference to the canonical issue.

set -euo pipefail

if [[ -z "${GH_TOKEN:-}" || -z "${REPO:-}" ]]; then
  echo "GH_TOKEN and REPO must be set" >&2
  exit 1
fi

issues=$(curl -s -H "Authorization: Bearer $GH_TOKEN" \
  -H "Accept: application/vnd.github+json" \
  "https://api.github.com/repos/${REPO}/issues?state=open&per_page=100")

# TODO: handle pagination if more than 100 open issues exist

# Sort issues by title before grouping so that group_by works correctly.
echo "$issues" | jq -c '[.[] | {number, title}] | sort_by(.title) | group_by(.title)[] | select(length>1)' | while read -r group; do
    canonical=$(echo "$group" | jq 'min_by(.number)')
    canonical_num=$(echo "$canonical" | jq -r '.number')

    echo "$group" | jq -c '.[]' | while read -r issue; do
        num=$(echo "$issue" | jq -r '.number')
        if [[ "$num" != "$canonical_num" ]]; then
            echo "Closing #$num as duplicate of #$canonical_num"
            curl -s -X PATCH \
              -H "Authorization: Bearer $GH_TOKEN" \
              -H "Accept: application/vnd.github+json" \
              "https://api.github.com/repos/${REPO}/issues/$num" \
              -d '{"state":"closed"}' > /dev/null
            curl -s -X POST \
              -H "Authorization: Bearer $GH_TOKEN" \
              -H "Accept: application/vnd.github+json" \
              "https://api.github.com/repos/${REPO}/issues/$num/comments" \
              -d "$(printf '{"body":"Duplicate of #%s"}' "$canonical_num")" > /dev/null
        fi
    done
  done
