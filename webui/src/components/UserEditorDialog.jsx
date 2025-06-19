// file: webui/src/components/UserEditorDialog.jsx
import { useEffect, useState } from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Button,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
} from '@mui/material';

/**
 * UserEditorDialog allows creation or modification of a user.
 *
 * @param {Object} props - Component properties
 * @param {boolean} props.open - Whether the dialog is visible
 * @param {Object|null} props.user - User being edited or null for new user
 * @param {Function} props.onClose - Callback when dialog closes
 * @param {Function} props.onSave - Callback with user data when saved
 */
export default function UserEditorDialog({ open, user, onClose, onSave }) {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [role, setRole] = useState('user');

  // Load user data when the dialog opens
  useEffect(() => {
    if (open && user) {
      setUsername(user.username || '');
      setEmail(user.email || '');
      setRole(user.role || 'user');
    } else if (open) {
      setUsername('');
      setEmail('');
      setRole('user');
    }
  }, [open, user]);

  const handleSave = () => {
    onSave({
      id: user?.id,
      username: username.trim(),
      email: email.trim(),
      role,
    });
  };

  const isValid = () => username.trim() !== '' && email.trim() !== '';

  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>{user ? 'Edit User' : 'Add User'}</DialogTitle>
      <DialogContent>
        <TextField
          autoFocus
          fullWidth
          margin="normal"
          label="Username"
          value={username}
          onChange={e => setUsername(e.target.value)}
        />
        <TextField
          fullWidth
          margin="normal"
          label="Email"
          type="email"
          value={email}
          onChange={e => setEmail(e.target.value)}
        />
        <FormControl fullWidth margin="normal">
          <InputLabel id="role-label">Role</InputLabel>
          <Select
            labelId="role-label"
            value={role}
            label="Role"
            onChange={e => setRole(e.target.value)}
          >
            <MenuItem value="user">User</MenuItem>
            <MenuItem value="admin">Admin</MenuItem>
          </Select>
        </FormControl>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button onClick={handleSave} variant="contained" disabled={!isValid()}>
          Save
        </Button>
      </DialogActions>
    </Dialog>
  );
}
