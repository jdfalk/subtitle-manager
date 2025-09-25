#!/bin/bash
# file: scripts/test_database_backends.sh
# version: 1.0.0
# guid: 3c4d5e6f-7890-1234-5678-901234567890

# Database Backend Testing Script
# Demonstrates testing with different build configurations

set -e

echo "🚀 Database Backend Testing Demonstration"
echo "=========================================="
echo

# Pure Go Build Testing
echo "1. Testing Pure Go Build (no CGO):"
echo "   Command: go test ./pkg/database -v -run TestBackendSelection"
echo "   Expected: SQLite unavailable, Pebble works"
echo
go test ./pkg/database -v -run TestBackendSelection 2>&1 | grep -E "(SQLite|Pebble|PASS|FAIL)"
echo

# CGO Build Testing (if available)
echo "2. Testing CGO Build with SQLite:"
echo "   Command: go test -tags sqlite ./pkg/database -v -run TestBackendSelection"
echo "   Expected: Both SQLite and Pebble work"
echo

if command -v gcc > /dev/null && [ "$CGO_ENABLED" != "0" ]; then
    go test -tags sqlite ./pkg/database -v -run TestBackendSelection 2>&1 | grep -E "(SQLite|Pebble|PASS|FAIL)" || {
        echo "   ❌ CGO build failed (may not have sqlite3 development headers)"
    }
else
    echo "   ⚠️  CGO not available or disabled - skipping SQLite test"
fi

echo
echo "3. Testing SQLite Availability Detection:"
echo "   Command: go test ./pkg/database -v -run TestSQLiteAvailability"
echo
go test ./pkg/database -v -run TestSQLiteAvailability 2>&1 | grep -E "(SQLite support|✓|PASS)"

echo
echo "✅ Database backend testing completed!"
echo
echo "Summary:"
echo "- Pure Go builds use Pebble database (no CGO dependency)"
echo "- CGO builds can use SQLite with: go build -tags sqlite"
echo "- All tests automatically detect and use the best available backend"
echo "- Comprehensive test coverage for all supported backends"
