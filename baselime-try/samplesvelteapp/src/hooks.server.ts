import { withOpenTelemetry } from '@baselime/sveltekit-opentelemetry-middleware'
import { BaselimeSDK, BetterHttpInstrumentation } from '@baselime/node-opentelemetry'
import type { Handle } from '@sveltejs/kit'
import { BASELIME_KEY } from '$env/static/private'

const sdk = new BaselimeSDK({
  baselimeKey: BASELIME_KEY,
  instrumentations: [
    new BetterHttpInstrumentation({
      plugins: [],
    }),
  ],
})

const tp = sdk.start()

export const handle: Handle = withOpenTelemetry(async ({ event, resolve }) => {
  return resolve(event)
}, {
  captureRequestBody: true,
})
