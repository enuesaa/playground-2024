# codetrailer
A CLI tool to capture stdin/stdout and generate a step-by-step document.

## コンセプト
コマンドの実行やキャプチャを行えるCLIツール

- コマンドを実行するとプロンプトが立ち上がり、レコーディングを開始する
- 同時にメニューバーが立ち上がり、そこでスクリーンショットの取得を行える
- 他に PDF の出力や HTML ファイルの出力を行える

## designdoc
- 各ステップは細かく刻む
- save .codetrailer/<name>.md

### Commands
- codetrailer write <name>
- codetrailer preview <name>
- codetrailer export <name>
