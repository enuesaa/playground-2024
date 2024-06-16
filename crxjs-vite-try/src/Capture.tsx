import { MouseEventHandler } from 'react'

export const Capture = () => {
  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()

    // const tab = await chrome.tabs.create({
    //   url: 'https://example.com',
    // })
    // console.log(tab)

    const tabs = await chrome.tabs.query({
      active: true,
      currentWindow: true,
    })
    if (tabs.length === 0) {
      return
    }
    const tab = tabs[0]
    const tabId = tab.id

    if (tabId === undefined) {
      return
    }

    const result = await chrome.scripting.executeScript({
      target: { tabId },
      files: ['contents.js'],
    })
    if (result.length === 0) {
      return
    }
    // @ts-ignore
    const width = result[0].result.width
    // @ts-ignore
    const height = result[0].result.height
    await chrome.debugger.attach({ tabId }, '1.3')

    // chrome devtool protocol
    // see https://zenn.dev/kakkoyakakko/articles/54fe29dc3751b9
    // see https://github.com/microsoft/playwright/blob/2ae2fb421c6b6ec115721a8b0a5215442e5411c3/packages/playwright-core/src/server/chromium/crPage.ts#L279
    // see https://github.com/microsoft/playwright/blob/2ae2fb421c6b6ec115721a8b0a5215442e5411c3/packages/playwright-core/src/server/screenshotter.ts#L187
    const layoutMetrics = await chrome.debugger.sendCommand({ tabId }, 'Page.getLayoutMetrics', {})
    console.log(layoutMetrics)

    // await chrome.debugger.sendCommand({tabId}, 'Emulation.setDeviceMetricsOverride', {
    //   // @ts-ignore
    //   width, //: layoutMetrics.contentSize.width,
    //   // @ts-ignore
    //   height, //: layoutMetrics.contentSize.height,
    //   deviceScaleFactor: 1,
    //   mobile: false,
    //   fitWindow: false,
    // })

    // await chrome.debugger.sendCommand({ tabId }, 'Emulation.setVisibleSize', {
    //   width,
    //   height,
    // })

    // see https://stackoverflow.com/questions/66259592/how-to-resize-chrome-browser-window-with-an-extension
    // await chrome.windows.update(tab.windowId, {
    //   width,
    //   height,
    // })

    const res = await chrome.debugger.sendCommand({ tabId }, 'Page.captureScreenshot', {
      format: 'png',
      clip: {
        "x": 0, 
        "y": 0,
        // "width": width,
        // "height": height,
      // @ts-ignore
      width: layoutMetrics.contentSize.width,
      // @ts-ignore
      height: layoutMetrics.contentSize.height,
        "scale": 1
      },
      // fromSurface: true,

      captureBeyondViewport: true,
    })
    console.log(res)

    // @ts-ignore
    const imageUrl: string = `data:image/png;base64,${res.data}`

    await chrome.downloads.download({
      url: imageUrl,
      filename: 'capture.png',
    })

    await chrome.debugger.detach({ tabId })
  };

  return <button onClick={handleClick}>capture</button>
}
