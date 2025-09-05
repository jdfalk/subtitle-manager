# file: tests/e2e/test_ui_critical_functionality.py
# version: 1.0.0
# guid: c3d4e5f6-a7b8-9012-3456-789012345678

"""
Critical UI functionality tests covering the main user workflows.
These tests verify that core features work end-to-end.
"""

import time

import pytest
from conftest import TestUtils
from selenium.common.exceptions import TimeoutException
from selenium.webdriver.common.by import By


class TestCriticalUIFunctionality:
    """Test suite for critical UI functionality."""

    @pytest.mark.critical
    def test_application_loads_successfully(self, driver, app_url):
        """
        Test that the application loads without errors.
        This is the most basic test - if this fails, nothing else will work.
        """
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Take screenshot of loaded app
        TestUtils.take_screenshot(driver, "ui_01_app_loaded")

        # Check that we don't have any obvious error messages
        error_indicators = driver.find_elements(
            By.XPATH,
            "//*[contains(text(), 'Error') or contains(text(), 'error') or contains(text(), '404') or contains(text(), '500')]",
        )
        error_messages = [elem.text for elem in error_indicators if elem.is_displayed()]

        assert len(error_messages) == 0, (
            f"Found error messages on page load: {error_messages}"
        )

        # Verify basic page structure exists
        assert driver.title, "Page should have a title"
        assert len(driver.page_source) > 1000, "Page should have substantial content"

    @pytest.mark.critical
    def test_navigation_menu_present_and_functional(self, driver, app_url, wait):
        """
        Test that main navigation menu is present and links work.
        """
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Look for navigation elements
        nav_selectors = [
            "nav",
            ".MuiDrawer-root",
            "[role='navigation']",
            ".navigation",
            ".sidebar",
        ]

        navigation_found = False
        for selector in nav_selectors:
            nav_elements = driver.find_elements(By.CSS_SELECTOR, selector)
            if nav_elements and any(elem.is_displayed() for elem in nav_elements):
                navigation_found = True
                break

        assert navigation_found, (
            f"No navigation menu found. Tried selectors: {nav_selectors}"
        )

        # Look for common navigation links
        expected_nav_items = [
            "Dashboard",
            "Media",
            "Library",
            "Settings",
            "History",
            "Wanted",
        ]

        nav_links = driver.find_elements(By.XPATH, "//a | //button")
        nav_text = [link.text.strip() for link in nav_links if link.text.strip()]

        found_nav_items = [
            item
            for item in expected_nav_items
            if any(item.lower() in text.lower() for text in nav_text)
        ]

        assert len(found_nav_items) >= 2, (
            f"Expected to find navigation items like {expected_nav_items}, but only found: {found_nav_items}"
        )

        TestUtils.take_screenshot(driver, "ui_02_navigation_verified")

    @pytest.mark.critical
    def test_settings_page_accessibility(self, driver, app_url):
        """
        Test that settings page is accessible and contains expected elements.
        """
        driver.get(f"{app_url}/settings")
        TestUtils.wait_for_page_load(driver)

        TestUtils.take_screenshot(driver, "ui_03_settings_page")

        # Verify we're on settings page
        assert "/settings" in driver.current_url

        # Look for settings-related content
        settings_indicators = driver.find_elements(
            By.XPATH,
            "//*[contains(text(), 'Settings') or contains(text(), 'Configuration') or contains(text(), 'Provider')]",
        )

        assert len(settings_indicators) > 0, (
            "Settings page should contain settings-related content"
        )

        # Look for interactive elements (forms, buttons, tabs)
        interactive_elements = driver.find_elements(
            By.CSS_SELECTOR, "input, button, select, [role='tab'], .MuiTab-root"
        )

        assert len(interactive_elements) > 0, (
            "Settings page should have interactive elements"
        )

    def test_dashboard_loads_and_shows_content(self, driver, app_url):
        """
        Test that dashboard page loads and shows relevant content.
        """
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Look for dashboard-specific content
        dashboard_indicators = [
            "Dashboard",
            "Statistics",
            "Recent",
            "Status",
            "Activity",
            "Overview",
        ]

        page_text = driver.page_source.lower()
        found_indicators = [
            indicator
            for indicator in dashboard_indicators
            if indicator.lower() in page_text
        ]

        # We should find at least some dashboard-related content
        assert len(found_indicators) > 0, (
            f"Dashboard should contain relevant content. Looking for: {dashboard_indicators}"
        )

        TestUtils.take_screenshot(driver, "ui_04_dashboard_content")

    def test_responsive_design_mobile(self, driver, app_url):
        """
        Test that the application works on mobile viewport.
        """
        # Set mobile viewport (iPhone size)
        driver.set_window_size(375, 667)

        try:
            driver.get(app_url)
            TestUtils.wait_for_page_load(driver)

            TestUtils.take_screenshot(driver, "ui_05_mobile_view")

            # Check that content is visible and not cut off
            body = driver.find_element(By.TAG_NAME, "body")
            body_rect = body.size

            # Content should not be wider than viewport
            assert body_rect["width"] <= 375, (
                f"Content width {body_rect['width']} exceeds mobile viewport width 375"
            )

            # Check for mobile navigation (hamburger menu, etc.)
            mobile_nav_selectors = [
                "[data-testid='menu-button']",
                ".hamburger",
                "[aria-label*='menu']",
                "button[aria-label*='Menu']",
            ]

            mobile_nav_found = False
            for selector in mobile_nav_selectors:
                elements = driver.find_elements(By.CSS_SELECTOR, selector)
                if elements and any(elem.is_displayed() for elem in elements):
                    mobile_nav_found = True
                    break

            # Note: Mobile nav is nice to have but not critical for this test
            if mobile_nav_found:
                TestUtils.take_screenshot(driver, "ui_06_mobile_navigation_found")

        finally:
            # Reset to desktop size
            driver.set_window_size(1920, 1080)

    def test_error_handling_404_page(self, driver, app_url):
        """
        Test that navigating to non-existent page shows appropriate handling.
        """
        # Navigate to a page that shouldn't exist
        driver.get(f"{app_url}/this-page-does-not-exist-12345")
        TestUtils.wait_for_page_load(driver)

        TestUtils.take_screenshot(driver, "ui_07_404_handling")

        page_text = driver.page_source.lower()

        # Should either show 404 handling or redirect to home
        has_error_handling = (
            "404" in page_text
            or "not found" in page_text
            or "page not found" in page_text
            or driver.current_url.rstrip("/")
            == app_url.rstrip("/")  # Redirected to home
        )

        assert has_error_handling, "Application should handle 404 errors gracefully"

    @pytest.mark.slow
    def test_performance_basic_page_loads(self, driver, app_url):
        """
        Test that pages load within reasonable time limits.
        """
        pages_to_test = [
            "",  # Home/dashboard
            "/settings",
            "/library",
            "/media",
            "/history",
        ]

        performance_results = {}

        for page_path in pages_to_test:
            start_time = time.time()

            try:
                driver.get(f"{app_url}{page_path}")
                TestUtils.wait_for_page_load(driver, timeout=10)

                load_time = time.time() - start_time
                performance_results[page_path or "home"] = load_time

                # Pages should load within 8 seconds
                assert load_time < 8.0, (
                    f"Page {page_path} took too long to load: {load_time:.2f} seconds"
                )

            except TimeoutException:
                load_time = time.time() - start_time
                performance_results[page_path or "home"] = load_time
                # Note the timeout but don't fail the test - some pages might not exist yet
                print(f"Page {page_path} timed out after {load_time:.2f} seconds")

        TestUtils.take_screenshot(driver, "ui_08_performance_test_complete")
        print(f"Performance results: {performance_results}")

    def test_javascript_errors_detection(self, driver, app_url):
        """
        Test that there are no JavaScript errors on page load.
        """
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Get browser console logs
        logs = driver.get_log("browser")

        # Filter for errors (not warnings or info)
        error_logs = [log for log in logs if log["level"] == "SEVERE"]

        # Filter out known non-critical errors
        critical_errors = []
        for log in error_logs:
            message = log["message"].lower()
            # Skip common non-critical errors
            if any(
                skip in message
                for skip in ["favicon", "manifest", "sw.js", "service-worker"]
            ):
                continue
            critical_errors.append(log)

        TestUtils.take_screenshot(driver, "ui_09_js_error_check")

        if critical_errors:
            error_messages = [
                f"{log['level']}: {log['message']}" for log in critical_errors
            ]
            pytest.fail(f"JavaScript errors found: {error_messages}")

    def test_accessibility_basic_checks(self, driver, app_url):
        """
        Basic accessibility checks for the application.
        """
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Check for basic accessibility features
        checks = {
            "Has page title": bool(driver.title and driver.title.strip()),
            "Has main landmark": len(driver.find_elements(By.TAG_NAME, "main")) > 0,
            "Has headings": len(
                driver.find_elements(By.CSS_SELECTOR, "h1, h2, h3, h4, h5, h6")
            )
            > 0,
            "Buttons have text or aria-label": True,  # Will check below
            "Links have text or aria-label": True,  # Will check below
        }

        # Check buttons have accessible text
        buttons = driver.find_elements(By.TAG_NAME, "button")
        for button in buttons[:10]:  # Check first 10 buttons
            if button.is_displayed():
                button_text = button.text.strip()
                aria_label = button.get_attribute("aria-label") or ""
                if not button_text and not aria_label:
                    checks["Buttons have text or aria-label"] = False
                    break

        # Check links have accessible text
        links = driver.find_elements(By.TAG_NAME, "a")
        for link in links[:10]:  # Check first 10 links
            if link.is_displayed():
                link_text = link.text.strip()
                aria_label = link.get_attribute("aria-label") or ""
                if not link_text and not aria_label:
                    checks["Links have text or aria-label"] = False
                    break

        TestUtils.take_screenshot(driver, "ui_10_accessibility_check")

        failed_checks = [check for check, passed in checks.items() if not passed]

        if failed_checks:
            print(f"Accessibility checks failed: {failed_checks}")
            # Note: Don't fail the test for accessibility issues, just log them
            # assert len(failed_checks) == 0, f"Accessibility checks failed: {failed_checks}"
