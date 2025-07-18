import {
  Archive as ExtractIcon,
  Folder as FolderIcon,
  Movie as MediaIcon,
  Subtitles as SubtitleIcon,
} from '@mui/icons-material';
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  LinearProgress,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Paper,
  TextField,
  Typography,
} from '@mui/material';
import { useState } from 'react';
import { apiService } from './services/api.js';

/**
 * Extract provides a simple form to request subtitle extraction for a media file.
 * The path to the media file is POSTed to `/api/extract` and the number of
 * extracted items is displayed.
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend service is available
 */
export default function Extract({ backendAvailable = true }) {
  const [path, setPath] = useState('');
  const [status, setStatus] = useState('');
  const [extracting, setExtracting] = useState(false);
  const [extractedItems, setExtractedItems] = useState([]);

  const doExtract = async () => {
    if (!path.trim()) {
      setStatus('Please enter a valid path');
      return;
    }

    setExtracting(true);
    setStatus('');
    setExtractedItems([]);

    try {
      const response = await apiService.subtitles.extract(path);

      if (response.ok) {
        const items = await response.json();
        setExtractedItems(items || []);
        setStatus(
          `Successfully extracted ${items?.length || 0} subtitle streams`
        );
      } else {
        const errorText = await response.text();
        setStatus(`Error: ${errorText || 'Failed to extract subtitles'}`);
      }
    } catch (error) {
      setStatus(
        `Network error: ${error.message || 'Please check your connection.'}`
      );
    } finally {
      setExtracting(false);
    }
  };

  const handleKeyPress = event => {
    if (event.key === 'Enter') {
      doExtract();
    }
  };

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        Extract Subtitles
      </Typography>

      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available. Subtitle extraction is currently
          disabled.
        </Alert>
      )}

      <Typography variant="body1" color="text.secondary" paragraph>
        Extract embedded subtitle streams from media files (MKV, MP4, etc.) into
        separate subtitle files.
      </Typography>

      <Card sx={{ maxWidth: 800, mx: 'auto' }}>
        <CardContent>
          <Box sx={{ mb: 3 }}>
            <TextField
              fullWidth
              label="Media File Path"
              placeholder="/path/to/media/file.mkv"
              value={path}
              onChange={e => setPath(e.target.value)}
              onKeyPress={handleKeyPress}
              disabled={!backendAvailable || extracting}
              InputProps={{
                startAdornment: (
                  <FolderIcon sx={{ mr: 1, color: 'action.active' }} />
                ),
              }}
              helperText="Enter the full path to a media file containing embedded subtitles"
            />
          </Box>

          <Button
            variant="contained"
            startIcon={extracting ? <LinearProgress /> : <ExtractIcon />}
            onClick={doExtract}
            disabled={extracting || !path.trim() || !backendAvailable}
            size="large"
            fullWidth
          >
            {!backendAvailable
              ? 'Backend Unavailable'
              : extracting
                ? 'Extracting...'
                : 'Extract Subtitles'}
          </Button>

          {extracting && (
            <Box mt={2}>
              <LinearProgress />
              <Typography
                variant="body2"
                color="text.secondary"
                align="center"
                mt={1}
              >
                Analyzing media file and extracting subtitle streams...
              </Typography>
            </Box>
          )}

          {status && (
            <Alert
              severity={
                status.includes('Error') || status.includes('error')
                  ? 'error'
                  : 'success'
              }
              sx={{ mt: 2 }}
            >
              {status}
            </Alert>
          )}

          {extractedItems.length > 0 && (
            <Box mt={3}>
              <Typography variant="h6" gutterBottom>
                Extracted Subtitle Files
              </Typography>
              <Paper variant="outlined">
                <List>
                  {extractedItems.map((item, index) => (
                    <ListItem
                      key={index}
                      divider={index < extractedItems.length - 1}
                    >
                      <ListItemIcon>
                        <SubtitleIcon color="primary" />
                      </ListItemIcon>
                      <ListItemText
                        primary={item.filename || `Stream ${index + 1}`}
                        secondary={
                          <Box
                            display="flex"
                            alignItems="center"
                            gap={1}
                            mt={1}
                          >
                            <Chip
                              label={item.language || 'Unknown'}
                              size="small"
                              variant="outlined"
                            />
                            <Chip
                              label={item.codec || 'Unknown format'}
                              size="small"
                              color="secondary"
                              variant="outlined"
                            />
                            {item.title && (
                              <Chip
                                label={item.title}
                                size="small"
                                variant="outlined"
                              />
                            )}
                          </Box>
                        }
                      />
                    </ListItem>
                  ))}
                </List>
              </Paper>
            </Box>
          )}
        </CardContent>
      </Card>

      <Card sx={{ mt: 3, maxWidth: 800, mx: 'auto' }}>
        <CardContent>
          <Typography variant="h6" gutterBottom>
            <MediaIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
            Supported Media Formats
          </Typography>
          <Typography variant="body2" color="text.secondary">
            This tool can extract subtitles from media files that contain
            embedded subtitle streams, including MKV, MP4, AVI, and other
            container formats. The extracted subtitles will be saved as separate
            files in the same directory as the source media.
          </Typography>
        </CardContent>
      </Card>
    </Box>
  );
}
