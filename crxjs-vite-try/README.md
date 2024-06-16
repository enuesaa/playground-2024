# crxjs-vite

- vite で chrome の拡張機能を作れる
- crxjs-vite という vite plugin を入れるだけで拡張機能として読み込めるっぽい
- manifest を plugin の設定へ渡す

## 拡張機能を試す方法
1. `pnpm dev`
2. chrome で拡張機能の管理ページでデベロッパーモードへ
3. 拡張機能を読み込むで dist を選択

実際に有効にして拡張機能を試してみると index.html が表示された

## Links
- https://crxjs.dev/vite-plugin/
- https://dev.classmethod.jp/articles/eetann-chrome-extension-by-crxjs/
- https://zenn.dev/kakkoyakakko/articles/54fe29dc3751b9
- https://github.com/microsoft/playwright/blob/2ae2fb421c6b6ec115721a8b0a5215442e5411c3/packages/playwright-core/src/server/chromium/crPage.ts#L279
- https://github.com/microsoft/playwright/blob/2ae2fb421c6b6ec115721a8b0a5215442e5411c3/packages/playwright-core/src/server/screenshotter.ts#L187
