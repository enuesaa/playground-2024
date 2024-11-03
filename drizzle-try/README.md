# drizzle

- よく聞くが触ったことがなかったのでトライ
- sql をそのまま ORM にした感じであり、細かい調整はしやすそうに感じた
- go の ent に近い印象
- drizzle に限らず Node.js アプリ全体に言えることだが、アプリケーション設計をどうするかちょっと悩む
- studio は便利だなあ

## Commands

```bash
# this generate migration file.
# see ./drizzle
npx drizzle-kit generate

# run migration
npx drizzle-kit migrate

# studio
npx drizzle-kit studio
```
