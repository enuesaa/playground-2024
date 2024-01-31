from playwright import sync_api
import time
from debugroll import debugroll

debugroll(sync_api)

def main():
    pw = sync_api.sync_playwright().start()
    browser = pw.chromium.launch(
        headless=False,
    )
    page = browser.new_page()
    debugroll(page)
    page.goto("http://playwright.dev")
    print(page.title())
    time.sleep(2)
    browser.close()
    pw.stop()
