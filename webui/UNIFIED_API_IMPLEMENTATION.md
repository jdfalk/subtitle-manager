# Unified API Implementation for React Frontend

## Overview

This document outlines the implementation of a unified API service for the Subtitle Manager React frontend, addressing issue #543. The unified API provides a consistent, semantic interface for all backend communication, replacing scattered `fetch()` calls with organized, well-documented methods.

## Key Improvements

### 1. **Unified API Structure**

The new `apiService` is organized into semantic modules:

```javascript
// Before: Direct fetch calls scattered throughout components
const res = await fetch('/api/convert', { method: 'POST', body: form });

// After: Semantic, unified API calls
const response = await apiService.subtitles.convert(file);
```

### 2. **Comprehensive Method Organization**

The API is organized into logical groupings:

- **Authentication & OAuth**: `apiService.auth.*`, `apiService.oauth.*`
- **Configuration**: `apiService.config.*`
- **Setup**: `apiService.setup.*`
- **Providers**: `apiService.providers.*`
- **Subtitles**: `apiService.subtitles.*`
- **Library**: `apiService.library.*`
- **System**: `apiService.system.*`
- **Database**: `apiService.database.*`
- **Users**: `apiService.users.*`
- **Tags**: `apiService.tags.*`
- **And more...** (40+ methods total)

### 3. **Enhanced File Upload Support**

```javascript
// Dedicated FormData handling for file uploads
async postFormData(url, formData, config = {}) {
  const { headers, ...otherConfig } = config;
  // Automatically removes Content-Type to let browser set boundary
  const cleanHeaders = { ...headers };
  delete cleanHeaders['Content-Type'];
  
  return apiClient(url, {
    method: 'POST',
    body: formData,
    headers: cleanHeaders,
    ...otherConfig,
  });
}
```

### 4. **Standardized Error Handling**

```javascript
// Utility methods for consistent error handling
async parseJsonResponse(response) {
  if (!response.ok) {
    let errorMessage = `HTTP ${response.status}: ${response.statusText}`;
    try {
      const errorData = await response.json();
      if (errorData.message) {
        errorMessage = errorData.message;
      } else if (errorData.error) {
        errorMessage = errorData.error;
      }
    } catch {
      // Fallback to status text if JSON parsing fails
    }
    throw new Error(errorMessage);
  }
  return await response.json();
}
```

### 5. **Built-in Download Helper**

```javascript
// Automated file download handling
async downloadFile(url, filename) {
  const response = await this.get(url);
  if (!response.ok) {
    throw new Error(`Download failed: ${response.status}`);
  }
  
  const blob = await response.blob();
  const downloadUrl = window.URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = downloadUrl;
  link.download = filename;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  window.URL.revokeObjectURL(downloadUrl);
}
```

## Components Updated

### Core Components Migrated:

1. **Convert.jsx**
   ```javascript
   // Before
   const res = await fetch('/api/convert', { method: 'POST', body: form });
   
   // After  
   const response = await apiService.subtitles.convert(file);
   ```

2. **Extract.jsx**
   ```javascript
   // Before
   const res = await fetch('/api/extract', {
     method: 'POST',
     headers: { 'Content-Type': 'application/json' },
     body: JSON.stringify({ path }),
   });
   
   // After
   const response = await apiService.subtitles.extract(path);
   ```

3. **Translate.jsx**
   ```javascript
   // Before
   const form = new FormData();
   form.append('file', file);
   form.append('lang', lang);
   const res = await fetch('/api/translate', { method: 'POST', body: form });
   
   // After
   const response = await apiService.subtitles.translate(file, lang);
   ```

4. **UserManagement.jsx**
   ```javascript
   // Before
   const res = await fetch('/api/users', {
     method: 'GET',
     headers: { 'Content-Type': 'application/json' },
     credentials: 'include',
   });
   
   // After
   const response = await apiService.users.list();
   ```

5. **Setup.jsx**
   ```javascript
   // Before
   const res = await fetch('/api/setup/bazarr', {
     method: 'POST',
     headers: { 'Content-Type': 'application/json' },
     body: JSON.stringify({ url: bazarrURL, api_key: bazarrAPIKey }),
   });
   
   // After
   const response = await apiService.setup.importBazarr(bazarrURL, bazarrAPIKey);
   ```

6. **DatabaseSettings.jsx**
   ```javascript
   // Before
   const response = await fetch('/api/database/info');
   const statsResponse = await fetch('/api/database/stats');
   
   // After
   const response = await apiService.database.getInfo();
   const statsResponse = await apiService.database.getStats();
   ```

## API Method Examples

### Authentication
```javascript
await apiService.auth.login(username, password);
await apiService.auth.logout();
```

### Subtitle Operations
```javascript
await apiService.subtitles.convert(file);
await apiService.subtitles.extract(filePath);
await apiService.subtitles.translate(file, targetLanguage, options);
await apiService.subtitles.download(downloadId);
```

### Library Management
```javascript
await apiService.library.browse(path);
await apiService.library.startScan(options);
await apiService.library.getScanStatus();
await apiService.library.getTags();
```

### System & Database
```javascript
await apiService.system.getInfo();
await apiService.database.backup();
await apiService.database.optimize();
```

### User Management
```javascript
await apiService.users.list();
await apiService.users.create(userData);
await apiService.users.update(userId, userData);
await apiService.users.resetPassword(userId);
```

## Benefits Achieved

### 1. **Consistency**
- All API calls now follow the same pattern
- Standardized error handling across the application
- Consistent response processing

### 2. **Maintainability**
- Centralized API logic in one place
- Easy to modify endpoints or add authentication
- Clear organization by functionality

### 3. **Developer Experience**
- Semantic method names that are self-documenting
- Comprehensive JSDoc documentation
- IDE autocompletion support

### 4. **Type Safety & Validation**
- Centralized parameter validation
- Consistent request/response handling
- Better error messages

### 5. **Future-Proofing**
- Easy to add new endpoints
- Simple to implement request/response interceptors
- Ready for TypeScript migration

## Testing & Validation

- ✅ **Linting**: All code passes ESLint validation
- ✅ **Build**: Successfully builds for production
- ✅ **Components**: Updated components use unified API correctly
- ✅ **Backward Compatibility**: Existing API endpoints unchanged

## Usage Examples

### Basic Usage
```javascript
import { apiService } from './services/api.js';

// Simple operations
const config = await apiService.config.get();
const providers = await apiService.providers.list();

// With error handling
try {
  const response = await apiService.subtitles.convert(file);
  const data = await apiService.parseJsonResponse(response);
  // Handle success
} catch (error) {
  console.error('Conversion failed:', error.message);
}
```

### File Operations
```javascript
// File uploads
const response = await apiService.subtitles.convert(file);

// File downloads
await apiService.downloadFile('/api/backup', 'backup.tar.gz');
```

### Complex Operations
```javascript
// Setup workflow
const setupStatus = await apiService.setup.getStatus();
if (setupStatus.needed) {
  await apiService.setup.initialize(setupData);
}

// Library management
await apiService.library.startScan({ paths: ['/media'] });
const status = await apiService.library.getScanStatus();
```

## Next Steps

While the core unified API is now implemented, potential future enhancements include:

1. **TypeScript Migration**: Add TypeScript types for all API methods
2. **Request Interceptors**: Add authentication token management
3. **Response Caching**: Implement caching for frequently accessed endpoints
4. **Retry Logic**: Add automatic retry for failed requests
5. **Progress Tracking**: Add upload/download progress support

## Conclusion

The unified API implementation successfully addresses issue #543 by providing a consistent, maintainable, and developer-friendly interface for all React frontend communication with the Subtitle Manager backend. The semantic organization and comprehensive documentation make the codebase more maintainable and easier for new developers to understand.