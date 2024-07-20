# pointmap

## Development Plan
### 課題感
- ブログに図を入れたい
- 図を描くツールをいろいろ見てみたが、自分のイメージしているものではなかった

### What I want..
- 文章のポイントを示せるものが欲しい。リッチな図ではなくシンプルな図を求めている
- プライベートで使うので、短時間で図を作成できるツール
- 全体像が分かるもの
- そもそも「図の作成」は高コストな作業であり、時間がかかるのでプライベートではモチベーションが上がらないのではないか

### Architecture
- Go + Sveltekit で開発し Go のビルド成果物(バイナリ)にアプリケーションを全部含めてしまう
- バイナリを実行するとUI(HTMLなど)やREST APIのレスポンス(JSON形式)が返ってくる
- これをLambdaにデプロイし関数URLを叩く。CloudFront の Origin に設定する

### Stack
- Go
- Sveltekit
- CloudFront
- Lambda (amazon linux 2023)
- DynamoDB

### Local Development
- DynamoDB Local
- invoke go binary from local. App normally serves http.

### Links
- https://github.com/awslabs/aws-lambda-go-api-proxy
- https://dev.classmethod.jp/articles/cloudfront-lambda-url-sigv4-signer/
