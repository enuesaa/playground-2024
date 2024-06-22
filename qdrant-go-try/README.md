# qdrant go
## create collections with curl
see https://github.com/qdrant/qdrant/blob/master/QUICK_START.md

```bash
curl -X PUT 'http://localhost:6333/collections/test_collection' \
    -H 'Content-Type: application/json' \
    --data-raw '{
        "vectors": {
          "size": 4,
          "distance": "Dot"
        }
    }'
```

```console
$ go run .
2024/06/23 01:13:12 [name:"test_collection"]
```

## 感想
- qdrant は rust 製らしい。GCないし運用楽そう
- ElasticSearch みたいな感じなのかな

## Links
- https://github.com/qdrant/qdrant/blob/master/QUICK_START.md
- https://dev.classmethod.jp/articles/qdrant-first-step/
- https://github.com/qdrant/go-client
- https://zenn.dev/yskaksk/articles/qdrant-intro
