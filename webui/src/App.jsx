import {
  ArrowBack as BackIcon,
  Transform as ConvertIcon,
  Brightness4 as DarkModeIcon,
  Dashboard as DashboardIcon,
  Archive as ExtractIcon,
  History as HistoryIcon,
  ChildFriendly as KidModeIcon,
  VideoLibrary as LibraryIcon,
  Brightness7 as LightModeIcon,
  Menu as MenuIcon,
  PushPin as PinIcon,
  Settings as SettingsIcon,
  BugReport as SystemIcon,
  Translate as TranslateIcon,
  Download as WantedIcon,
  Schedule as ScheduleIcon,
} from '@mui/icons-material';
import {
  Alert,
  alpha,
  AppBar,
  Box,
  Button,
  CircularProgress,
  Container,
  createTheme,
  CssBaseline,
  Divider,
  Drawer,
  Fab,
  IconButton,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  ListSubheader,
  Paper,
  TextField,
  ThemeProvider,
  Toolbar,
  Typography,
} from '@mui/material';
import { lazy, Suspense, useEffect, useState } from 'react';
import {
  NavLink,
  Route,
  BrowserRouter as Router,
  Routes,
  useLocation,
  useNavigate,
} from 'react-router-dom';
import './App.css';
import OfflineInfo from './OfflineInfo.jsx';
import LoadingComponent from './components/LoadingComponent.jsx';
import { apiService, getBasePath } from './services/api.js';

// Lazy load components for better performance
const Dashboard = lazy(() => import('./Dashboard.jsx'));
const MediaLibrary = lazy(() => import('./MediaLibrary.jsx'));
const Wanted = lazy(() => import('./Wanted.jsx'));
const History = lazy(() => import('./History.jsx'));
const Settings = lazy(() => import('./Settings.jsx'));
const System = lazy(() => import('./System.jsx'));
const Extract = lazy(() => import('./Extract.jsx'));
const Convert = lazy(() => import('./Convert.jsx'));
const Translate = lazy(() => import('./Translate.jsx'));
const Scheduling = lazy(() => import('./Scheduling.jsx'));
const Setup = lazy(() => import('./Setup.jsx'));
const MediaDetails = lazy(() => import('./MediaDetails.jsx'));

/**
 * Creates a Material Design 3 compliant theme with enhanced dark mode support
 * Following Google's Material Design guidelines for color schemes and accessibility
 * @param {boolean} isDarkMode - Whether to use dark mode
 * @param {boolean} kidMode - Increase font sizes for children if true
 * @returns {Theme} Material UI theme object
 */
const createAppTheme = (isDarkMode = true, kidMode = false) =>
  createTheme({
    palette: {
      mode: isDarkMode ? 'dark' : 'light',
      ...(isDarkMode
        ? {
            // Dark mode - Material Design 3 dark surface colors
            primary: {
              main: '#bb86fc',
              light: '#d0bcff',
              dark: '#9965f4',
              contrastText: '#000000',
            },
            secondary: {
              main: '#03dac6',
              light: '#66fff9',
              dark: '#00a896',
              contrastText: '#000000',
            },
            background: {
              default: '#121212',
              paper: '#1e1e1e',
            },
            surface: {
              main: '#1e1e1e',
              variant: '#2d2d30',
            },
            text: {
              primary: '#e1e1e1',
              secondary: '#b3b3b3',
            },
            error: {
              main: '#cf6679',
              light: '#ff94a3',
              dark: '#9a4052',
            },
            warning: {
              main: '#ffb74d',
              light: '#ffe97d',
              dark: '#c88719',
            },
            info: {
              main: '#64b5f6',
              light: '#9be7ff',
              dark: '#2286c3',
            },
            success: {
              main: '#81c784',
              light: '#b2fab4',
              dark: '#519657',
            },
            divider: alpha('#ffffff', 0.12),
          }
        : {
            // Light mode - Material Design 3 light surface colors
            primary: {
              main: '#6750a4',
              light: '#9a82db',
              dark: '#21005d',
              contrastText: '#ffffff',
            },
            secondary: {
              main: '#625b71',
              light: '#8e82a2',
              dark: '#463d52',
              contrastText: '#ffffff',
            },
            background: {
              default: '#fffbfe',
              paper: '#ffffff',
            },
            surface: {
              main: '#ffffff',
              variant: '#f4f0f7',
            },
            text: {
              primary: '#1c1b1f',
              secondary: '#49454f',
            },
            divider: alpha('#000000', 0.12),
          }),
    },
    typography: {
      fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
      htmlFontSize: kidMode ? 20 : 16,
      h1: {
        fontSize: '2.5rem',
        fontWeight: 400,
        letterSpacing: '-0.01562em',
      },
      h4: {
        fontSize: '2.125rem',
        fontWeight: 400,
        letterSpacing: '0.00735em',
      },
      h6: {
        fontSize: '1.25rem',
        fontWeight: 500,
        letterSpacing: '0.0075em',
      },
      body1: {
        fontSize: '1rem',
        letterSpacing: '0.00938em',
      },
      body2: {
        fontSize: '0.875rem',
        letterSpacing: '0.01071em',
      },
      button: {
        fontSize: '0.875rem',
        fontWeight: 500,
        letterSpacing: '0.02857em',
        textTransform: 'uppercase',
      },
    },
    shape: {
      borderRadius: 12, // More rounded corners for modern look
    },
    components: {
      MuiAppBar: {
        styleOverrides: {
          root: {
            backgroundColor: isDarkMode ? '#1e1e1e' : '#6750a4',
            backdropFilter: 'blur(8px)',
            boxShadow: '0 2px 8px rgba(0,0,0,0.1)',
          },
        },
      },
      MuiDrawer: {
        styleOverrides: {
          paper: {
            backgroundColor: isDarkMode ? '#1e1e1e' : '#ffffff',
            borderRight: `1px solid ${alpha(isDarkMode ? '#ffffff' : '#000000', 0.12)}`,
          },
        },
      },
      MuiCard: {
        styleOverrides: {
          root: {
            backgroundColor: isDarkMode ? '#1e1e1e' : '#ffffff',
            boxShadow: isDarkMode
              ? '0 4px 8px rgba(0,0,0,0.3)'
              : '0 2px 8px rgba(0,0,0,0.1)',
            borderRadius: 16,
            border: `1px solid ${alpha(isDarkMode ? '#ffffff' : '#000000', 0.08)}`,
          },
        },
      },
      MuiPaper: {
        styleOverrides: {
          root: {
            backgroundColor: isDarkMode ? '#1e1e1e' : '#ffffff',
            '&.MuiPaper-outlined': {
              border: `1px solid ${alpha(isDarkMode ? '#ffffff' : '#000000', 0.12)}`,
            },
          },
        },
      },
      MuiButton: {
        styleOverrides: {
          root: {
            borderRadius: 20,
            textTransform: 'none',
            fontWeight: 500,
          },
          contained: {
            boxShadow: '0 2px 4px rgba(0,0,0,0.2)',
            '&:hover': {
              boxShadow: '0 4px 8px rgba(0,0,0,0.3)',
            },
          },
        },
      },
      MuiTextField: {
        styleOverrides: {
          root: {
            '& .MuiOutlinedInput-root': {
              borderRadius: 12,
            },
          },
        },
      },
      MuiSvgIcon: {
        styleOverrides: {
          root: {
            fontSize: kidMode ? '2rem' : undefined,
          },
        },
      },
      MuiListItemButton: {
        styleOverrides: {
          root: {
            borderRadius: 12,
            margin: '4px 8px',
            '&.Mui-selected': {
              backgroundColor: alpha(isDarkMode ? '#bb86fc' : '#6750a4', 0.12),
              '&:hover': {
                backgroundColor: alpha(
                  isDarkMode ? '#bb86fc' : '#6750a4',
                  0.16
                ),
              },
            },
          },
        },
      },
    },
  });

function App() {
  const basePath = getBasePath();
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [status, setStatus] = useState('');
  const [authed, setAuthed] = useState(false);
  const [setupNeeded, setSetupNeeded] = useState(false);
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [drawerPinned, setDrawerPinned] = useState(() => {
    const saved = localStorage.getItem('sidebarPinned');
    return saved ? JSON.parse(saved) : false;
  });
  const [loading, setLoading] = useState(true);
  const [backendAvailable, setBackendAvailable] = useState(false);
  const [apiError, setApiError] = useState(null);
  const [darkMode, setDarkMode] = useState(() => {
    // Check localStorage or system preference for initial theme
    const saved = localStorage.getItem('darkMode');
    if (saved !== null) return JSON.parse(saved);
    return (
      (window.matchMedia &&
        window.matchMedia('(prefers-color-scheme: dark)')?.matches) ||
      false
    );
  });
  const [kidMode, setKidMode] = useState(() => {
    const saved = localStorage.getItem('kidMode');
    return saved ? JSON.parse(saved) : false;
  });

  // Create theme based on current dark mode state
  const theme = createAppTheme(darkMode, kidMode);

  /**
   * Toggle between light and dark mode themes
   * Saves preference to localStorage for persistence
   */
  const toggleDarkMode = () => {
    const newMode = !darkMode;
    setDarkMode(newMode);
    localStorage.setItem('darkMode', JSON.stringify(newMode));
  };

  /**
   * Toggle kid friendly mode which increases font sizes
   * for a simplified child interface.
   */
  const toggleKidMode = () => {
    const newMode = !kidMode;
    setKidMode(newMode);
    localStorage.setItem('kidMode', JSON.stringify(newMode));
  };

  /**
   * Toggle the drawer open/close behavior.
   * If the sidebar is pinned, unpin it instead of toggling.
   */
  const handleDrawerToggle = () => {
    if (drawerPinned) {
      setDrawerPinned(false);
      localStorage.setItem('sidebarPinned', 'false');
    } else {
      setDrawerOpen(!drawerOpen);
    }
  };

  /**
   * Pin or unpin the sidebar and persist state in localStorage.
   */
  const handleDrawerPin = () => {
    const newPinned = !drawerPinned;
    setDrawerPinned(newPinned);
    localStorage.setItem('sidebarPinned', JSON.stringify(newPinned));
    if (newPinned) {
      setDrawerOpen(true);
    }
  };

  const isDrawerOpen = drawerPinned || drawerOpen;

  const navigationItems = [
    {
      id: 'dashboard',
      label: 'Dashboard',
      icon: <DashboardIcon />,
      path: '/dashboard',
    },
    {
      id: 'library',
      label: 'Media Library',
      icon: <LibraryIcon />,
      path: '/library',
    },
    { id: 'wanted', label: 'Wanted', icon: <WantedIcon />, path: '/wanted' },
    {
      id: 'history',
      label: 'History',
      icon: <HistoryIcon />,
      path: '/history',
    },
    {
      id: 'settings',
      label: 'Settings',
      icon: <SettingsIcon />,
      path: '/settings',
    },
    { id: 'system', label: 'System', icon: <SystemIcon />, path: '/system' },
  ];

  const toolsItems = [
    {
      id: 'extract',
      label: 'Extract',
      icon: <ExtractIcon />,
      path: '/tools/extract',
    },
    {
      id: 'convert',
      label: 'Convert',
      icon: <ConvertIcon />,
      path: '/tools/convert',
    },
    {
      id: 'translate',
      label: 'Translate',
      icon: <TranslateIcon />,
      path: '/tools/translate',
    },
    {
      id: 'scheduling',
      label: 'Scheduling',
      icon: <ScheduleIcon />,
      path: '/tools/scheduling',
    },
  ];

  useEffect(() => {
    const checkBackend = async () => {
      try {
        // Check if backend is available
        const isBackendAvailable = await apiService.checkBackendHealth();
        setBackendAvailable(isBackendAvailable);

        if (isBackendAvailable) {
          // Backend is available, check auth and setup status
          setApiError(null);

          try {
            const configResponse = await apiService.get('/api/config');
            if (configResponse.ok) {
              setAuthed(true);
            }
          } catch (error) {
            if (import.meta.env.DEV) {
              console.warn('Config check failed:', error);
            }
          }

          try {
            const setupResponse = await apiService.get('/api/setup/status');
            if (setupResponse.ok) {
              const setupData = await setupResponse.json();
              setSetupNeeded(setupData.needed);
            }
          } catch (error) {
            if (import.meta.env.DEV) {
              console.warn('Setup status check failed:', error);
            }
          }
        } else {
          // Backend not available
          setApiError(
            'Backend service is not available. Some features may be limited.'
          );
          setAuthed(false);
          setSetupNeeded(false);
        }
      } catch (error) {
        if (import.meta.env.DEV) {
          console.error('Backend check failed:', error);
        }
        setBackendAvailable(false);
        setApiError('Failed to connect to backend service.');
      } finally {
        setLoading(false);
      }
    };

    checkBackend();
  }, []);

  const login = async () => {
    const form = new URLSearchParams({ username, password });
    const res = await fetch('/api/login', { method: 'POST', body: form });
    if (res.ok) {
      setStatus('logged in');
      setAuthed(true);
    } else {
      setStatus('login failed');
    }
  };

  // Show loading spinner during initial backend check
  if (loading) {
    return (
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Box
          sx={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'center',
            minHeight: '100vh',
            backgroundColor: 'background.default',
          }}
        >
          <CircularProgress size={60} sx={{ mb: 2 }} />
          <Typography variant="h6" color="text.secondary">
            Connecting to Subtitle Manager...
          </Typography>
        </Box>
      </ThemeProvider>
    );
  }

  // Show offline UI if backend is unavailable
  if (!backendAvailable) {
    return (
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Container component="main" maxWidth="md">
          <Box
            sx={{
              marginTop: 8,
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
              minHeight: '100vh',
              justifyContent: 'center',
            }}
          >
            <Paper
              elevation={0}
              sx={{
                p: 6,
                width: '100%',
                borderRadius: 3,
                border: '1px solid',
                borderColor: 'divider',
                backgroundColor: 'background.paper',
                textAlign: 'center',
              }}
            >
              <SystemIcon sx={{ fontSize: 64, color: 'error.main', mb: 2 }} />
              <Typography
                component="h1"
                variant="h4"
                gutterBottom
                color="error"
              >
                Backend Unavailable
              </Typography>
              <Typography variant="body1" color="text.secondary" sx={{ mb: 3 }}>
                The Subtitle Manager backend service is currently not available.
                Please check that the server is running and try again.
              </Typography>

              {apiError && (
                <Alert severity="error" sx={{ mb: 3, textAlign: 'left' }}>
                  {apiError}
                </Alert>
              )}

              <Box
                sx={{
                  display: 'flex',
                  gap: 2,
                  justifyContent: 'center',
                  flexWrap: 'wrap',
                }}
              >
                <Button
                  variant="contained"
                  onClick={() => window.location.reload()}
                  size="large"
                >
                  Retry Connection
                </Button>
                <Button
                  variant="outlined"
                  onClick={() => (window.location.href = '/offline-info')}
                  size="large"
                >
                  Offline Information
                </Button>
              </Box>
            </Paper>

            {/* Theme toggle for offline page */}
            <Box sx={{ mt: 2 }}>
              <IconButton
                onClick={toggleDarkMode}
                aria-label="toggle dark mode"
                sx={{
                  backgroundColor: 'background.paper',
                  border: '1px solid',
                  borderColor: 'divider',
                }}
              >
                {darkMode ? <LightModeIcon /> : <DarkModeIcon />}
              </IconButton>
              <IconButton
                onClick={toggleKidMode}
                aria-label="toggle kid mode"
                sx={{
                  backgroundColor: 'background.paper',
                  border: '1px solid',
                  borderColor: 'divider',
                  ml: 1,
                }}
              >
                <KidModeIcon />
              </IconButton>
            </Box>
          </Box>
        </Container>
      </ThemeProvider>
    );
  }

  if (!authed) {
    if (setupNeeded) {
      return (
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Setup backendAvailable={backendAvailable} />
        </ThemeProvider>
      );
    }
    return (
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Container component="main" maxWidth="sm">
          <Box
            sx={{
              marginTop: 8,
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
              minHeight: '100vh',
              justifyContent: 'center',
            }}
          >
            <Paper
              elevation={0}
              sx={{
                p: 6,
                width: '100%',
                borderRadius: 3,
                border: '1px solid',
                borderColor: 'divider',
                backgroundColor: 'background.paper',
              }}
            >
              <Box sx={{ textAlign: 'center', mb: 4 }}>
                <Typography
                  component="h1"
                  variant="h4"
                  gutterBottom
                  color="primary"
                >
                  Subtitle Manager
                </Typography>
                <Typography variant="body1" color="text.secondary">
                  Sign in to access your subtitle management dashboard
                </Typography>
              </Box>

              {apiError && (
                <Alert severity="warning" sx={{ mb: 3 }}>
                  {apiError}
                </Alert>
              )}

              <Box component="form" sx={{ mt: 1 }}>
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  id="username"
                  label="Username"
                  name="username"
                  autoComplete="username"
                  autoFocus
                  value={username}
                  onChange={e => setUsername(e.target.value)}
                  variant="outlined"
                  disabled={!backendAvailable}
                />
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="current-password"
                  value={password}
                  onChange={e => setPassword(e.target.value)}
                  variant="outlined"
                  disabled={!backendAvailable}
                />
                <Button
                  type="button"
                  fullWidth
                  variant="contained"
                  size="large"
                  sx={{ mt: 3, mb: 2, py: 1.5 }}
                  onClick={login}
                  disabled={!backendAvailable}
                >
                  Sign In
                </Button>
                {status && (
                  <Paper
                    sx={{
                      p: 2,
                      mt: 2,
                      backgroundColor: status.includes('failed')
                        ? 'error.main'
                        : 'success.main',
                      color: status.includes('failed')
                        ? 'error.contrastText'
                        : 'success.contrastText',
                      borderRadius: 2,
                    }}
                  >
                    <Typography variant="body2" align="center">
                      {status}
                    </Typography>
                  </Paper>
                )}
              </Box>
            </Paper>

            {/* Theme toggle for login page */}
            <Box sx={{ mt: 2 }}>
              <IconButton
                onClick={toggleDarkMode}
                aria-label="toggle dark mode"
                sx={{
                  backgroundColor: 'background.paper',
                  border: '1px solid',
                  borderColor: 'divider',
                }}
              >
                {darkMode ? <LightModeIcon /> : <DarkModeIcon />}
              </IconButton>
              <IconButton
                onClick={toggleKidMode}
                aria-label="toggle kid mode"
                sx={{
                  backgroundColor: 'background.paper',
                  border: '1px solid',
                  borderColor: 'divider',
                  ml: 1,
                }}
              >
                <KidModeIcon />
              </IconButton>
            </Box>
          </Box>
        </Container>
      </ThemeProvider>
    );
  }

  function AppContent() {
    const navigate = useNavigate();
    const location = useLocation();

    const handleBack = () => {
      navigate(-1);
    };

    const showBackButton =
      location.pathname !== '/' && location.pathname !== '/dashboard';

    return (
      <Box sx={{ display: 'flex' }}>
        <AppBar
          position="fixed"
          sx={{
            zIndex: theme => theme.zIndex.drawer + 1,
            transition: theme.transitions.create(['margin'], {
              easing: theme.transitions.easing.sharp,
              duration: theme.transitions.duration.leavingScreen,
            }),
          }}
        >
          <Toolbar>
            <IconButton
              color="inherit"
              aria-label="open drawer"
              edge="start"
              onClick={handleDrawerToggle}
              sx={{ mr: 2 }}
            >
              <MenuIcon />
            </IconButton>
            {showBackButton && (
              <IconButton color="inherit" onClick={handleBack} sx={{ mr: 2 }}>
                <BackIcon />
              </IconButton>
            )}
            <Typography
              variant="h6"
              noWrap
              component="div"
              sx={{ flexGrow: 1 }}
            >
              Subtitle Manager
            </Typography>
            <IconButton
              color="inherit"
              onClick={toggleDarkMode}
              aria-label="toggle dark mode"
            >
              {darkMode ? <LightModeIcon /> : <DarkModeIcon />}
            </IconButton>
            <IconButton
              color="inherit"
              onClick={toggleKidMode}
              aria-label="toggle kid mode"
            >
              <KidModeIcon />
            </IconButton>
          </Toolbar>
        </AppBar>

        <Drawer
          variant={drawerPinned ? 'permanent' : 'temporary'}
          anchor="left"
          open={isDrawerOpen}
          onClose={() => !drawerPinned && setDrawerOpen(false)}
          sx={{
            width: 280,
            flexShrink: 0,
            '& .MuiDrawer-paper': {
              width: 280,
              boxSizing: 'border-box',
            },
          }}
        >
          <Toolbar />
          <Box sx={{ overflow: 'auto', py: 1 }}>
            <Box sx={{ px: 2, pb: 1 }}>
              <Button
                fullWidth
                variant="outlined"
                size="small"
                onClick={handleDrawerPin}
                startIcon={<PinIcon />}
              >
                {drawerPinned ? 'Unpin Sidebar' : 'Pin Sidebar'}
              </Button>
            </Box>
            <List>
              {navigationItems.map(item => (
                <ListItem key={item.id} disablePadding>
                  <NavLink
                    to={item.path}
                    style={{
                      textDecoration: 'none',
                      color: 'inherit',
                      width: '100%',
                    }}
                    end={item.path !== '/dashboard'}
                  >
                    {({ isActive }) => (
                      <ListItemButton
                        onClick={() => setDrawerOpen(false)}
                        sx={{
                          ...(isActive && {
                            backgroundColor: theme =>
                              theme.palette.mode === 'dark'
                                ? alpha(theme.palette.primary.main, 0.12)
                                : alpha(theme.palette.primary.main, 0.08),
                            color: 'primary.main',
                            '& .MuiListItemIcon-root': {
                              color: 'primary.main',
                            },
                          }),
                        }}
                      >
                        <ListItemIcon sx={{ color: 'inherit' }}>
                          {item.icon}
                        </ListItemIcon>
                        <ListItemText primary={item.label} />
                      </ListItemButton>
                    )}
                  </NavLink>
                </ListItem>
              ))}

              <Divider sx={{ my: 1 }} />
              <ListSubheader component="div" sx={{ px: 2, py: 1 }}>
                Tools
              </ListSubheader>

              {toolsItems.map(item => (
                <ListItem key={item.id} disablePadding>
                  <ListItemButton
                    component={NavLink}
                    to={item.path}
                    onClick={() => {
                      setDrawerOpen(false);
                    }}
                    sx={{
                      pl: 4,
                      '&.active': {
                        backgroundColor: 'action.selected',
                        color: 'primary.main',
                        '& .MuiListItemIcon-root': {
                          color: 'primary.main',
                        },
                      },
                    }}
                  >
                    <ListItemIcon sx={{ color: 'inherit', minWidth: 36 }}>
                      {item.icon}
                    </ListItemIcon>
                    <ListItemText primary={item.label} />
                  </ListItemButton>
                </ListItem>
              ))}
            </List>
          </Box>
        </Drawer>

        <Box
          component="main"
          sx={{
            flexGrow: 1,
            p: 3,
            transition: theme.transitions.create('margin', {
              easing: theme.transitions.easing.sharp,
              duration: theme.transitions.duration.leavingScreen,
            }),
            marginLeft: 0,
            backgroundColor: 'background.default',
            minHeight: '100vh',
          }}
        >
          <Toolbar />

          {apiError && (
            <Alert severity="warning" sx={{ mb: 3 }}>
              {apiError}
            </Alert>
          )}

          <Suspense fallback={<LoadingComponent message="Loading page..." />}>
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
              <Route path="/details" element={<MediaDetails />} />
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
              <Route
                path="/tools/scheduling"
                element={<Scheduling backendAvailable={backendAvailable} />}
              />
              <Route path="/offline-info" element={<OfflineInfo />} />
            </Routes>
          </Suspense>
        </Box>

        {/* Floating action button for quick navigation on mobile */}
        {!isDrawerOpen && (
          <Fab
            color="primary"
            aria-label="menu"
            onClick={() => setDrawerOpen(true)}
            sx={{
              position: 'fixed',
              bottom: 16,
              right: 16,
              display: { xs: 'flex', md: 'none' },
            }}
          >
            <MenuIcon />
          </Fab>
        )}
      </Box>
    );
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router basename={basePath}>
        <AppContent />
      </Router>
    </ThemeProvider>
  );
}

export default App;
