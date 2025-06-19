// file: webui/src/components/DirectoryChooser.jsx
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  CircularProgress,
} from '@mui/material';
import {
  Folder as FolderIcon,
  ArrowBack as BackIcon,
} from '@mui/icons-material';
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
              ?.filter(item => item.isDirectory)
              .map(item => item.path) || [];
          setDirs(directories);
        } else {
          setDirs([]);
        }
      } catch (err) {
        console.error('Failed to browse directory:', err);
        setDirs([]);
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
