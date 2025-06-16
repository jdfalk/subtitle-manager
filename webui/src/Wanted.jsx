import {
    Add as AddIcon,
    Delete as DeleteIcon,
    Language as LanguageIcon,
    Movie as MediaIcon,
    CloudDownload as ProviderIcon,
    Search as SearchIcon,
    Download as WantedIcon,
} from '@mui/icons-material';
import {
    Alert,
    Box,
    Button,
    Card,
    CardContent,
    Chip,
    CircularProgress,
    FormControl,
    Grid,
    IconButton,
    InputLabel,
    List,
    ListItem,
    ListItemSecondaryAction,
    ListItemText,
    MenuItem,
    Paper,
    Select,
    TextField,
    Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';

/**
 * Wanted provides an interface for searching subtitles and maintaining
 * a list of wanted items. Results are fetched from `/api/search` and
 * selections are POSTed to `/api/wanted`.
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend service is available
 */
export default function Wanted({ backendAvailable = true }) {
  // Start with embedded provider available out of the box
  const [provider, setProvider] = useState('embedded');
  const [path, setPath] = useState('');
  const [lang, setLang] = useState('en');
  const [results, setResults] = useState([]);
  const [wanted, setWanted] = useState([]);
  const [searching, setSearching] = useState(false);
  const [loading, setLoading] = useState(true);

  const providers = [
    { value: 'generic', label: 'Generic' },
    { value: 'opensubtitles', label: 'OpenSubtitles' },
    { value: 'addic7ed', label: 'Addic7ed' },
    { value: 'podnapisi', label: 'Podnapisi' },
    { value: 'subscene', label: 'Subscene' },
  ];

  const languages = [
    { code: 'en', name: 'English' },
    { code: 'es', name: 'Spanish' },
    { code: 'fr', name: 'French' },
    { code: 'de', name: 'German' },
    { code: 'it', name: 'Italian' },
    { code: 'pt', name: 'Portuguese' },
    { code: 'ru', name: 'Russian' },
    { code: 'ja', name: 'Japanese' },
    { code: 'ko', name: 'Korean' },
    { code: 'zh', name: 'Chinese' },
    { code: 'zh-cn', name: 'Chinese (Simplified)' },
    { code: 'zh-hans', name: 'Chinese (Simplified)' },
    { code: 'zh-tw', name: 'Chinese (Taiwan)' },
    { code: 'zh-hk', name: 'Chinese (Hong Kong)' },
    { code: 'zh-hant', name: 'Chinese (Traditional)' },
  ];

  useEffect(() => {
    const loadWanted = async () => {
      setLoading(true);
      try {
        const res = await fetch('/api/wanted');
        if (res.ok) {
          const data = await res.json();
          setWanted(data || []);
        }
      } catch (error) {
        console.error('Failed to load wanted list:', error);
      } finally {
        setLoading(false);
      }
    };
    loadWanted();
  }, []);

  const search = async () => {
    if (!path.trim()) {
      return;
    }

    setSearching(true);
    setResults([]);

    try {
      const params = new URLSearchParams({ provider, path, lang });
      const res = await fetch(`/api/search?${params.toString()}`);
      if (res.ok) {
        const data = await res.json();
        setResults(data || []);
      }
    } catch (error) {
      console.error('Search failed:', error);
    } finally {
      setSearching(false);
    }
  };

  const add = async url => {
    try {
      const res = await fetch('/api/wanted', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ url }),
      });
      if (res.ok) {
        setWanted([...wanted, url]);
      }
    } catch (error) {
      console.error('Failed to add to wanted list:', error);
    }
  };

  const remove = async url => {
    try {
      const res = await fetch('/api/wanted', {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ url }),
      });
      if (res.ok) {
        setWanted(wanted.filter(item => item !== url));
      }
    } catch (error) {
      console.error('Failed to remove from wanted list:', error);
    }
  };

  const handleKeyPress = event => {
    if (event.key === 'Enter') {
      search();
    }
  };

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        Subtitle Search & Wanted List
      </Typography>

      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available. Search and wanted list features are currently disabled.
        </Alert>
      )}

      <Typography variant="body1" color="text.secondary" paragraph>
        Search for subtitles across multiple providers and maintain a wanted
        list for automated downloads.
      </Typography>

      <Grid container spacing={3}>
        {/* Search Form */}
        <Grid size={{ xs: 12 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                <SearchIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
                Search Subtitles
              </Typography>
              <Grid container spacing={2} alignItems="center">
                <Grid size={{ xs: 12, sm: 3 }}>
                  <FormControl fullWidth>
                    <InputLabel>Provider</InputLabel>
                    <Select
                      value={provider}
                      label="Provider"
                      onChange={e => setProvider(e.target.value)}
                      disabled={searching}
                      startAdornment={
                        <ProviderIcon sx={{ mr: 1, color: 'action.active' }} />
                      }
                    >
                      {providers.map(p => (
                        <MenuItem key={p.value} value={p.value}>
                          {p.label}
                        </MenuItem>
                      ))}
                    </Select>
                  </FormControl>
                </Grid>
                <Grid size={{ xs: 12, sm: 5 }}>
                  <TextField
                    fullWidth
                    label="Media File Path"
                    placeholder="/path/to/movie.mkv"
                    value={path}
                    onChange={e => setPath(e.target.value)}
                    onKeyPress={handleKeyPress}
                    disabled={searching}
                    InputProps={{
                      startAdornment: (
                        <MediaIcon sx={{ mr: 1, color: 'action.active' }} />
                      ),
                    }}
                  />
                </Grid>
                <Grid size={{ xs: 12, sm: 2 }}>
                  <FormControl fullWidth>
                    <InputLabel>Language</InputLabel>
                    <Select
                      value={lang}
                      label="Language"
                      onChange={e => setLang(e.target.value)}
                      disabled={searching}
                      startAdornment={
                        <LanguageIcon sx={{ mr: 1, color: 'action.active' }} />
                      }
                    >
                      {languages.map(l => (
                        <MenuItem key={l.code} value={l.code}>
                          <Box display="flex" alignItems="center">
                            <Chip
                              label={l.code.toUpperCase()}
                              size="small"
                              sx={{ mr: 1, minWidth: 40 }}
                            />
                            {l.name}
                          </Box>
                        </MenuItem>
                      ))}
                    </Select>
                  </FormControl>
                </Grid>
                <Grid size={{ xs: 12, sm: 2 }}>
                  <Button
                    variant="contained"
                    startIcon={
                      searching ? (
                        <CircularProgress size={20} />
                      ) : (
                        <SearchIcon />
                      )
                    }
                    onClick={search}
                    disabled={searching || !path.trim()}
                    fullWidth
                    size="large"
                  >
                    {searching ? 'Searching...' : 'Search'}
                  </Button>
                </Grid>
              </Grid>
            </CardContent>
          </Card>
        </Grid>

        {/* Search Results */}
        <Grid size={{ xs: 12, md: 8 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Search Results ({results.length})
              </Typography>
              {searching ? (
                <Box display="flex" justifyContent="center" p={3}>
                  <CircularProgress />
                </Box>
              ) : results.length === 0 ? (
                <Alert severity="info">
                  No search results. Try searching for a media file to find
                  subtitles.
                </Alert>
              ) : (
                <Paper
                  variant="outlined"
                  sx={{ maxHeight: 400, overflow: 'auto' }}
                >
                  <List>
                    {results.map((url, index) => (
                      <ListItem
                        key={index}
                        divider={index < results.length - 1}
                      >
                        <ListItemText
                          primary={
                            <Typography
                              variant="body2"
                              sx={{
                                fontFamily: 'monospace',
                                wordBreak: 'break-all',
                              }}
                            >
                              {url}
                            </Typography>
                          }
                          secondary={
                            <Chip
                              label={provider}
                              size="small"
                              variant="outlined"
                              sx={{ mt: 1 }}
                            />
                          }
                        />
                        <ListItemSecondaryAction>
                          <Button
                            variant="outlined"
                            size="small"
                            startIcon={<AddIcon />}
                            onClick={() => add(url)}
                            disabled={wanted.includes(url)}
                          >
                            {wanted.includes(url) ? 'Added' : 'Add to Wanted'}
                          </Button>
                        </ListItemSecondaryAction>
                      </ListItem>
                    ))}
                  </List>
                </Paper>
              )}
            </CardContent>
          </Card>
        </Grid>

        {/* Wanted List */}
        <Grid size={{ xs: 12, md: 4 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                <WantedIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
                Wanted List ({wanted.length})
              </Typography>
              {loading ? (
                <Box display="flex" justifyContent="center" p={2}>
                  <CircularProgress size={24} />
                </Box>
              ) : wanted.length === 0 ? (
                <Alert severity="info">
                  No items in wanted list. Add subtitles from search results.
                </Alert>
              ) : (
                <Paper
                  variant="outlined"
                  sx={{ maxHeight: 400, overflow: 'auto' }}
                >
                  <List dense>
                    {wanted.map((url, index) => (
                      <ListItem key={index} divider={index < wanted.length - 1}>
                        <ListItemText
                          primary={
                            <Typography
                              variant="body2"
                              sx={{
                                fontFamily: 'monospace',
                                fontSize: '0.75rem',
                                wordBreak: 'break-all',
                              }}
                            >
                              {url}
                            </Typography>
                          }
                        />
                        <ListItemSecondaryAction>
                          <IconButton
                            edge="end"
                            size="small"
                            color="error"
                            onClick={() => remove(url)}
                          >
                            <DeleteIcon fontSize="small" />
                          </IconButton>
                        </ListItemSecondaryAction>
                      </ListItem>
                    ))}
                  </List>
                </Paper>
              )}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Box>
  );
}
