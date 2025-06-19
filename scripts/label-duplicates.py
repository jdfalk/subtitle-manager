#!/usr/bin/env python3

import os
from collections import defaultdict

import requests

GITHUB_TOKEN = os.getenv('GITHUB_TOKEN')  # Set this in your environment
REPO = "jdfalk/subtitle-manager"
HEADERS = {
    "Authorization": f"token {GITHUB_TOKEN}",
    "Accept": "application/vnd.github.v3+json"
}

def get_closed_issues():
    issues = []
    page = 1
    while True:
        url = f"https://api.github.com/repos/{REPO}/issues"
        params = {
            "state": "closed",
            "per_page": 100,
            "page": page
        }
        r = requests.get(url, headers=HEADERS, params=params)
        data = r.json()
        if not data or r.status_code != 200:
            break
        issues.extend(data)
        if len(data) < 100:
            break
        page += 1
    # Exclude PRs
    return [issue for issue in issues if "pull_request" not in issue]

def find_duplicates(issues):
    title_map = defaultdict(list)
    for issue in issues:
        key = issue['title'].strip().lower()
        title_map[key].append(issue)
    # Only keep those with more than one
    return {k: v for k, v in title_map.items() if len(v) > 1}

def apply_duplicate_label(issue_number):
    url = f"https://api.github.com/repos/{REPO}/issues/{issue_number}/labels"
    r = requests.post(url, headers=HEADERS, json={"labels": ["duplicate"]})
    if r.status_code == 200 or r.status_code == 201:
        print(f"  ✓ Labeled issue #{issue_number} as duplicate.")
    else:
        print(f"  ✗ Failed to label issue #{issue_number}: {r.content}")

def main():
    issues = get_closed_issues()
    duplicates = find_duplicates(issues)
    for title, dups in duplicates.items():
        # Sort by issue number (older first)
        dups_sorted = sorted(dups, key=lambda i: i['number'])
        # Keep the first (oldest), label the rest as duplicates
        print(f"\nDuplicate Title: '{title}'")
        for issue in dups_sorted[1:]:
            labels = [lbl['name'] for lbl in issue['labels']]
            if "duplicate" in labels:
                print(f"  - Issue #{issue['number']} already labeled duplicate, skipping.")
            else:
                print(f"  - Labeling issue #{issue['number']} ({issue['html_url']}) as duplicate...")
                apply_duplicate_label(issue['number'])

if __name__ == "__main__":
    main()
