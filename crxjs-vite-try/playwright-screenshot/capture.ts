import { chromium } from 'playwright-chromium'

const browser = await chromium.launch({
  headless: false,
})
const context = await browser.newContext()
const page = await context.newPage()


await page.goto('')
await page.screenshot({
  path: 'screenshot.png',
  fullPage: true,
})
