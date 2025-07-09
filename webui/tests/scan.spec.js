// file: webui/tests/scan.spec.js
import { expect, test } from '@playwright/test';

/**
 * Log in by mocking authentication endpoints so the dashboard loads.
 * @param {import('@playwright/test').Page} page - Playwright page
 */
async function login(page) {
  await page.route('**/api/config', route => route.fulfill({ status: 401 }));
  await page.route('**/api/setup/status', route =>
    route.fulfill({ status: 200, body: JSON.stringify({ needed: false }) })
  );
  await page.route('**/api/login', async route => {
    if (route.request().method() === 'POST') {
      await page.unroute('**/api/config');
      await page.route('**/api/config', r =>
        r.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ user: 'tester', authenticated: true }),
        })
      );
      route.fulfill({ status: 200 });
    } else {
      route.continue();
    }
  });

  await page.goto('/');
  await page.waitForLoadState('networkidle');
  await page.locator('input[name="username"], input#username').fill('u');
  await page.locator('input[name="password"], input#password').fill('p');
  await page.getByRole('button', { name: 'Sign In' }).click();
  await page.waitForLoadState('networkidle');
}

/**
 * Verify library scanning triggers API calls and shows progress.
 */
test('library scan workflow', async ({ page }) => {
  await login(page);

  let scanCalled = false;
  await page.route('**/api/library/scan', route => {
    scanCalled = true;
    route.fulfill({ status: 200 });
  });
  await page.route('**/api/library/scan/status', route =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        running: false,
        completed: 5,
        files: ['movie.mkv'],
      }),
    })
  );

  await page.goto('/dashboard');
  await page.waitForLoadState('networkidle');

  const scanInput = page
    .getByPlaceholder(/directory.*scan/i)
    .or(page.locator('input[placeholder*="scan"]'))
    .or(page.locator('input[type="text"]'))
    .first();
  const scanButton = page
    .getByRole('button', { name: /Start Scan/i })
    .or(page.getByRole('button', { name: /Scan/i }));

  await expect(scanInput).toBeVisible();
  await expect(scanButton).toBeVisible();

  await scanInput.fill('/media');
  await scanButton.click();

  await expect.poll(() => scanCalled).toBe(true);
  await expect(page.getByText(/Processed 5 files/)).toBeVisible();
});
