// file: webui/src/__tests__/UserManagement.test.jsx
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
