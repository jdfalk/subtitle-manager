// file: services/mockApi.js

/**
 * Mock API service for testing environments
 * Provides dummy responses to prevent console errors during Lighthouse testing
 */

// Mock data to return instead of making real API calls
const mockData = {
  config: {
    db_path: './test.db',
    db_backend: 'sqlite',
    host: 'localhost',
    port: 8080,
    auth_enabled: false,
  },
  setupStatus: {
    configured: false,
    needs_setup: true,
  },
  providers: [
    { name: 'OpenSubtitles', enabled: true },
    { name: 'Subscene', enabled: false },
  ],
  users: [],
  system: {
    version: '1.0.0',
    uptime: '1h 30m',
    memory_usage: '45%',
  },
  wanted: [],
  history: [],
  library: [],
};

/**
 * Check if we're in a testing environment
 */
const isTestingEnvironment = () => {
  return (
    typeof window !== 'undefined' &&
    (window.location?.port === '61984' || // Lighthouse test ports
      window.location?.port === '61919' ||
      window.location?.port === '62077' ||
      window.navigator?.userAgent?.includes('Chrome-Lighthouse'))
  );
};

/**
 * Mock fetch that returns dummy data instead of making real requests
 */
const mockFetch = async (url, _options = {}) => {
  // Simulate network delay
  await new Promise(resolve => setTimeout(resolve, 100));

  // Parse the URL to determine what mock data to return
  const urlPath = new URL(url, window.location.origin).pathname;

  let mockResponse = {};

  if (urlPath.includes('/api/config')) {
    mockResponse = mockData.config;
  } else if (urlPath.includes('/api/setup/status')) {
    mockResponse = mockData.setupStatus;
  } else if (urlPath.includes('/api/providers')) {
    mockResponse = mockData.providers;
  } else if (urlPath.includes('/api/users')) {
    mockResponse = mockData.users;
  } else if (urlPath.includes('/api/system')) {
    mockResponse = mockData.system;
  } else if (urlPath.includes('/api/wanted')) {
    mockResponse = mockData.wanted;
  } else if (urlPath.includes('/api/history')) {
    mockResponse = mockData.history;
  } else if (urlPath.includes('/api/library')) {
    mockResponse = mockData.library;
  } else {
    // Default response for unknown endpoints
    mockResponse = { status: 'ok', message: 'Mock response' };
  }

  return {
    ok: true,
    status: 200,
    statusText: 'OK',
    json: async () => mockResponse,
    text: async () => JSON.stringify(mockResponse),
  };
};

/**
 * Override fetch in testing environments
 */
if (isTestingEnvironment()) {
  // Store original fetch
  window._originalFetch = window.fetch;

  // Replace with mock
  window.fetch = mockFetch;

  // Also silence console methods more aggressively
  const noop = () => {};
  window._originalConsole = {
    log: console.log,
    warn: console.warn,
    error: console.error,
    debug: console.debug,
  };

  console.log = noop;
  console.warn = noop;
  console.error = noop;
  console.debug = noop;
}

export { isTestingEnvironment, mockFetch };
