// file: webui/src/__tests__/BackButton.test.jsx
// version: 1.0.0
// guid: c1d2e3f4-a5b6-4c7d-8e9f-0123456789ab

import { render, screen, fireEvent } from '@testing-library/react';
import BackButton from '../components/BackButton.jsx';
import { BrowserRouter } from 'react-router-dom';

const mockNavigate = jest.fn();

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate,
}));

describe('BackButton', () => {
  test('navigates back when history exists', () => {
    Object.defineProperty(window, 'history', { value: { length: 2 } });
    render(<BackButton />, { wrapper: BrowserRouter });
    fireEvent.click(screen.getByRole('button', { name: /back/i }));
    expect(mockNavigate).toHaveBeenCalledWith(-1);
  });

  test('navigates home when no history', () => {
    mockNavigate.mockClear();
    Object.defineProperty(window, 'history', { value: { length: 1 } });
    render(<BackButton />, { wrapper: BrowserRouter });
    fireEvent.click(screen.getByRole('button', { name: /back/i }));
    expect(mockNavigate).toHaveBeenCalledWith('/');
  });
});
