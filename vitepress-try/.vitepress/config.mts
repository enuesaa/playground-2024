import { defineConfig } from 'vitepress'
import { getSidebar } from 'vitepress-plugin-auto-sidebar'

export default defineConfig({
  title: "mytestdoc",
  description: "my test doc",
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      // { text: 'Examples', link: '/markdown-examples' }
    ],
    // sidebar: [
    //   // {
    //   //   text: 'Examples',
    //   //   items: [
    //   //     { text: 'Markdown Examples', link: '/markdown-examples' },
    //   //     { text: 'Runtime API Examples', link: '/api-examples' }
    //   //   ]
    //   // }
    // ],
    sidebar: getSidebar({
      contentRoot: '/',
      contentDirs: ['docs'],
      collapsible: true,
      collapsed: true,
    }),
    socialLinks: [
      { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})
