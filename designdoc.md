# designdoc
## 課題感
- ブログに図を入れたい
- 文章のポイントを示せるものが欲しい。リッチな図ではなくシンプルな図を求めている
- プライベートで使うので、短時間で図を作成できるツール
- 図を描くツールをいろいろ見てみたが、自分のイメージしているものではなかった

## Architecture
- Go + Sveltekit で開発し Go のビルド成果物(バイナリ)にアプリケーションを全部含めてしまいたい
- バイナリを実行するとUI(HTMLなど)やREST APIのレスポンス(JSON形式)が返ってくるイメージ
- ただLambdaがPOSTしか受け付けないはずなので、間にAPI Gatewayを入れるか検討中

## Stack
- Go
- Sveltekit
- CloudFront
- Lambda (amazon linux 2023)
- DynamoDB

## Local Development
- DynamoDB Local
- Invoke local lambda function using docker image 
  https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/images-test.html#images-test-AWSbase
