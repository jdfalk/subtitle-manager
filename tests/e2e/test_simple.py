#!/usr/bin/env python3
# file: tests/e2e/test_simple.py
# version: 1.0.0
# guid: a1b2c3d4-e5f6-a7b8-c9d0-e1f2a3b4c5d6

"""
Simple test to verify our settings navigation fix works.
Uses basic selenium without the full test framework.
"""

import sys
import time
from pathlib import Path

# Add project root to path
PROJECT_ROOT = Path(__file__).parent.parent.parent
sys.path.insert(0, str(PROJECT_ROOT))


def test_settings_navigation_fix():
    """Test that our settings navigation fix works."""
    try:
        from selenium import webdriver
        from selenium.webdriver.chrome.options import Options
        from selenium.webdriver.common.by import By
        from selenium.webdriver.support import expected_conditions as EC
        from selenium.webdriver.support.ui import WebDriverWait
        from webdriver_manager.chrome import ChromeDriverManager
    except ImportError as e:
        print(f"Selenium not available: {e}")
        print("Install with: pip install selenium webdriver-manager")
        return False

    # Set up Chrome driver
    options = Options()
    options.add_argument("--headless")  # Run in background
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-dev-shm-usage")
    options.add_argument("--window-size=1920,1080")

    driver = None
    try:
        driver = webdriver.Chrome(ChromeDriverManager().install(), options=options)

        # Navigate to app
        app_url = "http://localhost:5173"
        print(f"Navigating to {app_url}")
        driver.get(app_url)

        # Wait for page to load
        WebDriverWait(driver, 10).until(
            lambda d: d.execute_script("return document.readyState") == "complete"
        )

        print("✓ Page loaded successfully")

        # Look for settings link
        settings_selectors = [
            "//a[contains(@href, '/settings')]",
            "//*[contains(text(), 'Settings')]",
            "//nav//a[contains(text(), 'Settings')]",
        ]

        settings_element = None
        for selector in settings_selectors:
            elements = driver.find_elements(By.XPATH, selector)
            if elements:
                settings_element = elements[0]
                break

        if not settings_element:
            print("✗ No settings navigation found")
            return False

        print("✓ Found settings navigation")

        # Click settings
        settings_element.click()
        time.sleep(2)  # Wait for navigation

        # Check URL
        current_url = driver.current_url
        print(f"Current URL after clicking Settings: {current_url}")

        if "/settings" not in current_url:
            print("✗ Settings navigation didn't work")
            return False

        print("✓ Successfully navigated to settings")

        # Check for tabbed interface (not 3-card overview)
        tab_elements = driver.find_elements(
            By.CSS_SELECTOR, "[role='tab'], .MuiTab-root"
        )
        card_elements = driver.find_elements(By.CSS_SELECTOR, ".MuiCard-root")

        print(f"Found {len(tab_elements)} tab elements")
        print(f"Found {len(card_elements)} card elements")

        if len(tab_elements) >= 3:
            print(
                "✓ PASS: Found tabbed interface - settings navigation fix is working!"
            )
            return True
        elif len(card_elements) == 3:
            print("✗ FAIL: Found 3-card overview page - routing fix not working")
            return False
        else:
            print("? UNCLEAR: Could not determine interface type")
            return False

    except Exception as e:
        print(f"✗ Test failed with error: {e}")
        return False

    finally:
        if driver:
            driver.quit()


if __name__ == "__main__":
    print("Testing settings navigation fix...")
    success = test_settings_navigation_fix()
    sys.exit(0 if success else 1)
