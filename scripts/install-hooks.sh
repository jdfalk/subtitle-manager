#!/bin/sh
# file: scripts/install-hooks.sh
#
# Script to install Git hooks for the subtitle-manager project
# This ensures consistent code quality across all contributors

set -e

echo "Installing Git hooks for subtitle-manager..."

# Create the pre-commit hook
cat > .git/hooks/pre-commit << 'EOF'
#!/bin/sh
# Pre-commit hook that runs gofmt and go vet on Go files
# This ensures code quality and consistency before commits

set -e

echo "Running pre-commit checks..."

# Check if there are any Go files staged for commit
go_files=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$' || true)

if [ -z "$go_files" ]; then
    echo "No Go files staged for commit."
    exit 0
fi

echo "Found Go files staged for commit:"
echo "$go_files"

# Check Go formatting
echo "Checking Go formatting..."
unformatted=$(echo "$go_files" | xargs gofmt -s -l)
if [ -n "$unformatted" ]; then
    echo "ERROR: The following files are not properly formatted:"
    echo "$unformatted"
    echo ""
    echo "Please run 'gofmt -s -w' on these files and try again:"
    echo "$unformatted" | xargs -I {} echo "  gofmt -s -w {}"
    echo ""
    echo "Or run: gofmt -s -w \$(git diff --cached --name-only --diff-filter=ACM | grep '\\.go\$')"
    exit 1
fi

# Run go vet on staged files
echo "Running go vet..."
if ! go vet ./...; then
    echo "ERROR: go vet found issues. Please fix them before committing."
    exit 1
fi

echo "All pre-commit checks passed!"
exit 0
EOF

# Make the hook executable
chmod +x .git/hooks/pre-commit

echo "✅ Pre-commit hook installed successfully!"
echo ""
echo "The hook will now run automatically before each commit and will:"
echo "  - Check Go file formatting (gofmt -s)"
echo "  - Run go vet for code analysis"
echo ""
echo "To bypass the hook temporarily, use: git commit --no-verify"
