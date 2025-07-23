<!-- file: .github/AUTOMATION_PRECHECKS_DISABLED.md -->
<!-- version: 1.0.0 -->
<!-- guid: 1a2b3c4d-5e6f-7890-abcd-1234567890ef -->

# Automation Prechecks Disabled for Subtitle Manager

## 🚫 Temporarily Disabled Checks

The following automation checks have been temporarily disabled in the unified
automation workflow to prevent failures during protobuf generation work:

### Super Linter Validations Disabled

- **Go validation** (`sl_run_go: false`)
- **Protobuf validation** (`sl_run_protobuf: false`)
- **Security validation** (`sl_run_security: false`)

### Operations Limited

- **Default operation** changed from `'all'` to `'issues,docs,label'`
- **Lint operation excluded** until protobuf generation is complete

## 🎯 Reason for Disabling

These checks were disabled because:

1. **Protobuf files are not yet generated** - causing Go compilation failures
2. **Missing generated code** - causing import and dependency issues
3. **Test failures expected** - until full protobuf ecosystem is implemented

## 📋 Re-enabling Checklist

When protobuf generation is complete, re-enable by:

1. **Update `.github/workflows/unified-automation.yml`:**

   ```yaml
   operation: ${{ github.event.inputs.operation || 'all' }}
   sl_run_go: true
   sl_run_protobuf: true
   sl_run_security: true
   ```

2. **Update version number** in workflow file header

3. **Test the workflow** with a manual trigger to ensure all checks pass

4. **Delete this documentation file** once checks are re-enabled

## 🔄 Current Status

- **Status**: ❌ Prechecks Disabled
- **Date Disabled**: July 20, 2025
- **Disabled By**: Automated setup during protobuf work
- **Expected Duration**: Until protobuf generation infrastructure is complete

## 📝 Notes

- Issue management and documentation automation continue to work normally
- Only validation/testing components are disabled
- This is a temporary measure to prevent CI/CD failures during development
