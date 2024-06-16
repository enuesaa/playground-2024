chrome.runtime.onInstalled.addListener(() => {
  console.log('Extension installed');
});

chrome.action.onClicked.addListener((tab) => {
  chrome.tabs.captureVisibleTab(null, {}, function (image) {
    chrome.downloads.download({
      url: image,
      filename: 'capture.png'
    });
  });
});
