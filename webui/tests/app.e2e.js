// file: webui/tests/app.e2e.js
import { test, expect } from '@playwright/test'

// Login flow should display dashboard when credentials are accepted.
test('login flow', async ({ page }) => {
  await page.route('/api/config', route => route.fulfill({ status: 401, body: '' }))
  await page.route('/api/login', route => route.fulfill({ status: 200 }))

  await page.goto('/')
  await page.fill('input[placeholder="Username"]', 'user')
  await page.fill('input[placeholder="Password"]', 'pass')
  await page.click('text=Login')
  await expect(page.locator('nav')).toBeVisible()
})
