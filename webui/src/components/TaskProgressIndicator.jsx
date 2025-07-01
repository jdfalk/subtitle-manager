// file: webui/src/components/TaskProgressIndicator.jsx
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174002

import { Box, LinearProgress, Typography, Chip } from '@mui/material';

/**
 * TaskProgressIndicator displays the progress of a background task
 * @param {Object} props - Component props
 * @param {Object} props.task - Task object with id, status, progress, etc.
 * @param {boolean} props.showDetails - Whether to show detailed information
 */
export default function TaskProgressIndicator({ task, showDetails = true }) {
  if (!task) return null;

  const getStatusColor = status => {
    switch (status) {
      case 'completed':
        return 'success';
      case 'failed':
        return 'error';
      case 'running':
        return 'primary';
      default:
        return 'default';
    }
  };

  const getProgressValue = () => {
    if (task.status === 'completed') return 100;
    if (task.status === 'failed') return 0;
    return task.progress || 0;
  };

  return (
    <Box sx={{ width: '100%', mb: 2 }}>
      {showDetails && (
        <Box
          display="flex"
          alignItems="center"
          justifyContent="space-between"
          mb={1}
        >
          <Typography variant="body2" color="text.primary">
            {task.id}
          </Typography>
          <Chip
            label={task.status}
            size="small"
            color={getStatusColor(task.status)}
            variant="outlined"
          />
        </Box>
      )}

      <LinearProgress
        variant="determinate"
        value={getProgressValue()}
        sx={{
          height: 8,
          borderRadius: 1,
          backgroundColor: 'grey.200',
          '& .MuiLinearProgress-bar': {
            borderRadius: 1,
            backgroundColor:
              task.status === 'failed' ? 'error.main' : undefined,
          },
        }}
      />

      {showDetails && (
        <Box
          display="flex"
          alignItems="center"
          justifyContent="space-between"
          mt={1}
        >
          <Typography variant="caption" color="text.secondary">
            {getProgressValue().toFixed(0)}% complete
          </Typography>
          {task.error && (
            <Typography variant="caption" color="error">
              Error: {task.error}
            </Typography>
          )}
        </Box>
      )}
    </Box>
  );
}
