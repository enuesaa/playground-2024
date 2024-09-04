# otel-desktop-viewer

- ローカルで otel の traces を確認できる
- 同じような課題感を持っており、こう言うツールをずっと探していたので、やっと見つけられて嬉しい
- なんか trace の更新がリアルタイムに反映されるから websocket でも使っているのかな。あるいは polling かも

## Links
- https://github.com/CtrlSpice/otel-desktop-viewer
- https://zenn.dev/ymtdzzz/articles/a3a809ca1ba440
- https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/labstack/echo/otelecho/example/server.go

## Commands
```bash
# install
brew tap CtrlSpice/homebrew-otel-desktop-viewer
brew install otel-desktop-viewer

# start up
otel-desktop-viewer
```

### configure app
```bash
OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318
```
