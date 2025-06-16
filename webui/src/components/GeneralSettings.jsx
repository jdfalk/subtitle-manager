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
  Alert,
} from '@mui/material';
import { useEffect, useState } from 'react';

/**
 * GeneralSettings provides Bazarr-compatible configuration options for the core
 * server settings. Values are loaded from the provided config object and
 * persisted through the onSave callback.
 *
 * @param {Object} props - Component properties
 * @param {Object} props.config - Current configuration values
 * @param {Function} props.onSave - Callback invoked with updated settings
 * @param {boolean} [props.backendAvailable=true] - Whether the backend API is reachable
 * @returns {JSX.Element} Form for managing general settings
 */
export default function GeneralSettings({
  config,
  onSave,
  backendAvailable = true,
}) {
  // Host settings
  const [address, setAddress] = useState('');
  const [port, setPort] = useState(8080);
  const [baseURL, setBaseURL] = useState('');

  // Proxy settings
  const [proxyEnabled, setProxyEnabled] = useState(false);
  const [proxyType, setProxyType] = useState('http');
  const [proxyHost, setProxyHost] = useState('');
  const [proxyPort, setProxyPort] = useState('');
  const [proxyUsername, setProxyUsername] = useState('');
  const [proxyPassword, setProxyPassword] = useState('');

  // Update settings
  const [autoUpdate, setAutoUpdate] = useState(false);
  const [updateBranch, setUpdateBranch] = useState('master');
  const [updateFrequency, setUpdateFrequency] = useState('daily');

  // Logging settings
  const [logLevel, setLogLevel] = useState('info');
  const [logFilter, setLogFilter] = useState('');
  const [logFilterRegex, setLogFilterRegex] = useState(false);
  const [logFilterIgnoreCase, setLogFilterIgnoreCase] = useState(true);

  // Backup settings
  const [backupEnabled, setBackupEnabled] = useState(true);
  const [backupFrequency, setBackupFrequency] = useState('weekly');
  const [backupRetention, setBackupRetention] = useState(5);
  const [backupLocation, setBackupLocation] = useState('');

  // Analytics settings
  const [analyticsEnabled, setAnalyticsEnabled] = useState(false);

  useEffect(() => {
    if (config) {
      // Host
      setAddress(config.address || '');
      setPort(config.port || 8080);
      setBaseURL(config.base_url || '');

      // Proxy
      setProxyEnabled(config.proxy_enabled || false);
      setProxyType(config.proxy_type || 'http');
      setProxyHost(config.proxy_host || '');
      setProxyPort(config.proxy_port || '');
      setProxyUsername(config.proxy_username || '');
      setProxyPassword(config.proxy_password || '');

      // Updates
      setAutoUpdate(config.auto_update || false);
      setUpdateBranch(config.update_branch || 'master');
      setUpdateFrequency(config.update_frequency || 'daily');

      // Logging
      setLogLevel(config.log_level || 'info');
      setLogFilter(config.log_filter || '');
      setLogFilterRegex(config.log_filter_regex || false);
      setLogFilterIgnoreCase(config.log_filter_ignore_case !== false);

      // Backups
      setBackupEnabled(config.backup_enabled !== false);
      setBackupFrequency(config.backup_frequency || 'weekly');
      setBackupRetention(config.backup_retention || 5);
      setBackupLocation(config.backup_location || '');

      // Analytics
      setAnalyticsEnabled(config.analytics_enabled || false);
    }
  }, [config]);

  const handleSave = () => {
    const newConfig = {
      address,
      port: parseInt(port, 10),
      base_url: baseURL,

      proxy_enabled: proxyEnabled,
      proxy_type: proxyType,
      proxy_host: proxyHost,
      proxy_port: proxyPort,
      proxy_username: proxyUsername,
      proxy_password: proxyPassword,

      auto_update: autoUpdate,
      update_branch: updateBranch,
      update_frequency: updateFrequency,

      log_level: logLevel,
      log_filter: logFilter,
      log_filter_regex: logFilterRegex,
      log_filter_ignore_case: logFilterIgnoreCase,

      backup_enabled: backupEnabled,
      backup_frequency: backupFrequency,
      backup_retention: parseInt(backupRetention, 10),
      backup_location: backupLocation,

      analytics_enabled: analyticsEnabled,
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

      {/* Host */}
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
                onChange={e => setAddress(e.target.value)}
                placeholder="0.0.0.0"
                helperText="IP address to bind to"
                disabled={!backendAvailable}
              />
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField
                fullWidth
                label="Port"
                type="number"
                value={port}
                onChange={e => setPort(e.target.value)}
                placeholder="8080"
                helperText="Web interface port"
                disabled={!backendAvailable}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="Base URL"
                value={baseURL}
                onChange={e => setBaseURL(e.target.value)}
                placeholder="/subtitles"
                helperText="Base URL for reverse proxy setups"
                disabled={!backendAvailable}
              />
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {/* Proxy */}
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
                  onChange={e => setProxyEnabled(e.target.checked)}
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
                    onChange={e => setProxyType(e.target.value)}
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
                  onChange={e => setProxyHost(e.target.value)}
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
                  onChange={e => setProxyPort(e.target.value)}
                  placeholder="8080"
                  disabled={!backendAvailable}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Username"
                  value={proxyUsername}
                  onChange={e => setProxyUsername(e.target.value)}
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
                  onChange={e => setProxyPassword(e.target.value)}
                  placeholder="Optional"
                  disabled={!backendAvailable}
                />
              </Grid>
            </Grid>
          )}
        </CardContent>
      </Card>

      {/* Updates */}
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
                      onChange={e => setAutoUpdate(e.target.checked)}
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
                      onChange={e => setUpdateBranch(e.target.value)}
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
                      onChange={e => setUpdateFrequency(e.target.value)}
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

      {/* Logging */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Logging
          </Typography>
          <Grid container spacing={2}>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel>Level</InputLabel>
                <Select
                  value={logLevel}
                  label="Level"
                  onChange={e => setLogLevel(e.target.value)}
                  disabled={!backendAvailable}
                >
                  <MenuItem value="debug">Debug</MenuItem>
                  <MenuItem value="info">Info</MenuItem>
                  <MenuItem value="warn">Warn</MenuItem>
                  <MenuItem value="error">Error</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} md={6}>
              <TextField
                fullWidth
                label="Filter"
                value={logFilter}
                onChange={e => setLogFilter(e.target.value)}
                placeholder="module=foo"
                helperText="Filter logs by module"
                disabled={!backendAvailable}
              />
            </Grid>
            <Grid item xs={12} md={4}>
              <FormControlLabel
                control={
                  <Switch
                    checked={logFilterRegex}
                    onChange={e => setLogFilterRegex(e.target.checked)}
                    disabled={!backendAvailable}
                  />
                }
                label="Regex"
              />
            </Grid>
            <Grid item xs={12} md={4}>
              <FormControlLabel
                control={
                  <Switch
                    checked={logFilterIgnoreCase}
                    onChange={e => setLogFilterIgnoreCase(e.target.checked)}
                    disabled={!backendAvailable}
                  />
                }
                label="Ignore Case"
              />
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {/* Backups */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Backups
          </Typography>
          <FormGroup sx={{ mb: 2 }}>
            <FormControlLabel
              control={
                <Switch
                  checked={backupEnabled}
                  onChange={e => setBackupEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              }
              label="Enable Backups"
            />
          </FormGroup>
          {backupEnabled && (
            <Grid container spacing={2}>
              <Grid item xs={12} md={4}>
                <FormControl fullWidth>
                  <InputLabel>Frequency</InputLabel>
                  <Select
                    value={backupFrequency}
                    label="Frequency"
                    onChange={e => setBackupFrequency(e.target.value)}
                    disabled={!backendAvailable}
                  >
                    <MenuItem value="daily">Daily</MenuItem>
                    <MenuItem value="weekly">Weekly</MenuItem>
                    <MenuItem value="monthly">Monthly</MenuItem>
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={12} md={4}>
                <TextField
                  fullWidth
                  label="Retention"
                  type="number"
                  value={backupRetention}
                  onChange={e => setBackupRetention(e.target.value)}
                  placeholder="5"
                  helperText="Number of backups to keep"
                  disabled={!backendAvailable}
                />
              </Grid>
              <Grid item xs={12} md={4}>
                <TextField
                  fullWidth
                  label="Location"
                  value={backupLocation}
                  onChange={e => setBackupLocation(e.target.value)}
                  placeholder="/path/to/backups"
                  disabled={!backendAvailable}
                />
              </Grid>
            </Grid>
          )}
        </CardContent>
      </Card>

      {/* Analytics */}
      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom color="primary">
            Analytics
          </Typography>
          <FormGroup>
            <FormControlLabel
              control={
                <Switch
                  checked={analyticsEnabled}
                  onChange={e => setAnalyticsEnabled(e.target.checked)}
                  disabled={!backendAvailable}
                />
              }
              label="Enable Anonymous Analytics"
            />
          </FormGroup>
        </CardContent>
      </Card>

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
