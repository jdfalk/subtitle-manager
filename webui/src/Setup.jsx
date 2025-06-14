import {
  CheckCircleOutlined,
  CloudDownloadOutlined,
  ConnectedTvOutlined,
  ExtensionOutlined,
  PersonAddOutlined,
  PhoneAndroidOutlined,
  SettingsOutlined,
  SubtitlesOutlined,
  TranslateOutlined,
} from "@mui/icons-material";
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  Checkbox,
  CircularProgress,
  Container,
  Divider,
  FormControlLabel,
  FormGroup,
  Grid,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Paper,
  Stack,
  Step,
  StepLabel,
  Stepper,
  TextField,
  Typography
} from "@mui/material";
import CssBaseline from "@mui/material/CssBaseline";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { useState } from "react";

// Material Design 3 theme
const theme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#6750a4',
    },
    secondary: {
      main: '#625b71',
    },
    background: {
      default: '#fffbfe',
      paper: '#ffffff',
    },
  },
  shape: {
    borderRadius: 16,
  },
  typography: {
    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
    h3: {
      fontWeight: 600,
    },
    h4: {
      fontWeight: 500,
    },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: 24,
          textTransform: 'none',
          fontSize: '0.9rem',
          fontWeight: 500,
          padding: '10px 24px',
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          borderRadius: 24,
          boxShadow: '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)',
        },
      },
    },
    MuiTextField: {
      styleOverrides: {
        root: {
          '& .MuiOutlinedInput-root': {
            borderRadius: 16,
          },
        },
      },
    },
  },
});

const steps = ['Welcome', 'Import from Bazarr', 'Create Admin', 'Server Settings'];

/**
 * Setup guides the user through initial configuration when no user exists.
 * Multi-step wizard with Material Design 3: Welcome -> Bazarr Import (optional) -> Admin User -> Server Settings -> Complete
 */
export default function Setup() {
  const [step, setStep] = useState(0);
  const [serverName, setServerName] = useState("Subtitle Manager");
  const [reverseProxy, setReverseProxy] = useState(false);
  const [adminUser, setAdminUser] = useState("");
  const [adminPass, setAdminPass] = useState("");

  // Bazarr import state
  const [bazarrURL, setBazarrURL] = useState("");
  const [bazarrAPIKey, setBazarrAPIKey] = useState("");
  const [bazarrSettings, setBazarrSettings] = useState(null);
  const [bazarrRawData, setBazarrRawData] = useState(null);
  const [selectedSettings, setSelectedSettings] = useState({});
  const [bazarrLoading, setBazarrLoading] = useState(false);
  const [bazarrError, setBazarrError] = useState("");
  const [showDebug, setShowDebug] = useState(false);

  const [status, setStatus] = useState("");
  const [loading, setLoading] = useState(false);

  const importBazarr = async () => {
    if (!bazarrURL || !bazarrAPIKey) {
      setBazarrError("Please enter both URL and API key");
      return;
    }

    setBazarrLoading(true);
    setBazarrError("");

    try {
      const res = await fetch("/api/setup/bazarr", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          url: bazarrURL,
          api_key: bazarrAPIKey
        }),
      });

      if (res.ok) {
        const data = await res.json();
        setBazarrRawData(data.raw_settings); // Store raw data for debugging
        setBazarrSettings(data.preview);
        // Pre-select all settings for import
        const selected = {};
        Object.keys(data.preview).forEach(key => {
          selected[key] = true;
        });
        setSelectedSettings(selected);
        console.log("Raw Bazarr data:", data.raw_settings);
        console.log("Mapped Bazarr data:", data.preview);
      } else {
        const errorText = await res.text();
        setBazarrError(errorText || "Failed to connect to Bazarr");
      }
    } catch (err) {
      setBazarrError("Network error: " + err.message);
    } finally {
      setBazarrLoading(false);
    }
  };

  const submit = async () => {
    setLoading(true);

    // Build the final configuration including Bazarr imports
    const integrations = {};
    const importedConfig = {};

    if (bazarrSettings) {
      Object.keys(selectedSettings).forEach(key => {
        if (selectedSettings[key]) {
          importedConfig[key] = bazarrSettings[key];
        }
      });
    }

    const body = {
      server_name: serverName,
      reverse_proxy: reverseProxy,
      admin_user: adminUser,
      admin_pass: adminPass,
      integrations,
      ...importedConfig // Merge Bazarr settings
    };

    try {
      const res = await fetch("/api/setup", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });

      if (res.ok) {
        setStatus("complete");
        setTimeout(() => window.location.reload(), 2000);
      } else {
        setStatus("error");
      }
    } catch (err) {
      console.error("Setup failed:", err);
      setStatus("error");
    } finally {
      setLoading(false);
    }
  };

  const next = () => setStep(step + 1);
  const prev = () => setStep(step - 1);
  const skip = () => {
    setBazarrSettings(null);
    setSelectedSettings({});
    next();
  };

  if (status === "complete") {
    return (
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Container maxWidth="sm" sx={{
          minHeight: '100vh',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          py: 4
        }}>
          <Card elevation={3} sx={{ width: '100%', textAlign: 'center', p: 4 }}>
            <CheckCircleOutlined color="success" sx={{ fontSize: 80, mb: 2 }} />
            <Typography variant="h3" color="success.main" gutterBottom>
              Setup Complete!
            </Typography>
            <Typography variant="body1" color="text.secondary" paragraph>
              Your Subtitle Manager is ready to use.
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Redirecting to login...
            </Typography>
            <Box sx={{ mt: 3 }}>
              <CircularProgress />
            </Box>
          </Card>
        </Container>
      </ThemeProvider>
    );
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Box sx={{
        minHeight: '100vh',
        background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        p: 2
      }}>
        <Container maxWidth="md">
          <Card elevation={8} sx={{ overflow: 'visible' }}>
            {/* Header with progress */}
            <Box sx={{
              background: 'linear-gradient(135deg, #6750a4 0%, #8b5cf6 100%)',
              color: 'white',
              p: 4,
              textAlign: 'center'
            }}>
              <Box sx={{ display: 'flex', justifyContent: 'center', mb: 3 }}>
                <Paper sx={{
                  width: 80,
                  height: 80,
                  borderRadius: 4,
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  background: 'rgba(255,255,255,0.2)',
                  backdropFilter: 'blur(10px)'
                }}>
                  <SubtitlesOutlined sx={{ fontSize: 40, color: 'white' }} />
                </Paper>
              </Box>

              <Typography variant="h4" gutterBottom>
                Subtitle Manager Setup
              </Typography>

              <Box sx={{ mt: 3 }}>
                <Stepper activeStep={step} alternativeLabel sx={{
                  '& .MuiStepLabel-label': { color: 'rgba(255,255,255,0.8)' },
                  '& .MuiStepLabel-label.Mui-active': { color: 'white' },
                  '& .MuiStepLabel-label.Mui-completed': { color: 'white' },
                  '& .MuiStepIcon-root': { color: 'rgba(255,255,255,0.5)' },
                  '& .MuiStepIcon-root.Mui-active': { color: 'white' },
                  '& .MuiStepIcon-root.Mui-completed': { color: 'white' },
                }}>
                  {steps.map((label) => (
                    <Step key={label}>
                      <StepLabel>{label}</StepLabel>
                    </Step>
                  ))}
                </Stepper>
              </Box>
            </Box>

            <CardContent sx={{ p: 5 }}>
              {step === 0 && (
                <Box>
                  <Typography variant="h3" gutterBottom>
                    Welcome to Subtitle Manager
                  </Typography>
                  <Typography variant="body1" color="text.secondary" paragraph>
                    A powerful, self-hosted subtitle management system that automates
                    subtitle downloading, translation, and organization for your media library.
                  </Typography>

                  <Grid container spacing={3} sx={{ my: 4 }}>
                    <Grid item xs={12} sm={6}>
                      <Card variant="outlined" sx={{ p: 3, height: '100%', textAlign: 'center' }}>
                        <CloudDownloadOutlined color="primary" sx={{ fontSize: 48, mb: 2 }} />
                        <Typography variant="h6" gutterBottom>
                          Automated Downloads
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          Automatically find and download subtitles for your movies and TV shows
                        </Typography>
                      </Card>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                      <Card variant="outlined" sx={{ p: 3, height: '100%', textAlign: 'center' }}>
                        <TranslateOutlined color="primary" sx={{ fontSize: 48, mb: 2 }} />
                        <Typography variant="h6" gutterBottom>
                          Translation
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          Translate subtitles to any language using Google Translate or OpenAI
                        </Typography>
                      </Card>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                      <Card variant="outlined" sx={{ p: 3, height: '100%', textAlign: 'center' }}>
                        <ExtensionOutlined color="primary" sx={{ fontSize: 48, mb: 2 }} />
                        <Typography variant="h6" gutterBottom>
                          Integrations
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          Works seamlessly with Sonarr, Radarr, Plex, and other media tools
                        </Typography>
                      </Card>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                      <Card variant="outlined" sx={{ p: 3, height: '100%', textAlign: 'center' }}>
                        <PhoneAndroidOutlined color="primary" sx={{ fontSize: 48, mb: 2 }} />
                        <Typography variant="h6" gutterBottom>
                          Modern UI
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          Clean, responsive web interface accessible from any device
                        </Typography>
                      </Card>
                    </Grid>
                  </Grid>

                  <Box sx={{ display: 'flex', justifyContent: 'flex-end', mt: 4 }}>
                    <Button variant="contained" size="large" onClick={next}>
                      Get Started
                    </Button>
                  </Box>
                </Box>
              )}

              {step === 1 && (
                <Box>
                  <Typography variant="h4" gutterBottom>
                    Import from Bazarr
                  </Typography>
                  <Typography variant="body1" color="text.secondary" paragraph>
                    Already using Bazarr? Import your existing configuration to get started quickly.
                  </Typography>

                  <Paper variant="outlined" sx={{ p: 3, mb: 3 }}>
                    <Stack spacing={3}>
                      <TextField
                        fullWidth
                        label="Bazarr URL"
                        placeholder="http://localhost:6767"
                        value={bazarrURL}
                        onChange={(e) => setBazarrURL(e.target.value)}
                        InputProps={{
                          startAdornment: <ConnectedTvOutlined sx={{ mr: 1, color: 'text.secondary' }} />,
                        }}
                      />

                      <TextField
                        fullWidth
                        label="API Key"
                        type="password"
                        placeholder="Your Bazarr API key"
                        value={bazarrAPIKey}
                        onChange={(e) => setBazarrAPIKey(e.target.value)}
                      />

                      {bazarrError && (
                        <Alert severity="error">{bazarrError}</Alert>
                      )}

                      <Stack direction="row" spacing={2} sx={{ alignSelf: 'flex-start' }}>
                        <Button
                          variant="outlined"
                          onClick={importBazarr}
                          disabled={bazarrLoading}
                          startIcon={bazarrLoading ? <CircularProgress size={20} /> : <CloudDownloadOutlined />}
                        >
                          {bazarrLoading ? "Connecting..." : "Connect to Bazarr"}
                        </Button>

                        <Button
                          variant="text"
                          onClick={() => setShowDebug(!showDebug)}
                          size="small"
                        >
                          {showDebug ? "Hide Debug" : "Show Debug"}
                        </Button>
                      </Stack>

                      {showDebug && bazarrRawData && (
                        <Paper variant="outlined" sx={{ p: 2, bgcolor: 'grey.50' }}>
                          <Typography variant="subtitle2" gutterBottom>
                            Raw Bazarr API Response:
                          </Typography>
                          <pre style={{ fontSize: '12px', overflow: 'auto', maxHeight: '200px' }}>
                            {JSON.stringify(bazarrRawData, null, 2)}
                          </pre>
                        </Paper>
                      )}

                      {bazarrSettings && (
                        <Box>
                          <Typography variant="h6" gutterBottom>
                            Found Settings
                          </Typography>
                          <Typography variant="body2" color="text.secondary" paragraph>
                            Select which settings to import:
                          </Typography>
                          <Paper variant="outlined" sx={{ maxHeight: 300, overflow: 'auto' }}>
                            <List>
                              {Object.entries(bazarrSettings).map(([key, value], index) => (
                                <div key={key}>
                                  <ListItem>
                                    <ListItemIcon>
                                      <Checkbox
                                        checked={selectedSettings[key] || false}
                                        onChange={(e) => setSelectedSettings({
                                          ...selectedSettings,
                                          [key]: e.target.checked
                                        })}
                                      />
                                    </ListItemIcon>
                                    <ListItemText
                                      primary={
                                        <Typography variant="body2" fontWeight="medium">
                                          {key}
                                        </Typography>
                                      }
                                      secondary={
                                        <Typography variant="caption" sx={{ fontFamily: 'monospace' }}>
                                          {typeof value === 'object' ? JSON.stringify(value) : String(value)}
                                        </Typography>
                                      }
                                    />
                                  </ListItem>
                                  {index < Object.keys(bazarrSettings).length - 1 && <Divider />}
                                </div>
                              ))}
                            </List>
                          </Paper>
                        </Box>
                      )}
                    </Stack>
                  </Paper>

                  <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 4 }}>
                    <Button variant="outlined" onClick={prev}>
                      Back
                    </Button>
                    <Stack direction="row" spacing={2}>
                      <Button variant="text" onClick={skip}>
                        Skip
                      </Button>
                      <Button
                        variant="contained"
                        onClick={next}
                        disabled={!bazarrSettings}
                      >
                        Continue
                      </Button>
                    </Stack>
                  </Box>
                </Box>
              )}

              {step === 2 && (
                <Box>
                  <Typography variant="h4" gutterBottom>
                    Create Admin Account
                  </Typography>
                  <Typography variant="body1" color="text.secondary" paragraph>
                    Create your administrator account to manage Subtitle Manager.
                  </Typography>

                  <Paper variant="outlined" sx={{ p: 3, mb: 3 }}>
                    <Stack spacing={3}>
                      <TextField
                        fullWidth
                        label="Username"
                        placeholder="admin"
                        value={adminUser}
                        onChange={(e) => setAdminUser(e.target.value)}
                        required
                        InputProps={{
                          startAdornment: <PersonAddOutlined sx={{ mr: 1, color: 'text.secondary' }} />,
                        }}
                      />

                      <TextField
                        fullWidth
                        label="Password"
                        type="password"
                        placeholder="Choose a secure password"
                        value={adminPass}
                        onChange={(e) => setAdminPass(e.target.value)}
                        required
                      />
                    </Stack>
                  </Paper>

                  <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 4 }}>
                    <Button variant="outlined" onClick={prev}>
                      Back
                    </Button>
                    <Button
                      variant="contained"
                      onClick={next}
                      disabled={!adminUser || !adminPass}
                    >
                      Continue
                    </Button>
                  </Box>
                </Box>
              )}

              {step === 3 && (
                <Box>
                  <Typography variant="h4" gutterBottom>
                    Server Configuration
                  </Typography>
                  <Typography variant="body1" color="text.secondary" paragraph>
                    Configure basic server settings for your deployment.
                  </Typography>

                  <Paper variant="outlined" sx={{ p: 3, mb: 3 }}>
                    <Stack spacing={3}>
                      <TextField
                        fullWidth
                        label="Server Name"
                        placeholder="Subtitle Manager"
                        value={serverName}
                        onChange={(e) => setServerName(e.target.value)}
                        InputProps={{
                          startAdornment: <SettingsOutlined sx={{ mr: 1, color: 'text.secondary' }} />,
                        }}
                      />

                      <FormGroup>
                        <FormControlLabel
                          control={
                            <Checkbox
                              checked={reverseProxy}
                              onChange={(e) => setReverseProxy(e.target.checked)}
                            />
                          }
                          label="Running behind a reverse proxy"
                        />
                        <Typography variant="caption" color="text.secondary" sx={{ ml: 4 }}>
                          Enable this if you're using nginx, Apache, or another reverse proxy
                        </Typography>
                      </FormGroup>
                    </Stack>
                  </Paper>

                  {status === "error" && (
                    <Alert severity="error" sx={{ mb: 3 }}>
                      Setup failed. Please check your settings and try again.
                    </Alert>
                  )}

                  <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 4 }}>
                    <Button variant="outlined" onClick={prev}>
                      Back
                    </Button>
                    <Button
                      variant="contained"
                      size="large"
                      onClick={submit}
                      disabled={loading}
                      startIcon={loading ? <CircularProgress size={20} /> : <CheckCircleOutlined />}
                    >
                      {loading ? "Setting up..." : "Complete Setup"}
                    </Button>
                  </Box>
                </Box>
              )}
            </CardContent>
          </Card>
        </Container>
      </Box>
    </ThemeProvider>
  );
}

