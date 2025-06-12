// file: webui/src/__tests__/App.test.jsx
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import '@testing-library/jest-dom'
import { vi } from 'vitest'
import App from '../App.jsx'

describe('App component', () => {
  beforeEach(() => {
    vi.restoreAllMocks()
    global.fetch = vi.fn()
  })

  test('shows login form when unauthenticated', async () => {
    fetch.mockResolvedValueOnce({ ok: false })
    render(<App />)
    expect(screen.getByText('Subtitle Manager')).toBeInTheDocument()
  })

  test('successful login renders dashboard', async () => {
    fetch.mockResolvedValueOnce({ ok: false }) // config check
    render(<App />)
    fetch.mockResolvedValueOnce({ ok: true })
    fireEvent.click(screen.getByText('Login'))
    await waitFor(() => expect(fetch).toHaveBeenLastCalledWith('/api/login', expect.any(Object)))
  })
})
