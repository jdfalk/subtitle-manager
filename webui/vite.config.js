import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  test: {
    environment: 'jsdom',
    globals: true,
    setupFiles: ['./src/test-setup.js'],
    // Exclude Playwright test files to avoid conflicts
    exclude: [
      '**/node_modules/**',
      '**/dist/**',
      '**/tests/**', // Exclude Playwright e2e tests
      '**/.{idea,git,cache,output,temp}/**',
      '**/{karma,rollup,webpack,vite,vitest,jest,ava,babel,nyc,cypress,tsup,build}.config.*',
    ],
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
          // Split router separately
          'react-router': ['react-router-dom'],
          // Create settings chunk for settings-related components
          settings: [
            './src/components/AuthSettings.jsx',
            './src/components/DatabaseSettings.jsx',
            './src/components/GeneralSettings.jsx',
            './src/components/NotificationSettings.jsx',
          ],
        },
      },
    },
    // Increase chunk size warning limit to 750kb
    chunkSizeWarningLimit: 750,
    // Enable minification and compression with esbuild (faster than terser)
    minify: 'esbuild',
    // Configure esbuild to drop console statements in production
    esbuild: {
      drop: ['console', 'debugger'],
      // Enable tree shaking
      treeShaking: true,
    },
    // Enable source maps for better debugging (disabled for production)
    sourcemap: false,
    // Optimize CSS
    cssMinify: true,
    // Optimize assets
    assetsInlineLimit: 4096,
  },
});
