# TASK-02-001: Fix UI Layout and Navigation

<!-- file: docs/tasks/02-ui-fixes/TASK-02-001-fix-layout-navigation.md -->
<!-- version: 1.0.0 -->
<!-- guid: e4f5g6h7-i8j9-0123-defg-456789012345 -->

## 🎯 Objective

Fix the current UI layout and navigation issues based on the TODO.md
requirements and implement the Bazarr-style interface improvements.

## 📋 Acceptance Criteria

- [ ] Fix user management display showing blank usernames
- [ ] Move user management to settings section
- [ ] Implement working back button navigation
- [ ] Add sidebar pinning functionality
- [ ] Reorganize navigation order: Dashboard → Media Library → Wanted → History
      → Settings → System
- [ ] Restructure tools: Move Extract/Translate/Convert to Tools section
- [ ] Fix provider configuration modals and dropdowns
- [ ] Implement card-based settings interface

## 🔍 Current State Analysis

### Known UI Issues from TODO.md

1. **User Management Display**: System/users shows blank usernames
2. **Navigation Structure**: Current order needs reorganization
3. **Missing Back Button**: Navigation history not working
4. **Sidebar UX**: No pinning functionality
5. **Provider Modals**: Configuration dropdowns not working properly
6. **Settings Layout**: Need card-based interface like Bazarr

### Current Navigation Structure

Current: Dashboard → Media Library → Wanted → History → System → Tools Target:
Dashboard → Media Library → Wanted → History → Settings → System

## 🔧 Implementation Steps

### Step 1: Analyze current UI structure

```bash
# Examine current React components
find webui/src -name "*.jsx" -o -name "*.tsx" | head -20
```

Examine these key files:

- `webui/src/App.jsx` - Main application structure
- `webui/src/components/Navigation.jsx` - Navigation sidebar
- `webui/src/components/Settings.jsx` - Settings page
- `webui/src/components/System.jsx` - System page
- `webui/src/components/UserManagement.jsx` - User management

### Step 2: Fix user management display

Update `webui/src/components/UserManagement.jsx`:

```jsx
// OLD - Likely missing user data display
function UserRow({ user }) {
  return (
    <tr>
      <td>{user.username || 'N/A'}</td> // Issue: blank display
      <td>{user.email}</td>
    </tr>
  );
}

// NEW - Fix blank username display
function UserRow({ user }) {
  return (
    <tr>
      <td>{user?.username || user?.name || user?.id || 'Unknown User'}</td>
      <td>{user?.email || 'No email'}</td>
      <td>{user?.role || 'user'}</td>
      <td>
        <span
          className={`badge ${user?.active ? 'badge-success' : 'badge-secondary'}`}
        >
          {user?.active ? 'Active' : 'Inactive'}
        </span>
      </td>
    </tr>
  );
}
```

### Step 3: Reorganize navigation structure

Update `webui/src/components/Navigation.jsx`:

```jsx
// OLD navigation order
const navigationItems = [
  { path: '/', label: 'Dashboard', icon: 'dashboard' },
  { path: '/media', label: 'Media Library', icon: 'video_library' },
  { path: '/wanted', label: 'Wanted', icon: 'search' },
  { path: '/history', label: 'History', icon: 'history' },
  { path: '/system', label: 'System', icon: 'settings' },
  { path: '/tools', label: 'Tools', icon: 'build' },
];

// NEW navigation order with proper organization
const navigationItems = [
  { path: '/', label: 'Dashboard', icon: 'dashboard' },
  { path: '/media', label: 'Media Library', icon: 'video_library' },
  { path: '/wanted', label: 'Wanted', icon: 'search' },
  { path: '/history', label: 'History', icon: 'history' },
  {
    path: '/settings',
    label: 'Settings',
    icon: 'settings',
    children: [
      { path: '/settings/general', label: 'General' },
      { path: '/settings/providers', label: 'Providers' },
      { path: '/settings/languages', label: 'Languages' },
      { path: '/settings/auth', label: 'Authentication' },
      { path: '/settings/users', label: 'Users' },
      { path: '/settings/notifications', label: 'Notifications' },
    ],
  },
  {
    path: '/tools',
    label: 'Tools',
    icon: 'build',
    children: [
      { path: '/tools/extract', label: 'Extract Subtitles' },
      { path: '/tools/translate', label: 'Translate' },
      { path: '/tools/convert', label: 'Convert Format' },
    ],
  },
  { path: '/system', label: 'System', icon: 'computer' },
];
```

### Step 4: Implement sidebar pinning

Add sidebar pinning functionality:

```jsx
// Add to Navigation.jsx
import { useState, useEffect } from 'react';

function Navigation() {
  const [isPinned, setIsPinned] = useState(false);
  const [isCollapsed, setIsCollapsed] = useState(false);

  useEffect(() => {
    const savedPinState = localStorage.getItem('sidebar-pinned');
    if (savedPinState) {
      setIsPinned(savedPinState === 'true');
    }
  }, []);

  const togglePin = () => {
    const newPinState = !isPinned;
    setIsPinned(newPinState);
    localStorage.setItem('sidebar-pinned', newPinState.toString());
  };

  return (
    <nav
      className={`sidebar ${isPinned ? 'pinned' : ''} ${isCollapsed ? 'collapsed' : ''}`}
    >
      <div className="sidebar-header">
        <h3>Subtitle Manager</h3>
        <button className="pin-button" onClick={togglePin}>
          <i className={`fas fa-thumbtack ${isPinned ? 'pinned' : ''}`}></i>
        </button>
      </div>
      {/* Navigation items */}
    </nav>
  );
}
```

### Step 5: Implement back button navigation

Add navigation history handling:

```jsx
// Create webui/src/hooks/useNavigationHistory.js
import { useNavigate, useLocation } from 'react-router-dom';
import { useState, useEffect } from 'react';

export function useNavigationHistory() {
  const navigate = useNavigate();
  const location = useLocation();
  const [history, setHistory] = useState([]);

  useEffect(() => {
    setHistory(prev => {
      const newHistory = [...prev];
      if (newHistory[newHistory.length - 1] !== location.pathname) {
        newHistory.push(location.pathname);
        // Keep only last 10 entries
        return newHistory.slice(-10);
      }
      return newHistory;
    });
  }, [location.pathname]);

  const goBack = () => {
    if (history.length > 1) {
      const previous = history[history.length - 2];
      navigate(previous);
    } else {
      navigate('/');
    }
  };

  const canGoBack = history.length > 1;

  return { goBack, canGoBack, history };
}

// Use in components
function PageHeader({ title }) {
  const { goBack, canGoBack } = useNavigationHistory();

  return (
    <div className="page-header">
      <button className="back-button" onClick={goBack} disabled={!canGoBack}>
        <i className="fas fa-arrow-left"></i> Back
      </button>
      <h1>{title}</h1>
    </div>
  );
}
```

### Step 6: Fix provider configuration modals

Update provider configuration modals:

```jsx
// Update webui/src/components/ProviderConfig.jsx
function ProviderConfigModal({ provider, isOpen, onClose, onSave }) {
  const [config, setConfig] = useState(provider?.config || {});
  const [availableProviders, setAvailableProviders] = useState([]);

  useEffect(() => {
    // Load available providers
    fetch('/api/providers/available')
      .then(res => res.json())
      .then(data => setAvailableProviders(data))
      .catch(err => console.error('Failed to load providers:', err));
  }, []);

  const handleSave = () => {
    fetch(`/api/providers/${provider.id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(config),
    })
      .then(res => res.json())
      .then(data => {
        onSave(data);
        onClose();
      })
      .catch(err => console.error('Failed to save provider:', err));
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <div className="provider-config-modal">
        <h3>Configure {provider?.name}</h3>

        <div className="form-group">
          <label>Provider Type</label>
          <select
            value={config.type || ''}
            onChange={e => setConfig({ ...config, type: e.target.value })}
          >
            <option value="">Select Provider</option>
            {availableProviders.map(p => (
              <option key={p.id} value={p.id}>
                {p.name}
              </option>
            ))}
          </select>
        </div>

        {/* Dynamic configuration fields based on provider type */}
        {config.type && (
          <ProviderSpecificConfig
            type={config.type}
            config={config}
            onChange={setConfig}
          />
        )}

        <div className="modal-actions">
          <button onClick={onClose}>Cancel</button>
          <button onClick={handleSave} className="btn-primary">
            Save
          </button>
        </div>
      </div>
    </Modal>
  );
}
```

### Step 7: Implement card-based settings interface

Create card-based settings layout:

```jsx
// Update webui/src/components/Settings.jsx
function Settings() {
  return (
    <div className="settings-page">
      <PageHeader title="Settings" />

      <div className="settings-grid">
        <SettingsCard
          title="General"
          description="Basic application settings and preferences"
          icon="settings"
          path="/settings/general"
        />

        <SettingsCard
          title="Providers"
          description="Configure subtitle providers and sources"
          icon="cloud_download"
          path="/settings/providers"
        />

        <SettingsCard
          title="Languages"
          description="Language preferences and profiles"
          icon="language"
          path="/settings/languages"
        />

        <SettingsCard
          title="Authentication"
          description="User authentication and security settings"
          icon="security"
          path="/settings/auth"
        />

        <SettingsCard
          title="Users"
          description="Manage user accounts and permissions"
          icon="people"
          path="/settings/users"
        />

        <SettingsCard
          title="Notifications"
          description="Configure alerts and notification channels"
          icon="notifications"
          path="/settings/notifications"
        />
      </div>
    </div>
  );
}

function SettingsCard({ title, description, icon, path }) {
  const navigate = useNavigate();

  return (
    <div className="settings-card" onClick={() => navigate(path)}>
      <div className="card-icon">
        <i className={`material-icons ${icon}`}></i>
      </div>
      <div className="card-content">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
      <div className="card-arrow">
        <i className="fas fa-chevron-right"></i>
      </div>
    </div>
  );
}
```

### Step 8: Add CSS styling

Create/update `webui/src/styles/components.css`:

```css
/* Sidebar pinning styles */
.sidebar {
  position: fixed;
  left: 0;
  top: 0;
  width: 250px;
  height: 100vh;
  background: #2c3e50;
  transition: transform 0.3s ease;
  z-index: 1000;
}

.sidebar:not(.pinned) {
  transform: translateX(-100%);
}

.sidebar:not(.pinned):hover {
  transform: translateX(0);
}

.sidebar.pinned {
  transform: translateX(0);
}

.pin-button {
  background: none;
  border: none;
  color: white;
  font-size: 16px;
  cursor: pointer;
  padding: 5px;
}

.pin-button i.pinned {
  transform: rotate(45deg);
  color: #3498db;
}

/* Settings card grid */
.settings-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  padding: 20px;
}

.settings-card {
  background: white;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition:
    transform 0.2s ease,
    box-shadow 0.2s ease;
  display: flex;
  align-items: center;
  gap: 15px;
}

.settings-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.card-icon {
  font-size: 48px;
  color: #3498db;
}

.card-content h3 {
  margin: 0 0 8px 0;
  color: #2c3e50;
}

.card-content p {
  margin: 0;
  color: #7f8c8d;
  font-size: 14px;
}

.card-arrow {
  margin-left: auto;
  color: #bdc3c7;
}

/* Back button */
.back-button {
  background: none;
  border: 1px solid #bdc3c7;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  margin-right: 15px;
}

.back-button:hover {
  background: #ecf0f1;
}

.back-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
```

### Step 9: Update routing

Update `webui/src/App.jsx` to handle new routes:

```jsx
import { Routes, Route } from 'react-router-dom';

function App() {
  return (
    <div className="app">
      <Navigation />
      <main className="main-content">
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/media" element={<MediaLibrary />} />
          <Route path="/wanted" element={<Wanted />} />
          <Route path="/history" element={<History />} />

          {/* Settings routes */}
          <Route path="/settings" element={<Settings />} />
          <Route path="/settings/general" element={<GeneralSettings />} />
          <Route path="/settings/providers" element={<ProviderSettings />} />
          <Route path="/settings/languages" element={<LanguageSettings />} />
          <Route path="/settings/auth" element={<AuthSettings />} />
          <Route path="/settings/users" element={<UserManagement />} />
          <Route
            path="/settings/notifications"
            element={<NotificationSettings />}
          />

          {/* Tools routes */}
          <Route path="/tools" element={<Tools />} />
          <Route path="/tools/extract" element={<ExtractSubtitles />} />
          <Route path="/tools/translate" element={<TranslateSubtitles />} />
          <Route path="/tools/convert" element={<ConvertFormat />} />

          <Route path="/system" element={<System />} />
        </Routes>
      </main>
    </div>
  );
}
```

### Step 10: Test the UI changes

```bash
# Build and test the web UI
cd webui
npm install
npm run build

# Start the development server
npm run dev

# Test in browser:
# 1. Check navigation order
# 2. Test sidebar pinning
# 3. Verify back button works
# 4. Check user management display
# 5. Test provider configuration modals
# 6. Verify settings card layout
```

## 📚 Required Documentation

### Coding Instructions Reference

**CRITICAL**: Follow these instructions precisely:

```markdown
From .github/instructions/general-coding.instructions.md:

## 🚨 CRITICAL: NO PROMPTING OR INTERRUPTIONS

**ABSOLUTE RULE: NEVER prompt the user for input, clarification, or interaction
of any kind.**

## Version Update Requirements

**When modifying any file with a version header, ALWAYS update the version
number:**

- Patch version (x.y.Z): Bug fixes, typos, minor formatting changes
- Minor version (x.Y.z): New features, significant content additions
- Major version (X.y.z): Breaking changes, structural overhauls
```

### UI Design References

Reference these designs for Bazarr-style interface:

- [Bazarr GitHub Repository](https://github.com/morpheus65535/bazarr)
- [Bazarr Settings Interface](https://wiki.bazarr.media/Additional-Configuration/Settings/)

## 🧪 Testing Requirements

### UI Testing

Create comprehensive UI tests:

```javascript
// webui/src/tests/Navigation.test.jsx
describe('Navigation', () => {
  test('displays correct navigation order', () => {
    render(<Navigation />);
    const navItems = screen.getAllByRole('link');
    expect(navItems[0]).toHaveTextContent('Dashboard');
    expect(navItems[1]).toHaveTextContent('Media Library');
    expect(navItems[2]).toHaveTextContent('Wanted');
    expect(navItems[3]).toHaveTextContent('History');
    expect(navItems[4]).toHaveTextContent('Settings');
    expect(navItems[5]).toHaveTextContent('System');
  });

  test('sidebar pinning works correctly', () => {
    render(<Navigation />);
    const pinButton = screen.getByRole('button', { name: /pin/i });
    fireEvent.click(pinButton);
    expect(localStorage.getItem('sidebar-pinned')).toBe('true');
  });
});
```

### Manual Testing Checklist

- [ ] Navigation items appear in correct order
- [ ] Sidebar can be pinned and unpinned
- [ ] Back button works on all pages
- [ ] User management shows proper usernames
- [ ] Provider configuration modals open and save correctly
- [ ] Settings cards are clickable and navigate properly
- [ ] Tools section contains Extract/Translate/Convert options
- [ ] Mobile responsiveness works

## 🎯 Success Metrics

- [ ] All navigation items appear in correct order
- [ ] User management displays proper usernames and data
- [ ] Sidebar pinning persists across sessions
- [ ] Back button works on all pages
- [ ] Provider configuration modals functional
- [ ] Settings interface matches Bazarr card-based design
- [ ] No console errors in browser
- [ ] Mobile-friendly responsive design

## 🚨 Common Pitfalls

1. **React Router Issues**: Ensure all routes are properly configured
2. **State Management**: Use proper state management for UI state
3. **CSS Conflicts**: Watch for existing CSS that might conflict
4. **Local Storage**: Handle cases where localStorage is not available
5. **Mobile Responsiveness**: Test on various screen sizes

## 📖 Additional Resources

- [React Router Documentation](https://reactrouter.com/docs)
- [CSS Grid Layout Guide](https://css-tricks.com/snippets/css/complete-guide-grid/)
- [Bazarr UI Screenshots](https://wiki.bazarr.media/)
- [General Coding Instructions](../../../.github/instructions/general-coding.instructions.md)

## 🔄 Related Tasks

- **TASK-02-002**: Implement Selenium-based E2E testing (will test these UI
  changes)
- **TASK-03-001**: Provider system improvements (depends on fixed provider
  modals)
- **TASK-04-001**: Authentication UI improvements (uses new settings structure)

## 📝 Notes for AI Agent

- Focus on user experience and navigation flow
- Standard web development practices - no special tooling required
- Test thoroughly in browser before marking complete
- Pay attention to responsive design for mobile devices
- Ensure all existing functionality remains working
- Follow React best practices for component structure
- If any step fails, document the error and continue with remaining steps where
  possible
