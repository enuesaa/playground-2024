# playwright-python
### install browser
```bash
poetry run playwright install
```

### memo
- sycn api というメソッド?くくり?は playwright-go にないので新鮮
- コード上からブラウザをインストールするメソッドは無いみたい. playwright-go には存在し便利なのだが..
- なんかplaywrightに関係ないが go に比べコード量は少ない
- ただし try except の方がエラー処理が雑だなあ、と感じた。ごっちゃ煮感
- goやrustみたいにエラーを返した方が、ちゃんとコードを書いている感、、考慮してます感がある
