import { defineConfig } from '@farmfe/core'
import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin'

export default defineConfig({
  plugins: ['@farmfe/plugin-react'],

  // see https://github.com/farm-fe/farm/blob/a5554b92bbcba77e8755390a4a73611ab58bdf14/examples/vanilla-extract/farm.config.ts#L4
  vitePlugins: [vanillaExtractPlugin()],

  // TODO: これはなんだ
  compilation: {
    presetEnv: false
  },
})
