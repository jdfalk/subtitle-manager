import {
    Assessment as LogIcon,
    Memory as MemoryIcon,
    Refresh as RefreshIcon,
    Storage as StorageIcon,
    BugReport as SystemIcon,
    Schedule as TaskIcon,
    Code as CodeIcon,
    ExpandMore as ExpandMoreIcon,
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
    Typography,
    Accordion,
    AccordionSummary,
    AccordionDetails,
    useTheme,
    alpha,
} from "@mui/material";
import { useEffect, useState } from "react";

/**
 * System component displays system information, logs, and running tasks.
 * Provides monitoring and debugging information for the subtitle manager.
 * Features Material Design 3 compliant UI with proper dark mode support.
 */
export default function System() {
  const [logs, setLogs] = useState([]);
  const [info, setInfo] = useState({});
  const [tasks, setTasks] = useState({});
  const [loading, setLoading] = useState(true);
  const [expandedRawData, setExpandedRawData] = useState(false);
  
  const theme = useTheme();
  const isDarkMode = theme.palette.mode === 'dark';

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
        <Typography variant="h4" component="h1" fontWeight={500}>
          System Monitor
        </Typography>
        <Tooltip title="Refresh system data">
          <IconButton 
            onClick={loadSystemData}
            sx={{ 
              backgroundColor: alpha(theme.palette.primary.main, 0.1),
              '&:hover': {
                backgroundColor: alpha(theme.palette.primary.main, 0.2),
              }
            }}
          >
            <RefreshIcon />
          </IconButton>
        </Tooltip>
      </Box>

      <Grid container spacing={3}>
        {/* System Information */}
        <Grid item xs={12} md={6}>
          <Card elevation={0}>
            <CardContent>
              <Typography variant="h6" gutterBottom sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <SystemIcon sx={{ mr: 1, color: 'primary.main' }} />
                System Information
              </Typography>
              <List dense>
                {Object.entries(info).map(([key, value]) => (
                  <ListItem key={key} divider sx={{ px: 0 }}>
                    <ListItemText
                      primary={
                        <Typography variant="body2" fontWeight={500} color="text.primary">
                          {key.charAt(0).toUpperCase() + key.slice(1).replace(/([A-Z])/g, ' $1')}
                        </Typography>
                      }
                      secondary={
                        <Typography variant="body2" color="text.secondary">
                          {key.includes('memory') || key.includes('size')
                            ? formatBytes(value)
                            : key.includes('uptime')
                            ? formatUptime(value)
                            : String(value)}
                        </Typography>
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
          <Card elevation={0}>
            <CardContent>
              <Typography variant="h6" gutterBottom sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <TaskIcon sx={{ mr: 1, color: 'primary.main' }} />
                Running Tasks
              </Typography>
              {Object.keys(tasks).length === 0 ? (
                <Alert severity="info" sx={{ borderRadius: 2 }}>
                  No running tasks
                </Alert>
              ) : (
                <List dense>
                  {Object.entries(tasks).map(([taskId, taskInfo]) => (
                    <ListItem key={taskId} divider sx={{ px: 0 }}>
                      <ListItemText
                        primary={
                          <Typography variant="body2" fontWeight={500} color="text.primary">
                            {taskId}
                          </Typography>
                        }
                        secondary={
                          <Box display="flex" alignItems="center" gap={1} mt={1}>
                            <Chip
                              label={taskInfo.status || 'Running'}
                              size="small"
                              color={taskInfo.status === 'completed' ? 'success' : 'primary'}
                              sx={{ fontSize: '0.75rem' }}
                            />
                            {taskInfo.progress && (
                              <Chip
                                label={`${taskInfo.progress}%`}
                                size="small"
                                variant="outlined"
                                sx={{ fontSize: '0.75rem' }}
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
          <Card elevation={0}>
            <CardContent>
              <Typography variant="h6" gutterBottom sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <LogIcon sx={{ mr: 1, color: 'primary.main' }} />
                Recent Logs
              </Typography>
              <Paper
                variant="outlined"
                sx={{
                  maxHeight: 400,
                  overflow: 'auto',
                  backgroundColor: isDarkMode ? '#0d1117' : '#f6f8fa',
                  borderRadius: 2,
                  border: '1px solid',
                  borderColor: 'divider',
                }}
              >
                <Box p={2}>
                  {logs.length === 0 ? (
                    <Typography color="text.secondary" sx={{ fontStyle: 'italic' }}>
                      No logs available
                    </Typography>
                  ) : (
                    <pre
                      data-testid="logs"
                      style={{
                        margin: 0,
                        whiteSpace: 'pre-wrap',
                        fontSize: '0.875rem',
                        lineHeight: 1.5,
                        color: isDarkMode ? '#e6edf3' : '#24292f',
                        fontFamily: '"Roboto Mono", "Consolas", "Monaco", monospace',
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

        {/* Raw Data Section - Collapsible */}
        <Grid item xs={12}>
          <Accordion 
            expanded={expandedRawData} 
            onChange={() => setExpandedRawData(!expandedRawData)}
            sx={{ 
              border: '1px solid',
              borderColor: 'divider',
              borderRadius: 2,
              '&:before': { display: 'none' },
            }}
          >
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
              <Typography variant="h6" sx={{ display: 'flex', alignItems: 'center' }}>
                <CodeIcon sx={{ mr: 1, color: 'primary.main' }} />
                Raw Data (Debug Information)
              </Typography>
            </AccordionSummary>
            <AccordionDetails>
              <Grid container spacing={3}>
                {/* Raw Tasks Data */}
                <Grid item xs={12} md={6}>
                  <Card variant="outlined">
                    <CardContent>
                      <Typography variant="subtitle1" gutterBottom sx={{ display: 'flex', alignItems: 'center' }}>
                        <StorageIcon sx={{ mr: 1, color: 'secondary.main' }} />
                        Tasks (Raw Data)
                      </Typography>
                      <Paper
                        variant="outlined"
                        sx={{
                          maxHeight: 300,
                          overflow: 'auto',
                          backgroundColor: isDarkMode ? '#0d1117' : '#f6f8fa',
                          borderRadius: 2,
                        }}
                      >
                        <Box p={2}>
                          <pre
                            data-testid="tasks"
                            style={{
                              margin: 0,
                              fontSize: '0.75rem',
                              whiteSpace: 'pre-wrap',
                              color: isDarkMode ? '#e6edf3' : '#24292f',
                              fontFamily: '"Roboto Mono", "Consolas", "Monaco", monospace',
                              lineHeight: 1.4,
                            }}
                          >
                            {JSON.stringify(tasks, null, 2)}
                          </pre>
                        </Box>
                      </Paper>
                    </CardContent>
                  </Card>
                </Grid>

                {/* Raw System Info */}
                <Grid item xs={12} md={6}>
                  <Card variant="outlined">
                    <CardContent>
                      <Typography variant="subtitle1" gutterBottom sx={{ display: 'flex', alignItems: 'center' }}>
                        <MemoryIcon sx={{ mr: 1, color: 'secondary.main' }} />
                        System Info (Raw Data)
                      </Typography>
                      <Paper
                        variant="outlined"
                        sx={{
                          maxHeight: 300,
                          overflow: 'auto',
                          backgroundColor: isDarkMode ? '#0d1117' : '#f6f8fa',
                          borderRadius: 2,
                        }}
                      >
                        <Box p={2}>
                          <pre
                            data-testid="info"
                            style={{
                              margin: 0,
                              fontSize: '0.75rem',
                              whiteSpace: 'pre-wrap',
                              color: isDarkMode ? '#e6edf3' : '#24292f',
                              fontFamily: '"Roboto Mono", "Consolas", "Monaco", monospace',
                              lineHeight: 1.4,
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
            </AccordionDetails>
          </Accordion>
        </Grid>
      </Grid>
    </Box>
  );
}
