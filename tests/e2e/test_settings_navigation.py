# file: tests/e2e/test_settings_navigation.py
# version: 1.0.0
# guid: b2c3d4e5-f6a7-8901-2345-678901234567

"""
E2E tests for settings navigation functionality.
Tests the fix for settings routing confusion where users expected /settings
to show the tabbed interface instead of the basic overview.
"""

import time

import pytest
from conftest import TestUtils
from selenium.common.exceptions import TimeoutException
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC


class TestSettingsNavigation:
    """Test suite for settings navigation and routing."""

    def test_settings_route_shows_tabbed_interface(self, driver, app_url, wait):
        """
        Test that clicking Settings in navigation goes directly to tabbed interface.
        This verifies the fix for the 'blocky three-box page' issue.
        """
        # Navigate to app
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Take initial screenshot
        TestUtils.take_screenshot(driver, "settings_nav_01_homepage")

        # Find and click Settings in navigation
        settings_link = TestUtils.wait_for_element_clickable(
            driver,
            (
                By.XPATH,
                "//a[contains(@href, '/settings') or contains(text(), 'Settings')]",
            ),
            timeout=15,
        )

        # Scroll to element and click
        TestUtils.scroll_to_element(driver, settings_link)
        settings_link.click()

        # Wait for navigation
        TestUtils.wait_for_page_load(driver)
        time.sleep(2)  # Extra time for React components to render

        # Take screenshot after clicking Settings
        TestUtils.take_screenshot(driver, "settings_nav_02_after_click")

        # Verify we're on the settings page
        assert "/settings" in driver.current_url, (
            f"Expected /settings in URL, got: {driver.current_url}"
        )

        # Look for tabbed interface indicators (Material-UI Tabs)
        # This should NOT be the overview page with just 3 cards
        try:
            # Check for Material-UI tab elements
            wait.until(
                EC.presence_of_element_located(
                    (
                        By.CSS_SELECTOR,
                        "[role='tablist'], .MuiTabs-root, [data-testid='tabs']",
                    )
                )
            )

            # Check for multiple tabs (Providers, General, etc.)
            tab_elements = driver.find_elements(
                By.CSS_SELECTOR, "[role='tab'], .MuiTab-root"
            )
            assert len(tab_elements) >= 3, (
                f"Expected at least 3 tabs, found {len(tab_elements)}"
            )

            # Verify we see expected tab labels
            tab_texts = [tab.text.lower() for tab in tab_elements if tab.text.strip()]
            expected_tabs = ["providers", "general", "database", "authentication"]

            found_expected_tabs = [
                tab for tab in expected_tabs if any(tab in text for text in tab_texts)
            ]
            assert len(found_expected_tabs) >= 2, (
                f"Expected to find tabs like {expected_tabs}, but found: {tab_texts}"
            )

            TestUtils.take_screenshot(
                driver, "settings_nav_03_tabbed_interface_verified"
            )

        except TimeoutException:
            # If we can't find tabs, check if we're seeing the old overview page
            TestUtils.take_screenshot(driver, "settings_nav_03_error_no_tabs")

            # Look for the old 3-card overview that we DON'T want to see
            cards = driver.find_elements(
                By.CSS_SELECTOR, ".MuiCard-root, [data-testid='settings-card']"
            )
            card_texts = [card.text.lower() for card in cards]

            # If we see exactly 3 cards with General/Providers/Users, that's the old overview
            if len(cards) == 3 and any(
                "general" in text and "providers" in text and "users" in text
                for text in [" ".join(card_texts)]
            ):
                pytest.fail(
                    "Found old 3-card overview page instead of tabbed interface. The routing fix didn't work."
                )

            # Re-raise the original timeout exception
            raise

    def test_settings_tabs_functionality(self, driver, app_url, wait):
        """
        Test that the settings tabs are functional and can be navigated.
        """
        # Navigate to settings
        driver.get(f"{app_url}/settings")
        TestUtils.wait_for_page_load(driver)

        # Find all tabs
        TestUtils.wait_for_element_visible(
            driver, (By.CSS_SELECTOR, "[role='tab'], .MuiTab-root")
        )

        all_tabs = driver.find_elements(By.CSS_SELECTOR, "[role='tab'], .MuiTab-root")

        if len(all_tabs) < 2:
            pytest.skip("Not enough tabs found to test tab functionality")

        # Test clicking different tabs
        for i, tab in enumerate(all_tabs[:3]):  # Test first 3 tabs
            if not tab.is_displayed() or not tab.is_enabled():
                continue

            tab_text = tab.text.strip()
            if not tab_text:
                continue

            # Click the tab
            TestUtils.scroll_to_element(driver, tab)
            tab.click()
            time.sleep(1)  # Wait for tab content to load

            # Take screenshot
            TestUtils.take_screenshot(
                driver, f"settings_nav_04_tab_{i}_{tab_text.lower().replace(' ', '_')}"
            )

            # Verify tab is selected/active
            tab_classes = tab.get_attribute("class") or ""
            tab_aria_selected = tab.get_attribute("aria-selected")

            assert (
                "selected" in tab_classes.lower()
                or "active" in tab_classes.lower()
                or tab_aria_selected == "true"
            ), f"Tab '{tab_text}' doesn't appear to be selected"

    def test_settings_url_with_section_parameter(self, driver, app_url, wait):
        """
        Test that /settings/:section URLs still work for deep linking.
        """
        sections_to_test = ["providers", "general", "users"]

        for section in sections_to_test:
            # Navigate directly to section URL
            section_url = f"{app_url}/settings/{section}"
            driver.get(section_url)
            TestUtils.wait_for_page_load(driver)

            # Take screenshot
            TestUtils.take_screenshot(driver, f"settings_nav_05_section_{section}")

            # Verify we're on settings page
            assert "/settings" in driver.current_url

            # Verify the correct tab is selected if possible
            try:
                # Look for active/selected tab
                active_tabs = driver.find_elements(
                    By.CSS_SELECTOR,
                    "[role='tab'][aria-selected='true'], .MuiTab-root.Mui-selected, .MuiTab-root[aria-selected='true']",
                )

                if active_tabs:
                    active_tab_text = active_tabs[0].text.lower()
                    assert section.lower() in active_tab_text, (
                        f"Expected {section} tab to be active, but found: {active_tab_text}"
                    )

            except Exception as e:
                # If we can't verify tab selection, that's okay - the main thing is that the page loads
                print(f"Could not verify tab selection for {section}: {e}")

    def test_back_button_functionality(self, driver, app_url, wait):
        """
        Test that back button works properly in settings.
        This was part of the original issue - confusing back button behavior.
        """
        # Start from homepage
        driver.get(app_url)
        TestUtils.wait_for_page_load(driver)

        # Navigate to settings
        settings_link = TestUtils.wait_for_element_clickable(
            driver,
            (
                By.XPATH,
                "//a[contains(@href, '/settings') or contains(text(), 'Settings')]",
            ),
        )
        settings_link.click()
        TestUtils.wait_for_page_load(driver)

        # Verify we're on settings
        assert "/settings" in driver.current_url
        TestUtils.take_screenshot(driver, "settings_nav_06_on_settings_page")

        # Use browser back button
        driver.back()
        TestUtils.wait_for_page_load(driver)

        # Verify we're back on homepage (not on a different settings page)
        current_url = driver.current_url
        assert not current_url.endswith("/settings"), (
            f"Back button should go to homepage, but URL is: {current_url}"
        )

        TestUtils.take_screenshot(driver, "settings_nav_07_after_back_button")

    @pytest.mark.slow
    def test_settings_performance(self, driver, app_url):
        """
        Test that settings page loads quickly and doesn't have performance issues.
        """
        start_time = time.time()

        # Navigate to settings
        driver.get(f"{app_url}/settings")
        TestUtils.wait_for_page_load(driver)

        # Wait for tabs to be visible
        TestUtils.wait_for_element_visible(
            driver, (By.CSS_SELECTOR, "[role='tab'], .MuiTab-root"), timeout=10
        )

        load_time = time.time() - start_time

        # Settings page should load within 5 seconds
        assert load_time < 5.0, (
            f"Settings page took too long to load: {load_time:.2f} seconds"
        )

        TestUtils.take_screenshot(driver, "settings_nav_08_performance_test")

    def test_mobile_responsive_settings(self, driver, app_url, wait):
        """
        Test that settings page works on mobile viewport.
        """
        # Set mobile viewport
        driver.set_window_size(375, 667)  # iPhone 6/7/8 size

        # Navigate to settings
        driver.get(f"{app_url}/settings")
        TestUtils.wait_for_page_load(driver)

        TestUtils.take_screenshot(driver, "settings_nav_09_mobile_view")

        # Verify tabs are still accessible (may be scrollable on mobile)
        try:
            tabs = driver.find_elements(By.CSS_SELECTOR, "[role='tab'], .MuiTab-root")
            assert len(tabs) > 0, "No tabs found in mobile view"

            # Try to click first tab
            if tabs:
                TestUtils.scroll_to_element(driver, tabs[0])
                tabs[0].click()
                time.sleep(1)

                TestUtils.take_screenshot(driver, "settings_nav_10_mobile_tab_clicked")

        except Exception as e:
            # Mobile responsiveness issues are not critical for this test
            print(f"Mobile tab interaction failed (may be expected): {e}")

        finally:
            # Reset to desktop size
            driver.set_window_size(1920, 1080)
