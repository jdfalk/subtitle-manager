import '@testing-library/jest-dom/vitest';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { describe, test, vi } from 'vitest';
import ProviderConfigDialog from '../components/ProviderConfigDialog.jsx';

vi.mock('../services/api.js', () => ({
  apiService: {
    get: vi.fn(),
  },
  getBasePath: () => '',
}));

describe('ProviderConfigDialog', () => {
  test('loads available providers when opened', async () => {
    const { apiService } = await import('../services/api.js');
    apiService.get.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve([{ name: 'opensubtitles' }]),
    });

    render(
      <ProviderConfigDialog open provider={null} onClose={() => {}} onSave={() => {}} />
    );

    await waitFor(() => {
      expect(apiService.get).toHaveBeenCalledWith('/api/providers/available');
    });

    // Open dropdown to verify option
    fireEvent.mouseDown(screen.getByRole('combobox'));
    expect(await screen.findByText('OpenSubtitles.org')).toBeInTheDocument();
  });

  test('calls onSave with entered configuration', async () => {
    const { apiService } = await import('../services/api.js');
    apiService.get.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve([{ name: 'opensubtitles' }]),
    });

    const onSave = vi.fn();
    render(
      <ProviderConfigDialog open provider={null} onClose={() => {}} onSave={onSave} />
    );

    // select provider
    fireEvent.mouseDown(screen.getByRole('combobox'));
    fireEvent.click(await screen.findByText('OpenSubtitles.org'));

    // fill required fields
    fireEvent.change(screen.getByLabelText('API Key *'), { target: { value: 'k' } });
    fireEvent.change(screen.getByLabelText('User Agent *'), { target: { value: 'ua' } });

    fireEvent.click(screen.getByRole('button', { name: 'Save Configuration' }));

    await waitFor(() => {
      expect(onSave).toHaveBeenCalledWith(
        expect.objectContaining({
          name: 'opensubtitles',
          config: expect.objectContaining({ api_key: 'k', user_agent: 'ua' }),
        })
      );
    });
  });
});
