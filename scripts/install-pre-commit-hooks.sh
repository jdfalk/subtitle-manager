#!/bin/bash
# file: scripts/install-pre-commit-hooks.sh

# Script to install pre-commit hooks for automatic code formatting
# This is optional for developers who want to format code before pushing

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Installing pre-commit hooks for subtitle-manager...${NC}"

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    echo -e "${RED}Error: This script must be run from the root of the git repository${NC}"
    exit 1
fi

# Create the pre-commit hook
HOOK_FILE=".git/hooks/pre-commit"

cat > "$HOOK_FILE" << 'EOF'
#!/bin/bash
# file: .git/hooks/pre-commit
# Pre-commit hook for automatic code formatting

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to format Go code
format_go() {
    echo "Formatting Go code..."

    # Find all staged Go files
    GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$' || true)

    if [ -n "$GO_FILES" ]; then
        # Format with gofmt
        gofmt -s -w $GO_FILES

        # Format with goimports if available
        if command_exists goimports; then
            goimports -w $GO_FILES
        else
            echo "goimports not found, installing..."
            go install golang.org/x/tools/cmd/goimports@latest
            goimports -w $GO_FILES
        fi

        # Re-stage the formatted files
        git add $GO_FILES
        echo "Go files formatted and re-staged"
    fi
}

# Function to format frontend code
format_frontend() {
    echo "Formatting frontend code..."

    # Check if we have frontend files to format
    FRONTEND_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep -E 'webui/.*\.(js|jsx|ts|tsx|css|json|md)$' || true)

    if [ -n "$FRONTEND_FILES" ] && [ -f "webui/package.json" ]; then
        cd webui

        # Check if prettier is available
        if npm list prettier > /dev/null 2>&1; then
            # Format the staged files
            echo "$FRONTEND_FILES" | sed 's|webui/||g' | xargs npm run format -- --write

            # Re-stage the formatted files
            cd ..
            git add $FRONTEND_FILES
            echo "Frontend files formatted and re-staged"
        else
            echo "Prettier not found in webui, skipping frontend formatting"
        fi

        cd ..
    fi
}

# Main execution
echo "Running pre-commit formatting..."

# Check if Go is available
if command_exists go; then
    format_go
else
    echo "Go not found, skipping Go formatting"
fi

# Check if Node.js is available
if command_exists npm && [ -f "webui/package.json" ]; then
    format_frontend
else
    echo "Node.js/npm not found or no webui directory, skipping frontend formatting"
fi

echo "Pre-commit formatting complete!"
EOF

# Make the hook executable
chmod +x "$HOOK_FILE"

echo -e "${GREEN}✅ Pre-commit hook installed successfully!${NC}"
echo -e "${YELLOW}The hook will now automatically format your Go and frontend code before each commit.${NC}"
echo ""
echo -e "${YELLOW}To disable the hook temporarily, use:${NC}"
echo "  git commit --no-verify"
echo ""
echo -e "${YELLOW}To uninstall the hook, delete:${NC}"
echo "  .git/hooks/pre-commit"
echo ""
echo -e "${GREEN}Happy coding! 🚀${NC}"
