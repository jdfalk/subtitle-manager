// file: webui/src/OfflineInfo.jsx
import {
    Build as BuildIcon,
    Cancel as CancelIcon,
    CheckCircle as CheckCircleIcon,
    CloudOff as CloudOffIcon,
    Info as InfoIcon,
    Storage as StorageIcon,
} from '@mui/icons-material';
import {
    Box,
    Card,
    CardContent,
    Chip,
    Divider,
    List,
    ListItem,
    ListItemIcon,
    ListItemText,
    Typography,
} from '@mui/material';

/**
 * OfflineInfo component that displays information about what features
 * are available when the backend is not accessible
 * @returns {JSX.Element} The offline information component
 */
function OfflineInfo() {
  return (
    <Box sx={{ maxWidth: 800, mx: 'auto' }}>
      <Typography variant="h4" gutterBottom sx={{ mb: 4, textAlign: 'center' }}>
        Offline Mode Information
      </Typography>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
            <CloudOffIcon sx={{ mr: 2, color: 'warning.main' }} />
            <Typography variant="h6">Current Status</Typography>
          </Box>
          <Typography variant="body1" color="text.secondary" sx={{ mb: 2 }}>
            The Subtitle Manager backend service is currently unavailable. This
            means most functionality that requires server communication is
            temporarily disabled.
          </Typography>
          <Chip
            label="Backend Offline"
            color="error"
            variant="outlined"
            icon={<CancelIcon />}
          />
        </CardContent>
      </Card>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
            <CheckCircleIcon sx={{ mr: 2, color: 'success.main' }} />
            <Typography variant="h6">Available Features</Typography>
          </Box>
          <List dense>
            <ListItem>
              <ListItemIcon>
                <CheckCircleIcon color="success" />
              </ListItemIcon>
              <ListItemText
                primary="Theme Switching"
                secondary="Dark/light mode and kid-friendly interface toggles"
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <CheckCircleIcon color="success" />
              </ListItemIcon>
              <ListItemText
                primary="User Interface"
                secondary="Full navigation and UI components remain functional"
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <CheckCircleIcon color="success" />
              </ListItemIcon>
              <ListItemText
                primary="Local Storage"
                secondary="Settings like theme preferences are preserved locally"
              />
            </ListItem>
          </List>
        </CardContent>
      </Card>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
            <CancelIcon sx={{ mr: 2, color: 'error.main' }} />
            <Typography variant="h6">Unavailable Features</Typography>
          </Box>
          <List dense>
            <ListItem>
              <ListItemIcon>
                <CancelIcon color="error" />
              </ListItemIcon>
              <ListItemText
                primary="Subtitle Management"
                secondary="Search, download, and conversion of subtitles"
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <CancelIcon color="error" />
              </ListItemIcon>
              <ListItemText
                primary="Media Library"
                secondary="Scanning and managing your media collection"
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <CancelIcon color="error" />
              </ListItemIcon>
              <ListItemText
                primary="Provider Integration"
                secondary="OpenSubtitles, Radarr, Sonarr, and Plex integration"
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <CancelIcon color="error" />
              </ListItemIcon>
              <ListItemText
                primary="Translation Services"
                secondary="AI-powered subtitle translation"
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <CancelIcon color="error" />
              </ListItemIcon>
              <ListItemText
                primary="System Configuration"
                secondary="Settings and system management"
              />
            </ListItem>
          </List>
        </CardContent>
      </Card>

      <Card>
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
            <BuildIcon sx={{ mr: 2, color: 'info.main' }} />
            <Typography variant="h6">Troubleshooting</Typography>
          </Box>
          <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
            To restore full functionality, please ensure:
          </Typography>
          <List dense>
            <ListItem>
              <ListItemIcon>
                <InfoIcon color="info" />
              </ListItemIcon>
              <ListItemText
                primary="Server is Running"
                secondary="The Subtitle Manager backend service is started and accessible"
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <InfoIcon color="info" />
              </ListItemIcon>
              <ListItemText
                primary="Network Connection"
                secondary="Check your network connection and firewall settings"
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <StorageIcon color="info" />
              </ListItemIcon>
              <ListItemText
                primary="Database Access"
                secondary="Ensure the database is accessible and properly configured"
              />
            </ListItem>
          </List>

          <Divider sx={{ my: 2 }} />

          <Typography variant="body2" color="text.secondary">
            Once the backend service is restored, refresh this page to regain
            full functionality.
          </Typography>
        </CardContent>
      </Card>
    </Box>
  );
}

export default OfflineInfo;
