import { type Config } from 'tailwindcss'

export default {
	content: ['src/**/*.svelte', 'src/app.html'],
	theme: {
		colors: {
			black: '#1a1a1a',
			blackgray: '#2a2a2a',
			gray: '#cccccc',
			graywhite: '#dadada',
			grayblacker: '#bfbfbf',
			grayblack: '#b3b3b3',
			white: '#fafafa',
		},
		fontFamily: {
			zenmaru: ['Zen Maru Gothic', 'serif'],
			ibmplex: ['"IBM Plex Sans JP"', 'sans-serif'],
		},
	},
	darkMode: 'class',
} satisfies Config
