# elastic apm

- 思いのほかオブザーバビリティに力を入れているよう
- docker-compose.esv7.yaml では trace を確認できた
- docker-compose.esv8.yaml では trace を確認できてない
- elastdocker でコンテナを立ち上げると trace がうまく elasticsearch に入り、kibana で確認できた
- see https://github.com/sherifabdlnaby/elastdocker

## Links
- https://medium.com/@sohaib278/distributed-tracing-with-elastic-apm-in-go-part-1-98ef94a514d7
- https://zenn.dev/fujimotoshinji/scraps/4fb4616976ee00
- https://adhon-rizky.medium.com/enable-apm-on-clustered-elasticsearch-stack-8-6-x-5201d48e3fa1
- https://inokara.hateblo.jp/entry/2022/03/13/192922

## Commands
```bash
ELASTIC_APM_SERVER_URL=http://localhost:8200 ELASTIC_APM_SERVICE_NAME=a ELASTIC_APM_LOG_LEVEL=debug ELASTIC_APM_ENVIRONMENT=a CGO_LDFLAGS="-Wl,-no_warn_duplicate_libraries" go run .
```