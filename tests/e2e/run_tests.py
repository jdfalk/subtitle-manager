#!/usr/bin/env python3
# file: tests/e2e/run_tests.py
# version: 1.0.0
# guid: e5f6a7b8-c9d0-1234-5678-901234567890

"""
Test runner script for Selenium E2E tests.
Provides convenient commands for running different test suites with proper setup.
"""

import argparse
import os
import signal
import subprocess
import sys
import time
from pathlib import Path

import requests

# Add project root to path
PROJECT_ROOT = Path(__file__).parent.parent.parent
sys.path.insert(0, str(PROJECT_ROOT))


class TestRunner:
    """Manages test execution with proper setup and teardown."""

    def __init__(self):
        self.tests_dir = Path(__file__).parent
        self.project_root = PROJECT_ROOT
        self.frontend_process = None
        self.backend_process = None

    def check_dependencies(self):
        """Check that required dependencies are installed."""
        try:
            import pytest
            import selenium

            print("✓ Selenium and pytest are available")
            return True
        except ImportError as e:
            print(f"✗ Missing dependencies: {e}")
            print("Install with: pip install -r tests/e2e/requirements.txt")
            return False

    def is_service_running(self, url: str, timeout: int = 5) -> bool:
        """Check if a service is running at the given URL."""
        try:
            response = requests.get(url, timeout=timeout)
            return response.status_code < 500
        except requests.RequestException:
            return False

    def start_frontend(self, port: int = 5173) -> bool:
        """Start the frontend development server."""
        webui_dir = self.project_root / "webui"

        if not webui_dir.exists():
            print("✗ Frontend directory not found")
            return False

        # Check if already running
        if self.is_service_running(f"http://localhost:{port}"):
            print(f"✓ Frontend already running on port {port}")
            return True

        print(f"Starting frontend server on port {port}...")

        try:
            # Start npm dev server
            self.frontend_process = subprocess.Popen(
                ["npm", "run", "dev", "--", "--port", str(port)],
                cwd=webui_dir,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                preexec_fn=os.setsid if os.name != "nt" else None,
            )

            # Wait for server to start
            for _ in range(30):  # Wait up to 30 seconds
                if self.is_service_running(f"http://localhost:{port}"):
                    print(f"✓ Frontend server started on port {port}")
                    return True
                time.sleep(1)

            print("✗ Frontend server failed to start within 30 seconds")
            return False

        except Exception as e:
            print(f"✗ Failed to start frontend server: {e}")
            return False

    def start_backend(self, port: int = 8080) -> bool:
        """Start the backend server."""
        # Check if already running
        if self.is_service_running(f"http://localhost:{port}"):
            print(f"✓ Backend already running on port {port}")
            return True

        print(f"Starting backend server on port {port}...")

        try:
            # Build and start Go server
            build_cmd = ["go", "build", "-o", "subtitle-manager", "."]
            subprocess.run(build_cmd, cwd=self.project_root, check=True)

            self.backend_process = subprocess.Popen(
                ["./subtitle-manager", "web", "--port", str(port)],
                cwd=self.project_root,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                preexec_fn=os.setsid if os.name != "nt" else None,
            )

            # Wait for server to start
            for _ in range(15):  # Wait up to 15 seconds
                if self.is_service_running(f"http://localhost:{port}"):
                    print(f"✓ Backend server started on port {port}")
                    return True
                time.sleep(1)

            print("✗ Backend server failed to start within 15 seconds")
            return False

        except Exception as e:
            print(f"✗ Failed to start backend server: {e}")
            return False

    def stop_services(self):
        """Stop any services started by this runner."""
        if self.frontend_process:
            try:
                if os.name != "nt":
                    os.killpg(os.getpgid(self.frontend_process.pid), signal.SIGTERM)
                else:
                    self.frontend_process.terminate()
                self.frontend_process.wait(timeout=5)
                print("✓ Frontend server stopped")
            except Exception as e:
                print(f"Warning: Failed to stop frontend server: {e}")

        if self.backend_process:
            try:
                if os.name != "nt":
                    os.killpg(os.getpgid(self.backend_process.pid), signal.SIGTERM)
                else:
                    self.backend_process.terminate()
                self.backend_process.wait(timeout=5)
                print("✓ Backend server stopped")
            except Exception as e:
                print(f"Warning: Failed to stop backend server: {e}")

    def run_tests(self, test_args: list = None) -> int:
        """Run the test suite."""
        if not self.check_dependencies():
            return 1

        # Change to tests directory
        os.chdir(self.tests_dir)

        # Build pytest command
        cmd = ["python", "-m", "pytest"]

        if test_args:
            cmd.extend(test_args)
        else:
            # Default test run
            cmd.extend(
                ["-v", "--html=reports/test_report.html", "--self-contained-html"]
            )

        print(f"Running: {' '.join(cmd)}")

        try:
            result = subprocess.run(cmd, check=False)
            return result.returncode
        except KeyboardInterrupt:
            print("\nTests interrupted by user")
            return 130


def main():
    parser = argparse.ArgumentParser(description="Run E2E tests for subtitle-manager")
    parser.add_argument(
        "--start-services",
        action="store_true",
        help="Start frontend and backend services before testing",
    )
    parser.add_argument(
        "--frontend-only",
        action="store_true",
        help="Start only frontend service (backend assumed running)",
    )
    parser.add_argument("--smoke", action="store_true", help="Run only smoke tests")
    parser.add_argument(
        "--critical", action="store_true", help="Run only critical tests"
    )
    parser.add_argument(
        "--settings", action="store_true", help="Run only settings navigation tests"
    )
    parser.add_argument(
        "--browser",
        choices=["chrome", "firefox"],
        default="chrome",
        help="Browser to use for testing",
    )
    parser.add_argument(
        "--headless", action="store_true", help="Run tests in headless mode"
    )
    parser.add_argument("--verbose", "-v", action="store_true", help="Verbose output")

    args = parser.parse_args()

    runner = TestRunner()

    try:
        # Set environment variables
        os.environ["TEST_BROWSER"] = args.browser
        if args.headless:
            os.environ["TEST_HEADLESS"] = "true"

        # Start services if requested
        if args.start_services:
            if not runner.start_frontend():
                return 1
            if not runner.start_backend():
                return 1
        elif args.frontend_only:
            if not runner.start_frontend():
                return 1

        # Build test arguments
        test_args = []

        if args.smoke:
            test_args.extend(["-m", "smoke"])
        elif args.critical:
            test_args.extend(["-m", "critical"])
        elif args.settings:
            test_args.append("test_settings_navigation.py")

        if args.verbose:
            test_args.append("-v")

        # Run tests
        return runner.run_tests(test_args)

    except KeyboardInterrupt:
        print("\nShutting down...")
        return 130
    finally:
        if args.start_services or args.frontend_only:
            runner.stop_services()


if __name__ == "__main__":
    sys.exit(main())
