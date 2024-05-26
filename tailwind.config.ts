import { type Config } from 'tailwindcss'

export default {
	content: ['src/**/*.svelte', 'src/app.html'],
	theme: {
		colors: {
			black: '#1a1a1a',
			blackgray: '#2a2a2a',
			gray: '#cccccc',
			graywhite: '#dadada',
			grayblack: '#b3b3b3',
			white: '#fafafa',
		},
		fontFamily: {
			zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
			zenmaru: ['Zen Maru Gothic', 'serif'],
		},
	},
	darkMode: 'class',
} satisfies Config
