# otel-desktop-viewer

- ローカルで otel の traces を確認できる
- 同じような課題感を持っており、こう言うツールをずっと探していたので、やっと見つけられて嬉しい
- なんか trace の更新がリアルタイムに反映されるから websocket でも使っているのかな。あるいは polling かも

## Links
- https://github.com/CtrlSpice/otel-desktop-viewer
- https://zenn.dev/ymtdzzz/articles/a3a809ca1ba440

## Commands
### Start
```bash
otel-desktop-viewer
```

```bash
OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318
```
