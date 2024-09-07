# typia

- Type を使ってバリデーションできる
- ビルド時に Type を静的解析して、バリデーションコード (type guard 的な) を生成する
  - mode がいくつかあり必ずしもビルド時に生成するわけではないらしい
- Go の struct validator (go-playground/validator) な風味を感じた
- Type に関する情報が一箇所にまとまるので、個人的には革命的に思える
- ビルドというステップが必要なため、vite や esbuild にプラグインを入れる必要がある
  - 導入ハードルが高い印象
- 下手にバリデーションするより抜け穴が少なくなりそうな印象
  - 例えば js の typeof null === 'object' という罠にハマらなさそうな。
  
## Links
- https://github.com/samchon/typia
- https://zenn.dev/ryoppippi/articles/c4775a3a5f3c11
