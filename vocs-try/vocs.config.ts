import { defineConfig } from 'vocs'

export default defineConfig({
  title: 'mydocs',
  rootDir: '.', 
  // banner: 'a',
  // basePath: '/docs', 
  // blogDir: './pages/writings',
  sidebar: [
    {
      text: 'Getting Started',
      link: '/getting-started',
    },
    {
      text: 'Example',
      link: '/example',
    },
    {
      text: 'Example2',
      link: '/example2',
    },    {
      text: 'hide sidebar',
      link: '/hide-sidebar',
    },
  ],
})
