// file: webui/src/components/AuthSettings.jsx
import { Box, Button, TextField, Typography } from "@mui/material";
import { useEffect, useState } from "react";

/**
 * AuthSettings configures authentication credentials and API key.
 *
 * @param {Object} props - Component properties
 * @param {Object} props.config - Current configuration values
 * @param {Function} props.onSave - Callback invoked with updated values
 */
export default function AuthSettings({ config, onSave }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [apiKey, setApiKey] = useState("");

  useEffect(() => {
    if (config && config.auth) {
      setUsername(config.auth.username || "");
      setPassword(config.auth.password || "");
      setApiKey(config.auth.api_key || "");
    }
  }, [config]);

  const handleSave = () => {
    onSave({
      "auth.username": username,
      "auth.password": password,
      "auth.api_key": apiKey,
    });
  };

  return (
    <Box sx={{ maxWidth: 500 }}>
      <Typography variant="h6" gutterBottom>
        Authentication
      </Typography>
      <TextField
        label="Username"
        fullWidth
        sx={{ mb: 2 }}
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <TextField
        label="Password"
        type="password"
        fullWidth
        sx={{ mb: 2 }}
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <TextField
        label="API Key"
        fullWidth
        sx={{ mb: 2 }}
        value={apiKey}
        onChange={(e) => setApiKey(e.target.value)}
      />
      <Button variant="contained" onClick={handleSave}>
        Save
      </Button>
    </Box>
  );
}
