import {
    Assessment as LogIcon,
    Memory as MemoryIcon,
    Refresh as RefreshIcon,
    Storage as StorageIcon,
    BugReport as SystemIcon,
    Schedule as TaskIcon,
} from "@mui/icons-material";
import {
    Alert,
    Box,
    Card,
    CardContent,
    Chip,
    CircularProgress,
    Grid,
    IconButton,
    List,
    ListItem,
    ListItemText,
    Paper,
    Tooltip,
    Typography
} from "@mui/material";
import { useEffect, useState } from "react";

/**
 * System component displays system information, logs, and running tasks.
 * Provides monitoring and debugging information for the subtitle manager.
 */

export default function System() {
  const [logs, setLogs] = useState([]);
  const [info, setInfo] = useState({});
  const [tasks, setTasks] = useState({});
  const [loading, setLoading] = useState(true);

  const loadSystemData = async () => {
    setLoading(true);
    try {
      const [logsRes, infoRes, tasksRes] = await Promise.all([
        fetch("/api/logs"),
        fetch("/api/system"),
        fetch("/api/tasks")
      ]);

      if (logsRes.ok) setLogs(await logsRes.json());
      if (infoRes.ok) setInfo(await infoRes.json());
      if (tasksRes.ok) setTasks(await tasksRes.json());
    } catch (error) {
      console.error("Failed to load system data:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadSystemData();
  }, []);

  const formatBytes = (bytes) => {
    if (!bytes) return 'N/A';
    const sizes = ['Bytes', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    return `${Math.round(bytes / Math.pow(1024, i) * 100) / 100} ${sizes[i]}`;
  };

  const formatUptime = (seconds) => {
    if (!seconds) return 'N/A';
    const days = Math.floor(seconds / 86400);
    const hours = Math.floor((seconds % 86400) / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);
    return `${days}d ${hours}h ${minutes}m`;
  };

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box>
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={3}>
        <Typography variant="h4" component="h1">
          System Monitor
        </Typography>
        <Tooltip title="Refresh system data">
          <IconButton onClick={loadSystemData}>
            <RefreshIcon />
          </IconButton>
        </Tooltip>
      </Box>

      <Grid container spacing={3}>
        {/* System Information */}
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                <SystemIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
                System Information
              </Typography>
              <List dense>
                {Object.entries(info).map(([key, value]) => (
                  <ListItem key={key} divider>
                    <ListItemText
                      primary={key.charAt(0).toUpperCase() + key.slice(1)}
                      secondary={
                        key.includes('memory') || key.includes('size')
                          ? formatBytes(value)
                          : key.includes('uptime')
                          ? formatUptime(value)
                          : String(value)
                      }
                    />
                  </ListItem>
                ))}
              </List>
            </CardContent>
          </Card>
        </Grid>

        {/* Running Tasks */}
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                <TaskIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
                Running Tasks
              </Typography>
              {Object.keys(tasks).length === 0 ? (
                <Alert severity="info">No running tasks</Alert>
              ) : (
                <List dense>
                  {Object.entries(tasks).map(([taskId, taskInfo]) => (
                    <ListItem key={taskId} divider>
                      <ListItemText
                        primary={taskId}
                        secondary={
                          <Box display="flex" alignItems="center" gap={1} mt={1}>
                            <Chip
                              label={taskInfo.status || 'Running'}
                              size="small"
                              color={taskInfo.status === 'completed' ? 'success' : 'primary'}
                            />
                            {taskInfo.progress && (
                              <Chip
                                label={`${taskInfo.progress}%`}
                                size="small"
                                variant="outlined"
                              />
                            )}
                          </Box>
                        }
                      />
                    </ListItem>
                  ))}
                </List>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* Recent Logs */}
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                <LogIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
                Recent Logs
              </Typography>
              <Paper
                variant="outlined"
                sx={{
                  maxHeight: 400,
                  overflow: 'auto',
                  backgroundColor: 'grey.900',
                  color: 'common.white',
                  fontFamily: 'monospace'
                }}
              >
                <Box p={2}>
                  {logs.length === 0 ? (
                    <Typography color="text.secondary">No logs available</Typography>
                  ) : (
                    <pre
                      data-testid="logs"
                      style={{
                        margin: 0,
                        whiteSpace: 'pre-wrap',
                        fontSize: '0.875rem',
                        lineHeight: 1.4
                      }}
                    >
                      {logs.join("\n")}
                    </pre>
                  )}
                </Box>
              </Paper>
            </CardContent>
          </Card>
        </Grid>

        {/* Raw Task Data (for debugging) */}
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                <StorageIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
                Tasks (Raw Data)
              </Typography>
              <Paper
                variant="outlined"
                sx={{
                  maxHeight: 300,
                  overflow: 'auto',
                  backgroundColor: 'grey.50'
                }}
              >
                <Box p={2}>
                  <pre
                    data-testid="tasks"
                    style={{
                      margin: 0,
                      fontSize: '0.75rem',
                      whiteSpace: 'pre-wrap'
                    }}
                  >
                    {JSON.stringify(tasks, null, 2)}
                  </pre>
                </Box>
              </Paper>
            </CardContent>
          </Card>
        </Grid>

        {/* Raw System Info (for debugging) */}
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                <MemoryIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
                System Info (Raw Data)
              </Typography>
              <Paper
                variant="outlined"
                sx={{
                  maxHeight: 300,
                  overflow: 'auto',
                  backgroundColor: 'grey.50'
                }}
              >
                <Box p={2}>
                  <pre
                    data-testid="info"
                    style={{
                      margin: 0,
                      fontSize: '0.75rem',
                      whiteSpace: 'pre-wrap'
                    }}
                  >
                    {JSON.stringify(info, null, 2)}
                  </pre>
                </Box>
              </Paper>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Box>
  );
}
