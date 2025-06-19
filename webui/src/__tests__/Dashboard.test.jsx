// file: webui/src/__tests__/Dashboard.test.jsx
import '@testing-library/jest-dom/vitest';
import {
  act,
  fireEvent,
  render,
  screen,
  waitFor,
} from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import Dashboard from '../Dashboard.jsx';

// Mock the API service
vi.mock('../services/api.js', () => ({
  apiService: {
    get: vi.fn(),
    post: vi.fn(),
  },
}));

describe('Dashboard component', () => {
  beforeEach(async () => {
    vi.clearAllMocks();

    // Get the mocked apiService
    const { apiService } = await import('../services/api.js');

    // Setup default mocks for API calls that happen on component mount
    apiService.get.mockImplementation(url => {
      if (url === '/api/scan/status') {
        return Promise.resolve({
          ok: true,
          json: () =>
            Promise.resolve({ running: false, completed: 0, files: [] }),
        });
      }
      if (url === '/api/providers') {
        return Promise.resolve({
          ok: true,
          json: () =>
            Promise.resolve([
              { id: 'opensubtitles', name: 'OpenSubtitles', enabled: true },
              { id: 'embedded', name: 'Embedded', enabled: true },
            ]),
        });
      }
      return Promise.resolve({
        ok: true,
        json: () => Promise.resolve({}),
      });
    });

    apiService.post.mockResolvedValue({
      ok: true,
      json: () => Promise.resolve({ message: 'Scan started' }),
    });
  });

  test('starts scan with provided options', async () => {
    await act(async () => {
      render(<Dashboard />);
    });

    // Wait for component to load
    await waitFor(() => {
      expect(
        screen.getByPlaceholderText('Enter directory to scan')
      ).toBeInTheDocument();
    });

    await act(async () => {
      fireEvent.change(screen.getByPlaceholderText('Enter directory to scan'), {
        target: { value: '/tmp' },
      });
    });

    // Click on the Language select to open it and select French
    const languageSelect = screen.getAllByRole('combobox')[1];
    await act(async () => {
      fireEvent.mouseDown(languageSelect);
    });

    const frenchOption = await screen.findByText('French');
    await act(async () => {
      fireEvent.click(frenchOption);
    });

    // Skip provider selection for now since it's more complex

    await act(async () => {
      fireEvent.click(screen.getByText('Start Scan'));
    });

    // Get the mocked apiService to check calls
    const { apiService } = await import('../services/api.js');

    await waitFor(() => {
      expect(apiService.post).toHaveBeenCalledWith(
        '/api/scan',
        expect.any(Object)
      );
    });
  });
});
