// file: webui/src/__tests__/Navigation.test.jsx
// version: 1.0.0
// guid: f0ecc3a3-dc59-42c9-b193-d6f44f6a1c4d

import '@testing-library/jest-dom/vitest';
import { fireEvent, render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import Navigation, { navigationItems } from '../components/Navigation.jsx';

describe('Navigation component', () => {
  test('displays correct navigation order', () => {
    render(
      <MemoryRouter>
        <Navigation />
      </MemoryRouter>
    );
    const links = screen.getAllByRole('link');
    const texts = links.map(link => link.textContent);
    expect(texts).toEqual(navigationItems.map(item => item.label));
  });

  test('sidebar pinning works correctly', () => {
    render(
      <MemoryRouter>
        <Navigation />
      </MemoryRouter>
    );
    const pinButton = screen.getByRole('button', { name: /pin/i });
    fireEvent.click(pinButton);
    expect(localStorage.getItem('sidebar-pinned')).toBe('true');
  });
});
