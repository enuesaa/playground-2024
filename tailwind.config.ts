import { type Config } from 'tailwindcss'

export default {
	content: ['src/**/*.svelte', 'src/app.html'],
	theme: {
		colors: {
			black: '#1a1a1a',
			gray: '#cccccc',
			graywhite: '#dadada',
			grayblack: '#b9b9b9',
			white: '#fafafa',
		},
		fontFamily: {
			zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
		},
	},
	darkMode: 'class',
} satisfies Config
