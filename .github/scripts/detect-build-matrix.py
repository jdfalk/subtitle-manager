#!/usr/bin/env python3
# file: .github/scripts/detect-build-matrix.py
# version: 1.0.0
# guid: f0a1b2c3-d4e5-6789-012f-345678901234

"""
Dynamic Build Matrix Detection for GitHub Actions

This script analyzes the repository structure and generates appropriate
build matrices for different technology stacks (Go, Python, Frontend, Docker).
It outputs GitHub Actions workflow outputs that can be consumed by matrix builds.
"""

import json
import os
from pathlib import Path
from typing import Any, Dict, List


def check_file_exists(patterns: List[str]) -> bool:
    """Check if any of the given file patterns exist in the repository."""
    for pattern in patterns:
        if list(Path(".").glob(pattern)):
            return True
    return False


def detect_go_project() -> Dict[str, Any]:
    """Detect Go project and generate build matrix."""
    go_files = [
        "go.mod",
        "go.sum",
        "*.go",
        "**/*.go",
        "cmd/**",
        "pkg/**",
        "internal/**",
    ]

    if not check_file_exists(go_files):
        return {"include": []}

    # Default Go matrix - can be customized based on project needs
    matrix = {
        "include": [
            {"go-version": "1.21", "os": "ubuntu-latest", "primary": True},
            {"go-version": "1.20", "os": "ubuntu-latest", "primary": False},
            {"go-version": "1.21", "os": "windows-latest", "primary": False},
            {"go-version": "1.21", "os": "macos-latest", "primary": False},
        ]
    }

    return matrix


def detect_python_project() -> Dict[str, Any]:
    """Detect Python project and generate build matrix."""
    python_files = [
        "requirements.txt",
        "pyproject.toml",
        "setup.py",
        "setup.cfg",
        "*.py",
        "**/*.py",
        "src/**/*.py",
    ]

    if not check_file_exists(python_files):
        return {"include": []}

    # Default Python matrix
    matrix = {
        "include": [
            {"python-version": "3.11", "os": "ubuntu-latest", "primary": True},
            {"python-version": "3.10", "os": "ubuntu-latest", "primary": False},
            {"python-version": "3.9", "os": "ubuntu-latest", "primary": False},
            {"python-version": "3.11", "os": "windows-latest", "primary": False},
            {"python-version": "3.11", "os": "macos-latest", "primary": False},
        ]
    }

    return matrix


def detect_frontend_project() -> Dict[str, Any]:
    """Detect frontend project and generate build matrix."""
    frontend_files = [
        "package.json",
        "package-lock.json",
        "yarn.lock",
        "webpack.config.js",
        "vite.config.js",
        "angular.json",
        "src/**/*.ts",
        "src/**/*.tsx",
        "src/**/*.js",
        "src/**/*.jsx",
    ]

    if not check_file_exists(frontend_files):
        return {"include": []}

    # Default Node.js matrix
    matrix = {
        "include": [
            {"node-version": "18", "os": "ubuntu-latest", "primary": True},
            {"node-version": "16", "os": "ubuntu-latest", "primary": False},
            {"node-version": "20", "os": "ubuntu-latest", "primary": False},
        ]
    }

    return matrix


def detect_docker_project() -> Dict[str, Any]:
    """Detect Docker project and generate build matrix."""
    docker_files = [
        "Dockerfile",
        "docker-compose.yml",
        "docker-compose.yaml",
        ".dockerignore",
        "Dockerfile.*",
    ]

    if not check_file_exists(docker_files):
        return {"include": []}

    # Default Docker matrix - multi-platform builds
    matrix = {
        "include": [
            {"platform": "linux/amd64", "os": "ubuntu-latest", "primary": True},
            {"platform": "linux/arm64", "os": "ubuntu-latest", "primary": False},
        ]
    }

    return matrix


def detect_protobuf_project() -> bool:
    """Detect if protobuf generation is needed."""
    protobuf_files = [
        "buf.yaml",
        "buf.gen.yaml",
        "buf.work.yaml",
        "proto/**/*.proto",
        "*.proto",
        "**/*.proto",
    ]

    return check_file_exists(protobuf_files)


def main():
    """Main function to detect build requirements and output matrices."""

    # Detect each technology stack
    go_matrix = detect_go_project()
    python_matrix = detect_python_project()
    frontend_matrix = detect_frontend_project()
    docker_matrix = detect_docker_project()
    protobuf_needed = detect_protobuf_project()

    # Check if each stack has any builds
    has_go = len(go_matrix["include"]) > 0
    has_python = len(python_matrix["include"]) > 0
    has_frontend = len(frontend_matrix["include"]) > 0
    has_docker = len(docker_matrix["include"]) > 0

    # Set GitHub Actions outputs
    outputs = {
        "go-matrix": json.dumps(go_matrix),
        "python-matrix": json.dumps(python_matrix),
        "frontend-matrix": json.dumps(frontend_matrix),
        "docker-matrix": json.dumps(docker_matrix),
        "protobuf-needed": str(protobuf_needed).lower(),
        "has-go": str(has_go).lower(),
        "has-python": str(has_python).lower(),
        "has-frontend": str(has_frontend).lower(),
        "has-docker": str(has_docker).lower(),
    }

    # Write to GitHub Actions outputs
    github_output = os.getenv("GITHUB_OUTPUT")
    if github_output:
        with open(github_output, "a") as f:
            for key, value in outputs.items():
                f.write(f"{key}={value}\n")

    # Also print for debugging
    print("Detected build requirements:")
    for key, value in outputs.items():
        print(f"  {key}: {value}")

    # Print summary
    technologies = []
    if has_go:
        technologies.append("Go")
    if has_python:
        technologies.append("Python")
    if has_frontend:
        technologies.append("Frontend")
    if has_docker:
        technologies.append("Docker")
    if protobuf_needed:
        technologies.append("Protobuf")

    if technologies:
        print(f"\n✅ Build matrix will include: {', '.join(technologies)}")
    else:
        print("\n❌ No recognized technology stacks found")


if __name__ == "__main__":
    main()
