import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.jsx';
import './index.css';

// Initialize i18n
import './i18n/i18n.js';

// Mock API service for testing environments (Lighthouse, etc.)
import './services/mockApi.js';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>
);
