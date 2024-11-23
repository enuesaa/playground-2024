# go-blueprint

- コマンドラインでサンプルアプリを作成できる
- 勉強になる
- echo.Echo が http.Handler を満たしているのは知らなかった。必ずしも echo.Echo の Start を呼ぶ必要ない
- 使うかは別の話

## Install
```bash
go install github.com/melkeydev/go-blueprint@latest
```

## Run
```bash
go-blueprint create --name sampleapp --framework echo --driver mysql --git skip
```
