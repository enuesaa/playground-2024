# mise

- asdf みたいなやつ
- カレンダーバージョニングを採用しているっぽい
- フラグがいろいろあり、いろんなことができそう
- 環境変数の管理やタスクランナー的な機能もある

## Commands
```console
# terraform cli は mise でインストールしたわけではないので、なぜ表示されているのか分からない
$ mise ls
Tool       Version          Config Source    Requested
terraform  1.2.9 (missing)  ~/.tool-versions 1.2.9

$ mise exec --cd ~/tmp -- echo a
a
```

## Links
- https://github.com/jdx/mise
- https://zenn.dev/euxn23/articles/b12f1f9b495d47

