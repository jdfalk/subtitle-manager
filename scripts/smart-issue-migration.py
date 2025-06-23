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
from datetime import datetime
from typing import Dict, List, Optional


class SmartMigrator:
    """Intelligent issue update migrator with duplicate prevention."""

    def __init__(self, base_dir: str):
        self.base_dir = base_dir
        self.pending_dir = os.path.join(base_dir, ".github/issue-updates")
        self.issue_updates_dir = self.pending_dir  # Alias for compatibility
        self.processed_dir = os.path.join(self.pending_dir, "processed")

        # Ensure directories exist
        os.makedirs(self.pending_dir, exist_ok=True)
        os.makedirs(self.processed_dir, exist_ok=True)

        # Load existing actions for duplicate detection
        self.existing_actions = self._load_existing_actions()
        print(
            f"📚 Loaded {len(self.existing_actions)} existing actions for duplicate detection"
        )

    def _load_existing_actions(self) -> List[Dict]:
        """Load all existing actions from processed and pending directories."""
        all_actions = []

        for directory in [self.processed_dir, self.pending_dir]:
            if not os.path.exists(directory):
                continue

            for filename in os.listdir(directory):
                if not filename.endswith(".json") or filename == "README.json":
                    continue

                file_path = os.path.join(directory, filename)
                try:
                    with open(file_path, "r", encoding="utf-8") as f:
                        data = json.load(f)

                    actions = data if isinstance(data, list) else [data]
                    for action in actions:
                        action["_source_file"] = file_path
                        all_actions.append(action)

                except Exception as e:
                    print(f"⚠️  Error reading {file_path}: {e}")

        return all_actions

    def _create_content_signature(self, action: Dict) -> str:
        """Create a unique signature for action content."""
        action_type = action.get("action", "")

        if action_type == "create":
            title = action.get("title", "").strip().lower()
            # Normalize whitespace and punctuation for better matching
            title = " ".join(title.split())
            return f"create:{title}"
        elif action_type in ["update", "comment", "close"]:
            number = action.get("number", "")
            body = action.get("body", "").strip()[:100]  # First 100 chars
            return f"{action_type}:{number}:{body}"

        return f"{action_type}:unknown"

    def _find_duplicate(self, action: Dict) -> Optional[Dict]:
        """Find if an action already exists (by GUID or content)."""
        guid = action.get("guid")

        # First check by GUID (exact match)
        if guid:
            for existing in self.existing_actions:
                if existing.get("guid") == guid:
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
            with open(json_file, "r", encoding="utf-8") as f:
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
            "errors": [],
        }

        for action in actions:
            try:
                result = self._process_action(action, force)

                if result["status"] == "created":
                    summary["created_files"].append(result["filename"])
                elif result["status"] == "duplicate":
                    summary["skipped_duplicates"].append(
                        {
                            "action": self._describe_action(action),
                            "existing_file": os.path.basename(
                                result["duplicate"]["_source_file"]
                            ),
                            "reason": result["reason"],
                        }
                    )
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
        sections = content.split("<<<<<<< HEAD")

        for section in sections:
            # Extract content between markers
            if "=======" in section:
                parts = section.split("=======")
                if len(parts) >= 2:
                    # Try both HEAD and incoming sections
                    for part in parts:
                        cleaned = (
                            part.split(">>>>>>> ")[0] if ">>>>>>> " in part else part
                        )
                        try:
                            data = json.loads(cleaned.strip())
                            section_actions = self._flatten_json_data(data)
                            actions.extend(section_actions)
                            print(
                                f"✅ Extracted {len(section_actions)} actions from merge section"
                            )
                        except json.JSONDecodeError:
                            continue
            else:
                # Regular section without conflict markers
                try:
                    data = json.loads(section.strip())
                    section_actions = self._flatten_json_data(data)
                    actions.extend(section_actions)
                    print(
                        f"✅ Extracted {len(section_actions)} actions from regular section"
                    )
                except json.JSONDecodeError:
                    continue

        # Remove duplicates by GUID
        seen_guids = set()
        unique_actions = []
        for action in actions:
            guid = action.get("guid")
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
                    if "action" in item:
                        actions.append(item)
                    else:
                        # Check if it's grouped format
                        for action_type in [
                            "create",
                            "update",
                            "comment",
                            "close",
                            "delete",
                        ]:
                            if action_type in item and isinstance(
                                item[action_type], list
                            ):
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
                "reason": "GUID match"
                if action.get("guid") == duplicate.get("guid")
                else "Content match",
            }

        # Generate filename and create file
        try:
            filename = f"{uuid.uuid4()}.json"
            file_path = os.path.join(self.pending_dir, filename)

            with open(file_path, "w", encoding="utf-8") as f:
                json.dump(action, f, indent=2)

            return {"status": "created", "filename": filename, "file_path": file_path}

        except Exception as e:
            return {"status": "error", "error": str(e)}

    def _describe_action(self, action: Dict) -> str:
        """Create a human-readable description of an action."""
        action_type = action.get("action", "unknown")

        if action_type == "create":
            title = action.get("title", "untitled")
            return f"create: {title[:50]}..."
        elif action_type in ["update", "comment", "close"]:
            number = action.get("number", "unknown")
            return f"{action_type}: issue #{number}"

        return f"{action_type}: unknown"

    def clean_duplicates(self, dry_run: bool = True) -> Dict:
        """Clean up duplicate files in pending directory."""
        print(f"🧹 Cleaning duplicates (dry_run={dry_run})")

        # Group pending files by content signature
        processed_signatures = set()

        # Get all processed signatures
        for action in self.existing_actions:
            if "processed" in action.get("_source_file", ""):
                signature = self._create_content_signature(action)
                processed_signatures.add(signature)

        # Find pending files that duplicate processed ones
        duplicate_files = []

        for filename in os.listdir(self.pending_dir):
            if not filename.endswith(".json") or filename == "README.json":
                continue

            file_path = os.path.join(self.pending_dir, filename)
            try:
                with open(file_path, "r", encoding="utf-8") as f:
                    data = json.load(f)

                actions = data if isinstance(data, list) else [data]
                for action in actions:
                    signature = self._create_content_signature(action)
                    if signature in processed_signatures:
                        duplicate_files.append(
                            {
                                "file": file_path,
                                "filename": filename,
                                "action": self._describe_action(action),
                                "signature": signature,
                            }
                        )
                        break

            except Exception as e:
                print(f"⚠️  Error reading {file_path}: {e}")

        summary = {"duplicate_files_found": len(duplicate_files), "removed_files": []}

        if duplicate_files:
            print(f"🔍 Found {len(duplicate_files)} duplicate files:")
            for dup in duplicate_files:
                print(f"  • {dup['filename']}: {dup['action']}")

                if not dry_run:
                    try:
                        os.remove(dup["file"])
                        summary["removed_files"].append(dup["filename"])
                        print("    ✅ Removed")
                    except Exception as e:
                        print(f"    ❌ Failed to remove: {e}")
                else:
                    print("    🔍 Would remove (dry run)")

        return summary

    def clear_source_file(
        self, json_file: str, backup: bool = True, dry_run: bool = False
    ) -> Dict:
        """
        Clear the source JSON file after migration, optionally creating a backup.

        Args:
            json_file: Path to the JSON file to clear
            backup: Create a backup before clearing
            dry_run: Show what would be done without making changes

        Returns:
            Summary of the clearing operation
        """
        print(f"🗑️  Clearing source file: {json_file} (dry_run={dry_run})")

        if not os.path.exists(json_file):
            return {"error": f"File {json_file} not found"}

        summary = {
            "source_file": json_file,
            "backup_created": False,
            "backup_file": None,
            "file_cleared": False,
            "actions_removed": 0,
            "error": None,
        }

        try:
            # Read current content to count actions
            with open(json_file, "r", encoding="utf-8") as f:
                content = f.read()

            # Count existing actions
            try:
                data = json.loads(content)
                actions = self._flatten_json_data(data)
                summary["actions_removed"] = len(actions)
            except json.JSONDecodeError:
                # Try to extract from merge conflict content
                actions = self._extract_actions_from_content(content)
                summary["actions_removed"] = len(actions)

            if dry_run:
                print(
                    f"🔍 Would remove {summary['actions_removed']} actions from {json_file}"
                )
                if backup:
                    backup_file = (
                        f"{json_file}.backup.{datetime.now().strftime('%Y%m%d_%H%M%S')}"
                    )
                    summary["backup_file"] = backup_file
                    print(f"🔍 Would create backup: {backup_file}")
                print("🔍 Would reset file to empty structure")
                return summary

            # Create backup if requested
            if backup:
                backup_file = (
                    f"{json_file}.backup.{datetime.now().strftime('%Y%m%d_%H%M%S')}"
                )
                with open(backup_file, "w", encoding="utf-8") as f:
                    f.write(content)
                summary["backup_created"] = True
                summary["backup_file"] = backup_file
                print(f"💾 Backup created: {backup_file}")

            # Clear the file with empty structure
            empty_structure = {
                "create": [],
                "update": [],
                "comment": [],
                "close": [],
                "delete": [],
            }

            with open(json_file, "w", encoding="utf-8") as f:
                json.dump(empty_structure, f, indent=2)

            summary["file_cleared"] = True
            print(f"✅ Cleared {summary['actions_removed']} actions from {json_file}")

        except Exception as e:
            summary["error"] = str(e)
            print(f"❌ Error clearing file: {e}")

        return summary

    def reset_workflow(
        self, json_file: str, dry_run: bool = False, keep_backup: bool = True
    ) -> Dict:
        """
        Complete workflow reset: migrate existing data then clear source file.

        Args:
            json_file: Path to the JSON file to process
            dry_run: Show what would be done without making changes
            keep_backup: Keep backup of original file

        Returns:
            Combined summary of migration and clearing
        """
        print(f"🔄 Starting workflow reset for {json_file}")
        print("=" * 60)

        summary = {
            "migration": {},
            "clearing": {},
            "total_actions_processed": 0,
            "success": False,
        }

        # Step 1: Migrate existing data
        print("📤 Step 1: Migrating existing data...")
        migration_summary = self.migrate_from_json(json_file, force=False)
        summary["migration"] = migration_summary

        if "error" in migration_summary:
            print(f"❌ Migration failed: {migration_summary['error']}")
            return summary

        summary["total_actions_processed"] = migration_summary.get("total_actions", 0)

        print(
            f"✅ Migration completed: {migration_summary.get('total_actions', 0)} actions processed"
        )
        print(f"   • Created: {len(migration_summary.get('created_files', []))}")
        print(f"   • Skipped: {len(migration_summary.get('skipped_duplicates', []))}")
        print(f"   • Errors: {len(migration_summary.get('errors', []))}")

        # Step 2: Clear source file (only if migration was successful)
        if migration_summary.get("total_actions", 0) > 0:
            print("\n🗑️  Step 2: Clearing source file...")
            clearing_summary = self.clear_source_file(
                json_file, backup=keep_backup, dry_run=dry_run
            )
            summary["clearing"] = clearing_summary

            if clearing_summary.get("error"):
                print(f"❌ Clearing failed: {clearing_summary['error']}")
                return summary

            if not dry_run:
                print("✅ Source file cleared and reset to empty structure")
                if clearing_summary.get("backup_created"):
                    print(
                        f"💾 Original data backed up to: {clearing_summary.get('backup_file')}"
                    )
        else:
            print("\n⏭️  Step 2: Skipping file clearing (no actions to migrate)")
            summary["clearing"] = {"skipped": True, "reason": "No actions to migrate"}

        summary["success"] = True
        return summary

    def analyze_guid_duplicates(self) -> Dict:
        """Analyze GUID duplicates across all issue update files."""
        guid_map = {}
        duplicates = []

        # Scan all files and track GUIDs
        all_files = []

        # Check main issue_updates.json if it exists
        issue_updates_file = os.path.join(self.base_dir, "issue_updates.json")
        if os.path.exists(issue_updates_file):
            all_files.append(("source", issue_updates_file))

        # Check distributed files
        for directory_type in ["pending", "processed"]:
            directory = (
                os.path.join(self.pending_dir, directory_type)
                if directory_type == "processed"
                else self.pending_dir
            )
            if os.path.exists(directory):
                for filename in os.listdir(directory):
                    if filename.endswith(".json") and filename != "README.json":
                        all_files.append(
                            (directory_type, os.path.join(directory, filename))
                        )

        # Analyze each file
        for file_type, file_path in all_files:
            try:
                with open(file_path, "r", encoding="utf-8") as f:
                    content = f.read()

                # Extract actions, handling merge conflicts
                actions = self._extract_actions_from_content(content)

                for action in actions:
                    guid = action.get("guid")
                    if not guid:
                        continue

                    if guid not in guid_map:
                        guid_map[guid] = []

                    guid_map[guid].append(
                        {
                            "file_path": file_path,
                            "file_type": file_type,
                            "action": action,
                            "signature": self._create_content_signature(action),
                        }
                    )

            except Exception as e:
                print(f"⚠️  Error analyzing {file_path}: {e}")

        # Find duplicates
        for guid, occurrences in guid_map.items():
            if len(occurrences) > 1:
                duplicates.append(
                    {
                        "guid": guid,
                        "count": len(occurrences),
                        "occurrences": occurrences,
                    }
                )

        return {
            "total_guids": len(guid_map),
            "duplicate_guids": len(duplicates),
            "duplicates": duplicates,
            "guid_map": guid_map,
        }

    def fix_guid_duplicates(
        self, dry_run: bool = True, strategy: str = "regenerate"
    ) -> Dict:
        """
        Fix GUID duplicates using specified strategy.

        Args:
            dry_run: If True, only show what would be changed
            strategy: "regenerate" (new GUIDs) or "merge" (combine duplicates)
        """
        analysis = self.analyze_guid_duplicates()

        if not analysis["duplicates"]:
            return {"message": "No GUID duplicates found", "changes": 0}

        changes_made = []

        for duplicate_info in analysis["duplicates"]:
            guid = duplicate_info["guid"]
            occurrences = duplicate_info["occurrences"]

            if strategy == "regenerate":
                # Keep the first occurrence, regenerate GUIDs for others
                for i, occurrence in enumerate(occurrences[1:], 1):
                    new_guid = str(uuid.uuid4())

                    change = {
                        "file": occurrence["file_path"],
                        "old_guid": guid,
                        "new_guid": new_guid,
                        "action_type": occurrence["action"].get("action", "unknown"),
                    }

                    if not dry_run:
                        # Update the file
                        self._update_guid_in_file(
                            occurrence["file_path"], guid, new_guid
                        )

                    changes_made.append(change)

            elif strategy == "merge":
                # More complex: merge duplicate actions
                # For now, we'll just report what would be merged
                change = {
                    "file": "multiple",
                    "action": "merge_duplicates",
                    "guid": guid,
                    "files": [occ["file_path"] for occ in occurrences],
                    "note": "Merge strategy not yet implemented",
                }
                changes_made.append(change)

        return {
            "duplicates_found": len(analysis["duplicates"]),
            "changes_made": len(changes_made),
            "changes": changes_made,
            "dry_run": dry_run,
        }

    def _update_guid_in_file(
        self, file_path: str, old_guid: str, new_guid: str
    ) -> bool:
        """Update a GUID in a specific file."""
        try:
            with open(file_path, "r", encoding="utf-8") as f:
                content = f.read()

            # Replace the GUID
            updated_content = content.replace(f'"{old_guid}"', f'"{new_guid}"')

            # Write back
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(updated_content)

            return True

        except Exception as e:
            print(f"❌ Error updating GUID in {file_path}: {e}")
            return False

    def generate_unique_guid(self, exclude_guids: set = None) -> str:
        """Generate a GUID that's unique across all existing files."""
        if exclude_guids is None:
            analysis = self.analyze_guid_duplicates()
            exclude_guids = set(analysis["guid_map"].keys())

        while True:
            new_guid = str(uuid.uuid4())
            if new_guid not in exclude_guids:
                return new_guid

    def validate_guid_uniqueness(self) -> Dict:
        """Validate that all GUIDs are unique across the project."""
        analysis = self.analyze_guid_duplicates()

        return {
            "is_valid": analysis["duplicate_guids"] == 0,
            "total_guids": analysis["total_guids"],
            "duplicate_count": analysis["duplicate_guids"],
            "details": analysis["duplicates"] if analysis["duplicates"] else None,
        }

    def migrate_to_dual_guid_format(self) -> Dict:
        """
        Migrate all issue update files to dual-GUID format.
        Preserves old GUID as 'legacy_guid' and adds new unique 'guid'.

        Returns:
            dict: Migration results with counts and any errors
        """
        results = {
            "files_processed": 0,
            "files_migrated": 0,
            "files_already_migrated": 0,
            "items_migrated": 0,
            "errors": [],
        }

        print("🔄 Migrating to dual-GUID format...")

        # Process main issue_updates.json
        main_file = os.path.join(self.base_dir, "issue_updates.json")
        if os.path.exists(main_file):
            try:
                result = self._migrate_file_to_dual_guid(main_file)
                results["files_processed"] += 1
                if result["migrated"]:
                    results["files_migrated"] += 1
                    results["items_migrated"] += result["items_migrated"]
                    print(f"✅ Migrated {main_file} ({result['items_migrated']} items)")
                else:
                    results["files_already_migrated"] += 1
                    print(f"⏭️  {main_file} already in dual-GUID format")
            except Exception as e:
                error_msg = f"Error processing {main_file}: {str(e)}"
                results["errors"].append(error_msg)
                print(f"❌ {error_msg}")

        # Process distributed files
        if os.path.exists(self.issue_updates_dir):
            for filename in os.listdir(self.issue_updates_dir):
                if filename.endswith(".json") and filename != "README.md":
                    filepath = os.path.join(self.issue_updates_dir, filename)
                    try:
                        result = self._migrate_file_to_dual_guid(filepath)
                        results["files_processed"] += 1
                        if result["migrated"]:
                            results["files_migrated"] += 1
                            results["items_migrated"] += result["items_migrated"]
                            print(
                                f"✅ Migrated {filename} ({result['items_migrated']} items)"
                            )
                        else:
                            results["files_already_migrated"] += 1
                    except Exception as e:
                        error_msg = f"Error processing {filepath}: {str(e)}"
                        results["errors"].append(error_msg)
                        print(f"❌ {error_msg}")

        # Process distributed processed files
        if os.path.exists(self.processed_dir):
            for filename in os.listdir(self.processed_dir):
                if filename.endswith(".json") and filename != "README.md":
                    filepath = os.path.join(self.processed_dir, filename)
                    try:
                        result = self._migrate_file_to_dual_guid(filepath)
                        results["files_processed"] += 1
                        if result["migrated"]:
                            results["files_migrated"] += 1
                            results["items_migrated"] += result["items_migrated"]
                            print(
                                f"✅ Migrated processed/{filename} ({result['items_migrated']} items)"
                            )
                        else:
                            results["files_already_migrated"] += 1
                    except Exception as e:
                        error_msg = f"Error processing {filepath}: {str(e)}"
                        results["errors"].append(error_msg)
                        print(f"❌ {error_msg}")

        return results

    def _migrate_file_to_dual_guid(self, filepath: str) -> Dict:
        """
        Migrate a single issue update file to dual-GUID format.

        Args:
            filepath: Path to the file to migrate

        Returns:
            dict: Migration result for this file
        """
        result = {"migrated": False, "items_migrated": 0}

        with open(filepath, "r", encoding="utf-8") as f:
            data = json.load(f)

        needs_migration = False

        # Handle both single items and arrays
        if isinstance(data, dict):
            if "action" in data:  # Single distributed file
                if self._migrate_item_to_dual_guid(data):
                    needs_migration = True
                    result["items_migrated"] += 1
            else:  # Main issue_updates.json format
                for action_type in ["create", "update", "delete"]:
                    if action_type in data and isinstance(data[action_type], list):
                        for item in data[action_type]:
                            if self._migrate_item_to_dual_guid(item):
                                needs_migration = True
                                result["items_migrated"] += 1

        if needs_migration:
            with open(filepath, "w", encoding="utf-8") as f:
                json.dump(data, f, indent=2, ensure_ascii=False)
            result["migrated"] = True

        return result

    def _migrate_item_to_dual_guid(self, item: Dict) -> bool:
        """
        Migrate a single issue item to dual-GUID format.

        Args:
            item: Issue item dictionary

        Returns:
            bool: True if item was migrated, False if already in new format
        """
        # Check if already migrated (has both 'guid' and 'legacy_guid')
        if "legacy_guid" in item:
            return False

        # If has old-style GUID, migrate it
        if "guid" in item:
            old_guid = item["guid"]

            # Generate new unique GUID
            new_guid = self.generate_unique_guid()

            # Update the item
            item["legacy_guid"] = old_guid
            item["guid"] = new_guid

            return True

        # If no GUID at all, add one (no legacy_guid since there wasn't one before)
        if "guid" not in item:
            new_guid = self.generate_unique_guid()
            item["guid"] = new_guid
            return True

        return False

    def validate_dual_guid_format(self) -> Dict:
        """
        Validate that all files are in the correct dual-GUID format.

        Returns:
            dict: Validation results
        """
        results = {
            "valid": True,
            "files_checked": 0,
            "files_with_legacy_guids": 0,
            "files_missing_guids": 0,
            "items_checked": 0,
            "items_with_legacy_guids": 0,
            "items_missing_guids": 0,
            "errors": [],
        }

        # Check main issue_updates.json
        main_file = os.path.join(self.base_dir, "issue_updates.json")
        if os.path.exists(main_file):
            try:
                file_result = self._validate_file_dual_guid_format(main_file)
                results["files_checked"] += 1
                results["items_checked"] += file_result["items_checked"]

                if file_result["has_legacy_guids"]:
                    results["files_with_legacy_guids"] += 1
                    results["items_with_legacy_guids"] += file_result[
                        "items_with_legacy_guids"
                    ]

                if file_result["missing_guids"]:
                    results["files_missing_guids"] += 1
                    results["items_missing_guids"] += file_result["items_missing_guids"]
                    results["valid"] = False

            except Exception as e:
                error_msg = f"Error validating {main_file}: {str(e)}"
                results["errors"].append(error_msg)
                results["valid"] = False

        # Check distributed files
        if os.path.exists(self.issue_updates_dir):
            for filename in os.listdir(self.issue_updates_dir):
                if filename.endswith(".json") and filename != "README.md":
                    filepath = os.path.join(self.issue_updates_dir, filename)
                    try:
                        file_result = self._validate_file_dual_guid_format(filepath)
                        results["files_checked"] += 1
                        results["items_checked"] += file_result["items_checked"]

                        if file_result["has_legacy_guids"]:
                            results["files_with_legacy_guids"] += 1
                            results["items_with_legacy_guids"] += file_result[
                                "items_with_legacy_guids"
                            ]

                        if file_result["missing_guids"]:
                            results["files_missing_guids"] += 1
                            results["items_missing_guids"] += file_result[
                                "items_missing_guids"
                            ]
                            results["valid"] = False

                    except Exception as e:
                        error_msg = f"Error validating {filepath}: {str(e)}"
                        results["errors"].append(error_msg)
                        results["valid"] = False

        return results

    def _validate_file_dual_guid_format(self, filepath: str) -> Dict:
        """Validate dual-GUID format for a single file."""
        result = {
            "items_checked": 0,
            "items_with_legacy_guids": 0,
            "items_missing_guids": 0,
            "has_legacy_guids": False,
            "missing_guids": False,
        }

        with open(filepath, "r", encoding="utf-8") as f:
            data = json.load(f)

        # Handle both single items and arrays
        if isinstance(data, dict):
            if "action" in data:  # Single distributed file
                self._validate_item_dual_guid_format(data, result)
            else:  # Main issue_updates.json format
                for action_type in ["create", "update", "delete"]:
                    if action_type in data and isinstance(data[action_type], list):
                        for item in data[action_type]:
                            self._validate_item_dual_guid_format(item, result)

        result["has_legacy_guids"] = result["items_with_legacy_guids"] > 0
        result["missing_guids"] = result["items_missing_guids"] > 0

        return result

    def _validate_item_dual_guid_format(self, item: Dict, result: Dict) -> None:
        """Validate dual-GUID format for a single item."""
        result["items_checked"] += 1

        if "legacy_guid" in item:
            result["items_with_legacy_guids"] += 1

        if "guid" not in item:
            result["items_missing_guids"] += 1


def generate_unique_guid_for_project(base_dir: str = None) -> str:
    """
    Utility function to generate a unique GUID for the project.
    This can be imported and used by other scripts.
    """
    if base_dir is None:
        base_dir = os.getcwd()

    migrator = SmartMigrator(base_dir)
    return migrator.generate_unique_guid()


def validate_guid_uniqueness_in_project(base_dir: str = None) -> bool:
    """
    Utility function to validate GUID uniqueness across the project.
    Returns True if all GUIDs are unique, False otherwise.
    """
    if base_dir is None:
        base_dir = os.getcwd()

    migrator = SmartMigrator(base_dir)
    validation = migrator.validate_guid_uniqueness()
    return validation["is_valid"]


def main():
    """Main function with CLI interface."""
    import argparse

    parser = argparse.ArgumentParser(description="Smart issue update migration tool")
    parser.add_argument(
        "command",
        choices=[
            "migrate",
            "clean",
            "analyze",
            "clear",
            "reset",
            "analyze-guids",
            "fix-guids",
            "validate-guids",
            "migrate-dual-guid",
            "validate-dual-guid",
        ],
        help="Command to execute",
    )
    parser.add_argument(
        "--file",
        "-f",
        default="issue_updates.json",
        help="JSON file to migrate (for migrate command)",
    )
    parser.add_argument(
        "--force", action="store_true", help="Force operation (for migrate command)"
    )
    parser.add_argument(
        "--base-dir",
        default=".",
        help="Base directory for operations (default: current directory)",
    )

    args = parser.parse_args()

    migrator = SmartMigrator(args.base_dir)

    if args.command == "migrate":
        print(f"📦 Starting migration from {args.file}")
        result = migrator.migrate_from_json(args.file, args.force)

        if "error" in result:
            print(f"❌ Migration failed: {result['error']}")
            exit(1)

        print(f"\n📊 Migration Summary:")
        print(f"   📋 Total actions: {result['total_actions']}")
        print(f"   ✅ Files created: {len(result['created_files'])}")
        print(f"   ⏭️  Skipped duplicates: {len(result['skipped_duplicates'])}")
        print(f"   ❌ Errors: {len(result['errors'])}")

        if result["created_files"]:
            print(f"\n📁 Created files:")
            for file in result["created_files"]:
                print(f"   • {file}")

        if result["skipped_duplicates"]:
            print(f"\n⏭️  Skipped duplicates:")
            for dup in result["skipped_duplicates"]:
                print(f"   • {dup['action']} (duplicate of {dup['existing_file']})")

        if result["errors"]:
            print(f"\n❌ Errors:")
            for error in result["errors"]:
                print(f"   • {error}")

    elif args.command == "clean":
        print("🧹 Cleaning duplicate files...")
        result = migrator.clean_duplicates()

        print(f"\n📊 Cleanup Summary:")
        print(f"   🔍 Files checked: {result['files_checked']}")
        print(f"   🗑️  Files removed: {len(result['removed_files'])}")
        print(f"   ❌ Errors: {len(result['errors'])}")

        if result["removed_files"]:
            print(f"\n🗑️  Removed files:")
            for file in result["removed_files"]:
                print(f"   • {file}")

        if result["errors"]:
            print(f"\n❌ Errors:")
            for error in result["errors"]:
                print(f"   • {error}")

    elif args.command == "analyze":
        print("🔍 Analyzing issue update files...")
        result = migrator.analyze_files()

        print(f"\n📊 Analysis Summary:")
        print(f"   📁 Total files: {result['total_files']}")
        print(f"   📋 Total actions: {result['total_actions']}")
        print(f"   🔄 Duplicate groups: {len(result['duplicates'])}")

        action_counts = result["action_counts"]
        print(f"   📝 Actions by type:")
        for action, count in action_counts.items():
            print(f"      • {action}: {count}")

        if result["duplicates"]:
            print(f"\n🔄 Duplicate groups found:")
            for i, group in enumerate(result["duplicates"], 1):
                print(f"   Group {i}:")
                for file_info in group:
                    print(f"      • {file_info['file']}: {file_info['summary']}")

    elif args.command == "clear":
        print("🗑️  Clearing all distributed issue update files...")
        result = migrator.clear_distributed_files()
        print(f"   Removed {result['removed_count']} files")
        if result["errors"]:
            print(f"   Errors: {result['errors']}")

    elif args.command == "reset":
        print("🔄 Resetting to original issue_updates.json...")
        result = migrator.reset_to_original()
        if result["success"]:
            print(f"   ✅ Reset complete. Removed {result['removed_count']} files.")
        else:
            print(f"   ❌ Reset failed: {result['error']}")

    elif args.command == "analyze-guids":
        print("🔍 Analyzing GUID usage...")
        result = migrator.analyze_guid_duplicates()

        print(f"\n📊 GUID Analysis:")
        print(f"   📋 Total GUIDs: {result['total_guids']}")
        print(f"   🔄 Duplicate GUIDs: {result['duplicate_guids']}")
        print(f"   ✅ Unique GUIDs: {result['unique_guids']}")

        if result["duplicates"]:
            print(f"\n🔄 Duplicate GUIDs found:")
            for guid, files in result["duplicates"].items():
                print(f"   GUID: {guid}")
                for file_info in files:
                    print(f"      • {file_info}")

    elif args.command == "fix-guids":
        print("🔧 Fixing duplicate GUIDs...")
        result = migrator.fix_guid_duplicates()

        print(f"\n📊 GUID Fix Summary:")
        print(f"   🔍 Files checked: {result['files_checked']}")
        print(f"   🔧 Files updated: {result['files_updated']}")
        print(f"   🔄 GUIDs fixed: {result['guids_fixed']}")

        if result["errors"]:
            print(f"\n❌ Errors:")
            for error in result["errors"]:
                print(f"   • {error}")

    elif args.command == "validate-guids":
        print("✅ Validating GUID uniqueness...")
        result = migrator.validate_guid_uniqueness()

        if result["is_valid"]:
            print(f"   ✅ All {result['total_guids']} GUIDs are unique!")
        else:
            print(f"   ❌ Found {result['duplicate_count']} duplicate GUIDs")
            if result["details"]:
                print("   Duplicate details:")
                for guid, files in result["details"].items():
                    print(f"      GUID: {guid}")
                    for file_info in files:
                        print(f"         • {file_info}")

    elif args.command == "migrate-dual-guid":
        print("🔄 Migrating to dual-GUID format...")
        result = migrator.migrate_to_dual_guid_format()

        print(f"\n📊 Dual-GUID Migration Summary:")
        print(f"   📁 Files processed: {result['files_processed']}")
        print(f"   ✅ Files migrated: {result['files_migrated']}")
        print(f"   ⏭️  Files already migrated: {result['files_already_migrated']}")
        print(f"   📋 Items migrated: {result['items_migrated']}")

        if result["errors"]:
            print(f"\n❌ Errors:")
            for error in result["errors"]:
                print(f"   • {error}")

    elif args.command == "validate-dual-guid":
        print("✅ Validating dual-GUID format...")
        result = migrator.validate_dual_guid_format()

        print(f"\n📊 Dual-GUID Validation:")
        print(f"   📁 Files checked: {result['files_checked']}")
        print(f"   📋 Items checked: {result['items_checked']}")
        print(f"   🔄 Items with legacy GUIDs: {result['items_with_legacy_guids']}")
        print(f"   ❌ Items missing GUIDs: {result['items_missing_guids']}")

        if result["valid"]:
            print("   ✅ All files are in valid dual-GUID format!")
        else:
            print("   ❌ Some files have missing GUIDs")

        if result["errors"]:
            print(f"\n❌ Errors:")
            for error in result["errors"]:
                print(f"   • {error}")


if __name__ == "__main__":
    main()
