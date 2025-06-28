// file: webui/src/LibraryScan.jsx
import FolderIcon from '@mui/icons-material/Folder';
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  CircularProgress,
  LinearProgress,
  List,
  ListItem,
  ListItemText,
  Paper,
  TextField,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';
import DirectoryChooser from './components/DirectoryChooser.jsx';
import { apiService } from './services/api.js';

export default function LibraryScan({ backendAvailable = true }) {
  const [dir, setDir] = useState('');
  const [status, setStatus] = useState({
    running: false,
    completed: 0,
    files: [],
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [chooserOpen, setChooserOpen] = useState(false);

  const poll = async () => {
    try {
      const resp = await apiService.get('/api/library/scan/status');
      if (resp.ok) {
        const data = await resp.json();
        setStatus({
          running: data.running || false,
          completed: data.completed || 0,
          files: data.files || [],
        });
        if (data.running) {
          setTimeout(poll, 1000);
        }
      }
    } catch (err) {
      // Log poll failure before notifying user
      console.error(err);
      setError('Failed to get scan status');
    }
  };

  const start = async () => {
    if (!dir) return;
    setLoading(true);
    setError(null);
    try {
      const resp = await apiService.post('/api/library/scan', { path: dir });
      if (resp.ok) {
        poll();
      } else {
        setError('Failed to start scan');
      }
    } catch {
      setError('Failed to get scan status');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    poll();
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Library Scan
      </Typography>
      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available.
        </Alert>
      )}
      {error && (
        <Alert severity="error" sx={{ mb: 3 }} onClose={() => setError(null)}>
          {error}
        </Alert>
      )}
      <Card sx={{ maxWidth: 600 }}>
        <CardContent>
          <Box sx={{ '& > :not(style)': { m: 1 } }}>
            <TextField
              fullWidth
              label="Directory Path"
              placeholder="/path/to/library"
              value={dir}
              onChange={e => setDir(e.target.value)}
              disabled={status.running}
              InputProps={{
                startAdornment: (
                  <FolderIcon
                    onClick={() => setChooserOpen(true)}
                    sx={{ mr: 1 }}
                  />
                ),
              }}
            />
            <Button
              variant="contained"
              startIcon={
                loading || status.running ? (
                  <CircularProgress size={20} />
                ) : null
              }
              onClick={start}
              disabled={!dir || loading || status.running || !backendAvailable}
              fullWidth
            >
              {status.running ? 'Scanning...' : 'Start Scan'}
            </Button>
          </Box>
          {status.running && (
            <Box sx={{ mt: 2 }}>
              <LinearProgress />
              <Typography variant="body2" color="text.secondary" mt={1}>
                Processed {status.completed} files
              </Typography>
            </Box>
          )}
          {status.files.length > 0 && (
            <Paper sx={{ mt: 2, maxHeight: 200, overflow: 'auto' }}>
              <List dense>
                {status.files.map((f, i) => (
                  <ListItem key={i} divider>
                    <ListItemText
                      primary={f}
                      sx={{ fontFamily: 'monospace' }}
                    />
                  </ListItem>
                ))}
              </List>
            </Paper>
          )}
        </CardContent>
      </Card>
      <DirectoryChooser
        open={chooserOpen}
        onClose={() => setChooserOpen(false)}
        onSelect={p => setDir(p)}
      />
    </Box>
  );
}
