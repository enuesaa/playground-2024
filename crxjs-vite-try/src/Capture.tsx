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
    console.log(tab)

    const a = await chrome.tabs.captureVisibleTab()
    console.log(a)

    await chrome.downloads.download({
      url: a,
      filename: 'capture.png',
    })
  };

  return <button onClick={handleClick}>capture</button>
}
