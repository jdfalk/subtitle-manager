// file: webui/src/__tests__/ConfigEditor.test.jsx
import '@testing-library/jest-dom';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { vi } from 'vitest';
import ConfigEditor from '../ConfigEditor.jsx';

describe('ConfigEditor component', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(url => {
      if (url === '/api/config') {
        return Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ test_key: 'value' }),
        });
      }
      return Promise.resolve({ ok: true });
    });
  });

  test('loads config and saves updates', async () => {
    render(<ConfigEditor />);
    await waitFor(() => expect(fetch).toHaveBeenCalledWith('/api/config'));
    expect(screen.getByTestId('config-editor').value).toContain('test_key');
    fireEvent.change(screen.getByTestId('config-editor'), {
      target: { value: 'test_key: new' },
    });
    fireEvent.click(screen.getByText('Save'));
    await waitFor(() => expect(fetch).toHaveBeenLastCalledWith('/api/config', expect.any(Object)));
  });
});

