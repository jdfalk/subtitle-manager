// file: webui/src/__tests__/DatabaseSettings.test.jsx
import '@testing-library/jest-dom/vitest';
import { render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import DatabaseSettings from '../components/DatabaseSettings.jsx';

describe('DatabaseSettings component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(url => {
      if (url === '/api/database/info') {
        return Promise.resolve({
          ok: true,
          json: () =>
            Promise.resolve({
              type: 'postgresql',
              version: '13',
              size: 1048576,
              path: '/db',
              connected: true,
            }),
        });
      }
      if (url === '/api/database/stats') {
        return Promise.resolve({
          ok: true,
          json: () =>
            Promise.resolve({
              totalRecords: 100,
              users: 5,
              downloads: 20,
              mediaItems: 30,
              lastBackup: '2024-05-01T00:00:00Z',
            }),
        });
      }
      return Promise.resolve({ ok: true, json: () => Promise.resolve({}) });
    });
  });

  test('displays database info from API', async () => {
    render(<DatabaseSettings config={{}} onSave={() => {}} backendAvailable />);

    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith('/api/database/info')
    );

    expect(await screen.findByText('postgresql')).toBeInTheDocument();
    expect(await screen.findByText('Connected')).toBeInTheDocument();
  });

  test('shows warning when backend unavailable', () => {
    render(
      <DatabaseSettings
        config={{}}
        onSave={() => {}}
        backendAvailable={false}
      />
    );

    expect(
      screen.getByText(/Backend service is not available/i)
    ).toBeInTheDocument();
  });
});
