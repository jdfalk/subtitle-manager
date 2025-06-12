// file: webui/src/__tests__/Extract.test.jsx
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import '@testing-library/jest-dom'
import { vi } from 'vitest'
import Extract from '../Extract.jsx'

describe('Extract component', () => {
  beforeEach(() => {
    vi.restoreAllMocks()
    global.fetch = vi.fn(() => Promise.resolve({ ok: true, json: () => Promise.resolve([]) }))
  })

  test('posts path and displays status', async () => {
    render(<Extract />)
    fireEvent.change(screen.getByPlaceholderText('/path/to/media'), { target: { value: '/movie.mkv' } })
    fireEvent.click(screen.getByText('Extract'))
    await waitFor(() => expect(fetch).toHaveBeenCalledWith('/api/extract', expect.any(Object)))
  })
})
