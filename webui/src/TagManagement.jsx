// file: webui/src/TagManagement.jsx
import { useCallback, useEffect, useState } from 'react';
import {
  Alert,
  Box,
  Button,
  CircularProgress,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TextField,
  Typography,
} from '@mui/material';

/**
 * TagManagement displays existing tags and allows creation/deletion.
 *
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend is reachable
 */
export default function TagManagement({ backendAvailable = true }) {
  const [tags, setTags] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [newTag, setNewTag] = useState('');
  const [editId, setEditId] = useState(null);
  const [editName, setEditName] = useState('');

  const loadTags = useCallback(async () => {
    if (!backendAvailable) {
      setLoading(false);
      return;
    }
    try {
      setLoading(true);
      setError(null);
      const res = await fetch('/api/tags');
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const data = await res.json();
      setTags(Array.isArray(data) ? data : []);
    } catch (err) {
      setError(`Failed to load tags: ${err.message}`);
    } finally {
      setLoading(false);
    }
  }, [backendAvailable]);

  const addTag = async () => {
    if (!newTag) return;
    try {
      const res = await fetch('/api/tags', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name: newTag }),
      });
      if (res.ok) {
        setNewTag('');
        loadTags();
      }
    } catch (err) {
      setError(`Failed to add tag: ${err.message}`);
    }
  };

  const startEdit = tag => {
    setEditId(tag.id);
    setEditName(tag.name);
  };

  const saveEdit = async () => {
    try {
      const res = await fetch(`/api/tags/${editId}`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name: editName }),
      });
      if (res.ok) {
        setEditId(null);
        setEditName('');
        loadTags();
      }
    } catch (err) {
      setError(`Failed to update tag: ${err.message}`);
    }
  };

  const cancelEdit = () => {
    setEditId(null);
    setEditName('');
  };

  const deleteTag = async id => {
    if (!window.confirm('Delete this tag?')) return;
    try {
      const res = await fetch(`/api/tags/${id}`, { method: 'DELETE' });
      if (res.ok) loadTags();
    } catch (err) {
      setError(`Failed to delete tag: ${err.message}`);
    }
  };

  useEffect(() => {
    loadTags();
  }, [loadTags]);

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
        Tag Management
      </Typography>
      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available.
        </Alert>
      )}
      {error && (
        <Alert severity="error" sx={{ mb: 3 }} onClose={() => setError(null)}>
          {error}
        </Alert>
      )}
      <Box display="flex" gap={2} mb={2}>
        <TextField
          label="New Tag"
          value={newTag}
          onChange={e => setNewTag(e.target.value)}
          size="small"
        />
        <Button variant="contained" onClick={addTag} disabled={!newTag}>
          Add
        </Button>
      </Box>
      <Paper variant="outlined">
        <Table size="small">
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {tags.map(tag => (
              <TableRow key={tag.id}>
                <TableCell>
                  {editId === tag.id ? (
                    <TextField
                      value={editName}
                      onChange={e => setEditName(e.target.value)}
                      size="small"
                    />
                  ) : (
                    tag.name
                  )}
                </TableCell>
                <TableCell>
                  {editId === tag.id ? (
                    <>
                      <Button size="small" onClick={saveEdit}>
                        Save
                      </Button>
                      <Button size="small" onClick={cancelEdit}>
                        Cancel
                      </Button>
                    </>
                  ) : (
                    <>
                      <Button size="small" onClick={() => startEdit(tag)}>
                        Edit
                      </Button>
                      <Button size="small" onClick={() => deleteTag(tag.id)}>
                        Delete
                      </Button>
                    </>
                  )}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Paper>
    </Box>
  );
}
