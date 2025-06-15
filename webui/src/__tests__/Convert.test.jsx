// file: webui/src/__tests__/Convert.test.jsx
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import { vi } from 'vitest';
import Convert from '../Convert.jsx';

describe('Convert component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(() =>
      Promise.resolve({ ok: true, blob: () => Promise.resolve(new Blob()) })
    );
  });

  test('uploads file to convert endpoint', async () => {
    render(<Convert />);
    const file = new File(['foo'], 'f.srt', { type: 'text/plain' });
    fireEvent.change(screen.getByTestId('file'), { target: { files: [file] } });
    fireEvent.click(screen.getByText('Convert'));
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith('/api/convert', expect.any(Object))
    );
  });
});
