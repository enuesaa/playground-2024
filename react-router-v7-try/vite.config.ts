import { reactRouter } from '@react-router/dev/vite'
import tsconfigPaths from 'vite-tsconfig-paths'
import { defineConfig } from 'vite'

export default defineConfig({
  plugins: [
    reactRouter({
      // ssr: true,

      async prerender() {
        return ['/', '/dashboard'];
      },
    }),
    tsconfigPaths()
  ],
})