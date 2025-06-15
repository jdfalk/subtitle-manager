// file: webui/src/components/AuthSettings.jsx
import { Box, Button, TextField, Typography } from "@mui/material";
import { useEffect, useState } from "react";

/**
 * AuthSettings configures OAuth integration settings such as GitHub client
 * credentials.
 *
 * @param {Object} props - Component properties
 * @param {Object} props.config - Current configuration values
 * @param {Function} props.onSave - Callback invoked with updated values
 */
export default function AuthSettings({ config, onSave }) {
  const [clientID, setClientID] = useState("");
  const [clientSecret, setClientSecret] = useState("");
  const [redirectURL, setRedirectURL] = useState("");

  useEffect(() => {
    if (config) {
      setClientID(config.github_client_id || "");
      setClientSecret(config.github_client_secret || "");
      setRedirectURL(config.github_redirect_url || "");
    }
  }, [config]);

  const handleSave = () => {
    onSave({
      github_client_id: clientID,
      github_client_secret: clientSecret,
      github_redirect_url: redirectURL,
    });
  };

  return (
    <Box sx={{ maxWidth: 500 }}>
      <Typography variant="h6" gutterBottom>
        OAuth Integrations
      </Typography>
      <TextField
        label="GitHub Client ID"
        fullWidth
        sx={{ mb: 2 }}
        value={clientID}
        onChange={(e) => setClientID(e.target.value)}
      />
      <TextField
        label="GitHub Client Secret"
        fullWidth
        sx={{ mb: 2 }}
        value={clientSecret}
        onChange={(e) => setClientSecret(e.target.value)}
      />
      <TextField
        label="GitHub Redirect URL"
        fullWidth
        sx={{ mb: 2 }}
        value={redirectURL}
        onChange={(e) => setRedirectURL(e.target.value)}
      />
      <Button variant="contained" onClick={handleSave}>
        Save
      </Button>
    </Box>
  );
}
