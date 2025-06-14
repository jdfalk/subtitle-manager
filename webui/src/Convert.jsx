import {
  Transform as ConvertIcon,
  Delete as DeleteIcon,
  FilePresent as FileIcon,
  CloudUpload as UploadIcon
} from "@mui/icons-material";
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  IconButton,
  LinearProgress,
  Paper,
  Snackbar,
  Typography,
} from "@mui/material";
import { useState } from "react";

/**
 * Convert provides a form to upload a subtitle file which is
 * converted to SRT format via the /api/convert endpoint.
 * The resulting file is downloaded by the browser.
 */
export default function Convert() {
  const [file, setFile] = useState(null);
  const [status, setStatus] = useState("");
  const [converting, setConverting] = useState(false);
  const [snackbarOpen, setSnackbarOpen] = useState(false);

  const doConvert = async () => {
    if (!file) return;
    setConverting(true);
    setStatus("");

    try {
      const form = new FormData();
      form.append("file", file);
      const res = await fetch("/api/convert", { method: "POST", body: form });

      if (res.ok) {
        const blob = await res.blob();
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = file.name.replace(/\.[^/.]+$/, ".srt");
        a.click();
        window.URL.revokeObjectURL(url);
        setStatus("File converted and downloaded successfully!");
        setSnackbarOpen(true);
      } else {
        setStatus("Error converting file. Please try again.");
      }
    } catch {
      setStatus("Network error. Please check your connection.");
    } finally {
      setConverting(false);
    }
  };

  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    setFile(selectedFile);
    setStatus("");
  };

  const handleRemoveFile = () => {
    setFile(null);
    setStatus("");
    // Reset the input element
    const fileInput = document.getElementById('subtitle-file-input');
    if (fileInput) fileInput.value = '';
  };

  const getSupportedFormats = () => [
    'VTT', 'ASS', 'SSA', 'SUB', 'SBV', 'TTML', 'DFXP'
  ];

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        Convert Subtitle
      </Typography>

      <Typography variant="body1" color="text.secondary" paragraph>
        Upload a subtitle file to convert it to SRT format. Supported formats include VTT, ASS, SSA, SUB, SBV, TTML, and DFXP.
      </Typography>

      <Card sx={{ maxWidth: 600, mx: 'auto' }}>
        <CardContent>
          <Box textAlign="center" p={3}>
            {!file ? (
              <Box>
                <input
                  accept=".vtt,.ass,.ssa,.sub,.sbv,.ttml,.dfxp"
                  style={{ display: 'none' }}
                  id="subtitle-file-input"
                  type="file"
                  onChange={handleFileChange}
                  data-testid="file"
                />
                <label htmlFor="subtitle-file-input">
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
                    <UploadIcon sx={{ fontSize: 48, color: 'primary.main', mb: 2 }} />
                    <Typography variant="h6" gutterBottom>
                      Click to upload subtitle file
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      Drag and drop or click to browse
                    </Typography>
                  </Paper>
                </label>
              </Box>
            ) : (
              <Box>
                <Paper variant="outlined" sx={{ p: 3, mb: 2 }}>
                  <Box display="flex" alignItems="center" justifyContent="space-between">
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

                <Button
                  variant="contained"
                  startIcon={converting ? <LinearProgress /> : <ConvertIcon />}
                  onClick={doConvert}
                  disabled={converting}
                  size="large"
                  fullWidth
                >
                  {converting ? 'Converting...' : 'Convert to SRT'}
                </Button>
              </Box>
            )}

            {converting && (
              <Box mt={2}>
                <LinearProgress />
                <Typography variant="body2" color="text.secondary" mt={1}>
                  Converting your subtitle file...
                </Typography>
              </Box>
            )}

            {status && (
              <Alert
                severity={status.includes('Error') || status.includes('error') ? 'error' : 'success'}
                sx={{ mt: 2 }}
              >
                {status}
              </Alert>
            )}
          </Box>
        </CardContent>
      </Card>

      <Card sx={{ mt: 3, maxWidth: 600, mx: 'auto' }}>
        <CardContent>
          <Typography variant="h6" gutterBottom>
            Supported Formats
          </Typography>
          <Box display="flex" flexWrap="wrap" gap={1}>
            {getSupportedFormats().map((format) => (
              <Chip key={format} label={format} variant="outlined" size="small" />
            ))}
          </Box>
        </CardContent>
      </Card>

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
