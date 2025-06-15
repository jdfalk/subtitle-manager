// file: webui/src/__tests__/Wanted.test.jsx
import '@testing-library/jest-dom';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { expect, vi } from 'vitest';
import Wanted from '../Wanted.jsx';

describe('Wanted component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn();
  });

  test('loads wanted list on mount and searches', async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve(['a']),
    });
    render(<Wanted />);
    await screen.findByText('a');

    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve(['u']),
    });
    fireEvent.change(screen.getByPlaceholderText('/path/to/movie.mkv'), {
      target: { value: 'f' },
    });
    fireEvent.click(screen.getByText('Search'));
    await waitFor(() =>
      expect(fetch).toHaveBeenLastCalledWith(
        '/api/search?provider=embedded&path=f&lang=en'
      )
    );
  });
});
