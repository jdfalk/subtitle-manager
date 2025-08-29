// file: webui/src/components/Navigation.jsx
// version: 1.0.0
// guid: 86da4298-39f0-4f62-ad2e-17bc3fffdf28

import { useEffect, useState } from 'react';
import { NavLink } from 'react-router-dom';

export const navigationItems = [
  { path: '/', label: 'Dashboard' },
  { path: '/media', label: 'Media Library' },
  { path: '/wanted', label: 'Wanted' },
  { path: '/history', label: 'History' },
  { path: '/settings', label: 'Settings' },
  { path: '/system', label: 'System' },
];

export default function Navigation() {
  const [pinned, setPinned] = useState(
    localStorage.getItem('sidebar-pinned') === 'true'
  );

  useEffect(() => {
    localStorage.setItem('sidebar-pinned', pinned ? 'true' : 'false');
  }, [pinned]);

  const togglePin = () => {
    setPinned(!pinned);
  };

  return (
    <aside className={pinned ? 'pinned' : ''}>
      <button aria-label="pin" onClick={togglePin}>
        {pinned ? 'Unpin' : 'Pin'}
      </button>
      <nav>
        <ul>
          {navigationItems.map(item => (
            <li key={item.path}>
              <NavLink to={item.path}>{item.label}</NavLink>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  );
}
