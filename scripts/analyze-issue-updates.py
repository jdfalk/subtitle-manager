#!/usr/bin/env python3
"""
# file: scripts/analyze-issue-updates.py
Analyze issue update files for duplicates and provide migration insights.

This script scans all issue update directories and identifies:
1. Duplicate GUIDs across files
2. Duplicate content (same action + same target)
3. Files that might be duplicates of processed files
4. Summary statistics
"""

import json
import os
from collections import defaultdict
from typing import Dict, List, Set, Tuple


def load_json_file(file_path: str) -> List[Dict]:
    """Load a JSON file and return actions as a list."""
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            data = json.load(f)

        if isinstance(data, list):
            return data
        elif isinstance(data, dict) and 'action' in data:
            return [data]
        else:
            print(f"⚠️  Unexpected format in {file_path}")
            return []
    except Exception as e:
        print(f"❌ Error reading {file_path}: {e}")
        return []


def analyze_duplicates(base_dir: str) -> Dict:
    """Analyze all issue update files for duplicates."""

    # Directories to scan
    directories = [
        os.path.join(base_dir, ".github/issue-updates"),
        os.path.join(base_dir, ".github/issue-updates/processed")
    ]

    all_actions = []
    guid_to_files = defaultdict(list)
    content_signatures = defaultdict(list)
    file_count = 0

    for directory in directories:
        if not os.path.exists(directory):
            continue

        print(f"📁 Scanning {directory}")

        for filename in os.listdir(directory):
            if not filename.endswith('.json') or filename == 'README.json':
                continue

            file_path = os.path.join(directory, filename)
            actions = load_json_file(file_path)
            file_count += 1

            for action in actions:
                # Track by GUID
                guid = action.get('guid')
                if guid:
                    guid_to_files[guid].append((file_path, action))

                # Create content signature for duplicate detection
                signature = create_content_signature(action)
                content_signatures[signature].append((file_path, action))

                # Add to all actions with file info
                action_with_file = action.copy()
                action_with_file['_source_file'] = file_path
                all_actions.append(action_with_file)

    # Find duplicates
    guid_duplicates = {guid: files for guid, files in guid_to_files.items() if len(files) > 1}
    content_duplicates = {sig: files for sig, files in content_signatures.items() if len(files) > 1}

    return {
        'total_files': file_count,
        'total_actions': len(all_actions),
        'guid_duplicates': guid_duplicates,
        'content_duplicates': content_duplicates,
        'all_actions': all_actions
    }


def create_content_signature(action: Dict) -> str:
    """Create a signature for an action to detect content duplicates."""
    action_type = action.get('action', '')

    if action_type == 'create':
        title = action.get('title', '').strip().lower()
        return f"create:{title}"
    elif action_type in ['update', 'comment', 'close']:
        number = action.get('number', '')
        body = action.get('body', '').strip()[:50]  # First 50 chars
        return f"{action_type}:{number}:{body}"

    return f"{action_type}:unknown"


def print_analysis_report(analysis: Dict):
    """Print a detailed analysis report."""
    print("\n" + "="*60)
    print("📊 ISSUE UPDATE ANALYSIS REPORT")
    print("="*60)

    print(f"\n📈 Summary:")
    print(f"  • Total files: {analysis['total_files']}")
    print(f"  • Total actions: {analysis['total_actions']}")
    print(f"  • GUID duplicates: {len(analysis['guid_duplicates'])}")
    print(f"  • Content duplicates: {len(analysis['content_duplicates'])}")

    # GUID Duplicates
    if analysis['guid_duplicates']:
        print(f"\n🔍 GUID Duplicates ({len(analysis['guid_duplicates'])}):")
        for guid, files in analysis['guid_duplicates'].items():
            print(f"\n  GUID: {guid}")
            for file_path, action in files:
                filename = os.path.basename(file_path)
                directory = "processed" if "processed" in file_path else "pending"
                print(f"    • {filename} ({directory}) - {action.get('action')} {action.get('number', action.get('title', ''))}")

    # Content Duplicates (excluding GUID duplicates)
    content_only_duplicates = {}
    for sig, files in analysis['content_duplicates'].items():
        # Check if any of these files share GUIDs
        guids = set()
        for file_path, action in files:
            guid = action.get('guid')
            if guid:
                guids.add(guid)

        # If all files have different GUIDs (or no GUIDs), it's a content duplicate
        if len(guids) == len(files) or len(guids) == 0:
            content_only_duplicates[sig] = files

    if content_only_duplicates:
        print(f"\n📝 Content Duplicates (different GUIDs) ({len(content_only_duplicates)}):")
        for sig, files in content_only_duplicates.items():
            print(f"\n  Signature: {sig}")
            for file_path, action in files:
                filename = os.path.basename(file_path)
                directory = "processed" if "processed" in file_path else "pending"
                guid = action.get('guid', 'no-guid')
                print(f"    • {filename} ({directory}) - GUID: {guid}")


def main():
    """Main analysis function."""
    base_dir = os.getcwd()

    if not os.path.exists(os.path.join(base_dir, ".github")):
        print("❌ Not in a repository with .github directory")
        return

    print("🔍 Analyzing issue update files for duplicates...")
    analysis = analyze_duplicates(base_dir)
    print_analysis_report(analysis)

    # Recommendations
    print(f"\n💡 Recommendations:")
    if analysis['guid_duplicates']:
        print("  • Remove duplicate GUID files (keep processed versions)")
    if len(analysis['content_duplicates']) > len(analysis['guid_duplicates']):
        print("  • Review content duplicates with different GUIDs")
    print("  • Run a migration script to consolidate and deduplicate")


if __name__ == "__main__":
    main()
