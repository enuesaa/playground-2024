# go screenshot

## メモ
- mac で動かすには、システム環境設定でターミナルへ権限を付与する必要あり
- 普通にスクリーンショットを撮れた

## 録画するには
- golang でそういうパッケージがあるか片手間に調べたがなさそう
- ffmpeg のコマンドを go から呼び出すとか
  - 下記のコマンドを終了するには q を打つ
  - やっぱ想定していたけど、処理が重たいなあ。

```bash
ffmpeg -f avfoundation -framerate 15 -i "1:none" -video_size 640x480 -t 1 -pix_fmt yuv420p output.mp4
```

## Links
- https://github.com/kbinani/screenshot
