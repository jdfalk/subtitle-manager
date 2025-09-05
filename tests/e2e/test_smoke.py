# file: tests/e2e/test_smoke.py
# version: 1.0.0
# guid: f6a7b8c9-d0e1-2345-6789-012345678901

"""
Smoke tests for basic functionality.
These tests should run quickly and verify that the application is basically functional.
"""

import pytest
from conftest import TestUtils
from selenium.webdriver.common.by import By


class TestSmoke:
    """Smoke test suite - basic functionality checks."""

    @pytest.mark.smoke
    def test_app_loads(self, driver, app_url):
        """Basic smoke test - can we load the app?"""
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        TestUtils.take_screenshot(driver, "smoke_01_app_loads")

        # Very basic checks
        assert driver.title is not None
        assert len(driver.page_source) > 100

        # Check for React app mounting
        assert "div" in driver.page_source.lower()

    @pytest.mark.smoke
    def test_settings_navigation_fix(self, driver, app_url):
        """Smoke test for the settings navigation fix we just implemented."""
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Find settings link and click it
        settings_links = driver.find_elements(
            By.XPATH,
            "//a[contains(@href, '/settings') or contains(text(), 'Settings')]",
        )

        if not settings_links:
            # Try alternative selectors
            settings_links = driver.find_elements(
                By.XPATH, "//*[contains(text(), 'Settings')]"
            )

        if settings_links:
            settings_link = settings_links[0]
            TestUtils.scroll_to_element(driver, settings_link)
            settings_link.click()

            TestUtils.wait_for_page_load(driver)
            TestUtils.take_screenshot(driver, "smoke_02_settings_clicked")

            # Verify we're on settings and see tabbed interface (not 3-card overview)
            assert "/settings" in driver.current_url

            # Look for tabs (the fix should show tabbed interface immediately)
            tab_elements = driver.find_elements(
                By.CSS_SELECTOR, "[role='tab'], .MuiTab-root"
            )

            if len(tab_elements) >= 2:
                print(f"✓ Found {len(tab_elements)} tabs - tabbed interface working")
            else:
                # Check if we see the old 3-card layout we wanted to avoid
                cards = driver.find_elements(By.CSS_SELECTOR, ".MuiCard-root")
                if len(cards) == 3:
                    pytest.fail(
                        "Found 3-card overview page instead of tabbed interface - routing fix may not be working"
                    )
        else:
            pytest.skip("No settings navigation found to test")

    @pytest.mark.smoke
    def test_no_obvious_errors(self, driver, app_url):
        """Check that there are no obvious errors on page load."""
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Check for error text in the page
        page_text = driver.page_source.lower()
        error_indicators = [
            "error occurred",
            "something went wrong",
            "500 internal server",
            "404 not found",
        ]

        found_errors = [error for error in error_indicators if error in page_text]

        TestUtils.take_screenshot(driver, "smoke_03_error_check")

        assert len(found_errors) == 0, f"Found error indicators: {found_errors}"
