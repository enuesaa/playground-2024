# fyne
- golang
- gui (mac native app)
- たぶん windows でも立ち上がるけど、事例が少ない？
- wails は webview ベースだけど、fyne は GUI のテンプレートみたいなのがあって、それを Go で組み立てる方式

## Links
- https://github.com/fyne-io/fyne

## Commands
### install fyne cli
```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
```

### package app
```bash
fyne package -os darwin -icon icon.png
```
