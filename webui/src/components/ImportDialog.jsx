// file: webui/src/components/ImportDialog.jsx

import {
  CloudDownloadOutlined,
  Download as ImportIcon,
  Refresh as RefreshIcon,
} from '@mui/icons-material';
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  Checkbox,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import { useMemo } from 'react';

/**
 * ImportDialog component for importing settings from Bazarr
 * @param {Object} props - Component props
 * @param {boolean} props.open - Whether the dialog is open
 * @param {Function} props.onClose - Function to close the dialog
 * @param {string} props.bazarrURL - Bazarr URL
 * @param {Function} props.setBazarrURL - Function to set Bazarr URL
 * @param {string} props.bazarrAPIKey - Bazarr API key
 * @param {Function} props.setBazarrAPIKey - Function to set Bazarr API key
 * @param {Array} props.bazarrMappings - Array of Bazarr settings mappings
 * @param {Object} props.selectedSettings - Object of selected settings
 * @param {Function} props.setSelectedSettings - Function to set selected settings
 * @param {Object} props.previewConfig - Preview configuration object
 * @param {boolean} props.importLoading - Whether import is loading
 * @param {string} props.importError - Import error message
 * @param {boolean} props.importing - Whether currently importing
 * @param {Function} props.onPreview - Function to preview Bazarr settings
 * @param {Function} props.onImport - Function to import selected settings
 */
export default function ImportDialog({
  open,
  onClose,
  bazarrURL,
  setBazarrURL,
  bazarrAPIKey,
  setBazarrAPIKey,
  bazarrMappings,
  selectedSettings,
  setSelectedSettings,
  previewConfig,
  importLoading,
  importError,
  importing,
  onPreview,
  onImport,
}) {
  // Memoize the grouping operation to avoid unnecessary recalculations
  const groupedMappings = useMemo(() => {
    return bazarrMappings.reduce((groups, mapping) => {
      const section = mapping.section || 'Other';
      if (!groups[section]) {
        groups[section] = [];
      }
      groups[section].push(mapping);
      return groups;
    }, {});
  }, [bazarrMappings]);

  /**
   * Select all settings
   */
  const handleSelectAll = () => {
    const selected = {};
    bazarrMappings.forEach(mapping => {
      selected[mapping.key] = true;
    });
    setSelectedSettings(selected);
  };

  /**
   * Deselect all settings
   */
  const handleDeselectAll = () => {
    setSelectedSettings({});
  };

  /**
   * Handle individual setting selection
   */
  const handleSettingToggle = (key, checked) => {
    setSelectedSettings({
      ...selectedSettings,
      [key]: checked,
    });
  };

  return (
    <Dialog open={open} onClose={onClose} maxWidth="md" fullWidth>
      <DialogTitle>Import Settings from Bazarr</DialogTitle>
      <DialogContent>
        <Stack spacing={3} sx={{ mt: 1 }}>
          <TextField
            label="Bazarr URL"
            fullWidth
            value={bazarrURL}
            onChange={e => setBazarrURL(e.target.value)}
          />
          <TextField
            label="API Key"
            type="password"
            fullWidth
            value={bazarrAPIKey}
            onChange={e => setBazarrAPIKey(e.target.value)}
          />
          {importError && <Alert severity="error">{importError}</Alert>}
          <Button
            variant="outlined"
            onClick={onPreview}
            disabled={importLoading}
            startIcon={
              importLoading ? (
                <CircularProgress size={20} />
              ) : (
                <CloudDownloadOutlined />
              )
            }
          >
            {importLoading ? 'Connecting...' : 'Connect'}
          </Button>

          {previewConfig && bazarrMappings.length > 0 && (
            <Box>
              <Typography variant="h6" gutterBottom>
                Found Settings ({bazarrMappings.length} items)
              </Typography>
              <Stack direction="row" spacing={2} sx={{ mb: 2 }}>
                <Button
                  variant="outlined"
                  size="small"
                  onClick={handleSelectAll}
                >
                  Select All
                </Button>
                <Button
                  variant="outlined"
                  size="small"
                  onClick={handleDeselectAll}
                >
                  Deselect All
                </Button>
              </Stack>

              {Object.entries(groupedMappings).map(([section, mappings]) => (
                <Card key={section} variant="outlined" sx={{ mb: 2 }}>
                  <CardContent>
                    <Typography variant="h6" color="primary" sx={{ mb: 1 }}>
                      {section}
                    </Typography>
                    <List dense>
                      {mappings.map((mapping, idx) => (
                        <div key={mapping.key}>
                          <ListItem sx={{ pl: 0 }}>
                            <ListItemIcon>
                              <Checkbox
                                checked={selectedSettings[mapping.key] || false}
                                onChange={e =>
                                  handleSettingToggle(
                                    mapping.key,
                                    e.target.checked
                                  )
                                }
                              />
                            </ListItemIcon>
                            <ListItemText
                              primary={
                                <Typography variant="body2" fontWeight="medium">
                                  {mapping.description}
                                </Typography>
                              }
                              secondary={
                                <Box>
                                  <Typography
                                    variant="caption"
                                    color="primary"
                                    sx={{
                                      fontFamily: 'monospace',
                                      display: 'block',
                                    }}
                                  >
                                    {mapping.key}
                                  </Typography>
                                  <Typography
                                    variant="caption"
                                    sx={{
                                      fontFamily: 'monospace',
                                      display: 'block',
                                    }}
                                  >
                                    Value: {String(mapping.value)}
                                  </Typography>
                                </Box>
                              }
                            />
                          </ListItem>
                          {idx < mappings.length - 1 && <Divider />}
                        </div>
                      ))}
                    </List>
                  </CardContent>
                </Card>
              ))}

              <Alert severity="info">
                <Typography variant="body2">
                  {Object.values(selectedSettings).filter(Boolean).length} of{' '}
                  {bazarrMappings.length} settings selected for import
                </Typography>
              </Alert>
            </Box>
          )}
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button
          onClick={onImport}
          variant="contained"
          disabled={importing || !previewConfig}
          startIcon={
            importing ? <RefreshIcon className="spin" /> : <ImportIcon />
          }
        >
          {importing ? 'Importing...' : 'Import Selected'}
        </Button>
      </DialogActions>
    </Dialog>
  );
}
