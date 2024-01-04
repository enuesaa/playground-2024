# pagefind-trial
## メモ
- 静的コンテンツに対してインデックスできる
- wasm で rust のコードを動かしているっぽい
- 開発元がCMSのSaaSっぽい

## Links
- https://github.com/cloudcannon/pagefind
- https://griponminds.jp/blog/try-pagefind/

## Commands
```bash
pnpm add pagefind
```

## `pagefind --site dist`
stdout of `pagefind --site dist`.

```console
$ pnpm register

> pagefind-trial@0.1.0 register ~/pagefind-trial
> pagefind --site dist


Running Pagefind v1.0.4 (Extended)
Running from: "~/pagefind-trial"
Source:       "dist"
Output:       "dist/pagefind"

[Walking source directory]
Found 2 files matching **/*.{html}

[Parsing files]
Did not find a data-pagefind-body element on the site.
↳ Indexing all <body> elements on the site.

[Reading languages]
Discovered 1 language: ja

[Building search indexes]
Total:
  Indexed 1 language
  Indexed 2 pages
  Indexed 11 words
  Indexed 0 filters
  Indexed 0 sorts

Finished in 0.010 seconds
```
