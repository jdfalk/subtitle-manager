// file: webui/src/test-setup.js
import * as matchers from '@testing-library/jest-dom/matchers';
import { cleanup } from '@testing-library/react';
import { afterEach, expect, vi } from 'vitest';

// Extend Vitest's expect with jest-dom matchers
expect.extend(matchers);

// Mock window.matchMedia for Material-UI components
Object.defineProperty(window, 'matchMedia', {
  writable: true,
  value: vi.fn().mockImplementation(query => ({
    matches: false,
    media: query,
    onchange: null,
    addListener: vi.fn(), // deprecated
    removeListener: vi.fn(), // deprecated
    addEventListener: vi.fn(),
    removeEventListener: vi.fn(),
    dispatchEvent: vi.fn(),
  })),
});

// Mock window.URL.createObjectURL for file download tests
Object.defineProperty(window.URL, 'createObjectURL', {
  writable: true,
  value: vi.fn(() => 'mocked-url'),
});

Object.defineProperty(window.URL, 'revokeObjectURL', {
  writable: true,
  value: vi.fn(),
});

// Mock ResizeObserver more thoroughly
class MockResizeObserver {
  constructor(callback) {
    this.callback = callback;
    this.observations = new Map();
  }

  observe(element) {
    this.observations.set(element, true);
    // Simulate a resize event immediately
    if (this.callback) {
      try {
        this.callback([
          {
            target: element,
            contentRect: { width: 100, height: 100 },
            borderBoxSize: [{ blockSize: 100, inlineSize: 100 }],
            contentBoxSize: [{ blockSize: 100, inlineSize: 100 }],
            devicePixelContentBoxSize: [{ blockSize: 100, inlineSize: 100 }]
          }
        ], this);
      } catch (error) {
        console.warn('ResizeObserver callback error:', error);
      }
    }
  }

  unobserve(element) {
    this.observations.delete(element);
  }

  disconnect() {
    this.observations.clear();
  }
}

global.ResizeObserver = MockResizeObserver;

// Cleanup after each test
afterEach(() => {
  cleanup();
});
