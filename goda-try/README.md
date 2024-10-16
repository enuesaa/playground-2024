# goda

- よそのソースコードを見ててパッケージ間の依存関係を明らかにしたいと感じトライ

## Links
- https://github.com/loov/goda

## Commands
```bash
go install github.com/loov/goda@latest

cd teatime
goda graph github.com/enuesaa/teatime/... | dot -Tsvg -o teatime.svg

cd cpbuf
goda graph github.com/enuesaa/cpbuf/... | dot -Tsvg -o cpbuf.svg
```

意図通り repository に全ての矢印が集まっていて、面白い
