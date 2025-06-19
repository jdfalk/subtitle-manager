// file: webui/src/__tests__/Settings.test.jsx
import '@testing-library/jest-dom/vitest';
import {
  act,
  fireEvent,
  render,
  screen,
  waitFor,
} from '@testing-library/react';
import { beforeEach, afterEach, describe, expect, test, vi } from 'vitest';
import Settings from '../Settings.jsx';

// Mock the API service
vi.mock('../services/api.js', () => ({
  apiService: {
    get: vi.fn(),
    post: vi.fn(),
  },
  getBasePath: () => '',
}));

describe('Settings component', () => {
  beforeEach(async () => {
    vi.clearAllMocks();

    global.fetch = vi.fn(url => {
      if (url === '/api/providers/available') {
        return Promise.resolve({
          ok: true,
          json: () =>
            Promise.resolve([{ name: 'opensubtitles' }, { name: 'subscene' }]),
        });
      }
      return Promise.resolve({ ok: true, json: () => Promise.resolve({}) });
    });

    // Get the mocked apiService
    const { apiService } = await import('../services/api.js');

    // Setup default mocks
    apiService.get.mockImplementation(url => {
      if (url === '/api/config') {
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ address: '0.0.0.0' }),
        });
      }
      if (url === '/api/providers') {
        return Promise.resolve({
          ok: true,
          json: () =>
            Promise.resolve([
              { name: 'opensubtitles', enabled: true, config: {} },
              { name: 'subscene', enabled: false, config: {} },
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
      json: () => Promise.resolve({}),
    });
  });

  afterEach(() => {
    delete global.fetch;
  });
  test('loads settings and renders tabs', async () => {
    await act(async () => {
      render(<Settings />);
    });

    // Wait for the settings to load and tabs to appear
    await waitFor(() => {
      expect(screen.getByText('Settings')).toBeInTheDocument();
      expect(screen.getByText('General')).toBeInTheDocument();
      expect(screen.getByText('Providers')).toBeInTheDocument();
    });

    // Verify API calls were made to load config
    const { apiService } = await import('../services/api.js');

    await waitFor(() => {
      expect(apiService.get).toHaveBeenCalledWith('/api/config');
      expect(apiService.get).toHaveBeenCalledWith('/api/providers');
    });

    // Click the General tab to verify it can be selected
    await act(async () => {
      fireEvent.click(screen.getByRole('tab', { name: /General/i }));
    });

    // Verify the General tab becomes selected
    await waitFor(() => {
      const generalTab = screen.getByRole('tab', { name: /General/i });
      expect(generalTab).toHaveAttribute('aria-selected', 'true');
    });
  });

  test('shows only enabled providers', async () => {
    await act(async () => {
      render(<Settings />);
    });

    await waitFor(() => {
      expect(screen.getByText('OpenSubtitles')).toBeInTheDocument();
    });

    expect(screen.queryByText('Subscene')).toBeNull();
  });

  test('add provider dialog lists all providers', async () => {
    await act(async () => {
      render(<Settings />);
    });

    const addButton = screen.getByText('Add Provider');
    await act(async () => {
      fireEvent.click(addButton);
    });

    await waitFor(() => {
      expect(screen.getByRole('dialog')).toBeInTheDocument();
    });

    const select = screen.getByRole('combobox');
    await act(async () => {
      fireEvent.mouseDown(select);
    });

    // All available providers should be listed
    expect(await screen.findByText('Subscene')).toBeInTheDocument();
  });
});
