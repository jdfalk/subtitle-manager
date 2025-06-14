// file: webui/src/components/ProviderConfigDialog.jsx

import {
    Close as CloseIcon,
    Save as SaveIcon,
    Visibility as VisibilityIcon,
    VisibilityOff as VisibilityOffIcon,
} from "@mui/icons-material";
import {
    Alert,
    Box,
    Button,
    Chip,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Divider,
    FormControl,
    FormControlLabel,
    Grid,
    IconButton,
    InputAdornment,
    InputLabel,
    MenuItem,
    Select,
    Switch,
    TextField,
    Typography,
} from "@mui/material";
import { useState } from "react";

/**
 * ProviderConfigDialog provides a comprehensive configuration interface
 * for subtitle providers. Handles different provider types and their
 * specific configuration requirements.
 *
 * @param {boolean} open - Whether dialog is open
 * @param {Object} provider - Provider to configure
 * @param {Function} onClose - Close callback
 * @param {Function} onSave - Save callback with configuration
 */
export default function ProviderConfigDialog({ open, provider, onClose, onSave }) {
  const [config, setConfig] = useState(provider?.config || {});
  const [showPasswords, setShowPasswords] = useState({});
  const [errors, setErrors] = useState({});

  const getProviderFields = (providerName) => {
    // Define configuration fields for different providers
    const fieldConfigs = {
      opensubtitles: [
        { name: 'apiKey', label: 'API Key', type: 'password', required: true,
          description: 'Get your API key from opensubtitles.com' },
        { name: 'username', label: 'Username', type: 'text', required: false },
        { name: 'rateLimit', label: 'Rate Limit (requests/minute)', type: 'number',
          defaultValue: 20, min: 1, max: 200 },
        { name: 'languages', label: 'Preferred Languages', type: 'multiselect',
          options: ['en', 'es', 'fr', 'de', 'it', 'pt', 'nl', 'ru', 'ja', 'ko'] },
      ],
      addic7ed: [
        { name: 'username', label: 'Username', type: 'text', required: true },
        { name: 'password', label: 'Password', type: 'password', required: true },
        { name: 'userAgent', label: 'User Agent', type: 'text',
          defaultValue: 'Mozilla/5.0 (compatible; SubtitleManager)' },
      ],
      subscene: [
        { name: 'userAgent', label: 'User Agent', type: 'text',
          defaultValue: 'Mozilla/5.0 (compatible; SubtitleManager)' },
        { name: 'timeout', label: 'Timeout (seconds)', type: 'number',
          defaultValue: 30, min: 5, max: 300 },
      ],
      whisper: [
        { name: 'model', label: 'Whisper Model', type: 'select',
          options: ['tiny', 'base', 'small', 'medium', 'large'], defaultValue: 'base' },
        { name: 'language', label: 'Force Language', type: 'select',
          options: ['auto', 'en', 'es', 'fr', 'de', 'it', 'pt'], defaultValue: 'auto' },
        { name: 'device', label: 'Device', type: 'select',
          options: ['auto', 'cpu', 'cuda'], defaultValue: 'auto' },
      ],
      generic: [
        { name: 'baseUrl', label: 'Base URL', type: 'url', required: true },
        { name: 'apiKey', label: 'API Key', type: 'password' },
        { name: 'timeout', label: 'Timeout (seconds)', type: 'number', defaultValue: 30 },
      ]
    };

    return fieldConfigs[providerName] || [
      { name: 'enabled', label: 'Enable Provider', type: 'boolean', defaultValue: true },
      { name: 'timeout', label: 'Timeout (seconds)', type: 'number', defaultValue: 30 },
    ];
  };

  const handleFieldChange = (fieldName, value) => {
    setConfig(prev => ({ ...prev, [fieldName]: value }));
    if (errors[fieldName]) {
      setErrors(prev => ({ ...prev, [fieldName]: null }));
    }
  };

  const togglePasswordVisibility = (fieldName) => {
    setShowPasswords(prev => ({ ...prev, [fieldName]: !prev[fieldName] }));
  };

  const validateConfig = () => {
    const fields = getProviderFields(provider?.name);
    const newErrors = {};

    fields.forEach(field => {
      if (field.required && !config[field.name]) {
        newErrors[field.name] = `${field.label} is required`;
      }
      if (field.type === 'url' && config[field.name] && !isValidUrl(config[field.name])) {
        newErrors[field.name] = 'Please enter a valid URL';
      }
      if (field.type === 'number' && config[field.name]) {
        const value = Number(config[field.name]);
        if (field.min && value < field.min) {
          newErrors[field.name] = `Minimum value is ${field.min}`;
        }
        if (field.max && value > field.max) {
          newErrors[field.name] = `Maximum value is ${field.max}`;
        }
      }
    });

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const isValidUrl = (string) => {
    try {
      new URL(string);
      return true;
    } catch {
      return false;
    }
  };

  const handleSave = () => {
    if (validateConfig()) {
      onSave({ ...provider, config });
    }
  };

  const renderField = (field) => {
    const value = config[field.name] ?? field.defaultValue ?? '';

    switch (field.type) {
      case 'boolean':
        return (
          <FormControlLabel
            control={
              <Switch
                checked={Boolean(value)}
                onChange={(e) => handleFieldChange(field.name, e.target.checked)}
              />
            }
            label={field.label}
          />
        );

      case 'select':
        return (
          <FormControl fullWidth error={Boolean(errors[field.name])}>
            <InputLabel>{field.label}</InputLabel>
            <Select
              value={value}
              label={field.label}
              onChange={(e) => handleFieldChange(field.name, e.target.value)}
            >
              {field.options?.map(option => (
                <MenuItem key={option} value={option}>
                  {option.charAt(0).toUpperCase() + option.slice(1)}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        );

      case 'multiselect':
        return (
          <FormControl fullWidth>
            <InputLabel>{field.label}</InputLabel>
            <Select
              multiple
              value={Array.isArray(value) ? value : []}
              label={field.label}
              onChange={(e) => handleFieldChange(field.name, e.target.value)}
              renderValue={(selected) => (
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {selected.map((value) => (
                    <Chip key={value} label={value.toUpperCase()} size="small" />
                  ))}
                </Box>
              )}
            >
              {field.options?.map(option => (
                <MenuItem key={option} value={option}>
                  {option.toUpperCase()}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        );

      case 'password':
        return (
          <TextField
            fullWidth
            label={field.label}
            type={showPasswords[field.name] ? 'text' : 'password'}
            value={value}
            onChange={(e) => handleFieldChange(field.name, e.target.value)}
            error={Boolean(errors[field.name])}
            helperText={errors[field.name] || field.description}
            required={field.required}
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton
                    onClick={() => togglePasswordVisibility(field.name)}
                    edge="end"
                  >
                    {showPasswords[field.name] ? <VisibilityOffIcon /> : <VisibilityIcon />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
          />
        );

      default:
        return (
          <TextField
            fullWidth
            label={field.label}
            type={field.type}
            value={value}
            onChange={(e) => handleFieldChange(field.name, e.target.value)}
            error={Boolean(errors[field.name])}
            helperText={errors[field.name] || field.description}
            required={field.required}
            inputProps={{
              min: field.min,
              max: field.max,
            }}
          />
        );
    }
  };

  if (!provider) return null;

  const fields = getProviderFields(provider.name);

  return (
    <Dialog open={open} onClose={onClose} maxWidth="md" fullWidth>
      <DialogTitle>
        <Box display="flex" alignItems="center" justifyContent="space-between">
          <Typography variant="h6">
            Configure {provider.displayName || provider.name}
          </Typography>
          <IconButton onClick={onClose}>
            <CloseIcon />
          </IconButton>
        </Box>
      </DialogTitle>

      <DialogContent>
        <Box sx={{ mb: 2 }}>
          <Alert severity="info" sx={{ mb: 2 }}>
            Configure this provider to enable subtitle downloads. Required fields are marked with *.
          </Alert>
        </Box>

        <Grid container spacing={3}>
          {fields.map((field) => (
            <Grid item xs={12} sm={field.type === 'boolean' ? 12 : 6} key={field.name}>
              {renderField(field)}
            </Grid>
          ))}
        </Grid>

        {provider.name === 'opensubtitles' && (
          <Box sx={{ mt: 3 }}>
            <Divider sx={{ mb: 2 }} />
            <Alert severity="warning">
              <Typography variant="body2">
                <strong>OpenSubtitles API Key Required:</strong><br />
                1. Visit <a href="https://www.opensubtitles.com/api" target="_blank" rel="noopener noreferrer">opensubtitles.com/api</a><br />
                2. Create an account and generate an API key<br />
                3. Enter your API key above to enable this provider
              </Typography>
            </Alert>
          </Box>
        )}
      </DialogContent>

      <DialogActions>
        <Button onClick={onClose}>
          Cancel
        </Button>
        <Button
          onClick={handleSave}
          variant="contained"
          startIcon={<SaveIcon />}
        >
          Save Configuration
        </Button>
      </DialogActions>
    </Dialog>
  );
}
