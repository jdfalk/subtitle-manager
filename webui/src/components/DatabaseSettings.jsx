// file: webui/src/components/DatabaseSettings.jsx
import {
  Box,
  Button,
  Card,
  CardContent,
  Typography,
  Grid,
  Chip,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Alert,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  CircularProgress,
  LinearProgress,
} from '@mui/material';
import {
  Storage as DatabaseIcon,
  Backup as BackupIcon,
  CleaningServices as OptimizeIcon,
  Assessment as StatsIcon,
} from '@mui/icons-material';
import { useEffect, useState } from 'react';

/**
 * Enhanced DatabaseSettings component with comprehensive database management.
 *
 * @param {Object} props - Component properties
 * @param {Object} props.config - Current configuration values
 * @param {Function} props.onSave - Callback invoked with updated values
 * @param {boolean} [props.backendAvailable=true] - Whether backend is reachable
 * @returns {JSX.Element} The rendered settings component
 */
export default function DatabaseSettings({
  config: _config,
  onSave: _onSave,
  backendAvailable = true,
}) {
  const [dbInfo, setDbInfo] = useState(null);
  const [loading, setLoading] = useState(false);
  const [backupDialog, setBackupDialog] = useState(false);
  const [optimizeDialog, setOptimizeDialog] = useState(false);
  const [stats, setStats] = useState(null);

  const loadDatabaseInfo = async () => {
    if (!backendAvailable) return;

    setLoading(true);
    try {
      const response = await fetch('/api/database/info');
      if (response.ok) {
        const data = await response.json();
        setDbInfo(data);
      }

      const statsResponse = await fetch('/api/database/stats');
      if (statsResponse.ok) {
        const statsData = await statsResponse.json();
        setStats(statsData);
      }
    } catch (error) {
      console.error('Failed to load database info:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadDatabaseInfo();
    // loadDatabaseInfo is defined in this component and does not need to be
    // included in the dependency array.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [backendAvailable]);

  const handleBackup = async () => {
    try {
      const response = await fetch('/api/database/backup', { method: 'POST' });
      if (response.ok) {
        const blob = await response.blob();
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'backup.tar.gz';
        document.body.appendChild(a);
        a.click();
        a.remove();
        window.URL.revokeObjectURL(url);
        setBackupDialog(false);
        loadDatabaseInfo();
      }
    } catch (error) {
      alert('Backup failed: ' + error.message);
    }
  };

  const handleOptimize = async () => {
    try {
      const response = await fetch('/api/database/optimize', {
        method: 'POST',
      });
      if (response.ok) {
        alert('Database optimization completed successfully');
        setOptimizeDialog(false);
        loadDatabaseInfo();
      }
    } catch (error) {
      alert('Optimization failed: ' + error.message);
    }
  };

  const formatBytes = bytes => {
    if (!bytes) return 'N/A';
    const sizes = ['Bytes', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    return Math.round((bytes / Math.pow(1024, i)) * 100) / 100 + ' ' + sizes[i];
  };

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="200px"
      >
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box sx={{ maxWidth: 1000 }}>
      <Typography variant="h6" gutterBottom>
        Database Settings
      </Typography>

      {!backendAvailable && (
        <Alert severity="warning" sx={{ mb: 3 }}>
          Backend service is not available. Database information cannot be
          loaded.
        </Alert>
      )}

      <Grid container spacing={3}>
        {/* Database Information */}
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography
                variant="h6"
                gutterBottom
                color="primary"
                sx={{ display: 'flex', alignItems: 'center' }}
              >
                <DatabaseIcon sx={{ mr: 1 }} />
                Database Information
              </Typography>

              {dbInfo ? (
                <TableContainer>
                  <Table size="small">
                    <TableBody>
                      <TableRow>
                        <TableCell>
                          <strong>Type</strong>
                        </TableCell>
                        <TableCell>
                          <Chip
                            label={dbInfo.type || 'Unknown'}
                            color={
                              dbInfo.type === 'postgresql'
                                ? 'primary'
                                : 'default'
                            }
                            size="small"
                          />
                        </TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Version</strong>
                        </TableCell>
                        <TableCell>{dbInfo.version || 'N/A'}</TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Size</strong>
                        </TableCell>
                        <TableCell>{formatBytes(dbInfo.size)}</TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Location</strong>
                        </TableCell>
                        <TableCell>{dbInfo.path || 'N/A'}</TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Connection Status</strong>
                        </TableCell>
                        <TableCell>
                          <Chip
                            label={
                              dbInfo.connected ? 'Connected' : 'Disconnected'
                            }
                            color={dbInfo.connected ? 'success' : 'error'}
                            size="small"
                          />
                        </TableCell>
                      </TableRow>
                    </TableBody>
                  </Table>
                </TableContainer>
              ) : (
                <Alert severity="info">
                  Database information not available
                </Alert>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* Database Statistics */}
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography
                variant="h6"
                gutterBottom
                color="primary"
                sx={{ display: 'flex', alignItems: 'center' }}
              >
                <StatsIcon sx={{ mr: 1 }} />
                Statistics
              </Typography>

              {stats ? (
                <TableContainer>
                  <Table size="small">
                    <TableBody>
                      <TableRow>
                        <TableCell>
                          <strong>Total Records</strong>
                        </TableCell>
                        <TableCell>
                          {stats.totalRecords?.toLocaleString() || 'N/A'}
                        </TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Users</strong>
                        </TableCell>
                        <TableCell>{stats.users || 0}</TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Downloads</strong>
                        </TableCell>
                        <TableCell>
                          {stats.downloads?.toLocaleString() || 0}
                        </TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Media Items</strong>
                        </TableCell>
                        <TableCell>
                          {stats.mediaItems?.toLocaleString() || 0}
                        </TableCell>
                      </TableRow>
                      <TableRow>
                        <TableCell>
                          <strong>Last Backup</strong>
                        </TableCell>
                        <TableCell>
                          {stats.lastBackup
                            ? new Date(stats.lastBackup).toLocaleDateString()
                            : 'Never'}
                        </TableCell>
                      </TableRow>
                    </TableBody>
                  </Table>
                </TableContainer>
              ) : (
                <Alert severity="info">Statistics not available</Alert>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* Database Management Actions */}
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom color="primary">
                Database Management
              </Typography>

              <Grid container spacing={2}>
                <Grid item>
                  <Button
                    variant="outlined"
                    startIcon={<BackupIcon />}
                    onClick={() => setBackupDialog(true)}
                    disabled={!backendAvailable}
                  >
                    Create Backup
                  </Button>
                </Grid>
                <Grid item>
                  <Button
                    variant="outlined"
                    startIcon={<OptimizeIcon />}
                    onClick={() => setOptimizeDialog(true)}
                    disabled={!backendAvailable}
                  >
                    Optimize Database
                  </Button>
                </Grid>
              </Grid>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      {/* Backup Confirmation Dialog */}
      <Dialog open={backupDialog} onClose={() => setBackupDialog(false)}>
        <DialogTitle>Create Database Backup</DialogTitle>
        <DialogContent>
          <DialogContentText>
            This will create a backup of your database. The process may take a
            few moments depending on the database size.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setBackupDialog(false)}>Cancel</Button>
          <Button onClick={handleBackup} variant="contained">
            Create Backup
          </Button>
        </DialogActions>
      </Dialog>

      {/* Optimize Confirmation Dialog */}
      <Dialog open={optimizeDialog} onClose={() => setOptimizeDialog(false)}>
        <DialogTitle>Optimize Database</DialogTitle>
        <DialogContent>
          <DialogContentText>
            This will optimize your database by rebuilding indexes and cleaning
            up unused space. This process may take several minutes.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOptimizeDialog(false)}>Cancel</Button>
          <Button onClick={handleOptimize} variant="contained">
            Optimize
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
