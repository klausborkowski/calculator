import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      '/package': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/packages': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/calculate': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})

