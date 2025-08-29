# TASK-02-002: Implement Selenium-Based E2E Testing

<!-- file: docs/tasks/02-ui-fixes/TASK-02-002-selenium-testing.md -->
<!-- version: 1.0.0 -->
<!-- guid: f5g6h7i8-j9k0-1234-5678-9abcdef01234 -->

## 🎯 Objective

Implement comprehensive end-to-end testing using Selenium WebDriver with video
recording capabilities for all user workflows in the subtitle-manager
application.

## 📋 Acceptance Criteria

- [ ] Set up Selenium WebDriver with Chrome/Firefox support
- [ ] Implement video recording for all test sessions
- [ ] Create comprehensive test suites for all user workflows
- [ ] Add parallel test execution capabilities
- [ ] Generate detailed test reports with screenshots
- [ ] Set up CI/CD integration for automated testing
- [ ] Create performance benchmarking tests
- [ ] Implement visual regression testing

## 🔍 Current State Analysis

### Existing Test Infrastructure

Current testing appears limited. Need to add:

1. **E2E Test Framework**: Selenium-based testing
2. **Video Recording**: Screen capture for debugging
3. **Test Coverage**: All user workflows
4. **CI Integration**: Automated test execution
5. **Performance Testing**: Load and stress testing

### Key User Workflows to Test

1. **Authentication Flow**: Login, logout, session management
2. **Media Management**: Add/remove/scan media libraries
3. **Subtitle Operations**: Search, download, extract, translate
4. **Provider Configuration**: Add/configure/test providers
5. **Settings Management**: User preferences, system settings
6. **System Monitoring**: Health checks, logs, statistics

## 🔧 Implementation Steps

### Step 1: Set up testing environment

```bash
# Create testing directory structure
mkdir -p tests/e2e/{tests,reports,recordings,screenshots,fixtures}
mkdir -p tests/e2e/utils/{drivers,helpers,pages}

# Create requirements file
cat > tests/e2e/requirements.txt << 'EOF'
selenium==4.15.2
pytest==7.4.3
pytest-html==4.1.1
pytest-xdist==3.3.1
pytest-rerunfailures==12.0
webdriver-manager==4.0.1
opencv-python==4.8.1.78
pillow==10.1.0
allure-pytest==2.13.2
requests==2.31.0
beautifulsoup4==4.12.2
EOF

# Install dependencies
cd tests/e2e
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
pip install -r requirements.txt
```

### Step 2: Create WebDriver configuration

Create `tests/e2e/utils/driver_manager.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/utils/driver_manager.py
# version: 1.0.0
# guid: a1b2c3d4-e5f6-7890-1234-567890abcdef

import os
import cv2
import time
import threading
from selenium import webdriver
from selenium.webdriver.chrome.service import Service as ChromeService
from selenium.webdriver.firefox.service import Service as FirefoxService
from selenium.webdriver.chrome.options import Options as ChromeOptions
from selenium.webdriver.firefox.options import Options as FirefoxOptions
from webdriver_manager.chrome import ChromeDriverManager
from webdriver_manager.firefox import GeckoDriverManager

class VideoRecorder:
    """Records screen during test execution"""

    def __init__(self, output_path, fps=10):
        self.output_path = output_path
        self.fps = fps
        self.recording = False
        self.video_writer = None
        self.thread = None

    def start_recording(self, screen_size=(1920, 1080)):
        """Start video recording"""
        fourcc = cv2.VideoWriter_fourcc(*'mp4v')
        self.video_writer = cv2.VideoWriter(
            self.output_path, fourcc, self.fps, screen_size
        )
        self.recording = True
        self.thread = threading.Thread(target=self._record_loop)
        self.thread.start()

    def stop_recording(self):
        """Stop video recording"""
        self.recording = False
        if self.thread:
            self.thread.join()
        if self.video_writer:
            self.video_writer.release()

    def _record_loop(self):
        """Recording loop"""
        import pyautogui
        while self.recording:
            screenshot = pyautogui.screenshot()
            frame = cv2.cvtColor(np.array(screenshot), cv2.COLOR_RGB2BGR)
            self.video_writer.write(frame)
            time.sleep(1.0 / self.fps)

class DriverManager:
    """Manages WebDriver instances with video recording"""

    def __init__(self, browser='chrome', headless=False, record_video=True):
        self.browser = browser.lower()
        self.headless = headless
        self.record_video = record_video
        self.driver = None
        self.video_recorder = None

    def get_driver(self):
        """Create and configure WebDriver"""
        if self.browser == 'chrome':
            return self._create_chrome_driver()
        elif self.browser == 'firefox':
            return self._create_firefox_driver()
        else:
            raise ValueError(f"Unsupported browser: {self.browser}")

    def _create_chrome_driver(self):
        """Create Chrome WebDriver"""
        options = ChromeOptions()

        if self.headless:
            options.add_argument('--headless')

        options.add_argument('--no-sandbox')
        options.add_argument('--disable-dev-shm-usage')
        options.add_argument('--disable-gpu')
        options.add_argument('--window-size=1920,1080')
        options.add_argument('--disable-blink-features=AutomationControlled')
        options.add_experimental_option("excludeSwitches", ["enable-automation"])
        options.add_experimental_option('useAutomationExtension', False)

        service = ChromeService(ChromeDriverManager().install())
        driver = webdriver.Chrome(service=service, options=options)
        driver.execute_script("Object.defineProperty(navigator, 'webdriver', {get: () => undefined})")

        return driver

    def _create_firefox_driver(self):
        """Create Firefox WebDriver"""
        options = FirefoxOptions()

        if self.headless:
            options.add_argument('--headless')

        options.add_argument('--width=1920')
        options.add_argument('--height=1080')

        service = FirefoxService(GeckoDriverManager().install())
        return webdriver.Firefox(service=service, options=options)

    def start_session(self, test_name):
        """Start WebDriver session with optional video recording"""
        self.driver = self.get_driver()

        if self.record_video and not self.headless:
            recording_path = f"recordings/{test_name}_{int(time.time())}.mp4"
            os.makedirs(os.path.dirname(recording_path), exist_ok=True)
            self.video_recorder = VideoRecorder(recording_path)
            self.video_recorder.start_recording()

        return self.driver

    def end_session(self):
        """End WebDriver session and stop recording"""
        if self.video_recorder:
            self.video_recorder.stop_recording()

        if self.driver:
            self.driver.quit()

# Global driver manager instance
driver_manager = None

def get_driver_manager(browser='chrome', headless=False, record_video=True):
    """Get driver manager instance"""
    global driver_manager
    if not driver_manager:
        driver_manager = DriverManager(browser, headless, record_video)
    return driver_manager
```

### Step 3: Create page object models

Create `tests/e2e/pages/base_page.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/pages/base_page.py
# version: 1.0.0
# guid: b2c3d4e5-f6g7-8901-2345-6789abcdef01

import time
import os
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.action_chains import ActionChains
from selenium.common.exceptions import TimeoutException, NoSuchElementException

class BasePage:
    """Base page class with common functionality"""

    def __init__(self, driver):
        self.driver = driver
        self.wait = WebDriverWait(driver, 10)
        self.base_url = os.getenv('SUBTITLE_MANAGER_URL', 'http://localhost:8080')

    def navigate_to(self, path=''):
        """Navigate to specified path"""
        url = f"{self.base_url}{path}"
        self.driver.get(url)
        self.wait_for_page_load()

    def wait_for_element(self, locator, timeout=10):
        """Wait for element to be present"""
        wait = WebDriverWait(self.driver, timeout)
        return wait.until(EC.presence_of_element_located(locator))

    def wait_for_clickable(self, locator, timeout=10):
        """Wait for element to be clickable"""
        wait = WebDriverWait(self.driver, timeout)
        return wait.until(EC.element_to_be_clickable(locator))

    def wait_for_visible(self, locator, timeout=10):
        """Wait for element to be visible"""
        wait = WebDriverWait(self.driver, timeout)
        return wait.until(EC.visibility_of_element_located(locator))

    def wait_for_page_load(self):
        """Wait for page to fully load"""
        self.wait.until(
            lambda driver: driver.execute_script("return document.readyState") == "complete"
        )

    def click_element(self, locator):
        """Click element with wait"""
        element = self.wait_for_clickable(locator)
        element.click()

    def type_text(self, locator, text):
        """Type text into element"""
        element = self.wait_for_element(locator)
        element.clear()
        element.send_keys(text)

    def get_text(self, locator):
        """Get text from element"""
        element = self.wait_for_element(locator)
        return element.text

    def take_screenshot(self, filename):
        """Take screenshot"""
        screenshot_path = f"screenshots/{filename}"
        os.makedirs(os.path.dirname(screenshot_path), exist_ok=True)
        self.driver.save_screenshot(screenshot_path)
        return screenshot_path

    def scroll_to_element(self, locator):
        """Scroll element into view"""
        element = self.wait_for_element(locator)
        self.driver.execute_script("arguments[0].scrollIntoView(true);", element)

    def hover_over_element(self, locator):
        """Hover over element"""
        element = self.wait_for_element(locator)
        ActionChains(self.driver).move_to_element(element).perform()

    def is_element_present(self, locator):
        """Check if element is present"""
        try:
            self.driver.find_element(*locator)
            return True
        except NoSuchElementException:
            return False

    def wait_for_url_change(self, current_url, timeout=10):
        """Wait for URL to change"""
        wait = WebDriverWait(self.driver, timeout)
        wait.until(lambda driver: driver.current_url != current_url)
```

Create `tests/e2e/pages/login_page.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/pages/login_page.py
# version: 1.0.0
# guid: c3d4e5f6-g7h8-9012-3456-789abcdef012

from selenium.webdriver.common.by import By
from .base_page import BasePage

class LoginPage(BasePage):
    """Login page object model"""

    # Locators
    USERNAME_INPUT = (By.ID, "username")
    PASSWORD_INPUT = (By.ID, "password")
    LOGIN_BUTTON = (By.XPATH, "//button[contains(text(), 'Login')]")
    ERROR_MESSAGE = (By.CLASS_NAME, "alert-danger")
    LOGOUT_BUTTON = (By.XPATH, "//a[contains(text(), 'Logout')]")

    def navigate_to_login(self):
        """Navigate to login page"""
        self.navigate_to('/login')

    def login(self, username, password):
        """Perform login"""
        self.type_text(self.USERNAME_INPUT, username)
        self.type_text(self.PASSWORD_INPUT, password)
        self.click_element(self.LOGIN_BUTTON)
        self.wait_for_page_load()

    def logout(self):
        """Perform logout"""
        self.click_element(self.LOGOUT_BUTTON)
        self.wait_for_page_load()

    def get_error_message(self):
        """Get login error message"""
        try:
            return self.get_text(self.ERROR_MESSAGE)
        except:
            return None

    def is_logged_in(self):
        """Check if user is logged in"""
        return self.is_element_present(self.LOGOUT_BUTTON)
```

### Step 4: Create comprehensive test suites

Create `tests/e2e/tests/test_authentication.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/tests/test_authentication.py
# version: 1.0.0
# guid: d4e5f6g7-h8i9-0123-4567-89abcdef0123

import pytest
import allure
from pages.login_page import LoginPage
from pages.dashboard_page import DashboardPage
from utils.test_data import TestData

@allure.feature("Authentication")
class TestAuthentication:
    """Authentication workflow tests"""

    @pytest.fixture(autouse=True)
    def setup(self, driver):
        """Set up test"""
        self.login_page = LoginPage(driver)
        self.dashboard_page = DashboardPage(driver)
        self.test_data = TestData()

    @allure.story("Valid Login")
    @pytest.mark.smoke
    def test_valid_login(self, driver):
        """Test successful login with valid credentials"""
        # Arrange
        self.login_page.navigate_to_login()

        # Act
        self.login_page.login(
            self.test_data.valid_username,
            self.test_data.valid_password
        )

        # Assert
        assert self.login_page.is_logged_in(), "User should be logged in"
        assert self.dashboard_page.is_dashboard_loaded(), "Dashboard should load"

    @allure.story("Invalid Login")
    @pytest.mark.negative
    def test_invalid_login(self, driver):
        """Test login with invalid credentials"""
        # Arrange
        self.login_page.navigate_to_login()

        # Act
        self.login_page.login("invalid_user", "invalid_password")

        # Assert
        error_message = self.login_page.get_error_message()
        assert error_message is not None, "Error message should be displayed"
        assert "invalid" in error_message.lower(), "Error should mention invalid credentials"
        assert not self.login_page.is_logged_in(), "User should not be logged in"

    @allure.story("Logout")
    @pytest.mark.smoke
    def test_logout(self, driver):
        """Test logout functionality"""
        # Arrange - Login first
        self.login_page.navigate_to_login()
        self.login_page.login(
            self.test_data.valid_username,
            self.test_data.valid_password
        )

        # Act
        self.login_page.logout()

        # Assert
        assert not self.login_page.is_logged_in(), "User should be logged out"
        assert "login" in driver.current_url.lower(), "Should redirect to login page"

    @allure.story("Session Management")
    @pytest.mark.regression
    def test_session_persistence(self, driver):
        """Test session persistence across page navigation"""
        # Arrange
        self.login_page.navigate_to_login()
        self.login_page.login(
            self.test_data.valid_username,
            self.test_data.valid_password
        )

        # Act - Navigate to different pages
        self.dashboard_page.navigate_to('/media')
        self.dashboard_page.navigate_to('/settings')
        self.dashboard_page.navigate_to('/')

        # Assert
        assert self.login_page.is_logged_in(), "Session should persist across navigation"
```

Create `tests/e2e/tests/test_media_management.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/tests/test_media_management.py
# version: 1.0.0
# guid: e5f6g7h8-i9j0-1234-5678-9abcdef01234

import pytest
import allure
import time
from pages.login_page import LoginPage
from pages.media_page import MediaPage
from utils.test_data import TestData

@allure.feature("Media Management")
class TestMediaManagement:
    """Media management workflow tests"""

    @pytest.fixture(autouse=True)
    def setup(self, driver):
        """Set up test with authentication"""
        self.login_page = LoginPage(driver)
        self.media_page = MediaPage(driver)
        self.test_data = TestData()

        # Login before each test
        self.login_page.navigate_to_login()
        self.login_page.login(
            self.test_data.valid_username,
            self.test_data.valid_password
        )

    @allure.story("Add Media Library")
    @pytest.mark.smoke
    def test_add_media_library(self, driver):
        """Test adding a new media library"""
        # Arrange
        self.media_page.navigate_to('/media')

        # Act
        self.media_page.click_add_library()
        library_data = {
            'name': 'Test Movies',
            'path': '/media/movies',
            'type': 'movies'
        }
        self.media_page.fill_library_form(library_data)
        self.media_page.save_library()

        # Assert
        assert self.media_page.is_library_present('Test Movies'), "Library should be added"
        assert self.media_page.get_library_path('Test Movies') == '/media/movies'

    @allure.story("Scan Media Library")
    @pytest.mark.regression
    def test_scan_media_library(self, driver):
        """Test scanning media library for content"""
        # Arrange
        self.media_page.navigate_to('/media')

        # Act
        self.media_page.scan_library('Test Movies')

        # Assert
        self.media_page.wait_for_scan_complete()
        scan_status = self.media_page.get_scan_status('Test Movies')
        assert 'complete' in scan_status.lower(), "Scan should complete successfully"

    @allure.story("Remove Media Library")
    @pytest.mark.destructive
    def test_remove_media_library(self, driver):
        """Test removing a media library"""
        # Arrange
        self.media_page.navigate_to('/media')
        initial_count = self.media_page.get_library_count()

        # Act
        self.media_page.remove_library('Test Movies')
        self.media_page.confirm_removal()

        # Assert
        final_count = self.media_page.get_library_count()
        assert final_count == initial_count - 1, "Library count should decrease"
        assert not self.media_page.is_library_present('Test Movies'), "Library should be removed"
```

### Step 5: Create visual regression testing

Create `tests/e2e/tests/test_visual_regression.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/tests/test_visual_regression.py
# version: 1.0.0
# guid: f6g7h8i9-j0k1-2345-6789-abcdef012345

import pytest
import allure
import cv2
import numpy as np
from PIL import Image, ImageChops
from pages.login_page import LoginPage
from pages.dashboard_page import DashboardPage
from utils.test_data import TestData

@allure.feature("Visual Regression")
class TestVisualRegression:
    """Visual regression testing suite"""

    @pytest.fixture(autouse=True)
    def setup(self, driver):
        """Set up test"""
        self.login_page = LoginPage(driver)
        self.dashboard_page = DashboardPage(driver)
        self.test_data = TestData()

        # Login
        self.login_page.navigate_to_login()
        self.login_page.login(
            self.test_data.valid_username,
            self.test_data.valid_password
        )

    def compare_screenshots(self, baseline_path, current_path, threshold=0.95):
        """Compare two screenshots for visual differences"""
        baseline = Image.open(baseline_path)
        current = Image.open(current_path)

        # Resize images to same size if needed
        if baseline.size != current.size:
            current = current.resize(baseline.size)

        # Calculate difference
        diff = ImageChops.difference(baseline, current)

        # Convert to numpy for analysis
        diff_array = np.array(diff)

        # Calculate similarity
        total_pixels = diff_array.size
        different_pixels = np.count_nonzero(diff_array)
        similarity = 1 - (different_pixels / total_pixels)

        return similarity >= threshold, similarity

    @allure.story("Dashboard Layout")
    @pytest.mark.visual
    def test_dashboard_visual_consistency(self, driver):
        """Test dashboard visual consistency"""
        # Arrange
        self.dashboard_page.navigate_to('/')

        # Act
        current_screenshot = self.dashboard_page.take_screenshot("dashboard_current.png")
        baseline_screenshot = "screenshots/baseline/dashboard_baseline.png"

        # Assert
        if os.path.exists(baseline_screenshot):
            is_similar, similarity = self.compare_screenshots(
                baseline_screenshot, current_screenshot
            )
            assert is_similar, f"Dashboard layout changed. Similarity: {similarity:.2%}"
        else:
            # Create baseline if it doesn't exist
            os.makedirs(os.path.dirname(baseline_screenshot), exist_ok=True)
            shutil.copy(current_screenshot, baseline_screenshot)

    @allure.story("Settings Page Layout")
    @pytest.mark.visual
    def test_settings_visual_consistency(self, driver):
        """Test settings page visual consistency"""
        # Arrange
        self.dashboard_page.navigate_to('/settings')

        # Act
        current_screenshot = self.dashboard_page.take_screenshot("settings_current.png")
        baseline_screenshot = "screenshots/baseline/settings_baseline.png"

        # Assert
        if os.path.exists(baseline_screenshot):
            is_similar, similarity = self.compare_screenshots(
                baseline_screenshot, current_screenshot
            )
            assert is_similar, f"Settings layout changed. Similarity: {similarity:.2%}"
```

### Step 6: Create performance testing

Create `tests/e2e/tests/test_performance.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/tests/test_performance.py
# version: 1.0.0
# guid: g7h8i9j0-k1l2-3456-7890-bcdef0123456

import pytest
import allure
import time
import requests
from selenium.webdriver.common.action_chains import ActionChains
from pages.login_page import LoginPage
from pages.dashboard_page import DashboardPage
from utils.test_data import TestData

@allure.feature("Performance")
class TestPerformance:
    """Performance testing suite"""

    @pytest.fixture(autouse=True)
    def setup(self, driver):
        """Set up test"""
        self.login_page = LoginPage(driver)
        self.dashboard_page = DashboardPage(driver)
        self.test_data = TestData()

    def measure_page_load_time(self, url):
        """Measure page load time"""
        start_time = time.time()
        self.dashboard_page.navigate_to(url)
        end_time = time.time()
        return end_time - start_time

    @allure.story("Page Load Performance")
    @pytest.mark.performance
    def test_dashboard_load_time(self, driver):
        """Test dashboard load time performance"""
        # Arrange
        self.login_page.navigate_to_login()
        self.login_page.login(
            self.test_data.valid_username,
            self.test_data.valid_password
        )

        # Act
        load_time = self.measure_page_load_time('/')

        # Assert
        assert load_time < 3.0, f"Dashboard should load in under 3 seconds, took {load_time:.2f}s"

    @allure.story("API Response Performance")
    @pytest.mark.performance
    def test_api_response_times(self, driver):
        """Test API endpoint response times"""
        # Test various API endpoints
        endpoints = [
            '/api/health',
            '/api/system/info',
            '/api/media/libraries',
            '/api/providers',
        ]

        for endpoint in endpoints:
            start_time = time.time()
            response = requests.get(f"{self.dashboard_page.base_url}{endpoint}")
            end_time = time.time()

            response_time = end_time - start_time
            assert response_time < 2.0, f"{endpoint} should respond in under 2 seconds"
            assert response.status_code == 200, f"{endpoint} should return 200 OK"

    @allure.story("UI Interaction Performance")
    @pytest.mark.performance
    def test_navigation_performance(self, driver):
        """Test navigation performance between pages"""
        # Login
        self.login_page.navigate_to_login()
        self.login_page.login(
            self.test_data.valid_username,
            self.test_data.valid_password
        )

        # Test navigation between pages
        pages = ['/', '/media', '/wanted', '/history', '/settings']

        for page in pages:
            start_time = time.time()
            self.dashboard_page.navigate_to(page)
            end_time = time.time()

            navigation_time = end_time - start_time
            assert navigation_time < 2.0, f"Navigation to {page} should take under 2 seconds"
```

### Step 7: Create test configuration

Create `tests/e2e/pytest.ini`:

```ini
[pytest]
testpaths = tests
python_files = test_*.py
python_classes = Test*
python_functions = test_*
addopts =
    --html=reports/report.html
    --self-contained-html
    --tb=short
    --strict-markers
    --disable-warnings
markers =
    smoke: Critical functionality tests
    regression: Full regression test suite
    negative: Negative test cases
    visual: Visual regression tests
    performance: Performance and load tests
    destructive: Tests that modify data
```

Create `tests/e2e/conftest.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/conftest.py
# version: 1.0.0
# guid: h8i9j0k1-l2m3-4567-8901-cdef01234567

import pytest
import os
import time
from utils.driver_manager import get_driver_manager

@pytest.fixture(scope="session")
def browser():
    """Browser type fixture"""
    return os.getenv("BROWSER", "chrome")

@pytest.fixture(scope="session")
def headless():
    """Headless mode fixture"""
    return os.getenv("HEADLESS", "false").lower() == "true"

@pytest.fixture
def driver(browser, headless, request):
    """WebDriver fixture with video recording"""
    test_name = request.node.name

    driver_mgr = get_driver_manager(browser, headless, record_video=True)
    driver = driver_mgr.start_session(test_name)

    yield driver

    # Take screenshot on failure
    if hasattr(request.node, 'rep_call') and request.node.rep_call.failed:
        timestamp = int(time.time())
        screenshot_path = f"screenshots/failure_{test_name}_{timestamp}.png"
        driver.save_screenshot(screenshot_path)

    driver_mgr.end_session()

@pytest.hookimpl(tryfirst=True, hookwrapper=True)
def pytest_runtest_makereport(item, call):
    """Add test result to item for screenshot on failure"""
    outcome = yield
    rep = outcome.get_result()
    setattr(item, f"rep_{rep.when}", rep)
```

### Step 8: Create CI/CD integration

Create `tests/e2e/run_tests.sh`:

```bash
#!/bin/bash
# file: tests/e2e/run_tests.sh
# version: 1.0.0
# guid: i9j0k1l2-m3n4-5678-9012-def012345678

set -e

# Configuration
BROWSER=${BROWSER:-chrome}
HEADLESS=${HEADLESS:-true}
PARALLEL=${PARALLEL:-2}
SUBTITLE_MANAGER_URL=${SUBTITLE_MANAGER_URL:-http://localhost:8080}

echo "Starting E2E Test Suite"
echo "Browser: $BROWSER"
echo "Headless: $HEADLESS"
echo "Parallel: $PARALLEL"
echo "URL: $SUBTITLE_MANAGER_URL"

# Create directories
mkdir -p reports screenshots recordings

# Wait for application to be ready
echo "Waiting for application to be ready..."
timeout 60 bash -c 'until curl -s $SUBTITLE_MANAGER_URL/health; do sleep 1; done'

# Activate virtual environment
source venv/bin/activate

# Run different test suites
echo "Running smoke tests..."
pytest tests/ -m smoke --html=reports/smoke_report.html --self-contained-html -v

echo "Running regression tests..."
pytest tests/ -m regression --html=reports/regression_report.html --self-contained-html -v

echo "Running performance tests..."
pytest tests/ -m performance --html=reports/performance_report.html --self-contained-html -v

echo "Running visual regression tests..."
pytest tests/ -m visual --html=reports/visual_report.html --self-contained-html -v

echo "Running full test suite with parallel execution..."
pytest tests/ -n $PARALLEL --html=reports/full_report.html --self-contained-html -v

echo "E2E Test Suite completed!"
echo "Reports generated in reports/ directory"
echo "Screenshots saved in screenshots/ directory"
echo "Video recordings saved in recordings/ directory"
```

### Step 9: Add test data management

Create `tests/e2e/utils/test_data.py`:

```python
#!/usr/bin/env python3
# file: tests/e2e/utils/test_data.py
# version: 1.0.0
# guid: j0k1l2m3-n4o5-6789-0123-ef0123456789

import os
import json
import random
import string

class TestData:
    """Test data management"""

    def __init__(self):
        self.load_config()

    def load_config(self):
        """Load test configuration"""
        config_file = os.getenv('TEST_CONFIG_FILE', 'fixtures/test_config.json')

        if os.path.exists(config_file):
            with open(config_file, 'r') as f:
                config = json.load(f)
        else:
            config = self.default_config()

        self.valid_username = config.get('valid_username', 'admin')
        self.valid_password = config.get('valid_password', 'admin')
        self.base_url = config.get('base_url', 'http://localhost:8080')
        self.test_media_path = config.get('test_media_path', '/tmp/test_media')

    def default_config(self):
        """Default test configuration"""
        return {
            'valid_username': 'admin',
            'valid_password': 'admin',
            'base_url': 'http://localhost:8080',
            'test_media_path': '/tmp/test_media'
        }

    def generate_random_string(self, length=10):
        """Generate random string"""
        letters = string.ascii_lowercase
        return ''.join(random.choice(letters) for i in range(length))

    def get_test_library_data(self):
        """Get test library data"""
        return {
            'name': f'Test Library {self.generate_random_string(5)}',
            'path': f'{self.test_media_path}/{self.generate_random_string(8)}',
            'type': random.choice(['movies', 'tv_shows']),
            'language': 'en'
        }

    def get_test_provider_data(self):
        """Get test provider data"""
        return {
            'name': f'Test Provider {self.generate_random_string(5)}',
            'type': 'opensubtitles',
            'enabled': True,
            'config': {
                'username': 'test_user',
                'password': 'test_password',
                'rate_limit': 10
            }
        }
```

### Step 10: Run and validate tests

```bash
# Make scripts executable
chmod +x tests/e2e/run_tests.sh

# Set up test environment
cd tests/e2e
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt

# Create test configuration
mkdir -p fixtures
cat > fixtures/test_config.json << 'EOF'
{
    "valid_username": "admin",
    "valid_password": "admin",
    "base_url": "http://localhost:8080",
    "test_media_path": "/tmp/test_media"
}
EOF

# Run tests
./run_tests.sh
```

## 📚 Required Documentation

### Coding Instructions Reference

**CRITICAL**: Follow these instructions precisely:

```markdown
From .github/instructions/general-coding.instructions.md:

## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS

**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction
of any kind.**

## Script Language Preference

**MANDATORY RULE: Prefer Python for scripts unless they are incredibly simple.**

Use Python for:

- API interactions (GitHub, REST APIs, etc.)
- JSON/YAML processing
- File manipulation beyond simple copying
- Error handling and logging
- Data parsing or transformation
- More than 20-30 lines of logic
```

## 🧪 Testing Requirements

### Test Coverage Requirements

- [ ] Authentication workflows (login, logout, sessions)
- [ ] Media management (add, scan, remove libraries)
- [ ] Subtitle operations (search, download, extract)
- [ ] Provider configuration and testing
- [ ] Settings management and persistence
- [ ] System health and monitoring
- [ ] Error handling and edge cases

### Video Recording Requirements

- [ ] All test sessions recorded when not in headless mode
- [ ] Recordings saved with test name and timestamp
- [ ] Failed test recordings highlighted for debugging
- [ ] Performance test recordings for analysis

## 🎯 Success Metrics

- [ ] All test suites execute successfully
- [ ] Video recordings generated for all tests
- [ ] Test reports generated with screenshots
- [ ] Performance benchmarks established
- [ ] CI/CD integration working
- [ ] Visual regression tests detecting layout changes
- [ ] Parallel test execution reducing overall time
- [ ] Test data management system functional

## 🚨 Common Pitfalls

1. **WebDriver Setup**: Ensure correct driver versions for browsers
2. **Test Dependencies**: Tests should be independent and idempotent
3. **Video Recording**: May impact performance in headless environments
4. **File Permissions**: Ensure test directories are writable
5. **Network Timeouts**: Handle network delays and API timeouts
6. **Browser Compatibility**: Test across different browsers and versions
7. **Test Data Cleanup**: Clean up test data after destructive tests

## 📖 Additional Resources

- [Selenium Documentation](https://selenium-python.readthedocs.io/)
- [pytest Documentation](https://docs.pytest.org/)
- [Allure Reporting](https://docs.qameta.io/allure/)
- [WebDriver Manager](https://github.com/SergeyPirogov/webdriver_manager)

## 🔄 Related Tasks

- **TASK-02-001**: UI Layout fixes (tests will validate these changes)
- **TASK-02-003**: Provider system improvements (will need updated tests)
- **TASK-04-001**: Authentication improvements (tests validate auth flows)

## 📝 Notes for AI Agent

- Focus on comprehensive test coverage for all user workflows
- Ensure tests are maintainable and not brittle
- Use page object pattern for maintainable test code
- Implement proper error handling and debugging aids
- Consider test execution time when designing test suites
- Standard Python testing tools - no special frameworks required
- If video recording fails, tests should still run successfully
- All tests should pass before marking this task complete
