#!/bin/bash
# file: .devcontainer/post-create.sh
# Post-creation script for the dev container

set -e

echo "🚀 Setting up Subtitle Manager development environment..."

# Ensure we're in the workspace directory
cd /workspace

# Install Go dependencies
echo "📦 Installing Go dependencies..."
go mod download

# Install Node.js dependencies for the web UI
echo "🌐 Installing Node.js dependencies..."
cd webui
npm ci --legacy-peer-deps
cd ..

# Install pre-commit hooks
echo "🔧 Installing pre-commit hooks..."
if [ -f "scripts/install-pre-commit-hooks.sh" ]; then
    ./scripts/install-pre-commit-hooks.sh
fi

# Build the web UI initially
echo "🏗️  Building web UI..."
cd webui
npm run build
cd ..

# Run go generate to ensure embedded assets are ready
echo "⚙️  Running go generate..."
go generate ./webui

# Set up git safe directory (since we're running as vscode user)
echo "🔒 Configuring git safe directory..."
git config --global --add safe.directory /workspace

# Create development directories
echo "📁 Creating development directories..."
mkdir -p /workspace/tmp
mkdir -p /workspace/logs

# Set proper permissions
sudo chown -R vscode:vscode /workspace

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