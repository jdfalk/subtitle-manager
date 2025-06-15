// file: webui/src/Dashboard.jsx
import {
  Folder as FolderIcon,
  PlayArrow as PlayIcon,
} from '@mui/icons-material';
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  CircularProgress,
  FormControl,
  Grid,
  InputLabel,
  LinearProgress,
  List,
  ListItem,
  ListItemText,
  MenuItem,
  Paper,
  Select,
  TextField,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';

/**
 * Dashboard component for managing subtitle scanning operations.
 * Provides controls for starting scans and monitoring progress.
 */

export default function Dashboard() {
  const [status, setStatus] = useState({
    running: false,
    completed: 0,
    files: [],
  });
  const [dir, setDir] = useState('');
  const [lang, setLang] = useState('en');
  // Default to embedded provider until others are added
  const [provider, setProvider] = useState('embedded');
  const [availableProviders, setAvailableProviders] = useState([]);

  useEffect(() => {
    poll();
    loadProviders();
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  const loadProviders = async () => {
    try {
      const response = await fetch('/api/providers');
      if (response.ok) {
        const data = await response.json();
        // Only show enabled providers
        const enabledProviders = data.filter(p => p.enabled);
        setAvailableProviders(enabledProviders);

        // Set default provider to first enabled one
        if (enabledProviders.length > 0 && !provider) {
          setProvider(enabledProviders[0].name);
        }
      }
    } catch (error) {
      console.error('Failed to load providers:', error);
      // Fallback to hardcoded providers if API fails
      setAvailableProviders([
        { name: 'opensubtitles', displayName: 'OpenSubtitles', enabled: true },
        { name: 'addic7ed', displayName: 'Addic7ed', enabled: true },
        { name: 'subscene', displayName: 'Subscene', enabled: true },
        { name: 'podnapisi', displayName: 'Podnapisi', enabled: true },
      ]);
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

  const poll = async () => {
    const res = await fetch('/api/scan/status');
    if (res.ok) {
      const data = await res.json();
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
  };

  const start = async () => {
    const body = { provider, directory: dir, lang };
    const res = await fetch('/api/scan', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    });
    if (res.ok) poll();
  };

  useEffect(() => {
    poll();
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        Dashboard
      </Typography>

      <Grid container spacing={3}>
        {/* Scan Controls */}
        <Grid size={{ xs: 12, md: 8 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Subtitle Scan
              </Typography>
              <Box component="form" sx={{ '& > :not(style)': { m: 1 } }}>
                <TextField
                  fullWidth
                  label="Directory Path"
                  placeholder="Enter directory to scan"
                  value={dir}
                  onChange={e => setDir(e.target.value)}
                  disabled={status.running}
                  InputProps={{
                    startAdornment: (
                      <FolderIcon sx={{ mr: 1, color: 'action.active' }} />
                    ),
                  }}
                />
                <FormControl fullWidth>
                  <InputLabel>Language</InputLabel>
                  <Select
                    value={lang}
                    label="Language"
                    onChange={e => setLang(e.target.value)}
                    disabled={status.running}
                  >
                    <MenuItem value="en">English</MenuItem>
                    <MenuItem value="es">Spanish</MenuItem>
                    <MenuItem value="fr">French</MenuItem>
                    <MenuItem value="de">German</MenuItem>
                    <MenuItem value="it">Italian</MenuItem>
                    <MenuItem value="pt">Portuguese</MenuItem>
                  </Select>
                </FormControl>
                <FormControl fullWidth>
                  <InputLabel>Provider</InputLabel>
                  <Select
                    value={provider}
                    label="Provider"
                    onChange={e => setProvider(e.target.value)}
                    disabled={status.running}
                  >
                    {availableProviders.length > 0
                      ? availableProviders.map(p => (
                          <MenuItem key={p.name} value={p.name}>
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
                          <MenuItem key="opensubtitles" value="opensubtitles">
                            OpenSubtitles
                          </MenuItem>,
                          <MenuItem key="addic7ed" value="addic7ed">
                            Addic7ed
                          </MenuItem>,
                          <MenuItem key="subscene" value="subscene">
                            Subscene
                          </MenuItem>,
                          <MenuItem key="podnapisi" value="podnapisi">
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
                    ) : (
                      <PlayIcon />
                    )
                  }
                  onClick={start}
                  disabled={status.running || !dir}
                  fullWidth
                  size="large"
                >
                  {status.running ? 'Scanning...' : 'Start Scan'}
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
    </Box>
  );
}
