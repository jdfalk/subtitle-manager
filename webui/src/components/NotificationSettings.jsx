// file: webui/src/components/NotificationSettings.jsx
import {
  Box,
  Button,
  FormControlLabel,
  Switch,
  TextField,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';

/**
 * NotificationSettings configures external notification services.
 *
 * @param {Object} props - Component properties
 * @param {Object} props.config - Current configuration values
 * @param {Function} props.onSave - Callback invoked with updated values
 */
export default function NotificationSettings({ config, onSave }) {
  const [enabled, setEnabled] = useState(false);
  const [discord, setDiscord] = useState('');
  const [telegramToken, setTelegramToken] = useState('');
  const [telegramChat, setTelegramChat] = useState('');
  const [emailURL, setEmailURL] = useState('');

  useEffect(() => {
    const n = (config && config.notifications) || {};
    setEnabled(n.enabled || false);
    setDiscord(n.discord_webhook || '');
    setTelegramToken(n.telegram_token || '');
    setTelegramChat(n.telegram_chat_id || '');
    setEmailURL(n.email_url || '');
  }, [config]);

  const handleSave = () => {
    onSave({
      'notifications.enabled': enabled,
      'notifications.discord_webhook': discord,
      'notifications.telegram_token': telegramToken,
      'notifications.telegram_chat_id': telegramChat,
      'notifications.email_url': emailURL,
    });
  };

  return (
    <Box sx={{ maxWidth: 500 }}>
      <Typography variant="h6" gutterBottom>
        Notifications
      </Typography>
      <FormControlLabel
        sx={{ mb: 2 }}
        control={
          <Switch
            checked={enabled}
            onChange={e => setEnabled(e.target.checked)}
          />
        }
        label="Enable notifications"
      />
      <TextField
        label="Discord Webhook"
        fullWidth
        sx={{ mb: 2 }}
        value={discord}
        onChange={e => setDiscord(e.target.value)}
      />
      <TextField
        label="Telegram Token"
        fullWidth
        sx={{ mb: 2 }}
        value={telegramToken}
        onChange={e => setTelegramToken(e.target.value)}
      />
      <TextField
        label="Telegram Chat ID"
        fullWidth
        sx={{ mb: 2 }}
        value={telegramChat}
        onChange={e => setTelegramChat(e.target.value)}
      />
      <TextField
        label="Email URL"
        fullWidth
        sx={{ mb: 2 }}
        value={emailURL}
        onChange={e => setEmailURL(e.target.value)}
      />
      <Button variant="contained" onClick={handleSave}>
        Save
      </Button>
    </Box>
  );
}
