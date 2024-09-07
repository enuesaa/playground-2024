import esbuild from 'esbuild'
import UnpluginTypia from '@ryoppippi/unplugin-typia/esbuild'

await esbuild.build({
  plugins: [
    UnpluginTypia(),
  ],
  bundle: true,
  entryPoints: [
    './src/index.ts',
  ],
  outdir: './dist',
  platform: 'node',
  format: 'esm',
})
