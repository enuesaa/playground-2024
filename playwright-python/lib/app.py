from playwright.sync_api import sync_playwright
import time

def main():
    pw = sync_playwright().start()
    browser = pw.chromium.launch(
        headless=False,
    )
    page = browser.new_page()
    page.goto("http://playwright.dev")
    print(page.title())
    time.sleep(10)
    browser.close()
    pw.stop()


    # with sync_playwright() as p:
    #     browser = p.chromium.launch(
    #         headless=False,
    #     )
    #     page = browser.new_page()
    #     page.goto("http://playwright.dev")
    #     print(page.title())
    #     time.sleep(10)
    #     browser.close()
