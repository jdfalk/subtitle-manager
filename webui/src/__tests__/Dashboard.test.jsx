// file: webui/src/__tests__/Dashboard.test.jsx
import '@testing-library/jest-dom/vitest';
import { act, render, screen, waitFor } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import Dashboard from '../Dashboard.jsx';

// Mock the API service
vi.mock('../services/api.js', () => ({
  apiService: {
    get: vi.fn(),
    post: vi.fn(),
  },
  getBasePath: () => '',
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
              {
                id: 'opensubtitles',
                name: 'OpenSubtitles',
                enabled: true,
                configured: true,
              },
              {
                id: 'embedded',
                name: 'Embedded',
                enabled: true,
                configured: false,
              },
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

  test('uses data-testid for resilient provider select element identification', async () => {
    await act(async () => {
      render(
        <MemoryRouter
          future={{
            v7_startTransition: true,
            v7_relativeSplatPath: true,
          }}
        >
          <Dashboard />
        </MemoryRouter>
      );
    });

    // Wait for component to load
    await waitFor(() => {
      expect(screen.getByTestId('provider-select')).toBeInTheDocument();
    });

    // Verify provider select is accessible via data-testid instead of brittle array index
    const providerSelect = screen.getByTestId('provider-select');
    expect(providerSelect).toBeInTheDocument();
    expect(providerSelect).toHaveAttribute('data-testid', 'provider-select');
  });

  test('uses data-testid for resilient language select element identification', async () => {
    await act(async () => {
      render(
        <MemoryRouter
          future={{
            v7_startTransition: true,
            v7_relativeSplatPath: true,
          }}
        >
          <Dashboard />
        </MemoryRouter>
      );
    });

    // Wait for component to load
    await waitFor(() => {
      expect(screen.getByTestId('language-select')).toBeInTheDocument();
    });

    // Verify language select is accessible via data-testid instead of brittle array index
    const languageSelect = screen.getByTestId('language-select');
    expect(languageSelect).toBeInTheDocument();
    expect(languageSelect).toHaveAttribute('data-testid', 'language-select');
  });
});
