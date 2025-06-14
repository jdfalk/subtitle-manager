import {
  Download as ImportIcon,
  Info as InfoIcon,
  Refresh as RefreshIcon,
  Save as SaveIcon,
} from "@mui/icons-material";
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Grid,
  IconButton,
  Paper,
  Snackbar,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";

/**
 * Settings component renders a configuration form with Bazarr import functionality.
 * Values are loaded from `/api/config` and submitted back via POST to the same endpoint.
 * Supports importing settings from Bazarr via `/api/bazarr/import`.
 */
export default function Settings() {
  const [config, setConfig] = useState(null);
  const [status, setStatus] = useState("");
  const [snackbarOpen, setSnackbarOpen] = useState(false);
  const [importDialogOpen, setImportDialogOpen] = useState(false);
  const [bazarrConfig, setBazarrConfig] = useState(null);
  const [importing, setImporting] = useState(false);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    loadConfig();
  }, []);

  const loadConfig = async () => {
    setLoading(true);
    try {
      const response = await fetch("/api/config");
      const data = await response.json();
      setConfig(data);
    } catch (error) {
      console.error("Failed to load configuration:", error);
      setStatus("Failed to load configuration");
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (key, value) => {
    setConfig({ ...config, [key]: value });
  };

  const save = async () => {
    try {
      const res = await fetch("/api/config", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(config),
      });
      if (res.ok) {
        setStatus("Configuration saved successfully");
        setSnackbarOpen(true);
      } else {
        setStatus("Failed to save configuration");
      }
    } catch (error) {
      setStatus("Error saving configuration");
    }
  };

  const openImportDialog = async () => {
    setImporting(true);
    try {
      const response = await fetch("/api/bazarr/config");
      if (response.ok) {
        const data = await response.json();
        setBazarrConfig(data);
        setImportDialogOpen(true);
      } else {
        setStatus("Failed to fetch Bazarr configuration");
      }
    } catch (error) {
      setStatus("Error connecting to Bazarr");
    } finally {
      setImporting(false);
    }
  };

  const importFromBazarr = async () => {
    setImporting(true);
    try {
      const response = await fetch("/api/bazarr/import", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
      });
      if (response.ok) {
        setStatus("Settings imported from Bazarr successfully");
        setImportDialogOpen(false);
        setSnackbarOpen(true);
        await loadConfig(); // Reload configuration
      } else {
        setStatus("Failed to import from Bazarr");
      }
    } catch (error) {
      setStatus("Error importing from Bazarr");
    } finally {
      setImporting(false);
    }
  };

  const getConfigSections = () => {
    if (!config) return {};

    const sections = {
      'Server': {},
      'Database': {},
      'Providers': {},
      'Authentication': {},
      'General': {}
    };

    Object.entries(config).forEach(([key, value]) => {
      if (key.toLowerCase().includes('server') || key.toLowerCase().includes('port')) {
        sections['Server'][key] = value;
      } else if (key.toLowerCase().includes('db') || key.toLowerCase().includes('database')) {
        sections['Database'][key] = value;
      } else if (key.toLowerCase().includes('provider') || key.toLowerCase().includes('api')) {
        sections['Providers'][key] = value;
      } else if (key.toLowerCase().includes('auth') || key.toLowerCase().includes('user') || key.toLowerCase().includes('password')) {
        sections['Authentication'][key] = value;
      } else {
        sections['General'][key] = value;
      }
    });

    return sections;
  };

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
        <Typography>Loading configuration...</Typography>
      </Box>
    );
  }

  if (!config) {
    return (
      <Alert severity="error">
        Failed to load configuration. Please try refreshing the page.
      </Alert>
    );
  }

  const sections = getConfigSections();

  return (
    <Box>
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={3}>
        <Typography variant="h4" component="h1">
          Settings
        </Typography>
        <Box>
          <Tooltip title="Import settings from Bazarr">
            <Button
              variant="outlined"
              startIcon={importing ? <RefreshIcon className="spin" /> : <ImportIcon />}
              onClick={openImportDialog}
              disabled={importing}
              sx={{ mr: 1 }}
            >
              Import from Bazarr
            </Button>
          </Tooltip>
          <Button
            variant="contained"
            startIcon={<SaveIcon />}
            onClick={save}
          >
            Save Configuration
          </Button>
        </Box>
      </Box>

      <Grid container spacing={3}>
        {Object.entries(sections).map(([sectionName, sectionConfig]) => {
          if (Object.keys(sectionConfig).length === 0) return null;

          return (
            <Grid item xs={12} key={sectionName}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    {sectionName}
                  </Typography>
                  <TableContainer component={Paper} variant="outlined">
                    <Table size="small">
                      <TableHead>
                        <TableRow>
                          <TableCell width="30%">Setting</TableCell>
                          <TableCell>Value</TableCell>
                          <TableCell width="10%">Info</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {Object.entries(sectionConfig).map(([key, value]) => (
                          <TableRow key={key}>
                            <TableCell>
                              <Typography variant="body2" fontWeight="medium">
                                {key}
                              </Typography>
                            </TableCell>
                            <TableCell>
                              <TextField
                                fullWidth
                                size="small"
                                variant="outlined"
                                value={value || ''}
                                onChange={(e) => handleChange(key, e.target.value)}
                                type={key.toLowerCase().includes('password') ? 'password' : 'text'}
                                placeholder={`Enter ${key}`}
                              />
                            </TableCell>
                            <TableCell>
                              <IconButton size="small">
                                <InfoIcon fontSize="small" />
                              </IconButton>
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </CardContent>
              </Card>
            </Grid>
          );
        })}
      </Grid>

      {/* Import Dialog */}
      <Dialog open={importDialogOpen} onClose={() => setImportDialogOpen(false)} maxWidth="md" fullWidth>
        <DialogTitle>
          Import Settings from Bazarr
        </DialogTitle>
        <DialogContent>
          <Alert severity="info" sx={{ mb: 2 }}>
            This will import provider configurations, API keys, and other settings from your Bazarr installation.
            Existing settings will be merged or overwritten.
          </Alert>
          {bazarrConfig && (
            <Box>
              <Typography variant="h6" gutterBottom>
                Preview of Bazarr Settings
              </Typography>
              <Paper sx={{ p: 2, maxHeight: 400, overflow: 'auto' }}>
                <Grid container spacing={1}>
                  {Object.entries(bazarrConfig).map(([key, value]) => (
                    <Grid item xs={12} sm={6} key={key}>
                      <Chip
                        label={`${key}: ${String(value).substring(0, 30)}${String(value).length > 30 ? '...' : ''}`}
                        variant="outlined"
                        size="small"
                      />
                    </Grid>
                  ))}
                </Grid>
              </Paper>
            </Box>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setImportDialogOpen(false)}>
            Cancel
          </Button>
          <Button
            onClick={importFromBazarr}
            variant="contained"
            disabled={importing}
            startIcon={importing ? <RefreshIcon className="spin" /> : <ImportIcon />}
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
          severity={status.includes('success') ? 'success' : 'error'}
          sx={{ width: '100%' }}
        >
          {status}
        </Alert>
      </Snackbar>
    </Box>
  );
}
