#!/bin/bash
# file: .devcontainer/post-create.sh
# version: 1.0.0
# guid: 47c8b2f5-e9d4-4c8a-b7f3-5a8c9d2e1f40
# Post-creation script for the dev container

set -e

echo "🚀 Setting up Subtitle Manager development environment..."

# Ensure we're in the workspace directory
cd /workspace

# Set up git safe directory first (before any git operations)
echo "🔒 Configuring git safe directory..."
git config --global --add safe.directory /workspace

# Install Go dependencies
echo "📦 Installing Go dependencies..."
if [ -f "go.mod" ]; then
    go mod download
else
    echo "⚠️  No go.mod found, skipping Go dependency installation"
fi

# Install Node.js dependencies for the web UI
echo "🌐 Installing Node.js dependencies..."
if [ -d "webui" ] && [ -f "webui/package.json" ]; then
    cd webui
    npm ci --legacy-peer-deps || npm install --legacy-peer-deps
    cd ..
else
    echo "⚠️  No webui directory or package.json found, skipping Node.js dependency installation"
fi

# Install pre-commit hooks
echo "🔧 Installing pre-commit hooks..."
if [ -f "scripts/install-pre-commit-hooks.sh" ]; then
    ./scripts/install-pre-commit-hooks.sh
else
    echo "⚠️  No pre-commit hooks script found, skipping"
fi

# Build the web UI initially
echo "🏗️  Building web UI..."
if [ -d "webui" ] && [ -f "webui/package.json" ]; then
    cd webui
    npm run build || echo "⚠️  Web UI build failed, continuing..."
    cd ..
else
    echo "⚠️  No webui directory found, skipping web UI build"
fi

# Run go generate to ensure embedded assets are ready
echo "⚙️  Running go generate..."
if [ -f "go.mod" ]; then
    go generate ./webui || echo "⚠️  go generate failed, continuing..."
else
    echo "⚠️  No go.mod found, skipping go generate"
fi

# Create development directories
echo "📁 Creating development directories..."
mkdir -p /workspace/tmp
mkdir -p /workspace/logs
mkdir -p /workspace/bin

# Set proper permissions
sudo chown -R vscode:vscode /workspace || true

# Final message
echo "✅ Development environment setup complete!"
echo ""
echo "🎯 Quick start commands:"
echo "  make build          # Build the application"
echo "  make test           # Run tests"
echo "  make test-sqlite    # Run tests with SQLite support"
echo "  make dev            # Start development server"
echo "  cd webui && npm run dev  # Start web UI development server"
echo ""
echo "🌐 Web UI will be available at http://localhost:3000"
echo "🖥️  Application will be available at http://localhost:8080"
