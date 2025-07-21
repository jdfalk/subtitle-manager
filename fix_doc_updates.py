#!/usr/bin/env python3
# file: fix_doc_updates.py
# version: 1.0.0
# guid: a1b2c3d4-e5f6-7890-abcd-ef1234567890

"""
Fix documentation updates that were moved to processed but not applied correctly.
This script will:
1. Parse all processed JSON files
2. Group updates by target file and mode
3. Apply updates in proper format
4. Clean up duplicates and formatting issues
"""

import json
import os
import re
from collections import defaultdict
from pathlib import Path


def load_processed_updates(processed_dir):
    """Load all processed JSON files and group by target file."""
    updates_by_file = defaultdict(list)

    for json_file in Path(processed_dir).glob("*.json"):
        try:
            with open(json_file, "r", encoding="utf-8") as f:
                update = json.load(f)
                target_file = update.get("file", "unknown")
                updates_by_file[target_file].append(update)
        except Exception as e:
            print(f"Error reading {json_file}: {e}")

    return updates_by_file


def process_changelog_updates(updates):
    """Process changelog updates and group by type."""
    added_items = []
    fixed_items = []
    changed_items = []

    for update in updates:
        content = update.get("content", "")
        if "### Added" in content:
            # Extract the added item
            item = content.replace("### Added\n\n- ", "").strip()
            added_items.append(item)
        elif "### Fixed" in content:
            # Extract the fixed item
            item = content.replace("### Fixed\n\n- ", "").strip()
            fixed_items.append(item)
        elif "### Changed" in content:
            # Extract the changed item
            item = content.replace("### Changed\n\n- ", "").strip()
            changed_items.append(item)

    return added_items, fixed_items, changed_items


def process_todo_updates(updates):
    """Process TODO updates by mode."""
    task_adds = []
    task_completes = []

    for update in updates:
        mode = update.get("mode", "")
        content = update.get("content", "").strip()

        if mode == "task-add":
            task_adds.append(content)
        elif mode == "task-complete":
            task_completes.append(content)

    return task_adds, task_completes


def process_readme_updates(updates):
    """Process README updates."""
    appends = []

    for update in updates:
        mode = update.get("mode", "")
        content = update.get("content", "").strip()

        if mode == "append":
            appends.append(content)

    return appends


def fix_changelog(file_path, updates):
    """Fix CHANGELOG.md with proper formatting."""
    print(f"Processing {len(updates)} changelog updates...")

    added_items, fixed_items, changed_items = process_changelog_updates(updates)

    # Read current changelog
    with open(file_path, "r", encoding="utf-8") as f:
        content = f.read()

    # Find the [Unreleased] section
    unreleased_pattern = r"## \[Unreleased\]\s*(.*?)(?=\n## \[|$)"
    match = re.search(unreleased_pattern, content, re.DOTALL)

    if not match:
        print("Could not find [Unreleased] section in CHANGELOG.md")
        return False

    # Build new unreleased section
    new_section = "## [Unreleased]\n\n"

    if added_items:
        new_section += "### Added\n\n"
        for item in sorted(set(added_items)):  # Remove duplicates and sort
            new_section += f"- {item}\n"
        new_section += "\n"

    if fixed_items:
        new_section += "### Fixed\n\n"
        for item in sorted(set(fixed_items)):  # Remove duplicates and sort
            new_section += f"- {item}\n"
        new_section += "\n"

    if changed_items:
        new_section += "### Changed\n\n"
        for item in sorted(set(changed_items)):  # Remove duplicates and sort
            new_section += f"- {item}\n"
        new_section += "\n"

    # Replace the unreleased section
    new_content = re.sub(
        unreleased_pattern, new_section.rstrip() + "\n", content, flags=re.DOTALL
    )

    # Write back
    with open(file_path, "w", encoding="utf-8") as f:
        f.write(new_content)

    print(
        f"Updated CHANGELOG.md with {len(added_items)} added, {len(fixed_items)} fixed, {len(changed_items)} changed items"
    )
    return True


def fix_todo(file_path, updates):
    """Fix TODO.md with new tasks."""
    print(f"Processing {len(updates)} TODO updates...")

    task_adds, task_completes = process_todo_updates(updates)

    if not task_adds and not task_completes:
        print("No TODO updates to apply")
        return True

    # Read current TODO
    with open(file_path, "r", encoding="utf-8") as f:
        content = f.read()

    # Find a good place to add tasks (look for "### " or "## " sections)
    lines = content.split("\n")
    insert_index = None

    # Look for a section to add tasks under
    for i, line in enumerate(lines):
        if (
            line.strip().startswith("## 🚧 Remaining Work")
            or line.strip().startswith("### 🎯")
            or line.strip().startswith("### 📝")
        ):
            # Find the end of this section
            for j in range(i + 1, len(lines)):
                if lines[j].strip().startswith("## ") or lines[j].strip().startswith(
                    "### "
                ):
                    insert_index = j
                    break
            break

    if insert_index is None:
        # Just append at the end
        insert_index = len(lines)

    # Add new tasks
    if task_adds:
        new_lines = []
        new_lines.append("")
        new_lines.append("### 📝 Recent Updates")
        new_lines.append("")
        for task in sorted(set(task_adds)):
            new_lines.append(f"- [ ] {task}")
        new_lines.append("")

        lines[insert_index:insert_index] = new_lines

    # Mark tasks as complete (find and check them off)
    for task_name in task_completes:
        for i, line in enumerate(lines):
            if task_name.lower() in line.lower() and "- [ ]" in line:
                lines[i] = line.replace("- [ ]", "- [x]")
                break

    # Write back
    with open(file_path, "w", encoding="utf-8") as f:
        f.write("\n".join(lines))

    print(
        f"Updated TODO.md with {len(task_adds)} new tasks, {len(task_completes)} completed tasks"
    )
    return True


def fix_readme(file_path, updates):
    """Fix README.md with appended content."""
    print(f"Processing {len(updates)} README updates...")

    appends = process_readme_updates(updates)

    if not appends:
        print("No README updates to apply")
        return True

    # Read current README
    with open(file_path, "r", encoding="utf-8") as f:
        content = f.read()

    # Add a "Recent Updates" section if it doesn't exist
    if "## Recent Updates" not in content:
        content += "\n\n## Recent Updates\n\n"

    # Add the new items
    for item in sorted(set(appends)):
        if item not in content:  # Avoid duplicates
            content += f"- {item}\n"

    # Write back
    with open(file_path, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"Updated README.md with {len(appends)} new items")
    return True


def main():
    repo_root = "/Users/jdfalk/repos/github.com/jdfalk/subtitle-manager"
    processed_dir = f"{repo_root}/.github/doc-updates/processed"

    print("Loading processed updates...")
    updates_by_file = load_processed_updates(processed_dir)

    print(f"Found updates for {len(updates_by_file)} files:")
    for file_name, updates in updates_by_file.items():
        print(f"  {file_name}: {len(updates)} updates")

    # Process each file type
    for file_name, updates in updates_by_file.items():
        file_path = f"{repo_root}/{file_name}"

        if not os.path.exists(file_path):
            print(f"Warning: Target file {file_path} does not exist, skipping...")
            continue

        print(f"\nProcessing {file_name}...")

        if file_name == "CHANGELOG.md":
            fix_changelog(file_path, updates)
        elif file_name == "TODO.md":
            fix_todo(file_path, updates)
        elif file_name == "README.md":
            fix_readme(file_path, updates)
        else:
            print("  Unknown file type, skipping...")

    print("\nDone!")


if __name__ == "__main__":
    main()
