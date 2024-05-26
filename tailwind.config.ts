import { type Config } from 'tailwindcss'

export default {
	content: ['src/**/*.svelte', 'src/app.html'],
	theme: {
		colors: {
			black: '#1a1a1a',
			blackgrayer: '#212121',
			blackgray: '#2a2a2a',
			gray: '#cccccc',
			// grayer: '#dddddd',
			graywhite: '#dddddd',
			grayblack: '#aaaaaa',
			white: '#fafafa',
		},
		fontFamily: {
			zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
		},
	},
	darkMode: 'class',
} satisfies Config
