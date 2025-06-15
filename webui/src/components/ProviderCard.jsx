// file: webui/src/components/ProviderCard.jsx

import {
  Add as AddIcon,
  Delete as DeleteIcon,
  CheckCircle as EnabledIcon,
  Settings as SettingsIcon,
} from '@mui/icons-material';
import {
  Avatar,
  Box,
  Card,
  CardActions,
  CardContent,
  Chip,
  FormControlLabel,
  IconButton,
  Switch,
  Tooltip,
  Typography,
} from '@mui/material';

/**
 * ProviderCard displays a single subtitle provider with enable/disable toggle
 * and configuration options. Follows Bazarr-style provider management.
 *
 * @param {Object} provider - Provider configuration object
 * @param {string} provider.name - Provider name (e.g., "opensubtitles")
 * @param {string} provider.displayName - Human-readable name
 * @param {boolean} provider.enabled - Whether provider is enabled
 * @param {string} provider.description - Provider description
 * @param {Array} provider.languages - Supported languages
 * @param {Function} onToggle - Callback when provider is enabled/disabled
 * @param {Function} onConfigure - Callback when configure button is clicked
 * @param {Function} onDelete - Callback when delete button is clicked (for custom providers)
 */
export default function ProviderCard({
  provider,
  onToggle,
  onConfigure,
  onDelete,
  isAddCard = false,
}) {
  if (isAddCard) {
    return (
      <Card
        sx={{
          height: 200,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          border: '2px dashed',
          borderColor: 'primary.main',
          backgroundColor: 'action.hover',
          cursor: 'pointer',
          transition: 'all 0.2s ease-in-out',
          '&:hover': {
            backgroundColor: 'action.selected',
            transform: 'translateY(-2px)',
          },
        }}
        onClick={onConfigure}
      >
        <CardContent sx={{ textAlign: 'center' }}>
          <Avatar
            sx={{
              width: 64,
              height: 64,
              mx: 'auto',
              mb: 1,
              backgroundColor: 'primary.main',
              fontSize: '2rem',
            }}
          >
            <AddIcon fontSize="large" />
          </Avatar>
          <Typography variant="h6" color="primary">
            Add Provider
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Configure a new subtitle provider
          </Typography>
        </CardContent>
      </Card>
    );
  }

  const getProviderIcon = name => {
    // Return first letter or custom icons for well-known providers
    const customIcons = {
      opensubtitles: '🎬',
      addic7ed: '🎭',
      subscene: '🎪',
      podnapisi: '🎨',
      whisper: '🎤',
      embedded: '📀',
    };
    return customIcons[name.toLowerCase()] || name.charAt(0).toUpperCase();
  };

  return (
    <Card
      sx={{
        height: 200,
        transition: 'all 0.2s ease-in-out',
        border: provider.enabled ? '2px solid' : '1px solid',
        borderColor: provider.enabled ? 'success.main' : 'divider',
        '&:hover': {
          transform: 'translateY(-2px)',
          boxShadow: 4,
        },
      }}
    >
      <CardContent sx={{ pb: 1 }}>
        <Box display="flex" alignItems="center" mb={1}>
          <Avatar
            sx={{
              width: 40,
              height: 40,
              mr: 1,
              backgroundColor: provider.enabled ? 'success.main' : 'grey.400',
              fontSize: '1.2rem',
            }}
          >
            {getProviderIcon(provider.name)}
          </Avatar>
          <Box flex={1}>
            <Typography variant="h6" noWrap>
              {provider.displayName || provider.name}
            </Typography>
            <FormControlLabel
              control={
                <Switch
                  checked={provider.enabled}
                  onChange={e => onToggle(provider.name, e.target.checked)}
                  size="small"
                />
              }
              label={provider.enabled ? 'Enabled' : 'Disabled'}
              sx={{ mr: 0 }}
            />
          </Box>
        </Box>

        <Typography
          variant="body2"
          color="text.secondary"
          sx={{
            mb: 1,
            overflow: 'hidden',
            textOverflow: 'ellipsis',
            display: '-webkit-box',
            WebkitLineClamp: 2,
            WebkitBoxOrient: 'vertical',
          }}
        >
          {provider.description ||
            `${provider.displayName || provider.name} subtitle provider`}
        </Typography>

        {provider.languages && provider.languages.length > 0 && (
          <Box display="flex" flexWrap="wrap" gap={0.5} mb={1}>
            {provider.languages.slice(0, 3).map(lang => (
              <Chip
                key={lang}
                label={lang.toUpperCase()}
                size="small"
                variant="outlined"
                sx={{ fontSize: '0.7rem', height: 20 }}
              />
            ))}
            {provider.languages.length > 3 && (
              <Chip
                label={`+${provider.languages.length - 3}`}
                size="small"
                variant="outlined"
                sx={{ fontSize: '0.7rem', height: 20 }}
              />
            )}
          </Box>
        )}
      </CardContent>

      <CardActions sx={{ pt: 0, justifyContent: 'space-between' }}>
        <Box>
          <Tooltip title="Configure provider">
            <IconButton
              size="small"
              onClick={() => onConfigure(provider)}
              color="primary"
            >
              <SettingsIcon fontSize="small" />
            </IconButton>
          </Tooltip>
        </Box>

        <Box display="flex" alignItems="center">
          {provider.configured && (
            <Tooltip title="Provider is configured">
              <EnabledIcon color="success" fontSize="small" sx={{ mr: 1 }} />
            </Tooltip>
          )}
          {provider.custom && onDelete && (
            <Tooltip title="Remove provider">
              <IconButton
                size="small"
                onClick={() => onDelete(provider.name)}
                color="error"
              >
                <DeleteIcon fontSize="small" />
              </IconButton>
            </Tooltip>
          )}
        </Box>
      </CardActions>
    </Card>
  );
}
