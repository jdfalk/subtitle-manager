// file: webui/src/utils/__tests__/security.test.js
import { describe, expect, test } from 'vitest';
import { validateAllowedHostname } from '../security.js';

describe('validateAllowedHostname', () => {
  test('accepts allowed host', () => {
    expect(
      validateAllowedHostname('https://omdbapi.com/?t=a', ['omdbapi.com'])
    ).toBe(true);
  });

  test('rejects disallowed host', () => {
    expect(validateAllowedHostname('https://evil.com/', ['omdbapi.com'])).toBe(
      false
    );
  });
});
