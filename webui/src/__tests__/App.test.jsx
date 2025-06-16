// file: webui/src/__tests__/App.test.jsx
import '@testing-library/jest-dom/vitest';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import App from '../App.jsx';

describe('App component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(() =>
      Promise.resolve({ ok: false, json: () => Promise.resolve({}) })
    );
  });

  test('shows login form when unauthenticated', async () => {
    fetch.mockResolvedValueOnce({ ok: false });
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve({ needed: false }),
    });
    render(<App />);
    expect(screen.getByText('Subtitle Manager')).toBeInTheDocument();
  });

  test('successful login renders dashboard', async () => {
    fetch.mockResolvedValueOnce({ ok: false }); // config check
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve({ needed: false }),
    });
    render(<App />);
    fetch.mockResolvedValueOnce({ ok: true });
    fireEvent.click(screen.getByText('Sign In'));
    await waitFor(() =>
      expect(fetch).toHaveBeenLastCalledWith('/api/login', expect.any(Object))
    );
  });

  test('pins sidebar and persists state', async () => {
    fetch.mockResolvedValueOnce({ ok: false });
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve({ needed: false }),
    });
    render(<App />);
    fetch.mockResolvedValueOnce({ ok: true });
    fireEvent.click(screen.getByText('Sign In'));
    await waitFor(() => screen.getByLabelText('open drawer'));

    fireEvent.click(screen.getByLabelText('open drawer'));
    const pinButton = await screen.findByRole('button', {
      name: 'Pin Sidebar',
    });
    fireEvent.click(pinButton);
    expect(localStorage.getItem('sidebarPinned')).toBe('true');
    fireEvent.click(screen.getByRole('button', { name: 'Unpin Sidebar' }));
    expect(localStorage.getItem('sidebarPinned')).toBe('false');
  });
});
