# grafana-tempo
grafana tempo でトレーシングができると聞きつけ試してみた

## Setup
```bash
docker compose up -d

# grafana を開いてデータソース設定
# 1. ブラウザで localhost:3000 を開く
# 2. ログイン
# 3. データソースの設定画面へ
# 4. URL に http://tempo:3200/ を入力
# 5. Save
# 6. Explorer で tempo を選択できる

go run .

# ブラウザで localhost:3001/users/aaa をたくさん開く
```

## メモ
- grafana tempo 単体でトレース情報をレシーブすることもできるが前段に otel collector を置く構成が多い
- ドキュメントが少ない
- example の docker-compose.yaml に記載のコンテナの数が多く、どれから手をつければいいか悩んだ

- 個人的には jaeger と大差ない印象
  - jaeger は比較的独自路線のはずなので、open telemetry に沿う意味合いでは、将来的にはこちらの方がいいかも
- ローカル開発環境では、まあ必要あれば使えるし有用だが、プロダクションでは悩む

- otel-collector はなくてもいい。アプリから tempo に直接送ることもできる

## Links
- https://techblog.goinc.jp/entry/2024/05/17/153729
- https://grafana.com/oss/tempo/

