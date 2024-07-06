# nextjs-biome

## Migration from prettier
biome migrate すると prettier の設定が biome に移植される  
prettier の設定は残ったままなので最後に消す

```bash
pnpm add --save-dev --save-exact @biomejs/biome
pnpm biome init
pnpm biome migrate prettier --write
```

## {ts,tsx} の記述はできない
formatter.includes で拡張子の or は書けないみたい。一個ずつ書くしかない