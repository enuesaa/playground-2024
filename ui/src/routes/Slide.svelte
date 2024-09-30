<script lang="ts">
	import SlideMd from './SlideMd.svelte'
	import DrawerSvg from '$lib/components/drawer/DrawerSvg.svelte'
	import type { Registry } from '$lib/registry'

	export let content: string
	export let registry: Registry
</script>

<section class="w-full flex h-full">
	<div class="w-[1000px] h-[800px] px-5 py-8 border-2 relative">
		<SlideMd {content} />
		<div class="absolute top-0 left-0">
			<DrawerSvg bind:registry />
		</div>
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
