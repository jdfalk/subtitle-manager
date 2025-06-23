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

    // Mock DOM methods to prevent navigation warnings
    global.URL.createObjectURL = vi.fn(() => 'mock-url');
    global.URL.revokeObjectURL = vi.fn();

    // Mock document.createElement and click to prevent navigation
    const originalCreateElement = global.document.createElement;
    const mockAnchor = {
      href: '',
      download: '',
      click: vi.fn(),
    };
    global.document.createElement = vi.fn(tagName => {
      if (tagName === 'a') {
        return mockAnchor;
      }
      return originalCreateElement.call(document, tagName);
    });
  });

  test('uploads file and language to translate endpoint', async () => {
    render(<Translate />);
    const file = new File(['foo'], 'f.srt', { type: 'text/plain' });
    fireEvent.change(screen.getByTestId('file'), { target: { files: [file] } });

    // Click on the select to open it and select a language
    const selectElement = screen.getByRole('combobox');
    fireEvent.mouseDown(selectElement);
    const esOptions = await screen.findAllByText('Spanish');
    fireEvent.click(esOptions[0]); // Click the first Spanish option

    fireEvent.click(screen.getByText(/Translate to/));
    await waitFor(() =>
      expect(fetch).toHaveBeenCalledWith(
        expect.stringContaining('/api/translate'),
        expect.any(Object)
      )
    );
  });
});
