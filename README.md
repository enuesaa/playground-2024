# webhook-receive-prototype-app
GitHub の Webhook 機能を試してみたくて作ったアプリ

## 機能
- Webhook を受信する
- localhost:3000 で立ち上がる

## Webサーバ
localhost:3000 を Cloudflare Tunnel で中継した

### Install
```bash
brew install cloudflared
```

### 立ち上げ
```bash
go run .
cloudflared tunnel --url localhost:3000
```

## 受信メッセージ
リクエストヘッダ X-GitHub-Event にイベントの名前が入る
### GitHub Repository に webhook を登録したとき
```json
{
  "zen": "Favor focus over features.", ...
}
```

### コミットしたとき
```json
{
  "ref":"refs/heads/main",
  "before":"25416b2b23f23b7cb6a576e262e7e9b676d2a2c4","after":"460cb7e69e4e7cd62da942c64a9a5fd87e24d9ed",
  "repository":{
    "name":"webhook-receive-prototype-app", ...
}
```

## Links
- https://docs.github.com/ja/webhooks/webhook-events-and-payloads
