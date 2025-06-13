#!/usr/bin/env python3
"""
Close duplicate GitHub issues by title, keeping the lowest numbered issue open.

Parameters are provided via environment variables:
  GH_TOKEN - GitHub token with repo permissions
  REPO     - owner/repo name, e.g. "user/project"

The script fetches all open issues, groups them by title, and closes
duplicates while commenting with a reference to the canonical issue.
"""

import os
import sys
from collections import defaultdict
from typing import Any, Dict, List

import requests


def main():
    # Get environment variables
    gh_token = os.environ.get('GH_TOKEN')
    repo = os.environ.get('REPO')

    if not gh_token or not repo:
        print("GH_TOKEN and REPO must be set", file=sys.stderr)
        sys.exit(1)

    # Set up headers for GitHub API
    headers = {
        'Authorization': f'Bearer {gh_token}',
        'Accept': 'application/vnd.github+json',
        'X-GitHub-Api-Version': '2022-11-28'
    }

    # Fetch all open issues
    print("Fetching open issues...")
    issues = fetch_all_issues(repo, headers)
    print(f"Found {len(issues)} open issues")

    # Group issues by title
    title_groups = group_issues_by_title(issues)

    # Find and close duplicates
    duplicates_found = False
    for title, issue_list in title_groups.items():
        if len(issue_list) > 1:
            duplicates_found = True
            print(f"Found {len(issue_list)} issues with title: '{title}'")
            close_duplicates(repo, headers, issue_list)

    if not duplicates_found:
        print("No duplicate issues found")


def fetch_all_issues(repo: str, headers: Dict[str, str]) -> List[Dict[str, Any]]:
    """Fetch all open issues from the repository with pagination support."""
    all_issues = []
    page = 1
    per_page = 100

    while True:
        url = f"https://api.github.com/repos/{repo}/issues"
        params = {
            'state': 'open',
            'per_page': per_page,
            'page': page
        }

        response = requests.get(url, headers=headers, params=params)
        response.raise_for_status()

        issues = response.json()
        if not issues:
            break

        # Filter out pull requests (they have a 'pull_request' key)
        issues = [issue for issue in issues if 'pull_request' not in issue]
        all_issues.extend(issues)

        print(f"Fetched page {page}: {len(issues)} issues")

        # Check if we got fewer issues than requested (last page)
        if len(issues) < per_page:
            break

        page += 1

    return all_issues


def group_issues_by_title(issues: List[Dict[str, Any]]) -> Dict[str, List[Dict[str, Any]]]:
    """Group issues by their title."""
    title_groups = defaultdict(list)

    for issue in issues:
        title = issue['title'].strip()
        title_groups[title].append({
            'number': issue['number'],
            'title': title,
            'url': issue['html_url']
        })

    return title_groups


def close_duplicates(repo: str, headers: Dict[str, str], issue_list: List[Dict[str, Any]]):
    """Close duplicate issues, keeping the one with the lowest number."""
    # Sort by issue number and keep the first (lowest numbered) as canonical
    issue_list.sort(key=lambda x: x['number'])
    canonical = issue_list[0]
    duplicates = issue_list[1:]

    print(f"Keeping issue #{canonical['number']} as canonical")

    for duplicate in duplicates:
        print(f"Closing issue #{duplicate['number']} as duplicate of #{canonical['number']}")

        # Close the issue
        close_url = f"https://api.github.com/repos/{repo}/issues/{duplicate['number']}"
        close_response = requests.patch(
            close_url,
            headers=headers,
            json={'state': 'closed'}
        )

        if close_response.status_code == 200:
            print(f"Successfully closed issue #{duplicate['number']}")
        else:
            print(f"Failed to close issue #{duplicate['number']}: {close_response.status_code}")
            print(f"Response: {close_response.text}")
            continue

        # Add a comment explaining it's a duplicate
        comment_url = f"https://api.github.com/repos/{repo}/issues/{duplicate['number']}/comments"
        comment_body = f"Duplicate of #{canonical['number']}"
        comment_response = requests.post(
            comment_url,
            headers=headers,
            json={'body': comment_body}
        )

        if comment_response.status_code == 201:
            print(f"Added duplicate comment to issue #{duplicate['number']}")
        else:
            print(f"Failed to add comment to issue #{duplicate['number']}: {comment_response.status_code}")
            print(f"Response: {comment_response.text}")


if __name__ == '__main__':
    main()
