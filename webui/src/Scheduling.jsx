// file: webui/src/Scheduling.jsx
import {
  Alarm as CleanupIcon,
  Storage as DiskIcon,
  Cached as RefreshIcon,
} from '@mui/icons-material';
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Select,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';
import { apiService } from './services/api.js';

/**
 * Scheduling page allows configuring automated maintenance tasks.
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend is reachable
 */
export default function Scheduling({ backendAvailable = true }) {
  const [config, setConfig] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const frequencies = ['never', 'hourly', 'daily', 'weekly', 'monthly'];

  const loadConfig = async () => {
    setLoading(true);
    try {
      const res = await apiService.get('/api/config');
      if (res.ok) {
        setConfig(await res.json());
      } else {
        setError('Failed to load configuration');
      }
    } catch {
      setError('Failed to load configuration');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadConfig();
  }, []);

  const saveConfig = async () => {
    setError(null);
    try {
      const res = await apiService.post('/api/config', {
        db_cleanup_frequency: config.db_cleanup_frequency,
        metadata_refresh_frequency: config.metadata_refresh_frequency,
        disk_scan_frequency: config.disk_scan_frequency,
      });
      if (!res.ok) {
        setError('Failed to save configuration');
      }
    } catch {
      setError('Failed to update schedule');
    }
  };

  if (loading) {
    return null;
  }

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Scheduling
      </Typography>
      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}
      <Grid container spacing={3}>
        <Grid item xs={12} md={4}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <CleanupIcon sx={{ mr: 1 }} />
                <Typography variant="h6">Database Cleanup</Typography>
              </Box>
              <FormControl fullWidth>
                <InputLabel>Frequency</InputLabel>
                <Select
                  value={config.db_cleanup_frequency || 'daily'}
                  label="Frequency"
                  onChange={e =>
                    setConfig({
                      ...config,
                      db_cleanup_frequency: e.target.value,
                    })
                  }
                >
                  {frequencies.map(f => (
                    <MenuItem key={f} value={f}>
                      {f}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <RefreshIcon sx={{ mr: 1 }} />
                <Typography variant="h6">Metadata Refresh</Typography>
              </Box>
              <FormControl fullWidth>
                <InputLabel>Frequency</InputLabel>
                <Select
                  value={config.metadata_refresh_frequency || 'weekly'}
                  label="Frequency"
                  onChange={e =>
                    setConfig({
                      ...config,
                      metadata_refresh_frequency: e.target.value,
                    })
                  }
                >
                  {frequencies.map(f => (
                    <MenuItem key={f} value={f}>
                      {f}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={4}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <DiskIcon sx={{ mr: 1 }} />
                <Typography variant="h6">Disk Scan</Typography>
              </Box>
              <FormControl fullWidth>
                <InputLabel>Frequency</InputLabel>
                <Select
                  value={config.disk_scan_frequency || 'weekly'}
                  label="Frequency"
                  onChange={e =>
                    setConfig({
                      ...config,
                      disk_scan_frequency: e.target.value,
                    })
                  }
                >
                  {frequencies.map(f => (
                    <MenuItem key={f} value={f}>
                      {f}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
      <Box sx={{ mt: 3 }}>
        <Button
          variant="contained"
          onClick={saveConfig}
          disabled={!backendAvailable}
        >
          Save
        </Button>
      </Box>
    </Box>
  );
}
