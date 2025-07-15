// file: webui/src/components/DirectoryChooser.jsx
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
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (open) {
      setCurrentPath('/');
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
            data.items
              ?.filter(
                item => item.isDirectory || item.type === 'directory'
              )
              .map(item => item.path) || [];
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
              key={dir}
              onClick={() => setCurrentPath(dir)}
              onDoubleClick={() => handleSelect(dir)}
            >
              <ListItemIcon>
                <FolderIcon />
              </ListItemIcon>
              <ListItemText primary={dir.split('/').pop()} />
            </ListItemButton>
          ))}
        </List>
      </DialogContent>
      <DialogActions>
        <Button onClick={() => handleSelect(currentPath)} variant="contained">
          Select
        </Button>
      </DialogActions>
    </Dialog>
  );
}
