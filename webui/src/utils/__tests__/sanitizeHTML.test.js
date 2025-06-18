// file: webui/src/utils/__tests__/sanitizeHTML.test.js
import { describe, expect, test } from 'vitest';
import { sanitizeHTML } from '../sanitizeHTML.js';

describe('sanitizeHTML utility', () => {
  test('removes dangerous HTML', () => {
    const dirty = '<img src=x onerror=alert(1)><p>ok</p>';
    const clean = sanitizeHTML(dirty);
    expect(clean).not.toMatch(/onerror/);
    expect(clean).toContain('<img src="x">');
    expect(clean).toContain('<p>ok</p>');
  });
});
