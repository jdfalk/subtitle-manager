import {
    Delete as DeleteIcon,
    FilePresent as FileIcon,
    Language as LanguageIcon,
    Translate as TranslateIcon,
    CloudUpload as UploadIcon,
} from '@mui/icons-material';
import {
    Alert,
    Box,
    Button,
    Card,
    CardContent,
    Chip,
    FormControl,
    Grid,
    IconButton,
    InputLabel,
    LinearProgress,
    MenuItem,
    Paper,
    Select,
    Snackbar,
    Typography,
} from '@mui/material';
import { useState } from 'react';

/**
 * Translate provides a form to upload a subtitle file and request
 * translation to a target language via the /api/translate endpoint.
 * The translated file is downloaded by the browser.
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend service is available
 */
export default function Translate({ backendAvailable = true }) {
  const [file, setFile] = useState(null);
  const [lang, setLang] = useState('es');
  const [status, setStatus] = useState('');
  const [translating, setTranslating] = useState(false);
  const [snackbarOpen, setSnackbarOpen] = useState(false);

  const supportedLanguages = [
    { code: 'es', name: 'Spanish' },
    { code: 'fr', name: 'French' },
    { code: 'de', name: 'German' },
    { code: 'it', name: 'Italian' },
    { code: 'pt', name: 'Portuguese' },
    { code: 'ru', name: 'Russian' },
    { code: 'ja', name: 'Japanese' },
    { code: 'ko', name: 'Korean' },
    { code: 'zh', name: 'Chinese' },
    { code: 'ar', name: 'Arabic' },
    { code: 'hi', name: 'Hindi' },
    { code: 'nl', name: 'Dutch' },
    { code: 'sv', name: 'Swedish' },
    { code: 'da', name: 'Danish' },
    { code: 'no', name: 'Norwegian' },
    { code: 'fi', name: 'Finnish' },
    { code: 'pl', name: 'Polish' },
    { code: 'cs', name: 'Czech' },
    { code: 'tr', name: 'Turkish' },
  ];

  const doTranslate = async () => {
    if (!file) return;
    setTranslating(true);
    setStatus('');

    try {
      const form = new FormData();
      form.append('file', file);
      form.append('lang', lang);
      const res = await fetch('/api/translate', { method: 'POST', body: form });

      if (res.ok) {
        const blob = await res.blob();
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        const selectedLang = supportedLanguages.find(l => l.code === lang);
        a.download = file.name.replace(/\.[^/.]+$/, `_${lang}.srt`);
        a.click();
        window.URL.revokeObjectURL(url);
        setStatus(
          `File translated to ${selectedLang?.name || lang} and downloaded successfully!`
        );
        setSnackbarOpen(true);
      } else {
        const errorText = await res.text();
        setStatus(`Translation failed: ${errorText || 'Unknown error'}`);
      }
    } catch (error) {
      setStatus(
        `Network error: ${error.message || 'Please check your connection.'}`
      );
    } finally {
      setTranslating(false);
    }
  };

  const handleFileChange = event => {
    const selectedFile = event.target.files[0];
    setFile(selectedFile);
    setStatus('');
  };

  const handleRemoveFile = () => {
    setFile(null);
    setStatus('');
    const fileInput = document.getElementById('translate-file-input');
    if (fileInput) fileInput.value = '';
  };

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        Translate Subtitle
      </Typography>

      {!backendAvailable && (
        <Alert severity="error" sx={{ mb: 3 }}>
          Backend service is not available. Translation features are currently disabled.
        </Alert>
      )}

      <Typography variant="body1" color="text.secondary" paragraph>
        Upload a subtitle file and translate it to your target language using AI
        translation services.
      </Typography>

      <Grid container spacing={3} justifyContent="center">
        <Grid size={{ xs: 12, md: 8 }}>
          <Card>
            <CardContent>
              <Box textAlign="center" p={2}>
                {!file ? (
                  <Box>
                    <input
                      accept=".srt,.vtt,.ass,.ssa,.sub"
                      style={{ display: 'none' }}
                      id="translate-file-input"
                      type="file"
                      onChange={handleFileChange}
                      data-testid="file"
                    />
                    <label htmlFor="translate-file-input">
                      <Paper
                        variant="outlined"
                        sx={{
                          p: 4,
                          cursor: 'pointer',
                          border: '2px dashed',
                          borderColor: 'primary.main',
                          '&:hover': {
                            backgroundColor: 'action.hover',
                          },
                        }}
                      >
                        <UploadIcon
                          sx={{ fontSize: 48, color: 'primary.main', mb: 2 }}
                        />
                        <Typography variant="h6" gutterBottom>
                          Upload subtitle file for translation
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          Supports SRT, VTT, ASS, SSA, and SUB formats
                        </Typography>
                      </Paper>
                    </label>
                  </Box>
                ) : (
                  <Box>
                    <Paper variant="outlined" sx={{ p: 3, mb: 3 }}>
                      <Box
                        display="flex"
                        alignItems="center"
                        justifyContent="space-between"
                      >
                        <Box display="flex" alignItems="center">
                          <FileIcon sx={{ mr: 2, color: 'primary.main' }} />
                          <Box>
                            <Typography variant="body1" fontWeight="medium">
                              {file.name}
                            </Typography>
                            <Typography variant="body2" color="text.secondary">
                              {(file.size / 1024).toFixed(1)} KB
                            </Typography>
                          </Box>
                        </Box>
                        <IconButton onClick={handleRemoveFile} color="error">
                          <DeleteIcon />
                        </IconButton>
                      </Box>
                    </Paper>

                    <FormControl fullWidth sx={{ mb: 3 }}>
                      <InputLabel>Target Language</InputLabel>
                      <Select
                        value={lang}
                        label="Target Language"
                        onChange={e => setLang(e.target.value)}
                        disabled={translating}
                        startAdornment={
                          <LanguageIcon
                            sx={{ mr: 1, color: 'action.active' }}
                          />
                        }
                      >
                        {supportedLanguages.map(language => (
                          <MenuItem key={language.code} value={language.code}>
                            <Box display="flex" alignItems="center">
                              <Chip
                                label={language.code.toUpperCase()}
                                size="small"
                                sx={{ mr: 1, minWidth: 40 }}
                              />
                              {language.name}
                            </Box>
                          </MenuItem>
                        ))}
                      </Select>
                    </FormControl>

                    <Button
                      variant="contained"
                      startIcon={
                        translating ? <LinearProgress /> : <TranslateIcon />
                      }
                      onClick={doTranslate}
                      disabled={translating}
                      size="large"
                      fullWidth
                    >
                      {translating
                        ? 'Translating...'
                        : `Translate to ${supportedLanguages.find(l => l.code === lang)?.name || lang}`}
                    </Button>
                  </Box>
                )}

                {translating && (
                  <Box mt={2}>
                    <LinearProgress />
                    <Typography variant="body2" color="text.secondary" mt={1}>
                      Translating your subtitle file. This may take a few
                      moments...
                    </Typography>
                  </Box>
                )}

                {status && (
                  <Alert
                    severity={
                      status.includes('failed') || status.includes('error')
                        ? 'error'
                        : 'success'
                    }
                    sx={{ mt: 2 }}
                  >
                    {status}
                  </Alert>
                )}
              </Box>
            </CardContent>
          </Card>
        </Grid>

        <Grid size={{ xs: 12, md: 4 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                <LanguageIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
                Translation Info
              </Typography>
              <Typography variant="body2" color="text.secondary" paragraph>
                Our AI-powered translation service supports over 15 languages
                with high accuracy for subtitle content.
              </Typography>
              <Typography variant="body2" color="text.secondary" paragraph>
                The translation preserves timing information and formatting
                while providing natural, context-aware translations.
              </Typography>
              <Box mt={2}>
                <Typography variant="body2" fontWeight="medium" gutterBottom>
                  Supported Languages:
                </Typography>
                <Box display="flex" flexWrap="wrap" gap={0.5}>
                  {supportedLanguages.slice(0, 8).map(language => (
                    <Chip
                      key={language.code}
                      label={language.code.toUpperCase()}
                      size="small"
                      variant="outlined"
                    />
                  ))}
                  <Chip
                    label={`+${supportedLanguages.length - 8} more`}
                    size="small"
                    variant="outlined"
                    color="primary"
                  />
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      <Snackbar
        open={snackbarOpen}
        autoHideDuration={6000}
        onClose={() => setSnackbarOpen(false)}
      >
        <Alert onClose={() => setSnackbarOpen(false)} severity="success">
          {status}
        </Alert>
      </Snackbar>
    </Box>
  );
}
