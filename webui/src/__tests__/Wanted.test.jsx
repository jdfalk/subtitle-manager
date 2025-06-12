// file: webui/src/__tests__/Wanted.test.jsx
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { vi, expect } from 'vitest'
import '@testing-library/jest-dom'
import Wanted from '../Wanted.jsx'

describe('Wanted component', () => {
  beforeEach(() => {
    vi.restoreAllMocks()
    global.fetch = vi.fn()
  })

  test('loads wanted list on mount and searches', async () => {
    fetch.mockResolvedValueOnce({ ok: true, json: () => Promise.resolve(['a']) })
    render(<Wanted />)
    await screen.findByText('a')

    fetch.mockResolvedValueOnce({ ok: true, json: () => Promise.resolve(['u']) })
    fireEvent.change(screen.getByPlaceholderText('Media path'), { target: { value: 'f' } })
    fireEvent.click(screen.getByText('Search'))
    await waitFor(() => expect(fetch).toHaveBeenLastCalledWith('/api/search?provider=generic&path=f&lang=en'))
  })
})
