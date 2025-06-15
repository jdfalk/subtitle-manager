import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  test: {
    environment: 'jsdom',
    globals: true,
    setupFiles: ['./src/test-setup.js'],
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          // Keep MUI packages together to avoid circular dependencies
          mui: [
            '@mui/material',
            '@mui/system',
            '@emotion/react',
            '@emotion/styled',
          ],
          'mui-icons': ['@mui/icons-material'],
          // Split React into separate chunk
          'react-vendor': ['react', 'react-dom'],
        },
      },
    },
    // Increase chunk size warning limit to 750kb
    chunkSizeWarningLimit: 750,
  },
});
