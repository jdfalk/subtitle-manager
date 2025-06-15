// file: webui/src/MediaLibrary.jsx

import {
  Download as DownloadIcon,
  Archive as ExtractIcon,
  Folder as FolderIcon,
  Info as InfoIcon,
  MoreVert as MoreIcon,
  Movie as MovieIcon,
  Refresh as RefreshIcon,
  CloudDownload as SearchIcon,
  Subtitles as SubtitleIcon,
  Translate as TranslateIcon,
  Tv as TvIcon,
} from '@mui/icons-material';
import {
  Box,
  Breadcrumbs,
  Button,
  Card,
  CardContent,
  Chip,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  Grid,
  IconButton,
  LinearProgress,
  Link,
  Menu,
  MenuItem,
  Paper,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';

/**
 * MediaLibrary provides integrated media file and subtitle management.
 * Shows media files with their available subtitles, allows searching,
 * downloading, extracting, and translating subtitles directly from the file view.
 */
export default function MediaLibrary() {
  const [currentPath, setCurrentPath] = useState('/');
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);
  const [actionMenu, setActionMenu] = useState({ anchor: null, file: null });
  const [bulkMode, setBulkMode] = useState(false);
  const [selectedFiles, setSelectedFiles] = useState(new Set());
  const [operationDialog, setOperationDialog] = useState({
    open: false,
    type: null,
    file: null,
  });
  const [progress, setProgress] = useState(null);

  const loadCurrentDirectory = async () => {
    setLoading(true);
    try {
      const response = await fetch(
        `/api/library/browse?path=${encodeURIComponent(currentPath)}`
      );
      if (response.ok) {
        const data = await response.json();
        setItems(data.items || []);
      }
    } catch (error) {
      console.error('Failed to load directory:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadCurrentDirectory();
  }, [currentPath]); // eslint-disable-line react-hooks/exhaustive-deps

  /**
   * Navigate to a subdirectory
   */
  const navigateToPath = path => {
    setCurrentPath(path);
    setSelectedFiles(new Set());
  };

  /**
   * Get breadcrumb items from current path
   */
  const getBreadcrumbs = () => {
    const parts = currentPath.split('/').filter(Boolean);
    const breadcrumbs = [{ name: 'Root', path: '/' }];

    let currentBreadcrumbPath = '';
    parts.forEach(part => {
      currentBreadcrumbPath += '/' + part;
      breadcrumbs.push({ name: part, path: currentBreadcrumbPath });
    });

    return breadcrumbs;
  };

  /**
   * Get file type icon
   */
  const getFileIcon = item => {
    if (item.type === 'directory') return <FolderIcon />;
    if (item.isVideo) return item.isTvShow ? <TvIcon /> : <MovieIcon />;
    if (item.isSubtitle) return <SubtitleIcon />;
    return <InfoIcon />;
  };

  /**
   * Handle file action menu
   */
  const handleActionMenu = (event, file) => {
    event.stopPropagation();
    setActionMenu({ anchor: event.currentTarget, file });
  };

  const closeActionMenu = () => {
    setActionMenu({ anchor: null, file: null });
  };

  /**
   * Handle file operations
   */
  const handleOperation = async (type, file) => {
    closeActionMenu();
    setOperationDialog({ open: true, type, file });
  };

  const executeOperation = async () => {
    const { type, file } = operationDialog;
    setProgress({ type, file: file.name, progress: 0 });

    try {
      let endpoint = '';
      let body = {};

      switch (type) {
        case 'extract':
          endpoint = '/api/extract';
          body = { path: file.path };
          break;
        case 'search':
          endpoint = '/api/search';
          body = { path: file.path, language: 'en' };
          break;
        case 'translate':
          endpoint = '/api/translate';
          body = { path: file.path, targetLanguage: 'es' };
          break;
        default:
          return;
      }

      const response = await fetch(endpoint, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
      });

      if (response.ok) {
        await loadCurrentDirectory(); // Refresh to show new files
      }
    } catch (error) {
      console.error(`Failed to ${type}:`, error);
    } finally {
      setProgress(null);
      setOperationDialog({ open: false, type: null, file: null });
    }
  };

  /**
   * Handle bulk operations
   */
  const handleBulkOperation = async type => {
    if (selectedFiles.size === 0) return;

    const files = Array.from(selectedFiles)
      .map(path => items.find(item => item.path === path))
      .filter(Boolean);

    setProgress({ type, file: `${files.length} files`, progress: 0 });

    try {
      await fetch('/api/bulk-operation', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          type,
          files: files.map(f => f.path),
          language: 'en',
        }),
      });

      await loadCurrentDirectory();
      setSelectedFiles(new Set());
    } catch (error) {
      console.error(`Failed bulk ${type}:`, error);
    } finally {
      setProgress(null);
    }
  };

  /**
   * Toggle file selection for bulk operations
   */
  const toggleFileSelection = filePath => {
    const newSelection = new Set(selectedFiles);
    if (newSelection.has(filePath)) {
      newSelection.delete(filePath);
    } else {
      newSelection.add(filePath);
    }
    setSelectedFiles(newSelection);
  };

  if (loading && items.length === 0) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="400px"
      >
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box>
      {/* Header */}
      <Box
        display="flex"
        justifyContent="space-between"
        alignItems="center"
        mb={3}
      >
        <Typography variant="h4" component="h1">
          Media Library
        </Typography>
        <Box>
          <Button
            variant={bulkMode ? 'contained' : 'outlined'}
            onClick={() => {
              setBulkMode(!bulkMode);
              setSelectedFiles(new Set());
            }}
            sx={{ mr: 1 }}
          >
            {bulkMode ? 'Exit Bulk Mode' : 'Bulk Operations'}
          </Button>
          <IconButton onClick={loadCurrentDirectory}>
            <RefreshIcon />
          </IconButton>
        </Box>
      </Box>

      {/* Breadcrumbs */}
      <Paper sx={{ p: 2, mb: 3 }}>
        <Breadcrumbs>
          {getBreadcrumbs().map((crumb, index) => (
            <Link
              key={index}
              component="button"
              variant="body1"
              onClick={() => navigateToPath(crumb.path)}
              sx={{ textDecoration: 'none' }}
            >
              {crumb.name}
            </Link>
          ))}
        </Breadcrumbs>
      </Paper>

      {/* Bulk Operations Bar */}
      {bulkMode && selectedFiles.size > 0 && (
        <Paper sx={{ p: 2, mb: 3, backgroundColor: 'action.selected' }}>
          <Box display="flex" alignItems="center" gap={2}>
            <Typography variant="body1">
              {selectedFiles.size} files selected
            </Typography>
            <Button
              startIcon={<ExtractIcon />}
              onClick={() => handleBulkOperation('extract')}
              size="small"
            >
              Extract Subtitles
            </Button>
            <Button
              startIcon={<SearchIcon />}
              onClick={() => handleBulkOperation('search')}
              size="small"
            >
              Search Subtitles
            </Button>
            <Button
              startIcon={<TranslateIcon />}
              onClick={() => handleBulkOperation('translate')}
              size="small"
            >
              Translate
            </Button>
          </Box>
        </Paper>
      )}

      {/* Progress Indicator */}
      {progress && (
        <Paper sx={{ p: 2, mb: 3 }}>
          <Typography variant="body2" gutterBottom>
            {progress.type.charAt(0).toUpperCase() + progress.type.slice(1)}ing:{' '}
            {progress.file}
          </Typography>
          <LinearProgress />
        </Paper>
      )}

      {/* File List */}
      <Grid container spacing={3}>
        {items.map(item => (
          <Grid item xs={12} key={item.path}>
            <Card
              sx={{
                cursor: item.type === 'directory' ? 'pointer' : 'default',
                border: selectedFiles.has(item.path)
                  ? '2px solid'
                  : '1px solid',
                borderColor: selectedFiles.has(item.path)
                  ? 'primary.main'
                  : 'divider',
              }}
              onClick={() => {
                if (bulkMode && item.type !== 'directory') {
                  toggleFileSelection(item.path);
                } else if (item.type === 'directory') {
                  navigateToPath(item.path);
                }
                // Individual file selection removed - use bulk mode for operations
              }}
            >
              <CardContent>
                <Box display="flex" alignItems="center">
                  <Box sx={{ mr: 2 }}>{getFileIcon(item)}</Box>

                  <Box flex={1}>
                    <Typography variant="h6" noWrap>
                      {item.name}
                    </Typography>

                    {item.isVideo && (
                      <Box display="flex" flexWrap="wrap" gap={1} mt={1}>
                        {item.subtitles?.map(sub => (
                          <Chip
                            key={sub.language}
                            label={`${sub.language.toUpperCase()} ${sub.format || 'SRT'}`}
                            size="small"
                            color={sub.embedded ? 'warning' : 'success'}
                            icon={<SubtitleIcon />}
                          />
                        ))}
                        {(!item.subtitles || item.subtitles.length === 0) && (
                          <Chip
                            label="No Subtitles"
                            size="small"
                            color="error"
                            variant="outlined"
                          />
                        )}
                      </Box>
                    )}

                    {item.metadata && (
                      <Typography
                        variant="body2"
                        color="text.secondary"
                        sx={{ mt: 1 }}
                      >
                        {item.metadata.resolution &&
                          `${item.metadata.resolution} • `}
                        {item.metadata.duration &&
                          `${item.metadata.duration} • `}
                        {item.size &&
                          `${(item.size / 1024 / 1024 / 1024).toFixed(1)} GB`}
                      </Typography>
                    )}
                  </Box>

                  {item.type !== 'directory' && (
                    <IconButton
                      onClick={e => handleActionMenu(e, item)}
                      size="small"
                    >
                      <MoreIcon />
                    </IconButton>
                  )}
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>

      {items.length === 0 && !loading && (
        <Paper sx={{ p: 4, textAlign: 'center' }}>
          <Typography variant="h6" color="text.secondary">
            No files found in this directory
          </Typography>
        </Paper>
      )}

      {/* Action Menu */}
      <Menu
        anchorEl={actionMenu.anchor}
        open={Boolean(actionMenu.anchor)}
        onClose={closeActionMenu}
      >
        {actionMenu.file?.isVideo && [
          <MenuItem
            key="extract"
            onClick={() => handleOperation('extract', actionMenu.file)}
          >
            <ExtractIcon sx={{ mr: 1 }} />
            Extract Embedded Subtitles
          </MenuItem>,
          <MenuItem
            key="search"
            onClick={() => handleOperation('search', actionMenu.file)}
          >
            <SearchIcon sx={{ mr: 1 }} />
            Search for Subtitles
          </MenuItem>,
          <Divider key="divider" />,
        ]}
        {actionMenu.file?.isSubtitle && (
          <MenuItem
            onClick={() => handleOperation('translate', actionMenu.file)}
          >
            <TranslateIcon sx={{ mr: 1 }} />
            Translate Subtitle
          </MenuItem>
        )}
        <MenuItem
          onClick={() =>
            window.open(
              `/api/download?path=${encodeURIComponent(actionMenu.file?.path)}`,
              '_blank'
            )
          }
        >
          <DownloadIcon sx={{ mr: 1 }} />
          Download File
        </MenuItem>
      </Menu>

      {/* Operation Dialog */}
      <Dialog
        open={operationDialog.open}
        onClose={() =>
          setOperationDialog({ open: false, type: null, file: null })
        }
      >
        <DialogTitle>
          {operationDialog.type &&
            `${operationDialog.type.charAt(0).toUpperCase() + operationDialog.type.slice(1)} Subtitles`}
        </DialogTitle>
        <DialogContent>
          <Typography>
            {operationDialog.type === 'extract' &&
              'Extract embedded subtitle streams from this video file?'}
            {operationDialog.type === 'search' &&
              'Search for subtitles for this video file using enabled providers?'}
            {operationDialog.type === 'translate' &&
              'Translate this subtitle file to another language?'}
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() =>
              setOperationDialog({ open: false, type: null, file: null })
            }
          >
            Cancel
          </Button>
          <Button onClick={executeOperation} variant="contained">
            {operationDialog.type &&
              operationDialog.type.charAt(0).toUpperCase() +
                operationDialog.type.slice(1)}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
