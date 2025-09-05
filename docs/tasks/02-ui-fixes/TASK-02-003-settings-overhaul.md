<!-- file: docs/tasks/02-ui-fixes/TASK-02-003-settings-overhaul.md -->
<!-- version: 1.0.0 -->
<!-- guid: b1c2d3e4-f5g6-h7i8-j9k0-l1m2n3o4p5q6 -->

# TASK-02-003: Settings Page Overhaul

## 🎯 Objective

Redesign the settings page with improved organization, modern UI components, and better user experience. Create a tabbed interface with logical grouping of settings and enhanced validation.

## 📋 Acceptance Criteria

- [ ] Fix settings navigation routing issue (settings overview vs tabbed interface)
- [ ] Implement tabbed settings interface with logical grouping
- [ ] Add real-time validation for all form fields
- [ ] Create modern form components with proper styling
- [ ] Implement settings search/filter functionality
- [ ] Add settings backup/restore functionality
- [ ] Include help text and tooltips for complex settings
- [ ] Add settings validation with error messages
- [ ] Implement settings change confirmation dialogs

## 🔍 Current State Analysis

### Critical Settings Navigation Issue

**PRIORITY 1: Routing Confusion**

Currently, there are two different settings components with confusing navigation:

1. **`/settings`** → `SettingsOverview` component
   - Shows a "blocky" three-box layout with cards for General, Providers, Users
   - Feels disconnected and limited
   - When clicking back, goes to the tabbed interface (confusing UX)

2. **`/settings/:section`** → `Settings` component  
   - Shows the proper tabbed interface with full functionality
   - Contains Providers, General, Database, Authentication, Notifications, Users, Tags, Webhooks tabs
   - This is the REAL settings page users want

**Problem**: When users click "Settings" in navigation, they get the overview page instead of the main tabbed interface.

**Solution**: Either:
- Option A: Make `/settings` route directly to the tabbed `Settings` component
- Option B: Improve the `SettingsOverview` to be a proper landing page with better UX
- Option C: Remove `SettingsOverview` entirely and always show tabbed interface

### Existing Settings Implementation

Current settings are likely scattered across different components. Need to evaluate:

1. **Settings Organization**: How settings are currently grouped
2. **Form Components**: Quality of form inputs and validation
3. **User Experience**: Ease of finding and modifying settings
4. **Data Persistence**: How settings are saved and loaded

### Settings Categories to Implement

1. **General Settings**
   - Application preferences
   - Language and locale
   - Theme and appearance
   - Notification preferences

2. **Media Library Settings**
   - Library paths configuration
   - Scan settings and frequency
   - File organization preferences
   - Metadata handling options

3. **Provider Settings**
   - Subtitle provider configuration
   - API keys and authentication
   - Provider priority and selection
   - Download preferences

4. **Integration Settings**
   - Sonarr/Radarr connection settings
   - Plex integration configuration
   - External service connections
   - Webhook configurations

5. **Advanced Settings**
   - Database configuration
   - Logging and debugging
   - Performance tuning
   - System maintenance

## 🎨 UI/UX Design

### Layout Structure

```
┌─────────────────────────────────────────────────────────┐
│ Settings                                    [Save] [Reset] │
├─────────────┬───────────────────────────────────────────┤
│ [General]   │ General Settings                           │
│ [Library]   │ ┌─────────────────────────────────────────┐ │
│ [Providers] │ │ Application Language: [Dropdown]       │ │
│ [Integration│ │ Theme: [Dark/Light/Auto] [Radio]       │ │
│ [Advanced]  │ │ Notifications: [✓] Enable              │ │
│             │ │ Auto-update: [✓] Check for updates     │ │
│             │ └─────────────────────────────────────────┘ │
│             │                                           │
│             │ [Apply Changes]                           │
└─────────────┴───────────────────────────────────────────┘
```

### Component Design

- **Tab Navigation**: Material-UI Tabs with icons
- **Form Sections**: Grouped with Paper/Card components
- **Input Validation**: Real-time feedback with error states
- **Action Buttons**: Consistent save/reset/cancel actions
- **Help System**: Tooltips and expandable help sections

## 🔧 Implementation Steps

### Step 0: Fix Settings Navigation Routing (PRIORITY)

**IMMEDIATE FIX NEEDED**: The current settings navigation is confusing users.

**Current Issue:**
- `/settings` → Shows basic 3-card overview (`SettingsOverview`)  
- `/settings/:section` → Shows full tabbed interface (`Settings`)
- Users expect `/settings` to show the main settings page

**Recommended Solution** (Option A):
```jsx
// webui/src/App.jsx - Update routing
<Route path="/settings" element={<Settings backendAvailable={backendAvailable} />} />
<Route path="/settings/:section" element={<Settings backendAvailable={backendAvailable} />} />
// Remove the SettingsOverview route entirely
```

**Alternative Solution** (Option B): 
```jsx
// webui/src/components/SettingsOverview.jsx - Make it a proper landing page
export default function SettingsOverview() {
  const navigate = useNavigate();
  
  // Immediately redirect to main settings with providers tab
  useEffect(() => {
    navigate('/settings/providers', { replace: true });
  }, [navigate]);
  
  return <LoadingComponent message="Loading Settings..." />;
}
```

**Files to modify:**
- `webui/src/App.jsx` (routing)
- Consider removing `webui/src/components/SettingsOverview.jsx` entirely

### Step 1: Create Settings Context and State Management

```jsx
// webui/src/contexts/SettingsContext.jsx
import { createContext, useContext, useReducer } from 'react';

const SettingsContext = createContext();

export const useSettings = () => {
  const context = useContext(SettingsContext);
  if (!context) {
    throw new Error('useSettings must be used within SettingsProvider');
  }
  return context;
};

const settingsReducer = (state, action) => {
  switch (action.type) {
    case 'LOAD_SETTINGS':
      return { ...state, ...action.payload, loading: false };
    case 'UPDATE_SETTING':
      return {
        ...state,
        [action.category]: {
          ...state[action.category],
          [action.key]: action.value
        },
        hasChanges: true
      };
    case 'SAVE_SETTINGS':
      return { ...state, hasChanges: false, saving: false };
    case 'RESET_SETTINGS':
      return { ...action.payload, hasChanges: false };
    default:
      return state;
  }
};

export const SettingsProvider = ({ children }) => {
  const [state, dispatch] = useReducer(settingsReducer, {
    general: {},
    library: {},
    providers: {},
    integration: {},
    advanced: {},
    loading: true,
    saving: false,
    hasChanges: false
  });

  // Settings API functions
  const loadSettings = async () => {
    try {
      const response = await fetch('/api/settings');
      const settings = await response.json();
      dispatch({ type: 'LOAD_SETTINGS', payload: settings });
    } catch (error) {
      console.error('Failed to load settings:', error);
    }
  };

  const updateSetting = (category, key, value) => {
    dispatch({ type: 'UPDATE_SETTING', category, key, value });
  };

  const saveSettings = async () => {
    dispatch({ type: 'SET_SAVING', payload: true });
    try {
      await fetch('/api/settings', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(state)
      });
      dispatch({ type: 'SAVE_SETTINGS' });
    } catch (error) {
      console.error('Failed to save settings:', error);
    }
  };

  return (
    <SettingsContext.Provider value={{
      ...state,
      loadSettings,
      updateSetting,
      saveSettings
    }}>
      {children}
    </SettingsContext.Provider>
  );
};
```

### Step 2: Create Main Settings Component

```jsx
// webui/src/components/Settings/SettingsPage.jsx
import React, { useState, useEffect } from 'react';
import {
  Box,
  Tabs,
  Tab,
  Paper,
  Button,
  Typography,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Alert
} from '@mui/material';
import {
  Settings as SettingsIcon,
  Folder as FolderIcon,
  CloudDownload as ProvidersIcon,
  Link as IntegrationIcon,
  Code as AdvancedIcon
} from '@mui/icons-material';

import { useSettings } from '../../contexts/SettingsContext';
import GeneralSettings from './GeneralSettings';
import LibrarySettings from './LibrarySettings';
import ProviderSettings from './ProviderSettings';
import IntegrationSettings from './IntegrationSettings';
import AdvancedSettings from './AdvancedSettings';

const TabPanel = ({ children, value, index, ...other }) => (
  <div
    role="tabpanel"
    hidden={value !== index}
    id={`settings-tabpanel-${index}`}
    aria-labelledby={`settings-tab-${index}`}
    {...other}
  >
    {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
  </div>
);

const SettingsPage = () => {
  const [activeTab, setActiveTab] = useState(0);
  const [showResetDialog, setShowResetDialog] = useState(false);
  const { hasChanges, saveSettings, resetSettings, loading } = useSettings();

  const tabs = [
    { label: 'General', icon: <SettingsIcon />, component: GeneralSettings },
    { label: 'Library', icon: <FolderIcon />, component: LibrarySettings },
    { label: 'Providers', icon: <ProvidersIcon />, component: ProviderSettings },
    { label: 'Integration', icon: <IntegrationIcon />, component: IntegrationSettings },
    { label: 'Advanced', icon: <AdvancedIcon />, component: AdvancedSettings }
  ];

  const handleSave = async () => {
    await saveSettings();
  };

  const handleReset = () => {
    setShowResetDialog(true);
  };

  const confirmReset = () => {
    resetSettings();
    setShowResetDialog(false);
  };

  if (loading) {
    return <Box sx={{ p: 3 }}>Loading settings...</Box>;
  }

  return (
    <Box sx={{ width: '100%', height: '100%' }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 2 }}>
        <Typography variant="h4">Settings</Typography>
        <Box>
          <Button 
            variant="outlined" 
            onClick={handleReset} 
            disabled={!hasChanges}
            sx={{ mr: 1 }}
          >
            Reset
          </Button>
          <Button 
            variant="contained" 
            onClick={handleSave}
            disabled={!hasChanges}
          >
            Save Changes
          </Button>
        </Box>
      </Box>

      {hasChanges && (
        <Alert severity="info" sx={{ mb: 2 }}>
          You have unsaved changes. Don't forget to save!
        </Alert>
      )}

      <Paper sx={{ width: '100%' }}>
        <Tabs
          value={activeTab}
          onChange={(e, newValue) => setActiveTab(newValue)}
          variant="scrollable"
          scrollButtons="auto"
        >
          {tabs.map((tab, index) => (
            <Tab
              key={index}
              icon={tab.icon}
              label={tab.label}
              iconPosition="start"
              id={`settings-tab-${index}`}
              aria-controls={`settings-tabpanel-${index}`}
            />
          ))}
        </Tabs>

        {tabs.map((tab, index) => (
          <TabPanel key={index} value={activeTab} index={index}>
            <tab.component />
          </TabPanel>
        ))}
      </Paper>

      <Dialog open={showResetDialog} onClose={() => setShowResetDialog(false)}>
        <DialogTitle>Reset Settings</DialogTitle>
        <DialogContent>
          Are you sure you want to reset all settings to their default values?
          This action cannot be undone.
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setShowResetDialog(false)}>Cancel</Button>
          <Button onClick={confirmReset} color="error">Reset</Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default SettingsPage;
```

### Step 3: Create Individual Settings Components

Create separate components for each settings category:

- `GeneralSettings.jsx` - App preferences, theme, language
- `LibrarySettings.jsx` - Media library configuration
- `ProviderSettings.jsx` - Subtitle provider setup
- `IntegrationSettings.jsx` - External service connections
- `AdvancedSettings.jsx` - System and debug settings

### Step 4: Implement Form Validation

```jsx
// webui/src/hooks/useSettingsValidation.js
import { useState, useCallback } from 'react';

export const useSettingsValidation = (initialValues = {}) => {
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});

  const validateField = useCallback((name, value, rules = {}) => {
    let error = '';

    if (rules.required && (!value || value.toString().trim() === '')) {
      error = 'This field is required';
    } else if (rules.minLength && value.length < rules.minLength) {
      error = `Minimum ${rules.minLength} characters required`;
    } else if (rules.pattern && !rules.pattern.test(value)) {
      error = rules.message || 'Invalid format';
    } else if (rules.custom && !rules.custom(value)) {
      error = rules.customMessage || 'Invalid value';
    }

    setErrors(prev => ({
      ...prev,
      [name]: error
    }));

    return error === '';
  }, []);

  const validateForm = useCallback((values, validationRules) => {
    let isValid = true;
    const newErrors = {};

    Object.keys(validationRules).forEach(field => {
      const fieldValid = validateField(field, values[field], validationRules[field]);
      if (!fieldValid) {
        isValid = false;
      }
    });

    return isValid;
  }, [validateField]);

  const setFieldTouched = useCallback((name) => {
    setTouched(prev => ({
      ...prev,
      [name]: true
    }));
  }, []);

  const resetValidation = useCallback(() => {
    setErrors({});
    setTouched({});
  }, []);

  return {
    errors,
    touched,
    validateField,
    validateForm,
    setFieldTouched,
    resetValidation
  };
};
```

### Step 5: Backend API Integration

Ensure the backend provides proper settings endpoints:

```go
// pkg/webserver/settings.go
func (s *Server) handleGetSettings(w http.ResponseWriter, r *http.Request) {
    // Load and return current settings
    settings := s.config.GetAllSettings()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(settings)
}

func (s *Server) handleUpdateSettings(w http.ResponseWriter, r *http.Request) {
    var settings map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&settings); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate and save settings
    if err := s.config.UpdateSettings(settings); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
```

## 🧪 Testing Requirements

### Unit Tests
- Settings context and state management
- Form validation logic
- Individual setting component behavior

### Integration Tests
- Settings API endpoints
- Data persistence
- Form submission workflows

### E2E Tests
- Complete settings configuration workflow
- Settings backup/restore functionality
- Cross-tab settings synchronization

## 📚 Documentation

### User Documentation
- Settings reference guide
- Configuration best practices
- Troubleshooting common issues

### Developer Documentation
- Settings architecture overview
- Adding new settings categories
- Validation patterns and examples

## 🔄 Future Enhancements

1. **Settings Templates**: Predefined configuration templates
2. **Import/Export**: Settings backup and sharing
3. **Advanced Search**: Search across all settings
4. **Settings History**: Track and revert changes
5. **Real-time Sync**: Multi-device settings synchronization

## ✅ Success Metrics

- Improved settings discoverability
- Reduced configuration errors
- Better user onboarding experience
- Increased settings adoption rates
- Faster support resolution for configuration issues

## 📋 Task Dependencies

- **Depends on**: TASK-02-001 (Navigation completed)
- **Blocks**: TASK-02-004 (Dashboard may use settings)
- **Related**: TASK-02-006 (User management settings)

## 🎯 Completion Checklist

- [ ] Settings context implementation
- [ ] Tabbed interface with all categories
- [ ] Form validation and error handling
- [ ] Backend API integration
- [ ] Settings persistence
- [ ] Help text and documentation
- [ ] Unit and integration tests
- [ ] User acceptance testing
- [ ] Performance optimization
- [ ] Accessibility compliance
