import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { crx, defineManifest } from '@crxjs/vite-plugin'

const manifest = defineManifest({
  manifest_version: 3,
  name: 'CRXJS React Vite Example',
  version: '0.0.1',
  permissions: ['activeTab', 'debugger', 'storage', 'downloads'],
  action: {
    default_popup: 'index.html',
  },
})

export default defineConfig({
  plugins: [
    react(),
    crx({ manifest }),
  ],
})
