#!/usr/bin/env python3
"""
# file: .github/scripts/copilot_tickets.py
Create or update GitHub issues based on Copilot review comments.

This script manages GitHub issues that track Copilot review comments across PRs.
It creates issues for new Copilot comments and closes them when comments are resolved
or when PRs are merged.

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
    try:
        response = requests.post(url, headers=headers, json=data, timeout=10)
        if response.status_code == 201:
            issue = response.json()
            print(f"Created issue #{issue['number']}")
        else:
            print(f"Failed to create issue: {response.status_code}", file=sys.stderr)
            print(response.text, file=sys.stderr)
    except requests.RequestException as e:
        print(f"Network error creating issue: {e}", file=sys.stderr)
    except Exception as e:
        print(f"Unexpected error creating issue: {e}", file=sys.stderr)


def update_issue(repo: str, headers: Dict[str, str], issue_number: int, body: str) -> None:
    """Update the body of an existing GitHub issue."""
    url = f"https://api.github.com/repos/{repo}/issues/{issue_number}"
    try:
        response = requests.patch(url, headers=headers, json={"body": body}, timeout=10)
        if response.status_code == 200:
            print(f"Updated issue #{issue_number}")
        else:
            print(f"Failed to update issue #{issue_number}: {response.status_code}", file=sys.stderr)
            print(response.text, file=sys.stderr)
    except requests.RequestException as e:
        print(f"Network error updating issue #{issue_number}: {e}", file=sys.stderr)
    except Exception as e:
        print(f"Unexpected error updating issue #{issue_number}: {e}", file=sys.stderr)


def close_issue(repo: str, headers: Dict[str, str], issue_number: int) -> None:
    """Close an issue."""
    url = f"https://api.github.com/repos/{repo}/issues/{issue_number}"
    try:
        response = requests.patch(url, headers=headers, json={"state": "closed"}, timeout=10)
        if response.status_code == 200:
            print(f"Closed issue #{issue_number}")
        else:
            print(f"Failed to close issue #{issue_number}: {response.status_code}", file=sys.stderr)
            print(response.text, file=sys.stderr)
    except requests.RequestException as e:
        print(f"Network error closing issue #{issue_number}: {e}", file=sys.stderr)
    except Exception as e:
        print(f"Unexpected error closing issue #{issue_number}: {e}", file=sys.stderr)


def search_issue(repo: str, headers: Dict[str, str], key: str) -> Dict[str, Any] | None:
    """Search for an open issue containing the given key."""
    query = f"repo:{repo} label:{LABEL} state:open {key}"
    url = "https://api.github.com/search/issues"
    params = {"q": query}
    try:
        response = requests.get(url, headers=headers, params=params, timeout=10)
        response.raise_for_status()
        items = response.json().get("items")
        return items[0] if items else None
    except requests.RequestException as e:
        print(f"Network error searching for issues: {e}", file=sys.stderr)
        return None
    except Exception as e:
        print(f"Unexpected error searching for issues: {e}", file=sys.stderr)
        return None


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

    try:
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
            comments = []

            if "<!-- copilot-data:" in body:
                try:
                    start = body.index("<!-- copilot-data:") + len("<!-- copilot-data:")
                    end = body.index("-->", start)
                    data = json.loads(body[start:end])
                    comments = data.get("comments", [])
                except (ValueError, json.JSONDecodeError) as e:
                    print(f"Error parsing existing issue data: {e}", file=sys.stderr)
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
    except (ValueError, json.JSONDecodeError) as e:
        print(f"Error parsing comment data: {e}", file=sys.stderr)
    except Exception as e:
        print(f"Unexpected error handling created comment: {e}", file=sys.stderr)


def handle_comment_deleted(repo: str, headers: Dict[str, str], comment: Dict[str, Any]) -> None:
    """Process a deleted Copilot review comment."""
    if comment.get("user", {}).get("login") != COPILOT_USER:
        print("Not a Copilot comment; skipping")
        return

    cid = comment["id"]
    search_key = f"id:{cid}"
    issue = search_issue(repo, headers, search_key)
    if not issue:
        print(f"No issue found for comment {cid}")
        return

    try:
        # Remove this comment from the issue
        body = issue["body"]
        if "<!-- copilot-data:" not in body:
            print("No copilot data found in issue")
            return

        start = body.index("<!-- copilot-data:") + len("<!-- copilot-data:")
        end = body.index("-->", start)
        data = json.loads(body[start:end])
        comments = [c for c in data.get("comments", []) if c["id"] != cid]

        if comments:
            # Update issue with remaining comments
            updated_body = build_body(comment, comments)
            update_issue(repo, headers, issue["number"], updated_body)
        else:
            # Close issue if no comments remain
            close_issue(repo, headers, issue["number"])
    except (ValueError, json.JSONDecodeError) as e:
        print(f"Error parsing issue data for comment {cid}: {e}", file=sys.stderr)
    except Exception as e:
        print(f"Unexpected error handling deleted comment {cid}: {e}", file=sys.stderr)


def handle_pr_merged(repo: str, headers: Dict[str, str], pr: Dict[str, Any]) -> None:
    """Close all open Copilot issues for a merged PR."""
    if not pr.get("merged", False):
        print("PR not merged, skipping")
        return

    pr_number = pr["number"]
    print(f"Processing merged PR #{pr_number}")

    try:
        # Search for all open copilot issues mentioning this PR
        query = f"repo:{repo} label:{LABEL} state:open {pr_number}"
        url = "https://api.github.com/search/issues"
        params = {"q": query}
        response = requests.get(url, headers=headers, params=params, timeout=10)
        response.raise_for_status()

        issues = response.json().get("items", [])
        print(f"Found {len(issues)} open Copilot issues for PR #{pr_number}")

        closed_count = 0
        for issue in issues:
            try:
                # Add a comment noting the PR was merged
                comment_url = f"https://api.github.com/repos/{repo}/issues/{issue['number']}/comments"
                comment_data = {
                    "body": f"Closing as PR #{pr_number} was merged. Copilot feedback has been addressed."
                }
                requests.post(comment_url, headers=headers, json=comment_data, timeout=10)
                close_issue(repo, headers, issue["number"])
                closed_count += 1
            except Exception as e:
                print(f"Failed to close issue #{issue['number']}: {e}", file=sys.stderr)

        if closed_count > 0:
            print(f"Closed {closed_count} issues due to PR #{pr_number} merge")
    except requests.RequestException as e:
        print(f"Network error processing merged PR #{pr_number}: {e}", file=sys.stderr)
    except Exception as e:
        print(f"Unexpected error processing merged PR #{pr_number}: {e}", file=sys.stderr)


def handle_push_to_main(repo: str, headers: Dict[str, str], push: Dict[str, Any]) -> None:
    """Handle pushes to main branch - close any stale Copilot issues."""
    ref = push.get("ref", "")
    if not ref.endswith("/main") and not ref.endswith("/master"):
        print(f"Push to {ref} - not main/master, skipping")
        return

    # Find old open copilot issues (older than the push)
    push_timestamp = push.get("head_commit", {}).get("timestamp")
    if not push_timestamp:
        print("No timestamp in push event, skipping stale issue cleanup")
        return

    print(f"Processing push to {ref} at {push_timestamp}")
    query = f"repo:{repo} label:{LABEL} state:open"
    url = "https://api.github.com/search/issues"
    params = {"q": query}

    try:
        response = requests.get(url, headers=headers, params=params, timeout=10)
        response.raise_for_status()
    except requests.RequestException as e:
        print(f"Failed to search for issues: {e}", file=sys.stderr)
        return

    issues = response.json().get("items", [])
    print(f"Found {len(issues)} open Copilot issues")

    closed_count = 0
    for issue in issues:
        # Close issues that might be stale after main branch updates
        # This is a conservative approach - only close issues older than the push
        if issue.get("created_at", "") < push_timestamp:
            try:
                comment_url = f"https://api.github.com/repos/{repo}/issues/{issue['number']}/comments"
                comment_data = {
                    "body": "Closing as main branch was updated. Please re-run Copilot review if issues persist."
                }
                requests.post(comment_url, headers=headers, json=comment_data, timeout=10)
                close_issue(repo, headers, issue["number"])
                closed_count += 1
            except Exception as e:
                print(f"Failed to close issue #{issue['number']}: {e}", file=sys.stderr)

    if closed_count > 0:
        print(f"Closed {closed_count} stale issues after main branch push")


def main() -> None:
    """
    Entry point for the script.

    Handles different GitHub webhook events:
    - pull_request_review_comment: Track individual Copilot comments (create/edit/delete)
    - pull_request_review: Log review submissions (minimal action currently)
    - pull_request (closed): Close all Copilot issues when PR is merged
    - push (to main/master): Close stale Copilot issues after main branch updates
    """
    token = os.environ.get("GH_TOKEN")
    repo = os.environ.get("REPO")
    event_name = os.environ.get("GITHUB_EVENT_NAME")

    if not token or not repo or not event_name:
        print("GH_TOKEN, REPO, and GITHUB_EVENT_NAME must be set", file=sys.stderr)
        sys.exit(1)

    print(f"Processing {event_name} event for {repo}")

    headers = get_headers(token)

    try:
        event = load_event()
    except Exception as e:
        print(f"Failed to load event: {e}", file=sys.stderr)
        sys.exit(1)

    action = event.get("action")
    print(f"Event action: {action}")

    try:
        if event_name == "pull_request_review_comment":
            if action == "created":
                handle_comment_created(repo, headers, event["comment"])
            elif action == "deleted":
                handle_comment_deleted(repo, headers, event["comment"])
            elif action == "edited":
                # For edited comments, treat as delete + create
                handle_comment_deleted(repo, headers, event["comment"])
                handle_comment_created(repo, headers, event["comment"])
            else:
                print(f"Unhandled pull_request_review_comment action: {action}")
        elif event_name == "pull_request_review":
            # We mainly track individual comments, but reviews can contain summary info
            print(f"Pull request review {action} - no specific action taken")
        elif event_name == "pull_request" and action == "closed":
            handle_pr_merged(repo, headers, event["pull_request"])
        elif event_name == "push":
            handle_push_to_main(repo, headers, event)
        else:
            print(f"Event not handled: {event_name}/{action}")
    except Exception as e:
        print(f"Error processing event: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()

