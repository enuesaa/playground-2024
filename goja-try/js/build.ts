import esbuild from 'esbuild'

await esbuild.build({
  bundle: true,
  target: 'es2017',
  entryPoints: [
    './index.ts',
  ],
  outdir: './dist',
  platform: 'node',
})
