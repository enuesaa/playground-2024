# pocketbase

## Install
- goreleaser を使っているっぽいので、普通に brew fomula の生成とかできると思うが、していないらしい
- なので Github Release からダウンロードする

```bash
mv pocketbase /usr/local/bin/pocketbase
```

## How to use
- pocketbase コマンドを実行すると pb_data というディレクトリが作成される
- 中身は SQLite と d.ts っぽい
- --dir フラグで変更可能

```bash
pocketbase --dir a
```

### Serve
```bash
pocketbase serve --http 0.0.0.0:3000
```
