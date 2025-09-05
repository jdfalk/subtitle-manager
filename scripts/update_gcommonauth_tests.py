#!/usr/bin/env python3
# file: scripts/update_gcommonauth_tests.py
# version: 1.0.0
# guid: a1b2c3d4-e5f6-7890-abcd-ef1234567890

"""
Script to update gcommonauth test files for gcommon migration.
Converts tests to handle gcommon Session and APIKey types.
"""

import re
import sys
from pathlib import Path


def update_auth_test():
    """Update auth_test.go to handle gcommon types."""
    file_path = Path("pkg/gcommonauth/auth_test.go")
    if not file_path.exists():
        print(f"Error: {file_path} not found")
        return False

    content = file_path.read_text()

    # Fix API key generation and validation patterns
    # Pattern 1: Generate and validate API key
    content = re.sub(
        r'apiKey(\d*), err := GenerateAPIKey\(db, (\d+)\)\s*if err != nil \{\s*t\.Fatalf\("generate API key: %v", err\)\s*\}\s*\n\s*// Test valid API key\s*userID(\d*), err := ValidateAPIKey\(db, apiKey(\d*)\)',
        lambda m: f"""apiKeyObj{m.group(1)}, err := GenerateAPIKey(db, {m.group(2)})
	if err != nil {{
		t.Fatalf("generate API key: %v", err)
	}}
	apiKeyStr{m.group(1)} := apiKeyObj{m.group(1)}.GetId()

	// Test valid API key
	validatedAPIKey{m.group(3)}, err := ValidateAPIKey(db, apiKeyStr{m.group(1)})""",
        content,
    )

    # Pattern 2: Check user ID from APIKey
    content = re.sub(
        r'if userID(\d*) != (\d+) \{\s*t\.Errorf\("expected user ID (\d+), got %d", userID(\d*)\)\s*\}',
        lambda m: f'''if userIdStr := validatedAPIKey{m.group(1)}.GetUserId(); userIdStr != "{m.group(2)}" {{
		t.Errorf("expected user ID '{m.group(3)}', got '%s'", userIdStr)
	}}''',
        content,
    )

    # Pattern 3: Multiple API key generation
    content = re.sub(
        r"apiKey1, err := GenerateAPIKey\(db, 1\)",
        "apiKeyObj1, err := GenerateAPIKey(db, 1)",
        content,
    )
    content = re.sub(
        r"apiKey2, err := GenerateAPIKey\(db, 2\)",
        "apiKeyObj2, err := GenerateAPIKey(db, 2)",
        content,
    )

    # Pattern 4: Multiple API key validation
    content = re.sub(
        r"userID1, err := ValidateAPIKey\(db, apiKey1\)",
        "validatedAPIKey1, err := ValidateAPIKey(db, apiKeyObj1.GetId())",
        content,
    )
    content = re.sub(
        r"userID2, err := ValidateAPIKey\(db, apiKey2\)",
        "validatedAPIKey2, err := ValidateAPIKey(db, apiKeyObj2.GetId())",
        content,
    )

    # Pattern 5: User ID comparisons for multiple keys
    content = re.sub(
        r'if userID1 != 1 \{\s*t\.Errorf\("expected user ID 1, got %d", userID1\)\s*\}',
        """if userIdStr := validatedAPIKey1.GetUserId(); userIdStr != "1" {
		t.Errorf("expected user ID '1', got '%s'", userIdStr)
	}""",
        content,
    )
    content = re.sub(
        r'if userID2 != 2 \{\s*t\.Errorf\("expected user ID 2, got %d", userID2\)\s*\}',
        """if userIdStr := validatedAPIKey2.GetUserId(); userIdStr != "2" {
		t.Errorf("expected user ID '2', got '%s'", userIdStr)
	}""",
        content,
    )

    file_path.write_text(content)
    print("Updated auth_test.go")
    return True


def update_session_test():
    """Update session_test.go to handle gcommon Session types."""
    file_path = Path("pkg/gcommonauth/session_test.go")
    if not file_path.exists():
        print(f"Error: {file_path} not found")
        return False

    content = file_path.read_text()

    # Pattern 1: Generate session
    content = re.sub(
        r"token, err := GenerateSession\(db, userID, 24\*time\.Hour\)",
        "sessionObj, err := GenerateSession(db, userID, 24*time.Hour)",
        content,
    )

    # Pattern 2: Validate session
    content = re.sub(
        r"validatedUserID, err := ValidateSession\(db, token\)",
        "validatedSession, err := ValidateSession(db, sessionObj.GetId())",
        content,
    )

    # Pattern 3: Check user ID from session
    content = re.sub(
        r'if validatedUserID != userID \{\s*t\.Errorf\("expected user ID %d, got %d", userID, validatedUserID\)\s*\}',
        """if userIdStr := validatedSession.GetUserId(); userIdStr != strconv.FormatInt(userID, 10) {
		t.Errorf("expected user ID %d, got %s", userID, userIdStr)
	}""",
        content,
    )

    # Pattern 4: Any remaining token usage in InvalidateSession
    content = re.sub(
        r"InvalidateSession\(db, token\)",
        "InvalidateSession(db, sessionObj.GetId())",
        content,
    )

    # Add strconv import if not present
    if "strconv" not in content and "strconv.FormatInt" in content:
        content = re.sub(
            r'import \(\s*"testing"\s*"time"',
            'import (\n\t"strconv"\n\t"testing"\n\t"time"',
            content,
        )

    file_path.write_text(content)
    print("Updated session_test.go")
    return True


def main():
    """Main function to update all test files."""
    print("Updating gcommonauth test files for gcommon migration...")

    success = True
    success &= update_auth_test()
    success &= update_session_test()

    if success:
        print("All test files updated successfully!")
        return 0
    else:
        print("Some updates failed!")
        return 1


if __name__ == "__main__":
    sys.exit(main())
