# localstack

- もともと知っていたが触れたことがなかったので try
- 内部的には moto を使っているっぽい
  - ぱっとみた感じではエラー系を localstack が実装していて、通過したリクエストに限り最後に moto を呼び出している
  - https://github.com/localstack/localstack/blob/38722cc3472c790862b547e6f7eb61e395ee2adf/localstack-core/localstack/services/secretsmanager/provider.py#L172
- localstack は「AWSライクなプログラムをローカルで動かしている」みたいなもん
- boto3 の endpoint url を変えるだけで普通に動くのは感動する
- 個人的にはメンタルモデルの構築に影響すると思うので、なるべく本物 (AWS) を使った方がいいとは思うが課題感は理解する

## Links
- https://github.com/localstack/localstack
- https://iga-ninja.hatenablog.com/entry/2020/10/24/011800

## Others
- devcontainer は以前より情報が増えた気がする。ただ面倒臭さは減っていない
