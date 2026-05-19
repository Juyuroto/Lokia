import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',
    port: 5000,
    watch: {
      usePolling: true,
      interval: 1000,
    }
  }
})