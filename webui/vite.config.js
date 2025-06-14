import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  test: {
    environment: 'jsdom',
    globals: true,
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          // Split MUI into separate chunks
          'mui-core': ['@mui/material', '@mui/system'],
          'mui-icons': ['@mui/icons-material'],
          'mui-styles': ['@emotion/react', '@emotion/styled'],
          // Split React into separate chunk
          'react-vendor': ['react', 'react-dom'],
        }
      }
    },
    // Increase chunk size warning limit to 750kb
    chunkSizeWarningLimit: 750,
  }
})
