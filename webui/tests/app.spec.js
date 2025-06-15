// file: webui/tests/app.spec.js
import { test, expect } from '@playwright/test';

// Login flow should display dashboard when credentials are accepted.
test('login flow', async ({ page }) => {
  // First call to /api/config should indicate unauthenticated
  await page.route(
    '**/api/config',
    route => route.fulfill({ status: 401, body: '' }),
    { times: 1 }
  );
  // Subsequent calls succeed so the dashboard can load
  await page.route('**/api/config', route => route.fulfill({ status: 200, body: '{}' }));
  await page.route('**/api/login', route => route.fulfill({ status: 200 }));
  await page.route('**/api/setup/status', route =>
    route.fulfill({ status: 200, body: JSON.stringify({ needed: false }) })
  );

  await page.goto('/');
  await page.fill('#username', 'user');
  await page.fill('#password', 'pass');
  await page.getByRole('button', { name: 'Sign In' }).click();
  await expect(page.locator('nav')).toBeVisible();
});
