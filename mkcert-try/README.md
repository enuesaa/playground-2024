# mkcert
- ローカル開発用の証明書を作れるCLIツール
- だいぶ前に使ったことあるかも
- Go 製

- フラグはGo標準の flag で実装しているっぽい
- ビルドターゲットをファイルの suffix につけたら build tag っぽいことができるらしい
  - https://github.com/FiloSottile/mkcert/blob/master/truststore_darwin.go
  - see https://pkg.go.dev/go/build#hdr-Build_Constraints
  - 知らなかった

- `mkcert -install` は作成した証明書をマシンに登録するみたい
  - see https://qiita.com/k_kind/items/b87777efa3d29dcc4467
  - だからこのマシンでないとブラウザで「安全ではない通信」と表示される認識

## Links
- https://github.com/FiloSottile/mkcert

## Commands
```bash
brew install mkcert
brew install nss
mkcert -install
mkcert 0.0.0.0
```
