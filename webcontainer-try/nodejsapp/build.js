import esbuild from 'esbuild'

await esbuild.build({
  bundle: true,
  entryPoints: [
    './index.js',
  ],
  outdir: './dist',
  platform: 'node',
  format: 'esm',
})
