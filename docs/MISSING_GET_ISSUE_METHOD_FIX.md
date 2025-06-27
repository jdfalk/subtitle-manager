# Missing get_issue Method Fix

## Issue Description

During the unified issue management workflow improvements, we discovered that the `GitHubAPI` class in `issue_manager.py` was missing a `get_issue` method that was being called in three places:

1. Line 1014: In `update_issue_action` - to get issue details after successful update
2. Line 1081: In `close_issue_action` - to get issue details after successful closure
3. Line 1355: In `delete_issue_action` - to get issue details before deletion for summary

## Root Cause

The `get_issue` method was referenced in the code but never implemented in the `GitHubAPI` class, causing runtime errors when the script attempted to update, close, or delete issues.

## Solution Implemented

Added the missing `get_issue` method to the `GitHubAPI` class with the following features:

### Method Signature

```python
def get_issue(self, issue_number: int) -> Optional[Dict[str, Any]]
```

### Functionality

- Fetches a single issue by its number using the GitHub REST API
- Returns issue data as a dictionary on success
- Returns `None` if the issue is not found (404) or if an error occurs
- Includes proper error handling and logging for different HTTP status codes
- Uses appropriate timeout and exception handling for network requests

### Implementation Details

- **Endpoint**: `GET /repos/{owner}/{repo}/issues/{issue_number}`
- **Success Response**: Returns full issue data dictionary (status 200)
- **Not Found Response**: Logs error and returns `None` (status 404)
- **Error Response**: Logs detailed error information and returns `None` (other status codes)
- **Network Errors**: Logs network exceptions and returns `None`

## Testing

Verified the implementation by:

1. Confirming script runs without syntax errors (`python issue_manager.py --help`)
2. Checking that the method is properly accessible in the GitHubAPI class
3. Validating the method signature matches expected usage patterns

## Files Modified

- `/Users/jdfalk/repos/github.com/jdfalk/ghcommon/scripts/issue_manager.py`
  - Added `get_issue` method to `GitHubAPI` class (lines added before `get_codeql_alerts` method)

## Impact

This fix resolves the runtime errors that were preventing the issue management workflow from properly:

- Updating issues and generating accurate summaries with titles and URLs
- Closing issues and recording them in operation summaries
- Deleting issues and capturing their details before removal

The workflow should now run completely without the missing method errors that were previously occurring.

## Next Steps

With this fix in place, the unified issue management workflow should be fully functional. The next recommended action is to:

1. Test the complete workflow end-to-end in a repository environment
2. Verify that issue updates, closures, and deletions work properly
3. Confirm that operation summaries include accurate titles and URLs
4. Validate that the enhanced PR management (avoiding duplicates) works as expected

## Related Documentation

- [UNIFIED_ISSUE_MANAGEMENT_SUMMARY_ENHANCEMENT.md](./UNIFIED_ISSUE_MANAGEMENT_SUMMARY_ENHANCEMENT.md) - Overall workflow improvements
- [DUPLICATE_PR_FIX_SUMMARY.md](./DUPLICATE_PR_FIX_SUMMARY.md) - Duplicate PR prevention fixes
