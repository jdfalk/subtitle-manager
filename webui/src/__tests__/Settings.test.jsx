// file: webui/src/__tests__/Settings.test.jsx
import '@testing-library/jest-dom/vitest';
import { act, fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import Settings from '../Settings.jsx';

// Mock the API service
vi.mock('../services/api.js', () => ({
  apiService: {
    get: vi.fn(),
    post: vi.fn(),
  },
}));

describe('Settings component', () => {
  beforeEach(async () => {
    vi.clearAllMocks();

    // Get the mocked apiService
    const { apiService } = await import('../services/api.js');

    // Setup default mocks
    apiService.get.mockImplementation((url) => {
      if (url === '/api/config') {
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ address: '0.0.0.0' }),
        });
      }
      if (url === '/api/providers') {
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve([]),
        });
      }
      return Promise.resolve({
        ok: true,
        json: () => Promise.resolve({}),
      });
    });

    apiService.post.mockResolvedValue({
      ok: true,
      json: () => Promise.resolve({}),
    });
  });

  test('edits general settings', async () => {
    await act(async () => {
      render(<Settings />);
    });

    // Wait for the settings to load and tabs to appear
    await waitFor(() => screen.getByText('General'));

    // Click the General tab
    await act(async () => {
      fireEvent.click(screen.getByRole('tab', { name: /General/i }));
    });

    // Wait for the general settings form to load
    await screen.findByDisplayValue('0.0.0.0');

    // Get the mocked apiService to mock additional calls
    const { apiService } = await import('../services/api.js');

    await act(async () => {
      fireEvent.change(screen.getByLabelText('Address'), {
        target: { value: '127.0.0.1' },
      });
    });

    await act(async () => {
      fireEvent.click(screen.getByText('Save'));
    });

    await waitFor(() => {
      expect(apiService.post).toHaveBeenCalledWith('/api/config', expect.any(Object));
    });
  });
});
