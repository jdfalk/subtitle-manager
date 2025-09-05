#!/usr/bin/env python3
# file: tests/e2e/test_live.py
# version: 1.0.0
# guid: f1e2d3c4-b5a6-7890-1234-567890abcdef

"""
Quick test runner for live development server.
Assumes server is already running - won't start/stop it.
"""

import sys
import time
from pathlib import Path

import requests


def check_server(url: str = "http://localhost:5173") -> bool:
    """Check if the development server is running."""
    try:
        response = requests.get(url, timeout=5)
        return response.status_code == 200
    except:
        return False


def main():
    """Run tests against live server."""
    print("🧪 Running tests against live server...")

    # Run pytest directly - assume server is running
    import subprocess

    result = subprocess.run(
        [
            sys.executable,
            "-m",
            "pytest",
            "-v",
            "--tb=short",
            "-x",  # Stop on first failure
            "test_simple.py",
            "test_settings_navigation.py",
        ],
        cwd=Path(__file__).parent,
    )

    return result.returncode


if __name__ == "__main__":
    sys.exit(main())
