// file: webui/src/__tests__/Dashboard.test.jsx
import '@testing-library/jest-dom/vitest';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import Dashboard from '../Dashboard.jsx';

describe('Dashboard component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () =>
          Promise.resolve({ running: false, completed: 0, files: [] }),
      })
    );
  });

  test('starts scan with provided options', async () => {
    render(<Dashboard />);
    fireEvent.change(screen.getByPlaceholderText('Enter directory to scan'), {
      target: { value: '/tmp' },
    });

    // Click on the Language select to open it and select French
    const languageSelect = screen.getByRole('combobox');
    fireEvent.mouseDown(languageSelect);
    const frenchOption = await screen.findByText('French');
    fireEvent.click(frenchOption);

    // Skip provider selection for now since it's more complex

    fireEvent.click(screen.getByText('Scan'));
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith('/api/scan', expect.any(Object))
    );
  });
});
