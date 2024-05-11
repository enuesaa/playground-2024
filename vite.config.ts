import { defineConfig } from 'vite'
import { svelte, vitePreprocess } from '@sveltejs/vite-plugin-svelte'
import tailwindcss from 'tailwindcss'
import path from 'node:path'

export default defineConfig({
  plugins: [
    svelte({
      preprocess: vitePreprocess(),
    }),
  ],
  resolve: {
		alias: {
			$lib: path.join(__dirname, './ui/lib'),
			$components: path.join(__dirname, './ui/components'),
		},
	},
  css: {
    postcss: {
      plugins: [
        tailwindcss('./tailwind.config.ts'),
      ]
    },
  },
})
