<script lang="ts">
	import type { Registry } from '$lib/registry'
	import { convertToSvg } from '$lib/convert'

	export let registry: Registry

	async function handleClick() {
		const data = convertToSvg(registry)
		const blob = new Blob([data], { type: 'image/svg+xml' })
		const url = URL.createObjectURL(blob)
		const a = document.createElement('a')
		a.href = url
		a.download = 'drawing.svg'
		a.click()
		URL.revokeObjectURL(url)
	}
</script>

<button on:click|preventDefault={handleClick}>Download</button>

<style lang="postcss">
	button {
		@apply text-sm border-solid border-[1px] rounded mb-1 p-1;
	}
</style>
