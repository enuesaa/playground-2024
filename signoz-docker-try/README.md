# signoz

- otel
- db は clickhouse
- 普通に docker で動く
- install.sh の実処理は docker compose
- デフォルトではサンプルアプリも作られる
- なので代わりに `docker compose -f docker/clickhouse-setup/docker-compose-minimal.yaml up -d` を実行する
  - see https://signoz.io/docs/operate/docker-standalone/#running-signoz-without-sample-application

## Links
- https://signoz.io/docs/install/docker/
