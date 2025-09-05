<!-- file: docs/tasks/02-ui-fixes/TASK-02-004-dashboard-enhancements.md -->
<!-- version: 1.0.0 -->
<!-- guid: c2d3e4f5-g6h7-i8j9-k0l1-m2n3o4p5q6r7 -->

# TASK-02-004: Dashboard Enhancements

## 🎯 Objective

Enhance the existing dashboard with modern widgets, better visual organization, real-time statistics, and improved user experience for monitoring and managing subtitle operations.

## 📋 Acceptance Criteria

- [ ] Add statistics cards showing key metrics (files processed, success rate, etc.)
- [ ] Implement real-time activity feed with recent operations
- [ ] Create overview widgets for different service statuses
- [ ] Add charts for subtitle download statistics and trends
- [ ] Implement quick action buttons for common operations (Sonarr/Radarr sync)
- [ ] Improve task monitoring with better visual indicators
- [ ] Add system health indicators and alerts
- [ ] Implement responsive grid layout for better organization

## 🔍 Current State Analysis

### Existing Dashboard Features (Good Foundation)

**Current Implementation** (`webui/src/Dashboard.jsx` - 614 lines):

1. **✅ Scan Controls**: Directory selection, provider selection, language settings
2. **✅ Task Monitoring**: Progress tracking with TaskProgressIndicator
3. **✅ Quick Links**: Navigation buttons to main sections
4. **✅ System Info**: Basic system information display
5. **✅ File Listing**: Shows files found during scans
6. **✅ Backend Availability**: Proper error handling for offline state

### Areas for Enhancement

1. **Visual Organization**: Current layout is functional but could be more visually appealing
2. **Statistics**: Missing key metrics and insights
3. **Activity Monitoring**: No recent activity overview
4. **Service Status**: No overview of Sonarr/Radarr/Plex integrations
5. **Performance Metrics**: No charts or trend analysis
6. **Quick Actions**: Limited one-click operations

## 🎨 Enhanced Dashboard Design

### Layout Structure

```
┌─────────────────────────────────────────────────────────────────┐
│ Dashboard                                          🔔 Alerts     │
├─────────────────────────────────────────────────────────────────┤
│ ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ │
│ │   Total     │ │  Success    │ │ Downloaded  │ │   Active    │ │
│ │   Files     │ │    Rate     │ │   Today     │ │    Tasks    │ │
│ │   1,234     │ │    98.5%    │ │     256     │ │      3      │ │
│ └─────────────┘ └─────────────┘ └─────────────┘ └─────────────┘ │
├─────────────────────────────────────────────────────────────────┤
│ ┌─── Quick Actions ─────────────┐ ┌─── Service Status ───────────┐ │
│ │ [🔄 Sync Sonarr] [🎬 Sync Radarr] │ │ Sonarr    ✅ Connected       │ │
│ │ [📂 Scan Library] [🔍 Search]     │ │ Radarr    ✅ Connected       │ │
│ │ [⚙️ Settings] [📊 Statistics]     │ │ Plex      ⚠️ Warning         │ │
│ └───────────────────────────────┘ └─────────────────────────────┘ │
├─────────────────────────────────────────────────────────────────┤
│ ┌─── Recent Activity ───────────┐ ┌─── Active Tasks ─────────────┐ │
│ │ 14:32 Downloaded "Movie.srt"  │ │ [████████░░] Scanning        │ │
│ │ 14:30 Synced with Sonarr      │ │ Progress: 80% (2.5m left)    │ │
│ │ 14:25 User "admin" logged in  │ │                              │ │
│ │ 14:20 Provider OpenSubs OK    │ │ [██████████] Completed       │ │
│ └───────────────────────────────┘ │ Download: TV Show.srt        │ │
│                                   └─────────────────────────────┘ │
│ ┌─── Current Scan Controls ────────────────────────────────────┐ │
│ │ [Original scan interface - enhanced styling]                │ │
│ └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## 🔧 Implementation Steps

### Step 1: Create Statistics Cards Component

```jsx
// webui/src/components/Dashboard/StatisticsCards.jsx
import React, { useEffect, useState } from 'react';
import {
  Card,
  CardContent,
  Typography,
  Grid,
  Chip,
  Skeleton,
  Box
} from '@mui/material';
import {
  TrendingUp as TrendIcon,
  CheckCircle as SuccessIcon,
  CloudDownload as DownloadIcon,
  PlayCircleOutline as ActiveIcon
} from '@mui/icons-material';
import { apiService } from '../../services/api.js';

const StatCard = ({ title, value, icon, color = 'primary', trend, loading }) => (
  <Card>
    <CardContent>
      <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
        <Box>
          <Typography color="textSecondary" gutterBottom variant="body2">
            {title}
          </Typography>
          {loading ? (
            <Skeleton width={60} height={32} />
          ) : (
            <Typography variant="h4" component="div" color={color}>
              {value}
            </Typography>
          )}
          {trend && !loading && (
            <Chip
              icon={<TrendIcon fontSize="small" />}
              label={trend}
              size="small"
              color={trend.startsWith('+') ? 'success' : 'error'}
              variant="outlined"
            />
          )}
        </Box>
        <Box sx={{ color: `${color}.main`, opacity: 0.7 }}>
          {icon}
        </Box>
      </Box>
    </CardContent>
  </Card>
);

export default function StatisticsCards() {
  const [stats, setStats] = useState({
    totalFiles: 0,
    successRate: 0,
    downloadedToday: 0,
    activeTasks: 0,
    trends: {}
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadStatistics();
    const interval = setInterval(loadStatistics, 30000); // Update every 30 seconds
    return () => clearInterval(interval);
  }, []);

  const loadStatistics = async () => {
    try {
      const response = await apiService.get('/api/statistics/dashboard');
      if (response.ok) {
        const data = await response.json();
        setStats(data);
      }
    } catch (error) {
      console.error('Failed to load statistics:', error);
    } finally {
      setLoading(false);
    }
  };

  const cardData = [
    {
      title: 'Total Files',
      value: stats.totalFiles.toLocaleString(),
      icon: <DownloadIcon fontSize="large" />,
      color: 'primary',
      trend: stats.trends.totalFiles
    },
    {
      title: 'Success Rate',
      value: `${stats.successRate}%`,
      icon: <SuccessIcon fontSize="large" />,
      color: stats.successRate > 90 ? 'success' : stats.successRate > 70 ? 'warning' : 'error',
      trend: stats.trends.successRate
    },
    {
      title: 'Downloaded Today',
      value: stats.downloadedToday.toLocaleString(),
      icon: <TrendIcon fontSize="large" />,
      color: 'info',
      trend: stats.trends.downloadedToday
    },
    {
      title: 'Active Tasks',
      value: stats.activeTasks.toString(),
      icon: <ActiveIcon fontSize="large" />,
      color: stats.activeTasks > 0 ? 'warning' : 'success',
      trend: stats.trends.activeTasks
    }
  ];

  return (
    <Grid container spacing={3} sx={{ mb: 3 }}>
      {cardData.map((card, index) => (
        <Grid item xs={12} sm={6} md={3} key={index}>
          <StatCard {...card} loading={loading} />
        </Grid>
      ))}
    </Grid>
  );
}
```

### Step 2: Create Quick Actions Component

```jsx
// webui/src/components/Dashboard/QuickActions.jsx
import React, { useState } from 'react';
import {
  Card,
  CardContent,
  Typography,
  Grid,
  Button,
  Alert,
  CircularProgress,
  Box
} from '@mui/material';
import {
  Sync as SyncIcon,
  Movie as MovieIcon,
  Tv as TvIcon,
  FolderOpen as ScanIcon,
  Search as SearchIcon,
  Settings as SettingsIcon,
  Assessment as StatsIcon
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { apiService } from '../../services/api.js';

export default function QuickActions() {
  const [loading, setLoading] = useState({});
  const [status, setStatus] = useState('');
  const navigate = useNavigate();

  const handleAction = async (actionType, endpoint) => {
    setLoading(prev => ({ ...prev, [actionType]: true }));
    try {
      const response = await apiService.post(endpoint);
      if (response.ok) {
        setStatus(`${actionType} started successfully`);
      } else {
        setStatus(`Failed to start ${actionType}`);
      }
    } catch (error) {
      console.error(`Failed to execute ${actionType}:`, error);
      setStatus(`Error: ${error.message}`);
    } finally {
      setLoading(prev => ({ ...prev, [actionType]: false }));
    }
  };

  const actions = [
    {
      label: 'Sync Sonarr',
      icon: <TvIcon />,
      color: 'primary',
      action: () => handleAction('Sonarr Sync', '/api/sonarr/sync'),
      disabled: loading.sonarr
    },
    {
      label: 'Sync Radarr',
      icon: <MovieIcon />,
      color: 'primary',
      action: () => handleAction('Radarr Sync', '/api/radarr/sync'),
      disabled: loading.radarr
    },
    {
      label: 'Scan Library',
      icon: <ScanIcon />,
      color: 'secondary',
      action: () => navigate('/scan'),
      disabled: false
    },
    {
      label: 'Search Subtitles',
      icon: <SearchIcon />,
      color: 'secondary',
      action: () => navigate('/search'),
      disabled: false
    },
    {
      label: 'Settings',
      icon: <SettingsIcon />,
      color: 'default',
      action: () => navigate('/settings'),
      disabled: false
    },
    {
      label: 'Statistics',
      icon: <StatsIcon />,
      color: 'default',
      action: () => navigate('/statistics'),
      disabled: false
    }
  ];

  return (
    <Card>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Quick Actions
        </Typography>

        {status && (
          <Alert severity="info" sx={{ mb: 2 }} onClose={() => setStatus('')}>
            {status}
          </Alert>
        )}

        <Grid container spacing={2}>
          {actions.map((action, index) => (
            <Grid item xs={6} sm={4} md={2} key={index}>
              <Button
                fullWidth
                variant="outlined"
                startIcon={
                  action.disabled ? <CircularProgress size={20} /> : action.icon
                }
                onClick={action.action}
                disabled={action.disabled}
                color={action.color}
              >
                {action.label}
              </Button>
            </Grid>
          ))}
        </Grid>
      </CardContent>
    </Card>
  );
}
```

### Step 3: Create Service Status Component

```jsx
// webui/src/components/Dashboard/ServiceStatus.jsx
import React, { useEffect, useState } from 'react';
import {
  Card,
  CardContent,
  Typography,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Chip,
  Box,
  IconButton,
  Tooltip
} from '@mui/material';
import {
  CheckCircle as SuccessIcon,
  Warning as WarningIcon,
  Error as ErrorIcon,
  Refresh as RefreshIcon,
  Settings as SettingsIcon
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { apiService } from '../../services/api.js';

export default function ServiceStatus() {
  const [services, setServices] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    loadServiceStatus();
    const interval = setInterval(loadServiceStatus, 60000); // Update every minute
    return () => clearInterval(interval);
  }, []);

  const loadServiceStatus = async () => {
    try {
      setLoading(true);
      const response = await apiService.get('/api/services/status');
      if (response.ok) {
        const data = await response.json();
        setServices(data);
      }
    } catch (error) {
      console.error('Failed to load service status:', error);
    } finally {
      setLoading(false);
    }
  };

  const getStatusIcon = (status) => {
    switch (status) {
      case 'connected':
        return <SuccessIcon color="success" />;
      case 'warning':
        return <WarningIcon color="warning" />;
      case 'error':
        return <ErrorIcon color="error" />;
      default:
        return <ErrorIcon color="disabled" />;
    }
  };

  const getStatusColor = (status) => {
    switch (status) {
      case 'connected':
        return 'success';
      case 'warning':
        return 'warning';
      case 'error':
        return 'error';
      default:
        return 'default';
    }
  };

  return (
    <Card>
      <CardContent>
        <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', mb: 2 }}>
          <Typography variant="h6">
            Service Status
          </Typography>
          <Tooltip title="Refresh Status">
            <IconButton onClick={loadServiceStatus} size="small">
              <RefreshIcon />
            </IconButton>
          </Tooltip>
        </Box>

        <List dense>
          {services.map((service, index) => (
            <ListItem
              key={index}
              secondaryAction={
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                  <Chip
                    label={service.status}
                    size="small"
                    color={getStatusColor(service.status)}
                    variant="outlined"
                  />
                  <Tooltip title="Configure">
                    <IconButton
                      size="small"
                      onClick={() => navigate(`/settings/${service.name.toLowerCase()}`)}
                    >
                      <SettingsIcon fontSize="small" />
                    </IconButton>
                  </Tooltip>
                </Box>
              }
            >
              <ListItemIcon sx={{ minWidth: 40 }}>
                {getStatusIcon(service.status)}
              </ListItemIcon>
              <ListItemText
                primary={service.displayName || service.name}
                secondary={service.message || 'No additional information'}
              />
            </ListItem>
          ))}
        </List>

        {services.length === 0 && !loading && (
          <Typography color="textSecondary" align="center">
            No services configured
          </Typography>
        )}
      </CardContent>
    </Card>
  );
}
```

### Step 4: Create Activity Feed Component

```jsx
// webui/src/components/Dashboard/ActivityFeed.jsx
import React, { useEffect, useState } from 'react';
import {
  Card,
  CardContent,
  Typography,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Avatar,
  Box,
  Chip
} from '@mui/material';
import {
  CloudDownload as DownloadIcon,
  Sync as SyncIcon,
  Person as UserIcon,
  CheckCircle as SuccessIcon,
  Error as ErrorIcon,
  Info as InfoIcon
} from '@mui/icons-material';
import { apiService } from '../../services/api.js';

export default function ActivityFeed() {
  const [activities, setActivities] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadActivities();
    const interval = setInterval(loadActivities, 30000); // Update every 30 seconds
    return () => clearInterval(interval);
  }, []);

  const loadActivities = async () => {
    try {
      const response = await apiService.get('/api/activities/recent?limit=10');
      if (response.ok) {
        const data = await response.json();
        setActivities(data);
      }
    } catch (error) {
      console.error('Failed to load activities:', error);
    } finally {
      setLoading(false);
    }
  };

  const getActivityIcon = (type) => {
    switch (type) {
      case 'download':
        return <DownloadIcon color="primary" />;
      case 'sync':
        return <SyncIcon color="secondary" />;
      case 'user':
        return <UserIcon color="default" />;
      case 'success':
        return <SuccessIcon color="success" />;
      case 'error':
        return <ErrorIcon color="error" />;
      default:
        return <InfoIcon color="info" />;
    }
  };

  const formatTime = (timestamp) => {
    const date = new Date(timestamp);
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
  };

  return (
    <Card>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Recent Activity
        </Typography>

        <List dense>
          {activities.map((activity, index) => (
            <ListItem key={index} divider={index < activities.length - 1}>
              <ListItemIcon>
                <Avatar sx={{ width: 32, height: 32 }}>
                  {getActivityIcon(activity.type)}
                </Avatar>
              </ListItemIcon>
              <ListItemText
                primary={
                  <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                    <Typography variant="body2">
                      {activity.message}
                    </Typography>
                    {activity.status && (
                      <Chip
                        label={activity.status}
                        size="small"
                        color={activity.status === 'success' ? 'success' : 'error'}
                        variant="outlined"
                      />
                    )}
                  </Box>
                }
                secondary={formatTime(activity.timestamp)}
              />
            </ListItem>
          ))}
        </List>

        {activities.length === 0 && !loading && (
          <Typography color="textSecondary" align="center">
            No recent activity
          </Typography>
        )}
      </CardContent>
    </Card>
  );
}
```

### Step 5: Update Main Dashboard Component

```jsx
// Update webui/src/Dashboard.jsx to include new components
import StatisticsCards from './components/Dashboard/StatisticsCards.jsx';
import QuickActions from './components/Dashboard/QuickActions.jsx';
import ServiceStatus from './components/Dashboard/ServiceStatus.jsx';
import ActivityFeed from './components/Dashboard/ActivityFeed.jsx';

// Add to return statement after the header:
return (
  <Box>
    <Typography variant="h4" component="h1" gutterBottom>
      Dashboard
    </Typography>

    {!backendAvailable && (
      <Alert severity="error" sx={{ mb: 3 }}>
        Backend service is not available. Features are currently disabled.
      </Alert>
    )}

    {error && (
      <Alert severity="error" sx={{ mb: 3 }} onClose={() => setError(null)}>
        {error}
      </Alert>
    )}

    {/* New Enhanced Dashboard Widgets */}
    <StatisticsCards />

    <Grid container spacing={3} sx={{ mb: 3 }}>
      <Grid item xs={12}>
        <QuickActions />
      </Grid>
    </Grid>

    <Grid container spacing={3} sx={{ mb: 3 }}>
      <Grid item xs={12} md={6}>
        <ServiceStatus />
      </Grid>
      <Grid item xs={12} md={6}>
        <ActivityFeed />
      </Grid>
    </Grid>

    {/* Existing Scan Controls (Enhanced) */}
    <Grid container spacing={3}>
      {/* ... existing scan controls code ... */}
    </Grid>
  </Box>
);
```

## 🔌 Backend API Requirements

### New Endpoints Needed

```go
// pkg/webserver/statistics.go
func (s *Server) handleDashboardStatistics(w http.ResponseWriter, r *http.Request) {
    stats := &DashboardStats{
        TotalFiles:      s.getTotalFiles(),
        SuccessRate:     s.getSuccessRate(),
        DownloadedToday: s.getDownloadedToday(),
        ActiveTasks:     s.getActiveTasks(),
        Trends:          s.getTrends(),
    }
    json.NewEncoder(w).Encode(stats)
}

func (s *Server) handleServiceStatus(w http.ResponseWriter, r *http.Request) {
    services := s.getServiceStatus()
    json.NewEncoder(w).Encode(services)
}

func (s *Server) handleRecentActivities(w http.ResponseWriter, r *http.Request) {
    limit := 10
    if l := r.URL.Query().Get("limit"); l != "" {
        if parsed, err := strconv.Atoi(l); err == nil {
            limit = parsed
        }
    }

    activities := s.getRecentActivities(limit)
    json.NewEncoder(w).Encode(activities)
}
```

## 🧪 Testing Requirements

### Component Tests
- Statistics cards data display
- Quick actions functionality
- Service status indicators
- Activity feed updates

### Integration Tests
- Dashboard API endpoints
- Real-time data updates
- Error handling scenarios

### E2E Tests
- Complete dashboard workflow
- Quick action execution
- Service status monitoring

## 📊 Success Metrics

- Improved dashboard engagement
- Faster access to common operations
- Better visibility of system status
- Reduced time to identify issues
- Increased user satisfaction

## 📋 Task Dependencies

- **Depends on**: TASK-02-001 (Navigation completed)
- **Enhances**: Dashboard user experience
- **Prepares for**: TASK-02-009 (Simplified mode)

## ✅ Completion Checklist

- [ ] Statistics cards implementation
- [ ] Quick actions component
- [ ] Service status monitoring
- [ ] Activity feed component
- [ ] Backend API endpoints
- [ ] Real-time data updates
- [ ] Error handling and loading states
- [ ] Responsive design
- [ ] Component testing
- [ ] Integration testing
- [ ] User acceptance testing
- [ ] Performance optimization
