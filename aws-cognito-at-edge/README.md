# aws-cognito-at-edge

- CloudFront でホストしているサイトにログイン機能をつけたいとき、cognito-at-edge を活用できる
- cognito-at-edge は awslabs が用意している NPM パッケージ
- Cognito をセッティングして Lambda@Edge をデプロイすればログイン機能を実現できる

## Links
- https://github.com/awslabs/cognito-at-edge

## 仕組み
- Lambda@Edge が起動したとき Cookie がなかったら Cognito の Hosted UI にリダイレクトする仕組み
- Cookie があったらそのままコンテンツを閲覧できる

## ユースケース
- 簡易的にログイン機能をつけたいときめちゃくちゃ便利
- 中規模チームのインターナルな管理画面には最適
- 実現できる機能は cognito-at-edge に縛られるので、画面をユーザに提供するようなサービスでは拡張性に問題があるかも
- 正直ユースケースは限られそうな気がしているが、これを参考に機能開発できるので、その意味では非常にありがたい
- あくまで Cognito を普通に使った場合に有用そう
