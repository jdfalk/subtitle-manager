<!-- file: docs/tasks/02-ui-fixes/TASK-02-009-simplified-mode-interface.md -->
<!-- version: 1.0.0 -->
<!-- guid: a9b8c7d6-e5f4-3210-9876-543210fedcba -->

# TASK-02-009: Implement Simplified Mode Interface

## 🎯 Objective

Create a "Simplified Mode" toggle that provides a streamlined, non-technical user interface with reduced options and simplified workflows for users who are not comfortable with complex technology interfaces.

## 📋 Acceptance Criteria

- [ ] Add "Simplified Mode" toggle in user preferences/settings
- [ ] Implement simplified dashboard with essential actions only
- [ ] Create wizard-style workflows for common operations
- [ ] Reduce navigation complexity and hide advanced features
- [ ] Provide guided setup and configuration processes
- [ ] Implement simplified terminology and help text
- [ ] Add visual progress indicators for all operations
- [ ] Create one-click automation for common tasks

## 🎨 Simplified Mode Features

### Dashboard Redesign
**Current**: Complex dashboard with multiple widgets and technical information
**Simplified**: Large, clear action buttons with simple descriptions

**Essential Actions:**
- 🔄 **"Sync with Sonarr"** - Large button with progress indicator
- 🎬 **"Sync with Radarr"** - Large button with progress indicator
- 💿 **"Scan My Library"** - Scan disk for new media files
- 📚 **"Browse My Media"** - Access to media library (simplified view)
- ⚙️ **"Quick Setup"** - Guided configuration wizard
- 📊 **"View Activity"** - Simple progress/status overview

### Navigation Simplification
**Hide in Simplified Mode:**
- System/Advanced settings
- Raw configuration options
- Developer tools (Extract/Convert/Translate)
- Complex provider configurations
- Advanced scheduling options

**Simplified Navigation:**
- 🏠 **Home** (Simplified Dashboard)
- 📺 **My Shows & Movies** (Media Library)
- ⚙️ **Settings** (Essential settings only)
- ❓ **Help** (Built-in guides and tutorials)

### Settings Simplification
**Essential Settings Only:**
- **Quick Setup Wizard**
  - Sonarr connection (simplified form)
  - Radarr connection (simplified form)
  - Library location picker
  - Language preference
- **Basic Preferences**
  - Subtitle languages (visual selector)
  - Auto-download toggle
  - Notification preferences (simple on/off)

### User Experience Enhancements
**Simplified Terminology:**
- "Providers" → "Subtitle Sources"
- "Sync" → "Update"
- "Extract" → "Get Subtitles"
- "Queue" → "Downloads"
- "History" → "Recent Activity"

**Visual Improvements:**
- Larger buttons and text
- Clear icons with labels
- Progress bars for all operations
- Success/error messages in plain language
- Tooltips with helpful explanations

## 🔧 Implementation Approach

### Step 1: Mode Toggle Implementation
```jsx
// Add to user preferences
const [simplifiedMode, setSimplifiedMode] = useState(
  localStorage.getItem('simplified-mode') === 'true'
);

// Mode context provider
const SimplifiedModeContext = createContext();
```

### Step 2: Simplified Dashboard Component
```jsx
// SimplifiedDashboard.jsx
const actions = [
  {
    title: "Sync with Sonarr",
    description: "Get new TV shows",
    icon: <TvIcon />,
    action: () => triggerSonarrSync(),
    color: "primary"
  },
  {
    title: "Sync with Radarr",
    description: "Get new movies",
    icon: <MovieIcon />,
    action: () => triggerRadarrSync(),
    color: "secondary"
  },
  {
    title: "Scan My Library",
    description: "Find new video files",
    icon: <ScanIcon />,
    action: () => triggerLibraryScan(),
    color: "success"
  },
  {
    title: "Browse My Media",
    description: "See all shows and movies",
    icon: <LibraryIcon />,
    action: () => navigate('/media'),
    color: "info"
  }
];
```

### Step 3: Wizard Components
```jsx
// QuickSetupWizard.jsx - Multi-step configuration
const steps = [
  "Welcome",
  "Connect Sonarr",
  "Connect Radarr",
  "Choose Languages",
  "Done"
];
```

### Step 4: Conditional UI Rendering
```jsx
// App.jsx modifications
{simplifiedMode ? (
  <SimplifiedNavigation />
) : (
  <AdvancedNavigation />
)}

{simplifiedMode ? (
  <SimplifiedDashboard />
) : (
  <AdvancedDashboard />
)}
```

## 🎪 Simplified Workflows

### Sonarr/Radarr Connection Wizard
**Step 1**: "Let's connect to your TV show manager"
- Simple form: Server address, API key
- Test connection button
- Success/failure with clear messaging

**Step 2**: "Let's connect to your movie manager"
- Same simple process for Radarr
- Auto-detection if possible

### Media Library Simplified View
**Features:**
- Large poster/thumbnail view
- Simple search box
- Filter: "Shows" / "Movies" / "All"
- Status indicators: ✅ Has Subtitles / ❌ Missing Subtitles
- One-click "Get Subtitles" button per item

### Simplified Settings
**Categories:**
1. **Connections** (Sonarr/Radarr setup)
2. **Languages** (Visual language picker)
3. **Notifications** (Simple toggles)
4. **Advanced** (Link to full settings - "Show all options")

## 📱 Mobile-First Design

### Touch-Friendly Interface
- Minimum 44px touch targets
- Large, clear buttons
- Simplified gesture navigation
- Responsive card layouts

### Progressive Disclosure
- Show essential options first
- "Show more" links for advanced features
- Collapsible sections
- Contextual help

## 🧪 User Testing Considerations

### Target Users
- Non-technical family members
- Users new to media management
- Users who find current interface overwhelming
- Mobile/tablet users

### Success Metrics
- Reduced time to complete common tasks
- Lower support requests
- Higher user adoption
- Positive user feedback

## 🔄 Implementation Phases

### Phase 1: Core Infrastructure
- [ ] Mode toggle implementation
- [ ] Context provider setup
- [ ] Basic simplified dashboard
- [ ] Navigation switching logic

### Phase 2: Essential Features
- [ ] Sonarr/Radarr sync buttons
- [ ] Library scan functionality
- [ ] Simplified media library view
- [ ] Basic settings interface

### Phase 3: Wizards and Guides
- [ ] Quick setup wizard
- [ ] Connection wizards
- [ ] Built-in help system
- [ ] Progress indicators

### Phase 4: Polish and Testing
- [ ] User testing with non-technical users
- [ ] Accessibility improvements
- [ ] Mobile optimization
- [ ] Documentation and tutorials

## 🎨 Visual Design Mockup Ideas

### Simplified Dashboard Layout
```
┌─────────────────────────────────────┐
│  Welcome to Subtitle Manager        │
│                                     │
│  ┌─────────┐  ┌─────────┐          │
│  │   📺    │  │   🎬    │          │
│  │ Sync TV │  │Sync     │          │
│  │ Shows   │  │Movies   │          │
│  └─────────┘  └─────────┘          │
│                                     │
│  ┌─────────┐  ┌─────────┐          │
│  │   💿    │  │   📚    │          │
│  │ Scan    │  │Browse   │          │
│  │Library  │  │Media    │          │
│  └─────────┘  └─────────┘          │
│                                     │
│  Recent Activity: ✅ 5 new subtitles│
└─────────────────────────────────────┘
```

## 🚀 Benefits

### For Non-Technical Users
- Reduced cognitive load
- Clear, obvious actions
- Guided workflows
- Plain language explanations
- Visual progress feedback

### For Power Users
- Optional mode - doesn't replace existing interface
- Quick toggle between modes
- Maintains all advanced functionality
- Can introduce features gradually

### For Support/Maintenance
- Reduced user confusion
- Fewer support requests
- Easier onboarding
- Better user adoption

## 🔗 Integration Points

### Existing Components to Enhance
- Dashboard.jsx → SimplifiedDashboard.jsx
- Navigation.jsx → SimplifiedNavigation.jsx
- Settings.jsx → SimplifiedSettings.jsx
- MediaLibrary.jsx → SimplifiedMediaLibrary.jsx

### New Components to Create
- QuickSetupWizard.jsx
- ConnectionWizard.jsx
- SimplifiedModeToggle.jsx
- HelpCenter.jsx
- ProgressOverlay.jsx

## 📊 Success Criteria

### Quantitative Metrics
- 50% reduction in time for first-time setup
- 75% reduction in common task completion time
- 90% of users can complete basic tasks without help
- <10% switch back to advanced mode

### Qualitative Feedback
- "Easy to understand"
- "Feels less overwhelming"
- "I know what to click"
- "It just works"

This simplified mode will make Subtitle Manager accessible to a much broader audience while maintaining the power and flexibility that advanced users need.
