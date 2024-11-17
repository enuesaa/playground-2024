# wire

- ずっと気になっていた
- ビルドタグ `//+build` を上手いこと使って DI するらしい
- wire_gen.go の InitializeEvent の実装は、依存関係を解決して勝手に組み立ててるっぽい
  - めっちゃ頭いい

## Install
```bash
go install github.com/google/wire/cmd/wire@latest
```

## Gen
```bash
# this generates wire_gen.go
wire
```

## Links
- https://github.com/google/wire
- https://github.com/google/wire/blob/main/_tutorial/README.md
