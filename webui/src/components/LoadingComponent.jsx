// file: webui/src/components/LoadingComponent.jsx

import { Box, CircularProgress, Typography } from '@mui/material';

/**
 * Reusable loading component for lazy-loaded components
 * Provides a consistent loading experience across the app
 * @param {Object} props - Component props
 * @param {string} props.message - Optional loading message
 * @param {string} props.size - Size of the progress indicator (small, medium, large)
 */
export default function LoadingComponent({
  message = 'Loading...',
  size = 'medium',
}) {
  const sizeMap = {
    small: 24,
    medium: 40,
    large: 60,
  };

  return (
    <Box
      display="flex"
      flexDirection="column"
      justifyContent="center"
      alignItems="center"
      minHeight="200px"
      gap={2}
    >
      <CircularProgress size={sizeMap[size]} />
      <Typography variant="body2" color="text.secondary">
        {message}
      </Typography>
    </Box>
  );
}
