// file: webui/src/components/SettingsOverview.jsx
// version: 1.0.0
// guid: 5f6e7d8c-9b0a-4c1d-8e2f-3a4b5c6d7e8f

import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardActionArea from '@mui/material/CardActionArea';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import SettingsIcon from '@mui/icons-material/Settings';
import StorageIcon from '@mui/icons-material/Storage';
import PeopleIcon from '@mui/icons-material/People';
import { useNavigate } from 'react-router-dom';
import BackButton from './BackButton.jsx';

/**
 * SettingsOverview displays high level configuration sections as cards.
 * Provides quick navigation to detailed settings pages.
 * @returns {JSX.Element} Settings overview grid
 */
export default function SettingsOverview() {
  const navigate = useNavigate();

  const cards = [
    { icon: <SettingsIcon />, label: 'General', path: '/settings/general' },
    { icon: <StorageIcon />, label: 'Providers', path: '/settings/providers' },
    { icon: <PeopleIcon />, label: 'Users', path: '/settings/users' },
  ];

  return (
    <>
      <BackButton />
      <Grid container spacing={2}>
        {cards.map(card => (
          <Grid item xs={12} sm={6} md={4} key={card.path}>
            <Card>
              <CardActionArea onClick={() => navigate(card.path)}>
                <CardContent sx={{ textAlign: 'center' }}>
                  {card.icon}
                  <Typography variant="h6" sx={{ mt: 1 }}>
                    {card.label}
                  </Typography>
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
        ))}
      </Grid>
    </>
  );
}
