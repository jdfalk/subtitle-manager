// file: webui/src/components/QuickLinks.jsx
import {
  History as HistoryIcon,
  VideoLibrary as LibraryIcon,
  Settings as SettingsIcon,
  BugReport as SystemIcon,
  CloudDownload as WantedIcon,
} from '@mui/icons-material';
import { Button, Card, CardContent, Grid, Typography } from '@mui/material';
import { NavLink } from 'react-router-dom';

/**
 * QuickLinks displays navigation shortcuts on the dashboard.
 * Provides buttons linking to common sections of the app.
 * @returns {JSX.Element} Quick links card
 */
export default function QuickLinks() {
  const links = [
    { label: 'Media Library', path: '/library', icon: <LibraryIcon /> },
    { label: 'Wanted', path: '/wanted', icon: <WantedIcon /> },
    { label: 'History', path: '/history', icon: <HistoryIcon /> },
    { label: 'Settings', path: '/settings', icon: <SettingsIcon /> },
    { label: 'System', path: '/system', icon: <SystemIcon /> },
  ];

  return (
    <Card sx={{ mb: 3 }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Quick Links
        </Typography>
        <Grid container spacing={2}>
          {links.map(link => (
            <Grid size={{ xs: 6, sm: 4, md: 2 }} key={link.path}>
              <Button
                component={NavLink}
                to={link.path}
                fullWidth
                variant="outlined"
                startIcon={link.icon}
              >
                {link.label}
              </Button>
            </Grid>
          ))}
        </Grid>
      </CardContent>
    </Card>
  );
}
