// file: webui/src/utils/__tests__/url.test.js
// version: 1.0.0
// guid: b10dabf3-a0db-474f-88f5-86cc4c8915f1

import { describe, expect, test } from 'vitest';
import { normalizeUrl } from '../url.js';

describe('normalizeUrl', () => {
  test('adds http to URLs without scheme', () => {
    const result = normalizeUrl('localhost:6767');
    expect(result).toBe('http://localhost:6767');
  });

  test('keeps existing http scheme', () => {
    const result = normalizeUrl('http://example.com');
    expect(result).toBe('http://example.com');
  });

  test('keeps existing https scheme', () => {
    const result = normalizeUrl('https://example.com');
    expect(result).toBe('https://example.com');
  });
});
