// file: webui/src/__tests__/Translate.test.jsx
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import '@testing-library/jest-dom'
import { vi } from 'vitest'
import Translate from '../Translate.jsx'

describe('Translate component', () => {
  beforeEach(() => {
    vi.restoreAllMocks()
    global.fetch = vi.fn(() => Promise.resolve({ ok: true, blob: () => Promise.resolve(new Blob()) }))
  })

  test('uploads file and language to translate endpoint', async () => {
    render(<Translate />)
    const file = new File(['foo'], 'f.srt', { type: 'text/plain' })
    fireEvent.change(screen.getByTestId('file'), { target: { files: [file] } })
    fireEvent.change(screen.getByPlaceholderText('Language'), { target: { value: 'es' } })
    fireEvent.click(screen.getByText('Translate'))
    await waitFor(() => expect(fetch).toHaveBeenCalledWith('/api/translate', expect.any(Object)))
  })
})
