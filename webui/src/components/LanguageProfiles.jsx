// file: webui/src/components/LanguageProfiles.jsx
// version: 1.0.0
// guid: 2c1d0e9f-8a7b-3c4d-6e5f-9a8b0c1d2e3f

import {
  Add as AddIcon,
  Delete as DeleteIcon,
  Edit as EditIcon,
  Star as DefaultIcon,
  StarBorder as NotDefaultIcon,
} from '@mui/icons-material';
import {
  Alert,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Chip,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  FormControl,
  FormControlLabel,
  Grid,
  IconButton,
  InputLabel,
  MenuItem,
  Select,
  Snackbar,
  Switch,
  TextField,
  Tooltip,
  Typography,
} from '@mui/material';
import { useEffect, useState } from 'react';
import { apiService } from '../services/api.js';

/**
 * Language Profiles Management Component
 * Allows users to create, edit, and manage language profiles similar to Bazarr.
 * @param {Object} props - Component props
 * @param {boolean} props.backendAvailable - Whether the backend service is available
 */
export default function LanguageProfiles({ backendAvailable = true }) {
  const [profiles, setProfiles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [dialogOpen, setDialogOpen] = useState(false);
  const [editingProfile, setEditingProfile] = useState(null);
  const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' });

  // Form state for profile editing
  const [formData, setFormData] = useState({
    name: '',
    cutoff_score: 75,
    languages: [{ language: 'en', priority: 1, forced: false, hi: false }],
  });

  // Common language codes
  const languageCodes = [
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
    { code: 'ar', name: 'Arabic' },
    { code: 'hi', name: 'Hindi' },
    { code: 'nl', name: 'Dutch' },
    { code: 'sv', name: 'Swedish' },
    { code: 'no', name: 'Norwegian' },
    { code: 'da', name: 'Danish' },
    { code: 'fi', name: 'Finnish' },
    { code: 'pl', name: 'Polish' },
    { code: 'cs', name: 'Czech' },
    { code: 'tr', name: 'Turkish' },
  ];

  /**
   * Load language profiles from the API
   */
  const loadProfiles = async () => {
    try {
      setLoading(true);
      const response = await apiService.get('/api/profiles');
      if (response.ok) {
        const data = await response.json();
        setProfiles(data || []);
      } else {
        throw new Error('Failed to load profiles');
      }
    } catch (err) {
      setError(err.message || 'Failed to load language profiles');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (backendAvailable) {
      loadProfiles();
    }
  }, [backendAvailable]);

  /**
   * Open dialog for creating new profile
   */
  const handleCreateProfile = () => {
    setEditingProfile(null);
    setFormData({
      name: '',
      cutoff_score: 75,
      languages: [{ language: 'en', priority: 1, forced: false, hi: false }],
    });
    setDialogOpen(true);
  };

  /**
   * Open dialog for editing existing profile
   */
  const handleEditProfile = (profile) => {
    setEditingProfile(profile);
    setFormData({
      name: profile.name,
      cutoff_score: profile.cutoff_score,
      languages: profile.languages || [],
    });
    setDialogOpen(true);
  };

  /**
   * Save profile (create or update)
   */
  const handleSaveProfile = async () => {
    try {
      const profileData = {
        ...formData,
        languages: formData.languages.map((lang, index) => ({
          ...lang,
          priority: index + 1, // Auto-assign priorities based on order
        })),
      };

      let response;
      if (editingProfile) {
        response = await apiService.put(`/api/profiles/${editingProfile.id}`, profileData);
      } else {
        response = await apiService.post('/api/profiles', profileData);
      }

      if (response.ok) {
        setSnackbar({
          open: true,
          message: `Profile ${editingProfile ? 'updated' : 'created'} successfully`,
          severity: 'success',
        });
        setDialogOpen(false);
        loadProfiles();
      } else {
        const errorData = await response.text();
        throw new Error(errorData || `Failed to ${editingProfile ? 'update' : 'create'} profile`);
      }
    } catch (err) {
      setSnackbar({
        open: true,
        message: err.message || 'Failed to save profile',
        severity: 'error',
      });
    }
  };

  /**
   * Delete a profile
   */
  const handleDeleteProfile = async (profileId) => {
    if (!confirm('Are you sure you want to delete this profile?')) {
      return;
    }

    try {
      const response = await apiService.delete(`/api/profiles/${profileId}`);
      if (response.ok) {
        setSnackbar({
          open: true,
          message: 'Profile deleted successfully',
          severity: 'success',
        });
        loadProfiles();
      } else {
        const errorData = await response.text();
        throw new Error(errorData || 'Failed to delete profile');
      }
    } catch (err) {
      setSnackbar({
        open: true,
        message: err.message || 'Failed to delete profile',
        severity: 'error',
      });
    }
  };

  /**
   * Set a profile as default
   */
  const handleSetDefault = async (profileId) => {
    try {
      const response = await apiService.post(`/api/profiles/${profileId}/default`);
      if (response.ok) {
        setSnackbar({
          open: true,
          message: 'Default profile updated successfully',
          severity: 'success',
        });
        loadProfiles();
      } else {
        throw new Error('Failed to set default profile');
      }
    } catch (err) {
      setSnackbar({
        open: true,
        message: err.message || 'Failed to set default profile',
        severity: 'error',
      });
    }
  };

  /**
   * Add a new language to the form
   */
  const handleAddLanguage = () => {
    setFormData({
      ...formData,
      languages: [
        ...formData.languages,
        { language: 'en', priority: formData.languages.length + 1, forced: false, hi: false },
      ],
    });
  };

  /**
   * Remove a language from the form
   */
  const handleRemoveLanguage = (index) => {
    const newLanguages = formData.languages.filter((_, i) => i !== index);
    setFormData({ ...formData, languages: newLanguages });
  };

  /**
   * Update a language in the form
   */
  const handleLanguageChange = (index, field, value) => {
    const newLanguages = [...formData.languages];
    newLanguages[index] = { ...newLanguages[index], [field]: value };
    setFormData({ ...formData, languages: newLanguages });
  };

  /**
   * Get language name from code
   */
  const getLanguageName = (code) => {
    const lang = languageCodes.find(l => l.code === code);
    return lang ? lang.name : code.toUpperCase();
  };

  if (!backendAvailable) {
    return (
      <Alert severity="warning">
        Backend service is not available. Language profiles cannot be managed at this time.
      </Alert>
    );
  }

  if (loading) {
    return <Typography>Loading language profiles...</Typography>;
  }

  if (error) {
    return <Alert severity="error">{error}</Alert>;
  }

  return (
    <Box>
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={3}>
        <Typography variant="h4" component="h1">
          Language Profiles
        </Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={handleCreateProfile}
          color="primary"
        >
          Create Profile
        </Button>
      </Box>

      <Typography variant="body1" color="text.secondary" mb={3}>
        Language profiles allow you to define preferred languages and quality thresholds for subtitle downloads.
        Each profile can contain multiple languages with priority ordering.
      </Typography>

      <Grid container spacing={3}>
        {profiles.map((profile) => (
          <Grid item xs={12} md={6} lg={4} key={profile.id}>
            <Card>
              <CardContent>
                <Box display="flex" alignItems="center" mb={2}>
                  <Typography variant="h6" component="h2" sx={{ flexGrow: 1 }}>
                    {profile.name}
                  </Typography>
                  <Tooltip title={profile.is_default ? 'Default Profile' : 'Not Default'}>
                    <IconButton
                      onClick={() => !profile.is_default && handleSetDefault(profile.id)}
                      disabled={profile.is_default}
                      color={profile.is_default ? 'primary' : 'default'}
                    >
                      {profile.is_default ? <DefaultIcon /> : <NotDefaultIcon />}
                    </IconButton>
                  </Tooltip>
                </Box>

                <Typography variant="body2" color="text.secondary" mb={2}>
                  Cutoff Score: {profile.cutoff_score}%
                </Typography>

                <Box mb={2}>
                  <Typography variant="subtitle2" mb={1}>Languages:</Typography>
                  <Box display="flex" flexWrap="wrap" gap={1}>
                    {(profile.languages || []).map((lang, index) => (
                      <Chip
                        key={index}
                        label={`${getLanguageName(lang.language)} (${lang.priority})`}
                        size="small"
                        color={lang.forced || lang.hi ? 'secondary' : 'default'}
                        variant={lang.forced || lang.hi ? 'filled' : 'outlined'}
                      />
                    ))}
                  </Box>
                </Box>
              </CardContent>

              <CardActions>
                <Button
                  size="small"
                  startIcon={<EditIcon />}
                  onClick={() => handleEditProfile(profile)}
                >
                  Edit
                </Button>
                <Button
                  size="small"
                  startIcon={<DeleteIcon />}
                  onClick={() => handleDeleteProfile(profile.id)}
                  disabled={profile.is_default}
                  color="error"
                >
                  Delete
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}

        {profiles.length === 0 && (
          <Grid item xs={12}>
            <Alert severity="info">
              No language profiles found. Create your first profile to get started.
            </Alert>
          </Grid>
        )}
      </Grid>

      {/* Profile Editor Dialog */}
      <Dialog open={dialogOpen} onClose={() => setDialogOpen(false)} maxWidth="md" fullWidth>
        <DialogTitle>
          {editingProfile ? 'Edit Language Profile' : 'Create Language Profile'}
        </DialogTitle>
        <DialogContent>
          <Box sx={{ pt: 1 }}>
            <TextField
              fullWidth
              label="Profile Name"
              value={formData.name}
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
              margin="normal"
              required
            />

            <TextField
              fullWidth
              label="Cutoff Score (%)"
              type="number"
              value={formData.cutoff_score}
              onChange={(e) => setFormData({ ...formData, cutoff_score: parseInt(e.target.value) })}
              margin="normal"
              inputProps={{ min: 0, max: 100 }}
              helperText="Minimum score required for subtitle downloads"
            />

            <Typography variant="h6" mt={3} mb={2}>
              Languages
            </Typography>

            {formData.languages.map((lang, index) => (
              <Box key={index} sx={{ border: 1, borderColor: 'divider', borderRadius: 1, p: 2, mb: 2 }}>
                <Grid container spacing={2} alignItems="center">
                  <Grid item xs={12} sm={4}>
                    <FormControl fullWidth>
                      <InputLabel>Language</InputLabel>
                      <Select
                        value={lang.language}
                        label="Language"
                        onChange={(e) => handleLanguageChange(index, 'language', e.target.value)}
                      >
                        {languageCodes.map((langCode) => (
                          <MenuItem key={langCode.code} value={langCode.code}>
                            {langCode.name} ({langCode.code})
                          </MenuItem>
                        ))}
                      </Select>
                    </FormControl>
                  </Grid>

                  <Grid item xs={12} sm={3}>
                    <FormControlLabel
                      control={
                        <Switch
                          checked={lang.forced}
                          onChange={(e) => handleLanguageChange(index, 'forced', e.target.checked)}
                        />
                      }
                      label="Forced"
                    />
                  </Grid>

                  <Grid item xs={12} sm={3}>
                    <FormControlLabel
                      control={
                        <Switch
                          checked={lang.hi}
                          onChange={(e) => handleLanguageChange(index, 'hi', e.target.checked)}
                        />
                      }
                      label="HI"
                    />
                  </Grid>

                  <Grid item xs={12} sm={2}>
                    <IconButton
                      onClick={() => handleRemoveLanguage(index)}
                      disabled={formData.languages.length === 1}
                      color="error"
                    >
                      <DeleteIcon />
                    </IconButton>
                  </Grid>
                </Grid>
              </Box>
            ))}

            <Button
              variant="outlined"
              startIcon={<AddIcon />}
              onClick={handleAddLanguage}
              sx={{ mt: 1 }}
            >
              Add Language
            </Button>
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDialogOpen(false)}>Cancel</Button>
          <Button
            onClick={handleSaveProfile}
            variant="contained"
            disabled={!formData.name || formData.languages.length === 0}
          >
            {editingProfile ? 'Update' : 'Create'}
          </Button>
        </DialogActions>
      </Dialog>

      {/* Snackbar for notifications */}
      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={() => setSnackbar({ ...snackbar, open: false })}
      >
        <Alert
          onClose={() => setSnackbar({ ...snackbar, open: false })}
          severity={snackbar.severity}
          sx={{ width: '100%' }}
        >
          {snackbar.message}
        </Alert>
      </Snackbar>
    </Box>
  );
}