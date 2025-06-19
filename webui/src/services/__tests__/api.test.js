import { describe, expect, test, afterEach } from 'vitest';
import { getBasePath } from '../api.js';

function setPath(path) {
  delete window.location;
  window.location = new URL(`http://localhost${path}`);
}

describe('getBasePath', () => {
  const originalLocation = window.location;

  afterEach(() => {
    // Restore original location after each test
    window.location = originalLocation;
  });

  test('returns empty string for known routes', () => {
    setPath('/library');
    expect(getBasePath()).toBe('');
    setPath('/details');
    expect(getBasePath()).toBe('');
  });

  test('returns prefix for unknown first segment', () => {
    setPath('/prefix/dashboard');
    expect(getBasePath()).toBe('/prefix');
  });
});
