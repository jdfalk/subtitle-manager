// file: webui/src/components/GeneralSettings.jsx
import { Box, Button, FormControlLabel, MenuItem, Switch, TextField, Typography } from "@mui/material";
import { useEffect, useState } from "react";

/**
 * GeneralSettings provides configuration fields for core server options.
 *
 * @param {Object} props - Component properties
 * @param {Object} props.config - Current configuration values
 * @param {Function} props.onSave - Callback invoked with updated values
 */
export default function GeneralSettings({ config, onSave }) {
  const [serverName, setServerName] = useState("");
  const [baseURL, setBaseURL] = useState("");
  const [logLevel, setLogLevel] = useState("info");
  const [reverseProxy, setReverseProxy] = useState(false);

  useEffect(() => {
    if (config) {
      setServerName(config.server_name || "");
      setBaseURL(config.base_url || "");
      setLogLevel(config.log_level || "info");
      setReverseProxy(config.reverse_proxy || false);
    }
  }, [config]);

  const handleSave = () => {
    onSave({
      server_name: serverName,
      base_url: baseURL,
      log_level: logLevel,
      reverse_proxy: reverseProxy,
    });
  };

  return (
    <Box sx={{ maxWidth: 500 }}>
      <Typography variant="h6" gutterBottom>
        General Settings
      </Typography>
      <TextField
        label="Server Name"
        fullWidth
        sx={{ mb: 2 }}
        value={serverName}
        onChange={(e) => setServerName(e.target.value)}
      />
      <TextField
        label="Base URL"
        fullWidth
        sx={{ mb: 2 }}
        value={baseURL}
        onChange={(e) => setBaseURL(e.target.value)}
        helperText="Path prefix when behind a reverse proxy"
      />
      <TextField
        select
        label="Log Level"
        fullWidth
        sx={{ mb: 2 }}
        value={logLevel}
        onChange={(e) => setLogLevel(e.target.value)}
      >
        {['debug', 'info', 'warn', 'error'].map((lvl) => (
          <MenuItem key={lvl} value={lvl}>
            {lvl}
          </MenuItem>
        ))}
      </TextField>
      <FormControlLabel
        sx={{ mb: 2 }}
        control={<Switch checked={reverseProxy} onChange={(e) => setReverseProxy(e.target.checked)} />}
        label="Running behind a reverse proxy"
      />
      <Button variant="contained" onClick={handleSave}>
        Save
      </Button>
    </Box>
  );
}
