// file: webui/src/MediaDetails.jsx
import {
  Box,
  CircularProgress,
  Typography,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Alert,
  Snackbar,
  Button,
  Card,
  CardContent,
  Chip,
  Grid,
} from '@mui/material';
import { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import { apiService } from './services/api.js';

/**
 * MediaDetails displays information about a selected movie or TV series.
 * It fetches basic metadata from the OMDb API using the provided title
 * and allows managing language profile assignments for the media item.
 */
export default function MediaDetails() {
  const [params] = useSearchParams();
  const title = params.get('title');
  const mediaPath = params.get('path'); // Use file path as media identifier
  const [info, setInfo] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  // Language profile management
  const [profiles, setProfiles] = useState([]);
  const [assignedProfile, setAssignedProfile] = useState(null);
  const [profileLoading, setProfileLoading] = useState(false);
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });

  // Load language profiles
  const loadProfiles = async () => {
    try {
      const response = await apiService.get('/api/profiles');
      if (response.ok) {
        const data = await response.json();
        setProfiles(data || []);
      }
    } catch (err) {
      console.error('Failed to load profiles:', err);
    }
  };

  // Load assigned profile for media item
  const loadAssignedProfile = async () => {
    if (!mediaPath) return;

    try {
      setProfileLoading(true);
      // URL encode the path to handle special characters and slashes
      const encodedPath = encodeURIComponent(mediaPath);
      const response = await apiService.get(
        `/api/media/profile/${encodedPath}`
      );
      if (response.ok) {
        const profile = await response.json();
        setAssignedProfile(profile);
      }
    } catch (err) {
      console.error('Failed to load assigned profile:', err);
    } finally {
      setProfileLoading(false);
    }
  };

  // Assign profile to media item
  const handleProfileChange = async profileId => {
    if (!mediaPath) return;

    try {
      setProfileLoading(true);
      // URL encode the path to handle special characters and slashes
      const encodedPath = encodeURIComponent(mediaPath);

      if (profileId === '') {
        // Remove profile assignment
        const response = await apiService.delete(
          `/api/media/profile/${encodedPath}`
        );
        if (response.ok) {
          setSnackbar({
            open: true,
            message: 'Language profile removed successfully',
            severity: 'success',
          });
          loadAssignedProfile(); // Reload to get default profile
        } else {
          throw new Error('Failed to remove profile assignment');
        }
      } else {
        // Assign new profile
        const response = await apiService.put(
          `/api/media/profile/${encodedPath}`,
          {
            profile_id: profileId,
          }
        );
        if (response.ok) {
          setSnackbar({
            open: true,
            message: 'Language profile updated successfully',
            severity: 'success',
          });
          loadAssignedProfile(); // Reload to show new assignment
        } else {
          throw new Error('Failed to assign profile');
        }
      }
    } catch (err) {
      setSnackbar({
        open: true,
        message: err.message || 'Failed to update language profile',
        severity: 'error',
      });
    } finally {
      setProfileLoading(false);
    }
  };

  useEffect(() => {
    const fetchInfo = async () => {
      if (!title) {
        setLoading(false);
        return;
      }
      try {
        const res = await fetch(
          `https://www.omdbapi.com/?t=${encodeURIComponent(title)}&apikey=thewdb`
        );
        if (res.ok) {
          const data = await res.json();
          if (data.Response === 'True') {
            setInfo(data);
          }
        }
      } catch {
        setError('Failed to fetch media details');
      } finally {
        setLoading(false);
      }
    };
    fetchInfo();
  }, [title]);

  useEffect(() => {
    loadProfiles();
    loadAssignedProfile();
  }, [mediaPath]);

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" mt={4}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box p={4}>
        <Typography variant="h5" color="error">
          {error}
        </Typography>
      </Box>
    );
  }

  // Get language name from code
  const getLanguageName = code => {
    const languageCodes = {
      en: 'English',
      es: 'Spanish',
      fr: 'French',
      de: 'German',
      it: 'Italian',
      pt: 'Portuguese',
      ru: 'Russian',
      ja: 'Japanese',
      ko: 'Korean',
      zh: 'Chinese',
      ar: 'Arabic',
      hi: 'Hindi',
      nl: 'Dutch',
      sv: 'Swedish',
      no: 'Norwegian',
      da: 'Danish',
      fi: 'Finnish',
      pl: 'Polish',
      cs: 'Czech',
      tr: 'Turkish',
    };
    return languageCodes[code] || code.toUpperCase();
  };

  return (
    <Box p={4}>
      <Grid container spacing={3}>
        {/* Media Information */}
        <Grid item xs={12} md={8}>
          {info ? (
            <>
              <Typography variant="h4" gutterBottom>
                {info.Title}
              </Typography>
              {info.Poster && info.Poster !== 'N/A' && (
                <Box
                  component="img"
                  src={info.Poster}
                  alt={info.Title}
                  sx={{ maxWidth: 300, mb: 2 }}
                />
              )}
              <Typography variant="body1" paragraph>
                {info.Plot}
              </Typography>
              {info.imdbRating && (
                <Typography variant="body2">
                  IMDB Rating: {info.imdbRating}
                </Typography>
              )}
            </>
          ) : title ? (
            <Typography variant="h4" gutterBottom>
              {title}
            </Typography>
          ) : (
            <Typography variant="h5">No details available</Typography>
          )}
        </Grid>

        {/* Language Profile Management */}
        {mediaPath && (
          <Grid item xs={12} md={4}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  Language Preferences
                </Typography>

                <Typography variant="body2" color="text.secondary" gutterBottom>
                  File: {mediaPath.split('/').pop()}
                </Typography>

                {profileLoading ? (
                  <Box display="flex" justifyContent="center" p={2}>
                    <CircularProgress size={24} />
                  </Box>
                ) : (
                  <>
                    {assignedProfile && (
                      <Box mb={2}>
                        <Typography variant="subtitle2" gutterBottom>
                          Current Profile: {assignedProfile.name}
                          {assignedProfile.is_default && (
                            <Chip
                              label="Default"
                              size="small"
                              color="primary"
                              sx={{ ml: 1 }}
                            />
                          )}
                        </Typography>

                        <Typography
                          variant="body2"
                          color="text.secondary"
                          gutterBottom
                        >
                          Cutoff Score: {assignedProfile.cutoff_score}%
                        </Typography>

                        <Box display="flex" flexWrap="wrap" gap={0.5} mt={1}>
                          {assignedProfile.languages?.map((lang, index) => (
                            <Chip
                              key={index}
                              label={`${getLanguageName(lang.language)} (${lang.priority})`}
                              size="small"
                              color={
                                lang.forced || lang.hi ? 'secondary' : 'default'
                              }
                              variant={
                                lang.forced || lang.hi ? 'filled' : 'outlined'
                              }
                            />
                          ))}
                        </Box>
                      </Box>
                    )}

                    <FormControl fullWidth size="small">
                      <InputLabel>Language Profile</InputLabel>
                      <Select
                        value={assignedProfile?.id || ''}
                        label="Language Profile"
                        onChange={e => handleProfileChange(e.target.value)}
                        disabled={profileLoading}
                      >
                        <MenuItem value="">
                          <em>Use Default Profile</em>
                        </MenuItem>
                        {profiles.map(profile => (
                          <MenuItem key={profile.id} value={profile.id}>
                            {profile.name}
                            {profile.is_default && ' (Default)'}
                          </MenuItem>
                        ))}
                      </Select>
                    </FormControl>

                    <Typography
                      variant="caption"
                      display="block"
                      mt={1}
                      color="text.secondary"
                    >
                      Select a language profile to customize subtitle
                      preferences for this title.
                    </Typography>
                  </>
                )}
              </CardContent>
            </Card>
          </Grid>
        )}
      </Grid>

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
