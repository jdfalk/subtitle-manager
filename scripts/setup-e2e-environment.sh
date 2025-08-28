#!/bin/bash
# file: scripts/setup-e2e-environment.sh
# version: 1.0.0
# guid: e2e-setup-12345-6789-abcd-ef01

# End-to-End Testing Environment Setup Script
# This script sets up the complete e2e testing environment

set -euo pipefail

# Colors for output
COLOR_RESET='\033[0m'
COLOR_BOLD='\033[1m'
COLOR_GREEN='\033[32m'
COLOR_YELLOW='\033[33m'
COLOR_BLUE='\033[34m'
COLOR_RED='\033[31m'

# Configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
TESTDIR="$PROJECT_DIR/testdir"
BINARY_PATH="$PROJECT_DIR/bin/subtitle-manager"

# Default configuration
DEFAULT_PORT=55327
DEFAULT_USERNAME="test"
DEFAULT_PASSWORD="test123"

print_header() {
    echo -e "${COLOR_BOLD}${COLOR_BLUE}=== Subtitle Manager E2E Testing Setup ===${COLOR_RESET}"
    echo ""
}

print_status() {
    echo -e "${COLOR_GREEN}✓${COLOR_RESET} $1"
}

print_warning() {
    echo -e "${COLOR_YELLOW}⚠${COLOR_RESET} $1"
}

print_error() {
    echo -e "${COLOR_RED}✗${COLOR_RESET} $1"
}

check_prerequisites() {
    echo -e "${COLOR_BLUE}Checking prerequisites...${COLOR_RESET}"

    # Check if binary exists
    if [[ ! -f "$BINARY_PATH" ]]; then
        print_warning "Binary not found at $BINARY_PATH"
        echo "Building subtitle-manager binary..."
        cd "$PROJECT_DIR"
        make build
        if [[ ! -f "$BINARY_PATH" ]]; then
            print_error "Failed to build binary"
            exit 1
        fi
    fi
    print_status "Binary found: $BINARY_PATH"

    # Check if testdir exists
    if [[ ! -d "$TESTDIR" ]]; then
        print_error "Test directory not found: $TESTDIR"
        print_error "Please run 'make setup-e2e-testdir' first"
        exit 1
    fi
    print_status "Test directory found: $TESTDIR"

    # Check for required tools
    for tool in curl jq; do
        if ! command -v "$tool" &> /dev/null; then
            print_warning "$tool not found - some tests may fail"
        else
            print_status "$tool available"
        fi
    done
}

start_subtitle_manager() {
    echo -e "${COLOR_BLUE}Starting Subtitle Manager...${COLOR_RESET}"

    # Kill any existing instances
    pkill -f subtitle-manager || true
    sleep 2

    # Start the server in background
    cd "$PROJECT_DIR"
    export SUBTITLE_MANAGER_PORT="$DEFAULT_PORT"
    export SUBTITLE_MANAGER_USERNAME="$DEFAULT_USERNAME"
    export SUBTITLE_MANAGER_PASSWORD="$DEFAULT_PASSWORD"
    export SUBTITLE_MANAGER_MEDIA_PATH="$TESTDIR"

    "$BINARY_PATH" web \
        --addr "127.0.0.1:$DEFAULT_PORT" \
        --admin-user "$DEFAULT_USERNAME" \
        --admin-password "$DEFAULT_PASSWORD" \
        --log-level debug \
        > "/tmp/subtitle-manager-e2e.log" 2>&1 &

    local PID=$!
    echo "$PID" > "/tmp/subtitle-manager-e2e.pid"

    # Wait for server to start
    echo "Starting server on port $DEFAULT_PORT..."
    for i in {1..30}; do
        if curl -s "http://127.0.0.1:$DEFAULT_PORT/health" > /dev/null 2>&1; then
            print_status "Server started successfully (PID: $PID)"
            return 0
        fi
        sleep 1
    done

    print_error "Server failed to start within 30 seconds"
    print_error "Check logs: tail -f /tmp/subtitle-manager-e2e.log"
    return 1
}

run_basic_tests() {
    echo -e "${COLOR_BLUE}Running basic functionality tests...${COLOR_RESET}"

    local BASE_URL="http://127.0.0.1:$DEFAULT_PORT"

    # Test health endpoint
    if curl -s "$BASE_URL/health" | jq -e '.status == "ok"' > /dev/null; then
        print_status "Health check passed"
    else
        print_error "Health check failed"
        return 1
    fi

    # Test API endpoints (if available)
    if curl -s "$BASE_URL/api/v1/subtitles" > /dev/null 2>&1; then
        print_status "API endpoints accessible"
    else
        print_warning "API endpoints not available or not implemented"
    fi

    # Test file discovery
    local subtitle_count=$(find "$TESTDIR" -name "*.srt" -o -name "*.ass" -o -name "*.sub" | wc -l)
    if [[ $subtitle_count -gt 0 ]]; then
        print_status "Found $subtitle_count subtitle files in test directory"
    else
        print_error "No subtitle files found in test directory"
        return 1
    fi
}

show_access_info() {
    echo ""
    echo -e "${COLOR_BOLD}${COLOR_GREEN}=== E2E Testing Environment Ready ===${COLOR_RESET}"
    echo ""
    echo -e "${COLOR_BLUE}Web Interface:${COLOR_RESET}"
    echo "  URL: http://127.0.0.1:$DEFAULT_PORT"
    echo "  Username: $DEFAULT_USERNAME"
    echo "  Password: $DEFAULT_PASSWORD"
    echo ""
    echo -e "${COLOR_BLUE}Test Data Location:${COLOR_RESET}"
    echo "  Path: $TESTDIR"
    echo "  Structure:"
    echo "    ├── movies/"
    echo "    │   ├── The Matrix (1999)/"
    echo "    │   └── Blade Runner 2049 (2017)/"
    echo "    ├── tv/"
    echo "    │   ├── Breaking Bad (2008)/"
    echo "    │   └── The Office (2005)/"
    echo "    └── anime/"
    echo "        ├── Attack on Titan (2013)/"
    echo "        └── Your Name (2016)/"
    echo ""
    echo -e "${COLOR_BLUE}Management:${COLOR_RESET}"
    echo "  Stop server: make stop-e2e"
    echo "  View logs: tail -f /tmp/subtitle-manager-e2e.log"
    echo "  Clean up: make clean-e2e"
    echo ""
    echo -e "${COLOR_YELLOW}Note: The server is running in the background.${COLOR_RESET}"
    echo -e "${COLOR_YELLOW}Use 'make stop-e2e' to stop it when done testing.${COLOR_RESET}"
}

cleanup_on_error() {
    print_error "Setup failed. Cleaning up..."
    pkill -f subtitle-manager || true
    rm -f "/tmp/subtitle-manager-e2e.pid"
    exit 1
}

main() {
    trap cleanup_on_error ERR

    print_header
    check_prerequisites
    start_subtitle_manager
    run_basic_tests
    show_access_info
}

# Allow script to be sourced for individual functions
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi
