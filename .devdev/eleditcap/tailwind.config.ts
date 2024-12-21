import { type Config } from 'tailwindcss'

export default {
	content: ['src/**/*.svelte', 'src/app.html'],
	theme: {
		// Do not override the theme, as this app also loads Tailwind in the browser dynamically.
	},
	darkMode: 'class',
} satisfies Config
