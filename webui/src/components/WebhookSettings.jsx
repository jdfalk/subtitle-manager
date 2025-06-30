// file: webui/src/components/WebhookSettings.jsx
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174008

import {
  Add as AddIcon,
  Delete as DeleteIcon,
  Edit as EditIcon,
  History as HistoryIcon,
  Send as TestIcon,
  Webhook as WebhookIcon,
} from '@mui/icons-material';
import {
  Alert,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Chip,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  FormControl,
  FormControlLabel,
  Grid,
  IconButton,
  InputLabel,
  MenuItem,
  Select,
  Snackbar,
  Switch,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TextField,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';
import { apiService } from '../services/api.js';

/**
 * WebhookSettings component for managing outgoing webhook endpoints.
 * Allows users to configure webhooks for subtitle events, test endpoints, and view history.
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend service is available
 */
export default function WebhookSettings({ backendAvailable = true }) {
  const [endpoints, setEndpoints] = useState([]);
  const [eventTypes, setEventTypes] = useState([]);
  const [history, setHistory] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);
  
  // Dialog states
  const [openCreate, setOpenCreate] = useState(false);
  const [openEdit, setOpenEdit] = useState(false);
  const [openHistory, setOpenHistory] = useState(false);
  const [openTest, setOpenTest] = useState(false);
  const [selectedEndpoint, setSelectedEndpoint] = useState(null);
  
  // Form state
  const [formData, setFormData] = useState({
    name: '',
    url: '',
    secret: '',
    events: [],
    headers: {},
    enabled: true,
  });

  useEffect(() => {
    if (backendAvailable) {
      loadData();
    }
  }, [backendAvailable]);

  const loadData = async () => {
    try {
      setLoading(true);
      const [endpointsRes, eventTypesRes] = await Promise.all([
        apiService.get('/api/webhooks/config'),
        apiService.get('/api/webhooks/event-types'),
      ]);
      
      setEndpoints(endpointsRes.endpoints || []);
      setEventTypes(eventTypesRes.event_types || []);
      setError(null);
    } catch (err) {
      setError('Failed to load webhook configuration');
      console.error('Error loading webhooks:', err);
    } finally {
      setLoading(false);
    }
  };

  const loadHistory = async () => {
    try {
      const response = await apiService.get('/api/webhooks/history?limit=50');
      setHistory(response.events || []);
    } catch (err) {
      setError('Failed to load webhook history');
      console.error('Error loading webhook history:', err);
    }
  };

  const handleCreate = () => {
    setFormData({
      name: '',
      url: '',
      secret: '',
      events: [],
      headers: {},
      enabled: true,
    });
    setOpenCreate(true);
  };

  const handleEdit = (endpoint) => {
    setSelectedEndpoint(endpoint);
    setFormData({
      name: endpoint.name,
      url: endpoint.url,
      secret: '', // Don't show existing secret
      events: endpoint.events || [],
      headers: endpoint.headers || {},
      enabled: endpoint.enabled,
    });
    setOpenEdit(true);
  };

  const handleSave = async () => {
    try {
      const payload = {
        name: formData.name,
        url: formData.url,
        events: formData.events,
        headers: formData.headers,
      };
      
      if (formData.secret) {
        payload.secret = formData.secret;
      }

      if (selectedEndpoint) {
        // Update existing endpoint
        payload.id = selectedEndpoint.id;
        payload.enabled = formData.enabled;
        await apiService.put('/api/webhooks/config', payload);
        setSuccess('Webhook endpoint updated successfully');
        setOpenEdit(false);
      } else {
        // Create new endpoint
        await apiService.post('/api/webhooks/config', payload);
        setSuccess('Webhook endpoint created successfully');
        setOpenCreate(false);
      }
      
      loadData();
      setSelectedEndpoint(null);
    } catch (err) {
      setError(err.response?.data || 'Failed to save webhook endpoint');
    }
  };

  const handleDelete = async (endpoint) => {
    if (!confirm(`Are you sure you want to delete the webhook "${endpoint.name}"?`)) {
      return;
    }
    
    try {
      await apiService.delete(`/api/webhooks/config?id=${endpoint.id}`);
      setSuccess('Webhook endpoint deleted successfully');
      loadData();
    } catch (err) {
      setError('Failed to delete webhook endpoint');
    }
  };

  const handleTest = async (endpoint) => {
    try {
      const payload = {
        url: endpoint.url,
        secret: endpoint.secret,
        headers: endpoint.headers,
      };
      
      await apiService.post('/api/webhooks/test', payload);
      setSuccess('Test webhook sent successfully');
    } catch (err) {
      setError('Failed to send test webhook');
    }
  };

  const handleTestDialog = () => {
    setOpenTest(true);
  };

  const handleTestCustom = async () => {
    try {
      const payload = {
        url: formData.url,
        secret: formData.secret,
        headers: formData.headers,
      };
      
      await apiService.post('/api/webhooks/test', payload);
      setSuccess('Test webhook sent successfully');
      setOpenTest(false);
    } catch (err) {
      setError('Failed to send test webhook');
    }
  };

  const handleEventChange = (event) => {
    const value = event.target.value;
    setFormData(prev => ({
      ...prev,
      events: typeof value === 'string' ? value.split(',') : value,
    }));
  };

  const formatTimestamp = (timestamp) => {
    if (!timestamp) return 'N/A';
    return new Date(timestamp).toLocaleString();
  };

  if (!backendAvailable) {
    return (
      <Alert severity="error" sx={{ mb: 3 }}>
        Backend service is not available. Webhook settings cannot be loaded.
      </Alert>
    );
  }

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" p={3}>
        <Typography>Loading webhook settings...</Typography>
      </Box>
    );
  }

  return (
    <Box>
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={3}>
        <Typography variant="h5" component="h2" sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
          <WebhookIcon />
          Webhook Settings
        </Typography>
        <Box>
          <Button
            variant="outlined"
            startIcon={<HistoryIcon />}
            onClick={() => {
              loadHistory();
              setOpenHistory(true);
            }}
            sx={{ mr: 1 }}
          >
            History
          </Button>
          <Button
            variant="outlined"
            startIcon={<TestIcon />}
            onClick={handleTestDialog}
            sx={{ mr: 1 }}
          >
            Test
          </Button>
          <Button
            variant="contained"
            startIcon={<AddIcon />}
            onClick={handleCreate}
          >
            Add Webhook
          </Button>
        </Box>
      </Box>

      <Typography variant="body2" color="text.secondary" mb={3}>
        Configure webhook endpoints to receive real-time notifications about subtitle events.
        Webhooks are sent as HTTP POST requests with JSON payloads.
      </Typography>

      {/* Webhook Endpoints Grid */}
      <Grid container spacing={3}>
        {endpoints.map((endpoint) => (
          <Grid item xs={12} md={6} lg={4} key={endpoint.id}>
            <Card>
              <CardHeader
                title={endpoint.name}
                subheader={endpoint.url}
                action={
                  <Switch
                    checked={endpoint.enabled}
                    onChange={(e) => {
                      const updatedEndpoint = { ...endpoint, enabled: e.target.checked };
                      handleEdit(updatedEndpoint);
                    }}
                  />
                }
              />
              <CardContent>
                <Box mb={2}>
                  <Typography variant="body2" color="text.secondary" gutterBottom>
                    Events:
                  </Typography>
                  <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                    {endpoint.events.map((event) => (
                      <Chip key={event} label={event} size="small" />
                    ))}
                  </Box>
                </Box>
                <Typography variant="body2" color="text.secondary">
                  Last Success: {formatTimestamp(endpoint.last_success)}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Retry Count: {endpoint.retry_count || 0}
                </Typography>
              </CardContent>
              <CardActions>
                <IconButton onClick={() => handleEdit(endpoint)} title="Edit">
                  <EditIcon />
                </IconButton>
                <IconButton onClick={() => handleTest(endpoint)} title="Test">
                  <TestIcon />
                </IconButton>
                <IconButton onClick={() => handleDelete(endpoint)} title="Delete" color="error">
                  <DeleteIcon />
                </IconButton>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>

      {endpoints.length === 0 && (
        <Box textAlign="center" py={4}>
          <Typography variant="h6" color="text.secondary" gutterBottom>
            No webhook endpoints configured
          </Typography>
          <Typography variant="body2" color="text.secondary" mb={2}>
            Add a webhook endpoint to receive notifications about subtitle events.
          </Typography>
          <Button variant="contained" startIcon={<AddIcon />} onClick={handleCreate}>
            Add Your First Webhook
          </Button>
        </Box>
      )}

      {/* Create/Edit Dialog */}
      <Dialog open={openCreate || openEdit} onClose={() => {
        setOpenCreate(false);
        setOpenEdit(false);
        setSelectedEndpoint(null);
      }} maxWidth="md" fullWidth>
        <DialogTitle>
          {selectedEndpoint ? 'Edit Webhook Endpoint' : 'Create Webhook Endpoint'}
        </DialogTitle>
        <DialogContent>
          <Grid container spacing={2} sx={{ mt: 1 }}>
            <Grid item xs={12} sm={6}>
              <TextField
                fullWidth
                label="Name"
                value={formData.name}
                onChange={(e) => setFormData(prev => ({ ...prev, name: e.target.value }))}
                placeholder="My Webhook"
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                fullWidth
                label="URL"
                value={formData.url}
                onChange={(e) => setFormData(prev => ({ ...prev, url: e.target.value }))}
                placeholder="https://example.com/webhook"
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="Secret (optional)"
                type="password"
                value={formData.secret}
                onChange={(e) => setFormData(prev => ({ ...prev, secret: e.target.value }))}
                placeholder="HMAC secret for signature validation"
                helperText="Leave empty to disable signature validation"
              />
            </Grid>
            <Grid item xs={12}>
              <FormControl fullWidth>
                <InputLabel>Event Types</InputLabel>
                <Select
                  multiple
                  value={formData.events}
                  onChange={handleEventChange}
                  label="Event Types"
                  renderValue={(selected) => (
                    <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                      {selected.map((value) => (
                        <Chip key={value} label={value} size="small" />
                      ))}
                    </Box>
                  )}
                >
                  <MenuItem value="*">All Events</MenuItem>
                  {eventTypes.map((eventType) => (
                    <MenuItem key={eventType.type} value={eventType.type}>
                      {eventType.type} - {eventType.description}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            {selectedEndpoint && (
              <Grid item xs={12}>
                <FormControlLabel
                  control={
                    <Switch
                      checked={formData.enabled}
                      onChange={(e) => setFormData(prev => ({ ...prev, enabled: e.target.checked }))}
                    />
                  }
                  label="Enabled"
                />
              </Grid>
            )}
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => {
            setOpenCreate(false);
            setOpenEdit(false);
            setSelectedEndpoint(null);
          }}>
            Cancel
          </Button>
          <Button onClick={handleSave} variant="contained">
            {selectedEndpoint ? 'Update' : 'Create'}
          </Button>
        </DialogActions>
      </Dialog>

      {/* Test Dialog */}
      <Dialog open={openTest} onClose={() => setOpenTest(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Test Webhook</DialogTitle>
        <DialogContent>
          <Grid container spacing={2} sx={{ mt: 1 }}>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="URL"
                value={formData.url}
                onChange={(e) => setFormData(prev => ({ ...prev, url: e.target.value }))}
                placeholder="https://example.com/webhook"
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="Secret (optional)"
                type="password"
                value={formData.secret}
                onChange={(e) => setFormData(prev => ({ ...prev, secret: e.target.value }))}
                placeholder="HMAC secret for signature validation"
              />
            </Grid>
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenTest(false)}>Cancel</Button>
          <Button onClick={handleTestCustom} variant="contained">
            Send Test
          </Button>
        </DialogActions>
      </Dialog>

      {/* History Dialog */}
      <Dialog open={openHistory} onClose={() => setOpenHistory(false)} maxWidth="lg" fullWidth>
        <DialogTitle>Webhook Event History</DialogTitle>
        <DialogContent>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Timestamp</TableCell>
                  <TableCell>Event Type</TableCell>
                  <TableCell>Source</TableCell>
                  <TableCell>Data</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {history.map((event, index) => (
                  <TableRow key={index}>
                    <TableCell>{formatTimestamp(event.timestamp)}</TableCell>
                    <TableCell>
                      <Chip label={event.type} size="small" />
                    </TableCell>
                    <TableCell>{event.source}</TableCell>
                    <TableCell>
                      <Typography variant="body2" sx={{ fontFamily: 'monospace', fontSize: '0.75rem' }}>
                        {JSON.stringify(event.data, null, 2).substring(0, 100)}...
                      </Typography>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          {history.length === 0 && (
            <Box textAlign="center" py={4}>
              <Typography color="text.secondary">
                No webhook events found
              </Typography>
            </Box>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenHistory(false)}>Close</Button>
        </DialogActions>
      </Dialog>

      {/* Snackbar for success/error messages */}
      <Snackbar
        open={!!success}
        autoHideDuration={6000}
        onClose={() => setSuccess(null)}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
      >
        <Alert severity="success" onClose={() => setSuccess(null)}>
          {success}
        </Alert>
      </Snackbar>
      
      <Snackbar
        open={!!error}
        autoHideDuration={6000}
        onClose={() => setError(null)}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
      >
        <Alert severity="error" onClose={() => setError(null)}>
          {error}
        </Alert>
      </Snackbar>
    </Box>
  );
}