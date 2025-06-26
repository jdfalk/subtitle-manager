// file: webui/src/Dashboard.jsx
import {
  Folder as FolderIcon,
  PlayArrow as PlayIcon,
} from '@mui/icons-material';
import {
  Alert,
  Autocomplete,
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  CircularProgress,
  FormControl,
  Grid,
  IconButton,
  InputAdornment,
  InputLabel,
  LinearProgress,
  List,
  ListItem,
  ListItemText,
  MenuItem,
  Paper,
  Select,
  Snackbar,
  TextField,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import DirectoryChooser from './components/DirectoryChooser.jsx';
import QuickLinks from './components/QuickLinks.jsx';
import { apiService } from './services/api.js';

/**
 * Dashboard component for managing subtitle scanning operations.
 * Provides controls for starting scans and monitoring progress.
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend service is available
 */

export default function Dashboard({ backendAvailable = true }) {
  const [status, setStatus] = useState({
    running: false,
    completed: 0,
    files: [],
  });
  const [dir, setDir] = useState('');
  const [lang, setLang] = useState('en');
  // Provider is selected from available options
  const [provider, setProvider] = useState('');
  const [availableProviders, setAvailableProviders] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);
  const [suggestions, setSuggestions] = useState([]);
  const [chooserOpen, setChooserOpen] = useState(false);
  const [systemInfo, setSystemInfo] = useState(null);
  const [configSnackOpen, setConfigSnackOpen] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    if (backendAvailable) {
      poll();
      loadProviders();
      loadSystemInfo();
    }
  }, [backendAvailable]); // eslint-disable-line react-hooks/exhaustive-deps

  const loadProviders = async () => {
    if (!backendAvailable) return;

    try {
      setLoading(true);
      setError(null);
      const response = await apiService.get('/api/providers');
      if (response.ok) {
        const data =
          typeof response.json === 'function' ? await response.json() : [];
        // Only show enabled providers
        const enabledProviders = (data || []).filter(p => p.enabled);
        setAvailableProviders(enabledProviders);

        // Set default provider to first enabled one
        if (enabledProviders.length > 0 && !provider) {
          setProvider(enabledProviders[0].name);
        }
      }
    } catch (error) {
      console.error('Failed to load providers:', error);
      setError('Failed to load providers');
      // Fallback to hardcoded providers if API fails
      setAvailableProviders([
        { name: 'opensubtitles', displayName: 'OpenSubtitles', enabled: true },
        { name: 'addic7ed', displayName: 'Addic7ed', enabled: true },
        { name: 'subscene', displayName: 'Subscene', enabled: true },
        { name: 'podnapisi', displayName: 'Podnapisi', enabled: true },
      ]);
    } finally {
      setLoading(false);
    }
  };

  const loadSystemInfo = async () => {
    if (!backendAvailable) return;
    try {
      const resp = await apiService.get('/api/system');
      if (resp.ok) {
        const data = typeof resp.json === 'function' ? await resp.json() : {};
        setSystemInfo(data);
      }
    } catch (err) {
      console.error('Failed to load system info:', err);
    }
  };

  const formatProviderName = name => {
    const specialNames = {
      opensubtitles: 'OpenSubtitles',
      opensubtitlescom: 'OpenSubtitles.com',
      opensubtitlesvip: 'OpenSubtitles VIP',
      addic7ed: 'Addic7ed',
      podnapisi: 'Podnapisi.NET',
      subscene: 'Subscene',
      yifysubtitles: 'YIFY Subtitles',
      turkcealtyazi: 'Türkçe Altyazı',
      greeksubtitles: 'Greek Subtitles',
      legendasdivx: 'Legendas DivX',
      legendasnet: 'Legendas.NET',
      napiprojekt: 'NapiProjekt',
    };

    return (
      specialNames[name] ||
      name
        .split(/(?=[A-Z])/)
        .map(word => word.charAt(0).toUpperCase() + word.slice(1))
        .join(' ')
    );
  };

  const fetchSuggestions = async prefix => {
    if (!backendAvailable) return;
    try {
      const lastSlash = prefix.lastIndexOf('/');
      const parent = lastSlash > 0 ? prefix.slice(0, lastSlash) || '/' : '/';
      const resp = await apiService.get(
        `/api/library/browse?path=${encodeURIComponent(parent)}`
      );
      if (resp.ok) {
        const data =
          typeof resp.json === 'function' ? await resp.json() : { items: [] };
        const dirs = (data.items || [])
          .filter(item => item.isDirectory)
          .map(item => item.path)
          .filter(p => p.startsWith(prefix));
        setSuggestions(dirs);
      } else {
        setSuggestions([]);
      }
    } catch {
      setSuggestions([]);
    }
  };

  const poll = async () => {
    if (!backendAvailable) return;

    try {
      const response = await apiService.get('/api/scan/status');
      if (response.ok) {
        const data =
          typeof response.json === 'function' ? await response.json() : {};
        // Ensure files is always an array to prevent null reference errors
        setStatus({
          running: data.running || false,
          completed: data.completed || 0,
          files: data.files || [],
        });
        if (data.running) {
          setTimeout(poll, 1000);
        }
      }
    } catch (error) {
      console.error('Failed to poll scan status:', error);
      setError('Failed to get scan status');
    }
  };

  const start = async () => {
    if (!backendAvailable) {
      setError('Backend service is not available');
      return;
    }

    try {
      setLoading(true);
      setError(null);
      const body = { provider, directory: dir, lang };
      const response = await apiService.post('/api/scan', body);
      if (response.ok) {
        poll();
      } else {
        setError('Failed to start scan');
      }
    } catch (error) {
      console.error('Failed to start scan:', error);
      setError('Failed to start scan');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    poll();
    loadSystemInfo();
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        Dashboard
      </Typography>

      {/* Backend availability warning */}
      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available. Subtitle scanning and management
          features are currently disabled.
        </Alert>
      )}

      {/* Error display */}
      {error && (
        <Alert severity="error" sx={{ mb: 3 }} onClose={() => setError(null)}>
          {error}
        </Alert>
      )}

      <QuickLinks />

      <Grid container spacing={3}>
        {/* Scan Controls */}
        <Grid size={{ xs: 12, md: 8 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Subtitle Scan
              </Typography>
              <Box component="form" sx={{ '& > :not(style)': { m: 1 } }}>
                <Autocomplete
                  freeSolo
                  options={suggestions}
                  inputValue={dir}
                  onInputChange={(e, value) => {
                    setDir(value);
                    fetchSuggestions(value);
                  }}
                  renderInput={params => (
                    <TextField
                      {...params}
                      fullWidth
                      label="Directory Path"
                      placeholder="Enter directory to scan"
                      disabled={status.running || !backendAvailable}
                      InputProps={{
                        ...params.InputProps,
                        startAdornment: (
                          <InputAdornment position="start">
                            <IconButton
                              data-testid="open-directory"
                              onClick={() => setChooserOpen(true)}
                              size="small"
                            >
                              <FolderIcon />
                            </IconButton>
                          </InputAdornment>
                        ),
                      }}
                    />
                  )}
                />
                <FormControl fullWidth>
                  <InputLabel>Language</InputLabel>
                  <Select
                    value={lang}
                    label="Language"
                    data-testid="language-select"
                    onChange={e => setLang(e.target.value)}
                    disabled={status.running || !backendAvailable}
                  >
                    <MenuItem value="en" data-testid="language-option-en">
                      English
                    </MenuItem>
                    <MenuItem value="es" data-testid="language-option-es">
                      Spanish
                    </MenuItem>
                    <MenuItem value="fr" data-testid="language-option-fr">
                      French
                    </MenuItem>
                    <MenuItem value="de" data-testid="language-option-de">
                      German
                    </MenuItem>
                    <MenuItem value="it" data-testid="language-option-it">
                      Italian
                    </MenuItem>
                    <MenuItem value="pt" data-testid="language-option-pt">
                      Portuguese
                    </MenuItem>
                  </Select>
                </FormControl>
                <FormControl fullWidth>
                  <InputLabel>Provider</InputLabel>
                  <Select
                    value={provider}
                    label="Provider"
                    data-testid="provider-select"
                    onChange={e => {
                      setProvider(e.target.value);
                      const p = availableProviders.find(
                        pr => pr.name === e.target.value
                      );
                      if (p && !p.configured) {
                        setConfigSnackOpen(true);
                      }
                    }}
                    disabled={status.running || !backendAvailable}
                  >
                    {availableProviders.length > 0
                      ? availableProviders.map(p => (
                          <MenuItem
                            key={p.name}
                            value={p.name}
                            data-testid={`provider-option-${p.name.toLowerCase()}`}
                          >
                            {formatProviderName(p.name)}
                            {!p.configured && (
                              <Chip
                                label="Config Required"
                                size="small"
                                color="warning"
                                sx={{ ml: 1 }}
                              />
                            )}
                          </MenuItem>
                        ))
                      : // Fallback options if providers haven't loaded
                        [
                          <MenuItem
                            key="opensubtitles"
                            value="opensubtitles"
                            data-testid="provider-option-opensubtitles"
                          >
                            OpenSubtitles
                          </MenuItem>,
                          <MenuItem
                            key="addic7ed"
                            value="addic7ed"
                            data-testid="provider-option-addic7ed"
                          >
                            Addic7ed
                          </MenuItem>,
                          <MenuItem
                            key="subscene"
                            value="subscene"
                            data-testid="provider-option-subscene"
                          >
                            Subscene
                          </MenuItem>,
                          <MenuItem
                            key="podnapisi"
                            value="podnapisi"
                            data-testid="provider-option-podnapisi"
                          >
                            Podnapisi
                          </MenuItem>,
                        ]}
                  </Select>
                  {availableProviders.length === 0 && (
                    <Alert
                      severity="warning"
                      sx={{ mt: 1, fontSize: '0.875rem' }}
                    >
                      No providers configured. Go to Settings → Providers to
                      enable subtitle providers.
                    </Alert>
                  )}
                </FormControl>
                <Button
                  variant="contained"
                  startIcon={
                    status.running ? (
                      <CircularProgress size={20} />
                    ) : loading ? (
                      <CircularProgress size={20} />
                    ) : (
                      <PlayIcon />
                    )
                  }
                  onClick={start}
                  disabled={
                    status.running || !dir || !backendAvailable || loading
                  }
                  fullWidth
                  size="large"
                >
                  {status.running
                    ? 'Scanning...'
                    : loading
                      ? 'Starting...'
                      : !backendAvailable
                        ? 'Backend Unavailable'
                        : 'Start Scan'}
                </Button>
              </Box>
            </CardContent>
          </Card>
        </Grid>

        {/* Status Panel */}
        <Grid size={{ xs: 12, md: 4 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Scan Status
              </Typography>
              <Box sx={{ mb: 2 }}>
                <Chip
                  label={status.running ? 'Running' : 'Idle'}
                  color={status.running ? 'primary' : 'default'}
                  variant={status.running ? 'filled' : 'outlined'}
                />
              </Box>
              {status.running && (
                <Box sx={{ mb: 2 }}>
                  <Typography
                    variant="body2"
                    color="text.secondary"
                    gutterBottom
                  >
                    Progress: {status.completed} files processed
                  </Typography>
                  <LinearProgress variant="indeterminate" sx={{ mb: 1 }} />
                </Box>
              )}
              {status.files.length > 0 && (
                <Alert severity="info" sx={{ mt: 2 }}>
                  Found {status.files.length} files to process
                </Alert>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* System Info */}
        {systemInfo && (
          <Grid size={{ xs: 12, md: 4 }}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  System Info
                </Typography>
                <List dense>
                  {Object.entries(systemInfo).map(([key, value]) => (
                    <ListItem key={key} divider>
                      <ListItemText primary={key} secondary={String(value)} />
                    </ListItem>
                  ))}
                </List>
              </CardContent>
            </Card>
          </Grid>
        )}

        {/* File List */}
        {status.files.length > 0 && (
          <Grid size={{ xs: 12 }}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  Files ({status.files.length})
                </Typography>
                <Paper sx={{ maxHeight: 300, overflow: 'auto' }}>
                  <List dense>
                    {status.files.map((file, index) => (
                      <ListItem key={index} divider>
                        <ListItemText
                          primary={file}
                          sx={{
                            '& .MuiListItemText-primary': {
                              fontSize: '0.875rem',
                              fontFamily: 'monospace',
                            },
                          }}
                        />
                      </ListItem>
                    ))}
                  </List>
                </Paper>
              </CardContent>
            </Card>
          </Grid>
        )}
      </Grid>
      <DirectoryChooser
        open={chooserOpen}
        onClose={() => setChooserOpen(false)}
        onSelect={path => setDir(path)}
      />
      <Snackbar
        open={configSnackOpen}
        autoHideDuration={6000}
        onClose={() => setConfigSnackOpen(false)}
        anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
      >
        <Alert
          onClose={() => setConfigSnackOpen(false)}
          severity="info"
          action={
            <Button
              color="inherit"
              size="small"
              onClick={() => navigate('/settings')}
            >
              Configure
            </Button>
          }
        >
          Provider requires configuration. Click Configure to edit settings.
        </Alert>
      </Snackbar>
    </Box>
  );
}
