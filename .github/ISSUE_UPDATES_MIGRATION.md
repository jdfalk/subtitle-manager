# Issue Updates Migration Plan

## Overview

We're migrating from a single `issue_updates.json` file to a distributed approach using individual UUID-named files in the `.github/issue-updates/` directory. This eliminates merge conflicts and allows parallel development.

## Current State (Legacy Format)

```json
{
  "create": [
    {
      "title": "Issue title",
      "body": "Issue body",
      "labels": ["label1"],
      "guid": "unique-id"
    }
  ],
  "update": [...],
  "comment": [...],
  "close": [...],
  "delete": [...]
}
```

## New Format (Per-File)

```json
{
  "action": "create",
  "title": "Issue title",
  "body": "Issue body",
  "labels": ["label1"],
  "guid": "unique-id"
}
```

## Migration Steps

### Phase 1: Dual Support ✅
- [x] Create `.github/issue-updates/` directory
- [x] Update workflow to process both formats
- [x] Add examples and documentation
- [x] Create helper script for new format

### Phase 2: Transition (Current)
- [ ] Convert existing items in `issue_updates.json` to individual files
- [ ] Update documentation to recommend new format
- [ ] Train team on new workflow

### Phase 3: Full Migration (Future)
- [ ] Remove legacy `issue_updates.json` support
- [ ] Archive old file
- [ ] Update workflow to only use new format

## Benefits

### ✅ Eliminated Merge Conflicts
- Each issue update is in its own file
- Parallel development without conflicts
- Atomic operations

### ✅ Better Organization
- UUID-based file naming prevents conflicts
- Clear action-based structure
- Easy to track and review changes

### ✅ Improved Workflow
- Simplified JSON structure
- Better git history tracking
- Easier rollback of individual changes

## Usage Examples

### Using the Helper Script

```bash
# Create new issue
./scripts/create-issue-update.sh create "Add WebSocket support" "Implement real-time updates" "enhancement,frontend"

# Update existing issue
./scripts/create-issue-update.sh update 123 "Updated implementation details" "enhancement,completed"

# Add comment
./scripts/create-issue-update.sh comment 123 "Testing completed successfully"

# Close issue
./scripts/create-issue-update.sh close 123 "completed"
```

### Manual File Creation

1. Generate UUID: `uuidgen` or use online generator
2. Create file: `.github/issue-updates/{uuid}.json`
3. Add content following the new format
4. Commit and push

## File Naming Convention

- **Format**: `{uuid}.json`
- **Example**: `bc03b7dc-eba7-4b95-9a90-a0224b274633.json`
- **UUID Generation**:
  - macOS/Linux: `uuidgen`
  - Online: https://www.uuidgenerator.net/
  - Python: `python3 -c "import uuid; print(uuid.uuid4())"`

## Workflow Integration

The unified issue management workflow now:

1. **Scans both locations**:
   - Legacy: `issue_updates.json`
   - New: `.github/issue-updates/*.json`

2. **Processes all files**: Combines all updates into a single operation

3. **Maintains compatibility**: Existing workflows continue to work

4. **Provides feedback**: Clear logging of which files are processed

## Best Practices

### File Organization
- One operation per file
- Use descriptive GUIDs
- Include timestamps in GUIDs
- Follow consistent naming

### Content Guidelines
- Keep actions atomic
- Use clear, descriptive titles
- Include comprehensive body text
- Apply appropriate labels

### Development Workflow
- Create files in feature branches
- Use meaningful commit messages
- Review files before merging
- Test with dry-run mode first

## Troubleshooting

### Common Issues

**Duplicate GUIDs**: Ensure each file has a unique GUID to prevent conflicts

**Invalid JSON**: Validate JSON syntax before committing

**Missing Actions**: Ensure each file includes a valid action field

**File Naming**: Use proper UUID format for file names

### Validation

```bash
# Check JSON syntax
jq empty .github/issue-updates/*.json

# List all files
ls -la .github/issue-updates/

# Check for duplicate GUIDs
grep -h '"guid"' .github/issue-updates/*.json | sort | uniq -c | grep -v "1 "
```

## Timeline

- **Now**: Dual format support active
- **Week 1**: Team training and adoption
- **Week 2-3**: Gradual migration of existing items
- **Month 1**: Full adoption of new format
- **Month 2**: Remove legacy support

## Support

For questions or issues:
1. Check this migration guide
2. Review example files in `.github/issue-updates/`
3. Use the helper script for new files
4. Test with dry-run mode first
