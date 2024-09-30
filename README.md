# codetrailer
A CLI tool to generate a step-by-step document.

## コンセプト
コマンドの実行やキャプチャを行えるCLIツール

- コマンドを実行するとウェブサーバが立ち上がり、編集できる
- アクションとして tree, prompt-recording がある
- スクリーンショットの取得もしたいが、ウェブベースだと編集画面が邪魔になってしまい仕様の検討が必要なのでスコープアウトする

### Commands
```bash
codetrailer write # start server
codetrailer preview
codetrailer export
```
