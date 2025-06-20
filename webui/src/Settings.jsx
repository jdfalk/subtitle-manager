// file: webui/src/Settings.jsx

import {
  Security as AuthIcon,
  Storage as DatabaseIcon,
  Settings as GeneralIcon,
  Download as ImportIcon,
  Notifications as NotificationIcon,
  CloudDownload as ProvidersIcon,
  Refresh as RefreshIcon,
  People as UsersIcon,
  Label as TagsIcon,
} from '@mui/icons-material';
import {
  Alert,
  Box,
  Button,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Grid,
  Paper,
  Snackbar,
  Tab,
  Tabs,
  Typography,
} from '@mui/material';
import { lazy, Suspense, useEffect, useState } from 'react';
import LoadingComponent from './components/LoadingComponent.jsx';
import ProviderCard from './components/ProviderCard.jsx';
import ProviderConfigDialog from './components/ProviderConfigDialog.jsx';
import { apiService } from './services/api.js';

// Lazy load settings components for better performance
const AuthSettings = lazy(() => import('./components/AuthSettings.jsx'));
const DatabaseSettings = lazy(
  () => import('./components/DatabaseSettings.jsx')
);
const GeneralSettings = lazy(() => import('./components/GeneralSettings.jsx'));
const NotificationSettings = lazy(
  () => import('./components/NotificationSettings.jsx')
);
const UserManagement = lazy(() => import('./UserManagement.jsx'));
const TagManagement = lazy(() => import('./TagManagement.jsx'));

/**
 * Settings component with modern tabbed interface for managing all aspects
 * of subtitle manager configuration. Includes provider management similar to Bazarr.
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend service is available
 */
export default function Settings({ backendAvailable = true }) {
  const [activeTab, setActiveTab] = useState(0);
  const [_config, setConfig] = useState(null);
  const [providers, setProviders] = useState([]);
  const [status, setStatus] = useState('');
  const [snackbarOpen, setSnackbarOpen] = useState(false);
  const [importDialogOpen, setImportDialogOpen] = useState(false);
  const [providerConfigDialog, setProviderConfigDialog] = useState({
    open: false,
    provider: null,
  });
  const [bazarrConfig, setBazarrConfig] = useState(null);
  const [importing, setImporting] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (backendAvailable) {
      loadConfig();
      loadProviders();
    }
  }, [backendAvailable]); // eslint-disable-line react-hooks/exhaustive-deps

  /**
   * Load current configuration from the server
   */
  const loadConfig = async () => {
    if (!backendAvailable) return;

    setLoading(true);
    try {
      setError(null);
      const response = await apiService.get('/api/config');
      if (response.ok) {
        const data = await response.json();
        setConfig(data);
      }
    } catch (error) {
      console.error('Failed to load configuration:', error);
      setStatus('Failed to load configuration');
      setError('Failed to load configuration');
    } finally {
      setLoading(false);
    }
  };

  /**
   * Load available providers and their current configuration
   */
  const loadProviders = async () => {
    if (!backendAvailable) return;

    try {
      setError(null);
      const response = await apiService.get('/api/providers');
      if (response.ok) {
        const data = await response.json();
        // Transform provider data to include display information
        const providersWithMetadata = data.map(provider => ({
          ...provider,
          displayName: formatProviderName(provider.name),
          description: getProviderDescription(provider.name),
          languages: getProviderLanguages(provider.name),
          configured: hasRequiredConfig(provider),
        }));
        setProviders(providersWithMetadata);
      }
    } catch (error) {
      console.error('Failed to load providers:', error);
      setError('Failed to load providers');
    }
  };

  /**
   * Format provider name for display (e.g., "opensubtitles" -> "OpenSubtitles")
   */
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

  /**
   * Get provider description based on provider type
   */
  const getProviderDescription = name => {
    const descriptions = {
      opensubtitles: 'Large community-driven subtitle database',
      addic7ed: 'TV shows subtitle provider with high quality subs',
      subscene: 'Popular subtitle site for movies and TV shows',
      whisper: 'AI-powered speech recognition for subtitle generation',
      embedded: 'Extract subtitles embedded in media files',
      generic: 'Generic HTTP/API-based subtitle provider',
    };

    return (
      descriptions[name] || `${formatProviderName(name)} subtitle provider`
    );
  };

  /**
   * Get commonly supported languages for a provider
   */
  const getProviderLanguages = name => {
    // This would ideally come from the provider metadata
    const commonLanguages = ['en', 'es', 'fr', 'de', 'it', 'pt'];
    const specialLanguages = {
      turkcealtyazi: ['tr', 'en'],
      greeksubtitles: ['el', 'en'],
      napiprojekt: ['pl', 'en'],
      legendasdivx: ['pt', 'es', 'en'],
    };

    return specialLanguages[name] || commonLanguages;
  };

  /**
   * Check if provider has required configuration
   */
  const hasRequiredConfig = provider => {
    if (!provider.config) return false;

    const requiredFields = {
      opensubtitles: ['apiKey'],
      addic7ed: ['username', 'password'],
      generic: ['baseUrl'],
    };

    const required = requiredFields[provider.name] || [];
    return required.every(field => provider.config[field]);
  };

  /**
   * Toggle provider enabled state
   */
  const handleProviderToggle = async (providerName, enabled) => {
    try {
      const response = await fetch(`/api/providers/${providerName}`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ enabled }),
      });

      if (response.ok) {
        setProviders(prev =>
          prev.map(p => (p.name === providerName ? { ...p, enabled } : p))
        );
        setStatus(
          `${formatProviderName(providerName)} ${enabled ? 'enabled' : 'disabled'}`
        );
        setSnackbarOpen(true);
      }
    } catch (error) {
      console.error('Failed to toggle provider:', error);
    }
  };

  /**
   * Open provider configuration dialog
   */
  const handleProviderConfigure = provider => {
    setProviderConfigDialog({ open: true, provider });
  };

  /**
   * Save provider configuration
   */
  const handleProviderSave = async provider => {
    try {
      const response = await fetch(`/api/providers/${provider.name}/config`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(provider.config),
      });

      if (response.ok) {
        setProviders(prev =>
          prev.map(p =>
            p.name === provider.name
              ? { ...provider, configured: hasRequiredConfig(provider) }
              : p
          )
        );
        setProviderConfigDialog({ open: false, provider: null });
        setStatus(`${formatProviderName(provider.name)} configuration saved`);
        setSnackbarOpen(true);
      }
    } catch (error) {
      console.error('Failed to save provider config:', error);
    }
  };

  /**
   * Save general configuration values via the API
   *
   * @param {Object} values - Key/value pairs to persist
   */
  const saveSettings = async values => {
    try {
      const response = await fetch('/api/config', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(values),
      });
      if (response.ok) {
        setStatus('Settings saved');
        setSnackbarOpen(true);
        await loadConfig();
      } else {
        setStatus('Failed to save settings');
        setSnackbarOpen(true);
      }
    } catch {
      setStatus('Error saving settings');
      setSnackbarOpen(true);
    }
  };

  /**
   * Import settings from Bazarr
   */
  const openImportDialog = async () => {
    setImporting(true);
    try {
      const response = await fetch('/api/bazarr/config');
      if (response.ok) {
        const data = await response.json();
        setBazarrConfig(data);
        setImportDialogOpen(true);
      } else {
        setStatus('Failed to fetch Bazarr configuration');
      }
    } catch {
      setStatus('Error connecting to Bazarr');
    } finally {
      setImporting(false);
    }
  };

  const importFromBazarr = async () => {
    setImporting(true);
    try {
      const response = await fetch('/api/bazarr/import', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
      });
      if (response.ok) {
        setStatus('Settings imported from Bazarr successfully');
        setImportDialogOpen(false);
        setSnackbarOpen(true);
        await loadConfig();
        await loadProviders();
      } else {
        setStatus('Failed to import from Bazarr');
      }
    } catch {
      setStatus('Error importing from Bazarr');
    } finally {
      setImporting(false);
    }
  };

  /**
   * Render provider management tab
   */
  const renderProvidersTab = () => (
    <Box>
      <Box
        display="flex"
        justifyContent="space-between"
        alignItems="center"
        mb={3}
      >
        <Typography variant="h5">Subtitle Providers</Typography>
        <Button
          variant="outlined"
          startIcon={
            importing ? <RefreshIcon className="spin" /> : <ImportIcon />
          }
          onClick={openImportDialog}
          disabled={importing}
        >
          Import from Bazarr
        </Button>
      </Box>

      <Alert severity="info" sx={{ mb: 3 }}>
        Enable and configure subtitle providers. Providers are checked in order,
        so arrange them by preference. Disabled providers are skipped during
        searches.
      </Alert>

      <Grid container spacing={3}>
        {providers
          .filter(p => p.enabled)
          .map(provider => (
            <Grid size={{ xs: 12, sm: 6, md: 4, lg: 3 }} key={provider.name}>
              <ProviderCard
                provider={provider}
                onToggle={handleProviderToggle}
                onConfigure={handleProviderConfigure}
              />
            </Grid>
          ))}

        {/* Add Provider Card */}
        <Grid size={{ xs: 12, sm: 6, md: 4, lg: 3 }}>
          <ProviderCard
            isAddCard
            onConfigure={() => {
              // Open dialog without a provider to show the full list
              setProviderConfigDialog({ open: true, provider: null });
            }}
          />
        </Grid>
      </Grid>
    </Box>
  );

  const tabs = [
    {
      label: 'Providers',
      icon: <ProvidersIcon />,
      component: renderProvidersTab,
    },
    {
      label: 'General',
      icon: <GeneralIcon />,
      component: () => (
        <Suspense
          fallback={<LoadingComponent message="Loading General Settings..." />}
        >
          <GeneralSettings
            config={_config}
            onSave={saveSettings}
            backendAvailable={backendAvailable}
          />
        </Suspense>
      ),
    },
    {
      label: 'Database',
      icon: <DatabaseIcon />,
      component: () => (
        <Suspense
          fallback={<LoadingComponent message="Loading Database Settings..." />}
        >
          <DatabaseSettings
            config={_config}
            onSave={saveSettings}
            backendAvailable={backendAvailable}
          />
        </Suspense>
      ),
    },
    {
      label: 'Authentication',
      icon: <AuthIcon />,
      component: () => (
        <Suspense
          fallback={
            <LoadingComponent message="Loading Authentication Settings..." />
          }
        >
          <AuthSettings
            config={_config}
            onSave={saveSettings}
            backendAvailable={backendAvailable}
          />
        </Suspense>
      ),
    },
    {
      label: 'Notifications',
      icon: <NotificationIcon />,
      component: () => (
        <Suspense
          fallback={
            <LoadingComponent message="Loading Notification Settings..." />
          }
        >
          <NotificationSettings
            config={_config}
            onSave={saveSettings}
            backendAvailable={backendAvailable}
          />
        </Suspense>
      ),
    },
    {
      label: 'Users',
      icon: <UsersIcon />,
      component: () => (
        <Suspense
          fallback={<LoadingComponent message="Loading User Management..." />}
        >
          <UserManagement backendAvailable={backendAvailable} />
        </Suspense>
      ),
    },
    {
      label: 'Tags',
      icon: <TagsIcon />,
      component: () => (
        <Suspense
          fallback={<LoadingComponent message="Loading Tag Management..." />}
        >
          <TagManagement backendAvailable={backendAvailable} />
        </Suspense>
      ),
    },
  ];

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="400px"
      >
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        Settings
      </Typography>

      {/* Backend availability warning */}
      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available. Settings cannot be loaded or
          modified at this time.
        </Alert>
      )}

      {/* Error display */}
      {error && (
        <Alert severity="error" sx={{ mb: 3 }} onClose={() => setError(null)}>
          {error}
        </Alert>
      )}

      <Paper sx={{ mb: 3 }}>
        <Tabs
          value={activeTab}
          onChange={(e, newValue) => setActiveTab(newValue)}
          variant="scrollable"
          scrollButtons="auto"
        >
          {tabs.map((tab, index) => (
            <Tab
              key={index}
              label={tab.label}
              icon={tab.icon}
              iconPosition="start"
              disabled={!backendAvailable}
            />
          ))}
        </Tabs>
      </Paper>

      <Box>{tabs[activeTab]?.component()}</Box>

      {/* Provider Configuration Dialog */}
      <ProviderConfigDialog
        open={providerConfigDialog.open}
        provider={providerConfigDialog.provider}
        onClose={() => setProviderConfigDialog({ open: false, provider: null })}
        onSave={handleProviderSave}
      />

      {/* Bazarr Import Dialog */}
      <Dialog
        open={importDialogOpen}
        onClose={() => setImportDialogOpen(false)}
        maxWidth="md"
        fullWidth
      >
        <DialogTitle>Import Settings from Bazarr</DialogTitle>
        <DialogContent>
          <Alert severity="info" sx={{ mb: 2 }}>
            This will import provider configurations, API keys, and other
            settings from your Bazarr installation. Existing settings will be
            merged or overwritten.
          </Alert>
          {bazarrConfig && (
            <Paper sx={{ p: 2, maxHeight: 300, overflow: 'auto' }}>
              <pre>{JSON.stringify(bazarrConfig, null, 2)}</pre>
            </Paper>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setImportDialogOpen(false)}>Cancel</Button>
          <Button
            onClick={importFromBazarr}
            variant="contained"
            disabled={importing}
            startIcon={
              importing ? <RefreshIcon className="spin" /> : <ImportIcon />
            }
          >
            {importing ? 'Importing...' : 'Import Settings'}
          </Button>
        </DialogActions>
      </Dialog>

      {/* Status Snackbar */}
      <Snackbar
        open={snackbarOpen}
        autoHideDuration={6000}
        onClose={() => setSnackbarOpen(false)}
      >
        <Alert
          onClose={() => setSnackbarOpen(false)}
          severity={
            status.includes('success') ||
            status.includes('enabled') ||
            status.includes('disabled') ||
            status.includes('saved')
              ? 'success'
              : 'error'
          }
          sx={{ width: '100%' }}
        >
          {status}
        </Alert>
      </Snackbar>
    </Box>
  );
}
