// file: webui/src/__tests__/History.test.jsx
import '@testing-library/jest-dom/vitest';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import History from '../History.jsx';

describe('History component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () =>
          Promise.resolve({
            translations: [
              { ID: '1', File: 'a.srt', Language: 'en', Service: 'g' },
            ],
            downloads: [
              { ID: '2', VideoFile: 'b.mkv', Language: 'en', Provider: 'os' },
            ],
          }),
      })
    );
  });

  test('loads and filters history', async () => {
    render(<History />);
    await screen.findByText('a.srt');
    fireEvent.change(screen.getByPlaceholderText('Filter by language (e.g., en, es, fr)'), {
      target: { value: 'fr' },
    });
    await waitFor(() =>
      expect(screen.queryByText('a.srt')).not.toBeInTheDocument()
    );
  });
});
