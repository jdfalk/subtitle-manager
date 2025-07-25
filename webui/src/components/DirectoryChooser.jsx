// file: webui/src/components/DirectoryChooser.jsx
// version: 1.0.2
// guid: d1836bb0-aa02-4af1-ad3d-9f5bb2948975
import {
  ArrowBack as BackIcon,
  Folder as FolderIcon,
} from '@mui/icons-material';
import {
  Button,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  TextField,
} from '@mui/material';
import { useEffect, useState } from 'react';
import { apiService } from '../services/api.js';

/**
 * DirectoryChooser provides a simple dialog to browse and select
 * directories on the server's file system using the library browse API.
 *
 * @param {boolean} open - Whether the dialog is open
 * @param {function} onClose - Callback when the dialog closes
 * @param {function} onSelect - Callback with the chosen directory path
 */
export default function DirectoryChooser({ open, onClose, onSelect }) {
  const [currentPath, setCurrentPath] = useState('/');
  const [dirs, setDirs] = useState([]);
  const [customPath, setCustomPath] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (open) {
      setCurrentPath('/');
      setCustomPath('');
    }
  }, [open]);

  useEffect(() => {
    if (!open) return;
    const load = async () => {
      setLoading(true);
      try {
        const resp = await apiService.get(
          `/api/library/browse?path=${encodeURIComponent(currentPath)}`
        );
        if (resp.ok) {
          const data = await resp.json();
          const directories =
            data.items?.filter(
              item => item.isDirectory || item.type === 'directory'
            ) || [];
          setDirs(directories);
        } else {
          setDirs([]);
        }
      } catch {
        setError('Failed to open directory');
      } finally {
        setLoading(false);
      }
    };
    load();
  }, [currentPath, open]);

  const goUp = () => {
    if (currentPath === '/' || currentPath === '') return;
    const parts = currentPath.split('/').filter(Boolean);
    parts.pop();
    setCurrentPath('/' + parts.join('/'));
  };

  const handleSelect = path => {
    onSelect(path);
    onClose();
  };

  return (
    <Dialog open={open} onClose={onClose} fullWidth maxWidth="sm">
      <DialogTitle>Select Directory</DialogTitle>
      <DialogContent dividers>
        {loading && <CircularProgress size={20} sx={{ mb: 2 }} />}
        {error && <div style={{ color: 'red' }}>{error}</div>}
        <List>
          {currentPath !== '/' && (
            <ListItemButton onClick={goUp}>
              <ListItemIcon>
                <BackIcon />
              </ListItemIcon>
              <ListItemText primary=".." />
            </ListItemButton>
          )}
          {dirs.map(dir => (
            <ListItemButton
              key={dir.path}
              onClick={() => setCurrentPath(dir.path)}
              onDoubleClick={() => handleSelect(dir.path)}
            >
              <ListItemIcon>
                <FolderIcon />
              </ListItemIcon>
              <ListItemText
                primary={
                  dir.name || dir.path.replace(/\\/g, '/').split('/').pop()
                }
              />
            </ListItemButton>
          ))}
        </List>
        <TextField
          fullWidth
          label="Directory Path"
          value={customPath}
          onChange={e => setCustomPath(e.target.value)}
          margin="normal"
          placeholder="Enter path manually"
        />
      </DialogContent>
      <DialogActions>
        <Button
          onClick={() => handleSelect(customPath || currentPath)}
          variant="contained"
        >
          {customPath ? 'Add' : 'Select'}
        </Button>
      </DialogActions>
    </Dialog>
  );
}
