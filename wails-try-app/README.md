# wails
- tauri に近い
- go + frontend app で native app を作成できる
- frontend app の sample app で svelte が一番初めに表記されていてびっくりした


## install
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## dev
```bash
wails dev
```

## build
```bash
wails build # おそらくホストマシンのアーキテクチャ向けにビルド
wails build --platform windows/amd64 -tags native_webview2loader # windows 向けにビルド
```
./build/bin 配下に置かれる

## フロントエンドアプリとのコミュニケーション方法
- Go の App にレシーバを生やして wails dev をすると d.ts とか生成されるっぽい
