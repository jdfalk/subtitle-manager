<!-- file: scripts/MIGRATION-NOTICE.md -->
<!-- version: 1.0.1 -->
<!-- guid: 9a8b7c6d-5e4f-3210-9876-fedcba098765 -->

# GitHub Project Scripts Migration Notice

## 🚨 IMPORTANT: Scripts have been migrated

The GitHub project management scripts previously located in this repository have been **migrated to the unified project manager** in the `ghcommon` repository.

### Old Scripts (REMOVED)

- ❌ `create-github-projects.sh` - **REMOVED** (replaced with new wrapper script)
- ❌ `setup-project-workflows.sh` - **REMOVED**
- ❌ `manage-project-structure.sh` - **REMOVED**
- ❌ `github_project_manager.py` - **REMOVED**

### New Unified Script

✅ **Use instead:** `ghcommon/scripts/unified_github_project_manager_v2.py`

The repository now provides a lightweight wrapper script
`scripts/create-github-projects.sh` which verifies your GitHub CLI
authentication scopes before delegating to the unified project manager.

## Usage

From anywhere in your workspace:

```bash
# Run full setup
python3 /path/to/ghcommon/scripts/unified_github_project_manager_v2.py

# Dry run to see what would be created
python3 /path/to/ghcommon/scripts/unified_github_project_manager_v2.py --dry-run

# List all projects
python3 /path/to/ghcommon/scripts/unified_github_project_manager_v2.py --list-projects

# Create labels only
python3 /path/to/ghcommon/scripts/unified_github_project_manager_v2.py --create-labels

# Create milestones only
python3 /path/to/ghcommon/scripts/unified_github_project_manager_v2.py --create-milestones
```

## Benefits of the Unified Script

- 🎯 **Single source of truth** for all project management
- 🔄 **Idempotent operations** - safe to run multiple times
- 🌐 **Cross-repository project support** - projects can span multiple repos
- 📊 **Comprehensive project structure** based on actual documentation analysis
- 🚀 **Advanced features** like auto-add workflow configuration
- 🛠️ **Better error handling** and logging

## Migration Date

July 10, 2025

See the unified script's README for complete documentation:
`ghcommon/scripts/README-unified-project-manager.md`
