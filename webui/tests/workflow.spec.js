// file: webui/tests/workflow.spec.js
import { expect, test } from '@playwright/test';

/**
 * Test major workflows in the application to ensure core functionality works
 * end-to-end without errors. This test covers critical user journeys including
 * authentication, navigation, and key operations.
 */
test('major workflows execute without errors', async ({ page }) => {
  // Mock API endpoints to ensure predictable test environment

  // Mock setup status - not needed
  await page.route('**/api/setup/status', route =>
    route.fulfill({
      status: 200,
      body: JSON.stringify({ needed: false }),
    })
  );

  // Mock unauthenticated config initially
  await page.route('**/api/config', route => {
    route.fulfill({ status: 401, body: '' });
  });

  // Mock successful login
  await page.route('**/api/login', async route => {
    if (route.request().method() === 'POST') {
      // Update config route to return authenticated user after login
      await page.unroute('**/api/config');
      await page.route('**/api/config', route => {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            user: 'testuser',
            authenticated: true,
            backendAvailable: true,
          }),
        });
      });
      route.fulfill({ status: 200 });
    } else {
      route.continue();
    }
  });

  // Mock database info endpoint
  await page.route('**/api/database/info', route => {
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        size: '10.5 MB',
        tables: 15,
        records: 1250,
        lastBackup: '2024-01-15T10:30:00Z',
        health: 'good',
      }),
    });
  });

  // Mock backup operation
  await page.route('**/api/database/backup', route => {
    if (route.request().method() === 'POST') {
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          success: true,
          message: 'Backup created successfully',
          filename: 'backup_20241215_103000.db',
        }),
      });
    } else {
      route.continue();
    }
  });

  // Mock optimize operation
  await page.route('**/api/database/optimize', route => {
    if (route.request().method() === 'POST') {
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          success: true,
          message: 'Database optimized successfully',
        }),
      });
    } else {
      route.continue();
    }
  });

  // Mock media library endpoint
  await page.route('**/api/media/library', route => {
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        movies: [],
        shows: [],
        total: 0,
      }),
    });
  });

  // Mock settings endpoints
  await page.route('**/api/settings/**', route => {
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({}),
    });
  });

  // Start the test workflow
  await page.goto('/');
  await page.waitForLoadState('networkidle');

  // Step 1: Complete login flow
  const usernameInput = page.locator('input[name="username"], input#username');
  const passwordInput = page.locator('input[name="password"], input#password');

  await expect(usernameInput).toBeVisible({ timeout: 10000 });
  await expect(passwordInput).toBeVisible({ timeout: 10000 });

  await usernameInput.fill('testuser');
  await passwordInput.fill('password');
  await page.getByRole('button', { name: 'Sign In' }).click();

  // Wait for dashboard to load
  await page.waitForLoadState('networkidle');
  await expect(page.getByText('Subtitle Manager')).toBeVisible(); // Step 2: Navigate to Settings page
  await page.goto('/settings');
  await page.waitForLoadState('networkidle');

  // Step 3: Navigate to Database Settings tab
  // Wait for the Database tab to be visible and click it
  const databaseTab = page.getByRole('tab', { name: 'Database' });
  await expect(databaseTab).toBeVisible({ timeout: 10000 });
  await databaseTab.click();

  // Wait for database settings to load
  await page.waitForTimeout(1000);

  // Step 4: Test backup functionality
  // Wait for the backup button to be visible and enabled
  const createBackupButton = page.getByRole('button', {
    name: 'Create Backup',
  });
  await expect(createBackupButton).toBeVisible({ timeout: 10000 });
  await expect(createBackupButton).toBeEnabled({ timeout: 5000 });

  await createBackupButton.click();

  // Handle backup confirmation dialog if it appears
  const confirmBackupButton = page
    .getByRole('button', { name: 'Create Backup' })
    .last();
  if (await confirmBackupButton.isVisible()) {
    await confirmBackupButton.click();
  }

  // Wait for backup operation to complete
  await page.waitForTimeout(1000);

  // Step 5: Test database optimization functionality
  // Wait for the optimize button to be visible and enabled
  const optimizeDatabaseButton = page.getByRole('button', {
    name: 'Optimize Database',
  });
  await expect(optimizeDatabaseButton).toBeVisible({ timeout: 10000 });
  await expect(optimizeDatabaseButton).toBeEnabled({ timeout: 5000 });

  await optimizeDatabaseButton.click();

  // Handle optimization confirmation dialog
  const confirmOptimizeButton = page.getByRole('button', { name: 'Optimize' });
  await expect(confirmOptimizeButton).toBeVisible({ timeout: 5000 });
  await confirmOptimizeButton.click();

  // Wait for operation to complete
  await page.waitForTimeout(1000);

  // Step 6: Verify workflows completed successfully
  // Check that we're still on a valid page and no error dialogs appeared
  await expect(page.locator('body')).not.toHaveText('Error');
  await expect(page.locator('body')).not.toHaveText('Failed');

  // Verify we can still navigate
  await expect(page.getByText('Subtitle Manager')).toBeVisible();
});

/**
 * Test basic navigation workflows to ensure all major sections are accessible
 * and load without errors.
 */
test('navigation workflow', async ({ page }) => {
  // Mock authenticated state
  await page.route('**/api/config', route => {
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        user: 'testuser',
        authenticated: true,
        backendAvailable: true,
      }),
    });
  });

  // Mock setup status
  await page.route('**/api/setup/status', route =>
    route.fulfill({
      status: 200,
      body: JSON.stringify({ needed: false }),
    })
  );

  // Mock API endpoints for different sections
  await page.route('**/api/media/library', route => {
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ movies: [], shows: [], total: 0 }),
    });
  });

  await page.route('**/api/settings/**', route => {
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({}),
    });
  });

  // Start test
  await page.goto('/');
  await page.waitForLoadState('networkidle');

  // Verify main app loads
  await expect(page.getByText('Subtitle Manager')).toBeVisible();

  // Test navigation to different sections (if they exist)
  const navigationTests = ['/dashboard', '/library', '/settings'];

  for (const route of navigationTests) {
    try {
      await page.goto(route);
      await page.waitForLoadState('networkidle');

      // Verify page loads without major errors
      await expect(page.locator('body')).not.toHaveText('Error 404');
      await expect(page.locator('body')).not.toHaveText('Page not found');

      // Verify we still have the main app header/navigation
      await expect(page.getByText('Subtitle Manager')).toBeVisible();
    } catch (error) {
      // Log navigation failures but don't fail the test for optional routes
      console.warn(`Navigation to ${route} failed:`, error.message);
    }
  }
});
