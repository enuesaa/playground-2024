from playwright import sync_api
from debugroll import debugroll

# debugroll(sync_api)

with sync_api.sync_playwright() as pw:
    browser = pw.chromium.launch(headless=True)
    page = browser.new_page()
    debugroll(page)
    page.goto("http://playwright.dev")
    browser.close()
