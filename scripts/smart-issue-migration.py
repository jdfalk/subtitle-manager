#!/usr/bin/env python3
"""
# file: scripts/smart-issue-migration.py
Smart issue update migration script that prevents duplicates.

This script provides intelligent migration from issue_updates.json to distributed
format while detecting and preventing duplicates based on:
1. GUID matching
2. Content similarity
3. Already processed actions

Features:
- Detects existing processed files
- Prevents duplicate creation
- Validates content before migration
- Provides detailed reporting
- Handles merge conflicts intelligently
"""

import json
import os
import uuid
from collections import defaultdict
from datetime import datetime
from typing import Dict, List, Set, Tuple, Optional


class SmartMigrator:
    """Intelligent issue update migrator with duplicate prevention."""

    def __init__(self, base_dir: str):
        self.base_dir = base_dir
        self.pending_dir = os.path.join(base_dir, ".github/issue-updates")
        self.processed_dir = os.path.join(self.pending_dir, "processed")

        # Ensure directories exist
        os.makedirs(self.pending_dir, exist_ok=True)
        os.makedirs(self.processed_dir, exist_ok=True)

        # Load existing actions for duplicate detection
        self.existing_actions = self._load_existing_actions()
        print(f"📚 Loaded {len(self.existing_actions)} existing actions for duplicate detection")

    def _load_existing_actions(self) -> List[Dict]:
        """Load all existing actions from processed and pending directories."""
        all_actions = []

        for directory in [self.processed_dir, self.pending_dir]:
            if not os.path.exists(directory):
                continue

            for filename in os.listdir(directory):
                if not filename.endswith('.json') or filename == 'README.json':
                    continue

                file_path = os.path.join(directory, filename)
                try:
                    with open(file_path, 'r', encoding='utf-8') as f:
                        data = json.load(f)

                    actions = data if isinstance(data, list) else [data]
                    for action in actions:
                        action['_source_file'] = file_path
                        all_actions.append(action)

                except Exception as e:
                    print(f"⚠️  Error reading {file_path}: {e}")

        return all_actions

    def _create_content_signature(self, action: Dict) -> str:
        """Create a unique signature for action content."""
        action_type = action.get('action', '')

        if action_type == 'create':
            title = action.get('title', '').strip().lower()
            # Normalize whitespace and punctuation for better matching
            title = ' '.join(title.split())
            return f"create:{title}"
        elif action_type in ['update', 'comment', 'close']:
            number = action.get('number', '')
            body = action.get('body', '').strip()[:100]  # First 100 chars
            return f"{action_type}:{number}:{body}"

        return f"{action_type}:unknown"

    def _find_duplicate(self, action: Dict) -> Optional[Dict]:
        """Find if an action already exists (by GUID or content)."""
        guid = action.get('guid')

        # First check by GUID (exact match)
        if guid:
            for existing in self.existing_actions:
                if existing.get('guid') == guid:
                    return existing

        # Then check by content signature
        signature = self._create_content_signature(action)
        for existing in self.existing_actions:
            if self._create_content_signature(existing) == signature:
                return existing

        return None

    def migrate_from_json(self, json_file: str, force: bool = False) -> Dict:
        """
        Migrate from issue_updates.json to distributed format.

        Args:
            json_file: Path to the JSON file to migrate
            force: If True, create files even if duplicates exist

        Returns:
            Migration summary dictionary
        """
        if not os.path.exists(json_file):
            return {"error": f"File {json_file} not found"}

        print(f"🔄 Migrating from {json_file}")

        try:
            with open(json_file, 'r', encoding='utf-8') as f:
                content = f.read()

            # Handle merge conflicts by extracting clean JSON sections
            actions = self._extract_actions_from_content(content)

        except Exception as e:
            return {"error": f"Failed to read {json_file}: {e}"}

        if not actions:
            return {"error": "No valid actions found in file"}

        # Process actions
        summary = {
            "total_actions": len(actions),
            "created_files": [],
            "skipped_duplicates": [],
            "errors": []
        }

        for action in actions:
            try:
                result = self._process_action(action, force)

                if result["status"] == "created":
                    summary["created_files"].append(result["filename"])
                elif result["status"] == "duplicate":
                    summary["skipped_duplicates"].append({
                        "action": self._describe_action(action),
                        "existing_file": os.path.basename(result["duplicate"]["_source_file"]),
                        "reason": result["reason"]
                    })
                elif result["status"] == "error":
                    summary["errors"].append(result["error"])

            except Exception as e:
                summary["errors"].append(f"Error processing action: {e}")

        return summary

    def _extract_actions_from_content(self, content: str) -> List[Dict]:
        """Extract actions from file content, handling merge conflicts."""
        actions = []

        # Try to parse as regular JSON first
        try:
            data = json.loads(content)
            actions = self._flatten_json_data(data)
            print(f"✅ Parsed as regular JSON: {len(actions)} actions")
            return actions
        except json.JSONDecodeError:
            pass

        # Handle merge conflict markers
        print("🔧 Handling merge conflicts...")

        # Split by merge conflict markers and try to parse sections
        sections = content.split('<<<<<<< HEAD')

        for section in sections:
            # Extract content between markers
            if '=======' in section:
                parts = section.split('=======')
                if len(parts) >= 2:
                    # Try both HEAD and incoming sections
                    for part in parts:
                        cleaned = part.split('>>>>>>> ')[0] if '>>>>>>> ' in part else part
                        try:
                            data = json.loads(cleaned.strip())
                            section_actions = self._flatten_json_data(data)
                            actions.extend(section_actions)
                            print(f"✅ Extracted {len(section_actions)} actions from merge section")
                        except json.JSONDecodeError:
                            continue
            else:
                # Regular section without conflict markers
                try:
                    data = json.loads(section.strip())
                    section_actions = self._flatten_json_data(data)
                    actions.extend(section_actions)
                    print(f"✅ Extracted {len(section_actions)} actions from regular section")
                except json.JSONDecodeError:
                    continue

        # Remove duplicates by GUID
        seen_guids = set()
        unique_actions = []
        for action in actions:
            guid = action.get('guid')
            if not guid or guid not in seen_guids:
                unique_actions.append(action)
                if guid:
                    seen_guids.add(guid)

        print(f"🎯 Final unique actions: {len(unique_actions)}")
        return unique_actions

    def _flatten_json_data(self, data) -> List[Dict]:
        """Convert JSON data to flat list of actions."""
        if isinstance(data, list):
            # Handle both flat format and nested format
            actions = []
            for item in data:
                if isinstance(item, dict):
                    if 'action' in item:
                        actions.append(item)
                    else:
                        # Check if it's grouped format
                        for action_type in ["create", "update", "comment", "close", "delete"]:
                            if action_type in item and isinstance(item[action_type], list):
                                for action_item in item[action_type]:
                                    action_item["action"] = action_type
                                    actions.append(action_item)
            return actions
        elif isinstance(data, dict):
            # Grouped format
            actions = []
            for action_type in ["create", "update", "comment", "close", "delete"]:
                if action_type in data and isinstance(data[action_type], list):
                    for item in data[action_type]:
                        item["action"] = action_type
                        actions.append(item)
            return actions

        return []

    def _process_action(self, action: Dict, force: bool) -> Dict:
        """Process a single action for migration."""

        # Check for duplicates
        duplicate = self._find_duplicate(action)
        if duplicate and not force:
            return {
                "status": "duplicate",
                "duplicate": duplicate,
                "reason": "GUID match" if action.get('guid') == duplicate.get('guid') else "Content match"
            }

        # Generate filename and create file
        try:
            filename = f"{uuid.uuid4()}.json"
            file_path = os.path.join(self.pending_dir, filename)

            with open(file_path, 'w', encoding='utf-8') as f:
                json.dump(action, f, indent=2)

            return {
                "status": "created",
                "filename": filename,
                "file_path": file_path
            }

        except Exception as e:
            return {
                "status": "error",
                "error": str(e)
            }

    def _describe_action(self, action: Dict) -> str:
        """Create a human-readable description of an action."""
        action_type = action.get('action', 'unknown')

        if action_type == 'create':
            title = action.get('title', 'untitled')
            return f"create: {title[:50]}..."
        elif action_type in ['update', 'comment', 'close']:
            number = action.get('number', 'unknown')
            return f"{action_type}: issue #{number}"

        return f"{action_type}: unknown"

    def clean_duplicates(self, dry_run: bool = True) -> Dict:
        """Clean up duplicate files in pending directory."""
        print(f"🧹 Cleaning duplicates (dry_run={dry_run})")

        # Group pending files by content signature
        pending_files = []
        processed_signatures = set()

        # Get all processed signatures
        for action in self.existing_actions:
            if "processed" in action.get("_source_file", ""):
                signature = self._create_content_signature(action)
                processed_signatures.add(signature)

        # Find pending files that duplicate processed ones
        duplicate_files = []

        for filename in os.listdir(self.pending_dir):
            if not filename.endswith('.json') or filename == 'README.json':
                continue

            file_path = os.path.join(self.pending_dir, filename)
            try:
                with open(file_path, 'r', encoding='utf-8') as f:
                    data = json.load(f)

                actions = data if isinstance(data, list) else [data]
                for action in actions:
                    signature = self._create_content_signature(action)
                    if signature in processed_signatures:
                        duplicate_files.append({
                            "file": file_path,
                            "filename": filename,
                            "action": self._describe_action(action),
                            "signature": signature
                        })
                        break

            except Exception as e:
                print(f"⚠️  Error reading {file_path}: {e}")

        summary = {
            "duplicate_files_found": len(duplicate_files),
            "removed_files": []
        }

        if duplicate_files:
            print(f"🔍 Found {len(duplicate_files)} duplicate files:")
            for dup in duplicate_files:
                print(f"  • {dup['filename']}: {dup['action']}")

                if not dry_run:
                    try:
                        os.remove(dup['file'])
                        summary["removed_files"].append(dup['filename'])
                        print(f"    ✅ Removed")
                    except Exception as e:
                        print(f"    ❌ Failed to remove: {e}")
                else:
                    print(f"    🔍 Would remove (dry run)")

        return summary


def main():
    """Main function with CLI interface."""
    import argparse

    parser = argparse.ArgumentParser(description="Smart issue update migration tool")
    parser.add_argument("command", choices=["migrate", "clean", "analyze"],
                       help="Command to execute")
    parser.add_argument("--file", "-f", default="issue_updates.json",
                       help="JSON file to migrate (for migrate command)")
    parser.add_argument("--force", action="store_true",
                       help="Force migration even if duplicates exist")
    parser.add_argument("--dry-run", action="store_true",
                       help="Show what would be done without making changes")

    args = parser.parse_args()

    migrator = SmartMigrator(os.getcwd())

    if args.command == "migrate":
        summary = migrator.migrate_from_json(args.file, force=args.force)

        print("\n" + "="*50)
        print("📊 MIGRATION SUMMARY")
        print("="*50)

        if "error" in summary:
            print(f"❌ Error: {summary['error']}")
            return

        print(f"📈 Total actions processed: {summary['total_actions']}")
        print(f"✅ Files created: {len(summary['created_files'])}")
        print(f"⏭️  Duplicates skipped: {len(summary['skipped_duplicates'])}")
        print(f"❌ Errors: {len(summary['errors'])}")

        if summary['skipped_duplicates']:
            print(f"\n🔍 Skipped duplicates:")
            for skip in summary['skipped_duplicates']:
                print(f"  • {skip['action']} (reason: {skip['reason']}, existing: {skip['existing_file']})")

        if summary['errors']:
            print(f"\n❌ Errors:")
            for error in summary['errors']:
                print(f"  • {error}")

    elif args.command == "clean":
        summary = migrator.clean_duplicates(dry_run=args.dry_run)

        print("\n" + "="*50)
        print("🧹 CLEANUP SUMMARY")
        print("="*50)

        print(f"🔍 Duplicate files found: {summary['duplicate_files_found']}")

        if not args.dry_run:
            print(f"🗑️  Files removed: {len(summary['removed_files'])}")
            if summary['removed_files']:
                for filename in summary['removed_files']:
                    print(f"  • {filename}")
        else:
            print("🔍 Dry run mode - no files were actually removed")

    elif args.command == "analyze":
        # Run the analysis script
        from analyze_issue_updates import analyze_duplicates, print_analysis_report
        analysis = analyze_duplicates(os.getcwd())
        print_analysis_report(analysis)


if __name__ == "__main__":
    main()
