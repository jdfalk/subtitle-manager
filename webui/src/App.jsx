import {
  Transform as ConvertIcon,
  Brightness4 as DarkModeIcon,
  Dashboard as DashboardIcon,
  Archive as ExtractIcon,
  History as HistoryIcon,
  VideoLibrary as LibraryIcon,
  Brightness7 as LightModeIcon,
  Menu as MenuIcon,
  Settings as SettingsIcon,
  BugReport as SystemIcon,
  Translate as TranslateIcon,
  Download as WantedIcon,
} from "@mui/icons-material";
import {
  alpha,
  AppBar,
  Box,
  Button,
  Container,
  createTheme,
  CssBaseline,
  Drawer,
  Fab,
  IconButton,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Paper,
  TextField,
  ThemeProvider,
  Toolbar,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";
import "./App.css";
import Convert from "./Convert.jsx";
import Dashboard from "./Dashboard.jsx";
import Extract from "./Extract.jsx";
import History from "./History.jsx";
import MediaLibrary from "./MediaLibrary.jsx";
import Settings from "./Settings.jsx";
import Setup from "./Setup.jsx";
import System from "./System.jsx";
import Translate from "./Translate.jsx";
import Wanted from "./Wanted.jsx";

/**
 * Creates a Material Design 3 compliant theme with enhanced dark mode support
 * Following Google's Material Design guidelines for color schemes and accessibility
 * @param {boolean} isDarkMode - Whether to use dark mode
 * @returns {Theme} Material UI theme object
 */
const createAppTheme = (isDarkMode = true) => createTheme({
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
    MuiListItemButton: {
      styleOverrides: {
        root: {
          borderRadius: 12,
          margin: '4px 8px',
          '&.Mui-selected': {
            backgroundColor: alpha(isDarkMode ? '#bb86fc' : '#6750a4', 0.12),
            '&:hover': {
              backgroundColor: alpha(isDarkMode ? '#bb86fc' : '#6750a4', 0.16),
            },
          },
        },
      },
    },
  },
});

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [status, setStatus] = useState("");
  const [authed, setAuthed] = useState(false);
  const [setupNeeded, setSetupNeeded] = useState(false);
  const [page, setPage] = useState("dashboard");
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [darkMode, setDarkMode] = useState(() => {
    // Check localStorage or system preference for initial theme
    const saved = localStorage.getItem('darkMode');
    if (saved !== null) return JSON.parse(saved);
    return window.matchMedia('(prefers-color-scheme: dark)').matches;
  });

  // Create theme based on current dark mode state
  const theme = createAppTheme(darkMode);

  /**
   * Toggle between light and dark mode themes
   * Saves preference to localStorage for persistence
   */
  const toggleDarkMode = () => {
    const newMode = !darkMode;
    setDarkMode(newMode);
    localStorage.setItem('darkMode', JSON.stringify(newMode));
  };

  const navigationItems = [
    { id: "dashboard", label: "Dashboard", icon: <DashboardIcon /> },
    { id: "library", label: "Media Library", icon: <LibraryIcon /> },
    { id: "settings", label: "Settings", icon: <SettingsIcon /> },
    { id: "extract", label: "Extract", icon: <ExtractIcon /> },
    { id: "history", label: "History", icon: <HistoryIcon /> },
    { id: "convert", label: "Convert", icon: <ConvertIcon /> },
    { id: "translate", label: "Translate", icon: <TranslateIcon /> },
    { id: "system", label: "System", icon: <SystemIcon /> },
    { id: "wanted", label: "Wanted", icon: <WantedIcon /> },
  ];

  useEffect(() => {
    fetch("/api/config").then((res) => {
      if (res.ok) setAuthed(true);
    });
    fetch("/api/setup/status")
      .then((r) => r.json())
      .then((d) => setSetupNeeded(d.needed));
  }, []);

  const login = async () => {
    const form = new URLSearchParams({ username, password });
    const res = await fetch("/api/login", { method: "POST", body: form });
    if (res.ok) {
      setStatus("logged in");
      setAuthed(true);
    } else {
      setStatus("login failed");
    }
  };

  if (!authed) {
    if (setupNeeded) {
      return (
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Setup />
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
                <Typography component="h1" variant="h4" gutterBottom color="primary">
                  Subtitle Manager
                </Typography>
                <Typography variant="body1" color="text.secondary">
                  Sign in to access your subtitle management dashboard
                </Typography>
              </Box>

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
                  onChange={(e) => setUsername(e.target.value)}
                  variant="outlined"
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
                  onChange={(e) => setPassword(e.target.value)}
                  variant="outlined"
                />
                <Button
                  type="button"
                  fullWidth
                  variant="contained"
                  size="large"
                  sx={{ mt: 3, mb: 2, py: 1.5 }}
                  onClick={login}
                >
                  Sign In
                </Button>
                {status && (
                  <Paper
                    sx={{
                      p: 2,
                      mt: 2,
                      backgroundColor: status.includes('failed') ? 'error.main' : 'success.main',
                      color: status.includes('failed') ? 'error.contrastText' : 'success.contrastText',
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
            </Box>
          </Box>
        </Container>
      </ThemeProvider>
    );
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Box sx={{ display: 'flex' }}>
        <AppBar
          position="fixed"
          sx={{
            zIndex: (theme) => theme.zIndex.drawer + 1,
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
              onClick={() => setDrawerOpen(!drawerOpen)}
              sx={{ mr: 2 }}
            >
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" noWrap component="div" sx={{ flexGrow: 1 }}>
              Subtitle Manager
            </Typography>
            <IconButton
              color="inherit"
              onClick={toggleDarkMode}
              aria-label="toggle dark mode"
            >
              {darkMode ? <LightModeIcon /> : <DarkModeIcon />}
            </IconButton>
          </Toolbar>
        </AppBar>

        <Drawer
          variant="persistent"
          anchor="left"
          open={drawerOpen}
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
            <List>
              {navigationItems.map((item) => (
                <ListItem key={item.id} disablePadding>
                  <ListItemButton
                    selected={page === item.id}
                    onClick={() => {
                      setPage(item.id);
                      setDrawerOpen(false);
                    }}
                  >
                    <ListItemIcon sx={{ color: 'inherit' }}>
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
            marginLeft: drawerOpen ? '280px' : 0,
            backgroundColor: 'background.default',
            minHeight: '100vh',
          }}
        >
          <Toolbar />
          {page === "library" ? (
            <MediaLibrary />
          ) : page === "settings" ? (
            <Settings />
          ) : page === "extract" ? (
            <Extract />
          ) : page === "history" ? (
            <History />
          ) : page === "convert" ? (
            <Convert />
          ) : page === "translate" ? (
            <Translate />
          ) : page === "system" ? (
            <System />
          ) : page === "wanted" ? (
            <Wanted />
          ) : (
            <Dashboard />
          )}
        </Box>

        {/* Floating action button for quick navigation on mobile */}
        {!drawerOpen && (
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
    </ThemeProvider>
  );
}

export default App;
