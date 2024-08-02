# dicedb
- Go製
- Redis の代替らしい
- Redisフォークは Valkey. DiceDBはフォークとかではなく全くの別物
- あまり詳しく見れてないが、コードベースが小さい気がしておりびっくりしている
- Go アプリに埋め込めないかな

## Install
### docker
docker-compose.yaml に記載
### マニュアル
```bash
$ go install github.com/dicedb/dice@latest
$ dice --help
Usage of dice:
  -host string
    	host for the dice server (default "0.0.0.0")
  -port int
    	port for the dice server (default 7379)
$ dice
2024/08/03 00:33:57 INFO starting an asynchronous TCP server on 0.0.0.0=7379
2024/08/03 00:33:57 INFO ready to accept connections
```

## Run Query
普通に redis
```bash
$ redis-cli -p 7379
127.0.0.1:7379> SET mykey "Hello, World!"
OK
127.0.0.1:7379> GET mykey
"Hello, World!"
```
