#!/usr/bin/env python3
"""
# file: .github/scripts/close-duplicates.py

DEPRECATED: This script has been superseded by the unified issue_manager.py

This script is kept for reference but functionality has been moved to:
- .github/scripts/issue_manager.py (close-duplicates command)
- .github/workflows/unified-issue-management.yml

The unified system provides:
- Better error handling and logging
- Integration with other issue management operations
- GUID-based duplicate prevention
- Comprehensive summary generation

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
    print("⚠️  DEPRECATION WARNING")
    print("This script has been superseded by issue_manager.py")
    print("Please use: python issue_manager.py close-duplicates")
    print("Or use the unified-issue-management.yml workflow")
    print("")
    print("Proceeding with legacy script execution...")
    print("")

    # Get environment variables
    gh_token = os.environ.get('GH_TOKEN')
    repo = os.environ.get('REPO')

    if not gh_token or not repo:
        print("GH_TOKEN and REPO must be set", file=sys.stderr)
        if not gh_token:
            print("GH_TOKEN is missing", file=sys.stderr)
        if not repo:
            print("REPO is missing", file=sys.stderr)
        sys.exit(1)

    print(f"Using repo: {repo}")
    print(f"Token length: {len(gh_token)} characters")

    # Detect token type and set appropriate authorization header
    # Fine-grained tokens start with "github_pat_"
    # Classic tokens start with "ghp_" or "gho_" or "ghs_"
    if gh_token.startswith('github_pat_'):
        auth_header = f'token {gh_token}'
        token_type = 'fine-grained'
    else:
        auth_header = f'Bearer {gh_token}'
        token_type = 'classic'

    print(f"Detected {token_type} token format")

    # Set up headers for GitHub API
    headers = {
        'Authorization': auth_header,
        'Accept': 'application/vnd.github+json',
        'X-GitHub-Api-Version': '2022-11-28'
    }

    # Test API access first
    test_url = f"https://api.github.com/repos/{repo}"
    try:
        test_response = requests.get(test_url, headers=headers)
        if test_response.status_code == 401:
            print("Error: Invalid or expired GitHub token", file=sys.stderr)
            print("Please check your GH_TOKEN environment variable", file=sys.stderr)
            print(f"Detected {token_type} token format", file=sys.stderr)
            if token_type == 'fine-grained':
                print("Fine-grained tokens need 'Contents' and 'Issues' permissions", file=sys.stderr)
            sys.exit(1)
        elif test_response.status_code == 404:
            print(f"Error: Repository '{repo}' not found or not accessible", file=sys.stderr)
            sys.exit(1)
        test_response.raise_for_status()
        print("✓ GitHub API access verified")
    except requests.exceptions.RequestException as e:
        print(f"Error testing GitHub API access: {e}", file=sys.stderr)
        sys.exit(1)

    # Fetch all open issues
    print("Fetching open issues...")
    try:
        issues = fetch_all_issues(repo, headers)
    except requests.exceptions.HTTPError as e:
        print(f"Error fetching issues: {e}", file=sys.stderr)
        sys.exit(1)
    print(f"Found {len(issues)} open issues")

    # Group issues by title
    title_groups = group_issues_by_title(issues)

    # Check for dry-run mode
    dry_run = os.environ.get('DRY_RUN', 'false').lower() in ('true', '1', 'yes')
    if dry_run:
        print("🔍 Running in DRY-RUN mode - no issues will be closed")

    # Find and close duplicates
    duplicates_found = False
    for title, issue_list in title_groups.items():
        if len(issue_list) > 1:
            duplicates_found = True
            print(f"Found {len(issue_list)} issues with title: '{title}'")
            if dry_run:
                print_duplicate_plan(issue_list)
            else:
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


def print_duplicate_plan(issue_list: List[Dict[str, Any]]):
    """Print what would be done in dry-run mode."""
    issue_list.sort(key=lambda x: x['number'])
    canonical = issue_list[0]
    duplicates = issue_list[1:]

    print(f"  📌 Would keep issue #{canonical['number']} as canonical")
    for duplicate in duplicates:
        print(f"  🚫 Would close issue #{duplicate['number']} as duplicate")


if __name__ == '__main__':
    main()
