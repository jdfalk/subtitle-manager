import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import System from '../System.jsx';
global.expect = expect;
global.beforeEach = beforeEach;

describe('System component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(url => {
      if (url.includes('/api/logs'))
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve(['log']),
        });
      if (url.includes('/api/system'))
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ go_version: 'go' }),
        });
      if (url.includes('/api/tasks'))
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ scan: {} }),
        });
      if (url.includes('/api/config'))
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ openai_api_key: 'sk-abcdef123456' }),
        });
      return Promise.resolve({ ok: true, json: () => Promise.resolve({}) });
    });
  });

  test('loads logs and info', async () => {
    render(<System />);
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith(
        expect.stringContaining('/api/logs'),
        expect.any(Object)
      )
    );
    expect(screen.getByTestId('logs').textContent).toBe('log');
  });

  test('masks sensitive config by default and reveals when toggled', async () => {
    render(<System />);
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith(
        expect.stringContaining('/api/config'),
        expect.any(Object)
      )
    );
    await waitFor(() =>
      expect(screen.getByTestId('config').textContent).toContain('****3456')
    );
    const toggle = screen.getByLabelText('Show Sensitive');
    fireEvent.click(toggle);
    await waitFor(() =>
      expect(screen.getByTestId('config').textContent).toContain(
        'sk-abcdef123456'
      )
    );
  });
});
