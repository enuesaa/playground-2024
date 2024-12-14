import { type Config } from 'tailwindcss'

export default {
	content: ['src/**/*.svelte', 'src/app.html'],
	theme: {
		// do not override theme as app uses tailwindcss in browser dynamically.
		// colors: {
		// 	white: '#fafafa',
		// 	black: '#3a3a3a',
		// 	gray: '#cccccc',
		// 	grayer: '#dddddd',
		// },
		// fontFamily: {
		// 	zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
		// },
	},
	darkMode: 'class',
} satisfies Config
