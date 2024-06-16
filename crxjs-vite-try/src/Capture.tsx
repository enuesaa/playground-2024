import { MouseEventHandler } from 'react'

export const Capture = () => {
  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()

    // const tab = await chrome.tabs.create({
    //   url: 'https://example.com',
    // })
    // console.log(tab)
    console.log('a')

    const tabs = await chrome.tabs.query({
      active: true,
      currentWindow: true,
    })
    if (tabs.length === 0) {
      return
    }
    const tab = tabs[0]
    const tabId = tab.id

    await chrome.debugger.attach({ tabId }, '1.3')

    // chrome devtool protocol
    // see https://zenn.dev/kakkoyakakko/articles/54fe29dc3751b9
    // see https://github.com/microsoft/playwright/blob/2ae2fb421c6b6ec115721a8b0a5215442e5411c3/packages/playwright-core/src/server/chromium/crPage.ts#L279
    const layoutMetrics = await chrome.debugger.sendCommand({ tabId }, 'Page.getLayoutMetrics', {})
    console.log(layoutMetrics)
    const res = await chrome.debugger.sendCommand({ tabId }, 'Page.captureScreenshot', {
      format: 'png',
      clip: {
        "x": 0, 
        "y": 0,
        // @ts-ignore
        "width": layoutMetrics.cssLayoutViewport.pageX,
        // @ts-ignore
        "height": layoutMetrics.cssLayoutViewport.pageY,
        "scale": 1
      },
      // captureBeyondViewport: true,
    })

    // @ts-ignore
    const imageUrl: string = `data:image/png;base64,${res.data}`
    console.log(imageUrl)

    await chrome.downloads.download({
      url: imageUrl,
      filename: 'capture.png',
    })

    await chrome.debugger.detach({ tabId })
  };

  return <button onClick={handleClick}>capture</button>
}
