import { defineConfig } from '@playwright/test';

export default defineConfig({
  // Only look for e2e tests under the `tests` directory. Without this
  // Playwright will try to execute unit tests inside `src/__tests__`
  // which import CSS files that Node cannot parse directly.
  testDir: 'tests',
  webServer: {
    // Run Vite dev server on a fixed port so Playwright can reliably connect.
    command: 'npm run dev -- --port 5173 --strictPort',
    port: 5173,
    reuseExistingServer: !process.env.CI,
  },
  use: {
    baseURL: 'http://localhost:5173',
  },
});
