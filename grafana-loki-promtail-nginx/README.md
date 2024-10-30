# grafana

- nginx のログを promtail で収集して loki + grafana で確認できるようにした
- nginx のログはファイルとして吐き出される想定。nginx の公式イメージは stdout に流すので、ここでは使用してない 

## Links
- https://grafana.com/docs/loki/latest/setup/install/docker/
- https://zenn.dev/hohner/articles/6b9ab871b422db
