import { type Config } from 'tailwindcss'

export default {
	content: [
		'ui/**/*.svelte',
		'index.html',
	],
	theme: {
		colors: {
			black: '#1a1a1a',
			blackgray: '#2a2a2a',
			gray: '#999999',
			grayblack: '#3a3a3a',
			graywhite: '#cfcfcf',
			white: '#fafafa',
		},
		fontFamily: {
			zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
		},
	},
} satisfies Config
