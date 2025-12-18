<!-- file: WORKFLOW_FAILURES_ANALYSIS.md -->
<!-- version: 1.0.0 -->
<!-- guid: a1b2c3d4-e5f6-7890-abcd-ef1234567890 -->

# Workflow Failures Analysis and Fix Documentation

## Problem Summary

The CI workflows are failing because of mismatched workflow configurations between the local repository and the upstream ghcommon repository.

## Root Cause Analysis

### Issue 1: Reusable Workflow Call Mismatch

**Location:** `.github/workflows/ci.yml` line 47

**Problem:**
```yaml
check-overrides:
  name: Check Commit Overrides
  uses: jdfalk/ghcommon/.github/workflows/commit-override-handler.yml@main
```

The CI workflow attempts to call a reusable workflow from ghcommon, but:

1. **The local `commit-override-handler.yml` has required inputs** that aren't being passed:
   - `override_token` (required)
   - `commit_message` (required)
   - `skip_validation` (optional, default: false)

2. **The ghcommon repository likely doesn't have this workflow** or has a different signature that doesn't match what's expected.

### Issue 2: Local Workflow Definition Conflict

**Location:** `.github/workflows/commit-override-handler.yml`

The local repository has its own `commit-override-handler.yml` workflow that defines a reusable workflow with specific inputs, but this is not being used - instead, the CI workflow is trying to call an external version from ghcommon.

## Solutions

### Option 1: Use Local Workflow (RECOMMENDED)

**Change in `.github/workflows/ci.yml`:**
```yaml
# OLD (line 45-47):
check-overrides:
  name: Check Commit Overrides
  uses: jdfalk/ghcommon/.github/workflows/commit-override-handler.yml@main

# NEW:
check-overrides:
  name: Check Commit Overrides
  uses: ./.github/workflows/commit-override-handler.yml
  with:
    override_token: ${{ secrets.GITHUB_TOKEN }}
    commit_message: ${{ github.event.head_commit.message || '' }}
    skip_validation: false
```

**Pros:**
- Uses the existing local workflow that's already properly configured
- No dependency on external repository
- Immediate fix without waiting for ghcommon updates

**Cons:**
- Diverges from centralized workflow management if that was the intent

### Option 2: Remove the Override Check Entirely

If the commit override functionality isn't actively used:

```yaml
# Remove the check-overrides job entirely and update the needs array
# in other jobs from:
needs: [detect-changes, check-overrides]

# to:
needs: [detect-changes]
```

Also update the `ci-summary` job's needs array.

**Pros:**
- Simplifies the workflow
- Removes a potential point of failure

**Cons:**
- Loses commit override functionality if it's needed

### Option 3: Create Stub Outputs (TEMPORARY FIX)

Add a simple local job that provides the expected outputs without external dependency:

```yaml
check-overrides:
  name: Check Commit Overrides
  runs-on: ubuntu-latest
  outputs:
    skip-tests: ${{ steps.check.outputs.skip-tests }}
    skip-build: ${{ steps.check.outputs.skip-build }}
  steps:
    - name: Check commit message for overrides
      id: check
      run: |
        COMMIT_MSG="${{ github.event.head_commit.message || '' }}"
        
        if [[ "$COMMIT_MSG" =~ \[skip-tests\] ]] || [[ "$COMMIT_MSG" =~ \[SKIP-TESTS\] ]]; then
          echo "skip-tests=true" >> $GITHUB_OUTPUT
        else
          echo "skip-tests=false" >> $GITHUB_OUTPUT
        fi
        
        if [[ "$COMMIT_MSG" =~ \[skip-build\] ]] || [[ "$COMMIT_MSG" =~ \[SKIP-BUILD\] ]]; then
          echo "skip-build=true" >> $GITHUB_OUTPUT
        else
          echo "skip-build=false" >> $GITHUB_OUTPUT
        fi
```

**Pros:**
- Quick fix that maintains expected outputs
- No external dependencies

**Cons:**
- Simplified functionality compared to full commit-override-handler

## What Needs to Change in ghcommon (if Option 1 is not chosen)

If the intention is to use a centralized workflow from ghcommon:

1. **Create the workflow in ghcommon repository** at:
   `jdfalk/ghcommon/.github/workflows/commit-override-handler.yml`

2. **The workflow signature should be:**
   ```yaml
   on:
     workflow_call:
       outputs:
         skip-tests:
           description: "Whether to skip tests"
           value: ${{ jobs.check.outputs.skip-tests }}
         skip-build:
           description: "Whether to skip build"
           value: ${{ jobs.check.outputs.skip-build }}
   ```

3. **Or simplify to not require inputs:**
   - Make it read from the caller's context directly
   - Provide outputs that the CI workflow expects

## Recommended Action

**Implement Option 1 (Use Local Workflow)** because:

1. The local workflow already exists and is properly configured
2. It provides immediate fix without external dependencies
3. The workflow can be tested and validated locally
4. No need to wait for ghcommon repository updates

After fixing, the workflow should:
- ✅ Successfully call the local commit-override-handler
- ✅ Provide the expected outputs (skip-tests, skip-build, etc.)
- ✅ Continue with the rest of the CI pipeline

## Testing the Fix

After implementing the fix:

1. Create a test PR to verify the workflow runs
2. Check that the `check-overrides` job completes successfully
3. Verify that subsequent jobs receive the correct outputs
4. Test with different commit messages:
   - Normal commit: should run all tests
   - Commit with `[skip-tests]`: should skip test jobs
   - Commit with `[skip ci]`: should skip entire CI

## Additional Notes

- The local `commit-override-handler.yml` appears to be a full-featured workflow with token validation, logging, and status creation
- Consider whether this functionality is needed or if a simpler stub is sufficient
- If ghcommon is meant to be a centralized source, coordinate with that repository's maintainers to align workflow definitions
