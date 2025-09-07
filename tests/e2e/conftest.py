# file: tests/e2e/conftest.py
# version: 1.0.0
# guid: a1b2c3d4-e5f6-7890-1234-567890abcdef

"""
Pytest configuration and fixtures for Selenium E2E testing.
Provides browser setup, video recording, screenshot capture, and test environment management.
"""

import os
import sys
import time
from datetime import datetime
from pathlib import Path
from typing import Any, Dict, Generator

# Optional imports for video recording
try:
    import cv2

    HAS_CV2 = True
except ImportError:
    HAS_CV2 = False

import pytest
from selenium import webdriver
from selenium.webdriver.chrome.options import Options as ChromeOptions
from selenium.webdriver.firefox.options import Options as FirefoxOptions
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import WebDriverWait
from webdriver_manager.chrome import ChromeDriverManager
from webdriver_manager.firefox import GeckoDriverManager

# Add project root to path for imports
PROJECT_ROOT = Path(__file__).parent.parent.parent
sys.path.insert(0, str(PROJECT_ROOT))

# Test configuration
TEST_CONFIG = {
    "base_url": os.getenv("TEST_BASE_URL", "http://127.0.0.1:5173"),
    "backend_url": os.getenv("TEST_BACKEND_URL", "http://localhost:8080"),
    "browser": os.getenv("TEST_BROWSER", "chrome"),
    "headless": os.getenv("TEST_HEADLESS", "false").lower() == "true",
    "window_size": os.getenv("TEST_WINDOW_SIZE", "1920,1080"),
    "timeout": int(os.getenv("TEST_TIMEOUT", "30")),
    "video_recording": os.getenv("TEST_VIDEO_RECORDING", "true").lower() == "true",
    "screenshot_on_failure": os.getenv("TEST_SCREENSHOT_ON_FAILURE", "true").lower()
    == "true",
    "parallel_workers": int(os.getenv("TEST_PARALLEL_WORKERS", "2")),
}

# Test directories
TESTS_DIR = Path(__file__).parent
SCREENSHOTS_DIR = TESTS_DIR / "screenshots"
RECORDINGS_DIR = TESTS_DIR / "recordings"
REPORTS_DIR = TESTS_DIR / "reports"
LOGS_DIR = TESTS_DIR / "logs"

# Ensure directories exist
for directory in [SCREENSHOTS_DIR, RECORDINGS_DIR, REPORTS_DIR, LOGS_DIR]:
    directory.mkdir(exist_ok=True)


class VideoRecorder:
    """Records video of test execution for debugging and documentation."""

    def __init__(self, output_path: str, fps: int = 10):
        self.output_path = output_path
        self.fps = fps
        self.writer = None
        self.recording = False

    def start_recording(self, window_size: tuple = (1920, 1080)):
        """Start video recording."""
        if not HAS_CV2:
            print("Video recording disabled - cv2 not available")
            return

        try:
            fourcc = cv2.VideoWriter_fourcc(*"mp4v")
            self.writer = cv2.VideoWriter(
                self.output_path, fourcc, self.fps, window_size
            )
            self.recording = True
            print(f"Started video recording: {self.output_path}")
        except Exception as e:
            print(f"Failed to start video recording: {e}")

    def add_frame(self, screenshot_path: str):
        """Add a frame to the video recording."""
        if not self.recording or not self.writer or not HAS_CV2:
            return

        try:
            img = cv2.imread(screenshot_path)
            if img is not None:
                self.writer.write(img)
        except Exception as e:
            print(f"Failed to add frame to video: {e}")

    def stop_recording(self):
        """Stop video recording."""
        if self.writer:
            self.writer.release()
            self.recording = False
            print(f"Stopped video recording: {self.output_path}")
        elif not HAS_CV2:
            print("Video recording was disabled - cv2 not available")


@pytest.fixture(scope="session")
def test_config() -> Dict[str, Any]:
    """Test configuration fixture."""
    return TEST_CONFIG


@pytest.fixture(scope="session")
def browser_type(test_config) -> str:
    """Browser type fixture."""
    return test_config["browser"]


@pytest.fixture
def driver(
    request, browser_type: str, test_config: Dict[str, Any]
) -> Generator[webdriver.Remote, None, None]:
    """
    WebDriver fixture with video recording and screenshot capabilities.
    Automatically handles browser setup, teardown, and test artifacts.
    """
    # Generate unique test identifier
    test_name = request.node.name
    timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
    test_id = f"{test_name}_{timestamp}"

    # Set up browser options
    driver_instance = None
    video_recorder = None

    try:
        if browser_type.lower() == "chrome":
            options = ChromeOptions()
            if test_config["headless"]:
                options.add_argument("--headless")
            options.add_argument("--no-sandbox")
            options.add_argument("--disable-dev-shm-usage")
            options.add_argument("--disable-gpu")
            options.add_argument(f"--window-size={test_config['window_size']}")
            options.add_argument("--start-maximized")

            driver_instance = webdriver.Chrome(
                service=webdriver.chrome.service.Service(
                    ChromeDriverManager().install()
                ),
                options=options,
            )
        else:  # Firefox
            options = FirefoxOptions()
            if test_config["headless"]:
                options.add_argument("--headless")
            options.add_argument(f"--width={test_config['window_size'].split(',')[0]}")
            options.add_argument(f"--height={test_config['window_size'].split(',')[1]}")

            driver_instance = webdriver.Firefox(
                service=webdriver.firefox.service.Service(
                    GeckoDriverManager().install()
                ),
                options=options,
            )

        # Configure implicit wait
        driver_instance.implicitly_wait(10)

        # Set up video recording
        if test_config["video_recording"]:
            video_path = str(RECORDINGS_DIR / f"{test_id}.mp4")
            window_size = tuple(map(int, test_config["window_size"].split(",")))
            video_recorder = VideoRecorder(video_path)
            video_recorder.start_recording(window_size)

        # Store test context
        driver_instance.test_id = test_id
        driver_instance.video_recorder = video_recorder
        driver_instance.test_config = test_config

        yield driver_instance

    except Exception as e:
        print(f"Failed to create WebDriver: {e}")
        raise

    finally:
        # Handle test failure artifacts
        if request.node.rep_call.failed if hasattr(request.node, "rep_call") else False:
            if test_config["screenshot_on_failure"] and driver_instance:
                screenshot_path = SCREENSHOTS_DIR / f"{test_id}_failure.png"
                driver_instance.save_screenshot(str(screenshot_path))
                print(f"Failure screenshot saved: {screenshot_path}")

        # Clean up
        if video_recorder:
            video_recorder.stop_recording()

        if driver_instance:
            driver_instance.quit()


@pytest.fixture
def wait(driver) -> WebDriverWait:
    """WebDriverWait fixture with configured timeout."""
    return WebDriverWait(driver, TEST_CONFIG["timeout"])


@pytest.fixture
def app_url(test_config: Dict[str, Any]) -> str:
    """Application URL fixture."""
    return test_config["base_url"]


@pytest.fixture
def api_url(test_config: Dict[str, Any]) -> str:
    """API URL fixture."""
    return test_config["backend_url"]


@pytest.hookimpl(tryfirst=True, hookwrapper=True)
def pytest_runtest_makereport(item, call):
    """Capture test results for failure handling."""
    outcome = yield
    rep = outcome.get_result()
    setattr(item, f"rep_{rep.when}", rep)


@pytest.fixture(autouse=True)
def test_logger(request):
    """Automatic test logging for debugging."""
    test_name = request.node.name
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    log_file = LOGS_DIR / f"{test_name}_{datetime.now().strftime('%Y%m%d')}.log"

    with open(log_file, "a") as f:
        f.write(f"\n=== {test_name} - {timestamp} ===\n")

    yield

    # Log test result
    if hasattr(request.node, "rep_call"):
        result = "PASSED" if request.node.rep_call.passed else "FAILED"
        with open(log_file, "a") as f:
            f.write(f"Result: {result}\n")


def pytest_configure(config):
    """Configure pytest with custom markers and settings."""
    config.addinivalue_line("markers", "smoke: mark test as smoke test")
    config.addinivalue_line("markers", "regression: mark test as regression test")
    config.addinivalue_line("markers", "slow: mark test as slow running")
    config.addinivalue_line("markers", "critical: mark test as critical functionality")


def pytest_collection_modifyitems(config, items):
    """Modify test collection to add markers based on file location."""
    for item in items:
        # Add markers based on file path
        if "smoke" in str(item.fspath):
            item.add_marker(pytest.mark.smoke)
        if "regression" in str(item.fspath):
            item.add_marker(pytest.mark.regression)
        if "critical" in str(item.fspath):
            item.add_marker(pytest.mark.critical)


# Utility functions for tests
class TestUtils:
    """Utility class for common test operations."""

    @staticmethod
    def wait_for_page_load(driver, timeout: int = 30):
        """Wait for page to fully load."""
        WebDriverWait(driver, timeout).until(
            lambda d: d.execute_script("return document.readyState") == "complete"
        )

    @staticmethod
    def wait_for_element_visible(driver, locator: tuple, timeout: int = 30):
        """Wait for element to be visible."""
        return WebDriverWait(driver, timeout).until(
            EC.visibility_of_element_located(locator)
        )

    @staticmethod
    def wait_for_element_clickable(driver, locator: tuple, timeout: int = 30):
        """Wait for element to be clickable."""
        return WebDriverWait(driver, timeout).until(EC.element_to_be_clickable(locator))

    @staticmethod
    def take_screenshot(driver, name: str = None) -> str:
        """Take a screenshot and return the path."""
        if not name:
            name = f"screenshot_{datetime.now().strftime('%Y%m%d_%H%M%S')}"

        screenshot_path = SCREENSHOTS_DIR / f"{name}.png"
        driver.save_screenshot(str(screenshot_path))

        # Add frame to video if recording
        if hasattr(driver, "video_recorder") and driver.video_recorder:
            driver.video_recorder.add_frame(str(screenshot_path))

        return str(screenshot_path)

    @staticmethod
    def scroll_to_element(driver, element):
        """Scroll element into view."""
        driver.execute_script("arguments[0].scrollIntoView(true);", element)
        time.sleep(0.5)  # Brief pause for scroll to complete
