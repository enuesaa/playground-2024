import { defineConfig } from 'vite'
import { sveltekit } from '@sveltejs/kit/vite'
import path from 'node:path'

export default defineConfig({
  plugins: [sveltekit()],
  resolve: {
    alias: {
      '$lib': path.join(__dirname, './src/lib'),
    },
  },
})
