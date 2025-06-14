import {
  Transform as ConvertIcon,
  Dashboard as DashboardIcon,
  Archive as ExtractIcon,
  History as HistoryIcon,
  Menu as MenuIcon,
  Settings as SettingsIcon,
  BugReport as SystemIcon,
  Translate as TranslateIcon,
  Download as WantedIcon,
} from "@mui/icons-material";
import {
  AppBar,
  Box,
  Button,
  Container,
  CssBaseline,
  Drawer,
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
  createTheme
} from "@mui/material";
import { useEffect, useState } from "react";
import "./App.css";
import Convert from "./Convert.jsx";
import Dashboard from "./Dashboard.jsx";
import Extract from "./Extract.jsx";
import History from "./History.jsx";
import Settings from "./Settings.jsx";
import Setup from "./Setup.jsx";
import System from "./System.jsx";
import Translate from "./Translate.jsx";
import Wanted from "./Wanted.jsx";

const theme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#90caf9',
    },
    secondary: {
      main: '#f48fb1',
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

  const navigationItems = [
    { id: "dashboard", label: "Dashboard", icon: <DashboardIcon /> },
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
        <Container component="main" maxWidth="xs">
          <Box
            sx={{
              marginTop: 8,
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
            }}
          >
            <Paper elevation={3} sx={{ p: 4, width: '100%' }}>
              <Typography component="h1" variant="h4" align="center" gutterBottom>
                Subtitle Manager
              </Typography>
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
                />
                <Button
                  type="button"
                  fullWidth
                  variant="contained"
                  sx={{ mt: 3, mb: 2 }}
                  onClick={login}
                >
                  Sign In
                </Button>
                {status && (
                  <Typography variant="body2" color="error" align="center">
                    {status}
                  </Typography>
                )}
              </Box>
            </Paper>
          </Box>
        </Container>
      </ThemeProvider>
    );
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Box sx={{ display: 'flex' }}>
        <AppBar position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
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
            <Typography variant="h6" noWrap component="div">
              Subtitle Manager
            </Typography>
          </Toolbar>
        </AppBar>
        <Drawer
          variant="persistent"
          anchor="left"
          open={drawerOpen}
          sx={{
            width: 240,
            flexShrink: 0,
            '& .MuiDrawer-paper': {
              width: 240,
              boxSizing: 'border-box',
            },
          }}
        >
          <Toolbar />
          <Box sx={{ overflow: 'auto' }}>
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
                    <ListItemIcon>{item.icon}</ListItemIcon>
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
            transition: (theme) => theme.transitions.create('margin', {
              easing: theme.transitions.easing.sharp,
              duration: theme.transitions.duration.leavingScreen,
            }),
            marginLeft: drawerOpen ? '240px' : 0,
          }}
        >
          <Toolbar />
          {page === "settings" ? (
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
      </Box>
    </ThemeProvider>
  );
}

export default App;
