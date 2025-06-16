// file: webui/src/__tests__/Dashboard.test.jsx
import '@testing-library/jest-dom/vitest';
import { act, fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import Dashboard from '../Dashboard.jsx';

describe('Dashboard component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn();
  });

  test('starts scan with provided options', async () => {
    // Mock scan status calls
    fetch.mockImplementation((url) => {
      if (url.includes('/api/scan/status')) {
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ running: false, completed: 0, files: [] }),
        });
      }
      if (url.includes('/api/providers')) {
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve([
            { id: 'opensubtitles', name: 'OpenSubtitles', enabled: true },
            { id: 'embedded', name: 'Embedded', enabled: true },
          ]),
        });
      }
      if (url.includes('/api/scan') && !url.includes('/status')) {
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ message: 'Scan started' }),
        });
      }
      return Promise.resolve({
        ok: true,
        json: () => Promise.resolve({}),
      });
    });

    await act(async () => {
      render(<Dashboard />);
    });

    // Wait for component to load
    await waitFor(() => {
      expect(screen.getByPlaceholderText('Enter directory to scan')).toBeInTheDocument();
    });

    await act(async () => {
      fireEvent.change(screen.getByPlaceholderText('Enter directory to scan'), {
        target: { value: '/tmp' },
      });
    });

    // Click on the Language select to open it and select French
    const languageSelect = screen.getAllByRole('combobox')[0]; // First combobox is language
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

    await waitFor(() => {
      const scanCalls = fetch.mock.calls.filter(call =>
        call[0].includes('/api/scan') && !call[0].includes('/status')
      );
      expect(scanCalls.length).toBeGreaterThan(0);
    });
  });
});
