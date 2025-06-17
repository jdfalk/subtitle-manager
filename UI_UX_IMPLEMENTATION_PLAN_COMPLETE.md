# UI/UX Implementation Plan - Complete Guide

This comprehensive guide combines all sections of the Subtitle Manager UI/UX implementation plan into a single document for easy reference during development.

**Generated on:** Mon Jun 16 18:42:53 EDT 2025
**Total Project Scope:** 82-114 hours for complete UI/UX overhaul
**Sections:** 6 (Summary + 5 Implementation Sections)

---

# Section 0: Project Summary and Overview

This document provides a comprehensive overview of all UI/UX improvements identified for the Subtitle Manager web application, organized into 5 manageable implementation sections.

## Overview

The Subtitle Manager backend is production-ready with comprehensive subtitle management capabilities. However, the frontend requires significant UI/UX improvements to match the backend's sophistication and provide a professional user experience comparable to Bazarr.

## Complete Issue List and Section Mapping

### Section 1: User Management and Navigation Basics (Issues 1-3)

**Time Estimate: 9-13 hours**

- **Fix user management display** - System/users shows blank usernames
- **Move user management to Settings** - Relocate users interface from System to Settings
- **Implement working back button** - Add proper navigation history and back functionality

### Section 2: Navigation Improvements (Issues 4-6)

**Time Estimate: 9-13 hours**

- **Add sidebar pinning** - Allow users to pin/unpin the sidebar for better UX
- **Reorganize navigation order** - Dashboard → Media Library → Wanted → History → Settings → System
- **Restructure tools** - Move Extract/Translate/Convert to Tools section or integrate into System

### Section 3: Settings Page Enhancement (Issues 7-8)

**Time Estimate: 14-20 hours**

- **Enhance General Settings** - Add Bazarr-compatible settings (Host, Proxy, Updates, Logging, Backups, Analytics)
- **Improve Database Settings** - Add comprehensive database information and management options

### Section 4: Authentication and Notifications (Issues 9-12)

**Time Estimate: 22-30 hours**

- **Redesign Authentication Page** - Card-based UI for each auth method with enable/disable toggles
- **Add OAuth2 management** - Generate/regenerate client ID/secret, reset to defaults
- **Enhance Notifications** - Card-based interface for each notification method with test buttons
- **Implement notification testing** - Add test buttons and reset functionality for notifications

### Section 5: Final Components and Provider Fixes (Issues 13-18)

**Time Estimate: 28-38 hours**

- **Create OAuth2 API endpoints** - Backend API support for credential management
- **Add Languages Page** - Global language settings for subtitle downloads (like Bazarr)
- **Add Scheduler Settings** - Integration into general settings or separate page
- **Fix provider configuration modals** - Proper provider selection dropdowns and configuration options
- **Improve embedded provider config** - Working dropdown and proper configuration display
- **Implement global language settings** - Move language settings from provider-level to global

## Implementation Files Created

Each section has its own detailed implementation plan:

- `UI_UX_IMPLEMENTATION_PLAN.md` - Section 1: User Management & Navigation Basics
- `UI_UX_IMPLEMENTATION_PLAN_SECTION_2.md` - Section 2: Navigation Improvements
- `UI_UX_IMPLEMENTATION_PLAN_SECTION_3.md` - Section 3: Settings Page Enhancement
- `UI_UX_IMPLEMENTATION_PLAN_SECTION_4.md` - Section 4: Authentication & Notifications
- `UI_UX_IMPLEMENTATION_PLAN_SECTION_5.md` - Section 5: Final Components & Provider Fixes

## Total Project Scope

**Total Time Estimate: 82-114 hours for junior developers**

### Priority Implementation Order

1. **High Priority** (Sections 1-2): Basic navigation and user management fixes

   - Essential for basic usability
   - Foundation for other improvements
   - 18-26 hours

2. **Medium Priority** (Sections 3-4): Settings and authentication enhancements

   - Significant UX improvements
   - Professional appearance
   - 36-50 hours

3. **Lower Priority** (Section 5): Advanced features and provider fixes
   - Nice-to-have features
   - Bazarr compatibility
   - 28-38 hours

## Key Features Delivered

### User Experience Improvements

- Professional card-based interfaces throughout
- Proper navigation with back button support
- Customizable sidebar pinning
- Logical navigation organization
- Comprehensive settings management

### Settings System Overhaul

- Bazarr-compatible general settings
- Database management interface
- Authentication system redesign
- OAuth2 credential management
- Notification system with testing
- Global language preferences
- Scheduler configuration

### Provider System Enhancement

- Fixed configuration modals
- Proper provider selection
- Global language settings
- Provider-specific configuration fields
- Better error handling and validation

### Technical Implementation

- Material Design 3 compliance
- Responsive design patterns
- Accessibility improvements
- Error handling and validation
- Backend API endpoint additions
- Proper state management

## Development Guidelines for Implementation

### Code Quality Standards

- Follow existing Material Design 3 patterns
- Maintain accessibility standards (WCAG 2.1 AA)
- Implement proper error handling
- Use consistent component patterns
- Add comprehensive JSDoc documentation

### Testing Requirements

- Test all components in light and dark modes
- Verify responsive behavior on mobile devices
- Test with backend unavailable scenarios
- Validate accessibility with screen readers
- Test error states and validation

### Backend Integration

- Use existing API patterns where possible
- Add new endpoints as specified in each section
- Maintain backwards compatibility
- Follow existing authentication patterns
- Use Viper for configuration management

## Success Criteria

Upon completion of all sections, the frontend will:

1. **Match Backend Quality**: Professional UI matching the sophisticated backend
2. **Bazarr Compatibility**: Feature parity with Bazarr's UI/UX patterns
3. **Production Ready**: Suitable for production deployment
4. **User Friendly**: Intuitive navigation and clear information architecture
5. **Accessible**: WCAG 2.1 AA compliance throughout
6. **Responsive**: Works across all device sizes
7. **Error Resilient**: Graceful handling of backend unavailability

## Maintenance and Future Enhancements

### Short-term Maintenance

- Monitor user feedback for additional UX issues
- Performance optimization based on usage patterns
- Browser compatibility testing
- Mobile app considerations

### Future Enhancement Opportunities

- Advanced theme variants
- Keyboard shortcuts
- Bulk operation improvements
- Advanced filtering and search
- Custom dashboard widgets
- Integration with additional media servers

## Conclusion

This comprehensive implementation plan addresses all identified UI/UX issues in the Subtitle Manager frontend. The modular approach allows for incremental implementation while maintaining functionality throughout the development process. Each section includes detailed code examples, testing procedures, and realistic time estimates to ensure successful completion by developers of varying experience levels.

The total scope represents a significant but manageable improvement that will transform the Subtitle Manager from a functional but basic interface to a professional, user-friendly application ready for production deployment.

---

# Section 1: User Management and Navigation Basics

This document provides a comprehensive implementation plan for resolving the UI/UX issues identified in the Subtitle Manager web interface. The plan is designed to help junior developers understand and implement the required changes step by step.

## Overview

The Subtitle Manager web UI requires significant improvements to match Bazarr's functionality and provide a better user experience. This implementation plan covers 17 main issues categorized into:

1. **Navigation and User Experience** (Issues 1-5)
2. **Settings Page Enhancements** (Issues 6-16)
3. **Provider System Improvements** (Issues 17-18)

## Prerequisites

Before starting implementation, ensure you have:

- Node.js 18+ installed
- React development environment set up
- Access to the existing codebase
- Basic understanding of React, Material-UI, and JavaScript/TypeScript

## Development Workflow

1. Create feature branches for each major section
2. Implement changes incrementally
3. Test each change thoroughly
4. Update relevant documentation
5. Create pull requests for review

---

## Section 1: Navigation and User Experience Fixes

### Issue 1: Fix User Management Display (Blank Usernames)

**Problem**: The system/users page shows users but usernames appear blank.

**Root Cause**: The UserManagement component is not properly fetching or displaying user data.

**Current File**: `webui/src/UserManagement.jsx`

**Solution**:

#### Code Changes Required

1. **Update UserManagement Component**:

```jsx
// file: webui/src/UserManagement.jsx
import { useEffect, useState } from "react";
import {
  Box,
  Button,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
  Alert,
  CircularProgress,
  Paper,
  Chip,
} from "@mui/material";

/**
 * UserManagement displays all users and allows password resets.
 * Fixed to properly display user data including usernames.
 */
export default function UserManagement({ backendAvailable = true }) {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const loadUsers = async () => {
    if (!backendAvailable) {
      setLoading(false);
      return;
    }

    try {
      setLoading(true);
      setError(null);

      // Add proper error handling and debugging
      const res = await fetch("/api/users", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include", // Ensure cookies are sent
      });

      if (!res.ok) {
        throw new Error(`HTTP ${res.status}: ${res.statusText}`);
      }

      const userData = await res.json();

      // Debug log to see what we're getting
      console.log("User data received:", userData);

      // Ensure we have an array
      const userArray = Array.isArray(userData)
        ? userData
        : userData.users || [];
      setUsers(userArray);
    } catch (error) {
      console.error("Failed to load users:", error);
      setError(`Failed to load users: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  const reset = async (id) => {
    if (!window.confirm("Reset password for this user?")) return;

    try {
      const res = await fetch(`/api/users/${id}/reset`, {
        method: "POST",
        credentials: "include",
      });

      if (res.ok) {
        alert("Password reset and emailed");
        // Reload users to reflect any changes
        await loadUsers();
      } else {
        const errorText = await res.text();
        alert(`Reset failed: ${errorText}`);
      }
    } catch (error) {
      alert(`Reset failed: ${error.message}`);
    }
  };

  useEffect(() => {
    loadUsers();
  }, [backendAvailable]);

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" p={4}>
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        User Management
      </Typography>

      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available. User management features are
          currently disabled.
        </Alert>
      )}

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      {users.length === 0 && !loading && !error ? (
        <Alert severity="info">
          No users found. This might indicate a backend connectivity issue.
        </Alert>
      ) : (
        <Paper variant="outlined">
          <Table size="small">
            <TableHead>
              <TableRow>
                <TableCell>
                  <strong>Username</strong>
                </TableCell>
                <TableCell>
                  <strong>Email</strong>
                </TableCell>
                <TableCell>
                  <strong>Role</strong>
                </TableCell>
                <TableCell>
                  <strong>Status</strong>
                </TableCell>
                <TableCell>
                  <strong>Actions</strong>
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {users.map((user, index) => (
                <TableRow key={user.id || index}>
                  <TableCell>{user.username || user.name || "N/A"}</TableCell>
                  <TableCell>{user.email || "N/A"}</TableCell>
                  <TableCell>
                    <Chip
                      label={user.role || "user"}
                      size="small"
                      color={user.role === "admin" ? "primary" : "default"}
                    />
                  </TableCell>
                  <TableCell>
                    <Chip
                      label={user.active !== false ? "Active" : "Inactive"}
                      size="small"
                      color={user.active !== false ? "success" : "default"}
                    />
                  </TableCell>
                  <TableCell>
                    <Button
                      size="small"
                      onClick={() => reset(user.id)}
                      disabled={!user.id}
                      variant="outlined"
                    >
                      Reset Password
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </Paper>
      )}

      {/* Debug information - remove in production */}
      {process.env.NODE_ENV === "development" && (
        <Box mt={2}>
          <Typography variant="caption" color="text.secondary">
            Debug: {users.length} users loaded
          </Typography>
        </Box>
      )}
    </Box>
  );
}
```

#### Testing Steps

1. Navigate to System → Users tab
2. Verify usernames are now displayed
3. Check that all user data fields show correctly
4. Test the reset password functionality
5. Verify loading and error states work properly

**Time Estimate**: 2-3 hours

---

### Issue 2: Move User Management to Settings

**Problem**: User management is currently in the System page but should be part of Settings.

**Solution**: Move the UserManagement component from System to Settings and update navigation.

#### Code Changes Required

1. **Update Settings.jsx to include User Management tab**:

```jsx
// file: webui/src/Settings.jsx

// Add import at the top
import UserManagement from "./UserManagement.jsx";
import { People as UsersIcon } from "@mui/icons-material";

// In the tabs array, add:
const tabs = [
  {
    label: "Providers",
    icon: <ProvidersIcon />,
    component: renderProvidersTab,
  },
  {
    label: "General",
    icon: <GeneralIcon />,
    component: () => (
      <GeneralSettings
        config={_config}
        onSave={saveSettings}
        backendAvailable={backendAvailable}
      />
    ),
  },
  // ... existing tabs ...
  {
    label: "Users",
    icon: <UsersIcon />,
    component: () => <UserManagement backendAvailable={backendAvailable} />,
  },
];
```

2. **Update System.jsx to remove Users tab**:

```jsx
// file: webui/src/System.jsx

// Remove UserManagement import and Users tab
// Update the tabs section:

return (
  <Box>
    {/* ... existing content ... */}

    <Tabs value={tab} onChange={(e, v) => setTab(v)} sx={{ mb: 3 }}>
      <Tab label="System" />
      {/* Remove Users tab */}
    </Tabs>

    {tab === 0 && (
      <Box
        display="grid"
        gridTemplateColumns={{ xs: "1fr", md: "1fr 1fr" }}
        gap={3}
      >
        {/* ... existing system content ... */}
      </Box>
    )}

    {/* Remove: {tab === 1 && <UserManagement />} */}
  </Box>
);
```

**Time Estimate**: 1-2 hours

---

## Next Steps

This concludes Section 1 of the implementation plan. The next sections will cover:

- **Section 2**: Navigation improvements (back button, sidebar pinning, reordering)
- **Section 3**: Settings page enhancements
- **Section 4**: Provider system improvements

Each section will include detailed code examples, testing procedures, and implementation guidance for junior developers.

---

# Section 2: Navigation Improvements

This section covers issues 3-5: implementing back button functionality, sidebar pinning, and reorganizing navigation order.

## Issue 3: Implement Working Back Button

**Problem**: The application lacks proper navigation history and back button functionality.

**Solution**: Implement browser history management and add a back button to the UI.

### Implementation Steps

1. **Add React Router for proper navigation**:

```bash
# Install React Router
npm install react-router-dom
```

2. **Update App.jsx to use React Router**:

```jsx
// file: webui/src/App.jsx

// Add imports
import {
  BrowserRouter as Router,
  Routes,
  Route,
  useNavigate,
  useLocation,
} from "react-router-dom";
import { ArrowBack as BackIcon } from "@mui/icons-material";

// Create a new component for the main app content
function AppContent() {
  const navigate = useNavigate();
  const location = useLocation();
  const [drawerOpen, setDrawerOpen] = useState(false);

  // ... existing state and functions ...

  // Add back button handler
  const handleBack = () => {
    navigate(-1);
  };

  // Show back button when not on dashboard
  const showBackButton =
    location.pathname !== "/" && location.pathname !== "/dashboard";

  return (
    <Box sx={{ display: "flex" }}>
      <AppBar
        position="fixed"
        sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}
      >
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            edge="start"
            onClick={() => setDrawerOpen(!drawerOpen)}
            sx={{ mr: 2 }}
          >
            <MenuIcon />
          </IconButton>

          {/* Add back button */}
          {showBackButton && (
            <IconButton color="inherit" onClick={handleBack} sx={{ mr: 2 }}>
              <BackIcon />
            </IconButton>
          )}

          <Typography variant="h6" noWrap component="div" sx={{ flexGrow: 1 }}>
            Subtitle Manager
          </Typography>

          {/* ... existing header content ... */}
        </Toolbar>
      </AppBar>

      {/* ... drawer content ... */}

      <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
        <Toolbar />
        <Routes>
          <Route
            path="/"
            element={<Dashboard backendAvailable={backendAvailable} />}
          />
          <Route
            path="/dashboard"
            element={<Dashboard backendAvailable={backendAvailable} />}
          />
          <Route
            path="/library"
            element={<MediaLibrary backendAvailable={backendAvailable} />}
          />
          <Route
            path="/wanted"
            element={<Wanted backendAvailable={backendAvailable} />}
          />
          <Route
            path="/history"
            element={<History backendAvailable={backendAvailable} />}
          />
          <Route
            path="/settings"
            element={<Settings backendAvailable={backendAvailable} />}
          />
          <Route
            path="/system"
            element={<System backendAvailable={backendAvailable} />}
          />
          <Route
            path="/tools/extract"
            element={<Extract backendAvailable={backendAvailable} />}
          />
          <Route
            path="/tools/convert"
            element={<Convert backendAvailable={backendAvailable} />}
          />
          <Route
            path="/tools/translate"
            element={<Translate backendAvailable={backendAvailable} />}
          />
        </Routes>
      </Box>
    </Box>
  );
}

// Update main App component
function App() {
  // ... existing authentication logic ...

  if (!authed) {
    return (
      <ThemeProvider theme={theme}>
        <CssBaseline />
        {/* ... login UI ... */}
      </ThemeProvider>
    );
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <AppContent />
      </Router>
    </ThemeProvider>
  );
}
```

3. **Update navigation items to use proper routing**:

```jsx
// Update the drawer navigation
{
  navigationItems.map((item) => (
    <ListItem key={item.id} disablePadding>
      <ListItemButton
        component={Link}
        to={item.path}
        selected={location.pathname === item.path}
        onClick={() => setDrawerOpen(false)}
      >
        <ListItemIcon sx={{ color: "inherit" }}>{item.icon}</ListItemIcon>
        <ListItemText primary={item.label} />
      </ListItemButton>
    </ListItem>
  ));
}
```

---

## Issue 4: Add Sidebar Pinning

**Problem**: Users cannot pin the sidebar for persistent navigation.

**Solution**: Add sidebar pinning functionality with persistent storage.

### Implementation Steps

1. **Add pinning state management**:

```jsx
// file: webui/src/App.jsx

function AppContent() {
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [drawerPinned, setDrawerPinned] = useState(() => {
    const saved = localStorage.getItem("sidebarPinned");
    return saved ? JSON.parse(saved) : false;
  });

  // Update drawer open logic
  const handleDrawerToggle = () => {
    if (drawerPinned) {
      // If pinned, just toggle the pinned state
      setDrawerPinned(false);
      localStorage.setItem("sidebarPinned", "false");
    } else {
      setDrawerOpen(!drawerOpen);
    }
  };

  const handleDrawerPin = () => {
    const newPinned = !drawerPinned;
    setDrawerPinned(newPinned);
    localStorage.setItem("sidebarPinned", JSON.stringify(newPinned));
    if (newPinned) {
      setDrawerOpen(true);
    }
  };

  // Determine if drawer should be open
  const isDrawerOpen = drawerPinned || drawerOpen;

  return (
    <Box sx={{ display: "flex" }}>
      {/* ... AppBar ... */}

      <Drawer
        variant={drawerPinned ? "permanent" : "temporary"}
        anchor="left"
        open={isDrawerOpen}
        onClose={() => !drawerPinned && setDrawerOpen(false)}
        sx={{
          width: 280,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: 280,
            boxSizing: "border-box",
          },
        }}
      >
        <Toolbar />
        <Box sx={{ overflow: "auto", py: 1 }}>
          {/* Add pin button at top of drawer */}
          <Box sx={{ px: 2, pb: 1 }}>
            <Button
              fullWidth
              variant="outlined"
              size="small"
              onClick={handleDrawerPin}
              startIcon={<PinIcon />}
            >
              {drawerPinned ? "Unpin Sidebar" : "Pin Sidebar"}
            </Button>
          </Box>

          <List>{/* ... navigation items ... */}</List>
        </Box>
      </Drawer>

      <Box
        component="main"
        sx={{
          flexGrow: 1,
          p: 3,
          marginLeft: drawerPinned ? "280px" : 0,
          transition: theme.transitions.create("margin", {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
          }),
        }}
      >
        {/* ... main content ... */}
      </Box>
    </Box>
  );
}
```

2. **Add pin icon import**:

```jsx
// file: webui/src/App.jsx
import { PushPin as PinIcon } from "@mui/icons-material";
```

---

## Issue 5: Reorganize Navigation Order

**Problem**: Current navigation order is not intuitive. Should be: Dashboard → Media Library → Wanted → History → Settings → System, with Extract/Translate/Convert moved to Tools.

**Solution**: Reorder navigation items and create Tools section.

### Implementation Steps

1. **Update navigation items array**:

```jsx
// file: webui/src/App.jsx

const navigationItems = [
  {
    id: "dashboard",
    label: "Dashboard",
    icon: <DashboardIcon />,
    path: "/dashboard",
  },
  {
    id: "library",
    label: "Media Library",
    icon: <LibraryIcon />,
    path: "/library",
  },
  {
    id: "wanted",
    label: "Wanted",
    icon: <WantedIcon />,
    path: "/wanted",
  },
  {
    id: "history",
    label: "History",
    icon: <HistoryIcon />,
    path: "/history",
  },
  {
    id: "settings",
    label: "Settings",
    icon: <SettingsIcon />,
    path: "/settings",
  },
  {
    id: "system",
    label: "System",
    icon: <SystemIcon />,
    path: "/system",
  },
];

const toolsItems = [
  {
    id: "extract",
    label: "Extract",
    icon: <ExtractIcon />,
    path: "/tools/extract",
  },
  {
    id: "convert",
    label: "Convert",
    icon: <ConvertIcon />,
    path: "/tools/convert",
  },
  {
    id: "translate",
    label: "Translate",
    icon: <TranslateIcon />,
    path: "/tools/translate",
  },
];
```

2. **Update drawer content to show sections**:

```jsx
// file: webui/src/App.jsx

<List>
  {navigationItems.map((item) => (
    <ListItem key={item.id} disablePadding>
      <ListItemButton
        component={Link}
        to={item.path}
        selected={location.pathname === item.path}
        onClick={() => !drawerPinned && setDrawerOpen(false)}
      >
        <ListItemIcon sx={{ color: "inherit" }}>{item.icon}</ListItemIcon>
        <ListItemText primary={item.label} />
      </ListItemButton>
    </ListItem>
  ))}

  {/* Tools Section */}
  <Divider sx={{ my: 1 }} />
  <ListSubheader component="div" sx={{ px: 2, py: 1 }}>
    Tools
  </ListSubheader>

  {toolsItems.map((item) => (
    <ListItem key={item.id} disablePadding>
      <ListItemButton
        component={Link}
        to={item.path}
        selected={location.pathname === item.path}
        onClick={() => !drawerPinned && setDrawerOpen(false)}
        sx={{ pl: 4 }} // Indent tools items
      >
        <ListItemIcon sx={{ color: "inherit", minWidth: 36 }}>
          {item.icon}
        </ListItemIcon>
        <ListItemText primary={item.label} />
      </ListItemButton>
    </ListItem>
  ))}
</List>
```

3. **Add required imports**:

```jsx
// file: webui/src/App.jsx
import { Divider, ListSubheader, Link } from "@mui/material";
import { Link } from "react-router-dom";
```

## Testing Procedures

### For Issue 3 (Back Button)

1. Navigate between pages
2. Verify back button appears when not on dashboard
3. Test back button functionality
4. Ensure browser history works correctly

### For Issue 4 (Sidebar Pinning)

1. Test pin/unpin functionality
2. Verify state persists across browser refreshes
3. Test drawer behavior in both pinned and unpinned modes
4. Check responsive behavior

### For Issue 5 (Navigation Order)

1. Verify navigation items appear in correct order
2. Test Tools section functionality
3. Ensure all routes work correctly
4. Verify visual hierarchy with Tools indentation

## Time Estimates

- Issue 3 (Back Button): 4-6 hours
- Issue 4 (Sidebar Pinning): 3-4 hours
- Issue 5 (Navigation Reorder): 2-3 hours

**Total Section 2**: 9-13 hours

---

# Section 3: Settings Page Enhancement

This section covers issues 6-16: enhancing the settings page to match Bazarr functionality and improve the overall configuration experience.

## Issue 6: Enhance General Settings (Bazarr Compatibility)

**Problem**: Current general settings are basic and lack important Bazarr-compatible options.

**Solution**: Expand GeneralSettings component to include Host, Proxy, Updates, Logging, Backups, and Analytics sections.

### Implementation Steps

1. **Update GeneralSettings.jsx with expanded functionality**:

```jsx
// file: webui/src/components/GeneralSettings.jsx
import {
  Box,
  Button,
  Card,
  CardContent,
  FormControl,
  FormControlLabel,
  FormGroup,
  Grid,
  InputLabel,
  MenuItem,
  Select,
  Switch,
  TextField,
  Typography,
  Divider,
  Alert,
  Accordion,
  AccordionSummary,
  AccordionDetails,
} from "@mui/material";
import { ExpandMore as ExpandIcon } from "@mui/icons-material";
import { useEffect, useState } from "react";

/**
 * Enhanced GeneralSettings component with Bazarr-compatible configuration options
 */
export default function GeneralSettings({
  config,
  onSave,
  backendAvailable = true,
}) {
  // Host Settings
  const [address, setAddress] = useState("");
  const [port, setPort] = useState(8080);
  const [baseURL, setBaseURL] = useState("");

  // Proxy Settings
  const [proxyEnabled, setProxyEnabled] = useState(false);
  const [proxyType, setProxyType] = useState("http");
  const [proxyHost, setProxyHost] = useState("");
  const [proxyPort, setProxyPort] = useState("");
  const [proxyUsername, setProxyUsername] = useState("");
  const [proxyPassword, setProxyPassword] = useState("");

  // Updates Settings
  const [autoUpdate, setAutoUpdate] = useState(false);
  const [updateBranch, setUpdateBranch] = useState("master");
  const [updateFrequency, setUpdateFrequency] = useState("daily");

  // Logging Settings
  const [logLevel, setLogLevel] = useState("info");
  const [logFilter, setLogFilter] = useState("");
  const [logFilterRegex, setLogFilterRegex] = useState(false);
  const [logFilterIgnoreCase, setLogFilterIgnoreCase] = useState(true);

  // Backup Settings
  const [backupEnabled, setBackupEnabled] = useState(true);
  const [backupFrequency, setBackupFrequency] = useState("weekly");
  const [backupRetention, setBackupRetention] = useState(5);
  const [backupLocation, setBackupLocation] = useState("");

  // Analytics Settings
  const [analyticsEnabled, setAnalyticsEnabled] = useState(false);

  // Scheduler Settings
  const [schedulerEnabled, setSchedulerEnabled] = useState(false);
  const [libraryScanFreq, setLibraryScanFreq] = useState("daily");
  const [wantedSearchFreq, setWantedSearchFreq] = useState("daily");
  const [libraryScanCron, setLibraryScanCron] = useState("");
  const [wantedSearchCron, setWantedSearchCron] = useState("");
  const [maxConcurrentDownloads, setMaxConcurrentDownloads] = useState(3);
  const [downloadTimeout, setDownloadTimeout] = useState(60);

  useEffect(() => {
    if (config) {
      // Host Settings
      setAddress(config.address || "");
      setPort(config.port || 8080);
      setBaseURL(config.base_url || "");

      // Proxy Settings
      setProxyEnabled(config.proxy_enabled || false);
      setProxyType(config.proxy_type || "http");
      setProxyHost(config.proxy_host || "");
      setProxyPort(config.proxy_port || "");
      setProxyUsername(config.proxy_username || "");
      setProxyPassword(config.proxy_password || "");

      // Updates Settings
      setAutoUpdate(config.auto_update || false);
      setUpdateBranch(config.update_branch || "master");
      setUpdateFrequency(config.update_frequency || "daily");

      // Logging Settings
      setLogLevel(config.log_level || "info");
      setLogFilter(config.log_filter || "");
      setLogFilterRegex(config.log_filter_regex || false);
      setLogFilterIgnoreCase(config.log_filter_ignore_case !== false);

      // Backup Settings
      setBackupEnabled(config.backup_enabled !== false);
      setBackupFrequency(config.backup_frequency || "weekly");
      setBackupRetention(config.backup_retention || 5);
      setBackupLocation(config.backup_location || "");

      // Analytics Settings
      setAnalyticsEnabled(config.analytics_enabled || false);

      // Scheduler Settings
      setSchedulerEnabled(config.scheduler_enabled || false);
      setLibraryScanFreq(config.library_scan_frequency || "daily");
      setWantedSearchFreq(config.wanted_search_frequency || "daily");
      setLibraryScanCron(config.library_scan_cron || "");
      setWantedSearchCron(config.wanted_search_cron || "");
      setMaxConcurrentDownloads(config.max_concurrent_downloads || 3);
      setDownloadTimeout(config.download_timeout || 60);
    }
  }, [config]);

  const handleSave = () => {
    const newConfig = {
      // Host Settings
      address,
      port: parseInt(port),
      base_url: baseURL,

      // Proxy Settings
      proxy_enabled: proxyEnabled,
      proxy_type: proxyType,
      proxy_host: proxyHost,
      proxy_port: proxyPort,
      proxy_username: proxyUsername,
      proxy_password: proxyPassword,

      // Updates Settings
      auto_update: autoUpdate,
      update_branch: updateBranch,
      update_frequency: updateFrequency,

      // Logging Settings
      log_level: logLevel,
      log_filter: logFilter,
      log_filter_regex: logFilterRegex,
      log_filter_ignore_case: logFilterIgnoreCase,

      // Backup Settings
      backup_enabled: backupEnabled,
      backup_frequency: backupFrequency,
      backup_retention: parseInt(backupRetention),
      backup_location: backupLocation,

      // Analytics Settings
      analytics_enabled: analyticsEnabled,

      // Scheduler Settings
      scheduler_enabled: schedulerEnabled,
      library_scan_frequency: libraryScanFreq,
      wanted_search_frequency: wantedSearchFreq,
      library_scan_cron: libraryScanCron,
      wanted_search_cron: wantedSearchCron,
      max_concurrent_downloads: parseInt(maxConcurrentDownloads),
      download_timeout: parseInt(downloadTimeout),
    };

    onSave(newConfig);
  };

  return (
    <Box sx={{ maxWidth: 800 }}>
      <Typography variant="h6" gutterBottom>
        General Settings
      </Typography>

      {!backendAvailable && (
        <Alert severity="warning" sx={{ mb: 3 }}>
          Backend service is not available. Settings cannot be modified.
        </Alert>
      )}

      {/* Host Configuration */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Host
          </Typography>
          <Grid container spacing={2}>
            <Grid item xs={12} md={6}>
              <TextField
                fullWidth
                label="Address"
                value={address}
                onChange={(e) => setAddress(e.target.value)}
                placeholder="0.0.0.0"
                helperText="IP address to bind to (0.0.0.0 for all interfaces)"
                disabled={!backendAvailable}
              />
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField
                fullWidth
                label="Port"
                type="number"
                value={port}
                onChange={(e) => setPort(e.target.value)}
                placeholder="8080"
                helperText="Port number for the web interface"
                disabled={!backendAvailable}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="Base URL"
                value={baseURL}
                onChange={(e) => setBaseURL(e.target.value)}
                placeholder="/subtitles"
                helperText="Base URL for reverse proxy setups (e.g., /subtitles)"
                disabled={!backendAvailable}
              />
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {/* Proxy Configuration */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Proxy
          </Typography>
          <FormGroup sx={{ mb: 2 }}>
            <FormControlLabel
              control={
                <Switch
                  checked={proxyEnabled}
                  onChange={(e) => setProxyEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              }
              label="Enable Proxy"
            />
          </FormGroup>

          {proxyEnabled && (
            <Grid container spacing={2}>
              <Grid item xs={12} md={3}>
                <FormControl fullWidth>
                  <InputLabel>Type</InputLabel>
                  <Select
                    value={proxyType}
                    label="Type"
                    onChange={(e) => setProxyType(e.target.value)}
                    disabled={!backendAvailable}
                  >
                    <MenuItem value="http">HTTP</MenuItem>
                    <MenuItem value="https">HTTPS</MenuItem>
                    <MenuItem value="socks5">SOCKS5</MenuItem>
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Host"
                  value={proxyHost}
                  onChange={(e) => setProxyHost(e.target.value)}
                  placeholder="proxy.example.com"
                  disabled={!backendAvailable}
                />
              </Grid>
              <Grid item xs={12} md={3}>
                <TextField
                  fullWidth
                  label="Port"
                  type="number"
                  value={proxyPort}
                  onChange={(e) => setProxyPort(e.target.value)}
                  placeholder="8080"
                  disabled={!backendAvailable}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Username"
                  value={proxyUsername}
                  onChange={(e) => setProxyUsername(e.target.value)}
                  placeholder="Optional"
                  disabled={!backendAvailable}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Password"
                  type="password"
                  value={proxyPassword}
                  onChange={(e) => setProxyPassword(e.target.value)}
                  placeholder="Optional"
                  disabled={!backendAvailable}
                />
              </Grid>
            </Grid>
          )}
        </CardContent>
      </Card>

      {/* Updates Configuration */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Updates
          </Typography>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <FormGroup>
                <FormControlLabel
                  control={
                    <Switch
                      checked={autoUpdate}
                      onChange={(e) => setAutoUpdate(e.target.checked)}
                      disabled={!backendAvailable}
                    />
                  }
                  label="Automatic Updates"
                />
              </FormGroup>
            </Grid>
            {autoUpdate && (
              <>
                <Grid item xs={12} md={6}>
                  <FormControl fullWidth>
                    <InputLabel>Branch</InputLabel>
                    <Select
                      value={updateBranch}
                      label="Branch"
                      onChange={(e) => setUpdateBranch(e.target.value)}
                      disabled={!backendAvailable}
                    >
                      <MenuItem value="master">Master (Stable)</MenuItem>
                      <MenuItem value="develop">Develop (Beta)</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12} md={6}>
                  <FormControl fullWidth>
                    <InputLabel>Frequency</InputLabel>
                    <Select
                      value={updateFrequency}
                      label="Frequency"
                      onChange={(e) => setUpdateFrequency(e.target.value)}
                      disabled={!backendAvailable}
                    >
                      <MenuItem value="daily">Daily</MenuItem>
                      <MenuItem value="weekly">Weekly</MenuItem>
                      <MenuItem value="monthly">Monthly</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
              </>
            )}
          </Grid>
        </CardContent>
      </Card>

      {/* Logging Configuration */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Logging
          </Typography>
          <Grid container spacing={2}>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel>Log Level</InputLabel>
                <Select
                  value={logLevel}
                  label="Log Level"
                  onChange={(e) => setLogLevel(e.target.value)}
                  disabled={!backendAvailable}
                >
                  <MenuItem value="error">Error</MenuItem>
                  <MenuItem value="warn">Warning</MenuItem>
                  <MenuItem value="info">Info</MenuItem>
                  <MenuItem value="debug">Debug</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField
                fullWidth
                label="Log Filter"
                value={logFilter}
                onChange={(e) => setLogFilter(e.target.value)}
                placeholder="Optional regex for filtering logs"
                disabled={!backendAvailable}
              />
            </Grid>
            <Grid item xs={12}>
              <FormGroup row>
                <FormControlLabel
                  control={
                    <Switch
                      checked={logFilterRegex}
                      onChange={(e) => setLogFilterRegex(e.target.checked)}
                      disabled={!backendAvailable}
                    />
                  }
                  label="Use Regex for Filter"
                />
                <FormControlLabel
                  control={
                    <Switch
                      checked={logFilterIgnoreCase}
                      onChange={(e) => setLogFilterIgnoreCase(e.target.checked)}
                      disabled={!backendAvailable}
                    />
                  }
                  label="Ignore Case"
                />
              </FormGroup>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {/* Backup Configuration */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Backups
          </Typography>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <FormGroup>
                <FormControlLabel
                  control={
                    <Switch
                      checked={backupEnabled}
                      onChange={(e) => setBackupEnabled(e.target.checked)}
                      disabled={!backendAvailable}
                    />
                  }
                  label="Enable Backups"
                />
              </FormGroup>
            </Grid>
            {backupEnabled && (
              <>
                <Grid item xs={12} md={6}>
                  <FormControl fullWidth>
                    <InputLabel>Frequency</InputLabel>
                    <Select
                      value={backupFrequency}
                      label="Frequency"
                      onChange={(e) => setBackupFrequency(e.target.value)}
                      disabled={!backendAvailable}
                    >
                      <MenuItem value="daily">Daily</MenuItem>
                      <MenuItem value="weekly">Weekly</MenuItem>
                      <MenuItem value="monthly">Monthly</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    label="Retention (number of backups)"
                    type="number"
                    value={backupRetention}
                    onChange={(e) => setBackupRetention(e.target.value)}
                    disabled={!backendAvailable}
                  />
                </Grid>
                <Grid item xs={12}>
                  <TextField
                    fullWidth
                    label="Backup Location"
                    value={backupLocation}
                    onChange={(e) => setBackupLocation(e.target.value)}
                    placeholder="/path/to/backup"
                    disabled={!backendAvailable}
                  />
                </Grid>
              </>
            )}
          </Grid>
        </CardContent>
      </Card>

      {/* Scheduler Configuration */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Scheduler
          </Typography>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <FormGroup>
                <FormControlLabel
                  control={
                    <Switch
                      checked={schedulerEnabled}
                      onChange={(e) => setSchedulerEnabled(e.target.checked)}
                      disabled={!backendAvailable}
                    />
                  }
                  label="Enable Automatic Scheduling"
                />
              </FormGroup>
            </Grid>
            {schedulerEnabled && (
              <>
                <Grid item xs={12} md={6}>
                  <FormControl fullWidth>
                    <InputLabel>Library Scan Frequency</InputLabel>
                    <Select
                      value={libraryScanFreq}
                      label="Library Scan Frequency"
                      onChange={(e) => setLibraryScanFreq(e.target.value)}
                      disabled={!backendAvailable}
                    >
                      <MenuItem value="never">Never</MenuItem>
                      <MenuItem value="hourly">Every Hour</MenuItem>
                      <MenuItem value="daily">Daily</MenuItem>
                      <MenuItem value="weekly">Weekly</MenuItem>
                      <MenuItem value="monthly">Monthly</MenuItem>
                      <MenuItem value="custom">Custom Cron</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12} md={6}>
                  <FormControl fullWidth>
                    <InputLabel>Wanted Search Frequency</InputLabel>
                    <Select
                      value={wantedSearchFreq}
                      label="Wanted Search Frequency"
                      onChange={(e) => setWantedSearchFreq(e.target.value)}
                      disabled={!backendAvailable}
                    >
                      <MenuItem value="never">Never</MenuItem>
                      <MenuItem value="15min">Every 15 Minutes</MenuItem>
                      <MenuItem value="30min">Every 30 Minutes</MenuItem>
                      <MenuItem value="hourly">Every Hour</MenuItem>
                      <MenuItem value="daily">Daily</MenuItem>
                      <MenuItem value="custom">Custom Cron</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
                {(libraryScanFreq === "custom" ||
                  wantedSearchFreq === "custom") && (
                  <Grid item xs={12}>
                    <Alert severity="info" sx={{ mb: 2 }}>
                      Cron expressions follow standard format: minute hour day
                      month weekday
                      <br />
                      Examples: "0 2 * * *" (daily at 2 AM), "*/15 * * * *"
                      (every 15 minutes)
                    </Alert>

                    {libraryScanFreq === "custom" && (
                      <TextField
                        fullWidth
                        label="Library Scan Cron Expression"
                        value={libraryScanCron}
                        onChange={(e) => setLibraryScanCron(e.target.value)}
                        placeholder="0 2 * * *"
                        sx={{ mb: 2 }}
                        disabled={!backendAvailable}
                      />
                    )}

                    {wantedSearchFreq === "custom" && (
                      <TextField
                        fullWidth
                        label="Wanted Search Cron Expression"
                        value={wantedSearchCron}
                        onChange={(e) => setWantedSearchCron(e.target.value)}
                        placeholder="*/30 * * * *"
                        disabled={!backendAvailable}
                      />
                    )}
                  </Grid>
                )}

                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    label="Max Concurrent Downloads"
                    type="number"
                    value={maxConcurrentDownloads}
                    onChange={(e) => setMaxConcurrentDownloads(e.target.value)}
                    inputProps={{ min: 1, max: 10 }}
                    helperText="Maximum number of simultaneous subtitle downloads"
                    disabled={!backendAvailable}
                  />
                </Grid>

                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    label="Download Timeout (seconds)"
                    type="number"
                    value={downloadTimeout}
                    onChange={(e) => setDownloadTimeout(e.target.value)}
                    inputProps={{ min: 10, max: 300 }}
                    helperText="Timeout for individual subtitle downloads"
                    disabled={!backendAvailable}
                  />
                </Grid>
              </>
            )}
          </Grid>
        </CardContent>
      </Card>

      {/* Continue in next part due to length... */}

      <Button
        variant="contained"
        onClick={handleSave}
        disabled={!backendAvailable}
        sx={{ mt: 2 }}
      >
        Save Settings
      </Button>
    </Box>
  );
}
```

**Note**: This is part 1 of the GeneralSettings component. Part 2 will include Logging, Backups, and Analytics sections.

## Issue 7: Improve Database Settings

**Problem**: Database page lacks comprehensive information and management options.

**Solution**: Create an enhanced DatabaseSettings component with detailed information and management capabilities.

### Implementation Steps

1. **Create enhanced DatabaseSettings.jsx**:

```jsx
// file: webui/src/components/DatabaseSettings.jsx
import {
  Box,
  Button,
  Card,
  CardContent,
  Typography,
  Grid,
  Chip,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Alert,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  CircularProgress,
  LinearProgress,
} from "@mui/material";
import {
  Storage as DatabaseIcon,
  Backup as BackupIcon,
  CleaningServices as OptimizeIcon,
  Assessment as StatsIcon,
} from "@mui/icons-material";
import { useEffect, useState } from "react";

/**
 * Enhanced DatabaseSettings component with comprehensive database management
 */
export default function DatabaseSettings({
  config,
  onSave,
  backendAvailable = true,
}) {
  const [dbInfo, setDbInfo] = useState(null);
  const [loading, setLoading] = useState(false);
  const [backupDialog, setBackupDialog] = useState(false);
  const [optimizeDialog, setOptimizeDialog] = useState(false);
  const [stats, setStats] = useState(null);

  const loadDatabaseInfo = async () => {
    if (!backendAvailable) return;

    setLoading(true);
    try {
      const response = await fetch("/api/database/info");
      if (response.ok) {
        const data = await response.json();
        setDbInfo(data);
      }

      const statsResponse = await fetch("/api/database/stats");
      if (statsResponse.ok) {
        const statsData = await statsResponse.json();
        setStats(statsData);
      }
    } catch (error) {
      console.error("Failed to load database info:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadDatabaseInfo();
  }, [backendAvailable]);

  const handleBackup = async () => {
    try {
      const response = await fetch("/api/database/backup", { method: "POST" });
      if (response.ok) {
        alert("Database backup completed successfully");
        setBackupDialog(false);
        loadDatabaseInfo();
      }
    } catch (error) {
      alert("Backup failed: " + error.message);
    }
  };

  const handleOptimize = async () => {
    try {
      const response = await fetch("/api/database/optimize", {
        method: "POST",
      });
      if (response.ok) {
        alert("Database optimization completed successfully");
        setOptimizeDialog(false);
        loadDatabaseInfo();
      }
    } catch (error) {
      alert("Optimization failed: " + error.message);
    }
  };

  const formatBytes = (bytes) => {
    if (!bytes) return "N/A";
    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    return Math.round((bytes / Math.pow(1024, i)) * 100) / 100 + " " + sizes[i];
  };

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="200px"
      >
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box sx={{ maxWidth: 1000 }}>
      <Typography variant="h6" gutterBottom>
        Database Settings
      </Typography>

      {!backendAvailable && (
        <Alert severity="warning" sx={{ mb: 3 }}>
          Backend service is not available. Database information cannot be
          loaded.
        </Alert>
      )}

      <Grid container spacing={3}>
        {/* Database Information */}
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography
                variant="h6"
                gutterBottom
                color="primary"
                sx={{ display: "flex", alignItems: "center" }}
              >
                <DatabaseIcon sx={{ mr: 1 }} />
                Database Information
              </Typography>

              {dbInfo ? (
                <TableContainer>
                  <Table size="small">
                    <TableBody>
                      <TableRow>
                        <TableCell>
                          <strong>Type</strong>
                        </TableCell>
                        <TableCell>
                          <Chip
                            label={dbInfo.type || "Unknown"}
                            color={
                              dbInfo.type === "postgresql"
                                ? "primary"
                                : "default"
                            }
                            size="small"
                          />
                        </TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Version</strong>
                        </TableCell>
                        <TableCell>{dbInfo.version || "N/A"}</TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Size</strong>
                        </TableCell>
                        <TableCell>{formatBytes(dbInfo.size)}</TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Location</strong>
                        </TableCell>
                        <TableCell>{dbInfo.path || "N/A"}</TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Connection Status</strong>
                        </TableCell>
                        <TableCell>
                          <Chip
                            label={
                              dbInfo.connected ? "Connected" : "Disconnected"
                            }
                            color={dbInfo.connected ? "success" : "error"}
                            size="small"
                          />
                        </TableCell>
                      </TableRow>
                    </TableBody>
                  </Table>
                </TableContainer>
              ) : (
                <Alert severity="info">
                  Database information not available
                </Alert>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* Database Statistics */}
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography
                variant="h6"
                gutterBottom
                color="primary"
                sx={{ display: "flex", alignItems: "center" }}
              >
                <StatsIcon sx={{ mr: 1 }} />
                Statistics
              </Typography>

              {stats ? (
                <TableContainer>
                  <Table size="small">
                    <TableBody>
                      <TableRow>
                        <TableCell>
                          <strong>Total Records</strong>
                        </TableCell>
                        <TableCell>
                          {stats.totalRecords?.toLocaleString() || "N/A"}
                        </TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Users</strong>
                        </TableCell>
                        <TableCell>{stats.users || 0}</TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Downloads</strong>
                        </TableCell>
                        <TableCell>
                          {stats.downloads?.toLocaleString() || 0}
                        </TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Media Items</strong>
                        </TableCell>
                        <TableCell>
                          {stats.mediaItems?.toLocaleString() || 0}
                        </TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Last Backup</strong>
                        </TableCell>
                        <TableCell>
                          {stats.lastBackup
                            ? new Date(stats.lastBackup).toLocaleDateString()
                            : "Never"}
                        </TableCell>
                      </TableRow>
                    </TableBody>
                  </Table>
                </TableContainer>
              ) : (
                <Alert severity="info">Statistics not available</Alert>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* Database Management Actions */}
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom color="primary">
                Database Management
              </Typography>

              <Grid container spacing={2}>
                <Grid item>
                  <Button
                    variant="outlined"
                    startIcon={<BackupIcon />}
                    onClick={() => setBackupDialog(true)}
                    disabled={!backendAvailable}
                  >
                    Create Backup
                  </Button>
                </Grid>
                <Grid item>
                  <Button
                    variant="outlined"
                    startIcon={<OptimizeIcon />}
                    onClick={() => setOptimizeDialog(true)}
                    disabled={!backendAvailable}
                  >
                    Optimize Database
                  </Button>
                </Grid>
              </Grid>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      {/* Backup Confirmation Dialog */}
      <Dialog open={backupDialog} onClose={() => setBackupDialog(false)}>
        <DialogTitle>Create Database Backup</DialogTitle>
        <DialogContent>
          <DialogContentText>
            This will create a backup of your database. The process may take a
            few moments depending on the database size.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setBackupDialog(false)}>Cancel</Button>
          <Button onClick={handleBackup} variant="contained">
            Create Backup
          </Button>
        </DialogActions>
      </Dialog>

      {/* Optimize Confirmation Dialog */}
      <Dialog open={optimizeDialog} onClose={() => setOptimizeDialog(false)}>
        <DialogTitle>Optimize Database</DialogTitle>
        <DialogContent>
          <DialogContentText>
            This will optimize your database by rebuilding indexes and cleaning
            up unused space. This process may take several minutes.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOptimizeDialog(false)}>Cancel</Button>
          <Button onClick={handleOptimize} variant="contained">
            Optimize
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
```

## Time Estimates

- Issue 6 (Enhanced General Settings): 8-12 hours
- Issue 7 (Database Settings): 6-8 hours

**Total for Issues 6-7**: 14-20 hours

## Next Section Preview

The next section will cover Issues 8-16, including:

- Authentication page redesign with cards
- OAuth2 management features
- Notifications system enhancement
- Languages page creation
- Scheduler settings integration

---

# Section 4: Authentication and Notifications

This section covers issues 8-16: redesigning authentication with cards, enhancing notifications, creating OAuth2 management, adding Languages page, and provider system improvements.

## Issue 8: Redesign Authentication Page with Cards

**Problem**: Current authentication page is basic and lacks visual organization for different auth methods.

**Solution**: Create a card-based UI for each authentication method with enable/disable toggles.

### Implementation Steps

1. **Create enhanced AuthSettings.jsx with card-based design**:

```jsx
// file: webui/src/components/AuthSettings.jsx
import {
  Box,
  Button,
  Card,
  CardContent,
  CardActions,
  FormControlLabel,
  Grid,
  Switch,
  TextField,
  Typography,
  Divider,
  Alert,
  IconButton,
  Tooltip,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
} from "@mui/material";
import {
  GitHub as GitHubIcon,
  Key as ApiKeyIcon,
  Password as PasswordIcon,
  Security as OAuthIcon,
  Refresh as RefreshIcon,
  Add as AddIcon,
  Delete as DeleteIcon,
  Visibility as ShowIcon,
  VisibilityOff as HideIcon,
} from "@mui/icons-material";
import { useEffect, useState } from "react";

/**
 * Enhanced AuthSettings with card-based UI for each authentication method
 */
export default function AuthSettings({
  config,
  onSave,
  backendAvailable = true,
}) {
  // Password Authentication
  const [passwordAuthEnabled, setPasswordAuthEnabled] = useState(true);
  const [requireStrongPasswords, setRequireStrongPasswords] = useState(true);
  const [passwordExpiry, setPasswordExpiry] = useState(90);

  // API Key Authentication
  const [apiKeyEnabled, setApiKeyEnabled] = useState(true);
  const [apiKeyExpiry, setApiKeyExpiry] = useState(0); // 0 = no expiry

  // GitHub OAuth
  const [githubEnabled, setGithubEnabled] = useState(false);
  const [githubClientId, setGithubClientId] = useState("");
  const [githubClientSecret, setGithubClientSecret] = useState("");
  const [githubRedirectUrl, setGithubRedirectUrl] = useState("");
  const [showGithubSecret, setShowGithubSecret] = useState(false);

  // Generic OAuth2
  const [genericOAuthEnabled, setGenericOAuthEnabled] = useState(false);
  const [oauthProvider, setOauthProvider] = useState("");
  const [oauthClientId, setOauthClientId] = useState("");
  const [oauthClientSecret, setOauthClientSecret] = useState("");
  const [oauthAuthUrl, setOauthAuthUrl] = useState("");
  const [oauthTokenUrl, setOauthTokenUrl] = useState("");
  const [oauthUserUrl, setOauthUserUrl] = useState("");
  const [showOAuthSecret, setShowOAuthSecret] = useState(false);

  // Dialog states
  const [resetGithubDialog, setResetGithubDialog] = useState(false);
  const [resetOAuthDialog, setResetOAuthDialog] = useState(false);

  useEffect(() => {
    if (config) {
      // Password settings
      setPasswordAuthEnabled(config.password_auth_enabled !== false);
      setRequireStrongPasswords(config.require_strong_passwords !== false);
      setPasswordExpiry(config.password_expiry || 90);

      // API Key settings
      setApiKeyEnabled(config.api_key_enabled !== false);
      setApiKeyExpiry(config.api_key_expiry || 0);

      // GitHub OAuth settings
      setGithubEnabled(config.github_oauth_enabled || false);
      setGithubClientId(config.github_client_id || "");
      setGithubClientSecret(config.github_client_secret || "");
      setGithubRedirectUrl(
        config.github_redirect_url ||
          `${window.location.origin}/api/oauth/github/callback`,
      );

      // Generic OAuth settings
      setGenericOAuthEnabled(config.generic_oauth_enabled || false);
      setOauthProvider(config.oauth_provider || "");
      setOauthClientId(config.oauth_client_id || "");
      setOauthClientSecret(config.oauth_client_secret || "");
      setOauthAuthUrl(config.oauth_auth_url || "");
      setOauthTokenUrl(config.oauth_token_url || "");
      setOauthUserUrl(config.oauth_user_url || "");
    }
  }, [config]);

  const handleSave = () => {
    const newConfig = {
      // Password authentication
      password_auth_enabled: passwordAuthEnabled,
      require_strong_passwords: requireStrongPasswords,
      password_expiry: parseInt(passwordExpiry),

      // API Key authentication
      api_key_enabled: apiKeyEnabled,
      api_key_expiry: parseInt(apiKeyExpiry),

      // GitHub OAuth
      github_oauth_enabled: githubEnabled,
      github_client_id: githubClientId,
      github_client_secret: githubClientSecret,
      github_redirect_url: githubRedirectUrl,

      // Generic OAuth
      generic_oauth_enabled: genericOAuthEnabled,
      oauth_provider: oauthProvider,
      oauth_client_id: oauthClientId,
      oauth_client_secret: oauthClientSecret,
      oauth_auth_url: oauthAuthUrl,
      oauth_token_url: oauthTokenUrl,
      oauth_user_url: oauthUserUrl,
    };

    onSave(newConfig);
  };

  const generateGithubCredentials = async () => {
    try {
      const response = await fetch("/api/oauth/github/generate", {
        method: "POST",
      });
      if (response.ok) {
        const data = await response.json();
        setGithubClientId(data.client_id);
        setGithubClientSecret(data.client_secret);
        alert("New GitHub OAuth credentials generated successfully");
      }
    } catch (error) {
      alert("Failed to generate credentials: " + error.message);
    }
  };

  const regenerateGithubSecret = async () => {
    try {
      const response = await fetch("/api/oauth/github/regenerate", {
        method: "POST",
      });
      if (response.ok) {
        const data = await response.json();
        setGithubClientSecret(data.client_secret);
        alert("GitHub client secret regenerated successfully");
      }
    } catch (error) {
      alert("Failed to regenerate secret: " + error.message);
    }
  };

  const resetGithubConfig = () => {
    setGithubClientId("");
    setGithubClientSecret("");
    setGithubRedirectUrl(`${window.location.origin}/api/oauth/github/callback`);
    setGithubEnabled(false);
    setResetGithubDialog(false);
    alert("GitHub OAuth configuration reset to defaults");
  };

  const resetOAuthConfig = () => {
    setOauthProvider("");
    setOauthClientId("");
    setOauthClientSecret("");
    setOauthAuthUrl("");
    setOauthTokenUrl("");
    setOauthUserUrl("");
    setGenericOAuthEnabled(false);
    setResetOAuthDialog(false);
    alert("Generic OAuth configuration reset to defaults");
  };

  return (
    <Box sx={{ maxWidth: 1200 }}>
      <Typography variant="h6" gutterBottom>
        Authentication Settings
      </Typography>

      {!backendAvailable && (
        <Alert severity="warning" sx={{ mb: 3 }}>
          Backend service is not available. Authentication settings cannot be
          modified.
        </Alert>
      )}

      <Grid container spacing={3}>
        {/* Password Authentication Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <PasswordIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  Password Authentication
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={passwordAuthEnabled}
                  onChange={(e) => setPasswordAuthEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Traditional username and password authentication for local
                users.
              </Typography>

              {passwordAuthEnabled && (
                <Box sx={{ mt: 2 }}>
                  <FormControlLabel
                    control={
                      <Switch
                        checked={requireStrongPasswords}
                        onChange={(e) =>
                          setRequireStrongPasswords(e.target.checked)
                        }
                        disabled={!backendAvailable}
                      />
                    }
                    label="Require Strong Passwords"
                    sx={{ mb: 1 }}
                  />

                  <TextField
                    fullWidth
                    label="Password Expiry (days)"
                    type="number"
                    value={passwordExpiry}
                    onChange={(e) => setPasswordExpiry(e.target.value)}
                    placeholder="Set to 0 for no expiry"
                    disabled={!backendAvailable}
                    sx={{ mb: 2 }}
                  />
                </Box>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* API Key Authentication Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <ApiKeyIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  API Key Authentication
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={apiKeyEnabled}
                  onChange={(e) => setApiKeyEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                API key-based authentication for automated access and
                integrations.
              </Typography>

              {apiKeyEnabled && (
                <TextField
                  fullWidth
                  label="API Key Expiry (days)"
                  type="number"
                  value={apiKeyExpiry}
                  onChange={(e) => setApiKeyExpiry(e.target.value)}
                  helperText="Set to 0 for no expiry"
                  disabled={!backendAvailable}
                  sx={{ mb: 2 }}
                />
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* GitHub OAuth Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <GitHubIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  GitHub OAuth
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={githubEnabled}
                  onChange={(e) => setGithubEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Allow users to authenticate using their GitHub accounts.
              </Typography>

              {githubEnabled && (
                <Box
                  sx={{
                    mt: 2,
                    display: "flex",
                    flexDirection: "column",
                    gap: 2,
                  }}
                >
                  <TextField
                    fullWidth
                    label="Client ID"
                    value={githubClientId}
                    onChange={(e) => setGithubClientId(e.target.value)}
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="Client Secret"
                    type={showGithubSecret ? "text" : "password"}
                    value={githubClientSecret}
                    onChange={(e) => setGithubClientSecret(e.target.value)}
                    disabled={!backendAvailable}
                    InputProps={{
                      endAdornment: (
                        <IconButton
                          onClick={() => setShowGithubSecret(!showGithubSecret)}
                          edge="end"
                        >
                          {showGithubSecret ? <VisibilityOff /> : <ShowIcon />}
                        </IconButton>
                      ),
                    }}
                  />

                  <TextField
                    fullWidth
                    label="Redirect URL"
                    value={githubRedirectUrl}
                    onChange={(e) => setGithubRedirectUrl(e.target.value)}
                    disabled={!backendAvailable}
                    helperText="Copy this URL to your GitHub OAuth app settings"
                  />
                </Box>
              )}
            </CardContent>

            {githubEnabled && (
              <CardActions>
                <Button
                  size="small"
                  startIcon={<AddIcon />}
                  onClick={generateGithubCredentials}
                  disabled={!backendAvailable}
                >
                  Generate New
                </Button>
                <Button
                  size="small"
                  startIcon={<RefreshIcon />}
                  onClick={regenerateGithubSecret}
                  disabled={!backendAvailable}
                >
                  Regenerate Secret
                </Button>
                <Button
                  size="small"
                  startIcon={<DeleteIcon />}
                  onClick={() => setResetGithubDialog(true)}
                  disabled={!backendAvailable}
                  color="error"
                >
                  Reset
                </Button>
              </CardActions>
            )}
          </Card>
        </Grid>

        {/* Generic OAuth Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <OAuthIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  Generic OAuth2
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={genericOAuthEnabled}
                  onChange={(e) => setGenericOAuthEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Configure authentication with any OAuth2-compatible provider.
              </Typography>

              {genericOAuthEnabled && (
                <Box
                  sx={{
                    mt: 2,
                    display: "flex",
                    flexDirection: "column",
                    gap: 2,
                  }}
                >
                  <TextField
                    fullWidth
                    label="Provider Name"
                    value={oauthProvider}
                    onChange={(e) => setOauthProvider(e.target.value)}
                    placeholder="e.g., Google, Microsoft, etc."
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="Client ID"
                    value={oauthClientId}
                    onChange={(e) => setOauthClientId(e.target.value)}
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="Client Secret"
                    type={showOAuthSecret ? "text" : "password"}
                    value={oauthClientSecret}
                    onChange={(e) => setOauthClientSecret(e.target.value)}
                    disabled={!backendAvailable}
                    InputProps={{
                      endAdornment: (
                        <IconButton
                          onClick={() => setShowOAuthSecret(!showOAuthSecret)}
                          edge="end"
                        >
                          {showOAuthSecret ? <VisibilityOff /> : <ShowIcon />}
                        </IconButton>
                      ),
                    }}
                  />

                  <TextField
                    fullWidth
                    label="Authorization URL"
                    value={oauthAuthUrl}
                    onChange={(e) => setOauthAuthUrl(e.target.value)}
                    placeholder="https://provider.com/oauth/authorize"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="Token URL"
                    value={oauthTokenUrl}
                    onChange={(e) => setOauthTokenUrl(e.target.value)}
                    placeholder="https://provider.com/oauth/token"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="User Info URL"
                    value={oauthUserUrl}
                    onChange={(e) => setOauthUserUrl(e.target.value)}
                    placeholder="https://provider.com/oauth/userinfo"
                    disabled={!backendAvailable}
                  />
                </Box>
              )}
            </CardContent>

            {genericOAuthEnabled && (
              <CardActions>
                <Button
                  size="small"
                  startIcon={<DeleteIcon />}
                  onClick={() => setResetOAuthDialog(true)}
                  disabled={!backendAvailable}
                  color="error"
                >
                  Reset Configuration
                </Button>
              </CardActions>
            )}
          </Card>
        </Grid>
      </Grid>

      {/* Save Button */}
      <Box sx={{ mt: 3, display: "flex", justifyContent: "flex-end" }}>
        <Button
          variant="contained"
          onClick={handleSave}
          disabled={!backendAvailable}
          size="large"
        >
          Save Authentication Settings
        </Button>
      </Box>

      {/* Reset Dialogs */}
      <Dialog
        open={resetGithubDialog}
        onClose={() => setResetGithubDialog(false)}
      >
        <DialogTitle>Reset GitHub OAuth Configuration</DialogTitle>
        <DialogContent>
          <Typography>
            This will reset all GitHub OAuth settings to their default values.
            This action cannot be undone.
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setResetGithubDialog(false)}>Cancel</Button>
          <Button onClick={resetGithubConfig} color="error" variant="contained">
            Reset
          </Button>
        </DialogActions>
      </Dialog>

      <Dialog
        open={resetOAuthDialog}
        onClose={() => setResetOAuthDialog(false)}
      >
        <DialogTitle>Reset Generic OAuth Configuration</DialogTitle>
        <DialogContent>
          <Typography>
            This will reset all generic OAuth settings to their default values.
            This action cannot be undone.
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setResetOAuthDialog(false)}>Cancel</Button>
          <Button onClick={resetOAuthConfig} color="error" variant="contained">
            Reset
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
```

---

## Issue 12: Enhance Notifications with Card Interface

**Problem**: Current notifications lack proper card-based interface and test functionality.

**Solution**: Create card-based notifications with enable/disable toggles and test buttons.

### Implementation Steps

1. **Create enhanced NotificationSettings.jsx**:

```jsx
// file: webui/src/components/NotificationSettings.jsx
import {
  Box,
  Button,
  Card,
  CardContent,
  CardActions,
  FormControlLabel,
  Grid,
  Switch,
  TextField,
  Typography,
  Alert,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  CircularProgress,
  Chip,
} from "@mui/material";
import {
  Email as EmailIcon,
  Chat as DiscordIcon,
  Telegram as TelegramIcon,
  Notifications as PushIcon,
  Webhook as WebhookIcon,
  Delete as DeleteIcon,
  Send as TestIcon,
} from "@mui/icons-material";
import { useEffect, useState } from "react";

/**
 * Enhanced NotificationSettings with card-based UI for each notification method
 */
export default function NotificationSettings({
  config,
  onSave,
  backendAvailable = true,
}) {
  // Email notifications
  const [emailEnabled, setEmailEnabled] = useState(false);
  const [smtpHost, setSmtpHost] = useState("");
  const [smtpPort, setSmtpPort] = useState(587);
  const [smtpUsername, setSmtpUsername] = useState("");
  const [smtpPassword, setSmtpPassword] = useState("");
  const [smtpFrom, setSmtpFrom] = useState("");
  const [smtpTo, setSmtpTo] = useState("");
  const [smtpTLS, setSmtpTLS] = useState(true);

  // Discord notifications
  const [discordEnabled, setDiscordEnabled] = useState(false);
  const [discordWebhook, setDiscordWebhook] = useState("");
  const [discordUsername, setDiscordUsername] = useState("Subtitle Manager");
  const [discordAvatar, setDiscordAvatar] = useState("");

  // Telegram notifications
  const [telegramEnabled, setTelegramEnabled] = useState(false);
  const [telegramToken, setTelegramToken] = useState("");
  const [telegramChatId, setTelegramChatId] = useState("");

  // Push notifications
  const [pushEnabled, setPushEnabled] = useState(false);
  const [pushoverToken, setPushoverToken] = useState("");
  const [pushoverUser, setPushoverUser] = useState("");

  // Webhook notifications
  const [webhookEnabled, setWebhookEnabled] = useState(false);
  const [webhookUrl, setWebhookUrl] = useState("");
  const [webhookMethod, setWebhookMethod] = useState("POST");
  const [webhookHeaders, setWebhookHeaders] = useState("");

  // Dialog states
  const [testDialog, setTestDialog] = useState({
    open: false,
    type: "",
    loading: false,
  });
  const [resetDialog, setResetDialog] = useState(false);

  useEffect(() => {
    if (config) {
      // Email settings
      setEmailEnabled(config.email_enabled || false);
      setSmtpHost(config.smtp_host || "");
      setSmtpPort(config.smtp_port || 587);
      setSmtpUsername(config.smtp_username || "");
      setSmtpPassword(config.smtp_password || "");
      setSmtpFrom(config.smtp_from || "");
      setSmtpTo(config.smtp_to || "");
      setSmtpTLS(config.smtp_tls !== false);

      // Discord settings
      setDiscordEnabled(config.discord_enabled || false);
      setDiscordWebhook(config.discord_webhook || "");
      setDiscordUsername(config.discord_username || "Subtitle Manager");
      setDiscordAvatar(config.discord_avatar || "");

      // Telegram settings
      setTelegramEnabled(config.telegram_enabled || false);
      setTelegramToken(config.telegram_token || "");
      setTelegramChatId(config.telegram_chat_id || "");

      // Push settings
      setPushEnabled(config.push_enabled || false);
      setPushoverToken(config.pushover_token || "");
      setPushoverUser(config.pushover_user || "");

      // Webhook settings
      setWebhookEnabled(config.webhook_enabled || false);
      setWebhookUrl(config.webhook_url || "");
      setWebhookMethod(config.webhook_method || "POST");
      setWebhookHeaders(config.webhook_headers || "");
    }
  }, [config]);

  const handleSave = () => {
    const newConfig = {
      // Email notifications
      email_enabled: emailEnabled,
      smtp_host: smtpHost,
      smtp_port: parseInt(smtpPort),
      smtp_username: smtpUsername,
      smtp_password: smtpPassword,
      smtp_from: smtpFrom,
      smtp_to: smtpTo,
      smtp_tls: smtpTLS,

      // Discord notifications
      discord_enabled: discordEnabled,
      discord_webhook: discordWebhook,
      discord_username: discordUsername,
      discord_avatar: discordAvatar,

      // Telegram notifications
      telegram_enabled: telegramEnabled,
      telegram_token: telegramToken,
      telegram_chat_id: telegramChatId,

      // Push notifications
      push_enabled: pushEnabled,
      pushover_token: pushoverToken,
      pushover_user: pushoverUser,

      // Webhook notifications
      webhook_enabled: webhookEnabled,
      webhook_url: webhookUrl,
      webhook_method: webhookMethod,
      webhook_headers: webhookHeaders,
    };

    onSave(newConfig);
  };

  const testNotification = async (type) => {
    setTestDialog({ open: true, type, loading: true });

    try {
      const response = await fetch(`/api/notifications/test/${type}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          message: "Test notification from Subtitle Manager",
        }),
      });

      if (response.ok) {
        alert(`${type} notification sent successfully!`);
      } else {
        const error = await response.text();
        alert(`Failed to send ${type} notification: ${error}`);
      }
    } catch (error) {
      alert(`Failed to send ${type} notification: ${error.message}`);
    } finally {
      setTestDialog({ open: false, type: "", loading: false });
    }
  };

  const resetAllNotifications = () => {
    setEmailEnabled(false);
    setSmtpHost("");
    setSmtpPort(587);
    setSmtpUsername("");
    setSmtpPassword("");
    setSmtpFrom("");
    setSmtpTo("");
    setSmtpTLS(true);

    setDiscordEnabled(false);
    setDiscordWebhook("");
    setDiscordUsername("Subtitle Manager");
    setDiscordAvatar("");

    setTelegramEnabled(false);
    setTelegramToken("");
    setTelegramChatId("");

    setPushEnabled(false);
    setPushoverToken("");
    setPushoverUser("");

    setWebhookEnabled(false);
    setWebhookUrl("");
    setWebhookMethod("POST");
    setWebhookHeaders("");

    setResetDialog(false);
    alert("All notification settings reset to defaults");
  };

  return (
    <Box sx={{ maxWidth: 1200 }}>
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          mb: 3,
        }}
      >
        <Typography variant="h6">Notification Settings</Typography>
        <Button
          variant="outlined"
          color="error"
          startIcon={<DeleteIcon />}
          onClick={() => setResetDialog(true)}
          disabled={!backendAvailable}
        >
          Reset All Notifications
        </Button>
      </Box>

      {!backendAvailable && (
        <Alert severity="warning" sx={{ mb: 3 }}>
          Backend service is not available. Notification settings cannot be
          modified.
        </Alert>
      )}

      <Grid container spacing={3}>
        {/* Email Notifications Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <EmailIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  Email (SMTP)
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={emailEnabled}
                  onChange={(e) => setEmailEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Send notifications via email using SMTP.
              </Typography>

              {emailEnabled && (
                <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
                  <Grid container spacing={2}>
                    <Grid item xs={8}>
                      <TextField
                        fullWidth
                        label="SMTP Host"
                        value={smtpHost}
                        onChange={(e) => setSmtpHost(e.target.value)}
                        placeholder="smtp.gmail.com"
                        size="small"
                        disabled={!backendAvailable}
                      />
                    </Grid>
                    <Grid item xs={4}>
                      <TextField
                        fullWidth
                        label="Port"
                        type="number"
                        value={smtpPort}
                        onChange={(e) => setSmtpPort(e.target.value)}
                        size="small"
                        disabled={!backendAvailable}
                      />
                    </Grid>
                  </Grid>

                  <TextField
                    fullWidth
                    label="Username"
                    value={smtpUsername}
                    onChange={(e) => setSmtpUsername(e.target.value)}
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="Password"
                    type="password"
                    value={smtpPassword}
                    onChange={(e) => setSmtpPassword(e.target.value)}
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="From Address"
                    value={smtpFrom}
                    onChange={(e) => setSmtpFrom(e.target.value)}
                    placeholder="subtitles@example.com"
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="To Address"
                    value={smtpTo}
                    onChange={(e) => setSmtpTo(e.target.value)}
                    placeholder="admin@example.com"
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <FormControlLabel
                    control={
                      <Switch
                        checked={smtpTLS}
                        onChange={(e) => setSmtpTLS(e.target.checked)}
                        disabled={!backendAvailable}
                      />
                    }
                    label="Use TLS/SSL"
                  />
                </Box>
              )}
            </CardContent>

            {emailEnabled && (
              <CardActions>
                <Button
                  size="small"
                  startIcon={<TestIcon />}
                  onClick={() => testNotification("email")}
                  disabled={
                    !backendAvailable || !smtpHost || !smtpFrom || !smtpTo
                  }
                >
                  Test Email
                </Button>
                <Chip
                  label={
                    smtpHost && smtpFrom && smtpTo ? "Configured" : "Incomplete"
                  }
                  size="small"
                  color={smtpHost && smtpFrom && smtpTo ? "success" : "warning"}
                />
              </CardActions>
            )}
          </Card>
        </Grid>

        {/* Discord Notifications Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <DiscordIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  Discord
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={discordEnabled}
                  onChange={(e) => setDiscordEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Send notifications to Discord channels via webhooks.
              </Typography>

              {discordEnabled && (
                <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
                  <TextField
                    fullWidth
                    label="Webhook URL"
                    value={discordWebhook}
                    onChange={(e) => setDiscordWebhook(e.target.value)}
                    placeholder="https://discord.com/api/webhooks/..."
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="Bot Username"
                    value={discordUsername}
                    onChange={(e) => setDiscordUsername(e.target.value)}
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="Avatar URL"
                    value={discordAvatar}
                    onChange={(e) => setDiscordAvatar(e.target.value)}
                    placeholder="https://example.com/avatar.png (optional)"
                    size="small"
                    disabled={!backendAvailable}
                  />
                </Box>
              )}
            </CardContent>

            {discordEnabled && (
              <CardActions>
                <Button
                  size="small"
                  startIcon={<TestIcon />}
                  onClick={() => testNotification("discord")}
                  disabled={!backendAvailable || !discordWebhook}
                >
                  Test Discord
                </Button>
                <Chip
                  label={discordWebhook ? "Configured" : "Incomplete"}
                  size="small"
                  color={discordWebhook ? "success" : "warning"}
                />
              </CardActions>
            )}
          </Card>
        </Grid>

        {/* Telegram Notifications Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <TelegramIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  Telegram
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={telegramEnabled}
                  onChange={(e) => setTelegramEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Send notifications to Telegram users or groups.
              </Typography>

              {telegramEnabled && (
                <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
                  <TextField
                    fullWidth
                    label="Bot Token"
                    value={telegramToken}
                    onChange={(e) => setTelegramToken(e.target.value)}
                    placeholder="123456789:ABCdefGhIJKlmNoPQRstuVWXyz"
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="Chat ID"
                    value={telegramChatId}
                    onChange={(e) => setTelegramChatId(e.target.value)}
                    placeholder="-1234567890"
                    size="small"
                    disabled={!backendAvailable}
                  />
                </Box>
              )}
            </CardContent>

            {telegramEnabled && (
              <CardActions>
                <Button
                  size="small"
                  startIcon={<TestIcon />}
                  onClick={() => testNotification("telegram")}
                  disabled={
                    !backendAvailable || !telegramToken || !telegramChatId
                  }
                >
                  Test Telegram
                </Button>
                <Chip
                  label={
                    telegramToken && telegramChatId
                      ? "Configured"
                      : "Incomplete"
                  }
                  size="small"
                  color={
                    telegramToken && telegramChatId ? "success" : "warning"
                  }
                />
              </CardActions>
            )}
          </Card>
        </Grid>

        {/* Push Notifications Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <PushIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  Push Notifications
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={pushEnabled}
                  onChange={(e) => setPushEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Send notifications via Pushover service.
              </Typography>

              {pushEnabled && (
                <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
                  <TextField
                    fullWidth
                    label="User Key"
                    value={pushoverUser}
                    onChange={(e) => setPushoverUser(e.target.value)}
                    placeholder="Your Pushover user key"
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <TextField
                    fullWidth
                    label="API Token"
                    value={pushoverToken}
                    onChange={(e) => setPushoverToken(e.target.value)}
                    placeholder="Your Pushover API token"
                    size="small"
                    disabled={!backendAvailable}
                  />
                </Box>
              )}
            </CardContent>

            {pushEnabled && (
              <CardActions>
                <Button
                  size="small"
                  startIcon={<TestIcon />}
                  onClick={() => testNotification("push")}
                  disabled={
                    !backendAvailable || !pushoverUser || !pushoverToken
                  }
                >
                  Test Push
                </Button>
                <Chip
                  label={
                    pushoverUser && pushoverToken ? "Configured" : "Incomplete"
                  }
                  size="small"
                  color={pushoverUser && pushoverToken ? "success" : "warning"}
                />
              </CardActions>
            )}
          </Card>
        </Grid>

        {/* Webhook Notifications Card */}
        <Grid item xs={12} md={6}>
          <Card sx={{ height: "100%" }}>
            <CardContent>
              <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
                <WebhookIcon sx={{ mr: 1, color: "primary.main" }} />
                <Typography variant="h6" color="primary">
                  Webhook
                </Typography>
                <Box sx={{ flexGrow: 1 }} />
                <Switch
                  checked={webhookEnabled}
                  onChange={(e) => setWebhookEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              </Box>

              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Send notifications to a custom URL via HTTP requests.
              </Typography>

              {webhookEnabled && (
                <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
                  <TextField
                    fullWidth
                    label="Webhook URL"
                    value={webhookUrl}
                    onChange={(e) => setWebhookUrl(e.target.value)}
                    placeholder="https://example.com/webhook"
                    size="small"
                    disabled={!backendAvailable}
                  />

                  <FormControl fullWidth>
                    <InputLabel>Method</InputLabel>
                    <Select
                      value={webhookMethod}
                      label="Method"
                      onChange={(e) => setWebhookMethod(e.target.value)}
                      disabled={!backendAvailable}
                    >
                      <MenuItem value="POST">POST</MenuItem>
                      <MenuItem value="GET">GET</MenuItem>
                    </Select>
                  </FormControl>

                  <TextField
                    fullWidth
                    label="Headers (JSON)"
                    value={webhookHeaders}
                    onChange={(e) => setWebhookHeaders(e.target.value)}
                    placeholder='{"Authorization": "Bearer token"}'
                    size="small"
                    disabled={!backendAvailable}
                  />
                </Box>
              )}
            </CardContent>

            {webhookEnabled && (
              <CardActions>
                <Button
                  size="small"
                  startIcon={<TestIcon />}
                  onClick={() => testNotification("webhook")}
                  disabled={!backendAvailable || !webhookUrl}
                >
                  Test Webhook
                </Button>
                <Chip
                  label={webhookUrl ? "Configured" : "Incomplete"}
                  size="small"
                  color={webhookUrl ? "success" : "warning"}
                />
              </CardActions>
            )}
          </Card>
        </Grid>
      </Grid>

      {/* Save Button */}
      <Box sx={{ mt: 3, display: "flex", justifyContent: "flex-end" }}>
        <Button
          variant="contained"
          onClick={handleSave}
          disabled={!backendAvailable}
          size="large"
        >
          Save Notification Settings
        </Button>
      </Box>

      {/* Test Dialog */}
      <Dialog
        open={testDialog.open}
        onClose={() =>
          !testDialog.loading &&
          setTestDialog({ open: false, type: "", loading: false })
        }
      >
        <DialogTitle>Testing {testDialog.type} Notification</DialogTitle>
        <DialogContent>
          {testDialog.loading ? (
            <Box sx={{ display: "flex", alignItems: "center", gap: 2, py: 2 }}>
              <CircularProgress size={24} />
              <Typography>Sending test notification...</Typography>
            </Box>
          ) : (
            <Typography>Test notification sent!</Typography>
          )}
        </DialogContent>
        {!testDialog.loading && (
          <DialogActions>
            <Button
              onClick={() =>
                setTestDialog({ open: false, type: "", loading: false })
              }
            >
              Close
            </Button>
          </DialogActions>
        )}
      </Dialog>

      {/* Reset Dialog */}
      <Dialog open={resetDialog} onClose={() => setResetDialog(false)}>
        <DialogTitle>Reset All Notification Settings</DialogTitle>
        <DialogContent>
          <Typography>
            This will reset all notification settings to their default values.
            This action cannot be undone.
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setResetDialog(false)}>Cancel</Button>
          <Button
            onClick={resetAllNotifications}
            color="error"
            variant="contained"
          >
            Reset All
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
```

---

## Issues 13-18: Final Components and Provider Fixes

This section covers the remaining issues 13-18: OAuth2 API endpoints, Languages page, Scheduler settings, and provider configuration fixes.

## Issue 13: Create OAuth2 API Endpoints

**Problem**: Frontend requires backend support for managing OAuth2 credentials (client ID and secret).

**Solution**: Implement backend API endpoints for creating, regenerating, and resetting OAuth2 credentials.

### Implementation Steps

1. **Create OAuth2 API endpoints in web server**:

```go
// file: pkg/webserver/oauth_management.go

// OAuthCredentials represents OAuth2 credentials
type OAuthCredentials struct {
    ClientID     string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
    RedirectURL  string `json:"redirect_url,omitempty"`
}

// generateSecureToken generates a cryptographically secure random token
func generateSecureToken(length int) (string, error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// HandleGitHubOAuthGenerate generates new GitHub OAuth credentials
// POST /api/oauth/github/generate
func (s *Server) HandleGitHubOAuthGenerate(w http.ResponseWriter, r *http.Request) {
    // Check admin permissions
    if !s.hasPermission(r, "admin") {
        http.Error(w, "Admin access required", http.StatusForbidden)
        return
    }

    // Generate new client ID and secret
    clientID, err := generateSecureToken(16) // 32 character hex string
    if err != nil {
        http.Error(w, "Failed to generate client ID", http.StatusInternalServerError)
        return
    }

    clientSecret, err := generateSecureToken(32) // 64 character hex string
    if err != nil {
        http.Error(w, "Failed to generate client secret", http.StatusInternalServerError)
        return
    }

    credentials := OAuthCredentials{
        ClientID:     "gh_" + clientID,
        ClientSecret: "ghs_" + clientSecret,
        RedirectURL:  r.Header.Get("Origin") + "/api/oauth/github/callback",
    };

    // Store in configuration
    s.config.GitHubClientID = credentials.ClientID
    s.config.GitHubClientSecret = credentials.ClientSecret
    s.config.GitHubRedirectURL = credentials.RedirectURL

    // Save configuration
    if err := s.saveConfig(); err != nil {
        http.Error(w, "Failed to save configuration", http.StatusInternalServerError)
        return
    }

    // Log the action
    s.logger.Info("Generated new GitHub OAuth credentials", "admin", s.getUserFromRequest(r))

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(credentials)
}

// HandleGitHubOAuthRegenerate regenerates GitHub OAuth client secret only
// POST /api/oauth/github/regenerate
func (s *Server) HandleGitHubOAuthRegenerate(w http.ResponseWriter, r *http.Request) {
    // Check admin permissions
    if !s.hasPermission(r, "admin") {
        http.Error(w, "Admin access required", http.StatusForbidden)
        return
    }

    // Validate that client ID exists
    if s.config.GitHubClientID == "" {
        http.Error(w, "No existing GitHub OAuth configuration found", http.StatusBadRequest)
        return
    }

    // Generate new client secret only
    clientSecret, err := generateSecureToken(32)
    if err != nil {
        http.Error(w, "Failed to generate client secret", http.StatusInternalServerError)
        return
    }

    credentials := OAuthCredentials{
        ClientID:     s.config.GitHubClientID, // Keep existing client ID
        ClientSecret: "ghs_" + clientSecret,
        RedirectURL:  s.config.GitHubRedirectURL,
    };

    // Update configuration with new secret only
    s.config.GitHubClientSecret = credentials.ClientSecret

    // Save configuration
    if err := s.saveConfig(); err != nil {
        http.Error(w, "Failed to save configuration", http.StatusInternalServerError)
        return
    }

    // Log the action
    s.logger.Info("Regenerated GitHub OAuth client secret", "admin", s.getUserFromRequest(r))

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(credentials)
}

// HandleGitHubOAuthReset resets GitHub OAuth configuration to defaults
// POST /api/oauth/github/reset
func (s *Server) HandleGitHubOAuthReset(w http.ResponseWriter, r *http.Request) {
    // Check admin permissions
    if !s.hasPermission(r, "admin") {
        http.Error(w, "Admin access required", http.StatusForbidden)
        return
    }

    // Reset to default values
    s.config.GitHubClientID = ""
    s.config.GitHubClientSecret = ""
    s.config.GitHubRedirectURL = ""
    s.config.GitHubOAuthEnabled = false

    // Save configuration
    if err := s.saveConfig(); err != nil {
        http.Error(w, "Failed to save configuration", http.StatusInternalServerError)
        return
    }

    // Log the action
    s.logger.Info("Reset GitHub OAuth configuration", "admin", s.getUserFromRequest(r))

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("GitHub OAuth configuration reset successfully"))
}

// Helper method to check user permissions
func (s *Server) hasPermission(r *http.Request, requiredRole string) bool {
                      divider
                      sx={{
                        border: "1px solid",
                        borderColor: "divider",
                        borderRadius: 1,
                        mb: 1,
                        bgcolor: profile.enabled
                          ? "background.paper"
                          : "action.disabled",
                      }}
                    >
                      <Box
                        sx={{ display: "flex", alignItems: "center", mr: 1 }}
                      >
                        <DragIcon sx={{ color: "text.secondary" }} />
                      </Box>

                      <ListItemText
                        primary={
                          <Box
                            sx={{
                              display: "flex",
                              alignItems: "center",
                              gap: 1,
                            }}
                          >
                            <Typography variant="subtitle2">
                              {profile.name}
                            </Typography>
                            <Chip
                              label={profile.enabled ? "Enabled" : "Disabled"}
                              size="small"
                              color={profile.enabled ? "success" : "default"}
                            />
                          </Box>
                        }
                        secondary={
                          <Box>
                            <Typography
                              variant="caption"
                              color="text.secondary"
                            >
                              {profile.description}
                            </Typography>
                            <Box sx={{ mt: 0.5 }}>
                              {profile.languages.map((langCode) => (
                                <Chip
                                  key={langCode}
                                  label={getLanguageName(langCode)}
                                  size="small"
                                  variant="outlined"
                                  sx={{ mr: 0.5, mb: 0.5 }}
                                />
                              ))}
                            </Box>
                          </Box>
                        }
                      />

                      <ListItemSecondaryAction>
                        <IconButton
                          edge="end"
                          onClick={() => handleEditProfile(profile)}
                          disabled={!backendAvailable}
                          size="small"
                        >
                          <EditIcon />
                        </IconButton>
                        <IconButton
                          edge="end"
                          onClick={() => handleDeleteProfile(profile.id)}
                          disabled={!backendAvailable}
                          size="small"
                          color="error"
                        >
                          <DeleteIcon />
                        </IconButton>
                      </ListItemSecondaryAction>
                    </ListItem>
                  ))}
                </List>
              )}
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      {/* Save Button */}
      <Box sx={{ mt: 3, display: "flex", justifyContent: "flex-end" }}>
        <Button
          variant="contained"
          onClick={handleSave}
          disabled={!backendAvailable}
          size="large"
        >
          Save Language Settings
        </Button>
      </Box>

      {/* Profile Dialog */}
      <Dialog
        open={profileDialog.open}
        onClose={() =>
          setProfileDialog({ open: false, profile: null, isEditing: false })
        }
        maxWidth="sm"
        fullWidth
      >
        <DialogTitle>
          {profileDialog.isEditing
            ? "Edit Language Profile"
            : "Create Language Profile"}
        </DialogTitle>
        <DialogContent>
          <Box sx={{ display: "flex", flexDirection: "column", gap: 2, pt: 1 }}>
            <TextField
              fullWidth
              label="Profile Name"
              value={profileDialog.profile?.name || ""}
              onChange={(e) =>
                setProfileDialog((prev) => ({
                  ...prev,
                  profile: { ...prev.profile, name: e.target.value },
                }))
              }
              placeholder="e.g., Movies, TV Shows, Anime"
            />

            <TextField
              fullWidth
              label="Description"
              value={profileDialog.profile?.description || ""}
              onChange={(e) =>
                setProfileDialog((prev) => ({
                  ...prev,
                  profile: { ...prev.profile, description: e.target.value },
                }))
              }
              placeholder="Brief description of this profile"
              multiline
              rows={2}
            />

            <Autocomplete
              multiple
              options={availableLanguages}
              getOptionLabel={(option) => `${option.name} (${option.code})`}
              value={availableLanguages.filter((lang) =>
                (profileDialog.profile?.languages || []).includes(lang.code),
              )}
              onChange={(event, newValue) => {
                setProfileDialog((prev) => ({
                  ...prev,
                  profile: {
                    ...prev.profile,
                    languages: newValue.map((v) => v.code),
                  },
                }));
              }}
              renderTags={(value, getTagProps) =>
                value.map((option, index) => (
                  <Chip
                    variant="outlined"
                    label={`${option.name} (${option.code})`}
                    {...getTagProps({ index })}
                    key={option.code}
                  />
                ))
              }
              renderInput={(params) => (
                <TextField
                  {...params}
                  label="Languages"
                  placeholder="Select languages for this profile"
                />
              )}
            />

            <FormControlLabel
              control={
                <Switch
                  checked={profileDialog.profile?.enabled !== false}
                  onChange={(e) =>
                    setProfileDialog((prev) => ({
                      ...prev,
                      profile: { ...prev.profile, enabled: e.target.checked },
                    }))
                  }
                />
              }
              label="Enable Profile"
            />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() =>
              setProfileDialog({ open: false, profile: null, isEditing: false })
            }
          >
            Cancel
          </Button>
          <Button
            onClick={handleSaveProfile}
            variant="contained"
            disabled={
              !profileDialog.profile?.name ||
              !profileDialog.profile?.languages?.length
            }
          >
            {profileDialog.isEditing ? "Update" : "Create"} Profile
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
```

2. **Add Languages tab to Settings.jsx**:

```jsx
// file: webui/src/Settings.jsx

// Add import
import LanguagesSettings from "./components/LanguagesSettings.jsx";
import { Translate as LanguagesIcon } from "@mui/icons-material";

// Add to tabs array
const tabs = [
  // ... existing tabs ...
  {
    label: "Languages",
    icon: <LanguagesIcon />,
    component: () => (
      <LanguagesSettings
        config={_config}
        onSave={saveSettings}
        backendAvailable={backendAvailable}
      />
    ),
  },
  // ... rest of tabs ...
];
```

---

## Issue 16: Add Scheduler Settings

**Problem**: Missing scheduler configuration options that Bazarr provides.

**Solution**: Integrate scheduler settings into the general settings page or create a separate section.

### Implementation Steps

1. **Add scheduler section to GeneralSettings.jsx**:

```jsx
// file: webui/src/components/GeneralSettings.jsx

// Add to existing GeneralSettings component after the Backup section:

{
  /* Scheduler Configuration */
}
<Card sx={{ mb: 3 }}>
  <CardContent>
    <Typography variant="h6" gutterBottom color="primary">
      Scheduler
    </Typography>
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <FormGroup>
          <FormControlLabel
            control={
              <Switch
                checked={schedulerEnabled}
                onChange={(e) => setSchedulerEnabled(e.target.checked)}
                disabled={!backendAvailable}
              />
            }
            label="Enable Automatic Scheduling"
          />
        </FormGroup>
      </Grid>

      {schedulerEnabled && (
        <>
          <Grid item xs={12} md={6}>
            <FormControl fullWidth>
              <InputLabel>Library Scan Frequency</InputLabel>
              <Select
                value={libraryScanFreq}
                label="Library Scan Frequency"
                onChange={(e) => setLibraryScanFreq(e.target.value)}
                disabled={!backendAvailable}
              >
                <MenuItem value="never">Never</MenuItem>
                <MenuItem value="hourly">Every Hour</MenuItem>
                <MenuItem value="daily">Daily</MenuItem>
                <MenuItem value="weekly">Weekly</MenuItem>
                <MenuItem value="monthly">Monthly</MenuItem>
                <MenuItem value="custom">Custom Cron</MenuItem>
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={12} md={6}>
            <FormControl fullWidth>
              <InputLabel>Wanted Search Frequency</InputLabel>
              <Select
                value={wantedSearchFreq}
                label="Wanted Search Frequency"
                onChange={(e) => setWantedSearchFreq(e.target.value)}
                disabled={!backendAvailable}
              >
                <MenuItem value="never">Never</MenuItem>
                <MenuItem value="15min">Every 15 Minutes</MenuItem>
                <MenuItem value="30min">Every 30 Minutes</MenuItem>
                <MenuItem value="hourly">Every Hour</MenuItem>
                <MenuItem value="daily">Daily</MenuItem>
                <MenuItem value="custom">Custom Cron</MenuItem>
              </Select>
            </FormControl>
          </Grid>

          {(libraryScanFreq === "custom" || wantedSearchFreq === "custom") && (
            <Grid item xs={12}>
              <Alert severity="info" sx={{ mb: 2 }}>
                Cron expressions follow standard format: minute hour day month
                weekday
                <br />
                Examples: "0 2 * * *" (daily at 2 AM), "*/15 * * * *" (every 15
                minutes)
              </Alert>

              {libraryScanFreq === "custom" && (
                <TextField
                  fullWidth
                  label="Library Scan Cron Expression"
                  value={libraryScanCron}
                  onChange={(e) => setLibraryScanCron(e.target.value)}
                  placeholder="0 2 * * *"
                  sx={{ mb: 2 }}
                  disabled={!backendAvailable}
                />
              )}

              {wantedSearchFreq === "custom" && (
                <TextField
                  fullWidth
                  label="Wanted Search Cron Expression"
                  value={wantedSearchCron}
                  onChange={(e) => setWantedSearchCron(e.target.value)}
                  placeholder="*/30 * * * *"
                  disabled={!backendAvailable}
                />
              )}
            </Grid>
          )}

          <Grid item xs={12} md={6}>
            <TextField
              fullWidth
              label="Max Concurrent Downloads"
              type="number"
              value={maxConcurrentDownloads}
              onChange={(e) => setMaxConcurrentDownloads(e.target.value)}
              inputProps={{ min: 1, max: 10 }}
              helperText="Maximum number of simultaneous subtitle downloads"
              disabled={!backendAvailable}
            />
          </Grid>

          <Grid item xs={12} md={6}>
            <TextField
              fullWidth
              label="Download Timeout (seconds)"
              type="number"
              value={downloadTimeout}
              onChange={(e) => setDownloadTimeout(e.target.value)}
              inputProps={{ min: 10, max: 300 }}
              helperText="Timeout for individual subtitle downloads"
              disabled={!backendAvailable}
            />
          </Grid>
        </>
      )}
    </Grid>
  </CardContent>
</Card>;
```

---

## Issues 17-18: Fix Provider Configuration

**Problem**: Provider configuration modals don't work properly - dropdowns are broken and text has typos.

**Solution**: Fix ProviderConfigDialog component with proper provider selection and configuration.

### Implementation Steps

1. **Fix ProviderConfigDialog.jsx**:

```jsx
// file: webui/src/components/ProviderConfigDialog.jsx
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Box,
  Typography,
  Alert,
  Grid,
  FormControlLabel,
  Switch,
  Autocomplete,
  Chip,
} from "@mui/material";
import { useEffect, useState } from "react";

/**
 * Fixed ProviderConfigDialog with proper provider selection and configuration options
 */
export default function ProviderConfigDialog({
  open,
  provider,
  onClose,
  onSave,
}) {
  const [selectedProvider, setSelectedProvider] = useState("");
  const [config, setConfig] = useState({});
  const [availableProviders, setAvailableProviders] = useState([]);

  // Load available providers when dialog opens
  useEffect(() => {
    if (open && !provider) {
      loadAvailableProviders();
    } else if (provider) {
      setSelectedProvider(provider.name);
      setConfig(provider.config || {});
    }
  }, [open, provider]);

  const loadAvailableProviders = async () => {
    try {
      const response = await fetch("/api/providers/available");
      if (response.ok) {
        const providers = await response.json();
        setAvailableProviders(providers);
      }
    } catch (error) {
      console.error("Failed to load available providers:", error);
    }
  };

  const getProviderDisplayName = (name) => {
    const displayNames = {
      opensubtitles: "OpenSubtitles.org",
      opensubtitlescom: "OpenSubtitles.com",
      addic7ed: "Addic7ed",
      subscene: "Subscene",
      podnapisi: "Podnapisi.NET",
      yifysubtitles: "YIFY Subtitles",
      embedded: "Embedded Subtitles",
      // Add more as needed
    };
    return displayNames[name] || name.charAt(0).toUpperCase() + name.slice(1);
  };

  const getProviderConfigFields = (providerName) => {
    const configs = {
      opensubtitles: [
        { key: "api_key", label: "API Key", type: "password", required: true },
        {
          key: "user_agent",
          label: "User Agent",
          type: "text",
          required: true,
        },
        { key: "enabled", label: "Enabled", type: "boolean", default: true },
      ],
      opensubtitlescom: [
        { key: "api_key", label: "API Key", type: "password", required: true },
        { key: "enabled", label: "Enabled", type: "boolean", default: true },
      ],
      addic7ed: [
        { key: "username", label: "Username", type: "text", required: true },
        {
          key: "password",
          label: "Password",
          type: "password",
          required: true,
        },
        { key: "enabled", label: "Enabled", type: "boolean", default: true },
      ],
      subscene: [
        { key: "enabled", label: "Enabled", type: "boolean", default: true },
        {
          key: "timeout",
          label: "Timeout (seconds)",
          type: "number",
          default: 30,
        },
      ],
      embedded: [
        { key: "enabled", label: "Enabled", type: "boolean", default: true },
        {
          key: "extract_mkv",
          label: "Extract from MKV",
          type: "boolean",
          default: true,
        },
        {
          key: "extract_mp4",
          label: "Extract from MP4",
          type: "boolean",
          default: true,
        },
        {
          key: "ffmpeg_path",
          label: "FFmpeg Path",
          type: "text",
          placeholder: "/usr/bin/ffmpeg",
        },
      ],
      // Add configurations for other providers
    };

    return (
      configs[providerName] || [
        { key: "enabled", label: "Enabled", type: "boolean", default: true },
      ]
    );
  };

  const handleProviderChange = (newProvider) => {
    setSelectedProvider(newProvider);
    // Reset config when provider changes
    const fields = getProviderConfigFields(newProvider);
    const newConfig = {};
    fields.forEach((field) => {
      if (field.default !== undefined) {
        newConfig[field.key] = field.default;
      }
    });
    setConfig(newConfig);
  };

  const handleConfigChange = (key, value) => {
    setConfig((prev) => ({
      ...prev,
      [key]: value,
    }));
  };

  const handleSave = () => {
    if (!selectedProvider) {
      alert("Please select a provider");
      return;
    }

    const providerData = {
      name: selectedProvider,
      displayName: getProviderDisplayName(selectedProvider),
      config: config,
      enabled: config.enabled !== false,
    };

    onSave(providerData);
    onClose();
  };

  const renderConfigField = (field) => {
    const value = config[field.key] ?? field.default ?? "";

    switch (field.type) {
      case "boolean":
        return (
          <FormControlLabel
            key={field.key}
            control={
              <Switch
                checked={!!value}
                onChange={(e) =>
                  handleConfigChange(field.key, e.target.checked)
                }
              />
            }
            label={field.label}
          />
        );

      case "number":
        return (
          <TextField
            key={field.key}
            fullWidth
            label={field.label}
            type="number"
            value={value}
            onChange={(e) =>
              handleConfigChange(field.key, parseInt(e.target.value) || 0)
            }
            required={field.required}
            placeholder={field.placeholder}
            sx={{ mb: 2 }}
          />
        );

      case "password":
        return (
          <TextField
            key={field.key}
            fullWidth
            label={field.label}
            type="password"
            value={value}
            onChange={(e) => handleConfigChange(field.key, e.target.value)}
            required={field.required}
            placeholder={field.placeholder}
            sx={{ mb: 2 }}
          />
        );

      default:
        return (
          <TextField
            key={field.key}
            fullWidth
            label={field.label}
            value={value}
            onChange={(e) => handleConfigChange(field.key, e.target.value)}
            required={field.required}
            placeholder={field.placeholder}
            sx={{ mb: 2 }}
          />
        );
    }
  };

  const isValid = () => {
    if (!selectedProvider) return false;

    const fields = getProviderConfigFields(selectedProvider);
    return fields.every((field) => {
      if (field.required) {
        const value = config[field.key];
        return value !== undefined && value !== null && value !== "";
      }
      return true;
    });
  };

  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>
        {provider ? "Configure Provider" : "Configure Custom Provider"}
      </DialogTitle>

      <DialogContent>
        <Box sx={{ pt: 1 }}>
          {!provider && (
            <>
              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Configure this provider to enable subtitle downloads. Required
                fields are marked with an asterisk (*).
              </Typography>

              <FormControl fullWidth sx={{ mb: 3 }}>
                <InputLabel>Provider</InputLabel>
                <Select
                  value={selectedProvider}
                  label="Provider"
                  onChange={(e) => handleProviderChange(e.target.value)}
                >
                  {availableProviders.map((p) => (
                    <MenuItem key={p.name} value={p.name}>
                      {getProviderDisplayName(p.name)}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </>
          )}

          {(provider || selectedProvider) && (
            <>
              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Configure this provider to enable subtitle downloads. Required
                fields are marked with an asterisk (*).
              </Typography>

              <Alert severity="info" sx={{ mb: 2 }}>
                Provider:{" "}
                <strong>
                  {getProviderDisplayName(provider?.name || selectedProvider)}
                </strong>
              </Alert>

              <Box>
                {getProviderConfigFields(
                  provider?.name || selectedProvider,
                ).map(renderConfigField)}
              </Box>
            </>
          )}
        </Box>
      </DialogContent>

      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button onClick={handleSave} variant="contained" disabled={!isValid()}>
          Save Configuration
        </Button>
      </DialogActions>
    </Dialog>
  );
}
```

## Time Estimates

- Issues 9-11 (OAuth2 API Endpoints): 6-8 hours
- Issue 15 (Languages Page): 12-16 hours
- Issue 16 (Scheduler Settings): 4-6 hours
- Issues 17-18 (Provider Config Fixes): 6-8 hours

**Total Section 5**: 28-38 hours

## Final Implementation Summary

### Total Project Time Estimate: 73-101 hours

**Section Breakdown:**

- Section 1 (User Management & Navigation): 9-13 hours
- Section 2 (Navigation Improvements): 9-13 hours
- Section 3 (Settings Enhancement): 14-20 hours
- Section 4 (Authentication & Notifications): 22-30 hours
- Section 5 (Final Components): 28-38 hours

### Priority Implementation Order

1. **High Priority**: Issues 1-5 (Navigation & User Management)
2. **Medium Priority**: Issues 6-12 (Settings & Authentication)
3. **Lower Priority**: Issues 13-18 (Languages, Scheduler, Providers)

The implementation plan provides comprehensive solutions for all identified UI/UX issues with detailed code examples, testing procedures, and realistic time estimates for junior developers.

---

## Document Information

**Generated:** Mon Jun 16 18:42:53 EDT 2025
**Source Files:** 6 implementation plan sections
**Total Length:** 3907 lines
**Repository:** subtitle-manager
**Purpose:** Complete UI/UX implementation reference guide

### Quick Navigation

- [Section 0: Project Summary](#section-0-project-summary-and-overview)
- [Section 1: User Management & Navigation](#section-1-user-management-and-navigation-basics)
- [Section 2: Navigation Improvements](#section-2-navigation-improvements)
- [Section 3: Settings Enhancement](#section-3-settings-page-enhancement)
- [Section 4: Authentication & Notifications](#section-4-authentication-and-notifications)
- [Section 5: Final Components & Provider Fixes](#section-5-final-components-and-provider-fixes)

### Implementation Notes

This combined document provides all necessary code samples, implementation steps, and testing procedures to complete the UI/UX overhaul of the Subtitle Manager frontend. Each section can be implemented independently while maintaining system functionality.

**For questions or clarifications, refer to the individual section files or the project documentation.**
