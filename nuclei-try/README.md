# nuclei
- ペネトレーションテストツール
- なんか事前にシナリオが定義されてあって、それを実行しているっぽい
- フラグでシナリオを渡さないと、テンプレートをダウンロードしてきてテストする
- なかなか恐ろしいツールだなと思った

## Install
- 依存パッケージが多く Go にしては install に時間がかかる

```bash
go install -v github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest
```

## Command
```console
$ nuclei -target example.com

[options-method] [http] [info] https://example.com ["OPTIONS, GET, HEAD, POST"]
[weak-cipher-suites:tls-1.1] [ssl] [low] example.com:443 ["[tls11 TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA]"]
[tls-version] [ssl] [info] example.com:443 ["tls12"]
[tls-version] [ssl] [info] example.com:443 ["tls13"]
```

## Links
- https://jpn.nec.com/cybersecurity/blog/211029/index.html
- https://github.com/projectdiscovery/nuclei
