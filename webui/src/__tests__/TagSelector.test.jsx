// file: webui/src/__tests__/TagSelector.test.jsx
import '@testing-library/jest-dom/vitest';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import TagSelector from '../components/TagSelector.jsx';

describe('TagSelector component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn();
  });

  test('loads tags for a media path', async () => {
    fetch
      .mockResolvedValueOnce({ ok: true, json: () => Promise.resolve([]) })
      .mockResolvedValueOnce({ ok: true, json: () => Promise.resolve([]) });
    render(<TagSelector path="/a.mkv" />);
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith('/api/library/tags?path=%2Fa.mkv')
    );
  });

  test('assigns a tag when checked', async () => {
    fetch
      .mockResolvedValueOnce({ ok: true, json: () => Promise.resolve([]) })
      .mockResolvedValueOnce({
        ok: true,
        json: () => Promise.resolve([{ id: '1', name: 'tag' }]),
      })
      .mockResolvedValue({ ok: true, json: () => Promise.resolve([]) });

    render(<TagSelector path="/b.mkv" />);
    fireEvent.click(screen.getByRole('button', { name: /manage tags/i }));
    await screen.findByText('tag');
    fireEvent.click(screen.getByRole('checkbox'));

    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith(
        '/api/library/tags',
        expect.objectContaining({ method: 'POST' })
      )
    );
  });
});
