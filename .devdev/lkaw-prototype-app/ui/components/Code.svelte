<script lang="ts">
	import { HighlightAuto } from 'svelte-highlight'
    import ashes from "svelte-highlight/styles/ashes"
	import { filesStore } from '../lib/files'

	let code = ''
	filesStore.subscribe(val => {
		if (val.open === undefined) {
			code = ''
			return
		}
		const file = val.files.find(f => f.filename === val.open)
		if (file === undefined) {
			code = ''
			return
		}
		code = file.content
	})
</script>

<svelte:head>
	{@html ashes}
</svelte:head>

<div>
	<HighlightAuto {code} />
</div>

<style lang="postcss">
	div {
		@apply px-5 py-3 bg-blackgray text-graywhite rounded-xl;
		@apply w-full h-full;
		font-size: 17px;
		line-height: 1.6;
	}
</style>
