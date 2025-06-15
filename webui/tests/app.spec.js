// file: webui/tests/app.spec.js
import { expect, test } from '@playwright/test';

// Login flow should display dashboard when credentials are accepted.
test('login flow', async ({ page }) => {
  // Initially, /api/config should indicate unauthenticated
  await page.route('**/api/config', route => {
    route.fulfill({ status: 401, body: '' });
  });

  // Mock setup status to indicate setup is not needed
  await page.route('**/api/setup/status', route =>
    route.fulfill({ status: 200, body: JSON.stringify({ needed: false }) })
  );

  // Mock login endpoint - when login succeeds, update config route to return success
  await page.route('**/api/login', async route => {
    const request = route.request();
    if (request.method() === 'POST') {
      // After successful login, change config to return authenticated user
      await page.unroute('**/api/config');
      await page.route('**/api/config', route => {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ user: 'testuser', authenticated: true })
        });
      });
      route.fulfill({ status: 200 });
    } else {
      route.continue();
    }
  });

  await page.goto('/');

  // Wait for page to load
  await page.waitForLoadState('networkidle');

  // Now look for login form elements
  const usernameInput = page.locator('input[name="username"], input#username');
  const passwordInput = page.locator('input[name="password"], input#password');

  await expect(usernameInput).toBeVisible({ timeout: 10000 });
  await expect(passwordInput).toBeVisible({ timeout: 10000 });

  await usernameInput.fill('user');
  await passwordInput.fill('pass');
  await page.getByRole('button', { name: 'Sign In' }).click();

  // Wait for navigation to dashboard
  await page.waitForLoadState('networkidle');

  // Look for the app bar (navigation) that appears after successful login
  await expect(page.locator('.MuiAppBar-root, [role="banner"], header')).toBeVisible();

  // Also check that we can see the main navigation elements
  await expect(page.getByText('Subtitle Manager')).toBeVisible();
});
