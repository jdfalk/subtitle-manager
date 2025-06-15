// file: webui/src/__tests__/Translate.test.jsx
import '@testing-library/jest-dom';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { vi } from 'vitest';
import Translate from '../Translate.jsx';

describe('Translate component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(() =>
      Promise.resolve({ ok: true, blob: () => Promise.resolve(new Blob()) })
    );
  });

  test('uploads file and language to translate endpoint', async () => {
    render(<Translate />);
    const file = new File(['foo'], 'f.srt', { type: 'text/plain' });
    fireEvent.change(screen.getByTestId('file'), { target: { files: [file] } });

    // Click on the select to open it and select a language
    const selectElement = screen.getByRole('combobox');
    fireEvent.mouseDown(selectElement);
    const esOption = await screen.findByText('Spanish');
    fireEvent.click(esOption);

    fireEvent.click(screen.getByText('Translate'));
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith('/api/translate', expect.any(Object))
    );
  });
});
