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
}
```
