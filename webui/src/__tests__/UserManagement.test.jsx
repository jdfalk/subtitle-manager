// file: webui/src/__tests__/UserManagement.test.jsx
// version: 1.1.0
// guid: 184acee4-a24c-4363-8893-b3d5394f8e5c
import '@testing-library/jest-dom/vitest';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import UserManagement from '../UserManagement.jsx';

describe('UserManagement component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn();
  });

  test('displays usernames from API', async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () =>
        Promise.resolve([
          { id: '1', username: 'alice', email: 'a@example.com', role: 'admin' },
        ]),
    });

    render(<UserManagement />);
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith('/api/users', expect.any(Object))
    );
    expect(await screen.findByText('alice')).toBeInTheDocument();
  });

  test('shows fallback when username missing', async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () =>
        Promise.resolve([
          {
            id: '42',
            email: 'no-name@example.com',
            role: 'user',
            active: true,
          },
        ]),
    });

    render(<UserManagement />);

    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith('/api/users', expect.any(Object))
    );

    expect(await screen.findByText('42')).toBeInTheDocument();
  });

  test('opens editor dialog when add user clicked', async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve([]),
    });

    render(<UserManagement />);
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith('/api/users', expect.any(Object))
    );

    fireEvent.click(screen.getByText('Add User'));

    const dialog = await screen.findByRole('dialog');
    expect(dialog).toBeInTheDocument();
    expect(screen.getByText('Save')).toBeInTheDocument();
  });
});
