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
		},
	},
  css: {
    postcss: {
      plugins: [
        tailwindcss({
          content: [
            'ui/**/*.svelte',
            'index.html',
          ],
          theme: {
            colors: {
              black: '#1a1a1a',
              blackgrayer: '#212121',
              blackgray: '#2a2a2a',
              gray: '#999999',
              graywhite: '#aaaaaa',
              grayblack: '#3a3a3a',
              white: '#cccccc',
            },
            fontFamily: {
              zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
            },
          },
        }),
      ]
    },
  },
})
