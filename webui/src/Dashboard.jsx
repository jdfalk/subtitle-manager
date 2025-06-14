// file: webui/src/Dashboard.jsx
import { Folder as FolderIcon, PlayArrow as PlayIcon } from "@mui/icons-material";
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  CircularProgress,
  FormControl,
  Grid,
  InputLabel,
  LinearProgress,
  List,
  ListItem,
  ListItemText,
  MenuItem,
  Paper,
  Select,
  TextField,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";

/**
 * Dashboard component for managing subtitle scanning operations.
 * Provides controls for starting scans and monitoring progress.
 */

export default function Dashboard() {
  const [status, setStatus] = useState({ running: false, completed: 0, files: [] });
  const [dir, setDir] = useState("");
  const [lang, setLang] = useState("en");
  const [provider, setProvider] = useState("generic");

  const poll = async () => {
    const res = await fetch("/api/scan/status");
    if (res.ok) {
      const data = await res.json();
      // Ensure files is always an array to prevent null reference errors
      setStatus({
        running: data.running || false,
        completed: data.completed || 0,
        files: data.files || []
      });
      if (data.running) {
        setTimeout(poll, 1000);
      }
    }
  };

  const start = async () => {
    const body = { provider, directory: dir, lang };
    const res = await fetch("/api/scan", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });
    if (res.ok) poll();
  };

  useEffect(() => {
    poll();
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        Dashboard
      </Typography>

      <Grid container spacing={3}>
        {/* Scan Controls */}
        <Grid item xs={12} md={8}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Subtitle Scan
              </Typography>
              <Box component="form" sx={{ '& > :not(style)': { m: 1 } }}>
                <TextField
                  fullWidth
                  label="Directory Path"
                  placeholder="Enter directory to scan"
                  value={dir}
                  onChange={(e) => setDir(e.target.value)}
                  disabled={status.running}
                  InputProps={{
                    startAdornment: <FolderIcon sx={{ mr: 1, color: 'action.active' }} />,
                  }}
                />
                <FormControl fullWidth>
                  <InputLabel>Language</InputLabel>
                  <Select
                    value={lang}
                    label="Language"
                    onChange={(e) => setLang(e.target.value)}
                    disabled={status.running}
                  >
                    <MenuItem value="en">English</MenuItem>
                    <MenuItem value="es">Spanish</MenuItem>
                    <MenuItem value="fr">French</MenuItem>
                    <MenuItem value="de">German</MenuItem>
                    <MenuItem value="it">Italian</MenuItem>
                    <MenuItem value="pt">Portuguese</MenuItem>
                  </Select>
                </FormControl>
                <FormControl fullWidth>
                  <InputLabel>Provider</InputLabel>
                  <Select
                    value={provider}
                    label="Provider"
                    onChange={(e) => setProvider(e.target.value)}
                    disabled={status.running}
                  >
                    <MenuItem value="generic">Generic</MenuItem>
                    <MenuItem value="opensubtitles">OpenSubtitles</MenuItem>
                    <MenuItem value="addic7ed">Addic7ed</MenuItem>
                    <MenuItem value="podnapisi">Podnapisi</MenuItem>
                  </Select>
                </FormControl>
                <Button
                  variant="contained"
                  startIcon={status.running ? <CircularProgress size={20} /> : <PlayIcon />}
                  onClick={start}
                  disabled={status.running || !dir}
                  fullWidth
                  size="large"
                >
                  {status.running ? "Scanning..." : "Start Scan"}
                </Button>
              </Box>
            </CardContent>
          </Card>
        </Grid>

        {/* Status Panel */}
        <Grid item xs={12} md={4}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Scan Status
              </Typography>
              <Box sx={{ mb: 2 }}>
                <Chip
                  label={status.running ? "Running" : "Idle"}
                  color={status.running ? "primary" : "default"}
                  variant={status.running ? "filled" : "outlined"}
                />
              </Box>
              {status.running && (
                <Box sx={{ mb: 2 }}>
                  <Typography variant="body2" color="text.secondary" gutterBottom>
                    Progress: {status.completed} files processed
                  </Typography>
                  <LinearProgress
                    variant="indeterminate"
                    sx={{ mb: 1 }}
                  />
                </Box>
              )}
              {status.files.length > 0 && (
                <Alert severity="info" sx={{ mt: 2 }}>
                  Found {status.files.length} files to process
                </Alert>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* File List */}
        {status.files.length > 0 && (
          <Grid item xs={12}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  Files ({status.files.length})
                </Typography>
                <Paper sx={{ maxHeight: 300, overflow: 'auto' }}>
                  <List dense>
                    {status.files.map((file, index) => (
                      <ListItem key={index} divider>
                        <ListItemText
                          primary={file}
                          sx={{
                            '& .MuiListItemText-primary': {
                              fontSize: '0.875rem',
                              fontFamily: 'monospace'
                            }
                          }}
                        />
                      </ListItem>
                    ))}
                  </List>
                </Paper>
              </CardContent>
            </Card>
          </Grid>
        )}
      </Grid>
    </Box>
  );
}
