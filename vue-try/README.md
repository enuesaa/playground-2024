# vue-try

- 改めて。
- https://vuejs.org/guide/quick-start

## Commands
言わずもがなだが vite を使う

```bash
pnpm create vue@latest
```

## npm-run-all2
run-p というコマンドが生える。
これはコマンドライン引数に npm scripts を渡して、それらを並列実行できるっぽい。

## devtool
- pnpm dev で devtool も立ち上がる
- cakephp の debugkit っぽさがする
- component の依存関係とか確認できる
- 熟練の開発者には不要そうだが、オンボーディングの時とか便利かも。

## tailwind
- 普通に react と変わらない
- svelte みたいに postcss の style block (scoped) を置けるのは綺麗だなあ。
- ただし scoped をつけないとグローバルになるっぽいので注意
