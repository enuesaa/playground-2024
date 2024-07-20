import { sveltekit } from '@sveltejs/kit/vite'
import { defineConfig } from 'vite'
import path from 'node:path'

export default defineConfig({
	plugins: [sveltekit()],
	resolve: {
		alias: {
			$lib: path.join(__dirname, './src/lib'),
		},
	},
})
