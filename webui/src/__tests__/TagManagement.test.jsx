// file: webui/src/__tests__/TagManagement.test.jsx
import '@testing-library/jest-dom/vitest';
import { render, screen, waitFor } from '@testing-library/react';
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
      expect(fetch).toHaveBeenCalledWith('/api/tags', expect.any(Object))
    );
    expect(await screen.findByText('english')).toBeInTheDocument();
  });
});
