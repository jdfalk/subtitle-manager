// file: webui/src/UserManagement.jsx
import { useCallback, useEffect, useState } from 'react';
import {
  Alert,
  Box,
  Button,
  Chip,
  CircularProgress,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
} from '@mui/material';
import UserEditorDialog from './components/UserEditorDialog.jsx';

/**
 * UserManagement displays all users and allows password resets.
 * The component now properly fetches and shows usernames.
 *
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend is reachable
 */
export default function UserManagement({ backendAvailable = true }) {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [editor, setEditor] = useState({ open: false, user: null });

  const loadUsers = useCallback(async () => {
    if (!backendAvailable) {
      setLoading(false);
      return;
    }

    try {
      setLoading(true);
      setError(null);

      const res = await fetch('/api/users', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      });

      if (!res.ok) {
        throw new Error(`HTTP ${res.status}: ${res.statusText}`);
      }

      const userData = await res.json();
      console.log('User data received:', userData);
      const userArray = Array.isArray(userData)
        ? userData
        : userData.users || [];
      setUsers(userArray);
    } catch (error) {
      console.error('Failed to load users:', error);
      setError(`Failed to load users: ${error.message}`);
    } finally {
      setLoading(false);
    }
  }, [backendAvailable]);

  const reset = async id => {
    if (!window.confirm('Reset password for this user?')) return;

    try {
      const res = await fetch(`/api/users/${id}/reset`, {
        method: 'POST',
        credentials: 'include',
      });

      if (res.ok) {
        alert('Password reset and emailed');
        await loadUsers();
      } else {
        const errorText = await res.text();
        alert(`Reset failed: ${errorText}`);
      }
    } catch (error) {
      alert(`Reset failed: ${error.message}`);
    }
  };

  const openEditor = user => setEditor({ open: true, user });
  const closeEditor = () => setEditor({ open: false, user: null });

  const saveUser = async data => {
    try {
      const res = await fetch(
        data.id ? `/api/users/${data.id}` : '/api/users',
        {
          method: data.id ? 'PUT' : 'POST',
          headers: { 'Content-Type': 'application/json' },
          credentials: 'include',
          body: JSON.stringify(data),
        }
      );
      if (!res.ok) {
        const txt = await res.text();
        throw new Error(txt);
      }
      await loadUsers();
      closeEditor();
    } catch (error) {
      alert(`Failed to save user: ${error.message}`);
    }
  };

  useEffect(() => {
    loadUsers();
  }, [loadUsers]);

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" p={4}>
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        User Management
      </Typography>

      <Box mb={2}>
        <Button variant="contained" onClick={() => openEditor(null)}>
          Add User
        </Button>
      </Box>

      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available. User management features are
          currently disabled.
        </Alert>
      )}

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      {users.length === 0 && !loading && !error ? (
        <Alert severity="info">
          No users found. This might indicate a backend connectivity issue.
        </Alert>
      ) : (
        <Paper variant="outlined">
          <Table size="small">
            <TableHead>
              <TableRow>
                <TableCell>
                  <strong>Username</strong>
                </TableCell>
                <TableCell>
                  <strong>Email</strong>
                </TableCell>
                <TableCell>
                  <strong>Role</strong>
                </TableCell>
                <TableCell>
                  <strong>Status</strong>
                </TableCell>
                <TableCell>
                  <strong>Actions</strong>
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {users.map((user, index) => (
                <TableRow key={user.id || index}>
                  <TableCell>{user.username || user.name || 'N/A'}</TableCell>
                  <TableCell>{user.email || 'N/A'}</TableCell>
                  <TableCell>
                    <Chip
                      label={user.role || 'user'}
                      size="small"
                      color={user.role === 'admin' ? 'primary' : 'default'}
                    />
                  </TableCell>
                  <TableCell>
                    <Chip
                      label={user.active !== false ? 'Active' : 'Inactive'}
                      size="small"
                      color={user.active !== false ? 'success' : 'default'}
                    />
                  </TableCell>
                  <TableCell>
                    <Button
                      size="small"
                      onClick={() => openEditor(user)}
                      variant="text"
                      sx={{ mr: 1 }}
                    >
                      Edit
                    </Button>
                    <Button
                      size="small"
                      onClick={() => reset(user.id)}
                      disabled={!user.id}
                      variant="outlined"
                    >
                      Reset Password
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </Paper>
      )}

      {process.env.NODE_ENV === 'development' && (
        <Box mt={2}>
          <Typography variant="caption" color="text.secondary">
            Debug: {users.length} users loaded
          </Typography>
        </Box>
      )}

      <UserEditorDialog
        open={editor.open}
        user={editor.user}
        onClose={closeEditor}
        onSave={saveUser}
      />
    </Box>
  );
}
