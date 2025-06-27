// file: webui/src/MediaDetails.jsx
import { Box, CircularProgress, Typography } from '@mui/material';
import { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';

/**
 * MediaDetails displays information about a selected movie or TV series.
 * It fetches basic metadata from the OMDb API using the provided title.
 */
export default function MediaDetails() {
  const [params] = useSearchParams();
  const title = params.get('title');
  const [info, setInfo] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

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

  if (!info) {
    return (
      <Box p={4}>
        <Typography variant="h5">No details available</Typography>
      </Box>
    );
  }

  return (
    <Box p={4}>
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
        <Typography variant="body2">IMDB Rating: {info.imdbRating}</Typography>
      )}
    </Box>
  );
}
