#!/usr/bin/env python3
"""
# file: .github/scripts/copilot_tickets.py
Create or update GitHub issues based on Copilot review comments.

Environment Variables:
  GH_TOKEN - GitHub token with repo access
  REPO - repository in owner/name format
  GITHUB_EVENT_NAME - webhook event name
  GITHUB_EVENT_PATH - path to the event payload
"""

import json
import os
import sys
from typing import Any, Dict, List

import requests


COPILOT_USER = "github-copilot[bot]"
LABEL = "copilot-review"
API_VERSION = "2022-11-28"


def load_event() -> Dict[str, Any]:
    """Load the GitHub event payload."""
    path = os.environ.get("GITHUB_EVENT_PATH")
    if not path:
        print("GITHUB_EVENT_PATH not set", file=sys.stderr)
        sys.exit(1)
    with open(path, "r", encoding="utf-8") as f:
        return json.load(f)


def get_headers(token: str) -> Dict[str, str]:
    """Return HTTP headers for the GitHub API."""
    return {
        "Authorization": f"Bearer {token}",
        "Accept": "application/vnd.github+json",
        "X-GitHub-Api-Version": API_VERSION,
    }


def create_issue(repo: str, headers: Dict[str, str], title: str, body: str) -> None:
    """Create a new GitHub issue."""
    url = f"https://api.github.com/repos/{repo}/issues"
    data = {"title": title, "body": body, "labels": [LABEL]}
    response = requests.post(url, headers=headers, json=data, timeout=10)
    if response.status_code == 201:
        issue = response.json()
        print(f"Created issue #{issue['number']}")
    else:
        print(f"Failed to create issue: {response.status_code}", file=sys.stderr)
        print(response.text, file=sys.stderr)


def update_issue(repo: str, headers: Dict[str, str], issue_number: int, body: str) -> None:
    """Update the body of an existing GitHub issue."""
    url = f"https://api.github.com/repos/{repo}/issues/{issue_number}"
    response = requests.patch(url, headers=headers, json={"body": body}, timeout=10)
    if response.status_code == 200:
        print(f"Updated issue #{issue_number}")
    else:
        print(f"Failed to update issue #{issue_number}: {response.status_code}", file=sys.stderr)
        print(response.text, file=sys.stderr)


def close_issue(repo: str, headers: Dict[str, str], issue_number: int) -> None:
    """Close an issue."""
    url = f"https://api.github.com/repos/{repo}/issues/{issue_number}"
    response = requests.patch(url, headers=headers, json={"state": "closed"}, timeout=10)
    if response.status_code == 200:
        print(f"Closed issue #{issue_number}")
    else:
        print(f"Failed to close issue #{issue_number}: {response.status_code}", file=sys.stderr)
        print(response.text, file=sys.stderr)


def search_issue(repo: str, headers: Dict[str, str], key: str) -> Dict[str, Any] | None:
    """Search for an open issue containing the given key."""
    query = f"repo:{repo} label:{LABEL} state:open {key}"
    url = "https://api.github.com/search/issues"
    params = {"q": query}
    response = requests.get(url, headers=headers, params=params, timeout=10)
    response.raise_for_status()
    items = response.json().get("items")
    return items[0] if items else None


def build_body(comment: Dict[str, Any], lines: List[Dict[str, Any]]) -> str:
    """Build the issue body from comment text and metadata."""
    snippet = comment["body"]
    bullet_lines = [
        f"- id:{item['id']} [{item['path']}#L{item['line']}]({item['url']})"
        for item in lines
    ]
    data = {"comments": lines}
    json_block = json.dumps(data, separators=(",", ":"))
    return (
        f"Generated from [Copilot review comment]({comment['url']}).\n\n"
        f"```text\n{snippet}\n```\n\n"
        + "\n".join(bullet_lines)
        + f"\n\n<!-- copilot-data:{json_block} -->"
    )


def handle_comment_created(repo: str, headers: Dict[str, str], comment: Dict[str, Any]) -> None:
    """Process a new Copilot review comment."""
    if comment.get("user", {}).get("login") != COPILOT_USER:
        print("Not a Copilot comment; skipping")
        return

    key = comment["body"].strip().split("\n", 1)[0]
    existing = search_issue(repo, headers, key)

    line_info = {
        "id": comment["id"],
        "path": comment.get("path", ""),
        "line": comment.get("line", 0),
        "url": comment.get("html_url", ""),
    }

    if existing:
        issue_number = existing["number"]
        body = existing["body"]
        if "<!-- copilot-data:" in body:
            start = body.index("<!-- copilot-data:") + len("<!-- copilot-data:")
            end = body.index("-->", start)
            data = json.loads(body[start:end])
            comments = data.get("comments", [])
        else:
            comments = []
        if any(item["id"] == line_info["id"] for item in comments):
            print("Comment already tracked")
            return
        comments.append(line_info)
        updated_body = build_body(comment, comments)
        update_issue(repo, headers, issue_number, updated_body)
    else:
        title = f"Copilot review: {key[:50]}"
        body = build_body(comment, [line_info])
        create_issue(repo, headers, title, body)


def handle_thread_resolved(repo: str, headers: Dict[str, str], thread: Dict[str, Any]) -> None:
    """Close or update issues when a review thread is resolved."""
    for comment in thread.get("comments", []):
        cid = comment["id"]
        search_key = f"id:{cid}"
        issue = search_issue(repo, headers, search_key)
        if not issue:
            continue
        body = issue["body"]
        start = body.index("<!-- copilot-data:") + len("<!-- copilot-data:")
        end = body.index("-->", start)
        data = json.loads(body[start:end])
        comments = [c for c in data.get("comments", []) if c["id"] != cid]
        if comments:
            updated_body = build_body(comment, comments)
            update_issue(repo, headers, issue["number"], updated_body)
        else:
            close_issue(repo, headers, issue["number"])


def main() -> None:
    """Entry point for the script."""
    token = os.environ.get("GH_TOKEN")
    repo = os.environ.get("REPO")
    event_name = os.environ.get("GITHUB_EVENT_NAME")
    if not token or not repo or not event_name:
        print("GH_TOKEN, REPO, and GITHUB_EVENT_NAME must be set", file=sys.stderr)
        sys.exit(1)
    headers = get_headers(token)
    event = load_event()
    action = event.get("action")

    if event_name == "pull_request_review_comment" and action == "created":
        handle_comment_created(repo, headers, event["comment"])
    elif event_name == "pull_request_review_thread" and action == "resolved":
        handle_thread_resolved(repo, headers, event["thread"])
    else:
        print("Event not handled")


if __name__ == "__main__":
    main()

