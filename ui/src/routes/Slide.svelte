<script lang="ts">
	import SlideMd from './SlideMd.svelte'
	import Drawer from '$lib/components/drawer/Drawer.svelte'
	import { createRegistry } from '$lib/drawing'

	let registry = createRegistry()

	export let content: string
</script>

<section class="w-full flex h-full">
	<div class="w-[1000px] h-[800px] px-5 py-8 border-2">
		<SlideMd {content} />
		<Drawer {registry} />
	</div>
	<div class="flex-auto">
		<textarea bind:value={content} />
	</div>
</section>

<style lang="postcss">
	textarea {
		@apply block h-full bg-[rgba(0,0,0,0)];
		@apply font-normal w-full p-3 text-black text-lg outline-none;
	}
</style>
