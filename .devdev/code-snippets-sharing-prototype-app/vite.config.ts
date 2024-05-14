import { sveltekit } from '@sveltejs/kit/vite'
import path from 'node:path'
import { defineConfig } from 'vite'

export default defineConfig({
	plugins: [sveltekit()],
	resolve: {
		alias: {
			$lib: path.join(__dirname, './lib'),
		},
	},
})
