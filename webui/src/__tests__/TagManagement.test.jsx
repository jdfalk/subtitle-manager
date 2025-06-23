// file: webui/src/__tests__/TagManagement.test.jsx
import '@testing-library/jest-dom/vitest';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, test, vi } from 'vitest';
import TagManagement from '../TagManagement.jsx';

describe('TagManagement component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn();
  });

  test('displays tags from API', async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve([{ id: '1', name: 'english' }]),
    });

    render(<TagManagement />);
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith(
        expect.stringContaining('/api/tags'),
        expect.any(Object)
      )
    );
    expect(await screen.findByText('english')).toBeInTheDocument();
  });

  test('allows editing a tag name', async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve([{ id: '1', name: 'english' }]),
    });

    render(<TagManagement />);
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith(expect.stringContaining('/api/tags'))
    );

    fetch.mockResolvedValueOnce({ ok: true });
    fireEvent.click(screen.getByRole('button', { name: /edit/i }));
    fireEvent.change(screen.getByDisplayValue('english'), {
      target: { value: 'spanish' },
    });
    fireEvent.click(screen.getByRole('button', { name: /save/i }));

    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith(
        expect.stringContaining('/api/tags/1'),
        expect.objectContaining({ method: 'PATCH' })
      )
    );
  });
});
